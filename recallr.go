package treezor

import "github.com/tifo/treezor-sdk/internal/types"

type RecallR struct {
	ID                    *types.Identifier     `json:"id,omitempty"`
	CxlID                 *string               `json:"cxl_id,omitempty"`
	StatusID              *types.Identifier     `json:"status_id,omitempty"` // NOTE: Can be an enum
	StatusLabel           *string               `json:"status_label,omitempty"`
	ReasonCode            *string               `json:"reason_code,omitempty"` // NOTE: Can be an enum
	AdditionalInformation *string               `json:"additional_information,omitempty"`
	UserID                *types.Identifier     `json:"user_id,omitempty"`
	UserName              *string               `json:"user_name,omitempty"`
	UserStatusID          *types.Identifier     `json:"user_status_id,omitempty"`
	WalletID              *types.Identifier     `json:"wallet_id,omitempty"`
	WalletStatusID        *types.Identifier     `json:"wallet_status_id,omitempty"`
	WalletActivationDate  *types.TimestampParis `json:"wallet_activation_date,omitempty"` // NOTE: Can be a types.Date with 00:00:00 suffix
	SctrID                *types.Identifier     `json:"sctr_id,omitempty"`
	SctrTxID              *types.Identifier     `json:"sctr_tx_id,omitempty"`
	SctrAmount            *types.Amount         `json:"sctr_amount,omitempty"`
	SctrCurrency          *Currency             `json:"sctr_currency,omitempty"`
	// NOTE: their is a typo in the example (settelment vs settlement), we need to check which one is right with real data
	SctrSettlementDate *types.TimestampParis `json:"sctr_settlement_date,omitempty"` // NOTE: Can be a types.Date if we add the 00:00:00 suffix
	SctrSettelmentDate *types.TimestampParis `json:"sctr_settelment_date,omitempty"` // NOTE: Can be a types.Date if we add the 00:00:00 suffix
	SctrDbtrName       *string               `json:"sctr_dbtr_name,omitempty"`
	ReceivedDate       *types.TimestampParis `json:"received_date,omitempty"` // NOTE: Can be a types.Date with 00:00:00 suffix
	PayinRefundID      *types.Identifier     `json:"payinrefund_id,omitempty"`
}

// TODO: Add RecallR API & Models (RecallR model in Events and API are different (snakeCase vs camlCase))
// TODO: Check as some "*Date" field in treezors API uses the 00:00:00 suffix, we might want to handle that with the same type.
// NOTE: https://docs.treezor.com/guide/transfers/sepa-recalls.html
