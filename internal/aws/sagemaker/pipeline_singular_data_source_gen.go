// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_sagemaker_pipeline", pipelineDataSourceType)
}

// pipelineDataSourceType returns the Terraform awscc_sagemaker_pipeline data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::SageMaker::Pipeline resource type.
func pipelineDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"parallelism_configuration": {
			// Property: ParallelismConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "MaxParallelExecutionSteps": {
			//       "description": "Maximum parallel execution steps",
			//       "minimum": 1,
			//       "type": "integer"
			//     }
			//   },
			//   "required": [
			//     "MaxParallelExecutionSteps"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"max_parallel_execution_steps": {
						// Property: MaxParallelExecutionSteps
						Description: "Maximum parallel execution steps",
						Type:        types.NumberType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"pipeline_definition": {
			// Property: PipelineDefinition
			// CloudFormation resource type schema:
			// {
			//   "properties": {
			//     "PipelineDefinitionBody": {
			//       "description": "A specification that defines the pipeline in JSON format.",
			//       "type": "string"
			//     },
			//     "PipelineDefinitionS3Location": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Bucket": {
			//           "description": "The name of the S3 bucket where the PipelineDefinition file is stored.",
			//           "type": "string"
			//         },
			//         "ETag": {
			//           "description": "The Amazon S3 ETag (a file checksum) of the PipelineDefinition file. If you don't specify a value, SageMaker skips ETag validation of your PipelineDefinition file.",
			//           "type": "string"
			//         },
			//         "Key": {
			//           "description": "The file name of the PipelineDefinition file (Amazon S3 object name).",
			//           "type": "string"
			//         },
			//         "Version": {
			//           "description": "For versioning-enabled buckets, a specific version of the PipelineDefinition file.",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "Bucket",
			//         "Key"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"pipeline_definition_body": {
						// Property: PipelineDefinitionBody
						Description: "A specification that defines the pipeline in JSON format.",
						Type:        types.StringType,
						Computed:    true,
					},
					"pipeline_definition_s3_location": {
						// Property: PipelineDefinitionS3Location
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"bucket": {
									// Property: Bucket
									Description: "The name of the S3 bucket where the PipelineDefinition file is stored.",
									Type:        types.StringType,
									Computed:    true,
								},
								"e_tag": {
									// Property: ETag
									Description: "The Amazon S3 ETag (a file checksum) of the PipelineDefinition file. If you don't specify a value, SageMaker skips ETag validation of your PipelineDefinition file.",
									Type:        types.StringType,
									Computed:    true,
								},
								"key": {
									// Property: Key
									Description: "The file name of the PipelineDefinition file (Amazon S3 object name).",
									Type:        types.StringType,
									Computed:    true,
								},
								"version": {
									// Property: Version
									Description: "For versioning-enabled buckets, a specific version of the PipelineDefinition file.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"pipeline_description": {
			// Property: PipelineDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the Pipeline.",
			//   "maxLength": 3072,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The description of the Pipeline.",
			Type:        types.StringType,
			Computed:    true,
		},
		"pipeline_display_name": {
			// Property: PipelineDisplayName
			// CloudFormation resource type schema:
			// {
			//   "description": "The display name of the Pipeline.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The display name of the Pipeline.",
			Type:        types.StringType,
			Computed:    true,
		},
		"pipeline_name": {
			// Property: PipelineName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Pipeline.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the Pipeline.",
			Type:        types.StringType,
			Computed:    true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Role Arn",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Role Arn",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::SageMaker::Pipeline",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::Pipeline").WithTerraformTypeName("awscc_sagemaker_pipeline")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bucket":                          "Bucket",
		"e_tag":                           "ETag",
		"key":                             "Key",
		"max_parallel_execution_steps":    "MaxParallelExecutionSteps",
		"parallelism_configuration":       "ParallelismConfiguration",
		"pipeline_definition":             "PipelineDefinition",
		"pipeline_definition_body":        "PipelineDefinitionBody",
		"pipeline_definition_s3_location": "PipelineDefinitionS3Location",
		"pipeline_description":            "PipelineDescription",
		"pipeline_display_name":           "PipelineDisplayName",
		"pipeline_name":                   "PipelineName",
		"role_arn":                        "RoleArn",
		"tags":                            "Tags",
		"value":                           "Value",
		"version":                         "Version",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
