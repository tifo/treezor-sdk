package treezor

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/tifo/treezor-sdk/internal/types"
)

type CardPermissionMask int8

// Available permissions for a card.
const (
	Foreign CardPermissionMask = 1 << iota
	Online
	ATM
	NFC
)

const (
	Noop CardPermissionMask = 0
	All  CardPermissionMask = Foreign | Online | ATM | NFC
)

// ConvertPermissions map binary field of card permission to
// an internal value at Treezor which groups those permissions.
//
// e.g.: ConvertPermissions(ATM|Foreign) returns TRZ-CU-006.
//       ConvertPermissions(All) returns TRZ-CU-016.
func ConvertPermissions(permissions CardPermissionMask) string {
	if permissions > All {
		return "TRZ-CU-016"
	}
	return fmt.Sprintf("TRZ-CU-%03d", permissions+1)
}

// CardService handles communication with the card related
// methods of the Treezor API.
//
// Treezor API docs: https://www.treezor.com/api-documentation/#/card
type CardService service

// CardResponse represents a list of cards.
// It may contain only one item.
type CardResponse struct {
	Cards []*Card `json:"cards"`
}

type CardStatus string

const (
	CardStatusUnlock CardStatus = "UNLOCK"
	CardStatusLock   CardStatus = "LOCK"
	CardStatusLost   CardStatus = "LOST"
	CardStatusStolen CardStatus = "STOLEN"
)

