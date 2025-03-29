//
//
// File generated from our OpenAPI spec
//
//

// Package payoutmethod provides the payoutmethod related APIs
package payoutmethod

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v81"
)

// Client is used to invoke payoutmethod related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a PayoutMethod object.
func (c Client) Get(id string, params *stripe.V2MoneyManagementPayoutMethodParams) (*stripe.V2MoneyManagementPayoutMethod, error) {
	path := stripe.FormatURLPath("/v2/money_management/payout_methods/%s", id)
	payoutmethod := &stripe.V2MoneyManagementPayoutMethod{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, payoutmethod)
	return payoutmethod, err
}

// Archive a PayoutMethod object. Archived objects cannot be used as payout methods
// and will not appear in the payout method list.
func (c Client) Archive(id string, params *stripe.V2MoneyManagementPayoutMethodArchiveParams) (*stripe.V2MoneyManagementPayoutMethod, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/payout_methods/%s/archive", id)
	payoutmethod := &stripe.V2MoneyManagementPayoutMethod{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, payoutmethod)
	return payoutmethod, err
}

// Unarchive an PayoutMethod object.
func (c Client) Unarchive(id string, params *stripe.V2MoneyManagementPayoutMethodUnarchiveParams) (*stripe.V2MoneyManagementPayoutMethod, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/payout_methods/%s/unarchive", id)
	payoutmethod := &stripe.V2MoneyManagementPayoutMethod{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, payoutmethod)
	return payoutmethod, err
}

// List objects that adhere to the PayoutMethod interface.
func (c Client) All(listParams *stripe.V2MoneyManagementPayoutMethodListParams) stripe.Seq2[stripe.V2MoneyManagementPayoutMethod, error] {
	return stripe.NewV2List("/v2/money_management/payout_methods", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[stripe.V2MoneyManagementPayoutMethod], error) {
		page := &stripe.V2Page[stripe.V2MoneyManagementPayoutMethod]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
