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

// TransferService handles communication with the transfer related
// methods of the Treezor API.
//
// Treezor API docs: https://www.treezor.com/api-documentation/#/transfer
type TransferService service

// TransferResponse represents a list of transfers.
// It may contain only one item.
type TransferResponse struct {
	Transfers []*Transfer `json:"transfers"`
}

// TransferType defines the type of transfer we're doing.
type TransferType int32

const (
	// Wallet2WalletTransfer is a transfer type used for peer-to-peer transaction.
	Wallet2WalletTransfer TransferType = 1
	// ClientFeesTransfer is transfer type used for client fees.
	ClientFeesTransfer TransferType = 3
	// CreditNoteTransfer is a transfer type used for credit note.
	CreditNoteTransfer TransferType = 4
)

func (t *TransferType) UnmarshalJSON(data []byte) error {
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	v, err := str.Int64()
	if err != nil {
		return err
	}
	*t = TransferType(v)
	return nil
}

// Transfer represents a transfer.
type Transfer struct {
	TransferID                 *types.Identifier `json:"transferId,omitempty"`
	TransferTypeID             *TransferType     `json:"transferTypeId,omitempty"`
	TransferTag                *string           `json:"transferTag,omitempty"`
	TransferStatus             *string           `json:"transferStatus,omitempty"` // TODO: can be an enum
	WalletID                   *types.Identifier `json:"walletId,omitempty"`
	ForeignID                  *types.Identifier `json:"foreignId,omitempty"`
	WalletTypeID               *types.Identifier `json:"walletTypeId,omitempty"` // TODO: can be an enum
	BeneficiaryWalletID        *types.Identifier `json:"beneficiaryWalletId,omitempty"`
	BeneficiaryWalletTypeID    *types.Identifier `json:"beneficiaryWalletTypeId,omitempty"` // TODO: can be an enum
	TransferDate               *types.Date       `json:"transferDate,omitempty"`
	Amount                     *types.Amount     `json:"amount,omitempty"`
	Currency                   *Currency         `json:"currency,omitempty"`
	Label                      *string           `json:"label,omitempty"`
	PartnerFee                 *types.Amount     `json:"partnerFee,omitempty"`
	WalletEventName            *string           `json:"walletEventName,omitempty"`
	WalletAlias                *string           `json:"walletAlias,omitempty"`
	BeneficiaryWalletEventName *string           `json:"beneficiaryWalletEventName,omitempty"`
	BeneficiaryWalletAlias     *string           `json:"beneficiaryWalletAlias,omitempty"`
	CreatedDate                *time.Time        `json:"createdDate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	ModifiedDate               *time.Time        `json:"modifiedDate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	TotalRows                  *types.Integer    `json:"totalRows,omitempty"`
	CodeStatus                 *types.Identifier `json:"codeStatus,omitempty"`        // Legacy field
	InformationStatus          *string           `json:"informationStatus,omitempty"` // Legacy field
}

type TransferCreateOptions struct {
	Access

	WalletID            *string       `url:"-" json:"walletId,omitempty"`            // Required
	BeneficiaryWalletID *string       `url:"-" json:"beneficiaryWalletId,omitempty"` // Required
	Amount              *float64      `url:"-" json:"amount,omitempty"`              // Required
	Currency            Currency      `url:"-" json:"currency,omitempty"`            // Required
	Label               *string       `url:"-" json:"label,omitempty"`               // Optional
	TransferTypeID      *TransferType `url:"-" json:"transferTypeId,omitempty"`      // Optional // TODO: create enum
}

// Create creates a Treezor transfer. Required: WalletID, BeneficiaryWalletID,Amount,Currency(ISO 4217)
func (s *TransferService) Create(ctx context.Context, opts *TransferCreateOptions) (*Transfer, *http.Response, error) {
	u := "transfers"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPost, u, opts)

	b := new(TransferResponse)
	resp, err := s.client.Do(ctx, req, b)
	if err != nil {
		return nil, resp, err
	}

	if len(b.Transfers) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one transfer: %d transfers returned", len(b.Transfers))
	}
	return b.Transfers[0], resp, nil
}

// Get returns a transfer.
func (s *TransferService) Get(ctx context.Context, transferID string) (*Transfer, *http.Response, error) {
	u := fmt.Sprintf("transfers/%s", transferID)
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	b := new(TransferResponse)
	resp, err := s.client.Do(ctx, req, b)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(b.Transfers) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one transfer: %d transfers returned", len(b.Transfers))
	}
	return b.Transfers[0], resp, nil
}

// TransferListOptions specifies the optional parameters to the TransferService.List.
type TransferListOptions struct {
	UserID              string `url:"userId,omitempty"`
	BeneficiaryUserID   string `url:"beneficiaryUserId,omitempty"`
	WalletID            string `url:"walletId,omitempty"`
	BeneficiaryWalletID string `url:"beneficiaryWalletId,omitempty"`
	TransferStatus      string `url:"transferStatus,omitempty"`
	TransferTypeID      string `url:"transferTypeId,omitempty"`
	TransferTag         string `url:"transferTag,omitempty"`
	Label               string `url:"label,omitempty"`
	CreatedDateFrom     string `url:"createdDateFrom,omitempty"`
	CreatedDateTo       string `url:"createdDateTo,omitempty"`

	ListOptions
}

// List the transfers for the authenticated user.s
func (s *TransferService) List(ctx context.Context, opt *TransferListOptions) (*TransferResponse, *http.Response, error) {
	u := "transfers"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	b := new(TransferResponse)
	resp, err := s.client.Do(ctx, req, b)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	return b, resp, errors.WithStack(err)
}

// Delete deletes a transfer. Change transfer's status to CANCELED. A validated transfer can't be cancelled.
func (s *TransferService) Delete(ctx context.Context, transferID string) (*Transfer, *http.Response, error) {
	u := fmt.Sprintf("transfers/%s", transferID)
	req, _ := s.client.NewRequest(http.MethodDelete, u, nil)

	b := new(TransferResponse)
	resp, err := s.client.Do(ctx, req, b)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(b.Transfers) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one transfer: %d transfers returned", len(b.Transfers))
	}
	return b.Transfers[0], resp, nil
}

// TODO: Update Transfer API
