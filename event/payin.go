package event

import (
	"github.com/tifo/treezor-sdk"
)

type PayinEvent struct {
	Payins []*treezor.Payin `json:"payins"`
}
