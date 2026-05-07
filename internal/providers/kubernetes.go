package providers

import (
	"fmt"
)

// extractKubernetesImportID returns the necessary import ID for a kubernetes resource
// based on its configuration extracted from the terraform plan.
func extractKubernetesImportID(ctx *ProviderContext, resourceType string, config map[string]any) string {
	// First, check if there's a custom resolver for this resource
	if id := resolveCustomextractKubernetesImportID(ctx, resourceType, config); id != "" {
		return id
	}

	// For kubernetes, most attributes are inside metadata
	var metadata map[string]any
	if mdBlock, ok := config["metadata"].([]any); ok && len(mdBlock) > 0 {
		metadata, _ = mdBlock[0].(map[string]any)
	}

	switch resourceType {
	case "kubernetes_annotations":
		// No standard import format found in documentation
		return ""
	case "kubernetes_api_service":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_api_service_v1":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_certificate_signing_request":
		// No standard import format found in documentation
		return ""
	case "kubernetes_certificate_signing_request_v1":
		// No standard import format found in documentation
		return ""
	case "kubernetes_cluster_role":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_cluster_role_binding":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_cluster_role_binding_v1":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_cluster_role_v1":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_config_map":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_config_map_v1":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_config_map_v1_data":
		// No standard import format found in documentation
		return ""
	case "kubernetes_cron_job":
		// No standard import format found in documentation
		return ""
	case "kubernetes_cron_job_v1":
		// No standard import format found in documentation
		return ""
	case "kubernetes_csi_driver":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_csi_driver_v1":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_daemon_set_v1":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_daemonset":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_default_service_account":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_default_service_account_v1":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_deployment":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_deployment_v1":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_endpoint_slice_v1":
		// No standard import format found in documentation
		return ""
	case "kubernetes_endpoints":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_endpoints_v1":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_env":
		// No standard import format found in documentation
		return ""
	case "kubernetes_horizontal_pod_autoscaler":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_horizontal_pod_autoscaler_v1":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_horizontal_pod_autoscaler_v2":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_horizontal_pod_autoscaler_v2beta2":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_ingress":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_ingress_class":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_ingress_class_v1":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_ingress_v1":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_job":
		// No standard import format found in documentation
		return ""
	case "kubernetes_job_v1":
		// No standard import format found in documentation
		return ""
	case "kubernetes_labels":
		// No standard import format found in documentation
		return ""
	case "kubernetes_limit_range":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_limit_range_v1":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_manifest":
		// No standard import format found in documentation
		return ""
	case "kubernetes_mutating_webhook_configuration":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_mutating_webhook_configuration_v1":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_namespace":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_namespace_v1":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_network_policy":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_network_policy_v1":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_node_taint":
		// No standard import format found in documentation
		return ""
	case "kubernetes_persistent_volume":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_persistent_volume_claim":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_persistent_volume_claim_v1":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_persistent_volume_v1":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_pod":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_pod_disruption_budget":
		// No standard import format found in documentation
		return ""
	case "kubernetes_pod_disruption_budget_v1":
		// No standard import format found in documentation
		return ""
	case "kubernetes_pod_security_policy":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_pod_security_policy_v1beta1":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_pod_v1":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_priority_class":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_priority_class_v1":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_replication_controller":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_replication_controller_v1":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_resource_quota":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_resource_quota_v1":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_role":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_role_binding":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_role_binding_v1":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_role_v1":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_runtime_class_v1":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_secret":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_secret_v1":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_secret_v1_data":
		// No standard import format found in documentation
		return ""
	case "kubernetes_service":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_service_account":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_service_account_v1":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_service_v1":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_stateful_set":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_stateful_set_v1":
		var parts []string
		if val, ok := metadata["namespace"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := metadata["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		return fmt.Sprintf("%s/%s", parts[0], parts[1])
	case "kubernetes_storage_class":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_storage_class_v1":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_token_request_v1":
		// No standard import format found in documentation
		return ""
	case "kubernetes_validating_admission_policy":
		// No standard import format found in documentation
		return ""
	case "kubernetes_validating_webhook_configuration":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "kubernetes_validating_webhook_configuration_v1":
		if val, ok := metadata["name"].(string); ok && val != "" {
			return val
		}
		return ""
	}
	return ""
}
