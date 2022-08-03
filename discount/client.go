//
//
// File generated from our OpenAPI spec
//
//

// Package discount provides the discount-related APIs
package discount

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
)

// Client is used to invoke discount-related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Del removes a discount from a customer.
func Del(id string, params *stripe.DiscountParams) (*stripe.Discount, error) {
	return getC().Del(id, params)
}

// Del removes a discount from a customer.
func (c Client) Del(id string, params *stripe.DiscountParams) (*stripe.Discount, error) {
	path := stripe.FormatURLPath("/v1/customers/%s/discount", id)
	discount := &stripe.Discount{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, discount)
	return discount, err
}

// DelSubscription removes a discount from a customer's subscription.
func DelSubscription(id string, params *stripe.DiscountParams) (*stripe.Discount, error) {
	return getC().DelSub(id, params)
}

// DelSub removes a discount from a customer's subscription.
func (c Client) DelSub(id string, params *stripe.DiscountParams) (*stripe.Discount, error) {
	path := stripe.FormatURLPath("/v1/subscriptions/%s/discount", id)
	discount := &stripe.Discount{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, discount)
	return discount, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
