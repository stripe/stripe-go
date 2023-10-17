package stripe

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/v76/form"
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
	client := Client{B: backend, Key: Key}

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

// Client is used to invoke /charges APIs.
type Client struct {
	B   Backend
	Key string
}

// SearchIter is an iterator for charges.
type TestSearchIter struct {
	*SearchIter
}

// Search returns a search result containing charges.
func (c Client) Search(params *SearchParams) *TestSearchIter {
	return &TestSearchIter{
		SearchIter: GetSearchIter(params, func(p *Params, b *form.Values) ([]interface{}, SearchContainer, error) {
			list := &TestSearchResult{}
			err := c.B.CallRaw(http.MethodGet, "/v1/something/search", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
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
