// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package autoscaling_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSAutoScalingWarmPoolDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::AutoScaling::WarmPool", "awscc_autoscaling_warm_pool", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSAutoScalingWarmPoolDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::AutoScaling::WarmPool", "awscc_autoscaling_warm_pool", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
