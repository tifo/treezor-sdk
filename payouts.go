package treezor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
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

// Payout represents a pay-out to a beneficiary.
type Payout struct {
	Access
	PayoutID               *string         `json:"payoutId,omitempty"`
	PayoutTag              *string         `json:"payoutTag,omitempty"`
	PayoutStatus           *string         `json:"payoutStatus,omitempty"`
	PayoutTypeID           *string         `json:"payoutTypeId,omitempty"`
	PayoutType             *string         `json:"payoutType,omitempty"`
	WalletID               *string         `json:"walletId,omitempty"`
	PayoutDate             *Date           `json:"payoutDate,omitempty"`
	WalletEventName        *string         `json:"walletEventName,omitempty"`
	WalletAlias            *string         `json:"walletAlias,omitempty"`
	UserFirstname          *string         `json:"userFirstname,omitempty"`
	UserLastname           *string         `json:"userLastname,omitempty"`
	UserID                 *string         `json:"userId,omitempty"`
	BeneficiaryID          *string         `json:"beneficiaryId,omitempty"`
	UniqueMandateReference *string         `json:"uniqueMandateReference,omitempty"`
	Label                  *string         `json:"label,omitempty"`
	Amount                 *float64        `json:"amount,string,omitempty"`
	Currency               Currency        `json:"currency,omitempty"`
	PartnerFee             *float64        `json:"partnerFee,string,omitempty"`
	CreatedDate            *TimestampParis `json:"createdDate,omitempty"`
	ModifiedDate           *TimestampParis `json:"modifiedDate,omitempty"`
	TotalRows              *int64          `json:"totalRows,string,omitempty"`
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
