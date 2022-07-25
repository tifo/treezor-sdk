package treezor

import (
	"time"

	"github.com/tifo/treezor-sdk/internal/types"
)

type MIDRestrictionGroup struct {
	ID           *types.Identifier   `json:"id,omitempty"`
	Name         *string             `json:"name,omitempty"`
	IsWhitelist  *types.Boolean      `json:"isWhitelist,omitempty"`
	Merchants    []*types.Identifier `json:"merchants,omitempty"`
	Status       *string             `json:"status,omitempty"` // NOTE: can be an enum
	StartDate    *time.Time          `json:"startDate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	CreatedDate  *time.Time          `json:"createdDate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
	ModifiedDate *time.Time          `json:"modifiedDate,omitempty" layout:"Treezor" loc:"Europe/Paris"`
}

// TODO: Add MerchantIDRestrictionGroup API
