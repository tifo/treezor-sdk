package event

import "github.com/tifo/treezor-sdk/types"

type PayinRefundEvent struct {
	PayinRefunds []*PayinRefund `json:"payinrefunds"`
}

type PayinRefund struct {
	PayinRefundID     *types.Identifier     `json:"payinrefundId,omitempty"`
	PayinRefundTag    *string               `json:"payinrefundTag,omitempty"`
	PayinRefundStatus *string               `json:"payinrefundStatus,omitempty"` // NOTE: can be an enum
	WalletID          *types.Identifier     `json:"walletId,omitempty"`
	PayinID           *types.Identifier     `json:"payinId,omitempty"`
	PayinRefundDate   *types.Date           `json:"payinrefundDate,omitempty"`
	Amount            *types.Amount         `json:"amount,omitempty"`
	Currency          *types.Currency       `json:"currency,omitempty"`
	CreatedDate       *types.TimestampParis `json:"createdDate,omitempty"`
	ModifiedDate      *types.TimestampParis `json:"modifiedDate,omitempty"`
	ReasonTms         *string               `json:"reasonTms,omitempty"`
	TotalRows         *types.Integer        `json:"totalRows,omitempty"`
	CodeStatus        *types.Identifier     `json:"codeStatus,omitempty"`        // NOTE: Legacy + Webhook
	InformationStatus *string               `json:"informationStatus,omitempty"` // NOTE: Legacy + Webhook
}

// TODO: Add PayinRefund API
