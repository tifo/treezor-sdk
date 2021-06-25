package event

import (
	"github.com/tifo/treezor-sdk"
)

type CardTransactionEvent struct {
	CardTransactions []*treezor.CardTransaction `json:"cardtransactions"`
}
