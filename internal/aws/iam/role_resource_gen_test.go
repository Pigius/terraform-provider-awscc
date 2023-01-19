// Code generated by generators/resource/main.go; DO NOT EDIT.

package iam_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSIAMRole_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IAM::Role", "awscc_iam_role", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
