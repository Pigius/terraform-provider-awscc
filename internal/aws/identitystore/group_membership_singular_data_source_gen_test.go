// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package identitystore_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSIdentityStoreGroupMembershipDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IdentityStore::GroupMembership", "awscc_identitystore_group_membership", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSIdentityStoreGroupMembershipDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IdentityStore::GroupMembership", "awscc_identitystore_group_membership", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
