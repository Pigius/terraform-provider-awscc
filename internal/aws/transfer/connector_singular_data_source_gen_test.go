// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package transfer_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSTransferConnectorDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Transfer::Connector", "awscc_transfer_connector", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSTransferConnectorDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Transfer::Connector", "awscc_transfer_connector", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
