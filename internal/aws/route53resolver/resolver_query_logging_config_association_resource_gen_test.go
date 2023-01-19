// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53resolver_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSRoute53ResolverResolverQueryLoggingConfigAssociation_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation", "awscc_route53resolver_resolver_query_logging_config_association", "test")

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

func TestAccAWSRoute53ResolverResolverQueryLoggingConfigAssociation_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation", "awscc_route53resolver_resolver_query_logging_config_association", "test")

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
