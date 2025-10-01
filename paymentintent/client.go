//
//
// File generated from our OpenAPI spec
//
//

// Package paymentintent provides the /v1/payment_intents APIs
package paymentintent

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
	"github.com/stripe/stripe-go/v83/form"
)

// Client is used to invoke /v1/payment_intents APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
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
func New(params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	return getC().New(params)
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
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	paymentintent := &stripe.PaymentIntent{}
	err := c.B.Call(
		http.MethodPost, "/v1/payment_intents", c.Key, params, paymentintent)
	return paymentintent, err
}

// Retrieves the details of a PaymentIntent that has previously been created.
//
// You can retrieve a PaymentIntent client-side using a publishable key when the client_secret is in the query string.
//
// If you retrieve a PaymentIntent with a publishable key, it only returns a subset of properties. Refer to the [payment intent](https://docs.stripe.com/api#payment_intent_object) object reference for more details.
func Get(id string, params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	return getC().Get(id, params)
}

// Retrieves the details of a PaymentIntent that has previously been created.
//
// You can retrieve a PaymentIntent client-side using a publishable key when the client_secret is in the query string.
//
// If you retrieve a PaymentIntent with a publishable key, it only returns a subset of properties. Refer to the [payment intent](https://docs.stripe.com/api#payment_intent_object) object reference for more details.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	path := stripe.FormatURLPath("/v1/payment_intents/%s", id)
	paymentintent := &stripe.PaymentIntent{}
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
func Update(id string, params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	return getC().Update(id, params)
}

// Updates properties on a PaymentIntent object without confirming.
//
// Depending on which properties you update, you might need to confirm the
// PaymentIntent again. For example, updating the payment_method
// always requires you to confirm the PaymentIntent again. If you prefer to
// update and confirm at the same time, we recommend updating properties through
// the [confirm API](https://docs.stripe.com/docs/api/payment_intents/confirm) instead.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	path := stripe.FormatURLPath("/v1/payment_intents/%s", id)
	paymentintent := &stripe.PaymentIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentintent)
	return paymentintent, err
}

// Manually reconcile the remaining amount for a customer_balance PaymentIntent.
func ApplyCustomerBalance(id string, params *stripe.PaymentIntentApplyCustomerBalanceParams) (*stripe.PaymentIntent, error) {
	return getC().ApplyCustomerBalance(id, params)
}

// Manually reconcile the remaining amount for a customer_balance PaymentIntent.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ApplyCustomerBalance(id string, params *stripe.PaymentIntentApplyCustomerBalanceParams) (*stripe.PaymentIntent, error) {
	path := stripe.FormatURLPath(
		"/v1/payment_intents/%s/apply_customer_balance", id)
	paymentintent := &stripe.PaymentIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentintent)
	return paymentintent, err
}

// You can cancel a PaymentIntent object when it's in one of these statuses: requires_payment_method, requires_capture, requires_confirmation, requires_action or, [in rare cases](https://docs.stripe.com/docs/payments/intents), processing.
//
// After it's canceled, no additional charges are made by the PaymentIntent and any operations on the PaymentIntent fail with an error. For PaymentIntents with a status of requires_capture, the remaining amount_capturable is automatically refunded.
//
// You can't cancel the PaymentIntent for a Checkout Session. [Expire the Checkout Session](https://docs.stripe.com/docs/api/checkout/sessions/expire) instead.
func Cancel(id string, params *stripe.PaymentIntentCancelParams) (*stripe.PaymentIntent, error) {
	return getC().Cancel(id, params)
}

