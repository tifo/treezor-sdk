package treezor

import (
	"time"

	"github.com/tifo/treezor-sdk/internal/types"
)

// Float64 is a helper routine that allocates a new float64 value to store v and returns a pointer to it.
func Float64(v float64) *float64 { return &v }

// Int64 is a helper routine that allocates a new int64 value to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value to store v and returns a pointer to it.
func String(v string) *string { return &v }

// Time is a helper routine that allocates a new Time value to store v and returns a pointer to it.
func Time(v time.Time) *time.Time { return &v }

// Bool is a helper routine that allocates a new bool value to store v and returns a pointer to it.
func Bool(v bool) *types.Boolean {
	d := types.Boolean(v)
	return &d
}

// Date is a helper routine that allocates a new Date value to store v and returns a pointer to it.
func Date(v time.Time) *types.Date {
	if v.IsZero() {
		return nil
	}
	return types.NewDate(v)
}

// NewDate creates a new Date value from a year, month, day set.
func NewDate(year int, month time.Month, day int) *types.Date {
	return &types.Date{
		Time: time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
	}
}

// Timestamp is a helper routine that allocates a new Timestamp value to store v and returns a pointer to it.
func Timestamp(v time.Time) *types.Timestamp {
	if v.IsZero() {
		return nil
	}
	return &types.Timestamp{
		Time: v,
	}
}
