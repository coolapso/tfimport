package providers

import (
	"fmt"
)

// resolveCustomExtractKubernetesImportID provides custom logic for extracting import IDs
// for specific kubernetes resources that don't follow the standard pattern.
func resolveCustomextractKubernetesImportID(ctx *ProviderContext, resourceType string, config map[string]any) string {
	if resourceType == "kubernetes_manifest" {
		manifestBlock, ok := config["manifest"].(map[string]any)
		if !ok {
			// Terraform also supports "object" block for some versions or states
			manifestBlock, ok = config["object"].(map[string]any)
		}

		if ok {
			apiVersion, _ := manifestBlock["apiVersion"].(string)
			kind, _ := manifestBlock["kind"].(string)

			if metadata, ok := manifestBlock["metadata"].(map[string]any); ok {
				name, hasName := metadata["name"].(string)
				namespace, hasNamespace := metadata["namespace"].(string)

				if apiVersion != "" && kind != "" && hasName && name != "" {
					if hasNamespace && namespace != "" {
						return fmt.Sprintf("apiVersion=%s,kind=%s,namespace=%s,name=%s", apiVersion, kind, namespace, name)
					}
					return fmt.Sprintf("apiVersion=%s,kind=%s,name=%s", apiVersion, kind, name)
				}
			}
		}
	}

	// Return "" to fall back to the default extraction logic.
	return ""
}
