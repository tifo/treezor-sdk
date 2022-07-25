package treezor

import (
	"time"

	"github.com/tifo/treezor-sdk/internal/types"
)

// CardChargebackEvent represents a card.aquiring.chargeback event
type CardChargebackEvent struct {
	Chargebacks []*Chargeback `json:"chargebacks"`
}

type Chargeback struct {
	UserID               *types.Identifier `json:"userId,omitempty"`
	WalletID             *types.Identifier `json:"walletId,omitempty"`
	PayinID              *types.Identifier `json:"payinId,omitempty"`
	TransactionReference *types.Identifier `json:"transactionReference,omitempty"`
	PayinRefundID        *types.Identifier `json:"payinRefundId,omitempty"`
	PaymentMethodId      *types.Identifier `json:"paymentMethodId,omitempty"`
	PaymentBrand         *string           `json:"paymentBrand,omitempty"`
	Currency             *Currency         `json:"currency,omitempty"`
	Amount               *string           `json:"amount,omitempty"`
	Country              *string           `json:"country,omitempty"`
	IsRefunded           *types.Boolean    `json:"isRefunded,omitempty"`
	ChargebackReason     *string           `json:"chargebackReason,omitempty"`
	// NOTE: Might be in Europe/London
	PayinCreatedDate       *time.Time `json:"payinCreatedDate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	PayinRefundCreatedDate *time.Time `json:"payinRefundCreatedDate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	ChargebackCreatedDate  *time.Time `json:"chargebackCreatedDate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
}
