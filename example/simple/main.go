package main

import (
	"fmt"

	"github.com/yuexcom/go-parasut/v1"
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

	contacts, _, err := client.Contacts.List(&parasut.ListOptions{
		TaxNumber: "1234567890",
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", contacts)
}
