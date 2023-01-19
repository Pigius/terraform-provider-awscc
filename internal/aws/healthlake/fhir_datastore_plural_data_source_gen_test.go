// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package healthlake_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSHealthLakeFHIRDatastoresDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::HealthLake::FHIRDatastore", "awscc_healthlake_fhir_datastores", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyDataSourceConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(fmt.Sprintf("data.%s", td.ResourceName), "ids.#"),
			),
		},
	})
}
