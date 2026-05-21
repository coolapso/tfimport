package providers

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	tfjson "github.com/hashicorp/terraform-json"
)

// resolveCustomextractAWSImportID provides a hook for hand-coded ID resolution.
// Return "" to fall back to the generated doc-cruncher logic.
func resolveCustomextractAWSImportID(ctx *ProviderContext, resourceType string, config map[string]any) string {
	switch resourceType {
	case "aws_iam_role_policy_attachment":
		role := resolveAttribute(ctx, config, "role")
		policyArn := resolveAttribute(ctx, config, "policy_arn")
		if role != "" && policyArn != "" {
			return fmt.Sprintf("%s/%s", role, policyArn)
		}
		return ""
	case "aws_iam_user_policy_attachment":
		user := resolveAttribute(ctx, config, "user")
		policyArn := resolveAttribute(ctx, config, "policy_arn")
		if user != "" && policyArn != "" {
			return fmt.Sprintf("%s/%s", user, policyArn)
		}
		return ""
	case "aws_iam_group_policy_attachment":
		group := resolveAttribute(ctx, config, "group")
		policyArn := resolveAttribute(ctx, config, "policy_arn")
		if group != "" && policyArn != "" {
			return fmt.Sprintf("%s/%s", group, policyArn)
		}
		return ""
	case "aws_iam_policy":
		name, _ := config["name"].(string)
		namePrefix, _ := config["name_prefix"].(string)
		path, _ := config["path"].(string)
		if path == "" {
			path = "/"
		}

		cleanPath := strings.TrimPrefix(path, "/")

		if name != "" {
			if ctx != nil {
				awsClient := ctx.GetAWSClient()
				if awsClient != nil && awsClient.AccountID != "" {
					return fmt.Sprintf("arn:%s:iam::%s:policy/%s%s", awsClient.Partition, awsClient.AccountID, cleanPath, name)
				}
			}
		} else if namePrefix != "" {
			if ctx != nil {
				awsClient := ctx.GetAWSClient()
				if awsClient != nil && awsClient.IAMClient != nil {
					res, err := awsClient.IAMClient.ListPolicies(ctx.Context, &iam.ListPoliciesInput{
						PathPrefix: aws.String(path),
						Scope:      "Local",
					})
					if err == nil {
						var matchedPolicyArn string
						matchCount := 0
						for _, p := range res.Policies {
							if p.PolicyName != nil && strings.HasPrefix(*p.PolicyName, namePrefix) {
								matchedPolicyArn = *p.Arn
								matchCount++
							}
						}
						if matchCount == 1 {
							return matchedPolicyArn
						} else if matchCount > 1 {
							log.Printf("Warning: found multiple IAM policies with prefix %q in path %q. Cannot unambiguously resolve ID.\n", namePrefix, path)
						}
					} else {
						log.Printf("Warning: failed to list IAM policies: %v\n", err)
					}
				}
			}
		}
	case "aws_sqs_queue":
		if ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.SQSClient != nil {
				return lookupSqsQueueUrl(ctx, config)
			}
		}
	case "aws_acm_certificate":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:acm:%s:%s:certificate/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_acmpca_certificate":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:acmpca:%s:%s:certificate/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_acmpca_certificate_authority":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:acmpca:%s:%s:certificate-authority/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_ami_launch_permission":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:ami:%s:%s:launch-permission/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_api_gateway_domain_name_access_association":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:api:%s:%s:gateway-domain-name-access-association/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_appfabric_app_bundle":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:appfabric:%s:%s:app-bundle/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_apprunner_auto_scaling_configuration_version":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:apprunner:%s:%s:auto-scaling-configuration-version/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_apprunner_observability_configuration":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:apprunner:%s:%s:observability-configuration/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_apprunner_service":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:apprunner:%s:%s:service/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_apprunner_vpc_connector":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:apprunner:%s:%s:vpc-connector/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_apprunner_vpc_ingress_connection":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:apprunner:%s:%s:vpc-ingress-connection/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_arcregionswitch_plan":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:arcregionswitch:%s:%s:plan/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_batch_job_definition":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:batch:%s:%s:job-definition/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_batch_job_queue":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:batch:%s:%s:job-queue/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_batch_scheduling_policy":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:batch:%s:%s:scheduling-policy/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_bcmdataexports_export":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:bcmdataexports:%s:%s:export/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_billing_view":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:billing:%s:%s:view/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_cloud9_environment_membership":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:cloud9:%s:%s:environment-membership/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_cloudformation_type":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:cloudformation:%s:%s:type/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_cloudfront_realtime_log_config":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:cloudfront:%s:%s:realtime-log-config/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_cloudtrail":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:cloudtrail:%s:%s:aws-cloudtrail/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_cloudtrail_event_data_store":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:cloudtrail:%s:%s:event-data-store/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_cloudwatch_log_anomaly_detector":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:cloudwatch:%s:%s:log-anomaly-detector/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_cloudwatch_log_resource_policy":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:cloudwatch:%s:%s:log-resource-policy/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_cloudwatch_query_definition":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:cloudwatch:%s:%s:query-definition/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_codeartifact_domain":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:codeartifact:%s:%s:domain/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_codeartifact_domain_permissions_policy":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:codeartifact:%s:%s:domain-permissions-policy/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_codeartifact_repository":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:codeartifact:%s:%s:repository/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_codeartifact_repository_permissions_policy":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:codeartifact:%s:%s:repository-permissions-policy/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_codebuild_report_group":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:codebuild:%s:%s:report-group/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_codebuild_resource_policy":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:codebuild:%s:%s:resource-policy/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_codebuild_source_credential":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:codebuild:%s:%s:source-credential/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_codeconnections_connection":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:codeconnections:%s:%s:connection/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_codeconnections_host":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:codeconnections:%s:%s:host/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_codepipeline_webhook":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:codepipeline:%s:%s:webhook/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_codestarconnections_connection":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:codestarconnections:%s:%s:connection/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_codestarconnections_host":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:codestarconnections:%s:%s:host/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_codestarnotifications_notification_rule":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:codestarnotifications:%s:%s:notification-rule/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_comprehend_document_classifier":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:comprehend:%s:%s:document-classifier/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_comprehend_entity_recognizer":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:comprehend:%s:%s:entity-recognizer/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_controltower_baseline":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:controltower:%s:%s:baseline/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_datasync_agent":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:datasync:%s:%s:agent/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_datasync_location_azure_blob":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:datasync:%s:%s:location-azure-blob/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_datasync_location_efs":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:datasync:%s:%s:location-efs/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_datasync_location_fsx_lustre_file_system":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:datasync:%s:%s:location-fsx-lustre-file-system/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_datasync_location_fsx_ontap_file_system":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:datasync:%s:%s:location-fsx-ontap-file-system/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_datasync_location_fsx_openzfs_file_system":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:datasync:%s:%s:location-fsx-openzfs-file-system/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_datasync_location_fsx_windows_file_system":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:datasync:%s:%s:location-fsx-windows-file-system/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_datasync_location_hdfs":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:datasync:%s:%s:location-hdfs/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_datasync_location_nfs":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:datasync:%s:%s:location-nfs/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_datasync_location_object_storage":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:datasync:%s:%s:location-object-storage/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_datasync_location_s3":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:datasync:%s:%s:location-s3/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_datasync_location_smb":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:datasync:%s:%s:location-smb/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_datasync_task":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:datasync:%s:%s:task/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_db_instance_automated_backups_replication":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:db:%s:%s:instance-automated-backups-replication/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_db_instance_role_association":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:db:%s:%s:instance-role-association/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_detective_graph":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:detective:%s:%s:graph/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_detective_invitation_accepter":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:detective:%s:%s:invitation-accepter/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_detective_member":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:detective:%s:%s:member/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_detective_organization_configuration":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:detective:%s:%s:organization-configuration/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_devicefarm_device_pool":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:devicefarm:%s:%s:device-pool/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_devicefarm_instance_profile":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:devicefarm:%s:%s:instance-profile/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_devicefarm_network_profile":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:devicefarm:%s:%s:network-profile/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_devicefarm_project":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:devicefarm:%s:%s:project/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_devicefarm_test_grid_project":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:devicefarm:%s:%s:test-grid-project/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_devicefarm_upload":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:devicefarm:%s:%s:upload/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_dms_replication_config":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:dms:%s:%s:replication-config/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_docdbelastic_cluster":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:docdbelastic:%s:%s:cluster/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_dynamodb_table_export":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:dynamodb:%s:%s:table-export/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_ebs_default_kms_key":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:ebs:%s:%s:default-kms-key/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_ecs_capacity_provider":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:ecs:%s:%s:capacity-provider/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_ecs_express_gateway_service":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:ecs:%s:%s:express-gateway-service/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_evidently_project":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:evidently:%s:%s:project/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_evidently_segment":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:evidently:%s:%s:segment/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_globalaccelerator_accelerator":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:globalaccelerator:%s:%s:accelerator/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_globalaccelerator_cross_account_attachment":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:globalaccelerator:%s:%s:cross-account-attachment/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_globalaccelerator_custom_routing_accelerator":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:globalaccelerator:%s:%s:custom-routing-accelerator/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_glue_registry":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:glue:%s:%s:registry/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_glue_schema":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:glue:%s:%s:schema/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_iam_openid_connect_provider":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:iam:%s:%s:openid-connect-provider/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_iam_saml_provider":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:iam:%s:%s:saml-provider/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_iam_service_linked_role":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:iam:%s:%s:service-linked-role/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_iam_virtual_mfa_device":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:iam:%s:%s:virtual-mfa-device/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_imagebuilder_component":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:imagebuilder:%s:%s:component/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_imagebuilder_container_recipe":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:imagebuilder:%s:%s:container-recipe/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_imagebuilder_distribution_configuration":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:imagebuilder:%s:%s:distribution-configuration/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_imagebuilder_image":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:imagebuilder:%s:%s:image/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_imagebuilder_image_pipeline":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:imagebuilder:%s:%s:image-pipeline/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_imagebuilder_image_recipe":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:imagebuilder:%s:%s:image-recipe/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_imagebuilder_infrastructure_configuration":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:imagebuilder:%s:%s:infrastructure-configuration/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_imagebuilder_lifecycle_policy":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:imagebuilder:%s:%s:lifecycle-policy/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_imagebuilder_workflow":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:imagebuilder:%s:%s:workflow/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_inspector2_filter":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:inspector2:%s:%s:filter/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_inspector_assessment_target":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:inspector:%s:%s:assessment-target/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_inspector_assessment_template":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:inspector:%s:%s:assessment-template/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_invoicing_invoice_unit":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:invoicing:%s:%s:invoice-unit/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_iot_topic_rule_destination":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:iot:%s:%s:topic-rule-destination/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_ivs_channel":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:ivs:%s:%s:channel/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_ivs_playback_key_pair":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:ivs:%s:%s:playback-key-pair/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_ivs_recording_configuration":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:ivs:%s:%s:recording-configuration/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_ivschat_logging_configuration":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:ivschat:%s:%s:logging-configuration/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_ivschat_room":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:ivschat:%s:%s:room/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_kinesis_analytics_application":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:kinesis:%s:%s:analytics-application/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_kinesis_firehose_delivery_stream":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:kinesis:%s:%s:firehose-delivery-stream/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_kinesis_stream_consumer":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:kinesis:%s:%s:stream-consumer/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_kinesis_video_stream":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:kinesis:%s:%s:video-stream/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_kinesisanalyticsv2_application":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:kinesisanalyticsv2:%s:%s:application/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_lambda_layer_version":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:lambda:%s:%s:layer-version/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_lb":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:lb:%s:%s:aws-lb/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_lb_listener":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:lb:%s:%s:listener/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_lb_listener_rule":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:lb:%s:%s:listener-rule/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_lb_target_group":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:lb:%s:%s:target-group/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_lb_trust_store":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:lb:%s:%s:trust-store/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_lb_trust_store_revocation":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:lb:%s:%s:trust-store-revocation/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_licensemanager_grant":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:licensemanager:%s:%s:grant/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_licensemanager_grant_accepter":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:licensemanager:%s:%s:grant-accepter/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_msk_cluster":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:msk:%s:%s:cluster/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_msk_configuration":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:msk:%s:%s:configuration/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_msk_replicator":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:msk:%s:%s:replicator/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_msk_serverless_cluster":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:msk:%s:%s:serverless-cluster/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_msk_vpc_connection":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:msk:%s:%s:vpc-connection/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_mskconnect_connector":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:mskconnect:%s:%s:connector/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_mskconnect_custom_plugin":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:mskconnect:%s:%s:custom-plugin/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_mskconnect_worker_configuration":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:mskconnect:%s:%s:worker-configuration/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_networkfirewall_firewall":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:networkfirewall:%s:%s:firewall/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_networkfirewall_firewall_policy":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:networkfirewall:%s:%s:firewall-policy/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_networkfirewall_rule_group":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:networkfirewall:%s:%s:rule-group/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_networkfirewall_tls_inspection_configuration":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:networkfirewall:%s:%s:tls-inspection-configuration/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_networkmanager_connection":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:networkmanager:%s:%s:connection/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_networkmanager_customer_gateway_association":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:networkmanager:%s:%s:customer-gateway-association/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_networkmanager_device":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:networkmanager:%s:%s:device/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_networkmanager_link":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:networkmanager:%s:%s:link/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_networkmanager_prefix_list_association":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:networkmanager:%s:%s:prefix-list-association/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_networkmanager_site":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:networkmanager:%s:%s:site/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_networkmanager_transit_gateway_connect_peer_association":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:networkmanager:%s:%s:transit-gateway-connect-peer-association/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_networkmanager_transit_gateway_registration":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:networkmanager:%s:%s:transit-gateway-registration/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_notifications_event_rule":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:notifications:%s:%s:event-rule/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_notifications_notification_configuration":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:notifications:%s:%s:notification-configuration/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_notificationscontacts_email_contact":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:notificationscontacts:%s:%s:email-contact/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_oam_link":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:oam:%s:%s:link/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_oam_sink":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:oam:%s:%s:sink/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_observabilityadmin_telemetry_pipeline":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:observabilityadmin:%s:%s:telemetry-pipeline/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_paymentcryptography_key":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:paymentcryptography:%s:%s:key/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_prometheus_rule_group_namespace":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:prometheus:%s:%s:rule-group-namespace/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_ram_resource_association":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:ram:%s:%s:resource-association/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_ram_resource_share":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:ram:%s:%s:resource-share/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_ram_resource_share_accepter":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:ram:%s:%s:resource-share-accepter/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_rds_cluster_role_association":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:rds:%s:%s:cluster-role-association/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_rds_integration":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:rds:%s:%s:integration/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_redshift_integration":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:redshift:%s:%s:integration/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_resiliencehub_resiliency_policy":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:resiliencehub:%s:%s:resiliency-policy/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_resourceexplorer2_index":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:resourceexplorer2:%s:%s:index/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_resourceexplorer2_view":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:resourceexplorer2:%s:%s:view/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_route53recoverycontrolconfig_cluster":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:route53recoverycontrolconfig:%s:%s:cluster/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_route53recoverycontrolconfig_control_panel":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:route53recoverycontrolconfig:%s:%s:control-panel/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_route53recoverycontrolconfig_routing_control":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:route53recoverycontrolconfig:%s:%s:routing-control/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_route53recoverycontrolconfig_safety_rule":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:route53recoverycontrolconfig:%s:%s:safety-rule/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_s3control_bucket":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:s3control:%s:%s:bucket/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_s3control_bucket_lifecycle_configuration":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:s3control:%s:%s:bucket-lifecycle-configuration/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_s3control_bucket_policy":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:s3control:%s:%s:bucket-policy/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_s3outposts_endpoint":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:s3outposts:%s:%s:endpoint/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_s3tables_table_bucket":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:s3tables:%s:%s:table-bucket/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_sagemaker_mlflow_app":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:sagemaker:%s:%s:mlflow-app/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_sagemaker_user_profile":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:sagemaker:%s:%s:user-profile/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_secretsmanager_secret":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:secretsmanager:%s:%s:secret:%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_secretsmanager_secret_policy":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:secretsmanager:%s:%s:secret:%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_secretsmanager_secret_rotation":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:secretsmanager:%s:%s:secret:%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_securityhub_account_v2":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:securityhub:%s:%s:account-v2/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_securityhub_action_target":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:securityhub:%s:%s:action-target/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_securityhub_automation_rule":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:securityhub:%s:%s:automation-rule/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_securityhub_finding_aggregator":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:securityhub:%s:%s:finding-aggregator/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_securityhub_insight":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:securityhub:%s:%s:insight/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_securityhub_standards_subscription":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:securityhub:%s:%s:standards-subscription/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_securitylake_data_lake":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:securitylake:%s:%s:data-lake/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_sfn_activity":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:sfn:%s:%s:activity/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_sfn_alias":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:sfn:%s:%s:alias/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_sfn_state_machine":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:sfn:%s:%s:state-machine/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_sns_platform_application":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:sns:%s:%s:%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_sns_topic":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:sns:%s:%s:%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_sns_topic_data_protection_policy":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:sns:%s:%s:%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_sns_topic_policy":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:sns:%s:%s:%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_sns_topic_subscription":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:sns:%s:%s:%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_ssmcontacts_contact":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:ssmcontacts:%s:%s:contact/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_ssmcontacts_contact_channel":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:ssmcontacts:%s:%s:contact-channel/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_ssmcontacts_plan":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:ssmcontacts:%s:%s:plan/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_ssmcontacts_rotation":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:ssmcontacts:%s:%s:rotation/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_ssmincidents_response_plan":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:ssmincidents:%s:%s:response-plan/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_ssoadmin_permission_set":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:ssoadmin:%s:%s:permission-set/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_storagegateway_cache":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:storagegateway:%s:%s:cache/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_storagegateway_cached_iscsi_volume":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:storagegateway:%s:%s:cached-iscsi-volume/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_storagegateway_file_system_association":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:storagegateway:%s:%s:file-system-association/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_storagegateway_gateway":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:storagegateway:%s:%s:gateway/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_storagegateway_nfs_file_share":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:storagegateway:%s:%s:nfs-file-share/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_storagegateway_smb_file_share":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:storagegateway:%s:%s:smb-file-share/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_storagegateway_stored_iscsi_volume":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:storagegateway:%s:%s:stored-iscsi-volume/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_storagegateway_tape_pool":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:storagegateway:%s:%s:tape-pool/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_storagegateway_upload_buffer":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:storagegateway:%s:%s:upload-buffer/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_storagegateway_working_storage":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:storagegateway:%s:%s:working-storage/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_timestreamquery_scheduled_query":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:timestreamquery:%s:%s:scheduled-query/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_wafv2_web_acl_logging_configuration":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:wafv2:%s:%s:web-acl-logging-configuration/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_xray_group":
		name := resolveAttribute(ctx, config, "name")
		if name != "" && ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.AccountID != "" && awsClient.Region != "" {
				return fmt.Sprintf("arn:%s:xray:%s:%s:group/%s", awsClient.Partition, awsClient.Region, awsClient.AccountID, name)
			}
		}
		return ""
	case "aws_vpc":
		if ctx != nil {
			awsClient := ctx.GetAWSClient()
			if awsClient != nil && awsClient.EC2Client != nil {
				return lookupVpcID(ctx, config)
			}
		}
	}
	return ""
}

