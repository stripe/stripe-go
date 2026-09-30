//
//
// File generated from our OpenAPI spec
//
//

// Package earnedcredit provides the earnedcredit related APIs
package earnedcredit

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke earnedcredit related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves an EarnedCredit.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementEarnedCreditParams) (*stripe.V2MoneyManagementEarnedCredit, error) {
	path := stripe.FormatURLPath("/v2/money_management/earned_credits/%s", id)
	earnedcredit := &stripe.V2MoneyManagementEarnedCredit{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, earnedcredit)
	return earnedcredit, err
}

// Returns a list of EarnedCredits.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementEarnedCreditListParams) stripe.Seq2[*stripe.V2MoneyManagementEarnedCredit, error] {
	if listParams == nil {
		listParams = &stripe.V2MoneyManagementEarnedCreditListParams{}
	}
	return stripe.NewV2List("/v2/money_management/earned_credits", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementEarnedCredit], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementEarnedCredit]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
