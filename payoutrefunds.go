package treezor

import (
	"time"

	"github.com/tifo/treezor-sdk/internal/types"
)

type PayoutRefund struct {
	ID                *types.Identifier `json:"id,omitempty"`
	Tag               *string           `json:"tag,omitempty"`
	PayoutID          *types.Identifier `json:"payoutId,omitempty"`
	RequestAmount     *types.Amount     `json:"requestAmount,omitempty"`
	RequestCurrency   *Currency         `json:"requestCurrency,omitempty"`
	RequestComment    *string           `json:"requestComment,omitempty"`
	ReasonCode        *string           `json:"reasonCode,omitempty"`
	RefundAmount      *types.Amount     `json:"refundAmount,omitempty"`
	RefundCurrency    *Currency         `json:"refundCurrency,omitempty"`
	RefundDate        *time.Time        `json:"refundDate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	RefundComment     *string           `json:"refundComment,omitempty"`
	CreatedDate       *time.Time        `json:"createdDate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	ModifiedDate      *time.Time        `json:"modifiedDate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	CodeStatus        *types.Identifier `json:"codeStatus,omitempty"`        // Legacy field
	InformationStatus *string           `json:"informationStatus,omitempty"` // Legacy field
}

// TODO: Add PayoutRefund API
