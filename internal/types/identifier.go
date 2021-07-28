package types

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"
)

type Identifier string

func (i Identifier) String() string {
	return string(i)
}

func (i *Identifier) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte(`""`)) {
		*i = Identifier("")
		return nil
	}
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return errors.Wrap(err, "treezor.Identifier")
	}
	*i = Identifier(str)
	return nil
}
