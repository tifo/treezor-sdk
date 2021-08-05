package json

import (
	"encoding/json"

	jsoniter "github.com/json-iterator/go"
)

var (
	JSON       = jsoniter.ConfigCompatibleWithStandardLibrary
	Marshal    = JSON.Marshal
	Unmarshal  = JSON.Unmarshal
	NewDecoder = JSON.NewDecoder
	NewEncoder = JSON.NewEncoder
)

type (
	Number             = jsoniter.Number
	RawMessage         = jsoniter.RawMessage
	UnmarshalTypeError = json.UnmarshalTypeError
)

func init() {
	JSON.RegisterExtension(&TimeExtension{})
}
