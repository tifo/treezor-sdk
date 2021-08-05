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

// BalanceSerachOptions specifies the optional parameters to the BalanceService.Search.
type BalanceSerachOptions struct {
	Access

	WalletID string `url:"walletId,omitempty"`
	UserID   string `url:"userId,omitempty"`

	ListOptions
}

// Search the balances for the authenticated user. If WalletID is provided,
// list one balance for the specified wallet; if UserID is provided, list all
// the balances for the user's wallets.
// See https://www.treezor.com/api-documentation/#/balance/getBalances
func (s *BalanceService) Search(ctx context.Context, opts *BalanceSerachOptions) (*BalanceResponse, *http.Response, error) {
	u := "balances"
	u, err := addOptions(u, opts)
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

// BalanceListOptions specifies the optional parameters to the BalanceService.List.
type BalanceListOptions struct {
	Access
	ListOptions
}

// List the balances for all the wallets of the specified userID.
func (s *BalanceService) List(ctx context.Context, userID string, opts *BalanceListOptions) ([]*Balance, *http.Response, error) {

	searchOpts := &BalanceSerachOptions{
		Access:      opts.Access,
		ListOptions: opts.ListOptions,
		UserID:      userID,
	}

	b, resp, err := s.Search(ctx, searchOpts)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	return b.Balances, resp, errors.WithStack(err)
}

// BalanceGetOptions specifies the optional parameters to the BalanceService.Get.
type BalanceGetOptions struct {
	Access
}

// Get the balances for the specified wallet.
func (s *BalanceService) Get(ctx context.Context, walletID string, opts *BalanceGetOptions) (*Balance, *http.Response, error) {

	searchOpts := &BalanceSerachOptions{
		Access:   opts.Access,
		WalletID: walletID,
	}

	b, resp, err := s.Search(ctx, searchOpts)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if l := len(b.Balances); l != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one balance: %d balance returned", l)
	}

	return b.Balances[0], resp, nil
}
