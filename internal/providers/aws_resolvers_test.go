package providers

import (
	"context"
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
)

// newS3Plan builds a plan mimicking the terraform-aws-modules/s3-bucket layout where
// every sub-resource references the bucket created in the same plan through
// `aws_s3_bucket.this[0].id`. Because that value is computed, terraform omits the
// `bucket` attribute from the sub-resource's planned values entirely.
func newS3Plan(subType string, after map[string]any) (*ProviderContext, map[string]any) {
	bucketRefs := []string{"aws_s3_bucket.this[0].id", "aws_s3_bucket.this[0]", "aws_s3_bucket.this"}
	plan := &tfjson.Plan{
		ResourceChanges: []*tfjson.ResourceChange{
			{
				Address: "aws_s3_bucket.this[0]",
				Type:    "aws_s3_bucket",
				Name:    "this",
				Index:   float64(0),
				Change:  &tfjson.Change{After: map[string]any{"bucket": "my-bucket"}},
			},
			{
				Address: subType + ".this[0]",
				Type:    subType,
				Name:    "this",
				Index:   float64(0),
				Change:  &tfjson.Change{After: after},
			},
		},
		Config: &tfjson.Config{
			RootModule: &tfjson.ConfigModule{
				Resources: []*tfjson.ConfigResource{
					{
						Address: "aws_s3_bucket.this",
						Type:    "aws_s3_bucket",
						Name:    "this",
						Expressions: map[string]*tfjson.Expression{
							"bucket": {ExpressionData: &tfjson.ExpressionData{ConstantValue: "my-bucket"}},
						},
					},
					{
						Address: subType + ".this",
						Type:    subType,
						Name:    "this",
						Expressions: map[string]*tfjson.Expression{
							"bucket": {ExpressionData: &tfjson.ExpressionData{References: bucketRefs}},
						},
					},
				},
			},
		},
	}

	ctx := NewProviderContext(context.Background(), plan)
	ctx.CurrentResource = plan.ResourceChanges[1]
	return ctx, after
}

func TestS3BucketSubResourcesWithComputedBucket(t *testing.T) {
	tests := []struct {
		resourceType string
		after        map[string]any
		want         string
	}{
		{"aws_s3_bucket_ownership_controls", map[string]any{}, "my-bucket"},
		{"aws_s3_bucket_public_access_block", map[string]any{"block_public_acls": true}, "my-bucket"},
		{"aws_s3_bucket_server_side_encryption_configuration", map[string]any{}, "my-bucket"},
		{"aws_s3_bucket_server_side_encryption_configuration", map[string]any{"expected_bucket_owner": "123456789012"}, "my-bucket,123456789012"},
		{"aws_s3_bucket_policy", map[string]any{"policy": "{}"}, "my-bucket"},
		{"aws_s3_bucket_versioning", map[string]any{}, "my-bucket"},
		{"aws_s3_bucket_acl", map[string]any{"acl": "private"}, "my-bucket,private"},
		{"aws_s3_bucket_acl", map[string]any{"expected_bucket_owner": "123456789012", "acl": "private"}, "my-bucket,123456789012,private"},
	}

	for _, tt := range tests {
		t.Run(tt.resourceType+"/"+tt.want, func(t *testing.T) {
			ctx, config := newS3Plan(tt.resourceType, tt.after)
			if got := GetImportID(ctx, tt.resourceType, config); got != tt.want {
				t.Fatalf("GetImportID(%s) = %q, want %q", tt.resourceType, got, tt.want)
			}
		})
	}
}

func TestS3BucketSubResourcesWithLiteralBucket(t *testing.T) {
	ctx := NewProviderContext(context.Background(), &tfjson.Plan{})
	config := map[string]any{"bucket": "literal-bucket"}
	for _, rt := range []string{
		"aws_s3_bucket_ownership_controls",
		"aws_s3_bucket_public_access_block",
		"aws_s3_bucket_server_side_encryption_configuration",
	} {
		if got := GetImportID(ctx, rt, config); got != "literal-bucket" {
			t.Fatalf("GetImportID(%s) = %q, want %q", rt, got, "literal-bucket")
		}
	}
}
