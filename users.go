package treezor

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/tifo/treezor-sdk/internal/types"
)

// UserService handles communication with the user related
// methods of the Treezor API.
//
// Treezor API docs: https://www.treezor.com/api-documentation/#/user
type UserService service

// UserResponse represents a list of users.
type UserResponse struct {
	Users []*User `json:"users"`
}

// UserType defines the type of a user.
type UserType int32

const (
	UserTypePhyisical                    UserType = 1 // Physical User and Anonymous User
	UserTypeBusiness                     UserType = 2 // Business User
	UserTypeNonGovernementalOrganization UserType = 3 // Non-governemental organization
	UserTypeGovernementalOrganization    UserType = 4 // Governemental organization
)

func (t *UserType) UnmarshalJSON(data []byte) error {
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	v, err := str.Int64()
	if err != nil {
		return err
	}
	*t = UserType(v)
	return nil
}

// ParentType defines the type of relation with the parent, used for actual family relations
type ParentType string

const (
	ParentTypeShareholder ParentType = "shareholder" // Shareholder of a Business User
	ParentTypeEmployee    ParentType = "employee"    // Employee of a Business User
	ParentTypeLeader      ParentType = "leader"      // Legal representative of a Business User
)

// ControllingPersonType defines the type of relation with the parent, used with Shareholders and Legal representatives
type ControllingPersonType int32

const (
	ControllingPersonTypeShareholder         ControllingPersonType = 1 // Shareholder of a Business User
	ControllingPersonTypeLegalRepresentative ControllingPersonType = 3 // Legal representative of a Business User
)

func (t *ControllingPersonType) UnmarshalJSON(data []byte) error {
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	v, err := str.Int64()
	if err != nil {
		return err
	}
	*t = ControllingPersonType(v)
	return nil
}

// EmployeeType defines the relationship between a user and its parent.
type EmployeeType int32

const (
	EmployeeTypeNone     EmployeeType = 0 // Not an employee
	EmployeeTypeLeader   EmployeeType = 1 // Legal representative
	EmployeeTypeEmployee EmployeeType = 2 // Employee
)

func (t *EmployeeType) UnmarshalJSON(data []byte) error {
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	v, err := str.Int64()
	if err != nil {
		return err
	}
	*t = EmployeeType(v)
	return nil
}

// EntityType defines the FATCA/CRS classification of a business user.
type EntityType int32

const (
	EntityTypeReportingFinancialInstitution     EntityType = 1 // Reporting Financial Institution
	EntityTypeNonReportingFinancialInstituition EntityType = 2 // Non-Reporting Financial Institution
	EntityTypeGovOrIntActiveNonFinancialEntity  EntityType = 3 // Active Non-Financial Entity - Governmental entities, Int. organizations
	EntityTypeOtherActiveNonFinancialEntity     EntityType = 4 // Active Non-Financial Entity - Other
	EntityTypePassiveNonFinancialEntity         EntityType = 5 // Passive Non-Financial Entity - Investment entity that is not Participating Jurisdiction FI
)

func (t *EntityType) UnmarshalJSON(data []byte) error {
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	v, err := str.Int64()
	if err != nil {
		return err
	}
	*t = EntityType(v)
	return nil
}

// KYCReview defines the review status of a user's kyc.
type KYCReview int32

const (
	KYCReviewNone      KYCReview = 0
	KYCReviewPending   KYCReview = 1
	KYCReviewValidated KYCReview = 2
	KYCReviewRefused   KYCReview = 3
)

var kycReviewNames = map[int32]string{
	0: "NONE",
	1: "PENDING",
	2: "VALIDATED",
	3: "REFUSED",
}

func (r *KYCReview) UnmarshalJSON(data []byte) error {
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	v, err := str.Int64()
	if err != nil {
		return err
	}
	*r = KYCReview(v)
	return nil
}

func (r KYCReview) String() string {
	s, ok := kycReviewNames[int32(r)]
	if ok {
		return s
	}
	return strconv.Itoa(int(r))
}

// KYCLevel defines the trust level of a user's kyc.
type KYCLevel int32

const (
	KYCLevelNone          KYCLevel = 0
	KYCLevelPending       KYCLevel = 1
	KYCLevelRegular       KYCLevel = 2
	KYCLevelStrong        KYCLevel = 3
	KYCLevelRefused       KYCLevel = 4
	KYCLevelInvestigating KYCLevel = 5
)

