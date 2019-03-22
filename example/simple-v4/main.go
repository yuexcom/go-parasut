package main

import (
	"encoding/json"
	"fmt"

	"github.com/yuexcom/go-parasut/v4"
)

const (
	clientID     = "CLIENT_ID"
	clientSecret = "CLIENT_SECRET"
	username     = "USERNAME"
	password     = "PASSWORD"
	companyID    = "COMPANY_ID"
)

func main() {

	httpClient, err := parasut.GetAuthHTTPClient(clientID, clientSecret, username, password)

	if err != nil {
		panic(err)
	}

	client, err := parasut.NewClient(httpClient, companyID)

	if err != nil {
		panic(err)
	}

	contacts, _, err := client.Contacts.List(&parasut.ContactListOptions{
		Filter: parasut.ContactFilterOptions{
			TaxNumber: "888888888",
		},
		Include: "category,contact_people",
	})

	if err != nil {
		panic(err)
	}

	prettyPrint(contacts)
}

func prettyPrint(v interface{}) {

	bytes, _ := json.MarshalIndent(v, "", "  ")

	fmt.Printf("%s", bytes)
}
