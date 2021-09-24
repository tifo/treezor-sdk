package types

import (
	"github.com/pkg/errors"

	json "github.com/tifo/treezor-sdk/internal/json"
)

type Boolean bool

func (b Boolean) Bool() bool {
	return bool(b)
}

// NOTE: should we implement accesor this way so we dont have to update gen_accessor when updating the default value of a custom type ?
// func (b *Boolean) Bool() bool {
// 	if b == nil {
// 		return false
// 	}
// 	return bool(*b)
// }

func (b *Boolean) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `true`, `1`:
		*b = true
	case `false`, `0`:
		*b = false
	default:
		var str json.Number
		err := json.Unmarshal(data, &str)
		if err != nil {
			return errors.Wrap(err, "treezor.Bool")
		}
		v, err := str.Int64()
		if err != nil {
			return errors.Wrap(err, "treezor.Bool")
		}
		*b = v > 0
	}
	return nil
}

func (b Boolean) MarshalJSON() ([]byte, error) {
	if !b {
		return []byte(`0`), nil
	}
	return []byte(`1`), nil
}
