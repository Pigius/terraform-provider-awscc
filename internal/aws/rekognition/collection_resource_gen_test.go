// Code generated by generators/resource/main.go; DO NOT EDIT.

package rekognition_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSRekognitionCollection_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Rekognition::Collection", "awscc_rekognition_collection", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
