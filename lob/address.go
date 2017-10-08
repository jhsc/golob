package lob

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// AddressService handles communication with the address related
// methods of the Lob API.
type AddressService struct {
	client *Client
}

// Address respresent Lob address object.
type Address struct {
	ID             string `json:"id,omitempty"`
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
	Company        string `json:"company,omitempty"`
	Phone          string `json:"phone,omitempty"`
	Email          string `json:"email,omitempty"`
	AddressLine1   string `json:"address_line1,omitempty"`
	AddressLine2   string `json:"address_line2,omitempty"`
	AddressCity    string `json:"address_city,omitempty"`
	AddressState   string `json:"address_state,omitempty"`
	AddressZip     string `json:"address_zip,omitempty"`
	AddressCountry string `json:"address_country,omitempty"`
	Metadata       struct {
	} `json:"metadata,omitempty"`
	DateCreated  time.Time `json:"date_created,omitempty"`
	DateModified time.Time `json:"date_modified,omitempty"`
	Object       string    `json:"object,omitempty"`
}

// AddressRequest represents a request to create/edit an address.
type AddressRequest struct {
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
	Company        string `json:"company,omitempty"`
	Phone          string `json:"phone,omitempty"`
	Email          string `json:"email,omitempty"`
	AddressLine1   string `json:"address_line1,omitempty"`
	AddressLine2   string `json:"address_line2,omitempty"`
	AddressCity    string `json:"address_city,omitempty"`
	AddressState   string `json:"address_state,omitempty"`
	AddressZip     string `json:"address_zip,omitempty"`
	AddressCountry string `json:"address_country,omitempty"`
	Metadata       struct {
	} `json:"metadata,omitempty"`
}

//DeleteAddressResponse ...
type DeleteAddressResponse struct {
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
}

// ListAddressesResponse ....
type ListAddressesResponse struct {
	Data        []Address `json:"data"`
	Object      string    `json:"object"`
	NextURL     string    `json:"next_url"`
	PreviousURL string    `json:"previous_url"`
	Count       int       `json:"count"`
}

// Create a new address object.
// Lob API docs: https://lob.com/docs/ruby#addresses_create
func (us *AddressService) Create(address *AddressRequest) (*AddressRequest, error) {
	body, err := json.Marshal(address)
	req, err := us.client.NewRequest("POST", "addresses", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	_, err = us.client.Do(req, address)
	if err != nil {
		return address, err
	}
	return address, err
}

// Get ...
func (us *AddressService) Get(id string) (*Address, error) {
	u := fmt.Sprintf("addresses/%v", id)
	req, err := us.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, err
	}

	address := new(Address)
	_, err = us.client.Do(req, address)
	return address, err
}

// Delete ...
func (us *AddressService) Delete(id string) error {
	u := fmt.Sprintf("addresses/%v", id)
	req, err := us.client.NewRequest("DELETE", u, nil)

	if err != nil {
		return err
	}

	delResp := new(DeleteAddressResponse)
	_, err = us.client.Do(req, delResp)

	if err != nil {
		return err
	}

	if !delResp.Deleted {
		return errors.New("Failed to delete address")
	}
	return nil
}

// List ...
func (us *AddressService) List(offset, limit int) (*ListAddressesResponse, error) {
	if offset <= 0 {
		offset = 0
	}
	if limit < 10 || limit > 100 {
		limit = 10
	}

	u := fmt.Sprintf("addresses/?limit=%v&offset=%v", limit, offset)
	req, err := us.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, err
	}

	address := new(ListAddressesResponse)
	_, err = us.client.Do(req, address)
	return address, err
}
