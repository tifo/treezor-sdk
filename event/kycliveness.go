package event

import "time"

// KYCLivenessEvent represents a kyc liveness event
type KYCLivenessEvent struct {
	StartedAt *time.Time           `json:"started-at,omitempty"`
	UpdatedAt *time.Time           `json:"updated-at,omitempty"`
	Identity  *KYCLivenessIdentity `json:"identity"`
	KycStatus *string              `json:"kyc-status,omitempty"`
	Comment   *string              `json:"comment,omitempty"`
	UserID    *string              `json:"user_id,omitempty"`
	Score     *int                 `json:"score,omitempty"`
}

// KYCLivenessIdentity represents the detected identity in a kyc liveness
type KYCLivenessIdentity struct {
	LastName  *string `json:"last-name,omitempty"`
	FirstName *string `json:"first-name,omitempty"`
	BirthDate *string `json:"birth-date,omitempty"`
}
