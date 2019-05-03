---
layout: "solidfire"
page_title: "SolidFire: solidfire_volume"
sidebar_current: "docs-solidfire-resource-volume"
description: |-
  Provides a SolidFire cluster volume resource. This can be used to create a new (empty) volume on the cluster. As soon
  as the volume creation is complete, the volume is available for connection via iSCSI.
---

# solidfire\_volume

Provides a SolidFire cluster volume resource. This can be used to create a new (empty) volume on the cluster. As soon
as the volume creation is complete, the volume is available for connection via iSCSI.

## Example Usages

**Create SolidFire cluster volume:**

```
resource "solidfire_volume" "main-volume" {
  name        = "main-volume"
  account_id  = "1"
  total_size  = 10000000000
  enable512e  = true
  min_iops    = 50
  max_iops    = 10000
  burst_iops  = 10000
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the SolidFire volume.
* `access` - (Optional) The type of access allowed for the volume. Options: readOnly, readWrite, locked, replicationTarget. Default: readWrite.
* `account_id` - (Required) The unique identifier of the SolidFire account owner.
* `attributes` - (Optional) List of name-value pairs of volume attributes.
* `total_size` - (Required) The total size of the volume, in bytes. Size is rounded up to the nearest 1MB size.
* `enable512e` - (Required) Whether to enable 512-byte sector emulation. The setting needs to be enabled if using VMWare.
* `min_iops` - (Optional) The minimum initial quality of service.
* `max_iops` - (Optional) The maximum initial quality of service.
* `burst_iops` - (Optional) The burst initial quality of service.
  
## Attributes Reference

The following attributes are exported in addition to the arguments listed above:

* `volume_id` - The unique identifier for the volume.
* `iqn` - The iSCSI Qualified Name of the volume.
* `block_size` - The size of blocks on the volume.
* `scsi_eui_device_id` - Globally unique SCSI device identifier for the volume in EUI-64 based 16-byte format.
* `scsi_naa_device_id` - Globally unique SCSI device identifier for the volume in NAA IEEE Registered Extended format.
* `status` - The current status of the volume.
* `virtual_volume_id` - The current sThe unique virtual volume ID associated with the volume, if any.





