package treezor

import (
	"context"
	"net/http"

	"github.com/pkg/errors"

	"github.com/tifo/treezor-sdk/internal/types"
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
	WalletID          *types.Identifier     `json:"walletId,omitempty"`
	CurrentBalance    *types.Amount         `json:"currentBalance,omitempty"`
	Authorizations    *types.Amount         `json:"authorizations,omitempty"`
	AuthorizedBalance *types.Amount         `json:"authorizedBalance,omitempty"`
	Currency          *Currency             `json:"currency,omitempty"`
	CalculationDate   *types.TimestampParis `json:"calculationDate,omitempty"`
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
