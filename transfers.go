package treezor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// TransferType defines the type of transfer we're doing.
type TransferType string

const (
	// Wallet2WalletTransfer is a transfer type used for peer-to-peer transaction.
	Wallet2WalletTransfer TransferType = "1"
	// ClientFeesTransfer is transfer type used for client fees.
	ClientFeesTransfer TransferType = "3"
	// CreditNoteTransfer is a transfer type used for credit note.
	CreditNoteTransfer TransferType = "4"
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

// Transfer represents a transfer.
type Transfer struct {
	Access
	TransferID                 *string         `json:"transferId,omitempty"`
	TransferStatus             *string         `json:"transferStatus,omitempty"`
	TransferTypeID             TransferType    `json:"transferTypeId,omitempty"`
	TransferTag                *string         `json:"transferTag,omitempty"`
	WalletID                   *string         `json:"walletId,omitempty"`
	WalletTypeID               *string         `json:"walletTypeId,omitempty"`
	BeneficiaryWalletID        *string         `json:"beneficiaryWalletId,omitempty"`
	BeneficiaryWalletTypeID    *string         `json:"beneficiaryWalletTypeId,omitempty"`
	TransferDate               *Date           `json:"transferDate,omitempty"`
	WalletEventName            *string         `json:"walletEventName,omitempty"`
	WalletAlias                *string         `json:"walletAlias,omitempty"`
	BeneficiaryWalletEventName *string         `json:"beneficiaryWalletEventName,omitempty"`
	BeneficiaryWalletAlias     *string         `json:"beneficiaryWalletAlias,omitempty"`
	Amount                     *float64        `json:"amount,string,omitempty"`
	Currency                   Currency        `json:"currency,omitempty"`
	Label                      *string         `json:"label,omitempty"`
	CreatedDate                *TimestampParis `json:"createdDate,omitempty"`
	ModifiedDate               *TimestampParis `json:"modifiedDate,omitempty"`
	TotalRows                  *int64          `json:"totalRows,string,omitempty"`
}

// Create creates a Treezor transfer. Required: WalletID, BeneficiaryWalletID,Amount,Currency(ISO 4217)
func (s *TransferService) Create(ctx context.Context, transfer *Transfer) (*Transfer, *http.Response, error) {
	req, _ := s.client.NewRequest(http.MethodPost, "transfers", transfer)

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
