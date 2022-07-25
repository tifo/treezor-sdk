package treezor

import (
	"fmt"
	"net/http"
)

// connectErrorResponse reports one or more errors caused by an API request.
type connectErrorResponse struct {
	Errors []ConnectAPIError `json:"errors,omitempty"`
	Error  *string           `json:"error,omitempty"`
}

// ConnectAPIError reports more details on an individual error in an connectApiErrorResponse.
type ConnectAPIError struct {
	Type    string `json:"type"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
	DocURL  string `json:"docUrl,omitempty"`
}

func (e *ConnectAPIError) Error() string {
	return fmt.Sprintf("%v error caused because: %v", e.Code, e.Message)
}

// ConnectError represents an all errors apiErrorResponse as API errors (transforming simple errors to ConnectAPIError).
type ConnectError struct {
	Response *http.Response    // HTTP response that caused this error
	Errors   []ConnectAPIError // Formatted Errors
}

func (r *ConnectError) Error() string {
	return fmt.Sprintf("%v %v: %d %+v",
		r.Response.Request.Method, sanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, r.Errors)
}
