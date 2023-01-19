// Code generated by generators/resource/main.go; DO NOT EDIT.

package eks_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSEKSFargateProfile_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EKS::FargateProfile", "awscc_eks_fargate_profile", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
