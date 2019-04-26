package solidfire

import (
	"fmt"
	"log"
	"strconv"

	"encoding/json"

	"github.com/fatih/structs"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sofixa/terraform-provider-solidfire/solidfire/element"
	"github.com/sofixa/terraform-provider-solidfire/solidfire/element/jsonrpc"
)

type CreateVolumeRequest struct {
	Name       string           `structs:"name"`
	AccountID  int              `structs:"accountID"`
	TotalSize  int              `structs:"totalSize"`
	Enable512E bool             `structs:"enable512e"`
	Attributes interface{}      `structs:"attributes"`
	QOS        QualityOfService `structs:"qos"`
}

type CreateVolumeResult struct {
	VolumeID int            `json:"volumeID"`
	Volume   element.Volume `json:"volume"`
}

type DeleteVolumeRequest struct {
	VolumeID int `structs:"volumeID"`
}

type ModifyVolumeRequest struct {
	VolumeID   int              `structs:"volumeID"`
	AccountID  int              `structs:"accountID"`
	Attributes interface{}      `structs:"attributes"`
	QOS        QualityOfService `structs:"qos"`
	TotalSize  int              `structs:"totalSize"`
}

type QualityOfService struct {
	MinIOPS   int `structs:"minIOPS"`
	MaxIOPS   int `structs:"maxIOPS"`
	BurstIOPS int `structs:"burstIOPS"`
}

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
			},
			"account_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
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
		},
	}
}

func resourceSolidFireVolumeCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("Creating volume: %#v", d)
	client := meta.(*element.Client)

	volume := CreateVolumeRequest{}

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

	resp, err := createVolume(client, volume)
	if err != nil {
		log.Print("Error creating volume")
		return err
	}

	d.SetId(fmt.Sprintf("%v", resp.VolumeID))
	d.Set("iqn", resp.Volume.Iqn)
	log.Printf("Created volume: %v %v", volume.Name, resp.VolumeID)

	return resourceSolidFireVolumeRead(d, meta)
}

func createVolume(client *element.Client, request CreateVolumeRequest) (CreateVolumeResult, error) {
	params := structs.Map(request)

	log.Printf("Parameters: %v", params)

	response, err := client.CallAPIMethod("CreateVolume", params)
	if err != nil {
		log.Print("CreateVolume request failed")
		return CreateVolumeResult{}, err
	}

	var result CreateVolumeResult
	if err := json.Unmarshal([]byte(*response), &result); err != nil {
		log.Print("Failed to unmarshall response from CreateVolume")
		return CreateVolumeResult{}, err
	}
	return result, nil
}

func resourceSolidFireVolumeRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("Reading volume: %#v", d)
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

func listVolumes(client *element.Client, request element.ListVolumesRequest) (element.ListVolumesResult, error) {
	params := structs.Map(request)

	response, err := client.CallAPIMethod("ListVolumes", params)
	if err != nil {
		log.Print("ListVolumes request failed")
		return element.ListVolumesResult{}, err
	}

	var result element.ListVolumesResult
	if err := json.Unmarshal([]byte(*response), &result); err != nil {
		log.Print("Failed to unmarshall response from ListVolumes")
		return element.ListVolumesResult{}, err
	}

	return result, nil
}

func resourceSolidFireVolumeUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("Updating volume %#v", d)
	client := meta.(*element.Client)

	volume := ModifyVolumeRequest{}

	id := d.Id()
	convID, convErr := strconv.Atoi(id)

	if convErr != nil {
		return fmt.Errorf("id argument is required")
	}
	volume.VolumeID = convID

	err := updateVolume(client, volume)
	if err != nil {
		return err
	}

	return nil
}

func updateVolume(client *element.Client, request ModifyVolumeRequest) error {
	params := structs.Map(request)

	_, err := client.CallAPIMethod("ModifyVolume", params)
	if err != nil {
		log.Print("ModifyVolume request failed")
		return err
	}

	return nil
}

func resourceSolidFireVolumeDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("Deleting volume: %#v", d)
	client := meta.(*element.Client)

	volume := DeleteVolumeRequest{}

	id := d.Id()
	convID, convErr := strconv.Atoi(id)

	if convErr != nil {
		return fmt.Errorf("id argument is required")
	}
	volume.VolumeID = convID

	deleteErr := deleteVolume(client, volume)
	if deleteErr != nil {
		return deleteErr
	}

	purgeErr := purgeDeletedVolume(client, volume)
	if purgeErr != nil {
		return purgeErr
	}

	return nil
}

func deleteVolume(client *element.Client, request DeleteVolumeRequest) error {
	params := structs.Map(request)

	_, err := client.CallAPIMethod("DeleteVolume", params)
	if err != nil {
		log.Print("DeleteVolume request failed")
		return err
	}

	return nil
}

func purgeDeletedVolume(client *element.Client, request DeleteVolumeRequest) error {
	params := structs.Map(request)

	_, err := client.CallAPIMethod("PurgeDeletedVolume", params)
	if err != nil {
		log.Print("PurgeDeletedVolume request failed")
		return err
	}

	return nil
}

func resourceSolidFireVolumeExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	log.Printf("Checking existence of volume: %#v", d)
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

	res, err := listVolumes(client, volumes)
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
