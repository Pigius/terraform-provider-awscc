// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudfront_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSCloudFrontContinuousDeploymentPolicy_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CloudFront::ContinuousDeploymentPolicy", "awscc_cloudfront_continuous_deployment_policy", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
