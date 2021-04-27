package treezor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// DocumentType represents the type of legal document.
type DocumentType int32

// Treezor document status
const (
	DocumentStatusPending   = "PENDING"
	DocumentStatusCanceled  = "CANCELED"
	DocumentStatusValidated = "VALIDATED"
)

// All the types of Document Treezor accepts.
const (
	PoliceRecord                DocumentType = 2
	CompanyRegistration                      = 4
	CV                                       = 6
	SwornStatement                           = 7
	Turnover                                 = 8
	IdentityCard                             = 9
	BankIdentityStatement                    = 11
	ProofOfAddress                           = 12
	MobilePhoneInvoice                       = 13
	Invoice                                  = 14
	ResidencePermit                          = 15
	DrivingLicense                           = 16
	Passport                                 = 17
	EmployeeProxy                            = 18
	OfficialCompanyRegistration              = 19
	TaxCertificate                           = 20
	EmployeePaymentNotice                    = 21
	UserBankStatement                        = 22
	BusinessLegalStatus                      = 23
	TaxStatement                             = 24
	ExemptionStatement                       = 25
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
	Access
	DocumentID        *string      `json:"documentId,omitempty"`
	DocumentStatus    *string      `json:"documentStatus,omitempty"`
	DocumentTypeID    DocumentType `json:"documentTypeId,string,omitempty"`
	DocumentType      *string      `json:"documentType,omitempty"`
	ClientID          *string      `json:"clientId,omitempty"`
	UserID            *string      `json:"userId,omitempty"`
	UserLastname      *string      `json:"userLastname,omitempty"`
	UserFirstname     *string      `json:"userFirstname,omitempty"`
	Filename          *string      `json:"fileName,omitempty"`
	FileSize          *int64       `json:"fileSize,omitempty"`
	Name              *string      `json:"name,omitempty"`
	FileContentBase64 string       `json:"fileContentBase64,omitempty"`
	TemporaryURL      *string      `json:"temporaryUrl,omitempty"`
	TemporaryURLThumb *string      `json:"temporaryUrlThumb,omitempty"`
	CreatedDate       *string      `json:"createdDate,omitempty"`
	ModifiedDate      *string      `json:"modifiedDate,omitempty"`
	TotalRows         *int64       `json:"totalRows,omitempty"`
	ResidenceID       *string      `json:"residenceId,omitempty"`
	InformationStatus *string      `json:"informationStatus,omitempty"`
	CodeStatus        *string      `json:"codeStatus,omitempty"`
}

// Send uploads the given file to Treezor for later KYC review.
func (s *DocumentService) Send(ctx context.Context, document *Document) (*Document, *http.Response, error) {
	req, _ := s.client.NewRequest(http.MethodPost, "documents", document)

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

// Get fetch document info from Treezor
func (s *DocumentService) Get(ctx context.Context, documentID string) (*Document, *http.Response, error) {
	route := fmt.Sprintf("documents/%s", documentID)
	req, _ := s.client.NewRequest(http.MethodGet, route, nil)

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

// Delete deletes a document in treezor
func (s *DocumentService) Delete(ctx context.Context, documentID string) (*http.Response, error) {
	route := fmt.Sprintf("documents/%s", documentID)
	req, _ := s.client.NewRequest(http.MethodDelete, route, nil)

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
