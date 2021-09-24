package treezor

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"

	json "github.com/tifo/treezor-sdk/internal/json"
	"github.com/tifo/treezor-sdk/internal/types"
)

type ConnectKYCService service

// Types

// DocumentStatus defines the state of a document
type KYCDocumentStatus int32

func (t *KYCDocumentStatus) UnmarshalJSON(data []byte) error {
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	v, err := str.Int64()
	if err != nil {
		return err
	}
	*t = KYCDocumentStatus(v)
	return nil
}

const (
	KYCDocumentStatusPending   KYCDocumentStatus = 0
	KYCDocumentStatusRefused   KYCDocumentStatus = 1
	KYCDocumentStatusValidated KYCDocumentStatus = 2
)

type ConnectKYCDocument struct {
	DocumentID   *string           `json:"doumentId"`
	DocumentType *DocumentType     `json:"documentType"`
	Status       KYCDocumentStatus `json:"status"` // Cest un nombre pas un document status standard :issou:
	UserID       *types.Identifier `json:"userId"`
	CreatedAt    *time.Time        `layout:"RFC3339" json:"createdAt"`
	UpdatedAt    *time.Time        `layout:"RFC3339" json:"updatedAt"`
	Comment      *string           `json:"comment"`
	Metadata     types.Metadata    `json:"metadata"` // Pas typesafe parfois une map parfois un array vide
}

type UploadDocumentTargetForm struct {
	Action  *string `json:"action,omitempty"`
	Method  *string `json:"method,omitempty"`
	EncType *string `json:"enctype,omitempty"`
}

type UploadDocumentTarget struct {
	DocumentID *string                   `json:"documentId,omitempty"`
	Form       *UploadDocumentTargetForm `json:"form,omitempty"`
	FormFields map[string]string         `json:"formFields,omitempty"`
	ExpireIn   *int32                    `json:"expireIn,omitempty"`
}

type PreviewDocumentTarget struct {
	URL         *string `json:"url,omitempty"`
	ContentType *string `json:"contentType,omitempty"`
	Duration    *int    `json:"duration,omitempty"`
}

// POST /core-connect/users/{userID}/kyc/document

type ConnectKYCUploadDocumentOptions struct {
	DocumentType DocumentType      `url:"-" json:"documentType,omitempty"` // Required
	Metadata     map[string]string `url:"-" json:"metadata,omitempty"`     // Optional
}

func (s *ConnectKYCService) UploadDocument(ctx context.Context, userID string, opts *ConnectKYCUploadDocumentOptions) (*UploadDocumentTarget, *http.Response, error) {
	u := fmt.Sprintf("core-connect/users/%s/kyc/document", userID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPost, u, opts)

	d := new(UploadDocumentTarget)
	resp, err := s.client.Do(ctx, req, d)

	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	return d, resp, nil
}

// GET /core-connect/documents/{documentID}/preview

type ConnectKYCPreviewDocumentOptions struct{}

func (s *ConnectKYCService) PreviewDocument(ctx context.Context, documentID string, opts *ConnectKYCPreviewDocumentOptions) (*PreviewDocumentTarget, *http.Response, error) {
	u := fmt.Sprintf("core-connect/kyc/documents/%s/preview", documentID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	d := new(PreviewDocumentTarget)
	resp, err := s.client.Do(ctx, req, d)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	return d, resp, nil
}

// PUT /core-connect/kyc/documents/{documentID}/status

type ConnectKYCReviewDocumentOptions struct {
	Status  KYCDocumentStatus `url:"-" json:"status,omitempty"`  // Required
	Comment *string           `url:"-" json:"comment,omitempty"` // Optional
}

func (s *ConnectKYCService) ReviewDocument(ctx context.Context, documentID string, opts *ConnectKYCReviewDocumentOptions) (*ConnectKYCDocument, *http.Response, error) {
	u := fmt.Sprintf("core-connect/kyc/documents/%s/status", documentID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPut, u, opts)

	d := new(ConnectKYCDocument)
	resp, err := s.client.Do(ctx, req, d)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	return d, resp, nil
}

// GET /core-connect/users/{userID}/kyc/document

type ConnectKYCListDocumentsOptions struct{}

func (s *ConnectKYCService) ListDocuments(ctx context.Context, userID string, opts *ConnectKYCListDocumentsOptions) ([]*ConnectKYCDocument, *http.Response, error) {
	u := fmt.Sprintf("core-connect/users/%s/kyc/document", userID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	d := make([]*ConnectKYCDocument, 0)
	resp, err := s.client.Do(ctx, req, &d)

	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	return d, resp, nil
}

// GET /core-connect/users/&s/kyc-review

type ConnectKYCPreReviewUser struct{}

func (s *ConnectKYCService) PreReviewUser(ctx context.Context, userID string, opts *ConnectKYCPreReviewUser) (*http.Response, error) {
	u := fmt.Sprintf("core-connect/users/%s/kyc-review", userID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPut, u, nil)

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, errors.WithStack(err)
	}

	return resp, nil
}
