package stripe

import (
	"reflect"

	"github.com/max-cape/stripe-go-test/form"
)

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
	list       ListContainer
	listParams ListParams
	meta       *ListMeta
	query      Query
	values     []interface{}
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

// List returns the current list object which the iterator is currently using.
// List objects will change as new API calls are made to continue pagination.
func (it *Iter) List() ListContainer {
	return it.list
}

// Meta returns the list metadata.
func (it *Iter) Meta() *ListMeta {
	return it.meta
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

func (it *Iter) getPage() {
	it.values, it.list, it.err = it.query(it.listParams.GetParams(), it.formValues)
	it.meta = it.list.GetListMeta()

	if it.listParams.EndingBefore != nil {
		// We are moving backward,
		// but items arrive in forward order.
		reverse(it.values)
	}
}

// Query is the function used to get a page listing.
type Query func(*Params, *form.Values) ([]interface{}, ListContainer, error)

// GetIter returns a new Iter for a given query and its options.
func GetIter(container ListParamsContainer, query Query) *Iter {
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
		query:      query,
	}

	iter.getPage()

	return iter
}

// v1List provides a convenient interface for iterating over the elements
// returned from paginated list API calls. It is meant to be an improvement
// over the Iter type, which was written before Go introduced generics and iter.Seq2.
// Calling the `All` allows you to iterate over all items in the list,
// with automatic pagination.
type v1List[T any] struct {
	cur           *T
	err           error
	formValues    *form.Values
	listContainer ListContainer
	listParams    ListParams
	listMeta      *ListMeta
	query         v1Query[T]
	values        []*T
}

// All returns a Seq2 that will be evaluated on each item in a v1List.
// The All function will continue to fetch pages of items as needed.
func (it *v1List[T]) All() Seq2[*T, error] {
	return func(yield func(*T, error) bool) {
		for it.next() {
			if !yield(it.cur, nil) {
				return
			}
		}
		if it.err != nil {
			if !yield(nil, it.err) {
				return
			}
		}
	}
}

// next advances the V1List to the next item in the list,
// which will then be available
// through the current method.
// It returns false when the iterator stops
// at the end of the list.
func (it *v1List[T]) next() bool {
	if len(it.values) == 0 && it.listMeta.HasMore && !it.listParams.Single {
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

func (it *v1List[T]) getPage() {
	it.values, it.listContainer, it.err = it.query(it.listParams.GetParams(), it.formValues)
	it.listMeta = it.listContainer.GetListMeta()

	if it.listParams.EndingBefore != nil {
		// We are moving backward,
		// but items arrive in forward order.
		reverse(it.values)
	}
}

// Query is the function used to get a page listing.
type v1Query[T any] func(*Params, *form.Values) ([]*T, ListContainer, error)

// newV1List returns a new v1List for a given query and its options, and initializes
// it by fetching the first page of items.
func newV1List[T any](container ListParamsContainer, query v1Query[T]) *v1List[T] {
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
	iter := &v1List[T]{
		formValues: formValues,
		listParams: *listParams,
		query:      query,
	}

	iter.getPage()

	return iter
}

func listItemID[T any](x T) string {
	return reflect.ValueOf(x).Elem().FieldByName("ID").String()
}

func reverse[T any](a []T) {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
}

// Seq2 is the same as the iter.Seq2 type in Go 1.23+. It is used as the return type
// of All methods. If you are using Go 1.23+, you can just range over the an All
// method directly, e.g.,
//
//	for event, err := range sc.V2Events.All() {
//		// check err and do something with event
//	}
//
// For older versions of Go, the yield function should return false
// to stop iteration or true to continue.
type Seq2[K, V any] func(yield func(K, V) bool)

// V2List contains a page of data received from a List API call,
// and the means to paginate to the next page of data via the fetch function.
type V2List[T any] struct {
	fetch       Fetch[T]
	params      ParamsContainer
	initialized bool
	// Page contains the items returned from the last API call.
	V2Page[T]
}

// V2Page is represents a single page returned from a List API call.
// Users will not ordinaily interact with this type directly.
type V2Page[T any] struct {
	APIResource
	Data            []T    `json:"data"`
	NextPageURL     string `json:"next_page_url"`
	PreviousPageURL string `json:"previous_page_url"`
}

// NewV2List creates a new V2List with the given path and fetch function.
func NewV2List[T any](path string, p ParamsContainer, fetch Fetch[T]) *V2List[T] {
	return &V2List[T]{
		fetch:  fetch,
		params: p,
		V2Page: V2Page[T]{NextPageURL: path},
	}
}

// Fetch is a function that fetches a page of items.
type Fetch[T any] func(path string, p ParamsContainer) (*V2Page[T], error)

// All returns a Seq2 that will be evaluated on each item in a V2List.
// The All function will continue to fetch pages of items as needed.
func (s *V2List[T]) All() Seq2[T, error] {
	return func(yield func(T, error) bool) {
		var fetchMore bool
		// fetch inital page
		err := s.page()
		if err != nil && !yield(*new(T), err) {
			return
		}
		s.initialized = true
		fetchMore = (s.NextPageURL != "")

		for len(s.Data) > 0 {
			for _, item := range s.Data {
				if !yield(item, nil) {
					return
				}
			}

			if !fetchMore {
				return
			}
			err := s.page()
			if err != nil && !yield(*new(T), err) {
				return
			}
			fetchMore = (s.NextPageURL != "")
		}
	}
}

// page fetches the next page of items and updates the Seq's state.
// It returns true if there exist more pages to fetch, and false if
// that was the last page.
func (s *V2List[T]) page() error {
	// if we've already fetched a page, the next page URL
	// already contains all of the query parameters
	var params ParamsContainer
	if s.initialized {
		params = &Params{}
	} else {
		params = s.params
	}

	next, err := s.fetch(s.NextPageURL, params)
	if err != nil {
		return err
	}

	s.V2Page = *next
	return nil
}