// Card represents a physical or virtual card.
type Card struct {
	CardID                     *types.Identifier             `json:"cardId,omitempty"`
	UserID                     *types.Identifier             `json:"userId,omitempty"`
	WalletID                   *types.Identifier             `json:"walletId,omitempty"`
	WalletCardtransactionID    *types.Identifier             `json:"walletCardtransactionId,omitempty"`
	MccRestrictionGroupID      *types.Identifier             `json:"mccRestrictionGroupId,omitempty"`
	MerchantRestrictionGroupID *types.Identifier             `json:"merchantRestrictionGroupId,omitempty"`
	CountryRestrictionGroupID  *types.Identifier             `json:"countryRestrictionGroupID,omitempty"`
	EventName                  *string                       `json:"eventName,omitempty"`
	EventAlias                 *string                       `json:"eventAlias,omitempty"`
	PublicToken                *string                       `json:"publicToken,omitempty"`
	CardTag                    *string                       `json:"cardTag,omitempty"`
	StatusCode                 *CardStatus                   `json:"statusCode,omitempty"`
	IsLive                     *types.Boolean                `json:"isLive,omitempty"`
	PINTryExceeds              *types.Boolean                `json:"pinTryExceeds,omitempty"`
	MaskedPan                  *string                       `json:"maskedPan,omitempty"`
	EmbossedName               *string                       `json:"embossedName,omitempty"`
	ExpiryDate                 *types.Date                   `json:"expiryDate,omitempty"`
	CVV                        *string                       `json:"CVV,omitempty"`
	StartDate                  *types.Date                   `json:"startDate,omitempty"`
	EndDate                    *types.Date                   `json:"endDate,omitempty"`
	CountryCode                *string                       `json:"countryCode,omitempty"`
	CurrencyCode               *Currency                     `json:"currencyCode,omitempty"`
	Lang                       *string                       `json:"lang,omitempty"`
	DeliveryTitle              *string                       `json:"deliveryTitle,omitempty"`
	DeliveryLastname           *string                       `json:"deliveryLastname,omitempty"`
	DeliveryFirstname          *string                       `json:"deliveryFirstname,omitempty"`
	DeliveryAddress1           *string                       `json:"deliveryAddress1,omitempty"`
	DeliveryAddress2           *string                       `json:"deliveryAddress2,omitempty"`
	DeliveryAddress3           *string                       `json:"deliveryAddress3,omitempty"`
	DeliveryCity               *string                       `json:"deliveryCity,omitempty"`
	DeliveryPostcode           *string                       `json:"deliveryPostcode,omitempty"`
	DeliveryCountry            *string                       `json:"deliveryCountry,omitempty"`
	MobileSent                 *string                       `json:"mobileSent,omitempty"`
	LimitsGroup                *string                       `json:"limitsGroup,omitempty"`
	PermsGroup                 *string                       `json:"permsGroup,omitempty"` // NOTE: could be a custom type using CardPermissionMask
	CardDesign                 *string                       `json:"cardDesign,omitempty"`
	VirtualConverted           *types.Boolean                `json:"virtualConverted,omitempty"`
	Physical                   *types.Boolean                `json:"physical,omitempty"`
	OptionATM                  *types.Boolean                `json:"optionAtm,omitempty"`
	OptionForeign              *types.Boolean                `json:"optionForeign,omitempty"`
	OptionOnline               *types.Boolean                `json:"optionOnline,omitempty"`
	OptionNFC                  *types.Boolean                `json:"optionNfc,omitempty"`
	LimitATMYear               *types.Integer                `json:"limitAtmYear,omitempty"`
	LimitATMMonth              *types.Integer                `json:"limitAtmMonth,omitempty"`
	LimitATMWeek               *types.Integer                `json:"limitAtmWeek,omitempty"`
	LimitATMDay                *types.Integer                `json:"limitAtmDay,omitempty"`
	LimitATMAll                *types.Integer                `json:"limitAtmAll,omitempty"`
	LimitPaymentYear           *types.Integer                `json:"limitPaymentYear,omitempty"`
	LimitPaymentMonth          *types.Integer                `json:"limitPaymentMonth,omitempty"`
	LimitPaymentWeek           *types.Integer                `json:"limitPaymentWeek,omitempty"`
	LimitPaymentDay            *types.Integer                `json:"limitPaymentDay,omitempty"`
	LimitPaymentAll            *types.Integer                `json:"limitPaymentAll,omitempty"`
	PaymentDailyLimit          *types.Amount                 `json:"paymentDailyLimit,omitempty"`
	RestrictionGroupLimits     []*CardRestrictionGroupLimits `json:"restrictionGroupLimits,omitempty"`
	TotalATMYear               *types.Amount                 `json:"totalAtmYear,omitempty"`
	TotalATMMonth              *types.Amount                 `json:"totalAtmMonth,omitempty"`
	TotalATMWeek               *types.Amount                 `json:"totalAtmWeek,omitempty"`
	TotalATMDay                *types.Amount                 `json:"totalAtmDay,omitempty"`
	TotalATMAll                *types.Amount                 `json:"totalAtmAll,omitempty"`
	TotalPaymentYear           *types.Amount                 `json:"totalPaymentYear,omitempty"`
	TotalPaymentMonth          *types.Amount                 `json:"totalPaymentMonth,omitempty"`
	TotalPaymentWeek           *types.Amount                 `json:"totalPaymentWeek,omitempty"`
	TotalPaymentDay            *types.Amount                 `json:"totalPaymentDay,omitempty"`
	TotalPaymentAll            *types.Amount                 `json:"totalPaymentAll,omitempty"`
	CreatedBy                  *types.Identifier             `json:"createdBy,omitempty"`
	CreatedDate                *time.Time                    `json:"createdDate,omitempty" layout:"Treezor" loc:"Europe/London"`
	ModifiedBy                 *types.Identifier             `json:"modifiedBy,omitempty"`
	ModifiedDate               *time.Time                    `json:"modifiedDate,omitempty" layout:"Treezor" loc:"Europe/London"`
	CancellationNumber         *types.Integer                `json:"cancellationNumber,omitempty"`
	TotalRows                  *types.Integer                `json:"totalRows,omitempty"`
}

type CardRestrictionGroupLimits struct {
	PaymentDailyLimit           *types.Amount     `json:"paymentDailyLimit,omitempty"`
	MccRestrictionGroups        *types.Identifier `json:"mccRestrictionGroups,omitempty"`        // NOTE: not sure if its an identifier or a random integer
	CountryRestrictionGroups    *types.Identifier `json:"countryRestrictionGroups,omitempty"`    // NOTE: not sure if its an identifier or a random integer
	MerchantIdRestrictionGroups *types.Identifier `json:"merchantIdRestrictionGroups,omitempty"` // NOTE: not sure if its an identifier or a random integer
}

type CardRestrictionGroups struct {
	MCCRestrictionGroupID      *string `url:"-" json:"mccRestrictionGroupId,omitempty"`
	MerchantRestrictionGroupID *string `url:"-" json:"merchantRestrictionGroupId,omitempty"`
	CountryRestrictionGroupID  *string `url:"-" json:"countryRestrictionGroupId,omitempty"`
}

