package event

import "github.com/tifo/treezor-sdk/types"

type TransferRefundEvent struct {
	TransferRefunds []*TransferRefund `json:"transferrefunds"`
}

type TransferRefund struct {
	TransferRefundID     *types.Identifier     `json:"transferrefundId,omitempty"`
	TransferRefundTag    *string               `json:"transferrefundTag,omitempty"`
	TransferRefundStatus *string               `json:"transferrefundStatus,omitempty"` // NOTE: can be an enum
	WalletID             *types.Identifier     `json:"walletId,omitempty"`
	TransferID           *types.Identifier     `json:"transferId,omitempty"`
	TransferRefundDate   *types.TimestampParis `json:"transferrefundDate,omitempty"`
	Amount               *types.Amount         `json:"amount,omitempty"`
	Currency             *types.Currency       `json:"currency,omitempty"`
	CreatedDate          *types.TimestampParis `json:"createdDate,omitempty"`
	ModifiedDate         *types.TimestampParis `json:"modifiedDate,omitempty"`
	TotalRows            *types.Integer        `json:"totalRows,omitempty"`
	CodeStatus           *types.Identifier     `json:"codeStatus,omitempty"`        // NOTE: Legacy + Webhook
	InformationStatus    *string               `json:"informationStatus,omitempty"` // NOTE: Legacy + Webhook
}

// TODO: Add TransferRefund API
