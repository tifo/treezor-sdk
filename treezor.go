//go:generate go run internal/scripts/gen_accessors.go -v

package treezor

import (
	"net/url"
	"time"

	"github.com/pkg/errors"

	"github.com/tifo/treezor-sdk/internal/query"
)

type service struct {
	client *HTTPClient
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
	CreatedFrom *time.Time `layout:"Treezor" loc:"Europe/Paris" url:"createdDateFrom,omitempty" json:"-"`
	// For paginated result sets, the creation date up to which you want to filter the request result.
	CreatedTo *time.Time `layout:"Treezor" loc:"Europe/Paris" url:"createdDateFrom,omitempty" json:"-"`
	// For paginated result sets, the modification date from which you want to filter the request result.
	UpdatedFrom *time.Time `layout:"Treezor" loc:"Europe/Paris" url:"createdDateFrom,omitempty" json:"-"`
	// For paginated result sets, the modification date up to which you want to filter the request result.
	UpdatedTo *time.Time `layout:"Treezor" loc:"Europe/Paris" url:"createdDateFrom,omitempty" json:"-"`
}

// NOTE: should we rename those to exactly their API names to be consistent with the Treezor documentation ?

// addOptions adds the parameters in opt as URL query parameters to s. opt
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	qs, err := query.Values(opt)
	if err != nil {
		return s, errors.WithStack(err)
	}
	if len(qs) == 0 {
		return s, nil
	}

	u, err := url.Parse(s)
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
