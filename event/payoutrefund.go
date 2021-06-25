package event

import "github.com/tifo/treezor-sdk/types"

type PayoutRefundEvent struct {
	PayoutRefunds []*PayoutRefund `json:"payoutRefunds"`
}

type PayoutRefund struct {
	ID                *types.Identifier     `json:"id,omitempty"`
	Tag               *string               `json:"tag,omitempty"`
	PayoutID          *types.Identifier     `json:"payoutId,omitempty"`
	RequestAmount     *types.Amount         `json:"requestAmount,omitempty"`
	RequestCurrency   *types.Currency       `json:"requestCurrency,omitempty"`
	RequestComment    *string               `json:"requestComment,omitempty"`
	ReasonCode        *string               `json:"reasonCode,omitempty"`
	RefundAmount      *types.Amount         `json:"refundAmount,omitempty"`
	RefundCurrency    *types.Currency       `json:"refundCurrency,omitempty"`
	RefundDate        *types.TimestampParis `json:"refundDate,omitempty"`
	RefundComment     *string               `json:"refundComment,omitempty"`
	CreatedDate       *types.TimestampParis `json:"createdDate,omitempty"`
	ModifiedDate      *types.TimestampParis `json:"modifiedDate,omitempty"`
	CodeStatus        *types.Identifier     `json:"codeStatus,omitempty"`        // NOTE: Legacy + Webhook
	InformationStatus *string               `json:"informationStatus,omitempty"` // NOTE: Legacy + Webhook
}

// TODO: Add PayoutRefund API
