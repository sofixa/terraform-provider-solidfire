package solidfire

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sofixa/terraform-provider-solidfire/solidfire/element"
	"github.com/sofixa/terraform-provider-solidfire/solidfire/element/jsonrpc"
)

func resourceSolidFireVolume() *schema.Resource {
	return &schema.Resource{
		Create: resourceSolidFireVolumeCreate,
		Read:   resourceSolidFireVolumeRead,
		Update: resourceSolidFireVolumeUpdate,
		Delete: resourceSolidFireVolumeDelete,
		Exists: resourceSolidFireVolumeExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"volume_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"iqn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"access": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "readWrite",
			},
			"account_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"attributes": schemaAttributes(),
			"block_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"enable512e": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"min_iops": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_iops": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"burst_iops": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scsi_eui_device_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scsi_naa_device_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_volume_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"purge_on_delete": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceSolidFireVolumeCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Creating volume: %#v", d)
	client := meta.(*element.Client)

	volume := element.CreateVolumeRequest{}

	if v, ok := d.GetOk("name"); ok {
		volume.Name = v.(string)
	} else {
		return fmt.Errorf("name argument is required")
	}

	if v, ok := d.GetOk("account_id"); ok {
		volume.AccountID = v.(int)
	} else {
		return fmt.Errorf("account_id argument is required")
	}

	if v, ok := d.GetOk("total_size"); ok {
		volume.TotalSize = v.(int)
	} else {
		return fmt.Errorf("total_size argument is required")
	}

	if v, ok := d.GetOk("enable512e"); ok {
		volume.Enable512E = v.(bool)
	} else {
		return fmt.Errorf("enable512e argument is required")
	}

	if v, ok := d.GetOk("min_iops"); ok {
		volume.QOS.MinIOPS = v.(int)
	}

	if v, ok := d.GetOk("max_iops"); ok {
		volume.QOS.MaxIOPS = v.(int)
	}

	if v, ok := d.GetOk("burst_iops"); ok {
		volume.QOS.BurstIOPS = v.(int)
	}
	volume.Attributes = make(map[string]string)
	if v, ok := d.GetOk("attributes"); ok {
		for key, val := range v.(map[string]interface{}) {
			volume.Attributes[key] = val.(string)
		}
	}

	if v, ok := d.GetOk("attributes"); ok {
		volume.Attributes = v.(map[string]string)
	}

	resp, err := client.CreateVolume(volume)
	if err != nil {
		log.Print("Error creating volume")
		return err
	}

	d.SetId(fmt.Sprintf("%v", resp.VolumeID))
	d.Set("iqn", resp.Volume.Iqn)
	log.Printf("[DEBUG] Created volume: %v %v", volume.Name, resp.VolumeID)

	return resourceSolidFireVolumeRead(d, meta)
}

func resourceSolidFireVolumeRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Reading volume: %#v", d)
	client := meta.(*element.Client)

	volume, err := client.GetVolumeByID(d.Id())
	if err != nil {
		return err
	}

	d.Set("name", volume.Name)
	d.Set("volume_id", volume.VolumeID)
	d.Set("iqn", volume.Iqn)
	d.Set("access", volume.Access)
	d.Set("account_id", volume.AccountID)

	d.Set("attributes", volume.Attributes)
	for key, val := range volume.Attributes {
		d.Set("attributes."+key, val)
	}

	d.Set("block_size", volume.BlockSize)
	d.Set("enable512e", volume.Enable512e)
	d.Set("min_iops", volume.QOS.MinIOPS)
	d.Set("max_iops", volume.QOS.MaxIOPS)
	d.Set("burst_iops", volume.QOS.BurstIOPS)
	d.Set("scsi_eui_device_id", volume.ScsiEUIDeviceID)
	d.Set("scsi_naa_device_id", volume.ScsiNAADeviceID)
	d.Set("status", volume.Status)
	d.Set("total_size", volume.TotalSize)
	d.Set("virtual_volumeID", volume.VirtualVolumeID)

	log.Printf("[DEBUG] %s: Read complete", volume.Name)
	return nil

}

func resourceSolidFireVolumeUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Updating volume %#v", d)
	client := meta.(*element.Client)

	volume := element.ModifyVolumeRequest{}

	id := d.Id()
	convID, convErr := strconv.Atoi(id)

	if convErr != nil {
		return fmt.Errorf("id argument is required")
	}
	volume.VolumeID = convID

	if v, ok := d.GetOk("account_id"); ok {
		volume.AccountID = v.(int)
	} else {
		return fmt.Errorf("account_id argument is required")
	}

	if v, ok := d.GetOk("total_size"); ok {
		volume.TotalSize = v.(int)
	} else {
		return fmt.Errorf("total_size argument is required")
	}

	if v, ok := d.GetOk("min_iops"); ok {
		volume.QOS.MinIOPS = v.(int)
	}

	if v, ok := d.GetOk("max_iops"); ok {
		volume.QOS.MaxIOPS = v.(int)
	}

	if v, ok := d.GetOk("burst_iops"); ok {
		volume.QOS.BurstIOPS = v.(int)
	}

	volume.Attributes = make(map[string]string)
	if v, ok := d.GetOk("attributes"); ok {
		for key, val := range v.(map[string]interface{}) {
			volume.Attributes[key] = val.(string)
		}
	}

	err := client.UpdateVolume(volume)
	if err != nil {
		return err
	}

	return nil
}

func resourceSolidFireVolumeDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Deleting volume: %#v", d)
	client := meta.(*element.Client)

	volume := element.DeleteVolumeRequest{}

	id := d.Id()
	convID, convErr := strconv.Atoi(id)

	if convErr != nil {
		return fmt.Errorf("id argument is required")
	}
	volume.VolumeID = convID

	deleteErr := client.DeleteVolume(volume)
	if deleteErr != nil {
		return deleteErr
	}
	// only purge the deleted volume if purge_on_delete is set to true
	if v, ok := d.GetOk("purge_on_delete"); ok {
		if v.(bool) {
			purgeErr := client.PurgeDeletedVolume(volume)
			if purgeErr != nil {
				return purgeErr
			}
		}
	}
	return nil
}

func resourceSolidFireVolumeExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	log.Printf("[DEBUG] Checking existence of volume: %#v", d)
	client := meta.(*element.Client)

	volumes := element.ListVolumesRequest{}

	id := d.Id()
	s := make([]int, 1)
	convID, convErr := strconv.Atoi(id)

	if convErr != nil {
		return false, fmt.Errorf("id argument is required")
	}

	s[0] = convID
	volumes.Volumes = s

	res, err := client.ListVolumes(volumes)
	if err != nil {
		if err, ok := err.(*jsonrpc.ResponseError); ok {
			if err.Name == "xUnknown" {
				d.SetId("")
				return false, nil
			}
		}
		return false, err
	}

	if len(res.Volumes) != 1 {
		d.SetId("")
		return false, nil
	}

	return true, nil
}
