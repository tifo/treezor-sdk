package event

import (
	"github.com/tifo/treezor-sdk"
	"github.com/tifo/treezor-sdk/internal/types"
)

type TransactionEvent struct {
	Transactions []*Transaction `json:"transactions"`
}

type Transaction struct {
	TransactionID       *types.Identifier     `json:"transactionId,omitempty"`
	WalletDebitID       *types.Identifier     `json:"walletDebitId,omitempty"`
	WalletCreditID      *types.Identifier     `json:"walletCreditId,omitempty"`
	TransactionType     *string               `json:"transactionType,omitempty"` // NOTE: Can be an enum
	ForeignID           *types.Identifier     `json:"foreignId,omitempty"`
	Name                *string               `json:"name,omitempty"`
	Description         *string               `json:"description,omitempty"`
	ValueDate           *types.Date           `json:"valueDate,omitempty"`
	ExecutionDate       *types.Date           `json:"executionDate,omitempty"`
	Amount              *types.Amount         `json:"amount,omitempty"`
	WalletDebitBalance  *types.Amount         `json:"walletDebitBalance,omitempty"`
	WalletCreditBalance *types.Amount         `json:"walletCreditBalance,omitempty"`
	Currency            *treezor.Currency     `json:"currency,omitempty"`
	CreatedDate         *types.TimestampParis `json:"createdDate,omitempty"`
	TotalRows           *types.Integer        `json:"totalRows,omitempty"`
}

// TODO: Add Transaction API
