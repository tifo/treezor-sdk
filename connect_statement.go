package treezor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type ConnectStatementService service

type ConnectStatementOptions struct {
	Month *string `url:"month" json:"-"` // Required, numeric representation of the month with leading 0 (Ex. 02)
	Year  *string `url:"year" json:"-"`  // Required, numeric representation of the year (Ex. 2021)
}

func (s *ConnectStatementService) Computed(ctx context.Context, walletID string, opts *ConnectStatementOptions) (interface{}, *http.Response, error) {

	u := fmt.Sprintf("v1/wallets/%s/statement/computed", walletID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	resp, err := s.client.Do(ctx, req, nil)

	if err != nil {
		return nil, resp, errors.WithStack(err)
	}
	return nil, resp, nil
}

func (s *ConnectStatementService) Raw(ctx context.Context, walletID string, opts *ConnectStatementOptions) (interface{}, *http.Response, error) {

	u := fmt.Sprintf("v1/wallet/%s/statement/raw", walletID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPost, u, opts)

	resp, err := s.client.Do(ctx, req, nil)

	if err != nil {
		return nil, resp, errors.WithStack(err)
	}
	return nil, resp, nil
}
