package treezor

import (
	"fmt"
	"net/http"
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
