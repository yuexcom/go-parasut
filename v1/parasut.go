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
	baseURL     = "https://api.parasut.com/v1"
	authURL     = "https://api.parasut.com/oauth/authorize"
	tokenURL    = "https://api.parasut.com/oauth/token"
	callbackURL = "urn:ietf:wg:oauth:2.0:oob"
)

// Parasut ...
type Parasut struct {
	ClientID     string
	ClientSecret string
	ClientHTTP   *http.Client
	BaseURL      string
	CompanyID    string
	Token        *oauth2.Token
}

// New ...
func New(companyID, clientID, clientSecret string) *Parasut {

	baseURL := fmt.Sprintf("%s/%s", baseURL, companyID)

	return &Parasut{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		BaseURL:      baseURL,
		CompanyID:    companyID,
	}
}

// Auth ...
func (p *Parasut) Auth(username, password string) (string, error) {

	cfg := oauth2.Config{
		ClientID:     p.ClientID,
		ClientSecret: p.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL,
		},
	}

	user := strings.TrimSpace(username)
	pass := strings.TrimSpace(password)

	token, err := cfg.PasswordCredentialsToken(oauth2.NoContext, user, pass)

	if err != nil {
		return "", err
	}

	p.Token = token

	p.ClientHTTP = cfg.Client(oauth2.NoContext, token)

	return token.AccessToken, nil
}

func prepareURL(endpoint string, opts interface{}) (string, error) {

	optsValue := reflect.ValueOf(opts)

	// pointer değilse ve pointer dolu değilse direkt aynı şekilde geri gönder
	if optsValue.Kind() == reflect.Ptr && optsValue.IsNil() {
		return endpoint, nil
	}

	// endpoint string'ini url struct'ına dönüştür
	u, err := url.Parse(endpoint)

	if err != nil {
		return "", err
	}

	// struct değerlerini query srting'e dönüştür
	qs, err := query.Values(opts)

	if err != nil {
		return "", err
	}

	u.RawQuery = qs.Encode()

	return u.String(), nil
}

func (p *Parasut) newRequest(method, urlStr string, body interface{}) (*http.Request, error) {

	endpoint := fmt.Sprintf("%s/%s", p.BaseURL, urlStr)

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

func (p *Parasut) doRequest(req *http.Request, v interface{}) (*http.Response, error) {

	resp, err := p.ClientHTTP.Do(req)

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
