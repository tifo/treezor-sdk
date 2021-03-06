package treezor

import (
	"time"
)

// TimestampParis Paris time zone
type TimestampParis timestamp

var (
	parisLocation *time.Location
)

func init() {
	var err error
	parisLocation, err = time.LoadLocation("Europe/Paris")
	if err != nil {
		panic(err)
	}
}

// NewTimestampParis create a timestamp with the right location for testing
func NewTimestampParis(t time.Time) *TimestampParis {
	return &TimestampParis{
		Time: t.In(parisLocation),
	}
}

// MarshalJSON implements the json.Marshaler interface.
// Time is return as "YYYY-MM-DD HH:mm:ss" or "YYYY-MM-DD"
// depending whether Timestamp.Short is true.
// time zone Europe/Paris
func (t *TimestampParis) MarshalJSON() ([]byte, error) {
	t.Time = t.In(parisLocation)

	return (*timestamp)(t).MarshalJSON()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// Time is expected in "YYYY-MM-DD HH:mm:ss" or "YYYY-MM-DD".
// Time zone Europe/Paris
func (t *TimestampParis) UnmarshalJSON(data []byte) error {
	t.Time = t.In(parisLocation)

	return (*timestamp)(t).UnmarshalJSON(data)
}
