package treezor

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/pkg/errors"

	json "github.com/tifo/treezor-sdk/internal/json"
)

const (
	// defaultStagingBaseURL    = "https://sandbox.treezor.com/v1/"
	// defaultProductionBaseURL = "https://treezor.com/v1/"
	userAgent = "go-treezor-sdk/0.1.0"
)

// ConnectClient allows access to both the legacy Treezor API as well as the new Connect endpoints
type ConnectClient struct {
	API     *apiClient
	Connect *connectClient
}

// LegacyClient allows access to the legacy Treezor API
type LegacyClient apiClient

type connectClient struct {
	common    service // Reuse a single struct instead of allocating one for each service on the heap.
	KYC       *ConnectKYCService
	Operation *ConnectOperationService
	Statement *ConnectStatementService
}

type apiClient struct {
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

func newAPIClient(httpClient *http.Client, apiBaseURL *url.URL) *apiClient {
	c := &apiClient{}
	c.common.client = &HTTPClient{client: httpClient, BaseURL: apiBaseURL, UserAgent: userAgent}
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

func newConnectClient(httpClient *http.Client, connectBaseURL *url.URL) *connectClient {
	c := &connectClient{}
	c.common.client = &HTTPClient{client: httpClient, BaseURL: connectBaseURL, UserAgent: userAgent}
	c.KYC = (*ConnectKYCService)(&c.common)
	c.Operation = (*ConnectOperationService)(&c.common)
	c.Statement = (*ConnectStatementService)(&c.common)
	return c
}

// // NewClient returns a new Treezor API client. If a nil httpClient is
// // provided, http.DefaultClient will be used. To use API methods which require
// // authentication, provide an http.Client that will perform the authentication
// // for you (such as that provided by the golang.org/x/oauth2 library).
// func NewClient(httpClient *http.Client, isProduction bool) *LegacyClient {
// 	if httpClient == nil {
// 		httpClient = http.DefaultClient
// 	}
// 	baseURL, _ := url.Parse(defaultStagingBaseURL)
// 	if isProduction {
// 		baseURL, _ = url.Parse(defaultProductionBaseURL)
// 	}
// 	return (*LegacyClient)(newAPIClient(httpClient, baseURL))
// }

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
	api := newAPIClient(httpClient, apiBaseURL)
	connect := newConnectClient(httpClient, baseURL)
	return &ConnectClient{
		API:     api,
		Connect: connect,
	}
}

type HTTPClient struct {
	// HTTP client used to communicate with the API.
	client *http.Client
	// Base URL for API requests. Defaults to the public Treezor API. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL
	// User agent used when communicating with the Treezor API.
	UserAgent string
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *HTTPClient) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {

	p := path.Join(c.BaseURL.Path, urlStr)
	u, err := c.BaseURL.Parse(p)
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
func (c *HTTPClient) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
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
			if u, err := url.Parse(e.URL); err == nil {
				e.URL = sanitizeURL(u).String()
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
