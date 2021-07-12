package treezor

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/tifo/treezor-sdk/internal/types"
)

// DocumentStatus defines the state of a document
type DocumentStatus string

const (
	DocumentStatusPending   DocumentStatus = "PENDING"
	DocumentStatusCanceled  DocumentStatus = "CANCELED"
	DocumentStatusValidated DocumentStatus = "VALIDATED"
)

// DocumentType represents the type of legal document.
type DocumentType int32

func (t *DocumentType) UnmarshalJSON(data []byte) error {
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	v, err := str.Int64()
	if err != nil {
		return err
	}
	*t = DocumentType(v)
	return nil
}

// All the types of Document Treezor accepts.
const (
	PoliceRecord                DocumentType = 2
	CompanyRegistration         DocumentType = 4
	CV                          DocumentType = 6
	SwornStatement              DocumentType = 7
	Turnover                    DocumentType = 8
	IdentityCard                DocumentType = 9
	BankIdentityStatement       DocumentType = 11
	ProofOfAddress              DocumentType = 12
	MobilePhoneInvoice          DocumentType = 13
	Invoice                     DocumentType = 14
	ResidencePermit             DocumentType = 15
	DrivingLicense              DocumentType = 16
	Passport                    DocumentType = 17
	EmployeeProxy               DocumentType = 18
	OfficialCompanyRegistration DocumentType = 19
	TaxCertificate              DocumentType = 20
	EmployeePaymentNotice       DocumentType = 21
	UserBankStatement           DocumentType = 22
	BusinessLegalStatus         DocumentType = 23
	TaxStatement                DocumentType = 24
	ExemptionStatement          DocumentType = 25
	LivenessResult              DocumentType = 26
	HealthInsuranceCard         DocumentType = 27
)

var documentTypeLookup = map[DocumentType]string{
	PoliceRecord:                "PoliceRecord",
	CompanyRegistration:         "CompanyRegistration",
	CV:                          "CV",
	SwornStatement:              "SwornStatement",
	Turnover:                    "Turnover",
	IdentityCard:                "IdentityCard",
	BankIdentityStatement:       "BankIdentityStatement",
	ProofOfAddress:              "ProofOfAddress",
	MobilePhoneInvoice:          "MobilePhoneInvoice",
	Invoice:                     "Invoice",
	ResidencePermit:             "ResidencePermit",
	DrivingLicense:              "DrivingLicense",
	Passport:                    "Passport",
	EmployeeProxy:               "EmployeeProxy",
	OfficialCompanyRegistration: "OfficialCompanyRegistration",
	TaxCertificate:              "TaxCertificate",
	EmployeePaymentNotice:       "EmployeePaymentNotice",
	UserBankStatement:           "UserBankStatement",
	BusinessLegalStatus:         "BusinessLegalStatus",
	TaxStatement:                "TaxStatement",
	ExemptionStatement:          "ExemptionStatement",
	LivenessResult:              "LivenessResult",
	HealthInsuranceCard:         "HealthInsuranceCard",
}

func (d DocumentType) String() string {
	return documentTypeLookup[d]
}

// DocumentService handles communication with the document related
// methods of the Treezor API.
//
// Treezor API docs: https://www.treezor.com/api-documentation/#/document
type DocumentService service

// DocumentResponse represents a list of KYC documents.
// It may contain only one item.
type DocumentResponse struct {
	Documents []*Document `json:"documents"`
}

// Document represents a KYC document.
type Document struct {
	DocumentID        *types.Identifier     `json:"documentId,omitempty"`
	DocumentTag       *string               `json:"documentTag,omitempty"`
	DocumentStatus    *DocumentStatus       `json:"documentStatus,omitempty"`
	DocumentTypeID    *DocumentType         `json:"documentTypeId,omitempty"`
	DocumentType      *string               `json:"documentType,omitempty"`
	ResidenceID       *types.Identifier     `json:"residenceId,omitempty"`
	ClientID          *types.Identifier     `json:"clientId,omitempty"`
	UserID            *types.Identifier     `json:"userId,omitempty"`
	UserLastname      *string               `json:"userLastname,omitempty"`
	UserFirstname     *string               `json:"userFirstname,omitempty"`
	Filename          *string               `json:"fileName,omitempty"`
	TemporaryURL      *string               `json:"temporaryUrl,omitempty"`
	TemporaryURLThumb *string               `json:"temporaryUrlThumb,omitempty"`
	CreatedDate       *types.TimestampParis `json:"createdDate,omitempty"`
	ModifiedDate      *types.TimestampParis `json:"modifiedDate,omitempty"`
	TotalRows         *types.Integer        `json:"totalRows,omitempty"`
	CodeStatus        *types.Identifier     `json:"codeStatus,omitempty"`        // Legacy field
	InformationStatus *string               `json:"informationStatus,omitempty"` // Legacy field
}

