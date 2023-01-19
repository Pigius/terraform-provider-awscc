// Code generated by generators/resource/main.go; DO NOT EDIT.

package logs_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSLogsLogStream_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Logs::LogStream", "awscc_logs_log_stream", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
