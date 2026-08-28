//
//
// File generated from our OpenAPI spec
//
//

// Package feedbackoption provides the /v1/billing/feedback_options APIs
package feedbackoption

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
	"github.com/stripe/stripe-go/v86/form"
)

// Client is used to invoke /v1/billing/feedback_options APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new feedback option.
func New(params *stripe.BillingFeedbackOptionParams) (*stripe.BillingFeedbackOption, error) {
	return getC().New(params)
}

// Creates a new feedback option.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.BillingFeedbackOptionParams) (*stripe.BillingFeedbackOption, error) {
	feedbackoption := &stripe.BillingFeedbackOption{}
	err := c.B.Call(
		http.MethodPost, "/v1/billing/feedback_options", c.Key, params, feedbackoption)
	return feedbackoption, err
}

// Retrieves a feedback options object given an ID.
func Get(id string, params *stripe.BillingFeedbackOptionParams) (*stripe.BillingFeedbackOption, error) {
	return getC().Get(id, params)
}

// Retrieves a feedback options object given an ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.BillingFeedbackOptionParams) (*stripe.BillingFeedbackOption, error) {
	path := stripe.FormatURLPath("/v1/billing/feedback_options/%s", id)
	feedbackoption := &stripe.BillingFeedbackOption{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, feedbackoption)
	return feedbackoption, err
}

// Updates the description of an existing feedback option.
func Update(id string, params *stripe.BillingFeedbackOptionParams) (*stripe.BillingFeedbackOption, error) {
	return getC().Update(id, params)
}

// Updates the description of an existing feedback option.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.BillingFeedbackOptionParams) (*stripe.BillingFeedbackOption, error) {
	path := stripe.FormatURLPath("/v1/billing/feedback_options/%s", id)
	feedbackoption := &stripe.BillingFeedbackOption{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, feedbackoption)
	return feedbackoption, err
}

// Deactivates a feedback option. Deactivated feedback options cannot be used in portal configurations.
func Deactivate(id string, params *stripe.BillingFeedbackOptionDeactivateParams) (*stripe.BillingFeedbackOption, error) {
	return getC().Deactivate(id, params)
}

// Deactivates a feedback option. Deactivated feedback options cannot be used in portal configurations.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Deactivate(id string, params *stripe.BillingFeedbackOptionDeactivateParams) (*stripe.BillingFeedbackOption, error) {
	path := stripe.FormatURLPath("/v1/billing/feedback_options/%s/deactivate", id)
	feedbackoption := &stripe.BillingFeedbackOption{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, feedbackoption)
	return feedbackoption, err
}

// An API method for listing the feedback options model
func List(params *stripe.BillingFeedbackOptionListParams) *Iter {
	return getC().List(params)
}

// An API method for listing the feedback options model
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.BillingFeedbackOptionListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.BillingFeedbackOptionList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/billing/feedback_options", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for billing feedback options.
type Iter struct {
	*stripe.Iter
}

// BillingFeedbackOption returns the billing feedback option which the iterator is currently pointing to.
func (i *Iter) BillingFeedbackOption() *stripe.BillingFeedbackOption {
	return i.Current().(*stripe.BillingFeedbackOption)
}

// BillingFeedbackOptionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) BillingFeedbackOptionList() *stripe.BillingFeedbackOptionList {
	return i.List().(*stripe.BillingFeedbackOptionList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
