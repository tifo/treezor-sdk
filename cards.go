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
	StatusCode                 *string                       `json:"statusCode,omitempty"`
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
	Physical                   *types.Boolean                `json:"physical,string,omitempty"`
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

type CreateVirtualCardOptions struct {
	// Card.CreateVirtual / Card.RequestPhysical only data
	CardPrint *string `json:"cardPrint,omitempty"`
	PIN       *string `json:"pin,omitempty"`
	// NOTE: might need to be stored in another struct as you don't need to send the whole card model when update a card status
	LockStatus *int64 `json:"lockStatus,omitempty"`
	// TODO: CreateCardRequest has much more field than cards and might require its own type (missing fields such as Anonymous, ...)
}

// CreateVirtual will create a virtual card.
func (s *CardService) CreateVirtual(ctx context.Context, card *Card) (*Card, *http.Response, error) {
	req, _ := s.client.NewRequest(http.MethodPost, "cards/CreateVirtual", card)

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

// RequestPhysical will request a physical card that will be sent to the user's address.
func (s *CardService) RequestPhysical(ctx context.Context, card *Card) (*Card, *http.Response, error) {
	req, _ := s.client.NewRequest(http.MethodPost, "cards/RequestPhysical", card)

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
	CardID string `url:"cardId,omitempty"`
}

// CardImagesResponse contains a list of virtual card images.
type CardImagesResponse struct {
	CardImages []*CardImage `json:"cardimages"`
}

// CardImage represents a virtual card image.
type CardImage struct {
	ID     *string `json:"id,omitempty"`
	CardID *string `json:"cardId,omitempty"`
	File   *string `json:"file,omitempty"`
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

// Get returns a card (virtual or physical).
func (s *CardService) Get(ctx context.Context, cardID string) (*Card, *http.Response, error) {
	u := fmt.Sprintf("cards/%s", cardID)
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
)

type CardLockUnlockOptions struct {
	LockStatus *int64 `json:"lockStatus,omitempty"`
}

// LockUnlock toggle the lock or unlock state of a card. If the card is locked, calling this function
// will unlock the card, and vice versa.
func (s *CardService) LockUnlock(ctx context.Context, cardID string, lockStatus LockStatus) (*Card, *http.Response, error) {
	u := fmt.Sprintf("cards/%s/LockUnlock/", cardID)
	req, _ := s.client.NewRequest(http.MethodPut, u, &CardLockUnlockOptions{
		LockStatus: Int64(int64(lockStatus)),
	})

	c := new(CardResponse)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	if len(c.Cards) < 1 {
		return nil, resp, errors.New("API returned no card")
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
	LimitATMYear      int64 `json:"limitAtmYear,omitempty"`
	LimitATMMonth     int64 `json:"limitAtmMonth,omitempty"`
	LimitATMWeek      int64 `json:"limitAtmWeek,omitempty"`
	LimitATMDay       int64 `json:"limitAtmDay,omitempty"`
	LimitATMAll       int64 `json:"limitAtmAll,omitempty"`
	LimitPaymentYear  int64 `json:"limitPaymentYear,omitempty"`
	LimitPaymentMonth int64 `json:"limitPaymentMonth,omitempty"`
	LimitPaymentWeek  int64 `json:"limitPaymentWeek,omitempty"`
	LimitPaymentDay   int64 `json:"limitPaymentDay,omitempty"`
	LimitPaymentAll   int64 `json:"limitPaymentAng,omitempty"`
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

// PIN is used to make PIN modification operations.
type PIN struct {
	Current      string `json:"currentPIN,omitempty"`
	New          string `json:"newPIN,omitempty"`
	Confirmation string `json:"confirmPIN,omitempty"`
}

// ChangePIN changes the card PIN. It needs the current PIN, the new one and a confirmation one.
func (s *CardService) ChangePIN(ctx context.Context, cardID string, pin *PIN) (*Card, *http.Response, error) {
	u := fmt.Sprintf("cards/%s/ChangePIN/", cardID)
	req, _ := s.client.NewRequest(http.MethodPut, u, pin)

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

// SetPIN sets the card PIN. It needs the the new PIN and a confirmation one. It is solely used by operators,
// not users.
func (s *CardService) SetPIN(ctx context.Context, cardID string, pin *PIN) (*Card, *http.Response, error) {
	u := fmt.Sprintf("cards/%s/setPIN/", cardID)
	req, _ := s.client.NewRequest(http.MethodPut, u, pin)

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

// Deactivate deactivates a card permanently.
func (s *CardService) Deactivate(ctx context.Context, cardID string) (*Card, *http.Response, error) {
	u := fmt.Sprintf("cards/%s", cardID)
	req, _ := s.client.NewRequest(http.MethodDelete, u, nil)

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
type Card3DS struct {
	CardID *string `json:"cardId,omitempty"`
}

// Register3DSecure will register a card to 3DSecure
func (s *CardService) Register3DSecure(ctx context.Context, cardID *Card3DS) (*Card, *http.Response, error) {
	card := &Card{}
	req, _ := s.client.NewRequest(http.MethodPost, "cards/Register3DS", cardID)

	c := new(CardResponse)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, errors.WithStack(err)
	}

	// TODO: Make sure the response is actually a single card or an empty array
	if len(c.GetCards()) > 0 {
		card = c.GetCards()[len(c.GetCards())-1]
	}
	/*if len(c.Cards) != 1 {
		return nil, resp, errors.Errorf("API did not returned exactly one card: %d cards returned", len(c.Cards))
	}*/
	return card, resp, nil
}

// TODO: Update Card API
