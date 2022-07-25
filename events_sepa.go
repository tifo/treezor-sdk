package treezor

import "github.com/tifo/treezor-sdk/internal/types"

type SepaSddrEvent struct {
	SepaSddrs []*SepaSddr
}

type SepaSddr struct {
	WalletID                  *types.Identifier `json:"wallet_id,omitempty"`
	VirtualIbanID             *types.Identifier `json:"virtual_iban_id,omitempty"`
	TransactionID             *types.Identifier `json:"transaction_id,omitempty"`
	SequenceType              *string           `json:"sequence_type,omitempty"` // NOTE: can be an enum
	RejectReasonCode          *string           `json:"reject_reason_code,omitempty"`
	InterbankSettlementAmount *types.Amount     `json:"interbank_settlement_amount,omitempty"`
	RequestedCollectionDate   *types.Date       `json:"requested_collection_date,omitempty"`
	MandateID                 *string           `json:"mandate_id,omitempty"`
	SepaCreditorIdentifier    *string           `json:"sepa_creditor_identifier,omitempty"`
	DateOfSignature           *types.Date       `json:"date_of_signature,omitempty"`
	DebitorName               *string           `json:"debitor_name,omitempty"`
	DebitorAddress            *string           `json:"debitor_address,omitempty"`
	DebitorCountry            *string           `json:"debitor_country,omitempty"`
	CreditorName              *string           `json:"creditor_name,omitempty"`
	CreditorAddress           *string           `json:"creditor_address,omitempty"`
	CreditorCountry           *string           `json:"creditor_country,omitempty"`
	UnstructuredField         *string           `json:"unstructured_field,omitempty"`
	BankaccountID             *types.Identifier `json:"bankaccount_id,omitempty"`
	BeneficiaryID             *types.Identifier `json:"beneficiary_id,omitempty"`
}

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
