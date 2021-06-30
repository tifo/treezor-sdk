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

func (a Amount) Float64() float64 {
	v, err := strconv.ParseFloat(string(a), 64)
	if err != nil {
		return 0.0
	}
	return v
}
