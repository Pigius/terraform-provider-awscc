// Code generated by generators/resource/main.go; DO NOT EDIT.

package ssm_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSSSMDocument_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::SSM::Document", "awscc_ssm_document", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
