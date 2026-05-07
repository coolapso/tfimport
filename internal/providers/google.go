package providers

import (
	"fmt"
)

// extractGoogleImportID returns the necessary import ID for a google resource
// based on its configuration extracted from the terraform plan.
func extractGoogleImportID(ctx *ProviderContext, resourceType string, config map[string]any) string {
	if id := resolveCustomextractGoogleImportID(ctx, resourceType, config); id != "" {
		return id
	}

	switch resourceType {
	case "google_access_context_manager_access_level":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_access_context_manager_access_level_condition":
		// No standard import format found in documentation
		return ""
	case "google_access_context_manager_access_levels":
		// Format: {{parent}}/accessLevels
		{
			v0, ok0 := config["parent"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("%s/accessLevels", v0)
			}
		}
		// Format: {{parent}}
		{
			v0, ok0 := config["parent"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_access_context_manager_access_policy":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_access_context_manager_access_policy_iam":
		// No standard import format found in documentation
		return ""
	case "google_access_context_manager_authorized_orgs_desc":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_access_context_manager_egress_policy":
		// Format: {{egress_policy_name}}/{{resource}}
		{
			v0, ok0 := config["egress_policy_name"].(string)
			v1, ok1 := config["resource"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_access_context_manager_gcp_user_access_binding":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_access_context_manager_ingress_policy":
		// Format: {{ingress_policy_name}}/{{resource}}
		{
			v0, ok0 := config["ingress_policy_name"].(string)
			v1, ok1 := config["resource"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_access_context_manager_service_perimeter":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_access_context_manager_service_perimeter_dry_run_egress_policy":
		// No standard import format found in documentation
		return ""
	case "google_access_context_manager_service_perimeter_dry_run_ingress_policy":
		// No standard import format found in documentation
		return ""
	case "google_access_context_manager_service_perimeter_dry_run_resource":
		// Format: {{perimeter_name}}/{{resource}}
		{
			v0, ok0 := config["perimeter_name"].(string)
			v1, ok1 := config["resource"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_access_context_manager_service_perimeter_egress_policy":
		// No standard import format found in documentation
		return ""
	case "google_access_context_manager_service_perimeter_ingress_policy":
		// No standard import format found in documentation
		return ""
	case "google_access_context_manager_service_perimeter_resource":
		// Format: {{perimeter_name}}/{{resource}}
		{
			v0, ok0 := config["perimeter_name"].(string)
			v1, ok1 := config["resource"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_access_context_manager_service_perimeters":
		// Format: {{parent}}/servicePerimeters
		{
			v0, ok0 := config["parent"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("%s/servicePerimeters", v0)
			}
		}
		// Format: {{parent}}
		{
			v0, ok0 := config["parent"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_active_directory_domain":
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{project}} {{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s %s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_active_directory_domain_trust":
		// Format: projects/{{project}}/locations/global/domains/{{domain}}/{{target_domain_name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["domain"].(string)
			v2, ok2 := config["target_domain_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/global/domains/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{domain}}/{{target_domain_name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["domain"].(string)
			v2, ok2 := config["target_domain_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{domain}}/{{target_domain_name}}
		{
			v0, ok0 := config["domain"].(string)
			v1, ok1 := config["target_domain_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_active_directory_peering":
		// No standard import format found in documentation
		return ""
	case "google_alloydb_backup":
		// Format: projects/{{project}}/locations/{{location}}/backups/{{backup_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["backup_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/backups/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{backup_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["backup_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{backup_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["backup_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_alloydb_cluster":
		// Format: projects/{{project}}/locations/{{location}}/clusters/{{cluster_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{cluster_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{cluster_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{cluster_id}}
		{
			v0, ok0 := config["cluster_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_alloydb_instance":
		// Format: projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/instances/{{instance_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster"].(string)
			v3, ok3 := config["instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/clusters/%s/instances/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{cluster}}/{{instance_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster"].(string)
			v3, ok3 := config["instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{cluster}}/{{instance_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["cluster"].(string)
			v2, ok2 := config["instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_alloydb_user":
		// Format: projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/users/{{user_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster"].(string)
			v3, ok3 := config["user_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/clusters/%s/users/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{cluster}}/{{user_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster"].(string)
			v3, ok3 := config["user_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{cluster}}/{{user_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["cluster"].(string)
			v2, ok2 := config["user_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_api_gateway_api":
		// Format: projects/{{project}}/locations/global/apis/{{api_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["api_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/apis/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{api_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["api_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{api_id}}
		{
			v0, ok0 := config["api_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_api_gateway_api_config":
		// Format: projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["api"].(string)
			v2, ok2 := config["api_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/global/apis/%s/configs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{api}}/{{api_config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["api"].(string)
			v2, ok2 := config["api_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{api}}/{{api_config_id}}
		{
			v0, ok0 := config["api"].(string)
			v1, ok1 := config["api_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_api_gateway_api_config_iam":
		// No standard import format found in documentation
		return ""
	case "google_api_gateway_api_iam":
		// No standard import format found in documentation
		return ""
	case "google_api_gateway_gateway":
		// Format: projects/{{project}}/locations/{{region}}/gateways/{{gateway_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["gateway_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/gateways/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{gateway_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["gateway_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{gateway_id}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["gateway_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{gateway_id}}
		{
			v0, ok0 := config["gateway_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_api_gateway_gateway_iam":
		// No standard import format found in documentation
		return ""
	case "google_apigee_addons_config":
		// Format: organizations/{{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("organizations/%s", v0)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_apigee_api":
		// Format: {{org_id}}/apis/{{name}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/apis/%s", v0, v1)
			}
		}
		// Format: {{org_id}}/{{name}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_api_deployment":
		// Format: organizations/{{org_id}}/environments/{{environment}}/apis/{{proxy_id}}/revisions/{{revision}}/deployments
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["environment"].(string)
			v2, ok2 := config["proxy_id"].(string)
			v3, ok3 := config["revision"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("organizations/%s/environments/%s/apis/%s/revisions/%s/deployments", v0, v1, v2, v3)
			}
		}
		// Format: organizations/{{org_id}}/environments/{{environment}}/apis/{{proxy_id}}/revisions/{{revision}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["environment"].(string)
			v2, ok2 := config["proxy_id"].(string)
			v3, ok3 := config["revision"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("organizations/%s/environments/%s/apis/%s/revisions/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{org_id}}/{{environment}}/{{proxy_id}}/{{revision}}/deployments
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["environment"].(string)
			v2, ok2 := config["proxy_id"].(string)
			v3, ok3 := config["revision"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s/deployments", v0, v1, v2, v3)
			}
		}
		// Format: {{org_id}}/{{environment}}/{{proxy_id}}/{{revision}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["environment"].(string)
			v2, ok2 := config["proxy_id"].(string)
			v3, ok3 := config["revision"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_apigee_api_product":
		// Format: {{org_id}}/apiproducts/{{name}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/apiproducts/%s", v0, v1)
			}
		}
		// Format: {{org_id}}/{{name}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_app_group":
		// Format: {{org_id}}/appgroups/{{name}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/appgroups/%s", v0, v1)
			}
		}
		// Format: {{org_id}}/{{name}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_control_plane_access":
		// Format: organizations/{{name}}/controlPlaneAccess
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("organizations/%s/controlPlaneAccess", v0)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_apigee_developer":
		// Format: {{org_id}}/developers/{{email}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["email"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/developers/%s", v0, v1)
			}
		}
		// Format: {{org_id}}/{{email}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["email"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_developer_app":
		// Format: {{org_id}}/developers/{{developer_email}}/apps/{{name}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["developer_email"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/developers/%s/apps/%s", v0, v1, v2)
			}
		}
		// Format: {{org_id}}/{{developer_email}}/{{name}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["developer_email"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_apigee_dns_zone":
		// Format: {{org_id}}/dnsZones/{{dns_zone_id}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["dns_zone_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/dnsZones/%s", v0, v1)
			}
		}
		// Format: {{org_id}}/{{dns_zone_id}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["dns_zone_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_endpoint_attachment":
		// Format: {{org_id}}/endpointAttachments/{{endpoint_attachment_id}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["endpoint_attachment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/endpointAttachments/%s", v0, v1)
			}
		}
		// Format: {{org_id}}/{{endpoint_attachment_id}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["endpoint_attachment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_env_keystore":
		// Format: {{env_id}}/keystores/{{name}}
		{
			v0, ok0 := config["env_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/keystores/%s", v0, v1)
			}
		}
		// Format: {{env_id}}/{{name}}
		{
			v0, ok0 := config["env_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_env_references":
		// Format: {{env_id}}/references/{{name}}
		{
			v0, ok0 := config["env_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/references/%s", v0, v1)
			}
		}
		// Format: {{env_id}}/{{name}}
		{
			v0, ok0 := config["env_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_envgroup":
		// Format: {{org_id}}/envgroups/{{name}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/envgroups/%s", v0, v1)
			}
		}
		// Format: {{org_id}}/{{name}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_envgroup_attachment":
		// Format: {{envgroup_id}}/attachments/{{name}}
		{
			v0, ok0 := config["envgroup_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/attachments/%s", v0, v1)
			}
		}
		// Format: {{envgroup_id}}/{{name}}
		{
			v0, ok0 := config["envgroup_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_environment":
		// Format: {{org_id}}/environments/{{name}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/environments/%s", v0, v1)
			}
		}
		// Format: {{org_id}}/{{name}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_environment_addons_config":
		// Format: {{env_id}}
		{
			v0, ok0 := config["env_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_apigee_environment_api_revision_deployment":
		// Format: organizations/{{org_id}}/environments/{{environment}}/apis/{{api}}/revisions/{{revision}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["environment"].(string)
			v2, ok2 := config["api"].(string)
			v3, ok3 := config["revision"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("organizations/%s/environments/%s/apis/%s/revisions/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{org_id}}/{{environment}}/{{api}}/{{revision}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["environment"].(string)
			v2, ok2 := config["api"].(string)
			v3, ok3 := config["revision"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{id}}
		{
			v0, ok0 := config["id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_apigee_environment_iam":
		// No standard import format found in documentation
		return ""
	case "google_apigee_environment_keyvaluemaps":
		// Format: {{env_id}}/keyvaluemaps/{{name}}
		{
			v0, ok0 := config["env_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/keyvaluemaps/%s", v0, v1)
			}
		}
		// Format: {{env_id}}/{{name}}
		{
			v0, ok0 := config["env_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_environment_keyvaluemaps_entries":
		// Format: {{env_keyvaluemap_id}}/entries/{{name}}
		{
			v0, ok0 := config["env_keyvaluemap_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/entries/%s", v0, v1)
			}
		}
		// Format: {{env_keyvaluemap_id}}/{{name}}
		{
			v0, ok0 := config["env_keyvaluemap_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_flowhook":
		// Format: organizations/{{org_id}}/environments/{{environment}}/flowhooks/{{flow_hook_point}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["environment"].(string)
			v2, ok2 := config["flow_hook_point"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("organizations/%s/environments/%s/flowhooks/%s", v0, v1, v2)
			}
		}
		// Format: {{org_id}}/{{environment}}/{{flow_hook_point}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["environment"].(string)
			v2, ok2 := config["flow_hook_point"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_apigee_instance":
		// Format: {{org_id}}/instances/{{name}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/instances/%s", v0, v1)
			}
		}
		// Format: {{org_id}}/{{name}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_instance_attachment":
		// Format: {{instance_id}}/attachments/{{name}}
		{
			v0, ok0 := config["instance_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/attachments/%s", v0, v1)
			}
		}
		// Format: {{instance_id}}/{{name}}
		{
			v0, ok0 := config["instance_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_keystores_aliases_key_cert_file":
		// Format: organizations/{{org_id}}/environments/{{environment}}/keystores/{{keystore}}/aliases/{{alias}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["environment"].(string)
			v2, ok2 := config["keystore"].(string)
			v3, ok3 := config["alias"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("organizations/%s/environments/%s/keystores/%s/aliases/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{org_id}}/{{environment}}/{{keystore}}/{{alias}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["environment"].(string)
			v2, ok2 := config["keystore"].(string)
			v3, ok3 := config["alias"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_apigee_keystores_aliases_pkcs12":
		// Format: organizations/{{org_id}}/environments/{{environment}}/keystores/{{keystore}}/aliases/{{alias}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["environment"].(string)
			v2, ok2 := config["keystore"].(string)
			v3, ok3 := config["alias"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("organizations/%s/environments/%s/keystores/%s/aliases/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{org_id}}/{{environment}}/{{keystore}}/{{alias}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["environment"].(string)
			v2, ok2 := config["keystore"].(string)
			v3, ok3 := config["alias"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_apigee_keystores_aliases_self_signed_cert":
		// Format: organizations/{{org_id}}/environments/{{environment}}/keystores/{{keystore}}/aliases/{{alias}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["environment"].(string)
			v2, ok2 := config["keystore"].(string)
			v3, ok3 := config["alias"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("organizations/%s/environments/%s/keystores/%s/aliases/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{org_id}}/{{environment}}/{{keystore}}/{{alias}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["environment"].(string)
			v2, ok2 := config["keystore"].(string)
			v3, ok3 := config["alias"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_apigee_nat_address":
		// Format: {{instance_id}}/natAddresses/{{name}}
		{
			v0, ok0 := config["instance_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/natAddresses/%s", v0, v1)
			}
		}
		// Format: {{instance_id}}/{{name}}
		{
			v0, ok0 := config["instance_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_organization":
		// Format: organizations/{{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("organizations/%s", v0)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_apigee_security_action":
		// Format: organizations/{{org_id}}/environments/{{env_id}}/securityActions/{{security_action_id}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["env_id"].(string)
			v2, ok2 := config["security_action_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("organizations/%s/environments/%s/securityActions/%s", v0, v1, v2)
			}
		}
		// Format: {{org_id}}/{{env_id}}/{{security_action_id}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["env_id"].(string)
			v2, ok2 := config["security_action_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_apigee_security_feedback":
		// Format: {{org_id}}/securityFeedback/{{feedback_id}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["feedback_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/securityFeedback/%s", v0, v1)
			}
		}
		// Format: {{org_id}}/{{feedback_id}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["feedback_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_security_monitoring_condition":
		// Format: {{org_id}}/securityMonitoringConditions/{{condition_id}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["condition_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/securityMonitoringConditions/%s", v0, v1)
			}
		}
		// Format: {{org_id}}/{{condition_id}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["condition_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_security_profile_v2":
		// Format: {{org_id}}/securityProfilesV2/{{profile_id}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["profile_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/securityProfilesV2/%s", v0, v1)
			}
		}
		// Format: {{org_id}}/{{profile_id}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["profile_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_sharedflow":
		// Format: {{org_id}}/sharedflows/{{name}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/sharedflows/%s", v0, v1)
			}
		}
		// Format: {{org_id}}/{{name}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_sharedflow_deployment":
		// Format: organizations/{{org_id}}/environments/{{environment}}/sharedflows/{{sharedflow_id}}/revisions/{{revision}}/deployments/{{name}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["environment"].(string)
			v2, ok2 := config["sharedflow_id"].(string)
			v3, ok3 := config["revision"].(string)
			v4, ok4 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("organizations/%s/environments/%s/sharedflows/%s/revisions/%s/deployments/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{org_id}}/{{environment}}/{{sharedflow_id}}/{{revision}}/{{name}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["environment"].(string)
			v2, ok2 := config["sharedflow_id"].(string)
			v3, ok3 := config["revision"].(string)
			v4, ok4 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("%s/%s/%s/%s/%s", v0, v1, v2, v3, v4)
			}
		}
		return ""
	case "google_apigee_space":
		// Format: {{org_id}}/spaces/{{space_id}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["space_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/spaces/%s", v0, v1)
			}
		}
		// Format: {{org_id}}/{{space_id}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["space_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apigee_sync_authorization":
		// Format: organizations/{{name}}/syncAuthorization
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("organizations/%s/syncAuthorization", v0)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_apigee_target_server":
		// Format: {{env_id}}/targetservers/{{name}}
		{
			v0, ok0 := config["env_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/targetservers/%s", v0, v1)
			}
		}
		// Format: {{env_id}}/{{name}}
		{
			v0, ok0 := config["env_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apihub_api_hub_instance":
		// Format: projects/{{project}}/locations/{{location}}/apiHubInstances/{{api_hub_instance_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["api_hub_instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/apiHubInstances/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{api_hub_instance_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["api_hub_instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{api_hub_instance_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["api_hub_instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apihub_curation":
		// Format: projects/{{project}}/locations/{{location}}/curations/{{curation_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["curation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/curations/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{curation_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["curation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{curation_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["curation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apihub_host_project_registration":
		// Format: projects/{{project}}/locations/{{location}}/hostProjectRegistrations/{{host_project_registration_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["host_project_registration_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/hostProjectRegistrations/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{host_project_registration_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["host_project_registration_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{host_project_registration_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["host_project_registration_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apihub_plugin":
		// Format: projects/{{project}}/locations/{{location}}/plugins/{{plugin_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["plugin_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/plugins/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{plugin_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["plugin_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{plugin_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["plugin_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apihub_plugin_instance":
		// Format: projects/{{project}}/locations/{{location}}/plugins/{{plugin}}/instances/{{plugin_instance_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["plugin"].(string)
			v3, ok3 := config["plugin_instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/plugins/%s/instances/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{plugin}}/{{plugin_instance_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["plugin"].(string)
			v3, ok3 := config["plugin_instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{plugin}}/{{plugin_instance_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["plugin"].(string)
			v2, ok2 := config["plugin_instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_apikeys_key":
		// Format: projects/{{project}}/locations/global/keys/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/keys/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_app_engine_application":
		// Format: {{project-id}}
		{
			v0, ok0 := config["project-id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_app_engine_application_url_dispatch_rules":
		// Format: {{project}}
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_app_engine_domain_mapping":
		// Format: apps/{{project}}/domainMappings/{{domain_name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["domain_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("apps/%s/domainMappings/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{domain_name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["domain_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{domain_name}}
		{
			v0, ok0 := config["domain_name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_app_engine_firewall_rule":
		// Format: apps/{{project}}/firewall/ingressRules/{{priority}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("apps/%s/firewall/ingressRules/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{priority}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{priority}}
		{
			v0, ok0 := config["priority"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_app_engine_flexible_app_version":
		// Format: apps/{{project}}/services/{{service}}/versions/{{version_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["service"].(string)
			v2, ok2 := config["version_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("apps/%s/services/%s/versions/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{service}}/{{version_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["service"].(string)
			v2, ok2 := config["version_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{service}}/{{version_id}}
		{
			v0, ok0 := config["service"].(string)
			v1, ok1 := config["version_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_app_engine_service_network_settings":
		// Format: apps/{{project}}/services/{{service}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["service"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("apps/%s/services/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{service}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["service"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{service}}
		{
			v0, ok0 := config["service"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_app_engine_service_split_traffic":
		// Format: apps/{{project}}/services/{{service}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["service"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("apps/%s/services/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{service}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["service"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{service}}
		{
			v0, ok0 := config["service"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_app_engine_standard_app_version":
		// Format: apps/{{project}}/services/{{service}}/versions/{{version_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["service"].(string)
			v2, ok2 := config["version_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("apps/%s/services/%s/versions/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{service}}/{{version_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["service"].(string)
			v2, ok2 := config["version_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{service}}/{{version_id}}
		{
			v0, ok0 := config["service"].(string)
			v1, ok1 := config["version_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apphub_application":
		// Format: projects/{{project}}/locations/{{location}}/applications/{{application_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["application_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/applications/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{application_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["application_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{application_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["application_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_apphub_boundary":
		// Format: projects/{{project}}/locations/{{location}}/boundary
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/boundary", v0, v1)
			}
		}
		// Format: {{project}}/{{location}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{location}}
		{
			v0, ok0 := config["location"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_apphub_service":
		// Format: projects/{{project}}/locations/{{location}}/applications/{{application_id}}/services/{{service_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["application_id"].(string)
			v3, ok3 := config["service_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/applications/%s/services/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{application_id}}/{{service_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["application_id"].(string)
			v3, ok3 := config["service_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{application_id}}/{{service_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["application_id"].(string)
			v2, ok2 := config["service_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_apphub_service_project_attachment":
		// Format: projects/{{project}}/locations/global/serviceProjectAttachments/{{service_project_attachment_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["service_project_attachment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/serviceProjectAttachments/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{service_project_attachment_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["service_project_attachment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{service_project_attachment_id}}
		{
			v0, ok0 := config["service_project_attachment_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_apphub_workload":
		// Format: projects/{{project}}/locations/{{location}}/applications/{{application_id}}/workloads/{{workload_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["application_id"].(string)
			v3, ok3 := config["workload_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/applications/%s/workloads/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{application_id}}/{{workload_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["application_id"].(string)
			v3, ok3 := config["workload_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{application_id}}/{{workload_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["application_id"].(string)
			v2, ok2 := config["workload_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_artifact_registry_repository":
		// Format: projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["repository_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/repositories/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{repository_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["repository_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{repository_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["repository_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_artifact_registry_repository_iam":
		// No standard import format found in documentation
		return ""
	case "google_artifact_registry_rule":
		// Format: projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}/rules/{{rule_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["repository_id"].(string)
			v3, ok3 := config["rule_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/repositories/%s/rules/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{repository_id}}/{{rule_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["repository_id"].(string)
			v3, ok3 := config["rule_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{repository_id}}/{{rule_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["repository_id"].(string)
			v2, ok2 := config["rule_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_artifact_registry_vpcsc_config":
		// Format: projects/{{project}}/locations/{{location}}/vpcscConfig/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/vpcscConfig/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_assured_workloads_workload":
		// Format: organizations/{{organization}}/locations/{{location}}/workloads/{{name}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("organizations/%s/locations/%s/workloads/%s", v0, v1, v2)
			}
		}
		// Format: {{organization}}/{{location}}/{{name}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_backup_dr_backup_plan":
		// Format: projects/{{project}}/locations/{{location}}/backupPlans/{{backup_plan_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["backup_plan_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{backup_plan_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["backup_plan_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{backup_plan_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["backup_plan_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_backup_dr_backup_plan_association":
		// Format: projects/{{project}}/locations/{{location}}/backupPlanAssociations/{{backup_plan_association_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["backup_plan_association_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/backupPlanAssociations/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{backup_plan_association_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["backup_plan_association_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{backup_plan_association_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["backup_plan_association_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_backup_dr_backup_vault":
		// Format: projects/{{project}}/locations/{{location}}/backupVaults/{{backup_vault_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["backup_vault_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{backup_vault_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["backup_vault_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{backup_vault_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["backup_vault_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_backup_dr_management_server":
		// Format: projects/{{project}}/locations/{{location}}/managementServers/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/managementServers/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_backup_dr_restore_workload":
		// Format: /{{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("/%s", v0)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_backup_dr_service_config":
		// No standard import format found in documentation
		return ""
	case "google_beyondcorp_app_connection":
		// Format: projects/{{project}}/locations/{{region}}/appConnections/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/appConnections/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_beyondcorp_app_connector":
		// Format: projects/{{project}}/locations/{{region}}/appConnectors/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/appConnectors/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_beyondcorp_app_gateway":
		// Format: projects/{{project}}/locations/{{region}}/appGateways/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/appGateways/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_beyondcorp_security_gateway":
		// Format: projects/{{project}}/locations/{{location}}/securityGateways/{{security_gateway_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["security_gateway_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/securityGateways/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{security_gateway_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["security_gateway_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{security_gateway_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["security_gateway_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_beyondcorp_security_gateway_application":
		// Format: projects/{{project}}/locations/global/securityGateways/{{security_gateway_id}}/applications/{{application_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["security_gateway_id"].(string)
			v2, ok2 := config["application_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/global/securityGateways/%s/applications/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{security_gateway_id}}/{{application_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["security_gateway_id"].(string)
			v2, ok2 := config["application_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{security_gateway_id}}/{{application_id}}
		{
			v0, ok0 := config["security_gateway_id"].(string)
			v1, ok1 := config["application_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_beyondcorp_security_gateway_application_iam":
		// No standard import format found in documentation
		return ""
	case "google_beyondcorp_security_gateway_iam":
		// No standard import format found in documentation
		return ""
	case "google_biglake_catalog":
		// Format: projects/{{project}}/locations/{{location}}/catalogs/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_biglake_database":
		// Format: {{catalog}}/databases/{{name}}
		{
			v0, ok0 := config["catalog"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/databases/%s", v0, v1)
			}
		}
		return ""
	case "google_biglake_iceberg_catalog":
		// Format: iceberg/v1/restcatalog/extensions/projects/{{project}}/catalogs/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("iceberg/v1/restcatalog/extensions/projects/%s/catalogs/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_biglake_iceberg_catalog_iam":
		// No standard import format found in documentation
		return ""
	case "google_biglake_iceberg_namespace":
		// Format: projects/{{project}}/catalogs/{{catalog}}/namespaces/{{namespace_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["catalog"].(string)
			v2, ok2 := config["namespace_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{catalog}}/{{namespace_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["catalog"].(string)
			v2, ok2 := config["namespace_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{catalog}}/{{namespace_id}}
		{
			v0, ok0 := config["catalog"].(string)
			v1, ok1 := config["namespace_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_biglake_iceberg_namespace_iam":
		// No standard import format found in documentation
		return ""
	case "google_biglake_iceberg_table":
		// Format: projects/{{project}}/catalogs/{{catalog}}/namespaces/{{namespace}}/tables/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["catalog"].(string)
			v2, ok2 := config["namespace"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s/tables/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{catalog}}/{{namespace}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["catalog"].(string)
			v2, ok2 := config["namespace"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{catalog}}/{{namespace}}/{{name}}
		{
			v0, ok0 := config["catalog"].(string)
			v1, ok1 := config["namespace"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_biglake_iceberg_table_iam":
		// No standard import format found in documentation
		return ""
	case "google_biglake_table":
		// Format: {{database}}/tables/{{name}}
		{
			v0, ok0 := config["database"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/tables/%s", v0, v1)
			}
		}
		return ""
	case "google_bigquery_analytics_hub_data_exchange":
		// Format: projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_exchange_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{data_exchange_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_exchange_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{data_exchange_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["data_exchange_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{data_exchange_id}}
		{
			v0, ok0 := config["data_exchange_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_bigquery_analytics_hub_data_exchange_iam":
		// No standard import format found in documentation
		return ""
	case "google_bigquery_analytics_hub_data_exchange_subscription":
		// Format: projects/{{project}}/locations/{{location}}/subscriptions/{{subscription_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["subscription_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/subscriptions/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{subscription_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["subscription_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{subscription_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["subscription_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_bigquery_analytics_hub_listing":
		// Format: projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}/listings/{{listing_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_exchange_id"].(string)
			v3, ok3 := config["listing_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/listings/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{data_exchange_id}}/{{listing_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_exchange_id"].(string)
			v3, ok3 := config["listing_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{data_exchange_id}}/{{listing_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["data_exchange_id"].(string)
			v2, ok2 := config["listing_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_bigquery_analytics_hub_listing_iam":
		// No standard import format found in documentation
		return ""
	case "google_bigquery_analytics_hub_listing_subscription":
		// Format: projects/{{project}}/locations/{{location}}/subscriptions/{{subscription_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["subscription_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/subscriptions/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{subscription_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["subscription_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{subscription_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["subscription_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_bigquery_bi_reservation":
		// Format: projects/{{project}}/locations/{{location}}/biReservation
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/biReservation", v0, v1)
			}
		}
		// Format: {{project}}/{{location}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{location}}
		{
			v0, ok0 := config["location"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_bigquery_capacity_commitment":
		// Format: projects/{{project}}/locations/{{location}}/capacityCommitments/{{capacity_commitment_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["capacity_commitment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/capacityCommitments/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{capacity_commitment_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["capacity_commitment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{capacity_commitment_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["capacity_commitment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_bigquery_connection":
		// Format: projects/{{project}}/locations/{{location}}/connections/{{connection_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["connection_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/connections/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{connection_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["connection_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{connection_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["connection_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_bigquery_connection_iam":
		// No standard import format found in documentation
		return ""
	case "google_bigquery_data_transfer_config":
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{project}} {{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s %s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_bigquery_datapolicy_data_policy":
		// Format: projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_policy_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{data_policy_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_policy_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{data_policy_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["data_policy_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_bigquery_datapolicy_data_policy_iam":
		// No standard import format found in documentation
		return ""
	case "google_bigquery_datapolicyv2_data_policy":
		// Format: projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_policy_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{data_policy_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_policy_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{data_policy_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["data_policy_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_bigquery_datapolicyv2_data_policy_iam":
		// No standard import format found in documentation
		return ""
	case "google_bigquery_dataset":
		// Format: projects/{{project}}/datasets/{{dataset_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["dataset_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/datasets/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{dataset_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["dataset_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{dataset_id}}
		{
			v0, ok0 := config["dataset_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_bigquery_dataset_access":
		// No standard import format found in documentation
		return ""
	case "google_bigquery_dataset_iam":
		// Format: "projects/{{project_id}}/datasets/{{dataset_id}} roles/viewer user:foo@example.com"
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["dataset_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("\"projects/%s/datasets/%s roles/viewer user:foo@example.com\"", v0, v1)
			}
		}
		return ""
	case "google_bigquery_job":
		// Format: projects/{{project}}/jobs/{{job_id}}/location/{{location}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["job_id"].(string)
			v2, ok2 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/jobs/%s/location/%s", v0, v1, v2)
			}
		}
		// Format: projects/{{project}}/jobs/{{job_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["job_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/jobs/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{job_id}}/{{location}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["job_id"].(string)
			v2, ok2 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{job_id}}/{{location}}
		{
			v0, ok0 := config["job_id"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{job_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["job_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{job_id}}
		{
			v0, ok0 := config["job_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_bigquery_reservation":
		// Format: projects/{{project}}/locations/{{location}}/reservations/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/reservations/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_bigquery_reservation_assignment":
		// Format: projects/{{project}}/locations/{{location}}/reservations/{{reservation}}/assignments/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["reservation"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/reservations/%s/assignments/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{reservation}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["reservation"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{reservation}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["reservation"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_bigquery_reservation_group":
		// Format: projects/{{project}}/locations/{{location}}/reservationGroups/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/reservationGroups/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_bigquery_routine":
		// Format: projects/{{project}}/datasets/{{dataset_id}}/routines/{{routine_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["dataset_id"].(string)
			v2, ok2 := config["routine_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/datasets/%s/routines/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{dataset_id}}/{{routine_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["dataset_id"].(string)
			v2, ok2 := config["routine_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{dataset_id}}/{{routine_id}}
		{
			v0, ok0 := config["dataset_id"].(string)
			v1, ok1 := config["routine_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_bigquery_row_access_policy":
		// Format: projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}/rowAccessPolicies/{{policy_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["dataset_id"].(string)
			v2, ok2 := config["table_id"].(string)
			v3, ok3 := config["policy_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/datasets/%s/tables/%s/rowAccessPolicies/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{dataset_id}}/{{table_id}}/{{policy_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["dataset_id"].(string)
			v2, ok2 := config["table_id"].(string)
			v3, ok3 := config["policy_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{dataset_id}}/{{table_id}}/{{policy_id}}
		{
			v0, ok0 := config["dataset_id"].(string)
			v1, ok1 := config["table_id"].(string)
			v2, ok2 := config["policy_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_bigquery_table":
		// Format: projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["dataset_id"].(string)
			v2, ok2 := config["table_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/datasets/%s/tables/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{dataset_id}}/{{table_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["dataset_id"].(string)
			v2, ok2 := config["table_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{dataset_id}}/{{table_id}}
		{
			v0, ok0 := config["dataset_id"].(string)
			v1, ok1 := config["table_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_bigquery_table_iam":
		// No standard import format found in documentation
		return ""
	case "google_bigtable_app_profile":
		// Format: projects/{{project}}/instances/{{instance}}/appProfiles/{{app_profile_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["app_profile_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/instances/%s/appProfiles/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{instance}}/{{app_profile_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["app_profile_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{instance}}/{{app_profile_id}}
		{
			v0, ok0 := config["instance"].(string)
			v1, ok1 := config["app_profile_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_bigtable_authorized_view":
		// Format: projects/{{project}}/instances/{{instance_name}}/tables/{{table_name}}/authorizedViews/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance_name"].(string)
			v2, ok2 := config["table_name"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/instances/%s/tables/%s/authorizedViews/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{instance_name}}/{{table_name}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance_name"].(string)
			v2, ok2 := config["table_name"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{instance_name}}/{{table_name}}/{{name}}
		{
			v0, ok0 := config["instance_name"].(string)
			v1, ok1 := config["table_name"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_bigtable_gc_policy":
		// No standard import format found in documentation
		return ""
	case "google_bigtable_instance":
		// Format: projects/{{project}}/instances/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/instances/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_bigtable_instance_iam":
		// Format: "projects/{project}/instances/{instance} roles/editor user:jane@example.com"
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("\"projects/%s/instances/%s roles/editor user:jane@example.com\"", v0, v1)
			}
		}
		return ""
	case "google_bigtable_logical_view":
		// Format: projects/{{project}}/instances/{{instance}}/logicalViews/{{logical_view_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["logical_view_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/instances/%s/logicalViews/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{instance}}/{{logical_view_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["logical_view_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{instance}}/{{logical_view_id}}
		{
			v0, ok0 := config["instance"].(string)
			v1, ok1 := config["logical_view_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_bigtable_materialized_view":
		// Format: projects/{{project}}/instances/{{instance}}/materializedViews/{{materialized_view_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["materialized_view_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/instances/%s/materializedViews/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{instance}}/{{materialized_view_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["materialized_view_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{instance}}/{{materialized_view_id}}
		{
			v0, ok0 := config["instance"].(string)
			v1, ok1 := config["materialized_view_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_bigtable_schema_bundle":
		// Format: projects/{{project}}/instances/{{instance}}/tables/{{table}}/schemaBundles/{{schema_bundle_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["table"].(string)
			v3, ok3 := config["schema_bundle_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/instances/%s/tables/%s/schemaBundles/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{instance}}/{{table}}/{{schema_bundle_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["table"].(string)
			v3, ok3 := config["schema_bundle_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{instance}}/{{table}}/{{schema_bundle_id}}
		{
			v0, ok0 := config["instance"].(string)
			v1, ok1 := config["table"].(string)
			v2, ok2 := config["schema_bundle_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_bigtable_table":
		// Format: projects/{{project}}/instances/{{instance_name}}/tables/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance_name"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/instances/%s/tables/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{instance_name}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance_name"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{instance_name}}/{{name}}
		{
			v0, ok0 := config["instance_name"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_bigtable_table_iam":
		// Format: "projects/{project}/instances/{instance}/tables/{table} roles/editor user:jane@example.com"
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["table"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("\"projects/%s/instances/%s/tables/%s roles/editor user:jane@example.com\"", v0, v1, v2)
			}
		}
		return ""
	case "google_billing_account_iam":
		// Format: "{{billing_account_id}} roles/billing.user user:jane@example.com"
		{
			v0, ok0 := config["billing_account_id"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("\"%s roles/billing.user user:jane@example.com\"", v0)
			}
		}
		return ""
	case "google_billing_budget":
		// Format: billingAccounts/{{billing_account}}/budgets/{{name}}
		{
			v0, ok0 := config["billing_account"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("billingAccounts/%s/budgets/%s", v0, v1)
			}
		}
		// Format: {{billing_account}}/{{name}}
		{
			v0, ok0 := config["billing_account"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_billing_project_info":
		// Format: projects/{{project}}
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("projects/%s", v0)
			}
		}
		// Format: {{project}}
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_binary_authorization_attestor":
		// Format: projects/{{project}}/attestors/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/attestors/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_binary_authorization_attestor_iam":
		// No standard import format found in documentation
		return ""
	case "google_binary_authorization_policy":
		// Format: projects/{{project}}
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("projects/%s", v0)
			}
		}
		// Format: {{project}}
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_blockchain_node_engine_blockchain_nodes":
		// Format: projects/{{project}}/locations/{{location}}/blockchainNodes/{{blockchain_node_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["blockchain_node_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/blockchainNodes/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{blockchain_node_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["blockchain_node_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{blockchain_node_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["blockchain_node_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_certificate_manager_certificate":
		// Format: projects/{{project}}/locations/{{location}}/certificates/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/certificates/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_certificate_manager_certificate_issuance_config":
		// Format: projects/{{project}}/locations/{{location}}/certificateIssuanceConfigs/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/certificateIssuanceConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_certificate_manager_certificate_map":
		// Format: projects/{{project}}/locations/global/certificateMaps/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/certificateMaps/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_certificate_manager_certificate_map_entry":
		// Format: projects/{{project}}/locations/global/certificateMaps/{{map}}/certificateMapEntries/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["map"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/global/certificateMaps/%s/certificateMapEntries/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{map}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["map"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{map}}/{{name}}
		{
			v0, ok0 := config["map"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_certificate_manager_dns_authorization":
		// Format: projects/{{project}}/locations/{{location}}/dnsAuthorizations/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/dnsAuthorizations/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_certificate_manager_trust_config":
		// Format: projects/{{project}}/locations/{{location}}/trustConfigs/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/trustConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_ces_agent":
		// Format: projects/{{project}}/locations/{{location}}/apps/{{app}}/agents/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["app"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/apps/%s/agents/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{app}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["app"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{app}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["app"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_ces_app":
		// Format: projects/{{project}}/locations/{{location}}/apps/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/apps/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_ces_app_root_agent_association":
		// Format: projects/{{project}}/locations/{{location}}/apps/{{app_id}}/agents/{{agent_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["app_id"].(string)
			v3, ok3 := config["agent_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/apps/%s/agents/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{app_id}}/{{agent_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["app_id"].(string)
			v3, ok3 := config["agent_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{app_id}}/{{agent_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["app_id"].(string)
			v2, ok2 := config["agent_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_ces_app_version":
		// Format: projects/{{project}}/locations/{{location}}/apps/{{app}}/versions/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["app"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/apps/%s/versions/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{app}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["app"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{app}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["app"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_ces_deployment":
		// Format: projects/{{project}}/locations/{{location}}/apps/{{app}}/deployments/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["app"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/apps/%s/deployments/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{app}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["app"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{app}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["app"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_ces_evaluation":
		// Format: projects/{{project}}/locations/{{location}}/apps/{{app}}/evaluations/{{evaluation_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["app"].(string)
			v3, ok3 := config["evaluation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/apps/%s/evaluations/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{app}}/{{evaluation_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["app"].(string)
			v3, ok3 := config["evaluation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{app}}/{{evaluation_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["app"].(string)
			v2, ok2 := config["evaluation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_ces_example":
		// Format: projects/{{project}}/locations/{{location}}/apps/{{app}}/examples/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["app"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/apps/%s/examples/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{app}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["app"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{app}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["app"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_ces_guardrail":
		// Format: projects/{{project}}/locations/{{location}}/apps/{{app}}/guardrails/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["app"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/apps/%s/guardrails/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{app}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["app"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{app}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["app"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_ces_tool":
		// Format: projects/{{project}}/locations/{{location}}/apps/{{app}}/tools/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["app"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/apps/%s/tools/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{app}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["app"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{app}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["app"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_ces_toolset":
		// Format: projects/{{project}}/locations/{{location}}/apps/{{app}}/toolsets/{{toolset_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["app"].(string)
			v3, ok3 := config["toolset_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/apps/%s/toolsets/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{app}}/{{toolset_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["app"].(string)
			v3, ok3 := config["toolset_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{app}}/{{toolset_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["app"].(string)
			v2, ok2 := config["toolset_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_chronicle_dashboard_chart":
		// Format: projects/{{project}}/locations/{{location}}/instances/{{instance}}/dashboardCharts/{{chart_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["chart_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s/dashboardCharts/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{instance}}/{{chart_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["chart_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{instance}}/{{chart_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["chart_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_chronicle_data_access_label":
		// Format: projects/{{project}}/locations/{{location}}/instances/{{instance}}/dataAccessLabels/{{data_access_label_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["data_access_label_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s/dataAccessLabels/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{instance}}/{{data_access_label_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["data_access_label_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{instance}}/{{data_access_label_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["data_access_label_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_chronicle_data_access_scope":
		// Format: projects/{{project}}/locations/{{location}}/instances/{{instance}}/dataAccessScopes/{{data_access_scope_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["data_access_scope_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s/dataAccessScopes/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{instance}}/{{data_access_scope_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["data_access_scope_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{instance}}/{{data_access_scope_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["data_access_scope_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_chronicle_data_table":
		// Format: projects/{{project}}/locations/{{location}}/instances/{{instance}}/dataTables/{{data_table_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["data_table_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s/dataTables/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{instance}}/{{data_table_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["data_table_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{instance}}/{{data_table_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["data_table_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_chronicle_data_table_row":
		// Format: projects/{{project}}/locations/{{location}}/instances/{{instance}}/dataTables/{{data_table_id}}/dataTableRows/{{data_table_row}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["data_table_id"].(string)
			v4, ok4 := config["data_table_row"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s/dataTables/%s/dataTableRows/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{project}}/{{location}}/{{instance}}/{{data_table_id}}/{{data_table_row}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["data_table_id"].(string)
			v4, ok4 := config["data_table_row"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("%s/%s/%s/%s/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{location}}/{{instance}}/{{data_table_id}}/{{data_table_row}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["data_table_id"].(string)
			v3, ok3 := config["data_table_row"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_chronicle_feed":
		// Format: projects/{{project}}/locations/{{location}}/instances/{{instance}}/feeds/{{feed}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["feed"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s/feeds/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{instance}}/{{feed}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["feed"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{instance}}/{{feed}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["feed"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_chronicle_native_dashboard":
		// Format: projects/{{project}}/locations/{{location}}/instances/{{instance}}/nativeDashboards/{{dashboard_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["dashboard_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s/nativeDashboards/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{instance}}/{{dashboard_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["dashboard_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{instance}}/{{dashboard_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["dashboard_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_chronicle_reference_list":
		// Format: projects/{{project}}/locations/{{location}}/instances/{{instance}}/referenceLists/{{reference_list_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["reference_list_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s/referenceLists/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{instance}}/{{reference_list_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["reference_list_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{instance}}/{{reference_list_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["reference_list_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_chronicle_retrohunt":
		// Format: projects/{{project}}/locations/{{location}}/instances/{{instance}}/rules/{{rule}}/retrohunts/{{retrohunt}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["rule"].(string)
			v4, ok4 := config["retrohunt"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s/rules/%s/retrohunts/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{project}}/{{location}}/{{instance}}/{{rule}}/{{retrohunt}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["rule"].(string)
			v4, ok4 := config["retrohunt"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("%s/%s/%s/%s/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{location}}/{{instance}}/{{rule}}/{{retrohunt}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["rule"].(string)
			v3, ok3 := config["retrohunt"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_chronicle_rule":
		// Format: projects/{{project}}/locations/{{location}}/instances/{{instance}}/rules/{{rule_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["rule_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s/rules/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{instance}}/{{rule_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["rule_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{instance}}/{{rule_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["rule_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_chronicle_rule_deployment":
		// Format: projects/{{project}}/locations/{{location}}/instances/{{instance}}/rules/{{rule}}/deployment
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["rule"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s/rules/%s/deployment", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{instance}}/{{rule}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["rule"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{instance}}/{{rule}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["rule"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_chronicle_watchlist":
		// Format: projects/{{project}}/locations/{{location}}/instances/{{instance}}/watchlists/{{watchlist_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["watchlist_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s/watchlists/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{instance}}/{{watchlist_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["watchlist_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{instance}}/{{watchlist_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["watchlist_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_cloud_asset_folder_feed":
		// Format: folders/{{folder_id}}/feeds/{{name}}
		{
			v0, ok0 := config["folder_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("folders/%s/feeds/%s", v0, v1)
			}
		}
		// Format: {{folder_id}}/{{name}}
		{
			v0, ok0 := config["folder_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_cloud_asset_organization_feed":
		// Format: organizations/{{org_id}}/feeds/{{name}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("organizations/%s/feeds/%s", v0, v1)
			}
		}
		// Format: {{org_id}}/{{name}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_cloud_asset_project_feed":
		// Format: projects/{{project}}/feeds/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/feeds/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_cloud_identity_group":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_cloud_identity_group_membership":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_cloud_identity_policy":
		// Format: policies/{{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("policies/%s", v0)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_cloud_ids_endpoint":
		// Format: projects/{{project}}/locations/{{location}}/endpoints/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_cloud_quotas_quota_adjuster_settings":
		// Format: {{parent}}/locations/global/quotaAdjusterSettings
		{
			v0, ok0 := config["parent"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("%s/locations/global/quotaAdjusterSettings", v0)
			}
		}
		return ""
	case "google_cloud_quotas_quota_preference":
		// Format: {{parent}}/locations/global/quotaPreferences/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/locations/global/quotaPreferences/%s", v0, v1)
			}
		}
		return ""
	case "google_cloud_run_domain_mapping":
		// Format: locations/{{location}}/namespaces/{{project}}/domainmappings/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["project"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("locations/%s/namespaces/%s/domainmappings/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{project}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["project"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_cloud_run_service":
		// Format: locations/{{location}}/namespaces/{{project}}/services/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["project"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("locations/%s/namespaces/%s/services/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{project}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["project"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_cloud_run_service_iam":
		// No standard import format found in documentation
		return ""
	case "google_cloud_run_v2_job":
		// Format: projects/{{project}}/locations/{{location}}/jobs/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/jobs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_cloud_run_v2_job_iam":
		// No standard import format found in documentation
		return ""
	case "google_cloud_run_v2_service":
		// Format: projects/{{project}}/locations/{{location}}/services/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/services/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_cloud_run_v2_service_iam":
		// No standard import format found in documentation
		return ""
	case "google_cloud_run_v2_worker_pool":
		// Format: projects/{{project}}/locations/{{location}}/workerPools/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/workerPools/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_cloud_run_v2_worker_pool_iam":
		// No standard import format found in documentation
		return ""
	case "google_cloud_scheduler_job":
		// Format: projects/{{project}}/locations/{{region}}/jobs/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/jobs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_cloud_security_compliance_cloud_control":
		// Format: organizations/{{organization}}/locations/{{location}}/cloudControls/{{cloud_control_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cloud_control_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("organizations/%s/locations/%s/cloudControls/%s", v0, v1, v2)
			}
		}
		// Format: {{organization}}/{{location}}/{{cloud_control_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cloud_control_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_cloud_security_compliance_framework":
		// Format: organizations/{{organization}}/locations/{{location}}/frameworks/{{framework_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["framework_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("organizations/%s/locations/%s/frameworks/%s", v0, v1, v2)
			}
		}
		// Format: {{organization}}/{{location}}/{{framework_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["framework_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_cloud_security_compliance_framework_deployment":
		// Format: organizations/{{organization}}/locations/{{location}}/frameworkDeployments/{{framework_deployment_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["framework_deployment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("organizations/%s/locations/%s/frameworkDeployments/%s", v0, v1, v2)
			}
		}
		// Format: {{organization}}/{{location}}/{{framework_deployment_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["framework_deployment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_cloud_tasks_queue":
		// Format: projects/{{project}}/locations/{{location}}/queues/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/queues/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_cloud_tasks_queue_iam":
		// No standard import format found in documentation
		return ""
	case "google_cloudbuild_bitbucket_server_config":
		// Format: projects/{{project}}/locations/{{location}}/bitbucketServerConfigs/{{config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/bitbucketServerConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{config_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_cloudbuild_trigger":
		// Format: projects/{{project}}/locations/{{location}}/triggers/{{trigger_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["trigger_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/triggers/%s", v0, v1, v2)
			}
		}
		// Format: projects/{{project}}/triggers/{{trigger_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["trigger_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/triggers/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{trigger_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["trigger_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{trigger_id}}
		{
			v0, ok0 := config["trigger_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_cloudbuild_worker_pool":
		// Format: projects/{{project}}/locations/{{location}}/workerPools/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/workerPools/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_cloudbuildv2_connection":
		// Format: projects/{{project}}/locations/{{location}}/connections/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/connections/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_cloudbuildv2_connection_iam":
		// No standard import format found in documentation
		return ""
	case "google_cloudbuildv2_repository":
		// Format: projects/{{project}}/locations/{{location}}/connections/{{parent_connection}}/repositories/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["parent_connection"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/connections/%s/repositories/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{parent_connection}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["parent_connection"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{parent_connection}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["parent_connection"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_clouddeploy_automation":
		// Format: projects/{{project}}/locations/{{location}}/deliveryPipelines/{{delivery_pipeline}}/automations/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["delivery_pipeline"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/automations/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{delivery_pipeline}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["delivery_pipeline"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{delivery_pipeline}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["delivery_pipeline"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_clouddeploy_custom_target_type":
		// Format: projects/{{project}}/locations/{{location}}/customTargetTypes/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/customTargetTypes/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_clouddeploy_custom_target_type_iam":
		// No standard import format found in documentation
		return ""
	case "google_clouddeploy_delivery_pipeline":
		// Format: projects/{{project}}/locations/{{location}}/deliveryPipelines/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_clouddeploy_delivery_pipeline_iam":
		// No standard import format found in documentation
		return ""
	case "google_clouddeploy_deploy_policy":
		// Format: projects/{{project}}/locations/{{location}}/deployPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/deployPolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_clouddeploy_target":
		// Format: projects/{{project}}/locations/{{location}}/targets/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/targets/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_clouddeploy_target_iam":
		// No standard import format found in documentation
		return ""
	case "google_clouddomains_registration":
		// Format: projects/{{project}}/locations/{{location}}/registrations/{{domain_name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["domain_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/registrations/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{domain_name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["domain_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{domain_name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["domain_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_cloudfunctions2_function":
		// Format: projects/{{project}}/locations/{{location}}/functions/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/functions/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_cloudfunctions2_function_iam":
		// No standard import format found in documentation
		return ""
	case "google_cloudfunctions_function":
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_cloudfunctions_function_iam":
		// No standard import format found in documentation
		return ""
	case "google_colab_notebook_execution":
		// Format: projects/{{project}}/locations/{{location}}/notebookExecutionJobs/{{notebook_execution_job_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["notebook_execution_job_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/notebookExecutionJobs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{notebook_execution_job_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["notebook_execution_job_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{notebook_execution_job_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["notebook_execution_job_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_colab_runtime":
		// Format: projects/{{project}}/locations/{{location}}/notebookRuntimes/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/notebookRuntimes/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_colab_runtime_template":
		// Format: projects/{{project}}/locations/{{location}}/notebookRuntimeTemplates/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/notebookRuntimeTemplates/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_colab_runtime_template_iam":
		// No standard import format found in documentation
		return ""
	case "google_colab_schedule":
		// Format: projects/{{project}}/locations/{{location}}/schedules/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/schedules/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_composer_environment":
		// Format: projects/{{project}}/locations/{{region}}/environments/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/environments/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_composer_user_workloads_config_map":
		// Format: projects/{{project}}/locations/{{region}}/environments/{{environment}}/userWorkloadsConfigMaps/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["environment"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/environments/%s/userWorkloadsConfigMaps/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{region}}/{{environment}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["environment"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{region}}/{{environment}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["environment"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{environment}}/{{name}}
		{
			v0, ok0 := config["environment"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_composer_user_workloads_secret":
		// Format: projects/{{project}}/locations/{{region}}/environments/{{environment}}/userWorkloadsSecrets/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["environment"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/environments/%s/userWorkloadsSecrets/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{region}}/{{environment}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["environment"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{environment}}/{{name}}
		{
			v0, ok0 := config["environment"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_address":
		// Format: projects/{{project}}/regions/{{region}}/addresses/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/addresses/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_attached_disk":
		// Format: projects/{{project}}/zones/{{zone}}/instances/{{instance.name}}/{{disk.name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["instance.name"].(string)
			v3, ok3 := config["disk.name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/instances/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{zone}}/{{instance.name}}/{{disk.name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["instance.name"].(string)
			v3, ok3 := config["disk.name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_compute_autoscaler":
		// Format: projects/{{project}}/zones/{{zone}}/autoscalers/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/autoscalers/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{zone}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{zone}}/{{name}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_backend_bucket":
		// Format: projects/{{project}}/global/backendBuckets/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/backendBuckets/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_backend_bucket_iam":
		// No standard import format found in documentation
		return ""
	case "google_compute_backend_bucket_signed_url_key":
		// No standard import format found in documentation
		return ""
	case "google_compute_backend_service":
		// Format: projects/{{project}}/global/backendServices/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/backendServices/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_backend_service_iam":
		// No standard import format found in documentation
		return ""
	case "google_compute_backend_service_signed_url_key":
		// No standard import format found in documentation
		return ""
	case "google_compute_cross_site_network":
		// Format: projects/{{project}}/global/crossSiteNetworks/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/crossSiteNetworks/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_disk":
		// Format: projects/{{project}}/zones/{{zone}}/disks/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/disks/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{zone}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{zone}}/{{name}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_disk_async_replication":
		// No standard import format found in documentation
		return ""
	case "google_compute_disk_iam":
		// No standard import format found in documentation
		return ""
	case "google_compute_disk_resource_policy_attachment":
		// Format: projects/{{project}}/zones/{{zone}}/disks/{{disk}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["disk"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/disks/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{zone}}/{{disk}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["disk"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{zone}}/{{disk}}/{{name}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["disk"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{disk}}/{{name}}
		{
			v0, ok0 := config["disk"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_external_vpn_gateway":
		// Format: projects/{{project}}/global/externalVpnGateways/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/externalVpnGateways/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_firewall":
		// Format: projects/{{project}}/global/firewalls/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/firewalls/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_firewall_policy":
		// Format: locations/global/firewallPolicies/{{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("locations/global/firewallPolicies/%s", v0)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_firewall_policy_association":
		// Format: locations/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}
		{
			v0, ok0 := config["firewall_policy"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("locations/global/firewallPolicies/%s/associations/%s", v0, v1)
			}
		}
		// Format: {{firewall_policy}}/{{name}}
		{
			v0, ok0 := config["firewall_policy"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_firewall_policy_rule":
		// Format: locations/global/firewallPolicies/{{firewall_policy}}/rules/{{priority}}
		{
			v0, ok0 := config["firewall_policy"].(string)
			v1, ok1 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("locations/global/firewallPolicies/%s/rules/%s", v0, v1)
			}
		}
		// Format: {{firewall_policy}}/{{priority}}
		{
			v0, ok0 := config["firewall_policy"].(string)
			v1, ok1 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_firewall_policy_with_rules":
		// Format: locations/global/firewallPolicies/{{policy_id}}
		{
			v0, ok0 := config["policy_id"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("locations/global/firewallPolicies/%s", v0)
			}
		}
		// Format: {{policy_id}}
		{
			v0, ok0 := config["policy_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_forwarding_rule":
		// Format: projects/{{project}}/regions/{{region}}/forwardingRules/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/forwardingRules/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_future_reservation":
		// Format: projects/{{project}}/zones/{{zone}}/futureReservations/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/futureReservations/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{zone}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{zone}}/{{name}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_global_address":
		// Format: projects/{{project}}/global/addresses/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/addresses/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_global_forwarding_rule":
		// Format: projects/{{project}}/global/forwardingRules/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/forwardingRules/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_global_network_endpoint":
		// Format: projects/{{project}}/global/networkEndpointGroups/{{global_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["global_network_endpoint_group"].(string)
			v2, ok2 := config["ip_address"].(string)
			v3, ok3 := config["fqdn"].(string)
			v4, ok4 := config["port"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("projects/%s/global/networkEndpointGroups/%s/%s/%s/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{project}}/{{global_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["global_network_endpoint_group"].(string)
			v2, ok2 := config["ip_address"].(string)
			v3, ok3 := config["fqdn"].(string)
			v4, ok4 := config["port"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("%s/%s/%s/%s/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{global_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}
		{
			v0, ok0 := config["global_network_endpoint_group"].(string)
			v1, ok1 := config["ip_address"].(string)
			v2, ok2 := config["fqdn"].(string)
			v3, ok3 := config["port"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_compute_global_network_endpoint_group":
		// Format: projects/{{project}}/global/networkEndpointGroups/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/networkEndpointGroups/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_ha_vpn_gateway":
		// Format: projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/vpnGateways/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_health_check":
		// Format: projects/{{project}}/global/healthChecks/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/healthChecks/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_http_health_check":
		// Format: projects/{{project}}/global/httpHealthChecks/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/httpHealthChecks/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_https_health_check":
		// Format: projects/{{project}}/global/httpsHealthChecks/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/httpsHealthChecks/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_image":
		// Format: projects/{{project}}/global/images/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/images/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_image_iam":
		// No standard import format found in documentation
		return ""
	case "google_compute_instance":
		// Format: projects/{{project}}/zones/{{zone}}/instances/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/instances/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{zone}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_instance_from_machine_image":
		// No standard import format found in documentation
		return ""
	case "google_compute_instance_from_template":
		// No standard import format found in documentation
		return ""
	case "google_compute_instance_group":
		// Format: projects/{{project_id}}/zones/{{zone}}/instanceGroups/{{instance_group_id}}
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["instance_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/instanceGroups/%s", v0, v1, v2)
			}
		}
		// Format: {{project_id}}/{{zone}}/{{instance_group_id}}
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["instance_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{zone}}/{{instance_group_id}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["instance_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_instance_group_manager":
		// Format: projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/instanceGroupManagers/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{zone}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_instance_group_membership":
		// Format: projects/{{project}}/zones/{{zone}}/instanceGroups/{{instance_group}}/{{instance}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["instance_group"].(string)
			v3, ok3 := config["instance"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/instanceGroups/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{zone}}/{{instance_group}}/{{instance}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["instance_group"].(string)
			v3, ok3 := config["instance"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{zone}}/{{instance_group}}/{{instance}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["instance_group"].(string)
			v2, ok2 := config["instance"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{instance_group}}/{{instance}}
		{
			v0, ok0 := config["instance_group"].(string)
			v1, ok1 := config["instance"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_instance_group_named_port":
		// Format: projects/{{project}}/zones/{{zone}}/instanceGroups/{{group}}/{{port}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["group"].(string)
			v3, ok3 := config["port"].(string)
			v4, ok4 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/instanceGroups/%s/%s/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{project}}/{{zone}}/{{group}}/{{port}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["group"].(string)
			v3, ok3 := config["port"].(string)
			v4, ok4 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("%s/%s/%s/%s/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{zone}}/{{group}}/{{port}}/{{name}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["group"].(string)
			v2, ok2 := config["port"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{group}}/{{port}}/{{name}}
		{
			v0, ok0 := config["group"].(string)
			v1, ok1 := config["port"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_compute_instance_iam":
		// No standard import format found in documentation
		return ""
	case "google_compute_instance_settings":
		// Format: projects/{{project}}/zones/{{zone}}/instanceSettings
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/instanceSettings", v0, v1)
			}
		}
		// Format: {{project}}/{{zone}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{zone}}
		{
			v0, ok0 := config["zone"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_instance_template":
		// Format: projects/{{project}}/global/instanceTemplates/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/instanceTemplates/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_instance_template_iam":
		// No standard import format found in documentation
		return ""
	case "google_compute_instant_snapshot":
		// Format: projects/{{project}}/zones/{{zone}}/instantSnapshots/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/instantSnapshots/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{zone}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{zone}}/{{name}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_instant_snapshot_iam":
		// No standard import format found in documentation
		return ""
	case "google_compute_interconnect":
		// Format: projects/{{project}}/global/interconnects/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/interconnects/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_interconnect_attachment":
		// Format: projects/{{project}}/regions/{{region}}/interconnectAttachments/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/interconnectAttachments/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_interconnect_attachment_group":
		// Format: projects/{{project}}/global/interconnectAttachmentGroups/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/interconnectAttachmentGroups/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_interconnect_group":
		// Format: projects/{{project}}/global/interconnectGroups/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/interconnectGroups/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_machine_image":
		// Format: projects/{{project}}/global/machineImages/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/machineImages/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_machine_image_iam":
		// No standard import format found in documentation
		return ""
	case "google_compute_managed_ssl_certificate":
		// Format: projects/{{project}}/global/sslCertificates/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/sslCertificates/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_network":
		// Format: projects/{{project}}/global/networks/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/networks/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_network_attachment":
		// Format: projects/{{project}}/regions/{{region}}/networkAttachments/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/networkAttachments/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_network_edge_security_service":
		// Format: projects/{{project}}/regions/{{region}}/networkEdgeSecurityServices/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/networkEdgeSecurityServices/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_network_endpoint":
		// Format: projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["network_endpoint_group"].(string)
			v3, ok3 := config["instance"].(string)
			v4, ok4 := config["ip_address"].(string)
			v5, ok5 := config["port"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" && ok5 && v5 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/networkEndpointGroups/%s/%s/%s/%s", v0, v1, v2, v3, v4, v5)
			}
		}
		// Format: {{project}}/{{zone}}/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["network_endpoint_group"].(string)
			v3, ok3 := config["instance"].(string)
			v4, ok4 := config["ip_address"].(string)
			v5, ok5 := config["port"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" && ok5 && v5 != "" {
				return fmt.Sprintf("%s/%s/%s/%s/%s/%s", v0, v1, v2, v3, v4, v5)
			}
		}
		// Format: {{zone}}/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["network_endpoint_group"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["ip_address"].(string)
			v4, ok4 := config["port"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("%s/%s/%s/%s/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}
		{
			v0, ok0 := config["network_endpoint_group"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["ip_address"].(string)
			v3, ok3 := config["port"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_compute_network_endpoint_group":
		// Format: projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/networkEndpointGroups/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{zone}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{zone}}/{{name}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_network_endpoints":
		// Format: projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["network_endpoint_group"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/networkEndpointGroups/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{zone}}/{{network_endpoint_group}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["network_endpoint_group"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{zone}}/{{network_endpoint_group}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["network_endpoint_group"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{network_endpoint_group}}
		{
			v0, ok0 := config["network_endpoint_group"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_network_firewall_policy":
		// Format: projects/{{project}}/global/firewallPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/firewallPolicies/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_network_firewall_policy_association":
		// Format: projects/{{project}}/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["firewall_policy"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/global/firewallPolicies/%s/associations/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{firewall_policy}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["firewall_policy"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{firewall_policy}}/{{name}}
		{
			v0, ok0 := config["firewall_policy"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_network_firewall_policy_packet_mirroring_rule":
		// Format: projects/{{project}}/global/firewallPolicies/{{firewall_policy}}/packetMirroringRules/{{priority}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["firewall_policy"].(string)
			v2, ok2 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/global/firewallPolicies/%s/packetMirroringRules/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{firewall_policy}}/{{priority}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["firewall_policy"].(string)
			v2, ok2 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{firewall_policy}}/{{priority}}
		{
			v0, ok0 := config["firewall_policy"].(string)
			v1, ok1 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_network_firewall_policy_rule":
		// Format: projects/{{project}}/global/firewallPolicies/{{firewall_policy}}/rules/{{priority}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["firewall_policy"].(string)
			v2, ok2 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/global/firewallPolicies/%s/rules/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{firewall_policy}}/{{priority}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["firewall_policy"].(string)
			v2, ok2 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{firewall_policy}}/{{priority}}
		{
			v0, ok0 := config["firewall_policy"].(string)
			v1, ok1 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_network_firewall_policy_with_rules":
		// Format: projects/{{project}}/global/firewallPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/firewallPolicies/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_network_peering":
		// Format: {{project_id}}/{{network_id}}/{{peering_id}}
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["network_id"].(string)
			v2, ok2 := config["peering_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_compute_network_peering_routes_config":
		// Format: projects/{{project}}/global/networks/{{network}}/networkPeerings/{{peering}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["network"].(string)
			v2, ok2 := config["peering"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/global/networks/%s/networkPeerings/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{network}}/{{peering}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["network"].(string)
			v2, ok2 := config["peering"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{network}}/{{peering}}
		{
			v0, ok0 := config["network"].(string)
			v1, ok1 := config["peering"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_node_group":
		// Format: projects/{{project}}/zones/{{zone}}/nodeGroups/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/nodeGroups/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{zone}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{zone}}/{{name}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_node_template":
		// Format: projects/{{project}}/regions/{{region}}/nodeTemplates/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/nodeTemplates/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_organization_security_policy":
		// Format: locations/global/securityPolicies/{{policy_id}}
		{
			v0, ok0 := config["policy_id"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("locations/global/securityPolicies/%s", v0)
			}
		}
		// Format: {{policy_id}}
		{
			v0, ok0 := config["policy_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_organization_security_policy_association":
		// Format: {{policy_id}}/association/{{name}}
		{
			v0, ok0 := config["policy_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/association/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_organization_security_policy_rule":
		// Format: {{policy_id}}/priority/{{priority}}
		{
			v0, ok0 := config["policy_id"].(string)
			v1, ok1 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/priority/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_packet_mirroring":
		// Format: projects/{{project}}/regions/{{region}}/packetMirrorings/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/packetMirrorings/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_per_instance_config":
		// Format: projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{instance_group_manager}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["instance_group_manager"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/instanceGroupManagers/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{zone}}/{{instance_group_manager}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["instance_group_manager"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{zone}}/{{instance_group_manager}}/{{name}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["instance_group_manager"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{instance_group_manager}}/{{name}}
		{
			v0, ok0 := config["instance_group_manager"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_preview_feature":
		// Format: projects/{{project}}/global/previewFeatures/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/previewFeatures/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_project_cloud_armor_tier":
		// Format: projects/{{project}}
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("projects/%s", v0)
			}
		}
		// Format: {{project}}
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_project_default_network_tier":
		// Format: {{project_id}}
		{
			v0, ok0 := config["project_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_project_metadata":
		// Format: {{project_id}}
		{
			v0, ok0 := config["project_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_project_metadata_item":
		// Format: {{key}}
		{
			v0, ok0 := config["key"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		// Format: projects/{{project}}/meta-data/{{key}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["key"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/meta-data/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_public_advertised_prefix":
		// Format: projects/{{project}}/global/publicAdvertisedPrefixes/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/publicAdvertisedPrefixes/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_public_delegated_prefix":
		// Format: projects/{{project}}/regions/{{region}}/publicDelegatedPrefixes/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/publicDelegatedPrefixes/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_autoscaler":
		// Format: projects/{{project}}/regions/{{region}}/autoscalers/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/autoscalers/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_backend_bucket":
		// Format: projects/{{project}}/regions/{{region}}/backendBuckets/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/backendBuckets/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_backend_bucket_iam":
		// No standard import format found in documentation
		return ""
	case "google_compute_region_backend_service":
		// Format: projects/{{project}}/regions/{{region}}/backendServices/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/backendServices/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_backend_service_iam":
		// No standard import format found in documentation
		return ""
	case "google_compute_region_commitment":
		// Format: projects/{{project}}/regions/{{region}}/commitments/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/commitments/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_composite_health_check":
		// Format: projects/{{project}}/regions/{{region}}/compositeHealthChecks/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/compositeHealthChecks/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_disk":
		// Format: projects/{{project}}/regions/{{region}}/disks/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/disks/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_disk_iam":
		// No standard import format found in documentation
		return ""
	case "google_compute_region_disk_resource_policy_attachment":
		// Format: projects/{{project}}/regions/{{region}}/disks/{{disk}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["disk"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/disks/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{region}}/{{disk}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["disk"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{region}}/{{disk}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["disk"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{disk}}/{{name}}
		{
			v0, ok0 := config["disk"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_region_health_aggregation_policy":
		// Format: projects/{{project}}/regions/{{region}}/healthAggregationPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/healthAggregationPolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_health_check":
		// Format: projects/{{project}}/regions/{{region}}/healthChecks/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/healthChecks/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_health_source":
		// Format: projects/{{project}}/regions/{{region}}/healthSources/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/healthSources/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_instance_group_manager":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_instance_template":
		// Format: projects/{{project}}/regions/{{region}}/instanceTemplates/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/instanceTemplates/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_network_endpoint":
		// Format: projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{region_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["region_network_endpoint_group"].(string)
			v3, ok3 := config["ip_address"].(string)
			v4, ok4 := config["fqdn"].(string)
			v5, ok5 := config["port"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" && ok5 && v5 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/networkEndpointGroups/%s/%s/%s/%s", v0, v1, v2, v3, v4, v5)
			}
		}
		// Format: {{project}}/{{region}}/{{region_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["region_network_endpoint_group"].(string)
			v3, ok3 := config["ip_address"].(string)
			v4, ok4 := config["fqdn"].(string)
			v5, ok5 := config["port"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" && ok5 && v5 != "" {
				return fmt.Sprintf("%s/%s/%s/%s/%s/%s", v0, v1, v2, v3, v4, v5)
			}
		}
		// Format: {{region}}/{{region_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["region_network_endpoint_group"].(string)
			v2, ok2 := config["ip_address"].(string)
			v3, ok3 := config["fqdn"].(string)
			v4, ok4 := config["port"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("%s/%s/%s/%s/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{region_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}
		{
			v0, ok0 := config["region_network_endpoint_group"].(string)
			v1, ok1 := config["ip_address"].(string)
			v2, ok2 := config["fqdn"].(string)
			v3, ok3 := config["port"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_compute_region_network_endpoint_group":
		// Format: projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/networkEndpointGroups/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_network_firewall_policy":
		// Format: projects/{{project}}/regions/{{region}}/firewallPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/firewallPolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_network_firewall_policy_association":
		// Format: projects/{{project}}/regions/{{region}}/firewallPolicies/{{firewall_policy}}/associations/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["firewall_policy"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/firewallPolicies/%s/associations/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{region}}/{{firewall_policy}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["firewall_policy"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{region}}/{{firewall_policy}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["firewall_policy"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{firewall_policy}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["firewall_policy"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{firewall_policy}}/{{name}}
		{
			v0, ok0 := config["firewall_policy"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_region_network_firewall_policy_rule":
		// Format: projects/{{project}}/regions/{{region}}/firewallPolicies/{{firewall_policy}}/{{priority}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["firewall_policy"].(string)
			v3, ok3 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/firewallPolicies/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{region}}/{{firewall_policy}}/{{priority}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["firewall_policy"].(string)
			v3, ok3 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{region}}/{{firewall_policy}}/{{priority}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["firewall_policy"].(string)
			v2, ok2 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{firewall_policy}}/{{priority}}
		{
			v0, ok0 := config["firewall_policy"].(string)
			v1, ok1 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_region_network_firewall_policy_with_rules":
		// Format: projects/{{project}}/regions/{{region}}/firewallPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/firewallPolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_per_instance_config":
		// Format: projects/{{project}}/regions/{{region}}/instanceGroupManagers/{{region_instance_group_manager}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["region_instance_group_manager"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/instanceGroupManagers/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{region}}/{{region_instance_group_manager}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["region_instance_group_manager"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{region}}/{{region_instance_group_manager}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["region_instance_group_manager"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region_instance_group_manager}}/{{name}}
		{
			v0, ok0 := config["region_instance_group_manager"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_region_resize_request":
		// Format: projects/{{project}}/regions/{{region}}/instanceGroupManagers/{{instance_group_manager}}/resizeRequests/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["instance_group_manager"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/instanceGroupManagers/%s/resizeRequests/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{region}}/{{instance_group_manager}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["instance_group_manager"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{region}}/{{instance_group_manager}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["instance_group_manager"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{instance_group_manager}}/{{name}}
		{
			v0, ok0 := config["instance_group_manager"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_region_security_policy":
		// Format: projects/{{project}}/regions/{{region}}/securityPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/securityPolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_security_policy_rule":
		// Format: projects/{{project}}/regions/{{region}}/securityPolicies/{{security_policy}}/priority/{{priority}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["security_policy"].(string)
			v3, ok3 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/securityPolicies/%s/priority/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{region}}/{{security_policy}}/{{priority}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["security_policy"].(string)
			v3, ok3 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{region}}/{{security_policy}}/{{priority}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["security_policy"].(string)
			v2, ok2 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{security_policy}}/{{priority}}
		{
			v0, ok0 := config["security_policy"].(string)
			v1, ok1 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_region_ssl_certificate":
		// Format: projects/{{project}}/regions/{{region}}/sslCertificates/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/sslCertificates/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_ssl_policy":
		// Format: projects/{{project}}/regions/{{region}}/sslPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/sslPolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_target_http_proxy":
		// Format: projects/{{project}}/regions/{{region}}/targetHttpProxies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/targetHttpProxies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_target_https_proxy":
		// Format: projects/{{project}}/regions/{{region}}/targetHttpsProxies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/targetHttpsProxies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_target_tcp_proxy":
		// Format: projects/{{project}}/regions/{{region}}/targetTcpProxies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/targetTcpProxies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_region_url_map":
		// Format: projects/{{project}}/regions/{{region}}/urlMaps/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/urlMaps/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_reservation":
		// Format: projects/{{project}}/zones/{{zone}}/reservations/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/reservations/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{zone}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{zone}}/{{name}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_resize_request":
		// Format: projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{instance_group_manager}}/resizeRequests/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["instance_group_manager"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/instanceGroupManagers/%s/resizeRequests/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{zone}}/{{instance_group_manager}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["instance_group_manager"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{zone}}/{{instance_group_manager}}/{{name}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["instance_group_manager"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{instance_group_manager}}/{{name}}
		{
			v0, ok0 := config["instance_group_manager"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_resource_policy":
		// Format: projects/{{project}}/regions/{{region}}/resourcePolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/resourcePolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_resource_policy_attachment":
		// Format: projects/{{project}}/zones/{{zone}}/instances/{{instance}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/instances/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{zone}}/{{instance}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{zone}}/{{instance}}/{{name}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{instance}}/{{name}}
		{
			v0, ok0 := config["instance"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_rollout_plan":
		// Format: projects/{{project}}/global/rolloutPlans/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/rolloutPlans/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_route":
		// Format: projects/{{project}}/global/routes/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/routes/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_router":
		// Format: projects/{{project}}/regions/{{region}}/routers/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/routers/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_router_interface":
		// Format: {{project_id}}/{{region}}/{{router}}/{{name}}
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["router"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{region}}/{{router}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["router"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_compute_router_named_set":
		// Format: {{project}}/{{region}}/{{router}}/namedSets/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["router"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/namedSets/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{region}}/{{router}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["router"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{region}}/{{router}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["router"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{router}}/{{name}}
		{
			v0, ok0 := config["router"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_router_nat":
		// Format: projects/{{project}}/regions/{{region}}/routers/{{router}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["router"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/routers/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{region}}/{{router}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["router"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{region}}/{{router}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["router"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{router}}/{{name}}
		{
			v0, ok0 := config["router"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_router_nat_address":
		// Format: projects/{{project}}/regions/{{region}}/routers/{{router}}/{{router_nat}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["router"].(string)
			v3, ok3 := config["router_nat"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/routers/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{region}}/{{router}}/{{router_nat}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["router"].(string)
			v3, ok3 := config["router_nat"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{region}}/{{router}}/{{router_nat}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["router"].(string)
			v2, ok2 := config["router_nat"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{router}}/{{router_nat}}
		{
			v0, ok0 := config["router"].(string)
			v1, ok1 := config["router_nat"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_router_peer":
		// Format: projects/{{project}}/regions/{{region}}/routers/{{router}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["router"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/routers/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{region}}/{{router}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["router"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{region}}/{{router}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["router"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{router}}/{{name}}
		{
			v0, ok0 := config["router"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_router_route_policy":
		// Format: {{project}}/{{region}}/{{router}}/routePolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["router"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/routePolicies/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{region}}/{{router}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["router"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{region}}/{{router}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["router"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{router}}/{{name}}
		{
			v0, ok0 := config["router"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_security_policy":
		// Format: projects/{{project}}/global/securityPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/securityPolicies/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_security_policy_rule":
		// Format: projects/{{project}}/global/securityPolicies/{{security_policy}}/priority/{{priority}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["security_policy"].(string)
			v2, ok2 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/global/securityPolicies/%s/priority/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{security_policy}}/{{priority}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["security_policy"].(string)
			v2, ok2 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{security_policy}}/{{priority}}
		{
			v0, ok0 := config["security_policy"].(string)
			v1, ok1 := config["priority"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_service_attachment":
		// Format: projects/{{project}}/regions/{{region}}/serviceAttachments/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/serviceAttachments/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_shared_vpc_host_project":
		// Format: {{project_id}}
		{
			v0, ok0 := config["project_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_shared_vpc_service_project":
		// Format: {{host_project}/{{service_project}}
		{
			v0, ok0 := config["host_project"].(string)
			v1, ok1 := config["service_project"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_compute_snapshot":
		// Format: projects/{{project}}/global/snapshots/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/snapshots/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_snapshot_iam":
		// No standard import format found in documentation
		return ""
	case "google_compute_snapshot_settings":
		// Format: projects/{{project}}/global/snapshotSettings/
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("projects/%s/global/snapshotSettings/", v0)
			}
		}
		// Format: {{project}}
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_ssl_certificate":
		// Format: projects/{{project}}/global/sslCertificates/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/sslCertificates/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_ssl_policy":
		// Format: projects/{{project}}/global/sslPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/sslPolicies/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_storage_pool":
		// Format: projects/{{project}}/zones/{{zone}}/storagePools/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/storagePools/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{zone}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{zone}}/{{name}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_storage_pool_iam":
		// No standard import format found in documentation
		return ""
	case "google_compute_subnetwork":
		// Format: projects/{{project}}/regions/{{region}}/subnetworks/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_subnetwork_iam":
		// No standard import format found in documentation
		return ""
	case "google_compute_target_grpc_proxy":
		// Format: projects/{{project}}/global/targetGrpcProxies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/targetGrpcProxies/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_target_http_proxy":
		// Format: projects/{{project}}/global/targetHttpProxies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/targetHttpProxies/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_target_https_proxy":
		// Format: projects/{{project}}/global/targetHttpsProxies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/targetHttpsProxies/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_target_instance":
		// Format: projects/{{project}}/zones/{{zone}}/targetInstances/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/zones/%s/targetInstances/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{zone}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{zone}}/{{name}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_target_pool":
		// Format: projects/{{project}}/regions/{{region}}/targetPools/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/targetPools/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_target_ssl_proxy":
		// Format: projects/{{project}}/global/targetSslProxies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/targetSslProxies/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_target_tcp_proxy":
		// Format: projects/{{project}}/global/targetTcpProxies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/targetTcpProxies/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_url_map":
		// Format: projects/{{project}}/global/urlMaps/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/global/urlMaps/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_vpn_gateway":
		// Format: projects/{{project}}/regions/{{region}}/targetVpnGateways/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/targetVpnGateways/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_vpn_tunnel":
		// Format: projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/regions/%s/vpnTunnels/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_compute_wire_group":
		// Format: projects/{{project}}/global/crossSiteNetworks/{{cross_site_network}}/wireGroups/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["cross_site_network"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/global/crossSiteNetworks/%s/wireGroups/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{cross_site_network}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["cross_site_network"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{cross_site_network}}/{{name}}
		{
			v0, ok0 := config["cross_site_network"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_contact_center_insights_analysis_rule":
		// Format: projects/{{project}}/locations/{{location}}/analysisRules/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/analysisRules/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_contact_center_insights_assessment_rule":
		// Format: projects/{{project}}/locations/{{location}}/assessmentRules/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/assessmentRules/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_contact_center_insights_auto_labeling_rule":
		// Format: projects/{{project}}/locations/{{location}}/autoLabelingRules/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/autoLabelingRules/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_contact_center_insights_qa_question":
		// Format: projects/{{project}}/locations/{{location}}/qaScorecards/{{qa_scorecard}}/revisions/{{revision}}/qaQuestions/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["qa_scorecard"].(string)
			v3, ok3 := config["revision"].(string)
			v4, ok4 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s/qaQuestions/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{project}}/{{location}}/{{qa_scorecard}}/{{revision}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["qa_scorecard"].(string)
			v3, ok3 := config["revision"].(string)
			v4, ok4 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("%s/%s/%s/%s/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{location}}/{{qa_scorecard}}/{{revision}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["qa_scorecard"].(string)
			v2, ok2 := config["revision"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_contact_center_insights_qa_scorecard":
		// Format: projects/{{project}}/locations/{{location}}/qaScorecards/{{qa_scorecard_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["qa_scorecard_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{qa_scorecard_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["qa_scorecard_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{qa_scorecard_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["qa_scorecard_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_contact_center_insights_qa_scorecard_revision":
		// Format: projects/{{project}}/locations/{{location}}/qaScorecards/{{qa_scorecard}}/revisions/{{qa_scorecard_revision_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["qa_scorecard"].(string)
			v3, ok3 := config["qa_scorecard_revision_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{qa_scorecard}}/{{qa_scorecard_revision_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["qa_scorecard"].(string)
			v3, ok3 := config["qa_scorecard_revision_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{qa_scorecard}}/{{qa_scorecard_revision_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["qa_scorecard"].(string)
			v2, ok2 := config["qa_scorecard_revision_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_contact_center_insights_view":
		// Format: projects/{{project}}/locations/{{location}}/views/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/views/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_container_analysis_note":
		// Format: projects/{{project}}/notes/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/notes/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_container_analysis_note_iam":
		// No standard import format found in documentation
		return ""
	case "google_container_analysis_occurrence":
		// Format: projects/{{project}}/occurrences/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/occurrences/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_container_attached_cluster":
		// Format: projects/{{project}}/locations/{{location}}/attachedClusters/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/attachedClusters/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_container_aws_cluster":
		// Format: projects/{{project}}/locations/{{location}}/awsClusters/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_container_aws_node_pool":
		// Format: projects/{{project}}/locations/{{location}}/awsClusters/{{cluster}}/awsNodePools/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s/awsNodePools/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{cluster}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{cluster}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["cluster"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_container_azure_client":
		// Format: projects/{{project}}/locations/{{location}}/azureClients/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/azureClients/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_container_azure_cluster":
		// Format: projects/{{project}}/locations/{{location}}/azureClusters/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_container_azure_node_pool":
		// Format: projects/{{project}}/locations/{{location}}/azureClusters/{{cluster}}/azureNodePools/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s/azureNodePools/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{cluster}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{cluster}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["cluster"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_container_cluster":
		// Format: projects/{{project_id}}/locations/{{location}}/clusters/{{cluster_id}}
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", v0, v1, v2)
			}
		}
		// Format: {{project_id}}/{{location}}/{{cluster_id}}
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{cluster_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_container_node_pool":
		// Format: {{project_id}}/{{location}}/{{cluster_id}}/{{pool_id}}
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster_id"].(string)
			v3, ok3 := config["pool_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{cluster_id}}/{{pool_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["cluster_id"].(string)
			v2, ok2 := config["pool_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_container_registry":
		// No standard import format found in documentation
		return ""
	case "google_data_catalog_entry":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_data_catalog_entry_group":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_data_catalog_entry_group_iam":
		// No standard import format found in documentation
		return ""
	case "google_data_catalog_policy_tag":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_data_catalog_policy_tag_iam":
		// No standard import format found in documentation
		return ""
	case "google_data_catalog_tag":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_data_catalog_tag_template":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_data_catalog_tag_template_iam":
		// No standard import format found in documentation
		return ""
	case "google_data_catalog_taxonomy":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_data_catalog_taxonomy_iam":
		// No standard import format found in documentation
		return ""
	case "google_data_fusion_instance":
		// Format: projects/{{project}}/locations/{{region}}/instances/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_data_fusion_instance_iam":
		// No standard import format found in documentation
		return ""
	case "google_data_lineage_config":
		// Format: {{parent}}/locations/{{location}}/config
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/locations/%s/config", v0, v1)
			}
		}
		return ""
	case "google_data_loss_prevention_deidentify_template":
		// Format: {{parent}}/deidentifyTemplates/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/deidentifyTemplates/%s", v0, v1)
			}
		}
		// Format: {{parent}}/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_data_loss_prevention_discovery_config":
		// Format: {{parent}}/discoveryConfigs/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/discoveryConfigs/%s", v0, v1)
			}
		}
		// Format: {{parent}}/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_data_loss_prevention_inspect_template":
		// Format: {{parent}}/inspectTemplates/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/inspectTemplates/%s", v0, v1)
			}
		}
		// Format: {{parent}}/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_data_loss_prevention_job_trigger":
		// Format: {{parent}}/jobTriggers/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/jobTriggers/%s", v0, v1)
			}
		}
		// Format: {{parent}}/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_data_loss_prevention_stored_info_type":
		// Format: {{parent}}/storedInfoTypes/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/storedInfoTypes/%s", v0, v1)
			}
		}
		// Format: {{parent}}/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_data_pipeline_pipeline":
		// Format: projects/{{project}}/locations/{{region}}/pipelines/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/pipelines/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_database_migration_service_connection_profile":
		// Format: projects/{{project}}/locations/{{location}}/connectionProfiles/{{connection_profile_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["connection_profile_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/connectionProfiles/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{connection_profile_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["connection_profile_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{connection_profile_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["connection_profile_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_database_migration_service_migration_job":
		// Format: projects/{{project}}/locations/{{location}}/migrationJobs/{{migration_job_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["migration_job_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{migration_job_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["migration_job_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{migration_job_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["migration_job_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_database_migration_service_private_connection":
		// Format: projects/{{project}}/locations/{{location}}/privateConnections/{{private_connection_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["private_connection_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{private_connection_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["private_connection_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{private_connection_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["private_connection_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dataflow_flex_template_job":
		// No standard import format found in documentation
		return ""
	case "google_dataflow_job":
		// Format: {{id}}
		{
			v0, ok0 := config["id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_dataform_config":
		// Format: projects/{{project}}/locations/{{region}}/config
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/config", v0, v1)
			}
		}
		// Format: {{project}}/{{region}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{region}}
		{
			v0, ok0 := config["region"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_dataform_folder":
		// Format: projects/{{project}}/locations/{{region}}/folders/{{folder_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["folder_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/folders/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{folder_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["folder_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{folder_id}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["folder_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{folder_id}}
		{
			v0, ok0 := config["folder_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_dataform_repository":
		// Format: projects/{{project}}/locations/{{region}}/repositories/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/repositories/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_dataform_repository_iam":
		// No standard import format found in documentation
		return ""
	case "google_dataform_repository_release_config":
		// Format: projects/{{project}}/locations/{{region}}/repositories/{{repository}}/releaseConfigs/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["repository"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/repositories/%s/releaseConfigs/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{region}}/{{repository}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["repository"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{region}}/{{repository}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["repository"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{repository}}/{{name}}
		{
			v0, ok0 := config["repository"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dataform_repository_workflow_config":
		// Format: projects/{{project}}/locations/{{region}}/repositories/{{repository}}/workflowConfigs/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["repository"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workflowConfigs/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{region}}/{{repository}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["repository"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{region}}/{{repository}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["repository"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{repository}}/{{name}}
		{
			v0, ok0 := config["repository"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dataform_team_folder":
		// Format: projects/{{project}}/locations/{{region}}/teamFolders/{{teamfolder_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["teamfolder_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/teamFolders/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{teamfolder_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["teamfolder_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{teamfolder_id}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["teamfolder_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{teamfolder_id}}
		{
			v0, ok0 := config["teamfolder_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_dataplex_aspect_type":
		// Format: projects/{{project}}/locations/{{location}}/aspectTypes/{{aspect_type_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["aspect_type_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/aspectTypes/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{aspect_type_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["aspect_type_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{aspect_type_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["aspect_type_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dataplex_aspect_type_iam":
		// No standard import format found in documentation
		return ""
	case "google_dataplex_asset":
		// Format: projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{dataplex_zone}}/assets/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["lake"].(string)
			v3, ok3 := config["dataplex_zone"].(string)
			v4, ok4 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/assets/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{project}}/{{location}}/{{lake}}/{{dataplex_zone}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["lake"].(string)
			v3, ok3 := config["dataplex_zone"].(string)
			v4, ok4 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("%s/%s/%s/%s/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{location}}/{{lake}}/{{dataplex_zone}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["lake"].(string)
			v2, ok2 := config["dataplex_zone"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_dataplex_asset_iam":
		// No standard import format found in documentation
		return ""
	case "google_dataplex_data_asset":
		// Format: projects/{{project}}/locations/{{location}}/dataProducts/{{data_product_id}}/dataAssets/{{data_asset_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_product_id"].(string)
			v3, ok3 := config["data_asset_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/dataProducts/%s/dataAssets/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{data_product_id}}/{{data_asset_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_product_id"].(string)
			v3, ok3 := config["data_asset_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{data_product_id}}/{{data_asset_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["data_product_id"].(string)
			v2, ok2 := config["data_asset_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_dataplex_data_product":
		// Format: projects/{{project}}/locations/{{location}}/dataProducts/{{data_product_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_product_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/dataProducts/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{data_product_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_product_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{data_product_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["data_product_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dataplex_data_product_data_asset":
		// Format: projects/{{project}}/locations/{{location}}/dataProducts/{{data_product_id}}/dataAssets/{{data_asset_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_product_id"].(string)
			v3, ok3 := config["data_asset_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/dataProducts/%s/dataAssets/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{data_product_id}}/{{data_asset_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_product_id"].(string)
			v3, ok3 := config["data_asset_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{data_product_id}}/{{data_asset_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["data_product_id"].(string)
			v2, ok2 := config["data_asset_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_dataplex_datascan":
		// Format: projects/{{project}}/locations/{{location}}/dataScans/{{data_scan_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_scan_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/dataScans/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{data_scan_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_scan_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{data_scan_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["data_scan_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{data_scan_id}}
		{
			v0, ok0 := config["data_scan_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_dataplex_datascan_iam":
		// No standard import format found in documentation
		return ""
	case "google_dataplex_entry":
		// Format: projects/{{project}}/locations/{{location}}/entryGroups/{{entry_group_id}}/entries/{{entry_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["entry_group_id"].(string)
			v3, ok3 := config["entry_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{entry_group_id}}/{{entry_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["entry_group_id"].(string)
			v3, ok3 := config["entry_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{entry_group_id}}/{{entry_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["entry_group_id"].(string)
			v2, ok2 := config["entry_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_dataplex_entry_group":
		// Format: projects/{{project}}/locations/{{location}}/entryGroups/{{entry_group_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["entry_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{entry_group_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["entry_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{entry_group_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["entry_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dataplex_entry_group_iam":
		// No standard import format found in documentation
		return ""
	case "google_dataplex_entry_link":
		// Format: projects/{{project}}/locations/{{location}}/entryGroups/{{entry_group_id}}/entryLinks/{{entry_link_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["entry_group_id"].(string)
			v3, ok3 := config["entry_link_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entryLinks/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{entry_group_id}}/{{entry_link_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["entry_group_id"].(string)
			v3, ok3 := config["entry_link_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{entry_group_id}}/{{entry_link_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["entry_group_id"].(string)
			v2, ok2 := config["entry_link_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_dataplex_entry_type":
		// Format: projects/{{project}}/locations/{{location}}/entryTypes/{{entry_type_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["entry_type_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/entryTypes/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{entry_type_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["entry_type_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{entry_type_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["entry_type_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dataplex_entry_type_iam":
		// No standard import format found in documentation
		return ""
	case "google_dataplex_glossary":
		// Format: projects/{{project}}/locations/{{location}}/glossaries/{{glossary_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["glossary_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/glossaries/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{glossary_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["glossary_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{glossary_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["glossary_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dataplex_glossary_category":
		// Format: projects/{{project}}/locations/{{location}}/glossaries/{{glossary_id}}/categories/{{category_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["glossary_id"].(string)
			v3, ok3 := config["category_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/glossaries/%s/categories/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{glossary_id}}/{{category_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["glossary_id"].(string)
			v3, ok3 := config["category_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{glossary_id}}/{{category_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["glossary_id"].(string)
			v2, ok2 := config["category_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_dataplex_glossary_iam":
		// No standard import format found in documentation
		return ""
	case "google_dataplex_glossary_term":
		// Format: projects/{{project}}/locations/{{location}}/glossaries/{{glossary_id}}/terms/{{term_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["glossary_id"].(string)
			v3, ok3 := config["term_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/glossaries/%s/terms/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{glossary_id}}/{{term_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["glossary_id"].(string)
			v3, ok3 := config["term_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{glossary_id}}/{{term_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["glossary_id"].(string)
			v2, ok2 := config["term_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_dataplex_lake":
		// Format: projects/{{project}}/locations/{{location}}/lakes/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/lakes/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dataplex_lake_iam":
		// No standard import format found in documentation
		return ""
	case "google_dataplex_task":
		// Format: projects/{{project}}/locations/{{location}}/lakes/{{lake}}/tasks/{{task_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["lake"].(string)
			v3, ok3 := config["task_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/lakes/%s/tasks/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{lake}}/{{task_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["lake"].(string)
			v3, ok3 := config["task_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{lake}}/{{task_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["lake"].(string)
			v2, ok2 := config["task_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_dataplex_task_iam":
		// No standard import format found in documentation
		return ""
	case "google_dataplex_zone":
		// Format: projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["lake"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{lake}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["lake"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{lake}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["lake"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_dataplex_zone_iam":
		// No standard import format found in documentation
		return ""
	case "google_dataproc_autoscaling_policy":
		// Format: projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["policy_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/autoscalingPolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{policy_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["policy_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{policy_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["policy_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dataproc_autoscaling_policy_iam":
		// No standard import format found in documentation
		return ""
	case "google_dataproc_batch":
		// Format: projects/{{project}}/locations/{{location}}/batches/{{batch_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["batch_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/batches/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{batch_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["batch_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{batch_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["batch_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dataproc_cluster":
		// No standard import format found in documentation
		return ""
	case "google_dataproc_cluster_iam":
		// Format: "projects/{project}/regions/{region}/clusters/{cluster} roles/editor user:jane@example.com"
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["cluster"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("\"projects/%s/regions/%s/clusters/%s roles/editor user:jane@example.com\"", v0, v1, v2)
			}
		}
		return ""
	case "google_dataproc_gdc_application_environment":
		// Format: projects/{{project}}/locations/{{location}}/serviceInstances/{{serviceinstance}}/applicationEnvironments/{{application_environment_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["serviceinstance"].(string)
			v3, ok3 := config["application_environment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/serviceInstances/%s/applicationEnvironments/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{serviceinstance}}/{{application_environment_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["serviceinstance"].(string)
			v3, ok3 := config["application_environment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{serviceinstance}}/{{application_environment_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["serviceinstance"].(string)
			v2, ok2 := config["application_environment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_dataproc_gdc_service_instance":
		// Format: projects/{{project}}/locations/{{location}}/serviceInstances/{{service_instance_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["service_instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/serviceInstances/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{service_instance_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["service_instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{service_instance_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["service_instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dataproc_gdc_spark_application":
		// Format: projects/{{project}}/locations/{{location}}/serviceInstances/{{serviceinstance}}/sparkApplications/{{spark_application_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["serviceinstance"].(string)
			v3, ok3 := config["spark_application_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/serviceInstances/%s/sparkApplications/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{serviceinstance}}/{{spark_application_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["serviceinstance"].(string)
			v3, ok3 := config["spark_application_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{serviceinstance}}/{{spark_application_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["serviceinstance"].(string)
			v2, ok2 := config["spark_application_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_dataproc_job":
		// No standard import format found in documentation
		return ""
	case "google_dataproc_job_iam":
		// Format: "projects/{project}/regions/{region}/jobs/{job_id} roles/editor user:jane@example.com"
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["job_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("\"projects/%s/regions/%s/jobs/%s roles/editor user:jane@example.com\"", v0, v1, v2)
			}
		}
		return ""
	case "google_dataproc_metastore_database_iam":
		// No standard import format found in documentation
		return ""
	case "google_dataproc_metastore_federation":
		// Format: projects/{{project}}/locations/{{location}}/federations/{{federation_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["federation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/federations/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{federation_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["federation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{federation_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["federation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dataproc_metastore_federation_iam":
		// No standard import format found in documentation
		return ""
	case "google_dataproc_metastore_service":
		// Format: projects/{{project}}/locations/{{location}}/services/{{service_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["service_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/services/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{service_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["service_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{service_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["service_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dataproc_metastore_service_iam":
		// No standard import format found in documentation
		return ""
	case "google_dataproc_metastore_table_iam":
		// No standard import format found in documentation
		return ""
	case "google_dataproc_session_template":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_dataproc_workflow_template":
		// Format: projects/{{project}}/locations/{{location}}/workflowTemplates/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/workflowTemplates/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_datastream_connection_profile":
		// Format: projects/{{project}}/locations/{{location}}/connectionProfiles/{{connection_profile_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["connection_profile_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/connectionProfiles/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{connection_profile_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["connection_profile_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{connection_profile_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["connection_profile_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_datastream_private_connection":
		// Format: projects/{{project}}/locations/{{location}}/privateConnections/{{private_connection_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["private_connection_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{private_connection_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["private_connection_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{private_connection_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["private_connection_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_datastream_stream":
		// Format: projects/{{project}}/locations/{{location}}/streams/{{stream_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["stream_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/streams/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{stream_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["stream_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{stream_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["stream_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_deployment_manager_deployment":
		// Format: projects/{{project}}/deployments/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/deployments/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_developer_connect_account_connector":
		// Format: projects/{{project}}/locations/{{location}}/accountConnectors/{{account_connector_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["account_connector_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/accountConnectors/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{account_connector_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["account_connector_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{account_connector_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["account_connector_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_developer_connect_connection":
		// Format: projects/{{project}}/locations/{{location}}/connections/{{connection_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["connection_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/connections/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{connection_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["connection_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{connection_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["connection_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_developer_connect_git_repository_link":
		// Format: projects/{{project}}/locations/{{location}}/connections/{{parent_connection}}/gitRepositoryLinks/{{git_repository_link_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["parent_connection"].(string)
			v3, ok3 := config["git_repository_link_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/connections/%s/gitRepositoryLinks/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{parent_connection}}/{{git_repository_link_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["parent_connection"].(string)
			v3, ok3 := config["git_repository_link_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{parent_connection}}/{{git_repository_link_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["parent_connection"].(string)
			v2, ok2 := config["git_repository_link_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_developer_connect_insights_config":
		// Format: projects/{{project}}/locations/{{location}}/insightsConfigs/{{insights_config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["insights_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/insightsConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{insights_config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["insights_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{insights_config_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["insights_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dialogflow_agent":
		// Format: {{project}}
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_dialogflow_conversation_profile":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_dialogflow_cx_agent":
		// Format: projects/{{project}}/locations/{{location}}/agents/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/agents/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dialogflow_cx_entity_type":
		// Format: {{parent}}/entityTypes/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/entityTypes/%s", v0, v1)
			}
		}
		// Format: {{parent}}/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dialogflow_cx_environment":
		// Format: {{parent}}/environments/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/environments/%s", v0, v1)
			}
		}
		// Format: {{parent}}/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dialogflow_cx_flow":
		// Format: {{parent}}/flows/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/flows/%s", v0, v1)
			}
		}
		// Format: {{parent}}/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dialogflow_cx_generative_settings":
		// Format: {{parent}}/generativeSettings
		{
			v0, ok0 := config["parent"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("%s/generativeSettings", v0)
			}
		}
		// Format: {{parent}}
		{
			v0, ok0 := config["parent"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_dialogflow_cx_generator":
		// Format: {{parent}}/generators/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/generators/%s", v0, v1)
			}
		}
		// Format: {{parent}}/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dialogflow_cx_intent":
		// Format: {{parent}}/intents/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/intents/%s", v0, v1)
			}
		}
		// Format: {{parent}}/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dialogflow_cx_page":
		// Format: {{parent}}/pages/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/pages/%s", v0, v1)
			}
		}
		// Format: {{parent}}/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dialogflow_cx_playbook":
		// Format: {{parent}}/playbooks/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/playbooks/%s", v0, v1)
			}
		}
		// Format: {{parent}}/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dialogflow_cx_security_settings":
		// Format: projects/{{project}}/locations/{{location}}/securitySettings/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/securitySettings/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dialogflow_cx_test_case":
		// Format: {{parent}}/testCases/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/testCases/%s", v0, v1)
			}
		}
		return ""
	case "google_dialogflow_cx_tool":
		// Format: {{parent}}/tools/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/tools/%s", v0, v1)
			}
		}
		// Format: {{parent}}/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dialogflow_cx_tool_version":
		// Format: {{parent}}/versions/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/versions/%s", v0, v1)
			}
		}
		// Format: {{parent}}/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dialogflow_cx_version":
		// Format: {{parent}}/versions/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/versions/%s", v0, v1)
			}
		}
		// Format: {{parent}}/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dialogflow_cx_webhook":
		// Format: {{parent}}/webhooks/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/webhooks/%s", v0, v1)
			}
		}
		// Format: {{parent}}/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dialogflow_encryption_spec":
		// No standard import format found in documentation
		return ""
	case "google_dialogflow_entity_type":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_dialogflow_environment":
		// Format: projects/{{project}}/locations/{{location}}/agent/environments/{{environmentid}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["environmentid"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/agent/environments/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{environmentid}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["environmentid"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{environmentid}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["environmentid"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dialogflow_fulfillment":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_dialogflow_generator":
		// Format: projects/{{project}}/locations/{{location}}/generators/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/generators/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_dialogflow_intent":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_dialogflow_version":
		// Format: {{parent}}/versions/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/versions/%s", v0, v1)
			}
		}
		// Format: {{parent}}/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_discovery_engine_acl_config":
		// Format: projects/{{project}}/locations/{{location}}/aclConfig
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/aclConfig", v0, v1)
			}
		}
		// Format: {{project}}/{{location}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{location}}
		{
			v0, ok0 := config["location"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_discovery_engine_assistant":
		// Format: projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines/{{engine_id}}/assistants/{{assistant_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["collection_id"].(string)
			v3, ok3 := config["engine_id"].(string)
			v4, ok4 := config["assistant_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/collections/%s/engines/%s/assistants/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{project}}/{{location}}/{{collection_id}}/{{engine_id}}/{{assistant_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["collection_id"].(string)
			v3, ok3 := config["engine_id"].(string)
			v4, ok4 := config["assistant_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("%s/%s/%s/%s/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{location}}/{{collection_id}}/{{engine_id}}/{{assistant_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["collection_id"].(string)
			v2, ok2 := config["engine_id"].(string)
			v3, ok3 := config["assistant_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_discovery_engine_chat_engine":
		// Format: projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines/{{engine_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["collection_id"].(string)
			v3, ok3 := config["engine_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/collections/%s/engines/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{collection_id}}/{{engine_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["collection_id"].(string)
			v3, ok3 := config["engine_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{collection_id}}/{{engine_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["collection_id"].(string)
			v2, ok2 := config["engine_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_discovery_engine_cmek_config":
		// Format: projects/{{project}}/locations/{{location}}/cmekConfigs/{{cmek_config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cmek_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/cmekConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{cmek_config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cmek_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{cmek_config_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["cmek_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_discovery_engine_control":
		// Format: projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines/{{engine_id}}/controls/{{control_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["collection_id"].(string)
			v3, ok3 := config["engine_id"].(string)
			v4, ok4 := config["control_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/collections/%s/engines/%s/controls/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{project}}/{{location}}/{{collection_id}}/{{engine_id}}/{{control_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["collection_id"].(string)
			v3, ok3 := config["engine_id"].(string)
			v4, ok4 := config["control_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("%s/%s/%s/%s/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{location}}/{{collection_id}}/{{engine_id}}/{{control_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["collection_id"].(string)
			v2, ok2 := config["engine_id"].(string)
			v3, ok3 := config["control_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_discovery_engine_data_connector":
		// Format: projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/dataConnector
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["collection_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataConnector", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{collection_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["collection_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{collection_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["collection_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_discovery_engine_data_store":
		// Format: projects/{{project}}/locations/{{location}}/collections/default_collection/dataStores/{{data_store_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_store_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/collections/default_collection/dataStores/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{data_store_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_store_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{data_store_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["data_store_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_discovery_engine_license_config":
		// Format: projects/{{project}}/locations/{{location}}/licenseConfigs/{{license_config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["license_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/licenseConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{license_config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["license_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{license_config_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["license_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_discovery_engine_recommendation_engine":
		// Format: projects/{{project}}/locations/{{location}}/collections/default_collection/engines/{{engine_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["engine_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/collections/default_collection/engines/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{engine_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["engine_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{engine_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["engine_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_discovery_engine_schema":
		// Format: projects/{{project}}/locations/{{location}}/collections/default_collection/dataStores/{{data_store_id}}/schemas/{{schema_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_store_id"].(string)
			v3, ok3 := config["schema_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/collections/default_collection/dataStores/%s/schemas/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{data_store_id}}/{{schema_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_store_id"].(string)
			v3, ok3 := config["schema_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{data_store_id}}/{{schema_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["data_store_id"].(string)
			v2, ok2 := config["schema_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_discovery_engine_search_engine":
		// Format: projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines/{{engine_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["collection_id"].(string)
			v3, ok3 := config["engine_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/collections/%s/engines/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{collection_id}}/{{engine_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["collection_id"].(string)
			v3, ok3 := config["engine_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{collection_id}}/{{engine_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["collection_id"].(string)
			v2, ok2 := config["engine_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_discovery_engine_serving_config":
		// Format: projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines/{{engine_id}}/servingConfigs/{{serving_config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["collection_id"].(string)
			v3, ok3 := config["engine_id"].(string)
			v4, ok4 := config["serving_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/collections/%s/engines/%s/servingConfigs/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{project}}/{{location}}/{{collection_id}}/{{engine_id}}/{{serving_config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["collection_id"].(string)
			v3, ok3 := config["engine_id"].(string)
			v4, ok4 := config["serving_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("%s/%s/%s/%s/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{location}}/{{collection_id}}/{{engine_id}}/{{serving_config_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["collection_id"].(string)
			v2, ok2 := config["engine_id"].(string)
			v3, ok3 := config["serving_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_discovery_engine_sitemap":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_discovery_engine_target_site":
		// Format: projects/{{project}}/locations/{{location}}/collections/default_collection/dataStores/{{data_store_id}}/siteSearchEngine/targetSites/{{target_site_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_store_id"].(string)
			v3, ok3 := config["target_site_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/collections/default_collection/dataStores/%s/siteSearchEngine/targetSites/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{data_store_id}}/{{target_site_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_store_id"].(string)
			v3, ok3 := config["target_site_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{data_store_id}}/{{target_site_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["data_store_id"].(string)
			v2, ok2 := config["target_site_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_discovery_engine_user_store":
		// Format: projects/{{project}}/locations/{{location}}/userStores/{{user_store_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["user_store_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/userStores/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{user_store_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["user_store_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{user_store_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["user_store_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_discovery_engine_widget_config":
		// Format: projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines/{{engine_id}}/widgetConfigs/{{widget_config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["collection_id"].(string)
			v3, ok3 := config["engine_id"].(string)
			v4, ok4 := config["widget_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/collections/%s/engines/%s/widgetConfigs/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{project}}/{{location}}/{{collection_id}}/{{engine_id}}/{{widget_config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["collection_id"].(string)
			v3, ok3 := config["engine_id"].(string)
			v4, ok4 := config["widget_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("%s/%s/%s/%s/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{location}}/{{collection_id}}/{{engine_id}}/{{widget_config_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["collection_id"].(string)
			v2, ok2 := config["engine_id"].(string)
			v3, ok3 := config["widget_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_dns_managed_zone":
		// Format: projects/{{project}}/managedZones/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/managedZones/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_dns_managed_zone_iam":
		// No standard import format found in documentation
		return ""
	case "google_dns_policy":
		// Format: projects/{{project}}/policies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/policies/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_dns_record_set":
		// Format: projects/{{project}}/managedZones/{{zone}}/rrsets/{{name}}/{{type}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			v3, ok3 := config["type"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/managedZones/%s/rrsets/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{zone}}/{{name}}/{{type}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			v3, ok3 := config["type"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{zone}}/{{name}}/{{type}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["name"].(string)
			v2, ok2 := config["type"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_dns_response_policy":
		// Format: projects/{{project}}/responsePolicies/{{response_policy_name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["response_policy_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/responsePolicies/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{response_policy_name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["response_policy_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{response_policy_name}}
		{
			v0, ok0 := config["response_policy_name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_dns_response_policy_rule":
		// Format: projects/{{project}}/responsePolicies/{{response_policy}}/rules/{{rule_name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["response_policy"].(string)
			v2, ok2 := config["rule_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/responsePolicies/%s/rules/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{response_policy}}/{{rule_name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["response_policy"].(string)
			v2, ok2 := config["rule_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{response_policy}}/{{rule_name}}
		{
			v0, ok0 := config["response_policy"].(string)
			v1, ok1 := config["rule_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_document_ai_processor":
		// Format: projects/{{project}}/locations/{{location}}/processors/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/processors/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_document_ai_processor_default_version":
		// Format: {{processor}}
		{
			v0, ok0 := config["processor"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_document_ai_schema":
		// Format: projects/{{project}}/locations/{{location}}/schemas/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/schemas/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_document_ai_warehouse_document_schema":
		// Format: projects/{{project_number}}/locations/{{location}}/documentSchemas/{{name}}
		{
			v0, ok0 := config["project_number"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/documentSchemas/%s", v0, v1, v2)
			}
		}
		// Format: {{project_number}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project_number"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_document_ai_warehouse_location":
		// No standard import format found in documentation
		return ""
	case "google_edgecontainer_cluster":
		// Format: projects/{{project}}/locations/{{location}}/clusters/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_edgecontainer_node_pool":
		// Format: projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/nodePools/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/clusters/%s/nodePools/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{cluster}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{cluster}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["cluster"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_edgecontainer_vpn_connection":
		// Format: projects/{{project}}/locations/{{location}}/vpnConnections/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/vpnConnections/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_edgenetwork_interconnect_attachment":
		// Format: projects/{{project}}/locations/{{location}}/zones/{{zone}}/interconnectAttachment/{{interconnect_attachment_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["zone"].(string)
			v3, ok3 := config["interconnect_attachment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/zones/%s/interconnectAttachment/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{zone}}/{{interconnect_attachment_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["zone"].(string)
			v3, ok3 := config["interconnect_attachment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{zone}}/{{interconnect_attachment_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["interconnect_attachment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{interconnect_attachment_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["interconnect_attachment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_edgenetwork_network":
		// Format: projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["zone"].(string)
			v3, ok3 := config["network_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/zones/%s/networks/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{zone}}/{{network_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["zone"].(string)
			v3, ok3 := config["network_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{zone}}/{{network_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["network_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{network_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["network_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_edgenetwork_subnet":
		// Format: projects/{{project}}/locations/{{location}}/zones/{{zone}}/subnets/{{subnet_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["zone"].(string)
			v3, ok3 := config["subnet_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/zones/%s/subnets/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{zone}}/{{subnet_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["zone"].(string)
			v3, ok3 := config["subnet_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{zone}}/{{subnet_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["subnet_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{subnet_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["subnet_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_endpoints_service":
		// No standard import format found in documentation
		return ""
	case "google_endpoints_service_consumers_iam":
		// No standard import format found in documentation
		return ""
	case "google_endpoints_service_iam":
		// No standard import format found in documentation
		return ""
	case "google_essential_contacts_contact":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_eventarc_channel":
		// Format: projects/{{project}}/locations/{{location}}/channels/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/channels/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_eventarc_enrollment":
		// Format: projects/{{project}}/locations/{{location}}/enrollments/{{enrollment_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["enrollment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/enrollments/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{enrollment_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["enrollment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{enrollment_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["enrollment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_eventarc_google_api_source":
		// Format: projects/{{project}}/locations/{{location}}/googleApiSources/{{google_api_source_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["google_api_source_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/googleApiSources/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{google_api_source_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["google_api_source_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{google_api_source_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["google_api_source_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_eventarc_google_channel_config":
		// Format: projects/{{project}}/locations/{{location}}/googleChannelConfig
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/googleChannelConfig", v0, v1)
			}
		}
		// Format: {{project}}/{{location}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{location}}
		{
			v0, ok0 := config["location"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_eventarc_message_bus":
		// Format: projects/{{project}}/locations/{{location}}/messageBuses/{{message_bus_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["message_bus_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/messageBuses/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{message_bus_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["message_bus_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{message_bus_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["message_bus_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_eventarc_pipeline":
		// Format: projects/{{project}}/locations/{{location}}/pipelines/{{pipeline_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["pipeline_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/pipelines/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{pipeline_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["pipeline_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{pipeline_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["pipeline_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_eventarc_trigger":
		// Format: projects/{{project}}/locations/{{location}}/triggers/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/triggers/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_filestore_backup":
		// Format: projects/{{project}}/locations/{{location}}/backups/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/backups/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_filestore_instance":
		// Format: projects/{{project}}/locations/{{location}}/instances/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_filestore_snapshot":
		// Format: projects/{{project}}/locations/{{location}}/instances/{{instance}}/snapshots/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s/snapshots/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{instance}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{instance}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_firebase_ai_logic_config":
		// Format: projects/{{project}}/locations/{{location}}/config
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/config", v0, v1)
			}
		}
		// Format: {{project}}/{{location}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{location}}
		{
			v0, ok0 := config["location"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firebase_ai_logic_prompt_template":
		// Format: projects/{{project}}/locations/{{location}}/templates/{{template_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["template_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/templates/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{template_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["template_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{template_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["template_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_firebase_ai_logic_prompt_template_lock":
		// Format: projects/{{project}}/locations/{{location}}/templates/{{template_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["template_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/templates/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{template_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["template_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{template_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["template_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_firebase_android_app":
		// Format: {{project}} projects/{{project}}/androidApps/{{app_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["project"].(string)
			v2, ok2 := config["app_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s projects/%s/androidApps/%s", v0, v1, v2)
			}
		}
		// Format: projects/{{project}}/androidApps/{{app_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["app_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/androidApps/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{project}}/{{app_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["project"].(string)
			v2, ok2 := config["app_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: androidApps/{{app_id}}
		{
			v0, ok0 := config["app_id"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("androidApps/%s", v0)
			}
		}
		// Format: {{app_id}}
		{
			v0, ok0 := config["app_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firebase_app_check_app_attest_config":
		// Format: projects/{{project}}/apps/{{app_id}}/appAttestConfig
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["app_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/apps/%s/appAttestConfig", v0, v1)
			}
		}
		// Format: {{project}}/{{app_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["app_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{app_id}}
		{
			v0, ok0 := config["app_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firebase_app_check_debug_token":
		// Format: projects/{{project}}/apps/{{app_id}}/debugTokens/{{debug_token_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["app_id"].(string)
			v2, ok2 := config["debug_token_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/apps/%s/debugTokens/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{app_id}}/{{debug_token_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["app_id"].(string)
			v2, ok2 := config["debug_token_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{app_id}}/{{debug_token_id}}
		{
			v0, ok0 := config["app_id"].(string)
			v1, ok1 := config["debug_token_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_firebase_app_check_device_check_config":
		// Format: projects/{{project}}/apps/{{app_id}}/deviceCheckConfig
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["app_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/apps/%s/deviceCheckConfig", v0, v1)
			}
		}
		// Format: {{project}}/{{app_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["app_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{app_id}}
		{
			v0, ok0 := config["app_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firebase_app_check_play_integrity_config":
		// Format: projects/{{project}}/apps/{{app_id}}/playIntegrityConfig
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["app_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/apps/%s/playIntegrityConfig", v0, v1)
			}
		}
		// Format: {{project}}/{{app_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["app_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{app_id}}
		{
			v0, ok0 := config["app_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firebase_app_check_recaptcha_enterprise_config":
		// Format: projects/{{project}}/apps/{{app_id}}/recaptchaEnterpriseConfig
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["app_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/apps/%s/recaptchaEnterpriseConfig", v0, v1)
			}
		}
		// Format: {{project}}/{{app_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["app_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{app_id}}
		{
			v0, ok0 := config["app_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firebase_app_check_recaptcha_v3_config":
		// Format: projects/{{project}}/apps/{{app_id}}/recaptchaV3Config
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["app_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/apps/%s/recaptchaV3Config", v0, v1)
			}
		}
		// Format: {{project}}/{{app_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["app_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{app_id}}
		{
			v0, ok0 := config["app_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firebase_app_check_resource_policy":
		// Format: projects/{{project}}/services/{{service_id}}/resourcePolicies/{{resource_policy_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["service_id"].(string)
			v2, ok2 := config["resource_policy_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/services/%s/resourcePolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{service_id}}/{{resource_policy_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["service_id"].(string)
			v2, ok2 := config["resource_policy_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{service_id}}/{{resource_policy_id}}
		{
			v0, ok0 := config["service_id"].(string)
			v1, ok1 := config["resource_policy_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_firebase_app_check_service_config":
		// Format: projects/{{project}}/services/{{service_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["service_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/services/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{service_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["service_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{service_id}}
		{
			v0, ok0 := config["service_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firebase_app_hosting_backend":
		// Format: projects/{{project}}/locations/{{location}}/backends/{{backend_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["backend_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/backends/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{backend_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["backend_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{backend_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["backend_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_firebase_app_hosting_build":
		// Format: projects/{{project}}/locations/{{location}}/backends/{{backend}}/builds/{{build_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["backend"].(string)
			v3, ok3 := config["build_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/backends/%s/builds/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{backend}}/{{build_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["backend"].(string)
			v3, ok3 := config["build_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{backend}}/{{build_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["backend"].(string)
			v2, ok2 := config["build_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_firebase_app_hosting_default_domain":
		// Format: projects/{{project}}/locations/{{location}}/backends/{{backend}}/domains/{{domain_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["backend"].(string)
			v3, ok3 := config["domain_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/backends/%s/domains/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{backend}}/{{domain_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["backend"].(string)
			v3, ok3 := config["domain_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{backend}}/{{domain_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["backend"].(string)
			v2, ok2 := config["domain_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_firebase_app_hosting_domain":
		// Format: projects/{{project}}/locations/{{location}}/backends/{{backend}}/domains/{{domain_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["backend"].(string)
			v3, ok3 := config["domain_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/backends/%s/domains/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{backend}}/{{domain_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["backend"].(string)
			v3, ok3 := config["domain_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{backend}}/{{domain_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["backend"].(string)
			v2, ok2 := config["domain_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_firebase_app_hosting_traffic":
		// Format: projects/{{project}}/locations/{{location}}/backends/{{backend}}/traffic
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["backend"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/backends/%s/traffic", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{backend}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["backend"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{backend}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["backend"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_firebase_apple_app":
		// Format: {{project}} projects/{{project}}/iosApps/{{app_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["project"].(string)
			v2, ok2 := config["app_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s projects/%s/iosApps/%s", v0, v1, v2)
			}
		}
		// Format: projects/{{project}}/iosApps/{{app_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["app_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/iosApps/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{project}}/{{app_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["project"].(string)
			v2, ok2 := config["app_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: iosApps/{{app_id}}
		{
			v0, ok0 := config["app_id"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("iosApps/%s", v0)
			}
		}
		// Format: {{app_id}}
		{
			v0, ok0 := config["app_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firebase_data_connect_service":
		// Format: projects/{{project}}/locations/{{location}}/services/{{service_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["service_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/services/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{service_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["service_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{service_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["service_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_firebase_database_instance":
		// Format: projects/{{project}}/locations/{{region}}/instances/{{instance_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{instance_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{instance_id}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{instance_id}}
		{
			v0, ok0 := config["instance_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firebase_extensions_instance":
		// Format: projects/{{project}}/instances/{{instance_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/instances/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{instance_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{instance_id}}
		{
			v0, ok0 := config["instance_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firebase_hosting_channel":
		// Format: sites/{{site_id}}/channels/{{channel_id}}
		{
			v0, ok0 := config["site_id"].(string)
			v1, ok1 := config["channel_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("sites/%s/channels/%s", v0, v1)
			}
		}
		// Format: {{site_id}}/{{channel_id}}
		{
			v0, ok0 := config["site_id"].(string)
			v1, ok1 := config["channel_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_firebase_hosting_custom_domain":
		// Format: projects/{{project}}/sites/{{site_id}}/customDomains/{{custom_domain}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["site_id"].(string)
			v2, ok2 := config["custom_domain"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/sites/%s/customDomains/%s", v0, v1, v2)
			}
		}
		// Format: sites/{{site_id}}/customDomains/{{custom_domain}}
		{
			v0, ok0 := config["site_id"].(string)
			v1, ok1 := config["custom_domain"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("sites/%s/customDomains/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{site_id}}/{{custom_domain}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["site_id"].(string)
			v2, ok2 := config["custom_domain"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{site_id}}/{{custom_domain}}
		{
			v0, ok0 := config["site_id"].(string)
			v1, ok1 := config["custom_domain"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_firebase_hosting_release":
		// Format: sites/{{site_id}}/channels/{{channel_id}}/releases/{{release_id}}
		{
			v0, ok0 := config["site_id"].(string)
			v1, ok1 := config["channel_id"].(string)
			v2, ok2 := config["release_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("sites/%s/channels/%s/releases/%s", v0, v1, v2)
			}
		}
		// Format: sites/{{site_id}}/releases/{{release_id}}
		{
			v0, ok0 := config["site_id"].(string)
			v1, ok1 := config["release_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("sites/%s/releases/%s", v0, v1)
			}
		}
		// Format: {{site_id}}/{{channel_id}}/{{release_id}}
		{
			v0, ok0 := config["site_id"].(string)
			v1, ok1 := config["channel_id"].(string)
			v2, ok2 := config["release_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{site_id}}/{{release_id}}
		{
			v0, ok0 := config["site_id"].(string)
			v1, ok1 := config["release_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_firebase_hosting_site":
		// Format: projects/{{project}}/sites/{{site_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["site_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/sites/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{site_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["site_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: sites/{{site_id}}
		{
			v0, ok0 := config["site_id"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("sites/%s", v0)
			}
		}
		// Format: {{site_id}}
		{
			v0, ok0 := config["site_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firebase_hosting_version":
		// Format: sites/{{site_id}}/versions/{{version_id}}
		{
			v0, ok0 := config["site_id"].(string)
			v1, ok1 := config["version_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("sites/%s/versions/%s", v0, v1)
			}
		}
		// Format: {{site_id}}/{{version_id}}
		{
			v0, ok0 := config["site_id"].(string)
			v1, ok1 := config["version_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_firebase_project":
		// Format: projects/{{project}}
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("projects/%s", v0)
			}
		}
		// Format: {{project}}
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firebase_remote_config_remote_config":
		// Format: projects/{{project}}/remoteConfig
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("projects/%s/remoteConfig", v0)
			}
		}
		// Format: {{project}}
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firebase_storage_bucket":
		// Format: projects/{{project}}/buckets/{{bucket_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["bucket_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/buckets/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{bucket_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["bucket_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{bucket_id}}
		{
			v0, ok0 := config["bucket_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firebase_storage_default_bucket":
		// Format: projects/{{project}}/defaultBucket
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("projects/%s/defaultBucket", v0)
			}
		}
		// Format: {{project}}
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firebase_web_app":
		// Format: {{project}} projects/{{project}}/webApps/{{app_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["project"].(string)
			v2, ok2 := config["app_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s projects/%s/webApps/%s", v0, v1, v2)
			}
		}
		// Format: projects/{{project}}/webApps/{{app_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["app_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/webApps/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{project}}/{{app_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["project"].(string)
			v2, ok2 := config["app_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: webApps/{{app_id}}
		{
			v0, ok0 := config["app_id"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("webApps/%s", v0)
			}
		}
		// Format: {{app_id}}
		{
			v0, ok0 := config["app_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firebaserules_release":
		// Format: projects/{{project}}/releases/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/releases/%s", v0, v1)
			}
		}
		return ""
	case "google_firebaserules_ruleset":
		// Format: projects/{{project}}/rulesets/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/rulesets/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firestore_backup_schedule":
		// Format: projects/{{project}}/databases/{{database}}/backupSchedules/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["database"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/databases/%s/backupSchedules/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{database}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["database"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{database}}/{{name}}
		{
			v0, ok0 := config["database"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_firestore_database":
		// Format: projects/{{project}}/databases/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/databases/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firestore_document":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firestore_field":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firestore_index":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_firestore_user_creds":
		// Format: projects/{{project}}/databases/{{database}}/userCreds/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["database"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/databases/%s/userCreds/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{database}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["database"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{database}}/{{name}}
		{
			v0, ok0 := config["database"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_folder_access_approval_settings":
		// Format: folders/{{folder_id}}/accessApprovalSettings
		{
			v0, ok0 := config["folder_id"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("folders/%s/accessApprovalSettings", v0)
			}
		}
		// Format: {{folder_id}}
		{
			v0, ok0 := config["folder_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_folder_service_identity":
		// No standard import format found in documentation
		return ""
	case "google_gemini_code_repository_index":
		// Format: projects/{{project}}/locations/{{location}}/codeRepositoryIndexes/{{code_repository_index_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["code_repository_index_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/codeRepositoryIndexes/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{code_repository_index_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["code_repository_index_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{code_repository_index_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["code_repository_index_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_gemini_code_tools_setting":
		// Format: projects/{{project}}/locations/{{location}}/codeToolsSettings/{{code_tools_setting_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["code_tools_setting_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/codeToolsSettings/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{code_tools_setting_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["code_tools_setting_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{code_tools_setting_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["code_tools_setting_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_gemini_code_tools_setting_binding":
		// Format: projects/{{project}}/locations/{{location}}/codeToolsSettings/{{code_tools_setting_id}}/settingBindings/{{setting_binding_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["code_tools_setting_id"].(string)
			v3, ok3 := config["setting_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/codeToolsSettings/%s/settingBindings/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{code_tools_setting_id}}/{{setting_binding_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["code_tools_setting_id"].(string)
			v3, ok3 := config["setting_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{code_tools_setting_id}}/{{setting_binding_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["code_tools_setting_id"].(string)
			v2, ok2 := config["setting_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_gemini_data_sharing_with_google_setting":
		// Format: projects/{{project}}/locations/{{location}}/dataSharingWithGoogleSettings/{{data_sharing_with_google_setting_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_sharing_with_google_setting_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/dataSharingWithGoogleSettings/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{data_sharing_with_google_setting_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_sharing_with_google_setting_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{data_sharing_with_google_setting_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["data_sharing_with_google_setting_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_gemini_data_sharing_with_google_setting_binding":
		// Format: projects/{{project}}/locations/{{location}}/dataSharingWithGoogleSettings/{{data_sharing_with_google_setting_id}}/settingBindings/{{setting_binding_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_sharing_with_google_setting_id"].(string)
			v3, ok3 := config["setting_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/dataSharingWithGoogleSettings/%s/settingBindings/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{data_sharing_with_google_setting_id}}/{{setting_binding_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["data_sharing_with_google_setting_id"].(string)
			v3, ok3 := config["setting_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{data_sharing_with_google_setting_id}}/{{setting_binding_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["data_sharing_with_google_setting_id"].(string)
			v2, ok2 := config["setting_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_gemini_gemini_gcp_enablement_setting":
		// Format: projects/{{project}}/locations/{{location}}/geminiGcpEnablementSettings/{{gemini_gcp_enablement_setting_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["gemini_gcp_enablement_setting_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/geminiGcpEnablementSettings/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{gemini_gcp_enablement_setting_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["gemini_gcp_enablement_setting_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{gemini_gcp_enablement_setting_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["gemini_gcp_enablement_setting_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_gemini_gemini_gcp_enablement_setting_binding":
		// Format: projects/{{project}}/locations/{{location}}/geminiGcpEnablementSettings/{{gemini_gcp_enablement_setting_id}}/settingBindings/{{setting_binding_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["gemini_gcp_enablement_setting_id"].(string)
			v3, ok3 := config["setting_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/geminiGcpEnablementSettings/%s/settingBindings/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{gemini_gcp_enablement_setting_id}}/{{setting_binding_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["gemini_gcp_enablement_setting_id"].(string)
			v3, ok3 := config["setting_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{gemini_gcp_enablement_setting_id}}/{{setting_binding_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["gemini_gcp_enablement_setting_id"].(string)
			v2, ok2 := config["setting_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_gemini_logging_setting":
		// Format: projects/{{project}}/locations/{{location}}/loggingSettings/{{logging_setting_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["logging_setting_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/loggingSettings/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{logging_setting_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["logging_setting_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{logging_setting_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["logging_setting_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_gemini_logging_setting_binding":
		// Format: projects/{{project}}/locations/{{location}}/loggingSettings/{{logging_setting_id}}/settingBindings/{{setting_binding_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["logging_setting_id"].(string)
			v3, ok3 := config["setting_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/loggingSettings/%s/settingBindings/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{logging_setting_id}}/{{setting_binding_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["logging_setting_id"].(string)
			v3, ok3 := config["setting_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{logging_setting_id}}/{{setting_binding_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["logging_setting_id"].(string)
			v2, ok2 := config["setting_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_gemini_release_channel_setting":
		// Format: projects/{{project}}/locations/{{location}}/releaseChannelSettings/{{release_channel_setting_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["release_channel_setting_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/releaseChannelSettings/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{release_channel_setting_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["release_channel_setting_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{release_channel_setting_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["release_channel_setting_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_gemini_release_channel_setting_binding":
		// Format: projects/{{project}}/locations/{{location}}/releaseChannelSettings/{{release_channel_setting_id}}/settingBindings/{{setting_binding_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["release_channel_setting_id"].(string)
			v3, ok3 := config["setting_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/releaseChannelSettings/%s/settingBindings/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{release_channel_setting_id}}/{{setting_binding_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["release_channel_setting_id"].(string)
			v3, ok3 := config["setting_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{release_channel_setting_id}}/{{setting_binding_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["release_channel_setting_id"].(string)
			v2, ok2 := config["setting_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_gemini_repository_group":
		// Format: projects/{{project}}/locations/{{location}}/codeRepositoryIndexes/{{code_repository_index}}/repositoryGroups/{{repository_group_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["code_repository_index"].(string)
			v3, ok3 := config["repository_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/codeRepositoryIndexes/%s/repositoryGroups/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{code_repository_index}}/{{repository_group_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["code_repository_index"].(string)
			v3, ok3 := config["repository_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{code_repository_index}}/{{repository_group_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["code_repository_index"].(string)
			v2, ok2 := config["repository_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_gemini_repository_group_iam":
		// No standard import format found in documentation
		return ""
	case "google_gke_backup_backup_channel":
		// Format: projects/{{project}}/locations/{{location}}/backupChannels/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/backupChannels/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_gke_backup_backup_plan":
		// Format: projects/{{project}}/locations/{{location}}/backupPlans/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_gke_backup_backup_plan_iam":
		// No standard import format found in documentation
		return ""
	case "google_gke_backup_restore_channel":
		// Format: projects/{{project}}/locations/{{location}}/restoreChannels/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/restoreChannels/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_gke_backup_restore_plan":
		// Format: projects/{{project}}/locations/{{location}}/restorePlans/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_gke_backup_restore_plan_iam":
		// No standard import format found in documentation
		return ""
	case "google_gke_hub_feature":
		// Format: projects/{{project}}/locations/{{location}}/features/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/features/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_gke_hub_feature_iam":
		// No standard import format found in documentation
		return ""
	case "google_gke_hub_feature_membership":
		// Format: projects/{{project}}/locations/{{location}}/features/{{feature}}/membershipId/{{membership}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["feature"].(string)
			v3, ok3 := config["membership"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/features/%s/membershipId/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{feature}}/{{membership}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["feature"].(string)
			v3, ok3 := config["membership"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{feature}}/{{membership}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["feature"].(string)
			v2, ok2 := config["membership"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_gke_hub_fleet":
		// Format: projects/{{project}}/locations/global/fleets/default
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("projects/%s/locations/global/fleets/default", v0)
			}
		}
		// Format: {{project}}
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_gke_hub_membership":
		// Format: projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["membership_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/memberships/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{membership_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["membership_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{membership_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["membership_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_gke_hub_membership_binding":
		// Format: projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}/bindings/{{membership_binding_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["membership_id"].(string)
			v3, ok3 := config["membership_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/memberships/%s/bindings/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{membership_id}}/{{membership_binding_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["membership_id"].(string)
			v3, ok3 := config["membership_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{membership_id}}/{{membership_binding_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["membership_id"].(string)
			v2, ok2 := config["membership_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_gke_hub_membership_iam":
		// No standard import format found in documentation
		return ""
	case "google_gke_hub_membership_rbac_role_binding":
		// Format: projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}/rbacrolebindings/{{membership_rbac_role_binding_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["membership_id"].(string)
			v3, ok3 := config["membership_rbac_role_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/memberships/%s/rbacrolebindings/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{membership_id}}/{{membership_rbac_role_binding_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["membership_id"].(string)
			v3, ok3 := config["membership_rbac_role_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{membership_id}}/{{membership_rbac_role_binding_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["membership_id"].(string)
			v2, ok2 := config["membership_rbac_role_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_gke_hub_namespace":
		// Format: projects/{{project}}/locations/global/scopes/{{scope_id}}/namespaces/{{scope_namespace_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["scope_id"].(string)
			v2, ok2 := config["scope_namespace_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/global/scopes/%s/namespaces/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{scope_id}}/{{scope_namespace_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["scope_id"].(string)
			v2, ok2 := config["scope_namespace_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{scope_id}}/{{scope_namespace_id}}
		{
			v0, ok0 := config["scope_id"].(string)
			v1, ok1 := config["scope_namespace_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_gke_hub_rollout_sequence":
		// Format: projects/{{project}}/locations/global/rolloutSequences/{{rollout_sequence_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["rollout_sequence_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/rolloutSequences/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{rollout_sequence_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["rollout_sequence_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{rollout_sequence_id}}
		{
			v0, ok0 := config["rollout_sequence_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_gke_hub_scope":
		// Format: projects/{{project}}/locations/global/scopes/{{scope_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["scope_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/scopes/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{scope_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["scope_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{scope_id}}
		{
			v0, ok0 := config["scope_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_gke_hub_scope_iam":
		// No standard import format found in documentation
		return ""
	case "google_gke_hub_scope_rbac_role_binding":
		// Format: projects/{{project}}/locations/global/scopes/{{scope_id}}/rbacrolebindings/{{scope_rbac_role_binding_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["scope_id"].(string)
			v2, ok2 := config["scope_rbac_role_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/global/scopes/%s/rbacrolebindings/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{scope_id}}/{{scope_rbac_role_binding_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["scope_id"].(string)
			v2, ok2 := config["scope_rbac_role_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{scope_id}}/{{scope_rbac_role_binding_id}}
		{
			v0, ok0 := config["scope_id"].(string)
			v1, ok1 := config["scope_rbac_role_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_gkeonprem_bare_metal_admin_cluster":
		// Format: projects/{{project}}/locations/{{location}}/bareMetalAdminClusters/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/bareMetalAdminClusters/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_gkeonprem_bare_metal_cluster":
		// Format: projects/{{project}}/locations/{{location}}/bareMetalClusters/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/bareMetalClusters/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_gkeonprem_bare_metal_node_pool":
		// Format: projects/{{project}}/locations/{{location}}/bareMetalClusters/{{bare_metal_cluster}}/bareMetalNodePools/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["bare_metal_cluster"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/bareMetalClusters/%s/bareMetalNodePools/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{bare_metal_cluster}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["bare_metal_cluster"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{bare_metal_cluster}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["bare_metal_cluster"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_gkeonprem_vmware_admin_cluster":
		// Format: projects/{{project}}/locations/{{location}}/vmwareAdminClusters/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/vmwareAdminClusters/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_gkeonprem_vmware_cluster":
		// Format: projects/{{project}}/locations/{{location}}/vmwareClusters/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/vmwareClusters/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_gkeonprem_vmware_node_pool":
		// Format: projects/{{project}}/locations/{{location}}/vmwareClusters/{{vmware_cluster}}/vmwareNodePools/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["vmware_cluster"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/vmwareClusters/%s/vmwareNodePools/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{vmware_cluster}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["vmware_cluster"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{vmware_cluster}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["vmware_cluster"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_google_billing_subaccount":
		// Format: billingAccounts/{billing_account_id}
		{
			v0, ok0 := config["billing_account_id"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("billingAccounts/%s", v0)
			}
		}
		return ""
	case "google_google_folder":
		// Format: folders/{{folder_id}}
		{
			v0, ok0 := config["folder_id"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("folders/%s", v0)
			}
		}
		// Format: {{folder_id}}
		{
			v0, ok0 := config["folder_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_google_folder_iam":
		// Format: "folders/{{folder_id}} roles/viewer user:foo@example.com"
		{
			v0, ok0 := config["folder_id"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("\"folders/%s roles/viewer user:foo@example.com\"", v0)
			}
		}
		return ""
	case "google_google_folder_organization_policy":
		// Format: folders/{{folder_id}}/constraints/serviceuser.services
		{
			v0, ok0 := config["folder_id"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("folders/%s/constraints/serviceuser.services", v0)
			}
		}
		// Format: {{folder_id}}/serviceuser.services
		{
			v0, ok0 := config["folder_id"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("%s/serviceuser.services", v0)
			}
		}
		return ""
	case "google_google_kms_crypto_key_iam":
		// Format: "{{project_id}}/{{location}}/{{key_ring_name}}/{{crypto_key_name}} roles/viewer user:foo@example.com"
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["key_ring_name"].(string)
			v3, ok3 := config["crypto_key_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("\"%s/%s/%s/%s roles/viewer user:foo@example.com\"", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_google_kms_key_ring_iam":
		// Format: "{{project_id}}/{{location}}/{{key_ring_name}} roles/viewer user:foo@example.com"
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["key_ring_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("\"%s/%s/%s roles/viewer user:foo@example.com\"", v0, v1, v2)
			}
		}
		return ""
	case "google_google_organization_iam":
		// Format: "{{org_id}} roles/viewer user:foo@example.com"
		{
			v0, ok0 := config["org_id"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("\"%s roles/viewer user:foo@example.com\"", v0)
			}
		}
		return ""
	case "google_google_organization_iam_custom_role":
		// No standard import format found in documentation
		return ""
	case "google_google_organization_policy":
		// Format: {{org_id}}/constraints/{{constraint}}
		{
			v0, ok0 := config["org_id"].(string)
			v1, ok1 := config["constraint"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/constraints/%s", v0, v1)
			}
		}
		return ""
	case "google_google_project":
		// Format: {{project_id}}
		{
			v0, ok0 := config["project_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_google_project_default_service_accounts":
		// No standard import format found in documentation
		return ""
	case "google_google_project_iam":
		// Format: "{{project_id}} roles/viewer user:foo@example.com"
		{
			v0, ok0 := config["project_id"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("\"%s roles/viewer user:foo@example.com\"", v0)
			}
		}
		return ""
	case "google_google_project_iam_custom_role":
		// Format: projects/{{project}}/roles/{{role_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["role_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/roles/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{role_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["role_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{role_id}}
		{
			v0, ok0 := config["role_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_google_project_iam_member_remove":
		// No standard import format found in documentation
		return ""
	case "google_google_project_organization_policy":
		// Format: projects/{{project_id}}:constraints/{{constraint}}
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["constraint"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s:constraints/%s", v0, v1)
			}
		}
		// Format: {{project_id}}:constraints/{{constraint}}
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["constraint"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s:constraints/%s", v0, v1)
			}
		}
		// Format: {{project_id}}:{{constraint}}
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["constraint"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s:%s", v0, v1)
			}
		}
		return ""
	case "google_google_project_service":
		// Format: {{project_id}}/{{service}}
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["service"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_google_service_account":
		// Format: projects/{{project_id}}/serviceAccounts/{{email}}
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["email"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/serviceAccounts/%s", v0, v1)
			}
		}
		return ""
	case "google_google_service_account_iam":
		// Format: "projects/{{project_id}}/serviceAccounts/{{service_account_email}} roles/editor user:foo@example.com"
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["service_account_email"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("\"projects/%s/serviceAccounts/%s roles/editor user:foo@example.com\"", v0, v1)
			}
		}
		return ""
	case "google_google_service_account_key":
		// No standard import format found in documentation
		return ""
	case "google_google_service_networking_peered_dns_domain":
		// Format: services/{service}/projects/{project}/global/networks/{network}/peeredDnsDomains/{name}
		{
			v0, ok0 := config["service"].(string)
			v1, ok1 := config["project"].(string)
			v2, ok2 := config["network"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("services/%s/projects/%s/global/networks/%s/peeredDnsDomains/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_google_tags_location_tag_binding":
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_healthcare_consent_store":
		// Format: {{dataset}}/consentStores/{{name}}
		{
			v0, ok0 := config["dataset"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/consentStores/%s", v0, v1)
			}
		}
		return ""
	case "google_healthcare_consent_store_iam":
		// No standard import format found in documentation
		return ""
	case "google_healthcare_dataset":
		// Format: projects/{{project}}/locations/{{location}}/datasets/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/datasets/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_healthcare_dataset_iam":
		// Format: "{{project_id}}/{{location}}/{{dataset}} roles/editor jane@example.com"
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["dataset"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("\"%s/%s/%s roles/editor jane@example.com\"", v0, v1, v2)
			}
		}
		return ""
	case "google_healthcare_dicom_store":
		// Format: {{dataset}}/dicomStores/{{name}}
		{
			v0, ok0 := config["dataset"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/dicomStores/%s", v0, v1)
			}
		}
		// Format: {{dataset}}/{{name}}
		{
			v0, ok0 := config["dataset"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_healthcare_dicom_store_iam":
		// Format: "{{project_id}}/{{location}}/{{dataset}}/{{dicom_store}} roles/editor jane@example.com"
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["dataset"].(string)
			v3, ok3 := config["dicom_store"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("\"%s/%s/%s/%s roles/editor jane@example.com\"", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_healthcare_fhir_store":
		// Format: {{dataset}}/fhirStores/{{name}}
		{
			v0, ok0 := config["dataset"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/fhirStores/%s", v0, v1)
			}
		}
		// Format: {{dataset}}/{{name}}
		{
			v0, ok0 := config["dataset"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_healthcare_fhir_store_iam":
		// Format: "{{project_id}}/{{location}}/{{dataset}}/{{fhir_store}} roles/editor jane@example.com"
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["dataset"].(string)
			v3, ok3 := config["fhir_store"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("\"%s/%s/%s/%s roles/editor jane@example.com\"", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_healthcare_hl7_v2_store":
		// Format: {{dataset}}/hl7V2Stores/{{name}}
		{
			v0, ok0 := config["dataset"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/hl7V2Stores/%s", v0, v1)
			}
		}
		// Format: {{dataset}}/{{name}}
		{
			v0, ok0 := config["dataset"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_healthcare_hl7_v2_store_iam":
		// Format: "{{project_id}}/{{location}}/{{dataset}}/{{hl7_v2_store}} roles/editor jane@example.com"
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["dataset"].(string)
			v3, ok3 := config["hl7_v2_store"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("\"%s/%s/%s/%s roles/editor jane@example.com\"", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_healthcare_pipeline_job":
		// Format: {{dataset}}/pipelineJobs/{{name}}
		{
			v0, ok0 := config["dataset"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/pipelineJobs/%s", v0, v1)
			}
		}
		// Format: {{dataset}}/pipelineJobs?pipelineJobId={{name}}
		{
			v0, ok0 := config["dataset"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/pipelineJobs?pipelineJobId=%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_healthcare_workspace":
		// Format: {{dataset}}/dataMapperWorkspaces/{{name}}
		{
			v0, ok0 := config["dataset"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/dataMapperWorkspaces/%s", v0, v1)
			}
		}
		return ""
	case "google_hypercomputecluster_cluster":
		// Format: projects/{{project}}/locations/{{location}}/clusters/{{cluster_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{cluster_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{cluster_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_iam_access_boundary_policy":
		// Format: {{parent}}/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_iam_deny_policy":
		// Format: {{parent}}/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_iam_folders_policy_binding":
		// Format: folders/{{folder}}/locations/{{location}}/policyBindings/{{policy_binding_id}}
		{
			v0, ok0 := config["folder"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["policy_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("folders/%s/locations/%s/policyBindings/%s", v0, v1, v2)
			}
		}
		// Format: {{folder}}/{{location}}/{{policy_binding_id}}
		{
			v0, ok0 := config["folder"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["policy_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_iam_oauth_client":
		// Format: projects/{{project}}/locations/{{location}}/oauthClients/{{oauth_client_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["oauth_client_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/oauthClients/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{oauth_client_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["oauth_client_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{oauth_client_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["oauth_client_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_iam_oauth_client_credential":
		// Format: projects/{{project}}/locations/{{location}}/oauthClients/{{oauthclient}}/credentials/{{oauth_client_credential_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["oauthclient"].(string)
			v3, ok3 := config["oauth_client_credential_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/oauthClients/%s/credentials/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{oauthclient}}/{{oauth_client_credential_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["oauthclient"].(string)
			v3, ok3 := config["oauth_client_credential_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{oauthclient}}/{{oauth_client_credential_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["oauthclient"].(string)
			v2, ok2 := config["oauth_client_credential_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_iam_organizations_policy_binding":
		// Format: organizations/{{organization}}/locations/{{location}}/policyBindings/{{policy_binding_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["policy_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("organizations/%s/locations/%s/policyBindings/%s", v0, v1, v2)
			}
		}
		// Format: {{organization}}/{{location}}/{{policy_binding_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["policy_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_iam_principal_access_boundary_policy":
		// Format: organizations/{{organization}}/locations/{{location}}/principalAccessBoundaryPolicies/{{principal_access_boundary_policy_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["principal_access_boundary_policy_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("organizations/%s/locations/%s/principalAccessBoundaryPolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{organization}}/{{location}}/{{principal_access_boundary_policy_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["principal_access_boundary_policy_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_iam_projects_policy_binding":
		// Format: projects/{{project}}/locations/{{location}}/policyBindings/{{policy_binding_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["policy_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/policyBindings/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{policy_binding_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["policy_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{policy_binding_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["policy_binding_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_iam_workforce_pool":
		// Format: locations/{{location}}/workforcePools/{{workforce_pool_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["workforce_pool_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("locations/%s/workforcePools/%s", v0, v1)
			}
		}
		// Format: {{location}}/{{workforce_pool_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["workforce_pool_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_iam_workforce_pool_iam":
		// No standard import format found in documentation
		return ""
	case "google_iam_workforce_pool_provider":
		// Format: locations/{{location}}/workforcePools/{{workforce_pool_id}}/providers/{{provider_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["workforce_pool_id"].(string)
			v2, ok2 := config["provider_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("locations/%s/workforcePools/%s/providers/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{workforce_pool_id}}/{{provider_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["workforce_pool_id"].(string)
			v2, ok2 := config["provider_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_iam_workforce_pool_provider_key":
		// Format: locations/{{location}}/workforcePools/{{workforce_pool_id}}/providers/{{provider_id}}/keys/{{key_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["workforce_pool_id"].(string)
			v2, ok2 := config["provider_id"].(string)
			v3, ok3 := config["key_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("locations/%s/workforcePools/%s/providers/%s/keys/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{workforce_pool_id}}/{{provider_id}}/{{key_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["workforce_pool_id"].(string)
			v2, ok2 := config["provider_id"].(string)
			v3, ok3 := config["key_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_iam_workforce_pool_provider_scim_tenant":
		// Format: locations/{{location}}/workforcePools/{{workforce_pool_id}}/providers/{{provider_id}}/scimTenants/{{scim_tenant_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["workforce_pool_id"].(string)
			v2, ok2 := config["provider_id"].(string)
			v3, ok3 := config["scim_tenant_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("locations/%s/workforcePools/%s/providers/%s/scimTenants/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{workforce_pool_id}}/{{provider_id}}/{{scim_tenant_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["workforce_pool_id"].(string)
			v2, ok2 := config["provider_id"].(string)
			v3, ok3 := config["scim_tenant_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_iam_workforce_pool_provider_scim_token":
		// Format: locations/{{location}}/workforcePools/{{workforce_pool_id}}/providers/{{provider_id}}/scimTenants/{{scim_tenant_id}}/tokens/{{scim_token_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["workforce_pool_id"].(string)
			v2, ok2 := config["provider_id"].(string)
			v3, ok3 := config["scim_tenant_id"].(string)
			v4, ok4 := config["scim_token_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("locations/%s/workforcePools/%s/providers/%s/scimTenants/%s/tokens/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{location}}/{{workforce_pool_id}}/{{provider_id}}/{{scim_tenant_id}}/{{scim_token_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["workforce_pool_id"].(string)
			v2, ok2 := config["provider_id"].(string)
			v3, ok3 := config["scim_tenant_id"].(string)
			v4, ok4 := config["scim_token_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("%s/%s/%s/%s/%s", v0, v1, v2, v3, v4)
			}
		}
		return ""
	case "google_iam_workload_identity_pool":
		// Format: projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["workload_identity_pool_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/workloadIdentityPools/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{workload_identity_pool_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["workload_identity_pool_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{workload_identity_pool_id}}
		{
			v0, ok0 := config["workload_identity_pool_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_iam_workload_identity_pool_iam":
		// No standard import format found in documentation
		return ""
	case "google_iam_workload_identity_pool_managed_identity":
		// Format: projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}/namespaces/{{workload_identity_pool_namespace_id}}/managedIdentities/{{workload_identity_pool_managed_identity_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["workload_identity_pool_id"].(string)
			v2, ok2 := config["workload_identity_pool_namespace_id"].(string)
			v3, ok3 := config["workload_identity_pool_managed_identity_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/global/workloadIdentityPools/%s/namespaces/%s/managedIdentities/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{workload_identity_pool_id}}/{{workload_identity_pool_namespace_id}}/{{workload_identity_pool_managed_identity_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["workload_identity_pool_id"].(string)
			v2, ok2 := config["workload_identity_pool_namespace_id"].(string)
			v3, ok3 := config["workload_identity_pool_managed_identity_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{workload_identity_pool_id}}/{{workload_identity_pool_namespace_id}}/{{workload_identity_pool_managed_identity_id}}
		{
			v0, ok0 := config["workload_identity_pool_id"].(string)
			v1, ok1 := config["workload_identity_pool_namespace_id"].(string)
			v2, ok2 := config["workload_identity_pool_managed_identity_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_iam_workload_identity_pool_namespace":
		// Format: projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}/namespaces/{{workload_identity_pool_namespace_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["workload_identity_pool_id"].(string)
			v2, ok2 := config["workload_identity_pool_namespace_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/global/workloadIdentityPools/%s/namespaces/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{workload_identity_pool_id}}/{{workload_identity_pool_namespace_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["workload_identity_pool_id"].(string)
			v2, ok2 := config["workload_identity_pool_namespace_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{workload_identity_pool_id}}/{{workload_identity_pool_namespace_id}}
		{
			v0, ok0 := config["workload_identity_pool_id"].(string)
			v1, ok1 := config["workload_identity_pool_namespace_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_iam_workload_identity_pool_provider":
		// Format: projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}/providers/{{workload_identity_pool_provider_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["workload_identity_pool_id"].(string)
			v2, ok2 := config["workload_identity_pool_provider_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/global/workloadIdentityPools/%s/providers/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{workload_identity_pool_id}}/{{workload_identity_pool_provider_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["workload_identity_pool_id"].(string)
			v2, ok2 := config["workload_identity_pool_provider_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{workload_identity_pool_id}}/{{workload_identity_pool_provider_id}}
		{
			v0, ok0 := config["workload_identity_pool_id"].(string)
			v1, ok1 := config["workload_identity_pool_provider_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_iap_agent_registry_iam":
		// No standard import format found in documentation
		return ""
	case "google_iap_app_engine_service_iam":
		// No standard import format found in documentation
		return ""
	case "google_iap_app_engine_version_iam":
		// No standard import format found in documentation
		return ""
	case "google_iap_brand":
		// Format: projects/{{project_id}}/brands/{{brand_id}}
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["brand_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/brands/%s", v0, v1)
			}
		}
		// Format: projects/{{project_number}}/brands/{{brand_id}}
		{
			v0, ok0 := config["project_number"].(string)
			v1, ok1 := config["brand_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/brands/%s", v0, v1)
			}
		}
		// Format: {{project_number}}/{{brand_id}}
		{
			v0, ok0 := config["project_number"].(string)
			v1, ok1 := config["brand_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_iap_client":
		// Format: {{brand}}/identityAwareProxyClients/{{client_id}}
		{
			v0, ok0 := config["brand"].(string)
			v1, ok1 := config["client_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/identityAwareProxyClients/%s", v0, v1)
			}
		}
		// Format: {{brand}}/{{client_id}}
		{
			v0, ok0 := config["brand"].(string)
			v1, ok1 := config["client_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_iap_location_web_iam":
		// No standard import format found in documentation
		return ""
	case "google_iap_settings":
		// Format: {{name}}/iapSettings
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("%s/iapSettings", v0)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_iap_tunnel_dest_group":
		// Format: projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["group_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/iap_tunnel/locations/%s/destGroups/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["group_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/iap_tunnel/locations/%s/destGroups/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{group_name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["group_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/destGroups/{{group_name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["group_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/destGroups/%s", v0, v1)
			}
		}
		// Format: {{region}}/{{group_name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["group_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{group_name}}
		{
			v0, ok0 := config["group_name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_iap_tunnel_dest_group_iam":
		// No standard import format found in documentation
		return ""
	case "google_iap_tunnel_iam":
		// No standard import format found in documentation
		return ""
	case "google_iap_tunnel_instance_iam":
		// No standard import format found in documentation
		return ""
	case "google_iap_web_backend_service_iam":
		// No standard import format found in documentation
		return ""
	case "google_iap_web_cloud_run_service_iam":
		// No standard import format found in documentation
		return ""
	case "google_iap_web_forwarding_rule_service_iam":
		// No standard import format found in documentation
		return ""
	case "google_iap_web_iam":
		// No standard import format found in documentation
		return ""
	case "google_iap_web_region_backend_service_iam":
		// No standard import format found in documentation
		return ""
	case "google_iap_web_region_forwarding_rule_service_iam":
		// No standard import format found in documentation
		return ""
	case "google_iap_web_type_app_engine_iam":
		// No standard import format found in documentation
		return ""
	case "google_iap_web_type_compute_iam":
		// No standard import format found in documentation
		return ""
	case "google_identity_platform_config":
		// Format: projects/{{project}}/config
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("projects/%s/config", v0)
			}
		}
		// Format: projects/{{project}}
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("projects/%s", v0)
			}
		}
		// Format: {{project}}
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_identity_platform_default_supported_idp_config":
		// Format: projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["idp_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/defaultSupportedIdpConfigs/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{idp_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["idp_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{idp_id}}
		{
			v0, ok0 := config["idp_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_identity_platform_inbound_saml_config":
		// Format: projects/{{project}}/inboundSamlConfigs/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/inboundSamlConfigs/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_identity_platform_oauth_idp_config":
		// Format: projects/{{project}}/oauthIdpConfigs/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/oauthIdpConfigs/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_identity_platform_tenant":
		// Format: projects/{{project}}/tenants/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/tenants/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_identity_platform_tenant_default_supported_idp_config":
		// Format: projects/{{project}}/tenants/{{tenant}}/defaultSupportedIdpConfigs/{{idp_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["tenant"].(string)
			v2, ok2 := config["idp_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/tenants/%s/defaultSupportedIdpConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{tenant}}/{{idp_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["tenant"].(string)
			v2, ok2 := config["idp_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{tenant}}/{{idp_id}}
		{
			v0, ok0 := config["tenant"].(string)
			v1, ok1 := config["idp_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_identity_platform_tenant_inbound_saml_config":
		// Format: projects/{{project}}/tenants/{{tenant}}/inboundSamlConfigs/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["tenant"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/tenants/%s/inboundSamlConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{tenant}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["tenant"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{tenant}}/{{name}}
		{
			v0, ok0 := config["tenant"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_identity_platform_tenant_oauth_idp_config":
		// Format: projects/{{project}}/tenants/{{tenant}}/oauthIdpConfigs/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["tenant"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/tenants/%s/oauthIdpConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{tenant}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["tenant"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{tenant}}/{{name}}
		{
			v0, ok0 := config["tenant"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_integration_connectors_connection":
		// Format: projects/{{project}}/locations/{{location}}/connections/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/connections/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_integration_connectors_endpoint_attachment":
		// Format: projects/{{project}}/locations/{{location}}/endpointAttachments/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/endpointAttachments/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_integration_connectors_managed_zone":
		// Format: projects/{{project}}/locations/global/managedZones/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/managedZones/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_integrations_auth_config":
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{project}} {{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s %s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_integrations_client":
		// Format: projects/{{project}}/locations/{{location}}/clients
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/clients", v0, v1)
			}
		}
		// Format: {{project}}/{{location}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{location}}
		{
			v0, ok0 := config["location"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_kms_autokey_config":
		// Format: folders/{{folder}}/autokeyConfig
		{
			v0, ok0 := config["folder"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("folders/%s/autokeyConfig", v0)
			}
		}
		// Format: {{folder}}
		{
			v0, ok0 := config["folder"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_kms_crypto_key":
		// Format: {{key_ring}}/cryptoKeys/{{name}}
		{
			v0, ok0 := config["key_ring"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/cryptoKeys/%s", v0, v1)
			}
		}
		// Format: {{key_ring}}/{{name}}
		{
			v0, ok0 := config["key_ring"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_kms_crypto_key_version":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_kms_ekm_connection":
		// Format: projects/{{project}}/locations/{{location}}/ekmConnections/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/ekmConnections/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_kms_ekm_connection_iam":
		// No standard import format found in documentation
		return ""
	case "google_kms_folder_kaj_policy_config":
		// Format: folders/{{folder}}/kajPolicyConfig
		{
			v0, ok0 := config["folder"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("folders/%s/kajPolicyConfig", v0)
			}
		}
		// Format: {{folder}}
		{
			v0, ok0 := config["folder"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_kms_key_handle":
		// Format: projects/{{project}}/locations/{{location}}/keyHandles/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/keyHandles/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_kms_key_ring":
		// Format: projects/{{project}}/locations/{{location}}/keyRings/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_kms_key_ring_import_job":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_kms_organization_kaj_policy_config":
		// Format: organizations/{{organization}}/kajPolicyConfig
		{
			v0, ok0 := config["organization"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("organizations/%s/kajPolicyConfig", v0)
			}
		}
		// Format: {{organization}}
		{
			v0, ok0 := config["organization"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_kms_project_autokey_config":
		// Format: projects/{{project}}/autokeyConfig
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("projects/%s/autokeyConfig", v0)
			}
		}
		// Format: {{project}}
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_kms_project_kaj_policy_config":
		// Format: projects/{{project}}/kajPolicyConfig
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("projects/%s/kajPolicyConfig", v0)
			}
		}
		// Format: {{project}}
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_kms_secret_ciphertext":
		// No standard import format found in documentation
		return ""
	case "google_logging_billing_account_bucket_config":
		// Format: billingAccounts/{{billingAccount}}/locations/{{location}}/buckets/{{bucket_id}}
		{
			v0, ok0 := config["billingAccount"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["bucket_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("billingAccounts/%s/locations/%s/buckets/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_logging_billing_account_exclusion":
		// Format: billingAccounts/{{billing_account}}/exclusions/{{name}}
		{
			v0, ok0 := config["billing_account"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("billingAccounts/%s/exclusions/%s", v0, v1)
			}
		}
		return ""
	case "google_logging_billing_account_sink":
		// Format: billingAccounts/{{billing_account_id}}/sinks/{{sink_id}}
		{
			v0, ok0 := config["billing_account_id"].(string)
			v1, ok1 := config["sink_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("billingAccounts/%s/sinks/%s", v0, v1)
			}
		}
		return ""
	case "google_logging_folder_bucket_config":
		// Format: folders/{{folder}}/locations/{{location}}/buckets/{{bucket_id}}
		{
			v0, ok0 := config["folder"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["bucket_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("folders/%s/locations/%s/buckets/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_logging_folder_exclusion":
		// Format: folders/{{folder}}/exclusions/{{name}}
		{
			v0, ok0 := config["folder"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("folders/%s/exclusions/%s", v0, v1)
			}
		}
		return ""
	case "google_logging_folder_settings":
		// Format: folders/{{folder}}/settings
		{
			v0, ok0 := config["folder"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("folders/%s/settings", v0)
			}
		}
		// Format: {{folder}}
		{
			v0, ok0 := config["folder"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_logging_folder_sink":
		// Format: folders/{{folder_id}}/sinks/{{name}}
		{
			v0, ok0 := config["folder_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("folders/%s/sinks/%s", v0, v1)
			}
		}
		return ""
	case "google_logging_linked_dataset":
		// Format: {{parent}}/locations/{{location}}/buckets/{{bucket}}/links/{{link_id}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["bucket"].(string)
			v3, ok3 := config["link_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/locations/%s/buckets/%s/links/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_logging_log_scope":
		// Format: {{parent}}/locations/{{location}}/logScopes/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/locations/%s/logScopes/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_logging_log_view":
		// Format: {{parent}}/locations/{{location}}/buckets/{{bucket}}/views/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["bucket"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/locations/%s/buckets/%s/views/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_logging_log_view_iam":
		// No standard import format found in documentation
		return ""
	case "google_logging_metric":
		// Format: {{project}} {{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s %s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_logging_organization_bucket_config":
		// Format: organizations/{{organization}}/locations/{{location}}/buckets/{{bucket_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["bucket_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("organizations/%s/locations/%s/buckets/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_logging_organization_exclusion":
		// Format: organizations/{{organization}}/exclusions/{{name}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("organizations/%s/exclusions/%s", v0, v1)
			}
		}
		return ""
	case "google_logging_organization_settings":
		// Format: organizations/{{organization}}/settings
		{
			v0, ok0 := config["organization"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("organizations/%s/settings", v0)
			}
		}
		// Format: {{organization}}
		{
			v0, ok0 := config["organization"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_logging_organization_sink":
		// Format: organizations/{{organization_id}}/sinks/{{sink_id}}
		{
			v0, ok0 := config["organization_id"].(string)
			v1, ok1 := config["sink_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("organizations/%s/sinks/%s", v0, v1)
			}
		}
		return ""
	case "google_logging_project_bucket_config":
		// Format: projects/{{project}}/locations/{{location}}/buckets/{{bucket_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["bucket_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/buckets/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_logging_project_exclusion":
		// Format: projects/{{project_id}}/exclusions/{{name}}
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/exclusions/%s", v0, v1)
			}
		}
		return ""
	case "google_logging_project_sink":
		// Format: projects/{{project_id}}/sinks/{{name}}
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/sinks/%s", v0, v1)
			}
		}
		return ""
	case "google_logging_saved_query":
		// Format: {{parent}}/locations/{{location}}/savedQueries/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/locations/%s/savedQueries/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_looker_instance":
		// Format: projects/{{project}}/locations/{{region}}/instances/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_lustre_instance":
		// Format: projects/{{project}}/locations/{{location}}/instances/{{instance_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{instance_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{instance_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_managed_kafka_acl":
		// Format: projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/acls/{{acl_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster"].(string)
			v3, ok3 := config["acl_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/clusters/%s/acls/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_managed_kafka_cluster":
		// Format: projects/{{project}}/locations/{{location}}/clusters/{{cluster_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{cluster_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{cluster_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_managed_kafka_connect_cluster":
		// Format: projects/{{project}}/locations/{{location}}/connectClusters/{{connect_cluster_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["connect_cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/connectClusters/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{connect_cluster_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["connect_cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{connect_cluster_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["connect_cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_managed_kafka_connector":
		// Format: projects/{{project}}/locations/{{location}}/connectClusters/{{connect_cluster}}/connectors/{{connector_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["connect_cluster"].(string)
			v3, ok3 := config["connector_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/connectClusters/%s/connectors/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{connect_cluster}}/{{connector_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["connect_cluster"].(string)
			v3, ok3 := config["connector_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{connect_cluster}}/{{connector_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["connect_cluster"].(string)
			v2, ok2 := config["connector_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_managed_kafka_topic":
		// Format: projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/topics/{{topic_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster"].(string)
			v3, ok3 := config["topic_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/clusters/%s/topics/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{cluster}}/{{topic_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cluster"].(string)
			v3, ok3 := config["topic_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{cluster}}/{{topic_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["cluster"].(string)
			v2, ok2 := config["topic_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_memcache_instance":
		// Format: projects/{{project}}/locations/{{region}}/instances/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_memorystore_instance":
		// Format: projects/{{project}}/locations/{{location}}/instances/{{instance_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{instance_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{instance_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_memorystore_instance_desired_user_created_endpoints":
		// Format: projects/{{project}}/locations/{{region}}/instances/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_migration_center_group":
		// Format: projects/{{project}}/locations/{{location}}/groups/{{group_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/groups/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{group_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{group_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_migration_center_preference_set":
		// Format: projects/{{project}}/locations/{{location}}/preferenceSets/{{preference_set_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["preference_set_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/preferenceSets/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{preference_set_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["preference_set_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{preference_set_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["preference_set_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_ml_engine_model":
		// Format: projects/{{project}}/models/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/models/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_model_armor_floorsetting":
		// Format: {{parent}}/locations/{{location}}/floorSetting
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/locations/%s/floorSetting", v0, v1)
			}
		}
		// Format: {{parent}}/{{location}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_model_armor_template":
		// Format: projects/{{project}}/locations/{{location}}/templates/{{template_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["template_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/templates/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{template_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["template_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{template_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["template_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_monitoring_alert_policy":
		// Format: projects/{{project}}/alertPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/alertPolicies/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_monitoring_custom_service":
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{project}} {{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s %s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_monitoring_dashboard":
		// Format: projects/{{project}}/dashboards/{{dashboard_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["dashboard_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/dashboards/%s", v0, v1)
			}
		}
		// Format: {{dashboard_id}}
		{
			v0, ok0 := config["dashboard_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_monitoring_group":
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{project}} {{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s %s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_monitoring_metric_descriptor":
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{project}} {{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s %s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_monitoring_monitored_project":
		// Format: v1/locations/global/metricsScopes/{{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("v1/locations/global/metricsScopes/%s", v0)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_monitoring_notification_channel":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_monitoring_service":
		// Format: projects/{{project}}/services/{{service_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["service_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/services/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{service_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["service_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{service_id}}
		{
			v0, ok0 := config["service_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_monitoring_slo":
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{project}} {{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s %s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_monitoring_uptime_check_config":
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{project}} {{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s %s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_netapp_active_directory":
		// Format: projects/{{project}}/locations/{{location}}/activeDirectories/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/activeDirectories/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_netapp_backup":
		// Format: projects/{{project}}/locations/{{location}}/backupVaults/{{vault_name}}/backups/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["vault_name"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s/backups/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{vault_name}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["vault_name"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{vault_name}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["vault_name"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_netapp_backup_policy":
		// Format: projects/{{project}}/locations/{{location}}/backupPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/backupPolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_netapp_backup_vault":
		// Format: projects/{{project}}/locations/{{location}}/backupVaults/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_netapp_host_group":
		// Format: projects/{{project}}/locations/{{location}}/hostGroups/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/hostGroups/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_netapp_kmsconfig":
		// Format: projects/{{project}}/locations/{{location}}/kmsConfigs/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/kmsConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_netapp_storage_pool":
		// Format: projects/{{project}}/locations/{{location}}/storagePools/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/storagePools/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_netapp_volume":
		// Format: projects/{{project}}/locations/{{location}}/volumes/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/volumes/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_netapp_volume_quota_rule":
		// Format: projects/{{project}}/locations/{{location}}/volumes/{{volume_name}}/quotaRules/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["volume_name"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/volumes/%s/quotaRules/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{volume_name}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["volume_name"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{volume_name}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["volume_name"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_netapp_volume_replication":
		// Format: projects/{{project}}/locations/{{location}}/volumes/{{volume_name}}/replications/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["volume_name"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/volumes/%s/replications/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{volume_name}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["volume_name"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{volume_name}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["volume_name"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_netapp_volume_snapshot":
		// Format: projects/{{project}}/locations/{{location}}/volumes/{{volume_name}}/snapshots/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["volume_name"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/volumes/%s/snapshots/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{volume_name}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["volume_name"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{volume_name}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["volume_name"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_network_connectivity_destination":
		// Format: projects/{{project}}/locations/{{location}}/multicloudDataTransferConfigs/{{multicloud_data_transfer_config}}/destinations/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["multicloud_data_transfer_config"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/multicloudDataTransferConfigs/%s/destinations/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{multicloud_data_transfer_config}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["multicloud_data_transfer_config"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{multicloud_data_transfer_config}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["multicloud_data_transfer_config"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_network_connectivity_gateway_advertised_route":
		// Format: projects/{{project}}/locations/{{location}}/spokes/{{spoke}}/gatewayAdvertisedRoutes/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["spoke"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/spokes/%s/gatewayAdvertisedRoutes/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{spoke}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["spoke"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{spoke}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["spoke"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_network_connectivity_group":
		// Format: projects/{{project}}/locations/global/hubs/{{hub}}/groups/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["hub"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/global/hubs/%s/groups/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{hub}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["hub"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{hub}}/{{name}}
		{
			v0, ok0 := config["hub"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_connectivity_hub":
		// Format: projects/{{project}}/locations/global/hubs/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/hubs/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_network_connectivity_hub_iam":
		// No standard import format found in documentation
		return ""
	case "google_network_connectivity_internal_range":
		// Format: projects/{{project}}/locations/global/internalRanges/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/internalRanges/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_network_connectivity_multicloud_data_transfer_config":
		// Format: projects/{{project}}/locations/{{location}}/multicloudDataTransferConfigs/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/multicloudDataTransferConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_connectivity_policy_based_route":
		// Format: projects/{{project}}/locations/global/policyBasedRoutes/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/policyBasedRoutes/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_network_connectivity_regional_endpoint":
		// Format: projects/{{project}}/locations/{{location}}/regionalEndpoints/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/regionalEndpoints/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_connectivity_service_connection_policy":
		// Format: projects/{{project}}/locations/{{location}}/serviceConnectionPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/serviceConnectionPolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_connectivity_spoke":
		// Format: projects/{{project}}/locations/{{location}}/spokes/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/spokes/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_connectivity_transport":
		// Format: projects/{{project}}/locations/{{region}}/transports/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/transports/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_network_management_connectivity_test":
		// Format: projects/{{project}}/locations/global/connectivityTests/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/connectivityTests/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_network_management_organization_vpc_flow_logs_config":
		// Format: organizations/{{organization}}/locations/{{location}}/vpcFlowLogsConfigs/{{vpc_flow_logs_config_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["vpc_flow_logs_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("organizations/%s/locations/%s/vpcFlowLogsConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{organization}}/{{location}}/{{vpc_flow_logs_config_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["vpc_flow_logs_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_network_management_vpc_flow_logs_config":
		// Format: projects/{{project}}/locations/{{location}}/vpcFlowLogsConfigs/{{vpc_flow_logs_config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["vpc_flow_logs_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/vpcFlowLogsConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{vpc_flow_logs_config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["vpc_flow_logs_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{vpc_flow_logs_config_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["vpc_flow_logs_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_security_address_group":
		// Format: {{parent}}/locations/{{location}}/addressGroups/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/locations/%s/addressGroups/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_network_security_address_group_iam":
		// No standard import format found in documentation
		return ""
	case "google_network_security_authorization_policy":
		// Format: projects/{{project}}/locations/{{location}}/authorizationPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/authorizationPolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_security_authz_policy":
		// Format: projects/{{project}}/locations/{{location}}/authzPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/authzPolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_network_security_backend_authentication_config":
		// Format: projects/{{project}}/locations/{{location}}/backendAuthenticationConfigs/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/backendAuthenticationConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_security_client_tls_policy":
		// Format: projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/clientTlsPolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_security_dns_threat_detector":
		// Format: projects/{{project}}/locations/{{location}}/dnsThreatDetectors/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/dnsThreatDetectors/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_network_security_firewall_endpoint":
		// Format: {{parent}}/locations/{{location}}/firewallEndpoints/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/locations/%s/firewallEndpoints/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_network_security_firewall_endpoint_association":
		// Format: {{parent}}/locations/{{location}}/firewallEndpointAssociations/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/locations/%s/firewallEndpointAssociations/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_network_security_gateway_security_policy":
		// Format: projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_security_gateway_security_policy_rule":
		// Format: projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{gateway_security_policy}}/rules/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["gateway_security_policy"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s/rules/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{gateway_security_policy}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["gateway_security_policy"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{gateway_security_policy}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["gateway_security_policy"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_network_security_intercept_deployment":
		// Format: projects/{{project}}/locations/{{location}}/interceptDeployments/{{intercept_deployment_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["intercept_deployment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/interceptDeployments/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{intercept_deployment_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["intercept_deployment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{intercept_deployment_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["intercept_deployment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_security_intercept_deployment_group":
		// Format: projects/{{project}}/locations/{{location}}/interceptDeploymentGroups/{{intercept_deployment_group_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["intercept_deployment_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/interceptDeploymentGroups/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{intercept_deployment_group_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["intercept_deployment_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{intercept_deployment_group_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["intercept_deployment_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_security_intercept_endpoint_group":
		// Format: projects/{{project}}/locations/{{location}}/interceptEndpointGroups/{{intercept_endpoint_group_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["intercept_endpoint_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/interceptEndpointGroups/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{intercept_endpoint_group_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["intercept_endpoint_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{intercept_endpoint_group_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["intercept_endpoint_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_security_intercept_endpoint_group_association":
		// Format: projects/{{project}}/locations/{{location}}/interceptEndpointGroupAssociations/{{intercept_endpoint_group_association_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["intercept_endpoint_group_association_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/interceptEndpointGroupAssociations/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{intercept_endpoint_group_association_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["intercept_endpoint_group_association_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{intercept_endpoint_group_association_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["intercept_endpoint_group_association_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_security_mirroring_deployment":
		// Format: projects/{{project}}/locations/{{location}}/mirroringDeployments/{{mirroring_deployment_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["mirroring_deployment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/mirroringDeployments/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{mirroring_deployment_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["mirroring_deployment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{mirroring_deployment_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["mirroring_deployment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_security_mirroring_deployment_group":
		// Format: projects/{{project}}/locations/{{location}}/mirroringDeploymentGroups/{{mirroring_deployment_group_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["mirroring_deployment_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/mirroringDeploymentGroups/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{mirroring_deployment_group_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["mirroring_deployment_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{mirroring_deployment_group_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["mirroring_deployment_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_security_mirroring_endpoint":
		// Format: projects/{{project}}/locations/{{location}}/mirroringEndpoints/{{mirroring_endpoint_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["mirroring_endpoint_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/mirroringEndpoints/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{mirroring_endpoint_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["mirroring_endpoint_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{mirroring_endpoint_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["mirroring_endpoint_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_security_mirroring_endpoint_group":
		// Format: projects/{{project}}/locations/{{location}}/mirroringEndpointGroups/{{mirroring_endpoint_group_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["mirroring_endpoint_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/mirroringEndpointGroups/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{mirroring_endpoint_group_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["mirroring_endpoint_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{mirroring_endpoint_group_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["mirroring_endpoint_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_security_mirroring_endpoint_group_association":
		// Format: projects/{{project}}/locations/{{location}}/mirroringEndpointGroupAssociations/{{mirroring_endpoint_group_association_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["mirroring_endpoint_group_association_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/mirroringEndpointGroupAssociations/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{mirroring_endpoint_group_association_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["mirroring_endpoint_group_association_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{mirroring_endpoint_group_association_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["mirroring_endpoint_group_association_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_security_sac_attachment":
		// Format: projects/{{project}}/locations/{{location}}/sacAttachments/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/sacAttachments/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_security_sac_realm":
		// Format: projects/{{project}}/locations/global/sacRealms/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/sacRealms/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_network_security_security_profile":
		// Format: {{parent}}/locations/{{location}}/securityProfiles/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/locations/%s/securityProfiles/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_network_security_security_profile_group":
		// Format: {{parent}}/locations/{{location}}/securityProfileGroups/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/locations/%s/securityProfileGroups/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_network_security_server_tls_policy":
		// Format: projects/{{project}}/locations/{{location}}/serverTlsPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/serverTlsPolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_security_tls_inspection_policy":
		// Format: projects/{{project}}/locations/{{location}}/tlsInspectionPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/tlsInspectionPolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_security_ull_mirroring_collector":
		// Format: projects/{{project}}/locations/{{location}}/ullMirroringCollectors/{{ull_mirroring_collector_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["ull_mirroring_collector_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/ullMirroringCollectors/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{ull_mirroring_collector_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["ull_mirroring_collector_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{ull_mirroring_collector_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["ull_mirroring_collector_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_security_ull_mirroring_collector_rule":
		// Format: projects/{{project}}/locations/{{location}}/ullMirroringCollectors/{{ull_mirroring_collector}}/rules/{{ull_mirroring_collector_rule_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["ull_mirroring_collector"].(string)
			v3, ok3 := config["ull_mirroring_collector_rule_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/ullMirroringCollectors/%s/rules/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{ull_mirroring_collector}}/{{ull_mirroring_collector_rule_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["ull_mirroring_collector"].(string)
			v3, ok3 := config["ull_mirroring_collector_rule_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{ull_mirroring_collector}}/{{ull_mirroring_collector_rule_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["ull_mirroring_collector"].(string)
			v2, ok2 := config["ull_mirroring_collector_rule_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_network_security_ull_mirroring_engine":
		// Format: projects/{{project}}/locations/{{location}}/ullMirroringEngines/{{ull_mirroring_engine_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["ull_mirroring_engine_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/ullMirroringEngines/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{ull_mirroring_engine_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["ull_mirroring_engine_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{ull_mirroring_engine_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["ull_mirroring_engine_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_security_url_lists":
		// Format: projects/{{project}}/locations/{{location}}/urlLists/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/urlLists/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_services_agent_gateway":
		// Format: projects/{{project}}/locations/{{location}}/agentGateways/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/agentGateways/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_services_authz_extension":
		// Format: projects/{{project}}/locations/{{location}}/authzExtensions/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/authzExtensions/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_network_services_edge_cache_keyset":
		// Format: projects/{{project}}/locations/global/edgeCacheKeysets/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/edgeCacheKeysets/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_network_services_edge_cache_origin":
		// Format: projects/{{project}}/locations/global/edgeCacheOrigins/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/edgeCacheOrigins/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_network_services_edge_cache_service":
		// Format: projects/{{project}}/locations/global/edgeCacheServices/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/edgeCacheServices/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_network_services_endpoint_policy":
		// Format: projects/{{project}}/locations/global/endpointPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/endpointPolicies/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_network_services_gateway":
		// Format: projects/{{project}}/locations/{{location}}/gateways/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/gateways/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_services_grpc_route":
		// Format: projects/{{project}}/locations/{{location}}/grpcRoutes/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/grpcRoutes/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_services_http_route":
		// Format: projects/{{project}}/locations/global/httpRoutes/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/httpRoutes/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_network_services_lb_edge_extension":
		// Format: projects/{{project}}/locations/{{location}}/lbEdgeExtensions/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/lbEdgeExtensions/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_services_lb_route_extension":
		// Format: projects/{{project}}/locations/{{location}}/lbRouteExtensions/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/lbRouteExtensions/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_services_lb_traffic_extension":
		// Format: projects/{{project}}/locations/{{location}}/lbTrafficExtensions/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/lbTrafficExtensions/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_services_mesh":
		// Format: projects/{{project}}/locations/{{location}}/meshes/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/meshes/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_services_multicast_consumer_association":
		// Format: projects/{{project}}/locations/{{location}}/multicastConsumerAssociations/{{multicast_consumer_association_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["multicast_consumer_association_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/multicastConsumerAssociations/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{multicast_consumer_association_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["multicast_consumer_association_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{multicast_consumer_association_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["multicast_consumer_association_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_services_multicast_domain":
		// Format: projects/{{project}}/locations/{{location}}/multicastDomains/{{multicast_domain_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["multicast_domain_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/multicastDomains/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{multicast_domain_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["multicast_domain_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{multicast_domain_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["multicast_domain_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_services_multicast_domain_activation":
		// Format: projects/{{project}}/locations/{{location}}/multicastDomainActivations/{{multicast_domain_activation_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["multicast_domain_activation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/multicastDomainActivations/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{multicast_domain_activation_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["multicast_domain_activation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{multicast_domain_activation_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["multicast_domain_activation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_services_multicast_domain_group":
		// Format: projects/{{project}}/locations/{{location}}/multicastDomainGroups/{{multicast_domain_group_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["multicast_domain_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/multicastDomainGroups/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{multicast_domain_group_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["multicast_domain_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{multicast_domain_group_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["multicast_domain_group_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_services_multicast_group_consumer_activation":
		// Format: projects/{{project}}/locations/{{location}}/multicastGroupConsumerActivations/{{multicast_group_consumer_activation_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["multicast_group_consumer_activation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/multicastGroupConsumerActivations/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{multicast_group_consumer_activation_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["multicast_group_consumer_activation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{multicast_group_consumer_activation_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["multicast_group_consumer_activation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_services_multicast_group_producer_activation":
		// Format: projects/{{project}}/locations/{{location}}/multicastGroupProducerActivations/{{multicast_group_producer_activation_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["multicast_group_producer_activation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/multicastGroupProducerActivations/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{multicast_group_producer_activation_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["multicast_group_producer_activation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{multicast_group_producer_activation_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["multicast_group_producer_activation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_services_multicast_group_range":
		// Format: projects/{{project}}/locations/{{location}}/multicastGroupRanges/{{multicast_group_range_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["multicast_group_range_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/multicastGroupRanges/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{multicast_group_range_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["multicast_group_range_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{multicast_group_range_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["multicast_group_range_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_services_multicast_group_range_activation":
		// Format: projects/{{project}}/locations/{{location}}/multicastGroupRangeActivations/{{multicast_group_range_activation_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["multicast_group_range_activation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/multicastGroupRangeActivations/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{multicast_group_range_activation_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["multicast_group_range_activation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{multicast_group_range_activation_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["multicast_group_range_activation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_services_multicast_producer_association":
		// Format: projects/{{project}}/locations/{{location}}/multicastProducerAssociations/{{multicast_producer_association_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["multicast_producer_association_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/multicastProducerAssociations/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{multicast_producer_association_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["multicast_producer_association_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{multicast_producer_association_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["multicast_producer_association_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_services_service_binding":
		// Format: projects/{{project}}/locations/global/serviceBindings/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/serviceBindings/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_network_services_service_lb_policies":
		// Format: projects/{{project}}/locations/{{location}}/serviceLbPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/serviceLbPolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_services_tcp_route":
		// Format: projects/{{project}}/locations/global/tcpRoutes/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/tcpRoutes/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_network_services_tls_route":
		// Format: projects/{{project}}/locations/{{location}}/tlsRoutes/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/tlsRoutes/%s", v0, v1, v2)
			}
		}
		// Format: projects/{{project}}/locations/global/tlsRoutes/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/tlsRoutes/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_network_services_wasm_plugin":
		// Format: projects/{{project}}/locations/{{location}}/wasmPlugins/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/wasmPlugins/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_notebooks_environment":
		// Format: projects/{{project}}/locations/{{location}}/environments/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/environments/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_notebooks_instance":
		// Format: projects/{{project}}/locations/{{location}}/instances/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_notebooks_instance_iam":
		// No standard import format found in documentation
		return ""
	case "google_notebooks_runtime":
		// Format: projects/{{project}}/locations/{{location}}/runtimes/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/runtimes/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_notebooks_runtime_iam":
		// No standard import format found in documentation
		return ""
	case "google_observability_folder_settings":
		// Format: folders/{{folder}}/locations/{{location}}/settings
		{
			v0, ok0 := config["folder"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("folders/%s/locations/%s/settings", v0, v1)
			}
		}
		// Format: {{folder}}/{{location}}
		{
			v0, ok0 := config["folder"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_observability_organization_settings":
		// Format: organizations/{{organization}}/locations/{{location}}/settings
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("organizations/%s/locations/%s/settings", v0, v1)
			}
		}
		// Format: {{organization}}/{{location}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_observability_project_settings":
		// Format: projects/{{project}}/locations/{{location}}/settings
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/settings", v0, v1)
			}
		}
		// Format: {{project}}/{{location}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{location}}
		{
			v0, ok0 := config["location"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_observability_trace_scope":
		// Format: projects/{{project}}/locations/{{location}}/traceScopes/{{trace_scope_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["trace_scope_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/traceScopes/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{trace_scope_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["trace_scope_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{trace_scope_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["trace_scope_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_oracle_database_autonomous_database":
		// Format: projects/{{project}}/locations/{{location}}/autonomousDatabases/{{autonomous_database_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["autonomous_database_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{autonomous_database_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["autonomous_database_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{autonomous_database_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["autonomous_database_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_oracle_database_cloud_exadata_infrastructure":
		// Format: projects/{{project}}/locations/{{location}}/cloudExadataInfrastructures/{{cloud_exadata_infrastructure_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cloud_exadata_infrastructure_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/cloudExadataInfrastructures/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{cloud_exadata_infrastructure_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cloud_exadata_infrastructure_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{cloud_exadata_infrastructure_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["cloud_exadata_infrastructure_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_oracle_database_cloud_vm_cluster":
		// Format: projects/{{project}}/locations/{{location}}/cloudVmClusters/{{cloud_vm_cluster_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cloud_vm_cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/cloudVmClusters/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{cloud_vm_cluster_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["cloud_vm_cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{cloud_vm_cluster_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["cloud_vm_cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_oracle_database_db_system":
		// Format: projects/{{project}}/locations/{{location}}/dbSystems/{{db_system_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["db_system_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/dbSystems/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{db_system_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["db_system_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{db_system_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["db_system_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_oracle_database_exadb_vm_cluster":
		// Format: projects/{{project}}/locations/{{location}}/exadbVmClusters/{{exadb_vm_cluster_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["exadb_vm_cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/exadbVmClusters/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{exadb_vm_cluster_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["exadb_vm_cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{exadb_vm_cluster_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["exadb_vm_cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_oracle_database_exascale_db_storage_vault":
		// Format: projects/{{project}}/locations/{{location}}/exascaleDbStorageVaults/{{exascale_db_storage_vault_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["exascale_db_storage_vault_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/exascaleDbStorageVaults/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{exascale_db_storage_vault_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["exascale_db_storage_vault_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{exascale_db_storage_vault_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["exascale_db_storage_vault_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_oracle_database_odb_network":
		// Format: projects/{{project}}/locations/{{location}}/odbNetworks/{{odb_network_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["odb_network_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/odbNetworks/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{odb_network_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["odb_network_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{odb_network_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["odb_network_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_oracle_database_odb_subnet":
		// Format: projects/{{project}}/locations/{{location}}/odbNetworks/{{odbnetwork}}/odbSubnets/{{odb_subnet_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["odbnetwork"].(string)
			v3, ok3 := config["odb_subnet_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/odbNetworks/%s/odbSubnets/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{odbnetwork}}/{{odb_subnet_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["odbnetwork"].(string)
			v3, ok3 := config["odb_subnet_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{odbnetwork}}/{{odb_subnet_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["odbnetwork"].(string)
			v2, ok2 := config["odb_subnet_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_org_policy_custom_constraint":
		// Format: {{parent}}/customConstraints/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/customConstraints/%s", v0, v1)
			}
		}
		return ""
	case "google_org_policy_policy":
		// Format: {{parent}}/policies/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/policies/%s", v0, v1)
			}
		}
		return ""
	case "google_organization_access_approval_settings":
		// Format: organizations/{{organization_id}}/accessApprovalSettings
		{
			v0, ok0 := config["organization_id"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("organizations/%s/accessApprovalSettings", v0)
			}
		}
		// Format: {{organization_id}}
		{
			v0, ok0 := config["organization_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_organization_service_identity":
		// No standard import format found in documentation
		return ""
	case "google_os_config_guest_policies":
		// Format: projects/{{project}}/guestPolicies/{{guest_policy_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["guest_policy_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/guestPolicies/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{guest_policy_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["guest_policy_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{guest_policy_id}}
		{
			v0, ok0 := config["guest_policy_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_os_config_os_policy_assignment":
		// Format: projects/{{project}}/locations/{{location}}/osPolicyAssignments/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/osPolicyAssignments/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_os_config_patch_deployment":
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{project}} {{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s %s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_os_config_v2_policy_orchestrator":
		// Format: projects/{{project}}/locations/global/policyOrchestrators/{{policy_orchestrator_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["policy_orchestrator_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/policyOrchestrators/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{policy_orchestrator_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["policy_orchestrator_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{policy_orchestrator_id}}
		{
			v0, ok0 := config["policy_orchestrator_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_os_config_v2_policy_orchestrator_for_folder":
		// Format: folders/{{folder_id}}/locations/global/policyOrchestrators/{{policy_orchestrator_id}}
		{
			v0, ok0 := config["folder_id"].(string)
			v1, ok1 := config["policy_orchestrator_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("folders/%s/locations/global/policyOrchestrators/%s", v0, v1)
			}
		}
		// Format: {{folder_id}}/{{policy_orchestrator_id}}
		{
			v0, ok0 := config["folder_id"].(string)
			v1, ok1 := config["policy_orchestrator_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_os_config_v2_policy_orchestrator_for_organization":
		// Format: organizations/{{organization_id}}/locations/global/policyOrchestrators/{{policy_orchestrator_id}}
		{
			v0, ok0 := config["organization_id"].(string)
			v1, ok1 := config["policy_orchestrator_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("organizations/%s/locations/global/policyOrchestrators/%s", v0, v1)
			}
		}
		// Format: {{organization_id}}/{{policy_orchestrator_id}}
		{
			v0, ok0 := config["organization_id"].(string)
			v1, ok1 := config["policy_orchestrator_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_os_login_ssh_public_key":
		// Format: users/{{user}}/sshPublicKeys/{{fingerprint}}
		{
			v0, ok0 := config["user"].(string)
			v1, ok1 := config["fingerprint"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("users/%s/sshPublicKeys/%s", v0, v1)
			}
		}
		// Format: {{user}}/{{fingerprint}}
		{
			v0, ok0 := config["user"].(string)
			v1, ok1 := config["fingerprint"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_parallelstore_instance":
		// Format: projects/{{project}}/locations/{{location}}/instances/{{instance_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{instance_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{instance_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_parameter_manager_parameter":
		// Format: projects/{{project}}/locations/global/parameters/{{parameter_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["parameter_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/parameters/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{parameter_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["parameter_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{parameter_id}}
		{
			v0, ok0 := config["parameter_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_parameter_manager_parameter_version":
		// Format: projects/{{project}}/locations/global/parameters/{{parameter_id}}/versions/{{parameter_version_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["parameter_id"].(string)
			v2, ok2 := config["parameter_version_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/global/parameters/%s/versions/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_parameter_manager_regional_parameter":
		// Format: projects/{{project}}/locations/{{location}}/parameters/{{parameter_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["parameter_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/parameters/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{parameter_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["parameter_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{parameter_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["parameter_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_parameter_manager_regional_parameter_version":
		// Format: projects/{{project}}/locations/{{location}}/parameters/{{parameter_id}}/versions/{{parameter_version_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["parameter_id"].(string)
			v3, ok3 := config["parameter_version_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/parameters/%s/versions/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_privateca_ca_pool":
		// Format: projects/{{project}}/locations/{{location}}/caPools/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/caPools/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_privateca_ca_pool_iam":
		// No standard import format found in documentation
		return ""
	case "google_privateca_certificate":
		// Format: projects/{{project}}/locations/{{location}}/caPools/{{pool}}/certificates/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["pool"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificates/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{pool}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["pool"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{pool}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["pool"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_privateca_certificate_authority":
		// Format: projects/{{project}}/locations/{{location}}/caPools/{{pool}}/certificateAuthorities/{{certificate_authority_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["pool"].(string)
			v3, ok3 := config["certificate_authority_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{pool}}/{{certificate_authority_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["pool"].(string)
			v3, ok3 := config["certificate_authority_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{pool}}/{{certificate_authority_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["pool"].(string)
			v2, ok2 := config["certificate_authority_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_privateca_certificate_template":
		// Format: projects/{{project}}/locations/{{location}}/certificateTemplates/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_privateca_certificate_template_iam":
		// No standard import format found in documentation
		return ""
	case "google_privileged_access_manager_entitlement":
		// Format: {{parent}}/locations/{{location}}/entitlements/{{entitlement_id}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["entitlement_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/locations/%s/entitlements/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_privileged_access_manager_settings":
		// Format: {{parent}}/locations/{{location}}/settings
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["location"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/locations/%s/settings", v0, v1)
			}
		}
		return ""
	case "google_project_access_approval_settings":
		// Format: projects/{{project_id}}/accessApprovalSettings
		{
			v0, ok0 := config["project_id"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("projects/%s/accessApprovalSettings", v0)
			}
		}
		// Format: {{project_id}}
		{
			v0, ok0 := config["project_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_project_service_identity":
		// No standard import format found in documentation
		return ""
	case "google_project_usage_export_bucket":
		// Format: {{project_id}}
		{
			v0, ok0 := config["project_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_public_ca_external_account_key":
		// No standard import format found in documentation
		return ""
	case "google_pubsub_lite_reservation":
		// Format: projects/{{project}}/locations/{{region}}/reservations/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/reservations/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_pubsub_lite_subscription":
		// Format: projects/{{project}}/locations/{{zone}}/subscriptions/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/subscriptions/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{zone}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{zone}}/{{name}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_pubsub_lite_topic":
		// Format: projects/{{project}}/locations/{{zone}}/topics/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/topics/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{zone}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{zone}}/{{name}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_pubsub_schema":
		// Format: projects/{{project}}/schemas/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/schemas/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_pubsub_schema_iam":
		// No standard import format found in documentation
		return ""
	case "google_pubsub_subscription":
		// Format: projects/{{project}}/subscriptions/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/subscriptions/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_pubsub_subscription_iam":
		// Format: "projects/{{project_id}}/subscriptions/{{subscription}} roles/editor jane@example.com"
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["subscription"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("\"projects/%s/subscriptions/%s roles/editor jane@example.com\"", v0, v1)
			}
		}
		return ""
	case "google_pubsub_topic":
		// Format: projects/{{project}}/topics/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/topics/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_pubsub_topic_iam":
		// No standard import format found in documentation
		return ""
	case "google_recaptcha_enterprise_key":
		// Format: projects/{{project}}/keys/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/keys/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_redis_cluster":
		// Format: projects/{{project}}/locations/{{region}}/clusters/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_redis_cluster_user_created_connections":
		// Format: projects/{{project}}/locations/{{region}}/clusters/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_redis_instance":
		// Format: projects/{{project}}/locations/{{region}}/instances/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_resource_manager_capability":
		// Format: {{parent}}/capabilities/{{capability_name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["capability_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/capabilities/%s", v0, v1)
			}
		}
		// Format: {{parent}}/{{capability_name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["capability_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_resource_manager_lien":
		// Format: {{parent}}/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_runtimeconfig_config":
		// Format: projects/{{project_id}}/configs/{{name}}
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/configs/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_runtimeconfig_config_iam":
		// No standard import format found in documentation
		return ""
	case "google_runtimeconfig_variable":
		// Format: projects/my-gcp-project/configs/{{config_id}}/variables/{{name}}
		{
			v0, ok0 := config["config_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/my-gcp-project/configs/%s/variables/%s", v0, v1)
			}
		}
		// Format: {{config_id}}/{{name}}
		{
			v0, ok0 := config["config_id"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_saas_runtime_release":
		// Format: projects/{{project}}/locations/{{location}}/releases/{{release_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["release_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/releases/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{release_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["release_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{release_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["release_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_saas_runtime_rollout_kind":
		// Format: projects/{{project}}/locations/{{location}}/rolloutKinds/{{rollout_kind_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["rollout_kind_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/rolloutKinds/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{rollout_kind_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["rollout_kind_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{rollout_kind_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["rollout_kind_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_saas_runtime_saas":
		// Format: projects/{{project}}/locations/{{location}}/saas/{{saas_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["saas_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/saas/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{saas_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["saas_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{saas_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["saas_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_saas_runtime_tenant":
		// Format: projects/{{project}}/locations/{{location}}/tenants/{{tenant_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["tenant_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/tenants/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{tenant_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["tenant_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{tenant_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["tenant_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_saas_runtime_unit":
		// Format: projects/{{project}}/locations/{{location}}/units/{{unit_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["unit_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/units/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{unit_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["unit_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{unit_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["unit_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_saas_runtime_unit_kind":
		// Format: projects/{{project}}/locations/{{location}}/unitKinds/{{unit_kind_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["unit_kind_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/unitKinds/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{unit_kind_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["unit_kind_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{unit_kind_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["unit_kind_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_saas_runtime_unit_operation":
		// Format: projects/{{project}}/locations/{{location}}/unitOperations/{{unit_operation_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["unit_operation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/unitOperations/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{unit_operation_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["unit_operation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{unit_operation_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["unit_operation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_scc_event_threat_detection_custom_module":
		// Format: organizations/{{organization}}/eventThreatDetectionSettings/customModules/{{name}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("organizations/%s/eventThreatDetectionSettings/customModules/%s", v0, v1)
			}
		}
		// Format: {{organization}}/{{name}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_scc_folder_custom_module":
		// Format: folders/{{folder}}/securityHealthAnalyticsSettings/customModules/{{name}}
		{
			v0, ok0 := config["folder"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("folders/%s/securityHealthAnalyticsSettings/customModules/%s", v0, v1)
			}
		}
		// Format: {{folder}}/{{name}}
		{
			v0, ok0 := config["folder"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_scc_folder_notification_config":
		// Format: folders/{{folder}}/notificationConfigs/{{config_id}}
		{
			v0, ok0 := config["folder"].(string)
			v1, ok1 := config["config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("folders/%s/notificationConfigs/%s", v0, v1)
			}
		}
		// Format: {{folder}}/{{config_id}}
		{
			v0, ok0 := config["folder"].(string)
			v1, ok1 := config["config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_scc_folder_scc_big_query_export":
		// Format: folders/{{folder}}/bigQueryExports/{{big_query_export_id}}
		{
			v0, ok0 := config["folder"].(string)
			v1, ok1 := config["big_query_export_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("folders/%s/bigQueryExports/%s", v0, v1)
			}
		}
		// Format: {{folder}}/{{big_query_export_id}}
		{
			v0, ok0 := config["folder"].(string)
			v1, ok1 := config["big_query_export_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_scc_management_folder_security_health_analytics_custom_module":
		// Format: folders/{{folder}}/locations/{{location}}/securityHealthAnalyticsCustomModules/{{name}}
		{
			v0, ok0 := config["folder"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("folders/%s/locations/%s/securityHealthAnalyticsCustomModules/%s", v0, v1, v2)
			}
		}
		// Format: {{folder}}/{{location}}/{{name}}
		{
			v0, ok0 := config["folder"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_scc_management_organization_event_threat_detection_custom_module":
		// Format: organizations/{{organization}}/locations/{{location}}/eventThreatDetectionCustomModules/{{name}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("organizations/%s/locations/%s/eventThreatDetectionCustomModules/%s", v0, v1, v2)
			}
		}
		// Format: {{organization}}/{{location}}/{{name}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_scc_management_organization_security_health_analytics_custom_module":
		// Format: organizations/{{organization}}/locations/{{location}}/securityHealthAnalyticsCustomModules/{{name}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("organizations/%s/locations/%s/securityHealthAnalyticsCustomModules/%s", v0, v1, v2)
			}
		}
		// Format: {{organization}}/{{location}}/{{name}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_scc_management_project_security_health_analytics_custom_module":
		// Format: projects/{{project}}/locations/{{location}}/securityHealthAnalyticsCustomModules/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/securityHealthAnalyticsCustomModules/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_scc_mute_config":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_scc_notification_config":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_scc_organization_custom_module":
		// Format: organizations/{{organization}}/securityHealthAnalyticsSettings/customModules/{{name}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("organizations/%s/securityHealthAnalyticsSettings/customModules/%s", v0, v1)
			}
		}
		// Format: {{organization}}/{{name}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_scc_organization_scc_big_query_export":
		// Format: organizations/{{organization}}/bigQueryExports/{{big_query_export_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["big_query_export_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("organizations/%s/bigQueryExports/%s", v0, v1)
			}
		}
		// Format: {{organization}}/{{big_query_export_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["big_query_export_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_scc_project_custom_module":
		// Format: projects/{{project}}/securityHealthAnalyticsSettings/customModules/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/securityHealthAnalyticsSettings/customModules/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_scc_project_notification_config":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_scc_project_scc_big_query_export":
		// Format: projects/{{project}}/bigQueryExports/{{big_query_export_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["big_query_export_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/bigQueryExports/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{big_query_export_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["big_query_export_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{big_query_export_id}}
		{
			v0, ok0 := config["big_query_export_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_scc_source":
		// Format: organizations/{{organization}}/sources/{{name}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("organizations/%s/sources/%s", v0, v1)
			}
		}
		// Format: {{organization}}/{{name}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_scc_source_iam":
		// No standard import format found in documentation
		return ""
	case "google_scc_v2_folder_mute_config":
		// Format: folders/{{folder}}/locations/{{location}}/muteConfigs/{{mute_config_id}}
		{
			v0, ok0 := config["folder"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["mute_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("folders/%s/locations/%s/muteConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{folder}}/{{location}}/{{mute_config_id}}
		{
			v0, ok0 := config["folder"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["mute_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_scc_v2_folder_notification_config":
		// Format: folders/{{folder}}/locations/{{location}}/notificationConfigs/{{config_id}}
		{
			v0, ok0 := config["folder"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("folders/%s/locations/%s/notificationConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{folder}}/{{location}}/{{config_id}}
		{
			v0, ok0 := config["folder"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_scc_v2_folder_scc_big_query_export":
		// Format: folders/{{folder}}/locations/{{location}}/bigQueryExports/{{big_query_export_id}}
		{
			v0, ok0 := config["folder"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["big_query_export_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("folders/%s/locations/%s/bigQueryExports/%s", v0, v1, v2)
			}
		}
		// Format: {{folder}}/{{location}}/{{big_query_export_id}}
		{
			v0, ok0 := config["folder"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["big_query_export_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_scc_v2_organization_mute_config":
		// Format: organizations/{{organization}}/locations/{{location}}/muteConfigs/{{mute_config_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["mute_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("organizations/%s/locations/%s/muteConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{organization}}/{{location}}/{{mute_config_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["mute_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_scc_v2_organization_notification_config":
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_scc_v2_organization_scc_big_query_export":
		// Format: organizations/{{organization}}/locations/{{location}}/bigQueryExports/{{big_query_export_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["big_query_export_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("organizations/%s/locations/%s/bigQueryExports/%s", v0, v1, v2)
			}
		}
		// Format: {{organization}}/{{location}}/{{big_query_export_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["big_query_export_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_scc_v2_organization_scc_big_query_exports":
		// Format: organizations/{{organization}}/locations/{{location}}/bigQueryExports/{{big_query_export_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["big_query_export_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("organizations/%s/locations/%s/bigQueryExports/%s", v0, v1, v2)
			}
		}
		// Format: {{organization}}/{{location}}/{{big_query_export_id}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["big_query_export_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_scc_v2_organization_source":
		// Format: organizations/{{organization}}/sources/{{name}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("organizations/%s/sources/%s", v0, v1)
			}
		}
		// Format: {{organization}}/{{name}}
		{
			v0, ok0 := config["organization"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_scc_v2_organization_source_iam":
		// No standard import format found in documentation
		return ""
	case "google_scc_v2_project_mute_config":
		// Format: projects/{{project}}/locations/{{location}}/muteConfigs/{{mute_config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["mute_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/muteConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{mute_config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["mute_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{mute_config_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["mute_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_scc_v2_project_notification_config":
		// Format: projects/{{project}}/locations/{{location}}/notificationConfigs/{{config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/notificationConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{config_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_scc_v2_project_scc_big_query_export":
		// Format: projects/{{project}}/locations/{{location}}/bigQueryExports/{{big_query_export_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["big_query_export_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/bigQueryExports/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{big_query_export_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["big_query_export_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{big_query_export_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["big_query_export_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_secret_manager_regional_secret":
		// Format: projects/{{project}}/locations/{{location}}/secrets/{{secret_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["secret_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/secrets/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{secret_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["secret_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{secret_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["secret_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_secret_manager_regional_secret_iam":
		// No standard import format found in documentation
		return ""
	case "google_secret_manager_regional_secret_version":
		// Format: projects/{{project}}/locations/{{location}}/secrets/{{secret_id}}/versions/{{version}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["secret_id"].(string)
			v3, ok3 := config["version"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/secrets/%s/versions/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_secret_manager_secret":
		// Format: projects/{{project}}/secrets/{{secret_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["secret_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/secrets/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{secret_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["secret_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{secret_id}}
		{
			v0, ok0 := config["secret_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_secret_manager_secret_iam":
		// No standard import format found in documentation
		return ""
	case "google_secret_manager_secret_version":
		// Format: projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["secret_id"].(string)
			v2, ok2 := config["version"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/secrets/%s/versions/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_secure_source_manager_branch_rule":
		// Format: projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}/branchRules/{{branch_rule_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["repository_id"].(string)
			v3, ok3 := config["branch_rule_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/repositories/%s/branchRules/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{repository_id}}/{{branch_rule_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["repository_id"].(string)
			v3, ok3 := config["branch_rule_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{repository_id}}/{{branch_rule_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["repository_id"].(string)
			v2, ok2 := config["branch_rule_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{branch_rule_id}}
		{
			v0, ok0 := config["branch_rule_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_secure_source_manager_hook":
		// Format: projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}/hooks/{{hook_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["repository_id"].(string)
			v3, ok3 := config["hook_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/repositories/%s/hooks/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{repository_id}}/{{hook_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["repository_id"].(string)
			v3, ok3 := config["hook_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{repository_id}}/{{hook_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["repository_id"].(string)
			v2, ok2 := config["hook_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{hook_id}}
		{
			v0, ok0 := config["hook_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_secure_source_manager_instance":
		// Format: projects/{{project}}/locations/{{location}}/instances/{{instance_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{instance_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{instance_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["instance_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{instance_id}}
		{
			v0, ok0 := config["instance_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_secure_source_manager_instance_iam":
		// No standard import format found in documentation
		return ""
	case "google_secure_source_manager_repository":
		// Format: projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["repository_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/repositories/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{repository_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["repository_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{repository_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["repository_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{repository_id}}
		{
			v0, ok0 := config["repository_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_secure_source_manager_repository_iam":
		// No standard import format found in documentation
		return ""
	case "google_security_scanner_scan_config":
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{project}} {{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s %s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_securityposture_posture":
		// Format: {{parent}}/locations/{{location}}/postures/{{posture_id}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["posture_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/locations/%s/postures/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_securityposture_posture_deployment":
		// Format: {{parent}}/locations/{{location}}/postureDeployments/{{posture_deployment_id}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["posture_deployment_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/locations/%s/postureDeployments/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_service_directory_endpoint":
		// Format: projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}}/endpoints/{{endpoint_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["namespace_id"].(string)
			v3, ok3 := config["service_id"].(string)
			v4, ok4 := config["endpoint_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s/endpoints/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{project}}/{{location}}/{{namespace_id}}/{{service_id}}/{{endpoint_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["namespace_id"].(string)
			v3, ok3 := config["service_id"].(string)
			v4, ok4 := config["endpoint_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("%s/%s/%s/%s/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{location}}/{{namespace_id}}/{{service_id}}/{{endpoint_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["namespace_id"].(string)
			v2, ok2 := config["service_id"].(string)
			v3, ok3 := config["endpoint_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_service_directory_namespace":
		// Format: projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["namespace_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{namespace_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["namespace_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{namespace_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["namespace_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_service_directory_namespace_iam":
		// No standard import format found in documentation
		return ""
	case "google_service_directory_service":
		// Format: projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["namespace_id"].(string)
			v3, ok3 := config["service_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{namespace_id}}/{{service_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["namespace_id"].(string)
			v3, ok3 := config["service_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{namespace_id}}/{{service_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["namespace_id"].(string)
			v2, ok2 := config["service_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_service_directory_service_iam":
		// No standard import format found in documentation
		return ""
	case "google_service_networking_connection":
		// Format: {{peering-network}}:{{service}}
		{
			v0, ok0 := config["peering-network"].(string)
			v1, ok1 := config["service"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s:%s", v0, v1)
			}
		}
		// Format: projects/{{project}}/global/networks/{{peering-network}}:{{service}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["peering-network"].(string)
			v2, ok2 := config["service"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/global/networks/%s:%s", v0, v1, v2)
			}
		}
		return ""
	case "google_service_networking_vpc_service_controls":
		// Format: services/{{service}}/projects/{{project}}/networks/{{network}}
		{
			v0, ok0 := config["service"].(string)
			v1, ok1 := config["project"].(string)
			v2, ok2 := config["network"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("services/%s/projects/%s/networks/%s", v0, v1, v2)
			}
		}
		// Format: {{service}}/{{project}}/{{network}}
		{
			v0, ok0 := config["service"].(string)
			v1, ok1 := config["project"].(string)
			v2, ok2 := config["network"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{service}}/{{network}}
		{
			v0, ok0 := config["service"].(string)
			v1, ok1 := config["network"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_service_usage_consumer_quota_override":
		// Format: projects/{{project}}/services/{{service}}/consumerQuotaMetrics/{{metric}}/limits/{{limit}}/consumerOverrides/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["service"].(string)
			v2, ok2 := config["metric"].(string)
			v3, ok3 := config["limit"].(string)
			v4, ok4 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("projects/%s/services/%s/consumerQuotaMetrics/%s/limits/%s/consumerOverrides/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: services/{{service}}/consumerQuotaMetrics/{{metric}}/limits/{{limit}}/consumerOverrides/{{name}}
		{
			v0, ok0 := config["service"].(string)
			v1, ok1 := config["metric"].(string)
			v2, ok2 := config["limit"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("services/%s/consumerQuotaMetrics/%s/limits/%s/consumerOverrides/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{service}}/{{metric}}/{{limit}}/{{name}}
		{
			v0, ok0 := config["service"].(string)
			v1, ok1 := config["metric"].(string)
			v2, ok2 := config["limit"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_site_verification_owner":
		// Format: webResource/{{web_resource_id}}/{{email}}
		{
			v0, ok0 := config["web_resource_id"].(string)
			v1, ok1 := config["email"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("webResource/%s/%s", v0, v1)
			}
		}
		return ""
	case "google_site_verification_web_resource":
		// Format: webResource/{{web_resource_id}}
		{
			v0, ok0 := config["web_resource_id"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("webResource/%s", v0)
			}
		}
		// Format: {{web_resource_id}}
		{
			v0, ok0 := config["web_resource_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_sourcerepo_repository":
		// Format: projects/{{project}}/repos/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/repos/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_sourcerepo_repository_iam":
		// No standard import format found in documentation
		return ""
	case "google_spanner_backup_schedule":
		// Format: projects/{{project}}/instances/{{instance}}/databases/{{database}}/backupSchedules/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["database"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/instances/%s/databases/%s/backupSchedules/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{instance}}/{{database}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["database"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{instance}}/{{database}}/{{name}}
		{
			v0, ok0 := config["instance"].(string)
			v1, ok1 := config["database"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_spanner_database":
		// Format: projects/{{project}}/instances/{{instance}}/databases/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/instances/%s/databases/%s", v0, v1, v2)
			}
		}
		// Format: instances/{{instance}}/databases/{{name}}
		{
			v0, ok0 := config["instance"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("instances/%s/databases/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{instance}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{instance}}/{{name}}
		{
			v0, ok0 := config["instance"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_spanner_database_iam":
		// Format: "{{project}}/{{instance}}/{{database}} roles/viewer user:foo@example.com"
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["database"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("\"%s/%s/%s roles/viewer user:foo@example.com\"", v0, v1, v2)
			}
		}
		return ""
	case "google_spanner_instance":
		// Format: projects/{{project}}/instances/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/instances/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_spanner_instance_config":
		// Format: projects/{{project}}/instanceConfigs/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/instanceConfigs/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_spanner_instance_iam":
		// Format: "{{project}}/{{instance}} roles/viewer user:foo@example.com"
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("\"%s/%s roles/viewer user:foo@example.com\"", v0, v1)
			}
		}
		return ""
	case "google_spanner_instance_partition":
		// Format: projects/{{project}}/instances/{{instance}}/instancePartitions/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/instances/%s/instancePartitions/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{instance}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{instance}}/{{name}}
		{
			v0, ok0 := config["instance"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_sql_database":
		// Format: projects/{{project}}/instances/{{instance}}/databases/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/instances/%s/databases/%s", v0, v1, v2)
			}
		}
		// Format: instances/{{instance}}/databases/{{name}}
		{
			v0, ok0 := config["instance"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("instances/%s/databases/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{instance}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{instance}}/{{name}}
		{
			v0, ok0 := config["instance"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_sql_database_instance":
		// Format: projects/{{project}}/instances/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/instances/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_sql_provision_script":
		// No standard import format found in documentation
		return ""
	case "google_sql_source_representation_instance":
		// Format: projects/{{project}}/instances/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/instances/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_sql_ssl_cert":
		// No standard import format found in documentation
		return ""
	case "google_sql_user":
		// Format: {{project_id}}/{{instance}}/{{host}}/{{name}}
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["instance"].(string)
			v2, ok2 := config["host"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_storage_anywhere_cache":
		// Format: b/{{bucket}}/anywhereCaches/{{anywhere_cache_id}}
		{
			v0, ok0 := config["bucket"].(string)
			v1, ok1 := config["anywhere_cache_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("b/%s/anywhereCaches/%s", v0, v1)
			}
		}
		// Format: {{bucket}}/{{anywhere_cache_id}}
		{
			v0, ok0 := config["bucket"].(string)
			v1, ok1 := config["anywhere_cache_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_storage_batch_operations_job":
		// Format: projects/{{project}}/locations/global/jobs/{{job_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["job_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/jobs/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{job_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["job_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{job_id}}
		{
			v0, ok0 := config["job_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_storage_bucket":
		// Format: {{project_id}}/{{bucket}}
		{
			v0, ok0 := config["project_id"].(string)
			v1, ok1 := config["bucket"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{bucket}}
		{
			v0, ok0 := config["bucket"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_storage_bucket_access_control":
		// Format: {{bucket}}/{{entity}}
		{
			v0, ok0 := config["bucket"].(string)
			v1, ok1 := config["entity"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_storage_bucket_acl":
		// No standard import format found in documentation
		return ""
	case "google_storage_bucket_iam":
		// No standard import format found in documentation
		return ""
	case "google_storage_bucket_object":
		// No standard import format found in documentation
		return ""
	case "google_storage_control_folder_intelligence_config":
		// Format: folders/{{name}}/locations/global/intelligenceConfig
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("folders/%s/locations/global/intelligenceConfig", v0)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_storage_control_organization_intelligence_config":
		// Format: organizations/{{name}}/locations/global/intelligenceConfig
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("organizations/%s/locations/global/intelligenceConfig", v0)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_storage_control_project_intelligence_config":
		// Format: projects/{{name}}/locations/global/intelligenceConfig
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("projects/%s/locations/global/intelligenceConfig", v0)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_storage_default_object_access_control":
		// Format: {{bucket}}/{{entity}}
		{
			v0, ok0 := config["bucket"].(string)
			v1, ok1 := config["entity"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_storage_default_object_acl":
		// No standard import format found in documentation
		return ""
	case "google_storage_folder":
		// Format: {{bucket}}/folders/{{name}}
		{
			v0, ok0 := config["bucket"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/folders/%s", v0, v1)
			}
		}
		// Format: {{bucket}}/{{name}}
		{
			v0, ok0 := config["bucket"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_storage_hmac_key":
		// Format: projects/{{project}}/hmacKeys/{{access_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["access_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/hmacKeys/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{access_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["access_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{access_id}}
		{
			v0, ok0 := config["access_id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_storage_insights_dataset_config":
		// Format: projects/{{project}}/locations/{{location}}/datasetConfigs/{{dataset_config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["dataset_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/datasetConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{dataset_config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["dataset_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{dataset_config_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["dataset_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_storage_insights_report_config":
		// Format: projects/{{project}}/locations/{{location}}/reportConfigs/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/reportConfigs/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_storage_managed_folder":
		// Format: {{bucket}}/managedFolders/{{name}}
		{
			v0, ok0 := config["bucket"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/managedFolders/%s", v0, v1)
			}
		}
		// Format: {{bucket}}/{{name}}
		{
			v0, ok0 := config["bucket"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_storage_managed_folder_iam":
		// No standard import format found in documentation
		return ""
	case "google_storage_notification":
		// Format: {{bucket_name}}/notificationConfigs/{{id}}
		{
			v0, ok0 := config["bucket_name"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/notificationConfigs/%s", v0, v1)
			}
		}
		return ""
	case "google_storage_object_access_control":
		// Format: {{bucket}}/{{object}}/{{entity}}
		{
			v0, ok0 := config["bucket"].(string)
			v1, ok1 := config["object"].(string)
			v2, ok2 := config["entity"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_storage_object_acl":
		// No standard import format found in documentation
		return ""
	case "google_storage_transfer_agent_pool":
		// Format: projects/{{project}}/agentPools/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/agentPools/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_storage_transfer_job":
		// No standard import format found in documentation
		return ""
	case "google_tags_tag_binding":
		// Format: tagBindings/{{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("tagBindings/%s", v0)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_tags_tag_key":
		// Format: tagKeys/{{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("tagKeys/%s", v0)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_tags_tag_key_iam":
		// No standard import format found in documentation
		return ""
	case "google_tags_tag_value":
		// Format: tagValues/{{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("tagValues/%s", v0)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_tags_tag_value_iam":
		// No standard import format found in documentation
		return ""
	case "google_tpu_v2_queued_resource":
		// Format: projects/{{project}}/locations/{{zone}}/queuedResources/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/queuedResources/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{zone}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{zone}}/{{name}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_tpu_v2_vm":
		// Format: projects/{{project}}/locations/{{zone}}/nodes/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/nodes/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{zone}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["zone"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{zone}}/{{name}}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_transcoder_job":
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{project}} {{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s %s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_transcoder_job_template":
		// Format: projects/{{project}}/locations/{{location}}/jobTemplates/{{job_template_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["job_template_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/jobTemplates/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{job_template_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["job_template_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{job_template_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["job_template_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_vector_search_collection":
		// Format: projects/{{project}}/locations/{{location}}/collections/{{collection_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["collection_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/collections/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{collection_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["collection_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{collection_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["collection_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_vertex_ai_cache_config":
		// Format: projects/{{project}}/cacheConfig
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return fmt.Sprintf("projects/%s/cacheConfig", v0)
			}
		}
		// Format: {{project}}
		{
			v0, ok0 := config["project"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_vertex_ai_dataset":
		// No standard import format found in documentation
		return ""
	case "google_vertex_ai_deployment_resource_pool":
		// Format: projects/{{project}}/locations/{{region}}/deploymentResourcePools/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/deploymentResourcePools/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_vertex_ai_endpoint":
		// Format: projects/{{project}}/locations/{{location}}/endpoints/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_vertex_ai_endpoint_iam":
		// No standard import format found in documentation
		return ""
	case "google_vertex_ai_endpoint_with_model_garden_deployment":
		// No standard import format found in documentation
		return ""
	case "google_vertex_ai_feature_group":
		// Format: projects/{{project}}/locations/{{region}}/featureGroups/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/featureGroups/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_vertex_ai_feature_group_feature":
		// Format: projects/{{project}}/locations/{{region}}/featureGroups/{{feature_group}}/features/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["feature_group"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/featureGroups/%s/features/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{region}}/{{feature_group}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["feature_group"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{region}}/{{feature_group}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["feature_group"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{feature_group}}/{{name}}
		{
			v0, ok0 := config["feature_group"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_vertex_ai_feature_group_iam":
		// No standard import format found in documentation
		return ""
	case "google_vertex_ai_feature_online_store":
		// Format: projects/{{project}}/locations/{{region}}/featureOnlineStores/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/featureOnlineStores/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_vertex_ai_feature_online_store_featureview":
		// Format: projects/{{project}}/locations/{{region}}/featureOnlineStores/{{feature_online_store}}/featureViews/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["feature_online_store"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/featureOnlineStores/%s/featureViews/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{region}}/{{feature_online_store}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["feature_online_store"].(string)
			v3, ok3 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{region}}/{{feature_online_store}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["feature_online_store"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{feature_online_store}}/{{name}}
		{
			v0, ok0 := config["feature_online_store"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_vertex_ai_feature_online_store_featureview_iam":
		// No standard import format found in documentation
		return ""
	case "google_vertex_ai_feature_online_store_iam":
		// No standard import format found in documentation
		return ""
	case "google_vertex_ai_featurestore":
		// Format: projects/{{project}}/locations/{{region}}/featurestores/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_vertex_ai_featurestore_entitytype":
		// Format: {{featurestore}}/entityTypes/{{name}}
		{
			v0, ok0 := config["featurestore"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/entityTypes/%s", v0, v1)
			}
		}
		return ""
	case "google_vertex_ai_featurestore_entitytype_feature":
		// Format: {{entitytype}}/features/{{name}}
		{
			v0, ok0 := config["entitytype"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/features/%s", v0, v1)
			}
		}
		return ""
	case "google_vertex_ai_featurestore_entitytype_iam":
		// No standard import format found in documentation
		return ""
	case "google_vertex_ai_featurestore_iam":
		// No standard import format found in documentation
		return ""
	case "google_vertex_ai_index":
		// Format: projects/{{project}}/locations/{{region}}/indexes/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/indexes/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_vertex_ai_index_endpoint":
		// Format: projects/{{project}}/locations/{{region}}/indexEndpoints/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/indexEndpoints/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_vertex_ai_index_endpoint_deployed_index":
		// Format: projects/{{project}}/locations/{{region}}/indexEndpoints/{{index_endpoint}}/deployedIndex/{{deployed_index_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["index_endpoint"].(string)
			v3, ok3 := config["deployed_index_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/indexEndpoints/%s/deployedIndex/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{region}}/{{index_endpoint}}/{{deployed_index_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["index_endpoint"].(string)
			v3, ok3 := config["deployed_index_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{region}}/{{index_endpoint}}/{{deployed_index_id}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["index_endpoint"].(string)
			v2, ok2 := config["deployed_index_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{index_endpoint}}/{{deployed_index_id}}
		{
			v0, ok0 := config["index_endpoint"].(string)
			v1, ok1 := config["deployed_index_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_vertex_ai_metadata_store":
		// Format: projects/{{project}}/locations/{{region}}/metadataStores/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_vertex_ai_rag_engine_config":
		// Format: projects/{{project}}/locations/{{region}}/ragEngineConfig
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/ragEngineConfig", v0, v1)
			}
		}
		// Format: {{project}}/{{region}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{region}}
		{
			v0, ok0 := config["region"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_vertex_ai_reasoning_engine":
		// Format: projects/{{project}}/locations/{{region}}/reasoningEngines/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/reasoningEngines/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_vertex_ai_reasoning_engine_iam":
		// No standard import format found in documentation
		return ""
	case "google_vertex_ai_tensorboard":
		// Format: projects/{{project}}/locations/{{region}}/tensorboards/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_vmwareengine_cluster":
		// Format: {{parent}}/clusters/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/clusters/%s", v0, v1)
			}
		}
		return ""
	case "google_vmwareengine_datastore":
		// Format: projects/{{project}}/locations/{{location}}/datastores/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/datastores/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_vmwareengine_external_access_rule":
		// Format: {{parent}}/externalAccessRules/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/externalAccessRules/%s", v0, v1)
			}
		}
		return ""
	case "google_vmwareengine_external_address":
		// Format: {{parent}}/externalAddresses/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/externalAddresses/%s", v0, v1)
			}
		}
		return ""
	case "google_vmwareengine_network":
		// Format: projects/{{project}}/locations/{{location}}/vmwareEngineNetworks/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/vmwareEngineNetworks/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_vmwareengine_network_peering":
		// Format: projects/{{project}}/locations/global/networkPeerings/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("projects/%s/locations/global/networkPeerings/%s", v0, v1)
			}
		}
		// Format: {{project}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_vmwareengine_network_policy":
		// Format: projects/{{project}}/locations/{{location}}/networkPolicies/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_vmwareengine_private_cloud":
		// Format: projects/{{project}}/locations/{{location}}/privateClouds/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_vmwareengine_subnet":
		// Format: {{parent}}/subnets/{{name}}
		{
			v0, ok0 := config["parent"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/subnets/%s", v0, v1)
			}
		}
		return ""
	case "google_vpc_access_connector":
		// Format: projects/{{project}}/locations/{{region}}/connectors/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/connectors/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{region}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["region"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{region}}/{{name}}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		// Format: {{name}}
		{
			v0, ok0 := config["name"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "google_workbench_instance":
		// Format: projects/{{project}}/locations/{{location}}/instances/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/instances/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{name}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{name}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_workbench_instance_iam":
		// No standard import format found in documentation
		return ""
	case "google_workflows_workflow":
		// No standard import format found in documentation
		return ""
	case "google_workload_identity_service_agent":
		// No standard import format found in documentation
		return ""
	case "google_workstations_workstation":
		// Format: projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["workstation_cluster_id"].(string)
			v3, ok3 := config["workstation_config_id"].(string)
			v4, ok4 := config["workstation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s/workstations/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{project}}/{{location}}/{{workstation_cluster_id}}/{{workstation_config_id}}/{{workstation_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["workstation_cluster_id"].(string)
			v3, ok3 := config["workstation_config_id"].(string)
			v4, ok4 := config["workstation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" && ok4 && v4 != "" {
				return fmt.Sprintf("%s/%s/%s/%s/%s", v0, v1, v2, v3, v4)
			}
		}
		// Format: {{location}}/{{workstation_cluster_id}}/{{workstation_config_id}}/{{workstation_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["workstation_cluster_id"].(string)
			v2, ok2 := config["workstation_config_id"].(string)
			v3, ok3 := config["workstation_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "google_workstations_workstation_cluster":
		// Format: projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["workstation_cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s", v0, v1, v2)
			}
		}
		// Format: {{project}}/{{location}}/{{workstation_cluster_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["workstation_cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		// Format: {{location}}/{{workstation_cluster_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["workstation_cluster_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "google_workstations_workstation_config":
		// Format: projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["workstation_cluster_id"].(string)
			v3, ok3 := config["workstation_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{project}}/{{location}}/{{workstation_cluster_id}}/{{workstation_config_id}}
		{
			v0, ok0 := config["project"].(string)
			v1, ok1 := config["location"].(string)
			v2, ok2 := config["workstation_cluster_id"].(string)
			v3, ok3 := config["workstation_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		// Format: {{location}}/{{workstation_cluster_id}}/{{workstation_config_id}}
		{
			v0, ok0 := config["location"].(string)
			v1, ok1 := config["workstation_cluster_id"].(string)
			v2, ok2 := config["workstation_config_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "google_workstations_workstation_config_iam":
		// No standard import format found in documentation
		return ""
	case "google_workstations_workstation_iam":
		// No standard import format found in documentation
		return ""
	}
	return ""
}
