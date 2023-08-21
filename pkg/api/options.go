package api

import (
	"fmt"
	"net/url"

	"github.com/clearchanneloutdoor/io-sdk-golang/internal"
)

const (
	defaultPageLimit  = 100
	defaultPageOffset = 0
	filterFMT         = "filter[%s]"
	pageFMT           = "page[%s]"
	pageLimit         = "limit"
	pageOffset        = "offset"
)

type Options struct {
	fields []string
	filter map[string][]string
	page   map[string]int
	sort   []string
}

// EmptyOptions returns a new options struct which can be used in
// conjunction with the various API methods to control the results
// returned from the API.
func EmptyOptions() *Options {
	return &Options{
		fields: []string{},
		filter: map[string][]string{},
		page:   map[string]int{},
		sort:   []string{},
	}
}

// AddField adds a field to the options used to control which fields
// are returned in the results for an API request.
func (o *Options) AddField(field ...string) *Options {
	for _, f := range field {
		if !internal.ContainsValue(o.fields, f) {
			o.fields = append(o.fields, f)
		}
	}

	return o
}

// AddFilter adds a field and value to the options used to control
// which results are returned for an API request.
func (o *Options) AddFilter(field string, value ...any) *Options {
	if len(value) == 0 {
		return o
	}

	for _, val := range value {
		var v string
		switch tv := val.(type) {
		case fmt.Stringer:
			v = tv.String()
		case string:
			v = tv
		default:
			v = fmt.Sprintf("%v", tv)
		}

		// check to see if the field already exists within the filters
		if f, ok := o.filter[field]; ok {
			// check to see if the value already exists within the filter
			if internal.ContainsValue(f, v) {
				return o
			}

			// append the filter and return
			o.filter[field] = append(f, v)
			return o
		}

		// add the filter and return
		o.filter[field] = []string{v}
	}

	return o
}

// AddPage sets pagination values for the desired pagination
// strategy... the PageDefauilt uses "offset" and "limit", but this method
// can be used to suggest a different strategy (e.g. "page" and "size") if
// desired (though, PageDefault is recommended).
func (o *Options) AddPage(field string, value int) *Options {
	o.page[field] = value
	return o
}

// AddSort adds a field to the options used to control which fields
// are used to sort the results for an API request. Provided a field
// name prefixed with a "-" will sort the results in descending order
// (e.g. "-name" or "-updatedAt").
func (o *Options) AddSort(field ...string) *Options {
	for _, f := range field {
		if !internal.ContainsValue(o.sort, f) {
			o.sort = append(o.sort, f)
		}
	}

	return o
}

// NextPage increments the offset value by the limit value for
// navigating to the next page of results. This method assumes
// the default pagination strategy of "offset" and "limit" is used.
func (o *Options) NextPage() *Options {
	// make sure the limit is set
	if _, ok := o.page[pageLimit]; !ok {
		o.page[pageLimit] = defaultPageLimit
	}

	// make sure the limit is valid
	if o.page[pageLimit] < 1 {
		return o
	}

	o.page[pageOffset] = o.page[pageOffset] + o.page[pageLimit]
	return o
}

// PageDefault resets the pagination options to the default values
// of "offset" and "limit". This method is useful for resetting the
// pagination options when switching between different pagination
// strategies.
func (o *Options) PageDefault() *Options {
	// reset the page options
	o.page = map[string]int{}
	o.page[pageOffset] = defaultPageOffset
	o.page[pageLimit] = defaultPageLimit
	return o
}

// PreviousPage decrements the offset value by the limit value for
// navigating to the previous page of results. This method assumes
// the default pagination strategy of "offset" and "limit" is used.
func (o *Options) PreviousPage() *Options {
	// make sure the limit is set
	if _, ok := o.page[pageLimit]; !ok {
		o.page[pageLimit] = defaultPageLimit
	}

	// make sure the limit is valid
	if o.page[pageLimit] < 1 {
		return o
	}

	v := o.page[pageOffset] - o.page[pageLimit]

	// ensure the value is not less than 0
	if v < 0 {
		v = 0
	}

	o.page[pageOffset] = v
	return o
}

// Query returns a url.Values representation of the options struct
// which can be used to construct a query string for an API request.
func (o *Options) Query() url.Values {
	query := url.Values{}

	// add fields to the query
	if len(o.fields) > 0 {
		query["fields"] = o.fields
	}

	// add filters to the query
	for field, values := range o.filter {
		fltr := fmt.Sprintf(filterFMT, field)
		query[fltr] = values
	}

	// add pagination to the query
	for field, value := range o.page {
		fltr := fmt.Sprintf(pageFMT, field)
		query.Set(fltr, fmt.Sprintf("%d", value))
	}

	// add pagination to the query
	if len(o.sort) > 0 {
		query["sort"] = o.sort
	}

	return query
}

// RemoveField removes a field from the options used to control which fields
// are returned in the results for an API request.
func (o *Options) RemoveField(field string) *Options {
	for i, f := range o.fields {
		if f == field {
			o.fields = append(o.fields[:i], o.fields[i+1:]...)
			return o
		}
	}

	return o
}

// RemoveFilter removes a field and value from the options used to control
// which results are returned for an API request.
func (o *Options) RemoveFilter(field string) *Options {
	delete(o.filter, field)
	return o
}

// RemovePage removes pagination values for the desired pagination
// strategy... the PageDefauilt uses "offset" and "limit", but this method
// can be used to suggest a different strategy (e.g. "page" and "size") if
// desired (though, PageDefault is recommended).
func (o *Options) RemovePage(field string) *Options {
	delete(o.page, field)
	return o
}

// RemoveSort removes a field from the options used to control which fields
// are used to sort the results for an API request.
func (o *Options) RemoveSort(field string) *Options {
	for i, f := range o.sort {
		if f == field {
			o.sort = append(o.sort[:i], o.sort[i+1:]...)
			return o
		}
	}

	return o
}
