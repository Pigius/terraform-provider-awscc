// Code generated by generators/resource/main.go; DO NOT EDIT.

package datasync_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSDataSyncLocationNFS_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::DataSync::LocationNFS", "awscc_datasync_location_nfs", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
