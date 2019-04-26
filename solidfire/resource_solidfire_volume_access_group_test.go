package solidfire

import (
	"strconv"
	"testing"

	"fmt"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sofixa/terraform-provider-solidfire/solidfire/element"
)

func TestVolumeAccessGroup_basic(t *testing.T) {
	var volumeAccessGroup element.VolumeAccessGroup
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSolidFireVolumeAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(
					testAccCheckSolidFireVolumeAccessGroupConfig,
					"terraform-acceptance-test",
					"Terraform-Acceptance-Volume-1",
					"1073741825",
					"true",
					"600",
					"8000",
					"8000",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSolidFireVolumeAccessGroupExists("solidfire_volume_access_group.terraform-acceptance-test-1", &volumeAccessGroup),
					resource.TestCheckResourceAttr("solidfire_volume_access_group.terraform-acceptance-test-1", "name", "terraform-acceptance-test"),
					resource.TestCheckResourceAttrSet("solidfire_volume_access_group.terraform-acceptance-test-1", "id"),
					resource.TestCheckResourceAttrSet("solidfire_volume_access_group.terraform-acceptance-test-1", "volumes"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "name", "erraform-Acceptance-Volume-1"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "total_size", "1073741825"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "enable512e", "true"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "min_iops", "600"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "max_iops", "8000"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "burst_iops", "8000"),
					resource.TestCheckResourceAttr("solidfire_account.terraform-acceptance-test-1", "username", "terraform-acceptance-test-vag"),
				),
			},
		},
	})
}

func TestVolumeAccessGroup_update(t *testing.T) {
	var volumeAccessGroup element.VolumeAccessGroup
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSolidFireVolumeAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(
					testAccCheckSolidFireVolumeAccessGroupConfig,
					"terraform-acceptance-test",
					"Terraform-Acceptance-Volume-1",
					"1073741825",
					"true",
					"600",
					"8000",
					"8000",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSolidFireVolumeAccessGroupExists("solidfire_volume_access_group.terraform-acceptance-test-1", &volumeAccessGroup),
					resource.TestCheckResourceAttr("solidfire_volume_access_group.terraform-acceptance-test-1", "name", "terraform-acceptance-test"),
					resource.TestCheckResourceAttrSet("solidfire_volume_access_group.terraform-acceptance-test-1", "id"),
					resource.TestCheckResourceAttrSet("solidfire_volume_access_group.terraform-acceptance-test-1", "volumes"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "name", "erraform-Acceptance-Volume-1"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "total_size", "1073741825"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "enable512e", "true"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "min_iops", "600"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "max_iops", "8000"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "burst_iops", "8000"),
					resource.TestCheckResourceAttr("solidfire_account.terraform-acceptance-test-1", "username", "terraform-acceptance-test-vag"),
				),
			},
			{
				Config: fmt.Sprintf(
					testAccCheckSolidFireVolumeAccessGroupConfigUpdate,
					"terraform-acceptance-test-update",
					"Terraform-Acceptance-Volume-1",
					"1093741825",
					"true",
					"600",
					"8000",
					"8000",
					"Terraform-Acceptance-Volume-2",
					"1093741825",
					"true",
					"600",
					"8000",
					"8000",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSolidFireVolumeAccessGroupExists("solidfire_volume_access_group.terraform-acceptance-test-1", &volumeAccessGroup),
					resource.TestCheckResourceAttr("solidfire_volume_access_group.terraform-acceptance-test-1", "name", "terraform-acceptance-test-update"),
					resource.TestCheckResourceAttrSet("solidfire_volume_access_group.terraform-acceptance-test-1", "id"),
					resource.TestCheckResourceAttrSet("solidfire_volume_access_group.terraform-acceptance-test-1", "volumes"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "name", "erraform-Acceptance-Volume-1"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "total_size", "1093741825"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "enable512e", "true"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "min_iops", "600"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "max_iops", "8000"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-1", "burst_iops", "8000"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-2", "name", "erraform-Acceptance-Volume-1"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-2", "total_size", "1093741825"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-2", "enable512e", "true"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-2", "min_iops", "600"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-2", "max_iops", "8000"),
					resource.TestCheckResourceAttr("solidfire_volume.terraform-acceptance-test-2", "burst_iops", "8000"),
					resource.TestCheckResourceAttr("solidfire_account.terraform-acceptance-test-1", "username", "terraform-acceptance-test-vag"),

				),
			},
		},
	})
}

 /*func TestVolumeAccessGroup_removeVolumes(t *testing.T) {
 	var volumeAccessGroup element.VolumeAccessGroup
 	resource.Test(t, resource.TestCase{
 		PreCheck:     func() { testAccPreCheck(t) },
 		Providers:    testAccProviders,
 		CheckDestroy: testAccCheckSolidFireVolumeAccessGroupDestroy,
 		Steps: []resource.TestStep{
 			{
 				Config: fmt.Sprintf(
 					testAccCheckSolidFireVolumeAccessGroupConfig,
					 "terraform-acceptance-test",
					 "Terraform-Acceptance-Volume-1",
					 "1073741825",
					 "true",
					 "600",
					 "8000",
					 "8000",
				  ),
 				Check: resource.ComposeTestCheckFunc(
 					testAccCheckSolidFireVolumeAccessGroupExists("solidfire_volume_access_group.terraform-acceptance-test-1", &volumeAccessGroup),
 					resource.TestCheckResourceAttr("solidfire_volume_access_group.terraform-acceptance-test-1", "name", "terraform-acceptance-test"),
 				),
 			},
 			{
 				Config: fmt.Sprintf(
 					testAccCheckSolidFireVolumeAccessGroupConfigRemoveVolumes,
 					"terraform-acceptance-test-remove",
 				),
 				Check: resource.ComposeTestCheckFunc(
 					testAccCheckSolidFireVolumeAccessGroupExists("solidfire_volume_access_group.terraform-acceptance-test-1", &volumeAccessGroup),
 					resource.TestCheckResourceAttr("solidfire_volume_access_group.terraform-acceptance-test-1", "name", "terraform-acceptance-test-update"),
 				),
 			},
 		},
 	})
 }
*/

