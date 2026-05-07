package providers

// extractAzurermImportID returns the necessary import ID for a azurerm resource
// based on its configuration extracted from the terraform plan.
func extractAzurermImportID(ctx *ProviderContext, resourceType string, config map[string]any) string {
	// First, check if there's a custom resolver for this resource
	if id := resolveCustomextractAzurermImportID(ctx, resourceType, config); id != "" {
		return id
	}

	switch resourceType {
	case "azurerm_aadb2c_directory":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureActiveDirectory/b2cDirectories/directory-name
		return ""
	case "azurerm_active_directory_domain_service":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AAD/domainServices/instance1/initialReplicaSetId/00000000-0000-0000-0000-000000000000
		return ""
	case "azurerm_active_directory_domain_service_replica_set":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AAD/domainServices/instance1/replicaSets/00000000-0000-0000-0000-000000000000
		return ""
	case "azurerm_active_directory_domain_service_trust":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AAD/domainServices/DomainService1/trusts/trust1
		return ""
	case "azurerm_advanced_threat_protection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/exampleResourceGroup/providers/Microsoft.Storage/storageAccounts/exampleaccount/providers/Microsoft.Security/advancedThreatProtectionSettings/default
		return ""
	case "azurerm_advisor_suppression":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Advisor/recommendations/00000000-0000-0000-0000-000000000000/suppressions/name
		return ""
	case "azurerm_ai_foundry":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.MachineLearningServices/workspaces/hub1
		return ""
	case "azurerm_ai_foundry_project":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.MachineLearningServices/workspaces/project1
		return ""
	case "azurerm_ai_services":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.CognitiveServices/accounts/account1
		return ""
	case "azurerm_analysis_services_server":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AnalysisServices/servers/server1
		return ""
	case "azurerm_api_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.Web/connections/example-connection
		return ""
	case "azurerm_api_management":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1
		return ""
	case "azurerm_api_management_api":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/apis/api1;rev=1
		return ""
	case "azurerm_api_management_api_diagnostic":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/instance1/apis/api1/diagnostics/diagnostic1
		return ""
	case "azurerm_api_management_api_operation":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/apis/api1/operations/operation1
		return ""
	case "azurerm_api_management_api_operation_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/instance1/apis/api1/operations/operation1
		return ""
	case "azurerm_api_management_api_operation_tag":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/apis/api1/operations/operation1/tags/tag1
		return ""
	case "azurerm_api_management_api_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/apis/exampleId
		return ""
	case "azurerm_api_management_api_release":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/apis/api1/releases/release1
		return ""
	case "azurerm_api_management_api_schema":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/apis/api1/schemas/schema1
		return ""
	case "azurerm_api_management_api_tag":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/apis/api1/tags/tag1
		return ""
	case "azurerm_api_management_api_tag_description":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/apis/api1/tagDescriptions/tagDescriptionId1
		return ""
	case "azurerm_api_management_api_version_set":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/apiVersionSets/set1
		return ""
	case "azurerm_api_management_authorization_server":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/authorizationServers/server1
		return ""
	case "azurerm_api_management_backend":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/backends/backend1
		return ""
	case "azurerm_api_management_certificate":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/certificates/certificate1
		return ""
	case "azurerm_api_management_custom_domain":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/customDomains/default
		return ""
	case "azurerm_api_management_diagnostic":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/diagnostics/applicationinsights
		return ""
	case "azurerm_api_management_email_template":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/instance1/templates/template1
		return ""
	case "azurerm_api_management_gateway":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/gateways/gateway1
		return ""
	case "azurerm_api_management_gateway_api":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/gateways/gateway1/apis/api1
		return ""
	case "azurerm_api_management_gateway_certificate_authority":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/gateways/gateway1/certificateAuthorities/cert1
		return ""
	case "azurerm_api_management_gateway_host_name_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/gateways/gateway1/hostnameConfigurations/hc1
		return ""
	case "azurerm_api_management_global_schema":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/schemas/schema1
		return ""
	case "azurerm_api_management_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.ApiManagement/service/example-apim/groups/example-apimg
		return ""
	case "azurerm_api_management_group_user":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/groups/groupId/users/user123
		return ""
	case "azurerm_api_management_identity_provider_aad":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/aad
		return ""
	case "azurerm_api_management_identity_provider_aadb2c":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/service1/identityProviders/aadB2C
		return ""
	case "azurerm_api_management_identity_provider_facebook":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/facebook
		return ""
	case "azurerm_api_management_identity_provider_google":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/google
		return ""
	case "azurerm_api_management_identity_provider_microsoft":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/microsoft
		return ""
	case "azurerm_api_management_identity_provider_twitter":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/twitter
		return ""
	case "azurerm_api_management_logger":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.ApiManagement/service/example-apim/loggers/example-logger
		return ""
	case "azurerm_api_management_named_value":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.ApiManagement/service/example-apim/namedValues/example-apimp
		return ""
	case "azurerm_api_management_notification_recipient_email":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/notifications/notificationName1/recipientEmails/email1
		return ""
	case "azurerm_api_management_notification_recipient_user":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/notifications/notificationName1/recipientUsers/userid1
		return ""
	case "azurerm_api_management_openid_connect_provider":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/openidConnectProviders/provider1
		return ""
	case "azurerm_api_management_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1
		return ""
	case "azurerm_api_management_policy_fragment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/instance1/policyFragments/policyFragment1
		return ""
	case "azurerm_api_management_product":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/products/myproduct
		return ""
	case "azurerm_api_management_product_api":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/products/exampleId/apis/apiId
		return ""
	case "azurerm_api_management_product_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/products/exampleId/groups/groupId
		return ""
	case "azurerm_api_management_product_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/products/product1
		return ""
	case "azurerm_api_management_product_tag":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/products/myproduct/tags/mytag
		return ""
	case "azurerm_api_management_redis_cache":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/caches/cache1
		return ""
	case "azurerm_api_management_standalone_gateway":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/gateways/gateway1
		return ""
	case "azurerm_api_management_subscription":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.ApiManagement/service/example-apim/subscriptions/subscription-name
		return ""
	case "azurerm_api_management_tag":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/tags/tag1
		return ""
	case "azurerm_api_management_user":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/users/abc123
		return ""
	case "azurerm_api_management_workspace":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/workspaces/workspace1
		return ""
	case "azurerm_api_management_workspace_api_version_set":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/workspaces/workspace1/apiVersionSets/set1
		return ""
	case "azurerm_api_management_workspace_certificate":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/workspaces/workspace1/certificates/certificate1
		return ""
	case "azurerm_api_management_workspace_named_value":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/workspaces/workspace1/namedValues/namedValue1
		return ""
	case "azurerm_api_management_workspace_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/workspaces/workspace1
		return ""
	case "azurerm_api_management_workspace_policy_fragment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/instance1/workspaces/workspace1/policyFragments/policyFragment1
		return ""
	case "azurerm_app_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1
		return ""
	case "azurerm_app_configuration_feature":
		// Azure Resource ID: https://appconfname1.azconfig.io/kv/.appconfig.featureflag%2FkeyName?label=labelName
		return ""
	case "azurerm_app_configuration_key":
		// Azure Resource ID: https://appconfname1.azconfig.io/kv/keyName?label=labelName
		return ""
	case "azurerm_app_service":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1
		return ""
	case "azurerm_app_service_active_slot":
		// No standard import format found in documentation
		return ""
	case "azurerm_app_service_certificate":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/certificates/certificate1
		return ""
	case "azurerm_app_service_certificate_binding":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/hostNameBindings/mywebsite.com|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/certificates/mywebsite.com"
		return ""
	case "azurerm_app_service_certificate_order":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.CertificateRegistration/certificateOrders/certificateorder1
		return ""
	case "azurerm_app_service_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Web/sites/webapp/providers/Microsoft.ServiceLinker/linkers/serviceconnector1
		return ""
	case "azurerm_app_service_custom_hostname_binding":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/hostNameBindings/mywebsite.com
		return ""
	case "azurerm_app_service_environment_v3":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Web/hostingEnvironments/myAppServiceEnv
		return ""
	case "azurerm_app_service_hybrid_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/exampleResourceGroup1/providers/Microsoft.Web/sites/exampleAppService1/hybridConnectionNamespaces/exampleRN1/relays/exampleRHC1
		return ""
	case "azurerm_app_service_managed_certificate":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Web/certificates/customhost.contoso.com
		return ""
	case "azurerm_app_service_plan":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/serverFarms/instance1
		return ""
	case "azurerm_app_service_public_certificate":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Web/sites/site1/publicCertificates/publicCertificate1
		return ""
	case "azurerm_app_service_slot":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/website1/slots/instance1
		return ""
	case "azurerm_app_service_slot_custom_hostname_binding":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/slots/staging/hostNameBindings/mywebsite.com
		return ""
	case "azurerm_app_service_slot_virtual_network_swift_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/slots/staging/config/virtualNetwork
		return ""
	case "azurerm_app_service_source_control":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1
		return ""
	case "azurerm_app_service_source_control_slot":
		// Azure Resource ID: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/slots/slot1"
		return ""
	case "azurerm_app_service_source_control_token":
		// Azure Resource ID: {type}
		return ""
	case "azurerm_app_service_virtual_network_swift_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/config/virtualNetwork
		return ""
	case "azurerm_application_gateway":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/applicationGateways/myGateway1
		return ""
	case "azurerm_application_insights":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Insights/components/instance1
		return ""
	case "azurerm_application_insights_analytics_item":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Insights/components/mycomponent1/analyticsItems/11111111-1111-1111-1111-111111111111
		return ""
	case "azurerm_application_insights_api_key":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Insights/components/instance1/apiKeys/00000000-0000-0000-0000-000000000000
		return ""
	case "azurerm_application_insights_smart_detection_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Insights/components/mycomponent1/proactiveDetectionConfigs/myrule1
		return ""
	case "azurerm_application_insights_standard_web_test":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Insights/webTests/appinsightswebtest
		return ""
	case "azurerm_application_insights_web_test":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Insights/webTests/my_test
		return ""
	case "azurerm_application_insights_workbook":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Insights/workbooks/resource1
		return ""
	case "azurerm_application_insights_workbook_template":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Insights/workbookTemplates/resource1
		return ""
	case "azurerm_application_load_balancer":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceNetworking/trafficControllers/alb1
		return ""
	case "azurerm_application_load_balancer_frontend":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceNetworking/trafficControllers/alb1/frontends/frontend1
		return ""
	case "azurerm_application_load_balancer_security_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ServiceNetworking/trafficControllers/alb/securityPolicies/sp1
		return ""
	case "azurerm_application_load_balancer_subnet_association":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ServiceNetworking/trafficControllers/alb1/associations/association1
		return ""
	case "azurerm_application_security_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/applicationSecurityGroups/securitygroup1
		return ""
	case "azurerm_arc_kubernetes_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Kubernetes/connectedClusters/cluster1
		return ""
	case "azurerm_arc_kubernetes_cluster_extension":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Kubernetes/connectedClusters/cluster1/providers/Microsoft.KubernetesConfiguration/extensions/extension1
		return ""
	case "azurerm_arc_kubernetes_flux_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Kubernetes/connectedClusters/cluster1/providers/Microsoft.KubernetesConfiguration/fluxConfigurations/fluxConfiguration1
		return ""
	case "azurerm_arc_kubernetes_provisioned_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/cluster1
		return ""
	case "azurerm_arc_machine":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.HybridCompute/machines/machine1
		return ""
	case "azurerm_arc_machine_automanage_configuration_assignment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.HybridCompute/machines/machine1/providers/Microsoft.AutoManage/configurationProfileAssignments/default
		return ""
	case "azurerm_arc_machine_extension":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.HybridCompute/machines/hcmachine1/extensions/ext1
		return ""
	case "azurerm_arc_private_link_scope":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScope1
		return ""
	case "azurerm_arc_resource_bridge_appliance":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ResourceConnector/appliances/appliancesExample
		return ""
	case "azurerm_attestation_provider":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Attestation/attestationProviders/provider1
		return ""
	case "azurerm_automanage_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AutoManage/configurationProfiles/configurationProfile1
		return ""
	case "azurerm_automation_account":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1
		return ""
	case "azurerm_automation_certificate":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/certificates/certificate1
		return ""
	case "azurerm_automation_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connections/conn1
		return ""
	case "azurerm_automation_connection_certificate":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connections/conn1
		return ""
	case "azurerm_automation_connection_classic_certificate":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connections/conn1
		return ""
	case "azurerm_automation_connection_service_principal":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connections/conn1
		return ""
	case "azurerm_automation_connection_type":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connectionTypes/type1
		return ""
	case "azurerm_automation_credential":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/credentials/credential1
		return ""
	case "azurerm_automation_dsc_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/configurations/configuration1
		return ""
	case "azurerm_automation_dsc_nodeconfiguration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/nodeConfigurations/configuration1
		return ""
	case "azurerm_automation_hybrid_runbook_worker":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/hybridRunbookWorkerGroups/group1/hybridRunbookWorkers/00000000-0000-0000-0000-000000000000
		return ""
	case "azurerm_automation_hybrid_runbook_worker_group":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/hybridRunbookWorkerGroups/grp1
		return ""
	case "azurerm_automation_job_schedule":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/schedules/schedule1|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/runbooks/runbook1"
		return ""
	case "azurerm_automation_module":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/modules/module1
		return ""
	case "azurerm_automation_powershell72_module":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/powerShell72Modules/module1
		return ""
	case "azurerm_automation_python3_package":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/python3Packages/pkg
		return ""
	case "azurerm_automation_runbook":
		// No standard import format found in documentation
		return ""
	case "azurerm_automation_runtime_environment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/runtimeEnvironments/env1
		return ""
	case "azurerm_automation_runtime_environment_package":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/automationAccount1/runtimeEnvironments/runtimeEnvironment1/packages/package1
		return ""
	case "azurerm_automation_schedule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/schedules/schedule1
		return ""
	case "azurerm_automation_software_update_configuration":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/softwareUpdateConfigurations/suc1
		return ""
	case "azurerm_automation_source_control":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/sourceControls/sc1
		return ""
	case "azurerm_automation_variable_bool":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tfex-example-rg/providers/Microsoft.Automation/automationAccounts/tfex-example-account/variables/tfex-example-var
		return ""
	case "azurerm_automation_variable_datetime":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tfex-example-rg/providers/Microsoft.Automation/automationAccounts/tfex-example-account/variables/tfex-example-var
		return ""
	case "azurerm_automation_variable_int":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tfex-example-rg/providers/Microsoft.Automation/automationAccounts/tfex-example-account/variables/tfex-example-var
		return ""
	case "azurerm_automation_variable_object":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tfex-example-rg/providers/Microsoft.Automation/automationAccounts/tfex-example-account/variables/tfex-example-var
		return ""
	case "azurerm_automation_variable_string":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tfex-example-rg/providers/Microsoft.Automation/automationAccounts/tfex-example-account/variables/tfex-example-var
		return ""
	case "azurerm_automation_watcher":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/watchers/watch1
		return ""
	case "azurerm_automation_webhook":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/webHooks/TestRunbook_webhook
		return ""
	case "azurerm_availability_set":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/availabilitySets/webAvailSet
		return ""
	case "azurerm_backup_container_storage_account":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/backupFabrics/Azure/protectionContainers/StorageContainer;storage;storage-rg-name;storage-account"
		return ""
	case "azurerm_backup_policy_file_share":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/example-recovery-vault/backupPolicies/policy1
		return ""
	case "azurerm_backup_policy_vm":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/example-recovery-vault/backupPolicies/policy1
		return ""
	case "azurerm_backup_policy_vm_workload":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/vault1/backupPolicies/policy1
		return ""
	case "azurerm_backup_protected_file_share":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/example-recovery-vault/backupFabrics/Azure/protectionContainers/StorageContainer;storage;group2;example-storage-account/protectedItems/AzureFileShare;3f6e3108a45793581bcbd1c61c87a3b2ceeb4ff4bc02a95ce9d1022b23722935"
		return ""
	case "azurerm_backup_protected_vm":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/example-recovery-vault/backupFabrics/Azure/protectionContainers/iaasvmcontainer;iaasvmcontainerv2;group1;vm1/protectedItems/vm;iaasvmcontainerv2;group1;vm1"
		return ""
	case "azurerm_bastion_host":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/bastionHosts/instance1
		return ""
	case "azurerm_batch_account":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Batch/batchAccounts/account1
		return ""
	case "azurerm_batch_application":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Batch/batchAccounts/exampleba/applications/example-batch-application
		return ""
	case "azurerm_batch_certificate":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Batch/batchAccounts/batch1/certificates/certificate1
		return ""
	case "azurerm_batch_job":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Batch/batchAccounts/account1/pools/pool1/jobs/job1
		return ""
	case "azurerm_batch_pool":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.Batch/batchAccounts/myBatchAccount1/pools/myBatchPool1
		return ""
	case "azurerm_billing_account_cost_management_export":
		// Azure Resource ID: /providers/Microsoft.Billing/billingAccounts/12345678/providers/Microsoft.CostManagement/exports/export1
		return ""
	case "azurerm_blueprint_assignment":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Blueprint/blueprintAssignments/assignSimpleBlueprint"
		return ""
	case "azurerm_bot_channel_alexa":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.BotService/botServices/botService1/channels/AlexaChannel
		return ""
	case "azurerm_bot_channel_direct_line_speech":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.BotService/botServices/botService1/channels/DirectLineSpeechChannel
		return ""
	case "azurerm_bot_channel_directline":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example/channels/DirectlineChannel
		return ""
	case "azurerm_bot_channel_email":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example/channels/EmailChannel
		return ""
	case "azurerm_bot_channel_facebook":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.BotService/botServices/botService1/channels/FacebookChannel
		return ""
	case "azurerm_bot_channel_line":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.BotService/botServices/botService1/channels/LineChannel
		return ""
	case "azurerm_bot_channel_ms_teams":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example/channels/MsTeamsChannel
		return ""
	case "azurerm_bot_channel_slack":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example/channels/SlackChannel
		return ""
	case "azurerm_bot_channel_sms":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.BotService/botServices/botService1/channels/SmsChannel
		return ""
	case "azurerm_bot_channel_web_chat":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.BotService/botServices/botService1/channels/WebChatChannel
		return ""
	case "azurerm_bot_channels_registration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example
		return ""
	case "azurerm_bot_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example/connections/example
		return ""
	case "azurerm_bot_service_azure_bot":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.BotService/botServices/botService1
		return ""
	case "azurerm_bot_web_app":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example
		return ""
	case "azurerm_capacity_reservation":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/capacityReservationGroups/capacityReservationGroup1/capacityReservations/capacityReservation1
		return ""
	case "azurerm_capacity_reservation_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/capacityReservationGroups/capacityReservationGroup1
		return ""
	case "azurerm_cdn_endpoint":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Cdn/profiles/myprofile1/endpoints/myendpoint1
		return ""
	case "azurerm_cdn_endpoint_custom_domain":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/customDomains/domain1
		return ""
	case "azurerm_cdn_frontdoor_custom_domain":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Cdn/profiles/profile1/customDomains/customDomain1
		return ""
	case "azurerm_cdn_frontdoor_custom_domain_association":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Cdn/profiles/profile1/associations/assoc1
		return ""
	case "azurerm_cdn_frontdoor_endpoint":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Cdn/profiles/profile1/afdEndpoints/endpoint1
		return ""
	case "azurerm_cdn_frontdoor_firewall_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/firewallPolicy1
		return ""
	case "azurerm_cdn_frontdoor_origin":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Cdn/profiles/profile1/originGroups/originGroup1/origins/origin1
		return ""
	case "azurerm_cdn_frontdoor_origin_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Cdn/profiles/profile1/originGroups/originGroup1
		return ""
	case "azurerm_cdn_frontdoor_profile":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Cdn/profiles/myprofile1
		return ""
	case "azurerm_cdn_frontdoor_route":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Cdn/profiles/profile1/afdEndpoints/endpoint1/routes/route1
		return ""
	case "azurerm_cdn_frontdoor_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Cdn/profiles/profile1/ruleSets/ruleSet1/rules/rule1
		return ""
	case "azurerm_cdn_frontdoor_rule_set":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Cdn/profiles/profile1/ruleSets/ruleSet1
		return ""
	case "azurerm_cdn_frontdoor_secret":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Cdn/profiles/profile1/secrets/secrets1
		return ""
	case "azurerm_cdn_frontdoor_security_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Cdn/profiles/profile1/securityPolicies/policy1
		return ""
	case "azurerm_cdn_profile":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Cdn/profiles/myprofile1
		return ""
	case "azurerm_chaos_studio_capability":
		// Azure Resource ID: /{scope}/providers/Microsoft.Chaos/targets/{targetName}/capabilities/{capabilityName}
		return ""
	case "azurerm_chaos_studio_experiment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Chaos/experiments/experiment1
		return ""
	case "azurerm_chaos_studio_target":
		// Azure Resource ID: /{scope}/providers/Microsoft.Chaos/targets/{targetName}
		return ""
	case "azurerm_cognitive_account":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.CognitiveServices/accounts/account1
		return ""
	case "azurerm_cognitive_account_customer_managed_key":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.CognitiveServices/accounts/account1
		return ""
	case "azurerm_cognitive_account_project":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.CognitiveServices/accounts/account1/projects/project1
		return ""
	case "azurerm_cognitive_account_rai_blocklist":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.CognitiveServices/accounts/account1/raiBlocklists/raiblocklist1
		return ""
	case "azurerm_cognitive_account_rai_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.CognitiveServices/accounts/account1/raiPolicies/policy1
		return ""
	case "azurerm_cognitive_deployment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.CognitiveServices/accounts/account1/deployments/deployment1
		return ""
	case "azurerm_communication_service":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Communication/communicationServices/communicationService1
		return ""
	case "azurerm_communication_service_email_domain_association":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Communication/communicationServices/communicationService1|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Communication/emailServices/emailCommunicationService1/domains/domain1"
		return ""
	case "azurerm_confidential_ledger":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-group/providers/Microsoft.ConfidentialLedger/ledgers/example-ledger
		return ""
	case "azurerm_consumption_budget_management_group":
		// Azure Resource ID: /providers/Microsoft.Management/managementGroups/00000000-0000-0000-0000-000000000000/providers/Microsoft.Consumption/budgets/budget1
		return ""
	case "azurerm_consumption_budget_resource_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Consumption/budgets/resourceGroup1
		return ""
	case "azurerm_consumption_budget_subscription":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Consumption/budgets/subscription1
		return ""
	case "azurerm_container_app":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.App/containerApps/myContainerApp"
		return ""
	case "azurerm_container_app_custom_domain":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.App/containerApps/myContainerApp/customDomainName/mycustomdomain.example.com"
		return ""
	case "azurerm_container_app_environment":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.App/managedEnvironments/myEnvironment"
		return ""
	case "azurerm_container_app_environment_certificate":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.App/managedEnvironments/myenv/certificates/mycertificate"
		return ""
	case "azurerm_container_app_environment_custom_domain":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.App/managedEnvironments/myEnvironment"
		return ""
	case "azurerm_container_app_environment_dapr_component":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.App/managedEnvironments/myenv/daprComponents/mydaprcomponent"
		return ""
	case "azurerm_container_app_environment_managed_certificate":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.App/managedEnvironments/myenv/managedCertificates/mycertificate"
		return ""
	case "azurerm_container_app_environment_storage":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.App/managedEnvironments/myEnvironment/storages/mystorage"
		return ""
	case "azurerm_container_app_job":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.App/jobs/example-container-app-job"
		return ""
	case "azurerm_container_connected_registry":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.ContainerRegistry/registries/registry1/connectedRegistries/registry1
		return ""
	case "azurerm_container_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ContainerInstance/containerGroups/myContainerGroup1
		return ""
	case "azurerm_container_registry":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ContainerRegistry/registries/myregistry1
		return ""
	case "azurerm_container_registry_agent_pool":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.ContainerRegistry/registries/registry1/agentPools/agentpool1
		return ""
	case "azurerm_container_registry_cache_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/cacheRules/myCacheRule
		return ""
	case "azurerm_container_registry_credential_set":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ContainerRegistry/registries/registry1/credentialSets/credentialSet1
		return ""
	case "azurerm_container_registry_scope_map":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ContainerRegistry/registries/myregistry1/scopeMaps/scopemap1
		return ""
	case "azurerm_container_registry_task":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.ContainerRegistry/registries/registry1/tasks/task1
		return ""
	case "azurerm_container_registry_task_schedule_run_now":
		// No standard import format found in documentation
		return ""
	case "azurerm_container_registry_token":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ContainerRegistry/registries/myregistry1/tokens/token1
		return ""
	case "azurerm_container_registry_token_password":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.ContainerRegistry/registries/registry1/tokens/token1/passwords/password
		return ""
	case "azurerm_container_registry_webhook":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ContainerRegistry/registries/myregistry1/webHooks/mywebhook1
		return ""
	case "azurerm_cosmosdb_account":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/databaseAccounts/account1
		return ""
	case "azurerm_cosmosdb_cassandra_cluster":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DocumentDB/cassandraClusters/cluster1
		return ""
	case "azurerm_cosmosdb_cassandra_datacenter":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DocumentDB/cassandraClusters/cluster1/dataCenters/dc1
		return ""
	case "azurerm_cosmosdb_cassandra_keyspace":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/cassandraKeyspaces/ks1
		return ""
	case "azurerm_cosmosdb_cassandra_table":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/cassandraKeyspaces/ks1/tables/table1
		return ""
	case "azurerm_cosmosdb_gremlin_database":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/gremlinDatabases/db1
		return ""
	case "azurerm_cosmosdb_gremlin_graph":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/gremlinDatabases/db1/graphs/graphs1
		return ""
	case "azurerm_cosmosdb_mongo_collection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/mongodbDatabases/db1/collections/collection1
		return ""
	case "azurerm_cosmosdb_mongo_database":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/mongodbDatabases/db1
		return ""
	case "azurerm_cosmosdb_mongo_role_definition":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/databaseAccounts/account1/mongodbRoleDefinitions/dbname1.rolename1
		return ""
	case "azurerm_cosmosdb_mongo_user_definition":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/databaseAccounts/account1/mongodbUserDefinitions/dbname1.username1
		return ""
	case "azurerm_cosmosdb_postgresql_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/cluster1
		return ""
	case "azurerm_cosmosdb_postgresql_coordinator_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/cluster1/coordinatorConfigurations/array_nulls
		return ""
	case "azurerm_cosmosdb_postgresql_firewall_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/cluster1/firewallRules/firewallRule1
		return ""
	case "azurerm_cosmosdb_postgresql_node_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/cluster1/nodeConfigurations/array_nulls
		return ""
	case "azurerm_cosmosdb_postgresql_role":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/cluster1/roles/role1
		return ""
	case "azurerm_cosmosdb_sql_container":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlDatabases/database1/containers/container1
		return ""
	case "azurerm_cosmosdb_sql_database":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlDatabases/db1
		return ""
	case "azurerm_cosmosdb_sql_dedicated_gateway":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.DocumentDB/databaseAccounts/account1/services/SqlDedicatedGateway
		return ""
	case "azurerm_cosmosdb_sql_function":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlDatabases/database1/containers/container1/userDefinedFunctions/userDefinedFunction1
		return ""
	case "azurerm_cosmosdb_sql_role_assignment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlRoleAssignments/9e007587-dbcd-4190-84cb-fcab5a09ca39
		return ""
	case "azurerm_cosmosdb_sql_role_definition":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlRoleDefinitions/28b3c337-f436-482b-a167-c2618dc52033
		return ""
	case "azurerm_cosmosdb_sql_stored_procedure":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlDatabases/db1/containers/c1/storedProcedures/sp1
		return ""
	case "azurerm_cosmosdb_sql_trigger":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlDatabases/database1/containers/container1/triggers/trigger1
		return ""
	case "azurerm_cosmosdb_table":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/tables/table1
		return ""
	case "azurerm_cost_anomaly_alert":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CostManagement/scheduledActions/dailyanomalybyresourcegroup
		return ""
	case "azurerm_cost_management_scheduled_action":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CostManagement/scheduledActions/scheduledaction1
		return ""
	case "azurerm_custom_ip_prefix":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/customIPPrefixes/customIPPrefix1
		return ""
	case "azurerm_custom_provider":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.CustomProviders/resourceProviders/example
		return ""
	case "azurerm_dashboard_grafana":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Dashboard/grafana/workspace1
		return ""
	case "azurerm_dashboard_grafana_managed_private_endpoint":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Dashboard/grafana/workspace1/managedPrivateEndpoints/endpoint1
		return ""
	case "azurerm_data_factory":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example
		return ""
	case "azurerm_data_factory_credential_service_principal":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.DataFactory/factories/example/credentials/credential1
		return ""
	case "azurerm_data_factory_credential_user_managed_identity":
		// Azure Resource ID: /subscriptions/1f3d6e58-feed-4bb6-87e5-a52305ad3375/resourceGroups/example-resources/providers/Microsoft.DataFactory/factories/example/credentials/credential1
		return ""
	case "azurerm_data_factory_custom_dataset":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
		return ""
	case "azurerm_data_factory_customer_managed_key":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example
		return ""
	case "azurerm_data_factory_data_flow":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/dataflows/example
		return ""
	case "azurerm_data_factory_dataset_azure_blob":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
		return ""
	case "azurerm_data_factory_dataset_azure_sql_table":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
		return ""
	case "azurerm_data_factory_dataset_binary":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
		return ""
	case "azurerm_data_factory_dataset_cosmosdb_sqlapi":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
		return ""
	case "azurerm_data_factory_dataset_delimited_text":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
		return ""
	case "azurerm_data_factory_dataset_http":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
		return ""
	case "azurerm_data_factory_dataset_json":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
		return ""
	case "azurerm_data_factory_dataset_mysql":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
		return ""
	case "azurerm_data_factory_dataset_parquet":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
		return ""
	case "azurerm_data_factory_dataset_postgresql":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
		return ""
	case "azurerm_data_factory_dataset_snowflake":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
		return ""
	case "azurerm_data_factory_dataset_sql_server_table":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
		return ""
	case "azurerm_data_factory_flowlet_data_flow":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/dataflows/example
		return ""
	case "azurerm_data_factory_integration_runtime_azure":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/integrationRuntimes/example
		return ""
	case "azurerm_data_factory_integration_runtime_azure_ssis":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/integrationRuntimes/example
		return ""
	case "azurerm_data_factory_integration_runtime_self_hosted":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/integrationRuntimes/example
		return ""
	case "azurerm_data_factory_linked_custom_service":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_azure_blob_storage":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_azure_databricks":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_azure_file_storage":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_azure_function":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_azure_search":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_azure_sql_database":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_azure_table_storage":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_cosmosdb":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_cosmosdb_mongoapi":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_data_lake_storage_gen2":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_key_vault":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_kusto":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_mysql":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_odata":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_odbc":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_postgresql":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_sftp":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_snowflake":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_sql_managed_instance":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.DataFactory/factories/example/linkedServices/example
		return ""
	case "azurerm_data_factory_linked_service_sql_server":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_synapse":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_linked_service_web":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
		return ""
	case "azurerm_data_factory_managed_private_endpoint":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/managedVirtualNetworks/default/managedPrivateEndpoints/endpoint1
		return ""
	case "azurerm_data_factory_pipeline":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/pipelines/example
		return ""
	case "azurerm_data_factory_trigger_blob_event":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/triggers/example
		return ""
	case "azurerm_data_factory_trigger_custom_event":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/triggers/example
		return ""
	case "azurerm_data_factory_trigger_schedule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/triggers/example
		return ""
	case "azurerm_data_factory_trigger_tumbling_window":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/triggers/example
		return ""
	case "azurerm_data_protection_backup_instance_blob_storage":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupInstances/backupInstance1
		return ""
	case "azurerm_data_protection_backup_instance_data_lake_storage":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupInstances/backupInstance1
		return ""
	case "azurerm_data_protection_backup_instance_disk":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupInstances/backupInstance1
		return ""
	case "azurerm_data_protection_backup_instance_kubernetes_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupInstances/backupInstance1
		return ""
	case "azurerm_data_protection_backup_instance_mysql_flexible_server":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupInstances/backupInstance1
		return ""
	case "azurerm_data_protection_backup_instance_postgresql":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupInstances/backupInstance1
		return ""
	case "azurerm_data_protection_backup_instance_postgresql_flexible_server":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupInstances/backupInstance1
		return ""
	case "azurerm_data_protection_backup_policy_blob_storage":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupPolicies/backupPolicy1
		return ""
	case "azurerm_data_protection_backup_policy_data_lake_storage":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupPolicies/backupPolicy1
		return ""
	case "azurerm_data_protection_backup_policy_disk":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupPolicies/backupPolicy1
		return ""
	case "azurerm_data_protection_backup_policy_kubernetes_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupPolicies/backupPolicy1
		return ""
	case "azurerm_data_protection_backup_policy_mysql_flexible_server":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupPolicies/backupPolicy1
		return ""
	case "azurerm_data_protection_backup_policy_postgresql":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupPolicies/backupPolicy1
		return ""
	case "azurerm_data_protection_backup_policy_postgresql_flexible_server":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupPolicies/backupPolicy1
		return ""
	case "azurerm_data_protection_backup_vault":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1
		return ""
	case "azurerm_data_protection_backup_vault_customer_managed_key":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1
		return ""
	case "azurerm_data_protection_resource_guard":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/resourceGuards/resourceGuard1
		return ""
	case "azurerm_data_share":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataShare/accounts/account1/shares/share1
		return ""
	case "azurerm_data_share_account":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataShare/accounts/account1
		return ""
	case "azurerm_data_share_dataset_blob_storage":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataShare/accounts/account1/shares/share1/dataSets/dataSet1
		return ""
	case "azurerm_data_share_dataset_data_lake_gen2":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataShare/accounts/account1/shares/share1/dataSets/dataSet1
		return ""
	case "azurerm_data_share_dataset_kusto_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataShare/accounts/account1/shares/share1/dataSets/dataSet1
		return ""
	case "azurerm_data_share_dataset_kusto_database":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataShare/accounts/account1/shares/share1/dataSets/dataSet1
		return ""
	case "azurerm_database_migration_project":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.DataMigration/services/example-dms/projects/project1
		return ""
	case "azurerm_database_migration_service":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.DataMigration/services/database_migration_service1
		return ""
	case "azurerm_databox_edge_device":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/device1
		return ""
	case "azurerm_databricks_access_connector":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Databricks/accessConnectors/connector1
		return ""
	case "azurerm_databricks_virtual_network_peering":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Databricks/workspaces/workspace1/virtualNetworkPeerings/peering1
		return ""
	case "azurerm_databricks_workspace":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Databricks/workspaces/workspace1
		return ""
	case "azurerm_databricks_workspace_root_dbfs_customer_managed_key":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Databricks/workspaces/workspace1
		return ""
	case "azurerm_datadog_monitor":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Datadog/monitors/monitor1
		return ""
	case "azurerm_datadog_monitor_sso_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Datadog/monitors/monitor1/singleSignOnConfigurations/default
		return ""
	case "azurerm_datadog_monitor_tag_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Datadog/monitors/monitor1/tagRules/default
		return ""
	case "azurerm_dedicated_hardware_security_module":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/hsm1
		return ""
	case "azurerm_dedicated_host":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/hostGroups/group1/hosts/host1
		return ""
	case "azurerm_dedicated_host_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/hostGroups/group1
		return ""
	case "azurerm_dev_center":
		// Azure Resource ID: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devCenters/{devCenterName}
		return ""
	case "azurerm_dev_center_attached_network":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevCenter/devCenters/dc1/attachedNetworks/et1
		return ""
	case "azurerm_dev_center_catalog":
		// Azure Resource ID: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devCenters/{devCenterName}/catalogs/{catalogName}
		return ""
	case "azurerm_dev_center_dev_box_definition":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevCenter/devCenters/dc1/devBoxDefinitions/et1
		return ""
	case "azurerm_dev_center_environment_type":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevCenter/devCenters/dc1/environmentTypes/et1
		return ""
	case "azurerm_dev_center_gallery":
		// Azure Resource ID: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devCenters/{devCenterName}/galleries/{galleryName}
		return ""
	case "azurerm_dev_center_network_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevCenter/networkConnections/networkConnection1
		return ""
	case "azurerm_dev_center_project":
		// Azure Resource ID: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}
		return ""
	case "azurerm_dev_center_project_environment_type":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevCenter/projects/project1/environmentTypes/et1
		return ""
	case "azurerm_dev_center_project_pool":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevCenter/projects/project1/pools/pool1
		return ""
	case "azurerm_dev_test_global_vm_shutdown_schedule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-rg/providers/Microsoft.DevTestLab/schedules/shutdown-computevm-SampleVM
		return ""
	case "azurerm_dev_test_lab":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1
		return ""
	case "azurerm_dev_test_linux_virtual_machine":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1/virtualMachines/machine1
		return ""
	case "azurerm_dev_test_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1/policySets/default/policies/policy1
		return ""
	case "azurerm_dev_test_schedule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DevTestLab/labs/myDevTestLab/schedules/labvmautostart
		return ""
	case "azurerm_dev_test_virtual_network":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1/virtualNetworks/network1
		return ""
	case "azurerm_dev_test_windows_virtual_machine":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1/virtualMachines/machine1
		return ""
	case "azurerm_digital_twins_endpoint_eventgrid":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DigitalTwins/digitalTwinsInstances/dt1/endpoints/ep1
		return ""
	case "azurerm_digital_twins_endpoint_eventhub":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DigitalTwins/digitalTwinsInstances/dt1/endpoints/ep1
		return ""
	case "azurerm_digital_twins_endpoint_servicebus":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DigitalTwins/digitalTwinsInstances/dt1/endpoints/ep1
		return ""
	case "azurerm_digital_twins_instance":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DigitalTwins/digitalTwinsInstances/dt1
		return ""
	case "azurerm_digital_twins_time_series_database_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DigitalTwins/digitalTwinsInstances/dt1/timeSeriesDatabaseConnections/connection1
		return ""
	case "azurerm_disk_access":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/diskAccesses/diskAccess1
		return ""
	case "azurerm_disk_encryption_set":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/diskEncryptionSets/encryptionSet1
		return ""
	case "azurerm_dns_a_record":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnsZones/zone1/A/myrecord1
		return ""
	case "azurerm_dns_aaaa_record":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnsZones/zone1/AAAA/myrecord1
		return ""
	case "azurerm_dns_caa_record":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnsZones/zone1/CAA/myrecord1
		return ""
	case "azurerm_dns_cname_record":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnsZones/zone1/CNAME/myrecord1
		return ""
	case "azurerm_dns_mx_record":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnsZones/zone1/MX/myrecord1
		return ""
	case "azurerm_dns_ns_record":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnsZones/zone1/NS/myrecord1
		return ""
	case "azurerm_dns_ptr_record":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnsZones/zone1/PTR/myrecord1
		return ""
	case "azurerm_dns_srv_record":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnsZones/zone1/SRV/myrecord1
		return ""
	case "azurerm_dns_txt_record":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnsZones/zone1/TXT/myrecord1
		return ""
	case "azurerm_dns_zone":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnsZones/zone1
		return ""
	case "azurerm_dynatrace_monitor":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Dynatrace.Observability/monitors/monitor1
		return ""
	case "azurerm_dynatrace_tag_rules":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Dynatrace.Observability/monitors/monitor1/tagRules/tagRules1
		return ""
	case "azurerm_elastic_cloud_elasticsearch":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Elastic/monitors/monitor1
		return ""
	case "azurerm_elastic_san":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ElasticSan/elasticSans/esan1
		return ""
	case "azurerm_elastic_san_volume":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ElasticSan/elasticSans/esan1/volumeGroups/vg1/volumes/vol1
		return ""
	case "azurerm_elastic_san_volume_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ElasticSan/elasticSans/esan1/volumeGroups/vg1
		return ""
	case "azurerm_email_communication_service":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Communication/emailServices/emailCommunicationService1
		return ""
	case "azurerm_email_communication_service_domain":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Communication/emailServices/emailCommunicationService1/domains/domain1
		return ""
	case "azurerm_email_communication_service_domain_sender_username":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Communication/emailServices/service1/domains/domain1/senderUsernames/username1
		return ""
	case "azurerm_eventgrid_domain":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/domains/domain1
		return ""
	case "azurerm_eventgrid_domain_topic":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/domains/domain1/topics/topic1
		return ""
	case "azurerm_eventgrid_event_subscription":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/topics/topic1/providers/Microsoft.EventGrid/eventSubscriptions/eventSubscription1
		return ""
	case "azurerm_eventgrid_namespace":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/namespaces/namespace1
		return ""
	case "azurerm_eventgrid_namespace_topic":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/namespaces/eventgrid1/topics/topic1
		return ""
	case "azurerm_eventgrid_partner_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1
		return ""
	case "azurerm_eventgrid_partner_namespace":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.EventGrid/partnerNamespaces/example
		return ""
	case "azurerm_eventgrid_partner_registration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.EventGrid/partnerRegistrations/example
		return ""
	case "azurerm_eventgrid_system_topic":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/systemTopics/systemTopic1
		return ""
	case "azurerm_eventgrid_system_topic_event_subscription":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/systemTopics/topic1/eventSubscriptions/subscription1
		return ""
	case "azurerm_eventgrid_topic":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/topics/topic1
		return ""
	case "azurerm_eventhub":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/eventhubs/eventhub1
		return ""
	case "azurerm_eventhub_authorization_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/eventhubs/eventhub1/authorizationRules/rule1
		return ""
	case "azurerm_eventhub_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/clusters/cluster1
		return ""
	case "azurerm_eventhub_consumer_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/eventhubs/eventhub1/consumerGroups/consumerGroup1
		return ""
	case "azurerm_eventhub_namespace":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1
		return ""
	case "azurerm_eventhub_namespace_authorization_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/authorizationRules/rule1
		return ""
	case "azurerm_eventhub_namespace_customer_managed_key":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1
		return ""
	case "azurerm_eventhub_namespace_disaster_recovery_config":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/disasterRecoveryConfigs/config1
		return ""
	case "azurerm_eventhub_namespace_schema_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/schemaGroups/group1
		return ""
	case "azurerm_express_route_circuit":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/expressRouteCircuits/myExpressRoute
		return ""
	case "azurerm_express_route_circuit_authorization":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/expressRouteCircuits/myExpressRoute/authorizations/auth1
		return ""
	case "azurerm_express_route_circuit_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/expressRouteCircuits/circuit1/peerings/peering1/connections/connection1
		return ""
	case "azurerm_express_route_circuit_peering":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/expressRouteCircuits/myExpressRoute/peerings/peering1
		return ""
	case "azurerm_express_route_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/expressRouteGateways/expressRouteGateway1/expressRouteConnections/connection1
		return ""
	case "azurerm_express_route_gateway":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/expressRouteGateways/myExpressRouteGateway
		return ""
	case "azurerm_express_route_port":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/expressRoutePorts/port1
		return ""
	case "azurerm_express_route_port_authorization":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/expressRoutePorts/myExpressPort/authorizations/auth1
		return ""
	case "azurerm_extended_location_custom_location":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.ExtendedLocation/customLocations/example-custom-location
		return ""
	case "azurerm_fabric_capacity":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Fabric/capacities/capacity1
		return ""
	case "azurerm_federated_identity_credential":
		// Azure Resource ID: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{parentIdentityName}/federatedIdentityCredentials/{resourceName}
		return ""
	case "azurerm_firewall":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/azureFirewalls/testfirewall
		return ""
	case "azurerm_firewall_application_rule_collection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/azureFirewalls/myfirewall/applicationRuleCollections/mycollection
		return ""
	case "azurerm_firewall_nat_rule_collection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/azureFirewalls/myfirewall/natRuleCollections/mycollection
		return ""
	case "azurerm_firewall_network_rule_collection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/azureFirewalls/myfirewall/networkRuleCollections/mycollection
		return ""
	case "azurerm_firewall_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/firewallPolicies/policy1
		return ""
	case "azurerm_firewall_policy_rule_collection_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/firewallPolicies/policy1/ruleCollectionGroups/gruop1
		return ""
	case "azurerm_fluid_relay_server":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.FluidRelay/fluidRelayServers/server1
		return ""
	case "azurerm_frontdoor":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/frontDoors/frontdoor1
		return ""
	case "azurerm_frontdoor_custom_https_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/frontDoors/frontdoor1/customHttpsConfiguration/endpoint1
		return ""
	case "azurerm_frontdoor_firewall_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/examplefdwafpolicy
		return ""
	case "azurerm_frontdoor_rules_engine":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Network/frontdoors/frontdoor1/rulesEngines/rule1
		return ""
	case "azurerm_function_app":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/functionapp1
		return ""
	case "azurerm_function_app_active_slot":
		// Azure Resource ID: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1"
		return ""
	case "azurerm_function_app_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Web/sites/webapp/providers/Microsoft.ServiceLinker/linkers/serviceconnector1
		return ""
	case "azurerm_function_app_flex_consumption":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1
		return ""
	case "azurerm_function_app_function":
		// Azure Resource ID: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/functions/function1"
		return ""
	case "azurerm_function_app_hybrid_connection":
		// Azure Resource ID: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/hybridConnectionNamespaces/hybridConnectionNamespace1/relays/relay1"
		return ""
	case "azurerm_function_app_slot":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/functionapp1/slots/staging
		return ""
	case "azurerm_gallery_application":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/galleries/gallery1/applications/galleryApplication1
		return ""
	case "azurerm_gallery_application_version":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/galleries/gallery1/applications/galleryApplication1/versions/galleryApplicationVersion1
		return ""
	case "azurerm_graph_services_account":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.GraphServices/accounts/account1
		return ""
	case "azurerm_hdinsight_hadoop_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1
		return ""
	case "azurerm_hdinsight_hbase_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1
		return ""
	case "azurerm_hdinsight_interactive_query_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1
		return ""
	case "azurerm_hdinsight_kafka_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1
		return ""
	case "azurerm_hdinsight_spark_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1
		return ""
	case "azurerm_healthbot":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.HealthBot/healthBots/bot1
		return ""
	case "azurerm_healthcare_dicom_service":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.HealthcareApis/workspaces/workspace1/dicomServices/service1
		return ""
	case "azurerm_healthcare_fhir_service":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.HealthcareApis/workspaces/workspace1/fhirServices/service1
		return ""
	case "azurerm_healthcare_medtech_service":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.HealthcareApis/workspaces/workspace1/iotConnectors/iotconnector1
		return ""
	case "azurerm_healthcare_medtech_service_fhir_destination":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.HealthcareApis/workspaces/workspace1/iotConnectors/iotconnector1/fhirDestinations/destination1
		return ""
	case "azurerm_healthcare_service":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource_group/providers/Microsoft.HealthcareApis/services/service_name
		return ""
	case "azurerm_healthcare_workspace":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.HealthcareApis/workspaces/workspace1
		return ""
	case "azurerm_hpc_cache":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.StorageCache/caches/cacheName
		return ""
	case "azurerm_hpc_cache_access_policy":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.StorageCache/caches/cache1/cacheAccessPolicies/policy1
		return ""
	case "azurerm_hpc_cache_blob_nfs_target":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StorageCache/caches/cache1/storageTargets/target1
		return ""
	case "azurerm_hpc_cache_blob_target":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StorageCache/caches/cache1/storageTargets/target1
		return ""
	case "azurerm_hpc_cache_nfs_target":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StorageCache/caches/cache1/storageTargets/target1
		return ""
	case "azurerm_image":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/images/image1
		return ""
	case "azurerm_iot_security_device_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Devices/iotHubs/hub1/providers/Microsoft.Security/deviceSecurityGroups/group1
		return ""
	case "azurerm_iot_security_solution":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Security/iotSecuritySolutions/solution1
		return ""
	case "azurerm_iotcentral_application":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.IoTCentral/iotApps/app1
		return ""
	case "azurerm_iotcentral_application_network_rule_set":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.IoTCentral/iotApps/app1
		return ""
	case "azurerm_iotcentral_organization":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.IoTCentral/iotApps/example/organizations/example
		return ""
	case "azurerm_iothub":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/iotHubs/hub1
		return ""
	case "azurerm_iothub_certificate":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/iotHubs/example/certificates/example
		return ""
	case "azurerm_iothub_consumer_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/iotHubs/hub1/eventHubEndpoints/events/consumerGroups/group1
		return ""
	case "azurerm_iothub_device_update_account":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.DeviceUpdate/accounts/account1
		return ""
	case "azurerm_iothub_device_update_instance":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.DeviceUpdate/accounts/account1/instances/instance1
		return ""
	case "azurerm_iothub_dps":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/provisioningServices/example
		return ""
	case "azurerm_iothub_dps_certificate":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/provisioningServices/example/certificates/example
		return ""
	case "azurerm_iothub_dps_shared_access_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/provisioningServices/dps1/keys/shared_access_policy1
		return ""
	case "azurerm_iothub_endpoint_cosmosdb_account":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/iotHubs/hub1/endpoints/cosmosDBAccountEndpoint1
		return ""
	case "azurerm_iothub_endpoint_eventhub":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/iotHubs/hub1/endpoints/eventhub_endpoint1
		return ""
	case "azurerm_iothub_endpoint_servicebus_queue":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/iotHubs/hub1/endpoints/servicebusqueue_endpoint1
		return ""
	case "azurerm_iothub_endpoint_servicebus_topic":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/iotHubs/hub1/endpoints/servicebustopic_endpoint1
		return ""
	case "azurerm_iothub_endpoint_storage_container":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/iotHubs/hub1/endpoints/storage_container_endpoint1
		return ""
	case "azurerm_iothub_enrichment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/iotHubs/hub1/enrichments/enrichment1
		return ""
	case "azurerm_iothub_fallback_route":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/iotHubs/hub1/fallbackRoute/default
		return ""
	case "azurerm_iothub_file_upload":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/iotHubs/hub1
		return ""
	case "azurerm_iothub_route":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/iotHubs/hub1/routes/route1
		return ""
	case "azurerm_iothub_shared_access_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/iotHubs/hub1/iotHubKeys/shared_access_policy1
		return ""
	case "azurerm_ip_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/ipGroups/myIpGroup
		return ""
	case "azurerm_ip_group_cidr":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/ipGroups/test-ipgroup/cidrs/10.1.0.0_24
		return ""
	case "azurerm_key_vault":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.KeyVault/vaults/vault1
		return ""
	case "azurerm_key_vault_access_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.KeyVault/vaults/test-vault/objectId/11111111-1111-1111-1111-111111111111/applicationId/22222222-2222-2222-2222-222222222222
		return ""
	case "azurerm_key_vault_certificate":
		// Azure Resource ID: "https://example-keyvault.vault.azure.net/certificates/example/fdf067c93bbb4b22bff4d8b7a9a56217"
		return ""
	case "azurerm_key_vault_certificate_contacts":
		// Azure Resource ID: https://example-keyvault.vault.azure.net/certificates/contacts
		return ""
	case "azurerm_key_vault_certificate_issuer":
		// Azure Resource ID: "https://key-vault-name.vault.azure.net/certificates/issuers/example"
		return ""
	case "azurerm_key_vault_key":
		// Azure Resource ID: "https://example-keyvault.vault.azure.net/keys/example/fdf067c93bbb4b22bff4d8b7a9a56217"
		return ""
	case "azurerm_key_vault_managed_hardware_security_module":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.KeyVault/managedHSMs/hsm1
		return ""
	case "azurerm_key_vault_managed_hardware_security_module_key":
		// Azure Resource ID: https://exampleHSM.managedhsm.azure.net/keys/exampleKey
		return ""
	case "azurerm_key_vault_managed_hardware_security_module_key_rotation_policy":
		// Azure Resource ID: https://example-hsm.managedhsm.azure.net/keys/example
		return ""
	case "azurerm_key_vault_managed_hardware_security_module_role_assignment":
		// Azure Resource ID: https://0000.managedhsm.azure.net///RoleAssignment/00000000-0000-0000-0000-000000000000
		return ""
	case "azurerm_key_vault_managed_hardware_security_module_role_definition":
		// Azure Resource ID: https://0000.managedhsm.azure.net///RoleDefinition/00000000-0000-0000-0000-000000000000
		return ""
	case "azurerm_key_vault_managed_storage_account":
		// Azure Resource ID: https://example-keyvault.vault.azure.net/storage/exampleStorageAcc01
		return ""
	case "azurerm_key_vault_managed_storage_account_sas_token_definition":
		// Azure Resource ID: https://example-keyvault.vault.azure.net/storage/exampleStorageAcc01/sas/exampleSasDefinition01
		return ""
	case "azurerm_key_vault_secret":
		// Azure Resource ID: "https://example-keyvault.vault.azure.net/secrets/example/fdf067c93bbb4b22bff4d8b7a9a56217"
		return ""
	case "azurerm_kubernetes_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ContainerService/managedClusters/cluster1
		return ""
	case "azurerm_kubernetes_cluster_deployment_safeguard":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ContainerService/managedClusters/cluster1
		return ""
	case "azurerm_kubernetes_cluster_extension":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.ContainerService/managedClusters/cluster1/providers/Microsoft.KubernetesConfiguration/extensions/extension1
		return ""
	case "azurerm_kubernetes_cluster_node_pool":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ContainerService/managedClusters/cluster1/agentPools/pool1
		return ""
	case "azurerm_kubernetes_cluster_trusted_access_role_binding":
		// Azure Resource ID: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{managedClusterName}/trustedAccessRoleBindings/{trustedAccessRoleBindingName}
		return ""
	case "azurerm_kubernetes_fleet_manager":
		// Azure Resource ID: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}
		return ""
	case "azurerm_kubernetes_fleet_member":
		// Azure Resource ID: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/members/{memberName}
		return ""
	case "azurerm_kubernetes_fleet_update_run":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.ContainerService/fleets/fleet1/updateRuns/updateRun1
		return ""
	case "azurerm_kubernetes_fleet_update_strategy":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.ContainerService/fleets/fleet1/updateStrategies/updateStrategy1
		return ""
	case "azurerm_kubernetes_flux_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.ContainerService/managedClusters/cluster1/providers/Microsoft.KubernetesConfiguration/fluxConfigurations/fluxConfiguration1
		return ""
	case "azurerm_kusto_attached_database_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/clusters/cluster1/attachedDatabaseConfigurations/configuration1
		return ""
	case "azurerm_kusto_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/clusters/cluster1
		return ""
	case "azurerm_kusto_cluster_customer_managed_key":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/clusters/cluster1
		return ""
	case "azurerm_kusto_cluster_managed_private_endpoint":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/clusters/cluster1/managedPrivateEndpoints/managedPrivateEndpoint1
		return ""
	case "azurerm_kusto_cluster_principal_assignment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/clusters/cluster1/principalAssignments/assignment1
		return ""
	case "azurerm_kusto_cosmosdb_data_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/clusters/cluster1/databases/database1/dataConnections/dataConnection1
		return ""
	case "azurerm_kusto_database":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/clusters/cluster1/databases/database1
		return ""
	case "azurerm_kusto_database_principal_assignment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/clusters/cluster1/databases/database1/principalAssignments/assignment1
		return ""
	case "azurerm_kusto_eventgrid_data_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/clusters/cluster1/databases/database1/dataConnections/dataConnection1
		return ""
	case "azurerm_kusto_eventhub_data_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/clusters/cluster1/databases/database1/dataConnections/eventHubConnection1
		return ""
	case "azurerm_kusto_iothub_data_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/clusters/cluster1/databases/database1/dataConnections/dataConnection1
		return ""
	case "azurerm_kusto_script":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Kusto/clusters/cluster1/databases/database1/scripts/script1
		return ""
	case "azurerm_lb":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1
		return ""
	case "azurerm_lb_backend_address_pool":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/backendAddressPools/pool1
		return ""
	case "azurerm_lb_backend_address_pool_address":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/loadBalancer1/backendAddressPools/backendAddressPool1/addresses/address1
		return ""
	case "azurerm_lb_nat_pool":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatPools/pool1
		return ""
	case "azurerm_lb_nat_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/rule1
		return ""
	case "azurerm_lb_outbound_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/outboundRules/rule1
		return ""
	case "azurerm_lb_probe":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/probes/probe1
		return ""
	case "azurerm_lb_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/loadBalancingRules/rule1
		return ""
	case "azurerm_lighthouse_assignment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ManagedServices/registrationAssignments/00000000-0000-0000-0000-000000000000
		return ""
	case "azurerm_lighthouse_definition":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ManagedServices/registrationDefinitions/00000000-0000-0000-0000-000000000000
		return ""
	case "azurerm_linux_function_app":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1
		return ""
	case "azurerm_linux_function_app_slot":
		// Azure Resource ID: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/slots/slot1"
		return ""
	case "azurerm_linux_virtual_machine":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachines/machine1
		return ""
	case "azurerm_linux_virtual_machine_scale_set":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachineScaleSets/scaleset1
		return ""
	case "azurerm_linux_web_app":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1
		return ""
	case "azurerm_linux_web_app_slot":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/slots/slot1
		return ""
	case "azurerm_load_test":
		// Azure Resource ID: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LoadTestService/loadTests/{loadTestName}
		return ""
	case "azurerm_local_network_gateway":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/localNetworkGateways/lng1
		return ""
	case "azurerm_log_analytics_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/clusters/cluster1
		return ""
	case "azurerm_log_analytics_cluster_customer_managed_key":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/clusters/cluster1
		return ""
	case "azurerm_log_analytics_data_export_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/dataExports/dataExport1
		return ""
	case "azurerm_log_analytics_datasource_windows_event":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/dataSources/datasource1
		return ""
	case "azurerm_log_analytics_datasource_windows_performance_counter":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/dataSources/datasource1
		return ""
	case "azurerm_log_analytics_linked_service":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/linkedServices/Automation
		return ""
	case "azurerm_log_analytics_linked_storage_account":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/linkedStorageAccounts/{dataSourceType}
		return ""
	case "azurerm_log_analytics_query_pack":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.OperationalInsights/queryPacks/queryPack1
		return ""
	case "azurerm_log_analytics_query_pack_query":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.OperationalInsights/queryPacks/queryPack1/queries/15b49e87-8555-4d92-8a7b-2014b469a9df
		return ""
	case "azurerm_log_analytics_saved_search":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/savedSearches/search1
		return ""
	case "azurerm_log_analytics_solution":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.OperationsManagement/solutions/solution1
		return ""
	case "azurerm_log_analytics_storage_insights":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/storageInsightConfigs/storageInsight1
		return ""
	case "azurerm_log_analytics_workspace":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1
		return ""
	case "azurerm_log_analytics_workspace_table":
		// No standard import format found in documentation
		return ""
	case "azurerm_log_analytics_workspace_table_custom_log":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/tables/table1
		return ""
	case "azurerm_logic_app_action_custom":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1/actions/custom1
		return ""
	case "azurerm_logic_app_action_http":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1/actions/webhook1
		return ""
	case "azurerm_logic_app_integration_account":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1
		return ""
	case "azurerm_logic_app_integration_account_agreement":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/agreements/agreement1
		return ""
	case "azurerm_logic_app_integration_account_assembly":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/assemblies/assembly1
		return ""
	case "azurerm_logic_app_integration_account_batch_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/batchConfigurations/batchConfiguration1
		return ""
	case "azurerm_logic_app_integration_account_certificate":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/certificates/certificate1
		return ""
	case "azurerm_logic_app_integration_account_map":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/maps/map1
		return ""
	case "azurerm_logic_app_integration_account_partner":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/partners/partner1
		return ""
	case "azurerm_logic_app_integration_account_schema":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/schemas/schema1
		return ""
	case "azurerm_logic_app_integration_account_session":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/sessions/session1
		return ""
	case "azurerm_logic_app_standard":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/logicapp1
		return ""
	case "azurerm_logic_app_trigger_custom":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1/triggers/custom1
		return ""
	case "azurerm_logic_app_trigger_http_request":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1/triggers/request1
		return ""
	case "azurerm_logic_app_trigger_recurrence":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1/triggers/daily
		return ""
	case "azurerm_logic_app_workflow":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1
		return ""
	case "azurerm_machine_learning_compute_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/cluster1
		return ""
	case "azurerm_machine_learning_compute_instance":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/compute1
		return ""
	case "azurerm_machine_learning_datastore_blobstorage":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.MachineLearningServices/workspaces/mlw1/dataStores/datastore1
		return ""
	case "azurerm_machine_learning_datastore_datalake_gen2":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.MachineLearningServices/workspaces/mlw1/dataStores/datastore1
		return ""
	case "azurerm_machine_learning_datastore_fileshare":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.MachineLearningServices/workspaces/mlw1/dataStores/datastore1
		return ""
	case "azurerm_machine_learning_inference_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/cluster1
		return ""
	case "azurerm_machine_learning_synapse_spark":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/compute1
		return ""
	case "azurerm_machine_learning_workspace":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.MachineLearningServices/workspaces/workspace1
		return ""
	case "azurerm_machine_learning_workspace_network_outbound_rule_fqdn":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/outboundRules/rule1
		return ""
	case "azurerm_machine_learning_workspace_network_outbound_rule_private_endpoint":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/outboundRules/rule1
		return ""
	case "azurerm_machine_learning_workspace_network_outbound_rule_service_tag":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/outboundRules/rule1
		return ""
	case "azurerm_maintenance_assignment_dedicated_host":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Compute/hostGroups/group1/hosts/host1/providers/Microsoft.Maintenance/configurationAssignments/assign1
		return ""
	case "azurerm_maintenance_assignment_dynamic_scope":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Maintenance/configurationAssignments/assignmentName
		return ""
	case "azurerm_maintenance_assignment_virtual_machine":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.Maintenance/configurationAssignments/assign1
		return ""
	case "azurerm_maintenance_assignment_virtual_machine_scale_set":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/providers/Microsoft.Maintenance/configurationAssignments/assign1
		return ""
	case "azurerm_maintenance_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Maintenance/maintenanceConfigurations/example-mc
		return ""
	case "azurerm_managed_application":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Solutions/applications/app1
		return ""
	case "azurerm_managed_application_definition":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Solutions/applicationDefinitions/appDefinition1
		return ""
	case "azurerm_managed_devops_pool":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevOpsInfrastructure/pools/pool1
		return ""
	case "azurerm_managed_disk":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/disks/manageddisk1
		return ""
	case "azurerm_managed_disk_sas_token":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/disks/manageddisk1
		return ""
	case "azurerm_managed_lustre_file_system":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.StorageCache/amlFilesystems/amlFilesystem1
		return ""
	case "azurerm_managed_redis":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/redisEnterprise/cluster1
		return ""
	case "azurerm_managed_redis_access_policy_assignment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/redisEnterprise/redis1/databases/default/accessPolicyAssignments/00000000-0000-0000-0000-000000000000
		return ""
	case "azurerm_managed_redis_geo_replication":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/redisEnterprise/cluster1
		return ""
	case "azurerm_management_group":
		// Azure Resource ID: /providers/Microsoft.Management/managementGroups/group1
		return ""
	case "azurerm_management_group_policy_assignment":
		// Azure Resource ID: /providers/Microsoft.Management/managementGroups/group1/providers/Microsoft.Authorization/policyAssignments/assignment1
		return ""
	case "azurerm_management_group_policy_exemption":
		// Azure Resource ID: /providers/Microsoft.Management/managementGroups/group1/providers/Microsoft.Authorization/policyExemptions/exemption1
		return ""
	case "azurerm_management_group_policy_remediation":
		// Azure Resource ID: /providers/Microsoft.Management/managementGroups/my-mgmt-group-id/providers/Microsoft.PolicyInsights/remediations/remediation1
		return ""
	case "azurerm_management_group_policy_set_definition":
		// Azure Resource ID: /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName
		return ""
	case "azurerm_management_group_subscription_association":
		// Azure Resource ID: /providers/Microsoft.Management/managementGroups/MyManagementGroup/subscriptions/12345678-1234-1234-1234-123456789012
		return ""
	case "azurerm_management_group_template_deployment":
		// Azure Resource ID: /providers/Microsoft.Management/managementGroups/my-management-group-id/providers/Microsoft.Resources/deployments/deploy1
		return ""
	case "azurerm_management_lock":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Authorization/locks/lock1
		return ""
	case "azurerm_maps_account":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Maps/accounts/my-maps-account
		return ""
	case "azurerm_maps_creator":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Maps/accounts/account1/creators/creator1
		return ""
	case "azurerm_marketplace_agreement":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.MarketplaceOrdering/agreements/publisher1/offers/offer1/plans/plan1
		return ""
	case "azurerm_marketplace_role_assignment":
		// Azure Resource ID: /providers/Microsoft.Marketplace/providers/Microsoft.Authorization/roleAssignments/00000000-0000-0000-0000-000000000000
		return ""
	case "azurerm_mongo_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/mongoClusters/myMongoCluster
		return ""
	case "azurerm_mongo_cluster_firewall_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/mongoClusters/cluster1/firewallRules/rule1
		return ""
	case "azurerm_mongo_cluster_user":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/mongoClusters/cluster1/users/user1
		return ""
	case "azurerm_monitor_aad_diagnostic_setting":
		// Azure Resource ID: /providers/Microsoft.AADIAM/diagnosticSettings/setting1
		return ""
	case "azurerm_monitor_action_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Insights/actionGroups/myagname
		return ""
	case "azurerm_monitor_activity_log_alert":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Insights/activityLogAlerts/myalertname
		return ""
	case "azurerm_monitor_alert_processing_rule_action_group":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.AlertsManagement/actionRules/actionRule1
		return ""
	case "azurerm_monitor_alert_processing_rule_suppression":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.AlertsManagement/actionRules/actionRule1
		return ""
	case "azurerm_monitor_alert_prometheus_rule_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AlertsManagement/prometheusRuleGroups/ruleGroup1
		return ""
	case "azurerm_monitor_autoscale_setting":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Insights/autoScaleSettings/setting1
		return ""
	case "azurerm_monitor_data_collection_endpoint":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Insights/dataCollectionEndpoints/endpoint1
		return ""
	case "azurerm_monitor_data_collection_rule":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Insights/dataCollectionRules/rule1
		return ""
	case "azurerm_monitor_data_collection_rule_association":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.Insights/dataCollectionRuleAssociations/dca1
		return ""
	case "azurerm_monitor_diagnostic_setting":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.KeyVault/vaults/vault1|logMonitoring1"
		return ""
	case "azurerm_monitor_metric_alert":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.Insights/metricAlerts/example-metricalert
		return ""
	case "azurerm_monitor_private_link_scope":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Insights/privateLinkScopes/pls1
		return ""
	case "azurerm_monitor_private_link_scoped_service":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Insights/privateLinkScopes/pls1/scopedResources/sr1
		return ""
	case "azurerm_monitor_scheduled_query_rules_alert":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Insights/scheduledQueryRules/myrulename
		return ""
	case "azurerm_monitor_scheduled_query_rules_alert_v2":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Insights/scheduledQueryRules/rule1
		return ""
	case "azurerm_monitor_scheduled_query_rules_log":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Insights/scheduledQueryRules/myrulename
		return ""
	case "azurerm_monitor_smart_detector_alert_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AlertsManagement/smartDetectorAlertRules/rule1
		return ""
	case "azurerm_monitor_workspace":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Monitor/accounts/azureMonitorWorkspace1
		return ""
	case "azurerm_mssql_database":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/databases/example1
		return ""
	case "azurerm_mssql_database_extended_auditing_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlServer1/databases/db1/extendedAuditingSettings/default
		return ""
	case "azurerm_mssql_database_vulnerability_assessment_rule_baseline":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/acceptanceTestResourceGroup1/providers/Microsoft.Sql/servers/mssqlserver/databases/mysqldatabase/vulnerabilityAssessments/Default/rules/VA2065/baselines/master
		return ""
	case "azurerm_mssql_elasticpool":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/elasticPools/myelasticpoolname
		return ""
	case "azurerm_mssql_failover_group":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Sql/servers/server1/failoverGroups/failoverGroup1
		return ""
	case "azurerm_mssql_firewall_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/firewallRules/rule1
		return ""
	case "azurerm_mssql_job":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Sql/servers/myserver1/jobAgents/myjobagent1/jobs/myjob1
		return ""
	case "azurerm_mssql_job_agent":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Sql/servers/myserver1/jobAgents/myjobagent1
		return ""
	case "azurerm_mssql_job_credential":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Sql/servers/myserver1/jobAgents/myjobagent1/credentials/credential1
		return ""
	case "azurerm_mssql_job_schedule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Sql/servers/myserver1/jobAgents/myjobagent1/jobs/myjob1
		return ""
	case "azurerm_mssql_job_step":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Sql/servers/myserver1/jobAgents/myjobagent1/jobs/myjob1/steps/myjobstep1
		return ""
	case "azurerm_mssql_job_target_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Sql/servers/myserver1/jobAgents/myjobagent1/targetGroups/mytargetgroup1
		return ""
	case "azurerm_mssql_managed_database":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/managedInstances/myserver/databases/mydatabase
		return ""
	case "azurerm_mssql_managed_instance":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/managedInstances/myserver
		return ""
	case "azurerm_mssql_managed_instance_active_directory_administrator":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/managedInstances/mymanagedinstance/administrators/activeDirectory
		return ""
	case "azurerm_mssql_managed_instance_failover_group":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Sql/locations/Location/instanceFailoverGroups/failoverGroup1
		return ""
	case "azurerm_mssql_managed_instance_security_alert_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/acceptanceTestResourceGroup1/providers/Microsoft.Sql/managedInstances/instance1/securityAlertPolicies/Default
		return ""
	case "azurerm_mssql_managed_instance_start_stop_schedule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Sql/managedInstances/managedInstance1/startStopSchedules/default
		return ""
	case "azurerm_mssql_managed_instance_transparent_data_encryption":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Sql/managedInstances/instance1/encryptionProtector/current
		return ""
	case "azurerm_mssql_managed_instance_vulnerability_assessment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/acceptanceTestResourceGroup1/providers/Microsoft.Sql/managedInstances/instance1/vulnerabilityAssessments/Default
		return ""
	case "azurerm_mssql_outbound_firewall_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/outboundFirewallRules/fqdn1
		return ""
	case "azurerm_mssql_server":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver
		return ""
	case "azurerm_mssql_server_dns_alias":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/dnsAliases/default
		return ""
	case "azurerm_mssql_server_extended_auditing_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlServer1/extendedAuditingSettings/default
		return ""
	case "azurerm_mssql_server_microsoft_support_auditing_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlServer1/devOpsAuditingSettings/default
		return ""
	case "azurerm_mssql_server_security_alert_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/acceptanceTestResourceGroup1/providers/Microsoft.Sql/servers/mssqlserver/securityAlertPolicies/Default
		return ""
	case "azurerm_mssql_server_transparent_data_encryption":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/encryptionProtector/current
		return ""
	case "azurerm_mssql_server_vulnerability_assessment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/acceptanceTestResourceGroup1/providers/Microsoft.Sql/servers/mssqlserver/vulnerabilityAssessments/Default
		return ""
	case "azurerm_mssql_virtual_machine":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/example1
		return ""
	case "azurerm_mssql_virtual_machine_availability_group_listener":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/vmgroup1/availabilityGroupListeners/listener1
		return ""
	case "azurerm_mssql_virtual_machine_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/vmgroup1
		return ""
	case "azurerm_mssql_virtual_network_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/virtualNetworkRules/vnetrulename
		return ""
	case "azurerm_mysql_flexible_database":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforMySQL/flexibleServers/flexibleserver1/databases/database1
		return ""
	case "azurerm_mysql_flexible_server":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforMySQL/flexibleServers/flexibleServer1
		return ""
	case "azurerm_mysql_flexible_server_active_directory_administrator":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.DBforMySQL/flexibleServers/server1/administrators/ActiveDirectory
		return ""
	case "azurerm_mysql_flexible_server_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforMySQL/flexibleServers/flexibleServer1/configurations/interactive_timeout
		return ""
	case "azurerm_mysql_flexible_server_firewall_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforMySQL/flexibleServers/flexibleServer1/firewallRules/firewallRule1
		return ""
	case "azurerm_nat_gateway":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/natGateways/gateway1
		return ""
	case "azurerm_nat_gateway_public_ip_association":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/natGateways/natGateway1|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/publicIPAddresses/publicIPAddress1"
		return ""
	case "azurerm_nat_gateway_public_ip_prefix_association":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/natGateways/natGateway1|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/publicIPPrefixes/publicIPPrefix1"
		return ""
	case "azurerm_netapp_account":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1
		return ""
	case "azurerm_netapp_account_encryption":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1
		return ""
	case "azurerm_netapp_backup_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1/backupPolicies/backuppolicy1
		return ""
	case "azurerm_netapp_backup_vault":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1/backupVaults/backupVault1
		return ""
	case "azurerm_netapp_pool":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1
		return ""
	case "azurerm_netapp_snapshot":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1/snapshots/snapshot1
		return ""
	case "azurerm_netapp_snapshot_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1/snapshotPolicies/snapshotpolicy1
		return ""
	case "azurerm_netapp_volume":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1
		return ""
	case "azurerm_netapp_volume_group_oracle":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mytest-rg/providers/Microsoft.NetApp/netAppAccounts/netapp-account-test/volumeGroups/netapp-volumegroup-test
		return ""
	case "azurerm_netapp_volume_group_sap_hana":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mytest-rg/providers/Microsoft.NetApp/netAppAccounts/netapp-account-test/volumeGroups/netapp-volumegroup-test
		return ""
	case "azurerm_netapp_volume_quota_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/vol1/volumeQuotaRules/quota1
		return ""
	case "azurerm_network_connection_monitor":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/networkWatchers/watcher1/connectionMonitors/connectionMonitor1
		return ""
	case "azurerm_network_ddos_protection_plan":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/ddosProtectionPlans/testddospplan
		return ""
	case "azurerm_network_function_azure_traffic_collector":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.NetworkFunction/azureTrafficCollectors/azureTrafficCollector1
		return ""
	case "azurerm_network_function_collector_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.NetworkFunction/azureTrafficCollectors/azureTrafficCollector1/collectorPolicies/collectorPolicy1
		return ""
	case "azurerm_network_interface":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkInterfaces/nic1
		return ""
	case "azurerm_network_interface_application_gateway_backend_address_pool_association":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkInterfaces/nic1/ipConfigurations/example|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/applicationGateways/gateway1/backendAddressPools/pool1
		return ""
	case "azurerm_network_interface_application_security_group_association":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkInterfaces/nic1|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/applicationSecurityGroups/securityGroup1"
		return ""
	case "azurerm_network_interface_backend_address_pool_association":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkInterfaces/nic1/ipConfigurations/example|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/backendAddressPools/pool1"
		return ""
	case "azurerm_network_interface_nat_rule_association":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkInterfaces/nic1/ipConfigurations/example|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/rule1
		return ""
	case "azurerm_network_interface_security_group_association":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkInterfaces/example|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/networkSecurityGroups/group1"
		return ""
	case "azurerm_network_manager":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/networkManagers/networkManager1
		return ""
	case "azurerm_network_manager_admin_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/networkManagers/networkManager1/securityAdminConfigurations/configuration1/ruleCollections/ruleCollection1/rules/rule1
		return ""
	case "azurerm_network_manager_admin_rule_collection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/networkManagers/networkManager1/securityAdminConfigurations/configuration1/ruleCollections/ruleCollection1
		return ""
	case "azurerm_network_manager_connectivity_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/networkManagers/networkManager1/connectivityConfigurations/configuration1
		return ""
	case "azurerm_network_manager_deployment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/networkManagers/networkManager1/commit|eastus|Connectivity
		return ""
	case "azurerm_network_manager_ipam_pool":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/manager1/ipamPools/pool1
		return ""
	case "azurerm_network_manager_ipam_pool_static_cidr":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/manager1/ipamPools/pool1/staticCidrs/cidr1
		return ""
	case "azurerm_network_manager_management_group_connection":
		// Azure Resource ID: /providers/Microsoft.Management/managementGroups/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/networkManagerConnections/networkManagerConnection1
		return ""
	case "azurerm_network_manager_network_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/networkManagers/networkManager1/networkGroups/networkGroup1
		return ""
	case "azurerm_network_manager_routing_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/manager1/routingConfigurations/conf1
		return ""
	case "azurerm_network_manager_routing_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/manager1/routingConfigurations/conf1/ruleCollections/collection1/rules/rule1
		return ""
	case "azurerm_network_manager_routing_rule_collection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/manager1/routingConfigurations/conf1/ruleCollections/collection1
		return ""
	case "azurerm_network_manager_scope_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/networkManagers/networkManager1/scopeConnections/scopeConnection1
		return ""
	case "azurerm_network_manager_security_admin_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/networkManagers/networkManager1/securityAdminConfigurations/configuration1
		return ""
	case "azurerm_network_manager_static_member":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/networkManagers/networkManager1/networkGroups/networkGroup1/staticMembers/staticMember1
		return ""
	case "azurerm_network_manager_subscription_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/networkManagerConnections/networkManagerConnection1
		return ""
	case "azurerm_network_manager_verifier_workspace":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/manager1/verifierWorkspaces/workspace1
		return ""
	case "azurerm_network_manager_verifier_workspace_reachability_analysis_intent":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/manager1/verifierWorkspaces/workspace1/reachabilityAnalysisIntents/intent1
		return ""
	case "azurerm_network_profile":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/networkProfiles/examplenetprofile
		return ""
	case "azurerm_network_security_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkSecurityGroups/mySecurityGroup
		return ""
	case "azurerm_network_security_perimeter":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Network/networkSecurityPerimeters/example-nsp
		return ""
	case "azurerm_network_security_perimeter_access_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Network/networkSecurityPerimeters/example-nsp/profiles/defaultProfile/accessRules/example-accessrule
		return ""
	case "azurerm_network_security_perimeter_association":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Network/networkSecurityPerimeters/example-nsp/resourceAssociations/example-assoc
		return ""
	case "azurerm_network_security_perimeter_profile":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Network/networkSecurityPerimeters/example-nsp/profiles/defaultProfile
		return ""
	case "azurerm_network_security_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkSecurityGroups/mySecurityGroup/securityRules/rule1
		return ""
	case "azurerm_network_watcher":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkWatchers/watcher1
		return ""
	case "azurerm_network_watcher_flow_log":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkWatchers/watcher1/flowLogs/log1
		return ""
	case "azurerm_new_relic_monitor":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/NewRelic.Observability/monitors/monitor1
		return ""
	case "azurerm_new_relic_tag_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/NewRelic.Observability/monitors/monitor1/tagRules/ruleSet1
		return ""
	case "azurerm_nginx_api_key":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Nginx.NginxPlus/nginxDeployments/deploy1/apiKeys/key1
		return ""
	case "azurerm_nginx_certificate":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Nginx.NginxPlus/nginxDeployments/deploy1/certificates/cer1
		return ""
	case "azurerm_nginx_configuration":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Nginx.NginxPlus/nginxDeployments/dep1/configurations/default
		return ""
	case "azurerm_nginx_deployment":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Nginx.NginxPlus/nginxDeployments/dep1
		return ""
	case "azurerm_notification_hub":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.NotificationHubs/namespaces/namespace1/notificationHubs/hub1
		return ""
	case "azurerm_notification_hub_authorization_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.NotificationHubs/namespaces/namespace1/notificationHubs/hub1/authorizationRules/rule1
		return ""
	case "azurerm_notification_hub_namespace":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.NotificationHubs/namespaces/namespace1
		return ""
	case "azurerm_oracle_autonomous_database":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup/providers/Oracle.Database/autonomousDatabases/autonomousDatabases1
		return ""
	case "azurerm_oracle_autonomous_database_backup":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup/providers/Oracle.Database/autonomousDatabases/autonomousDatabase1/autonomousDatabaseBackups/autonomousDatabaseBackup1
		return ""
	case "azurerm_oracle_autonomous_database_clone_from_backup":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Oracle.Database/autonomousDatabases/example
		return ""
	case "azurerm_oracle_autonomous_database_clone_from_database":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Oracle.Database/autonomousDatabases/example
		return ""
	case "azurerm_oracle_cloud_vm_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup/providers/Oracle.Database/cloudVmClusters/cloudVmClusters1
		return ""
	case "azurerm_oracle_exadata_infrastructure":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup/providers/Oracle.Database/cloudExadataInfrastructures/cloudExadataInfrastructures1
		return ""
	case "azurerm_oracle_exascale_database_storage_vault":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup/providers/Oracle.Database/exascaleDbStorageVaults/exascaleDbStorageVaults1
		return ""
	case "azurerm_oracle_resource_anchor":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Oracle.Database/resourceAnchors/example
		return ""
	case "azurerm_orbital_contact":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Orbital/spacecrafts/spacecraft1/contacts/contact1
		return ""
	case "azurerm_orbital_contact_profile":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Orbital/contactProfiles/contactProfile1
		return ""
	case "azurerm_orbital_spacecraft":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Orbital/spacecrafts/spacecraft1
		return ""
	case "azurerm_orchestrated_virtual_machine_scale_set":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/virtualMachineScaleSets/scaleset1
		return ""
	case "azurerm_palo_alto_local_rulestack":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/myLocalRulestack
		return ""
	case "azurerm_palo_alto_local_rulestack_certificate":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/myLocalRulestack/certificates/myCertificate
		return ""
	case "azurerm_palo_alto_local_rulestack_fqdn_list":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/myLocalRulestack/fqdnLists/myFQDNList1
		return ""
	case "azurerm_palo_alto_local_rulestack_outbound_trust_certificate_association":
		// No standard import format found in documentation
		return ""
	case "azurerm_palo_alto_local_rulestack_outbound_untrust_certificate_association":
		// No standard import format found in documentation
		return ""
	case "azurerm_palo_alto_local_rulestack_prefix_list":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/myLocalRulestack/prefixLists/myFQDNList1
		return ""
	case "azurerm_palo_alto_local_rulestack_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/myLocalRulestack/localRules/myRule1
		return ""
	case "azurerm_palo_alto_next_generation_firewall_virtual_hub_local_rulestack":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/PaloAltoNetworks.Cloudngfw/firewalls/myVhubRulestackFW
		return ""
	case "azurerm_palo_alto_next_generation_firewall_virtual_hub_panorama":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/PaloAltoNetworks.Cloudngfw/firewalls/myVhubPanoramaFW
		return ""
	case "azurerm_palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/PaloAltoNetworks.Cloudngfw/firewalls/myVNetStrataCloudManagerFW
		return ""
	case "azurerm_palo_alto_next_generation_firewall_virtual_network_local_rulestack":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/PaloAltoNetworks.Cloudngfw/firewalls/myVNetRulestackFW
		return ""
	case "azurerm_palo_alto_next_generation_firewall_virtual_network_panorama":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/PaloAltoNetworks.Cloudngfw/firewalls/myVNetPanoramaFW
		return ""
	case "azurerm_palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/PaloAltoNetworks.Cloudngfw/firewalls/myVNetStrataCloudManagerFW
		return ""
	case "azurerm_palo_alto_virtual_network_appliance":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkVirtualAppliances/myPANetworkVirtualAppliance
		return ""
	case "azurerm_pim_active_role_assignment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000|/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/roleDefinitions/00000000-0000-0000-0000-000000000000|00000000-0000-0000-0000-000000000000
		return ""
	case "azurerm_pim_eligible_role_assignment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000|/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/roleDefinitions/00000000-0000-0000-0000-000000000000|00000000-0000-0000-0000-000000000000
		return ""
	case "azurerm_point_to_site_vpn_gateway":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/p2sVpnGateways/gateway1
		return ""
	case "azurerm_policy_definition":
		// Azure Resource ID: /subscriptions/<SUBSCRIPTION_ID>/providers/Microsoft.Authorization/policyDefinitions/<POLICY_NAME>
		return ""
	case "azurerm_policy_set_definition":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName
		return ""
	case "azurerm_policy_virtual_machine_configuration_assignment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/assignment1
		return ""
	case "azurerm_portal_dashboard":
		// No standard import format found in documentation
		return ""
	case "azurerm_portal_tenant_configuration":
		// Azure Resource ID: /providers/Microsoft.Portal/tenantConfigurations/default
		return ""
	case "azurerm_postgresql_active_directory_administrator":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.DBforPostgreSQL/servers/myserver
		return ""
	case "azurerm_postgresql_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/servers/server1/configurations/backslash_quote
		return ""
	case "azurerm_postgresql_database":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/servers/server1/databases/database1
		return ""
	case "azurerm_postgresql_firewall_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/servers/server1/firewallRules/rule1
		return ""
	case "azurerm_postgresql_flexible_server":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/flexibleServers/server1
		return ""
	case "azurerm_postgresql_flexible_server_active_directory_administrator":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/myserver/administrators/objectId
		return ""
	case "azurerm_postgresql_flexible_server_backup":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforPostgreSQL/flexibleServers/fs1/backups/backup1
		return ""
	case "azurerm_postgresql_flexible_server_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/flexibleServers/server1/configurations/configuration1
		return ""
	case "azurerm_postgresql_flexible_server_database":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DBforPostgreSQL/flexibleServers/flexibleServer1/databases/database1
		return ""
	case "azurerm_postgresql_flexible_server_firewall_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforPostgreSQL/flexibleServers/flexibleServer1/firewallRules/firewallRule1
		return ""
	case "azurerm_postgresql_flexible_server_virtual_endpoint":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforPostgreSQL/flexibleServers/sourceServerName/virtualEndpoints/endpointName|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforPostgreSQL/flexibleServers/replicaServerName/virtualEndpoints/endpointName"
		return ""
	case "azurerm_postgresql_server":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/servers/server1
		return ""
	case "azurerm_postgresql_server_key":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforPostgreSQL/servers/server1/keys/keyvaultname_key-name_keyversion
		return ""
	case "azurerm_postgresql_virtual_network_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.DBforPostgreSQL/servers/myserver/virtualNetworkRules/vnetrulename
		return ""
	case "azurerm_powerbi_embedded":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.PowerBIDedicated/capacities/capacity1
		return ""
	case "azurerm_private_dns_a_record":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/zone1/A/myrecord1
		return ""
	case "azurerm_private_dns_aaaa_record":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/zone1/AAAA/myrecord1
		return ""
	case "azurerm_private_dns_cname_record":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/zone1/CNAME/myrecord1
		return ""
	case "azurerm_private_dns_mx_record":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/contoso.com/MX/@
		return ""
	case "azurerm_private_dns_ptr_record":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/2.0.192.in-addr.arpa/PTR/15
		return ""
	case "azurerm_private_dns_resolver":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/dnsResolvers/dnsResolver1
		return ""
	case "azurerm_private_dns_resolver_dns_forwarding_ruleset":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/dnsForwardingRulesets/dnsForwardingRuleset1
		return ""
	case "azurerm_private_dns_resolver_forwarding_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/dnsForwardingRulesets/dnsForwardingRuleset1/forwardingRules/forwardingRule1
		return ""
	case "azurerm_private_dns_resolver_inbound_endpoint":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/dnsResolvers/dnsResolver1/inboundEndpoints/inboundEndpoint1
		return ""
	case "azurerm_private_dns_resolver_outbound_endpoint":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/dnsResolvers/dnsResolver1/outboundEndpoints/outboundEndpoint1
		return ""
	case "azurerm_private_dns_resolver_virtual_network_link":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/dnsForwardingRulesets/dnsForwardingRuleset1/virtualNetworkLinks/virtualNetworkLink1
		return ""
	case "azurerm_private_dns_srv_record":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/contoso.com/SRV/test
		return ""
	case "azurerm_private_dns_txt_record":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/contoso.com/TXT/test
		return ""
	case "azurerm_private_dns_zone":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/zone1
		return ""
	case "azurerm_private_dns_zone_virtual_network_link":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/zone1.com/virtualNetworkLinks/myVnetLink1
		return ""
	case "azurerm_private_endpoint":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/privateEndpoints/endpoint1
		return ""
	case "azurerm_private_endpoint_application_security_group_association":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/privateEndpoints/endpoints1|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/applicationSecurityGroups/securityGroup1",
		return ""
	case "azurerm_private_link_service":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/privateLinkServices/service1
		return ""
	case "azurerm_proximity_placement_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/proximityPlacementGroups/example-ppg
		return ""
	case "azurerm_public_ip":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/publicIPAddresses/myPublicIpAddress1
		return ""
	case "azurerm_public_ip_prefix":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/publicIPPrefixes/myPublicIpPrefix1
		return ""
	case "azurerm_purview_account":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Purview/accounts/account1
		return ""
	case "azurerm_qumulo_file_system":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Qumulo.Storage/fileSystems/example
		return ""
	case "azurerm_recovery_services_vault":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/vault1
		return ""
	case "azurerm_recovery_services_vault_resource_guard_association":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/vault1/backupResourceGuardProxies/proxy1
		return ""
	case "azurerm_redhat_openshift_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.RedHatOpenShift/openShiftClusters/cluster1
		return ""
	case "azurerm_redis_cache":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/redis/cache1
		return ""
	case "azurerm_redis_cache_access_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/redis/cache1/accessPolicies/policy1
		return ""
	case "azurerm_redis_cache_access_policy_assignment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/redis/cache1/accessPolicyAssignments/assignment1
		return ""
	case "azurerm_redis_enterprise_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/redisEnterprise/cluster1
		return ""
	case "azurerm_redis_enterprise_database":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/redisEnterprise/cluster1/databases/database1
		return ""
	case "azurerm_redis_firewall_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/redis/cache1/firewallRules/rule1
		return ""
	case "azurerm_redis_linked_server":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/redis/cache1/linkedServers/cache2
		return ""
	case "azurerm_relay_hybrid_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Relay/namespaces/relay1/hybridConnections/hconn1
		return ""
	case "azurerm_relay_hybrid_connection_authorization_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Relay/namespaces/namespace1/hybridConnections/connection1/authorizationRules/rule1
		return ""
	case "azurerm_relay_namespace":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Relay/namespaces/relay1
		return ""
	case "azurerm_relay_namespace_authorization_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Relay/namespaces/namespace1/authorizationRules/rule1
		return ""
	case "azurerm_resource_deployment_script_azure_cli":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Resources/deploymentScripts/script1
		return ""
	case "azurerm_resource_deployment_script_azure_power_shell":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Resources/deploymentScripts/script1
		return ""
	case "azurerm_resource_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1
		return ""
	case "azurerm_resource_group_cost_management_export":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.CostManagement/exports/export1
		return ""
	case "azurerm_resource_group_cost_management_view":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.CostManagement/views/costmanagementview
		return ""
	case "azurerm_resource_group_policy_assignment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Authorization/policyAssignments/assignment1
		return ""
	case "azurerm_resource_group_policy_exemption":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Authorization/policyExemptions/exemption1
		return ""
	case "azurerm_resource_group_policy_remediation":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.PolicyInsights/remediations/remediation1
		return ""
	case "azurerm_resource_group_template_deployment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Resources/deployments/template1
		return ""
	case "azurerm_resource_management_private_link":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Authorization/resourceManagementPrivateLinks/link1
		return ""
	case "azurerm_resource_management_private_link_association":
		// Azure Resource ID: /providers/Microsoft.Management/managementGroups/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/privateLinkAssociations/00000000-0000-0000-0000-000000000000
		return ""
	case "azurerm_resource_policy_assignment":
		// Azure Resource ID: "{resource}/providers/Microsoft.Authorization/policyAssignments/assignment1"
		return ""
	case "azurerm_resource_policy_exemption":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Authorization/policyExemptions/exemption1
		return ""
	case "azurerm_resource_policy_remediation":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.PolicyInsights/remediations/remediation1
		return ""
	case "azurerm_resource_provider_feature_registration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Features/providers/{ResourceProviderName}/features/{FeatureName}
		return ""
	case "azurerm_resource_provider_registration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.PolicyInsights
		return ""
	case "azurerm_role_assignment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/roleAssignments/00000000-0000-0000-0000-000000000000
		return ""
	case "azurerm_role_definition":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/roleDefinitions/00000000-0000-0000-0000-000000000000|/subscriptions/00000000-0000-0000-0000-000000000000"
		return ""
	case "azurerm_role_management_policy":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/roleDefinitions/00000000-0000-0000-0000-000000000000|<scope>"
		return ""
	case "azurerm_route":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/routeTables/mytable1/routes/myroute1
		return ""
	case "azurerm_route_filter":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/routeFilters/routeFilter1
		return ""
	case "azurerm_route_map":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1
		return ""
	case "azurerm_route_server":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/routeServer1
		return ""
	case "azurerm_route_server_bgp_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/routeServer1/bgpConnections/connection1
		return ""
	case "azurerm_route_table":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/routeTables/mytable1
		return ""
	case "azurerm_search_service":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Search/searchServices/service1
		return ""
	case "azurerm_search_shared_private_link_service":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Search/searchServices/service1/sharedPrivateLinkResources/resource1
		return ""
	case "azurerm_security_center_assessment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/providers/Microsoft.Security/assessments/00000000-0000-0000-0000-000000000000
		return ""
	case "azurerm_security_center_assessment_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Security/assessmentMetadata/metadata1
		return ""
	case "azurerm_security_center_auto_provisioning":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Security/autoProvisioningSettings/default
		return ""
	case "azurerm_security_center_automation":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Security/automations/automation1
		return ""
	case "azurerm_security_center_contact":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Security/securityContacts/default1
		return ""
	case "azurerm_security_center_server_vulnerability_assessment_virtual_machine":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.Compute/virtualMachines/vm-name/providers/Microsoft.Security/serverVulnerabilityAssessments/Default
		return ""
	case "azurerm_security_center_server_vulnerability_assessments_setting":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Security/serverVulnerabilityAssessmentsSettings/AzureServersSetting
		return ""
	case "azurerm_security_center_setting":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Security/settings/<setting_name>
		return ""
	case "azurerm_security_center_storage_defender":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Storage/storageAccounts/storageacc
		return ""
	case "azurerm_security_center_subscription_pricing":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Security/pricings/<resource_type>
		return ""
	case "azurerm_security_center_workspace":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Security/workspaceSettings/default
		return ""
	case "azurerm_sentinel_alert_rule_anomaly_built_in":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/securityMLAnalyticsSettings/setting1
		return ""
	case "azurerm_sentinel_alert_rule_anomaly_duplicate":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/securityMLAnalyticsSettings/setting1
		return ""
	case "azurerm_sentinel_alert_rule_fusion":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
		return ""
	case "azurerm_sentinel_alert_rule_machine_learning_behavior_analytics":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
		return ""
	case "azurerm_sentinel_alert_rule_ms_security_incident":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
		return ""
	case "azurerm_sentinel_alert_rule_nrt":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
		return ""
	case "azurerm_sentinel_alert_rule_scheduled":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
		return ""
	case "azurerm_sentinel_alert_rule_threat_intelligence":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
		return ""
	case "azurerm_sentinel_automation_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/automationRules/rule1
		return ""
	case "azurerm_sentinel_data_connector_aws_cloud_trail":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
		return ""
	case "azurerm_sentinel_data_connector_aws_s3":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
		return ""
	case "azurerm_sentinel_data_connector_azure_active_directory":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
		return ""
	case "azurerm_sentinel_data_connector_azure_advanced_threat_protection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
		return ""
	case "azurerm_sentinel_data_connector_azure_security_center":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
		return ""
	case "azurerm_sentinel_data_connector_dynamics_365":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
		return ""
	case "azurerm_sentinel_data_connector_iot":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
		return ""
	case "azurerm_sentinel_data_connector_microsoft_cloud_app_security":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
		return ""
	case "azurerm_sentinel_data_connector_microsoft_defender_advanced_threat_protection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
		return ""
	case "azurerm_sentinel_data_connector_microsoft_threat_intelligence":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
		return ""
	case "azurerm_sentinel_data_connector_microsoft_threat_protection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
		return ""
	case "azurerm_sentinel_data_connector_office_365":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
		return ""
	case "azurerm_sentinel_data_connector_office_365_project":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
		return ""
	case "azurerm_sentinel_data_connector_office_atp":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
		return ""
	case "azurerm_sentinel_data_connector_office_irm":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
		return ""
	case "azurerm_sentinel_data_connector_office_power_bi":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
		return ""
	case "azurerm_sentinel_data_connector_threat_intelligence":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
		return ""
	case "azurerm_sentinel_data_connector_threat_intelligence_taxii":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
		return ""
	case "azurerm_sentinel_log_analytics_workspace_onboarding":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/onboardingStates/defaults
		return ""
	case "azurerm_sentinel_metadata":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/metadata/metadata1
		return ""
	case "azurerm_sentinel_threat_intelligence_indicator":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/threatIntelligence/main/indicators/indicator1
		return ""
	case "azurerm_sentinel_watchlist":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/watchlists/list1
		return ""
	case "azurerm_sentinel_watchlist_item":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/watchlists/list1/watchlistItems/item1
		return ""
	case "azurerm_service_fabric_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceFabric/clusters/cluster1
		return ""
	case "azurerm_service_fabric_managed_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.ServiceFabric/managedClusters/clusterName1
		return ""
	case "azurerm_service_plan":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/serverFarms/farm1
		return ""
	case "azurerm_servicebus_namespace":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceBus/namespaces/sbns1
		return ""
	case "azurerm_servicebus_namespace_authorization_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceBus/namespaces/namespace1/authorizationRules/rule1
		return ""
	case "azurerm_servicebus_namespace_customer_managed_key":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceBus/namespaces/sbns1
		return ""
	case "azurerm_servicebus_namespace_disaster_recovery_config":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceBus/namespaces/namespace1/disasterRecoveryConfigs/config1
		return ""
	case "azurerm_servicebus_queue":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceBus/namespaces/sbns1/queues/snqueue1
		return ""
	case "azurerm_servicebus_queue_authorization_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceBus/namespaces/namespace1/queues/queue1/authorizationRules/rule1
		return ""
	case "azurerm_servicebus_subscription":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceBus/namespaces/sbns1/topics/sntopic1/subscriptions/sbsub1
		return ""
	case "azurerm_servicebus_subscription_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceBus/namespaces/sbns1/topics/sntopic1/subscriptions/sbsub1/rules/sbrule1
		return ""
	case "azurerm_servicebus_topic":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceBus/namespaces/sbns1/topics/sntopic1
		return ""
	case "azurerm_servicebus_topic_authorization_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceBus/namespaces/namespace1/topics/topic1/authorizationRules/rule1
		return ""
	case "azurerm_shared_image":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/galleries/gallery1/images/image1
		return ""
	case "azurerm_shared_image_gallery":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/galleries/gallery1
		return ""
	case "azurerm_shared_image_version":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/galleries/gallery1/images/image1/versions/1.2.3
		return ""
	case "azurerm_signalr_service":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/terraform-signalr/providers/Microsoft.SignalRService/signalR/tfex-signalr
		return ""
	case "azurerm_signalr_service_custom_certificate":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/signalR/signalr1/customCertificates/cert1
		return ""
	case "azurerm_signalr_service_custom_domain":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/signalR/signalr1/customDomains/customDomain1
		return ""
	case "azurerm_signalr_service_network_acl":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/signalR/signalr1
		return ""
	case "azurerm_signalr_shared_private_link_resource":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/signalR/signalr1/sharedPrivateLinkResources/resource1
		return ""
	case "azurerm_site_recovery_fabric":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/fabric-name
		return ""
	case "azurerm_site_recovery_hyperv_network_mapping":
		// No standard import format found in documentation
		return ""
	case "azurerm_site_recovery_hyperv_replication_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationPolicies/policy-name
		return ""
	case "azurerm_site_recovery_hyperv_replication_policy_association":
		// No standard import format found in documentation
		return ""
	case "azurerm_site_recovery_network_mapping":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/primary-fabric-name/replicationNetworks/azureNetwork/replicationNetworkMappings/mapping-name
		return ""
	case "azurerm_site_recovery_protection_container":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/fabric-name/replicationProtectionContainers/protection-container-name
		return ""
	case "azurerm_site_recovery_protection_container_mapping":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/fabric1/replicationProtectionContainers/container1/replicationProtectionContainerMappings/mapping1
		return ""
	case "azurerm_site_recovery_replicated_vm":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/fabric-name/replicationProtectionContainers/protection-container-name/replicationProtectedItems/vm-replication-name
		return ""
	case "azurerm_site_recovery_replication_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationPolicies/policy-name
		return ""
	case "azurerm_site_recovery_replication_recovery_plan":
		// No standard import format found in documentation
		return ""
	case "azurerm_site_recovery_services_vault_hyperv_site":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric1
		return ""
	case "azurerm_site_recovery_vmware_replicated_vm":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/fabric-name/replicationProtectionContainers/protection-container-name/replicationProtectedItems/vm-replication-name
		return ""
	case "azurerm_site_recovery_vmware_replication_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/vault1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/policy1
		return ""
	case "azurerm_site_recovery_vmware_replication_policy_association":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/site-name/replicationProtectionContainers/container-name/replicationProtectionContainerMappings/mapping-name
		return ""
	case "azurerm_snapshot":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/snapshots/snapshot1
		return ""
	case "azurerm_source_control_token":
		// Azure Resource ID: /providers/Microsoft.Web/sourceControls/GitHub
		return ""
	case "azurerm_spring_cloud_accelerator":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/applicationAccelerators/default
		return ""
	case "azurerm_spring_cloud_active_deployment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/spring/service1/apps/app1
		return ""
	case "azurerm_spring_cloud_api_portal":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/apiPortals/apiPortal1
		return ""
	case "azurerm_spring_cloud_api_portal_custom_domain":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/apiPortals/apiPortal1/domains/domain1
		return ""
	case "azurerm_spring_cloud_app":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.AppPlatform/spring/myservice/apps/myapp
		return ""
	case "azurerm_spring_cloud_app_cosmosdb_association":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/spring/service1/apps/app1/bindings/bind1
		return ""
	case "azurerm_spring_cloud_app_dynamics_application_performance_monitoring":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AppPlatform/spring/service1/apms/apm1
		return ""
	case "azurerm_spring_cloud_app_mysql_association":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/spring/service1/apps/app1/bindings/bind1
		return ""
	case "azurerm_spring_cloud_app_redis_association":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.AppPlatform/spring/myservice/apps/myapp/bindings/bind1
		return ""
	case "azurerm_spring_cloud_application_insights_application_performance_monitoring":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AppPlatform/spring/service1/apms/apm1
		return ""
	case "azurerm_spring_cloud_application_live_view":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/applicationLiveViews/default
		return ""
	case "azurerm_spring_cloud_build_deployment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/deployments/deploy1
		return ""
	case "azurerm_spring_cloud_build_pack_binding":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/buildServices/buildService1/builders/builder1/buildPackBindings/binding1
		return ""
	case "azurerm_spring_cloud_builder":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/buildServices/buildService1/builders/builder1
		return ""
	case "azurerm_spring_cloud_certificate":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/spring/spring1/certificates/cert1
		return ""
	case "azurerm_spring_cloud_configuration_service":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/configurationServices/configurationService1
		return ""
	case "azurerm_spring_cloud_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AppPlatform/Spring/springcloud/apps/springcloudapp/deployments/deployment/providers/Microsoft.ServiceLinker/linkers/serviceconnector1
		return ""
	case "azurerm_spring_cloud_container_deployment":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/deployments/deploy1
		return ""
	case "azurerm_spring_cloud_custom_domain":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/domains/domain.com
		return ""
	case "azurerm_spring_cloud_customized_accelerator":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/applicationAccelerators/default/customizedAccelerators/customizedAccelerator1
		return ""
	case "azurerm_spring_cloud_dev_tool_portal":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/Spring/service1/DevToolPortals/default
		return ""
	case "azurerm_spring_cloud_dynatrace_application_performance_monitoring":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AppPlatform/spring/service1/apms/apm1
		return ""
	case "azurerm_spring_cloud_elastic_application_performance_monitoring":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AppPlatform/spring/service1/apms/apm1
		return ""
	case "azurerm_spring_cloud_gateway":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/gateways/gateway1
		return ""
	case "azurerm_spring_cloud_gateway_custom_domain":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/gateways/gateway1/domains/domain1
		return ""
	case "azurerm_spring_cloud_gateway_route_config":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/gateways/gateway1/routeConfigs/routeConfig1
		return ""
	case "azurerm_spring_cloud_java_deployment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/spring/service1/apps/app1/deployments/deploy1
		return ""
	case "azurerm_spring_cloud_new_relic_application_performance_monitoring":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AppPlatform/spring/service1/apms/apm1
		return ""
	case "azurerm_spring_cloud_service":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AppPlatform/spring/spring1
		return ""
	case "azurerm_spring_cloud_storage":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/storages/storage1
		return ""
	case "azurerm_ssh_public_key":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/sshPublicKeys/mySshPublicKeyName1
		return ""
	case "azurerm_stack_hci_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AzureStackHCI/clusters/cluster1
		return ""
	case "azurerm_stack_hci_deployment_setting":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.AzureStackHCI/clusters/clus1/deploymentSettings/default
		return ""
	case "azurerm_stack_hci_extension":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AzureStackHCI/clusters/cluster1/arcSettings/default/extensions/extension1
		return ""
	case "azurerm_stack_hci_logical_network":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AzureStackHCI/logicalNetworks/ln1
		return ""
	case "azurerm_stack_hci_marketplace_gallery_image":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.AzureStackHCI/marketplaceGalleryImages/image1
		return ""
	case "azurerm_stack_hci_network_interface":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AzureStackHCI/networkInterfaces/ni1
		return ""
	case "azurerm_stack_hci_storage_path":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AzureStackHCI/storageContainers/storage1
		return ""
	case "azurerm_stack_hci_virtual_hard_disk":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.AzureStackHCI/virtualHardDisks/disk1
		return ""
	case "azurerm_static_site":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Web/staticSites/my-static-site1
		return ""
	case "azurerm_static_site_custom_domain":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Web/staticSites/my-static-site1/customDomains/name.contoso.com
		return ""
	case "azurerm_static_web_app":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Web/staticSites/my-static-site1
		return ""
	case "azurerm_static_web_app_custom_domain":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Web/staticSites/my-static-site1/customDomains/name.contoso.com
		return ""
	case "azurerm_static_web_app_function_app_registration":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Web/staticSites/my-static-site1/userProvidedFunctionApps/myFunctionApp
		return ""
	case "azurerm_storage_account":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Storage/storageAccounts/myaccount
		return ""
	case "azurerm_storage_account_customer_managed_key":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Storage/storageAccounts/myaccount
		return ""
	case "azurerm_storage_account_local_user":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Storage/storageAccounts/storageAccount1/localUsers/user1
		return ""
	case "azurerm_storage_account_network_rules":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Storage/storageAccounts/myaccount
		return ""
	case "azurerm_storage_account_queue_properties":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Storage/storageAccounts/myaccount
		return ""
	case "azurerm_storage_account_static_website":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Storage/storageAccounts/myaccount
		return ""
	case "azurerm_storage_blob":
		// Azure Resource ID: https://example.blob.core.windows.net/container/blob.vhd
		return ""
	case "azurerm_storage_blob_inventory_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Storage/storageAccounts/storageAccount1
		return ""
	case "azurerm_storage_container":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Storage/storageAccounts/myaccount/blobServices/default/containers/mycontainer
		return ""
	case "azurerm_storage_container_immutability_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Storage/storageAccounts/myaccount/blobServices/default/containers/mycontainer/immutabilityPolicies/default
		return ""
	case "azurerm_storage_data_lake_gen2_filesystem":
		// Azure Resource ID: https://account1.dfs.core.windows.net/fileSystem1
		return ""
	case "azurerm_storage_data_lake_gen2_path":
		// Azure Resource ID: https://account1.dfs.core.windows.net/fileSystem1/path
		return ""
	case "azurerm_storage_encryption_scope":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Storage/storageAccounts/account1/encryptionScopes/scope1
		return ""
	case "azurerm_storage_management_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Storage/storageAccounts/myaccountname/managementPolicies/default
		return ""
	case "azurerm_storage_mover":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.StorageMover/storageMovers/storageMover1
		return ""
	case "azurerm_storage_mover_agent":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.StorageMover/storageMovers/storageMover1/agents/agent1
		return ""
	case "azurerm_storage_mover_job_definition":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.StorageMover/storageMovers/storageMover1/projects/project1/jobDefinitions/jobDefinition1
		return ""
	case "azurerm_storage_mover_project":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.StorageMover/storageMovers/storageMover1/projects/project1
		return ""
	case "azurerm_storage_mover_source_endpoint":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.StorageMover/storageMovers/storageMover1/endpoints/endpoint1
		return ""
	case "azurerm_storage_mover_target_endpoint":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.StorageMover/storageMovers/storageMover1/endpoints/endpoint1
		return ""
	case "azurerm_storage_object_replication":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Storage/storageAccounts/storageAccount1/objectReplicationPolicies/objectReplicationPolicy1;/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group2/providers/Microsoft.Storage/storageAccounts/storageAccount2/objectReplicationPolicies/objectReplicationPolicy2
		return ""
	case "azurerm_storage_queue":
		// Azure Resource ID: https://example.queue.core.windows.net/queue1
		return ""
	case "azurerm_storage_share":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Storage/storageAccounts/myAccount/fileServices/default/shares/exampleShare
		return ""
	case "azurerm_storage_share_directory":
		// Azure Resource ID: https://tomdevsa20.file.core.windows.net/share1/directory1
		return ""
	case "azurerm_storage_share_file":
		// Azure Resource ID: https://account1.file.core.windows.net/share1/file1
		return ""
	case "azurerm_storage_sync":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StorageSync/storageSyncServices/sync1
		return ""
	case "azurerm_storage_sync_cloud_endpoint":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StorageSync/storageSyncServices/sync1/syncGroups/syncgroup1/cloudEndpoints/cloudEndpoint1
		return ""
	case "azurerm_storage_sync_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.StorageSync/storageSyncServices/sync1/syncGroups/group1
		return ""
	case "azurerm_storage_sync_server_endpoint":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StorageSync/storageSyncServices/sync1/syncGroups/syncGroup1/serverEndpoints/endpoint1
		return ""
	case "azurerm_storage_table":
		// Azure Resource ID: "https://example.table.core.windows.net/Tables('replace-with-table-name')"
		return ""
	case "azurerm_storage_table_entity":
		// Azure Resource ID: https://example.table.core.windows.net/table1(PartitionKey='samplepartition',RowKey='samplerow')
		return ""
	case "azurerm_stream_analytics_cluster":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.StreamAnalytics/clusters/cluster1
		return ""
	case "azurerm_stream_analytics_function_javascript_uda":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/functions/func1
		return ""
	case "azurerm_stream_analytics_function_javascript_udf":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/functions/func1
		return ""
	case "azurerm_stream_analytics_job":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1
		return ""
	case "azurerm_stream_analytics_job_schedule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/schedule/default
		return ""
	case "azurerm_stream_analytics_job_storage_account":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1
		return ""
	case "azurerm_stream_analytics_managed_private_endpoint":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.StreamAnalytics/clusters/cluster1/privateEndpoints/endpoint1
		return ""
	case "azurerm_stream_analytics_output_blob":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/outputs/output1
		return ""
	case "azurerm_stream_analytics_output_cosmosdb":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/outputs/output1
		return ""
	case "azurerm_stream_analytics_output_eventhub":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/outputs/output1
		return ""
	case "azurerm_stream_analytics_output_function":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/outputs/output1
		return ""
	case "azurerm_stream_analytics_output_mssql":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/outputs/output1
		return ""
	case "azurerm_stream_analytics_output_powerbi":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/outputs/output1
		return ""
	case "azurerm_stream_analytics_output_servicebus_queue":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/outputs/output1
		return ""
	case "azurerm_stream_analytics_output_servicebus_topic":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/outputs/output1
		return ""
	case "azurerm_stream_analytics_output_synapse":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/outputs/output1
		return ""
	case "azurerm_stream_analytics_output_table":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/outputs/output1
		return ""
	case "azurerm_stream_analytics_reference_input_blob":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/inputs/input1
		return ""
	case "azurerm_stream_analytics_reference_input_mssql":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/inputs/input1
		return ""
	case "azurerm_stream_analytics_stream_input_blob":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/inputs/input1
		return ""
	case "azurerm_stream_analytics_stream_input_eventhub":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/inputs/input1
		return ""
	case "azurerm_stream_analytics_stream_input_eventhub_v2":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/inputs/input1
		return ""
	case "azurerm_stream_analytics_stream_input_iothub":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/inputs/input1
		return ""
	case "azurerm_subnet":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1/subnets/mysubnet1
		return ""
	case "azurerm_subnet_nat_gateway_association":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1/subnets/mysubnet1
		return ""
	case "azurerm_subnet_network_security_group_association":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1/subnets/mysubnet1
		return ""
	case "azurerm_subnet_route_table_association":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1/subnets/mysubnet1
		return ""
	case "azurerm_subnet_service_endpoint_storage_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/serviceEndpointPolicies/policy1
		return ""
	case "azurerm_subscription":
		// Azure Resource ID: "/providers/Microsoft.Subscription/aliases/subscription1"
		return ""
	case "azurerm_subscription_cost_management_export":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CostManagement/exports/export1
		return ""
	case "azurerm_subscription_cost_management_view":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CostManagement/views/costmanagementview
		return ""
	case "azurerm_subscription_policy_assignment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-000000000000/providers/Microsoft.Authorization/policyAssignments/assignment1
		return ""
	case "azurerm_subscription_policy_exemption":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-000000000000/providers/Microsoft.Authorization/policyExemptions/exemption1
		return ""
	case "azurerm_subscription_policy_remediation":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.PolicyInsights/remediations/remediation1
		return ""
	case "azurerm_subscription_template_deployment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Resources/deployments/template1
		return ""
	case "azurerm_synapse_firewall_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.Synapse/workspaces/workspace1/firewallRules/rule1
		return ""
	case "azurerm_synapse_integration_runtime_azure":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/integrationRuntimes/IntegrationRuntime1
		return ""
	case "azurerm_synapse_integration_runtime_self_hosted":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/integrationRuntimes/IntegrationRuntime1
		return ""
	case "azurerm_synapse_linked_service":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/linkedServices/linkedservice1
		return ""
	case "azurerm_synapse_managed_private_endpoint":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/managedVirtualNetworks/default/managedPrivateEndpoints/endpoint1
		return ""
	case "azurerm_synapse_private_link_hub":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/privateLinkHubs/privateLinkHub1
		return ""
	case "azurerm_synapse_role_assignment":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1|000000000000"
		return ""
	case "azurerm_synapse_spark_pool":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/bigDataPools/sparkPool1
		return ""
	case "azurerm_synapse_sql_pool":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1
		return ""
	case "azurerm_synapse_sql_pool_extended_auditing_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/extendedAuditingSettings/default
		return ""
	case "azurerm_synapse_sql_pool_security_alert_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/securityAlertPolicies/default
		return ""
	case "azurerm_synapse_sql_pool_vulnerability_assessment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/vulnerabilityAssessments/default
		return ""
	case "azurerm_synapse_sql_pool_vulnerability_assessment_baseline":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/vulnerabilityAssessments/default/rules/rule1/baselines/baseline1
		return ""
	case "azurerm_synapse_sql_pool_workload_classifier":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/workloadGroups/workloadGroup1/workloadClassifiers/workloadClassifier1
		return ""
	case "azurerm_synapse_sql_pool_workload_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/workloadGroups/workloadGroup1
		return ""
	case "azurerm_synapse_workspace":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1
		return ""
	case "azurerm_synapse_workspace_aad_admin":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Synapse/workspaces/workspace1/administrators/activeDirectory
		return ""
	case "azurerm_synapse_workspace_extended_auditing_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/extendedAuditingSettings/default
		return ""
	case "azurerm_synapse_workspace_key":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/keys/key1
		return ""
	case "azurerm_synapse_workspace_security_alert_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/securityAlertPolicies/Default
		return ""
	case "azurerm_synapse_workspace_sql_aad_admin":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlAdministrators/activeDirectory
		return ""
	case "azurerm_synapse_workspace_vulnerability_assessment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/vulnerabilityAssessments/default
		return ""
	case "azurerm_system_center_virtual_machine_manager_availability_set":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.ScVmm/availabilitySets/availabilitySet1
		return ""
	case "azurerm_system_center_virtual_machine_manager_cloud":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.ScVmm/clouds/cloud1
		return ""
	case "azurerm_system_center_virtual_machine_manager_server":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.ScVmm/vmmServers/vmmServer1
		return ""
	case "azurerm_system_center_virtual_machine_manager_virtual_machine_instance":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.HybridCompute/machines/machine1/providers/Microsoft.ScVmm/virtualMachineInstances/default
		return ""
	case "azurerm_system_center_virtual_machine_manager_virtual_machine_instance_guest_agent":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.HybridCompute/machines/machine1/providers/Microsoft.ScVmm/virtualMachineInstances/default/guestAgents/default
		return ""
	case "azurerm_system_center_virtual_machine_manager_virtual_machine_template":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.ScVmm/virtualMachineTemplates/virtualMachineTemplate1
		return ""
	case "azurerm_system_center_virtual_machine_manager_virtual_network":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.ScVmm/virtualNetworks/virtualNetwork1
		return ""
	case "azurerm_tenant_template_deployment":
		// Azure Resource ID: /providers/Microsoft.Resources/deployments/deploy1
		return ""
	case "azurerm_traffic_manager_azure_endpoint":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.Network/trafficManagerProfiles/example-profile/AzureEndpoints/example-endpoint
		return ""
	case "azurerm_traffic_manager_external_endpoint":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-group/providers/Microsoft.Network/trafficManagerProfiles/example-profile/ExternalEndpoints/example-endpoint
		return ""
	case "azurerm_traffic_manager_nested_endpoint":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.Network/trafficManagerProfiles/example-profile/NestedEndpoints/example-endpoint
		return ""
	case "azurerm_traffic_manager_profile":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/trafficManagerProfiles/mytrafficmanagerprofile1
		return ""
	case "azurerm_trusted_signing_account":
		// Azure Resource ID: /subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.CodeSigning/codeSigningAccounts/example-account
		return ""
	case "azurerm_user_assigned_identity":
		// Azure Resource ID: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{userAssignedIdentityName}
		return ""
	case "azurerm_video_indexer_account":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.VideoIndexer/accounts/example-account-name
		return ""
	case "azurerm_virtual_desktop_application":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.DesktopVirtualization/applicationGroups/myapplicationgroup/applications/myapplication
		return ""
	case "azurerm_virtual_desktop_application_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.DesktopVirtualization/applicationGroups/myapplicationgroup
		return ""
	case "azurerm_virtual_desktop_host_pool":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.DesktopVirtualization/hostPools/myhostpool
		return ""
	case "azurerm_virtual_desktop_host_pool_registration_info":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DesktopVirtualization/hostPools/pool1/registrationInfo/default
		return ""
	case "azurerm_virtual_desktop_scaling_plan":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DesktopVirtualization/scalingPlans/plan1
		return ""
	case "azurerm_virtual_desktop_scaling_plan_host_pool_association":
		// Azure Resource ID: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DesktopVirtualization/scalingPlans/plan1|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.DesktopVirtualization/hostPools/myhostpool"
		return ""
	case "azurerm_virtual_desktop_workspace":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.DesktopVirtualization/workspaces/myworkspace
		return ""
	case "azurerm_virtual_desktop_workspace_application_group_association":
		// Azure Resource ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.DesktopVirtualization/workspaces/myworkspace|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.DesktopVirtualization/applicationGroups/myapplicationgroup"
		return ""
	case "azurerm_virtual_hub":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/hub1
		return ""
	case "azurerm_virtual_hub_bgp_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/virtualHub1/bgpConnections/connection1
		return ""
	case "azurerm_virtual_hub_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/hub1/hubVirtualNetworkConnections/connection1
		return ""
	case "azurerm_virtual_hub_ip":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/virtualHub1/ipConfigurations/ipConfig1
		return ""
	case "azurerm_virtual_hub_route_table":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/routeTable1
		return ""
	case "azurerm_virtual_hub_route_table_route":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/routeTable1/routes/routeName
		return ""
	case "azurerm_virtual_hub_routing_intent":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/virtualHubs/virtualHub1/routingIntent/routingIntent1
		return ""
	case "azurerm_virtual_hub_security_partner_provider":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/securityPartnerProviders/securityPartnerProvider1
		return ""
	case "azurerm_virtual_machine":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachines/machine1
		return ""
	case "azurerm_virtual_machine_automanage_configuration_assignment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.AutoManage/configurationProfileAssignments/default
		return ""
	case "azurerm_virtual_machine_data_disk_attachment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachines/machine1/dataDisks/disk1
		return ""
	case "azurerm_virtual_machine_extension":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachines/myVM/extensions/extensionName
		return ""
	case "azurerm_virtual_machine_gallery_application_assignment":
		// Azure Resource ID: subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachines/machine1|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/galleries/gallery1/applications/galleryApplication1/versions/galleryApplicationVersion1
		return ""
	case "azurerm_virtual_machine_implicit_data_disk_from_source":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachines/machine1/dataDisks/disk1
		return ""
	case "azurerm_virtual_machine_packet_capture":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkWatchers/watcher1/packetCaptures/capture1
		return ""
	case "azurerm_virtual_machine_restore_point":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/restorePointCollections/collection1/restorePoints/restorePoint1
		return ""
	case "azurerm_virtual_machine_restore_point_collection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/restorePointCollections/collection1
		return ""
	case "azurerm_virtual_machine_run_command":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachines/vm1/runCommands/rc1
		return ""
	case "azurerm_virtual_machine_scale_set":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachineScaleSets/scaleset1
		return ""
	case "azurerm_virtual_machine_scale_set_extension":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/virtualMachineScaleSets/scaleSet1/extensions/extension1
		return ""
	case "azurerm_virtual_machine_scale_set_packet_capture":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkWatchers/watcher1/packetCaptures/capture1
		return ""
	case "azurerm_virtual_machine_scale_set_standby_pool":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/standbyVirtualMachinePool1
		return ""
	case "azurerm_virtual_network":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1
		return ""
	case "azurerm_virtual_network_dns_servers":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1/dnsServers/default
		return ""
	case "azurerm_virtual_network_gateway":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.Network/virtualNetworkGateways/myGateway1
		return ""
	case "azurerm_virtual_network_gateway_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.Network/connections/myConnection1
		return ""
	case "azurerm_virtual_network_gateway_nat_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Network/virtualNetworkGateways/gw1/natRules/rule1
		return ""
	case "azurerm_virtual_network_peering":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1/virtualNetworkPeerings/myvnet1peering
		return ""
	case "azurerm_virtual_wan":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualWans/testvwan
		return ""
	case "azurerm_vmware_cluster":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AVS/privateClouds/privateCloud1/clusters/cluster1
		return ""
	case "azurerm_vmware_express_route_authorization":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AVS/privateClouds/privateCloud1/authorizations/authorization1
		return ""
	case "azurerm_vmware_netapp_volume_attachment":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/privateCloud1/clusters/Cluster1/dataStores/datastore1
		return ""
	case "azurerm_vmware_private_cloud":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/privateCloud1
		return ""
	case "azurerm_voice_services_communications_gateway":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.VoiceServices/communicationsGateways/communicationsGateway1
		return ""
	case "azurerm_voice_services_communications_gateway_test_line":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.VoiceServices/communicationsGateways/communicationsGateway1/testLines/testLine1
		return ""
	case "azurerm_vpn_gateway":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/vpnGateways/gateway1
		return ""
	case "azurerm_vpn_gateway_connection":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/vpnGateways/gateway1/vpnConnections/conn1
		return ""
	case "azurerm_vpn_gateway_nat_rule":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Network/vpnGateways/vpnGateway1/natRules/natRule1
		return ""
	case "azurerm_vpn_server_configuration":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/vpnServerConfigurations/config1
		return ""
	case "azurerm_vpn_server_configuration_policy_group":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Network/vpnServerConfigurations/serverConfiguration1/configurationPolicyGroups/configurationPolicyGroup1
		return ""
	case "azurerm_vpn_site":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/vpnSites/site1
		return ""
	case "azurerm_web_app_active_slot":
		// Azure Resource ID: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1"
		return ""
	case "azurerm_web_app_hybrid_connection":
		// Azure Resource ID: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/hybridConnectionNamespaces/hybridConnectionNamespace1/relays/relay1"
		return ""
	case "azurerm_web_application_firewall_policy":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Network/applicationGatewayWebApplicationFirewallPolicies/example-wafpolicy
		return ""
	case "azurerm_web_pubsub":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubSub/pubsub1
		return ""
	case "azurerm_web_pubsub_custom_certificate":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubSub/WebPubsub1/customCertificates/cert1
		return ""
	case "azurerm_web_pubsub_custom_domain":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubSub/webpubsub1/customDomains/customDomain1
		return ""
	case "azurerm_web_pubsub_hub":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubSub/webPubSub1/hubs/webPubSubhub1
		return ""
	case "azurerm_web_pubsub_network_acl":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubSub/webpubsub1
		return ""
	case "azurerm_web_pubsub_shared_private_link_resource":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubSub/webPubSub1/sharedPrivateLinkResources/resource1
		return ""
	case "azurerm_web_pubsub_socketio":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubSub/pubsub1
		return ""
	case "azurerm_windows_function_app":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1
		return ""
	case "azurerm_windows_function_app_slot":
		// Azure Resource ID: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/slots/slot1"
		return ""
	case "azurerm_windows_virtual_machine":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachines/machine1
		return ""
	case "azurerm_windows_virtual_machine_scale_set":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachineScaleSets/scaleset1
		return ""
	case "azurerm_windows_web_app":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1
		return ""
	case "azurerm_windows_web_app_slot":
		// Azure Resource ID: /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/slots/slot1
		return ""
	case "azurerm_workloads_sap_discovery_virtual_instance":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Workloads/sapVirtualInstances/vis1
		return ""
	case "azurerm_workloads_sap_single_node_virtual_instance":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Workloads/sapVirtualInstances/vis1
		return ""
	case "azurerm_workloads_sap_three_tier_virtual_instance":
		// Azure Resource ID: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Workloads/sapVirtualInstances/vis1
		return ""
	}
	return ""
}
