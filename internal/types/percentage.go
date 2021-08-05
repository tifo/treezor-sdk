package types

import (
	"bytes"

	"github.com/pkg/errors"

	json "github.com/tifo/treezor-sdk/internal/json"
)

type Percentage float64

func (p Percentage) Float64() float64 {
	return float64(p)
}

func (p *Percentage) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte(`""`)) {
		*p = Percentage(0.0)
		return nil
	}
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return errors.Wrap(err, "treezor.Percentage")
	}
	f, err := str.Float64()
	if err != nil {
		return errors.Wrap(err, "treezor.Percentage")
	}
	*p = Percentage(f)
	return nil
}
