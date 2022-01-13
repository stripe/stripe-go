//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Creates a session of the customer portal.
type BillingPortalSessionParams struct {
	Params        `form:"*"`
	Configuration *string `form:"configuration"`
	Customer      *string `form:"customer"`
	Locale        *string `form:"locale"`
	OnBehalfOf    *string `form:"on_behalf_of"`
	ReturnURL     *string `form:"return_url"`
}

// The Billing customer portal is a Stripe-hosted UI for subscription and
// billing management.
//
// A portal configuration describes the functionality and features that you
// want to provide to your customers through the portal.
//
// A portal session describes the instantiation of the customer portal for
// a particular customer. By visiting the session's URL, the customer
// can manage their subscriptions and billing details. For security reasons,
// sessions are short-lived and will expire if the customer does not visit the URL.
// Create sessions on-demand when customers intend to manage their subscriptions
// and billing details.
//
// Learn more in the [integration guide](https://stripe.com/docs/billing/subscriptions/integrating-customer-portal).
type BillingPortalSession struct {
	APIResource
	Configuration *BillingPortalConfiguration `json:"configuration"`
	Created       int64                       `json:"created"`
	Customer      string                      `json:"customer"`
	ID            string                      `json:"id"`
	Livemode      bool                        `json:"livemode"`
	Locale        string                      `json:"locale"`
	Object        string                      `json:"object"`
	OnBehalfOf    string                      `json:"on_behalf_of"`
	ReturnURL     string                      `json:"return_url"`
	URL           string                      `json:"url"`
}

// UnmarshalJSON handles deserialization of a BillingPortalSession.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (b *BillingPortalSession) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		b.ID = id
		return nil
	}

	type billingPortalSession BillingPortalSession
	var v billingPortalSession
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*b = BillingPortalSession(v)
	return nil
}
