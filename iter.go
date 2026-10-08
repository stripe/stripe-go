package stripe

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"reflect"

	"github.com/stripe/stripe-go/v87/form"
)

var errEmptyPageWithHasMore = errors.New("stripe: invalid list response: has_more is true but data is empty")

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
	if len(it.values) == 0 && it.err == nil && it.meta.HasMore && !it.listParams.Single {
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
	if it.err == nil {
		it.err = validateV1ListPage(len(it.values), it.meta.HasMore, it.listParams.Single)
	}

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

// V1List provides a convenient interface for iterating over the elements
// returned from paginated list API calls. It is meant to be an improvement
// over the Iter type, which was written before Go introduced generics and iter.Seq2.
// Calling the `All` allows you to iterate over all items in the list,
// with automatic pagination.
type V1List[T any] struct {
	err        error
	formValues *form.Values
	listParams ListParams
	query      v1Query[T]
	backward   bool
	v1Page     *v1Page[T]
}

// v1Page represents a single page returned from a V1 List API call.
// The internal state will be updated by the parent V1List when the
// Page method is called.
type v1Page[T any] struct {
	APIResource
	ListMeta
	Data []T `json:"data"`
}

// All returns a Seq2 that will be evaluated on each item in a V1List.
// The All function will continue to fetch pages of items as needed.
func (l *V1List[T]) All(ctx context.Context) Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for {
			for _, item := range l.Data() {
				if !yield(item, nil) {
					return
				}
			}
			if l.err != nil {
				if !yield(*new(T), l.Err()) {
					return
				}
			}
			if !l.hasMore() {
				return
			}
			l.page(ctx)
		}
	}
}

// Data returns the data for the current page.
func (l *V1List[T]) Data() []T {
	return l.v1Page.Data
}

// Err returns the error for the current page.
func (l *V1List[T]) Err() error {
	return l.err
}

// Meta returns the metadata for the current page.
func (l *V1List[T]) Meta() ListMeta {
	return l.v1Page.ListMeta
}

// LastResponse returns the last response for the current page.
func (l *V1List[T]) LastResponse() *APIResponse {
	return l.v1Page.LastResponse
}

// page updates the V1List's state by fetching the next page of items.
func (l *V1List[T]) page(ctx context.Context) {
	if len(l.Data()) > 0 && l.backward {
		l.listParams.EndingBefore = String(listItemID(l.Data()[len(l.Data())-1]))
		l.formValues.Set(EndingBefore, *l.listParams.EndingBefore)
	} else if len(l.Data()) > 0 {
		l.listParams.StartingAfter = String(listItemID(l.Data()[len(l.Data())-1]))
		l.formValues.Set(StartingAfter, *l.listParams.StartingAfter)
	}
	page, err := l.query(ctx, l.listParams.GetParams(), l.formValues)
	l.v1Page = page
	if err != nil {
		l.err = err
		return
	}
	if err := validateV1ListPage(len(page.Data), page.HasMore, l.listParams.Single); err != nil {
		l.err = err
		return
	}
	if err := maybeAddLastResponseV1(page); err != nil {
		l.err = err
		return
	}

	if l.backward {
		// We are moving backward,
		// but items arrive in forward order.
		reverse(l.Data())
	}
}

// hasMore returns true if there is another page of items to fetch.
func (l *V1List[T]) hasMore() bool {
	if l == nil || l.err != nil {
		return false
	}
	return l.v1Page.HasMore && !l.listParams.Single
}

func validateV1ListPage(dataLen int, hasMore, single bool) error {
	if dataLen == 0 && hasMore && !single {
		return errEmptyPageWithHasMore
	}
	return nil
}

// maybeAddLastResponseV1 adds the LastResponse to the items in the page.
// It parses the page's JSON and adds each `data` item's JSON to the
// LastResponse of the corresponding resource. Note that not
// every resource implements the LastResponseSetter interface.
func maybeAddLastResponseV1[T any](page *v1Page[T]) error {
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

// v1Query is the function used to get a page listing.
type v1Query[T any] func(context.Context, *Params, *form.Values) (*v1Page[T], error)

// newV1List returns a new v1List for a given query and its options, and initializes
// it by fetching the first page of items.
func newV1List[T any](ctx context.Context, container ListParamsContainer, query v1Query[T]) *V1List[T] {
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
	iter := &V1List[T]{
		formValues: formValues,
		listParams: *listParams,
		query:      query,
		backward:   listParams.EndingBefore != nil,
		v1Page:     &v1Page[T]{},
	}

	iter.page(ctx)

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
// of List methods. If you are using Go 1.23+, you can just range over the an List
// method directly, e.g.,
//
//	for event, err := range sc.V2CoreEvents.List(...) {
//		// check err and do something with event
//	}
//
// For older versions of Go, the yield function should return false
// to stop iteration or true to continue.
type Seq2[K, V any] func(yield func(K, V) bool)

// V2List contains a page of data received from a List API call,
// and the means to paginate to the next page of data via the fetch function.
type V2List[T any] struct {
	fetch       v2Query[T]
	params      ParamsContainer
	initialized bool
	err         error
	// Page contains the items returned from the last API call.
	v2Page *V2Page[T]
}

// V2Page is represents a single page returned from a V2 List API call.
type V2Page[T any] struct {
	APIResource
	V2ListMeta
	Data []T `json:"data"`
}

// Data returns the data for the current page.
func (l *V2List[T]) Data() []T {
	return l.v2Page.Data
}

// Err returns the error for the current page.
func (l *V2List[T]) Err() error {
	return l.err
}

// Meta returns the metadata for the current page.
func (l *V2List[T]) Meta() V2ListMeta {
	return l.v2Page.V2ListMeta
}

// LastResponse returns the last response for the current page.
func (l *V2List[T]) LastResponse() *APIResponse {
	return l.v2Page.LastResponse
}

// All returns a Seq2 that will be evaluated on each item in a V2List.
// The All function will continue to fetch pages of items as needed.
func (l *V2List[T]) All(ctx context.Context) Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for {
			for _, item := range l.Data() {
				if !yield(item, nil) {
					return
				}
			}
			if l.err != nil {
				if !yield(*new(T), l.err) {
					return
				}
			}
			if !l.hasMore() {
				return
			}
			l.page(ctx)
		}
	}
}

