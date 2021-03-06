package treezor

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

// BalanceService handles communication with the balance related
// methods of the Treezor API.
//
// Treezor API docs: https://www.treezor.com/api-documentation/#/balance
type BalanceService service

// BalanceResponse represents a list of balances on multiple wallets.
// It may contain only one item.
type BalanceResponse struct {
	Balances []*Balance `json:"balances"`
}

// Balance represents the balance on a wallet.
type Balance struct {
	Access
	WalletID          *string         `json:"walletId,omitempty"`
	CurrentBalance    *float64        `json:"currentBalance,string,omitempty"`
	Authorizations    *float64        `json:"authorizations,string,omitempty"`
	AuthorizedBalance *float64        `json:"authorizedBalance,string,omitempty"`
	Currency          Currency        `json:"currency,omitempty"`
	CalculationDate   *TimestampParis `json:"calculationDate,omitempty"`
}

// BalanceOptions specifies the optional parameters to the BalanceService.List.
type BalanceOptions struct {
	ListOptions

	WalletID string `url:"walletId,omitempty"`
	UserID   string `url:"userId,omitempty"`
}

// List the balances for the authenticated user. If WalletID is provided,
// list one balance for the specified wallet; if UserID is provided, list all
// the balances for the user's wallets.
func (s *BalanceService) List(ctx context.Context, opt *BalanceOptions) (*BalanceResponse, *http.Response, error) {
	u := "balances"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	b := new(BalanceResponse)
	resp, err := s.client.Do(ctx, req, b)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	return b, resp, errors.WithStack(err)
}
