package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func getArgocdStrategy() ProviderStrategy {
	return ProviderStrategy{
		MatchRegex: regexp.MustCompile(`(?i)imported using (.*?)\.`),
		ExtractFunc: func(match []string) ([]string, bool, string) {
			usingStr := strings.TrimSpace(match[1])
			lower := strings.ToLower(usingStr)

			if strings.Contains(lower, "{name}:{namespace}") {
				return []string{"name", "namespace"}, false, ""
			} else if strings.Contains(lower, "server url") {
				return []string{"server"}, false, ""
			} else if strings.Contains(lower, "project name") {
				return []string{"name"}, false, ""
			} else if strings.Contains(lower, "repository url") {
				// Depending on resource, it could be `repo` (for repository) or `url` (for repository_credentials)
				return []string{"URL_OR_REPO"}, false, ""
			} else if strings.Contains(lower, "key id") {
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

// %s returns the necessary import ID for an %s resource
// based on its configuration extracted from the terraform plan.
func %s(ctx *ProviderContext, resourceType string, config map[string]any) string {
	// First, check if there's a custom resolver for this resource
	if id := resolveCustom%s(ctx, resourceType, config); id != "" {
		return id
	}

	// For ArgoCD, some attributes are inside metadata
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
						if key == "URL_OR_REPO" {
							// Special logic for repository url which can be `repo` or `url` depending on resource
							switch m.Resource {
							case "argocd_repository":
								fmt.Fprintf(&buf, "\t\tif val, ok := config[\"repo\"].(string); ok && val != \"\" {\n")
								fmt.Fprintf(&buf, "\t\t\treturn val\n")
								fmt.Fprintf(&buf, "\t\t}\n")
							case "argocd_repository_credentials":
								fmt.Fprintf(&buf, "\t\tif val, ok := config[\"url\"].(string); ok && val != \"\" {\n")
								fmt.Fprintf(&buf, "\t\t\treturn val\n")
								fmt.Fprintf(&buf, "\t\t}\n")
							}
							fmt.Fprintf(&buf, "\t\treturn \"\"\n")
						} else if key == "name" && (m.Resource == "argocd_project" || m.Resource == "argocd_application") {
							fmt.Fprintf(&buf, "\t\tif val, ok := metadata[\"%s\"].(string); ok && val != \"\" {\n", key)
							fmt.Fprintf(&buf, "\t\t\treturn val\n")
							fmt.Fprintf(&buf, "\t\t}\n")
							fmt.Fprintf(&buf, "\t\treturn \"\"\n")
						} else {
							fmt.Fprintf(&buf, "\t\tif val, ok := config[\"%s\"].(string); ok && val != \"\" {\n", key)
							fmt.Fprintf(&buf, "\t\t\treturn val\n")
							fmt.Fprintf(&buf, "\t\t}\n")
							fmt.Fprintf(&buf, "\t\treturn \"\"\n")
						}
					} else {
						// This is for application: name and namespace
						fmt.Fprintf(&buf, "\t\tvar parts []string\n")
						for _, key := range m.Keys {
							fmt.Fprintf(&buf, "\t\tif val, ok := metadata[\"%s\"].(string); ok && val != \"\" {\n", key)
							fmt.Fprintf(&buf, "\t\t\tparts = append(parts, val)\n")
							fmt.Fprintf(&buf, "\t\t} else {\n")
							fmt.Fprintf(&buf, "\t\t\treturn \"\"\n")
							fmt.Fprintf(&buf, "\t\t}\n")
						}
						fmt.Fprintf(&buf, "\t\treturn fmt.Sprintf(\"%%s:%%s\", parts[0], parts[1])\n")
					}
				}
			}

			fmt.Fprintf(&buf, "\t}\n\treturn \"\"\n}\n")
			return buf.String()
		},
	}
}
