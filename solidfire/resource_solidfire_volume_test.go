package solidfire

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sofixa/terraform-provider-solidfire/solidfire/element"
	"strconv"
	"testing"
)

func TestAccResourceVolume_basic(t *testing.T) {
	var volume element.Volume
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSolidFireVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(
					testAccCheckResourceSolidFireVolumeConfig,
					"terraform-acceptance-test",
					"1080033280",
					"true",
					"500",
					"8000",
					"10000",
					"true",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSolidFireVolumeExists("solidfire_volume.terraform-acceptance-test-1", &volume),
					testAccCheckSolidFireVolumeAttributes(&volume),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "name", "terraform-acceptance-test"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "total_size", "1080033280"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "enable512e", "true"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "min_iops", "500"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "max_iops", "8000"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "burst_iops", "10000"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "access", "readWrite"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "status", "active"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "purge_on_delete", "true"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "attributes.%", "1"),
					resource.TestCheckResourceAttrSet("solidfire_volume.terraform-acceptance-test-1", "volume_id"),
					resource.TestCheckResourceAttrSet("solidfire_volume.terraform-acceptance-test-1", "iqn"),
					resource.TestCheckResourceAttrSet("solidfire_volume.terraform-acceptance-test-1", "block_size"),
					resource.TestCheckResourceAttrSet("solidfire_volume.terraform-acceptance-test-1", "scsi_eui_device_id"),
					resource.TestCheckResourceAttrSet("solidfire_volume.terraform-acceptance-test-1", "scsi_naa_device_id"),
				),
			},
		},
	})
}

func TestAccResourceVolume_update(t *testing.T) {
	var volume element.Volume
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSolidFireVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(
					testAccCheckResourceSolidFireVolumeConfig,
					"terraform-acceptance-test",
					"1080033280",
					"true",
					"500",
					"8000",
					"10000",
					"false",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSolidFireVolumeExists("solidfire_volume.terraform-acceptance-test-1", &volume),
					testAccCheckSolidFireVolumeAttributes(&volume),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "name", "terraform-acceptance-test"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "total_size", "1080033280"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "enable512e", "true"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "min_iops", "500"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "max_iops", "8000"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "burst_iops", "10000"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "access", "readWrite"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "status", "active"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "purge_on_delete", "false"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "attributes.%", "1"),
					resource.TestCheckResourceAttrSet("solidfire_volume.terraform-acceptance-test-1", "volume_id"),
					resource.TestCheckResourceAttrSet("solidfire_volume.terraform-acceptance-test-1", "iqn"),
					resource.TestCheckResourceAttrSet("solidfire_volume.terraform-acceptance-test-1", "block_size"),
					resource.TestCheckResourceAttrSet("solidfire_volume.terraform-acceptance-test-1", "scsi_eui_device_id"),
					resource.TestCheckResourceAttrSet("solidfire_volume.terraform-acceptance-test-1", "scsi_naa_device_id"),
				),
			},
			{
				Config: fmt.Sprintf(
					testAccCheckResourceSolidFireVolumeConfigUpdate,
					"terraform-acceptance-test",
					"1090519040",
					"true",
					"650",
					"8600",
					"9600",
					"true",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSolidFireVolumeExists("solidfire_volume.terraform-acceptance-test-1", &volume),
					testAccCheckSolidFireVolumeAttributesUpdate(&volume),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "name", "terraform-acceptance-test"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "total_size", "1090519040"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "enable512e", "true"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "purge_on_delete", "true"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "min_iops", "650"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "max_iops", "8600"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "burst_iops", "9600"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "attributes.%", "3"),
				),
			},
		},
	})
}


