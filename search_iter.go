package stripe

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/stripe/stripe-go/v82/form"
)

//
// Public constants
//

// Contains constants for the names of parameters used for pagination in search APIs.
const (
	Page = "page"
)

//
// Public types
//

// SearchIter provides a convenient interface
// for iterating over the elements
// returned from paginated search API calls.
// Successive calls to the Next method
// will step through each item in the search results,
// fetching pages of items as needed.
// Iterators are not thread-safe, so they should not be consumed
// across multiple goroutines.
type SearchIter struct {
	cur             interface{}
	err             error
	formValues      *form.Values
	searchContainer SearchContainer
	searchParams    SearchParams
	meta            *SearchMeta
	query           SearchQuery
	values          []interface{}
}

// Current returns the most recent item
// visited by a call to Next.
func (it *SearchIter) Current() interface{} {
	return it.cur
}

// Err returns the error, if any,
// that caused the SearchIter to stop.
// It must be inspected
// after Next returns false.
func (it *SearchIter) Err() error {
	return it.err
}

// SearchResult returns the current search result container which the iterator is currently using.
// Objects will change as new API calls are made to continue pagination.
func (it *SearchIter) SearchResult() SearchContainer {
	return it.searchContainer
}

// Meta returns the search metadata.
func (it *SearchIter) Meta() *SearchMeta {
	return it.meta
}

// Next advances the SearchIter to the next item in the search results,
// which will then be available
// through the Current method.
// It returns false when the iterator stops
// at the end of the search results.
func (it *SearchIter) Next() bool {
	if len(it.values) == 0 && it.meta.HasMore && !it.searchParams.Single {
		if it.meta.NextPage != nil {
			it.formValues.Set(Page, *it.meta.NextPage)
			it.getPage()
		}
	}
	if len(it.values) == 0 {
		return false
	}
	it.cur = it.values[0]
	it.values = it.values[1:]
	return true
}

func (it *SearchIter) getPage() {
	it.values, it.searchContainer, it.err = it.query(it.searchParams.GetParams(), it.formValues)
	it.meta = it.searchContainer.GetSearchMeta()
}

// SearchQuery is the function used to get search results.
type SearchQuery func(*Params, *form.Values) ([]interface{}, SearchContainer, error)

//
// Public functions
//

// GetSearchIter returns a new SearchIter for a given query and its options.
func GetSearchIter(container SearchParamsContainer, query SearchQuery) *SearchIter {
	var searchParams *SearchParams
	formValues := &form.Values{}

	if container != nil {
		reflectValue := reflect.ValueOf(container)

		// See the comment on Call in stripe.go.
		if reflectValue.Kind() == reflect.Ptr && !reflectValue.IsNil() {
			searchParams = container.GetSearchParams()
			form.AppendTo(formValues, container)
		}
	}

	if searchParams == nil {
		searchParams = &SearchParams{}
	}
	iter := &SearchIter{
		formValues:   formValues,
		searchParams: *searchParams,
		query:        query,
	}

	iter.getPage()

	return iter
}

// v1SearchList provides a convenient interface for iterating over the elements
// returned from paginated list API calls. It is meant to be an improvement
// over the SearchIter type, which was written before Go introduced generics and iter.Seq2.
// Calling the `All` allows you to iterate over all items in the list,
// with automatic pagination.
type v1SearchList[T any] struct {
	cur          T
	err          error
	formValues   *form.Values
	searchParams SearchParams
	query        v1SearchQuery[T]
	*v1SearchPage[T]
}

type v1SearchPage[T any] struct {
	APIResource
	SearchMeta
	Data []T `json:"data"`
}

func (it *v1SearchList[T]) All() Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for it.next() {
			if !yield(it.cur, nil) {
				return
			}
		}
		if it.err != nil {
			if !yield(*new(T), it.err) {
				return
			}
		}
	}
}

// next advances the v1SearchList to the next item in the list,
// which will then be available
// through the current method.
// It returns false when the iterator stops
// at the end of the list.
func (it *v1SearchList[T]) next() bool {
	// Refresh the page if there is an more data to fetch
	if len(it.Data) == 0 && it.HasMore && !it.searchParams.Single && it.NextPage != nil {
		it.formValues.Set(Page, *it.NextPage)
		it.getPage()
	}
	// If there was no new data after fetching, return false
	if len(it.Data) == 0 {
		return false
	}
	it.cur = it.Data[0]
	it.Data = it.Data[1:]
	return true
}

func (it *v1SearchList[T]) getPage() {
	page, err := it.query(it.searchParams.GetParams(), it.formValues)
	it.v1SearchPage = page
	if err != nil {
		it.err = err
		return
	}
	if err := maybeAddLastResponseSearch(page); err != nil {
		it.err = err
		return
	}
}

// maybeAddLastResponse adds the LastResponse to the items in the page.
// It parses the page's JSON and adds each `data` item's JSON to the
// LastResponse of the corresponding resource. Note that not
// every resource implements the LastResponseSetter interface.
func maybeAddLastResponseSearch[T any](page *v1SearchPage[T]) error {
	if page.LastResponse == nil {
		return nil
	}
	lastResponse := page.LastResponse

	var pageData struct {
		Data []json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(lastResponse.RawJSON, &pageData); err != nil {
		return err
	}

	if len(pageData.Data) != len(page.Data) {
		return fmt.Errorf("mismatch in data length for requestID %s", lastResponse.RequestID)
	}

	for i, item := range page.Data {
		// Note that not every resource implements the LastResponseSetter interface
		// (e.g. CreditNoteLineItem).
		if item, ok := any(item).(LastResponseSetter); ok {
			// Create a copy of the original response with individual item's raw JSON
			itemResponse := &APIResponse{
				Header:         lastResponse.Header,
				IdempotencyKey: lastResponse.IdempotencyKey,
				RawJSON:        []byte(pageData.Data[i]),
				RequestID:      lastResponse.RequestID,
				Status:         lastResponse.Status,
				StatusCode:     lastResponse.StatusCode,
			}
			item.SetLastResponse(itemResponse)
		}
	}
	return nil
}

// SearchQuery is the function used to get search results.
type v1SearchQuery[T any] func(*Params, *form.Values) (*v1SearchPage[T], error)

//
// Public functions
//

// GetSearchIter returns a new SearchIter for a given query and its options.
func newV1SearchList[T any](container SearchParamsContainer, query v1SearchQuery[T]) *v1SearchList[T] {
	var searchParams *SearchParams
	formValues := &form.Values{}

	if container != nil {
		reflectValue := reflect.ValueOf(container)

		// This is a little unfortunate, but Go makes it impossible to compare
		// an interface value to nil without the use of the reflect package and
		// its true disciples insist that this is a feature and not a bug.
		//
		// Here we do invoke reflect because (1) we have to reflect anyway to
		// use encode with the form package, and (2) the corresponding removal
		// of boilerplate that this enables makes the small performance penalty
		// worth it.
		if reflectValue.Kind() == reflect.Ptr && !reflectValue.IsNil() {
			searchParams = container.GetSearchParams()
			form.AppendTo(formValues, container)
		}
	}

	if searchParams == nil {
		searchParams = &SearchParams{}
	}
	iter := &v1SearchList[T]{
		formValues:   formValues,
		searchParams: *searchParams,
		query:        query,
	}

	iter.getPage()

	return iter
}
