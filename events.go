package treezor

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Event represent an event that occurred on Treezor's infrastructure
// and is sent as a webhook to us.
type Event struct {
	ID               *string          `json:"webhook_id,omitempty"`
	Type             *string          `json:"webhook,omitempty"`
	Object           *string          `json:"object,omitempty"`
	ObjectID         *string          `json:"object_id,omitempty"`
	RawPayload       *json.RawMessage `json:"object_payload,omitempty"`
	PayloadSignature *string          `json:"object_payload_signature,omitempty"`
}

func (e Event) String() string {
	return Stringify(e)
}

// ParsePayload parses the event payload. For recognized event types,
// a value of the corresponding struct type will be returned.
func (e *Event) ParsePayload() (payload interface{}, err error) {
	switch e.GetType() {
	case "balance.update":
	case "bankaccount.create":
	case "bankaccount.update":
	case "bankaccount.cancel":
	case "beneficiary.create":
	case "beneficiary.update":
	case "card.requestphysical":
		payload = &CardRequestPhysicalEvent{}
	case "card.createvirtual":
		payload = &CardCreateVirtualEvent{}
	case "card.convertvirtual":
		payload = &CardConvertVirtualEvent{}
	case "card.changepin":
		payload = &CardChangePINEvent{}
	case "card.activate":
		payload = &CardActivateEvent{}
	case "card.renew":
		payload = &CardRenewEvent{}
	case "card.regenerate":
		payload = &CardRegenerateEvent{}
	case "card.update":
		payload = &CardUpdateEvent{}
	case "card.limits":
		payload = &CardLimitsEvent{}
	case "card.options":
		payload = &CardOptionsEvent{}
	case "card.setpin":
		payload = &CardSetPINEvent{}
	case "card.unblockpin":
		payload = &CardUnblockPINEvent{}
	case "card.lockunlock":
		payload = &CardLockUnlockEvent{}
	case "cardDigitalization.update":
	case "cardtransaction.create":
		payload = &CardTransactionCreateEvent{}
	case "countryGroup.create":
	case "countryGroup.update":
	case "countryGroup.cancel":
	case "document.create":
	case "document.update":
	case "document.cancel":
	case "mandate.create":
	case "mandate.sign":
	case "mandate.cancel":
	case "merchantIdGroup.create":
	case "mccGroup.create":
	case "mccGroup.cancel":
	case "mccGroup.update":
	case "oneclickcard.create":
	case "oneclickcard.update":
	case "oneclickcard.cancel":
	case "payin.create":
		payload = &PayinCreateEvent{}
	case "payin.update":
		payload = &PayinUpdateEvent{}
	case "payin.cancel":
		payload = &PayinCancelEvent{}
	case "payinrefund.create":
	case "payinrefund.update":
	case "payinrefund.cancel":
	case "payout.create":
		payload = &PayoutCreateEvent{}
	case "payout.update":
		payload = &PayoutUpdateEvent{}
	case "payout.cancel":
		payload = &PayoutCancelEvent{}
	case "sepa.return_sctr":
	case "sepa.reject_sddr_core":
		payload = &SepaSddrCoreRejectEvent{}
	case "sepa.reject_sddr_b2b":
	case "transaction.create":
	case "transfer.create":
	case "transfer.update":
		payload = &TransferUpdateEvent{}
	case "transfer.cancel":
	case "transferrefund.create":
	case "transferrefund.update":
	case "transferrefund.cancel":
	case "user.create":
		payload = &UserCreateEvent{}
	case "user.update":
		payload = &UserUpdateEvent{}
	case "user.cancel":
		payload = &UserCancelEvent{}
	case "user.kycrequest":
		payload = &UserKYCRequestEvent{}
	case "user.kycreview":
		payload = &UserKYCReviewEvent{}
	case "wallet.create":
		payload = &WalletCreateEvent{}
	case "wallet.update":
		payload = &WalletUpdateEvent{}
	case "wallet.cancel":
		payload = &WalletCancelEvent{}
	case "kycliveness.create":
		payload = &KycLivenessCreateEvent{}
	case "kycliveness.update":
		payload = &KycLivenessUpdateEvent{}
	}
	err = json.Unmarshal(e.GetRawPayload(), &payload)
	return payload, errors.WithStack(err)
}
