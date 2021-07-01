//go:generate go run internal/scripts/gen_accessors.go -v

package treezor

import (
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"

	"github.com/tifo/treezor-sdk/internal/types"
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
	Page *int64 `url:"pageNumber,omitempty" json:"-"`
	// For paginated result sets, the number of results to include per page.
	PerPage *int64 `url:"pageCount,omitempty" json:"-"`
	// For paginated result sets, the resource element you want to sort the list with.
	SortBy *string `url:"sortBy,omitempty" json:"-"`
	// For paginated result sets, The order you want to sort the list.
	// Possible values: DESC and ASC.
	SortOrder SortOrder `url:"sortOrder,omitempty" json:"-"`
	// For paginated result sets, the creation date from which you want to filter the request result.
	// Format: YYYY-MM-DD HH:MM:SS
	CreatedFrom *types.Timestamp `layout:"2006-01-02 15:04:05" url:"createdDateFrom,omitempty" json:"-"`
	// For paginated result sets, the creation date up to which you want to filter the request result.
	// Format: YYYY-MM-DD HH:MM:SS
	CreatedTo *types.Timestamp `layout:"2006-01-02 15:04:05" url:"createdDateFrom,omitempty" json:"-"`
	// For paginated result sets, the modification date from which you want to filter the request result.
	// Format: YYYY-MM-DD HH:MM:SS
	UpdatedFrom *types.Timestamp `layout:"2006-01-02 15:04:05" url:"createdDateFrom,omitempty" json:"-"`
	// For paginated result sets, the modification date up to which you want to filter the request result.
	// Format: YYYY-MM-DD HH:MM:SS
	UpdatedTo *types.Timestamp `layout:"2006-01-02 15:04:05" url:"createdDateFrom,omitempty" json:"-"`
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

// TODO: see presence of totalRows, codeStatus and informationStatus in all models
