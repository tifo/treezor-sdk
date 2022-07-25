package xtime

import "time"

const (
	Treezor     = "Treezor"
	ANSIC       = "ANSIC"
	UnixDate    = "UnixDate"
	RubyDate    = "RubyDate"
	RFC822      = "RFC822"
	RFC822Z     = "RFC822Z"
	RFC850      = "RFC850"
	RFC1123     = "RFC1123"
	RFC1123Z    = "RFC1123Z"
	RFC3339     = "RFC3339"
	RFC3339Nano = "RFC3339Nano"
	Kitchen     = "Kitchen"
	Stamp       = "Stamp"
	StampMilli  = "StampMilli"
	StampMicro  = "StampMicro"
	StampNano   = "StampNano"
)

var timeFormats = map[string]string{
	Treezor:     "2006-01-02 15:04:05",
	ANSIC:       time.ANSIC,
	UnixDate:    time.UnixDate,
	RubyDate:    time.RubyDate,
	RFC822:      time.RFC822,
	RFC822Z:     time.RFC822Z,
	RFC850:      time.RFC850,
	RFC1123:     time.RFC1123,
	RFC1123Z:    time.RFC1123Z,
	RFC3339:     time.RFC3339,
	RFC3339Nano: time.RFC3339Nano,
	Kitchen:     time.Kitchen,
	Stamp:       time.Stamp,
	StampMilli:  time.StampMilli,
	StampMicro:  time.StampMicro,
	StampNano:   time.StampNano,
}

const (
	Local  = "Local"
	Paris  = "Europe/Paris"
	London = "Europe/London"
	UTC    = "UTC"
)

var timeLocations = map[string]*time.Location{
	Local:  time.Local,
	Paris:  mustLoadLocation("Europe/Paris"),
	London: mustLoadLocation("Europe/London"),
	UTC:    time.UTC,
}

const (
	DefaultFormat     = time.RFC3339
	StructTagFormat   = "layout"
	StructTagLocation = "loc"
)

func GetLocation(loc string) (*time.Location, bool) {
	l, ok := timeLocations[loc]
	return l, ok
}

func GetFormat(fmt string) string {
	f, ok := timeFormats[fmt]
	if !ok {
		return fmt
	}
	return f
}

func mustLoadLocation(loc string) *time.Location {
	l, err := time.LoadLocation(loc)
	if err != nil {
		panic(err)
	}
	return l
}
