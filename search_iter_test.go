package stripe

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/v83/form"
)

var nextPageTestToken = "next_page_test_token"

func TestSearchIterEmpty(t *testing.T) {
	tq := testSearchQuery{{nil, &SearchMeta{}, nil}}
	g, gerr := collect(GetSearchIter(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, 0, len(g))
	assert.NoError(t, gerr)
}

func TestSearchIterEmptyErr(t *testing.T) {
	tq := testSearchQuery{{nil, &SearchMeta{}, errTest}}
	g, gerr := collect(GetSearchIter(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, 0, len(g))
	assert.Equal(t, errTest, gerr)
}

func TestSearchIterOne(t *testing.T) {
	tq := testSearchQuery{{[]interface{}{1}, &SearchMeta{}, nil}}
	want := []interface{}{1}
	g, gerr := collect(GetSearchIter(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestSearchIterOneErr(t *testing.T) {
	tq := testSearchQuery{{[]interface{}{1}, &SearchMeta{}, errTest}}
	want := []interface{}{1}
	g, gerr := collect(GetSearchIter(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestSearchIterPage2Empty(t *testing.T) {
	tq := testSearchQuery{
		{[]interface{}{&item{"x"}}, &SearchMeta{HasMore: true, URL: "", NextPage: &nextPageTestToken}, nil},
		{nil, &SearchMeta{}, nil},
	}
	want := []interface{}{&item{"x"}}
	g, gerr := collect(GetSearchIter(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestSearchIterPage2EmptyErr(t *testing.T) {
	tq := testSearchQuery{
		{[]interface{}{&item{"x"}}, &SearchMeta{HasMore: true, URL: "", NextPage: &nextPageTestToken}, nil},
		{nil, &SearchMeta{}, errTest},
	}
	want := []interface{}{&item{"x"}}
	g, gerr := collect(GetSearchIter(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestSearchIterTwoPages(t *testing.T) {
	tq := testSearchQuery{
		{[]interface{}{&item{"x"}}, &SearchMeta{HasMore: true, URL: "", NextPage: &nextPageTestToken}, nil},
		{[]interface{}{2}, &SearchMeta{HasMore: false, URL: ""}, nil},
	}
	want := []interface{}{&item{"x"}, 2}
	g, gerr := collect(GetSearchIter(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestSearchIterTwoPagesErr(t *testing.T) {
	tq := testSearchQuery{
		{[]interface{}{&item{"x"}}, &SearchMeta{HasMore: true, URL: "", NextPage: &nextPageTestToken}, nil},
		{[]interface{}{2}, &SearchMeta{HasMore: false, URL: ""}, errTest},
	}
	want := []interface{}{&item{"x"}, 2}
	g, gerr := collect(GetSearchIter(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestSearchIterListAndMeta(t *testing.T) {
	type listType struct {
		SearchMeta
	}
	listMeta := &SearchMeta{HasMore: true, URL: "", NextPage: &nextPageTestToken}
	list := &listType{SearchMeta: *listMeta}

	tq := testSearchQuery{{nil, list, nil}}
	it := GetSearchIter(nil, tq.query)
	assert.Equal(t, list, it.SearchResult())
	assert.Equal(t, listMeta, it.Meta())
}

func TestSearchIterMultiplePages(t *testing.T) {
	// Create an ephemeral test server so that we can inspect request attributes.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RawQuery == "query=my+query" {
			w.Write([]byte(`{"data":[{"id": "1"}, {"id":"2"}], "has_more":true, "next_page":"page2", "total_count": 4 }`))
			return
		} else if r.URL.RawQuery == "query=my+query&page=page2" {
			w.Write([]byte(`{"data":[{"id": "3"}, {"id":"4"}], "has_more":false, "next_page":null }`))
			return
		}
		assert.Fail(t, "shouldn't be hit")
	}))
	defer ts.Close()

	// Configure the stripe client to use the ephemeral backend.
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{
		URL: String(ts.URL),
	})
	client := TestClient{B: backend, Key: Key}

	iter := client.Search(&SearchParams{
		Query: "my query",
	})
	assert.Equal(t, *iter.Meta().TotalCount, uint32(4))
	cnt := 0
	for iter.Next() {
		e := iter.Current().(*TestEntity)
		cnt += 1
		assert.Equal(t, fmt.Sprint(cnt), e.ID)
	}

	assert.Equal(t, 4, cnt)
}

func TestV1SearchListEmpty(t *testing.T) {
	tq := testV1SearchQuery[*item]{{v: &V1SearchPage[*item]{}, e: nil}}
	g, gerr := collectSearchList(newV1SearchList(context.Background(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, 0, len(g))
	assert.NoError(t, gerr)
}

func TestV1SearchListEmptyErr(t *testing.T) {
	tq := testV1SearchQuery[*item]{{v: &V1SearchPage[*item]{}, e: errTest}}
	g, gerr := collectSearchList(newV1SearchList(context.Background(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, 0, len(g))
	assert.Equal(t, errTest, gerr)
}

func TestV1SearchListOne(t *testing.T) {
	tq := testV1SearchQuery[*item]{{v: &V1SearchPage[*item]{Data: []*item{{"1"}}}, e: nil}}
	want := []*item{{"1"}}
	g, gerr := collectSearchList(newV1SearchList(context.Background(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestV1SearchListOneErr(t *testing.T) {
	tq := testV1SearchQuery[*item]{{v: &V1SearchPage[*item]{Data: []*item{{"1"}}}, e: errTest}}
	want := []*item{{"1"}}
	g, gerr := collectSearchList(newV1SearchList(context.Background(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestV1SearchListPage2Empty(t *testing.T) {
	tq := testV1SearchQuery[*item]{
		{v: &V1SearchPage[*item]{Data: []*item{{"x"}}, SearchMeta: SearchMeta{HasMore: true, URL: "", NextPage: &nextPageTestToken}}, e: nil},
		{v: &V1SearchPage[*item]{}, e: nil},
	}
	want := []*item{{"x"}}
	g, gerr := collectSearchList(newV1SearchList(context.Background(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestV1SearchListPage2EmptyErr(t *testing.T) {
	tq := testV1SearchQuery[*item]{
		{v: &V1SearchPage[*item]{Data: []*item{{"x"}}, SearchMeta: SearchMeta{HasMore: true, URL: "", NextPage: &nextPageTestToken}}, e: nil},
		{v: &V1SearchPage[*item]{}, e: errTest},
	}
	want := []*item{{"x"}}
	g, gerr := collectSearchList(newV1SearchList(context.Background(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestV1SearchListTwoPages(t *testing.T) {
	tq := testV1SearchQuery[*item]{
		{v: &V1SearchPage[*item]{Data: []*item{{"x"}}, SearchMeta: SearchMeta{HasMore: true, URL: "", NextPage: &nextPageTestToken}}, e: nil},
		{v: &V1SearchPage[*item]{Data: []*item{{"y"}}, SearchMeta: SearchMeta{HasMore: false, URL: ""}}, e: nil},
	}
	want := []*item{{"x"}, {"y"}}
	g, gerr := collectSearchList(newV1SearchList(context.Background(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestV1SearchListTwoPagesErr(t *testing.T) {
	tq := testV1SearchQuery[*item]{
		{v: &V1SearchPage[*item]{Data: []*item{{"x"}}, SearchMeta: SearchMeta{HasMore: true, URL: "", NextPage: &nextPageTestToken}}, e: nil},
		{v: &V1SearchPage[*item]{Data: []*item{{"y"}}, SearchMeta: SearchMeta{HasMore: false, URL: ""}}, e: errTest},
	}
	want := []*item{{"x"}, {"y"}}
	g, gerr := collectSearchList(newV1SearchList(context.Background(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestV1SearchListMultiplePages(t *testing.T) {
	// Create an ephemeral test server so that we can inspect request attributes.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RawQuery == "query=my+query" {
			w.Write([]byte(`{"data":[{"id": "1"}, {"id":"2"}], "has_more":true, "next_page":"page2", "total_count": 4 }`))
			return
		} else if r.URL.RawQuery == "query=my+query&page=page2" {
			w.Write([]byte(`{"data":[{"id": "3"}, {"id":"4"}], "has_more":false, "next_page":null }`))
			return
		}
		assert.Fail(t, "shouldn't be hit")
	}))
	defer ts.Close()

	// Configure the stripe client to use the ephemeral backend.
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{
		URL: String(ts.URL),
	})
	client := TestServer{B: backend, Key: Key}

	iter := client.Search(context.Background(), &SearchParams{
		Query: "my query",
	})

	cnt := 0
	var err error
	iter(func(t *TestEntity, e error) bool {
		if e != nil {
			err = e
			return false
		}
		cnt++
		return true
	})
	assert.NoError(t, err)
	assert.Equal(t, 4, cnt)
}

type testSearchQuery []struct {
	v []interface{}
	m SearchContainer
	e error
}

func (tq *testSearchQuery) query(*Params, *form.Values) ([]interface{}, SearchContainer, error) {
	x := (*tq)[0]
	*tq = (*tq)[1:]
	return x.v, x.m, x.e
}

type testV1SearchQuery[T LastResponseSetter] []struct {
	v *V1SearchPage[T]
	e error
}

func (tq *testV1SearchQuery[T]) query(context.Context, *Params, *form.Values) (*V1SearchPage[T], error) {
	x := (*tq)[0]
	*tq = (*tq)[1:]
	return x.v, x.e
}

// TestClient is used to invoke /charges APIs.
type TestClient struct {
	B   Backend
	Key string
}

// SearchIter is an iterator for charges.
type TestSearchIter struct {
	*SearchIter
}

// Search returns a search result containing charges.
func (c TestClient) Search(params *SearchParams) *TestSearchIter {
	return &TestSearchIter{
		SearchIter: GetSearchIter(params, func(p *Params, b *form.Values) ([]interface{}, SearchContainer, error) {
			list := &TestSearchResult{}
			err := c.B.CallRaw(http.MethodGet, "/v1/something/search", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

type TestServer struct {
	B   Backend
	Key string
}

func (c TestServer) Search(ctx context.Context, params *SearchParams) Seq2[*TestEntity, error] {
	return newV1SearchList(context.Background(), params, func(ctx context.Context, p *Params, b *form.Values) (*V1SearchPage[*TestEntity], error) {
		list := &V1SearchPage[*TestEntity]{}
		err := c.B.CallRaw(http.MethodGet, "/v1/something/search", c.Key, []byte(b.Encode()), p, list)
		ret := make([]*TestEntity, len(list.Data))
		copy(ret, list.Data)
		return list, err
	}).All(ctx)
}

type TestEntity struct {
	APIResource
	// Amount intended to be collected by this payment. A positive integer representing how much to charge in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency). The minimum amount is $0.50 US or [equivalent in charge currency](https://stripe.com/docs/currencies#minimum-and-maximum-charge-amounts). The amount value supports up to eight digits (e.g., a value of 99999999 for a USD charge of $999,999.99).
	ID string `json:"id"`
}

type TestSearchResult struct {
	APIResource
	SearchMeta
	Data []*TestEntity `json:"data"`
}

func collectSearchList[T LastResponseSetter](it *V1SearchList[T]) ([]T, error) {
	var tt []T
	var err error
	it.All(context.TODO())(func(t T, e error) bool {
		if e != nil {
			err = e
			return false
		}
		tt = append(tt, t)
		return true
	})
	return tt, err
}

type testSearchItemWithResponse struct {
	ID           string
	Name         string
	lastResponse *APIResponse
}

func (t *testSearchItemWithResponse) SetLastResponse(response *APIResponse) {
	t.lastResponse = response
}

type testSearchItemSimple struct {
	ID string
}

func (t *testSearchItemSimple) SetLastResponse(response *APIResponse) {}

type simpleSearchItem struct {
	ID string
}

func TestMaybeAddLastResponseSearchIndividualJSON(t *testing.T) {
	// Test that each item gets its corresponding raw JSON from the data array
	pageRawJSON := `{
		"object": "search_result",
		"url": "/v1/customers/search",
		"has_more": false,
		"data": [
			{"id": "cus_1", "name": "Customer 1"},
			{"id": "cus_2", "name": "Customer 2"}
		]
	}`

	item1 := &testSearchItemWithResponse{ID: "cus_1", Name: "Customer 1"}
	item2 := &testSearchItemWithResponse{ID: "cus_2", Name: "Customer 2"}

	page := &V1SearchPage[*testSearchItemWithResponse]{
		APIResource: APIResource{
			LastResponse: &APIResponse{
				RawJSON: []byte(pageRawJSON),
			},
		},
		Data: []*testSearchItemWithResponse{item1, item2},
	}

	// Call the function
	err := maybeAddLastResponseSearch(page)
	assert.NoError(t, err)

	// Verify each item has its corresponding JSON
	expectedJSON1 := `{"id": "cus_1", "name": "Customer 1"}`
	expectedJSON2 := `{"id": "cus_2", "name": "Customer 2"}`

	assert.NotNil(t, item1.lastResponse)
	assert.JSONEq(t, expectedJSON1, string(item1.lastResponse.RawJSON))

	assert.NotNil(t, item2.lastResponse)
	assert.JSONEq(t, expectedJSON2, string(item2.lastResponse.RawJSON))

	// Verify other fields are copied from the original response
	assert.Equal(t, page.LastResponse.Header, item1.lastResponse.Header)
	assert.Equal(t, page.LastResponse.IdempotencyKey, item1.lastResponse.IdempotencyKey)
	assert.Equal(t, page.LastResponse.RequestID, item1.lastResponse.RequestID)
	assert.Equal(t, page.LastResponse.Status, item1.lastResponse.Status)
	assert.Equal(t, page.LastResponse.StatusCode, item1.lastResponse.StatusCode)
}

func TestMaybeAddLastResponseSearchMismatchedLengths(t *testing.T) {
	// Test error when data array length doesn't match page.Data length
	pageRawJSON := `{
		"object": "search_result",
		"data": [
			{"id": "cus_1"}
		]
	}`

	page := &V1SearchPage[*testSearchItemSimple]{
		APIResource: APIResource{
			LastResponse: &APIResponse{
				RawJSON:   []byte(pageRawJSON),
				RequestID: "req_search_test123",
			},
		},
		Data: []*testSearchItemSimple{{"cus_1"}, {"cus_2"}}, // 2 items but only 1 in JSON data array
	}

	err := maybeAddLastResponseSearch(page)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "mismatch in data length for requestID req_search_test123")
}

func TestMaybeAddLastResponseSearchInvalidJSON(t *testing.T) {
	// Test error when page JSON is invalid
	pageRawJSON := `{invalid json`

	page := &V1SearchPage[*testSearchItemSimple]{
		APIResource: APIResource{
			LastResponse: &APIResponse{
				RawJSON: []byte(pageRawJSON),
			},
		},
		Data: []*testSearchItemSimple{{"cus_1"}},
	}

	err := maybeAddLastResponseSearch(page)
	assert.Error(t, err)
}

func TestMaybeAddLastResponseSearchWithNonLastResponseSetter(t *testing.T) {
	// Test with items that don't implement LastResponseSetter
	pageRawJSON := `{
		"object": "search_result",
		"data": [
			{"id": "item_1"},
			{"id": "item_2"}
		]
	}`

	// Note: simpleSearchItem does NOT implement LastResponseSetter
	page := &V1SearchPage[*simpleSearchItem]{
		APIResource: APIResource{
			LastResponse: &APIResponse{
				RawJSON: []byte(pageRawJSON),
			},
		},
		Data: []*simpleSearchItem{{"item_1"}, {"item_2"}},
	}

	// Should not error even though items don't implement LastResponseSetter
	err := maybeAddLastResponseSearch(page)
	assert.NoError(t, err)
}

func TestMaybeAddLastResponseSearchNilLastResponse(t *testing.T) {
	// Test when LastResponse is nil - should return nil without error
	page := &V1SearchPage[*testSearchItemSimple]{
		APIResource: APIResource{
			LastResponse: nil, // nil LastResponse
		},
		Data: []*testSearchItemSimple{{"cus_1"}},
	}

	err := maybeAddLastResponseSearch(page)
	assert.NoError(t, err)
}
