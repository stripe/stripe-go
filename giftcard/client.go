//
//
// File generated from our OpenAPI spec
//
//

// Package giftcard provides the /v1/gift_cards APIs
package giftcard

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/gift_cards APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a gift card object.
func New(params *stripe.GiftCardParams) (*stripe.GiftCard, error) {
	return getC().New(params)
}

// Creates a gift card object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.GiftCardParams) (*stripe.GiftCard, error) {
	giftcard := &stripe.GiftCard{}
	err := c.B.Call(http.MethodPost, "/v1/gift_cards", c.Key, params, giftcard)
	return giftcard, err
}

// Retrieves a third-party gift card object.
func Get(id string, params *stripe.GiftCardParams) (*stripe.GiftCard, error) {
	return getC().Get(id, params)
}

// Retrieves a third-party gift card object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.GiftCardParams) (*stripe.GiftCard, error) {
	path := stripe.FormatURLPath("/v1/gift_cards/%s", id)
	giftcard := &stripe.GiftCard{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, giftcard)
	return giftcard, err
}

// Activates a third-party gift card and optionally sets its balance.
func Activate(id string, params *stripe.GiftCardActivateParams) (*stripe.GiftCardOperation, error) {
	return getC().Activate(id, params)
}

// Activates a third-party gift card and optionally sets its balance.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Activate(id string, params *stripe.GiftCardActivateParams) (*stripe.GiftCardOperation, error) {
	path := stripe.FormatURLPath("/v1/gift_cards/%s/activate", id)
	giftcardoperation := &stripe.GiftCardOperation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, giftcardoperation)
	return giftcardoperation, err
}

// Cashout a third-party gift card by zeroing its balance.
func Cashout(id string, params *stripe.GiftCardCashoutParams) (*stripe.GiftCardOperation, error) {
	return getC().Cashout(id, params)
}

// Cashout a third-party gift card by zeroing its balance.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cashout(id string, params *stripe.GiftCardCashoutParams) (*stripe.GiftCardOperation, error) {
	path := stripe.FormatURLPath("/v1/gift_cards/%s/cashout", id)
	giftcardoperation := &stripe.GiftCardOperation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, giftcardoperation)
	return giftcardoperation, err
}

// Checks the balance of a third-party gift card.
func CheckBalance(id string, params *stripe.GiftCardCheckBalanceParams) (*stripe.GiftCardOperation, error) {
	return getC().CheckBalance(id, params)
}

// Checks the balance of a third-party gift card.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) CheckBalance(id string, params *stripe.GiftCardCheckBalanceParams) (*stripe.GiftCardOperation, error) {
	path := stripe.FormatURLPath("/v1/gift_cards/%s/check_balance", id)
	giftcardoperation := &stripe.GiftCardOperation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, giftcardoperation)
	return giftcardoperation, err
}

// Reloads a third-party gift card by adding the specified amount to its balance.
func Reload(id string, params *stripe.GiftCardReloadParams) (*stripe.GiftCardOperation, error) {
	return getC().Reload(id, params)
}

// Reloads a third-party gift card by adding the specified amount to its balance.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Reload(id string, params *stripe.GiftCardReloadParams) (*stripe.GiftCardOperation, error) {
	path := stripe.FormatURLPath("/v1/gift_cards/%s/reload", id)
	giftcardoperation := &stripe.GiftCardOperation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, giftcardoperation)
	return giftcardoperation, err
}

// Voids a previously performed gift card operation.
func VoidOperation(id string, params *stripe.GiftCardVoidOperationParams) (*stripe.GiftCardOperation, error) {
	return getC().VoidOperation(id, params)
}

// Voids a previously performed gift card operation.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) VoidOperation(id string, params *stripe.GiftCardVoidOperationParams) (*stripe.GiftCardOperation, error) {
	path := stripe.FormatURLPath("/v1/gift_cards/%s/void_operation", id)
	giftcardoperation := &stripe.GiftCardOperation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, giftcardoperation)
	return giftcardoperation, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
