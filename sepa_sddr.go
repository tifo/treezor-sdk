package treezor

//SepaSddrResponse represent data send when a sepa.reject_sdd is received
type SepaSddrResponse struct {
	SepaSddrs []*SepaSddr `json:"sepaSddrs"`
}

// SepaSddr represent a sdd reject object
type SepaSddr struct {
	Access
	WalletID                  *int64  `json:"wallet_id,omitempty"`
	VirtualIbanID             *string `json:"virtual_iban_id,omitempty"`
	TransactionID             *string `json:"transaction_id,omitempty"`
	SequenceType              *string `json:"sequence_type,omitempty"`
	RejectReasonCode          *string `json:"reject_reason_code,omitempty"`
	InterbankSettlementAmount *string `json:"interbank_settlement_amount,omitempty"`
	RequestedCollectionDate   *string `json:"requested_collection_date,omitempty"`
	MandateID                 *string `json:"mandate_id,omitempty"`
	SepaCreditorIdentifier    *string `json:"sepa_creditor_identifier,omitempty"`
	DateOfSignature           *string `json:"date_of_signature,omitempty"`
	DebitorName               *string `json:"debitor_name,omitempty"`
	DebitorAddress            *string `json:"debitor_address,omitempty"`
	DebitorCountry            *string `json:"debitor_country,omitempty"`
	CreditorName              *string `json:"creditor_name,omitempty"`
	CreditorAddress           *string `json:"creditor_address,omitempty"`
	CreditorCountry           *string `json:"creditor_country,omitempty"`
	UnstructuredField         *string `json:"unstructured_field,omitempty"`
	BankaccountID             *string `json:"bankaccount_id,omitempty"`
	BeneficiaryID             *string `json:"beneficiary_id,omitempty"`
}
