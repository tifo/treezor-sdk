package treezor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
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
	Access
	WalletID          *string         `json:"walletId,omitempty"`
	WalletTypeID      *string         `json:"walletTypeId,omitempty"`
	WalletStatus      *string         `json:"walletStatus,omitempty"`
	WalletTag         *string         `json:"walletTag,omitempty"`
	UserID            *string         `json:"userId,omitempty"`
	Name              *string         `json:"eventName,omitempty"`
	Description       *string         `json:"eventMessage,omitempty"`
	Date              *Date           `json:"eventDate,omitempty"`
	PayinEndDate      *Date           `json:"eventPayinEndDate,omitempty"`
	PayinStartDate    *Date           `json:"eventPayinStartDate,omitempty"`
	Currency          Currency        `json:"currency,omitempty"`
	JointUserID       *string         `json:"jointUserId,omitempty"`
	Alias             *string         `json:"eventAlias,omitempty"`
	ContractSigned    *string         `json:"contractSigned,omitempty"`
	URLImage          *string         `json:"urlImage,omitempty"`
	CreatedDate       *TimestampParis `json:"createdDate,omitempty"`
	ModifiedDate      *TimestampParis `json:"modifiedDate,omitempty"`
	UserFirstname     *string         `json:"userFirstname,omitempty"`
	UserLastname      *string         `json:"userLastname,omitempty"`
	CodeStatus        *string         `json:"codeStatus,omitempty"`
	TariffID          *string         `json:"tariffId,omitempty"`
	InformationStatus *string         `json:"informationStatus,omitempty"`
	PayinCount        *int64          `json:"payinCount,string,omitempty"`
	PayoutCount       *int64          `json:"payoutCount,string,omitempty"`
	TransferCount     *int64          `json:"transferCount,string,omitempty"`
	Solde             *float64        `json:"solde,string,omitempty"`
	AuthorizedBalance *float64        `json:"authorizedBalance,string,omitempty"`
	BIC               *string         `json:"bic,omitempty"`
	IBAN              *string         `json:"iban,omitempty"`
	TotalRows         *int64          `json:"totalRows,omitempty"`
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
