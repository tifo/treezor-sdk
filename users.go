package treezor

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/tifo/treezor-sdk/types"
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

// User represents a Treezor User.
type User struct {
	UserID                     *types.Identifier     `json:"userId,omitempty"`
	UserTypeID                 *types.Identifier     `json:"userTypeId,omitempty"` // NOTE: Can be an enum
	UserStatus                 *string               `json:"userStatus,omitempty"`
	ClientID                   *types.Identifier     `json:"clientId,omitempty"` // NOTE: Legacy + Webhook
	UserTag                    *string               `json:"userTag,omitempty"`
	ParentUserID               *types.Identifier     `json:"parentUserId,omitempty"`          // NOTE: Can be an enum
	ParentType                 *string               `json:"parentType,omitempty"`            // NOTE: Can be an enum
	ControllingPersonType      *types.Identifier     `json:"controllingPersonType,omitempty"` // NOTE: Can be an enum
	EmployeeType               *types.Identifier     `json:"employeeType,omitempty"`          // NOTE: Can be an enum
	EntityType                 *types.Identifier     `json:"entityType,omitempty"`            // NOTE: Can be an enum
	SpecifiedUSPerson          *types.Boolean        `json:"specifiedUSPerson,omitempty"`
	Title                      *string               `json:"title,omitempty"`
	Firstname                  *string               `json:"firstname,omitempty"`
	Lastname                   *string               `json:"lastname,omitempty"`
	MiddleNames                *string               `json:"middleNames,omitempty"`
	Birthday                   *types.Date           `json:"birthday,omitempty"`
	Email                      *string               `json:"email,omitempty"`
	Address1                   *string               `json:"address1,omitempty"`
	Address2                   *string               `json:"address2,omitempty"`
	Address3                   *string               `json:"address3,omitempty"`
	Postcode                   *string               `json:"postcode,omitempty"`
	City                       *string               `json:"city,omitempty"`
	State                      *string               `json:"state,omitempty"`
	Country                    *string               `json:"country,omitempty"`
	CountryName                *string               `json:"countryName,omitempty"`
	Phone                      *string               `json:"phone,omitempty"`
	Mobile                     *string               `json:"mobile,omitempty"`
	Nationality                *string               `json:"nationality,omitempty"`
	NationalityOther           *string               `json:"nationalityOther,omitempty"`
	PlaceOfBirth               *string               `json:"placeOfBirth,omitempty"`
	BirthCountry               *string               `json:"birthCountry,omitempty"`
	Occupation                 *string               `json:"occupation,omitempty"`
	IncomeRange                *string               `json:"incomeRange,omitempty"`
	LegalName                  *string               `json:"legalName,omitempty"`
	LegalNameEmbossed          *string               `json:"legalNameEmbossed,omitempty"`
	LegalRegistrationNumber    *string               `json:"legalRegistrationNumber,omitempty"`
	LegalTvaNumber             *string               `json:"legalTvaNumber,omitempty"`
	LegalRegistrationDate      *types.Date           `json:"legalRegistrationDate,omitempty"`
	LegalForm                  *string               `json:"legalForm,omitempty"`
	LegalShareCapital          *types.Integer        `json:"legalShareCapital,omitempty"`
	LegalSector                *string               `json:"legalSector,omitempty"`
	LegalAnnualTurnOver        *string               `json:"legalAnnualTurnOver,omitempty"`
	LegalNetIncomeRange        *string               `json:"legalNetIncomeRange,omitempty"`
	LegalNumberOfEmployeeRange *string               `json:"legalNumberOfEmployeeRange,omitempty"`
	EffectiveBeneficiary       *types.Identifier     `json:"effectiveBeneficiary,omitempty"`
	KycLevel                   *types.Level          `json:"kycLevel,omitempty"`
	KycReview                  *types.Review         `json:"kycReview,omitempty"`
	KycReviewComment           *string               `json:"kycReviewComment,omitempty"`
	IsFreezed                  *types.Boolean        `json:"isFreezed,omitempty"`
	IsFrozen                   *types.Boolean        `json:"isFrozen,omitempty"` // NOTE: Not populated
	Language                   *string               `json:"language,omitempty"`
	OptInMailing               *types.Boolean        `json:"optInMailing,omitempty"`
	SepaCreditorIdentifier     *string               `json:"sepaCreditorIdentifier,omitempty"`
	TaxNumber                  *string               `json:"taxNumber,omitempty"`
	TaxResidence               *string               `json:"taxResidence,omitempty"`
	Position                   *string               `json:"position,omitempty"`
	PersonalAssets             *string               `json:"personalAssets,omitempty"`
	ActivityOutsideEu          *types.Boolean        `json:"activityOutsideEu,omitempty"`
	EconomicSanctions          *types.Boolean        `json:"economicSanctions,omitempty"`
	ResidentCountriesSanctions *types.Boolean        `json:"residentCountriesSanctions,omitempty"`
	InvolvedSanctions          *types.Boolean        `json:"involvedSanctions,omitempty"`
	SanctionsQuestionnaireDate *types.Date           `json:"sanctionsQuestionnaireDate,omitempty"`
	Timezone                   *string               `json:"timezone,omitempty"`
	CreatedDate                *types.TimestampParis `json:"createdDate,omitempty"`
	ModifiedDate               *types.TimestampParis `json:"modifiedDate,omitempty"`
	WalletCount                *types.Integer        `json:"walletCount,omitempty"`
	PayinCount                 *types.Integer        `json:"payinCount,omitempty"`
	TotalRows                  *types.Integer        `json:"totalRows,omitempty"`
	CodeStatus                 *types.Identifier     `json:"codeStatus,omitempty"`        // NOTE: Legacy + Webhook
	InformationStatus          *string               `json:"informationStatus,omitempty"` // NOTE: Legacy + Webhook
}

