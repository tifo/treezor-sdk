package event

import (
	"github.com/tifo/treezor-sdk"
	"github.com/tifo/treezor-sdk/types"
)

// CardEvent represents a card event
type CardEvent struct {
	Cards []*treezor.Card `json:"cards"`
}

// CardChargebackEvent represents a card.aquiring.chargeback event
type CardChargebackEvent struct {
	Chargebacks []*Chargeback `json:"chargebacks"`
}

type Chargeback struct {
	UserID                 *types.Identifier     `json:"userId,omitempty"`
	WalletID               *types.Identifier     `json:"walletId,omitempty"`
	PayinID                *types.Identifier     `json:"payinId,omitempty"`
	TransactionReference   *types.Identifier     `json:"transactionReference,omitempty"`
	PayinRefundID          *types.Identifier     `json:"payinRefundId,omitempty"`
	PaymentMethodId        *types.Identifier     `json:"paymentMethodId,omitempty"`
	PaymentBrand           *string               `json:"paymentBrand,omitempty"`
	Currency               *types.Currency       `json:"currency,omitempty"`
	Amount                 *string               `json:"amount,omitempty"`
	Country                *string               `json:"country,omitempty"`
	IsRefunded             *types.Boolean        `json:"isRefunded,omitempty"`
	ChargebackReason       *string               `json:"chargebackReason,omitempty"`
	PayinCreatedDate       *types.TimestampParis `json:"payinCreatedDate,omitempty"`       // Might be a TimestampLondon
	PayinRefundCreatedDate *types.TimestampParis `json:"payinRefundCreatedDate,omitempty"` // Might be a TimestampLondon
	ChargebackCreatedDate  *types.TimestampParis `json:"chargebackCreatedDate,omitempty"`  // Might be a TimestampLondon
}
