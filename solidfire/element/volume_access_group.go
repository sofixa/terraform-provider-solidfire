package element

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/fatih/structs"
	"github.com/solidfire/terraform-provider-solidfire/solidfire/element"
)

type ListVolumeAccessGroupsRequest struct {
	VolumeAccessGroups []int `structs:"volumeAccessGroups"`
}

type ListVolumeAccessGroupsResult struct {
	VolumeAccessGroups         []VolumeAccessGroup `json:"volumeAccessGroups"`
	VolumeAccessGroupsNotFound []int               `json:"volumeAccessGroupsNotFound"`
}

type VolumeAccessGroup struct {
	VolumeAccessGroupID int      `json:"volumeAccessGroupID"`
	Name                string   `json:"name"`
	Initiators          []string `json:"initiators"`
	Volumes             []int    `json:"volumes"`
	ID                  int      `json:"id"`
}

type CreateVolumeAccessGroupRequest struct {
	Name       string      `structs:"name"`
	Volumes    []int       `structs:"volumes"`
	Attributes interface{} `structs:"attributes"`
	ID         int         `structs:"id"`
}

type CreateVolumeAccessGroupResult struct {
	VolumeAccessGroupID int `json:"volumeAccessGroupID"`
	element.VolumeAccessGroup
}

type DeleteVolumeAccessGroupRequest struct {
	VolumeAccessGroupID    int  `structs:"volumeAccessGroupID"`
	DeleteOrphanInitiators bool `structs:"deleteOrphanInitiators"`
	Force                  bool `structs:"force"`
}

type ModifyVolumeAccessGroupRequest struct {
	VolumeAccessGroupID    int         `structs:"volumeAccessGroupID"`
	Name                   string      `structs:"name"`
	Attributes             interface{} `structs:"attributes"`
	DeleteOrphanInitiators bool        `structs:"deleteOrphanInitiators"`
	Volumes                []int       `structs:"volumes"`
}

func (c *Client) GetVolumeAccessGroupByID(id string) (VolumeAccessGroup, error) {
	convID, err := strconv.Atoi(id)
	if err != nil {
		return VolumeAccessGroup{}, err
	}

	vagIDs := make([]int, 1)
	vagIDs[0] = convID

	params := structs.Map(ListVolumeAccessGroupsRequest{VolumeAccessGroups: vagIDs})

	response, err := c.CallAPIMethod("ListVolumeAccessGroups", params)
	if err != nil {
		log.Print("ListVolumeAccessGroups request failed")
		return VolumeAccessGroup{}, err
	}

	var result ListVolumeAccessGroupsResult
	if err := json.Unmarshal([]byte(*response), &result); err != nil {
		log.Print("Failed to unmarshal respone from ListVolumeAccessGroups")
		return VolumeAccessGroup{}, err
	}

	if len(result.VolumeAccessGroupsNotFound) > 0 {
		return VolumeAccessGroup{}, errors.New(fmt.Sprintf("Unable to find Volume Access Groups with the ID of %v", result.VolumeAccessGroupsNotFound))
	}

	if len(result.VolumeAccessGroups) != 1 {
		return VolumeAccessGroup{}, errors.New(fmt.Sprintf("Expected one Volume Access Group to be found. Response contained %v results", len(result.VolumeAccessGroups)))
	}

	return result.VolumeAccessGroups[0], nil
}

func (c *Client) CreateVolumeAccessGroup(request CreateVolumeAccessGroupRequest) (CreateVolumeAccessGroupResult, error) {
	params := structs.Map(request)

	log.Printf("[DEBUG] Parameters: %v", params)

	response, err := c.CallAPIMethod("CreateVolumeAccessGroup", params)
	if err != nil {
		log.Print("CreateVolumeAccessGroup request failed")
		return CreateVolumeAccessGroupResult{}, err
	}

	var result CreateVolumeAccessGroupResult
	if err := json.Unmarshal([]byte(*response), &result); err != nil {
		log.Print("Failed to unmarshall response from CreateVolumeAccessGroup")
		return CreateVolumeAccessGroupResult{}, err
	}
	return result, nil
}

func (c *Client) ListVolumeAccessGroups(request ListVolumeAccessGroupsRequest) (ListVolumeAccessGroupsResult, error) {
	params := structs.Map(request)

	response, err := c.CallAPIMethod("ListVolumeAccessGroups", params)
	if err != nil {
		log.Print("ListVolumeAccessGroups request failed")
		return ListVolumeAccessGroupsResult{}, err
	}

	var result ListVolumeAccessGroupsResult
	if err := json.Unmarshal([]byte(*response), &result); err != nil {
		log.Print("Failed to unmarshall response from ListVolumeAccessGroups")
		return ListVolumeAccessGroupsResult{}, err
	}

	return result, nil
}

func (c *Client) ModifyVolumeAccessGroup(request ModifyVolumeAccessGroupRequest) error {
	params := structs.Map(request)

	_, err := c.CallAPIMethod("ModifyVolumeAccessGroup", params)
	if err != nil {
		log.Print("ModifyVolumeAccessGroup request failed")
		return err
	}

	return nil
}

func (c *Client) DeleteVolumeAccessGroup(request DeleteVolumeAccessGroupRequest) error {
	params := structs.Map(request)

	_, err := c.CallAPIMethod("DeleteVolumeAccessGroup", params)
	if err != nil {
		log.Print("DeleteVolumeAccessGroup request failed")
		return err
	}

	return nil
}
