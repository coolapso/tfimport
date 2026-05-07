package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func getGoogleStrategy() ProviderStrategy {
	return ProviderStrategy{
		MatchRegex: regexp.MustCompile(`(?s)## Import.*?\n((?:\*\s+` + "`" + `[^` + "`" + `]+` + "`" + `\n)+)`),
		ExtractFunc: func(match []string) ([]string, bool, string) {
			lines := strings.Split(strings.TrimSpace(match[1]), "\n")
			var formats []string
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.HasPrefix(line, "* ") {
					line = strings.TrimPrefix(line, "* ")
					line = strings.Trim(line, "`")
					formats = append(formats, line)
				}
			}
			if len(formats) > 0 {
				// We return formats in keys
				return formats, false, ""
			}
			return nil, true, "No valid import formats found"
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
					for _, formatStr := range m.Keys {
						// formatStr might be "projects/{{project}}/instances/{{name}}"
						// Extract variables
						varRegex := regexp.MustCompile(`\{{1,2}([^}]+)\}{1,2}`)
						vars := varRegex.FindAllStringSubmatch(formatStr, -1)

						if len(vars) == 0 {
							// Literal format? Just return it.
							escapedLiteral := strings.ReplaceAll(formatStr, "\"", "\\\"")
							fmt.Fprintf(&buf, "\t\treturn \"%s\"\n", escapedLiteral)
							continue
						}

						// Build the fmt.Sprintf template
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
					}
					fmt.Fprintf(&buf, "\t\treturn \"\"\n")
				}
			}

			fmt.Fprintf(&buf, "\t}\n\treturn \"\"\n}\n")
			return buf.String()
		},
	}
}
