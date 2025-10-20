package stripe

import (
	"context"
	"errors"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/v83/form"
)

func TestIterEmpty(t *testing.T) {
	tq := testQuery{{nil, &ListMeta{}, nil}}
	g, gerr := collect(GetIter(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, 0, len(g))
	assert.NoError(t, gerr)
}

func TestIterEmptyErr(t *testing.T) {
	tq := testQuery{{nil, &ListMeta{}, errTest}}
	g, gerr := collect(GetIter(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, 0, len(g))
	assert.Equal(t, errTest, gerr)
}

func TestIterOne(t *testing.T) {
	tq := testQuery{{[]interface{}{1}, &ListMeta{}, nil}}
	want := []interface{}{1}
	g, gerr := collect(GetIter(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestIterOneErr(t *testing.T) {
	tq := testQuery{{[]interface{}{1}, &ListMeta{}, errTest}}
	want := []interface{}{1}
	g, gerr := collect(GetIter(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestIterPage2Empty(t *testing.T) {
	tq := testQuery{
		{[]interface{}{&item{"x"}}, &ListMeta{HasMore: true, TotalCount: 0, URL: ""}, nil},
		{nil, &ListMeta{}, nil},
	}
	want := []interface{}{&item{"x"}}
	g, gerr := collect(GetIter(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestIterPage2EmptyErr(t *testing.T) {
	tq := testQuery{
		{[]interface{}{&item{"x"}}, &ListMeta{HasMore: true, TotalCount: 0, URL: ""}, nil},
		{nil, &ListMeta{}, errTest},
	}
	want := []interface{}{&item{"x"}}
	g, gerr := collect(GetIter(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestIterTwoPages(t *testing.T) {
	tq := testQuery{
		{[]interface{}{&item{"x"}}, &ListMeta{HasMore: true, TotalCount: 0, URL: ""}, nil},
		{[]interface{}{2}, &ListMeta{HasMore: false, TotalCount: 0, URL: ""}, nil},
	}
	want := []interface{}{&item{"x"}, 2}
	g, gerr := collect(GetIter(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestIterTwoPagesErr(t *testing.T) {
	tq := testQuery{
		{[]interface{}{&item{"x"}}, &ListMeta{HasMore: true, TotalCount: 0, URL: ""}, nil},
		{[]interface{}{2}, &ListMeta{HasMore: false, TotalCount: 0, URL: ""}, errTest},
	}
	want := []interface{}{&item{"x"}, 2}
	g, gerr := collect(GetIter(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestIterReversed(t *testing.T) {
	tq := testQuery{{[]interface{}{1, 2}, &ListMeta{}, nil}}
	want := []interface{}{2, 1}
	g, gerr := collect(GetIter(&ListParams{EndingBefore: String("x")}, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestIterReversedTwoPages(t *testing.T) {
	tq := testQuery{
		{[]interface{}{&item{"3"}, 4}, &ListMeta{HasMore: true, TotalCount: 0, URL: ""}, nil},
		{[]interface{}{1, 2}, &ListMeta{}, nil},
	}
	want := []interface{}{4, &item{"3"}, 2, 1}
	g, gerr := collect(GetIter(&ListParams{EndingBefore: String("x")}, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestReverse(t *testing.T) {
	var cases = [][]interface{}{
		{},
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4},
	}
	for _, a := range cases {
		b := make([]interface{}, len(a))
		copy(b, a)
		reverse(b)
		for i, g := range b {
			want := a[len(a)-1-i]
			assert.Equal(t, want, g)
		}
	}
}

func TestIterListAndMeta(t *testing.T) {
	type listType struct {
		ListMeta
	}
	listMeta := &ListMeta{HasMore: true, TotalCount: 0, URL: ""}
	list := &listType{ListMeta: *listMeta}

	tq := testQuery{{nil, list, nil}}
	it := GetIter(nil, tq.query)
	assert.Equal(t, list, it.List())
	assert.Equal(t, listMeta, it.Meta())
}

func TestV1ListEmpty(t *testing.T) {
	tq := testV1Query[*item]{{v: &V1Page[*item]{}, e: nil}}
	g, gerr := collectList(newV1List(context.TODO(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, 0, len(g))
	assert.NoError(t, gerr)
}

func TestV1ListEmptyErr(t *testing.T) {
	tq := testV1Query[*item]{{v: &V1Page[*item]{}, e: errTest}}
	g, gerr := collectList(newV1List(context.TODO(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, 0, len(g))
	assert.Equal(t, errTest, gerr)
}

func TestV1ListOne(t *testing.T) {
	tq := testV1Query[*item]{{v: &V1Page[*item]{Data: []*item{{"1"}}}, e: nil}}
	want := []*item{{"1"}}
	g, gerr := collectList(newV1List(context.TODO(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestV1ListOneErr(t *testing.T) {
	tq := testV1Query[*item]{{v: &V1Page[*item]{Data: []*item{{"1"}}}, e: errTest}}
	want := []*item{{"1"}}
	g, gerr := collectList(newV1List(context.TODO(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestV1ListPage2EmptyErr(t *testing.T) {
	tq := testV1Query[*item]{
		{v: &V1Page[*item]{Data: []*item{{"x"}}, ListMeta: ListMeta{HasMore: true}}, e: nil},
		{v: &V1Page[*item]{}, e: errTest},
	}
	want := []*item{{"x"}}
	g, gerr := collectList(newV1List(context.TODO(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestV1ListTwoPages(t *testing.T) {
	tq := testV1Query[*item]{
		{v: &V1Page[*item]{Data: []*item{{"x"}}, ListMeta: ListMeta{HasMore: true}}, e: nil},
		{v: &V1Page[*item]{Data: []*item{{"y"}}}, e: nil},
	}
	want := []*item{{"x"}, {"y"}}
	g, gerr := collectList(newV1List(context.TODO(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestV1ListTwoPagesErr(t *testing.T) {
	tq := testV1Query[*item]{
		{v: &V1Page[*item]{Data: []*item{{"x"}}, ListMeta: ListMeta{HasMore: true}}, e: nil},
		{v: &V1Page[*item]{Data: []*item{{"y"}}}, e: errTest},
	}
	want := []*item{{"x"}, {"y"}}
	g, gerr := collectList(newV1List(context.TODO(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestV1ListReversed(t *testing.T) {
	tq := testV1Query[*item]{{v: &V1Page[*item]{Data: []*item{{"1"}, {"2"}}}, e: nil}}
	want := []*item{{"2"}, {"1"}}
	g, gerr := collectList(newV1List(context.TODO(), &ListParams{EndingBefore: String("x")}, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestV1ListReversedTwoPages(t *testing.T) {
	tq := testV1Query[*item]{
		{v: &V1Page[*item]{Data: []*item{{"3"}, {"4"}}, ListMeta: ListMeta{HasMore: true}}, e: nil},
		{v: &V1Page[*item]{Data: []*item{{"1"}, {"2"}}}, e: nil},
	}
	want := []*item{{"4"}, {"3"}, {"2"}, {"1"}}
	g, gerr := collectList(newV1List(context.TODO(), &ListParams{EndingBefore: String("x")}, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestV2ListEmpty(t *testing.T) {
	tq := testV2Query[*item]{{v: &V2Page[*item]{}, e: nil}}
	g, gerr := collectV2List(newV2List(context.TODO(), "/test", nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, 0, len(g))
	assert.NoError(t, gerr)
}

func TestV2ListEmptyErr(t *testing.T) {
	tq := testV2Query[*item]{{v: &V2Page[*item]{}, e: errTest}}
	g, gerr := collectV2List(newV2List(context.TODO(), "/test", nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, 0, len(g))
	assert.Equal(t, errTest, gerr)
}

func TestV2ListOne(t *testing.T) {
	tq := testV2Query[*item]{{v: &V2Page[*item]{Data: []*item{{"1"}}}, e: nil}}
	want := []*item{{"1"}}
	g, gerr := collectV2List(newV2List(context.TODO(), "/test", nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestV2ListOneErr(t *testing.T) {
	tq := testV2Query[*item]{{v: &V2Page[*item]{Data: []*item{{"1"}}}, e: errTest}}
	want := []*item{{"1"}}
	g, gerr := collectV2List(newV2List(context.TODO(), "/test", nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestV2ListPage2EmptyErr(t *testing.T) {
	tq := testV2Query[*item]{
		{v: &V2Page[*item]{Data: []*item{{"x"}}, NextPageURL: "/test?page=2"}, e: nil},
		{v: &V2Page[*item]{}, e: errTest},
	}
	want := []*item{{"x"}}
	g, gerr := collectV2List(newV2List(context.TODO(), "/test", nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestV2ListTwoPages(t *testing.T) {
	tq := testV2Query[*item]{
		{v: &V2Page[*item]{Data: []*item{{"x"}}, NextPageURL: "/test?page=2"}, e: nil},
		{v: &V2Page[*item]{Data: []*item{{"y"}}}, e: nil},
	}
	want := []*item{{"x"}, {"y"}}
	g, gerr := collectV2List(newV2List(context.TODO(), "/test", nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestV2ListTwoPagesErr(t *testing.T) {
	tq := testV2Query[*item]{
		{v: &V2Page[*item]{Data: []*item{{"x"}}, NextPageURL: "/test?page=2"}, e: nil},
		{v: &V2Page[*item]{Data: []*item{{"y"}}}, e: errTest},
	}
	want := []*item{{"x"}, {"y"}}
	g, gerr := collectV2List(newV2List(context.TODO(), "/test", nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

//
// ---
//

var errTest = errors.New("test error")

type item struct {
	ID string
}

func (i *item) SetLastResponse(response *APIResponse) {}

type testQuery []struct {
	v []interface{}
	m ListContainer
	e error
}

func (tq *testQuery) query(*Params, *form.Values) ([]interface{}, ListContainer, error) {
	x := (*tq)[0]
	*tq = (*tq)[1:]
	return x.v, x.m, x.e
}

type testV1Query[T LastResponseSetter] []struct {
	v *V1Page[T]
	e error
}

func (tq *testV1Query[T]) query(context.Context, *Params, *form.Values) (*V1Page[T], error) {
	x := (*tq)[0]
	*tq = (*tq)[1:]
	return x.v, x.e
}

type testV2Query[T any] []struct {
	v *V2Page[T]
	e error
}

func (tq *testV2Query[T]) query(context.Context, string, ParamsContainer) (*V2Page[T], error) {
	x := (*tq)[0]
	*tq = (*tq)[1:]
	return x.v, x.e
}

func collectList[T LastResponseSetter](it *V1List[T]) ([]T, error) {
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

func collectV2List[T any](it *V2List[T]) ([]T, error) {
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

type collectable interface {
	Next() bool
	Current() interface{}
	Err() error
}

func collect(it collectable) ([]interface{}, error) {
	var g []interface{}
	for it.Next() {
		g = append(g, it.Current())
	}
	return g, it.Err()
}

type testItemWithResponse struct {
	ID           string
	Name         string
	lastResponse *APIResponse
}

func (t *testItemWithResponse) SetLastResponse(response *APIResponse) {
	t.lastResponse = response
}

type testItemSimple struct {
	ID string
}

func (t *testItemSimple) SetLastResponse(response *APIResponse) {}

type simpleItem struct {
	ID string
}

func TestMaybeAddLastResponseIndividualJSON(t *testing.T) {
	// Test that each item gets its corresponding raw JSON from the data array
	pageRawJSON := `{
		"object": "list",
		"url": "/v1/customers",
		"has_more": false,
		"data": [
			{"id": "cus_1", "name": "Customer 1"},
			{"id": "cus_2", "name": "Customer 2"}
		]
	}`

	item1 := &testItemWithResponse{ID: "cus_1", Name: "Customer 1"}
	item2 := &testItemWithResponse{ID: "cus_2", Name: "Customer 2"}

	page := &V1Page[*testItemWithResponse]{
		APIResource: APIResource{
			LastResponse: &APIResponse{
				RawJSON: []byte(pageRawJSON),
			},
		},
		Data: []*testItemWithResponse{item1, item2},
	}

	// Call the function
	err := maybeAddLastResponseV1(page)
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

func TestMaybeAddLastResponseMismatchedLengths(t *testing.T) {
	// Test error when data array length doesn't match page.Data length
	pageRawJSON := `{
		"object": "list",
		"data": [
			{"id": "cus_1"}
		]
	}`

	page := &V1Page[*testItemSimple]{
		APIResource: APIResource{
			LastResponse: &APIResponse{
				RawJSON:   []byte(pageRawJSON),
				RequestID: "req_test123",
			},
		},
		Data: []*testItemSimple{{"cus_1"}, {"cus_2"}}, // 2 items but only 1 in JSON data array
	}

	err := maybeAddLastResponseV1(page)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "mismatch in data length for requestID req_test123")
}

func TestMaybeAddLastResponseInvalidJSON(t *testing.T) {
	// Test error when page JSON is invalid
	pageRawJSON := `{invalid json`

	page := &V1Page[*testItemSimple]{
		APIResource: APIResource{
			LastResponse: &APIResponse{
				RawJSON: []byte(pageRawJSON),
			},
		},
		Data: []*testItemSimple{{"cus_1"}},
	}

	err := maybeAddLastResponseV1(page)
	assert.Error(t, err)
}

func TestMaybeAddLastResponseWithNonLastResponseSetter(t *testing.T) {
	// Test with items that don't implement LastResponseSetter
	pageRawJSON := `{
		"object": "list",
		"data": [
			{"id": "item_1"},
			{"id": "item_2"}
		]
	}`

	// Note: simpleItem does NOT implement LastResponseSetter
	page := &V1Page[*simpleItem]{
		APIResource: APIResource{
			LastResponse: &APIResponse{
				RawJSON: []byte(pageRawJSON),
			},
		},
		Data: []*simpleItem{{"item_1"}, {"item_2"}},
	}

	// Should not error even though items don't implement LastResponseSetter
	err := maybeAddLastResponseV1(page)
	assert.NoError(t, err)
}
