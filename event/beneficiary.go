package event

import (
	"github.com/tifo/treezor-sdk"
)

type BeneficiaryEvent struct {
	Beneficiaries []*treezor.Beneficiary `json:"beneficiaries"`
}
