package providers

import (
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type AWSClientContext struct {
	Config    *aws.Config
	AccountID string
	Region    string
	Partition string
	EC2Client *ec2.Client
	IAMClient *iam.Client
	SQSClient *sqs.Client
}

func (p *ProviderContext) GetAWSClient() *AWSClientContext {
	p.awsOnce.Do(func() {
		awsCfg, err := config.LoadDefaultConfig(p.Context)
		if err != nil {
			log.Printf("Warning: failed to load AWS config: %v", err)
			return
		}

		clientCtx := &AWSClientContext{
			Config:    &awsCfg,
			Region:    awsCfg.Region,
			Partition: "aws", // Default fallback
			EC2Client: ec2.NewFromConfig(awsCfg),
			IAMClient: iam.NewFromConfig(awsCfg),
			SQSClient: sqs.NewFromConfig(awsCfg),
		}

		if os.Getenv("AWS_PARTITION") != "" {
			clientCtx.Partition = os.Getenv("AWS_PARTITION")
		}

		if accID := os.Getenv("AWS_ACCOUNT_ID"); accID != "" {
			clientCtx.AccountID = accID
		} else {
			stsClient := sts.NewFromConfig(awsCfg)
			identity, err := stsClient.GetCallerIdentity(p.Context, &sts.GetCallerIdentityInput{})
			if err == nil && identity.Account != nil && identity.Arn != nil {
				clientCtx.AccountID = *identity.Account
				parts := strings.Split(*identity.Arn, ":")
				if len(parts) > 1 {
					clientCtx.Partition = parts[1]
				}
			} else {
				log.Printf("Warning: failed to get AWS Caller Identity: %v", err)
			}
		}

		p.awsClient = clientCtx
	})
	return p.awsClient
}

// extractAWSImportID returns the necessary import ID for a aws resource
// based on its configuration extracted from the terraform plan.
func extractAWSImportID(ctx *ProviderContext, resourceType string, config map[string]any) string {
	// First, check if there's a custom resolver for this resource
	if id := resolveCustomextractAWSImportID(ctx, resourceType, config); id != "" {
		return id
	}

	switch resourceType {
	case "aws_accessanalyzer_analyzer":
		if val, ok := config["analyzer_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_accessanalyzer_archive_rule":
		var parts []string
		if val, ok := config["analyzer_name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["rule_name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_account_alternate_contact":
		// No standard import format found in documentation
		return ""
	case "aws_account_primary_contact":
		if val, ok := config["account_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_account_region":
		// No standard import format found in documentation
		return ""
	case "aws_acm_certificate":
		// Computed ID format: their ARN
		return ""
	case "aws_acm_certificate_validation":
		// No standard import format found in documentation
		return ""
	case "aws_acmpca_certificate":
		// Computed ID format: their ARN
		return ""
	case "aws_acmpca_certificate_authority":
		// Computed ID format: the certificate authority ARN
		return ""
	case "aws_acmpca_certificate_authority_certificate":
		// No standard import format found in documentation
		return ""
	case "aws_acmpca_permission":
		// No standard import format found in documentation
		return ""
	case "aws_acmpca_policy":
		if val, ok := config["resource_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ami":
		// Computed ID format: the ID of the AMI
		return ""
	case "aws_ami_copy":
		// No standard import format found in documentation
		return ""
	case "aws_ami_from_instance":
		// No standard import format found in documentation
		return ""
	case "aws_ami_launch_permission":
		// Computed or complex ID format: `[ACCOUNT-ID|GROUP-NAME|ORGANIZATION-ARN|ORGANIZATIONAL-UNIT-ARN]/IMAGE-ID`
		return ""
	case "aws_amplify_app":
		// Computed ID format: Amplify App ID (appId)
		return ""
	case "aws_amplify_backend_environment":
		if val, ok := config["app_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_amplify_branch":
		if val, ok := config["app_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_amplify_domain_association":
		if val, ok := config["app_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_amplify_webhook":
		// Computed ID format: a webhook ID
		return ""
	case "aws_api_gateway_account":
		// Computed ID format: the account ID
		return ""
	case "aws_api_gateway_api_key":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_api_gateway_authorizer":
		// Computed or complex ID format: the `REST-API-ID/AUTHORIZER-ID`
		return ""
	case "aws_api_gateway_base_path_mapping":
		// Computed ID format: the domain name and base path or domain name, base path and domain name ID (for private custom domain names)
		return ""
	case "aws_api_gateway_client_certificate":
		// Computed ID format: the id
		return ""
	case "aws_api_gateway_deployment":
		// Computed or complex ID format: `REST-API-ID/DEPLOYMENT-ID`
		return ""
	case "aws_api_gateway_documentation_part":
		// Computed or complex ID format: `REST-API-ID/DOC-PART-ID`
		return ""
	case "aws_api_gateway_documentation_version":
		// Computed or complex ID format: `REST-API-ID/VERSION`
		return ""
	case "aws_api_gateway_domain_name":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_api_gateway_domain_name_access_association":
		// Computed or complex ID format: their `arn`
		return ""
	case "aws_api_gateway_gateway_response":
		// Computed or complex ID format: `REST-API-ID/RESPONSE-TYPE`
		return ""
	case "aws_api_gateway_integration":
		// Computed or complex ID format: `REST-API-ID/RESOURCE-ID/HTTP-METHOD`
		return ""
	case "aws_api_gateway_integration_response":
		// Computed or complex ID format: `REST-API-ID/RESOURCE-ID/HTTP-METHOD/STATUS-CODE`
		return ""
	case "aws_api_gateway_method":
		// Computed or complex ID format: `REST-API-ID/RESOURCE-ID/HTTP-METHOD`
		return ""
	case "aws_api_gateway_method_response":
		// Computed or complex ID format: `REST-API-ID/RESOURCE-ID/HTTP-METHOD/STATUS-CODE`
		return ""
	case "aws_api_gateway_method_settings":
		// Computed or complex ID format: `REST-API-ID/STAGE-NAME/METHOD-PATH`
		return ""
	case "aws_api_gateway_model":
		// Computed or complex ID format: `REST-API-ID/NAME`
		return ""
	case "aws_api_gateway_request_validator":
		// Computed or complex ID format: `REST-API-ID/REQUEST-VALIDATOR-ID`
		return ""
	case "aws_api_gateway_resource":
		// Computed or complex ID format: `REST-API-ID/RESOURCE-ID`
		return ""
	case "aws_api_gateway_rest_api":
		// Computed ID format: the REST API ID
		return ""
	case "aws_api_gateway_rest_api_policy":
		// Computed ID format: the REST API ID
		return ""
	case "aws_api_gateway_rest_api_put":
		if val, ok := config["rest_api_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_api_gateway_stage":
		// Computed or complex ID format: `REST-API-ID/STAGE-NAME`
		return ""
	case "aws_api_gateway_usage_plan":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_api_gateway_usage_plan_key":
		// Computed or complex ID format: the `USAGE-PLAN-ID/USAGE-PLAN-KEY-ID`
		return ""
	case "aws_api_gateway_vpc_link":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_apigatewayv2_api":
		// Computed ID format: the API identifier
		return ""
	case "aws_apigatewayv2_api_mapping":
		// Computed ID format: the API mapping identifier and domain name
		return ""
	case "aws_apigatewayv2_authorizer":
		// Computed ID format: the API identifier and authorizer identifier
		return ""
	case "aws_apigatewayv2_deployment":
		// Computed ID format: the API identifier and deployment identifier
		return ""
	case "aws_apigatewayv2_domain_name":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_apigatewayv2_integration":
		// Computed ID format: the API identifier and integration identifier
		return ""
	case "aws_apigatewayv2_integration_response":
		// Computed ID format: the API identifier, integration identifier and integration response identifier
		return ""
	case "aws_apigatewayv2_model":
		// Computed ID format: the API identifier and model identifier
		return ""
	case "aws_apigatewayv2_route":
		if val, ok := config["api_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_apigatewayv2_route_response":
		// Computed ID format: the API identifier, route identifier and route response identifier
		return ""
	case "aws_apigatewayv2_routing_rule":
		if val, ok := config["routing_rule_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_apigatewayv2_stage":
		// Computed ID format: the API identifier and stage name
		return ""
	case "aws_apigatewayv2_vpc_link":
		// Computed ID format: the VPC Link identifier
		return ""
	case "aws_app_cookie_stickiness_policy":
		// Computed or complex ID format: the ELB name, port, and policy name separated by colons (`:`)
		return ""
	case "aws_appautoscaling_policy":
		// Computed or complex ID format: the `service-namespace` , `resource-id`, `scalable-dimension` and `policy-name` separated by `/`
		return ""
	case "aws_appautoscaling_scheduled_action":
		// No standard import format found in documentation
		return ""
	case "aws_appautoscaling_target":
		// Computed or complex ID format: the `service-namespace` , `resource-id` and `scalable-dimension` separated by `/`
		return ""
	case "aws_appconfig_application":
		// Computed ID format: their application ID
		return ""
	case "aws_appconfig_configuration_profile":
		// Computed or complex ID format: the configuration profile ID and application ID separated by a colon (`:`)
		return ""
	case "aws_appconfig_deployment":
		// Computed or complex ID format: the application ID, environment ID, and deployment number separated by a slash (`/`)
		return ""
	case "aws_appconfig_deployment_strategy":
		// Computed ID format: their deployment strategy ID
		return ""
	case "aws_appconfig_environment":
		// Computed or complex ID format: the environment ID and application ID separated by a colon (`:`)
		return ""
	case "aws_appconfig_extension":
		// Computed ID format: their extension ID
		return ""
	case "aws_appconfig_extension_association":
		// Computed ID format: their extension association ID
		return ""
	case "aws_appconfig_hosted_configuration_version":
		// Computed or complex ID format: the application ID, configuration profile ID, and version number separated by a slash (`/`)
		return ""
	case "aws_appfabric_app_authorization":
		// No standard import format found in documentation
		return ""
	case "aws_appfabric_app_authorization_connection":
		// No standard import format found in documentation
		return ""
	case "aws_appfabric_app_bundle":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_appfabric_ingestion":
		if val, ok := config["app_bundle_identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appfabric_ingestion_destination":
		// No standard import format found in documentation
		return ""
	case "aws_appflow_connector_profile":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appflow_flow":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appintegrations_data_integration":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_appintegrations_event_integration":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_applicationinsights_application":
		if val, ok := config["resource_group_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appmesh_gateway_route":
		if val, ok := config["mesh_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appmesh_mesh":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appmesh_route":
		if val, ok := config["mesh_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appmesh_virtual_gateway":
		if val, ok := config["mesh_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appmesh_virtual_node":
		if val, ok := config["mesh_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appmesh_virtual_router":
		if val, ok := config["mesh_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appmesh_virtual_service":
		if val, ok := config["mesh_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_apprunner_auto_scaling_configuration_version":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_apprunner_connection":
		if val, ok := config["connection_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_apprunner_custom_domain_association":
		if val, ok := config["domain_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_apprunner_default_auto_scaling_configuration_version":
		// Unknown ID format: the current Region
		return ""
	case "aws_apprunner_deployment":
		// No standard import format found in documentation
		return ""
	case "aws_apprunner_observability_configuration":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_apprunner_service":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_apprunner_vpc_connector":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_apprunner_vpc_ingress_connection":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_appstream_directory_config":
		// Computed ID format: the id
		return ""
	case "aws_appstream_fleet":
		// Computed ID format: the id
		return ""
	case "aws_appstream_fleet_stack_association":
		if val, ok := config["fleet_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appstream_image_builder":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appstream_stack":
		// Computed ID format: the id
		return ""
	case "aws_appstream_user":
		if val, ok := config["user_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appstream_user_stack_association":
		if val, ok := config["user_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appsync_api":
		if val, ok := config["api_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appsync_api_cache":
		// Computed ID format: the AppSync API ID
		return ""
	case "aws_appsync_api_key":
		// Computed or complex ID format: the AppSync API ID and key separated by `:`
		return ""
	case "aws_appsync_channel_namespace":
		if val, ok := config["api_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appsync_datasource":
		if val, ok := config["api_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appsync_domain_name":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appsync_domain_name_api_association":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appsync_function":
		// Computed or complex ID format: the AppSync API ID and Function ID separated by `-`
		return ""
	case "aws_appsync_graphql_api":
		// Computed ID format: the GraphQL API ID
		return ""
	case "aws_appsync_resolver":
		if val, ok := config["api_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appsync_source_api_association":
		if val, ok := config["association_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_appsync_type":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_arcregionswitch_plan":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_arczonalshift_autoshift_observer_notification_status":
		// Unknown ID format: the AWS region
		return ""
	case "aws_athena_capacity_reservation":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_athena_data_catalog":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_athena_database":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_athena_named_query":
		// Computed ID format: the query ID
		return ""
	case "aws_athena_prepared_statement":
		// Computed or complex ID format: the `WORKGROUP-NAME/STATEMENT-NAME`
		return ""
	case "aws_athena_workgroup":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_auditmanager_account_registration":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_auditmanager_assessment":
		// Computed or complex ID format: the assessment `id`
		return ""
	case "aws_auditmanager_assessment_delegation":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_auditmanager_assessment_report":
		// Computed or complex ID format: the assessment report `id`
		return ""
	case "aws_auditmanager_control":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_auditmanager_framework":
		// Computed or complex ID format: the framework `id`
		return ""
	case "aws_auditmanager_framework_share":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_auditmanager_organization_admin_account_registration":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_autoscaling_attachment":
		// No standard import format found in documentation
		return ""
	case "aws_autoscaling_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_autoscaling_group_tag":
		// Computed or complex ID format: the ASG name and key, separated by a comma (`,`)
		return ""
	case "aws_autoscaling_lifecycle_hook":
		if val, ok := config["autoscaling_group_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_autoscaling_notification":
		// No standard import format found in documentation
		return ""
	case "aws_autoscaling_policy":
		if val, ok := config["autoscaling_group_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_autoscaling_schedule":
		if val, ok := config["autoscaling_group_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_autoscaling_traffic_source_attachment":
		if val, ok := config["autoscaling_group_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_autoscalingplans_scaling_plan":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_backup_framework":
		// Computed or complex ID format: the `id` which corresponds to the name of the Backup Framework
		return ""
	case "aws_backup_global_settings":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_backup_logically_air_gapped_vault":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_backup_plan":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_backup_region_settings":
		if val, ok := config["region"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_backup_report_plan":
		// Computed or complex ID format: the `id` which corresponds to the name of the Backup Report Plan
		return ""
	case "aws_backup_restore_testing_plan":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_backup_restore_testing_selection":
		// Computed or complex ID format: `name:restore_testing_plan_name`
		return ""
	case "aws_backup_selection":
		// Computed or complex ID format: the role plan_id and id separated by `|`
		return ""
	case "aws_backup_vault":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_backup_vault_lock_configuration":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_backup_vault_notifications":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_backup_vault_policy":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_batch_compute_environment":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_batch_job_definition":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_batch_job_queue":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_batch_scheduling_policy":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_bcmdataexports_export":
		// Computed ID format: the export ARN
		return ""
	case "aws_bedrock_custom_model":
		if val, ok := config["job_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_bedrock_guardrail":
		if val, ok := config["guardrail_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_bedrock_guardrail_version":
		if val, ok := config["guardrail_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_bedrock_inference_profile":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_bedrock_model_invocation_logging_configuration":
		// Computed or complex ID format: the `id` set to the AWS Region
		return ""
	case "aws_bedrock_provisioned_model_throughput":
		if val, ok := config["provisioned_model_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_bedrockagent_agent":
		// Computed ID format: the agent ID
		return ""
	case "aws_bedrockagent_agent_action_group":
		// No standard import format found in documentation
		return ""
	case "aws_bedrockagent_agent_alias":
		// Computed or complex ID format: the alias ID and the agent ID separated by `,`
		return ""
	case "aws_bedrockagent_agent_collaborator":
		if val, ok := config["agent_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_bedrockagent_agent_knowledge_base_association":
		// Computed or complex ID format: the agent ID, the agent version, and the knowledge base ID separated by `,`
		return ""
	case "aws_bedrockagent_data_source":
		// Computed ID format: the data source ID and the knowledge base ID
		return ""
	case "aws_bedrockagent_flow":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_bedrockagent_knowledge_base":
		// Computed ID format: the knowledge base ID
		return ""
	case "aws_bedrockagent_prompt":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_bedrockagentcore_agent_runtime":
		if val, ok := config["agent_runtime_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_bedrockagentcore_agent_runtime_endpoint":
		if val, ok := config["agent_runtime_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_bedrockagentcore_api_key_credential_provider":
		// Computed ID format: the provider name
		return ""
	case "aws_bedrockagentcore_browser":
		// Computed ID format: the browser ID
		return ""
	case "aws_bedrockagentcore_code_interpreter":
		// Computed ID format: the code interpreter ID
		return ""
	case "aws_bedrockagentcore_gateway":
		// Computed ID format: the gateway ID
		return ""
	case "aws_bedrockagentcore_gateway_target":
		// Computed ID format: the gateway identifier and target ID separated by a comma
		return ""
	case "aws_bedrockagentcore_memory":
		// Computed ID format: the memory ID
		return ""
	case "aws_bedrockagentcore_memory_strategy":
		var parts []string
		if val, ok := config["memory_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["strategy_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_bedrockagentcore_oauth2_credential_provider":
		// Computed ID format: the provider name
		return ""
	case "aws_bedrockagentcore_token_vault_cmk":
		// Computed ID format: the token vault ID
		return ""
	case "aws_bedrockagentcore_workload_identity":
		// Computed ID format: the workload identity name
		return ""
	case "aws_billing_view":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_budgets_budget":
		// Computed or complex ID format: `AccountID:BudgetName`
		return ""
	case "aws_budgets_budget_action":
		// Computed or complex ID format: `AccountID:ActionID:BudgetName`
		return ""
	case "aws_ce_anomaly_monitor":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ce_anomaly_subscription":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ce_cost_allocation_tag":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ce_cost_category":
		// Computed ID format: the id
		return ""
	case "aws_chatbot_slack_channel_configuration":
		if val, ok := config["chat_configuration_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_chatbot_teams_channel_configuration":
		if val, ok := config["team_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_chime_voice_connector":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_chime_voice_connector_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_chime_voice_connector_logging":
		if val, ok := config["voice_connector_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_chime_voice_connector_origination":
		if val, ok := config["voice_connector_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_chime_voice_connector_streaming":
		if val, ok := config["voice_connector_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_chime_voice_connector_termination":
		if val, ok := config["voice_connector_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_chime_voice_connector_termination_credentials":
		if val, ok := config["voice_connector_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_chimesdkmediapipelines_media_insights_pipeline_configuration":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_chimesdkvoice_global_settings":
		// Computed or complex ID format: the `id` (AWS account ID)
		return ""
	case "aws_chimesdkvoice_sip_media_application":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_chimesdkvoice_sip_rule":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_chimesdkvoice_voice_profile_domain":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cleanrooms_collaboration":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cleanrooms_configured_table":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cleanrooms_membership":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cloud9_environment_ec2":
		// No standard import format found in documentation
		return ""
	case "aws_cloud9_environment_membership":
		// Computed or complex ID format: the `environment-id#user-arn`
		return ""
	case "aws_cloudcontrolapi_resource":
		// No standard import format found in documentation
		return ""
	case "aws_cloudformation_stack":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudformation_stack_instances":
		if val, ok := config["call_as"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudformation_stack_set":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudformation_stack_set_instance":
		// Computed or complex ID format: the StackSet name, target AWS account ID, and target AWS Region separated by commas (`,`)
		return ""
	case "aws_cloudformation_type":
		// Computed ID format: the type version Amazon Resource Name (ARN)
		return ""
	case "aws_cloudfront_anycast_ip_list":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cloudfront_cache_policy":
		// Computed or complex ID format: the `id` of the cache policy
		return ""
	case "aws_cloudfront_connection_function":
		// Computed ID format: the function ID
		return ""
	case "aws_cloudfront_connection_group":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cloudfront_continuous_deployment_policy":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cloudfront_distribution":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cloudfront_distribution_tenant":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cloudfront_field_level_encryption_config":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cloudfront_field_level_encryption_profile":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cloudfront_function":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudfront_key_group":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cloudfront_key_value_store":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudfront_monitoring_subscription":
		// Computed ID format: the id
		return ""
	case "aws_cloudfront_multitenant_distribution":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cloudfront_origin_access_control":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cloudfront_origin_access_identity":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cloudfront_origin_request_policy":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cloudfront_public_key":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cloudfront_realtime_log_config":
		// Computed ID format: the ARN
		return ""
	case "aws_cloudfront_response_headers_policy":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cloudfront_trust_store":
		// Computed ID format: the trust store ID
		return ""
	case "aws_cloudfront_vpc_origin":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cloudfrontkeyvaluestore_key":
		if val, ok := config["key_value_store_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudfrontkeyvaluestore_keys_exclusive":
		if val, ok := config["key_value_store_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudhsm_v2_cluster":
		// Computed or complex ID format: the cluster `id`
		return ""
	case "aws_cloudhsm_v2_hsm":
		// Computed ID format: their HSM ID
		return ""
	case "aws_cloudsearch_domain":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudsearch_domain_service_access_policy":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudtrail":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_cloudtrail_event_data_store":
		// Computed or complex ID format: their `arn`
		return ""
	case "aws_cloudtrail_organization_delegated_admin_account":
		// Computed or complex ID format: the delegate account `id`
		return ""
	case "aws_cloudwatch_alarm_mute_rule":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_composite_alarm":
		if val, ok := config["alarm_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_contributor_insight_rule":
		if val, ok := config["rule_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_contributor_managed_insight_rule":
		if val, ok := config["resource_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_dashboard":
		if val, ok := config["dashboard_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_event_api_destination":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_event_archive":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_event_bus":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_event_bus_policy":
		if val, ok := config["event_bus_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_event_connection":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_event_endpoint":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_event_permission":
		var parts []string
		if val, ok := config["event_bus_name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["statement_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_cloudwatch_event_rule":
		var parts []string
		if val, ok := config["event_bus_name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["rule_name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_cloudwatch_event_target":
		// Computed or complex ID format: `event_bus_name/rule-name/target-id` (if you omit `event_bus_name`, the `default` event bus will be used)
		return ""
	case "aws_cloudwatch_log_account_policy":
		if val, ok := config["policy_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_log_anomaly_detector":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_cloudwatch_log_data_protection_policy":
		if val, ok := config["log_group_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_log_delivery":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cloudwatch_log_delivery_destination":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_log_delivery_destination_policy":
		if val, ok := config["delivery_destination_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_log_delivery_source":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_log_destination":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_log_destination_policy":
		if val, ok := config["destination_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_log_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_log_index_policy":
		if val, ok := config["log_group_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_log_metric_filter":
		if val, ok := config["log_group_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_log_resource_policy":
		// Computed ID format: the policy name for account-scoped policies, or the ARN of the CloudWatch Logs resource to which the policy is attached for resource-scoped policies
		return ""
	case "aws_cloudwatch_log_stream":
		if val, ok := config["log_group_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_log_subscription_filter":
		if val, ok := config["log_group_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_log_transformer":
		if val, ok := config["log_group_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_metric_alarm":
		if val, ok := config["alarm_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_metric_stream":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cloudwatch_otel_enrichment":
		// Unknown ID format: the region
		return ""
	case "aws_cloudwatch_query_definition":
		// Computed ID format: the query definition ARN. The ARN can be found on the "Edit Query" page for the query in the AWS Console
		return ""
	case "aws_codeartifact_domain":
		// Computed ID format: the CodeArtifact Domain arn
		return ""
	case "aws_codeartifact_domain_permissions_policy":
		// Computed ID format: the CodeArtifact Domain ARN
		return ""
	case "aws_codeartifact_repository":
		// Computed ID format: the CodeArtifact Repository ARN
		return ""
	case "aws_codeartifact_repository_permissions_policy":
		// Computed ID format: the CodeArtifact Repository ARN
		return ""
	case "aws_codebuild_fleet":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_codebuild_project":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_codebuild_report_group":
		// Computed ID format: the CodeBuild Report Group arn
		return ""
	case "aws_codebuild_resource_policy":
		// Computed ID format: the CodeBuild Resource Policy arn
		return ""
	case "aws_codebuild_source_credential":
		// Computed ID format: the CodeBuild Source Credential arn
		return ""
	case "aws_codebuild_webhook":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_codecatalyst_dev_environment":
		// No standard import format found in documentation
		return ""
	case "aws_codecatalyst_project":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_codecatalyst_source_repository":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_codecommit_approval_rule_template":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_codecommit_approval_rule_template_association":
		if val, ok := config["approval_rule_template_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_codecommit_repository":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_codecommit_trigger":
		// No standard import format found in documentation
		return ""
	case "aws_codeconnections_connection":
		// Computed ID format: the ARN
		return ""
	case "aws_codeconnections_host":
		// Computed ID format: the ARN
		return ""
	case "aws_codedeploy_app":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_codedeploy_deployment_config":
		if val, ok := config["deployment_config_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_codedeploy_deployment_group":
		if val, ok := config["app_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_codeguruprofiler_profiling_group":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_codegurureviewer_repository_association":
		// No standard import format found in documentation
		return ""
	case "aws_codepipeline":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_codepipeline_custom_action_type":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_codepipeline_webhook":
		// Computed ID format: their ARN
		return ""
	case "aws_codestarconnections_connection":
		// Computed ID format: the ARN
		return ""
	case "aws_codestarconnections_host":
		// Computed ID format: the ARN
		return ""
	case "aws_codestarnotifications_notification_rule":
		// Computed ID format: the ARN
		return ""
	case "aws_cognito_identity_pool":
		// Computed ID format: its ID
		return ""
	case "aws_cognito_identity_pool_provider_principal_tag":
		// Computed ID format: the Identity Pool ID and provider name
		return ""
	case "aws_cognito_identity_pool_roles_attachment":
		// Computed ID format: the Identity Pool ID
		return ""
	case "aws_cognito_identity_provider":
		// Computed ID format: their User Pool ID and Provider Name
		return ""
	case "aws_cognito_log_delivery_configuration":
		if val, ok := config["user_pool_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cognito_managed_login_branding":
		if val, ok := config["user_pool_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cognito_managed_user_pool_client":
		// Computed or complex ID format: the `id` of the Cognito User Pool and the `id` of the Cognito User Pool Client
		return ""
	case "aws_cognito_resource_server":
		// Computed ID format: their User Pool ID and Identifier
		return ""
	case "aws_cognito_risk_configuration":
		// No standard import format found in documentation
		return ""
	case "aws_cognito_user":
		if val, ok := config["user_pool_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cognito_user_group":
		if val, ok := config["user_pool_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cognito_user_in_group":
		if val, ok := config["user_pool_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cognito_user_pool":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cognito_user_pool_client":
		// Computed or complex ID format: the `id` of the Cognito User Pool, and the `id` of the Cognito User Pool Client
		return ""
	case "aws_cognito_user_pool_domain":
		if val, ok := config["domain"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_cognito_user_pool_ui_customization":
		if val, ok := config["user_pool_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_comprehend_document_classifier":
		// Computed ID format: the ARN
		return ""
	case "aws_comprehend_entity_recognizer":
		// Computed ID format: the ARN
		return ""
	case "aws_computeoptimizer_enrollment_status":
		// Computed ID format: the account ID
		return ""
	case "aws_computeoptimizer_recommendation_preferences":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_config_aggregate_authorization":
		// Computed or complex ID format: `account_id:authorized_aws_region`
		return ""
	case "aws_config_config_rule":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_config_configuration_aggregator":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_config_configuration_recorder":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_config_configuration_recorder_status":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_config_conformance_pack":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_config_delivery_channel":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_config_organization_conformance_pack":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_config_organization_custom_policy_rule":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_config_organization_custom_rule":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_config_organization_managed_rule":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_config_remediation_configuration":
		if val, ok := config["config_rule_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_config_retention_configuration":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_connect_bot_association":
		// Computed or complex ID format: the Amazon Connect instance ID, Lex (V1) bot name, and Lex (V1) bot region separated by colons (`:`)
		return ""
	case "aws_connect_contact_flow":
		if val, ok := config["instance_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_connect_contact_flow_module":
		if val, ok := config["instance_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_connect_hours_of_operation":
		if val, ok := config["instance_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_connect_instance":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_connect_instance_storage_config":
		if val, ok := config["instance_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_connect_lambda_function_association":
		if val, ok := config["instance_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_connect_phone_number":
		// Computed or complex ID format: its `id`
		return ""
	case "aws_connect_phone_number_contact_flow_association":
		if val, ok := config["phone_number_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_connect_queue":
		if val, ok := config["instance_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_connect_quick_connect":
		if val, ok := config["instance_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_connect_routing_profile":
		if val, ok := config["instance_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_connect_security_profile":
		if val, ok := config["instance_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_connect_user":
		if val, ok := config["instance_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_connect_user_hierarchy_group":
		if val, ok := config["instance_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_connect_user_hierarchy_structure":
		if val, ok := config["instance_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_connect_vocabulary":
		if val, ok := config["instance_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_controltower_baseline":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_controltower_control":
		var parts []string
		if val, ok := config["organizational_unit_arn"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["control_identifier"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_controltower_landing_zone":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_costoptimizationhub_enrollment_status":
		// Computed ID format: your AWS account ID
		return ""
	case "aws_costoptimizationhub_preferences":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_cur_report_definition":
		if val, ok := config["report_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_customer_gateway":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_customerprofiles_domain":
		// Computed or complex ID format: the resource `id`
		return ""
	case "aws_customerprofiles_profile":
		// Computed or complex ID format: the resource `id`
		return ""
	case "aws_dataexchange_data_set":
		// Computed or complex ID format: their `id`
		return ""
	case "aws_dataexchange_event_action":
		// Computed ID format: the id
		return ""
	case "aws_dataexchange_revision":
		// Computed or complex ID format: their `data-set-id:revision-id`
		return ""
	case "aws_dataexchange_revision_assets":
		// No standard import format found in documentation
		return ""
	case "aws_datapipeline_pipeline":
		// Computed ID format: the id (Pipeline ID)
		return ""
	case "aws_datapipeline_pipeline_definition":
		// Computed ID format: the id
		return ""
	case "aws_datasync_agent":
		// Computed ID format: the DataSync Agent Amazon Resource Name (ARN)
		return ""
	case "aws_datasync_location_azure_blob":
		// Computed ID format: the Amazon Resource Name (ARN)
		return ""
	case "aws_datasync_location_efs":
		// Computed ID format: the DataSync Task Amazon Resource Name (ARN)
		return ""
	case "aws_datasync_location_fsx_lustre_file_system":
		// Computed or complex ID format: the `DataSync-ARN#FSx-Lustre-ARN`
		return ""
	case "aws_datasync_location_fsx_ontap_file_system":
		// Computed or complex ID format: the `DataSync-ARN#FSx-ontap-svm-ARN`
		return ""
	case "aws_datasync_location_fsx_openzfs_file_system":
		// Computed or complex ID format: the `DataSync-ARN#FSx-openzfs-ARN`
		return ""
	case "aws_datasync_location_fsx_windows_file_system":
		// Computed or complex ID format: the `DataSync-ARN#FSx-Windows-ARN`
		return ""
	case "aws_datasync_location_hdfs":
		// Computed ID format: the Amazon Resource Name (ARN)
		return ""
	case "aws_datasync_location_nfs":
		// Computed ID format: the DataSync Task Amazon Resource Name (ARN)
		return ""
	case "aws_datasync_location_object_storage":
		// Computed ID format: the Amazon Resource Name (ARN)
		return ""
	case "aws_datasync_location_s3":
		// Computed ID format: the DataSync Task Amazon Resource Name (ARN)
		return ""
	case "aws_datasync_location_smb":
		// Computed ID format: the Amazon Resource Name (ARN)
		return ""
	case "aws_datasync_task":
		// Computed ID format: the DataSync Task Amazon Resource Name (ARN)
		return ""
	case "aws_datazone_asset_type":
		var parts []string
		if val, ok := config["domain_identifier"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_datazone_domain":
		if val, ok := config["domain_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_datazone_environment":
		// Computed or complex ID format: the `domain_idntifier,id`
		return ""
	case "aws_datazone_environment_blueprint_configuration":
		if val, ok := config["domain_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_datazone_environment_profile":
		// Computed or complex ID format: a comma-delimited string combining `id` and `domain_identifier`
		return ""
	case "aws_datazone_form_type":
		if val, ok := config["domain_identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_datazone_glossary":
		// Computed ID format: the import Datazone Glossary using a comma-delimited string combining the domain id, glossary id, and the id of the project it's under
		return ""
	case "aws_datazone_glossary_term":
		if val, ok := config["domain_identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_datazone_project":
		if val, ok := config["domain_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_datazone_user_profile":
		var parts []string
		if val, ok := config["user_identifier"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["domain_identifier"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["type"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_dax_cluster":
		if val, ok := config["cluster_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_dax_parameter_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_dax_subnet_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_db_cluster_snapshot":
		// Computed ID format: the cluster snapshot identifier
		return ""
	case "aws_db_event_subscription":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_db_instance":
		if val, ok := config["identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_db_instance_automated_backups_replication":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_db_instance_role_association":
		// Computed or complex ID format: the DB Instance Identifier and IAM Role ARN separated by a comma (`,`)
		return ""
	case "aws_db_option_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_db_parameter_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_db_proxy":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_db_proxy_default_target_group":
		if val, ok := config["db_proxy_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_db_proxy_endpoint":
		// Computed or complex ID format: the `DB-PROXY-NAME/DB-PROXY-ENDPOINT-NAME`
		return ""
	case "aws_db_proxy_target":
		// No standard import format found in documentation
		return ""
	case "aws_db_snapshot":
		// Computed ID format: the snapshot identifier
		return ""
	case "aws_db_snapshot_copy":
		// Computed ID format: the snapshot identifier
		return ""
	case "aws_db_subnet_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_default_network_acl":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_default_route_table":
		if val, ok := config["vpc_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_default_security_group":
		// Computed or complex ID format: the security group `id`
		return ""
	case "aws_default_subnet":
		// Computed or complex ID format: the subnet `id`
		return ""
	case "aws_default_vpc":
		// Computed or complex ID format: the VPC `id`
		return ""
	case "aws_default_vpc_dhcp_options":
		// Computed or complex ID format: the DHCP Options `id`
		return ""
	case "aws_detective_graph":
		// Computed ID format: the ARN
		return ""
	case "aws_detective_invitation_accepter":
		// Computed ID format: the graph ARN
		return ""
	case "aws_detective_member":
		// Computed ID format: the ARN of the graph followed by the account ID of the member account
		return ""
	case "aws_detective_organization_admin_account":
		if val, ok := config["account_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_detective_organization_configuration":
		// Computed ID format: the behavior graph ARN
		return ""
	case "aws_devicefarm_device_pool":
		// Computed ID format: their ARN
		return ""
	case "aws_devicefarm_instance_profile":
		// Computed ID format: their ARN
		return ""
	case "aws_devicefarm_network_profile":
		// Computed ID format: their ARN
		return ""
	case "aws_devicefarm_project":
		// Computed ID format: their ARN
		return ""
	case "aws_devicefarm_test_grid_project":
		// Computed ID format: their ARN
		return ""
	case "aws_devicefarm_upload":
		// Computed ID format: their ARN
		return ""
	case "aws_devopsguru_event_sources_config":
		// Unknown ID format: the region
		return ""
	case "aws_devopsguru_notification_channel":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_devopsguru_resource_collection":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_devopsguru_service_integration":
		// Unknown ID format: the region
		return ""
	case "aws_directory_service_conditional_forwarder":
		// Computed ID format: the directory id and remote_domain_name
		return ""
	case "aws_directory_service_directory":
		// Computed or complex ID format: the directory `id`
		return ""
	case "aws_directory_service_log_subscription":
		// Computed ID format: the directory id
		return ""
	case "aws_directory_service_radius_settings":
		// Computed ID format: the directory ID
		return ""
	case "aws_directory_service_region":
		// Computed ID format: directory ID,Region name
		return ""
	case "aws_directory_service_shared_directory":
		// Computed ID format: the owner directory ID/shared directory ID
		return ""
	case "aws_directory_service_shared_directory_accepter":
		// Computed ID format: the shared directory ID
		return ""
	case "aws_directory_service_trust":
		// Computed or complex ID format: the directory ID and remote domain name, separated by a `/`
		return ""
	case "aws_dlm_lifecycle_policy":
		// Computed ID format: their policy ID
		return ""
	case "aws_dms_certificate":
		if val, ok := config["certificate_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_dms_endpoint":
		if val, ok := config["endpoint_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_dms_event_subscription":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_dms_replication_config":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_dms_replication_instance":
		if val, ok := config["replication_instance_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_dms_replication_subnet_group":
		if val, ok := config["replication_subnet_group_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_dms_replication_task":
		if val, ok := config["replication_task_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_dms_s3_endpoint":
		if val, ok := config["endpoint_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_docdb_cluster":
		if val, ok := config["cluster_identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_docdb_cluster_instance":
		if val, ok := config["identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_docdb_cluster_parameter_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_docdb_cluster_snapshot":
		// Computed ID format: the cluster snapshot identifier
		return ""
	case "aws_docdb_event_subscription":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_docdb_global_cluster":
		// Computed ID format: the Global Cluster identifier
		return ""
	case "aws_docdb_subnet_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_docdbelastic_cluster":
		// Computed or complex ID format: the `arn` argument. For example,
		return ""
	case "aws_drs_replication_configuration_template":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_dsql_cluster":
		if val, ok := config["identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_dsql_cluster_peering":
		if val, ok := config["identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_dx_bgp_peer":
		// No standard import format found in documentation
		return ""
	case "aws_dx_connection":
		// Computed or complex ID format: the connection `id`
		return ""
	case "aws_dx_connection_association":
		// No standard import format found in documentation
		return ""
	case "aws_dx_connection_confirmation":
		// No standard import format found in documentation
		return ""
	case "aws_dx_gateway":
		// Computed or complex ID format: the gateway `id`
		return ""
	case "aws_dx_gateway_association":
		if val, ok := config["dx_gateway_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_dx_gateway_association_proposal":
		// No standard import format found in documentation
		return ""
	case "aws_dx_hosted_connection":
		// No standard import format found in documentation
		return ""
	case "aws_dx_hosted_private_virtual_interface":
		// Computed or complex ID format: the VIF `id`
		return ""
	case "aws_dx_hosted_private_virtual_interface_accepter":
		// Computed or complex ID format: the VIF `id`
		return ""
	case "aws_dx_hosted_public_virtual_interface":
		// Computed or complex ID format: the VIF `id`
		return ""
	case "aws_dx_hosted_public_virtual_interface_accepter":
		// Computed or complex ID format: the VIF `id`
		return ""
	case "aws_dx_hosted_transit_virtual_interface":
		// Computed or complex ID format: the VIF `id`
		return ""
	case "aws_dx_hosted_transit_virtual_interface_accepter":
		// Computed or complex ID format: the VIF `id`
		return ""
	case "aws_dx_lag":
		// Computed or complex ID format: the LAG `id`
		return ""
	case "aws_dx_macsec_key_association":
		// No standard import format found in documentation
		return ""
	case "aws_dx_private_virtual_interface":
		// Computed or complex ID format: the VIF `id`
		return ""
	case "aws_dx_public_virtual_interface":
		// Computed or complex ID format: the VIF `id`
		return ""
	case "aws_dx_transit_virtual_interface":
		// Computed or complex ID format: the VIF `id`
		return ""
	case "aws_dynamodb_contributor_insights":
		// Computed or complex ID format: the format `name:table_name/index:index_name`, followed by the account number
		return ""
	case "aws_dynamodb_global_secondary_index":
		if val, ok := config["table_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_dynamodb_global_table":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_dynamodb_kinesis_streaming_destination":
		if val, ok := config["table_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_dynamodb_resource_policy":
		if val, ok := config["resource_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_dynamodb_table":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_dynamodb_table_export":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_dynamodb_table_item":
		// No standard import format found in documentation
		return ""
	case "aws_dynamodb_table_replica":
		// Computed or complex ID format: the `table-name:main-region`
		return ""
	case "aws_dynamodb_tag":
		// Computed or complex ID format: the DynamoDB resource identifier and key, separated by a comma (`,`)
		return ""
	case "aws_ebs_default_kms_key":
		// Computed ID format: the KMS key ARN
		return ""
	case "aws_ebs_encryption_by_default":
		// No standard import format found in documentation
		return ""
	case "aws_ebs_fast_snapshot_restore":
		if val, ok := config["availability_zone"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ebs_snapshot":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ebs_snapshot_block_public_access":
		// No standard import format found in documentation
		return ""
	case "aws_ebs_snapshot_copy":
		// No standard import format found in documentation
		return ""
	case "aws_ebs_snapshot_import":
		// No standard import format found in documentation
		return ""
	case "aws_ebs_volume":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ebs_volume_copy":
		// Computed ID format: the volume ID
		return ""
	case "aws_ec2_allowed_images_settings":
		// No standard import format found in documentation
		return ""
	case "aws_ec2_availability_zone_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ec2_capacity_block_reservation":
		// No standard import format found in documentation
		return ""
	case "aws_ec2_capacity_reservation":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ec2_carrier_gateway":
		// Computed ID format: the carrier gateway's ID
		return ""
	case "aws_ec2_client_vpn_authorization_rule":
		// No standard import format found in documentation
		return ""
	case "aws_ec2_client_vpn_endpoint":
		// Computed or complex ID format: the `id` value found via `aws ec2 describe-client-vpn-endpoints`
		return ""
	case "aws_ec2_client_vpn_network_association":
		// Computed or complex ID format: the endpoint ID and the association ID. Values are separated by a `,`
		return ""
	case "aws_ec2_client_vpn_route":
		// Computed or complex ID format: the endpoint ID, target subnet ID, and destination CIDR block. All values are separated by a `,`
		return ""
	case "aws_ec2_default_credit_specification":
		if val, ok := config["instance_family"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ec2_fleet":
		// Computed ID format: the Fleet identifier
		return ""
	case "aws_ec2_host":
		// Computed or complex ID format: the host `id`
		return ""
	case "aws_ec2_image_block_public_access":
		// No standard import format found in documentation
		return ""
	case "aws_ec2_instance_connect_endpoint":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ec2_instance_metadata_defaults":
		// No standard import format found in documentation
		return ""
	case "aws_ec2_instance_state":
		if val, ok := config["instance_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ec2_local_gateway_route":
		if val, ok := config["_"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ec2_local_gateway_route_table_vpc_association":
		// Computed ID format: the Local Gateway Route Table VPC Association identifier
		return ""
	case "aws_ec2_managed_prefix_list":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ec2_managed_prefix_list_entry":
		if val, ok := config["prefix_list_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ec2_network_insights_access_scope":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ec2_network_insights_analysis":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ec2_network_insights_path":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ec2_secondary_network":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ec2_secondary_subnet":
		// Computed ID format: the secondary subnet ID
		return ""
	case "aws_ec2_serial_console_access":
		// No standard import format found in documentation
		return ""
	case "aws_ec2_subnet_cidr_reservation":
		// Computed or complex ID format: `SUBNET_ID:RESERVATION_ID`
		return ""
	case "aws_ec2_tag":
		// Computed or complex ID format: the EC2 resource identifier and key, separated by a comma (`,`)
		return ""
	case "aws_ec2_traffic_mirror_filter":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ec2_traffic_mirror_filter_rule":
		if val, ok := config["traffic_mirror_filter_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ec2_traffic_mirror_session":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ec2_traffic_mirror_target":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ec2_transit_gateway":
		// Computed ID format: the EC2 Transit Gateway identifier
		return ""
	case "aws_ec2_transit_gateway_connect":
		// Computed ID format: the EC2 Transit Gateway Connect identifier
		return ""
	case "aws_ec2_transit_gateway_connect_peer":
		// Computed ID format: the EC2 Transit Gateway Connect Peer identifier
		return ""
	case "aws_ec2_transit_gateway_default_route_table_association":
		// No standard import format found in documentation
		return ""
	case "aws_ec2_transit_gateway_default_route_table_propagation":
		// No standard import format found in documentation
		return ""
	case "aws_ec2_transit_gateway_metering_policy":
		if val, ok := config["transit_gateway_metering_policy_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ec2_transit_gateway_metering_policy_entry":
		// Computed ID format: the composite identifier
		return ""
	case "aws_ec2_transit_gateway_multicast_domain":
		// Computed ID format: the EC2 Transit Gateway Multicast Domain identifier
		return ""
	case "aws_ec2_transit_gateway_multicast_domain_association":
		// No standard import format found in documentation
		return ""
	case "aws_ec2_transit_gateway_multicast_group_member":
		// No standard import format found in documentation
		return ""
	case "aws_ec2_transit_gateway_multicast_group_source":
		// No standard import format found in documentation
		return ""
	case "aws_ec2_transit_gateway_peering_attachment":
		// Computed ID format: the EC2 Transit Gateway Attachment identifier
		return ""
	case "aws_ec2_transit_gateway_peering_attachment_accepter":
		// Computed ID format: the EC2 Transit Gateway Attachment identifier
		return ""
	case "aws_ec2_transit_gateway_policy_table":
		// Computed ID format: the EC2 Transit Gateway Policy Table identifier
		return ""
	case "aws_ec2_transit_gateway_policy_table_association":
		// Computed ID format: the EC2 Transit Gateway Policy Table identifier, an underscore, and the EC2 Transit Gateway Attachment identifier
		return ""
	case "aws_ec2_transit_gateway_prefix_list_reference":
		if val, ok := config["_"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ec2_transit_gateway_route":
		// Unknown ID format: the EC2 Transit Gateway Route Table, an underscore, and the destination
		return ""
	case "aws_ec2_transit_gateway_route_table":
		// Computed ID format: the EC2 Transit Gateway Route Table identifier
		return ""
	case "aws_ec2_transit_gateway_route_table_association":
		// Computed ID format: the EC2 Transit Gateway Route Table identifier, an underscore, and the EC2 Transit Gateway Attachment identifier
		return ""
	case "aws_ec2_transit_gateway_route_table_propagation":
		// Computed ID format: the EC2 Transit Gateway Route Table identifier, an underscore, and the EC2 Transit Gateway Attachment identifier
		return ""
	case "aws_ec2_transit_gateway_vpc_attachment":
		// Computed ID format: the EC2 Transit Gateway Attachment identifier
		return ""
	case "aws_ec2_transit_gateway_vpc_attachment_accepter":
		// Computed ID format: the EC2 Transit Gateway Attachment identifier
		return ""
	case "aws_ecr_account_setting":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ecr_lifecycle_policy":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ecr_pull_through_cache_rule":
		if val, ok := config["ecr_repository_prefix"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ecr_pull_time_update_exclusion":
		if val, ok := config["principal_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ecr_registry_policy":
		// Computed ID format: the registry id
		return ""
	case "aws_ecr_registry_scanning_configuration":
		if val, ok := config["registry_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ecr_replication_configuration":
		if val, ok := config["registry_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ecr_repository":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ecr_repository_creation_template":
		if val, ok := config["prefix"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ecr_repository_policy":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ecrpublic_repository":
		if val, ok := config["repository_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ecrpublic_repository_policy":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ecs_account_setting_default":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ecs_capacity_provider":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_ecs_cluster":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ecs_cluster_capacity_providers":
		if val, ok := config["cluster_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ecs_express_gateway_service":
		// Computed ID format: the service ARN
		return ""
	case "aws_ecs_service":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ecs_tag":
		// Computed or complex ID format: the ECS resource identifier and key, separated by a comma (`,`)
		return ""
	case "aws_ecs_task_definition":
		// Computed ID format: their ARNs
		return ""
	case "aws_ecs_task_set":
		if val, ok := config["task_set_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_efs_access_point":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_efs_backup_policy":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_efs_file_system":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_efs_file_system_policy":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_efs_mount_target":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_efs_replication_configuration":
		if val, ok := config["availability_zone_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_egress_only_internet_gateway":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_eip":
		// Computed ID format: their Allocation ID
		return ""
	case "aws_eip_association":
		// Computed ID format: their association IDs
		return ""
	case "aws_eip_domain_name":
		// Computed ID format: their association IDs
		return ""
	case "aws_eks_access_entry":
		if val, ok := config["cluster_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_eks_access_policy_association":
		if val, ok := config["cluster_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_eks_addon":
		if val, ok := config["cluster_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_eks_capability":
		if val, ok := config["cluster_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_eks_cluster":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_eks_fargate_profile":
		if val, ok := config["cluster_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_eks_identity_provider_config":
		if val, ok := config["cluster_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_eks_node_group":
		if val, ok := config["cluster_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_eks_pod_identity_association":
		if val, ok := config["cluster_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_elastic_beanstalk_application":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_elastic_beanstalk_application_version":
		// No standard import format found in documentation
		return ""
	case "aws_elastic_beanstalk_configuration_template":
		// No standard import format found in documentation
		return ""
	case "aws_elastic_beanstalk_environment":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_elasticache_cluster":
		if val, ok := config["cluster_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_elasticache_global_replication_group":
		if val, ok := config["global_replication_group_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_elasticache_parameter_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_elasticache_replication_group":
		if val, ok := config["replication_group_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_elasticache_reserved_cache_node":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_elasticache_serverless_cache":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_elasticache_subnet_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_elasticache_user":
		if val, ok := config["user_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_elasticache_user_group":
		if val, ok := config["user_group_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_elasticache_user_group_association":
		if val, ok := config["user_group_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_elasticsearch_domain":
		if val, ok := config["domain_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_elasticsearch_domain_policy":
		// No standard import format found in documentation
		return ""
	case "aws_elasticsearch_domain_saml_options":
		if val, ok := config["domain_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_elasticsearch_vpc_endpoint":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_elastictranscoder_pipeline":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_elastictranscoder_preset":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_elb":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_elb_attachment":
		// No standard import format found in documentation
		return ""
	case "aws_emr_block_public_access_configuration":
		// No standard import format found in documentation
		return ""
	case "aws_emr_cluster":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_emr_instance_fleet":
		// Computed or complex ID format: the EMR Cluster identifier and Instance Fleet identifier separated by a forward slash (`/`)
		return ""
	case "aws_emr_instance_group":
		// Computed or complex ID format: their EMR Cluster id and Instance Group id separated by a forward-slash `/`
		return ""
	case "aws_emr_managed_scaling_policy":
		// Computed ID format: the EMR Cluster identifier
		return ""
	case "aws_emr_security_configuration":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_emr_studio":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_emr_studio_session_mapping":
		// Computed or complex ID format: `studio-id:identity-type:identity-id`
		return ""
	case "aws_emrcontainers_job_template":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_emrcontainers_virtual_cluster":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_emrserverless_application":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_evidently_feature":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_evidently_launch":
		// No standard import format found in documentation
		return ""
	case "aws_evidently_project":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_evidently_segment":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_finspace_kx_cluster":
		// Computed or complex ID format: the `id` (environment ID and cluster name, comma-delimited)
		return ""
	case "aws_finspace_kx_database":
		// Computed or complex ID format: the `id` (environment ID and database name, comma-delimited)
		return ""
	case "aws_finspace_kx_dataview":
		// Computed or complex ID format: the `id` (environment ID and cluster name, comma-delimited)
		return ""
	case "aws_finspace_kx_environment":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_finspace_kx_scaling_group":
		// Computed or complex ID format: the `id` (environment ID and scaling group name, comma-delimited)
		return ""
	case "aws_finspace_kx_user":
		// Computed or complex ID format: the `id` (environment ID and user name, comma-delimited)
		return ""
	case "aws_finspace_kx_volume":
		// Computed or complex ID format: the `id` (environment ID and volume name, comma-delimited)
		return ""
	case "aws_fis_experiment_template":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_fis_target_account_configuration":
		var parts []string
		if val, ok := config["account_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["experiment_template_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_flow_log":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_fms_admin_account":
		// Computed ID format: the account ID
		return ""
	case "aws_fms_policy":
		// Computed ID format: the policy ID
		return ""
	case "aws_fms_resource_set":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_fsx_backup":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_fsx_data_repository_association":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_fsx_file_cache":
		// Computed or complex ID format: the resource `id`
		return ""
	case "aws_fsx_lustre_file_system":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_fsx_ontap_file_system":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_fsx_ontap_storage_virtual_machine":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_fsx_ontap_volume":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_fsx_openzfs_file_system":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_fsx_openzfs_snapshot":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_fsx_openzfs_volume":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_fsx_s3_access_point_attachment":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_fsx_windows_file_system":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_gamelift_alias":
		// Computed ID format: the ID
		return ""
	case "aws_gamelift_build":
		// Computed ID format: the ID
		return ""
	case "aws_gamelift_fleet":
		// Computed ID format: the ID
		return ""
	case "aws_gamelift_game_server_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_gamelift_game_session_queue":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_gamelift_script":
		// Computed ID format: the ID
		return ""
	case "aws_glacier_vault":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_glacier_vault_lock":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_globalaccelerator_accelerator":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_globalaccelerator_cross_account_attachment":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_globalaccelerator_custom_routing_accelerator":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_globalaccelerator_custom_routing_endpoint_group":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_globalaccelerator_custom_routing_listener":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_globalaccelerator_endpoint_group":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_globalaccelerator_listener":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_glue_catalog_database":
		// Computed or complex ID format: the `catalog_id:name`. If you have not set a Catalog ID specify the AWS Account ID that the database is in
		return ""
	case "aws_glue_catalog_table":
		// Computed ID format: the catalog ID (usually AWS account ID), database name, and table name
		return ""
	case "aws_glue_catalog_table_optimizer":
		var parts []string
		if val, ok := config["catalog_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["database_name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["table_name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["type"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_glue_classifier":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_glue_connection":
		// Computed or complex ID format: the `CATALOG-ID` (AWS account ID if not custom) and `NAME`
		return ""
	case "aws_glue_crawler":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_glue_data_catalog_encryption_settings":
		// Computed or complex ID format: `CATALOG-ID` (AWS account ID if not custom)
		return ""
	case "aws_glue_data_quality_ruleset":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_glue_dev_endpoint":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_glue_job":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_glue_ml_transform":
		// Computed or complex ID format: `id`
		return ""
	case "aws_glue_partition":
		// Computed ID format: the catalog ID (usually AWS account ID), database name, table name and partition values
		return ""
	case "aws_glue_partition_index":
		// Computed ID format: the catalog ID (usually AWS account ID), database name, table name, and index name
		return ""
	case "aws_glue_registry":
		// Computed or complex ID format: `arn`
		return ""
	case "aws_glue_resource_policy":
		// Computed ID format: the region where the resource resides
		return ""
	case "aws_glue_schema":
		// Computed or complex ID format: `arn`
		return ""
	case "aws_glue_security_configuration":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_glue_trigger":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_glue_user_defined_function":
		// Computed or complex ID format: the `catalog_id:database_name:function_name`. If you have not set a Catalog ID specify the AWS Account ID that the database is in
		return ""
	case "aws_glue_workflow":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_grafana_license_association":
		// Computed or complex ID format: the workspace's `id`
		return ""
	case "aws_grafana_role_association":
		// No standard import format found in documentation
		return ""
	case "aws_grafana_workspace":
		// Computed or complex ID format: the workspace's `id`
		return ""
	case "aws_grafana_workspace_api_key":
		// No standard import format found in documentation
		return ""
	case "aws_grafana_workspace_saml_configuration":
		// Computed or complex ID format: the workspace's `id`
		return ""
	case "aws_grafana_workspace_service_account":
		if val, ok := config["workspace_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_grafana_workspace_service_account_token":
		// No standard import format found in documentation
		return ""
	case "aws_guardduty_detector":
		// Computed ID format: the detector ID
		return ""
	case "aws_guardduty_detector_feature":
		// No standard import format found in documentation
		return ""
	case "aws_guardduty_filter":
		// Computed ID format: the detector ID and filter's name separated by a colon
		return ""
	case "aws_guardduty_invite_accepter":
		// Computed ID format: the member GuardDuty detector ID
		return ""
	case "aws_guardduty_ipset":
		// Computed ID format: the primary GuardDuty detector ID and IPSet ID
		return ""
	case "aws_guardduty_malware_protection_plan":
		// Computed ID format: their IDs
		return ""
	case "aws_guardduty_member":
		// Computed ID format: the primary GuardDuty detector ID and member AWS account ID
		return ""
	case "aws_guardduty_member_detector_feature":
		// No standard import format found in documentation
		return ""
	case "aws_guardduty_organization_admin_account":
		// Computed ID format: the AWS account ID
		return ""
	case "aws_guardduty_organization_configuration":
		// Computed ID format: the GuardDuty Detector ID
		return ""
	case "aws_guardduty_organization_configuration_feature":
		// No standard import format found in documentation
		return ""
	case "aws_guardduty_publishing_destination":
		// Computed ID format: the master GuardDuty detector ID and PublishingDestinationID
		return ""
	case "aws_guardduty_threatintelset":
		// Computed ID format: the primary GuardDuty detector ID and ThreatIntelSetID
		return ""
	case "aws_iam_access_key":
		// Computed ID format: the identifier
		return ""
	case "aws_iam_account_alias":
		if val, ok := config["account_alias"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iam_account_password_policy":
		// Computed or complex ID format: the word `iam-account-password-policy`
		return ""
	case "aws_iam_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iam_group_membership":
		// No standard import format found in documentation
		return ""
	case "aws_iam_group_policies_exclusive":
		if val, ok := config["group_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iam_group_policy":
		// Computed or complex ID format: the `group_name:group_policy_name`
		return ""
	case "aws_iam_group_policy_attachment":
		// Computed or complex ID format: the group name and policy arn separated by `/`
		return ""
	case "aws_iam_group_policy_attachments_exclusive":
		if val, ok := config["group_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iam_instance_profile":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iam_openid_connect_provider":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_iam_organizations_features":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_iam_outbound_web_identity_federation":
		// Computed ID format: the AWS account ID
		return ""
	case "aws_iam_policy":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_iam_policy_attachment":
		// No standard import format found in documentation
		return ""
	case "aws_iam_role":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iam_role_policies_exclusive":
		if val, ok := config["role_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iam_role_policy":
		// Computed or complex ID format: the `role_name:role_policy_name`
		return ""
	case "aws_iam_role_policy_attachment":
		// Computed or complex ID format: the role name and policy arn separated by `/`
		return ""
	case "aws_iam_role_policy_attachments_exclusive":
		if val, ok := config["role_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iam_saml_provider":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_iam_security_token_service_preferences":
		// No standard import format found in documentation
		return ""
	case "aws_iam_server_certificate":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iam_service_linked_role":
		// Computed ID format: role ARN
		return ""
	case "aws_iam_service_specific_credential":
		// Computed or complex ID format: the `service_name:user_name:service_specific_credential_id`
		return ""
	case "aws_iam_signing_certificate":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_iam_user":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iam_user_group_membership":
		// Computed or complex ID format: the user name and group names separated by `/`
		return ""
	case "aws_iam_user_login_profile":
		// No standard import format found in documentation
		return ""
	case "aws_iam_user_policies_exclusive":
		if val, ok := config["user_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iam_user_policy":
		// Computed or complex ID format: the `user_name:user_policy_name`
		return ""
	case "aws_iam_user_policy_attachment":
		// Computed or complex ID format: the user name and policy arn separated by `/`
		return ""
	case "aws_iam_user_policy_attachments_exclusive":
		if val, ok := config["user_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iam_user_ssh_key":
		if val, ok := config["username"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iam_virtual_mfa_device":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_identitystore_group":
		var parts []string
		if val, ok := config["identity_store_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["group_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_identitystore_group_membership":
		var parts []string
		if val, ok := config["identity_store_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["membership_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_identitystore_user":
		var parts []string
		if val, ok := config["identity_store_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["user_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_imagebuilder_component":
		// Computed ID format: the Amazon Resource Name (ARN)
		return ""
	case "aws_imagebuilder_container_recipe":
		// Computed ID format: the Amazon Resource Name (ARN)
		return ""
	case "aws_imagebuilder_distribution_configuration":
		// Computed ID format: the Amazon Resource Name (ARN)
		return ""
	case "aws_imagebuilder_image":
		// Computed ID format: the Amazon Resource Name (ARN)
		return ""
	case "aws_imagebuilder_image_pipeline":
		// Computed ID format: the Amazon Resource Name (ARN)
		return ""
	case "aws_imagebuilder_image_recipe":
		// Computed ID format: the Amazon Resource Name (ARN)
		return ""
	case "aws_imagebuilder_infrastructure_configuration":
		// Computed ID format: the Amazon Resource Name (ARN)
		return ""
	case "aws_imagebuilder_lifecycle_policy":
		// Computed ID format: the Amazon Resource Name (ARN)
		return ""
	case "aws_imagebuilder_workflow":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_inspector2_delegated_admin_account":
		if val, ok := config["account_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_inspector2_enabler":
		if val, ok := config["account_ids"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_inspector2_filter":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_inspector2_member_association":
		if val, ok := config["account_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_inspector2_organization_configuration":
		// No standard import format found in documentation
		return ""
	case "aws_inspector_assessment_target":
		// Computed ID format: their Amazon Resource Name (ARN)
		return ""
	case "aws_inspector_assessment_template":
		// Computed ID format: the template assessment ARN
		return ""
	case "aws_inspector_resource_group":
		// No standard import format found in documentation
		return ""
	case "aws_instance":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_internet_gateway":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_internet_gateway_attachment":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_internetmonitor_monitor":
		if val, ok := config["monitor_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_invoicing_invoice_unit":
		// Computed ID format: the ARN
		return ""
	case "aws_iot_authorizer":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iot_billing_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iot_ca_certificate":
		// No standard import format found in documentation
		return ""
	case "aws_iot_certificate":
		// No standard import format found in documentation
		return ""
	case "aws_iot_domain_configuration":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iot_event_configurations":
		// Unknown ID format: the AWS Region
		return ""
	case "aws_iot_indexing_configuration":
		// No standard import format found in documentation
		return ""
	case "aws_iot_logging_options":
		// No standard import format found in documentation
		return ""
	case "aws_iot_policy":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iot_policy_attachment":
		// No standard import format found in documentation
		return ""
	case "aws_iot_provisioning_template":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iot_role_alias":
		// Unknown ID format: the alias
		return ""
	case "aws_iot_thing":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iot_thing_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iot_thing_group_membership":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iot_thing_principal_attachment":
		// No standard import format found in documentation
		return ""
	case "aws_iot_thing_type":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iot_topic_rule":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_iot_topic_rule_destination":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_ivs_channel":
		// Computed ID format: the ARN
		return ""
	case "aws_ivs_playback_key_pair":
		// Computed ID format: the ARN
		return ""
	case "aws_ivs_recording_configuration":
		// Computed ID format: the ARN
		return ""
	case "aws_ivschat_logging_configuration":
		// Computed ID format: the ARN
		return ""
	case "aws_ivschat_room":
		// Computed ID format: the ARN
		return ""
	case "aws_kendra_data_source":
		// Computed or complex ID format: the unique identifiers of the data_source and index separated by a slash (`/`)
		return ""
	case "aws_kendra_experience":
		// Computed or complex ID format: the unique identifiers of the experience and index separated by a slash (`/`)
		return ""
	case "aws_kendra_faq":
		// Computed or complex ID format: the unique identifiers of the FAQ and index separated by a slash (`/`)
		return ""
	case "aws_kendra_index":
		// Computed or complex ID format: its `id`
		return ""
	case "aws_kendra_query_suggestions_block_list":
		// Computed or complex ID format: the unique identifiers of the block list and index separated by a slash (`/`)
		return ""
	case "aws_kendra_thesaurus":
		// Computed or complex ID format: the unique identifiers of the thesaurus and index separated by a slash (`/`)
		return ""
	case "aws_key_pair":
		if val, ok := config["key_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_keyspaces_keyspace":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_keyspaces_table":
		if val, ok := config["keyspace_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_kinesis_analytics_application":
		// Computed ID format: ARN
		return ""
	case "aws_kinesis_firehose_delivery_stream":
		// Computed ID format: the stream ARN
		return ""
	case "aws_kinesis_resource_policy":
		if val, ok := config["resource_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_kinesis_stream":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_kinesis_stream_consumer":
		// Computed ID format: the Amazon Resource Name (ARN)
		return ""
	case "aws_kinesis_video_stream":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_kinesisanalyticsv2_application":
		// Computed ID format: the application ARN
		return ""
	case "aws_kinesisanalyticsv2_application_snapshot":
		if val, ok := config["application_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_kms_alias":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_kms_ciphertext":
		// No standard import format found in documentation
		return ""
	case "aws_kms_custom_key_store":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_kms_external_key":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_kms_grant":
		// Computed or complex ID format: the Key ID and Grant ID separated by a colon (`:`)
		return ""
	case "aws_kms_key":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_kms_key_policy":
		if val, ok := config["key_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_kms_replica_external_key":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_kms_replica_key":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_lakeformation_data_cells_filter":
		if val, ok := config["database_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lakeformation_data_lake_settings":
		// No standard import format found in documentation
		return ""
	case "aws_lakeformation_identity_center_configuration":
		if val, ok := config["catalog_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lakeformation_lf_tag":
		// Computed or complex ID format: the `catalog_id:key`. If you have not set a Catalog ID specify the AWS Account ID that the database is in
		return ""
	case "aws_lakeformation_lf_tag_expression":
		var parts []string
		if val, ok := config["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["catalog_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_lakeformation_opt_in":
		// No standard import format found in documentation
		return ""
	case "aws_lakeformation_permissions":
		// No standard import format found in documentation
		return ""
	case "aws_lakeformation_resource":
		// No standard import format found in documentation
		return ""
	case "aws_lakeformation_resource_lf_tag":
		// No standard import format found in documentation
		return ""
	case "aws_lakeformation_resource_lf_tags":
		// No standard import format found in documentation
		return ""
	case "aws_lambda_alias":
		// No standard import format found in documentation
		return ""
	case "aws_lambda_capacity_provider":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lambda_code_signing_config":
		// No standard import format found in documentation
		return ""
	case "aws_lambda_event_source_mapping":
		// Computed or complex ID format: the `UUID` (event source mapping identifier)
		return ""
	case "aws_lambda_function":
		if val, ok := config["function_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lambda_function_event_invoke_config":
		// No standard import format found in documentation
		return ""
	case "aws_lambda_function_recursion_config":
		// No standard import format found in documentation
		return ""
	case "aws_lambda_function_url":
		if val, ok := config["function_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lambda_invocation":
		var parts []string
		if val, ok := config["function_name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["qualifier"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["result_hash"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_lambda_layer_version":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_lambda_layer_version_permission":
		// No standard import format found in documentation
		return ""
	case "aws_lambda_permission":
		// No standard import format found in documentation
		return ""
	case "aws_lambda_provisioned_concurrency_config":
		if val, ok := config["function_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lambda_runtime_management_config":
		if val, ok := config["function_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_launch_configuration":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_launch_template":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_lb":
		// Computed ID format: their ARN
		return ""
	case "aws_lb_cookie_stickiness_policy":
		// No standard import format found in documentation
		return ""
	case "aws_lb_listener":
		// Computed ID format: their ARN
		return ""
	case "aws_lb_listener_certificate":
		if val, ok := config["_"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lb_listener_rule":
		// Computed ID format: their ARN
		return ""
	case "aws_lb_ssl_negotiation_policy":
		// No standard import format found in documentation
		return ""
	case "aws_lb_target_group":
		// Computed ID format: their ARN
		return ""
	case "aws_lb_target_group_attachment":
		// Unknown ID format: the same format
		return ""
	case "aws_lb_trust_store":
		// Computed ID format: their ARN
		return ""
	case "aws_lb_trust_store_revocation":
		// Computed ID format: their ARN
		return ""
	case "aws_lex_bot":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lex_bot_alias":
		// Computed or complex ID format: an ID with the format `bot_name:bot_alias_name`
		return ""
	case "aws_lex_intent":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lex_slot_type":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lexv2models_bot":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_lexv2models_bot_locale":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_lexv2models_bot_version":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_lexv2models_intent":
		// Computed or complex ID format: the `intent_id:bot_id:bot_version:locale_id`
		return ""
	case "aws_lexv2models_slot":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_lexv2models_slot_type":
		if val, ok := config["bot_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_licensemanager_association":
		var parts []string
		if val, ok := config["resource_arn"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["license_configuration_arn"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_licensemanager_grant":
		// Computed ID format: the grant arn
		return ""
	case "aws_licensemanager_grant_accepter":
		// Computed ID format: the grant arn
		return ""
	case "aws_licensemanager_license_configuration":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_lightsail_bucket":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lightsail_bucket_access_key":
		// Computed or complex ID format: the `id` attribute
		return ""
	case "aws_lightsail_bucket_resource_access":
		// Computed or complex ID format: the `id` attribute
		return ""
	case "aws_lightsail_certificate":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lightsail_container_service":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lightsail_container_service_deployment_version":
		if val, ok := config["service_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lightsail_database":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lightsail_disk":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lightsail_disk_attachment":
		// Computed ID format: the id attribute
		return ""
	case "aws_lightsail_distribution":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lightsail_domain":
		// No standard import format found in documentation
		return ""
	case "aws_lightsail_domain_entry":
		// Computed ID format: the id attribute
		return ""
	case "aws_lightsail_instance":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lightsail_instance_public_ports":
		// No standard import format found in documentation
		return ""
	case "aws_lightsail_key_pair":
		// No standard import format found in documentation
		return ""
	case "aws_lightsail_lb":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lightsail_lb_attachment":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lightsail_lb_certificate":
		// Computed ID format: the id attribute
		return ""
	case "aws_lightsail_lb_certificate_attachment":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lightsail_lb_https_redirection_policy":
		if val, ok := config["lb_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lightsail_lb_stickiness_policy":
		if val, ok := config["lb_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lightsail_static_ip":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_lightsail_static_ip_attachment":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_load_balancer_backend_server_policy":
		// No standard import format found in documentation
		return ""
	case "aws_load_balancer_listener_policy":
		// No standard import format found in documentation
		return ""
	case "aws_load_balancer_policy":
		// No standard import format found in documentation
		return ""
	case "aws_location_geofence_collection":
		if val, ok := config["collection_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_location_map":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_location_place_index":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_location_route_calculator":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_location_tracker":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_location_tracker_association":
		// Computed or complex ID format: the `tracker_name|consumer_arn`
		return ""
	case "aws_m2_application":
		// Computed or complex ID format: the `01234567890abcdef012345678`
		return ""
	case "aws_m2_deployment":
		// Computed or complex ID format: the `APPLICATION-ID,DEPLOYMENT-ID`
		return ""
	case "aws_m2_environment":
		// Computed or complex ID format: the `01234567890abcdef012345678`
		return ""
	case "aws_macie2_account":
		// Computed ID format: the id
		return ""
	case "aws_macie2_classification_export_configuration":
		// Unknown ID format: the region
		return ""
	case "aws_macie2_classification_job":
		// Computed ID format: the id
		return ""
	case "aws_macie2_custom_data_identifier":
		// Computed ID format: the id
		return ""
	case "aws_macie2_findings_filter":
		// Computed ID format: the id
		return ""
	case "aws_macie2_invitation_accepter":
		// Computed ID format: the admin account ID
		return ""
	case "aws_macie2_member":
		// Computed ID format: the account ID of the member account
		return ""
	case "aws_macie2_organization_admin_account":
		// Computed ID format: the id
		return ""
	case "aws_macie2_organization_configuration":
		// No standard import format found in documentation
		return ""
	case "aws_main_route_table_association":
		// No standard import format found in documentation
		return ""
	case "aws_media_convert_queue":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_media_package_channel":
		// Computed ID format: the channel ID
		return ""
	case "aws_media_packagev2_channel_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_media_store_container":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_media_store_container_policy":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_medialive_channel":
		if val, ok := config["channel_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_medialive_input":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_medialive_input_security_group":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_medialive_multiplex":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_medialive_multiplex_program":
		// Computed or complex ID format: the `id`, or a combination of "`program_name`/`multiplex_id`"
		return ""
	case "aws_memorydb_acl":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_memorydb_cluster":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_memorydb_multi_region_cluster":
		if val, ok := config["multi_region_cluster_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_memorydb_parameter_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_memorydb_snapshot":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_memorydb_subnet_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_memorydb_user":
		if val, ok := config["user_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_mq_broker":
		// Computed ID format: their broker id
		return ""
	case "aws_mq_configuration":
		// Computed ID format: the configuration ID
		return ""
	case "aws_msk_cluster":
		// Computed ID format: the cluster ARN
		return ""
	case "aws_msk_cluster_policy":
		if val, ok := config["cluster_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_msk_configuration":
		// Computed ID format: the configuration ARN
		return ""
	case "aws_msk_replicator":
		// Computed ID format: the replicator ARN
		return ""
	case "aws_msk_scram_secret_association":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_msk_serverless_cluster":
		// Computed ID format: the cluster ARN
		return ""
	case "aws_msk_single_scram_secret_association":
		if val, ok := config["cluster_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_msk_topic":
		if val, ok := config["cluster_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_msk_vpc_connection":
		// Computed ID format: the configuration ARN
		return ""
	case "aws_mskconnect_connector":
		// Computed or complex ID format: the connector's `arn`
		return ""
	case "aws_mskconnect_custom_plugin":
		// Computed or complex ID format: the plugin's `arn`
		return ""
	case "aws_mskconnect_worker_configuration":
		// Computed or complex ID format: the plugin's `arn`
		return ""
	case "aws_mwaa_environment":
		// Computed or complex ID format: `Name`
		return ""
	case "aws_nat_gateway":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_nat_gateway_eip_association":
		var parts []string
		if val, ok := config["nat_gateway_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["allocation_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_neptune_cluster":
		// Computed ID format: the cluster identifier
		return ""
	case "aws_neptune_cluster_endpoint":
		// Computed or complex ID format: the `cluster-identifier:endpoint-identfier`
		return ""
	case "aws_neptune_cluster_instance":
		// Computed ID format: the instance identifier
		return ""
	case "aws_neptune_cluster_parameter_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_neptune_cluster_snapshot":
		// Computed ID format: the cluster snapshot identifier
		return ""
	case "aws_neptune_event_subscription":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_neptune_global_cluster":
		// Computed ID format: the Global Cluster identifier
		return ""
	case "aws_neptune_parameter_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_neptune_subnet_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_neptunegraph_graph":
		// Computed ID format: the graph identifier
		return ""
	case "aws_network_acl":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_network_acl_association":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_network_acl_rule":
		// Computed or complex ID format: `NETWORK_ACL_ID:RULE_NUMBER:PROTOCOL:EGRESS`, where `PROTOCOL` can be a decimal (such as `"6"`) or string (such as `"tcp"`) value
		return ""
	case "aws_network_interface":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_network_interface_attachment":
		// Computed ID format: its Attachment ID
		return ""
	case "aws_network_interface_permission":
		if val, ok := config["network_interface_permission_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_network_interface_sg_attachment":
		if val, ok := config["_"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_networkfirewall_firewall":
		// Computed or complex ID format: their `arn`
		return ""
	case "aws_networkfirewall_firewall_policy":
		// Computed or complex ID format: their `arn`
		return ""
	case "aws_networkfirewall_firewall_transit_gateway_attachment_accepter":
		if val, ok := config["transit_gateway_attachment_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_networkfirewall_logging_configuration":
		if val, ok := config["firewall_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_networkfirewall_resource_policy":
		if val, ok := config["resource_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_networkfirewall_rule_group":
		// Computed or complex ID format: their `arn`
		return ""
	case "aws_networkfirewall_tls_inspection_configuration":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_networkfirewall_vpc_endpoint_association":
		if val, ok := config["vpc_endpoint_association_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_networkflowmonitor_monitor":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_networkflowmonitor_scope":
		// Computed ID format: the scope ID
		return ""
	case "aws_networkmanager_attachment_accepter":
		// No standard import format found in documentation
		return ""
	case "aws_networkmanager_attachment_routing_policy_label":
		// Computed or complex ID format: the core network ID and attachment ID separated by a comma (`,`)
		return ""
	case "aws_networkmanager_connect_attachment":
		// Computed ID format: the attachment ID
		return ""
	case "aws_networkmanager_connect_peer":
		// Computed ID format: the connect peer ID
		return ""
	case "aws_networkmanager_connection":
		// Computed ID format: the connection ARN
		return ""
	case "aws_networkmanager_core_network":
		// Computed ID format: the core network ID
		return ""
	case "aws_networkmanager_core_network_policy_attachment":
		// Computed ID format: the core network ID
		return ""
	case "aws_networkmanager_customer_gateway_association":
		// Computed ID format: the global network ID and customer gateway ARN
		return ""
	case "aws_networkmanager_device":
		// Computed ID format: the device ARN
		return ""
	case "aws_networkmanager_dx_gateway_attachment":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_networkmanager_global_network":
		// Computed ID format: the global network ID
		return ""
	case "aws_networkmanager_link":
		// Computed ID format: the link ARN
		return ""
	case "aws_networkmanager_link_association":
		// Computed ID format: the global network ID, link ID and device ID
		return ""
	case "aws_networkmanager_prefix_list_association":
		// Computed or complex ID format: the core network ID and prefix list ARN separated by a comma (`,`)
		return ""
	case "aws_networkmanager_site":
		// Computed ID format: the site ARN
		return ""
	case "aws_networkmanager_site_to_site_vpn_attachment":
		// Computed ID format: the attachment ID
		return ""
	case "aws_networkmanager_transit_gateway_connect_peer_association":
		// Computed ID format: the global network ID and Connect peer ARN
		return ""
	case "aws_networkmanager_transit_gateway_peering":
		// Computed ID format: the peering ID
		return ""
	case "aws_networkmanager_transit_gateway_registration":
		// Computed ID format: the global network ID and transit gateway ARN
		return ""
	case "aws_networkmanager_transit_gateway_route_table_attachment":
		// Computed ID format: the attachment ID
		return ""
	case "aws_networkmanager_vpc_attachment":
		// Computed ID format: the attachment ID
		return ""
	case "aws_networkmonitor_monitor":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_networkmonitor_probe":
		// Computed ID format: the monitor name and probe id
		return ""
	case "aws_notifications_channel_association":
		var parts []string
		if val, ok := config["notification_configuration_arn"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["channel_arn"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_notifications_event_rule":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_notifications_managed_notification_account_contact_association":
		var parts []string
		if val, ok := config["managed_notification_configuration_arn"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["contact_identifier"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_notifications_managed_notification_additional_channel_association":
		var parts []string
		if val, ok := config["managed_notification_arn"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["channel_arn"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_notifications_notification_configuration":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_notifications_notification_hub":
		if val, ok := config["notification_hub_region"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_notifications_organizational_unit_association":
		var parts []string
		if val, ok := config["notification_configuration_arn"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["organizational_unit_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_notifications_organizations_access":
		// Computed ID format: the AWS account ID
		return ""
	case "aws_notificationscontacts_email_contact":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_oam_link":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_oam_sink":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_oam_sink_policy":
		if val, ok := config["sink_identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_observabilityadmin_centralization_rule_for_organization":
		if val, ok := config["rule_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_observabilityadmin_telemetry_enrichment":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_observabilityadmin_telemetry_pipeline":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_odb_cloud_autonomous_vm_cluster":
		// No standard import format found in documentation
		return ""
	case "aws_odb_cloud_exadata_infrastructure":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_odb_cloud_vm_cluster":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_odb_network":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_odb_network_peering_connection":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_opensearch_application":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_opensearch_authorize_vpc_endpoint_access":
		if val, ok := config["domain_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_opensearch_domain":
		if val, ok := config["domain_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_opensearch_domain_policy":
		if val, ok := config["domain_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_opensearch_domain_saml_options":
		if val, ok := config["domain_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_opensearch_inbound_connection_accepter":
		// Computed ID format: the Inbound Connection ID
		return ""
	case "aws_opensearch_outbound_connection":
		// Computed ID format: the Outbound Connection ID
		return ""
	case "aws_opensearch_package":
		// Computed ID format: the Package ID
		return ""
	case "aws_opensearch_package_association":
		// No standard import format found in documentation
		return ""
	case "aws_opensearch_vpc_endpoint":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_opensearchserverless_access_policy":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_opensearchserverless_collection":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_opensearchserverless_collection_group":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_opensearchserverless_lifecycle_policy":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_opensearchserverless_security_config":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_opensearchserverless_security_policy":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_opensearchserverless_vpc_endpoint":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_organizations_account":
		if val, ok := config["account_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_organizations_aws_service_access":
		if val, ok := config["service_principal"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_organizations_delegated_administrator":
		// Computed ID format: the account ID and its service principal
		return ""
	case "aws_organizations_organization":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_organizations_organizational_unit":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_organizations_policy":
		// Computed ID format: the policy ID
		return ""
	case "aws_organizations_policy_attachment":
		// Computed ID format: the target ID and policy ID
		return ""
	case "aws_organizations_resource_policy":
		// Computed ID format: the resource policy ID
		return ""
	case "aws_organizations_tag":
		// Computed or complex ID format: the Organizations resource identifier and key, separated by a comma (`,`)
		return ""
	case "aws_osis_pipeline":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_paymentcryptography_key":
		// Computed or complex ID format: the `arn:aws:payment-cryptography:us-east-1:123456789012:key/qtbojf64yshyvyzf`
		return ""
	case "aws_paymentcryptography_key_alias":
		// Computed or complex ID format: the `alias/4681482429376900170`
		return ""
	case "aws_pinpoint_adm_channel":
		// Computed or complex ID format: the `application-id`
		return ""
	case "aws_pinpoint_apns_channel":
		// Computed or complex ID format: the `application-id`
		return ""
	case "aws_pinpoint_apns_sandbox_channel":
		// Computed or complex ID format: the `application-id`
		return ""
	case "aws_pinpoint_apns_voip_channel":
		// Computed or complex ID format: the `application-id`
		return ""
	case "aws_pinpoint_apns_voip_sandbox_channel":
		// Computed or complex ID format: the `application-id`
		return ""
	case "aws_pinpoint_app":
		// Computed or complex ID format: the `application-id`
		return ""
	case "aws_pinpoint_baidu_channel":
		// Computed or complex ID format: the `application-id`
		return ""
	case "aws_pinpoint_email_channel":
		// Computed or complex ID format: the `application-id`
		return ""
	case "aws_pinpoint_email_template":
		if val, ok := config["template_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_pinpoint_event_stream":
		// Computed or complex ID format: the `application-id`
		return ""
	case "aws_pinpoint_gcm_channel":
		// Computed or complex ID format: the `application-id`
		return ""
	case "aws_pinpoint_sms_channel":
		if val, ok := config["application_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_pinpointsmsvoicev2_configuration_set":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_pinpointsmsvoicev2_opt_out_list":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_pinpointsmsvoicev2_phone_number":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_pipes_pipe":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_placement_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_prometheus_alert_manager_definition":
		// Computed ID format: the workspace identifier
		return ""
	case "aws_prometheus_query_logging_configuration":
		// No standard import format found in documentation
		return ""
	case "aws_prometheus_resource_policy":
		// Computed ID format: the workspace ID
		return ""
	case "aws_prometheus_rule_group_namespace":
		// Computed ID format: the arn
		return ""
	case "aws_prometheus_scraper":
		// Computed ID format: its identifier
		return ""
	case "aws_prometheus_workspace":
		// Computed ID format: the identifier
		return ""
	case "aws_prometheus_workspace_configuration":
		if val, ok := config["workspace_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_proxy_protocol_policy":
		// No standard import format found in documentation
		return ""
	case "aws_qbusiness_application":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_qldb_ledger":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_qldb_stream":
		// No standard import format found in documentation
		return ""
	case "aws_quicksight_account_settings":
		// Computed ID format: the AWS account ID
		return ""
	case "aws_quicksight_account_subscription":
		if val, ok := config["aws_account_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_quicksight_analysis":
		// Computed or complex ID format: the AWS account ID and analysis ID separated by a comma (`,`)
		return ""
	case "aws_quicksight_custom_permissions":
		// Computed or complex ID format: the AWS account ID and custom permissions profile name separated by a comma (`,`)
		return ""
	case "aws_quicksight_dashboard":
		// Computed or complex ID format: the AWS account ID and dashboard ID separated by a comma (`,`)
		return ""
	case "aws_quicksight_data_set":
		// Computed or complex ID format: the AWS account ID and data set ID separated by a comma (`,`)
		return ""
	case "aws_quicksight_data_source":
		// Computed or complex ID format: the AWS account ID, and data source ID separated by a slash (`/`)
		return ""
	case "aws_quicksight_folder":
		// Computed or complex ID format: the AWS account ID and folder ID name separated by a comma (`,`)
		return ""
	case "aws_quicksight_folder_membership":
		// Computed or complex ID format: the AWS account ID, folder ID, member type, and member ID separated by commas (`,`)
		return ""
	case "aws_quicksight_group":
		// Computed or complex ID format: the aws account id, namespace and group name separated by `/`
		return ""
	case "aws_quicksight_group_membership":
		// Computed or complex ID format: the AWS account ID, namespace, group name and member name separated by `/`
		return ""
	case "aws_quicksight_iam_policy_assignment":
		// Computed or complex ID format: the AWS account ID, namespace, and assignment name separated by commas (`,`)
		return ""
	case "aws_quicksight_ingestion":
		// Computed or complex ID format: the AWS account ID, data set ID, and ingestion ID separated by commas (`,`)
		return ""
	case "aws_quicksight_ip_restriction":
		// Computed ID format: the AWS account ID
		return ""
	case "aws_quicksight_key_registration":
		// Computed ID format: the AWS account ID
		return ""
	case "aws_quicksight_namespace":
		// Computed or complex ID format: the AWS account ID and namespace separated by commas (`,`)
		return ""
	case "aws_quicksight_refresh_schedule":
		// Computed or complex ID format: the AWS account ID, data set ID and schedule ID separated by commas (`,`)
		return ""
	case "aws_quicksight_role_custom_permission":
		if val, ok := config["aws_account_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_quicksight_role_membership":
		if val, ok := config["aws_account_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_quicksight_template":
		// Computed or complex ID format: the AWS account ID and template ID separated by a comma (`,`)
		return ""
	case "aws_quicksight_template_alias":
		// Computed or complex ID format: the AWS account ID, template ID, and alias name separated by a comma (`,`)
		return ""
	case "aws_quicksight_theme":
		// Computed or complex ID format: the AWS account ID and theme ID separated by a comma (`,`)
		return ""
	case "aws_quicksight_user":
		// No standard import format found in documentation
		return ""
	case "aws_quicksight_user_custom_permission":
		if val, ok := config["aws_account_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_quicksight_vpc_connection":
		// Computed or complex ID format: the AWS account ID and VPC connection ID separated by commas (`,`)
		return ""
	case "aws_ram_permission":
		if val, ok := config["example_id_arg"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ram_principal_association":
		if val, ok := config["principal"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ram_resource_association":
		// Computed ID format: their Resource Share ARN and Resource ARN separated by a comma
		return ""
	case "aws_ram_resource_share":
		// Computed or complex ID format: the `arn` of the resource share
		return ""
	case "aws_ram_resource_share_accepter":
		// Computed ID format: the resource share ARN
		return ""
	case "aws_ram_resource_share_associations_exclusive":
		if val, ok := config["resource_share_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ram_sharing_with_organization":
		// Computed ID format: the current AWS account ID
		return ""
	case "aws_rbin_rule":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_rds_certificate":
		if val, ok := config["region"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_rds_cluster":
		if val, ok := config["cluster_identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_rds_cluster_activity_stream":
		if val, ok := config["resource_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_rds_cluster_endpoint":
		if val, ok := config["cluster_endpoint_identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_rds_cluster_instance":
		if val, ok := config["identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_rds_cluster_parameter_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_rds_cluster_role_association":
		// Computed or complex ID format: the DB Cluster Identifier and IAM Role ARN separated by a comma (`,`)
		return ""
	case "aws_rds_cluster_snapshot_copy":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_rds_custom_db_engine_version":
		if val, ok := config["engine"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_rds_export_task":
		if val, ok := config["export_task_identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_rds_global_cluster":
		// Computed ID format: the RDS Global Cluster identifier
		return ""
	case "aws_rds_instance_state":
		if val, ok := config["identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_rds_integration":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_rds_reserved_instance":
		if val, ok := config["instance_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_rds_shard_group":
		if val, ok := config["db_shard_group_identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_redshift_authentication_profile":
		// No standard import format found in documentation
		return ""
	case "aws_redshift_cluster":
		if val, ok := config["cluster_identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_redshift_cluster_iam_roles":
		if val, ok := config["cluster_identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_redshift_cluster_snapshot":
		if val, ok := config["snapshot_identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_redshift_data_share_authorization":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_redshift_data_share_consumer_association":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_redshift_endpoint_access":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_redshift_endpoint_authorization":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_redshift_event_subscription":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_redshift_hsm_client_certificate":
		if val, ok := config["hsm_client_certificate_identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_redshift_hsm_configuration":
		if val, ok := config["hsm_configuration_identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_redshift_idc_application":
		if val, ok := config["redshift_idc_application_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_redshift_integration":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_redshift_logging":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_redshift_parameter_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_redshift_partner":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_redshift_resource_policy":
		if val, ok := config["resource_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_redshift_scheduled_action":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_redshift_snapshot_copy":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_redshift_snapshot_copy_grant":
		// No standard import format found in documentation
		return ""
	case "aws_redshift_snapshot_schedule":
		if val, ok := config["identifier"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_redshift_snapshot_schedule_association":
		// Computed or complex ID format: the `<cluster-identifier>/<schedule-identifier>`
		return ""
	case "aws_redshift_subnet_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_redshift_usage_limit":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_redshiftdata_statement":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_redshiftserverless_custom_domain_association":
		if val, ok := config["workgroup_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_redshiftserverless_endpoint_access":
		if val, ok := config["endpoint_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_redshiftserverless_namespace":
		if val, ok := config["namespace_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_redshiftserverless_resource_policy":
		if val, ok := config["resource_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_redshiftserverless_snapshot":
		if val, ok := config["snapshot_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_redshiftserverless_usage_limit":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_redshiftserverless_workgroup":
		if val, ok := config["workgroup_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_rekognition_collection":
		if val, ok := config["collection_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_rekognition_project":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_rekognition_stream_processor":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_resiliencehub_resiliency_policy":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_resourceexplorer2_index":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_resourceexplorer2_view":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_resourcegroups_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_resourcegroups_resource":
		if val, ok := config["group_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_rolesanywhere_profile":
		// Computed or complex ID format: its `id`
		return ""
	case "aws_rolesanywhere_trust_anchor":
		// Computed or complex ID format: its `id`
		return ""
	case "aws_route":
		// No standard import format found in documentation
		return ""
	case "aws_route53_cidr_collection":
		// Computed ID format: their ID
		return ""
	case "aws_route53_cidr_location":
		// Computed ID format: their the CIDR collection ID and location name
		return ""
	case "aws_route53_delegation_set":
		// Computed or complex ID format: the delegation set `id`
		return ""
	case "aws_route53_health_check":
		// Computed or complex ID format: the health check `id`
		return ""
	case "aws_route53_hosted_zone_dnssec":
		// Computed ID format: the Route 53 Hosted Zone identifier
		return ""
	case "aws_route53_key_signing_key":
		// Computed or complex ID format: the Route 53 Hosted Zone identifier and KMS Key identifier, separated by a comma (`,`)
		return ""
	case "aws_route53_query_log":
		// Computed ID format: their ID
		return ""
	case "aws_route53_record":
		// No standard import format found in documentation
		return ""
	case "aws_route53_records_exclusive":
		if val, ok := config["zone_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_route53_resolver_config":
		// Computed ID format: the Route 53 Resolver config ID
		return ""
	case "aws_route53_resolver_dnssec_config":
		// Computed ID format: the Route 53 Resolver DNSSEC config ID
		return ""
	case "aws_route53_resolver_endpoint":
		// Computed ID format: the Route 53 Resolver endpoint ID
		return ""
	case "aws_route53_resolver_firewall_config":
		// Computed ID format: the Route 53 Resolver DNS Firewall config ID
		return ""
	case "aws_route53_resolver_firewall_domain_list":
		// Computed ID format: the Route 53 Resolver DNS Firewall domain list ID
		return ""
	case "aws_route53_resolver_firewall_rule":
		// Computed ID format: the Route 53 Resolver DNS Firewall rule group ID and domain list ID (for standard rules) or threat protection ID (for advanced rules) separated by ':'
		return ""
	case "aws_route53_resolver_firewall_rule_group":
		// Computed ID format: the Route 53 Resolver DNS Firewall rule group ID
		return ""
	case "aws_route53_resolver_firewall_rule_group_association":
		// Computed ID format: the Route 53 Resolver DNS Firewall rule group association ID
		return ""
	case "aws_route53_resolver_query_log_config":
		// Computed ID format: the Route 53 Resolver query logging configuration ID
		return ""
	case "aws_route53_resolver_query_log_config_association":
		// Computed ID format: the Route 53 Resolver query logging configuration association ID
		return ""
	case "aws_route53_resolver_rule":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_route53_resolver_rule_association":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_route53_traffic_policy":
		// Computed or complex ID format: the `id` and `version`
		return ""
	case "aws_route53_traffic_policy_instance":
		// Computed ID format: its id
		return ""
	case "aws_route53_vpc_association_authorization":
		// Computed or complex ID format: the Hosted Zone ID and VPC ID, separated by a colon (`:`)
		return ""
	case "aws_route53_zone":
		// Computed or complex ID format: the zone `id`
		return ""
	case "aws_route53_zone_association":
		// No standard import format found in documentation
		return ""
	case "aws_route53domains_delegation_signer_record":
		// Computed or complex ID format: the domain name and DNSSEC key ID, separated by a comma (`,`)
		return ""
	case "aws_route53domains_domain":
		if val, ok := config["domain_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_route53domains_registered_domain":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_route53profiles_association":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_route53profiles_profile":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_route53profiles_resource_association":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_route53recoverycontrolconfig_cluster":
		// Computed ID format: the cluster ARN
		return ""
	case "aws_route53recoverycontrolconfig_control_panel":
		// Computed ID format: the control panel arn
		return ""
	case "aws_route53recoverycontrolconfig_routing_control":
		// Computed ID format: the routing control arn
		return ""
	case "aws_route53recoverycontrolconfig_safety_rule":
		// Computed ID format: the safety rule ARN
		return ""
	case "aws_route53recoveryreadiness_cell":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_route53recoveryreadiness_readiness_check":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_route53recoveryreadiness_recovery_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_route53recoveryreadiness_resource_set":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_route_table":
		// Computed or complex ID format: the route table `id`
		return ""
	case "aws_route_table_association":
		// No standard import format found in documentation
		return ""
	case "aws_rum_app_monitor":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_rum_metrics_destination":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_s3_access_point":
		// No standard import format found in documentation
		return ""
	case "aws_s3_account_public_access_block":
		// Computed ID format: the AWS account ID
		return ""
	case "aws_s3_bucket":
		if val, ok := config["bucket"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3_bucket_abac":
		if val, ok := config["bucket"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3_bucket_accelerate_configuration":
		// No standard import format found in documentation
		return ""
	case "aws_s3_bucket_acl":
		// No standard import format found in documentation
		return ""
	case "aws_s3_bucket_analytics_configuration":
		// Computed or complex ID format: `bucket:analytics`
		return ""
	case "aws_s3_bucket_cors_configuration":
		// No standard import format found in documentation
		return ""
	case "aws_s3_bucket_intelligent_tiering_configuration":
		// Computed or complex ID format: `bucket:name`
		return ""
	case "aws_s3_bucket_inventory":
		// Computed or complex ID format: `bucket:inventory`
		return ""
	case "aws_s3_bucket_lifecycle_configuration":
		if val, ok := config["bucket"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3_bucket_logging":
		// No standard import format found in documentation
		return ""
	case "aws_s3_bucket_metadata_configuration":
		// No standard import format found in documentation
		return ""
	case "aws_s3_bucket_metric":
		// Computed or complex ID format: `bucket:metric`
		return ""
	case "aws_s3_bucket_notification":
		if val, ok := config["bucket"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3_bucket_object":
		// No standard import format found in documentation
		return ""
	case "aws_s3_bucket_object_lock_configuration":
		// No standard import format found in documentation
		return ""
	case "aws_s3_bucket_ownership_controls":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3_bucket_policy":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3_bucket_public_access_block":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3_bucket_replication_configuration":
		if val, ok := config["bucket"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3_bucket_request_payment_configuration":
		// No standard import format found in documentation
		return ""
	case "aws_s3_bucket_server_side_encryption_configuration":
		// No standard import format found in documentation
		return ""
	case "aws_s3_bucket_versioning":
		// No standard import format found in documentation
		return ""
	case "aws_s3_bucket_website_configuration":
		// No standard import format found in documentation
		return ""
	case "aws_s3_directory_bucket":
		if val, ok := config["bucket"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3_object":
		// No standard import format found in documentation
		return ""
	case "aws_s3_object_copy":
		// No standard import format found in documentation
		return ""
	case "aws_s3control_access_grant":
		if val, ok := config["account_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3control_access_grants_instance":
		if val, ok := config["account_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3control_access_grants_instance_resource_policy":
		if val, ok := config["account_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3control_access_grants_location":
		if val, ok := config["account_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3control_access_point_policy":
		if val, ok := config["access_point_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3control_bucket":
		// Computed ID format: Amazon Resource Name (ARN)
		return ""
	case "aws_s3control_bucket_lifecycle_configuration":
		// Computed ID format: the Amazon Resource Name (ARN)
		return ""
	case "aws_s3control_bucket_policy":
		// Computed ID format: the Amazon Resource Name (ARN)
		return ""
	case "aws_s3control_directory_bucket_access_point_scope":
		// Computed or complex ID format: access point name and AWS account ID separated by a colon (`,`)
		return ""
	case "aws_s3control_multi_region_access_point":
		if val, ok := config["account_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3control_multi_region_access_point_policy":
		if val, ok := config["account_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3control_object_lambda_access_point":
		if val, ok := config["account_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3control_object_lambda_access_point_policy":
		if val, ok := config["account_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3control_storage_lens_configuration":
		if val, ok := config["account_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3files_access_point":
		// Computed or complex ID format: `id`
		return ""
	case "aws_s3files_file_system":
		// Computed or complex ID format: `id`
		return ""
	case "aws_s3files_file_system_policy":
		if val, ok := config["file_system_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3files_mount_target":
		// Computed or complex ID format: `id`
		return ""
	case "aws_s3files_synchronization_configuration":
		if val, ok := config["file_system_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3outposts_endpoint":
		// Computed or complex ID format: Amazon Resource Name (ARN), EC2 Security Group identifier, and EC2 Subnet identifier, separated by commas (`,`)
		return ""
	case "aws_s3tables_namespace":
		if val, ok := config["table_bucket_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3tables_table":
		if val, ok := config["table_bucket_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3tables_table_bucket":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_s3tables_table_bucket_policy":
		if val, ok := config["table_bucket_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3tables_table_bucket_replication":
		if val, ok := config["table_bucket_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3tables_table_policy":
		if val, ok := config["table_bucket_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3tables_table_replication":
		if val, ok := config["table_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3vectors_index":
		if val, ok := config["index_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3vectors_vector_bucket":
		if val, ok := config["vector_bucket_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_s3vectors_vector_bucket_policy":
		if val, ok := config["vector_bucket_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_algorithm":
		if val, ok := config["algorithm_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_app":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_sagemaker_app_image_config":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_code_repository":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_data_quality_job_definition":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_device":
		// Computed or complex ID format: the `device-fleet-name/device-name`
		return ""
	case "aws_sagemaker_device_fleet":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_domain":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_sagemaker_endpoint":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_endpoint_configuration":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_feature_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_flow_definition":
		if val, ok := config["flow_definition_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_hub":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_human_task_ui":
		if val, ok := config["human_task_ui_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_hyper_parameter_tuning_job":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_image":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_image_version":
		if val, ok := config["image_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_labeling_job":
		if val, ok := config["labeling_job_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_mlflow_app":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_sagemaker_mlflow_tracking_server":
		if val, ok := config["workteam_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_model":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_model_card":
		if val, ok := config["model_card_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_model_card_export_job":
		if val, ok := config["model_card_export_job_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_model_package_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_model_package_group_policy":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_monitoring_schedule":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_notebook_instance":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_notebook_instance_lifecycle_configuration":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_pipeline":
		if val, ok := config["pipeline_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_project":
		if val, ok := config["project_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_servicecatalog_portfolio_status":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_sagemaker_space":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_sagemaker_studio_lifecycle_config":
		if val, ok := config["studio_lifecycle_config_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_training_job":
		if val, ok := config["training_job_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_user_profile":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_sagemaker_workforce":
		if val, ok := config["workforce_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sagemaker_workteam":
		if val, ok := config["workteam_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_savingsplans_savings_plan":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_scheduler_schedule":
		var parts []string
		if val, ok := config["group_name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_scheduler_schedule_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_schemas_discoverer":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_schemas_registry":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_schemas_registry_policy":
		if val, ok := config["registry_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_schemas_schema":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_secretsmanager_secret":
		// Computed ID format: the secret Amazon Resource Name (ARN)
		return ""
	case "aws_secretsmanager_secret_policy":
		// Computed ID format: the secret Amazon Resource Name (ARN)
		return ""
	case "aws_secretsmanager_secret_rotation":
		// Computed ID format: the secret Amazon Resource Name (ARN)
		return ""
	case "aws_secretsmanager_secret_version":
		// Computed ID format: the secret ID and version ID
		return ""
	case "aws_secretsmanager_tag":
		// Computed or complex ID format: the AWS Secrets Manager secret identifier and key, separated by a comma (`,`)
		return ""
	case "aws_security_group":
		// Computed or complex ID format: the security group `id`
		return ""
	case "aws_security_group_rule":
		// No standard import format found in documentation
		return ""
	case "aws_securityhub_account":
		// Computed ID format: the AWS account ID
		return ""
	case "aws_securityhub_account_v2":
		// Computed or complex ID format: `arn`
		return ""
	case "aws_securityhub_action_target":
		// Computed or complex ID format: `arn`
		return ""
	case "aws_securityhub_automation_rule":
		// Computed or complex ID format: `arn`
		return ""
	case "aws_securityhub_configuration_policy":
		// Computed or complex ID format: `id`
		return ""
	case "aws_securityhub_configuration_policy_association":
		if val, ok := config["target_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_securityhub_finding_aggregator":
		// Computed or complex ID format: `arn`
		return ""
	case "aws_securityhub_insight":
		// Computed or complex ID format: `arn`
		return ""
	case "aws_securityhub_invite_accepter":
		// Computed ID format: the account ID
		return ""
	case "aws_securityhub_member":
		if val, ok := config["account_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_securityhub_organization_admin_account":
		if val, ok := config["admin_account_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_securityhub_organization_configuration":
		// Computed ID format: the AWS account ID
		return ""
	case "aws_securityhub_product_subscription":
		if val, ok := config["product_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_securityhub_standards_control":
		if val, ok := config["standards_control_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_securityhub_standards_control_association":
		if val, ok := config["security_control_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_securityhub_standards_subscription":
		// Computed or complex ID format: `arn`
		return ""
	case "aws_securitylake_aws_log_source":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_securitylake_custom_log_source":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_securitylake_data_lake":
		// Computed ID format: the standards subscription ARN
		return ""
	case "aws_securitylake_subscriber":
		// Computed ID format: the subscriber ID
		return ""
	case "aws_securitylake_subscriber_notification":
		// No standard import format found in documentation
		return ""
	case "aws_serverlessapplicationrepository_cloudformation_stack":
		// Computed or complex ID format: the CloudFormation Stack name (with or without the `serverlessrepo-` prefix) or the CloudFormation Stack ID
		return ""
	case "aws_service_discovery_http_namespace":
		// Computed ID format: the namespace ID
		return ""
	case "aws_service_discovery_instance":
		// Computed ID format: the service ID and instance ID
		return ""
	case "aws_service_discovery_private_dns_namespace":
		// Computed ID format: the namespace ID and VPC ID
		return ""
	case "aws_service_discovery_public_dns_namespace":
		// Computed ID format: the namespace ID
		return ""
	case "aws_service_discovery_service":
		// Computed ID format: the service ID
		return ""
	case "aws_servicecatalog_budget_resource_association":
		// Computed ID format: the budget name and resource ID
		return ""
	case "aws_servicecatalog_constraint":
		// Computed ID format: the constraint ID
		return ""
	case "aws_servicecatalog_organizations_access":
		// No standard import format found in documentation
		return ""
	case "aws_servicecatalog_portfolio":
		// Computed or complex ID format: the Service Catalog Portfolio `id`
		return ""
	case "aws_servicecatalog_portfolio_share":
		// Computed ID format: the portfolio share ID
		return ""
	case "aws_servicecatalog_principal_portfolio_association":
		if val, ok := config["accept_language"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_servicecatalog_product":
		// Computed ID format: the product ID
		return ""
	case "aws_servicecatalog_product_portfolio_association":
		// Computed ID format: the accept language, portfolio ID, and product ID
		return ""
	case "aws_servicecatalog_provisioned_product":
		// Computed ID format: the provisioned product ID
		return ""
	case "aws_servicecatalog_provisioning_artifact":
		// Computed ID format: the provisioning artifact ID and product ID separated by a colon
		return ""
	case "aws_servicecatalog_service_action":
		// Computed ID format: the service action ID
		return ""
	case "aws_servicecatalog_tag_option":
		// Computed ID format: the tag option ID
		return ""
	case "aws_servicecatalog_tag_option_resource_association":
		// Computed ID format: the tag option ID and resource ID
		return ""
	case "aws_servicecatalogappregistry_application":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_servicecatalogappregistry_attribute_group":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_servicecatalogappregistry_attribute_group_association":
		if val, ok := config["application_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_servicequotas_auto_management":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_servicequotas_service_quota":
		// Computed or complex ID format: the service code and quota code, separated by a front slash (`/`)
		return ""
	case "aws_servicequotas_template":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_servicequotas_template_association":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ses_active_receipt_rule_set":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ses_configuration_set":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ses_domain_dkim":
		if val, ok := config["domain"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ses_domain_identity":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ses_domain_identity_verification":
		// No standard import format found in documentation
		return ""
	case "aws_ses_domain_mail_from":
		if val, ok := config["domain"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ses_email_identity":
		// Unknown ID format: the email address
		return ""
	case "aws_ses_event_destination":
		if val, ok := config["configuration_set_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ses_identity_notification_topic":
		// Computed or complex ID format: the ID of the record. The ID is made up as `IDENTITY|TYPE` where `IDENTITY` is the SES Identity and `TYPE` is the Notification Type
		return ""
	case "aws_ses_identity_policy":
		// Computed or complex ID format: the identity and policy name, separated by a pipe character (`|`)
		return ""
	case "aws_ses_receipt_filter":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ses_receipt_rule":
		// Computed or complex ID format: the ruleset name and rule name separated by `:`
		return ""
	case "aws_ses_receipt_rule_set":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ses_template":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sesv2_account_suppression_attributes":
		// Computed ID format: the account ID
		return ""
	case "aws_sesv2_account_vdm_attributes":
		// Computed or complex ID format: the word `ses-account-vdm-attributes`
		return ""
	case "aws_sesv2_configuration_set":
		if val, ok := config["configuration_set_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sesv2_configuration_set_event_destination":
		// Computed or complex ID format: the `id` (`configuration_set_name|event_destination_name`)
		return ""
	case "aws_sesv2_contact_list":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_sesv2_dedicated_ip_assignment":
		// Computed or complex ID format: the `id`, which is a comma-separated string made up of `ip` and `destination_pool_name`
		return ""
	case "aws_sesv2_dedicated_ip_pool":
		if val, ok := config["pool_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sesv2_email_identity":
		if val, ok := config["email_identity"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sesv2_email_identity_feedback_attributes":
		if val, ok := config["email_identity"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sesv2_email_identity_mail_from_attributes":
		if val, ok := config["email_identity"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sesv2_email_identity_policy":
		if val, ok := config["email_identity"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sesv2_tenant":
		if val, ok := config["tenant_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sesv2_tenant_resource_association":
		if val, ok := config["tenant_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sfn_activity":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_sfn_alias":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_sfn_state_machine":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_shield_application_layer_automatic_response":
		// No standard import format found in documentation
		return ""
	case "aws_shield_drt_access_log_bucket_association":
		if val, ok := config["log_bucket"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_shield_drt_access_role_arn_association":
		// Computed ID format: the AWS account ID
		return ""
	case "aws_shield_proactive_engagement":
		// Computed ID format: the AWS account ID
		return ""
	case "aws_shield_protection":
		// Computed ID format: specifying their ID
		return ""
	case "aws_shield_protection_group":
		// Computed ID format: their protection group id
		return ""
	case "aws_shield_protection_health_check_association":
		if val, ok := config["shield_protection_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_shield_subscription":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_signer_signing_job":
		if val, ok := config["job_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_signer_signing_profile":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_signer_signing_profile_permission":
		// Computed ID format: profile_name/statement_id
		return ""
	case "aws_snapshot_create_volume_permission":
		// No standard import format found in documentation
		return ""
	case "aws_sns_platform_application":
		// Computed ID format: the ARN
		return ""
	case "aws_sns_sms_preferences":
		// No standard import format found in documentation
		return ""
	case "aws_sns_topic":
		// Computed or complex ID format: the topic `arn`
		return ""
	case "aws_sns_topic_data_protection_policy":
		// Computed ID format: the topic ARN
		return ""
	case "aws_sns_topic_policy":
		// Computed ID format: the topic ARN
		return ""
	case "aws_sns_topic_subscription":
		// Computed or complex ID format: the subscription `arn`
		return ""
	case "aws_spot_datafeed_subscription":
		// Computed or complex ID format: the word `spot-datafeed-subscription`
		return ""
	case "aws_spot_fleet_request":
		// Computed or complex ID format: `id`
		return ""
	case "aws_spot_instance_request":
		// No standard import format found in documentation
		return ""
	case "aws_sqs_queue":
		if val, ok := config["url"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_sqs_queue_policy":
		// Unknown ID format: the queue URL
		return ""
	case "aws_sqs_queue_redrive_allow_policy":
		// Unknown ID format: the queue URL
		return ""
	case "aws_sqs_queue_redrive_policy":
		// Unknown ID format: the queue URL
		return ""
	case "aws_ssm_activation":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ssm_association":
		if val, ok := config["association_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ssm_default_patch_baseline":
		// No standard import format found in documentation
		return ""
	case "aws_ssm_document":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ssm_maintenance_window":
		// Computed or complex ID format: the maintenance window `id`
		return ""
	case "aws_ssm_maintenance_window_target":
		// Computed or complex ID format: `WINDOW_ID/WINDOW_TARGET_ID`
		return ""
	case "aws_ssm_maintenance_window_task":
		if val, ok := config["window_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ssm_parameter":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ssm_patch_baseline":
		// Computed ID format: their baseline ID
		return ""
	case "aws_ssm_patch_group":
		if val, ok := config["patch_group"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ssm_resource_data_sync":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ssm_service_setting":
		if val, ok := config["setting_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ssmcontacts_contact":
		// Computed or complex ID format: the `ARN`
		return ""
	case "aws_ssmcontacts_contact_channel":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_ssmcontacts_plan":
		// Computed ID format: the Contact ARN
		return ""
	case "aws_ssmcontacts_rotation":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_ssmincidents_replication_set":
		// No standard import format found in documentation
		return ""
	case "aws_ssmincidents_response_plan":
		// Computed ID format: the response plan ARN. You can find the response plan ARN in the AWS Management Console
		return ""
	case "aws_ssmquicksetup_configuration_manager":
		if val, ok := config["manager_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ssoadmin_account_assignment":
		if val, ok := config["principal_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ssoadmin_application":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ssoadmin_application_access_scope":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ssoadmin_application_assignment":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ssoadmin_application_assignment_configuration":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_ssoadmin_customer_managed_policy_attachment":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ssoadmin_customer_managed_policy_attachments_exclusive":
		if val, ok := config["instance_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ssoadmin_instance_access_control_attributes":
		if val, ok := config["instance_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ssoadmin_managed_policy_attachment":
		if val, ok := config["managed_policy_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ssoadmin_managed_policy_attachments_exclusive":
		if val, ok := config["instance_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ssoadmin_permission_set":
		// Computed or complex ID format: the `arn` and `instance_arn` separated by a comma (`,`)
		return ""
	case "aws_ssoadmin_permission_set_inline_policy":
		if val, ok := config["permission_set_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ssoadmin_permissions_boundary_attachment":
		if val, ok := config["permission_set_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_ssoadmin_trusted_token_issuer":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_storagegateway_cache":
		// Computed or complex ID format: the gateway Amazon Resource Name (ARN) and local disk identifier separated with a colon (`:`)
		return ""
	case "aws_storagegateway_cached_iscsi_volume":
		// Computed ID format: the volume Amazon Resource Name (ARN)
		return ""
	case "aws_storagegateway_file_system_association":
		// Computed ID format: the FSx file system association Amazon Resource Name (ARN)
		return ""
	case "aws_storagegateway_gateway":
		// Computed ID format: the gateway Amazon Resource Name (ARN)
		return ""
	case "aws_storagegateway_nfs_file_share":
		// Computed ID format: the NFS File Share Amazon Resource Name (ARN)
		return ""
	case "aws_storagegateway_smb_file_share":
		// Computed ID format: the SMB File Share Amazon Resource Name (ARN)
		return ""
	case "aws_storagegateway_stored_iscsi_volume":
		// Computed ID format: the volume Amazon Resource Name (ARN)
		return ""
	case "aws_storagegateway_tape_pool":
		// Computed ID format: the volume Amazon Resource Name (ARN)
		return ""
	case "aws_storagegateway_upload_buffer":
		// Computed or complex ID format: the gateway Amazon Resource Name (ARN) and local disk identifier separated with a colon (`:`)
		return ""
	case "aws_storagegateway_working_storage":
		// Computed or complex ID format: the gateway Amazon Resource Name (ARN) and local disk identifier separated with a colon (`:`)
		return ""
	case "aws_subnet":
		// Computed or complex ID format: the subnet `id`
		return ""
	case "aws_swf_domain":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_synthetics_canary":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_synthetics_group":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_synthetics_group_association":
		var parts []string
		if val, ok := config["canary_arn"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["group_name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_timestreaminfluxdb_db_cluster":
		// Computed ID format: its identifier
		return ""
	case "aws_timestreaminfluxdb_db_instance":
		// Computed ID format: its identifier
		return ""
	case "aws_timestreamquery_scheduled_query":
		// Computed or complex ID format: the `arn`
		return ""
	case "aws_timestreamwrite_database":
		if val, ok := config["database_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_timestreamwrite_table":
		if val, ok := config["table_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_transcribe_language_model":
		if val, ok := config["model_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_transcribe_medical_vocabulary":
		if val, ok := config["vocabulary_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_transcribe_vocabulary":
		if val, ok := config["vocabulary_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_transcribe_vocabulary_filter":
		if val, ok := config["vocabulary_filter_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_transfer_access":
		if val, ok := config["server_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_transfer_agreement":
		var parts []string
		if val, ok := config["server_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["agreement_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_transfer_certificate":
		if val, ok := config["certificate_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_transfer_connector":
		if val, ok := config["connector_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_transfer_host_key":
		if val, ok := config["server_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_transfer_profile":
		if val, ok := config["profile_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_transfer_server":
		// Computed or complex ID format: the server `id`
		return ""
	case "aws_transfer_ssh_key":
		if val, ok := config["server_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_transfer_tag":
		// Computed or complex ID format: the Transfer Family resource identifier and key, separated by a comma (`,`)
		return ""
	case "aws_transfer_user":
		if val, ok := config["server_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_transfer_web_app":
		if val, ok := config["web_app_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_transfer_web_app_customization":
		if val, ok := config["web_app_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_transfer_workflow":
		if val, ok := config["worflow_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_uxc_account_customizations":
		// Computed ID format: the AWS account ID
		return ""
	case "aws_verifiedaccess_endpoint":
		// Computed or complex ID format: the  `id`
		return ""
	case "aws_verifiedaccess_group":
		// No standard import format found in documentation
		return ""
	case "aws_verifiedaccess_instance":
		// Computed or complex ID format: the  `id`
		return ""
	case "aws_verifiedaccess_instance_logging_configuration":
		// Computed or complex ID format: the Verified Access Instance `id`
		return ""
	case "aws_verifiedaccess_instance_trust_provider_attachment":
		if val, ok := config["verifiedaccess_instance_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_verifiedaccess_trust_provider":
		// Computed or complex ID format: the  `id`
		return ""
	case "aws_verifiedpermissions_identity_source":
		// Computed or complex ID format: the `policy_store_id:identity_source_id`
		return ""
	case "aws_verifiedpermissions_policy":
		var parts []string
		if val, ok := config["policy_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["policy_store_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_verifiedpermissions_policy_store":
		if val, ok := config["policy_store_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_verifiedpermissions_policy_template":
		// Computed or complex ID format: the `policy_store_id:policy_template_id`
		return ""
	case "aws_verifiedpermissions_schema":
		if val, ok := config["policy_store_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_volume_attachment":
		// Computed or complex ID format: `DEVICE_NAME:VOLUME_ID:INSTANCE_ID`
		return ""
	case "aws_vpc":
		// Computed or complex ID format: the VPC `id`
		return ""
	case "aws_vpc_block_public_access_exclusion":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_vpc_block_public_access_options":
		if val, ok := config["aws_region"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_vpc_dhcp_options":
		// Computed or complex ID format: the DHCP Options `id`
		return ""
	case "aws_vpc_dhcp_options_association":
		// Computed ID format: the VPC ID associated with the options
		return ""
	case "aws_vpc_encryption_control":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_vpc_endpoint":
		// Computed or complex ID format: the VPC endpoint `id`
		return ""
	case "aws_vpc_endpoint_connection_accepter":
		// Computed or complex ID format: ID of the connection, which is the `VPC Endpoint Service ID` and `VPC Endpoint ID` separated by underscore (`_`).
		return ""
	case "aws_vpc_endpoint_connection_notification":
		// Computed or complex ID format: the VPC endpoint connection notification `id`
		return ""
	case "aws_vpc_endpoint_policy":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_vpc_endpoint_private_dns":
		if val, ok := config["vpc_endpoint_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_vpc_endpoint_route_table_association":
		if val, ok := config["vpc_endpoint_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_vpc_endpoint_security_group_association":
		if val, ok := config["vpc_endpoint_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_vpc_endpoint_service":
		// Computed or complex ID format: the VPC endpoint service `id`
		return ""
	case "aws_vpc_endpoint_service_allowed_principal":
		// No standard import format found in documentation
		return ""
	case "aws_vpc_endpoint_service_private_dns_verification":
		// No standard import format found in documentation
		return ""
	case "aws_vpc_endpoint_subnet_association":
		if val, ok := config["vpc_endpoint_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_vpc_ipam":
		// Computed or complex ID format: the IPAM `id`
		return ""
	case "aws_vpc_ipam_organization_admin_account":
		// Computed or complex ID format: the delegate account `id`
		return ""
	case "aws_vpc_ipam_pool":
		// Computed or complex ID format: the IPAM pool `id`
		return ""
	case "aws_vpc_ipam_pool_cidr":
		// Computed or complex ID format: the `<cidr>_<ipam-pool-id>`
		return ""
	case "aws_vpc_ipam_pool_cidr_allocation":
		// Computed or complex ID format: the allocation `id` and `pool id`, separated by `_`
		return ""
	case "aws_vpc_ipam_preview_next_cidr":
		// No standard import format found in documentation
		return ""
	case "aws_vpc_ipam_resource_discovery":
		// Computed or complex ID format: the IPAM resource discovery `id`
		return ""
	case "aws_vpc_ipam_resource_discovery_association":
		// Computed or complex ID format: the IPAM resource discovery association `id`
		return ""
	case "aws_vpc_ipam_scope":
		if val, ok := config["scope_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_vpc_ipv4_cidr_block_association":
		// Computed ID format: the VPC CIDR association ID and optionally the IPv4 IPAM pool ID and netmask length
		return ""
	case "aws_vpc_ipv6_cidr_block_association":
		// Computed ID format: the VPC CIDR association ID and optionally the IPv6 IPAM pool ID and netmask length
		return ""
	case "aws_vpc_network_performance_metric_subscription":
		// No standard import format found in documentation
		return ""
	case "aws_vpc_peering_connection":
		// Computed or complex ID format: the VPC peering `id`
		return ""
	case "aws_vpc_peering_connection_accepter":
		// Computed ID format: the Peering Connection ID
		return ""
	case "aws_vpc_peering_connection_options":
		// Computed or complex ID format: the VPC peering `id`
		return ""
	case "aws_vpc_route_server":
		if val, ok := config["route_server_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_vpc_route_server_endpoint":
		if val, ok := config["route_server_endpoint_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_vpc_route_server_peer":
		if val, ok := config["route_server_peer_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_vpc_route_server_propagation":
		// No standard import format found in documentation
		return ""
	case "aws_vpc_route_server_vpc_association":
		// No standard import format found in documentation
		return ""
	case "aws_vpc_security_group_egress_rule":
		if val, ok := config["security_group_rule_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_vpc_security_group_ingress_rule":
		if val, ok := config["security_group_rule_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_vpc_security_group_rules_exclusive":
		if val, ok := config["security_group_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_vpc_security_group_vpc_association":
		if val, ok := config["security_group_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_vpclattice_access_log_subscription":
		// Computed ID format: the access log subscription ID
		return ""
	case "aws_vpclattice_auth_policy":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_vpclattice_domain_verification":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_vpclattice_listener":
		if val, ok := config["listener_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_vpclattice_listener_rule":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_vpclattice_resource_configuration":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_vpclattice_resource_gateway":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_vpclattice_resource_policy":
		if val, ok := config["resource_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_vpclattice_service":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_vpclattice_service_network":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_vpclattice_service_network_resource_association":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_vpclattice_service_network_service_association":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_vpclattice_service_network_vpc_association":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_vpclattice_target_group":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_vpclattice_target_group_attachment":
		// No standard import format found in documentation
		return ""
	case "aws_vpn_concentrator":
		// Computed ID format: the VPN concentrator ID
		return ""
	case "aws_vpn_connection":
		// Computed or complex ID format: the VPN connection `id`
		return ""
	case "aws_vpn_connection_route":
		// No standard import format found in documentation
		return ""
	case "aws_vpn_gateway":
		// Computed or complex ID format: the VPN gateway `id`
		return ""
	case "aws_vpn_gateway_attachment":
		// No standard import format found in documentation
		return ""
	case "aws_vpn_gateway_route_propagation":
		// No standard import format found in documentation
		return ""
	case "aws_waf_byte_match_set":
		// Computed ID format: the id
		return ""
	case "aws_waf_geo_match_set":
		// Computed ID format: their ID
		return ""
	case "aws_waf_ipset":
		// Computed ID format: their ID
		return ""
	case "aws_waf_rate_based_rule":
		// Computed ID format: the id
		return ""
	case "aws_waf_regex_match_set":
		// Computed ID format: their ID
		return ""
	case "aws_waf_regex_pattern_set":
		// Computed ID format: their ID
		return ""
	case "aws_waf_rule":
		// Computed ID format: the id
		return ""
	case "aws_waf_rule_group":
		// Computed ID format: the id
		return ""
	case "aws_waf_size_constraint_set":
		// Computed ID format: their ID
		return ""
	case "aws_waf_sql_injection_match_set":
		// Computed ID format: their ID
		return ""
	case "aws_waf_web_acl":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_waf_xss_match_set":
		// Computed ID format: their ID
		return ""
	case "aws_wafregional_byte_match_set":
		// Computed ID format: the id
		return ""
	case "aws_wafregional_geo_match_set":
		// Computed ID format: the id
		return ""
	case "aws_wafregional_ipset":
		// Computed ID format: their ID
		return ""
	case "aws_wafregional_rate_based_rule":
		// Computed ID format: the id
		return ""
	case "aws_wafregional_regex_match_set":
		// Computed ID format: the id
		return ""
	case "aws_wafregional_regex_pattern_set":
		// Computed ID format: the id
		return ""
	case "aws_wafregional_rule":
		// Computed ID format: the id
		return ""
	case "aws_wafregional_rule_group":
		// Computed ID format: the id
		return ""
	case "aws_wafregional_size_constraint_set":
		// Computed ID format: the id
		return ""
	case "aws_wafregional_sql_injection_match_set":
		// Computed ID format: the id
		return ""
	case "aws_wafregional_web_acl":
		// Computed ID format: the id
		return ""
	case "aws_wafregional_web_acl_association":
		// Computed or complex ID format: their `web_acl_id:resource_arn`
		return ""
	case "aws_wafregional_xss_match_set":
		// Computed or complex ID format: the `id`
		return ""
	case "aws_wafv2_api_key":
		var parts []string
		if val, ok := config["api_key"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["scope"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_wafv2_ip_set":
		// Computed or complex ID format: `ID/name/scope`
		return ""
	case "aws_wafv2_regex_pattern_set":
		// Computed or complex ID format: `ID/name/scope`
		return ""
	case "aws_wafv2_rule_group":
		// Computed or complex ID format: `ID/name/scope`
		return ""
	case "aws_wafv2_web_acl":
		// Computed or complex ID format: `ID/Name/Scope`
		return ""
	case "aws_wafv2_web_acl_association":
		// Computed or complex ID format: `WEB_ACL_ARN,RESOURCE_ARN`
		return ""
	case "aws_wafv2_web_acl_logging_configuration":
		// Computed ID format: the ARN of the WAFv2 Web ACL
		return ""
	case "aws_wafv2_web_acl_rule":
		if val, ok := config["web_acl_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_wafv2_web_acl_rule_group_association":
		// Computed or complex ID format: `WebACLARN,RuleName,RuleGroupType,RuleGroupARN`, where `RuleGroupType` is `custom`
		return ""
	case "aws_workmail_default_domain":
		// Computed ID format: the organization ID
		return ""
	case "aws_workmail_domain":
		var parts []string
		if val, ok := config["organization_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["domain_name"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_workmail_group":
		var parts []string
		if val, ok := config["organization_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["group_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_workmail_organization":
		if val, ok := config["organization_id"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_workmail_user":
		var parts []string
		if val, ok := config["organization_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["user_id"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_workspaces_connection_alias":
		// Computed ID format: the connection alias ID
		return ""
	case "aws_workspaces_directory":
		// Computed ID format: the directory ID
		return ""
	case "aws_workspaces_ip_group":
		// Computed ID format: their GroupID
		return ""
	case "aws_workspaces_workspace":
		// Computed ID format: their ID
		return ""
	case "aws_workspacesweb_browser_settings":
		if val, ok := config["browser_settings_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_workspacesweb_browser_settings_association":
		var parts []string
		if val, ok := config["browser_settings_arn"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["portal_arn"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_workspacesweb_data_protection_settings":
		if val, ok := config["data_protection_settings_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_workspacesweb_data_protection_settings_association":
		// No standard import format found in documentation
		return ""
	case "aws_workspacesweb_identity_provider":
		if val, ok := config["identity_provider_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_workspacesweb_ip_access_settings":
		if val, ok := config["ip_access_settings_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_workspacesweb_ip_access_settings_association":
		// No standard import format found in documentation
		return ""
	case "aws_workspacesweb_network_settings":
		if val, ok := config["network_settings_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_workspacesweb_network_settings_association":
		// No standard import format found in documentation
		return ""
	case "aws_workspacesweb_portal":
		if val, ok := config["portal_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_workspacesweb_session_logger":
		if val, ok := config["session_logger_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_workspacesweb_session_logger_association":
		var parts []string
		if val, ok := config["session_logger_arn"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["portal_arn"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_workspacesweb_trust_store":
		if val, ok := config["trust_store_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_workspacesweb_trust_store_association":
		var parts []string
		if val, ok := config["trust_store_arn"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		if val, ok := config["portal_arn"].(string); ok && val != "" {
			parts = append(parts, val)
		} else {
			return ""
		}
		// Note: Composite separator might need adjustment based on exact docs
		return strings.Join(parts, ",")
	case "aws_workspacesweb_user_access_logging_settings":
		if val, ok := config["user_access_logging_settings_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_workspacesweb_user_access_logging_settings_association":
		// No standard import format found in documentation
		return ""
	case "aws_workspacesweb_user_settings":
		if val, ok := config["user_settings_arn"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_workspacesweb_user_settings_association":
		// No standard import format found in documentation
		return ""
	case "aws_xray_encryption_config":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_xray_group":
		// Computed ID format: the ARN
		return ""
	case "aws_xray_resource_policy":
		if val, ok := config["policy_name"].(string); ok && val != "" {
			return val
		}
		return ""
	case "aws_xray_sampling_rule":
		if val, ok := config["name"].(string); ok && val != "" {
			return val
		}
		return ""
	}
	return ""
}
