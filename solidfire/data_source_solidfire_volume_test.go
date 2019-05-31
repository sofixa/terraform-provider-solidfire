package solidfire

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/sofixa/terraform-provider-solidfire/solidfire/element"
	"testing"
)

func TestAccDataSourceVolume_basic(t *testing.T) {
	var volume element.Volume
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSolidFireVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(
					testAccCheckDataSourceSolidFireVolumeConfig,
					"terraform-acceptance-test",
					"1080033280",
					"true",
					"500",
					"8000",
					"10000",
					"true",
					"terraform-acceptance-test",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSolidFireVolumeExists("solidfire_volume.terraform-acceptance-test-1", &volume),
					testAccCheckSolidFireVolumeAttributes(&volume),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "name", "data.solidfire_volume","name"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "total_size", "data.solidfire_volume","total_size"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "enable512e", "data.solidfire_volume","enable512e"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "min_iops", "data.solidfire_volume","min_iops"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "max_iops", "data.solidfire_volume","max_iops"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "burst_iops", "data.solidfire_volume","burst_iops"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "access", "data.solidfire_volume","access"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "status", "data.solidfire_volume","status"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "purge_on_delete", "data.solidfire_volume","purge_on_delete"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "attributes", "data.solidfire_volume","attributes"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "volume_id", "data.solidfire_volume","volume_id"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "iqn", "data.solidfire_volume","iqn"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "block_size", "data.solidfire_volume","block_size"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "scsi_eui_device_id", "data.solidfire_volume","scsi_eui_device_id"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "scsi_naa_device_id", "data.solidfire_volume","scsi_naa_device_id"),
				),
			},
		},
	})
}

func TestAccDataSourceVolume_update(t *testing.T) {
	var volume element.Volume
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSolidFireVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(
					testAccCheckDataSourceSolidFireVolumeConfig,
					"terraform-acceptance-test",
					"1080033280",
					"true",
					"500",
					"8000",
					"10000",
					"false",
					"terraform-acceptance-test",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSolidFireVolumeExists("solidfire_volume.terraform-acceptance-test-1", &volume),
					testAccCheckSolidFireVolumeAttributes(&volume),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "name", "data.solidfire_volume","name"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "total_size", "data.solidfire_volume","total_size"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "enable512e", "data.solidfire_volume","enable512e"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "min_iops", "data.solidfire_volume","min_iops"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "max_iops", "data.solidfire_volume","max_iops"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "burst_iops", "data.solidfire_volume","burst_iops"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "access", "data.solidfire_volume","access"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "status", "data.solidfire_volume","status"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "purge_on_delete", "data.solidfire_volume","purge_on_delete"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "attributes", "data.solidfire_volume","attributes"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "volume_id", "data.solidfire_volume","volume_id"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "iqn", "data.solidfire_volume","iqn"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "block_size", "data.solidfire_volume","block_size"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "scsi_eui_device_id", "data.solidfire_volume","scsi_eui_device_id"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "scsi_naa_device_id", "data.solidfire_volume","scsi_naa_device_id"),
				),
			},
			{
				Config: fmt.Sprintf(
					testAccCheckDataSourceSolidFireVolumeConfigUpdate,
					"terraform-acceptance-test",
					"1090519040",
					"true",
					"650",
					"8600",
					"9600",
					"true",
					"terraform-acceptance-test",

				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSolidFireVolumeExists("solidfire_volume.terraform-acceptance-test-1", &volume),
					testAccCheckSolidFireVolumeAttributesUpdate(&volume),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "name", "data.solidfire_volume","name"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "total_size", "data.solidfire_volume","total_size"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "enable512e", "data.solidfire_volume","enable512e"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "min_iops", "data.solidfire_volume","min_iops"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "max_iops", "data.solidfire_volume","max_iops"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "burst_iops", "data.solidfire_volume","burst_iops"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "access", "data.solidfire_volume","access"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "status", "data.solidfire_volume","status"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "purge_on_delete", "data.solidfire_volume","purge_on_delete"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "attributes", "data.solidfire_volume","attributes"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "volume_id", "data.solidfire_volume","volume_id"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "iqn", "data.solidfire_volume","iqn"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "block_size", "data.solidfire_volume","block_size"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "scsi_eui_device_id", "data.solidfire_volume","scsi_eui_device_id"),
					resource.TestCheckResourceAttrPair("solidfire_volume.terraform-acceptance-test-1", "scsi_naa_device_id", "data.solidfire_volume","scsi_naa_device_id"),
				),
			},
		},
	})
}



const testAccCheckDataSourceSolidFireVolumeConfig = `
resource "solidfire_volume" "terraform-acceptance-test-1" {
	name = "%s"
	account_id = "${solidfire_account.terraform-acceptance-test-1.id}"
	total_size = "%s"
	enable512e = "%s"
	min_iops = "%s"
	max_iops = "%s"
	burst_iops = "%s"
	purge_on_delete = "%s"
	attributes = {
		test = "test"
	}
}

data "solidfire_volume" "terraform-acceptance-test-1" {
	name_filter = "%s"
}

resource "solidfire_account" "terraform-acceptance-test-1" {
	username = "terraform-acceptance-test-volume"
}
`
const testAccCheckDataSourceSolidFireVolumeConfigUpdate = `
resource "solidfire_volume" "terraform-acceptance-test-1" {
	name = "%s"
	account_id = "${solidfire_account.terraform-acceptance-test-1.id}"
	total_size = "%s"
	enable512e = "%s"
	min_iops = "%s"
	max_iops = "%s"
	burst_iops = "%s"
	purge_on_delete = "%s"
	attributes = {
		test = "test"
		test2 = "test2"
		test3 = "test3"
	}
}

data "solidfire_volume" "terraform-acceptance-test-1" {
	name_filter = "%s"
}

resource "solidfire_account" "terraform-acceptance-test-1" {
	username = "terraform-acceptance-test-volume"
}
`