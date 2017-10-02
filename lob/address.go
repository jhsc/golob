package lob

import "time"
import "encoding/json"
import "bytes"
import "fmt"

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
