// Code generated by generators/resource/main.go; DO NOT EDIT.

package elasticbeanstalk_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSElasticBeanstalkApplication_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ElasticBeanstalk::Application", "awscc_elasticbeanstalk_application", "test")

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

func TestAccAWSElasticBeanstalkApplication_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ElasticBeanstalk::Application", "awscc_elasticbeanstalk_application", "test")

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
