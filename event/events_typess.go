package event

import "github.com/tifo/treezor-sdk"

// BalanceEvent represents a balance event
type BalanceEvent struct {
	Balances []*treezor.Balance `json:"balances,omitempty"`
}

type BankAccountEvent struct {
	BankAccounts []*treezor.BankAccount `json:"bankaccounts"`
}

type BeneficiaryEvent struct {
	Beneficiaries []*treezor.Beneficiary `json:"beneficiaries"`
}

// CardEvent represents a card event
type CardEvent struct {
	Cards []*treezor.Card `json:"cards"`
}

type CardTransactionEvent struct {
	CardTransactions []*treezor.CardTransaction `json:"cardtransactions"`
}

type CountryRestrictionGroupEvent struct {
	CountryRestrictionGroups []*treezor.CountryRestrictionGroup `json:"countryRestrictionGroups"`
}

// DocumentEvent represents a document event
type DocumentEvent struct {
	Documents []*treezor.Document `json:"documents"`
}

type MandateEvent struct {
	Mandates []*treezor.Mandate `json:"mandates"`
}

type MCCRestrictionGroupEvent struct {
	MCCRestrictionGroups []*treezor.MCCRestrictionGroup `json:"mccIdRestrictionGroups"`
}

type MIDRestrictionGroupEvent struct {
	MerchantIDRestrictionGroups []*treezor.MIDRestrictionGroup `json:"merchantIdRestrictionGroups"`
}

type PayinEvent struct {
	Payins []*treezor.Payin `json:"payins"`
}

type PayinRefundEvent struct {
	PayinRefunds []*treezor.PayinRefund `json:"payinrefunds"`
}

type PayoutEvent struct {
	Payouts []*treezor.Payout `json:"payouts"`
}

type PayoutRefundEvent struct {
	PayoutRefunds []*treezor.PayoutRefund `json:"payoutRefunds"`
}

type RecallREvent struct {
	RecallRs []*treezor.RecallR `json:"recallrs"`
}

type TransactionEvent struct {
	Transactions []*treezor.Transaction `json:"transactions"`
}

type TransferEvent struct {
	Transfers []*treezor.Transfer `json:"transfers"`
}

type TransferRefundEvent struct {
	TransferRefunds []*treezor.TransferRefund `json:"transferrefunds"`
}

// UserEvent represents a user event
type UserEvent struct {
	Users []*treezor.User `json:"users"`
}

// WalletEvent represents a wallet event
type WalletEvent struct {
	Wallets []*treezor.Wallet `json:"wallets"`
}