var kycLevelNames = map[int32]string{
	0: "LEVEL_NONE",
	1: "LEVEL_PENDING",
	2: "LEVEL_REGULAR",
	3: "LEVEL_STRONG",
	4: "LEVEL_REFUSED",
	5: "LEVEL_INVESTIGATING",
}

func (l *KYCLevel) UnmarshalJSON(data []byte) error {
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	v, err := str.Int64()
	if err != nil {
		return err
	}
	*l = KYCLevel(v)
	return nil
}

func (l KYCLevel) String() string {
	s, ok := kycLevelNames[int32(l)]
	if ok {
		return s
	}
	return strconv.Itoa(int(l))
}

// User represents a Treezor User.
type User struct {
	UserID                     *types.Identifier      `json:"userId,omitempty"`
	UserTypeID                 *UserType              `json:"userTypeId,omitempty"`
	UserStatus                 *string                `json:"userStatus,omitempty"`
	ClientID                   *types.Identifier      `json:"clientId,omitempty"`
	UserTag                    *string                `json:"userTag,omitempty"`
	ParentUserID               *types.Identifier      `json:"parentUserId,omitempty"`
	ParentType                 *ParentType            `json:"parentType,omitempty"`
	ControllingPersonType      *ControllingPersonType `json:"controllingPersonType,omitempty"`
	EmployeeType               *EmployeeType          `json:"employeeType,omitempty"`
	EntityType                 *EntityType            `json:"entityType,omitempty"`
	SpecifiedUSPerson          *types.Boolean         `json:"specifiedUSPerson,omitempty"`
	Title                      *string                `json:"title,omitempty"`
	Firstname                  *string                `json:"firstname,omitempty"`
	Lastname                   *string                `json:"lastname,omitempty"`
	MiddleNames                *string                `json:"middleNames,omitempty"`
	Birthday                   *types.Date            `json:"birthday,omitempty"`
	Email                      *string                `json:"email,omitempty"`
	Address1                   *string                `json:"address1,omitempty"`
	Address2                   *string                `json:"address2,omitempty"`
	Address3                   *string                `json:"address3,omitempty"`
	Postcode                   *string                `json:"postcode,omitempty"`
	City                       *string                `json:"city,omitempty"`
	State                      *string                `json:"state,omitempty"`
	Country                    *string                `json:"country,omitempty"`
	CountryName                *string                `json:"countryName,omitempty"`
	Phone                      *string                `json:"phone,omitempty"`
	Mobile                     *string                `json:"mobile,omitempty"`
	Nationality                *string                `json:"nationality,omitempty"`
	NationalityOther           *string                `json:"nationalityOther,omitempty"`
	PlaceOfBirth               *string                `json:"placeOfBirth,omitempty"`
	BirthCountry               *string                `json:"birthCountry,omitempty"`
	Occupation                 *string                `json:"occupation,omitempty"`
	IncomeRange                *string                `json:"incomeRange,omitempty"`
	LegalName                  *string                `json:"legalName,omitempty"`
	LegalNameEmbossed          *string                `json:"legalNameEmbossed,omitempty"`
	LegalRegistrationNumber    *string                `json:"legalRegistrationNumber,omitempty"`
	LegalTVANumber             *string                `json:"legalTvaNumber,omitempty"`
	LegalRegistrationDate      *types.Date            `json:"legalRegistrationDate,omitempty"`
	LegalForm                  *string                `json:"legalForm,omitempty"`
	LegalShareCapital          *types.Integer         `json:"legalShareCapital,omitempty"`
	LegalSector                *string                `json:"legalSector,omitempty"`
	LegalAnnualTurnOver        *string                `json:"legalAnnualTurnOver,omitempty"`
	LegalNetIncomeRange        *string                `json:"legalNetIncomeRange,omitempty"`
	LegalNumberOfEmployeeRange *string                `json:"legalNumberOfEmployeeRange,omitempty"`
	EffectiveBeneficiary       *types.Percentage      `json:"effectiveBeneficiary,omitempty"`
	KycLevel                   *KYCLevel              `json:"kycLevel,omitempty"`
	KycReview                  *KYCReview             `json:"kycReview,omitempty"`
	KycReviewComment           *string                `json:"kycReviewComment,omitempty"`
	IsFreezed                  *types.Boolean         `json:"isFreezed,omitempty"` // Deprecated
	IsFrozen                   *types.Boolean         `json:"isFrozen,omitempty"`
	Language                   *string                `json:"language,omitempty"`
	OptInMailing               *types.Boolean         `json:"optInMailing,omitempty"`
	SepaCreditorIdentifier     *string                `json:"sepaCreditorIdentifier,omitempty"`
	TaxNumber                  *string                `json:"taxNumber,omitempty"`
	TaxResidence               *string                `json:"taxResidence,omitempty"`
	Position                   *string                `json:"position,omitempty"`
	PersonalAssets             *string                `json:"personalAssets,omitempty"`
	ActivityOutsideEu          *types.Boolean         `json:"activityOutsideEu,omitempty"`
	EconomicSanctions          *types.Boolean         `json:"economicSanctions,omitempty"`
	ResidentCountriesSanctions *types.Boolean         `json:"residentCountriesSanctions,omitempty"`
	InvolvedSanctions          *types.Boolean         `json:"involvedSanctions,omitempty"`
	SanctionsQuestionnaireDate *types.Date            `json:"sanctionsQuestionnaireDate,omitempty"`
	Timezone                   *string                `json:"timezone,omitempty"`
	CreatedDate                *types.TimestampParis  `json:"createdDate,omitempty"`
	ModifiedDate               *types.TimestampParis  `json:"modifiedDate,omitempty"`
	WalletCount                *types.Integer         `json:"walletCount,omitempty"`
	PayinCount                 *types.Integer         `json:"payinCount,omitempty"`
	TotalRows                  *types.Integer         `json:"totalRows,omitempty"`
	CodeStatus                 *types.Identifier      `json:"codeStatus,omitempty"`        // Legacy field
	InformationStatus          *string                `json:"informationStatus,omitempty"` // Legacy field
}

