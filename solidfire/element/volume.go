package element

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/fatih/structs"
)

type ListVolumesRequest struct {
	Volumes               []int `structs:"volumeIDs"`
	IncludeVirtualVolumes bool  `structs:"includeVirtualVolumes"`
}

type ListVolumesByAccountsRequest struct {
	Accounts              []int `structs:"volumeIDs"`
	IncludeVirtualVolumes bool  `structs:"includeVirtualVolumes"`
}

type ListVolumesResult struct {
	Volumes []Volume `json:"volumes"`
}

type QualityOfService struct {
	MinIOPS   int `structs:"minIOPS"`
	MaxIOPS   int `structs:"maxIOPS"`
	BurstIOPS int `structs:"burstIOPS"`
}

type CreateVolumeRequest struct {
	Name       string            `structs:"name"`
	AccountID  int               `structs:"accountID"`
	TotalSize  int               `structs:"totalSize"`
	Enable512E bool              `structs:"enable512e"`
	Attributes map[string]string `structs:"attributes"`
	QOS        QualityOfService  `structs:"qos"`
}

type CreateVolumeResult struct {
	VolumeID int    `json:"volumeID"`
	Volume   Volume `json:"volume"`
}

type ModifyVolumeRequest struct {
	VolumeID   int               `structs:"volumeID"`
	AccountID  int               `structs:"accountID"`
	Attributes map[string]string `structs:"attributes"`
	QOS        QualityOfService  `structs:"qos"`
	TotalSize  int               `structs:"totalSize"`
}

type DeleteVolumeRequest struct {
	VolumeID int `structs:"volumeID"`
}

type Volume struct {
	Name               string            `json:"name"`
	VolumeID           int               `json:"volumeID"`
	Iqn                string            `json:"iqn"`
	Access             string            `json:"access"`
	AccountID          int               `json:"accountID"`
	Attributes         map[string]string `structs:"attributes"`
	BlockSize          int               `json:"blockSize"`
	Enable512e         bool              `json:"enable512e"`
	QOS                QualityOfService  `structs:"qos"`
	ScsiEUIDeviceID    string            `json:"scsiEUIDeviceID"`
	ScsiNAADeviceID    string            `json:"ScsiNAADeviceID"`
	Status             string            `json:"status"`
	TotalSize          int               `structs:"totalSize"`
	VirtualVolumeID    int               `json:"virtualVolumeID"`
	VolumeAccessGroups []int             `json:"volumeAccessGroups"`
}

func (c *Client) GetVolumeByID(id string) (Volume, error) {
	convID, err := strconv.Atoi(id)
	if err != nil {
		return Volume{}, err
	}

	volIDs := make([]int, 1)
	volIDs[0] = convID

	volumes, err := c.getVolumes(structs.Map(ListVolumesRequest{Volumes: volIDs}))
	if err != nil {
		return Volume{}, err
	}

	if len(volumes) != 1 {
		return Volume{}, errors.New(fmt.Sprintf("Expected one Volume to be found. Response contained %v results", len(volumes)))
	}

	return volumes[0], nil

}

func (c *Client) getVolumes(params map[string]interface{}) ([]Volume, error) {
	response, err := c.CallAPIMethod("ListVolumes", params)
	if err != nil {
		log.Print("ListVolumes request failed")
		return []Volume{}, err
	}

	var result ListVolumesResult
	if err := json.Unmarshal([]byte(*response), &result); err != nil {
		log.Print("Failed to unmarshal response from ListVolumes")
		return []Volume{}, err
	}
	return result.Volumes, nil
}

func (c *Client) GetVolumeByAccount(id string) ([]Volume, error) {
	convID, err := strconv.Atoi(id)
	if err != nil {
		return []Volume{}, err
	}

	accountIDs := make([]int, 1)
	accountIDs[0] = convID

	params := structs.Map(ListVolumesByAccountsRequest{Accounts: accountIDs})

	response, err := c.CallAPIMethod("ListVolumes", params)
	if err != nil {
		log.Print("ListVolumes request failed")
		return []Volume{}, err
	}

	var result ListVolumesResult
	if err := json.Unmarshal([]byte(*response), &result); err != nil {
		log.Print("Failed to unmarshal response from ListVolumes")
		return []Volume{}, err
	}

	return result.Volumes, nil
}

func (c *Client) CreateVolume(request CreateVolumeRequest) (CreateVolumeResult, error) {
	params := structs.Map(request)

	log.Printf("[DEBUG] Parameters: %v", params)

	response, err := c.CallAPIMethod("CreateVolume", params)
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

func (c *Client) UpdateVolume(request ModifyVolumeRequest) error {
	params := structs.Map(request)

	_, err := c.CallAPIMethod("ModifyVolume", params)
	if err != nil {
		log.Print("ModifyVolume request failed")
		return err
	}

	return nil
}

func (c *Client) ListVolumes(request ListVolumesRequest) (ListVolumesResult, error) {
	params := structs.Map(request)

	response, err := c.CallAPIMethod("ListVolumes", params)
	if err != nil {
		log.Print("ListVolumes request failed")
		return ListVolumesResult{}, err
	}

	var result ListVolumesResult
	if err := json.Unmarshal([]byte(*response), &result); err != nil {
		log.Print("Failed to unmarshall response from ListVolumes")
		return ListVolumesResult{}, err
	}

	return result, nil
}

func (c *Client) DeleteVolume(request DeleteVolumeRequest) error {
	params := structs.Map(request)

	_, err := c.CallAPIMethod("DeleteVolume", params)
	if err != nil {
		log.Print("DeleteVolume request failed")
		return err
	}

	return nil
}

func (c *Client) PurgeDeletedVolume(request DeleteVolumeRequest) error {
	params := structs.Map(request)

	_, err := c.CallAPIMethod("PurgeDeletedVolume", params)
	if err != nil {
		log.Print("PurgeDeletedVolume request failed")
		return err
	}

	return nil
}
