package treezor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/tifo/treezor-sdk/types"
)

// WalletService handles communication with the wallet related
// methods of the Treezor API.
//
// Treezor API docs: https://www.treezor.com/api-documentation/#/wallet
type WalletService service

// WalletResponse represents a list of wallets.
// It may contain only one item.
type WalletResponse struct {
	Wallets []*Wallet `json:"wallets"`
}

// Wallet represents a Treezor wallet.
type Wallet struct {
	WalletID            *types.Identifier     `json:"walletId,omitempty"`
	WalletTypeID        *types.Identifier     `json:"walletTypeId,omitempty"` // NOTE: Can be an enum
	WalletStatus        *string               `json:"walletStatus,omitempty"` // NOTE: Can be an enum
	WalletTag           *string               `json:"walletTag,omitempty"`
	UserID              *types.Identifier     `json:"userId,omitempty"`
	UserFirstname       *string               `json:"userFirstname,omitempty"`
	UserLastname        *string               `json:"userLastname,omitempty"`
	JointUserID         *types.Identifier     `json:"jointUserId,omitempty"`
	TariffID            *types.Identifier     `json:"tariffId,omitempty"`
	EventName           *string               `json:"eventName,omitempty"`
	EventAlias          *string               `json:"eventAlias,omitempty"`
	EventMessage        *string               `json:"eventMessage,omitempty"`
	EventDate           *types.Date           `json:"eventDate,omitempty"`
	EventPayinEndDate   *types.Date           `json:"eventPayinEndDate,omitempty"`
	EventPayinStartDate *types.Date           `json:"eventPayinStartDate,omitempty"`
	ContractSigned      *types.Boolean        `json:"contractSigned,omitempty"`
	BIC                 *string               `json:"bic,omitempty"`
	IBAN                *string               `json:"iban,omitempty"`
	URLImage            *string               `json:"urlImage,omitempty"`
	Currency            *types.Currency       `json:"currency,omitempty"`
	CreatedDate         *types.TimestampParis `json:"createdDate,omitempty"`
	ModifiedDate        *types.TimestampParis `json:"modifiedDate,omitempty"`
	PayinCount          *types.Integer        `json:"payinCount,omitempty"`
	PayoutCount         *types.Integer        `json:"payoutCount,omitempty"`
	TransferCount       *types.Integer        `json:"transferCount,omitempty"`
	Solde               *types.Amount         `json:"solde,omitempty"`
	AuthorizedBalance   *types.Amount         `json:"authorizedBalance,omitempty"`
	TotalRows           *types.Integer        `json:"totalRows,omitempty"`
	CodeStatus          *types.Identifier     `json:"codeStatus,omitempty"`        // NOTE: Legacy + Webhook
	InformationStatus   *string               `json:"informationStatus,omitempty"` // NOTE: Legacy + Webhook
}

// Create creates a Treezor wallet.
func (s *WalletService) Create(ctx context.Context, wallet *Wallet) (*Wallet, *http.Response, error) {
	req, _ := s.client.NewRequest(http.MethodPost, "wallets", wallet)

	w := new(WalletResponse)
	resp, err := s.client.Do(ctx, req, w)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(w.Wallets) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one wallet: %d wallets returned", len(w.Wallets))
	}
	return w.Wallets[0], resp, nil
}

// Get fetches a wallet from Treezor.
func (s *WalletService) Get(ctx context.Context, walletID string) (*Wallet, *http.Response, error) {
	u := fmt.Sprintf("wallets/%s", walletID)
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	w := new(WalletResponse)
	resp, err := s.client.Do(ctx, req, w)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(w.Wallets) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one wallet: %d wallets returned", len(w.Wallets))
	}
	return w.Wallets[0], resp, nil
}

// WalletListOptions contains options for listing wallets.
type WalletListOptions struct {
	WalletID            string               `url:"walletId,omitempty"`
	WalletStatus        string               `url:"walletStatus,omitempty"`
	UserID              string               `url:"userId,omitempty"`
	ParentUserID        string               `url:"parentUserId,omitempty"`
	WalletTag           string               `url:"walletTag,omitempty"`
	WalletTypeID        string               `url:"walletTypeId,omitempty"`
	EventAlias          string               `url:"eventAlias,omitempty"`
	EventPayinStartDate types.TimestampParis `url:"eventPayinStartDate,omitempty"`
	EventPayinEndDate   types.Date           `url:"eventPayinEndDate,omitempty"`
	TariffID            string               `url:"tariffId,omitempty"`
	PayinCount          int                  `url:"payinCount,omitempty"`

	ListOptions
}

// List returns a list of wallets.
func (s *WalletService) List(ctx context.Context, opt *WalletListOptions) (*WalletResponse, *http.Response, error) {
	u := "wallets"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	w := new(WalletResponse)
	resp, err := s.client.Do(ctx, req, w)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	return w, resp, errors.WithStack(err)
}

// Edit updates a wallet.
func (s *WalletService) Edit(ctx context.Context, walletID string, wallet *Wallet) (*Wallet, *http.Response, error) {
	u := fmt.Sprintf("wallets/%s", walletID)
	req, _ := s.client.NewRequest(http.MethodPut, u, wallet)

	w := new(WalletResponse)
	resp, err := s.client.Do(ctx, req, w)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(w.Wallets) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one wallet: %d wallets returned", len(w.Wallets))
	}
	return w.Wallets[0], resp, nil
}

// WalletCancelOptions contains options for deletion of a wallet.
// Origin can be of value OPERATOR or USER.
type WalletCancelOptions struct {
	Origin Origin `url:"origin,omitempty"`
}

// Cancel makes a User cancelled, meaning all future operation for that wallet
// will be refused.
func (s *WalletService) Cancel(ctx context.Context, walletID string, opt *WalletCancelOptions) (*Wallet, *http.Response, error) {
	u := fmt.Sprintf("wallets/%s", walletID)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	req, _ := s.client.NewRequest(http.MethodDelete, u, nil)

	w := new(WalletResponse)
	resp, err := s.client.Do(ctx, req, w)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(w.Wallets) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one wallet: %d wallets returned", len(w.Wallets))
	}
	return w.Wallets[0], resp, nil
}
