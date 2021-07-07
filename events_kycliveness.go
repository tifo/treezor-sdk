package treezor

import (
	"time"

	"github.com/tifo/treezor-sdk/internal/types"
)

// KYCLivenessEvent represents a kyc liveness event
type KYCLivenessEvent struct {
	StartedAt *time.Time           `json:"started-at,omitempty"`
	UpdatedAt *time.Time           `json:"updated-at,omitempty"`
	Identity  *KYCLivenessIdentity `json:"identity"`
	KYCStatus *string              `json:"kyc-status,omitempty"`
	Comment   *string              `json:"comment,omitempty"`
	UserID    *string              `json:"user_id,omitempty"`
	Score     *types.Integer       `json:"score,omitempty"`
}

// KYCLivenessIdentity represents the detected identity in a kyc liveness
type KYCLivenessIdentity struct {
	Lastname  *string `json:"last-name,omitempty"`
	Firstname *string `json:"first-name,omitempty"`
	Birthdate *string `json:"birth-date,omitempty"`
}
