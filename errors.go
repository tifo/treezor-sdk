package treezor

import (
	"io/ioutil"
	"net/http"

	json "github.com/tifo/treezor-sdk/internal/json"
)

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
		errorResponse.Errors = append(errorResponse.Errors, APIError{Message: *errorResponse.Error})
	}

	return &Error{
		Response: r,
		Errors:   errorResponse.Errors,
	}
}
