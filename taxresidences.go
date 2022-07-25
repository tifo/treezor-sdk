package treezor

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/tifo/treezor-sdk/internal/types"
)

// TaxResidencesService handles the communication with
// tax residences related kyc methods
//
// Treezor API docs: https://www.treezor.com/api-documentation/#!/taxResidence/
type TaxResidencesService service

// TaxResidence represents a kyc TaxResidences
type TaxResidence struct {
	ID              *types.Identifier `json:"id,omitempty"`
	UserID          *types.Identifier `json:"userId,omitempty"`
	Country         *string           `json:"country,omitempty"`
	TaxPayerID      *string           `json:"taxPayerId,omitempty"`
	LiabilityWaiver *bool             `json:"liabilityWaiver,omitempty"`

	CreatedDate *time.Time `json:"createdDate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	LastUpdate  *time.Time `json:"lastUpdate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	DeletedDate *time.Time `json:"deletedDate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
}

// TaxResidencesResponse returns an array of TaxResidence. Array
// may contain only one item.
type TaxResidencesResponse struct {
	TaxResidences []*TaxResidence `json:"taxResidences,omitempty"`
}

type TaxResidenceCreateOptions struct {
	Access

	UserID          *string `url:"-" json:"userId,omitempty"`          // Required
	Country         *string `url:"-" json:"country,omitempty"`         // Required
	TaxPayerID      *string `url:"-" json:"taxPayerId,omitempty"`      // Optional
	LiabilityWaiver *bool   `url:"-" json:"liabilityWaiver,omitempty"` // Optional
}

// Create tax residences in Treezor.
func (s *TaxResidencesService) Create(ctx context.Context, opts *TaxResidenceCreateOptions) (*TaxResidence, *http.Response, error) {
	u := "taxResidences"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPost, u, opts)

	t := new(TaxResidencesResponse)
	resp, err := s.client.Do(ctx, req, t)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(t.TaxResidences) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one tax residence: %d tax residences returned", len(t.TaxResidences))
	}
	return t.TaxResidences[0], resp, nil
}

type TaxResidenceGetOptions struct {
	Access

	ID     *string `url:"-" json:"id,omitempty"`     // Optional
	UserID *string `url:"-" json:"userId,omitempty"` // Optional
}

// Get tax residences in Treezor.
func (s *TaxResidencesService) Get(ctx context.Context, opts *TaxResidenceGetOptions) (*TaxResidence, *http.Response, error) {
	u := "taxResidences"

	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, opts)

	t := new(TaxResidencesResponse)
	resp, err := s.client.Do(ctx, req, t)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(t.TaxResidences) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one tax residence: %d tax residences returned", len(t.TaxResidences))
	}
	return t.TaxResidences[0], resp, nil
}

type TaxResidenceEditOptions struct {
	Access

	UserID          *string `url:"-" json:"userId,omitempty"`          // Optional
	Country         *string `url:"-" json:"country,omitempty"`         // Optional
	TaxPayerID      *string `url:"-" json:"taxPayerId,omitempty"`      // Optional
	LiabilityWaiver *bool   `url:"-" json:"liabilityWaiver,omitempty"` // Optional
}

// Edit tax residences in Treezor.
func (s *TaxResidencesService) Edit(ctx context.Context, taxResidenceID string, opts *TaxResidenceEditOptions) (*TaxResidence, *http.Response, error) {
	u := fmt.Sprintf("taxResidences/%s", taxResidenceID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPut, u, opts)

	t := new(TaxResidencesResponse)
	resp, err := s.client.Do(ctx, req, t)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(t.TaxResidences) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one tax residence: %d tax residences returned", len(t.TaxResidences))
	}
	return t.TaxResidences[0], resp, nil
}

type TaxResidenceDeleteOptions struct {
	Access
}

// Delete tax residences in Treezor.
func (s *TaxResidencesService) Delete(ctx context.Context, taxResidenceID string, opts *TaxResidenceDeleteOptions) (*TaxResidence, *http.Response, error) {
	u := fmt.Sprintf("taxResidences/%s", taxResidenceID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodDelete, u, opts)

	t := new(TaxResidencesResponse)
	resp, err := s.client.Do(ctx, req, t)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(t.TaxResidences) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one tax residence: %d tax residences returned", len(t.TaxResidences))
	}
	return t.TaxResidences[0], resp, nil
}
