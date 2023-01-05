// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package evidently

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_evidently_experiment", experimentDataSource)
}

// experimentDataSource returns the Terraform awscc_evidently_experiment data source.
// This Terraform data source corresponds to the CloudFormation AWS::Evidently::Experiment resource.
func experimentDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "arn:[^:]*:[^:]*:[^:]*:[^:]*:project/[-a-zA-Z0-9._]*/experiment/[-a-zA-Z0-9._]*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 160,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: MetricGoals
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "DesiredChange": {
		//	        "enum": [
		//	          "INCREASE",
		//	          "DECREASE"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "EntityIdKey": {
		//	        "description": "The JSON path to reference the entity id in the event.",
		//	        "type": "string"
		//	      },
		//	      "EventPattern": {
		//	        "description": "Event patterns have the same structure as the events they match. Rules use event patterns to select events. An event pattern either matches an event or it doesn't.",
		//	        "type": "string"
		//	      },
		//	      "MetricName": {
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "pattern": "^[\\S]+$",
		//	        "type": "string"
		//	      },
		//	      "UnitLabel": {
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": ".*",
		//	        "type": "string"
		//	      },
		//	      "ValueKey": {
		//	        "description": "The JSON path to reference the numerical metric value in the event.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "MetricName",
		//	      "EntityIdKey",
		//	      "ValueKey",
		//	      "DesiredChange"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 3,
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"metric_goals": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: DesiredChange
					"desired_change": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: EntityIdKey
					"entity_id_key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The JSON path to reference the entity id in the event.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: EventPattern
					"event_pattern": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Event patterns have the same structure as the events they match. Rules use event patterns to select events. An event pattern either matches an event or it doesn't.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: MetricName
					"metric_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: UnitLabel
					"unit_label": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: ValueKey
					"value_key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The JSON path to reference the numerical metric value in the event.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 127,
		//	  "minLength": 1,
		//	  "pattern": "[-a-zA-Z0-9._]*",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: OnlineAbConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "ControlTreatmentName": {
		//	      "maxLength": 127,
		//	      "minLength": 1,
		//	      "pattern": "[-a-zA-Z0-9._]*",
		//	      "type": "string"
		//	    },
		//	    "TreatmentWeights": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "SplitWeight": {
		//	            "maximum": 100000,
		//	            "minimum": 0,
		//	            "type": "integer"
		//	          },
		//	          "Treatment": {
		//	            "maxLength": 127,
		//	            "minLength": 1,
		//	            "pattern": "[-a-zA-Z0-9._]*",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Treatment",
		//	          "SplitWeight"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"online_ab_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ControlTreatmentName
				"control_treatment_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: TreatmentWeights
				"treatment_weights": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: SplitWeight
							"split_weight": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Treatment
							"treatment": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Project
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 0,
		//	  "pattern": "([-a-zA-Z0-9._]*)|(arn:[^:]*:[^:]*:[^:]*:[^:]*:project/[-a-zA-Z0-9._]*)",
		//	  "type": "string"
		//	}
		"project": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RandomizationSalt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 127,
		//	  "minLength": 0,
		//	  "pattern": ".*",
		//	  "type": "string"
		//	}
		"randomization_salt": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RemoveSegment
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"remove_segment": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RunningStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Start Experiment. Default is False",
		//	  "oneOf": [
		//	    {
		//	      "required": [
		//	        "Status",
		//	        "AnalysisCompleteTime"
		//	      ]
		//	    },
		//	    {
		//	      "required": [
		//	        "Status",
		//	        "Reason",
		//	        "DesiredState"
		//	      ]
		//	    }
		//	  ],
		//	  "properties": {
		//	    "AnalysisCompleteTime": {
		//	      "description": "Provide the analysis Completion time for an experiment",
		//	      "type": "string"
		//	    },
		//	    "DesiredState": {
		//	      "description": "Provide CANCELLED or COMPLETED desired state when stopping an experiment",
		//	      "pattern": "^(CANCELLED|COMPLETED)",
		//	      "type": "string"
		//	    },
		//	    "Reason": {
		//	      "description": "Reason is a required input for stopping the experiment",
		//	      "type": "string"
		//	    },
		//	    "Status": {
		//	      "description": "Provide START or STOP action to apply on an experiment",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"running_status": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AnalysisCompleteTime
				"analysis_complete_time": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Provide the analysis Completion time for an experiment",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: DesiredState
				"desired_state": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Provide CANCELLED or COMPLETED desired state when stopping an experiment",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Reason
				"reason": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Reason is a required input for stopping the experiment",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Status
				"status": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Provide START or STOP action to apply on an experiment",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Start Experiment. Default is False",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SamplingRate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maximum": 100000,
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"sampling_rate": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Segment
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 0,
		//	  "pattern": "([-a-zA-Z0-9._]*)|(arn:[^:]*:[^:]*:[^:]*:[^:]*:segment/[-a-zA-Z0-9._]*)",
		//	  "type": "string"
		//	}
		"segment": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
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
		//	        "pattern": "",
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
		// Property: Treatments
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Description": {
		//	        "type": "string"
		//	      },
		//	      "Feature": {
		//	        "pattern": "([-a-zA-Z0-9._]*)|(arn:[^:]*:[^:]*:[^:]*:[^:]*:.*)",
		//	        "type": "string"
		//	      },
		//	      "TreatmentName": {
		//	        "maxLength": 127,
		//	        "minLength": 1,
		//	        "pattern": "[-a-zA-Z0-9._]*",
		//	        "type": "string"
		//	      },
		//	      "Variation": {
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "pattern": "[-a-zA-Z0-9._]*",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "TreatmentName",
		//	      "Feature",
		//	      "Variation"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 5,
		//	  "minItems": 2,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"treatments": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Feature
					"feature": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: TreatmentName
					"treatment_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Variation
					"variation": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Evidently::Experiment",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Evidently::Experiment").WithTerraformTypeName("awscc_evidently_experiment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"analysis_complete_time": "AnalysisCompleteTime",
		"arn":                    "Arn",
		"control_treatment_name": "ControlTreatmentName",
		"description":            "Description",
		"desired_change":         "DesiredChange",
		"desired_state":          "DesiredState",
		"entity_id_key":          "EntityIdKey",
		"event_pattern":          "EventPattern",
		"feature":                "Feature",
		"key":                    "Key",
		"metric_goals":           "MetricGoals",
		"metric_name":            "MetricName",
		"name":                   "Name",
		"online_ab_config":       "OnlineAbConfig",
		"project":                "Project",
		"randomization_salt":     "RandomizationSalt",
		"reason":                 "Reason",
		"remove_segment":         "RemoveSegment",
		"running_status":         "RunningStatus",
		"sampling_rate":          "SamplingRate",
		"segment":                "Segment",
		"split_weight":           "SplitWeight",
		"status":                 "Status",
		"tags":                   "Tags",
		"treatment":              "Treatment",
		"treatment_name":         "TreatmentName",
		"treatment_weights":      "TreatmentWeights",
		"treatments":             "Treatments",
		"unit_label":             "UnitLabel",
		"value":                  "Value",
		"value_key":              "ValueKey",
		"variation":              "Variation",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
