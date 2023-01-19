// Code generated by generators/resource/main.go; DO NOT EDIT.

package billingconductor_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSBillingConductorCustomLineItem_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::BillingConductor::CustomLineItem", "awscc_billingconductor_custom_line_item", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
