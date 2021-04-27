package treezor

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
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

type AdditionalDataOneOf struct {
	AdditionalData struct {
		Card struct {
			ExternalProvider struct {
				TransactionReference string `json:"transactionReference"`
			} `json:"externalProvider"`
		} `json:"card"`
	} `json:"additionalData"`
	Message string
}

// Payin represents a pay-in to a beneficiary.
type Payin struct {
	Access
	PayinID              *string         `json:"payinId,omitempty"`
	PayinTag             *string         `json:"payinTag,omitempty"`
	PayinStatus          *string         `json:"payinStatus,omitempty"`
	CodeStatus           *string         `json:"codeStatus,omitempty"`
	InformationStatus    *string         `json:"informationStatus,omitempty"`
	WalletID             *string         `json:"walletId,omitempty"`
	UserID               *string         `json:"userId,omitempty"`
	WalletEventName      *string         `json:"walletEventName,omitempty"`
	WalletAlias          *string         `json:"walletAlias,omitempty"`
	UserFirstname        *string         `json:"userFirstname,omitempty"`
	UserLastname         *string         `json:"userLastname,omitempty"`
	MessageToUser        *string         `json:"messageToUser,omitempty"`
	PaymentMethodID      *string         `json:"paymentMethodId,omitempty"`
	SubtotalItems        *float64        `json:"subtotalItems,string,omitempty"`
	SubtotalServices     *float64        `json:"subtotalServices,string,omitempty"`
	SubtotalTax          *float64        `json:"subtotalTax,string,omitempty"`
	Amount               *float64        `json:"amount,string,omitempty"`
	Currency             Currency        `json:"currency,omitempty"`
	DistributorFee       *float64        `json:"distributorFee,string,omitempty"`
	CreatedDate          *TimestampParis `json:"createdDate,omitempty"`
	CreatedIP            *string         `json:"createdIp,omitempty"`
	PaymentHTML          *string         `json:"paymentHtml,omitempty"`
	PaymentLanguage      *string         `json:"paymentLanguage,omitempty"`
	PaymentPostURL       *string         `json:"paymentPostUrl,omitempty"`
	PaymentPostDataURL   *string         `json:"paymentPostDataUrl,omitempty"`
	PaymentAcceptedURL   *string         `json:"paymentAcceptedUrl,omitempty"`
	PaymentWaitingURL    *string         `json:"paymentWaitingUrl,omitempty"`
	PaymentRefusedURL    *string         `json:"paymentRefusedUrl,omitempty"`
	PaymentCanceledURL   *string         `json:"paymentCanceledUrl,omitempty"`
	PaymentExceptionURL  *string         `json:"paymentExceptionUrl,omitempty"`
	DebitorIBAN          *string         `json:"DbtrIBAN,omitempty"`
	IBANFullname         *string         `json:"ibanFullname,omitempty"`
	IBANID               *string         `json:"ibanId,omitempty"`
	IBANBIC              *string         `json:"ibanBic,omitempty"`
	IBANTxEndToEndID     *string         `json:"ibanTxEndToEndId,omitempty"`
	IBANTxID             *string         `json:"ibanTxId,omitempty"`
	RefundAmount         *float64        `json:"refundAmount,string,omitempty"`
	TotalRows            *int64          `json:"totalRows,string,omitempty"`
	ForwardURL           *string         `json:"forwardUrl,omitempty"`
	PayinDate            *Date           `json:"payinDate,omitempty"`
	MandateID            *string         `json:"mandateId,omitempty"`
	CreditorName         *string         `json:"creditorName,omitempty"`
	CreditorAddressLine  *string         `json:"creditorAddressLine,omitempty"`
	CreditorCountry      *string         `json:"creditorCountry,omitempty"`
	CreditorIBAN         *string         `json:"creditorIban,omitempty"`
	CreditorBIC          *string         `json:"creditorBIC,omitempty"`
	VirtualIBANID        *string         `json:"virtualIbanId,omitempty"`
	VirtualIBANReference *string         `json:"virtualIbanReference,omitempty"`
	AdditionalData       *AdditionalDataOneOf
}

func (t *AdditionalDataOneOf) UnmarshalJSON(data []byte) error {
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
