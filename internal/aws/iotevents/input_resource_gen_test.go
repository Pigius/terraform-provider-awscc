// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotevents_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSIoTEventsInput_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IoTEvents::Input", "awscc_iotevents_input", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
