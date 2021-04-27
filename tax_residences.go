package treezor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// TaxResidencesService handles the communication with
// tax residences related kyc methods
//
// Treezor API docs: https://www.treezor.com/api-documentation/#!/taxResidence/
type TaxResidencesService service

// TaxResidence represents a kyc TaxResidences
type TaxResidence struct {
	ID              *int64  `json:"id,omitempty"`
	UserID          *int64  `json:"userId,omitempty"`
	Country         *string `json:"country,omitempty"`
	TaxPayerID      *string `json:"taxPayerId,omitempty"`
	LiabilityWaiver *bool   `json:"liabilityWaiver,omitempty"`
}

// TaxResidencesResponse returns an array of TaxResidence. Array
// may contain only one item.
type TaxResidencesResponse struct {
	TaxResidences []*TaxResidence `json:"taxResidences,omitempty"`
}

// Create tax residences in Treezor.
func (s *TaxResidencesService) Create(ctx context.Context, taxResidence *TaxResidence) (*TaxResidence, *http.Response, error) {
	c, _ := s.client.NewRequest(http.MethodPost, "taxResidences", taxResidence)

	t := new(TaxResidencesResponse)
	resp, err := s.client.Do(ctx, c, t)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(t.TaxResidences) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one tax residence: %d tax residences returned", len(t.TaxResidences))
	}
	return t.TaxResidences[0], resp, nil
}

// Edit updates a tax residences.
func (s *TaxResidencesService) Edit(ctx context.Context, taxResidenceID int64, taxResidence *TaxResidence) (*TaxResidence, *http.Response, error) {
	id := fmt.Sprintf("taxResidences/%d", taxResidenceID)
	c, _ := s.client.NewRequest(http.MethodPut, id, taxResidence)

	t := new(TaxResidencesResponse)
	resp, err := s.client.Do(ctx, c, t)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(t.TaxResidences) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one tax residence: %d tax residences returned", len(t.TaxResidences))
	}
	return t.TaxResidences[0], resp, nil
}
