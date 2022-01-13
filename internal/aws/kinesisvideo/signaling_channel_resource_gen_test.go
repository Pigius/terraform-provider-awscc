// Code generated by generators/resource/main.go; DO NOT EDIT.

package kinesisvideo_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSKinesisVideoSignalingChannel_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::KinesisVideo::SignalingChannel", "awscc_kinesisvideo_signaling_channel", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
			),
		},
		{
			ResourceName:      td.ResourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
	})
}

func TestAccAWSKinesisVideoSignalingChannel_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::KinesisVideo::SignalingChannel", "awscc_kinesisvideo_signaling_channel", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				td.DeleteResource(),
			),
			ExpectNonEmptyPlan: true,
		},
	})
}
