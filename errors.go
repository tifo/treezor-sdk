package treezor

import (
	"fmt"
	"io/ioutil"
	"net/http"

	json "github.com/tifo/treezor-sdk/internal/json"
)

// Error represents an all errors apiErrorResponse as API errors (transforming simple errors to APIErrors).
type Error struct {
	Response *http.Response    // HTTP response that caused this error
	Errors   []ConnectAPIError // Formatted Errors
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
	errorResponse := &connectErrorResponse{}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		_ = json.Unmarshal(data, errorResponse)
	}
	if errorResponse.Error != nil {
		errorResponse.Errors = append(errorResponse.Errors, ConnectAPIError{Message: *errorResponse.Error})
	}

	return &Error{
		Response: r,
		Errors:   errorResponse.Errors,
	}
}
