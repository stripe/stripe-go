package stripe

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/v87/form"
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

func TestIterEmptyWithHasMore(t *testing.T) {
	tq := testQuery{{nil, &ListMeta{HasMore: true}, nil}}
	g, gerr := collect(GetIter(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, 0, len(g))
	assert.ErrorIs(t, gerr, errEmptyPageWithHasMore)
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

func TestIterPage2EmptyWithHasMore(t *testing.T) {
	tq := testQuery{
		{[]interface{}{&item{"x"}}, &ListMeta{HasMore: true}, nil},
		{nil, &ListMeta{HasMore: true}, nil},
	}
	want := []interface{}{&item{"x"}}
	g, gerr := collect(GetIter(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.ErrorIs(t, gerr, errEmptyPageWithHasMore)
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
	tq := testV1Query[*item]{{v: &v1Page[*item]{}, e: nil}}
	g, gerr := collectList(newV1List(context.TODO(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, 0, len(g))
	assert.NoError(t, gerr)
}

func TestV1ListEmptyErr(t *testing.T) {
	tq := testV1Query[*item]{{v: &v1Page[*item]{}, e: errTest}}
	g, gerr := collectList(newV1List(context.TODO(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, 0, len(g))
	assert.Equal(t, errTest, gerr)
}

func TestV1ListEmptyWithHasMore(t *testing.T) {
	tq := testV1Query[*item]{{v: &v1Page[*item]{ListMeta: ListMeta{HasMore: true}}, e: nil}}
	g, gerr := collectList(newV1List(context.TODO(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, 0, len(g))
	assert.ErrorIs(t, gerr, errEmptyPageWithHasMore)
}

func TestV1ListOne(t *testing.T) {
	tq := testV1Query[*item]{{v: &v1Page[*item]{Data: []*item{{"1"}}}, e: nil}}
	want := []*item{{"1"}}
	g, gerr := collectList(newV1List(context.TODO(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestV1ListOneErr(t *testing.T) {
	tq := testV1Query[*item]{{v: &v1Page[*item]{Data: []*item{{"1"}}}, e: errTest}}
	want := []*item{{"1"}}
	g, gerr := collectList(newV1List(context.TODO(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestV1ListPage2EmptyErr(t *testing.T) {
	tq := testV1Query[*item]{
		{v: &v1Page[*item]{Data: []*item{{"x"}}, ListMeta: ListMeta{HasMore: true}}, e: nil},
		{v: &v1Page[*item]{}, e: errTest},
	}
	want := []*item{{"x"}}
	g, gerr := collectList(newV1List(context.TODO(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestV1ListPage2EmptyWithHasMore(t *testing.T) {
	tq := testV1Query[*item]{
		{v: &v1Page[*item]{Data: []*item{{"x"}}, ListMeta: ListMeta{HasMore: true}}, e: nil},
		{v: &v1Page[*item]{ListMeta: ListMeta{HasMore: true}}, e: nil},
	}
	want := []*item{{"x"}}
	g, gerr := collectList(newV1List(context.TODO(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.ErrorIs(t, gerr, errEmptyPageWithHasMore)
}

func TestV1ListTwoPages(t *testing.T) {
	tq := testV1Query[*item]{
		{v: &v1Page[*item]{Data: []*item{{"x"}}, ListMeta: ListMeta{HasMore: true}}, e: nil},
		{v: &v1Page[*item]{Data: []*item{{"y"}}}, e: nil},
	}
	want := []*item{{"x"}, {"y"}}
	g, gerr := collectList(newV1List(context.TODO(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestV1ListTwoPagesErr(t *testing.T) {
	tq := testV1Query[*item]{
		{v: &v1Page[*item]{Data: []*item{{"x"}}, ListMeta: ListMeta{HasMore: true}}, e: nil},
		{v: &v1Page[*item]{Data: []*item{{"y"}}}, e: errTest},
	}
	want := []*item{{"x"}, {"y"}}
	g, gerr := collectList(newV1List(context.TODO(), nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestV1ListReversed(t *testing.T) {
	tq := testV1Query[*item]{{v: &v1Page[*item]{Data: []*item{{"1"}, {"2"}}}, e: nil}}
	want := []*item{{"2"}, {"1"}}
	g, gerr := collectList(newV1List(context.TODO(), &ListParams{EndingBefore: String("x")}, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestV1ListReversedTwoPages(t *testing.T) {
	tq := testV1Query[*item]{
		{v: &v1Page[*item]{Data: []*item{{"3"}, {"4"}}, ListMeta: ListMeta{HasMore: true}}, e: nil},
		{v: &v1Page[*item]{Data: []*item{{"1"}, {"2"}}}, e: nil},
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
		{v: &V2Page[*item]{Data: []*item{{"x"}}, V2ListMeta: V2ListMeta{NextPageURL: "/test?page=2"}}, e: nil},
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
		{v: &V2Page[*item]{Data: []*item{{"x"}}, V2ListMeta: V2ListMeta{NextPageURL: "/test?page=2"}}, e: nil},
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
		{v: &V2Page[*item]{Data: []*item{{"x"}}, V2ListMeta: V2ListMeta{NextPageURL: "/test?page=2"}}, e: nil},
		{v: &V2Page[*item]{Data: []*item{{"y"}}}, e: errTest},
	}
	want := []*item{{"x"}, {"y"}}
	g, gerr := collectV2List(newV2List(context.TODO(), "/test", nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestV2ListJSONRoundTrip(t *testing.T) {
	type resource struct {
		Items *V2List[*item] `json:"items,omitempty"`
	}
	input := `{"items":{"data":[{"ID":"first"}],"next_page_url":"/items?page=2","previous_page_url":""}}`
	var resourceWithItems resource
	assert.NoError(t, json.Unmarshal([]byte(input), &resourceWithItems))
	assert.Equal(t, []*item{{"first"}}, resourceWithItems.Items.Data())
	assert.Equal(t, "/items?page=2", resourceWithItems.Items.Meta().NextPageURL)

	encoded, err := json.Marshal(resourceWithItems)
	assert.NoError(t, err)
	assert.JSONEq(t, input, string(encoded))
}

func TestV2ListJSONAbsentNullAndEmpty(t *testing.T) {
	type resource struct {
		Items *V2List[*item] `json:"items"`
	}

	var absent resource
	assert.NoError(t, json.Unmarshal([]byte(`{}`), &absent))
	assert.Nil(t, absent.Items)

	var null resource
	assert.NoError(t, json.Unmarshal([]byte(`{"items":null}`), &null))
	assert.Nil(t, null.Items)

	var empty resource
	assert.NoError(t, json.Unmarshal([]byte(`{"items":{}}`), &empty))
	assert.NotNil(t, empty.Items)
	assert.Nil(t, empty.Items.Data())
	assert.Equal(t, V2ListMeta{}, empty.Items.Meta())
}

func TestV2ListZeroValueMarshal(t *testing.T) {
	var list V2List[*item]
	encoded, err := json.Marshal(list)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"data":null,"next_page_url":"","previous_page_url":""}`, string(encoded))
}

func TestV2ListIncludedPagePagination(t *testing.T) {
	type contextKey string
	ctx := context.WithValue(context.Background(), contextKey("key"), "value")
	backend := &v2ListTestBackend{
		responses: map[string]string{
			"/items?page=2": `{"data":[{"ID":"second"}],"next_page_url":"/items?page=3"}`,
			"/items?page=3": `{"data":[{"ID":"third"}]}`,
		},
	}
	var list V2List[*item]
	assert.NoError(t, json.Unmarshal([]byte(`{"data":[{"ID":"first"}],"next_page_url":"/items?page=2"}`), &list))
	InitializeV2Lists(&list, backend, "sk_test_custom")

	assert.Equal(t, []*item{{"first"}}, list.Data())
	assert.Empty(t, backend.paths)
	assert.Equal(t, []*item{{"first"}, {"second"}, {"third"}}, collectV2ListValues(t, &list, ctx))
	assert.Equal(t, []string{"/items?page=2", "/items?page=3"}, backend.paths)
	assert.Equal(t, []string{"sk_test_custom", "sk_test_custom"}, backend.keys)
	assert.Same(t, ctx, backend.contexts[0])
	assert.Same(t, ctx, backend.contexts[1])
}

func TestV2ListInitializesNestedListsOnLaterPages(t *testing.T) {
	type parentItem struct {
		ID       string         `json:"id"`
		Children *V2List[*item] `json:"children"`
	}
	backend := &v2ListTestBackend{
		responses: map[string]string{
			"/parents?page=2":  `{"data":[{"id":"parent","children":{"data":[{"ID":"child-1"}],"next_page_url":"/children?page=2"}}]}`,
			"/children?page=2": `{"data":[{"ID":"child-2"}]}`,
		},
	}
	var list V2List[*parentItem]
	assert.NoError(t, json.Unmarshal([]byte(`{"data":[],"next_page_url":"/parents?page=2"}`), &list))
	InitializeV2Lists(&list, backend, "key")

	parents := collectV2ListValues(t, &list, context.Background())
	assert.Len(t, parents, 1)
	assert.Equal(t, []*item{{"child-1"}, {"child-2"}}, collectV2ListValues(t, parents[0].Children, context.Background()))
	assert.Equal(t, []string{"/parents?page=2", "/children?page=2"}, backend.paths)
}

func TestV2ListMissingInitializerError(t *testing.T) {
	var list V2List[*item]
	assert.NoError(t, json.Unmarshal([]byte(`{"data":[{"ID":"first"}],"next_page_url":"/items?page=2"}`), &list))
	values, err := collectV2List(&list)
	assert.Equal(t, []*item{{"first"}}, values)
	assert.EqualError(t, err, "stripe: V2List cannot fetch its next page because it is not associated with a Stripe client")
}

func TestV2ListPreservesBackendError(t *testing.T) {
	backend := &v2ListTestBackend{
		responses: map[string]string{"/items?page=2": `{"data":[{"ID":"partial"}]}`},
		errors:    map[string]error{"/items?page=2": errTest},
	}
	var list V2List[*item]
	assert.NoError(t, json.Unmarshal([]byte(`{"data":[],"next_page_url":"/items?page=2"}`), &list))
	InitializeV2Lists(&list, backend, "key")

	values, err := collectV2List(&list)
	assert.Equal(t, []*item{{"partial"}}, values)
	assert.ErrorIs(t, err, errTest)
}

func TestInitializeV2ListsHandlesCyclesAndInaccessibleFields(t *testing.T) {
	type resource struct {
		Next   *resource
		Items  *V2List[*item]
		hidden *V2List[*item]
	}
	backend := &v2ListTestBackend{responses: map[string]string{"/items?page=2": `{"data":[]}`}}
	value := &resource{Items: &V2List[*item]{}, hidden: &V2List[*item]{}}
	value.Next = value
	assert.NoError(t, json.Unmarshal([]byte(`{"data":[],"next_page_url":"/items?page=2"}`), value.Items))

	InitializeV2Lists(value, backend, "key")
	assert.Empty(t, collectV2ListValues(t, value.Items, context.Background()))
	assert.Equal(t, []string{"/items?page=2"}, backend.paths)
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
	v *v1Page[T]
	e error
}

func (tq *testV1Query[T]) query(context.Context, *Params, *form.Values) (*v1Page[T], error) {
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

type v2ListTestBackend struct {
	Backend
	responses map[string]string
	errors    map[string]error
	paths     []string
	keys      []string
	contexts  []context.Context
}

func (b *v2ListTestBackend) Call(_ string, path, key string, params ParamsContainer, v LastResponseSetter) error {
	b.paths = append(b.paths, path)
	b.keys = append(b.keys, key)
	b.contexts = append(b.contexts, params.GetParams().Context)
	if response, ok := b.responses[path]; ok {
		if err := json.Unmarshal([]byte(response), v); err != nil {
			return err
		}
	}
	return b.errors[path]
}

func collectV2ListValues[T any](t *testing.T, list *V2List[T], ctx context.Context) []T {
	t.Helper()
	var values []T
	list.All(ctx)(func(value T, err error) bool {
		assert.NoError(t, err)
		values = append(values, value)
		return true
	})
	return values
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

	page := &v1Page[*testItemWithResponse]{
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

	page := &v1Page[*testItemSimple]{
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

	page := &v1Page[*testItemSimple]{
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
	page := &v1Page[*simpleItem]{
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
