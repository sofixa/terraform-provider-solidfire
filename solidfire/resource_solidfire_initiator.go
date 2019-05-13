package solidfire

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sofixa/terraform-provider-solidfire/solidfire/element"
	"github.com/sofixa/terraform-provider-solidfire/solidfire/element/jsonrpc"
)

func resourceSolidFireInitiator() *schema.Resource {
	return &schema.Resource{
		Create: resourceSolidFireInitiatorCreate,
		Read:   resourceSolidFireInitiatorRead,
		Update: resourceSolidFireInitiatorUpdate,
		Delete: resourceSolidFireInitiatorDelete,
		Exists: resourceSolidFireInitiatorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"alias": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"volume_access_group_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"iqns": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceSolidFireInitiatorCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Creating initiator: %#v", d)
	client := meta.(*element.Client)

	initiators := element.CreateInitiatorsRequest{}
	newInitiator := make([]element.Initiator, 1)
	var iqns []string

	if v, ok := d.GetOk("name"); ok {
		newInitiator[0].Name = v.(string)
	} else {
		return fmt.Errorf("name argument is required")
	}

	if v, ok := d.GetOk("alias"); ok {
		newInitiator[0].Alias = v.(string)
	}

	if v, ok := d.GetOk("volume_access_group_id"); ok {
		newInitiator[0].VolumeAccessGroupID = v.(int)
	}

	if v, ok := d.GetOk("iqns"); ok {

		if a, ok := v.([]interface{}); ok {
			for i := range a {
				iqns = append(iqns, a[i].(string))
			}
		}
	}

	initiators.Initiators = newInitiator

	resp, err := client.CreateInitiators(initiators)
	if err != nil {
		log.Print("Error creating initiator")
		return err
	}

	d.SetId(fmt.Sprintf("%v", resp.Initiators[0].ID))
	log.Printf("[DEBUG] Created initiator: %v %v", newInitiator[0].Name, resp.Initiators[0].ID)

	return resourceSolidFireInitiatorRead(d, meta)
}

func resourceSolidFireInitiatorRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Reading initiator: %#v", d)
	client := meta.(*element.Client)

	initiators := element.ListInitiatorRequest{}

	id := d.Id()
	s := make([]int, 1)
	convID, convErr := strconv.Atoi(id)

	if convErr != nil {
		return fmt.Errorf("id argument is required")
	}

	s[0] = convID
	initiators.Initiators = s

	res, err := client.ListInitiators(initiators)
	if err != nil {
		return err
	}

	if len(res.Initiators) != 1 {
		return fmt.Errorf("Expected one Initiator to be found. Response contained %v results", len(res.Initiators))
	}

	d.Set("name", res.Initiators[0].Name)
	d.Set("alias", res.Initiators[0].Alias)
	d.Set("attributes", res.Initiators[0].Attributes)

	if len(res.Initiators[0].VolumeAccessGroups) == 1 {
		d.Set("volume_access_group_id", res.Initiators[0].VolumeAccessGroups[0])
	}

	return nil
}

func resourceSolidFireInitiatorUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Updating initiator: %#v", d)
	client := meta.(*element.Client)

	initiators := ModifyInitiatorsRequest{}
	initiator := make([]element.Initiator, 1)

	id := d.Id()
	convID, convErr := strconv.Atoi(id)

	if convErr != nil {
		return fmt.Errorf("id argument is required")
	}

	initiator[0].InitiatorID = convID

	if v, ok := d.GetOk("alias"); ok {
		initiator[0].Alias = v.(string)
	}

	if v, ok := d.GetOk("volume_access_group_id"); ok {
		initiator[0].VolumeAccessGroupID = v.(int)
	}

	initiators.Initiators = initiator

	err := client.ModifyInitiators(initiators)
	if err != nil {
		return err
	}

	return nil
}

func resourceSolidFireInitiatorDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Deleting initiator: %#v", d)
	client := meta.(*element.Client)

	initiators := element.DeleteInitiatorsRequest{}

	id := d.Id()
	s := make([]int, 1)
	convID, convErr := strconv.Atoi(id)

	if convErr != nil {
		return fmt.Errorf("id argument is required")
	}

	s[0] = convID
	initiators.Initiators = s

	err := client.DeleteInitiator(client)
	if err != nil {
		return err
	}

	return nil
}

func resourceSolidFireInitiatorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	log.Printf("[DEBUG] Checking existence of initiator: %#v", d)
	client := meta.(*element.Client)

	initiators := element.ListInitiatorRequest{}

	id := d.Id()
	s := make([]int, 1)
	convID, convErr := strconv.Atoi(id)

	if convErr != nil {
		return false, fmt.Errorf("id argument is required")
	}

	s[0] = convID
	initiators.Initiators = s

	res, err := client.ListInitiators(initiators)
	if err != nil {
		if err, ok := err.(*jsonrpc.ResponseError); ok {
			if err.Name == "xUnknown" {
				d.SetId("")
				return false, nil
			}
		}
		return false, err
	}

	if len(res.Initiators) != 1 {
		d.SetId("")
		return false, nil
	}

	return true, nil
}
