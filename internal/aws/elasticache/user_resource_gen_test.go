// Code generated by generators/resource/main.go; DO NOT EDIT.

package elasticache_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSElastiCacheUser_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ElastiCache::User", "awscc_elasticache_user", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
