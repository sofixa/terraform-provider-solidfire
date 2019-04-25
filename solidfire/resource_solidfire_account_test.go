package solidfire

import (
	"strconv"
	"testing"

	"fmt"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/solidfire/terraform-provider-solidfire/solidfire/element"
)

func TestAccount_basic(t *testing.T) {
	var account element.Account
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSolidFireAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(
					testAccCheckSolidFireAccountConfig,
					"terraform-acceptance-test",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSolidFireAccountExists("solidfire_account.terraform-acceptance-account-1", &account),
					resource.TestCheckResourceAttr("solidfire_account.terraform-acceptance-account-1", "username", "terraform-acceptance-test"),
					resource.TestCheckResourceAttrSet("solidfire_account.terraform-acceptance-account-1", "target_secret"),
					resource.TestCheckResourceAttrSet("solidfire_account.terraform-acceptance-account-1", "initiator_secret"),
					resource.TestCheckResourceAttrSet("solidfire_account.terraform-acceptance-account-1", "id"),
				),
			},
		},
	})
}

func TestAccount_secrets(t *testing.T) {
	var account element.Account
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSolidFireAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(
					testAccCheckSolidFireAccountConfigSecrets,
					"terraform-acceptance-test",
					"ABC123456XYZ",
					"SecretSecret1",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSolidFireAccountExists("solidfire_account.terraform-acceptance-account-1", &account),
					resource.TestCheckResourceAttr("solidfire_account.terraform-acceptance-account-1", "username", "terraform-acceptance-test"),
					resource.TestCheckResourceAttr("solidfire_account.terraform-acceptance-account-1", "target_secret", "ABC123456XYZ"),
					resource.TestCheckResourceAttr("solidfire_account.terraform-acceptance-account-1", "initiator_secret", "SecretSecret1"),
				),
			},
		},
	})
}

func TestAccount_update(t *testing.T) {
	var account element.Account
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSolidFireAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(
					testAccCheckSolidFireAccountConfigSecrets,
					"terraform-acceptance-test",
					"ABC123456XYZ",
					"SecretSecret1",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSolidFireAccountExists("solidfire_account.terraform-acceptance-account-1", &account),
					resource.TestCheckResourceAttr("solidfire_account.terraform-acceptance-account-1", "username", "terraform-acceptance-test"),
					resource.TestCheckResourceAttr("solidfire_account.terraform-acceptance-account-1", "target_secret", "ABC123456XYZ"),
					resource.TestCheckResourceAttr("solidfire_account.terraform-acceptance-account-1", "initiator_secret", "SecretSecret1"),
				),
			},
			{
				Config: fmt.Sprintf(
					testAccCheckSolidFireAccountConfigUpdate,
					"terraform-acceptance-test-update",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSolidFireAccountExists("solidfire_account.terraform-acceptance-account-1", &account),
					resource.TestCheckResourceAttr("solidfire_account.terraform-acceptance-account-1", "username", "terraform-acceptance-test-update"),
					resource.TestCheckResourceAttr("solidfire_account.terraform-acceptance-account-1", "target_secret", "ABC123456XYZU"),
					resource.TestCheckResourceAttr("solidfire_account.terraform-acceptance-account-1", "initiator_secret", "SecretSecret1U"),
				),
			},
		},
	})
}

func testAccCheckSolidFireAccountDestroy(s *terraform.State) error {
	virConn := testAccProvider.Meta().(*element.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "solidfire_account" {
			continue
		}

		convID, convErr := strconv.Atoi(rs.Primary.ID)
		if convErr != nil {
			return convErr
		}

		_, err := virConn.GetAccountByID(convID)
		if err == nil {
			return fmt.Errorf("Error waiting for volume (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckSolidFireAccountExists(n string, account *element.Account) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		virConn := testAccProvider.Meta().(*element.Client)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SolidFire account key ID is set")
		}

		convID, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}

		retrievedAcc, err := virConn.GetAccountByID(convID)
		if err != nil {
			return err
		}

		if retrievedAcc.AccountID != convID {
			return fmt.Errorf("Resource ID and account ID do not match")
		}

		*account = retrievedAcc

		return nil
	}
}

const testAccCheckSolidFireAccountConfig = `
resource "solidfire_account" "terraform-acceptance-account-1" {
	username = "%s"
}
`

const testAccCheckSolidFireAccountConfigSecrets = `
resource "solidfire_account" "terraform-acceptance-account-1" {
	username = "%s"
	target_secret = "%s"
	initiator_secret = "%s"
}
`

const testAccCheckSolidFireAccountConfigUpdate = `
resource "solidfire_account" "terraform-acceptance-account-1" {
	username = "%s"
}
`
