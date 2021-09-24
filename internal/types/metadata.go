package types

import (
	"bytes"

	json "github.com/tifo/treezor-sdk/internal/json"
)

type Metadata map[string]string

func (m Metadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string(m))
}

func (m *Metadata) UnmarshalJSON(data []byte) error {
	meta := make(map[string]string)
	if bytes.Equal(data, []byte("{}")) || bytes.Equal(data, []byte("[]")) {
		*m = meta
		return nil
	}
	err := json.Unmarshal(data, &meta)
	if err != nil {
		return err
	}
	*m = meta
	return nil
}
