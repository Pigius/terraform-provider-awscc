// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudformation_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSCloudFormationResourceVersionDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CloudFormation::ResourceVersion", "awscc_cloudformation_resource_version", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSCloudFormationResourceVersionDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CloudFormation::ResourceVersion", "awscc_cloudformation_resource_version", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
