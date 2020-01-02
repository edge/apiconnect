// Edge Network
// (c) 2020 Edge Network technologies Ltd.

package apiconnect

import (
	"fmt"

	"github.com/edge/utils/pkg/strhelp"
)

const (
	requestTimeout       = 5
	Regex                = "$regex"
	LessThan             = "$lt"
	LessThanOrEqualTo    = "$lte"
	GreaterThan          = "$gt"
	GreaterThanOrEqualTo = "$gte"
	Equals               = "$eq"
	Not                  = "$not"
	NotEquals            = "$ne"
	NotIn                = "$nin"
	In                   = "$in"
)

// FieldFilter is an interface for a filter.
type FieldFilter interface {
	// String returns a stringified filter.
	String() string
}

// BooleanFilter stores boolean operator and value.
type BooleanFilter struct {
	Field    string
	Operator string
	Value    bool
}

// String returns a stringified filter.
func (b *BooleanFilter) String() string {
	return fmt.Sprintf(`"%s": {"%s": %t}`, b.Field, b.Operator, b.Value)
}

// BoolFilter returns a new instance of BooleanFilter.
func BoolFilter(f, o string, v bool) FieldFilter {
	return &BooleanFilter{
		Field:    f,
		Operator: o,
		Value:    v,
	}
}

// NumberFilter stores number operator and value.
type NumberFilter struct {
	Field    string
	Operator string
	Value    uint64
}

// String returns a stringified filter.
func (n *NumberFilter) String() string {
	return fmt.Sprintf(`"%s": {"%s": %d}`, n.Field, n.Operator, n.Value)
}

// NumFilter returns a new instance of NumberFilter.
func NumFilter(f, o string, v uint64) FieldFilter {
	return &NumberFilter{
		Field:    f,
		Operator: o,
		Value:    v,
	}
}

// StringFilter stores number operator and value.
type StringFilter struct {
	Field    string
	Operator string
	Value    string
}

// String returns a stringified filter.
func (s *StringFilter) String() string {
	return fmt.Sprintf(`"%s": {"%s": "%s"}`, s.Field, s.Operator, s.Value)
}

// StrFilter returns a new instance of StringFilter.
func StrFilter(f, o, v string) FieldFilter {
	return &StringFilter{
		Field:    f,
		Operator: o,
		Value:    v,
	}
}

// Filters is a wrapper for field filters.
type Filters struct {
	Value []FieldFilter
}

// ToQueryString converts filter to a query string.
func (f *Filters) ToQueryString() string {
	if f.Value == nil {
		return ""
	}

	filters := make([]string, 0)

	for _, f := range f.Value {
		filters = append(filters, f.String())
	}

	return fmt.Sprintf(`{%s}`, strhelp.Join(",", 1, filters...))
}

// NewFilters returns a new instance of Filters.
func NewFilters(filters ...FieldFilter) *Filters {
	f := &Filters{
		Value: make([]FieldFilter, 0),
	}

	f.Value = append(f.Value, filters...)
	return f
}
