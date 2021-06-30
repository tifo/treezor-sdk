package treezor

import (
	"context"
	"encoding/json"
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

// WalletType defines the type of a wallet.
type WalletType int32

const (
	ElectronicMoneyWallet     WalletType = 9
	PaymentAccountWallet      WalletType = 10
	MirrorWallet              WalletType = 13
	ElectronicMoneyCardWallet WalletType = 14 // Internal Only
)

func (t *WalletType) UnmarshalJSON(data []byte) error {
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	v, err := str.Int64()
	if err != nil {
		return err
	}
	*t = WalletType(v)
	return nil
}

// Wallet represents a Treezor wallet.
type Wallet struct {
	WalletID            *types.Identifier     `json:"walletId,omitempty"`
	WalletTypeID        *WalletType           `json:"walletTypeId,omitempty"`
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
	Currency            *Currency             `json:"currency,omitempty"`
	CreatedDate         *types.TimestampParis `json:"createdDate,omitempty"`
	ModifiedDate        *types.TimestampParis `json:"modifiedDate,omitempty"`
	PayinCount          *types.Integer        `json:"payinCount,omitempty"`
	PayoutCount         *types.Integer        `json:"payoutCount,omitempty"`
	TransferCount       *types.Integer        `json:"transferCount,omitempty"`
	Solde               *types.Amount         `json:"solde,omitempty"`
	AuthorizedBalance   *types.Amount         `json:"authorizedBalance,omitempty"`
	TotalRows           *types.Integer        `json:"totalRows,omitempty"`
	CodeStatus          *types.Identifier     `json:"codeStatus,omitempty"`        // Legacy field
	InformationStatus   *string               `json:"informationStatus,omitempty"` // Legacy field
}

type WalletCreateOptions struct {
	Access

	WalletTypeID        *WalletType `url:"-" json:"walletTypeId"`                  // Required
	TariffID            *string     `url:"-" json:"tariffId"`                      // Required
	UserID              *string     `url:"-" json:"userId"`                        // Required
	JointUserID         *string     `url:"-" json:"jointUserId,omitempty"`         // Optional
	WalletTag           *string     `url:"-" json:"walletTag,omitempty"`           // Optional
	Currency            *Currency   `url:"-" json:"currency"`                      // Required
	EventName           *string     `url:"-" json:"eventName,omitempty"`           // Optional
	EventAlias          *string     `url:"-" json:"eventAlias,omitempty"`          // Optional
	EventDate           *types.Date `url:"-" json:"eventDate,omitempty"`           // Optional
	EventMessage        *string     `url:"-" json:"eventMessage,omitempty"`        // Optional
	EventPayinStartDate *types.Date `url:"-" json:"eventPayinStartDate,omitempty"` // Deprecated, Optional
	EventPayinEndDate   *types.Date `url:"-" json:"eventPayinEndDate,omitempty"`   // Deprecated, Optional
}

// Create creates a Treezor wallet.
func (s *WalletService) Create(ctx context.Context, opts *WalletCreateOptions) (*Wallet, *http.Response, error) {
	u := "wallets"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPost, u, opts)

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

type WalletGetOptions struct {
	Access
}

// Get fetches a wallet from Treezor.
func (s *WalletService) Get(ctx context.Context, walletID string, opts *WalletGetOptions) (*Wallet, *http.Response, error) {
	u := fmt.Sprintf("wallets/%s", walletID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
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
	Access

	WalletID            *string          `url:"walletId,omitempty" json:"-"`
	WalletStatus        *string          `url:"walletStatus,omitempty" json:"-"` // NOTE: can be an enum (need to see if VALIDATED or Validated)
	UserID              *string          `url:"userId,omitempty" json:"-"`
	ParentUserID        *string          `url:"parentUserId,omitempty" json:"-"`
	WalletTag           *string          `url:"walletTag,omitempty" json:"-"`
	WalletTypeID        *WalletType      `url:"walletTypeId,omitempty" json:"-"`
	EventAlias          *string          `url:"eventAlias,omitempty" json:"-"`
	EventPayinStartDate *types.Timestamp `url:"eventPayinStartDate,omitempty" json:"-"`
	EventPayinEndDate   *types.Date      `url:"eventPayinEndDate,omitempty" json:"-"`
	TariffID            *string          `url:"tariffId,omitempty" json:"-"`
	PayinCount          *int             `url:"payinCount,omitempty" json:"-"`

	ListOptions
}

// List returns a list of wallets.
func (s *WalletService) List(ctx context.Context, opts *WalletListOptions) (*WalletResponse, *http.Response, error) {
	u := "wallets"
	u, err := addOptions(u, opts)
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

type WalletEditOptions struct {
	Access

	WalletTypeID        *WalletType `url:"-" json:"walletTypeId,omitempty"`        // Optional
	EventName           *string     `url:"-" json:"eventName,omitempty"`           // Optional
	EventAlias          *string     `url:"-" json:"eventAlias,omitempty"`          // Optional
	EventDate           *types.Date `url:"-" json:"eventDate,omitempty"`           // Optional
	EventMessage        *string     `url:"-" json:"eventMessage,omitempty"`        // Optional
	EventPayinStartDate *types.Date `url:"-" json:"eventPayinStartDate,omitempty"` // Deprecated, Optional
	EventPayinEndDate   *types.Date `url:"-" json:"eventPayinEndDate,omitempty"`   // Deprecated, Optional
	URLImage            *string     `url:"-" json:"urlImage,omitempty"`            // Optional
	ImageName           *string     `url:"-" json:"imageName,omitempty"`           // Optional
	TariffID            *string     `url:"-" json:"tariffId"`                      // Required
}

// Edit updates a wallet.
func (s *WalletService) Edit(ctx context.Context, walletID string, opts *WalletEditOptions) (*Wallet, *http.Response, error) {
	u := fmt.Sprintf("wallets/%s", walletID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPut, u, opts)

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
	Access

	Origin *Origin `url:"origin"` // Required
}

// Cancel makes a Wallet cancelled, meaning all future operation for that wallet
// will be refused.
func (s *WalletService) Cancel(ctx context.Context, walletID string, opts *WalletCancelOptions) (*Wallet, *http.Response, error) {
	u := fmt.Sprintf("wallets/%s", walletID)
	u, err := addOptions(u, opts)
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
