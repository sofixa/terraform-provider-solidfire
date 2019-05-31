package solidfire

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sofixa/terraform-provider-solidfire/solidfire/element"
	"log"
	"strconv"
)

func testAccCheckSolidFireVolumeDestroy(s *terraform.State) error {
	virConn := testAccProvider.Meta().(*element.Client)
	var volume element.Volume
	var err error

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "solidfire_volume" {
			continue
		}

		volume, err = virConn.GetVolumeByID(rs.Primary.ID)
		// volume should have been purged
		if rs.Primary.Attributes["purge_on_delete"] == "true" {
			if err == nil {
				return fmt.Errorf("Error waiting for volume %s to be destroyed, it should have been purged", rs.Primary.ID)
			}
		} else {
			// if there isn't an error, the volume still exists, as it should
			if err == nil {
				// if the volume's status isn't deleted, it isn't marked as to be purged
				if volume.Status != "deleted" {
					return fmt.Errorf("Volume %s still exists and status isn't deleted, it's %s", rs.Primary.ID, volume.Status)
					// everything is working fine (volume was marked as deleted and will be by the SF), launch an explicit purge to make place for future tests
				} else {
					log.Printf("[DEBUG] Volume %s wasn't purged due to purge_on_delete=false, purging explicitly to clean up", rs.Primary.ID)
					delVolume := element.DeleteVolumeRequest{}
					convID, convErr := strconv.Atoi(rs.Primary.ID)

					if convErr != nil {
						return fmt.Errorf("id argument is required")
					}
					delVolume.VolumeID = convID
					err := virConn.PurgeDeletedVolume(delVolume)
					if err != nil {
						return fmt.Errorf("Failed purging volume %s due to error %s", rs.Primary.ID, err)
					}
				}
			} else {
				return fmt.Errorf("Volume %s doesn't exist anymore, but it shouldn't have been purged", rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccCheckSolidFireVolumeExists(n string, volume *element.Volume) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		virConn := testAccProvider.Meta().(*element.Client)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SolidFire volume key ID is set")
		}

		retrievedVol, err := virConn.GetVolumeByID(rs.Primary.ID)
		if err != nil {
			return err
		}

		convID, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}

		if retrievedVol.VolumeID != convID {
			return fmt.Errorf("Resource ID and volume ID do not match")
		}

		*volume = retrievedVol

		return nil
	}
}