// resolveAttribute is an implementation of the Expression Tracer Pattern.
// When dealing with composite IDs (e.g. role_name/policy_arn), one of the parts might
// be "computed" (unknown) if the referenced resource is being created in the same plan.
// In this scenario, the value is omitted from the Change.After map.
// To bypass this, this function traces the terraform configuration Expressions to find
// the referenced resource (e.g., aws_iam_policy.my_policy) and recursively calls
// GetImportID on it to hydrate the missing attribute.
func resolveAttribute(ctx *ProviderContext, config map[string]any, attr string) string {
	if val, ok := config[attr].(string); ok && val != "" {
		return val
	}

	if ctx == nil || ctx.Plan == nil || ctx.CurrentResource == nil {
		return ""
	}

	// The CurrentResource Address includes indices (e.g. aws_iam_role_policy_attachment.att["sandbox"])
	// but the ConfigResource Address does not. We need to strip all indices.
	re := regexp.MustCompile(`\[[^\]]+\]`)
	baseAddress := re.ReplaceAllString(ctx.CurrentResource.Address, "")

	cfgRes := findConfigResource(ctx.Plan.Config.RootModule, baseAddress)
	if cfgRes == nil {
		return ""
	}

	expr, ok := cfgRes.Expressions[attr]
	if !ok || expr == nil || expr.ExpressionData == nil {
		return ""
	}

	if expr.ConstantValue != nil {
		if s, ok := expr.ConstantValue.(string); ok {
			return s
		}
	}

	// Try to find referenced resource
	for _, ref := range expr.References {
		// Filter out terraform keywords
		if ref == "each.key" || ref == "each.value" || ref == "count.index" || strings.HasPrefix(ref, "var.") || strings.HasPrefix(ref, "local.") {
			continue
		}

		// example 1: aws_iam_policy.policy.arn
		// example 2: aws_iam_policy.policy
		parts := strings.Split(ref, ".")
		if len(parts) >= 2 {
			var refAddress string
			var expectedAttr string

			// Does it end with an attribute we recognize?
			lastPart := parts[len(parts)-1]
			if lastPart == "arn" || lastPart == "id" || lastPart == "name" {
				refAddress = strings.Join(parts[:len(parts)-1], ".")
				expectedAttr = lastPart
			} else {
				refAddress = ref
				// Default to standard GetImportID hydration logic which usually returns ARN/ID/Name
				expectedAttr = ""
			}

			// We only support resolving attributes from other resources created in same plan
			for _, rc := range ctx.Plan.ResourceChanges {
				// We need to match the base address.
				// e.g. rc.Address is aws_iam_policy.policy["sandbox"] and refAddress is aws_iam_policy.policy
				rcBaseAddress := re.ReplaceAllString(rc.Address, "")

				if rcBaseAddress == refAddress {
					// Check if this resource shares the same index, or if it doesn't have an index
					if rc.Index != nil && ctx.CurrentResource.Index != nil {
						if rc.Index != ctx.CurrentResource.Index {
							continue
						}
					}

					after, ok := rc.Change.After.(map[string]any)
					if ok {
						if expectedAttr != "" {
							if attrVal, ok := after[expectedAttr].(string); ok && attrVal != "" {
								return attrVal
							}
						}

						// Compute the ID using our normal hydration logic
						id := GetImportID(ctx, rc.Type, after)
						if id != "" && id != MessageProviderNotSupported {
							return id
						}
					}
				}
			}
		}
	}

	return ""
}

