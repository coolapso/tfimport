package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func getScalewayStrategy() ProviderStrategy {
	return ProviderStrategy{
		MatchRegex: regexp.MustCompile(`(?i)imported using the ([^,\n]+?)(?:, e\.g\.| argument| as shown below)`),
		ExtractFunc: func(match []string) ([]string, bool, string) {
			formatStr := strings.TrimSpace(match[1])
			// remove backticks
			formatStr = strings.Trim(formatStr, "`")
			// remove surrounding quotes if any
			formatStr = strings.Trim(formatStr, "\"")

			// Some are just "id", others are "{zone}/{id}" or "{region}/{deployment_id}/{name}"
			if formatStr == "id" {
				formatStr = "{id}"
			}

			// check if it resembles {something}
			if strings.Contains(formatStr, "{") && strings.Contains(formatStr, "}") {
				return []string{formatStr}, false, ""
			}

			return nil, true, fmt.Sprintf("Unknown format: %s", formatStr)
		},
		GenerateFunc: func(mappings []mapping, funcName, providerName string) string {
			var buf bytes.Buffer
			fmt.Fprintf(&buf, `package providers

import (
	"fmt"
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
					formatStr := m.Keys[0] // e.g. "{zone}/{id}"

					varRegex := regexp.MustCompile(`\{([^}]+)\}`)
					vars := varRegex.FindAllStringSubmatch(formatStr, -1)

					if len(vars) == 0 {
						fmt.Fprintf(&buf, "\t\treturn \"%s\"\n", formatStr)
						continue
					}

					templateStr := varRegex.ReplaceAllString(formatStr, "%s")

					fmt.Fprintf(&buf, "\t\t// Format: %s\n", formatStr)
					fmt.Fprintf(&buf, "\t\t{\n")

					var checks []string
					var varNames []string

					for i, v := range vars {
						varName := v[1]
						varSafe := fmt.Sprintf("v%d", i)
						fmt.Fprintf(&buf, "\t\t\t%s, ok%d := config[\"%s\"].(string)\n", varSafe, i, varName)
						checks = append(checks, fmt.Sprintf("ok%d && %s != \"\"", i, varSafe))
						varNames = append(varNames, varSafe)
					}

					fmt.Fprintf(&buf, "\t\t\tif %s {\n", strings.Join(checks, " && "))

					if templateStr == "%s" {
						fmt.Fprintf(&buf, "\t\t\t\treturn %s\n", varNames[0])
					} else {
						escapedTemplate := strings.ReplaceAll(templateStr, "\"", "\\\"")
						fmt.Fprintf(&buf, "\t\t\t\treturn fmt.Sprintf(\"%s\", %s)\n", escapedTemplate, strings.Join(varNames, ", "))
					}

					fmt.Fprintf(&buf, "\t\t\t}\n")
					fmt.Fprintf(&buf, "\t\t}\n")
					fmt.Fprintf(&buf, "\t\treturn \"\"\n")
				}
			}

			fmt.Fprintf(&buf, "\t}\n\treturn \"\"\n}\n")
			return buf.String()
		},
	}
}
