// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package supportapp_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSSupportAppSlackChannelConfigurationDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::SupportApp::SlackChannelConfiguration", "awscc_supportapp_slack_channel_configuration", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSSupportAppSlackChannelConfigurationDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::SupportApp::SlackChannelConfiguration", "awscc_supportapp_slack_channel_configuration", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
