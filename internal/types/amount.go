package types

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
)

type Amount string

func (a *Amount) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte(`""`)) {
		*a = Amount("0.0")
		return nil
	}
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return errors.Wrap(err, "treezor.Amount")
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
