// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package s3_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSS3AccessPointsDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::S3::AccessPoint", "awscc_s3_access_points", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyDataSourceConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(fmt.Sprintf("data.%s", td.ResourceName), "ids.#"),
			),
		},
	})
}
