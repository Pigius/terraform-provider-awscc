// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSApiGatewayUsagePlanKeyDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ApiGateway::UsagePlanKey", "awscc_apigateway_usage_plan_key", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSApiGatewayUsagePlanKeyDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ApiGateway::UsagePlanKey", "awscc_apigateway_usage_plan_key", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
