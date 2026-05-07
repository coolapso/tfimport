package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func getKubernetesStrategy() ProviderStrategy {
	return ProviderStrategy{
		MatchRegex: regexp.MustCompile(`(?i)can be imported using (.*?)(?:,|:)`),
		ExtractFunc: func(match []string) ([]string, bool, string) {
			usingStr := strings.TrimSpace(match[1])
			lower := strings.ToLower(usingStr)

			if strings.Contains(lower, "namespace and name") {
				return []string{"namespace", "name"}, false, ""
			} else if strings.Contains(lower, "name") {
				return []string{"name"}, false, ""
			} else if strings.Contains(lower, "id") || strings.Contains(lower, "arn") {
				return nil, true, fmt.Sprintf("Computed ID format: %s", usingStr)
			}
			return nil, true, fmt.Sprintf("Unknown ID format: %s", usingStr)
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

	// For kubernetes, most attributes are inside metadata
	var metadata map[string]any
	if mdBlock, ok := config["metadata"].([]any); ok && len(mdBlock) > 0 {
		metadata, _ = mdBlock[0].(map[string]any)
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
						fmt.Fprintf(&buf, "\t\tif val, ok := metadata[\"%s\"].(string); ok && val != \"\" {\n", key)
						fmt.Fprintf(&buf, "\t\t\treturn val\n")
						fmt.Fprintf(&buf, "\t\t}\n")
						fmt.Fprintf(&buf, "\t\treturn \"\"\n")
					} else {
						fmt.Fprintf(&buf, "\t\tvar parts []string\n")
						for _, key := range m.Keys {
							fmt.Fprintf(&buf, "\t\tif val, ok := metadata[\"%s\"].(string); ok && val != \"\" {\n", key)
							fmt.Fprintf(&buf, "\t\t\tparts = append(parts, val)\n")
							fmt.Fprintf(&buf, "\t\t} else {\n")
							fmt.Fprintf(&buf, "\t\t\treturn \"\"\n")
							fmt.Fprintf(&buf, "\t\t}\n")
						}
						fmt.Fprintf(&buf, "\t\treturn fmt.Sprintf(\"%%s/%%s\", parts[0], parts[1])\n")
					}
				}
			}

			fmt.Fprintf(&buf, "\t}\n\treturn \"\"\n}\n")
			return buf.String()
		},
	}
}
