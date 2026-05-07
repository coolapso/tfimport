package providers

import (
	"fmt"
)

// extractScalewayImportID returns the necessary import ID for a scaleway resource
// based on its configuration extracted from the terraform plan.
func extractScalewayImportID(ctx *ProviderContext, resourceType string, config map[string]any) string {
	// First, check if there's a custom resolver for this resource
	if id := resolveCustomextractScalewayImportID(ctx, resourceType, config); id != "" {
		return id
	}

	switch resourceType {
	case "scaleway_account_project":
		// Format: {id}
		{
			v0, ok0 := config["id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "scaleway_account_ssh_key":
		// Format: {id}
		{
			v0, ok0 := config["id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "scaleway_apple_silicon_runner":
		// Format: {zone}/{id}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_apple_silicon_server":
		// Format: {zone}/{id}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_autoscaling_instance_group":
		// No standard import format found in documentation
		return ""
	case "scaleway_autoscaling_instance_policy":
		// No standard import format found in documentation
		return ""
	case "scaleway_autoscaling_instance_template":
		// No standard import format found in documentation
		return ""
	case "scaleway_baremetal_server":
		// Format: {zone}/{id}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_block_snapshot":
		// No standard import format found in documentation
		return ""
	case "scaleway_block_volume":
		// No standard import format found in documentation
		return ""
	case "scaleway_cockpit":
		// No standard import format found in documentation
		return ""
	case "scaleway_cockpit_alert_manager":
		// No standard import format found in documentation
		return ""
	case "scaleway_cockpit_exporter":
		// No standard import format found in documentation
		return ""
	case "scaleway_cockpit_grafana_user":
		// No standard import format found in documentation
		return ""
	case "scaleway_cockpit_source":
		// No standard import format found in documentation
		return ""
	case "scaleway_cockpit_token":
		// No standard import format found in documentation
		return ""
	case "scaleway_container":
		// No standard import format found in documentation
		return ""
	case "scaleway_container_cron":
		// No standard import format found in documentation
		return ""
	case "scaleway_container_domain":
		// No standard import format found in documentation
		return ""
	case "scaleway_container_namespace":
		// No standard import format found in documentation
		return ""
	case "scaleway_container_token":
		// No standard import format found in documentation
		return ""
	case "scaleway_container_trigger":
		// No standard import format found in documentation
		return ""
	case "scaleway_datawarehouse_database":
		// Format: {region}/{deployment_id}/{name}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["deployment_id"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "scaleway_datawarehouse_deployment":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_datawarehouse_user":
		// Format: {region}/{deployment_id}/{name}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["deployment_id"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "scaleway_domain_record":
		// No standard import format found in documentation
		return ""
	case "scaleway_domain_registration":
		// No standard import format found in documentation
		return ""
	case "scaleway_domain_zone":
		// No standard import format found in documentation
		return ""
	case "scaleway_edge_services_backend_stage":
		// Format: {id}
		{
			v0, ok0 := config["id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "scaleway_edge_services_cache_stage":
		// Format: {id}
		{
			v0, ok0 := config["id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "scaleway_edge_services_dns_stage":
		// Format: {id}
		{
			v0, ok0 := config["id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "scaleway_edge_services_head_stage":
		// Format: {id}
		{
			v0, ok0 := config["id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "scaleway_edge_services_pipeline":
		// Format: {id}
		{
			v0, ok0 := config["id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "scaleway_edge_services_plan":
		// No standard import format found in documentation
		return ""
	case "scaleway_edge_services_route_stage":
		// Format: {id}
		{
			v0, ok0 := config["id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "scaleway_edge_services_tls_stage":
		// Format: {id}
		{
			v0, ok0 := config["id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "scaleway_edge_services_waf_stage":
		// Format: {id}
		{
			v0, ok0 := config["id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "scaleway_file_filesystem":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_flexible_ip":
		// Format: {zone}/{id}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_flexible_ip_mac_address":
		// Format: {zone}/{id}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_function":
		// No standard import format found in documentation
		return ""
	case "scaleway_function_cron":
		// No standard import format found in documentation
		return ""
	case "scaleway_function_domain":
		// No standard import format found in documentation
		return ""
	case "scaleway_function_namespace":
		// No standard import format found in documentation
		return ""
	case "scaleway_function_token":
		// No standard import format found in documentation
		return ""
	case "scaleway_function_trigger":
		// No standard import format found in documentation
		return ""
	case "scaleway_iam_api_key":
		// Format: {id}
		{
			v0, ok0 := config["id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "scaleway_iam_application":
		// Format: {id}
		{
			v0, ok0 := config["id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "scaleway_iam_group":
		// Format: {id}
		{
			v0, ok0 := config["id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "scaleway_iam_group_membership":
		// No standard import format found in documentation
		return ""
	case "scaleway_iam_policy":
		// Format: {id}
		{
			v0, ok0 := config["id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "scaleway_iam_saml":
		// No standard import format found in documentation
		return ""
	case "scaleway_iam_saml_certificate":
		// No standard import format found in documentation
		return ""
	case "scaleway_iam_scim":
		// No standard import format found in documentation
		return ""
	case "scaleway_iam_ssh_key":
		// Format: {id}
		{
			v0, ok0 := config["id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "scaleway_iam_user":
		// Format: {id}
		{
			v0, ok0 := config["id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "scaleway_inference_deployment":
		// No standard import format found in documentation
		return ""
	case "scaleway_inference_model":
		// No standard import format found in documentation
		return ""
	case "scaleway_instance_image":
		// Format: {zone}/{id}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_instance_ip":
		// Format: {zone}/{id}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_instance_ip_reverse_dns":
		// Format: {zone}/{id}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_instance_placement_group":
		// Format: {zone}/{id}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_instance_private_nic":
		// Format: {zone}/{server_id}/{private_nic_id}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["server_id"].(string)
			v2, ok2 := config["private_nic_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "scaleway_instance_security_group":
		// Format: {zone}/{id}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_instance_security_group_rules":
		// Format: {zone}/{id}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_instance_server":
		// Format: {zone}/{id}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_instance_snapshot":
		// Format: {zone}/{id}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_instance_user_data":
		// Format: {zone}/{key}/{server_id}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["key"].(string)
			v2, ok2 := config["server_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "scaleway_instance_volume":
		// Format: {zone}/{id}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_interlink_link":
		// No standard import format found in documentation
		return ""
	case "scaleway_interlink_routing_policy":
		// No standard import format found in documentation
		return ""
	case "scaleway_iot_device":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_iot_hub":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_iot_network":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_iot_route":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_ipam_ip":
		// No standard import format found in documentation
		return ""
	case "scaleway_ipam_ip_reverse_dns":
		// No standard import format found in documentation
		return ""
	case "scaleway_job_definition":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_k8s_acl":
		// Format: {region}/{cluster-id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["cluster-id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_k8s_cluster":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_k8s_pool":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_kafka_cluster":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_key_manager_key":
		// No standard import format found in documentation
		return ""
	case "scaleway_lb":
		// No standard import format found in documentation
		return ""
	case "scaleway_lb_acl":
		// No standard import format found in documentation
		return ""
	case "scaleway_lb_backend":
		// No standard import format found in documentation
		return ""
	case "scaleway_lb_certificate":
		// Format: {zone}/{id}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_lb_frontend":
		// No standard import format found in documentation
		return ""
	case "scaleway_lb_ip":
		// No standard import format found in documentation
		return ""
	case "scaleway_lb_private_network":
		// No standard import format found in documentation
		return ""
	case "scaleway_lb_route":
		// No standard import format found in documentation
		return ""
	case "scaleway_mnq_nats_account":
		// No standard import format found in documentation
		return ""
	case "scaleway_mnq_nats_credentials":
		// No standard import format found in documentation
		return ""
	case "scaleway_mnq_sns":
		// No standard import format found in documentation
		return ""
	case "scaleway_mnq_sns_credentials":
		// No standard import format found in documentation
		return ""
	case "scaleway_mnq_sns_topic":
		// No standard import format found in documentation
		return ""
	case "scaleway_mnq_sns_topic_subscription":
		// No standard import format found in documentation
		return ""
	case "scaleway_mnq_sqs":
		// Format: {region}/{project_id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["project_id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_mnq_sqs_credentials":
		// No standard import format found in documentation
		return ""
	case "scaleway_mnq_sqs_queue":
		// No standard import format found in documentation
		return ""
	case "scaleway_mongodb_instance":
		// Format: {id}
		{
			v0, ok0 := config["id"].(string)
			if ok0 && v0 != "" {
				return v0
			}
		}
		return ""
	case "scaleway_mongodb_snapshot":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_mongodb_user":
		// Format: {region}/{instance_id}/{name}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["instance_id"].(string)
			v2, ok2 := config["name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "scaleway_object":
		// No standard import format found in documentation
		return ""
	case "scaleway_object_bucket":
		// No standard import format found in documentation
		return ""
	case "scaleway_object_bucket_acl":
		// No standard import format found in documentation
		return ""
	case "scaleway_object_bucket_lock_configuration":
		// No standard import format found in documentation
		return ""
	case "scaleway_object_bucket_policy":
		// No standard import format found in documentation
		return ""
	case "scaleway_object_bucket_server_side_encryption_configuration":
		// No standard import format found in documentation
		return ""
	case "scaleway_object_bucket_website_configuration":
		// No standard import format found in documentation
		return ""
	case "scaleway_opensearch_deployment":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_rdb_acl":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_rdb_database":
		// Format: {region}/{id}/{DBNAME}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			v2, ok2 := config["DBNAME"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" {
				return fmt.Sprintf("%s/%s/%s", v0, v1, v2)
			}
		}
		return ""
	case "scaleway_rdb_database_backup":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_rdb_instance":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_rdb_privilege":
		// Format: {region}/{instance_id}/{database_name}/{user_name}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["instance_id"].(string)
			v2, ok2 := config["database_name"].(string)
			v3, ok3 := config["user_name"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" && ok2 && v2 != "" && ok3 && v3 != "" {
				return fmt.Sprintf("%s/%s/%s/%s", v0, v1, v2, v3)
			}
		}
		return ""
	case "scaleway_rdb_read_replica":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_rdb_snapshot":
		// No standard import format found in documentation
		return ""
	case "scaleway_rdb_user":
		// No standard import format found in documentation
		return ""
	case "scaleway_redis_cluster":
		// Format: {zone}/{id}
		{
			v0, ok0 := config["zone"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_registry_namespace":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_s2s_vpn_connection":
		// No standard import format found in documentation
		return ""
	case "scaleway_s2s_vpn_customer_gateway":
		// No standard import format found in documentation
		return ""
	case "scaleway_s2s_vpn_gateway":
		// No standard import format found in documentation
		return ""
	case "scaleway_s2s_vpn_routing_policy":
		// No standard import format found in documentation
		return ""
	case "scaleway_sdb_sql_database":
		// No standard import format found in documentation
		return ""
	case "scaleway_secret":
		// No standard import format found in documentation
		return ""
	case "scaleway_secret_version":
		// No standard import format found in documentation
		return ""
	case "scaleway_tem_blocked_list":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_tem_domain":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_tem_domain_validation":
		// No standard import format found in documentation
		return ""
	case "scaleway_tem_webhook":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	case "scaleway_vpc":
		// No standard import format found in documentation
		return ""
	case "scaleway_vpc_acl":
		// No standard import format found in documentation
		return ""
	case "scaleway_vpc_connector":
		// No standard import format found in documentation
		return ""
	case "scaleway_vpc_gateway_network":
		// No standard import format found in documentation
		return ""
	case "scaleway_vpc_private_network":
		// No standard import format found in documentation
		return ""
	case "scaleway_vpc_public_gateway":
		// No standard import format found in documentation
		return ""
	case "scaleway_vpc_public_gateway_dhcp":
		// No standard import format found in documentation
		return ""
	case "scaleway_vpc_public_gateway_dhcp_reservation":
		// No standard import format found in documentation
		return ""
	case "scaleway_vpc_public_gateway_ip":
		// No standard import format found in documentation
		return ""
	case "scaleway_vpc_public_gateway_ip_reverse_dns":
		// No standard import format found in documentation
		return ""
	case "scaleway_vpc_public_gateway_pat_rule":
		// No standard import format found in documentation
		return ""
	case "scaleway_vpc_route":
		// No standard import format found in documentation
		return ""
	case "scaleway_webhosting":
		// Format: {region}/{id}
		{
			v0, ok0 := config["region"].(string)
			v1, ok1 := config["id"].(string)
			if ok0 && v0 != "" && ok1 && v1 != "" {
				return fmt.Sprintf("%s/%s", v0, v1)
			}
		}
		return ""
	}
	return ""
}
