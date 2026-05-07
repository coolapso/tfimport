package providers

import (
	"context"
	"strings"
	"sync"

	tfjson "github.com/hashicorp/terraform-json"
)

type ProviderContext struct {
	Context context.Context

	awsClient *AWSClientContext
	awsOnce   sync.Once

	Plan            *tfjson.Plan
	CurrentResource *tfjson.ResourceChange
}

func NewProviderContext(ctx context.Context, plan *tfjson.Plan) *ProviderContext {
	return &ProviderContext{
		Context: ctx,
		Plan:    plan,
	}
}

var (
	MessageProviderNotSupported = "provider not supported"
)

// GetImportID returns the necessary import ID for a given resource based on its configuration
// extracted from the terraform plan.
func GetImportID(ctx *ProviderContext, resourceType string, config map[string]any) string {
	provider := strings.Split(resourceType, "_")[0]
	switch provider {
	case "aws":
		return extractAWSImportID(ctx, resourceType, config)
	case "kubernetes":
		return extractKubernetesImportID(ctx, resourceType, config)
	case "argocd":
		return extractArgocdImportID(ctx, resourceType, config)
	case "google":
		return extractGoogleImportID(ctx, resourceType, config)
	case "azurerm":
		return extractAzurermImportID(ctx, resourceType, config)
	case "scaleway":
		return extractScalewayImportID(ctx, resourceType, config)
	default:
		return MessageProviderNotSupported
	}
}
