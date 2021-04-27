package treezor

type CardRequestPhysicalEvent struct {
	CardResponse
}

type CardCreateVirtualEvent struct {
	CardResponse
}

type CardConvertVirtualEvent struct {
	CardResponse
}

type CardChangePINEvent struct {
	CardResponse
}

type CardActivateEvent struct {
	CardResponse
}

type CardRenewEvent struct {
	CardResponse
}

type CardRegenerateEvent struct {
	CardResponse
}

type CardUpdateEvent struct {
	CardResponse
}

type CardLimitsEvent struct {
	CardResponse
}

type CardOptionsEvent struct {
	CardResponse
}

type CardSetPINEvent struct {
	CardResponse
}

type CardUnblockPINEvent struct {
	CardResponse
}

type CardLockUnlockEvent struct {
	CardResponse
}

type CardTransactionCreateEvent struct {
	CardTransactionResponse
}

type PayinCreateEvent struct {
	PayinResponse
}

type PayinUpdateEvent struct {
	PayinResponse
}

type PayinCancelEvent struct {
	PayinResponse
}

type PayoutCreateEvent struct {
	PayoutResponse
}

type PayoutUpdateEvent struct {
	PayoutResponse
}

type PayoutCancelEvent struct {
	PayoutResponse
}

type TransferCreateEvent struct {
	TransferResponse
}

type TransferUpdateEvent struct {
	TransferResponse
}

type TransferCancelEvent struct {
	TransferResponse
}

type UserCreateEvent struct {
	UserResponse
}

type UserUpdateEvent struct {
	UserResponse
}

type UserCancelEvent struct {
	UserResponse
}

type UserKYCReviewEvent struct {
	UserResponse
}

type UserKYCRequestEvent struct {
	UserResponse
}

type WalletCreateEvent struct {
	WalletResponse
}

type WalletUpdateEvent struct {
	WalletResponse
}

type WalletCancelEvent struct {
	WalletResponse
}

type KycLivenessCreateEvent struct {
	*KycLiveness
}

type KycLivenessUpdateEvent struct {
	*KycLiveness
}

type SepaSddrCoreRejectEvent struct {
	SepaSddrResponse
}
