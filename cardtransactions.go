package treezor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// CardTransactionService handles communication with the card transaction related
// methods of the Treezor API.
//
// Treezor API docs: https://www.treezor.com/api-documentation/#/cardtransaction
type CardTransactionService service

// CardTransactionResponse represents a list of card transactions.
// It may contain only one item.
type CardTransactionResponse struct {
	CardTransactions []*CardTransaction `json:"cardtransactions"`
}

// CardTransaction represents a card transaction made at a PoS.
type CardTransaction struct {
	CardTransactionID         *string          `json:"cardtransactionId,omitempty"`
	CardID                    *string          `json:"cardId,omitempty"`
	WalletID                  *string          `json:"walletId,omitempty"`
	WalletCurrency            *string          `json:"walletCurrency,omitempty"`
	MerchantID                *string          `json:"merchantId,omitempty"`
	MerchantName              *string          `json:"merchantName,omitempty"`
	MerchantCity              *string          `json:"merchantCity,omitempty"`
	MerchantCountry           *string          `json:"merchantCountry,omitempty"`
	PaymentLocalTime          *string          `json:"paymentLocalTime,omitempty"`
	PublicToken               *string          `json:"publicToken,omitempty"`
	PaymentAmount             *float64         `json:"paymentAmount,string,omitempty"`
	PaymentCurrency           *string          `json:"paymentCurrency,omitempty"`
	Fees                      *float64         `json:"fees,string,omitempty"`
	Is3DS                     *string          `json:"is3DS,omitempty"`
	PaymentCountry            *string          `json:"paymentCountry,omitempty"`
	PaymentID                 *string          `json:"paymentId,omitempty"`
	PaymentStatus             *string          `json:"paymentStatus,omitempty"`
	PaymentLocalAmount        *float64         `json:"paymentLocalAmount,string,omitempty"`
	PosCardholderPresence     *string          `json:"posCardholderPresence,omitempty"`
	PosPostcode               *string          `json:"posPostcode,omitempty"`
	PosCountry                *string          `json:"posCountry,omitempty"`
	PosTerminalID             *string          `json:"posTerminalId,omitempty"`
	PosCardPresence           *string          `json:"posCardPresence,omitempty"`
	PanEntryMethod            *string          `json:"panEntryMethod,omitempty"`
	AuthorizationNote         *string          `json:"authorizationNote,omitempty"`
	AuthorizationResponseCode *string          `json:"authorizationResponseCode,omitempty"`
	AuthorizationIssuerID     *string          `json:"authorizationIssuerId,omitempty"`
	AuthorizationIssuerTime   *TimestampLondon `json:"authorizationIssuerTime,omitempty"`
	AuthorizationMti          *string          `json:"authorizationMti,omitempty"`
	AuthorizedBalance         *float64         `json:"authorizedBalance,string,omitempty"`
	LimitATMYear              *int64           `json:"limitAtmYear,string,omitempty"`
	LimitATMMonth             *int64           `json:"limitAtmMonth,string,omitempty"`
	LimitATMWeek              *int64           `json:"limitAtmWeek,string,omitempty"`
	LimitATMDay               *int64           `json:"limitAtmDay,string,omitempty"`
	LimitATMAll               *int64           `json:"limitAtmAll,string,omitempty"`
	LimitPaymentYear          *int64           `json:"limitPaymentYear,string,omitempty"`
	LimitPaymentMonth         *int64           `json:"limitPaymentMonth,string,omitempty"`
	LimitPaymentWeek          *int64           `json:"limitPaymentWeek,string,omitempty"`
	LimitPaymentDay           *int64           `json:"limitPaymentDay,string,omitempty"`
	LimitPaymentAll           *int64           `json:"limitPaymentAll,string,omitempty"`
	TotalLimitATMYear         *float64         `json:"totalLimitAtmYear,string,omitempty"`
	TotalLimitATMMonth        *float64         `json:"totalLimitAtmMonth,string,omitempty"`
	TotalLimitATMWeek         *float64         `json:"totalLimitAtmWeek,string,omitempty"`
	TotalLimitATMDay          *float64         `json:"totalLimitAtmDay,string,omitempty"`
	TotalLimitATMAll          *float64         `json:"totalLimitAtmAll,string,omitempty"`
	TotalLimitPaymentYear     *float64         `json:"totalLimitPaymentYear,string,omitempty"`
	TotalLimitPaymentMonth    *float64         `json:"totalLimitPaymentMonth,string,omitempty"`
	TotalLimitPaymentWeek     *float64         `json:"totalLimitPaymentWeek,string,omitempty"`
	TotalLimitPaymentDay      *float64         `json:"totalLimitPaymentDay,string,omitempty"`
	TotalLimitPaymentAll      *float64         `json:"totalLimitPaymentAll,string,omitempty"`
	MccCode                   *string          `json:"mccCode,omitempty"`
}

// Get fetches a CardTransaction from Treezor.
func (s *CardTransactionService) Get(ctx context.Context, cardTransactionID string) (*CardTransaction, *http.Response, error) {
	u := fmt.Sprintf("cardtransactions/%s", cardTransactionID)
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	ct := new(CardTransactionResponse)
	resp, err := s.client.Do(ctx, req, ct)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(ct.CardTransactions) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one cardtransaction: %d cardtransactions returned", len(ct.CardTransactions))
	}
	return ct.CardTransactions[0], resp, nil
}

// CardTransactionsListOptions defines options to be passed as
// query parameters when performing a List operation
type CardTransactionsListOptions struct {
	PaymentID       string `url:"paymentId,omitempty"`
	UserID          string `url:"userId,omitempty"`
	WalletID        string `url:"walletId,omitempty"`
	CreatedDateFrom string `url:"createdDateFrom,omitempty"`
	CreatedDateTo   string `url:"createdDateTo,omitempty"`

	ListOptions
}

// List the pay-ins for the authenticated user.
func (s *CardTransactionService) List(ctx context.Context, opt *CardTransactionsListOptions) (*CardTransactionResponse, *http.Response, error) {
	u := "cardtransactions"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	b := new(CardTransactionResponse)
	resp, err := s.client.Do(ctx, req, b)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	return b, resp, errors.WithStack(err)
}
