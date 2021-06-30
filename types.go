package treezor

import (
	"time"

	"github.com/tifo/treezor-sdk/internal/types"
)

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *types.Boolean {
	d := types.Boolean(v)
	return &d
}

// Float64 is a helper routine that allocates a new float64 value
// to store v and returns a pointer to it.
func Float64(v float64) *float64 { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// Date is a helper routine that allocates a new Date value
// to store v and returns a pointer to it.
func Date(date time.Time) *types.Date {
	return types.NewDate(date)
}

// Timestamp is a helper routine that allocates a new Timestamp value
// to store v and returns a pointer to it.
func Timestamp(v time.Time) *types.Timestamp {
	return &types.Timestamp{
		Time: v,
	}
}
