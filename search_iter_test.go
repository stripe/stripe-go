package stripe

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/v72/form"
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
