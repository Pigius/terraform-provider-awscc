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
	registry.AddDataSourceTypeFactory("awscc_lightsail_alarm", alarmDataSourceType)
}

// alarmDataSourceType returns the Terraform awscc_lightsail_alarm data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Lightsail::Alarm resource type.
func alarmDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"alarm_arn": {
			// Property: AlarmArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"alarm_name": {
			// Property: AlarmName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name for the alarm. Specify the name of an existing alarm to update, and overwrite the previous configuration of the alarm.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name for the alarm. Specify the name of an existing alarm to update, and overwrite the previous configuration of the alarm.",
			Type:        types.StringType,
			Computed:    true,
		},
		"comparison_operator": {
			// Property: ComparisonOperator
			// CloudFormation resource type schema:
			// {
			//   "description": "The arithmetic operation to use when comparing the specified statistic to the threshold. The specified statistic value is used as the first operand.",
			//   "type": "string"
			// }
			Description: "The arithmetic operation to use when comparing the specified statistic to the threshold. The specified statistic value is used as the first operand.",
			Type:        types.StringType,
			Computed:    true,
		},
		"contact_protocols": {
			// Property: ContactProtocols
			// CloudFormation resource type schema:
			// {
			//   "description": "The contact protocols to use for the alarm, such as Email, SMS (text messaging), or both.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The contact protocols to use for the alarm, such as Email, SMS (text messaging), or both.",
			Type:        types.SetType{ElemType: types.StringType},
			Computed:    true,
		},
		"datapoints_to_alarm": {
			// Property: DatapointsToAlarm
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of data points that must be not within the specified threshold to trigger the alarm. If you are setting an \"M out of N\" alarm, this value (datapointsToAlarm) is the M.",
			//   "type": "integer"
			// }
			Description: "The number of data points that must be not within the specified threshold to trigger the alarm. If you are setting an \"M out of N\" alarm, this value (datapointsToAlarm) is the M.",
			Type:        types.NumberType,
			Computed:    true,
		},
		"evaluation_periods": {
			// Property: EvaluationPeriods
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of most recent periods over which data is compared to the specified threshold. If you are setting an \"M out of N\" alarm, this value (evaluationPeriods) is the N.",
			//   "type": "integer"
			// }
			Description: "The number of most recent periods over which data is compared to the specified threshold. If you are setting an \"M out of N\" alarm, this value (evaluationPeriods) is the N.",
			Type:        types.NumberType,
			Computed:    true,
		},
		"metric_name": {
			// Property: MetricName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the metric to associate with the alarm.",
			//   "type": "string"
			// }
			Description: "The name of the metric to associate with the alarm.",
			Type:        types.StringType,
			Computed:    true,
		},
		"monitored_resource_name": {
			// Property: MonitoredResourceName
			// CloudFormation resource type schema:
			// {
			//   "description": "The validation status of the SSL/TLS certificate.",
			//   "type": "string"
			// }
			Description: "The validation status of the SSL/TLS certificate.",
			Type:        types.StringType,
			Computed:    true,
		},
		"notification_enabled": {
			// Property: NotificationEnabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether the alarm is enabled. Notifications are enabled by default if you don't specify this parameter.",
			//   "type": "boolean"
			// }
			Description: "Indicates whether the alarm is enabled. Notifications are enabled by default if you don't specify this parameter.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"notification_triggers": {
			// Property: NotificationTriggers
			// CloudFormation resource type schema:
			// {
			//   "description": "The alarm states that trigger a notification.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The alarm states that trigger a notification.",
			Type:        types.SetType{ElemType: types.StringType},
			Computed:    true,
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "The current state of the alarm.",
			//   "type": "string"
			// }
			Description: "The current state of the alarm.",
			Type:        types.StringType,
			Computed:    true,
		},
		"threshold": {
			// Property: Threshold
			// CloudFormation resource type schema:
			// {
			//   "description": "The value against which the specified statistic is compared.",
			//   "type": "number"
			// }
			Description: "The value against which the specified statistic is compared.",
			Type:        types.NumberType,
			Computed:    true,
		},
		"treat_missing_data": {
			// Property: TreatMissingData
			// CloudFormation resource type schema:
			// {
			//   "description": "Sets how this alarm will handle missing data points.",
			//   "type": "string"
			// }
			Description: "Sets how this alarm will handle missing data points.",
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
		Description: "Data Source schema for AWS::Lightsail::Alarm",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lightsail::Alarm").WithTerraformTypeName("awscc_lightsail_alarm")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alarm_arn":               "AlarmArn",
		"alarm_name":              "AlarmName",
		"comparison_operator":     "ComparisonOperator",
		"contact_protocols":       "ContactProtocols",
		"datapoints_to_alarm":     "DatapointsToAlarm",
		"evaluation_periods":      "EvaluationPeriods",
		"metric_name":             "MetricName",
		"monitored_resource_name": "MonitoredResourceName",
		"notification_enabled":    "NotificationEnabled",
		"notification_triggers":   "NotificationTriggers",
		"state":                   "State",
		"threshold":               "Threshold",
		"treat_missing_data":      "TreatMissingData",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