type Level int32

func (l *Level) UnmarshalJSON(data []byte) error {
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	v, err := str.Int64()
	if err != nil {
		return err
	}
	*l = Level(v)
	return nil
}

const (
	LevelNone          Level = 0
	LevelPending       Level = 1
	LevelRegular       Level = 2
	LevelStrong        Level = 3
	LevelRefused       Level = 4
	LevelInvestigating Level = 5
)

var levelNames = map[int32]string{
	0: "LEVEL_NONE",
	1: "LEVEL_PENDING",
	2: "LEVEL_REGULAR",
	3: "LEVEL_STRONG",
	4: "LEVEL_REFUSED",
	5: "LEVEL_INVESTIGATING",
}

func (l Level) String() string {
	s, ok := levelNames[int32(l)]
	if ok {
		return s
	}
	return strconv.Itoa(int(l))
}

type Review int32

func (r *Review) UnmarshalJSON(data []byte) error {
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	v, err := str.Int64()
	if err != nil {
		return err
	}
	*r = Review(v)
	return nil
}

const (
	ReviewNone      Review = 0
	ReviewPending   Review = 1
	ReviewValidated Review = 2
	ReviewRefused   Review = 3
)

var reviewNames = map[int32]string{
	0: "REVIEW_NONE",
	1: "REVIEW_PENDING",
	2: "REVIEW_VALIDATED",
	3: "REVIEW_REFUSED",
}

func (r Review) String() string {
	s, ok := reviewNames[int32(r)]
	if ok {
		return s
	}
	return strconv.Itoa(int(r))
}

// Create creates a Treezor user.
func (s *UserService) Create(ctx context.Context, user *User) (*User, *http.Response, error) {
	req, _ := s.client.NewRequest(http.MethodPost, "users", user)

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

// Get fetches a user from Treezor.
func (s *UserService) Get(ctx context.Context, userID string) (*User, *http.Response, error) {
	u := fmt.Sprintf("users/%s", userID)
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
	UserID                string `url:"userId,omitempty"`
	UserTypeID            string `url:"userTypeId,omitempty"`
	UserStatus            string `url:"userStatus,omitempty"`
	UserTag               string `url:"userTag,omitempty"`
	SpecifiedUSPerson     string `url:"specifiedUSPerson,omitempty"`
	ControllingPersonType string `url:"controllingPersonType,omitempty"`
	EmployeeType          string `url:"employeeType,omitempty"`
	Email                 string `url:"email,omitempty"`
	Name                  string `url:"name,omitempty"`
	LegalName             string `url:"legalName,omitempty"`
	ParentUserID          string `url:"parentUserId,omitempty"`

	ListOptions
}

// List returns a list of users.
func (s *UserService) List(ctx context.Context, opt *UserListOptions) (*UserResponse, *http.Response, error) {
	u := "users"
	u, err := addOptions(u, opt)
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

// Edit updates a user.
func (s *UserService) Edit(ctx context.Context, userID string, user *User) (*User, *http.Response, error) {
	u := fmt.Sprintf("users/%s", userID)
	req, _ := s.client.NewRequest(http.MethodPut, u, user)

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

// ReviewKYC asks Treezor to do a KYC review against that user.
func (s *UserService) ReviewKYC(ctx context.Context, userID string) (*User, *http.Response, error) {
	u := fmt.Sprintf("users/%s/Kycreview/", userID)
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
func (s *UserService) RequestKYCLiveness(ctx context.Context, treezorUserID string) (*Identification, *http.Response, error) {
	u := fmt.Sprintf("users/%s/kycliveness", treezorUserID)
	req, _ := s.client.NewRequest(http.MethodPut, u, nil)

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

// ReviewKYCLiveness asks Treezor to do a KYC review against that user.
func (s *UserService) ReviewKYCLiveness(ctx context.Context, treezorUserID string) (*http.Response, error) {
	u := fmt.Sprintf("users/%s/kycliveness", treezorUserID)
	req, _ := s.client.NewRequest(http.MethodPut, u, nil)

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, errors.WithStack(err)
	}

	return resp, nil
}

// UserCancelOptions contains options for deletion of a user.
// Origin can be of value OPERATOR or USER.
type UserCancelOptions struct {
	Origin Origin `url:"origin,omitempty"`
}

// Cancel makes a User cancelled, meaning all future operation for that user
// will be refused.
func (s *UserService) Cancel(ctx context.Context, userID string, opt *UserCancelOptions) (*User, *http.Response, error) {
	u := fmt.Sprintf("users/%s", userID)
	u, err := addOptions(u, opt)
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