func findConfigResource(mod *tfjson.ConfigModule, address string) *tfjson.ConfigResource {
	if mod == nil {
		return nil
	}
	for _, res := range mod.Resources {
		if res.Address == address {
			return res
		}
	}
	for _, call := range mod.ModuleCalls {
		if found := findConfigResource(call.Module, address); found != nil {
			return found
		}
	}
	return nil
}

// lookupVpcID attempts to find an existing VPC in AWS using the Name tag.
func lookupVpcID(ctx *ProviderContext, config map[string]any) string {
	tags, ok := config["tags"].(map[string]any)
	if !ok {
		return ""
	}
	nameTag, ok := tags["Name"].(string)
	if !ok || nameTag == "" {
		return ""
	}

	res, err := ctx.GetAWSClient().EC2Client.DescribeVpcs(ctx.Context, &ec2.DescribeVpcsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("tag:Name"),
				Values: []string{nameTag},
			},
		},
	})
	if err != nil {
		log.Printf("Warning: failed to describe VPCs for name %q: %v\n", nameTag, err)
		return ""
	}

	if len(res.Vpcs) == 1 {
		return aws.ToString(res.Vpcs[0].VpcId)
	} else if len(res.Vpcs) > 1 {
		log.Printf("Warning: found multiple VPCs with Name=%q. Cannot unambiguously resolve ID.\n", nameTag)
		return ""
	}

	return ""
}

