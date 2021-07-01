package treezor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/tifo/treezor-sdk/internal/types"
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
	CardTransactionID         *types.Identifier      `json:"cardtransactionId,omitempty"`
	CardID                    *types.Identifier      `json:"cardId,omitempty"`
	WalletID                  *types.Identifier      `json:"walletId,omitempty"`
	WalletCurrency            *string                `json:"walletCurrency,omitempty"` // NOTE: Numeric identifier for the Currency note the ISO Code
	MerchantID                *string                `json:"merchantId,omitempty"`     // NOTE: MID not an identifier
	MerchantName              *string                `json:"merchantName,omitempty"`
	MerchantCity              *string                `json:"merchantCity,omitempty"`
	MerchantCountry           *string                `json:"merchantCountry,omitempty"`
	MccCode                   *string                `json:"mccCode,omitempty"`
	PaymentLocalTime          *types.TimestampLondon `json:"paymentLocalTime,omitempty"`
	PublicToken               *string                `json:"publicToken,omitempty"`
	PaymentAmount             *types.Amount          `json:"paymentAmount,omitempty"`
	PaymentCurrency           *string                `json:"paymentCurrency,omitempty"` // NOTE: Numeric identifier for the Currency note the ISO Code
	Fees                      *types.Amount          `json:"fees,omitempty"`
	PaymentCountry            *string                `json:"paymentCountry,omitempty"`
	PaymentID                 *types.Identifier      `json:"paymentId,omitempty"`
	PaymentStatus             *string                `json:"paymentStatus,omitempty"` // NOTE: can be an enum
	PaymentLocalAmount        *types.Amount          `json:"paymentLocalAmount,omitempty"`
	PaymentLocalDate          *types.Date            `json:"paymentLocalDate,omitempty"`
	Is3DS                     *types.Boolean         `json:"is3DS,omitempty"`
	PosCardholderPresence     *types.Boolean         `json:"posCardholderPresence,omitempty"`
	PosPostcode               *string                `json:"posPostcode,omitempty"`
	PosCountry                *string                `json:"posCountry,omitempty"`
	PosTerminalID             *string                `json:"posTerminalId,omitempty"`
	PosCardPresence           *types.Boolean         `json:"posCardPresence,omitempty"`
	PanEntryMethod            *string                `json:"panEntryMethod,omitempty"` // NOTE: can be an enum
	AuthorizationNote         *string                `json:"authorizationNote,omitempty"`
	AuthorizationResponseCode *string                `json:"authorizationResponseCode,omitempty"`
	AuthorizationIssuerID     *string                `json:"authorizationIssuerId,omitempty"`
	AuthorizationIssuerTime   *types.TimestampLondon `json:"authorizationIssuerTime,omitempty"`
	AuthorizationMTI          *string                `json:"authorizationMti,omitempty"` // NOTE: see ISO8583
	AuthorizedBalance         *types.Amount          `json:"authorizedBalance,omitempty"`
	LimitATMYear              *types.Integer         `json:"limitAtmYear,omitempty"`
	LimitATMMonth             *types.Integer         `json:"limitAtmMonth,omitempty"`
	LimitATMWeek              *types.Integer         `json:"limitAtmWeek,omitempty"`
	LimitATMDay               *types.Integer         `json:"limitAtmDay,omitempty"`
	LimitATMAll               *types.Integer         `json:"limitAtmAll,omitempty"`
	LimitPaymentYear          *types.Integer         `json:"limitPaymentYear,omitempty"`
	LimitPaymentMonth         *types.Integer         `json:"limitPaymentMonth,omitempty"`
	LimitPaymentWeek          *types.Integer         `json:"limitPaymentWeek,omitempty"`
	LimitPaymentDay           *types.Integer         `json:"limitPaymentDay,omitempty"`
	LimitPaymentAll           *types.Integer         `json:"limitPaymentAll,omitempty"`
	TotalLimitATMYear         *types.Amount          `json:"totalLimitAtmYear,omitempty"`
	TotalLimitATMMonth        *types.Amount          `json:"totalLimitAtmMonth,omitempty"`
	TotalLimitATMWeek         *types.Amount          `json:"totalLimitAtmWeek,omitempty"`
	TotalLimitATMDay          *types.Amount          `json:"totalLimitAtmDay,omitempty"`
	TotalLimitATMAll          *types.Amount          `json:"totalLimitAtmAll,omitempty"`
	TotalLimitPaymentYear     *types.Amount          `json:"totalLimitPaymentYear,omitempty"`
	TotalLimitPaymentMonth    *types.Amount          `json:"totalLimitPaymentMonth,omitempty"`
	TotalLimitPaymentWeek     *types.Amount          `json:"totalLimitPaymentWeek,omitempty"`
	TotalLimitPaymentDay      *types.Amount          `json:"totalLimitPaymentDay,omitempty"`
	TotalLimitPaymentAll      *types.Amount          `json:"totalLimitPaymentAll,omitempty"`
	TotalRows                 *types.Integer         `json:"totalRows,omitempty"`
	// NOTE: see about totalRows, codeStatus and informationStatus
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
	CardID    string `url:"paymentId,omitempty"`
	PaymentID string `url:"paymentId,omitempty"`
	UserID    string `url:"userId,omitempty"`
	WalletID  string `url:"walletId,omitempty"`

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

// TODO: Update CardTransaction API
