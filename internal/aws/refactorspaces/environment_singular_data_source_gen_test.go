// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package refactorspaces_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSRefactorSpacesEnvironmentDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::RefactorSpaces::Environment", "awscc_refactorspaces_environment", "test")

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

func TestAccAWSRefactorSpacesEnvironmentDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::RefactorSpaces::Environment", "awscc_refactorspaces_environment", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