type CreateVirtualCardOptions struct {
	Access

	UserID       *string        `url:"-" json:"userId"`     // Required
	WalletID     *string        `url:"-" json:"walletId"`   // Required
	PermsGroup   *string        `url:"-" json:"permsGroup"` // Required
	CardPrint    *string        `url:"-" json:"cardPrint"`  // Required
	CardTag      *string        `url:"-" json:"cardTag,omitempty"`
	PIN          *string        `url:"-" json:"pin,omitempty"`
	Anonymous    *types.Boolean `url:"-" json:"anonymous,omitempty"`
	SendToParent *types.Boolean `url:"-" json:"sendToParent,omitempty"`

	CardLimits
	CardRestrictionGroups

	EmbossLegalName *types.Boolean `url:"-" json:"embossLegalName,omitempty"`
	LogoID          *string        `url:"-" json:"logoId,omitempty"`
	DesignCode      *string        `url:"-" json:"designCode,omitempty"`
	PackageID       *string        `url:"-" json:"packageId,omitempty"`
	CustomizedInfo  *string        `url:"-" json:"customizedInfo,omitempty"`
	CardLanguages   *string        `url:"-" json:"cardLanguages,omitempty"`
}

// CreateVirtual will create a virtual card.
func (s *CardService) CreateVirtual(ctx context.Context, opts *CreateVirtualCardOptions) (*Card, *http.Response, error) {
	u := "cards/CreateVirtual"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPost, u, opts)
	c := new(CardResponse)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(c.Cards) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one card: %d cards returned", len(c.Cards))
	}
	return c.Cards[0], resp, nil
}

type RequestPhysicalCardOptions struct {
	Access

	UserID       *string        `url:"-" json:"userId"`     // Required
	WalletID     *string        `url:"-" json:"walletId"`   // Required
	PermsGroup   *string        `url:"-" json:"permsGroup"` // Required
	CardPrint    *string        `url:"-" json:"cardPrint"`  // Required
	CardTag      *string        `url:"-" json:"cardTag,omitempty"`
	PIN          *string        `url:"-" json:"pin,omitempty"`
	Anonymous    *types.Boolean `url:"-" json:"anonymous,omitempty"`
	SendToParent *types.Boolean `url:"-" json:"sendToParent,omitempty"`

	CardLimits
	CardRestrictionGroups

	EmbossLegalName *types.Boolean `url:"-" json:"embossLegalName,omitempty"`
	LogoID          *string        `url:"-" json:"logoId,omitempty"`
	DesignCode      *string        `url:"-" json:"designCode,omitempty"`
	PackageID       *string        `url:"-" json:"packageId,omitempty"`
	CustomizedInfo  *string        `url:"-" json:"customizedInfo,omitempty"`
	CardLanguages   *string        `url:"-" json:"cardLanguages,omitempty"`
}

// RequestPhysical will request a physical card that will be sent to the user's address.
func (s *CardService) RequestPhysical(ctx context.Context, opts *RequestPhysicalCardOptions) (*Card, *http.Response, error) {
	u := "cards/RequestPhysical"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPost, u, opts)

	c := new(CardResponse)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(c.Cards) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one card: %d cards returned", len(c.Cards))
	}
	return c.Cards[0], resp, nil
}

// CardGetImagesOptions contains options when getting a card image.
type CardGetImagesOptions struct {
	Access
	CardID              *string `url:"cardId,omitempty" json:"-"` // Required
	EncryptionMethod    *string `url:"encryptionMethod,omitempty" json:"-"`
	EncryptionPublicKey *string `url:"encryptionPublicKey,omitempty" json:"-"`
}

// CardImagesResponse contains a list of virtual card images.
type CardImagesResponse struct {
	CardImages []*CardImage `json:"cardimages"`
}

// CardImage represents a virtual card image.
type CardImage struct {
	ID     *string `json:"id"` // Required
	CardID *string `json:"cardId,omitempty"`
	File   []byte  `json:"file,omitempty"`
}

// GetImage returns the provided virtual card image.
func (s *CardService) GetImage(ctx context.Context, opt *CardGetImagesOptions) (*CardImage, *http.Response, error) {
	u := "cardimages"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	c := new(CardImagesResponse)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(c.CardImages) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one card: %d cards returned", len(c.CardImages))
	}

	return c.CardImages[0], resp, errors.WithStack(err)
}

type CardGetOptions struct {
}

// Get returns a card (virtual or physical).
func (s *CardService) Get(ctx context.Context, cardID string, opts *CardGetOptions) (*Card, *http.Response, error) {
	u := fmt.Sprintf("cards/%s", cardID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)
	c := new(CardResponse)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(c.Cards) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one card: %d cards returned", len(c.Cards))
	}
	return c.Cards[0], resp, nil
}

