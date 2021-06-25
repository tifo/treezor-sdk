package event

import "github.com/tifo/treezor-sdk"

type PayoutEvent struct {
	Payouts []*treezor.Payout `json:"payouts"`
}
