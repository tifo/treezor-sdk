package treezor

import (
	"time"

	"github.com/tifo/treezor-sdk/internal/types"
)

type TransferRefund struct {
	TransferRefundID     *types.Identifier `json:"transferrefundId,omitempty"`
	TransferRefundTag    *string           `json:"transferrefundTag,omitempty"`
	TransferRefundStatus *string           `json:"transferrefundStatus,omitempty"` // NOTE: can be an enum
	WalletID             *types.Identifier `json:"walletId,omitempty"`
	TransferID           *types.Identifier `json:"transferId,omitempty"`
	TransferRefundDate   *time.Time        `json:"transferrefundDate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	Amount               *types.Amount     `json:"amount,omitempty"`
	Currency             *Currency         `json:"currency,omitempty"`
	CreatedDate          *time.Time        `json:"createdDate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	ModifiedDate         *time.Time        `json:"modifiedDate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	TotalRows            *types.Integer    `json:"totalRows,omitempty"`
	CodeStatus           *types.Identifier `json:"codeStatus,omitempty"`        // Legacy field
	InformationStatus    *string           `json:"informationStatus,omitempty"` // Legacy field
}

// TODO: Add TransferRefund API
