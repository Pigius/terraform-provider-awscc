// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_iotfleetwise_campaign", campaignResource)
}

// campaignResource returns the Terraform awscc_iotfleetwise_campaign resource.
// This Terraform resource corresponds to the CloudFormation AWS::IoTFleetWise::Campaign resource.
func campaignResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Action
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "APPROVE",
		//	    "SUSPEND",
		//	    "RESUME",
		//	    "UPDATE"
		//	  ],
		//	  "type": "string"
		//	}
		"action": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"APPROVE",
					"SUSPEND",
					"RESUME",
					"UPDATE",
				),
			}, /*END VALIDATORS*/
			// Action is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CollectionScheme
		// CloudFormation resource type schema:
		//
		//	{
		//	  "properties": {
		//	    "ConditionBasedCollectionScheme": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "ConditionLanguageVersion": {
		//	          "minimum": 1,
		//	          "type": "integer"
		//	        },
		//	        "Expression": {
		//	          "maxLength": 2048,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "MinimumTriggerIntervalMs": {
		//	          "maximum": 4294967295,
		//	          "minimum": 0,
		//	          "type": "number"
		//	        },
		//	        "TriggerMode": {
		//	          "enum": [
		//	            "ALWAYS",
		//	            "RISING_EDGE"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Expression"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "TimeBasedCollectionScheme": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "PeriodMs": {
		//	          "maximum": 60000,
		//	          "minimum": 10000,
		//	          "type": "number"
		//	        }
		//	      },
		//	      "required": [
		//	        "PeriodMs"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"collection_scheme": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ConditionBasedCollectionScheme
				"condition_based_collection_scheme": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ConditionLanguageVersion
						"condition_language_version": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.Int64{ /*START VALIDATORS*/
								int64validator.AtLeast(1),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Expression
						"expression": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 2048),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: MinimumTriggerIntervalMs
						"minimum_trigger_interval_ms": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.Float64{ /*START VALIDATORS*/
								float64validator.Between(0.000000, 4294967295.000000),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
								float64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: TriggerMode
						"trigger_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"ALWAYS",
									"RISING_EDGE",
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TimeBasedCollectionScheme
				"time_based_collection_scheme": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: PeriodMs
						"period_ms": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.Float64{ /*START VALIDATORS*/
								float64validator.Between(10000.000000, 60000.000000),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Compression
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "OFF",
		//	  "enum": [
		//	    "OFF",
		//	    "SNAPPY"
		//	  ],
		//	  "type": "string"
		//	}
		"compression": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"OFF",
					"SNAPPY",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				generic.StringDefaultValue("OFF"),
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DataDestinationConfigs
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "properties": {
		//	      "S3Config": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "BucketArn": {
		//	            "maxLength": 100,
		//	            "minLength": 16,
		//	            "pattern": "^arn:(aws[a-zA-Z0-9-]*):s3:::.+$",
		//	            "type": "string"
		//	          },
		//	          "DataFormat": {
		//	            "enum": [
		//	              "JSON",
		//	              "PARQUET"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "Prefix": {
		//	            "maxLength": 512,
		//	            "minLength": 1,
		//	            "pattern": "^[a-zA-Z0-9-_:./!*'()]+$",
		//	            "type": "string"
		//	          },
		//	          "StorageCompressionFormat": {
		//	            "enum": [
		//	              "NONE",
		//	              "GZIP"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "BucketArn"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "TimestreamConfig": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "ExecutionRoleArn": {
		//	            "maxLength": 2048,
		//	            "minLength": 20,
		//	            "pattern": "",
		//	            "type": "string"
		//	          },
		//	          "TimestreamTableArn": {
		//	            "maxLength": 2048,
		//	            "minLength": 20,
		//	            "pattern": "^arn:(aws[a-zA-Z0-9-]*):timestream:[a-zA-Z0-9-]+:[0-9]{12}:database\\/[a-zA-Z0-9_.-]+\\/table\\/[a-zA-Z0-9_.-]+$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "TimestreamTableArn",
		//	          "ExecutionRoleArn"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 1,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"data_destination_configs": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: S3Config
					"s3_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: BucketArn
							"bucket_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(16, 100),
									stringvalidator.RegexMatches(regexp.MustCompile("^arn:(aws[a-zA-Z0-9-]*):s3:::.+$"), ""),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: DataFormat
							"data_format": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"JSON",
										"PARQUET",
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Prefix
							"prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 512),
									stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9-_:./!*'()]+$"), ""),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: StorageCompressionFormat
							"storage_compression_format": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"NONE",
										"GZIP",
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: TimestreamConfig
					"timestream_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ExecutionRoleArn
							"execution_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(20, 2048),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: TimestreamTableArn
							"timestream_table_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(20, 2048),
									stringvalidator.RegexMatches(regexp.MustCompile("^arn:(aws[a-zA-Z0-9-]*):timestream:[a-zA-Z0-9-]+:[0-9]{12}:database\\/[a-zA-Z0-9_.-]+\\/table\\/[a-zA-Z0-9_.-]+$"), ""),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 1),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DataExtraDimensions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 150,
		//	    "minLength": 1,
		//	    "pattern": "^[a-zA-Z0-9_.]+$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 5,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"data_extra_dimensions": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 5),
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(1, 150),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_.]+$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 2048),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DiagnosticsMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "OFF",
		//	  "enum": [
		//	    "OFF",
		//	    "SEND_ACTIVE_DTCS"
		//	  ],
		//	  "type": "string"
		//	}
		"diagnostics_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"OFF",
					"SEND_ACTIVE_DTCS",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				generic.StringDefaultValue("OFF"),
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ExpiryTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "253402214400",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"expiry_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				validate.IsRFC3339Time(),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				generic.StringDefaultValue("253402214400"),
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LastModificationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"last_modification_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z\\d\\-_:]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 100),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_:]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PostTriggerCollectionDuration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": 0,
		//	  "maximum": 4294967295,
		//	  "minimum": 0,
		//	  "type": "number"
		//	}
		"post_trigger_collection_duration": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.Float64{ /*START VALIDATORS*/
				float64validator.Between(0.000000, 4294967295.000000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
				generic.Float64DefaultValue(0.000000),
				float64planmodifier.UseStateForUnknown(),
				float64planmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Priority
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": 0,
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"priority": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.AtLeast(0),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				generic.Int64DefaultValue(0),
				int64planmodifier.UseStateForUnknown(),
				int64planmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SignalCatalogArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"signal_catalog_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SignalsToCollect
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "MaxSampleCount": {
		//	        "maximum": 4294967295,
		//	        "minimum": 1,
		//	        "type": "number"
		//	      },
		//	      "MinimumSamplingIntervalMs": {
		//	        "maximum": 4294967295,
		//	        "minimum": 0,
		//	        "type": "number"
		//	      },
		//	      "Name": {
		//	        "maxLength": 150,
		//	        "minLength": 1,
		//	        "pattern": "^[\\w|*|-]+(\\.[\\w|*|-]+)*$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Name"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 1000,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"signals_to_collect": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: MaxSampleCount
					"max_sample_count": schema.Float64Attribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.Float64{ /*START VALIDATORS*/
							float64validator.Between(1.000000, 4294967295.000000),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
							float64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: MinimumSamplingIntervalMs
					"minimum_sampling_interval_ms": schema.Float64Attribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.Float64{ /*START VALIDATORS*/
							float64validator.Between(0.000000, 4294967295.000000),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
							float64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 150),
							stringvalidator.RegexMatches(regexp.MustCompile("^[\\w|*|-]+(\\.[\\w|*|-]+)*$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 1000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SpoolingMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "OFF",
		//	  "enum": [
		//	    "OFF",
		//	    "TO_DISK"
		//	  ],
		//	  "type": "string"
		//	}
		"spooling_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"OFF",
					"TO_DISK",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				generic.StringDefaultValue("OFF"),
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StartTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "0",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"start_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				validate.IsRFC3339Time(),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				generic.StringDefaultValue("0"),
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "CREATING",
		//	    "WAITING_FOR_APPROVAL",
		//	    "RUNNING",
		//	    "SUSPENDED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
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
		//	  "maxItems": 50,
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeBetween(0, 50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TargetArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"target_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Definition of AWS::IoTFleetWise::Campaign Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTFleetWise::Campaign").WithTerraformTypeName("awscc_iotfleetwise_campaign")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                            "Action",
		"arn":                               "Arn",
		"bucket_arn":                        "BucketArn",
		"collection_scheme":                 "CollectionScheme",
		"compression":                       "Compression",
		"condition_based_collection_scheme": "ConditionBasedCollectionScheme",
		"condition_language_version":        "ConditionLanguageVersion",
		"creation_time":                     "CreationTime",
		"data_destination_configs":          "DataDestinationConfigs",
		"data_extra_dimensions":             "DataExtraDimensions",
		"data_format":                       "DataFormat",
		"description":                       "Description",
		"diagnostics_mode":                  "DiagnosticsMode",
		"execution_role_arn":                "ExecutionRoleArn",
		"expiry_time":                       "ExpiryTime",
		"expression":                        "Expression",
		"key":                               "Key",
		"last_modification_time":            "LastModificationTime",
		"max_sample_count":                  "MaxSampleCount",
		"minimum_sampling_interval_ms":      "MinimumSamplingIntervalMs",
		"minimum_trigger_interval_ms":       "MinimumTriggerIntervalMs",
		"name":                              "Name",
		"period_ms":                         "PeriodMs",
		"post_trigger_collection_duration":  "PostTriggerCollectionDuration",
		"prefix":                            "Prefix",
		"priority":                          "Priority",
		"s3_config":                         "S3Config",
		"signal_catalog_arn":                "SignalCatalogArn",
		"signals_to_collect":                "SignalsToCollect",
		"spooling_mode":                     "SpoolingMode",
		"start_time":                        "StartTime",
		"status":                            "Status",
		"storage_compression_format":        "StorageCompressionFormat",
		"tags":                              "Tags",
		"target_arn":                        "TargetArn",
		"time_based_collection_scheme":      "TimeBasedCollectionScheme",
		"timestream_config":                 "TimestreamConfig",
		"timestream_table_arn":              "TimestreamTableArn",
		"trigger_mode":                      "TriggerMode",
		"value":                             "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Action",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
