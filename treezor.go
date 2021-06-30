//go:generate go run scripts/gen_accessors.go -v

package treezor

import (
	"net/url"
	"reflect"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

type service struct {
	client *Client
}

type SortOrder string

const (
	SortASC  SortOrder = "ASC"
	SortDESC SortOrder = "DESC"
)

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	Page *int `url:"pageNumber,omitempty" json:"-"`
	// For paginated result sets, the number of results to include per page.
	PerPage *int `url:"pageCount,omitempty" json:"-"`
	// For paginated result sets, the resource element you want to sort the list with.
	SortBy *string `url:"sortBy,omitempty" json:"-"`
	// For paginated result sets, The order you want to sort the list.
	// Possible values: DESC and ASC.
	SortOrder *SortOrder `url:"sortOrder,omitempty" json:"-"`
	// For paginated result sets, the creation date from which you want to filter the request result.
	// Format: YYYY-MM-DD HH:MM:SS
	CreatedFrom *time.Time `layout:"2006-01-02 15:04:05" url:"createdDateFrom,omitempty" json:"-"` // NOTE: Should we use types.TimestampParis ?
	// For paginated result sets, the creation date up to which you want to filter the request result.
	// Format: YYYY-MM-DD HH:MM:SS
	CreatedTo *time.Time `layout:"2006-01-02 15:04:05" url:"createdDateFrom,omitempty" json:"-"` // NOTE: Should we use types.TimestampParis ?
	// For paginated result sets, the modification date from which you want to filter the request result.
	// Format: YYYY-MM-DD HH:MM:SS
	UpdatedFrom *time.Time `layout:"2006-01-02 15:04:05" url:"createdDateFrom,omitempty" json:"-"` // NOTE: Should we use types.TimestampParis ?
	// For paginated result sets, the modification date up to which you want to filter the request result.
	// Format: YYYY-MM-DD HH:MM:SS
	UpdatedTo *time.Time `layout:"2006-01-02 15:04:05" url:"createdDateFrom,omitempty" json:"-"` // NOTE: Should we use types.TimestampParis ?
}

// addOptions adds the parameters in opt as URL query parameters to s. opt
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, errors.WithStack(err)
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, errors.WithStack(err)
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// Access contains global keys to all Treezor objects.
type Access struct {
	// AccessSignature can be mandatory for specific context. Treezor will contact you if so.
	AccessSignature *string `url:"accessSignature,omitempty" json:"-"`
	// AccessTag is used for idem potency query.
	AccessTag *string `url:"accessTag,omitempty" json:"-"`
	// AccessUserID is used for user's action restriction.
	AccessUserID *string `url:"accessUserId,omitempty" json:"-"`
	// AccessUserIP is used for user's action restriction.
	AccessUserIP *string `url:"accessUserIp,omitempty" json:"-"`
}

// Origin represents who made the request.
type Origin string

const (
	// OperatorOrigin represents the support
	OperatorOrigin Origin = "OPERATOR"
	// UserOrigin represents the end user.
	UserOrigin Origin = "USER"
)

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// Float64 is a helper routine that allocates a new float64 value
// to store v and returns a pointer to it.
func Float64(v float64) *float64 { return &v }

// TODO: See how to handle "Access" setup for Get, List and Delete requests
