package types

import (
	"encoding/json"
)

type Percentage float64

func (p Percentage) Float64() float64 {
	return float64(p)
}

func (p *Percentage) UnmarshalJSON(data []byte) error {
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	f, err := str.Float64()
	if err != nil {
		return err
	}
	*p = Percentage(f)
	return nil
}