// lookupSqsQueueUrl attempts to find an existing SQS Queue URL.
func lookupSqsQueueUrl(ctx *ProviderContext, config map[string]any) string {
	name := resolveAttribute(ctx, config, "name")

	if name != "" && ctx != nil {
		awsClient := ctx.GetAWSClient()
		if awsClient != nil && awsClient.Region != "" && awsClient.AccountID != "" {
			return fmt.Sprintf("https://sqs.%s.amazonaws.com/%s/%s", awsClient.Region, awsClient.AccountID, name)
		}
	}

	if name == "" {
		namePrefix := resolveAttribute(ctx, config, "name_prefix")
		if namePrefix != "" {
			res, err := ctx.GetAWSClient().SQSClient.ListQueues(ctx.Context, &sqs.ListQueuesInput{
				QueueNamePrefix: aws.String(namePrefix),
			})
			if err == nil {
				if len(res.QueueUrls) == 1 {
					return res.QueueUrls[0]
				} else if len(res.QueueUrls) > 1 {
					log.Printf("Warning: found multiple SQS queues with prefix %q. Cannot unambiguously resolve ID.\n", namePrefix)
				}
			} else {
				log.Printf("Warning: failed to list SQS queues for prefix %q: %v\n", namePrefix, err)
			}
		}
		return ""
	}

	res, err := ctx.GetAWSClient().SQSClient.GetQueueUrl(ctx.Context, &sqs.GetQueueUrlInput{
		QueueName: aws.String(name),
	})
	if err == nil && res.QueueUrl != nil {
		return *res.QueueUrl
	} else if err != nil {
		log.Printf("Warning: failed to get SQS queue URL for name %q: %v\n", name, err)
	}

	return ""
}
