package treezor

import "time"

// KycLiveness represent a kyc liveness event
type KycLiveness struct {
	Access
	StartedAt *time.Time           `json:"started-at,omitempty"`
	UpdatedAt *time.Time           `json:"updated-at,omitempty"`
	Identity  *KycLivenessIdentity `json:"identity"`
	KycStatus *string              `json:"kyc-status,omitempty"`
	Comment   *string              `json:"comment,omitempty"`
	UserID    *string              `json:"user_id,omitempty"`
	Score     *int                 `json:"score,omitempty"`
}

// KycLivenessIdentity represent a detect identity in a kyc liveness
type KycLivenessIdentity struct {
	LastName  *string `json:"last-name,omitempty"`
	FirstName *string `json:"first-name,omitempty"`
	BirthDate *string `json:"birth-date,omitempty"`
}
