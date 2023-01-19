// Code generated by generators/resource/main.go; DO NOT EDIT.

package ses_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSSESEmailIdentity_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::SES::EmailIdentity", "awscc_ses_email_identity", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