// CardListOptions contains URL options for listing cards.
type CardListOptions struct {
	Access

	CardID                     *string        `url:"cardId,omitempty" json:"-"`
	UserID                     *string        `url:"userId,omitempty" json:"-"`
	EmbossedName               *string        `url:"embossedName,omitempty" json:"-"`
	PublicToken                *string        `url:"publicToken,omitempty" json:"-"`
	PermsGroup                 *string        `url:"permsGroup,omitempty" json:"-"`
	IsPhysical                 *types.Boolean `url:"isPhysical,omitempty" json:"-"`
	IsConverted                *types.Boolean `url:"isConverted,omitempty" json:"-"`
	LockStatus                 *LockStatus    `url:"lockStatus,omitempty" json:"-"`
	MCCRestrictionGroupID      *string        `url:"mccRestrictionGroupId,omitempty" json:"-"`
	MerchantRestrictionGroupID *string        `url:"merchantRestrictionGroupId,omitempty" json:"-"`
	CountryRestrictionGroupID  *string        `url:"countryRestrictionGroupId,omitempty" json:"-"`

	ListOptions
}

// List the cards for the authenticated user.
func (s *CardService) List(ctx context.Context, opt *CardListOptions) (*CardResponse, *http.Response, error) {
	u := "cards"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodGet, u, nil)

	c := new(CardResponse)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	return c, resp, errors.WithStack(err)
}

// Edit updates the referenced card (with cardID) in parameter.
func (s *CardService) Edit(ctx context.Context, cardID string, card *Card) (*Card, *http.Response, error) {
	u := fmt.Sprintf("cards/%s", cardID)
	req, _ := s.client.NewRequest(http.MethodPut, u, card)

	c := new(CardResponse)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(c.Cards) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one card: %d cards returned", len(c.Cards))
	}
	return c.Cards[0], resp, nil
}

// Activate enable a card to make payments. It needs to be done only once.
func (s *CardService) Activate(ctx context.Context, cardID string) (*Card, *http.Response, error) {
	u := fmt.Sprintf("cards/%s/Activate/", cardID)
	req, _ := s.client.NewRequest(http.MethodPut, u, nil)

	c := new(CardResponse)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(c.Cards) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one card: %d cards returned", len(c.Cards))
	}
	return c.Cards[0], resp, nil
}

// LockStatus is the card status.
type LockStatus int32

// All the possible card lock statuses.
const (
	Unlocked LockStatus = iota
	Locked
	Lost
	Stolen
	Destroyed
)

type CardLockUnlockOptions struct {
	Access
	LockStatus LockStatus `url:"-" json:"lockStatus"`
}

// LockUnlock updates a card status
func (s *CardService) LockUnlock(ctx context.Context, cardID string, opts *CardLockUnlockOptions) (*Card, *http.Response, error) {
	u := fmt.Sprintf("cards/%s/LockUnlock/", cardID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPut, u, opts)

	c := new(CardResponse)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(c.Cards) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one card: %d cards returned", len(c.Cards))
	}
	return c.Cards[0], resp, nil
}

// CardOptions contains a card options.
type CardOptions struct {
	Foreign int `json:"foreign"`
	Online  int `json:"online"`
	ATM     int `json:"atm"`
	NFC     int `json:"nfc"`
}

// ChangeOptions change a card' options with the provided options.
func (s *CardService) ChangeOptions(ctx context.Context, cardID string, options *CardOptions) (*Card, *http.Response, error) {
	u := fmt.Sprintf("cards/%s/Options/", cardID)
	req, _ := s.client.NewRequest(http.MethodPut, u, options)

	c := new(CardResponse)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(c.Cards) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one card: %d cards returned", len(c.Cards))
	}
	return c.Cards[0], resp, nil
}

// CardLimits contains a card limit.
type CardLimits struct {
	LimitATMYear      int64 `url:"-" json:"limitAtmYear,omitempty"`
	LimitATMMonth     int64 `url:"-" json:"limitAtmMonth,omitempty"`
	LimitATMWeek      int64 `url:"-" json:"limitAtmWeek,omitempty"`
	LimitATMDay       int64 `url:"-" json:"limitAtmDay,omitempty"`
	LimitATMAll       int64 `url:"-" json:"limitAtmAll,omitempty"`
	LimitPaymentYear  int64 `url:"-" json:"limitPaymentYear,omitempty"`
	LimitPaymentMonth int64 `url:"-" json:"limitPaymentMonth,omitempty"`
	LimitPaymentWeek  int64 `url:"-" json:"limitPaymentWeek,omitempty"`
	LimitPaymentDay   int64 `url:"-" json:"limitPaymentDay,omitempty"`
	LimitPaymentAll   int64 `url:"-" json:"limitPaymentAng,omitempty"`
}

