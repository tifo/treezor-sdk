package types

import (
	"time"
)

const (
	fullFormat = "2006-01-02 15:04:05"
)

// Timestamp represents a time that can be unmarshalled from a JSON string
// formatted as "YYYY-MM-DD HH:mm:ss".
// This is necessary for some fields since the Treezor API is inconsistent
// in how it represents times. All exported methods of time.Time can be called on Timestamp.
type timestamp struct {
	time.Time
	OriginalPayload string
}

// MarshalJSON implements the json.Marshaler interface.
// Time is return as "YYYY-MM-DD HH:mm:ss"
func (t *timestamp) MarshalJSON() ([]byte, error) {
	var str string

	if t.Time.IsZero() {
		return []byte(`"0000-00-00 00:00:00"`), nil
	}

	str = t.Format(fullFormat)

	return []byte(`"` + str + `"`), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// Time is expected in "YYYY-MM-DD HH:mm:ss".
func (t *timestamp) UnmarshalJSON(data []byte) error {
	str := string(data)
	t.OriginalPayload = str

	if str == `"0000-00-00 00:00:00"` {
		return nil
	}

	tt, err := time.ParseInLocation(`"`+fullFormat+`"`, str, t.Location())
	if err != nil {
		// NOTE: the error is ignored here as we sometime have weird negative dates coming in, instead we handle it as ZeroTime
		return nil
	}

	t.Time = tt

	return nil
}

// TODO(0rax): see if we want / need to move those types to a `types` subpackage and just creer helper functions like treezor.String for those types in the main package.
// TODO(0rax): see about creating custom enums for some of the TypeId's available and use those types in root and event pkg
