// Code generated by generators/resource/main.go; DO NOT EDIT.

package config_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSConfigStoredQuery_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Config::StoredQuery", "awscc_config_stored_query", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
