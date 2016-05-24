package stripe

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

const (
	startafter = "starting_after"
	endbefore  = "ending_before"
)

// Form is an abstraction of a type that can have a series of values added to
// it and then subsequently be encoded for use in an HTTP request.
type Form interface {
	Add(key, value string)
	Encode() string
}

// StableForm is a Form implementation that allows duplicate keys and encodes
// its entries in the same order that they were added.
type StableForm struct {
	values []formValue
}

// Add adds a key/value tuple to the form.
func (f *StableForm) Add(key, val string) {
	f.values = append(f.values, formValue{key, val})
}

// Encode encodes the form's values into “URL encoded” form
// ("bar=baz&foo=quux").
func (f *StableForm) Encode() string {
	var buf bytes.Buffer
	for _, v := range f.values {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(url.QueryEscape(v.Key))
		buf.WriteString("=")
		buf.WriteString(url.QueryEscape(v.Value))
	}
	return buf.String()
}

// A key/value tuple for use in the StableForm type.
type formValue struct {
	Key   string
	Value string
}

// Params is the structure that contains the common properties
// of any *Params structure.
type Params struct {
	Exp            []string
	Meta           map[string]string
	Extra          url.Values
	IdempotencyKey string

	// StripeAccount may contain the ID of a connected account. By including
	// this field, the request is made as if it originated from the connected
	// account instead of under the account of the owner of the configured
	// Stripe key.
	StripeAccount string

	// Account is deprecated form of StripeAccount that will do the same thing.
	// Please use StripeAccount instead.
	Account string
}

// ListParams is the structure that contains the common properties
// of any *ListParams structure.
type ListParams struct {
	Exp        []string
	Start, End string
	Limit      int
	Filters    Filters
	// By default, listing through an iterator will automatically grab
	// additional pages as the query progresses. To change this behavior
	// and just load a single page, set this to true.
	Single bool

	// StripeAccount may contain the ID of a connected account. By including
	// this field, the request is made as if it originated from the connected
	// account instead of under the account of the owner of the configured
	// Stripe key.
	StripeAccount string
}

// ListMeta is the structure that contains the common properties
// of List iterators. The Count property is only populated if the
// total_count include option is passed in (see tests for example).
type ListMeta struct {
	Count uint32 `json:"total_count"`
	More  bool   `json:"has_more"`
	URL   string `json:"url"`
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

// NewIdempotencyKey generates a new idempotency key that
// can be used on a request.
func NewIdempotencyKey() string {
	now := time.Now().UnixNano()
	buf := make([]byte, 4)
	rand.Read(buf)
	return fmt.Sprintf("%v_%v", now, base64.URLEncoding.EncodeToString(buf)[:6])
}

// SetAccount sets a value for the Stripe-Account header.
func (p *Params) SetAccount(val string) {
	p.Account = val
	p.StripeAccount = val
}

// SetStripeAccount sets a value for the Stripe-Account header.
func (p *Params) SetStripeAccount(val string) {
	p.StripeAccount = val
}

// Expand appends a new field to expand.
func (p *Params) Expand(f string) {
	p.Exp = append(p.Exp, f)
}

// AddMeta adds a new key-value pair to the Metadata.
func (p *Params) AddMeta(key, value string) {
	if p.Meta == nil {
		p.Meta = make(map[string]string)
	}

	p.Meta[key] = value
}

// AddExtra adds a new arbitrary key-value pair to the request data
func (p *Params) AddExtra(key, value string) {
	if p.Extra == nil {
		p.Extra = make(url.Values)
	}

	p.Extra.Add(key, value)
}

// AddFilter adds a new filter with a given key, op and value.
func (f *Filters) AddFilter(key, op, value string) {
	filter := &filter{Key: key, Op: op, Val: value}
	f.f = append(f.f, filter)
}

// AppendTo adds the common parameters to the query string values.
func (p *Params) AppendTo(body Form) {
	for k, v := range p.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	for _, v := range p.Exp {
		body.Add("expand[]", v)
	}

	for k, vs := range p.Extra {
		for _, v := range vs {
			body.Add(k, v)
		}
	}
}

// Expand appends a new field to expand.
func (p *ListParams) Expand(f string) {
	p.Exp = append(p.Exp, f)
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

	for _, v := range p.Exp {
		body.Add("expand[]", v)
	}
}

// ToParams converts a ListParams to a Params by moving over any fields that
// have valid targets in the new type. This is useful because fields in
// Params can be injected directly into an http.Request while generally
// ListParams is only used to build a set of parameters.
func (p *ListParams) ToParams() *Params {
	return &Params{
		StripeAccount: p.StripeAccount,
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
