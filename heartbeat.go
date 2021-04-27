package treezor

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

// HearthbeatService handles communication with the heartbeat related
// methods of the Treezor API.
//
// Treezor API docs: https://www.treezor.com/api-documentation/#/heartbeat
type HearthbeatService service

// Ping will try to reach the Treezor API. Returns true if the API is healthy, otherwise
// it returns false.
func (s *HearthbeatService) Ping(ctx context.Context) (bool, *http.Response, error) {
	u := "heartbeats"
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return false, resp, errors.WithStack(err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return false, resp, errors.WithStack(err)
	}

	return true, resp, errors.WithStack(err)
}
