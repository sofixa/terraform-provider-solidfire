package solidfire

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/sofixa/terraform-provider-solidfire/solidfire/element"
)

func dataSourceSolidFireVolume() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSolidFireVolumeRead,
		Schema: map[string]*schema.Schema{
			"account_filter": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "The ID of the account to filter on",
			},
			"name_filter": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "A regular expression to filter the name of volumes",
				ValidateFunc: validation.ValidateRegexp,
			},

			"name": {
				Type:     schema.TypeString,
				Computed: true,
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
				Computed: true,
			},
			"attributes": schemaDataSourceAttributes(),
			"block_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"enable512e": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"min_iops": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_iops": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"burst_iops": {
				Type:     schema.TypeInt,
				Computed: true,
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
				Computed: true,
			},
		},
	}
}

func dataSourceSolidFireVolumeRead(d *schema.ResourceData, meta interface{}) error {
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
	d.Set("virtual_volume_id", volume.VirtualVolumeID)

	log.Printf("[DEBUG] %s: Read complete", volume.Name)
	return nil
}
