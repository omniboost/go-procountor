package procountor

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	scope                = "*"
	oauthStateString     = ""
	authorizationTimeout = 60 * time.Second
	tokenTimeout         = 5 * time.Second
)

type Oauth2Config struct {
	oauth2.Config
	APIKey string
}

func NewOauth2Config() *Oauth2Config {
	config := &Oauth2Config{
		Config: oauth2.Config{
			RedirectURL:  "",
			ClientID:     "",
			ClientSecret: "",
			Scopes:       []string{scope},
			Endpoint: oauth2.Endpoint{
				TokenURL: "https://api.procountor.com/api/oauth/token",
			},
		},
	}

	return config
}

// type Oauth2M2MConfig struct {
// 	clientcredentials.Config
// 	APIKey string
// }

// func NewOauth2M2MConfig() *Oauth2M2MConfig {
// 	config := &Oauth2M2MConfig{
// 		Config: clientcredentials.Config{
// 			ClientID:       "",
// 			ClientSecret:   "",
// 			Scopes:         []string{scope},
// 			TokenURL:       "https://api.procountor.com/api/oauth/token",
// 			EndpointParams: url.Values{"asdf": []string{"asdfaf"}},
// 		},
// 	}

// 	return config
// }

// Client returns an HTTP client using the provided token.

// The token will auto-refresh as necessary. The underlying
// HTTP transport will be obtained using the provided context.
// The returned client and its Transport should not be modified.
func (c *Oauth2Config) Client(ctx context.Context, token *oauth2.Token) *http.Client {
	if c.APIKey != "" {
		config := clientcredentials.Config{
			ClientID:       c.ClientID,
			ClientSecret:   c.ClientSecret,
			Scopes:         c.Scopes,
			TokenURL:       c.Endpoint.TokenURL,
			EndpointParams: url.Values{"api_key": []string{c.APIKey}},
		}
		return config.Client(ctx)
	}

	config := oauth2.Config{
		RedirectURL:  c.RedirectURL,
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		Scopes:       c.Scopes,
		Endpoint:     c.Endpoint,
	}
	return config.Client(ctx, token)
}
