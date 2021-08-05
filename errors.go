package treezor

import (
	"fmt"
	"io/ioutil"
	"net/http"

	json "github.com/tifo/treezor-sdk/internal/json"
)

// Error code for given status
const (
	ErrCodeInsuficientFunds = 15030
	ErrCodeCardWrongPIN     = 32056
	ErrCodeCardLost         = 32095
	ErrCodeCardStolen       = 32096
	ErrCodeCardBlocked      = 32111
)

// An apiErrorResponse reports one or more errors caused by an API request.
type apiErrorResponse struct {
	Errors []APIError `json:"errors,omitempty"` // Formatted error
	Error  *string    `json:"error,omitempty"`  // Simple error used by liveness endpoint
}

// APIError reports more details on an individual error in an apiErrorResponse.
type APIError struct {
	Code                  int      `json:"errorCode,omitempty"`
	Message               string   `json:"errorMessage"`
	AdditionalInformation []string `json:"additionalInformation,omitempty"`
}

// TODO: errors in legacy API with legacy client do not have the same format as connect ones (see https://docs.treezor.com/guide/api-basics/response-codes.html#error-attributes).
// NOTE: we might need to decide to specialized this package for connect only, if it starts to be come a pain and rename it to `treezor-connect-sdk`.

func (e *APIError) Error() string {
	return fmt.Sprintf("%v error caused because: %v", e.Code, e.Message)
}

// Error represents an all errors apiErrorResponse as API errors (transforming simple errors to APIErrors).
type Error struct {
	Response *http.Response // HTTP response that caused this error
	Errors   []APIError     // Formatted Errors
}

func (r *Error) Error() string {
	return fmt.Sprintf("%v %v: %d %+v",
		r.Response.Request.Method, sanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, r.Errors)
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range.
// API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse. Any other
// response body will be silently ignored.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= http.StatusOK && c < http.StatusBadRequest {
		return nil
	}
	errorResponse := &apiErrorResponse{}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		_ = json.Unmarshal(data, errorResponse)
	}
	if errorResponse.Error != nil {
		errorResponse.Errors = append(errorResponse.Errors, APIError{Message: *errorResponse.Error})
	}

	return &Error{
		Response: r,
		Errors:   errorResponse.Errors,
	}
}
