//
//
// File generated from our OpenAPI spec
//
//

// Package subscription provides the /v1/subscriptions APIs
package subscription

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/form"
)

// Client is used to invoke /v1/subscriptions APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new subscription on an existing customer. Each customer can have up to 500 active or scheduled subscriptions.
//
// When you create a subscription with collection_method=charge_automatically, the first invoice is finalized as part of the request.
// The payment_behavior parameter determines the exact behavior of the initial payment.
//
// To start subscriptions where the first invoice always begins in a draft status, use [subscription schedules](https://docs.stripe.com/docs/billing/subscriptions/subscription-schedules#managing) instead.
// Schedules provide the flexibility to model more complex billing configurations that change over time.
func New(params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	return getC().New(params)
}

// Creates a new subscription on an existing customer. Each customer can have up to 500 active or scheduled subscriptions.
//
// When you create a subscription with collection_method=charge_automatically, the first invoice is finalized as part of the request.
// The payment_behavior parameter determines the exact behavior of the initial payment.
//
// To start subscriptions where the first invoice always begins in a draft status, use [subscription schedules](https://docs.stripe.com/docs/billing/subscriptions/subscription-schedules#managing) instead.
// Schedules provide the flexibility to model more complex billing configurations that change over time.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	subscription := &stripe.Subscription{}
	err := c.B.Call(
		http.MethodPost, "/v1/subscriptions", c.Key, params, subscription)
	return subscription, err
}

// Retrieves the subscription with the given ID.
func Get(id string, params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	return getC().Get(id, params)
}

// Retrieves the subscription with the given ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	path := stripe.FormatURLPath("/v1/subscriptions/%s", id)
	subscription := &stripe.Subscription{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, subscription)
	return subscription, err
}

// Updates an existing subscription to match the specified parameters.
// When changing prices or quantities, we optionally prorate the price we charge next month to make up for any price changes.
// To preview how the proration is calculated, use the [create preview](https://docs.stripe.com/docs/api/invoices/create_preview) endpoint.
//
// By default, we prorate subscription changes. For example, if a customer signs up on May 1 for a 100 price, they'll be billed 100 immediately. If on May 15 they switch to a 200 price, then on June 1 they'll be billed 250 (200 for a renewal of her subscription, plus a 50 prorating adjustment for half of the previous month's 100 difference). Similarly, a downgrade generates a credit that is applied to the next invoice. We also prorate when you make quantity changes.
//
// Switching prices does not normally change the billing date or generate an immediate charge unless:
//
// The billing interval is changed (for example, from monthly to yearly).
// The subscription moves from free to paid.
// A trial starts or ends.
//
// In these cases, we apply a credit for the unused time on the previous price, immediately charge the customer using the new price, and reset the billing date. Learn about how [Stripe immediately attempts payment for subscription changes](https://docs.stripe.com/docs/billing/subscriptions/upgrade-downgrade#immediate-payment).
//
// If you want to charge for an upgrade immediately, pass proration_behavior as always_invoice to create prorations, automatically invoice the customer for those proration adjustments, and attempt to collect payment. If you pass create_prorations, the prorations are created but not automatically invoiced. If you want to bill the customer for the prorations before the subscription's renewal date, you need to manually [invoice the customer](https://docs.stripe.com/docs/api/invoices/create).
//
// If you don't want to prorate, set the proration_behavior option to none. With this option, the customer is billed 100 on May 1 and 200 on June 1. Similarly, if you set proration_behavior to none when switching between different billing intervals (for example, from monthly to yearly), we don't generate any credits for the old subscription's unused time. We still reset the billing date and bill immediately for the new subscription.
//
// Updating the quantity on a subscription many times in an hour may result in [rate limiting. If you need to bill for a frequently changing quantity, consider integrating <a href="/docs/billing/subscriptions/usage-based">usage-based billing](https://docs.stripe.com/docs/rate-limits) instead.
func Update(id string, params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	return getC().Update(id, params)
}

