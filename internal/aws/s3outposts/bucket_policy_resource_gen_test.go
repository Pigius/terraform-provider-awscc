// Code generated by generators/resource/main.go; DO NOT EDIT.

package s3outposts_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSS3OutpostsBucketPolicy_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::S3Outposts::BucketPolicy", "awscc_s3outposts_bucket_policy", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
