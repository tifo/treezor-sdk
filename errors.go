package treezor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ErrCodeInsuficientFunds = 15030
)

// An ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Errors   []Error        `json:"errors"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %+v",
		r.Response.Request.Method, sanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, r.Errors)
}

// An Error reports more details on an individual error in an ErrorResponse.
type Error struct {
	Code    int    `json:"errorCode"`
	Message string `json:"errorMessage"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v error caused because: %v", e.Code, e.Message)
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
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	return errorResponse
}
