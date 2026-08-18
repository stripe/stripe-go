//
//
// File generated from our OpenAPI spec
//
//

// Package paymentplan provides the /v1/payment_plans APIs
package paymentplan

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
	"github.com/stripe/stripe-go/v86/form"
)

// Client is used to invoke /v1/payment_plans APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a payment plan that splits a single invoice obligation into installments with their own due dates and amounts.
func New(params *stripe.PaymentPlanParams) (*stripe.PaymentPlan, error) {
	return getC().New(params)
}

// Creates a payment plan that splits a single invoice obligation into installments with their own due dates and amounts.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.PaymentPlanParams) (*stripe.PaymentPlan, error) {
	paymentplan := &stripe.PaymentPlan{}
	err := c.B.Call(
		http.MethodPost, "/v1/payment_plans", c.Key, params, paymentplan)
	return paymentplan, err
}

// Retrieves the payment plan with the given ID.
func Get(id string, params *stripe.PaymentPlanParams) (*stripe.PaymentPlan, error) {
	return getC().Get(id, params)
}

// Retrieves the payment plan with the given ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.PaymentPlanParams) (*stripe.PaymentPlan, error) {
	path := stripe.FormatURLPath("/v1/payment_plans/%s", id)
	paymentplan := &stripe.PaymentPlan{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentplan)
	return paymentplan, err
}

// Updates the schedule or metadata of an existing payment plan. Only unpaid installments can be updated.
func Update(id string, params *stripe.PaymentPlanParams) (*stripe.PaymentPlan, error) {
	return getC().Update(id, params)
}

// Updates the schedule or metadata of an existing payment plan. Only unpaid installments can be updated.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.PaymentPlanParams) (*stripe.PaymentPlan, error) {
	path := stripe.FormatURLPath("/v1/payment_plans/%s", id)
	paymentplan := &stripe.PaymentPlan{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentplan)
	return paymentplan, err
}

// Returns a list of payment plans.
func List(params *stripe.PaymentPlanListParams) *Iter {
	return getC().List(params)
}

// Returns a list of payment plans.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.PaymentPlanListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.PaymentPlanList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/payment_plans", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for payment plans.
type Iter struct {
	*stripe.Iter
}

// PaymentPlan returns the payment plan which the iterator is currently pointing to.
func (i *Iter) PaymentPlan() *stripe.PaymentPlan {
	return i.Current().(*stripe.PaymentPlan)
}

// PaymentPlanList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PaymentPlanList() *stripe.PaymentPlanList {
	return i.List().(*stripe.PaymentPlanList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
