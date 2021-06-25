package types

import "encoding/json"

type Identifier string

func (i *Identifier) UnmarshalJSON(data []byte) error {
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	*i = Identifier(str)
	return nil
}
