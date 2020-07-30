package stripe

import (
	"errors"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/v71/form"
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

//
// ---
//

var errTest = errors.New("test error")

type item struct {
	ID string
}

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

func collect(it *Iter) ([]interface{}, error) {
	var g []interface{}
	for it.Next() {
		g = append(g, it.Current())
	}
	return g, it.Err()
}
