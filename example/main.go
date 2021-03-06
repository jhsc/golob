package main

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/jhsc/golob/lob"
)

func main() {
	client := lob.NewClient("KEY", nil)

	testAddress := &lob.AddressRequest{
		Name:           "Lobster Test",
		Email:          "lobtest@example.com",
		Phone:          "5555555555",
		AddressLine1:   "1005 W Burnside St",
		AddressCity:    "Portland",
		AddressState:   "OR",
		AddressZip:     "97209",
		AddressCountry: "US",
	}

	spew.Dump(testAddress)

	address, err := client.Address.Create(testAddress)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}

	spew.Dump(address)
}
