package solidfire

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sofixa/terraform-provider-solidfire/solidfire/element"
	"github.com/sofixa/terraform-provider-solidfire/solidfire/element/jsonrpc"
)

func resourceSolidFireAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceSolidFireAccountCreate,
		Read:   resourceSolidFireAccountRead,
		Update: resourceSolidFireAccountUpdate,
		Delete: resourceSolidFireAccountDelete,
		Exists: resourceSolidFireAccountExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"initiator_secret": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target_secret": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceSolidFireAccountCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Creating account: %#v", d)
	client := meta.(*element.Client)

	acct := element.CreateAccountRequest{}

	if v, ok := d.GetOk("username"); ok {
		acct.Username = v.(string)
	} else {
		return fmt.Errorf("username argument is required")
	}

	if v, ok := d.GetOk("initiator_secret"); ok {
		acct.InitiatorSecret = v.(string)
	}

	if v, ok := d.GetOk("target_secret"); ok {
		acct.TargetSecret = v.(string)
	}

	resp, err := client.CreateAccount(acct)
	if err != nil {
		log.Print("Error creating account")
		return err
	}

	d.SetId(fmt.Sprintf("%v", resp.Account.AccountID))

	log.Printf("[DEBUG] Created account: %v %v", acct.Username, resp.Account.AccountID)

	return resourceSolidFireAccountRead(d, meta)
}

func resourceSolidFireAccountRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Reading account: %#v", d)
	client := meta.(*element.Client)

	id := d.Id()
	convID, convErr := strconv.Atoi(id)

	if convErr != nil {
		return fmt.Errorf("id argument is required")
	}

	res, err := client.GetAccountByID(convID)
	if err != nil {
		log.Print("GetAccountByID failed")
		return err
	}
	d.Set("username", res.Username)
	d.Set("initiator_secret", res.InitiatorSecret)
	d.Set("target_secret", res.TargetSecret)

	return nil
}

func resourceSolidFireAccountUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Updating account %#v", d)
	client := meta.(*element.Client)

	acct := element.ModifyAccountRequest{}

	id := d.Id()
	convID, convErr := strconv.Atoi(id)

	if convErr != nil {
		return fmt.Errorf("id argument is required")
	}
	acct.AccountID = convID

	if v, ok := d.GetOk("username"); ok {
		acct.Username = v.(string)
	}

	if v, ok := d.GetOk("initiator_secret"); ok {
		acct.InitiatorSecret = v.(string)
	}

	if v, ok := d.GetOk("target_secret"); ok {
		acct.TargetSecret = v.(string)
	}

	err := client.ModifyAccount(acct)
	if err != nil {
		return err
	}

	return nil
}

func resourceSolidFireAccountDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Deleting account: %#v", d)
	client := meta.(*element.Client)

	acct := element.RemoveAccountRequest{}

	id := d.Id()
	convID, convErr := strconv.Atoi(id)

	if convErr != nil {
		return fmt.Errorf("id argument is required")
	}
	acct.AccountID = convID

	err := client.RemoveAccount(acct)
	if err != nil {
		return err
	}

	return nil
}

func resourceSolidFireAccountExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	log.Printf("[DEBUG] Checking existence of account: %#v", d)
	client := meta.(*element.Client)

	id := d.Id()
	convID, convErr := strconv.Atoi(id)

	if convErr != nil {
		return false, fmt.Errorf("id argument is required")
	}

	_, err := client.GetAccountByID(convID)
	if err != nil {
		if err, ok := err.(*jsonrpc.ResponseError); ok {
			if err.Name == "xUnknownAccount" {
				d.SetId("")
				return false, nil
			}
		}
		log.Print("AccountExists failed")
		return false, err
	}

	return true, nil
}
