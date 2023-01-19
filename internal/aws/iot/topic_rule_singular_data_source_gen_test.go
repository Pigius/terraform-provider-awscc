// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iot_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSIoTTopicRuleDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IoT::TopicRule", "awscc_iot_topic_rule", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSIoTTopicRuleDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IoT::TopicRule", "awscc_iot_topic_rule", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
