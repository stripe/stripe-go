package stripe

import (
	"net/url"
	"reflect"
)

// Query is the function used to get a page listing.
type Query func(url.Values) ([]interface{}, ListMeta, error)

// Iter provides a convenient interface
// for iterating over the elements
// returned from paginated list API calls.
// Successive calls to the Next method
// will step through each item in the list,
// fetching pages of items as needed.
// Iterators are not thread-safe, so they should not be consumed
// across multiple goroutines.
type Iter struct {
	query  Query
	qs     url.Values
	values []interface{}
	meta   ListMeta
	params ListParams
	err    error
	cur    interface{}
}

// GetIter returns a new Iter for a given query and its options.
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

	iter.getPage()
	return iter
}

func (it *Iter) getPage() {
	it.values, it.meta, it.err = it.query(it.qs)
	if it.params.End != "" {
		// We are moving backward,
		// but items arrive in forward order.
		reverse(it.values)
	}
}

// Next advances the Iter to the next item in the list,
// which will then be available
// through the Current method.
// It returns false when the iterator stops
// at the end of the list.
func (it *Iter) Next() bool {
	if len(it.values) == 0 && it.meta.More && !it.params.Single {
		// determine if we're moving forward or backwards in paging
		if it.params.End != "" {
			it.params.End = listItemID(it.cur)
			it.qs.Set(endbefore, it.params.End)
		} else {
			it.params.Start = listItemID(it.cur)
			it.qs.Set(startafter, it.params.Start)
		}
		it.getPage()
	}
	if len(it.values) == 0 {
		return false
	}
	it.cur = it.values[0]
	it.values = it.values[1:]
	return true
}

// Current returns the most recent item
// visited by a call to Next.
func (it *Iter) Current() interface{} {
	return it.cur
}

// Err returns the error, if any,
// that caused the Iter to stop.
// It must be inspected
// after Next returns false.
func (it *Iter) Err() error {
	return it.err
}

// Meta returns the list metadata.
func (it *Iter) Meta() *ListMeta {
	return &it.meta
}

func listItemID(x interface{}) string {
	return reflect.ValueOf(x).Elem().FieldByName("ID").String()
}

func reverse(a []interface{}) {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
}
