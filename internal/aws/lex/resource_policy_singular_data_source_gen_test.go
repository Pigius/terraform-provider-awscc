// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lex_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSLexResourcePolicyDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Lex::ResourcePolicy", "awscc_lex_resource_policy", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSLexResourcePolicyDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Lex::ResourcePolicy", "awscc_lex_resource_policy", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
