//
//
// File generated from our OpenAPI spec
//
//

// Package giftcardoperation provides the /v1/gift_card_operations APIs
package giftcardoperation

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke /v1/gift_card_operations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a third-party gift card operation object.
func Get(id string, params *stripe.GiftCardOperationParams) (*stripe.GiftCardOperation, error) {
	return getC().Get(id, params)
}

// Retrieves a third-party gift card operation object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.GiftCardOperationParams) (*stripe.GiftCardOperation, error) {
	path := stripe.FormatURLPath("/v1/gift_card_operations/%s", id)
	giftcardoperation := &stripe.GiftCardOperation{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, giftcardoperation)
	return giftcardoperation, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
