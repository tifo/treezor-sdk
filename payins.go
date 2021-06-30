package treezor

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/tifo/treezor-sdk/internal/types"
)

// PayinService handles communication with the payin related
// methods of the Treezor API.
//
// Treezor API docs: https://www.treezor.com/api-documentation/#/payin
type PayinService service

// PayinResponse represents a list of payins.
// It may contain only one item.
type PayinResponse struct {
	Payins []*Payin `json:"payins"`
}

type Payin struct {
	PayinID              *types.Identifier     `json:"payinId,omitempty"`
	PayinTag             *string               `json:"payinTag,omitempty"`
	PayinStatus          *string               `json:"payinStatus,omitempty"` // NOTE: can be an enum
	WalletID             *types.Identifier     `json:"walletId,omitempty"`
	UserID               *types.Identifier     `json:"userId,omitempty"`
	CartID               *types.Identifier     `json:"cartId,omitempty"`
	WalletEventName      *string               `json:"walletEventName,omitempty"`
	WalletAlias          *string               `json:"walletAlias,omitempty"`
	UserFirstname        *string               `json:"userFirstname,omitempty"`
	UserLastname         *string               `json:"userLastname,omitempty"`
	MessageToUser        *string               `json:"messageToUser,omitempty"`
	PaymentMethodID      *types.Identifier     `json:"paymentMethodId,omitempty"` // NOTE: can be an enum
	SubtotalItems        *types.Amount         `json:"subtotalItems,omitempty"`
	SubtotalServices     *types.Amount         `json:"subtotalServices,omitempty"`
	SubtotalTax          *types.Amount         `json:"subtotalTax,omitempty"`
	Amount               *types.Amount         `json:"amount,omitempty"`
	Currency             *Currency             `json:"currency,omitempty"`
	DistributorFee       *types.Amount         `json:"distributorFee,omitempty"`
	CreatedDate          *types.TimestampParis `json:"createdDate,omitempty"`
	CreatedIP            *string               `json:"createdIp,omitempty"`
	PaymentHTML          *string               `json:"paymentHtml,omitempty"`
	PaymentLanguage      *string               `json:"paymentLanguage,omitempty"`
	PaymentPostURL       *string               `json:"paymentPostUrl,omitempty"`
	PaymentPostDataURL   *string               `json:"paymentPostDataUrl,omitempty"`
	PaymentAcceptedURL   *string               `json:"paymentAcceptedUrl,omitempty"`
	PaymentWaitingURL    *string               `json:"paymentWaitingUrl,omitempty"`
	PaymentRefusedURL    *string               `json:"paymentRefusedUrl,omitempty"`
	PaymentCanceledURL   *string               `json:"paymentCanceledUrl,omitempty"`
	PaymentExceptionURL  *string               `json:"paymentExceptionUrl,omitempty"`
	IBANFullname         *string               `json:"ibanFullname,omitempty"`
	IBANID               *string               `json:"ibanId,omitempty"`
	IBANBIC              *string               `json:"ibanBic,omitempty"`
	IBANTxEndToEndID     *string               `json:"ibanTxEndToEndId,omitempty"`
	IBANTxID             *string               `json:"ibanTxId,omitempty"`
	RefundAmount         *types.Amount         `json:"refundAmount,omitempty"`
	DbtrIBAN             *string               `json:"DbtrIBAN,omitempty"`
	ForwardURL           *string               `json:"forwardUrl,omitempty"`
	PayinDate            *types.Date           `json:"payinDate,omitempty"`
	MandateID            *types.Identifier     `json:"mandateId,omitempty"`
	CreditorName         *string               `json:"creditorName,omitempty"`
	CreditorAddressLine  *string               `json:"creditorAddressLine,omitempty"`
	CreditorCountry      *string               `json:"creditorCountry,omitempty"`
	CreditorIBAN         *string               `json:"creditorIban,omitempty"`
	CreditorBIC          *string               `json:"creditorBIC,omitempty"`
	VirtualIBANID        *types.Identifier     `json:"virtualIbanId,omitempty"`
	VirtualIBANReference *string               `json:"virtualIbanReference,omitempty"`
	AdditionalData       *PayinAdditionalData
	TotalRows            *types.Integer    `json:"totalRows,omitempty"`
	CodeStatus           *types.Identifier `json:"codeStatus,omitempty"`        // Legacy field
	InformationStatus    *string           `json:"informationStatus,omitempty"` // Legacy field
}

