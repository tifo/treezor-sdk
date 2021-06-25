package event

import "github.com/tifo/treezor-sdk"

// WalletEvent represents a wallet event
type WalletEvent struct {
	Wallets []*treezor.Wallet `json:"wallets"`
}
