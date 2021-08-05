package treezor

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"

	json "github.com/tifo/treezor-sdk/internal/json"
	"github.com/tifo/treezor-sdk/internal/types"
)

// PayoutService handles communication with the payout related
// methods of the Treezor API.
//
// Treezor API docs: https://www.treezor.com/api-documentation/#/payout
type PayoutService service

// PayoutResponse represents a list of payouts.
// It may contain only one item.
type PayoutResponse struct {
	Payouts []*Payout `json:"payouts"`
}

// PayoutType defines the type of payout we're doing.
type PayoutType int32

const (
	// CreditTransferPayout is a payout type used for sepa transfer transaction.
	CreditTransferPayout PayoutType = 1
	// DirectDebitPayout is payout type used for direct debit transaction.
	DirectDebitPayout PayoutType = 2
)

func (t *PayoutType) UnmarshalJSON(data []byte) error {
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	v, err := str.Int64()
	if err != nil {
		return err
	}
	*t = PayoutType(v)
	return nil
}

// Payout represents a pay-out to a beneficiary.
type Payout struct {
	PayoutID               *types.Identifier `json:"payoutId,omitempty"`
	PayoutTag              *string           `json:"payoutTag,omitempty"`
	PayoutStatus           *string           `json:"payoutStatus,omitempty"` // NOTE: can be an enum
	PayoutTypeID           *PayoutType       `json:"payoutTypeId,omitempty"`
	PayoutType             *string           `json:"payoutType,omitempty"`
	WalletID               *types.Identifier `json:"walletId,omitempty"`
	PayoutDate             *types.Date       `json:"payoutDate,omitempty"`
	WalletEventName        *string           `json:"walletEventName,omitempty"`
	WalletAlias            *string           `json:"walletAlias,omitempty"`
	UserFirstname          *string           `json:"userFirstname,omitempty"`
	UserLastname           *string           `json:"userLastname,omitempty"`
	UserID                 *types.Identifier `json:"userId,omitempty"`
	BankAccountID          *types.Identifier `json:"bankaccountId,omitempty"`
	BeneficiaryID          *types.Identifier `json:"beneficiaryId,omitempty"`
	UniqueMandateReference *string           `json:"uniqueMandateReference,omitempty"`
	BankAccountIBAN        *string           `json:"bankaccountIBAN,omitempty"`
	Label                  *string           `json:"label,omitempty"`
	Amount                 *types.Amount     `json:"amount,omitempty"`
	Currency               *Currency         `json:"currency,omitempty"`
	PartnerFee             *types.Amount     `json:"partnerFee,omitempty"`
	CreatedDate            *time.Time        `json:"createdDate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	ModifiedDate           *time.Time        `json:"modifiedDate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	TotalRows              *types.Integer    `json:"totalRows,omitempty"`
	CodeStatus             *types.Identifier `json:"codeStatus,omitempty"`        // Legacy field
	InformationStatus      *string           `json:"informationStatus,omitempty"` // Legacy field
}

// Create creates a Treezor pay-out.
// The required field are WalletID, BeneficiaryID, Amount, Currency(ISO 4217).
func (s *PayoutService) Create(ctx context.Context, payout *Payout) (*Payout, *http.Response, error) {
	req, _ := s.client.NewRequest(http.MethodPost, "payouts", payout)

	b := new(PayoutResponse)
	resp, err := s.client.Do(ctx, req, b)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(b.Payouts) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one pay-out: %d pay-outs returned", len(b.Payouts))
	}
	return b.Payouts[0], resp, nil
}

// Get returns a pay-out.
func (s *PayoutService) Get(ctx context.Context, payoutID string) (*Payout, *http.Response, error) {
	u := fmt.Sprintf("payouts/%s", payoutID)
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	b := new(PayoutResponse)
	resp, err := s.client.Do(ctx, req, b)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(b.Payouts) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one pay-out: %d pay-outs returned", len(b.Payouts))
	}
	return b.Payouts[0], resp, nil
}

// PayoutListOptions specifies the optional parameters to the PayoutService.List.
type PayoutListOptions struct {
	PayoutStatus    string `url:"payoutStatus,omitempty"`
	UserID          string `url:"userId,omitempty"`
	WalletID        string `url:"walletId,omitempty"`
	PayoutTypeID    string `url:"payoutId,omitempty"`
	CreatedDateFrom string `url:"createdDateFrom,omitempty"`
	CreatedDateTo   string `url:"createdDateTo,omitempty"`

	ListOptions
}

// List the pay-outs for the authenticated user.
func (s *PayoutService) List(ctx context.Context, opt *PayoutListOptions) (*PayoutResponse, *http.Response, error) {
	u := "payouts"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	b := new(PayoutResponse)
	resp, err := s.client.Do(ctx, req, b)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	return b, resp, errors.WithStack(err)
}

// Delete deletes a payout. Change payout's status to CANCELED. A validated payout can't be cancelled.
func (s *PayoutService) Delete(ctx context.Context, payoutID string) (*Payout, *http.Response, error) {
	u := fmt.Sprintf("payouts/%s", payoutID)
	req, _ := s.client.NewRequest(http.MethodDelete, u, nil)

	b := new(PayoutResponse)
	resp, err := s.client.Do(ctx, req, b)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(b.Payouts) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one pay-out: %d pay-outs returned", len(b.Payouts))
	}
	return b.Payouts[0], resp, nil
}

// TODO: Update Payout API
