package main

import (
	"fmt"

	"github.com/yuexcom/go-parasut/v1"
)

func main() {

	prst := parasut.New("COMPANY_ID", "CLIENT_ID", "CLIENT_SECRET")

	err := prst.Auth("username", "password")

	if err != nil {
		panic(err)
	}

	contactList, err := prst.GetContacts(&parasut.ListOptions{
		TaxNumber: "1234567890",
	})

	fmt.Printf("%+v", contactList)
}
