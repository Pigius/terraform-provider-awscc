// Code generated by generators/resource/main.go; DO NOT EDIT.

package xray_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSXRaySamplingRule_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::XRay::SamplingRule", "awscc_xray_sampling_rule", "test")

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

func TestAccAWSXRaySamplingRule_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::XRay::SamplingRule", "awscc_xray_sampling_rule", "test")

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
