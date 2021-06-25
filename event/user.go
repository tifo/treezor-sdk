package event

import (
	"github.com/tifo/treezor-sdk"
)

// UserEvent represents a user event
type UserEvent struct {
	Users []*treezor.User `json:"users"`
}
