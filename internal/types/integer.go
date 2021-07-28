package types

import (
	"bytes"
	"encoding/json"
)

type Integer int64

func (i Integer) Int64() int64 {
	return int64(i)
}

func (i *Integer) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte(`""`)) {
		*i = Integer(0)
	}
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	v, err := str.Int64()
	if err != nil {
		return err
	}
	*i = Integer(v)
	return nil
}
