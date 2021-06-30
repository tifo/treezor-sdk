package treezor

import "github.com/tifo/treezor-sdk/internal/types"

type BankAccount struct {
	BankAccountID           *types.Identifier     `json:"bankaccountId,omitempty"`
	BankAccountTag          *string               `json:"bankaccountTag,omitempty"`
	BankAccountStatus       *string               `json:"bankaccountStatus,omitempty"` // NOTE: can be an enum
	UserID                  *types.Identifier     `json:"userId,omitempty"`
	Name                    *string               `json:"name,omitempty"`
	BankAccountOwnerName    *string               `json:"bankaccountOwnerName,omitempty"`
	BankAccountOwnerAddress *string               `json:"bankaccountOwnerAddress,omitempty"`
	BankAccountIBAN         *string               `json:"bankaccountIBAN,omitempty"`
	BankAccountBIC          *string               `json:"bankaccountBIC,omitempty"`
	BankAccountType         *string               `json:"bankaccountType,omitempty"`
	CreatedDate             *types.TimestampParis `json:"createdDate,omitempty"`
	ModifiedDate            *types.TimestampParis `json:"modifiedDate,omitempty"`
	TotalRows               *types.Integer        `json:"totalRows,omitempty"`
	CodeStatus              *types.Identifier     `json:"codeStatus,omitempty"`        // Legacy field
	InformationStatus       *string               `json:"informationStatus,omitempty"` // Legacy field
}

// TODO: Add BankAccount API
