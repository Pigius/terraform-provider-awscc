// Code generated by generators/resource/main.go; DO NOT EDIT.

package acmpca_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSACMPCACertificateAuthority_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ACMPCA::CertificateAuthority", "awscc_acmpca_certificate_authority", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
