package treezor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/tifo/treezor-sdk/internal/types"
)

type ConnectStatementService service

type ConnectStatementOptions struct {
	Month *string `url:"month" json:"-"` // Required, numeric representation of the month with leading 0 (Ex. 02)
	Year  *string `url:"year" json:"-"`  // Required, numeric representation of the year (Ex. 2021)
}

type ConnectStatementListOptions struct{}

type StatementInfo struct {
	Year     int    `json:"year,omitempty"`
	Month    int    `json:"month,omitempty"`
	Raw      string `json:"raw,omitempty"`
	Computed string `json:"computed,omitempty"`
}

type ComputedStatement struct {
	Link     string `json:"link,omitempty"`
	ExpireIn int    `json:"expireIn,omitempty"`
}

type RawStatement struct {
	BalanceStart *StatementBalance `json:"balanceStart,omitempty"`
	BalanceEnd   *StatementBalance `json:"balanceEnd,omitempty"`
	Operations   []*Operation      `json:"oprations,omitempty"`
	User         *User             `json:"user,omitempty"`
	Wallet       *Wallet           `json:"wallet,omitempty"`
}

type StatementBalance struct {
	Amount    *int64      `json:"amount,omitempty"`
	Currency  *Currency   `json:"currency,omitempty"`
	Direction *string     `json:"direction,omitempty"`
	Date      *types.Date `json:"date,omitempty"`
}

func (s *ConnectStatementService) List(ctx context.Context, walletID string, opts *ConnectStatementListOptions) ([]*StatementInfo, *http.Response, error) {
	u := fmt.Sprintf("wallets/%s/statements", walletID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	var d []*StatementInfo
	resp, err := s.client.Do(ctx, req, d)

	if err != nil {
		return nil, resp, errors.WithStack(err)
	}
	return d, resp, nil
}

func (s *ConnectStatementService) Computed(ctx context.Context, walletID string, opts *ConnectStatementOptions) (*ComputedStatement, *http.Response, error) {

	u := fmt.Sprintf("wallets/%s/statement/computed", walletID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	d := new(ComputedStatement)
	resp, err := s.client.Do(ctx, req, d)

	if err != nil {
		return nil, resp, errors.WithStack(err)
	}
	return d, resp, nil
}

func (s *ConnectStatementService) Raw(ctx context.Context, walletID string, opts *ConnectStatementOptions) (*RawStatement, *http.Response, error) {

	u := fmt.Sprintf("wallet/%s/statement/raw", walletID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPost, u, opts)

	d := new(RawStatement)
	resp, err := s.client.Do(ctx, req, d)

	if err != nil {
		return nil, resp, errors.WithStack(err)
	}
	return d, resp, nil
}
