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

// v1PaymentIntentService is used to invoke /v1/payment_intents APIs.
type v1PaymentIntentService struct {
	B   Backend
	Key string
}

// Creates a PaymentIntent object.
//
// After the PaymentIntent is created, attach a payment method and [confirm](https://docs.stripe.com/docs/api/payment_intents/confirm)
// to continue the payment. Learn more about <a href="/docs/payments/payment-intents">the available payment flows
// with the Payment Intents API.
//
// When you use confirm=true during creation, it's equivalent to creating
// and confirming the PaymentIntent in the same call. You can use any parameters
// available in the [confirm API](https://docs.stripe.com/docs/api/payment_intents/confirm) when you supply
// confirm=true.
func (c v1PaymentIntentService) Create(ctx context.Context, params *PaymentIntentCreateParams) (*PaymentIntent, error) {
	if params == nil {
		params = &PaymentIntentCreateParams{}
	}
	params.Context = ctx
	paymentintent := &PaymentIntent{}
	err := c.B.Call(
		http.MethodPost, "/v1/payment_intents", c.Key, params, paymentintent)
	return paymentintent, err
}

// Retrieves the details of a PaymentIntent that has previously been created.
//
// You can retrieve a PaymentIntent client-side using a publishable key when the client_secret is in the query string.
//
// If you retrieve a PaymentIntent with a publishable key, it only returns a subset of properties. Refer to the [payment intent](https://docs.stripe.com/api#payment_intent_object) object reference for more details.
func (c v1PaymentIntentService) Retrieve(ctx context.Context, id string, params *PaymentIntentRetrieveParams) (*PaymentIntent, error) {
	if params == nil {
		params = &PaymentIntentRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_intents/%s", id)
	paymentintent := &PaymentIntent{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentintent)
	return paymentintent, err
}

// Updates properties on a PaymentIntent object without confirming.
//
// Depending on which properties you update, you might need to confirm the
// PaymentIntent again. For example, updating the payment_method
// always requires you to confirm the PaymentIntent again. If you prefer to
// update and confirm at the same time, we recommend updating properties through
// the [confirm API](https://docs.stripe.com/docs/api/payment_intents/confirm) instead.
func (c v1PaymentIntentService) Update(ctx context.Context, id string, params *PaymentIntentUpdateParams) (*PaymentIntent, error) {
	if params == nil {
		params = &PaymentIntentUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_intents/%s", id)
	paymentintent := &PaymentIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentintent)
	return paymentintent, err
}

// Manually reconcile the remaining amount for a customer_balance PaymentIntent.
func (c v1PaymentIntentService) ApplyCustomerBalance(ctx context.Context, id string, params *PaymentIntentApplyCustomerBalanceParams) (*PaymentIntent, error) {
	if params == nil {
		params = &PaymentIntentApplyCustomerBalanceParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_intents/%s/apply_customer_balance", id)
	paymentintent := &PaymentIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentintent)
	return paymentintent, err
}

// You can cancel a PaymentIntent object when it's in one of these statuses: requires_payment_method, requires_capture, requires_confirmation, requires_action or, [in rare cases](https://docs.stripe.com/docs/payments/intents), processing.
//
// After it's canceled, no additional charges are made by the PaymentIntent and any operations on the PaymentIntent fail with an error. For PaymentIntents with a status of requires_capture, the remaining amount_capturable is automatically refunded.
//
// You can't cancel the PaymentIntent for a Checkout Session. [Expire the Checkout Session](https://docs.stripe.com/docs/api/checkout/sessions/expire) instead.
func (c v1PaymentIntentService) Cancel(ctx context.Context, id string, params *PaymentIntentCancelParams) (*PaymentIntent, error) {
	if params == nil {
		params = &PaymentIntentCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_intents/%s/cancel", id)
	paymentintent := &PaymentIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentintent)
	return paymentintent, err
}

// Capture the funds of an existing uncaptured PaymentIntent when its status is requires_capture.
//
// Uncaptured PaymentIntents are cancelled a set number of days (7 by default) after their creation.
//
// Learn more about [separate authorization and capture](https://docs.stripe.com/docs/payments/capture-later).
func (c v1PaymentIntentService) Capture(ctx context.Context, id string, params *PaymentIntentCaptureParams) (*PaymentIntent, error) {
	if params == nil {
		params = &PaymentIntentCaptureParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_intents/%s/capture", id)
	paymentintent := &PaymentIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentintent)
	return paymentintent, err
}

// Confirm that your customer intends to pay with current or provided
// payment method. Upon confirmation, the PaymentIntent will attempt to initiate
// a payment.
//
// If the selected payment method requires additional authentication steps, the
// PaymentIntent will transition to the requires_action status and
// suggest additional actions via next_action. If payment fails,
// the PaymentIntent transitions to the requires_payment_method status or the
// canceled status if the confirmation limit is reached. If
// payment succeeds, the PaymentIntent will transition to the succeeded
// status (or requires_capture, if capture_method is set to manual).
//
// If the confirmation_method is automatic, payment may be attempted
// using our [client SDKs](https://docs.stripe.com/docs/stripe-js/reference#stripe-handle-card-payment)
// and the PaymentIntent's [client_secret](https://docs.stripe.com/api#payment_intent_object-client_secret).
// After next_actions are handled by the client, no additional
// confirmation is required to complete the payment.
//
// If the confirmation_method is manual, all payment attempts must be
// initiated using a secret key.
//
// If any actions are required for the payment, the PaymentIntent will
// return to the requires_confirmation state
// after those actions are completed. Your server needs to then
// explicitly re-confirm the PaymentIntent to initiate the next payment
// attempt.
//
// There is a variable upper limit on how many times a PaymentIntent can be confirmed.
// After this limit is reached, any further calls to this endpoint will
// transition the PaymentIntent to the canceled state.
func (c v1PaymentIntentService) Confirm(ctx context.Context, id string, params *PaymentIntentConfirmParams) (*PaymentIntent, error) {
	if params == nil {
		params = &PaymentIntentConfirmParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_intents/%s/confirm", id)
	paymentintent := &PaymentIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentintent)
	return paymentintent, err
}

// Perform an incremental authorization on an eligible
// [PaymentIntent](https://docs.stripe.com/docs/api/payment_intents/object). To be eligible, the
// PaymentIntent's status must be requires_capture and
// [incremental_authorization_supported](https://docs.stripe.com/docs/api/charges/object#charge_object-payment_method_details-card_present-incremental_authorization_supported)
// must be true.
//
// Incremental authorizations attempt to increase the authorized amount on
// your customer's card to the new, higher amount provided. Similar to the
// initial authorization, incremental authorizations can be declined. A
// single PaymentIntent can call this endpoint multiple times to further
// increase the authorized amount.
//
// If the incremental authorization succeeds, the PaymentIntent object
// returns with the updated
// [amount](https://docs.stripe.com/docs/api/payment_intents/object#payment_intent_object-amount).
// If the incremental authorization fails, a
// [card_declined](https://docs.stripe.com/docs/error-codes#card-declined) error returns, and no other
// fields on the PaymentIntent or Charge update. The PaymentIntent
// object remains capturable for the previously authorized amount.
//
// Each PaymentIntent can have a maximum of 10 incremental authorization attempts, including declines.
// After it's captured, a PaymentIntent can no longer be incremented.
//
// Learn more about [incremental authorizations](https://docs.stripe.com/docs/terminal/features/incremental-authorizations).
func (c v1PaymentIntentService) IncrementAuthorization(ctx context.Context, id string, params *PaymentIntentIncrementAuthorizationParams) (*PaymentIntent, error) {
	if params == nil {
		params = &PaymentIntentIncrementAuthorizationParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_intents/%s/increment_authorization", id)
	paymentintent := &PaymentIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentintent)
	return paymentintent, err
}

// Verifies microdeposits on a PaymentIntent object.
func (c v1PaymentIntentService) VerifyMicrodeposits(ctx context.Context, id string, params *PaymentIntentVerifyMicrodepositsParams) (*PaymentIntent, error) {
	if params == nil {
		params = &PaymentIntentVerifyMicrodepositsParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_intents/%s/verify_microdeposits", id)
	paymentintent := &PaymentIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentintent)
	return paymentintent, err
}

// Returns a list of PaymentIntents.
func (c v1PaymentIntentService) List(ctx context.Context, listParams *PaymentIntentListParams) Seq2[*PaymentIntent, error] {
	if listParams == nil {
		listParams = &PaymentIntentListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*PaymentIntent, ListContainer, error) {
		list := &PaymentIntentList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/payment_intents", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}

// Search for PaymentIntents you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
func (c v1PaymentIntentService) Search(ctx context.Context, params *PaymentIntentSearchParams) Seq2[*PaymentIntent, error] {
	if params == nil {
		params = &PaymentIntentSearchParams{}
	}
	params.Context = ctx
	return newV1SearchList(params, func(p *Params, b *form.Values) ([]*PaymentIntent, SearchContainer, error) {
		list := &PaymentIntentSearchResult{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/payment_intents/search", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
