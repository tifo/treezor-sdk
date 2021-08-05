package treezor

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"

	"github.com/pkg/errors"

	json "github.com/tifo/treezor-sdk/internal/json"
	"github.com/tifo/treezor-sdk/internal/types"
)

// WebhookEvent represent an event that occurred on Treezor's infrastructure and is
// sent as a webhook to us. See https://docs.treezor.com/guide/webhooks/events-descriptions.html.
type WebhookEvent struct {
	WebhookID              *string          `json:"webhook_id,omitempty"`
	Webhook                *string          `json:"webhook,omitempty"`
	Object                 *string          `json:"object,omitempty"`
	ObjectID               *string          `json:"object_id,omitempty"`
	ObjectPayload          *json.RawMessage `json:"object_payload,omitempty" swaggertype:"object"`
	ObjectPayloadSignature *string          `json:"object_payload_signature,omitempty"`
}

func (e WebhookEvent) String() string {
	return types.Stringify(e)
}

// genMAC generates the HMAC signature for a message provided the secret key
// and hashFunc.
func genMAC(message, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	_, _ = mac.Write(message)
	return mac.Sum(nil)
}

func (e *WebhookEvent) Validate(secretKey []byte) (bool, error) {

	if len(e.GetObjectPayload()) == 0 {
		return false, errors.New("Webhook request has missing payload")
	}
	if e.GetObjectPayloadSignature() == "" {
		return false, errors.New("Webhook request has missing signature")
	}

	messageSignature, err := base64.StdEncoding.DecodeString(*e.ObjectPayloadSignature)
	if err != nil {
		return false, errors.Errorf("error decoding signature %q: %v", *e.ObjectPayloadSignature, err)
	}
	expectedSignature := genMAC(*e.ObjectPayload, secretKey)

	return hmac.Equal(messageSignature, expectedSignature), nil
}

