package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func getAWSStrategy() ProviderStrategy {
	return ProviderStrategy{
		MatchRegex: regexp.MustCompile("(?i)Using `terraform import`, import .*? using (.*)"),
		ExtractFunc: func(match []string) ([]string, bool, string) {
			usingStr := match[1]
			usingStr = strings.TrimSuffix(usingStr, " For example:")
			usingStr = strings.TrimSuffix(usingStr, ".")

			backtickRegex := regexp.MustCompile("`([^`]+)`")
			btMatches := backtickRegex.FindAllStringSubmatch(usingStr, -1)

			if len(btMatches) > 0 {
				inner := btMatches[0][1]
				inner = strings.ReplaceAll(inner, "/", ",")
				parts := strings.Split(inner, ",")

				validSnakeCase := regexp.MustCompile(`^[a-z_][a-z0-9_]*$`)
				validKeys := true
				var keys []string

				for _, p := range parts {
					p = strings.TrimSpace(p)
					if p == "" || p == "id" || p == "arn" || !validSnakeCase.MatchString(p) {
						validKeys = false
						break
					}
					keys = append(keys, p)
				}

				if validKeys && len(keys) > 0 {
					return keys, false, ""
				}
				return nil, true, fmt.Sprintf("Computed or complex ID format: %s", usingStr)
			}

			lower := strings.ToLower(usingStr)
			if strings.Contains(lower, "id") || strings.Contains(lower, "arn") {
				return nil, true, fmt.Sprintf("Computed ID format: %s", usingStr)
			} else if strings.Contains(lower, "name") {
				return []string{"name"}, false, ""
			}
			return nil, true, fmt.Sprintf("Unknown ID format: %s", usingStr)
		},
		GenerateFunc: func(mappings []mapping, funcName, providerName string) string {
			var buf bytes.Buffer
			fmt.Fprintf(&buf, `package providers

import (
	"strings"
)

// %s returns the necessary import ID for a %s resource
// based on its configuration extracted from the terraform plan.
func %s(ctx *ProviderContext, resourceType string, config map[string]any) string {
	// First, check if there's a custom resolver for this resource
	if id := resolveCustom%s(ctx, resourceType, config); id != "" {
		return id
	}

	switch resourceType {
`, funcName, providerName, funcName, funcName)

			for _, m := range mappings {
				fmt.Fprintf(&buf, "\tcase \"%s\":\n", m.Resource)
				if m.IsComputed {
					fmt.Fprintf(&buf, "\t\t// %s\n", m.Comment)
					fmt.Fprintf(&buf, "\t\treturn \"\"\n")
				} else {
					if len(m.Keys) == 1 {
						key := m.Keys[0]
						fmt.Fprintf(&buf, "\t\tif val, ok := config[\"%s\"].(string); ok && val != \"\" {\n", key)
						fmt.Fprintf(&buf, "\t\t\treturn val\n")
						fmt.Fprintf(&buf, "\t\t}\n")
						fmt.Fprintf(&buf, "\t\treturn \"\"\n")
					} else {
						fmt.Fprintf(&buf, "\t\tvar parts []string\n")
						for _, key := range m.Keys {
							fmt.Fprintf(&buf, "\t\tif val, ok := config[\"%s\"].(string); ok && val != \"\" {\n", key)
							fmt.Fprintf(&buf, "\t\t\tparts = append(parts, val)\n")
							fmt.Fprintf(&buf, "\t\t} else {\n")
							fmt.Fprintf(&buf, "\t\t\treturn \"\"\n")
							fmt.Fprintf(&buf, "\t\t}\n")
						}
						fmt.Fprintf(&buf, "\t\t// Note: Composite separator might need adjustment based on exact docs\n")
						fmt.Fprintf(&buf, "\t\treturn strings.Join(parts, \",\")\n")
					}
				}
			}

			fmt.Fprintf(&buf, "\t}\n\treturn \"\"\n}\n")
			return buf.String()
		},
	}
}
