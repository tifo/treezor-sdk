package types

import "encoding/json"

type Boolean bool

func (b *Boolean) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `true`:
		*b = true
	case `false`:
		*b = false
	default:
		var str json.Number
		err := json.Unmarshal(data, &str)
		if err != nil {
			return err
		}
		v, err := str.Int64()
		if err != nil {
			return err
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
