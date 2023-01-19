// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudformation_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSCloudFormationResourceDefaultVersionDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CloudFormation::ResourceDefaultVersion", "awscc_cloudformation_resource_default_version", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.DataSourceWithEmptyResourceConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrPair(fmt.Sprintf("data.%s", td.ResourceName), "id", td.ResourceName, "id"),
				resource.TestCheckResourceAttrPair(fmt.Sprintf("data.%s", td.ResourceName), "arn", td.ResourceName, "arn"),
			),
		},
	})
}

func TestAccAWSCloudFormationResourceDefaultVersionDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CloudFormation::ResourceDefaultVersion", "awscc_cloudformation_resource_default_version", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
