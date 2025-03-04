// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_verified_access_group", verifiedAccessGroupDataSource)
}

// verifiedAccessGroupDataSource returns the Terraform awscc_ec2_verified_access_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::VerifiedAccessGroup resource.
func verifiedAccessGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Time this Verified Access Group was created.",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Time this Verified Access Group was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description for the AWS Verified Access group.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description for the AWS Verified Access group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastUpdatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Time this Verified Access Group was last updated.",
		//	  "type": "string"
		//	}
		"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Time this Verified Access Group was last updated.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Owner
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS account number that owns the group.",
		//	  "type": "string"
		//	}
		"owner": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS account number that owns the group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyDocument
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS Verified Access policy document.",
		//	  "type": "string"
		//	}
		"policy_document": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS Verified Access policy document.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the Verified Access policy.",
		//	  "type": "boolean"
		//	}
		"policy_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the Verified Access policy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SseSpecification
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration options for customer provided KMS encryption.",
		//	  "properties": {
		//	    "CustomerManagedKeyEnabled": {
		//	      "description": "Whether to encrypt the policy with the provided key or disable encryption",
		//	      "type": "boolean"
		//	    },
		//	    "KmsKeyArn": {
		//	      "description": "KMS Key Arn used to encrypt the group policy",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"sse_specification": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CustomerManagedKeyEnabled
				"customer_managed_key_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Whether to encrypt the policy with the provided key or disable encryption",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: KmsKeyArn
				"kms_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "KMS Key Arn used to encrypt the group policy",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration options for customer provided KMS encryption.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VerifiedAccessGroupArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the Verified Access group.",
		//	  "type": "string"
		//	}
		"verified_access_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the Verified Access group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VerifiedAccessGroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the AWS Verified Access group.",
		//	  "type": "string"
		//	}
		"verified_access_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the AWS Verified Access group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VerifiedAccessInstanceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the AWS Verified Access instance.",
		//	  "type": "string"
		//	}
		"verified_access_instance_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the AWS Verified Access instance.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::VerifiedAccessGroup",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::VerifiedAccessGroup").WithTerraformTypeName("awscc_ec2_verified_access_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"creation_time":                "CreationTime",
		"customer_managed_key_enabled": "CustomerManagedKeyEnabled",
		"description":                  "Description",
		"key":                          "Key",
		"kms_key_arn":                  "KmsKeyArn",
		"last_updated_time":            "LastUpdatedTime",
		"owner":                        "Owner",
		"policy_document":              "PolicyDocument",
		"policy_enabled":               "PolicyEnabled",
		"sse_specification":            "SseSpecification",
		"tags":                         "Tags",
		"value":                        "Value",
		"verified_access_group_arn":    "VerifiedAccessGroupArn",
		"verified_access_group_id":     "VerifiedAccessGroupId",
		"verified_access_instance_id":  "VerifiedAccessInstanceId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
