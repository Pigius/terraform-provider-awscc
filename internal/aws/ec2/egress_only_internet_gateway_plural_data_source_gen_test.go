// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package ec2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSEC2EgressOnlyInternetGatewaysDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EC2::EgressOnlyInternetGateway", "awscc_ec2_egress_only_internet_gateways", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyDataSourceConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(fmt.Sprintf("data.%s", td.ResourceName), "ids.#"),
			),
		},
	})
}
