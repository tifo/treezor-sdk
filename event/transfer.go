package event

import "github.com/tifo/treezor-sdk"

type TransferEvent struct {
	Transfers []*treezor.Transfer `json:"transfers"`
}
