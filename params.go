package stripe

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/stripe/stripe-go/form"
)

const (
	startafter = "starting_after"
	endbefore  = "ending_before"
)

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

	// Headers may be used to provide extra header lines on the HTTP request.
	Headers http.Header
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

// RangeQueryParams are a set of generic request parameters that are used on
// list endpoints to filter their results by some timestamp.
type RangeQueryParams struct {
	// GreaterThan specifies that values should be a greater than this
	// timestamp.
	GreaterThan int64 `form:"gt"`

	// GreaterThanOrEqual specifies that values should be greater than or equal
	// to this timestamp.
	GreaterThanOrEqual int64 `form:"gte"`

	// LesserThan specifies that values should be lesser than this timetamp.
	LesserThan int64 `form:"lt"`

	// LesserThanOrEqual specifies that values should be lesser than or
	// equalthis timetamp.
	LesserThanOrEqual int64 `form:"lte"`
}

// AppendTo adds the range query parametes to a set of request values.
func (r *RangeQueryParams) AppendTo(values *form.Values, name string) {
	if r.GreaterThan > 0 {
		values.Add(name+"[gt]", strconv.FormatInt(r.GreaterThan, 10))
	}

	if r.GreaterThanOrEqual > 0 {
		values.Add(name+"[gte]", strconv.FormatInt(r.GreaterThanOrEqual, 10))
	}

	if r.LesserThan > 0 {
		values.Add(name+"[lt]", strconv.FormatInt(r.LesserThan, 10))
	}

	if r.LesserThanOrEqual > 0 {
		values.Add(name+"[lte]", strconv.FormatInt(r.LesserThanOrEqual, 10))
	}
}

// Filters is a structure that contains a collection of filters for list-related APIs.
type Filters struct {
	f []*filter
}

// AppendTo adds the list of filters to the query string values.
func (f *Filters) AppendTo(values *form.Values) {
	for _, v := range f.f {
		if len(v.Op) > 0 {
			values.Add(fmt.Sprintf("%v[%v]", v.Key, v.Op), v.Val)
		} else {
			values.Add(v.Key, v.Val)
		}
	}
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
func (p *Params) AppendTo(body *form.Values) {
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
func (p *ListParams) AppendTo(body *form.Values) {
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