// You can cancel a PaymentIntent object when it's in one of these statuses: requires_payment_method, requires_capture, requires_confirmation, requires_action or, [in rare cases](https://docs.stripe.com/docs/payments/intents), processing.
//
// After it's canceled, no additional charges are made by the PaymentIntent and any operations on the PaymentIntent fail with an error. For PaymentIntents with a status of requires_capture, the remaining amount_capturable is automatically refunded.
//
// You can't cancel the PaymentIntent for a Checkout Session. [Expire the Checkout Session](https://docs.stripe.com/docs/api/checkout/sessions/expire) instead.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.PaymentIntentCancelParams) (*stripe.PaymentIntent, error) {
	path := stripe.FormatURLPath("/v1/payment_intents/%s/cancel", id)
	paymentintent := &stripe.PaymentIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentintent)
	return paymentintent, err
}

// Capture the funds of an existing uncaptured PaymentIntent when its status is requires_capture.
//
// Uncaptured PaymentIntents are cancelled a set number of days (7 by default) after their creation.
//
// Learn more about [separate authorization and capture](https://docs.stripe.com/docs/payments/capture-later).
func Capture(id string, params *stripe.PaymentIntentCaptureParams) (*stripe.PaymentIntent, error) {
	return getC().Capture(id, params)
}

// Capture the funds of an existing uncaptured PaymentIntent when its status is requires_capture.
//
// Uncaptured PaymentIntents are cancelled a set number of days (7 by default) after their creation.
//
// Learn more about [separate authorization and capture](https://docs.stripe.com/docs/payments/capture-later).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Capture(id string, params *stripe.PaymentIntentCaptureParams) (*stripe.PaymentIntent, error) {
	path := stripe.FormatURLPath("/v1/payment_intents/%s/capture", id)
	paymentintent := &stripe.PaymentIntent{}
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
func Confirm(id string, params *stripe.PaymentIntentConfirmParams) (*stripe.PaymentIntent, error) {
	return getC().Confirm(id, params)
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
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Confirm(id string, params *stripe.PaymentIntentConfirmParams) (*stripe.PaymentIntent, error) {
	path := stripe.FormatURLPath("/v1/payment_intents/%s/confirm", id)
	paymentintent := &stripe.PaymentIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentintent)
	return paymentintent, err
}

// Perform a decremental authorization on an eligible
// [PaymentIntent](https://docs.stripe.com/docs/api/payment_intents/object). To be eligible, the
// PaymentIntent's status must be requires_capture and
// [decremental_authorization.status](https://docs.stripe.com/docs/api/charges/object#charge_object-payment_method_details-card-decremental_authorization)
// must be available.
//
// Decremental authorizations decrease the authorized amount on your customer's card
// to the new, lower amount provided. A single PaymentIntent can call this endpoint multiple times to further decrease the authorized amount.
//
// After decrement, the PaymentIntent object
// returns with the updated
// [amount](https://docs.stripe.com/docs/api/payment_intents/object#payment_intent_object-amount).
// The PaymentIntent will now be capturable up to the new authorized amount.
//
// Each PaymentIntent can have a maximum of 10 decremental or incremental authorization attempts, including declines.
// After it's fully captured, a PaymentIntent can no longer be decremented.
func DecrementAuthorization(id string, params *stripe.PaymentIntentDecrementAuthorizationParams) (*stripe.PaymentIntent, error) {
	return getC().DecrementAuthorization(id, params)
}

