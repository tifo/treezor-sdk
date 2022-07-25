package treezor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// HearthbeatService handles communication with the heartbeat related
// methods of the Treezor API.
//
// Treezor API docs: https://www.treezor.com/api-documentation/#/heartbeat
type HearthbeatService service

type HeartbeatPingOptions struct {
	Access
}

// Ping will try to reach the Treezor API. Returns true if the API is healthy, otherwise
// it returns false. (Legacy clients only)
func (s *HearthbeatService) Ping(ctx context.Context, opts *HeartbeatPingOptions) (bool, *http.Response, error) {
	u := fmt.Sprintf("heartbeats")
	u, err := addOptions(u, opts)
	if err != nil {
		return false, nil, errors.WithStack(err)
	}
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

// NOTE: Heartbeats endpoint is not available for Treezor Connect
