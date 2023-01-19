// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53recoverycontrol_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSRoute53RecoveryControlControlPanel_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Route53RecoveryControl::ControlPanel", "awscc_route53recoverycontrol_control_panel", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