// ParsePayload parses the event payload. For recognized event types,
// a value of the corresponding struct type will be returned.
//nolint:gocyclo
func (e *WebhookEvent) ParsePayload() (payload interface{}, err error) {
	switch e.GetWebhook() {

	// Balance (https://docs.treezor.com/guide/wallets/events.html#balances)
	case "balance.update":
		payload = &BalanceEvent{}

	// Beneficiary (https://docs.treezor.com/guide/transfers/events.html#beneficiaries)
	case "beneficiary.create":
		payload = &BeneficiaryEvent{}
	case "beneficiary.update":
		payload = &BeneficiaryEvent{}

	// SepaSctr
	case "sepa.return_sctr":
		payload = &SepaSctrEvent{}

	// Card (https://docs.treezor.com/guide/cards/events.html#cards)
	case "card.requestphysical":
		payload = &CardEvent{}
	case "card.createvirtual":
		payload = &CardEvent{}
	case "card.convertvirtual":
		payload = &CardEvent{}
	case "card.changepin":
		payload = &CardEvent{}
	case "card.activate":
		payload = &CardEvent{}
	case "card.renew":
		payload = &CardEvent{}
	case "card.regenerate":
		payload = &CardEvent{}
	case "card.update":
		payload = &CardEvent{}
	case "card.limits":
		payload = &CardEvent{}
	case "card.options":
		payload = &CardEvent{}
	case "card.setpin":
		payload = &CardEvent{}
	case "card.unblockpin":
		payload = &CardEvent{}
	case "card.lockunlock":
		payload = &CardEvent{}
	case "card.register3DS":
		payload = &CardEvent{}

	// CardDigitalization (https://docs.treezor.com/guide/cards/events.html#carddigitalization)
	case "cardDigitalization.create":
		payload = &CardDigitalizationEvent{}
	case "cardDigitalization.update":
		payload = &CardDigitalizationEvent{}
	case "cardDigitalization.activation":
		payload = &CardDigitalizationEvent{}
	case "cardDigitalization.deactivation":
		payload = &CardDigitalizationEvent{}
	case "cardDigitalization.complete":
		payload = &CardDigitalizationEvent{}

	// CardTransaction (https://docs.treezor.com/guide/cards/events.html#cardtransactions)
	case "cardtransaction.create":
		payload = &CardTransactionEvent{}

	// Card Aquiring (https://docs.treezor.com/guide/cards/events.html#cards)
	case "card.acquiring.chargeback.create":
		payload = &CardChargebackEvent{}

	// CountryGroup (https://docs.treezor.com/guide/cards/events.html#country-group)
	case "countryGroup.create":
		payload = &CountryRestrictionGroupEvent{}
	case "countryGroup.update":
		payload = &CountryRestrictionGroupEvent{}
	case "countryGroup.cancel":
		payload = &CountryRestrictionGroupEvent{}

	// Document (https://docs.treezor.com/guide/users/events.html#documents)
	case "document.create":
		payload = &DocumentEvent{}
	case "document.update":
		payload = &DocumentEvent{}
	case "document.cancel":
		payload = &DocumentEvent{}

	// Mandate (https://docs.treezor.com/guide/transfers/events.html#mandates)
	case "mandate.create":
		payload = &MandateEvent{}
	case "mandate.sign":
		payload = &MandateEvent{}
	case "mandate.cancel":
		payload = &MandateEvent{}

	// MCCGroup (https://docs.treezor.com/guide/cards/events.html#mcc-group)
	case "mccGroup.create":
		payload = &MCCRestrictionGroupEvent{}
	case "mccGroup.update":
		payload = &MCCRestrictionGroupEvent{}
	case "mccGroup.cancel":
		payload = &MCCRestrictionGroupEvent{}

	// MIDGroup (https://docs.treezor.com/guide/cards/events.html#mid-group)
	case "merchantIdGroup.create":
		payload = &MIDRestrictionGroupEvent{}
	case "merchantIdGroup.update":
		payload = &MIDRestrictionGroupEvent{}
	case "merchantIdGroup.cancel":
		payload = &MIDRestrictionGroupEvent{}

	// SepaSddr (https://docs.treezor.com/guide/transfers/events.html#sepa)
	case "sepa.return_sddr":
		payload = &SepaSddrEvent{}
	case "sepa.reject_sddr_core":
		payload = &SepaSddrEvent{}
	case "sepa.reject_sddr_b2b":
		payload = &SepaSddrEvent{}

	// RecalR (https://docs.treezor.com/guide/transfers/events.html#sepa)
	case "recallR.need_response":
		payload = &RecallREvent{}

	// Payin (https://docs.treezor.com/guide/cards/events.html#payin & https://docs.treezor.com/guide/cheques/events.html)
	case "payin.create":
		payload = &PayinEvent{}
	case "payin.update":
		payload = &PayinEvent{}
	case "payin.cancel":
		payload = &PayinEvent{}

	// PayinRefund (https://docs.treezor.com/guide/cheques/events.html)
	case "payinrefund.create":
		payload = &PayinRefundEvent{}
	case "payinrefund.update":
		payload = &PayinRefundEvent{}
	case "payinrefund.cancel":
		payload = &PayinRefundEvent{}

	// Payout (https://docs.treezor.com/guide/transfers/events.html#payins-payouts)
	case "payout.create":
		payload = &PayoutEvent{}
	case "payout.update":
		payload = &PayoutEvent{}
	case "payout.cancel":
		payload = &PayoutEvent{}

	// PayoutRefund (https://docs.treezor.com/guide/transfers/events.html#payins-payouts)
	case "payoutRefund.create":
		payload = &PayoutRefundEvent{}
	case "payoutRefund.update":
		payload = &PayoutRefundEvent{}
	case "payoutRefund.cancel":
		payload = &PayoutRefundEvent{}

	// Transactiona (https://docs.treezor.com/guide/transfers/events.html#transfers)
	case "transaction.create":
		payload = &TransactionEvent{}

	// Transfer (https://docs.treezor.com/guide/transfers/events.html#transfers)
	case "transfer.create":
		payload = &TransferEvent{}
	case "transfer.update":
		payload = &TransferEvent{}
	case "transfer.cancel":
		payload = &TransferEvent{}

	// TransferRefund (https://docs.treezor.com/guide/transfers/events.html#transfers)
	case "transferrefund.create":
		payload = &TransferRefundEvent{}
	case "transferrefund.update":
		payload = &TransferRefundEvent{}
	case "transferrefund.cancel":
		payload = &TransferRefundEvent{}

	// User (https://docs.treezor.com/guide/users/events.html#users)
	case "user.create":
		payload = &UserEvent{}
	case "user.update":
		payload = &UserEvent{}
	case "user.cancel":
		payload = &UserEvent{}
	case "user.kycrequest":
		payload = &UserEvent{}
	case "user.kycreview":
		payload = &UserEvent{}

	// KYCLiveness (https://docs.treezor.com/guide/users/events.html#kyc-liveness)
	case "kycliveness.create":
		payload = &KYCLivenessEvent{}
	case "kycliveness.update":
		payload = &KYCLivenessEvent{}

	// Wallet (https://docs.treezor.com/guide/wallets/events.html#wallets)
	case "wallet.create":
		payload = &WalletEvent{}
	case "wallet.update":
		payload = &WalletEvent{}
	case "wallet.cancel":
		payload = &WalletEvent{}

	// BankAccount (https://www.treezor.com/api-documentation/#/bankaccount)
	case "bankaccount.create":
		payload = &BankAccountEvent{}
	case "bankaccount.update":
		payload = &BankAccountEvent{}
	case "bankaccount.cancel":
		payload = &BankAccountEvent{}

	// SepaSdde (Undocumented, https://docs.treezor.com/guide/transfers/events.html#sepa)
	case "sepa.reject_sdde":

	// OneClickCard (Undocumented)
	case "oneclickcard.create":
	case "oneclickcard.update":
	case "oneclickcard.cancel":
	}

	err = json.Unmarshal(e.GetObjectPayload(), &payload)
	return payload, errors.WithStack(err)
}
