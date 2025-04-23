package stripe

import (
	"errors"
	"testing"

	"github.com/max-cape/stripe-go-test/form"
	assert "github.com/stretchr/testify/require"
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
	tq := testV1Query[*int]{{nil, &ListMeta{}, nil}}
	g, gerr := collectList(newV1List(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, 0, len(g))
	assert.NoError(t, gerr)
}

func TestV1ListEmptyErr(t *testing.T) {
	tq := testV1Query[*int]{{nil, &ListMeta{}, errTest}}
	g, gerr := collectList(newV1List(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, 0, len(g))
	assert.Equal(t, errTest, gerr)
}

func TestV1ListOne(t *testing.T) {
	tq := testV1Query[*int]{{[]*int{intPtr(1)}, &ListMeta{}, nil}}
	want := []*int{intPtr(1)}
	g, gerr := collectList(newV1List(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestV1ListOneErr(t *testing.T) {
	tq := testV1Query[*int]{{[]*int{intPtr(1)}, &ListMeta{}, errTest}}
	want := []*int{intPtr(1)}
	g, gerr := collectList(newV1List(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestV1ListPage2EmptyErr(t *testing.T) {
	tq := testV1Query[*item]{
		{[]*item{{"x"}}, &ListMeta{HasMore: true, TotalCount: 0, URL: ""}, nil},
		{nil, &ListMeta{}, errTest},
	}
	want := []*item{{"x"}}
	g, gerr := collectList(newV1List(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestV1ListTwoPages(t *testing.T) {
	tq := testV1Query[*item]{
		{[]*item{{"x"}}, &ListMeta{HasMore: true, TotalCount: 0, URL: ""}, nil},
		{[]*item{{"y"}}, &ListMeta{}, nil},
	}
	want := []*item{{"x"}, {"y"}}
	g, gerr := collectList(newV1List(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestV1ListTwoPagesErr(t *testing.T) {
	tq := testV1Query[*item]{
		{[]*item{{"x"}}, &ListMeta{HasMore: true, TotalCount: 0, URL: ""}, nil},
		{[]*item{{"y"}}, &ListMeta{}, errTest},
	}
	want := []*item{{"x"}, {"y"}}
	g, gerr := collectList(newV1List(nil, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.Equal(t, errTest, gerr)
}

func TestV1ListReversed(t *testing.T) {
	tq := testV1Query[*int]{{[]*int{intPtr(1), intPtr(2)}, &ListMeta{}, nil}}
	want := []*int{intPtr(2), intPtr(1)}
	g, gerr := collectList(newV1List(&ListParams{EndingBefore: String("x")}, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
}

func TestV1ListReversedTwoPages(t *testing.T) {
	tq := testV1Query[*item]{
		{[]*item{{"3"}, {"4"}}, &ListMeta{HasMore: true, TotalCount: 0, URL: ""}, nil},
		{[]*item{{"1"}, {"2"}}, &ListMeta{}, nil},
	}
	want := []*item{{"4"}, {"3"}, {"2"}, {"1"}}
	g, gerr := collectList(newV1List(&ListParams{EndingBefore: String("x")}, tq.query))
	assert.Equal(t, 0, len(tq))
	assert.Equal(t, want, g)
	assert.NoError(t, gerr)
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

type testV1Query[T any] []struct {
	v []T
	m ListContainer
	e error
}

func (tq *testV1Query[T]) query(*Params, *form.Values) ([]T, ListContainer, error) {
	x := (*tq)[0]
	*tq = (*tq)[1:]
	return x.v, x.m, x.e
}

func collectList[T any](it *v1List[T]) ([]*T, error) {
	var tt []*T
	var err error
	it.All()(func(t *T, e error) bool {
		if e != nil {
			err = e
			return false
		}
		tt = append(tt, t)
		return true
	})
	return tt, err
}

func intPtr(i int) *int {
	intPtr := new(int)
	*intPtr = i
	return intPtr
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
