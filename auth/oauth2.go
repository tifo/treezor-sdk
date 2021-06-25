package auth

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// NewOAuth2Client
func NewOAuth2Client(clientID, clientSecret, tokenURL string, httpClient *http.Client) *http.Client {
	oauthCtx := context.WithValue(context.Background(), oauth2.HTTPClient, httpClient)
	oauthCfg := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
	}
	return oauth2.NewClient(oauthCtx, oauthCfg.TokenSource(oauthCtx))
}
