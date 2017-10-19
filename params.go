package stripe

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stripe/stripe-go/form"
)

const (
	endbefore  = "ending_before"
	startafter = "starting_after"
)

// Params is the structure that contains the common properties
// of any *Params structure.
type Params struct {
	// Account is deprecated form of StripeAccount that will do the same thing.
	// Please use StripeAccount instead.
	Account string `form:"-"` // Passed as header

	// Context used for request. It may carry deadlines, cancelation signals,
	// and other request-scoped values across API boundaries and between
	// processes.
	//
	// Note that a cancelled or timed out context does not provide any
	// guarantee whether the operation was or was not completed on Stripe's API
	// servers. For certainty, you must either retry with the same idempotency
	// key or query the state of the API.
	Context context.Context `form:"-"`

	Exp   []string     `form:"expand"`
	Extra *ExtraValues `form:"*"`

	// Headers may be used to provide extra header lines on the HTTP request.
	Headers http.Header `form:"-"`

	IdempotencyKey string            `form:"-"` // Passed as header
	Meta           map[string]string `form:"metadata"`

	// StripeAccount may contain the ID of a connected account. By including
	// this field, the request is made as if it originated from the connected
	// account instead of under the account of the owner of the configured
	// Stripe key.
	StripeAccount string `form:"-"` // Passed as header
}

// ExtraValues are extra parameters that are attached to an API request.
// They're implemented as a custom type so that they can have their own
// AppendTo implementation.
type ExtraValues struct {
	url.Values `form:"-"` // See custom AppendTo implementation
}

// AppendTo implements custom form encoding for extra parameter values.
func (v ExtraValues) AppendTo(body *form.Values, keyParts []string) {
	for k, vs := range v.Values {
		for _, v := range vs {
			body.Add(form.FormatKey(append(keyParts, k)), v)
		}
	}
}

// ListParams is the structure that contains the common properties
// of any *ListParams structure.
type ListParams struct {
	// Context used for request. It may carry deadlines, cancelation signals,
	// and other request-scoped values across API boundaries and between
	// processes.
	//
	// Note that a cancelled or timed out context does not provide any
	// guarantee whether the operation was or was not completed on Stripe's API
	// servers. For certainty, you must either retry with the same idempotency
	// key or query the state of the API.
	Context context.Context `form:"-"`

	End     string   `form:"ending_before"`
	Exp     []string `form:"expand"`
	Filters Filters  `form:"*"`
	Limit   int      `form:"limit"`

	// Single specifies whether this is a single page iterator. By default,
	// listing through an iterator will automatically grab additional pages as
	// the query progresses. To change this behavior and just load a single
	// page, set this to true.
	Single bool `form:"-"` // Not an API parameter

	Start string `form:"starting_after"`

	// StripeAccount may contain the ID of a connected account. By including
	// this field, the request is made as if it originated from the connected
	// account instead of under the account of the owner of the configured
	// Stripe key.
	StripeAccount string `form:"-"` // Passed as header
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

// Filters is a structure that contains a collection of filters for list-related APIs.
type Filters struct {
	f []*filter `form:"-"` // See custom AppendTo implementation
}

// AddFilter adds a new filter with a given key, op and value.
func (f *Filters) AddFilter(key, op, value string) {
	filter := &filter{Key: key, Op: op, Val: value}
	f.f = append(f.f, filter)
}

// AppendTo implements custom form encoding for filters.
func (f Filters) AppendTo(body *form.Values, keyParts []string) {
	if len(f.f) > 0 {
		for _, v := range f.f {
			if len(v.Op) > 0 {
				body.Add(form.FormatKey(append(keyParts, v.Key, v.Op)), v.Val)
			} else {
				body.Add(form.FormatKey(append(keyParts, v.Key)), v.Val)
			}
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
		p.Extra = &ExtraValues{Values: make(url.Values)}
	}

	p.Extra.Add(key, value)
}

// Expand appends a new field to expand.
func (p *ListParams) Expand(f string) {
	p.Exp = append(p.Exp, f)
}

// ToParams converts a ListParams to a Params by moving over any fields that
// have valid targets in the new type. This is useful because fields in
// Params can be injected directly into an http.Request while generally
// ListParams is only used to build a set of parameters.
func (p *ListParams) ToParams() *Params {
	return &Params{
		Context:       p.Context,
		StripeAccount: p.StripeAccount,
	}
}
