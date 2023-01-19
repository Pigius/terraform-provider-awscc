// Code generated by generators/resource/main.go; DO NOT EDIT.

package kendraranking_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSKendraRankingExecutionPlan_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::KendraRanking::ExecutionPlan", "awscc_kendraranking_execution_plan", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
