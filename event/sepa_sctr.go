package event

import "github.com/tifo/treezor-sdk/internal/types"

type SepaSctrEvent struct {
	SepaSctrs []*SepaSctr
}

type SepaSctr struct {
	WalletID                  *types.Identifier `json:"wallet_id,omitempty"`
	VirtualIBANID             *types.Identifier `json:"virtual_iban_id,omitempty"`
	TransactionID             *types.Identifier `json:"transaction_id,omitempty"`
	InterbankSettlementAmount *types.Amount     `json:"interbank_settlement_amount,omitempty"`
	DebitorName               *string           `json:"debitor_name,omitempty"`
	DebitorAddress            *string           `json:"debitor_address,omitempty"`
	DebitorCountry            *string           `json:"debitor_country,omitempty"`
	CreditorName              *string           `json:"creditor_name,omitempty"`
	CreditorAddress           *string           `json:"creditor_address,omitempty"`
	CreditorCountry           *string           `json:"creditor_country,omitempty"`
	UnstructuredField         *string           `json:"unstructured_field,omitempty"`
	ReturnReasonCode          *string           `json:"return_reason_code,omitempty"`
}
