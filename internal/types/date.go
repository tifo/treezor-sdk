package types

import (
	"time"
)

const (
	shortFormat = "2006-01-02"
)

// Date represents a date that can be unmarshalled from a JSON string
// formatted as "YYYY-MM-DD".
// This is necessary for some fields since the Treezor API is inconsistent
// in how it represents dates.
type Date struct {
	time.Time
	OriginalPayload string
}

// NewDate create a date with the right format for testing
func NewDate(t time.Time) *Date {
	return &Date{
		Time: time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC),
	}
}

// String implements string default behavior
func (d Date) String() string {
	if d.Time.IsZero() {
		return "0000-00-00"
	}
	return d.Time.Format(shortFormat)
}

// MarshalJSON implements the json.Marshaler interface.
// Time is return as "YYYY-MM-DD"
func (d *Date) MarshalJSON() ([]byte, error) {
	var str string

	if d.Time.IsZero() {
		return []byte(`"0000-00-00"`), nil
	}

	str = d.Format(shortFormat)

	return []byte(`"` + str + `"`), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// Time is expected in "YYYY-MM-DD".
func (d *Date) UnmarshalJSON(data []byte) error {
	str := string(data)
	d.OriginalPayload = str

	if str == `"0000-00-00"` {
		return nil
	}

	tt, err := time.ParseInLocation(`"`+shortFormat+`"`, str, time.UTC)
	if err != nil {
		return err
	}

	d.Time = tt

	return nil
}

// NOTE: we might need to allow dates to include a time just in case we failed in setting up the type or the date includes 00:00:00