// Updates an existing subscription to match the specified parameters.
// When changing prices or quantities, we optionally prorate the price we charge next month to make up for any price changes.
// To preview how the proration is calculated, use the [create preview](https://docs.stripe.com/docs/api/invoices/create_preview) endpoint.
//
// By default, we prorate subscription changes. For example, if a customer signs up on May 1 for a 100 price, they'll be billed 100 immediately. If on May 15 they switch to a 200 price, then on June 1 they'll be billed 250 (200 for a renewal of her subscription, plus a 50 prorating adjustment for half of the previous month's 100 difference). Similarly, a downgrade generates a credit that is applied to the next invoice. We also prorate when you make quantity changes.
//
// Switching prices does not normally change the billing date or generate an immediate charge unless:
//
// The billing interval is changed (for example, from monthly to yearly).
// The subscription moves from free to paid.
// A trial starts or ends.
//
// In these cases, we apply a credit for the unused time on the previous price, immediately charge the customer using the new price, and reset the billing date. Learn about how [Stripe immediately attempts payment for subscription changes](https://docs.stripe.com/docs/billing/subscriptions/upgrade-downgrade#immediate-payment).
//
// If you want to charge for an upgrade immediately, pass proration_behavior as always_invoice to create prorations, automatically invoice the customer for those proration adjustments, and attempt to collect payment. If you pass create_prorations, the prorations are created but not automatically invoiced. If you want to bill the customer for the prorations before the subscription's renewal date, you need to manually [invoice the customer](https://docs.stripe.com/docs/api/invoices/create).
//
// If you don't want to prorate, set the proration_behavior option to none. With this option, the customer is billed 100 on May 1 and 200 on June 1. Similarly, if you set proration_behavior to none when switching between different billing intervals (for example, from monthly to yearly), we don't generate any credits for the old subscription's unused time. We still reset the billing date and bill immediately for the new subscription.
//
// Updating the quantity on a subscription many times in an hour may result in [rate limiting. If you need to bill for a frequently changing quantity, consider integrating <a href="/docs/billing/subscriptions/usage-based">usage-based billing](https://docs.stripe.com/docs/rate-limits) instead.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	path := stripe.FormatURLPath("/v1/subscriptions/%s", id)
	subscription := &stripe.Subscription{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, subscription)
	return subscription, err
}

// Cancels a customer's subscription immediately. The customer won't be charged again for the subscription. After it's canceled, you can no longer update the subscription or its [metadata](https://docs.stripe.com/metadata).
//
// Any pending invoice items that you've created are still charged at the end of the period, unless manually [deleted](https://docs.stripe.com/api#delete_invoiceitem). If you've set the subscription to cancel at the end of the period, any pending prorations are also left in place and collected at the end of the period. But if the subscription is set to cancel immediately, pending prorations are removed if invoice_now and prorate are both set to true.
//
// By default, upon subscription cancellation, Stripe stops automatic collection of all finalized invoices for the customer. This is intended to prevent unexpected payment attempts after the customer has canceled a subscription. However, you can resume automatic collection of the invoices manually after subscription cancellation to have us proceed. Or, you could check for unpaid invoices before allowing the customer to cancel the subscription at all.
func Cancel(id string, params *stripe.SubscriptionCancelParams) (*stripe.Subscription, error) {
	return getC().Cancel(id, params)
}

// Cancels a customer's subscription immediately. The customer won't be charged again for the subscription. After it's canceled, you can no longer update the subscription or its [metadata](https://docs.stripe.com/metadata).
//
// Any pending invoice items that you've created are still charged at the end of the period, unless manually [deleted](https://docs.stripe.com/api#delete_invoiceitem). If you've set the subscription to cancel at the end of the period, any pending prorations are also left in place and collected at the end of the period. But if the subscription is set to cancel immediately, pending prorations are removed if invoice_now and prorate are both set to true.
//
// By default, upon subscription cancellation, Stripe stops automatic collection of all finalized invoices for the customer. This is intended to prevent unexpected payment attempts after the customer has canceled a subscription. However, you can resume automatic collection of the invoices manually after subscription cancellation to have us proceed. Or, you could check for unpaid invoices before allowing the customer to cancel the subscription at all.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.SubscriptionCancelParams) (*stripe.Subscription, error) {
	path := stripe.FormatURLPath("/v1/subscriptions/%s", id)
	subscription := &stripe.Subscription{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, subscription)
	return subscription, err
}