func testAccCheckSolidFireVolumeAccessGroupDestroy(s *terraform.State) error {
	virConn := testAccProvider.Meta().(*element.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "solidfire_volume_access_group" {
			continue
		}

		_, err := virConn.GetVolumeAccessGroupByID(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Error waiting for volume access group (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckSolidFireVolumeAccessGroupExists(n string, volume *element.VolumeAccessGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		virConn := testAccProvider.Meta().(*element.Client)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SolidFire volume access group key ID is set")
		}

		retrievedVAG, err := virConn.GetVolumeAccessGroupByID(rs.Primary.ID)
		if err != nil {
			return err
		}

		convID, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}

		if retrievedVAG.VolumeAccessGroupID != convID {
			return fmt.Errorf("Resource ID and volume access group ID do not match")
		}

		*volume = retrievedVAG

		return nil
	}
}

const testAccCheckSolidFireVolumeAccessGroupConfig = `
resource "solidfire_volume_access_group" "terraform-acceptance-test-1" {
	name = "%s"
	volumes = ["${solidfire_volume.terraform-acceptance-test-1.id}"]
}
resource "solidfire_volume" "terraform-acceptance-test-1" {
	name = "%s"
	account_id = "${solidfire_account.terraform-acceptance-test-1.id}"
	total_size = "%s"
	enable512e = "%s"
	min_iops = "%s"
	max_iops = "%s"
	burst_iops = "%s"
}
resource "solidfire_account" "terraform-acceptance-test-1" {
	username = "terraform-acceptance-test-vag"
}
`

const testAccCheckSolidFireVolumeAccessGroupConfigUpdate = `
resource "solidfire_volume_access_group" "terraform-acceptance-test-1" {
	name = "%s"
	volumes = ["${solidfire_volume.terraform-acceptance-test-1.id}", "${solidfire_volume.terraform-acceptance-test-2.id}"]
}
resource "solidfire_volume" "terraform-acceptance-test-1" {
	name = "%s"
	account_id = "${solidfire_account.terraform-acceptance-test-1.id}"
	total_size = "%s"
	enable512e = "%s"
	min_iops = "%s"
	max_iops = "%s"
	burst_iops = "%s"
}
resource "solidfire_volume" "terraform-acceptance-test-2" {
	name = "%s"
	account_id = "${solidfire_account.terraform-acceptance-test-1.id}"
	total_size = "%s"
	enable512e = "%s"
	min_iops = "%s"
	max_iops = "%s"
	burst_iops = "%s"
}
resource "solidfire_account" "terraform-acceptance-test-1" {
	username = "terraform-acceptance-test-vag"
}
`

const testAccCheckSolidFireVolumeAccessGroupConfigRemoveVolumes = `
resource "solidfire_volume_access_group" "terraform-acceptance-test-1" {
	name = "%s"
	volumes = ["${solidfire_volume.terraform-acceptance-test-1.id}"]
}
`
