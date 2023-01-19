// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package route53_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSRoute53HostedZonesDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Route53::HostedZone", "awscc_route53_hosted_zones", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyDataSourceConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(fmt.Sprintf("data.%s", td.ResourceName), "ids.#"),
			),
		},
	})
}
