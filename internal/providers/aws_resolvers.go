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
