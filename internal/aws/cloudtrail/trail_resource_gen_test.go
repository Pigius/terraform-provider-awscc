// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudtrail_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSCloudTrailTrail_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CloudTrail::Trail", "awscc_cloudtrail_trail", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