// Perform a decremental authorization on an eligible
// [PaymentIntent](https://docs.stripe.com/docs/api/payment_intents/object). To be eligible, the
// PaymentIntent's status must be requires_capture and
// [decremental_authorization.status](https://docs.stripe.com/docs/api/charges/object#charge_object-payment_method_details-card-decremental_authorization)
// must be available.
//
// Decremental authorizations decrease the authorized amount on your customer's card
// to the new, lower amount provided. A single PaymentIntent can call this endpoint multiple times to further decrease the authorized amount.
//
// After decrement, the PaymentIntent object
// returns with the updated
// [amount](https://docs.stripe.com/docs/api/payment_intents/object#payment_intent_object-amount).
// The PaymentIntent will now be capturable up to the new authorized amount.
//
// Each PaymentIntent can have a maximum of 10 decremental or incremental authorization attempts, including declines.
// After it's fully captured, a PaymentIntent can no longer be decremented.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) DecrementAuthorization(id string, params *stripe.PaymentIntentDecrementAuthorizationParams) (*stripe.PaymentIntent, error) {
	path := stripe.FormatURLPath(
		"/v1/payment_intents/%s/decrement_authorization", id)
	paymentintent := &stripe.PaymentIntent{}
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
func IncrementAuthorization(id string, params *stripe.PaymentIntentIncrementAuthorizationParams) (*stripe.PaymentIntent, error) {
	return getC().IncrementAuthorization(id, params)
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
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) IncrementAuthorization(id string, params *stripe.PaymentIntentIncrementAuthorizationParams) (*stripe.PaymentIntent, error) {
	path := stripe.FormatURLPath(
		"/v1/payment_intents/%s/increment_authorization", id)
	paymentintent := &stripe.PaymentIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentintent)
	return paymentintent, err
}

// Trigger an external action on a PaymentIntent.
func TriggerAction(id string, params *stripe.PaymentIntentTriggerActionParams) (*stripe.PaymentIntent, error) {
	return getC().TriggerAction(id, params)
}

// Trigger an external action on a PaymentIntent.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) TriggerAction(id string, params *stripe.PaymentIntentTriggerActionParams) (*stripe.PaymentIntent, error) {
	path := stripe.FormatURLPath("/v1/test/payment_intents/%s/trigger_action", id)
	paymentintent := &stripe.PaymentIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentintent)
	return paymentintent, err
}

// Verifies microdeposits on a PaymentIntent object.
func VerifyMicrodeposits(id string, params *stripe.PaymentIntentVerifyMicrodepositsParams) (*stripe.PaymentIntent, error) {
	return getC().VerifyMicrodeposits(id, params)
}

// Verifies microdeposits on a PaymentIntent object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) VerifyMicrodeposits(id string, params *stripe.PaymentIntentVerifyMicrodepositsParams) (*stripe.PaymentIntent, error) {
	path := stripe.FormatURLPath(
		"/v1/payment_intents/%s/verify_microdeposits", id)
	paymentintent := &stripe.PaymentIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentintent)
	return paymentintent, err
}

// Returns a list of PaymentIntents.
func List(params *stripe.PaymentIntentListParams) *Iter {
	return getC().List(params)
}

// Returns a list of PaymentIntents.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.PaymentIntentListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.PaymentIntentList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/payment_intents", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for payment intents.
type Iter struct {
	*stripe.Iter
}

// PaymentIntent returns the payment intent which the iterator is currently pointing to.
func (i *Iter) PaymentIntent() *stripe.PaymentIntent {
	return i.Current().(*stripe.PaymentIntent)
}

// PaymentIntentList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PaymentIntentList() *stripe.PaymentIntentList {
	return i.List().(*stripe.PaymentIntentList)
}

// Search for PaymentIntents you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
func Search(params *stripe.PaymentIntentSearchParams) *SearchIter {
	return getC().Search(params)
}

// Search for PaymentIntents you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Search(params *stripe.PaymentIntentSearchParams) *SearchIter {
	return &SearchIter{
		SearchIter: stripe.GetSearchIter(params, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.SearchContainer, error) {
			list := &stripe.PaymentIntentSearchResult{}
			err := c.B.CallRaw(http.MethodGet, "/v1/payment_intents/search", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// SearchIter is an iterator for payment intents.
type SearchIter struct {
	*stripe.SearchIter
}

// PaymentIntent returns the payment intent which the iterator is currently pointing to.
func (i *SearchIter) PaymentIntent() *stripe.PaymentIntent {
	return i.Current().(*stripe.PaymentIntent)
}

// PaymentIntentSearchResult returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *SearchIter) PaymentIntentSearchResult() *stripe.PaymentIntentSearchResult {
	return i.SearchResult().(*stripe.PaymentIntentSearchResult)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