type PayinAdditionalData struct {
	AdditionalData struct {
		Card *struct {
			ExternalProvider struct {
				TransactionReference string `json:"transactionReference"`
			} `json:"externalProvider"`
		} `json:"card,omitempty"`
		Cheque *struct {
			CMC7 struct {
				A string `json:"a"`
				B string `json:"b"`
				C string `json:"c"`
			} `json:"cmc7"`
			RLMCKey    string `json:"RLMCKey"`
			DrawerData struct {
				Email           *string `json:"email,omitempty"`
				FirstName       *string `json:"firstName,omitempty"`
				LastName        *string `json:"lastName,omitempty"`
				Address         *string `json:"address,omitempty"`
				Address2        *string `json:"address2,omitempty"`
				ZipCode         *string `json:"zipCode,omitempty"`
				City            *string `json:"city,omitempty"`
				IsNaturalPerson *bool   `json:"isNaturalPerson,omitempty"`
			}
		} `json:"cheque,omitempty"`
	} `json:"additionalData"`
	Message string
}

// NOTE: See about setting this as a propper "one-of" and have each oneof type nullable ?

func (t *PayinAdditionalData) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &t.AdditionalData); err != nil {
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			if merr := json.Unmarshal(data, &t.Message); merr != nil {
				return merr
			} else {
				if derr := json.Unmarshal([]byte(t.Message), &t.AdditionalData); derr != nil {
					if _, ok := err.(*json.UnmarshalTypeError); ok {
						return nil
					}
					return derr
				}
			}
		} else {
			return err
		}
	}
	return nil
}

// Create creates a Treezor pay-in.
// The required field are WalletID, BeneficiaryID, Amount, Currency(ISO 4217).
func (s *PayinService) Create(ctx context.Context, payin *Payin) (*Payin, *http.Response, error) {
	req, _ := s.client.NewRequest(http.MethodPost, "payins", payin)

	b := new(PayinResponse)
	resp, err := s.client.Do(ctx, req, b)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(b.Payins) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one pay-in: %d pay-ins returned", len(b.Payins))
	}
	return b.Payins[0], resp, nil
}

// Get returns a pay-in.
func (s *PayinService) Get(ctx context.Context, payinID string) (*Payin, *http.Response, error) {
	u := fmt.Sprintf("payins/%s", payinID)
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	b := new(PayinResponse)
	resp, err := s.client.Do(ctx, req, b)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(b.Payins) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one pay-in: %d pay-ins returned", len(b.Payins))
	}
	return b.Payins[0], resp, nil
}

// PayinListOptions specifies the optional parameters to the PayinService.List.
type PayinListOptions struct {
	PayinStatus     string `url:"payinStatus,omitempty"`
	UserID          string `url:"userId,omitempty"`
	WalletID        string `url:"walletId,omitempty"`
	CreatedDateFrom string `url:"createdDateFrom,omitempty"`
	CreatedDateTo   string `url:"createdDateTo,omitempty"`

	ListOptions
}

// List the pay-ins for the authenticated user.
func (s *PayinService) List(ctx context.Context, opt *PayinListOptions) (*PayinResponse, *http.Response, error) {
	u := "payins"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	b := new(PayinResponse)
	resp, err := s.client.Do(ctx, req, b)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	return b, resp, errors.WithStack(err)
}

// Delete deletes a payin. Change payin's status to CANCELED. A validated payin can't be cancelled.
func (s *PayinService) Delete(ctx context.Context, payinID string) (*Payin, *http.Response, error) {
	u := fmt.Sprintf("payins/%s", payinID)
	req, _ := s.client.NewRequest(http.MethodDelete, u, nil)

	b := new(PayinResponse)
	resp, err := s.client.Do(ctx, req, b)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(b.Payins) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one pay-in: %d pay-ins returned", len(b.Payins))
	}
	return b.Payins[0], resp, nil
}
