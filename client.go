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
	"path"

	"github.com/pkg/errors"
)

const (
	defaultStagingBaseURL    = "https://sandbox.treezor.com/v1/"
	defaultProductionBaseURL = "https://treezor.com/v1/"
	userAgent                = "go-treezor/"
)

type ConnectClient struct {
	API     *Client
	BaseURL *url.URL
}

// A Client manages communication with the Treezor API.
type Client struct {
	client *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. Defaults to the public Treezor API. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL

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

func newAPIClient(httpClient *http.Client, apiBaseURL *url.URL) *Client {
	c := &Client{client: httpClient, BaseURL: apiBaseURL, UserAgent: userAgent}
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

// NewClient returns a new Treezor API client. If a nil httpClient is
// provided, http.DefaultClient will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client, isProduction bool) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultStagingBaseURL)
	if isProduction {
		baseURL, _ = url.Parse(defaultProductionBaseURL)
	}
	return newAPIClient(httpClient, baseURL)
}

// NewConnectClient returns a new Treezor API client using the Base URL
// passed as parameters. If a nil httpClient is provided, http.DefaultClient
// will be used. To use API methods which require authentication, provide an
// http.Client that will perform the authentication for you (such as that
// provided by the golang.org/x/oauth2 library).
func NewConnectClient(httpClient *http.Client, connectURL string) *ConnectClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(connectURL)
	apiBaseURL := &url.URL{
		Scheme: baseURL.Scheme,
		Host:   baseURL.Host,
		Path:   path.Join(baseURL.Path, "/v1/"),
	}
	return &ConnectClient{
		API:     newAPIClient(httpClient, apiBaseURL),
		BaseURL: baseURL,
	}
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {

	path := path.Join(c.BaseURL.Path, urlStr)
	u, err := c.BaseURL.Parse(path)
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
		_, _ = io.CopyN(ioutil.Discard, resp.Body, 512)
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
			_, _ = io.Copy(w, resp.Body)
		} else {
			buf := &bytes.Buffer{}
			_, _ = buf.ReadFrom(resp.Body)
			fmt.Println(buf.String())
			err = json.NewDecoder(buf).Decode(v)
			// err = json.NewDecoder(resp.Body).Decode(v)
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