// page fetches the next page of items and updates the V2List's state.
// It returns an error if the fetch fails.
func (l *V2List[T]) page(ctx context.Context) {
	// if we've already fetched a page, the next page URL
	// already contains all of the query parameters
	var params ParamsContainer
	if l.initialized {
		params = &Params{}
	} else {
		params = l.params
	}

	next, err := l.fetch(ctx, l.v2Page.NextPageURL, params)
	l.v2Page = next
	if err != nil {
		l.err = err
		return
	}

	if err := maybeAddLastResponseV2(next); err != nil {
		l.err = err
		return
	}
}

// V2SearchList contains API v2 search results and replays the original body when paging.
type V2SearchList[T any] struct {
	fetch       v2SearchQuery[T]
	params      V2SearchParamsContainer
	err         error
	page        *V2SearchPage[T]
	initialized bool
}

// V2SearchPage is a single API v2 search result page.
type V2SearchPage[T any] struct {
	APIResource
	V2ListMeta
	Data       []T   `json:"data"`
	TotalCount int64 `json:"total_count"`
}

func (l *V2SearchList[T]) Data() []T                  { return l.page.Data }
func (l *V2SearchList[T]) Err() error                 { return l.err }
func (l *V2SearchList[T]) Meta() V2ListMeta           { return l.page.V2ListMeta }
func (l *V2SearchList[T]) TotalCount() int64          { return l.page.TotalCount }
func (l *V2SearchList[T]) LastResponse() *APIResponse { return l.page.LastResponse }

func (l *V2SearchList[T]) All(ctx context.Context) Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for {
			for _, item := range l.Data() {
				if !yield(item, nil) {
					return
				}
			}
			if l.err != nil {
				yield(*new(T), l.err)
				return
			}
			if l.page.NextPageURL == "" {
				return
			}
			l.fetchPage(ctx, l.page.NextPageURL)
		}
	}
}

func (l *V2SearchList[T]) fetchPage(ctx context.Context, path string) {
	path, params, err := splitV2SearchLimit(path, l.params, !l.initialized)
	if err != nil {
		l.page = &V2SearchPage[T]{}
		l.err = err
		return
	}
	next, err := l.fetch(ctx, path, params)
	l.initialized = true
	if next == nil {
		next = &V2SearchPage[T]{}
		if err == nil {
			err = errors.New("invalid v2 search response: nil page")
		}
	}
	l.page = next
	if err != nil {
		l.err = err
		return
	}
	if err := maybeAddLastResponseV2Search(next); err != nil {
		l.err = err
	}
}

type v2SearchQuery[T any] func(context.Context, string, ParamsContainer) (*V2SearchPage[T], error)

func cloneV2SearchParams(p V2SearchParamsContainer) (V2SearchParamsContainer, error) {
	if p == nil {
		return nil, nil
	}
	value := reflect.ValueOf(p)
	if value.Kind() != reflect.Ptr || value.IsNil() || value.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("v2 search params must be a non-nil pointer to a struct, got %T", p)
	}
	cloneValue := reflect.New(value.Elem().Type())
	cloneValue.Elem().Set(value.Elem())
	clone, ok := cloneValue.Interface().(V2SearchParamsContainer)
	if !ok || p.GetV2SearchParams() == nil || clone.GetV2SearchParams() == nil {
		return nil, fmt.Errorf("v2 search params %T do not expose V2SearchParams", p)
	}
	originalSearch, clonedSearch := p.GetV2SearchParams(), clone.GetV2SearchParams()
	if originalSearch.Limit != nil {
		limit := *originalSearch.Limit
		clonedSearch.Limit = &limit
	}
	originalBase, cloneBase := &originalSearch.Params, &clonedSearch.Params
	cloneBase.Headers = originalBase.Headers.Clone()
	if originalBase.Extra != nil {
		extra := make(map[string][]string, len(originalBase.Extra.Values))
		for key, values := range originalBase.Extra.Values {
			extra[key] = append([]string(nil), values...)
		}
		cloneBase.Extra = &ExtraValues{Values: extra}
	}
	if originalBase.Metadata != nil {
		cloneBase.Metadata = make(map[string]string, len(originalBase.Metadata))
		for key, value := range originalBase.Metadata {
			cloneBase.Metadata[key] = value
		}
	}
	cloneBase.Expand = append([]*string(nil), originalBase.Expand...)
	cloneBase.usage = append([]string(nil), originalBase.usage...)
	return clone, nil
}

