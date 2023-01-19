// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package cloudformation_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSCloudFormationTypeActivationsDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CloudFormation::TypeActivation", "awscc_cloudformation_type_activations", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyDataSourceConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(fmt.Sprintf("data.%s", td.ResourceName), "ids.#"),
			),
		},
	})
}
