package auth

import (
	"net/http"
)

// BearerAuthTransport is an http.RoundTripper that authenticates all requests
// using HTTP Bearer Authentication with the provided access token.
type BearerAuthTransport struct {
	AccessToken string // Treezor AccessToken
	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

// RoundTrip implements the RoundTripper interface.
func (t *BearerAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// To set extra headers, we must make a copy of the Request so
	// that we don't modify the Request we were given. This is required by the
	// specification of http.RoundTripper.
	//
	// Since we are going to modify only req.Header here, we only need a deep copy
	// of req.Header.
	req2 := new(http.Request)
	*req2 = *req
	req2.Header = make(http.Header, len(req.Header))
	for k, s := range req.Header {
		req2.Header[k] = append([]string(nil), s...)
	}

	authorization := "Bearer " + t.AccessToken

	req2.Header.Set("Authorization", authorization)

	return t.Transport.RoundTrip(req2)
}

// NewBearerAuthClient
func NewBearerAuthClient(accessToken string, httpClient *http.Client) *http.Client {
	client := http.DefaultClient
	transport := http.DefaultTransport
	if httpClient != nil {
		client = httpClient
		if httpClient.Transport != nil {
			transport = httpClient.Transport
		}
	}
	return &http.Client{
		Transport: &BearerAuthTransport{
			AccessToken: accessToken,
			Transport:   transport,
		},
		CheckRedirect: client.CheckRedirect,
		Jar:           client.Jar,
		Timeout:       client.Timeout,
	}
}
