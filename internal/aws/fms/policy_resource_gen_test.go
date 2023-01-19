// Code generated by generators/resource/main.go; DO NOT EDIT.

package fms_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSFMSPolicy_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::FMS::Policy", "awscc_fms_policy", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
