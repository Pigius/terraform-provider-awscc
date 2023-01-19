// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package dynamodb_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSDynamoDBGlobalTableDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::DynamoDB::GlobalTable", "awscc_dynamodb_global_table", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSDynamoDBGlobalTableDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::DynamoDB::GlobalTable", "awscc_dynamodb_global_table", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
