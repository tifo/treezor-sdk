package treezor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/tifo/treezor-sdk/internal/types"
)

// BeneficiaryService handles communication with the beneficiary related
// methods of the Treezor API.
//
// Treezor API docs: https://www.treezor.com/api-documentation/#/beneficiary
type BeneficiaryService service

// BeneficiaryResponse represents a list of beneficiaries.
// It may contain only one item.
type BeneficiaryResponse struct {
	Beneficiaries []*Beneficiary `json:"beneficiaries"`
}

// Beneficiary represents a beneficiary.
type Beneficiary struct {
	BeneficiaryID                      *types.Identifier     `json:"id,omitempty"`
	Tag                                *string               `json:"tag,omitempty"`
	UserID                             *types.Identifier     `json:"userId,omitempty"`
	NickName                           *string               `json:"nickName,omitempty"`
	Name                               *string               `json:"name,omitempty"`
	Address                            *string               `json:"address,omitempty"`
	IBAN                               *string               `json:"iban,omitempty"`
	BIC                                *string               `json:"bic,omitempty"`
	SepaCreditorIdentifier             *string               `json:"sepaCreditorIdentifier,omitempty"`
	SDDB2BWhitelist                    []*SDDB2BWhitelist    `json:"sddB2bWhitelist,omitempty"`
	SDDCoreBlacklist                   []string              `json:"sddCoreBlacklist,omitempty"`
	UsableForSCT                       *types.Boolean        `json:"usableForSct,omitempty"`
	SDDCoreKnownUniqueMandateReference []string              `json:"sddCoreKnownUniqueMandateReference,omitempty"`
	IsActive                           *types.Boolean        `json:"isActive,omitempty"`
	CreatedDate                        *types.TimestampParis `json:"createdDate,omitempty"`
	ModifiedDate                       *types.TimestampParis `json:"modifiedDate,omitempty"`
	// NOTE: see about totalRows, codeStatus and informationStatus
}

// SDDB2BWhitelist is a whitelisted company for B2B SEPA mandates.
type SDDB2BWhitelist struct {
	UniqueMandateReference *string           `json:"uniqueMandateReference,omitempty"`
	IsRecurrent            *types.Boolean    `json:"isRecurrent,omitempty"`
	WalletID               *types.Identifier `json:"walletId,omitempty"`
}

// BeneficiaryRequest represents a request to create/edit a beneficiary.
// It is separate from Beneficiary above because otherwise IDs
// fail to serialize to the correct JSON.
type BeneficiaryRequest struct {
	Access
	BeneficiaryID                      *string            `json:"id,omitempty"`
	UserID                             *string            `json:"userId,omitempty"`
	Tag                                *string            `json:"tag,omitempty"`
	NickName                           *string            `json:"nickName,omitempty"`
	Name                               *string            `json:"name,omitempty"`
	Address                            *string            `json:"address,omitempty"`
	IBAN                               *string            `json:"iban,omitempty"`
	BIC                                *string            `json:"bic,omitempty"`
	SepaCreditorIdentifier             *string            `json:"sepaCreditorIdentifier,omitempty"`
	UsableForSCT                       *bool              `json:"usableForSct,omitempty"`
	SDDCoreKnownUniqueMandateReference *[]string          `json:"sddCoreKnownUniqueMandateReference,omitempty"`
	SDDCoreBlacklist                   *[]string          `json:"sddCoreBlacklist,omitempty"`
	SDDB2BWhitelist                    []*SDDB2BWhitelist `json:"sddB2bWhitelist,omitempty"`
}

// Create creates a Treezor beneficiary.
func (s *BeneficiaryService) Create(ctx context.Context, beneficiary *BeneficiaryRequest) (*Beneficiary, *http.Response, error) {
	req, _ := s.client.NewRequest(http.MethodPost, "beneficiaries", beneficiary)

	b := new(BeneficiaryResponse)
	resp, err := s.client.Do(ctx, req, b)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(b.Beneficiaries) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one beneficiary: %d beneficiaries returned", len(b.Beneficiaries))
	}
	return b.Beneficiaries[0], resp, nil
}

// Get returns a beneficiary.
func (s *BeneficiaryService) Get(ctx context.Context, beneficiaryID string) (*Beneficiary, *http.Response, error) {
	u := fmt.Sprintf("beneficiaries/%s", beneficiaryID)
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	b := new(BeneficiaryResponse)
	resp, err := s.client.Do(ctx, req, b)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(b.Beneficiaries) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one beneficiary: %d beneficiaries returned", len(b.Beneficiaries))
	}
	return b.Beneficiaries[0], resp, nil
}

// BeneficiaryOptions specifies the optional parameters to the BeneficiaryService.List.
type BeneficiaryOptions struct {
	UserID string `url:"userId,omitempty"`
}

// List the beneficiaries for the authenticated user.s
func (s *BeneficiaryService) List(ctx context.Context, opt *BeneficiaryOptions) (*BeneficiaryResponse, *http.Response, error) {
	u := "beneficiaries"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	b := new(BeneficiaryResponse)
	resp, err := s.client.Do(ctx, req, b)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	return b, resp, errors.WithStack(err)
}

// Edit updates a beneficiary.
func (s *BeneficiaryService) Edit(ctx context.Context, beneficiaryID string, beneficiary *BeneficiaryRequest) (*Beneficiary, *http.Response, error) {
	u := fmt.Sprintf("beneficiaries/%s", beneficiaryID)
	req, _ := s.client.NewRequest(http.MethodPut, u, beneficiary)

	b := new(BeneficiaryResponse)
	resp, err := s.client.Do(ctx, req, b)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(b.Beneficiaries) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one beneficiary: %d beneficiaries returned", len(b.Beneficiaries))
	}
	return b.Beneficiaries[0], resp, nil
}

// TODO: Update Beneficiary API
