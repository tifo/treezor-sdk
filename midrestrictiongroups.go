package treezor

import "github.com/tifo/treezor-sdk/internal/types"

type MIDRestrictionGroup struct {
	ID           *types.Identifier     `json:"id,omitempty"`
	Name         *string               `json:"name,omitempty"`
	IsWhitelist  *types.Boolean        `json:"isWhitelist,omitempty"`
	Merchants    []*types.Identifier   `json:"merchants,omitempty"`
	Status       *string               `json:"status,omitempty"` // NOTE: can be an enum
	StartDate    *types.TimestampParis `json:"startDate,omitempty"`
	CreatedDate  *types.TimestampParis `json:"createdDate,omitempty"`
	ModifiedDate *types.TimestampParis `json:"modifiedDate,omitempty"`
}

// TODO: Add MerchantIDRestrictionGroup API
