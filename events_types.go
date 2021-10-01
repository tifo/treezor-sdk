package treezor

// BalanceEvent represents a balance event
type BalanceEvent struct {
	Balances []*Balance `json:"balances,omitempty"`
}

func (evt *BalanceEvent) GetBalance() *Balance {
	if len(evt.Balances) == 0 {
		return nil
	}
	return evt.Balances[0]
}

// BankAccountEvent
type BankAccountEvent struct {
	BankAccounts []*BankAccount `json:"bankaccounts"`
}

func (evt *BankAccountEvent) GetBankAccount() *BankAccount {
	if len(evt.BankAccounts) == 0 {
		return nil
	}
	return evt.BankAccounts[0]
}

// BeneficiaryEvent
type BeneficiaryEvent struct {
	Beneficiaries []*Beneficiary `json:"beneficiaries"`
}

func (evt *BeneficiaryEvent) GetBeneficiary() *Beneficiary {
	if len(evt.Beneficiaries) == 0 {
		return nil
	}
	return evt.Beneficiaries[0]
}

// CardEvent represents a card event
type CardEvent struct {
	Cards []*Card `json:"cards"`
}

func (evt *CardEvent) GetCard() *Card {
	if len(evt.Cards) == 0 {
		return nil
	}
	return evt.Cards[0]
}

// CardTransactionEvent
type CardTransactionEvent struct {
	CardTransactions []*CardTransaction `json:"cardtransactions"`
}

func (evt *CardTransactionEvent) GetCardTransaction() *CardTransaction {
	if len(evt.CardTransactions) == 0 {
		return nil
	}
	return evt.CardTransactions[0]
}

// CountryRestrictionGroupEvent
type CountryRestrictionGroupEvent struct {
	CountryRestrictionGroups []*CountryRestrictionGroup `json:"countryRestrictionGroups"`
}

func (evt *CountryRestrictionGroupEvent) GetCountryRestrictionGroup() *CountryRestrictionGroup {
	if len(evt.CountryRestrictionGroups) == 0 {
		return nil
	}
	return evt.CountryRestrictionGroups[0]
}

// DocumentEvent represents a document event
type DocumentEvent struct {
	Documents []*Document `json:"documents"`
}

func (evt *DocumentEvent) GetDocument() *Document {
	if len(evt.Documents) == 0 {
		return nil
	}
	return evt.Documents[0]
}

// MandateEvent
type MandateEvent struct {
	Mandates []*Mandate `json:"mandates"`
}

func (evt *MandateEvent) GetMandate() *Mandate {
	if len(evt.Mandates) == 0 {
		return nil
	}
	return evt.Mandates[0]
}

// MCCRestrictionGroupEvent
type MCCRestrictionGroupEvent struct {
	MCCRestrictionGroups []*MCCRestrictionGroup `json:"mccIdRestrictionGroups"`
}

func (evt *MCCRestrictionGroupEvent) GetMCCRestrictionGroup() *MCCRestrictionGroup {
	if len(evt.MCCRestrictionGroups) == 0 {
		return nil
	}
	return evt.MCCRestrictionGroups[0]
}

// MIDRestrictionGroupEvent
type MIDRestrictionGroupEvent struct {
	MerchantIDRestrictionGroups []*MIDRestrictionGroup `json:"merchantIdRestrictionGroups"`
}

func (evt *MIDRestrictionGroupEvent) GetMerchantIDRestrictionGroup() *MIDRestrictionGroup {
	if len(evt.MerchantIDRestrictionGroups) == 0 {
		return nil
	}
	return evt.MerchantIDRestrictionGroups[0]
}

// PayinEvent
type PayinEvent struct {
	Payins []*Payin `json:"payins"`
}

func (evt *PayinEvent) GetPayin() *Payin {
	if len(evt.Payins) == 0 {
		return nil
	}
	return evt.Payins[0]
}

// PayinRefundEvent
type PayinRefundEvent struct {
	PayinRefunds []*PayinRefund `json:"payinrefunds"`
}

func (evt *PayinRefundEvent) GetPayinRefund() *PayinRefund {
	if len(evt.PayinRefunds) == 0 {
		return nil
	}
	return evt.PayinRefunds[0]
}

// PayoutEvent
type PayoutEvent struct {
	Payouts []*Payout `json:"payouts"`
}

func (evt *PayoutEvent) GetPayout() *Payout {
	if len(evt.Payouts) == 0 {
		return nil
	}
	return evt.Payouts[0]
}

// PayoutRefundEvent
type PayoutRefundEvent struct {
	PayoutRefunds []*PayoutRefund `json:"payoutRefunds"`
}

func (evt *PayoutRefundEvent) GetPayoutRefund() *PayoutRefund {
	if len(evt.PayoutRefunds) == 0 {
		return nil
	}
	return evt.PayoutRefunds[0]
}

// RecallREvent
type RecallREvent struct {
	RecallRs []*RecallR `json:"recallrs"`
}

func (evt *RecallREvent) GetRecallR() *RecallR {
	if len(evt.RecallRs) == 0 {
		return nil
	}
	return evt.RecallRs[0]
}

// TransactionEvent
type TransactionEvent struct {
	Transactions []*Transaction `json:"transactions"`
}

func (evt *TransactionEvent) GetTransaction() *Transaction {
	if len(evt.Transactions) == 0 {
		return nil
	}
	return evt.Transactions[0]
}

// TransferEvent
type TransferEvent struct {
	Transfers []*Transfer `json:"transfers"`
}

func (evt *TransferEvent) GetTransfer() *Transfer {
	if len(evt.Transfers) == 0 {
		return nil
	}
	return evt.Transfers[0]
}

// TransferRefundEvent
type TransferRefundEvent struct {
	TransferRefunds []*TransferRefund `json:"transferrefunds"`
}

func (evt *TransferRefundEvent) GetTransferRefund() *TransferRefund {
	if len(evt.TransferRefunds) == 0 {
		return nil
	}
	return evt.TransferRefunds[0]
}

// UserEvent represents a user event
type UserEvent struct {
	Users []*User `json:"users"`
}

func (evt *UserEvent) GetUser() *User {
	if len(evt.Users) == 0 {
		return nil
	}
	return evt.Users[0]
}

// WalletEvent represents a wallet event
type WalletEvent struct {
	Wallets []*Wallet `json:"wallets"`
}

func (evt *WalletEvent) GetWallet() *Wallet {
	if len(evt.Wallets) == 0 {
		return nil
	}
	return evt.Wallets[0]
}
