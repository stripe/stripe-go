// Package discount provides the discount-related APIs
package discount

import (
	"fmt"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke discount-related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Del removes a discount from a customer.
// For more details see https://stripe.com/docs/api#delete_discount.
func Del(customerID string, params *stripe.DiscountParams) (*stripe.Discount, error) {
	return getC().Del(customerID, params)
}

func (c Client) Del(customerID string, params *stripe.DiscountParams) (*stripe.Discount, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		commonParams = &params.Params
	}

	discount := &stripe.Discount{}
	err := c.B.Call("DELETE", fmt.Sprintf("/customers/%v/discount", customerID), c.Key, body, commonParams, discount)

	return discount, err
}

// DelSub removes a discount from a customer's subscription.
// For more details see https://stripe.com/docs/api#delete_subscription_discount.
func DelSub(subscriptionID string, params *stripe.DiscountParams) (*stripe.Discount, error) {
	return getC().DelSub(subscriptionID, params)
}

func (c Client) DelSub(subscriptionID string, params *stripe.DiscountParams) (*stripe.Discount, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		commonParams = &params.Params
	}

	discount := &stripe.Discount{}
	err := c.B.Call("DELETE", fmt.Sprintf("/subscriptions/%v/discount", subscriptionID), c.Key, body, commonParams, discount)

	return discount, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
