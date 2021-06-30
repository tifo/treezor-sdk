package event

import "github.com/tifo/treezor-sdk/internal/types"

type MandateEvent struct {
	Mandates []*Mandate `json:"mandates"`
}

type Mandate struct {
	MandateID                           *types.Identifier     `json:"mandateId,omitempty"`
	Title                               *string               `json:"title,omitempty"`
	LegalInformations                   *string               `json:"legalInformations,omitempty"`
	UniqueMandateReference              *string               `json:"uniqueMandateReference,omitempty"`
	MandateStatus                       *string               `json:"mandateStatus,omitempty"` // NOTE: can be an en enum
	UserID                              *types.Identifier     `json:"userId,omitempty"`
	DebtorName                          *string               `json:"debtorName,omitempty"`
	DebtorAddress                       *string               `json:"debtorAddress,omitempty"`
	DebtorCity                          *string               `json:"debtorCity,omitempty"`
	DebtorZipCode                       *string               `json:"debtorZipCode,omitempty"`
	DebtorCountry                       *string               `json:"debtorCountry,omitempty"`
	DebtorIBAN                          *string               `json:"debtorIban,omitempty"`
	DebtorBIC                           *string               `json:"debtorBic,omitempty"`
	SequenceType                        *string               `json:"sequenceType,omitempty"` // NOTE: can be an enum
	CreditorName                        *string               `json:"creditorName,omitempty"`
	SepaCreditorIdentifier              *string               `json:"sepaCreditorIdentifier,omitempty"`
	CreditorAddress                     *string               `json:"creditorAddress,omitempty"`
	CreditorCity                        *string               `json:"creditorCity,omitempty"`
	CreditorZipCode                     *string               `json:"creditorZipCode,omitempty"`
	CreditorCountry                     *string               `json:"creditorCountry,omitempty"`
	SignatureDate                       *types.Date           `json:"signatureDate,omitempty"`
	DebtorSignatureIP                   *string               `json:"debtorSignatureIp,omitempty"`
	Signed                              *types.Boolean        `json:"signed,omitempty"`
	RevocationSignatureDate             *types.Date           `json:"revocationSignatureDate,omitempty"` // NOTE: might be a timestamp
	DebtorIdentificationCode            *string               `json:"debtorIdentificationCode,omitempty"`
	DebtorReferencePartyName            *string               `json:"debtorReferencePartyName,omitempty"`
	DebtorReferenceIdentificationCode   *string               `json:"debtorReferenceIdentificationCode,omitempty"`
	CreditorReferencePartyName          *string               `json:"creditorReferencePartyName,omitempty"`
	CreditorReferenceIdentificationCode *string               `json:"creditorReferenceIdentificationCode,omitempty"`
	ContractIdentificationNumber        *string               `json:"contractIdentificationNumber,omitempty"`
	ContractDescription                 *string               `json:"contractDescription,omitempty"`
	IsPaper                             *types.Boolean        `json:"isPaper,omitempty"`
	SDDType                             *string               `json:"sddType,omitempty"` // NOTE: can be an enum
	UserIDUltimateCreditor              *types.Identifier     `json:"userIdUltimateCreditor,omitempty"`
	CreatedIP                           *string               `json:"createdIp,omitempty"`
	CreatedDate                         *types.TimestampParis `json:"createdDate,omitempty"`
	ModifiedDate                        *types.TimestampParis `json:"modifiedDate,omitempty"`
	CodeStatus                          *types.Identifier     `json:"codeStatus,omitempty"`        // Legacy field
	InformationStatus                   *string               `json:"informationStatus,omitempty"` // Legacy field
}

// TODO: Add Mandate API & Models
// TODO: see presence of totalRows, codeStatus and informationStatus in all models

// NOTE: we might need to allow dates to include a time just in case we failed in setting up the type or the date includes 00:00:00
