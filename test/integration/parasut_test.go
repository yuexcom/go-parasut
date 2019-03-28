package integration

import (
	"fmt"
	"os"

	"github.com/yuexcom/go-parasut/parasut"
)

var (
	err    error
	client *parasut.Client
)

const (
	token     = "TOKEN"
	companyID = "COMPANY_ID"
)

func init() {

	envs := []string{token, companyID}

	for _, env := range envs {
		e := os.Getenv(env)
		if e == "" {
			print(fmt.Sprintf("missing environment variable: %s!\n", env))
			os.Exit(1)
		}
	}

	httpClient := parasut.GetHTTPClientWithToken(os.Getenv(token))

	client, err = parasut.NewClient(httpClient, os.Getenv(companyID))

	if err != nil {
		panic(err)
	}
}
