package treezor

import (
	"time"

	"github.com/tifo/treezor-sdk/internal/types"
)

type RecallR struct {
	ID                    *types.Identifier `json:"id,omitempty"`
	CxlID                 *string           `json:"cxl_id,omitempty"`
	StatusID              *types.Identifier `json:"status_id,omitempty"` // NOTE: Can be an enum
	StatusLabel           *string           `json:"status_label,omitempty"`
	ReasonCode            *string           `json:"reason_code,omitempty"` // NOTE: Can be an enum
	AdditionalInformation *string           `json:"additional_information,omitempty"`
	UserID                *types.Identifier `json:"user_id,omitempty"`
	UserName              *string           `json:"user_name,omitempty"`
	UserStatusID          *types.Identifier `json:"user_status_id,omitempty"`
	WalletID              *types.Identifier `json:"wallet_id,omitempty"`
	WalletStatusID        *types.Identifier `json:"wallet_status_id,omitempty"`
	WalletActivationDate  *time.Time        `json:"wallet_activation_date,omitempty" layout:"Treezor" loc:"Europe/Paris"` // NOTE: Can be a types.Date with 00:00:00 suffix
	SctrID                *types.Identifier `json:"sctr_id,omitempty"`
	SctrTxID              *types.Identifier `json:"sctr_tx_id,omitempty"`
	SctrAmount            *types.Amount     `json:"sctr_amount,omitempty"`
	SctrCurrency          *Currency         `json:"sctr_currency,omitempty"`
	// NOTE: *Date could be types.Date if we add authorize a 00:00:00 suffix to them
	// NOTE: their is a typo in the example (settelment vs settlement), we need to check which one is right with real data
	SctrSettlementDate *time.Time        `json:"sctr_settlement_date,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	SctrSettelmentDate *time.Time        `json:"sctr_settelment_date,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	SctrDbtrName       *string           `json:"sctr_dbtr_name,omitempty"`
	ReceivedDate       *time.Time        `json:"received_date,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	PayinRefundID      *types.Identifier `json:"payinrefund_id,omitempty"`
}

// TODO: Add RecallR API & Models (RecallR model in Events and API are different (snakeCase vs camlCase))
// TODO: Check as some "*Date" field in treezors API uses the 00:00:00 suffix, we might want to handle that with the same type.
// NOTE: https://docs.treezor.com/guide/transfers/sepa-recalls.html