type UserCreateOptions struct {
	Access

	UserTypeID                 *UserType              `url:"-" json:"userTypeId,omitempty"`                 // Optional
	UserTag                    *string                `url:"-" json:"userTag,omitempty"`                    // Optional
	ParentUserID               *string                `url:"-" json:"parentUserId,omitempty"`               // Optional
	ParentType                 *ParentType            `url:"-" json:"parentType,omitempty"`                 // Optional
	SpecifiedUSPerson          *types.Boolean         `url:"-" json:"specifiedUSPerson,omitempty"`          // Required
	ControllingPersonType      *ControllingPersonType `url:"-" json:"controllingPersonType,omitempty"`      // Optional
	EmployeeType               *EmployeeType          `url:"-" json:"employeeType,omitempty"`               // Optional
	EntityType                 *EntityType            `url:"-" json:"entityType,omitempty"`                 // Optional
	Title                      *string                `url:"-" json:"title,omitempty"`                      // Optional
	Firstname                  *string                `url:"-" json:"firstname,omitempty"`                  // Optional
	Lastname                   *string                `url:"-" json:"lastname,omitempty"`                   // Optional
	MiddleNames                *string                `url:"-" json:"middleNames,omitempty"`                // Optional
	Birthday                   *types.Date            `url:"-" json:"birthday,omitempty"`                   // Optional
	Email                      *string                `url:"-" json:"email,omitempty"`                      // Required
	Address1                   *string                `url:"-" json:"address1,omitempty"`                   // Optional
	Address2                   *string                `url:"-" json:"address2,omitempty"`                   // Optional
	Address3                   *string                `url:"-" json:"address3,omitempty"`                   // Optional
	Postcode                   *string                `url:"-" json:"postcode,omitempty"`                   // Optional
	City                       *string                `url:"-" json:"city,omitempty"`                       // Optional
	State                      *string                `url:"-" json:"state,omitempty"`                      // Optional
	Country                    *string                `url:"-" json:"country,omitempty"`                    // Optional
	Phone                      *string                `url:"-" json:"phone,omitempty"`                      // Optional
	Mobile                     *string                `url:"-" json:"mobile,omitempty"`                     // Optional
	Nationality                *string                `url:"-" json:"nationality,omitempty"`                // Optional
	NationalityOther           *string                `url:"-" json:"nationalityOther,omitempty"`           // Optional
	PlaceOfBirth               *string                `url:"-" json:"placeOfBirth,omitempty"`               // Optional
	BirthCountry               *string                `url:"-" json:"birthCountry,omitempty"`               // Optional
	Occupation                 *string                `url:"-" json:"occupation,omitempty"`                 // Optional
	IncomeRange                *string                `url:"-" json:"incomeRange,omitempty"`                // Optional
	LegalName                  *string                `url:"-" json:"legalName,omitempty"`                  // Optional
	LegalRegistrationNumber    *string                `url:"-" json:"legalRegistrationNumber,omitempty"`    // Optional
	LegalTVANumber             *string                `url:"-" json:"legalTvaNumber,omitempty"`             // Optional
	LegalRegistrationDate      *types.Date            `url:"-" json:"legalRegistrationDate,omitempty"`      // Optional
	LegalForm                  *string                `url:"-" json:"legalForm,omitempty"`                  // Optional
	LegalShareCapital          *int64                 `url:"-" json:"legalShareCapital,omitempty"`          // Optional
	LegalSector                *string                `url:"-" json:"legalSector,omitempty"`                // Optional
	LegalAnnualTurnOver        *string                `url:"-" json:"legalAnnualTurnOver,omitempty"`        // Optional
	LegalNetIncomeRange        *string                `url:"-" json:"legalNetIncomeRange,omitempty"`        // Optional
	LegalNumberOfEmployeeRange *string                `url:"-" json:"legalNumberOfEmployeeRange,omitempty"` // Optional
	EffectiveBeneficiary       *float64               `url:"-" json:"effectiveBeneficiary,omitempty"`       // Optional
	Language                   *string                `url:"-" json:"language,omitempty"`                   // Optional
	TaxNumber                  *string                `url:"-" json:"taxNumber,omitempty"`                  // Optional
	TaxResidence               *string                `url:"-" json:"taxResidence,omitempty"`               // Optional
	Position                   *string                `url:"-" json:"position,omitempty"`                   // Optional
	PersonalAssets             *string                `url:"-" json:"personalAssets,omitempty"`             // Optional
	ActivityOutsideEu          *types.Boolean         `url:"-" json:"activityOutsideEu,omitempty"`          // Optional
	EconomicSanctions          *types.Boolean         `url:"-" json:"economicSanctions,omitempty"`          // Optional
	ResidentCountriesSanctions *types.Boolean         `url:"-" json:"residentCountriesSanctions,omitempty"` // Optional
	InvolvedSanctions          *types.Boolean         `url:"-" json:"involvedSanctions,omitempty"`          // Optional
	Timezone                   *string                `url:"-" json:"timezone,omitempty"`                   // Optional
}

