package treezor

import (
	"time"

	"github.com/tifo/treezor-sdk/internal/types"
)

type PayinRefund struct {
	PayinRefundID     *types.Identifier `json:"payinrefundId,omitempty"`
	PayinRefundTag    *string           `json:"payinrefundTag,omitempty"`
	PayinRefundStatus *string           `json:"payinrefundStatus,omitempty"` // NOTE: can be an enum
	WalletID          *types.Identifier `json:"walletId,omitempty"`
	PayinID           *types.Identifier `json:"payinId,omitempty"`
	PayinRefundDate   *types.Date       `json:"payinrefundDate,omitempty"`
	Amount            *types.Amount     `json:"amount,omitempty"`
	Currency          *Currency         `json:"currency,omitempty"`
	CreatedDate       *time.Time        `json:"createdDate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	ModifiedDate      *time.Time        `json:"modifiedDate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	ReasonTms         *string           `json:"reasonTms,omitempty"`
	TotalRows         *types.Integer    `json:"totalRows,omitempty"`
	CodeStatus        *types.Identifier `json:"codeStatus,omitempty"`        // Legacy field
	InformationStatus *string           `json:"informationStatus,omitempty"` // Legacy field
}

// TODO: Add PayinRefund API
