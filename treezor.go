//go:generate go run gen_accessors.go -v

package treezor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/pkg/errors"
)

const (
	defaultStagingBaseURL             = "https://sandbox.treezor.com/v1/index.php/"
	defaultStagingBaseURLWithoutIndex = "https://sandbox.treezor.com/v1/"
	defaultBaseURL                    = "https://treezor.com/v1/index.php/"
	defaultBaseURLWithoutIndex        = "https://treezor.com/v1/"
	userAgent                         = "go-treezor"
)

// A Client manages communication with the Treezor API.
type Client struct {
	client *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. Defaults to the public Treezor API. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL
	// Base URL without index used for endpoints that doesn't have index.php prefix.
	// For example : https://sandbox.treezor.com/v1/index.php/users/{id}/kycliveness
	BaseURLWithoutIndex *url.URL

	// User agent used when communicating with the Treezor API.
	UserAgent string

	common          service // Reuse a single struct instead of allocating one for each service on the heap.
	User            *UserService
	Wallet          *WalletService
	Card            *CardService
	CardTransaction *CardTransactionService
	Balance         *BalanceService
	Document        *DocumentService
	Beneficiary     *BeneficiaryService
	Transfer        *TransferService
	Payin           *PayinService
	Payout          *PayoutService
	Hearthbeat      *HearthbeatService
	TaxResidences   *TaxResidencesService
}

type service struct {
	client *Client
}

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	Page int `url:"pageNumber,omitempty"`

	// For paginated result sets, the number of results to include per page.
	PerPage int `url:"pageCount,omitempty"`

	// For paginated result sets, the resource element you want to sort the list with.
	SortBy string `url:"sortBy,omitempty"`

	// For paginated result sets, The order you want to sort the list.
	// Possible values: DESC and ASC.
	SortOrder string `url:"sortOrder,omitempty"`

	// For paginated result sets, the creation date from which you want to filter the request result.
	// Format: YYYY-MM-DD HH:MM:SS
	CreatedFrom string `url:"createdDateFrom,omitempty"`

	// For paginated result sets, the creation date up to which you want to filter the request result.
	// Format: YYYY-MM-DD HH:MM:SS
	CreatedTo string `url:"createdDateFrom,omitempty"`

	// For paginated result sets, the modification date from which you want to filter the request result.
	// Format: YYYY-MM-DD HH:MM:SS
	UpdatedFrom string `url:"createdDateFrom,omitempty"`

	// For paginated result sets, the modification date up to which you want to filter the request result.
	// Format: YYYY-MM-DD HH:MM:SS
	UpdatedTo string `url:"createdDateFrom,omitempty"`
}

// addOptions adds the parameters in opt as URL query parameters to s. opt
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, errors.WithStack(err)
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, errors.WithStack(err)
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// NewClient returns a new Treezor API client. If a nil httpClient is
// provided, http.DefaultClient will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client, isProduction bool) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultStagingBaseURL)
	baseURLWithoutIndex, _ := url.Parse(defaultStagingBaseURLWithoutIndex)
	if isProduction {
		baseURL, _ = url.Parse(defaultBaseURL)
		baseURLWithoutIndex, _ = url.Parse(defaultBaseURLWithoutIndex)
	}

	c := &Client{client: httpClient, BaseURL: baseURL, BaseURLWithoutIndex: baseURLWithoutIndex, UserAgent: userAgent}
	c.common.client = c
	c.User = (*UserService)(&c.common)
	c.Wallet = (*WalletService)(&c.common)
	c.Card = (*CardService)(&c.common)
	c.CardTransaction = (*CardTransactionService)(&c.common)
	c.Balance = (*BalanceService)(&c.common)
	c.Document = (*DocumentService)(&c.common)
	c.Beneficiary = (*BeneficiaryService)(&c.common)
	c.Transfer = (*TransferService)(&c.common)
	c.Payin = (*PayinService)(&c.common)
	c.Payout = (*PayoutService)(&c.common)
	c.Hearthbeat = (*HearthbeatService)(&c.common)
	c.TaxResidences = (*TaxResidencesService)(&c.common)
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// NewRequestWithoutIndex do the same as NewRequest but without /index.php/
func (c *Client) NewRequestWithoutIndex(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURLWithoutIndex.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURLWithoutIndex)
	}
	u, err := c.BaseURLWithoutIndex.Parse(urlStr)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
//
// The provided ctx must be non-nil. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// If the error type is *url.Error, sanitize its URL before returning.
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = sanitizeURL(url).String()
				return nil, e
			}
		}

		return nil, errors.WithStack(err)
	}

	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return resp, errors.WithStack(err)
}

// sanitizeURL redacts the client_secret parameter from the URL which may be
// exposed to the user.
func sanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("client_secret")) > 0 {
		params.Set("client_secret", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}

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

	return t.transport().RoundTrip(req2)
}

// Client returns an *http.Client that makes requests that are authenticated
// using HTTP Basic Authentication.
func (t *BearerAuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *BearerAuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return cleanhttp.DefaultTransport()
}

// Access contains global keys to all Treezor objects.
type Access struct {
	IdempotencyKey *string `json:"accessTag,omitempty"`
	UserID         *string `json:"accessUserId,omitempty"`
	UserIP         *string `json:"accessUserIp,omitempty"`
}

// Origin represents who made the request.
type Origin string

const (
	// OperatorOrigin represents the support
	OperatorOrigin Origin = "OPERATOR"
	// UserOrigin represents the end user.
	UserOrigin Origin = "USER"
)

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// Float64 is a helper routine that allocates a new float64 value
// to store v and returns a pointer to it.
func Float64(v float64) *float64 { return &v }