// Create creates a Treezor user.
func (s *UserService) Create(ctx context.Context, opts *UserCreateOptions) (*User, *http.Response, error) {
	u := "users"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPost, u, opts)

	ur := new(UserResponse)
	resp, err := s.client.Do(ctx, req, ur)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(ur.Users) != 1 {
		return nil, resp, errors.Errorf("API did not return exactly one user: %d users returned", len(ur.Users))
	}
	return ur.Users[0], resp, nil
}

type UserGetOptions struct {
	Access
}

// Get fetches a user from Treezor.
func (s *UserService) Get(ctx context.Context, userID string, opts *UserGetOptions) (*User, *http.Response, error) {
	u := fmt.Sprintf("users/%s", userID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	ur := new(UserResponse)
	resp, err := s.client.Do(ctx, req, ur)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(ur.Users) != 1 {
		return nil, resp, errors.Errorf("API did not return exactly one user: %d users returned", len(ur.Users))
	}
	return ur.Users[0], resp, nil
}

// UserListOptions contains options for listing users.
type UserListOptions struct {
	Access

	UserID                *string                `url:"userId,omitempty" json:"-"`
	UserTypeID            *UserType              `url:"userTypeId,omitempty" json:"-"`
	UserStatus            *string                `url:"userStatus,omitempty" json:"-"` // NOTE: can be an enum (need to see if VALIDATED or Validated)
	UserTag               *string                `url:"userTag,omitempty" json:"-"`
	SpecifiedUSPerson     *types.Boolean         `url:"specifiedUSPerson,omitempty" json:"-"`
	ControllingPersonType *ControllingPersonType `url:"controllingPersonType,omitempty" json:"-"`
	EmployeeType          *EmployeeType          `url:"employeeType,omitempty" json:"-"`
	Email                 *string                `url:"email,omitempty" json:"-"`
	Name                  *string                `url:"name,omitempty" json:"-"`
	LegalName             *string                `url:"legalName,omitempty" json:"-"`
	ParentUserID          *string                `url:"parentUserId,omitempty" json:"-"`

	ListOptions
}

// List returns a list of users.
func (s *UserService) List(ctx context.Context, opts *UserListOptions) (*UserResponse, *http.Response, error) {
	u := "users"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	ur := new(UserResponse)
	resp, err := s.client.Do(ctx, req, ur)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	return ur, resp, nil
}

type UserEditOptions struct {
	Access

	UserTag                    *string                `url:"-" json:"userTag,omitempty"`                    // Optional
	SpecifiedUSPerson          *types.Boolean         `url:"-" json:"specifiedUSPerson,omitempty"`          // Required
	ControllingPersonType      *ControllingPersonType `url:"-" json:"controllingPersonType,omitempty"`      // Optional
	EmployeeType               *EmployeeType          `url:"-" json:"employeeType,omitempty"`               // Optional
	Title                      *string                `url:"-" json:"title,omitempty"`                      // Optional
	Firstname                  *string                `url:"-" json:"firstname,omitempty"`                  // Optional
	Lastname                   *string                `url:"-" json:"lastname,omitempty"`                   // Optional
	MiddleNames                *string                `url:"-" json:"middleNames,omitempty"`                // Optional
	Birthday                   *types.Date            `url:"-" json:"birthday,omitempty"`                   // Optional
	Email                      *string                `url:"-" json:"email,omitempty"`                      // Optional
	Address1                   *string                `url:"-" json:"address1,omitempty"`                   // Optional
	Address2                   *string                `url:"-" json:"address2,omitempty"`                   // Optional
	Address3                   *string                `url:"-" json:"address3,omitempty"`                   // Optional
	Postcode                   *string                `url:"-" json:"postcode,omitempty"`                   // Optional
	City                       *string                `url:"-" json:"city,omitempty"`                       // Optional
	State                      *string                `url:"-" json:"state,omitempty"`                      // Optional
	Country                    *string                `url:"-" json:"country,omitempty"`                    // Optional
	Phone                      *string                `url:"-" json:"phone,omitempty"`                      // Optional
	Mobile                     *string                `url:"-" json:"mobile,omitempty"`                     // Optional
	Nationality                *string                `url:"-" json:"nationality,omitempty"`                // Optional
	NationalityOther           *string                `url:"-" json:"nationalityOther,omitempty"`           // Optional
	PlaceOfBirth               *string                `url:"-" json:"placeOfBirth,omitempty"`               // Optional
	BirthCountry               *string                `url:"-" json:"birthCountry,omitempty"`               // Optional
	Occupation                 *string                `url:"-" json:"occupation,omitempty"`                 // Optional
	IncomeRange                *string                `url:"-" json:"incomeRange,omitempty"`                // Optional
	LegalName                  *string                `url:"-" json:"legalName,omitempty"`                  // Optional
	LegalRegistrationNumber    *string                `url:"-" json:"legalRegistrationNumber,omitempty"`    // Optional
	LegalTVANumber             *string                `url:"-" json:"legalTvaNumber,omitempty"`             // Optional
	LegalRegistrationDate      *types.Date            `url:"-" json:"legalRegistrationDate,omitempty"`      // Optional
	LegalForm                  *string                `url:"-" json:"legalForm,omitempty"`                  // Optional
	LegalShareCapital          *int64                 `url:"-" json:"legalShareCapital,omitempty"`          // Optional
	LegalSector                *string                `url:"-" json:"legalSector,omitempty"`                // Optional
	LegalAnnualTurnOver        *string                `url:"-" json:"legalAnnualTurnOver,omitempty"`        // Optional
	LegalNetIncomeRange        *string                `url:"-" json:"legalNetIncomeRange,omitempty"`        // Optional
	LegalNumberOfEmployeeRange *string                `url:"-" json:"legalNumberOfEmployeeRange,omitempty"` // Optional
	EffectiveBeneficiary       *float64               `url:"-" json:"effectiveBeneficiary,omitempty"`       // Optional
	Language                   *string                `url:"-" json:"language,omitempty"`                   // Optional
	TaxNumber                  *string                `url:"-" json:"taxNumber,omitempty"`                  // Optional
	TaxResidence               *string                `url:"-" json:"taxResidence,omitempty"`               // Optional
	Position                   *string                `url:"-" json:"position,omitempty"`                   // Optional
	PersonalAssets             *string                `url:"-" json:"personalAssets,omitempty"`             // Optional
	ActivityOutsideEu          *types.Boolean         `url:"-" json:"activityOutsideEu,omitempty"`          // Optional
	EconomicSanctions          *types.Boolean         `url:"-" json:"economicSanctions,omitempty"`          // Optional
	ResidentCountriesSanctions *types.Boolean         `url:"-" json:"residentCountriesSanctions,omitempty"` // Optional
	InvolvedSanctions          *types.Boolean         `url:"-" json:"involvedSanctions,omitempty"`          // Optional
	Timezone                   *string                `url:"-" json:"timezone,omitempty"`                   // Optional
}

// Edit updates a user.
func (s *UserService) Edit(ctx context.Context, userID string, opts *UserEditOptions) (*User, *http.Response, error) {
	u := fmt.Sprintf("users/%s", userID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPut, u, opts)

	ur := new(UserResponse)
	resp, err := s.client.Do(ctx, req, ur)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(ur.Users) != 1 {
		return nil, resp, errors.Errorf("API did not return exactly one user: %d users returned", len(ur.Users))
	}
	return ur.Users[0], resp, nil
}

// UserCancelOptions contains options for deletion of a user.
// Origin can be of value OPERATOR or USER.
type UserCancelOptions struct {
	Access

	Origin *Origin `url:"origin"` // Required
}

// Cancel makes a User cancelled, meaning all future operation for that user
// will be refused.
func (s *UserService) Cancel(ctx context.Context, userID string, opts *UserCancelOptions) (*User, *http.Response, error) {
	u := fmt.Sprintf("users/%s", userID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodDelete, u, nil)

	ur := new(UserResponse)
	resp, err := s.client.Do(ctx, req, ur)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(ur.Users) != 1 {
		return nil, resp, errors.Errorf("API did not return exactly one user: %d users returned", len(ur.Users))
	}
	return ur.Users[0], resp, nil
}

type UserReviewKYCOptions struct {
	Access
}

// ReviewKYC asks Treezor to do a KYC review against that user.
func (s *UserService) ReviewKYC(ctx context.Context, userID string, opts *UserReviewKYCOptions) (*User, *http.Response, error) {
	u := fmt.Sprintf("users/%s/Kycreview/", userID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPut, u, nil)

	ur := new(UserResponse)
	resp, err := s.client.Do(ctx, req, ur)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(ur.Users) != 1 {
		return nil, resp, errors.Errorf("API did not return exactly one user: %d users returned", len(ur.Users))
	}
	return ur.Users[0], resp, nil
}

type UserRequestKYCLivenessOptions struct {
	Access

	RedirectURL *string `url:"-" json:"redirect_url"`
}

// IdentificationResponse represent a list of identification
type IdentificationResponse struct {
	Identification *Identification `json:"identification,omitempty"`
}

// Identification represent an identification returned by Treezor for a kycLiveness request
type Identification struct {
	IdentificationID  *string `json:"identification-id,omitempty"`
	IdentificationURL *string `json:"identification-url,omitempty"`
}

// RequestKYCLiveness makes a kyc url request for the kycliveness process.
func (s *UserService) RequestKYCLiveness(ctx context.Context, userID string, opts *UserRequestKYCLivenessOptions) (*Identification, *http.Response, error) {
	u := fmt.Sprintf("users/%s/kycliveness", userID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPost, u, opts)

	k := new(IdentificationResponse)
	resp, err := s.client.Do(ctx, req, k)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if k.Identification == nil {
		return nil, resp, errors.New("API did not return a valid identification")
	}
	return k.Identification, resp, nil
}

type UserReviewKYCLivenessOptions struct {
	Access
}

// ReviewKYCLiveness asks Treezor to do a KYC review against that user.
func (s *UserService) ReviewKYCLiveness(ctx context.Context, userID string, opts *UserReviewKYCLivenessOptions) (*http.Response, error) {
	u := fmt.Sprintf("users/%s/kycliveness", userID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPut, u, opts)

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, errors.WithStack(err)
	}

	return resp, nil
}
