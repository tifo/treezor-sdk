package event

import (
	"time"

	"github.com/tifo/treezor-sdk/internal/types"
)

type CardDigitalizationEvent struct {
	DeviceName                   *string           `json:"deviceName,omitempty"`
	DeviceType                   *string           `json:"deviceType,omitempty"`     // NOTE: can be an enum
	TokenRequestor               *string           `json:"tokenRequestor,omitempty"` // NOTE: can be an enum
	CardDigitalizationExternalID *string           `json:"cardDigitalizationExternalId,omitempty"`
	CardID                       *types.Identifier `json:"cardId,omitempty"`
	ActivationCode               *string           `json:"activactionCode,omitempty"`
	ActivationCodeExpiry         *time.Time        `json:"activationCodeExpiry,omitempty"`
	ActivationMethod             *string           `json:"activationMethod,omitempty"` // NOTE: can be an enum
	ExpirationDate               *types.Date       `json:"expirationDate,omitempty"`
}

// TODO: Add CardDigitalization API