func splitV2SearchLimit(path string, p V2SearchParamsContainer, addToPath bool) (string, V2SearchParamsContainer, error) {
	params, err := cloneV2SearchParams(p)
	if err != nil {
		return "", nil, err
	}
	if params == nil {
		return path, nil, nil
	}
	searchParams := params.GetV2SearchParams()
	if searchParams.Limit == nil {
		return path, params, nil
	}
	if addToPath {
		parsed, err := url.Parse(path)
		if err != nil {
			return "", nil, err
		}
		query := parsed.Query()
		query.Set("limit", fmt.Sprint(*searchParams.Limit))
		parsed.RawQuery = query.Encode()
		path = parsed.String()
	}
	searchParams.Limit = nil
	return path, params, nil
}

func newV2SearchList[T any](ctx context.Context, path string, p V2SearchParamsContainer, fetch v2SearchQuery[T]) *V2SearchList[T] {
	list := &V2SearchList[T]{fetch: fetch, page: &V2SearchPage[T]{V2ListMeta: V2ListMeta{NextPageURL: path}}}
	var err error
	list.params, err = cloneV2SearchParams(p)
	if err != nil {
		list.err = err
		return list
	}
	list.fetchPage(ctx, path)
	return list
}

// SearchFetch fetches an API v2 search result page.
type SearchFetch[T any] func(string, ParamsContainer) (*V2SearchPage[T], error)

// NewV2SearchList creates an API v2 search iterator.
func NewV2SearchList[T any](path string, p V2SearchParamsContainer, fetch SearchFetch[T]) *V2SearchList[T] {
	ctx := context.Background()
	if p != nil && reflect.ValueOf(p).Kind() == reflect.Ptr && !reflect.ValueOf(p).IsNil() && p.GetParams() != nil && p.GetParams().Context != nil {
		ctx = p.GetParams().Context
	}
	return newV2SearchList(ctx, path, p, func(ctx context.Context, path string, p ParamsContainer) (*V2SearchPage[T], error) {
		if p != nil && p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		return fetch(path, p)
	})
}

func maybeAddLastResponseV2Search[T any](page *V2SearchPage[T]) error {
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
		if item, ok := any(item).(LastResponseSetter); ok {
			response := *lastResponse
			response.RawJSON = pageData.Data[i]
			item.SetLastResponse(&response)
		}
	}
	return nil
}

// maybeAddLastResponseV2 adds the LastResponse to the items in the page.
// It parses the page's JSON and adds each `data` item's JSON to the
// LastResponse of the corresponding resource. Note that not
// every resource implements the LastResponseSetter interface.
func maybeAddLastResponseV2[T any](page *V2Page[T]) error {
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

// hasMore returns true if there is another page of items to fetch.
func (l *V2List[T]) hasMore() bool {
	if l == nil {
		return false
	}
	return l.v2Page.NextPageURL != ""
}

// newV2List creates a new V2List with the given path and fetch function.
func newV2List[T any](ctx context.Context, path string, p ParamsContainer, fetch v2Query[T]) *V2List[T] {
	list := &V2List[T]{
		fetch:  fetch,
		params: p,
		v2Page: &V2Page[T]{V2ListMeta: V2ListMeta{NextPageURL: path}},
	}
	list.page(ctx)
	list.initialized = true
	return list
}

// v2Query is a function that fetches a page of items.
type v2Query[T any] func(ctx context.Context, path string, p ParamsContainer) (*V2Page[T], error)

// Fetch is a function that fetches a page of items.
// Deprecated: This type is intended for internal use only, and will be removed in a future version.
type Fetch[T any] func(path string, p ParamsContainer) (*V2Page[T], error)

// NewV2List creates a new V2List with the given path and fetch function.
// Deprecated: This function is intended for internal use only, and will be removed in a future version.
func NewV2List[T any](path string, p ParamsContainer, fetch Fetch[T]) *V2List[T] {
	var ctx context.Context
	if p.GetParams() != nil {
		ctx = p.GetParams().Context
	} else {
		ctx = context.Background()
	}
	v2Query := func(ctx context.Context, path string, p ParamsContainer) (*V2Page[T], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		return fetch(path, p)
	}
	return newV2List(ctx, path, p, v2Query)
}
