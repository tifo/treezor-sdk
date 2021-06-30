package event

import "github.com/tifo/treezor-sdk/internal/types"

type CountryRestrictionGroupEvent struct {
	CountryRestrictionGroups []*CountryRestrictionGroup `json:"countryRestrictionGroups"`
}

type CountryRestrictionGroup struct {
	ID           *types.Identifier     `json:"id,omitempty"`
	Name         *string               `json:"name,omitempty"`
	IsWhitelist  *types.Boolean        `json:"isWhitelist,omitempty"`
	Countries    []*types.Identifier   `json:"merchants,omitempty"`
	Status       *string               `json:"status,omitempty"` // NOTE: can be an enum
	StartDate    *types.TimestampParis `json:"startDate,omitempty"`
	CreatedDate  *types.TimestampParis `json:"createdDate,omitempty"`
	ModifiedDate *types.TimestampParis `json:"modifiedDate,omitempty"`
}

// TODO: Add CountryRestrictionGroup API
