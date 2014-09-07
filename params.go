package stripe

import (
	"fmt"
	"net/url"
	"strconv"
)

const (
	startafter = "starting_after"
	endbefore  = "ending_before"
)

// Params is the structure that contains the common properties
// of any *Params structure.
type Params struct {
	Exp  []string
	Meta map[string]string
}

// ListParams is the structure that contains the common properties
// of any *ListParams structure.
type ListParams struct {
	Start, End string
	Limit      int
	Filters    Filters
	// By default, listing through an iterator will automatically grab
	// additional pages as the query progresses. To change this behavior
	// and just load a single page, set this to true.
	Single bool
}

// ListMeta is the structure that contains the common properties
// of List iterators. The Count property is only populated if the
// total_count include option is passed in (see tests for example).
type ListMeta struct {
	Count uint16 `json:"total_count"`
	More  bool   `json:"has_more"`
	Url   string `json:"url"`
}

// Filters is a structure that contains a collection of filters for list-related APIs.
type Filters struct {
	f []*filter
}

// filter is the structure that contains a filter for list-related APIs.
// It ends up passing query string parameters in the format key[op]=value.
type filter struct {
	Key, Op, Val string
}

// Expand appends a new field to expand.
func (p *Params) Expand(f string) {
	p.Exp = append(p.Exp, f)
}

// AddFilter adds a new filter with a given key, op and value.
func (f *Filters) AddFilter(key, op, value string) {
	filter := &filter{Key: key, Op: op, Val: value}
	f.f = append(f.f, filter)
}

// AppendTo adds the common parameters to the query string values.
func (p *Params) AppendTo(body *url.Values) {
	for k, v := range p.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	for _, v := range p.Exp {
		body.Add("expand[]", v)
	}
}

// AppendTo adds the common parameters to the query string values.
func (p *ListParams) AppendTo(body *url.Values) {
	if len(p.Filters.f) > 0 {
		p.Filters.AppendTo(body)
	}

	if len(p.Start) > 0 {
		body.Add(startafter, p.Start)
	}

	if len(p.End) > 0 {
		body.Add(endbefore, p.End)
	}

	if p.Limit > 0 {
		if p.Limit > 100 {
			p.Limit = 100
		}

		body.Add("limit", strconv.Itoa(p.Limit))
	}
}

// AppendTo adds the list of filters to the query string values.
func (f *Filters) AppendTo(values *url.Values) {
	for _, v := range f.f {
		if len(v.Op) > 0 {
			values.Add(fmt.Sprintf("%v[%v]", v.Key, v.Op), v.Val)
		} else {
			values.Add(v.Key, v.Val)
		}
	}
}
