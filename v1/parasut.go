package parasut

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	tokenURL = "https://api.parasut.com/oauth/token"
	authURL  = "https://api.parasut.com/oauth/authorize"
	baseURL  = "https://api.parasut.com/v1"
)

type service struct {
	client *Client
}

// Client ...
type Client struct {
	client  *http.Client
	baseURL *url.URL

	common service

	Contacts      *ContactsService
	SalesInvoices *SalesInvoicesService

	companyID string
}

// GetAuthHTTPClient ...
func GetAuthHTTPClient(clientID, clientSecret, username, password string) (*http.Client, error) {

	cfg := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL,
		},
	}

	user := strings.TrimSpace(username)
	pass := strings.TrimSpace(password)

	token, err := cfg.PasswordCredentialsToken(oauth2.NoContext, user, pass)

	if err != nil {
		return nil, err
	}

	return cfg.Client(oauth2.NoContext, token), nil
}

// NewClient ...
func NewClient(httpClient *http.Client, companyID string) (*Client, error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	companyBaseURL, err := url.Parse(fmt.Sprintf("%s/%s", baseURL, companyID))

	if err != nil {
		return nil, err
	}

	c := &Client{
		client:  httpClient,
		baseURL: companyBaseURL,
	}

	c.common.client = c

	c.Contacts = (*ContactsService)(&c.common)
	c.SalesInvoices = (*SalesInvoicesService)(&c.common)

	c.companyID = companyID

	return c, nil
}

func addOptions(s string, opt interface{}) (string, error) {

	v := reflect.ValueOf(opt)

	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)

	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)

	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()

	return u.String(), nil
}

// NewRequest ...
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {

	endpoint := fmt.Sprintf("%s/%s", c.baseURL.String(), urlStr)

	// baseUrl'ye endpoint ekleniyor
	u, err := url.Parse(endpoint)

	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// http isteği gönderiliyor
	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// Do ...
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = checkResponse(resp)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, v)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func checkResponse(resp *http.Response) error {

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	data, err := ioutil.ReadAll(resp.Body)

	errorResponse := &ResponseError{}

	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	// http response'u errorResponse'a iliştir
	errorResponse.Response = resp

	switch {
	case len(errorResponse.Errors) > 0:
		return &CommonResponseError{
			Errors: errorResponse.Errors,
		}
	case len(errorResponse.Base) > 0:
		return &BaseResponseError{
			Base: errorResponse.Base,
		}
	case errorResponse.Type == "e_archive":
		return &EDocumentStatusResponseError{
			SalesInvoiceID: errorResponse.SalesInvoiceID,
			Type:           errorResponse.Type,
			Status:         errorResponse.Status,
			Errors:         errorResponse.Errors,
			PDF:            errorResponse.PDF,
		}
	default:
		return errorResponse
	}
}