// Compare the actual attributes as present on the SolidFire cluster via the SolidFire API
// to check there's no difference between the reality and TF's state
func testAccCheckSolidFireVolumeAttributes(volume *element.Volume) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		// Check all attributes are correct
		if volume.Name != "terraform-acceptance-test" {
			return fmt.Errorf("Volume name is %s, was expecting %s", volume.Name, "terraform-acceptance-test")
		}
		if volume.TotalSize != 1080033280 {
			return fmt.Errorf("Volume size is %d, was expecting %d", volume.TotalSize, 1080033280)
		}
		if volume.Enable512e != true {
			return fmt.Errorf("Volume 512e isn't enabled")
		}
		if volume.QOS.MinIOPS != 500 {
			return fmt.Errorf("Volume min_iops is %d, was expecting %d", volume.QOS.MinIOPS, 500)
		}
		if volume.QOS.MaxIOPS != 8000 {
			return fmt.Errorf("Volume max_iops is %d, was expecting %d", volume.QOS.MaxIOPS, 8000)
		}
		if volume.QOS.BurstIOPS != 10000 {
			return fmt.Errorf("Volume burst_iops is %d, was expecting %d", volume.QOS.BurstIOPS, 10000)
		}
		if volume.Access != "readWrite" {
			return fmt.Errorf("Volume access isn't readWrite")
		}
		if volume.Status != "active" {
			return fmt.Errorf("Volume is not active")
		}
		if len(volume.Attributes) != 1 {
			return fmt.Errorf("Number of attributes is %d, was expecting %d", len(volume.Attributes), 1)
		}
		for k, v := range volume.Attributes {
			if k != "test" || v != "test" {
				return fmt.Errorf("Attributes key and value are %s and %s, was expecting %s and %s", k, v, "test", "test")
			}
		}

		// Check volume's account_id and volume_access_group_id are correct
		for _, rs := range s.RootModule().Resources {
			if rs.Type == "solidfire_account" {
				convID, err := strconv.Atoi(rs.Primary.ID)
				if err != nil {
					return err
				}
				if convID != volume.AccountID {
					return fmt.Errorf("Volume account_id is %d, was expecting %d", volume.AccountID, convID)
				}
			}
		}
		return nil
	}
}

// Compare the actual attributes as present on the SolidFire cluster via the SolidFire API
// to check there's no difference between the reality and TF's state after the volume has been updated
func testAccCheckSolidFireVolumeAttributesUpdate(volume *element.Volume) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		// Check all attributes are correct
		if volume.Name != "terraform-acceptance-test" {
			return fmt.Errorf("Volume name is %s, was expecting %s", volume.Name, "terraform-acceptance-test")
		}
		if volume.TotalSize != 1090519040 {
			return fmt.Errorf("Volume size is %d, was expecting %d", volume.TotalSize, 1090519040)
		}
		if volume.Enable512e != true {
			return fmt.Errorf("Volume 512e isn't enabled")
		}
		if volume.QOS.MinIOPS != 650 {
			return fmt.Errorf("Volume min_iops is %d, was expecting %d", volume.QOS.MinIOPS, 650)
		}
		if volume.QOS.MaxIOPS != 8600 {
			return fmt.Errorf("Volume max_iops is %d, was expecting %d", volume.QOS.MaxIOPS, 8600)
		}
		if volume.QOS.BurstIOPS != 9600 {
			return fmt.Errorf("Volume burst_iops is %d, was expecting %d", volume.QOS.BurstIOPS, 9600)
		}
		if volume.Access != "readWrite" {
			return fmt.Errorf("Volume access isn't readWrite")
		}
		if volume.Status != "active" {
			return fmt.Errorf("Volume is not active")
		}
		if len(volume.Attributes) != 3 {
			return fmt.Errorf("Number of attributes is %d, was expecting %d", len(volume.Attributes), 3)
		}
		for k, v := range volume.Attributes {
			if k != v {
				return fmt.Errorf("Attributes key and value are %s and %s, was expecting them to be equal", k, v)
			}
		}

		// Check volume's account_id and volume_access_group_id are correct
		for _, rs := range s.RootModule().Resources {
			if rs.Type == "solidfire_account" {
				convID, err := strconv.Atoi(rs.Primary.ID)
				if err != nil {
					return err
				}
				if convID != volume.AccountID {
					return fmt.Errorf("Volume account_id is %d, was expecting %d", volume.AccountID, convID)
				}
			}
		}
		return nil
	}
}

const testAccCheckResourceSolidFireVolumeConfig = `
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

resource "solidfire_account" "terraform-acceptance-test-1" {
	username = "terraform-acceptance-test-volume"
}
`
const testAccCheckResourceSolidFireVolumeConfigUpdate = `
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

resource "solidfire_account" "terraform-acceptance-test-1" {
	username = "terraform-acceptance-test-volume"
}
`
