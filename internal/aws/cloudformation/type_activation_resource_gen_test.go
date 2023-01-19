// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudformation_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSCloudFormationTypeActivation_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CloudFormation::TypeActivation", "awscc_cloudformation_type_activation", "test")

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

func TestAccAWSCloudFormationTypeActivation_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CloudFormation::TypeActivation", "awscc_cloudformation_type_activation", "test")

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
