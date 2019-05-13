package element

import (
	"encoding/json"

	"github.com/fatih/structs"
	"github.com/solidfire/terraform-provider-solidfire/solidfire/element"
)

type GetAccountByIDRequest struct {
	AccountID int `structs:"accountID"`
}

type GetAccountByIDResult struct {
	Account Account `json:"account"`
}

type Account struct {
	AccountID       int         `json:"accountID"`
	Attributes      interface{} `json:"attributes"`
	InitiatorSecret string      `json:"initiatorSecret"`
	Status          string      `json:"status"`
	TargetSecret    string      `json:"targetSecret"`
	Username        string      `json:"username"`
}

type CreateAccountRequest struct {
	Username        string      `structs:"username"`
	InitiatorSecret string      `structs:"initiatorSecret,omitempty"`
	TargetSecret    string      `structs:"targetSecret,omitempty"`
	Attributes      interface{} `structs:"attributes,omitempty"`
}

type CreateAccountResult struct {
	Account element.Account `json:"account"`
}

type ModifyAccountRequest struct {
	AccountID       int         `structs:"accountID"`
	InitiatorSecret string      `structs:"initiatorSecret,omitempty"`
	TargetSecret    string      `structs:"targetSecret,omitempty"`
	Attributes      interface{} `structs:"attributes,omitempty"`
	Username        string      `structs:"username,omitempty"`
}

type RemoveAccountRequest struct {
	AccountID int `structs:"accountID"`
}

func (c *Client) GetAccountByID(id int) (Account, error) {
	params := structs.Map(GetAccountByIDRequest{AccountID: id})

	response, err := c.CallAPIMethod("GetAccountByID", params)
	if err != nil {
		log.Print("GetAccountByID request failed")
		return Account{}, err
	}

	var result GetAccountByIDResult
	if err := json.Unmarshal([]byte(*response), &result); err != nil {
		log.Print("Failed to unmarshal response from GetAccountByID")
		return Account{}, err
	}

	return result.Account, nil
}
func (c *Client) CreateAccount(request CreateAccountRequest) (CreateAccountResult, error) {
	params := structs.Map(request)

	log.Printf("[DEBUG] Parameters: %v", params)

	response, err := c.CallAPIMethod("AddAccount", params)
	if err != nil {
		log.Print("CreateAccount request failed")
		return CreateAccountResult{}, err
	}

	var result CreateAccountResult
	if err := json.Unmarshal([]byte(*response), &result); err != nil {
		log.Print("Failed to unmarshall response from CreateAccount")
		return CreateAccountResult{}, err
	}
	return result, nil
}

func (c *Client) ModifyAccount(request ModifyAccountRequest) error {
	params := structs.Map(request)

	_, err := c.CallAPIMethod("ModifyAccount", params)
	if err != nil {
		log.Print("ModifyAccount request failed")
		return err
	}

	return nil
}

func (c *Client) RemoveAccount(request RemoveAccountRequest) error {
	params := structs.Map(request)

	_, err := c.CallAPIMethod("RemoveAccount", params)
	if err != nil {
		log.Print("DeleteAccount request failed")
		return err
	}

	return nil
}
