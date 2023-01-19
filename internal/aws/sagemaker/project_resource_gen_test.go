// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSSageMakerProject_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::SageMaker::Project", "awscc_sagemaker_project", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
