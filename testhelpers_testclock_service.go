//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1TestHelpersTestClockService is used to invoke /v1/test_helpers/test_clocks APIs.
type v1TestHelpersTestClockService struct {
	B   Backend
	Key string
}

// Creates a new test clock that can be attached to new customers and quotes.
func (c v1TestHelpersTestClockService) Create(ctx context.Context, params *TestHelpersTestClockCreateParams) (*TestHelpersTestClock, error) {
	if params == nil {
		params = &TestHelpersTestClockCreateParams{}
	}
	params.Context = ctx
	testclock := &TestHelpersTestClock{}
	err := c.B.Call(
		http.MethodPost, "/v1/test_helpers/test_clocks", c.Key, params, testclock)
	return testclock, err
}

// Retrieves a test clock.
func (c v1TestHelpersTestClockService) Retrieve(ctx context.Context, id string, params *TestHelpersTestClockRetrieveParams) (*TestHelpersTestClock, error) {
	if params == nil {
		params = &TestHelpersTestClockRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/test_helpers/test_clocks/%s", id)
	testclock := &TestHelpersTestClock{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, testclock)
	return testclock, err
}

// Deletes a test clock.
func (c v1TestHelpersTestClockService) Delete(ctx context.Context, id string, params *TestHelpersTestClockDeleteParams) (*TestHelpersTestClock, error) {
	if params == nil {
		params = &TestHelpersTestClockDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/test_helpers/test_clocks/%s", id)
	testclock := &TestHelpersTestClock{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, testclock)
	return testclock, err
}

// Starts advancing a test clock to a specified time in the future. Advancement is done when status changes to Ready.
func (c v1TestHelpersTestClockService) Advance(ctx context.Context, id string, params *TestHelpersTestClockAdvanceParams) (*TestHelpersTestClock, error) {
	if params == nil {
		params = &TestHelpersTestClockAdvanceParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/test_helpers/test_clocks/%s/advance", id)
	testclock := &TestHelpersTestClock{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, testclock)
	return testclock, err
}

// Returns a list of your test clocks.
func (c v1TestHelpersTestClockService) List(ctx context.Context, listParams *TestHelpersTestClockListParams) Seq2[*TestHelpersTestClock, error] {
	if listParams == nil {
		listParams = &TestHelpersTestClockListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*TestHelpersTestClock, ListContainer, error) {
		list := &TestHelpersTestClockList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/test_helpers/test_clocks", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
