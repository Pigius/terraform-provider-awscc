// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package apigateway_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSApiGatewayDomainNamesDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ApiGateway::DomainName", "awscc_apigateway_domain_names", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyDataSourceConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(fmt.Sprintf("data.%s", td.ResourceName), "ids.#"),
			),
		},
	})
}
