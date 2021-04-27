package treezor

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
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
	Access
	UserID                     *string         `json:"userId,omitempty"`
	UserTypeID                 *string         `json:"userTypeId,omitempty"`
	UserStatus                 *string         `json:"userStatus,omitempty"`
	ParentUserID               *string         `json:"parentUserId,omitempty"`
	ParentType                 *string         `json:"parentType,omitempty"`
	ControllingPersonType      *string         `json:"controllingPersonType,omitempty"`
	EmployeeType               *string         `json:"employeeType,omitempty"`
	ClientID                   *string         `json:"clientId,omitempty"`
	UserTag                    *string         `json:"userTag,omitempty"`
	SpecifiedUSPerson          *string         `json:"specifiedUSPerson,omitempty"`
	Title                      *string         `json:"title,omitempty"`
	Firstname                  *string         `json:"firstname,omitempty"`
	Lastname                   *string         `json:"lastname,omitempty"`
	MiddleNames                *string         `json:"middleNames,omitempty"`
	Birthday                   *Date           `json:"birthday,omitempty"`
	Email                      *string         `json:"email,omitempty"`
	Address1                   *string         `json:"address1,omitempty"`
	Address2                   *string         `json:"address2,omitempty"`
	Address3                   *string         `json:"address3,omitempty"`
	Postcode                   *string         `json:"postcode,omitempty"`
	City                       *string         `json:"city,omitempty"`
	State                      *string         `json:"state,omitempty"`
	Country                    *string         `json:"country,omitempty"`
	CountryName                *string         `json:"countryName,omitempty"`
	Phone                      *string         `json:"phone,omitempty"`
	Mobile                     *string         `json:"mobile,omitempty"`
	Nationality                *string         `json:"nationality,omitempty"`
	NationalityOther           *string         `json:"nationalityOther,omitempty"`
	PlaceOfBirth               *string         `json:"placeOfBirth,omitempty"`
	BirthCountry               *string         `json:"birthCountry,omitempty"`
	Occupation                 *string         `json:"occupation,omitempty"`
	Position                   *string         `json:"position,omitempty"`
	IncomeRange                *string         `json:"incomeRange,omitempty"`
	PersonalAssets             *string         `json:"personalAssets,omitempty"`
	LegalName                  *string         `json:"legalName,omitempty"`
	LegalNameEmbossed          *string         `json:"legalNameEmbossed,omitempty"`
	LegalRegistrationNumber    *string         `json:"legalRegistrationNumber,omitempty"`
	LegalTvaNumber             *string         `json:"legalTvaNumber,omitempty"`
	LegalRegistrationDate      *Date           `json:"legalRegistrationDate,omitempty"`
	LegalForm                  *string         `json:"legalForm,omitempty"`
	LegalShareCapital          *string         `json:"legalShareCapital,omitempty"`
	LegalSector                *string         `json:"legalSector,omitempty"`
	LegalAnnualTurnOver        *string         `json:"legalAnnualTurnOver,omitempty"`
	LegalNetIncomeRange        *string         `json:"legalNetIncomeRange,omitempty"`
	LegalNumberOfEmployeeRange *string         `json:"legalNumberOfEmployeeRange,omitempty"`
	EffectiveBeneficiary       *string         `json:"effectiveBeneficiary,omitempty"`
	KycLevel                   *Level          `json:"kycLevel,string,omitempty"`
	KycReview                  *Review         `json:"kycReview,string,omitempty"`
	KycReviewComment           *string         `json:"kycReviewComment,omitempty"`
	IsFreezed                  *int64          `json:"isFreezed,string,omitempty"`
	Language                   *string         `json:"language,omitempty"`
	SepaCreditorIdentifier     *string         `json:"sepaCreditorIdentifier,omitempty"`
	CreatedDate                *TimestampParis `json:"createdDate,omitempty"`
	ModifiedDate               *TimestampParis `json:"modifiedDate,omitempty"`
	CodeStatus                 *string         `json:"codeStatus,omitempty"`
	TaxNumber                  *string         `json:"taxNumber,omitempty"`
	TaxResidence               *string         `json:"taxResidence,omitempty"`
	ActivityOutsideEu          *string         `json:"activityOutsideEu,omitempty"`
	EconomicSanctions          *string         `json:"economicSanctions,omitempty"`
	ResidentCountriesSanctions *string         `json:"residentCountriesSanctions,omitempty"`
	InvolvedSanctions          *string         `json:"involvedSanctions,omitempty"`
	SanctionsQuestionnaireDate *string         `json:"sanctionsQuestionnaireDate,omitempty"`
	InformationStatus          *string         `json:"informationStatus,omitempty"`
	EntityType                 *string         `json:"entityType,omitempty"`
	WalletCount                *int64          `json:"walletCount,string,omitempty"`
	PayinCount                 *int64          `json:"payinCount,string,omitempty"`
	TotalRows                  *int64          `json:"totalRows,string,omitempty"`
}

type Level int32

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

	return ur, resp, errors.WithStack(err)
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

// ReviewKYCLiveness asks Treezor to do a KYC review against that user.
func (s *UserService) ReviewKYCLiveness(ctx context.Context, treezorUserID string) (*http.Response, error) {
	u := fmt.Sprintf("users/%s/kycliveness", treezorUserID)
	req, _ := s.client.NewRequestWithoutIndex(http.MethodPut, u, nil)

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
	req, _ := s.client.NewRequestWithoutIndex(http.MethodPost, u, nil)

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
