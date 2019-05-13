package solidfire

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sofixa/terraform-provider-solidfire/solidfire/element"
	"github.com/sofixa/terraform-provider-solidfire/solidfire/element/jsonrpc"
)

func resourceSolidFireVolumeAccessGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSolidFireVolumeAccessGroupCreate,
		Read:   resourceSolidFireVolumeAccessGroupRead,
		Update: resourceSolidFireVolumeAccessGroupUpdate,
		Delete: resourceSolidFireVolumeAccessGroupDelete,
		Exists: resourceSolidFireVolumeAccessGroupExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"volumes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
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

func resourceSolidFireVolumeAccessGroupCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Creating volume access group: %#v", d)
	client := meta.(*element.Client)

	vag := element.CreateVolumeAccessGroupRequest{}

	if v, ok := d.GetOk("name"); ok {
		vag.Name = v.(string)
	} else {
		return fmt.Errorf("name argument is required")
	}

	if raw, ok := d.GetOk("volumes"); ok {
		for _, v := range raw.([]interface{}) {
			vag.Volumes = append(vag.Volumes, v.(int))
		}
	}

	resp, err := client.CreateVolumeAccessGroup(vag)
	if err != nil {
		log.Print("Error creating volume access group")
		return err
	}

	d.SetId(fmt.Sprintf("%v", resp.VolumeAccessGroupID))
	log.Printf("[DEBUG] Created volume access group: %v %v", vag.Name, resp.VolumeAccessGroupID)

	return resourceSolidFireVolumeAccessGroupRead(d, meta)
}

func resourceSolidFireVolumeAccessGroupRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Reading volume access group: %#v", d)
	client := meta.(*element.Client)

	vags := element.ListVolumeAccessGroupsRequest{}

	id := d.Id()
	s := make([]int, 1)
	convID, convErr := strconv.Atoi(id)

	if convErr != nil {
		return fmt.Errorf("id argument is required")
	}

	s[0] = convID
	vags.VolumeAccessGroups = s

	res, err := client.ListVolumeAccessGroups(vags)
	if err != nil {
		return err
	}

	if len(res.VolumeAccessGroupsNotFound) > 0 {
		return fmt.Errorf("Unable to find Volume Access Groups with the ID of %v", res.VolumeAccessGroupsNotFound)
	}

	if len(res.VolumeAccessGroups) != 1 {
		return fmt.Errorf("Expected one Volume Access Group to be found. Response contained %v results", len(res.VolumeAccessGroups))
	}

	d.Set("name", res.VolumeAccessGroups[0].Name)
	d.Set("volumes", res.VolumeAccessGroups[0].Volumes)

	return nil
}

func resourceSolidFireVolumeAccessGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Updating volume access group %#v", d)
	client := meta.(*element.Client)

	vag := element.ModifyVolumeAccessGroupRequest{}

	id := d.Id()
	convID, convErr := strconv.Atoi(id)

	if convErr != nil {
		return fmt.Errorf("id argument is required")
	}
	vag.VolumeAccessGroupID = convID

	if v, ok := d.GetOk("name"); ok {
		vag.Name = v.(string)

	} else {
		return fmt.Errorf("name argument is required during update")
	}

	if raw, ok := d.GetOk("volumes"); ok {
		for _, v := range raw.([]interface{}) {
			vag.Volumes = append(vag.Volumes, v.(int))
		}
	}

	err := client.ModifyVolumeAccessGroup(vag)
	if err != nil {
		return err
	}

	return nil
}

func resourceSolidFireVolumeAccessGroupDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Deleting volume access group: %#v", d)
	client := meta.(*element.Client)

	vag := element.DeleteVolumeAccessGroupRequest{}

	id := d.Id()
	convID, convErr := strconv.Atoi(id)

	if convErr != nil {
		return fmt.Errorf("id argument is required")
	}
	vag.VolumeAccessGroupID = convID

	err := client.DeleteVolumeAccessGroup(vag)
	if err != nil {
		return err
	}

	return nil
}

func resourceSolidFireVolumeAccessGroupExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	log.Printf("[DEBUG] Checking existence of volume access group: %#v", d)
	client := meta.(*element.Client)

	vags := element.ListVolumeAccessGroupsRequest{}

	id := d.Id()
	s := make([]int, 1)
	convID, convErr := strconv.Atoi(id)

	if convErr != nil {
		return false, fmt.Errorf("id argument is required")
	}

	s[0] = convID
	vags.VolumeAccessGroups = s

	res, err := client.ListVolumeAccessGroups(vags)
	if err != nil {
		if err, ok := err.(*jsonrpc.ResponseError); ok {
			if err.Name == "xUnknown" {
				d.SetId("")
				return false, nil
			}
		}
		return false, err
	}

	if len(res.VolumeAccessGroupsNotFound) > 0 {
		d.SetId("")
		return false, nil
	}

	if len(res.VolumeAccessGroups) != 1 {
		d.SetId("")
		return false, nil
	}

	return true, nil
}
