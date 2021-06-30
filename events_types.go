package treezor

// BalanceEvent represents a balance event
type BalanceEvent struct {
	Balances []*Balance `json:"balances,omitempty"`
}

type BankAccountEvent struct {
	BankAccounts []*BankAccount `json:"bankaccounts"`
}

type BeneficiaryEvent struct {
	Beneficiaries []*Beneficiary `json:"beneficiaries"`
}

// CardEvent represents a card event
type CardEvent struct {
	Cards []*Card `json:"cards"`
}

type CardTransactionEvent struct {
	CardTransactions []*CardTransaction `json:"cardtransactions"`
}

type CountryRestrictionGroupEvent struct {
	CountryRestrictionGroups []*CountryRestrictionGroup `json:"countryRestrictionGroups"`
}

// DocumentEvent represents a document event
type DocumentEvent struct {
	Documents []*Document `json:"documents"`
}

type MandateEvent struct {
	Mandates []*Mandate `json:"mandates"`
}

type MCCRestrictionGroupEvent struct {
	MCCRestrictionGroups []*MCCRestrictionGroup `json:"mccIdRestrictionGroups"`
}

type MIDRestrictionGroupEvent struct {
	MerchantIDRestrictionGroups []*MIDRestrictionGroup `json:"merchantIdRestrictionGroups"`
}

type PayinEvent struct {
	Payins []*Payin `json:"payins"`
}

type PayinRefundEvent struct {
	PayinRefunds []*PayinRefund `json:"payinrefunds"`
}

type PayoutEvent struct {
	Payouts []*Payout `json:"payouts"`
}

type PayoutRefundEvent struct {
	PayoutRefunds []*PayoutRefund `json:"payoutRefunds"`
}

type RecallREvent struct {
	RecallRs []*RecallR `json:"recallrs"`
}

type TransactionEvent struct {
	Transactions []*Transaction `json:"transactions"`
}

type TransferEvent struct {
	Transfers []*Transfer `json:"transfers"`
}

type TransferRefundEvent struct {
	TransferRefunds []*TransferRefund `json:"transferrefunds"`
}

// UserEvent represents a user event
type UserEvent struct {
	Users []*User `json:"users"`
}

// WalletEvent represents a wallet event
type WalletEvent struct {
	Wallets []*Wallet `json:"wallets"`
}
