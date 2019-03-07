package parasut

import (
	"net/http"

	"golang.org/x/oauth2"
)

const (
	baseURL     = "https://api.parasut.com/v4"
	authURL     = "https://api.parasut.com/oauth/authorize"
	tokenURL    = "https://api.parasut.com/oauth/token"
	callbackURL = "urn:ietf:wg:oauth:2.0:oob"
)

// Parasut ...
type Parasut struct {
	ClientID     string
	ClientSecret string
	Token        *oauth2.Token
	ClientHTTP   *http.Client
}

// New ...
func New(clientID, clientSecret string) *Parasut {

	return &Parasut{
		ClientID:     clientID,
		ClientSecret: clientSecret,
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

	token, err := cfg.PasswordCredentialsToken(oauth2.NoContext, username, password)

	if err != nil {
		return "", err
	}

	p.Token = token

	p.ClientHTTP = cfg.Client(oauth2.NoContext, token)

	return token.AccessToken, nil
}
