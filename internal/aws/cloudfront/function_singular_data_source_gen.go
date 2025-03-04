// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudfront_function", functionDataSource)
}

// functionDataSource returns the Terraform awscc_cloudfront_function data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudFront::Function resource.
func functionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AutoPublish
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"auto_publish": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: FunctionARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"function_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: FunctionCode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"function_code": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: FunctionConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Comment": {
		//	      "type": "string"
		//	    },
		//	    "Runtime": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Comment",
		//	    "Runtime"
		//	  ],
		//	  "type": "object"
		//	}
		"function_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Comment
				"comment": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Runtime
				"runtime": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: FunctionMetadata
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "FunctionARN": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"function_metadata": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: FunctionARN
				"function_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Stage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"stage": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CloudFront::Function",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::Function").WithTerraformTypeName("awscc_cloudfront_function")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auto_publish":      "AutoPublish",
		"comment":           "Comment",
		"function_arn":      "FunctionARN",
		"function_code":     "FunctionCode",
		"function_config":   "FunctionConfig",
		"function_metadata": "FunctionMetadata",
		"name":              "Name",
		"runtime":           "Runtime",
		"stage":             "Stage",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
