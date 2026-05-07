package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func getAzurermStrategy() ProviderStrategy {
	return ProviderStrategy{
		MatchRegex: regexp.MustCompile(`terraform import azurerm_[a-zA-Z0-9_]+(?:\.[a-zA-Z0-9_]+)?\s+(.*)`),
		ExtractFunc: func(match []string) ([]string, bool, string) {
			idStr := match[1]
			return nil, true, fmt.Sprintf("Azure Resource ID: %s", idStr)
		},
		GenerateFunc: func(mappings []mapping, funcName, providerName string) string {
			var buf bytes.Buffer
			fmt.Fprintf(&buf, `package providers

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
					// Fallback for non-computed, though Azure typically won't hit this.
					fmt.Fprintf(&buf, "\t\treturn \"\"\n")
				}
			}

			fmt.Fprintf(&buf, "\t}\n\treturn \"\"\n}\n")
			return buf.String()
		},
	}
}
