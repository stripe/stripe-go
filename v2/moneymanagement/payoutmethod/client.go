//
//
// File generated from our OpenAPI spec
//
//

// Package payoutmethod provides the payoutmethod related APIs
package payoutmethod

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke payoutmethod related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a PayoutMethod object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementPayoutMethodParams) (*stripe.V2MoneyManagementPayoutMethod, error) {
	path := stripe.FormatURLPath("/v2/money_management/payout_methods/%s", id)
	payoutmethod := &stripe.V2MoneyManagementPayoutMethod{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, payoutmethod)
	return payoutmethod, err
}

// Archive a `PayoutMethod`. Archiving prevents the Payout Method from being used for outbound payments
// or transfers and omits it from normal list results. To restore list visibility, use the
// [unarchive endpoint](https://docs.stripe.com/api/v2/money-management/payout-methods/unarchive).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Archive(id string, params *stripe.V2MoneyManagementPayoutMethodArchiveParams) (*stripe.V2MoneyManagementPayoutMethod, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/payout_methods/%s/archive", id)
	payoutmethod := &stripe.V2MoneyManagementPayoutMethod{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, payoutmethod)
	return payoutmethod, err
}

// Disable a `PayoutMethod`. Disabling temporarily prevents the Payout Method from being used for outbound
// payments or transfers while keeping it in normal list results. To re-enable it, complete setup again by
// [creating an Outbound Setup Intent](https://docs.stripe.com/api/v2/money-management/outbound-setup-intents/create).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Disable(id string, params *stripe.V2MoneyManagementPayoutMethodDisableParams) (*stripe.V2MoneyManagementPayoutMethod, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/payout_methods/%s/disable", id)
	payoutmethod := &stripe.V2MoneyManagementPayoutMethod{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, payoutmethod)
	return payoutmethod, err
}

// Unarchive a `PayoutMethod`. Unarchiving restores the Payout Method to normal list results and clears
// only its archived state. It doesn't guarantee that the Payout Method can be used.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Unarchive(id string, params *stripe.V2MoneyManagementPayoutMethodUnarchiveParams) (*stripe.V2MoneyManagementPayoutMethod, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/payout_methods/%s/unarchive", id)
	payoutmethod := &stripe.V2MoneyManagementPayoutMethod{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, payoutmethod)
	return payoutmethod, err
}

// List objects that adhere to the PayoutMethod interface.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementPayoutMethodListParams) stripe.Seq2[*stripe.V2MoneyManagementPayoutMethod, error] {
	if listParams == nil {
		listParams = &stripe.V2MoneyManagementPayoutMethodListParams{}
	}
	return stripe.NewV2List("/v2/money_management/payout_methods", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementPayoutMethod], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementPayoutMethod]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
