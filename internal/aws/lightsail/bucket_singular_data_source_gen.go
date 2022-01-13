// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_lightsail_bucket", bucketDataSourceType)
}

// bucketDataSourceType returns the Terraform awscc_lightsail_bucket data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Lightsail::Bucket resource type.
func bucketDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"able_to_update_bundle": {
			// Property: AbleToUpdateBundle
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether the bundle that is currently applied to a bucket can be changed to another bundle. You can update a bucket's bundle only one time within a monthly AWS billing cycle.",
			//   "type": "boolean"
			// }
			Description: "Indicates whether the bundle that is currently applied to a bucket can be changed to another bundle. You can update a bucket's bundle only one time within a monthly AWS billing cycle.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"access_rules": {
			// Property: AccessRules
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "An object that sets the public accessibility of objects in the specified bucket.",
			//   "properties": {
			//     "AllowPublicOverrides": {
			//       "description": "A Boolean value that indicates whether the access control list (ACL) permissions that are applied to individual objects override the getObject option that is currently specified.",
			//       "type": "boolean"
			//     },
			//     "GetObject": {
			//       "description": "Specifies the anonymous access to all objects in a bucket.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "An object that sets the public accessibility of objects in the specified bucket.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"allow_public_overrides": {
						// Property: AllowPublicOverrides
						Description: "A Boolean value that indicates whether the access control list (ACL) permissions that are applied to individual objects override the getObject option that is currently specified.",
						Type:        types.BoolType,
						Computed:    true,
					},
					"get_object": {
						// Property: GetObject
						Description: "Specifies the anonymous access to all objects in a bucket.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"bucket_arn": {
			// Property: BucketArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"bucket_name": {
			// Property: BucketName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name for the bucket.",
			//   "maxLength": 54,
			//   "minLength": 3,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name for the bucket.",
			Type:        types.StringType,
			Computed:    true,
		},
		"bundle_id": {
			// Property: BundleId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the bundle to use for the bucket.",
			//   "type": "string"
			// }
			Description: "The ID of the bundle to use for the bucket.",
			Type:        types.StringType,
			Computed:    true,
		},
		"object_versioning": {
			// Property: ObjectVersioning
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies whether to enable or disable versioning of objects in the bucket.",
			//   "type": "boolean"
			// }
			Description: "Specifies whether to enable or disable versioning of objects in the bucket.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"read_only_access_accounts": {
			// Property: ReadOnlyAccessAccounts
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of strings to specify the AWS account IDs that can access the bucket.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of strings to specify the AWS account IDs that can access the bucket.",
			Type:        types.SetType{ElemType: types.StringType},
			Computed:    true,
		},
		"resources_receiving_access": {
			// Property: ResourcesReceivingAccess
			// CloudFormation resource type schema:
			// {
			//   "description": "The names of the Lightsail resources for which to set bucket access.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The names of the Lightsail resources for which to set bucket access.",
			Type:        types.SetType{ElemType: types.StringType},
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
		"url": {
			// Property: Url
			// CloudFormation resource type schema:
			// {
			//   "description": "The URL of the bucket.",
			//   "type": "string"
			// }
			Description: "The URL of the bucket.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Lightsail::Bucket",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lightsail::Bucket").WithTerraformTypeName("awscc_lightsail_bucket")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"able_to_update_bundle":      "AbleToUpdateBundle",
		"access_rules":               "AccessRules",
		"allow_public_overrides":     "AllowPublicOverrides",
		"bucket_arn":                 "BucketArn",
		"bucket_name":                "BucketName",
		"bundle_id":                  "BundleId",
		"get_object":                 "GetObject",
		"key":                        "Key",
		"object_versioning":          "ObjectVersioning",
		"read_only_access_accounts":  "ReadOnlyAccessAccounts",
		"resources_receiving_access": "ResourcesReceivingAccess",
		"tags":                       "Tags",
		"url":                        "Url",
		"value":                      "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
