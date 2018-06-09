package stripe

import (
	"reflect"

	"github.com/stripe/stripe-go/form"
)

// Query is the function used to get a page listing.
type Query func(*form.Values) ([]interface{}, ListMeta, error)

// Query2 is the function used to get a page listing.
type Query2 func(*Params, *form.Values) ([]interface{}, ListMeta, error)

// Iter provides a convenient interface
// for iterating over the elements
// returned from paginated list API calls.
// Successive calls to the Next method
// will step through each item in the list,
// fetching pages of items as needed.
// Iterators are not thread-safe, so they should not be consumed
// across multiple goroutines.
type Iter struct {
	cur        interface{}
	err        error
	formValues *form.Values
	listParams ListParams
	meta       ListMeta
	query      Query
	query2     Query2
	values     []interface{}
}

// GetIter returns a new Iter for a given query and its options.
func GetIter(listParams *ListParams, formValues *form.Values, query Query) *Iter {
	iter := &Iter{}
	iter.query = query

	p := listParams
	if p == nil {
		p = &ListParams{}
	}
	iter.listParams = *p

	q := formValues
	if q == nil {
		q = &form.Values{}
	}
	iter.formValues = q

	iter.getPage()
	return iter
}

// TODO: After every list API call uses GetIter2, remove GetIter, then rename
// all instances of GetIter2 to GetIter. This only exists as a separate method
// to keep the build/tests working while we refactor.
func GetIter2(container ListParamsContainer, query Query2) *Iter {
	var listParams *ListParams
	formValues := &form.Values{}

	if container != nil {
		reflectValue := reflect.ValueOf(container)

		// See the comment on Call in stripe.go.
		if reflectValue.Kind() == reflect.Ptr && !reflectValue.IsNil() {
			listParams = container.GetListParams()
			form.AppendTo(formValues, container)
		}
	}

	if listParams == nil {
		listParams = &ListParams{}
	}
	iter := &Iter{
		formValues: formValues,
		listParams: *listParams,
		query2:     query,
	}

	iter.getPage()

	return iter
}

func (it *Iter) getPage() {
	if it.query != nil {
		it.values, it.meta, it.err = it.query(it.formValues)
	} else {
		it.values, it.meta, it.err = it.query2(it.listParams.GetParams(), it.formValues)
	}

	if it.listParams.EndingBefore != nil {
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
	if len(it.values) == 0 && it.meta.HasMore && !it.listParams.Single {
		// determine if we're moving forward or backwards in paging
		if it.listParams.EndingBefore != nil {
			it.listParams.EndingBefore = String(listItemID(it.cur))
			it.formValues.Set(EndingBefore, *it.listParams.EndingBefore)
		} else {
			it.listParams.StartingAfter = String(listItemID(it.cur))
			it.formValues.Set(StartingAfter, *it.listParams.StartingAfter)
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