type DocumentCreateOptions struct {
	Access

	DocumentTag       *string      `url:"-" json:"documentTag,omitempty"`       // Optional
	UserID            *string      `url:"-" json:"userId,omitempty"`            // Required
	ResidenceID       *string      `url:"-" json:"residenceId,omitempty"`       // Required when DocumentTypeID is set to 24 (TaxStatement) or 25 (ExemptionStatement)
	DocumentTypeID    DocumentType `url:"-" json:"documentTypeId,omitempty"`    // Required
	Name              *string      `url:"-" json:"name,omitempty"`              // Required
	FileContentBase64 *string      `url:"-" json:"fileContentBase64,omitempty"` // Required
}

// Create uploads the given file to Treezor for later KYC review.
func (s *DocumentService) Create(ctx context.Context, opts *DocumentCreateOptions) (*Document, *http.Response, error) {
	u := "documents"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPost, u, opts)

	d := new(DocumentResponse)
	resp, err := s.client.Do(ctx, req, d)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(d.Documents) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one document: %d documents returned", len(d.Documents))
	}
	return d.Documents[0], resp, nil
}

type DocumentEditOptions struct {
	Access
}

// Edit updates a give document.
// NOTE: this method seems to do nothing (see https://www.treezor.com/api-documentation/#/document/putDocument)
func (s *DocumentService) Edit(ctx context.Context, documentID string, opts *DocumentEditOptions) (*Document, *http.Response, error) {
	u := fmt.Sprintf("documents/%s", documentID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPut, u, opts)

	d := new(DocumentResponse)
	resp, err := s.client.Do(ctx, req, d)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(d.Documents) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one document: %d documents returned", len(d.Documents))
	}
	return d.Documents[0], resp, nil
}

// DocumentListOptions contains options for listing documents.
type DocumentListOptions struct {
	Access

	DocumentID     *string        `url:"documentId,omitempty" json:"-"`     // Optional
	DocumentTag    *string        `url:"documentTag,omitempty" json:"-"`    // Optional
	DocumentStatus DocumentStatus `url:"documentStatus,omitempty" json:"-"` // Optional
	DocumentTypeID DocumentType   `url:"documentTypeId,omitempty" json:"-"` // Optional
	DocumentType   *string        `url:"documentType,omitempty" json:"-"`   // Optional
	UserID         *string        `url:"userId,omitempty" json:"-"`         // Optional
	UserName       *string        `url:"userName,omitempty" json:"-"`       // Optional
	UserEmail      *string        `url:"userEmail,omitempty" json:"-"`      // Optional
	FileName       *string        `url:"fileName,omitempty" json:"-"`       // Optional
	FileSize       *string        `url:"fileSize,omitempty" json:"-"`       // Optional
	FileType       *string        `url:"fileType,omitempty" json:"-"`       // Optional
	IsAgent        *string        `url:"isAgent,omitempty" json:"-"`        // Optional

	ListOptions
}

// List returns a list of documents.
func (s *DocumentService) List(ctx context.Context, opts *DocumentListOptions) (*DocumentResponse, *http.Response, error) {
	u := "documents"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	dr := new(DocumentResponse)
	resp, err := s.client.Do(ctx, req, dr)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	return dr, resp, nil
}

type DocumentGetOptions struct {
	Access
}

// Get fetch document info from Treezor
func (s *DocumentService) Get(ctx context.Context, documentID string, opts *DocumentGetOptions) (*Document, *http.Response, error) {
	u := fmt.Sprintf("documents/%s", documentID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	docs := new(DocumentResponse)
	resp, err := s.client.Do(ctx, req, docs)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(docs.Documents) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one document: %d document returned", len(docs.Documents))
	}
	return docs.Documents[0], resp, nil
}

type DocumentDeleteOptions struct {
	Access
}

// Delete deletes a document in treezor
func (s *DocumentService) Delete(ctx context.Context, documentID string, opts *DocumentDeleteOptions) (*http.Response, error) {
	u := fmt.Sprintf("documents/%s", documentID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodDelete, u, nil)

	docs := new(DocumentResponse)
	resp, err := s.client.Do(ctx, req, docs)
	if err != nil {
		return resp, errors.WithStack(err)
	}

	if len(docs.Documents) != 1 {
		return resp, errors.Errorf("API did not returned exactly one document: %d document returned", len(docs.Documents))
	}

	return resp, nil
}
