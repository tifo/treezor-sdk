package treezor

import "time"

// TimestampLondon london time zone
type TimestampLondon timestamp

var (
	londonLocation *time.Location
)

func init() {
	var err error
	londonLocation, err = time.LoadLocation("Europe/London")
	if err != nil {
		panic(err)
	}
}

// NewTimestampLondon create a timestamp with the right location for testing
func NewTimestampLondon(t time.Time) *TimestampLondon {
	return &TimestampLondon{
		Time: t.In(londonLocation),
	}
}

// MarshalJSON implements the json.Marshaler interface.
// Time is return as "YYYY-MM-DD HH:mm:ss" or "YYYY-MM-DD"
// depending whether Timestamp.Short is true.
// time zone Europe/Paris
func (t *TimestampLondon) MarshalJSON() ([]byte, error) {
	t.Time = t.In(londonLocation)
	// t.Location = londonLocation

	return (*timestamp)(t).MarshalJSON()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// Time is expected in "YYYY-MM-DD HH:mm:ss" or "YYYY-MM-DD".
// Time zone Europe/Paris
func (t *TimestampLondon) UnmarshalJSON(data []byte) error {
	t.Time = t.In(londonLocation)
	// t.Location = londonLocation

	return (*timestamp)(t).UnmarshalJSON(data)
}
