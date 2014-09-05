package stripe

import (
	"net/url"
	"reflect"
)

// Query is the function used to get a page listing.
type Query func(url.Values) ([]interface{}, ListResponse, error)

// Iter represents an interator used for list pagination.
// Iterators are not thread-safe at this time, so they should not be consumed
// across multiple goroutines.
type Iter struct {
	end         bool
	query       Query
	qs          url.Values
	values      []interface{}
	meta        ListResponse
	params      ListParams
	cur, pstart int
	err         error
}

// GetIter returns a new iterator for a given query and its options.
func GetIter(params *ListParams, qs *url.Values, query Query) *Iter {
	iter := &Iter{}
	iter.query = query

	p := params
	if p == nil {
		p = &ListParams{}
	}
	iter.params = *p

	q := qs
	if q == nil {
		q = &url.Values{}
	}
	iter.qs = *q

	return iter
}

// Next returns the next entry in the page.
// By default, this loads a new page each time it's done with the current one.
func (i *Iter) Next() (interface{}, *ListResponse, error) {
	// if the previous invocation resulted in us being complete,
	// no point in doing any work
	if i.end {
		return nil, nil, i.err
	}

	// initial run when we have no page loaded
	if len(i.values) == 0 {
		i.values, i.meta, i.err = i.query(i.qs)
	}

	// either there was a failure or no results were returned
	if i.err != nil || i.cur == len(i.values) {
		i.end = true
		return nil, nil, i.err
	}

	ret := i.values[i.cur]
	i.cur += 1

	// we've reached the end of the page
	if i.cur == len(i.values) {
		// if there are no more results or we're supposed to stop after a
		// single page just bail
		if i.params.Single || !i.meta.More {
			i.end = true
		} else {
			// determine if we're moving forward or backwards in paging
			if len(i.params.End) != 0 {
				i.params.End = reflect.ValueOf(i.values[i.pstart]).Elem().FieldByName("Id").String()
				i.qs.Set(endbefore, i.params.End)
			} else {
				i.params.Start = reflect.ValueOf(i.values[i.cur-1]).Elem().FieldByName("Id").String()
				i.qs.Set(startafter, i.params.Start)
			}

			// now actually load the next page
			var page []interface{}
			page, i.meta, i.err = i.query(i.qs)

			if i.err != nil {
				i.end = true
				return nil, nil, i.err
			}

			i.pstart = len(i.values)
			i.values = append(i.values, page...)
		}
	}

	return ret, &i.meta, i.err
}

// End returns true if there are no more iterations to be performed.
func (i *Iter) End() bool {
	return i.end
}
