package treezor

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
		return err
	}

	t.Time = tt

	return nil
}