// ChangeLimits change a card' limits with the provided limits.
func (s *CardService) ChangeLimits(ctx context.Context, cardID string, limits *CardLimits) (*Card, *http.Response, error) {
	u := fmt.Sprintf("cards/%s/Limits/", cardID)
	req, _ := s.client.NewRequest(http.MethodPut, u, limits)

	c := new(CardResponse)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(c.Cards) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one card: %d cards returned", len(c.Cards))
	}
	return c.Cards[0], resp, nil
}

// Regenerate will recreate or re-order the card given in parameter with the exact same configuration.
func (s *CardService) Regenerate(ctx context.Context, cardID string) (*Card, *http.Response, error) {
	u := fmt.Sprintf("cards/%s/Regenerate/", cardID)
	req, _ := s.client.NewRequest(http.MethodPut, u, nil)

	c := new(CardResponse)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(c.Cards) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one card: %d cards returned", len(c.Cards))
	}
	return c.Cards[0], resp, nil
}

// ConvertVirtual will convert a virtual card to a physical one.
func (s *CardService) ConvertVirtual(ctx context.Context, cardID string) (*Card, *http.Response, error) {
	u := fmt.Sprintf("cards/%s/ConvertVirtual/", cardID)
	req, _ := s.client.NewRequest(http.MethodPut, u, nil)

	c := new(CardResponse)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(c.Cards) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one card: %d cards returned", len(c.Cards))
	}
	return c.Cards[0], resp, nil
}

type ChangePINOptions struct {
	CurrentPIN string `url:"-" json:"currentPIN,omitempty"`
	NewPIN     string `url:"-" json:"newPIN,omitempty"`
	ConfirmPIN string `url:"-" json:"confirmPIN,omitempty"`
}

// ChangePIN changes the card PIN. It needs the current PIN, the new one and a confirmation one.
func (s *CardService) ChangePIN(ctx context.Context, cardID string, opts *ChangePINOptions) (*Card, *http.Response, error) {
	u := fmt.Sprintf("cards/%s/ChangePIN/", cardID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPut, u, opts)

	c := new(CardResponse)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(c.Cards) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one card: %d cards returned", len(c.Cards))
	}
	return c.Cards[0], resp, nil
}

type SetPINOptions struct {
	Access
	NewPIN     string `url:"-" json:"newPIN,omitempty"`
	ConfirmPIN string `url:"-" json:"confirmPIN,omitempty"`
}

// SetPIN sets the card PIN. It needs the the new PIN and a confirmation one. It is solely used by operators,
// not users.
func (s *CardService) SetPIN(ctx context.Context, cardID string, opts *SetPINOptions) (*Card, *http.Response, error) {
	u := fmt.Sprintf("cards/%s/setPIN/", cardID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPut, u, opts)

	c := new(CardResponse)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(c.Cards) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one card: %d cards returned", len(c.Cards))
	}
	return c.Cards[0], resp, nil
}

// UnblockPIN unlocks the card PIN if it was blocked because of 3 failed attempts.
func (s *CardService) UnblockPIN(ctx context.Context, cardID string) (*Card, *http.Response, error) {
	u := fmt.Sprintf("cards/%s/UnblockPIN/", cardID)
	req, _ := s.client.NewRequest(http.MethodPut, u, nil)

	c := new(CardResponse)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(c.Cards) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one card: %d cards returned", len(c.Cards))
	}
	return c.Cards[0], resp, nil
}

// Card3DS is used to make register a 3D Secure card.
type RegisterCard3DSOptions struct {
	CardID *string `url:"-" json:"cardId"`
}

// Register3DS will register a card to 3DS
func (s *CardService) Register3DS(ctx context.Context, opts *RegisterCard3DSOptions) (*Card, *http.Response, error) {
	u := "cards/Register3DS"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	req, _ := s.client.NewRequest(http.MethodPost, u, opts)

	c := new(CardResponse)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(c.Cards) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one card: %d cards returned", len(c.Cards))
	}
	return c.Cards[0], resp, nil
}