// Removes the currently applied discount on a subscription.
func DeleteDiscount(id string, params *stripe.SubscriptionDeleteDiscountParams) (*stripe.Subscription, error) {
	return getC().DeleteDiscount(id, params)
}

// Removes the currently applied discount on a subscription.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) DeleteDiscount(id string, params *stripe.SubscriptionDeleteDiscountParams) (*stripe.Subscription, error) {
	path := stripe.FormatURLPath("/v1/subscriptions/%s/discount", id)
	subscription := &stripe.Subscription{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, subscription)
	return subscription, err
}

// Upgrade the billing_mode of an existing subscription.
func Migrate(id string, params *stripe.SubscriptionMigrateParams) (*stripe.Subscription, error) {
	return getC().Migrate(id, params)
}

// Upgrade the billing_mode of an existing subscription.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Migrate(id string, params *stripe.SubscriptionMigrateParams) (*stripe.Subscription, error) {
	path := stripe.FormatURLPath("/v1/subscriptions/%s/migrate", id)
	subscription := &stripe.Subscription{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, subscription)
	return subscription, err
}

// Initiates resumption of a paused subscription, optionally resetting the billing cycle anchor and creating prorations. If no resumption invoice is generated, the subscription becomes active immediately. If a resumption invoice is generated, the subscription remains paused until the invoice is paid or marked uncollectible. If the invoice is not paid by the expiration date, it is voided and the subscription remains paused.
func Resume(id string, params *stripe.SubscriptionResumeParams) (*stripe.Subscription, error) {
	return getC().Resume(id, params)
}

// Initiates resumption of a paused subscription, optionally resetting the billing cycle anchor and creating prorations. If no resumption invoice is generated, the subscription becomes active immediately. If a resumption invoice is generated, the subscription remains paused until the invoice is paid or marked uncollectible. If the invoice is not paid by the expiration date, it is voided and the subscription remains paused.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Resume(id string, params *stripe.SubscriptionResumeParams) (*stripe.Subscription, error) {
	path := stripe.FormatURLPath("/v1/subscriptions/%s/resume", id)
	subscription := &stripe.Subscription{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, subscription)
	return subscription, err
}

// By default, returns a list of subscriptions that have not been canceled. In order to list canceled subscriptions, specify status=canceled.
func List(params *stripe.SubscriptionListParams) *Iter {
	return getC().List(params)
}

// By default, returns a list of subscriptions that have not been canceled. In order to list canceled subscriptions, specify status=canceled.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.SubscriptionListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.SubscriptionList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/subscriptions", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for subscriptions.
type Iter struct {
	*stripe.Iter
}

// Subscription returns the subscription which the iterator is currently pointing to.
func (i *Iter) Subscription() *stripe.Subscription {
	return i.Current().(*stripe.Subscription)
}

// SubscriptionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) SubscriptionList() *stripe.SubscriptionList {
	return i.List().(*stripe.SubscriptionList)
}

// Search for subscriptions you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
func Search(params *stripe.SubscriptionSearchParams) *SearchIter {
	return getC().Search(params)
}

// Search for subscriptions you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Search(params *stripe.SubscriptionSearchParams) *SearchIter {
	return &SearchIter{
		SearchIter: stripe.GetSearchIter(params, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.SearchContainer, error) {
			list := &stripe.SubscriptionSearchResult{}
			err := c.B.CallRaw(http.MethodGet, "/v1/subscriptions/search", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// SearchIter is an iterator for subscriptions.
type SearchIter struct {
	*stripe.SearchIter
}

// Subscription returns the subscription which the iterator is currently pointing to.
func (i *SearchIter) Subscription() *stripe.Subscription {
	return i.Current().(*stripe.Subscription)
}

// SubscriptionSearchResult returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *SearchIter) SubscriptionSearchResult() *stripe.SubscriptionSearchResult {
	return i.SearchResult().(*stripe.SubscriptionSearchResult)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
