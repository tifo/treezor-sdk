package treezor

import (
	"context"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/tifo/treezor-sdk/internal/types"
)

type ConnectOperationService service

type OperationAmount struct {
	Amount   *int64    `json:"amount,omitempty"`
	Currency *Currency `json:"currency,omitempty"`
}

// NOTE: It's not clear in the documentation if snake_case or camlCase is used here (see https://docs.treezor.com/guide/operations/introduction.html)
const (
	// Standard Operation
	OperationTypeCheck           = "check"
	OperationTypeBankTransfer    = "bankTransfer"
	OperationTypeBankDirectDebit = "bankDirectDebit"
	OperationTypeCardTopup       = "cardTopup"
	OperationTypeCardTransaction = "cardTransaction"
	// Refund Operation
	OperationTypeCheckRefund           = "checkRefund"
	OperationTypeBankTransferRefund    = "bankTransferRefund"
	OperationTypeBankDirectDebitRefund = "bankDirectDebitRefund"
	OperationTypeCardTopupRefund       = "cardTopupRefund"
	OperationTypeCardTransactionRefund = "cardTransactionRefund"
)

type OperationDirection string

const (
	OperationDirectionDebit  OperationDirection = "DEBIT"
	OperationDirectionCredit OperationDirection = "CREDIT"
)

type OperationStatus string

const (
	OperationStatusPending   OperationStatus = "PENDING"
	OperationStatusCanceled  OperationStatus = "CANCELED"
	OperationStatusConfirmed OperationStatus = "CONFIRMED"
	OperationStatusSettled   OperationStatus = "SETTLED"
)

type OperationSettlement string

const (
	OperationSettlementConfirmed OperationSettlement = "CONFIRMED"
	OperationSettlementSettled   OperationSettlement = "SETTLED"
)

type OperationMetadata struct {
	CardPayment *CardPaymentMetadata `json:"cardPayment,omitempty"`
}

type CardPaymentMetadata struct {
	MCC               *MCCMetadata     `json:"mcc,omitempty"`
	MID               *MIDMetadata     `json:"mid,omitempty"`
	LocalAmount       *OperationAmount `json:"localAmount,omitempty"`
	AuthorizationNote *string          `json:"authorizationNote,omitempty"`
}

type MCCMetadata struct {
	Code *types.Identifier `json:"code,omitempty"`
}

type MIDMetadata struct {
	Value *string `json:"value,omitempty"`
}

type OperationDate struct {
	Creation   time.Time `layout:"RFC3339" json:"creation,omitempty"`
	Settlement time.Time `layout:"RFC3339" json:"settlement,omitempty"`
}

type Operation struct {
	OperationType     *string              `json:"operationType,omitempty"`
	Amount            *OperationAmount     `json:"amount,omitempty"`
	WalletID          *types.Identifier    `json:"walletId,omitempty"`
	Settlement        *OperationSettlement `json:"settlement,omitempty"`
	Direction         *OperationDirection  `json:"direction,omitempty"`
	ObjectID          *types.Identifier    `json:"objectId,omitempty"`
	Label             *string              `json:"label,omitempty"`
	ExternalReference *string              `json:"externalReference,omitempty"`
	Metadata          *OperationMetadata   `json:"metadata,omitempty"`
	Status            *OperationStatus     `json:"status,omitempty"`
	Date              *OperationDate       `json:"date,omitempty"`
}

type OperationList struct {
	Data   []*Operation
	Cursor *ConnectCursor
}

type ConnectOperationListOptions struct {
	WalletID *string   `url:"walletId,omitempty" json:"-"`                  // Required
	DateTo   time.Time `layout:"RFC3339" url:"dateTo,omitempty" json:"-"`   // Required
	DateFrom time.Time `layout:"RFC3339" url:"dateFrom,omitempty" json:"-"` // Required
	Cursor   *string   `url:"cursor,omitempty" json:"-"`                    // Optional
}

func (s *ConnectOperationService) List(ctx context.Context, opts *ConnectOperationListOptions) (*OperationList, *http.Response, error) {
	u := "core-connect/operations"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	d := new(OperationList)
	resp, err := s.client.Do(ctx, req, d)

	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	return d, resp, nil
}

type ConnectOperationListAllOptions struct {
	WalletID *string   `url:"walletId,omitempty" json:"-"`                                    // Required
	DateTo   time.Time `layout:"2006-01-02T15:04:05Z07:00" url:"dateTo,omitempty" json:"-"`   // Required
	DateFrom time.Time `layout:"2006-01-02T15:04:05Z07:00" url:"dateFrom,omitempty" json:"-"` // Required
}

func (s *ConnectOperationService) ListAll(ctx context.Context, opts *ConnectOperationListAllOptions) ([]*Operation, *http.Response, error) {

	var ops []*Operation
	var cursor *string
	first := true
	for first || cursor != nil {
		oplist, resp, err := s.List(ctx, &ConnectOperationListOptions{
			WalletID: opts.WalletID,
			DateTo:   opts.DateTo,
			DateFrom: opts.DateFrom,
			Cursor:   cursor,
		})
		if err != nil {
			return ops, resp, err
		}
		ops = append(ops, oplist.GetData()...)
		first = false
		cursor = oplist.Cursor.Next
	}
	return ops, nil, nil
}
