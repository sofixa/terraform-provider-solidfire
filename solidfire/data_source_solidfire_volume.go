package solidfire

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
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
			"attributes": {
				Type:     schema.TypeList,
				Computed: true,
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
	return nil
}
