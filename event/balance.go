package event

import (
	"github.com/tifo/treezor-sdk"
)

// BalanceEvent represents a balance event
type BalanceEvent struct {
	Balances []*treezor.Balance `json:"balances,omitempty"`
}
