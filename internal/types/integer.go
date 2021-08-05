package types

import (
	"bytes"

	"github.com/pkg/errors"

	json "github.com/tifo/treezor-sdk/internal/json"
)

type Integer int64

func (i Integer) Int64() int64 {
	return int64(i)
}

func (i *Integer) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte(`""`)) {
		*i = Integer(0)
		return nil
	}
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return errors.Wrap(err, "treezor.Integer")
	}
	v, err := str.Int64()
	if err != nil {
		return errors.Wrap(err, "treezor.Integer")
	}
	*i = Integer(v)
	return nil
}
