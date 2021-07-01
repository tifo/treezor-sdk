package treezor

import "github.com/tifo/treezor-sdk/internal/types"

type PayoutRefund struct {
	ID                *types.Identifier     `json:"id,omitempty"`
	Tag               *string               `json:"tag,omitempty"`
	PayoutID          *types.Identifier     `json:"payoutId,omitempty"`
	RequestAmount     *types.Amount         `json:"requestAmount,omitempty"`
	RequestCurrency   *Currency             `json:"requestCurrency,omitempty"`
	RequestComment    *string               `json:"requestComment,omitempty"`
	ReasonCode        *string               `json:"reasonCode,omitempty"`
	RefundAmount      *types.Amount         `json:"refundAmount,omitempty"`
	RefundCurrency    *Currency             `json:"refundCurrency,omitempty"`
	RefundDate        *types.TimestampParis `json:"refundDate,omitempty"`
	RefundComment     *string               `json:"refundComment,omitempty"`
	CreatedDate       *types.TimestampParis `json:"createdDate,omitempty"`
	ModifiedDate      *types.TimestampParis `json:"modifiedDate,omitempty"`
	CodeStatus        *types.Identifier     `json:"codeStatus,omitempty"`        // Legacy field
	InformationStatus *string               `json:"informationStatus,omitempty"` // Legacy field
}

// TODO: Add PayoutRefund API