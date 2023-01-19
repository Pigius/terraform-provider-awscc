// Code generated by generators/resource/main.go; DO NOT EDIT.

package iot_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSIoTTopicRuleDestination_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IoT::TopicRuleDestination", "awscc_iot_topic_rule_destination", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
			),
		},
		{
			ResourceName:      td.ResourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
	})
}

func TestAccAWSIoTTopicRuleDestination_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IoT::TopicRuleDestination", "awscc_iot_topic_rule_destination", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				td.DeleteResource(),
			),
			ExpectNonEmptyPlan: true,
		},
	})
}
