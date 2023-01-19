// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotwireless_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSIoTWirelessWirelessDeviceDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IoTWireless::WirelessDevice", "awscc_iotwireless_wireless_device", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSIoTWirelessWirelessDeviceDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IoTWireless::WirelessDevice", "awscc_iotwireless_wireless_device", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
