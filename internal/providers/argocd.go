package providers

import (
	"fmt"
)

// extractArgocdImportID returns the necessary import ID for an argocd resource
// based on its configuration extracted from the terraform plan.
func extractArgocdImportID(ctx *ProviderContext, resourceType string, config map[string]any) string {
	// First, check if there's a custom resolver for this resource
	if id := resolveCustomextractArgocdImportID(ctx, resourceType, config); id != "" {
		return id
	}

	// For ArgoCD, some attributes are inside metadata
	var metadata map[string]any
	if mdBlock, ok := config["metadata"].([]any); ok && len(mdBlock) > 0 {
		metadata, _ = mdBlock[0].(map[string]any)
	}

	switch resourceType {
	case "argocd_account_token":
		// No standard import format found in documentation
		return ""
	case "argocd_application":
		var parts []string
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s:%s", parts[0], parts[1])
	case "argocd_application_set":
		// No standard import format found in documentation
		return ""
	case "argocd_cluster":
		if val, ok := config["server"].(string); ok && val != "" {
			return val
		}
		return ""
	case "argocd_gpg_key":
		// Computed ID format: the key ID
		return ""
	case "argocd_project":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "argocd_project_token":
		// No standard import format found in documentation
		return ""
	case "argocd_repository":
		if val, ok := config["repo"].(string); ok && val != "" {
			return val
		}
		return ""
	case "argocd_repository_certificate":
		// No standard import format found in documentation
		return ""
	case "argocd_repository_credentials":
		if val, ok := config["url"].(string); ok && val != "" {
			return val
		}
		return ""
	}
	return ""
}
