package types

import (
	"encoding/json"
	"strconv"
)

type Amount string

func (a *Amount) UnmarshalJSON(data []byte) error {
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	*a = Amount(str)
	return nil
}

func (a Amount) Float64() (float64, error) {
	return strconv.ParseFloat(string(a), 64)
}

// NOTE: an idea could be to alias "pointer type" and manage teh access directly
