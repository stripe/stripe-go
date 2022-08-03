//
//
// File generated from our OpenAPI spec
//
//

// Package testclock provides the /test_helpers/test_clocks APIs
package testclock

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /test_helpers/test_clocks APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new test helpers test clock.
func New(params *stripe.TestHelpersTestClockParams) (*stripe.TestHelpersTestClock, error) {
	return getC().New(params)
}

// New creates a new test helpers test clock.
func (c Client) New(params *stripe.TestHelpersTestClockParams) (*stripe.TestHelpersTestClock, error) {
	testclock := &stripe.TestHelpersTestClock{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/test_helpers/test_clocks",
		c.Key,
		params,
		testclock,
	)
	return testclock, err
}

// Get returns the details of a test helpers test clock.
func Get(id string, params *stripe.TestHelpersTestClockParams) (*stripe.TestHelpersTestClock, error) {
	return getC().Get(id, params)
}

// Get returns the details of a test helpers test clock.
func (c Client) Get(id string, params *stripe.TestHelpersTestClockParams) (*stripe.TestHelpersTestClock, error) {
	path := stripe.FormatURLPath("/v1/test_helpers/test_clocks/%s", id)
	testclock := &stripe.TestHelpersTestClock{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, testclock)
	return testclock, err
}

// Del removes a test helpers test clock.
func Del(id string, params *stripe.TestHelpersTestClockParams) (*stripe.TestHelpersTestClock, error) {
	return getC().Del(id, params)
}

// Del removes a test helpers test clock.
func (c Client) Del(id string, params *stripe.TestHelpersTestClockParams) (*stripe.TestHelpersTestClock, error) {
	path := stripe.FormatURLPath("/v1/test_helpers/test_clocks/%s", id)
	testclock := &stripe.TestHelpersTestClock{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, testclock)
	return testclock, err
}

// Advance is the method for the `POST /v1/test_helpers/test_clocks/{test_clock}/advance` API.
func Advance(id string, params *stripe.TestHelpersTestClockAdvanceParams) (*stripe.TestHelpersTestClock, error) {
	return getC().Advance(id, params)
}

// Advance is the method for the `POST /v1/test_helpers/test_clocks/{test_clock}/advance` API.
func (c Client) Advance(id string, params *stripe.TestHelpersTestClockAdvanceParams) (*stripe.TestHelpersTestClock, error) {
	path := stripe.FormatURLPath("/v1/test_helpers/test_clocks/%s/advance", id)
	testclock := &stripe.TestHelpersTestClock{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, testclock)
	return testclock, err
}

// List returns a list of test helpers test clocks.
func List(params *stripe.TestHelpersTestClockListParams) *Iter {
	return getC().List(params)
}

// List returns a list of test helpers test clocks.
func (c Client) List(listParams *stripe.TestHelpersTestClockListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TestHelpersTestClockList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/test_helpers/test_clocks", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for test helpers test clocks.
type Iter struct {
	*stripe.Iter
}

// TestHelpersTestClock returns the test helpers test clock which the iterator is currently pointing to.
func (i *Iter) TestHelpersTestClock() *stripe.TestHelpersTestClock {
	return i.Current().(*stripe.TestHelpersTestClock)
}

// TestHelpersTestClockList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TestHelpersTestClockList() *stripe.TestHelpersTestClockList {
	return i.List().(*stripe.TestHelpersTestClockList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
