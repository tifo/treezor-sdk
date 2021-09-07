package treezor

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/tifo/treezor-sdk/internal/types"
)

type ConnectKYCService service

type ConnectKYCUploadDocumentOptions struct {
	DocumentType DocumentType      `url:"-" json:"documentType,omitempty"` // Required
	Metadata     map[string]string `url:"-" json:"metadata,omitempty"`     // Optional
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

type ConnectKYCPreviewDocumentOptions struct{}

type PreviewDocumentTarget struct {
	URL         *string `json:"url,omitempty"`
	ContentType *string `json:"contentType,omitempty"`
	Duration    *int    `json:"duration,omitempty"`
}

func (s *ConnectKYCService) PreviewDocument(ctx context.Context, documentID string, opts *ConnectKYCPreviewDocumentOptions) (*PreviewDocumentTarget, *http.Response, error) {
	u := fmt.Sprintf("core-connect/kyc/documents/%s/preview", documentID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPut, u, nil)

	d := new(PreviewDocumentTarget)
	resp, err := s.client.Do(ctx, req, d)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	return d, resp, nil
}

type DocumentReviewStatus int32

const (
	DocumentReviewRefused   DocumentReviewStatus = 1
	DocumentReviewValidated DocumentReviewStatus = 2
)

type ConnectKYCReviewDocumentOptions struct {
	Status  DocumentReviewStatus `url:"-" json:"status,omitempty"`  // Required
	Comment *string              `url:"-" json:"comment,omitempty"` // Optional
}

type DocumentReview struct {
	DocumentID   *types.Identifier     `json:"documentId,omitempty"`
	DocumentType *DocumentType         `json:"documentType,omitempty"`
	Status       *DocumentReviewStatus `json:"status,omitempty"`
	UserID       *types.Identifier     `json:"userId,omitempty"`
	Metadata     map[string]string     `json:"metadata,omitempty"`
	Comment      *string               `json:"comment,omitempty"`
	CreatedAt    *time.Time            `json:"createdAt,omitempty"`
	UpdatedAt    *time.Time            `json:"updatedAt,omitempty"`
}

func (s *ConnectKYCService) ReviewDocument(ctx context.Context, documentID string, opts *ConnectKYCReviewDocumentOptions) (*DocumentReview, *http.Response, error) {
	u := fmt.Sprintf("core-connect/kyc/documents/%s/status", documentID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPut, u, opts)

	d := new(DocumentReview)
	resp, err := s.client.Do(ctx, req, d)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	return d, resp, nil
}
