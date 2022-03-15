//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Creates a session of the customer portal.
type BillingPortalSessionParams struct {
	Params `form:"*"`
	// The ID of an existing [configuration](https://stripe.com/docs/api/customer_portal/configuration) to use for this session, describing its functionality and features. If not specified, the session uses the default configuration.
	Configuration *string `form:"configuration"`
	// The ID of an existing customer.
	Customer *string `form:"customer"`
	// The IETF language tag of the locale Customer Portal is displayed in. If blank or auto, the customer's `preferred_locales` or browser's locale is used.
	Locale *string `form:"locale"`
	// The `on_behalf_of` account to use for this session. When specified, only subscriptions and invoices with this `on_behalf_of` account appear in the portal. For more information, see the [docs](https://stripe.com/docs/connect/charges-transfers#on-behalf-of). Use the [Accounts API](https://stripe.com/docs/api/accounts/object#account_object-settings-branding) to modify the `on_behalf_of` account's branding settings, which the portal displays.
	OnBehalfOf *string `form:"on_behalf_of"`
	// The default URL to redirect customers to when they click on the portal's link to return to your website.
	ReturnURL *string `form:"return_url"`
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
	// The configuration used by this session, describing the features available.
	Configuration *BillingPortalConfiguration `json:"configuration"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The ID of the customer for this session.
	Customer string `json:"customer"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The IETF language tag of the locale Customer Portal is displayed in. If blank or auto, the customer's `preferred_locales` or browser's locale is used.
	Locale string `json:"locale"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The account for which the session was created on behalf of. When specified, only subscriptions and invoices with this `on_behalf_of` account appear in the portal. For more information, see the [docs](https://stripe.com/docs/connect/charges-transfers#on-behalf-of). Use the [Accounts API](https://stripe.com/docs/api/accounts/object#account_object-settings-branding) to modify the `on_behalf_of` account's branding settings, which the portal displays.
	OnBehalfOf string `json:"on_behalf_of"`
	// The URL to redirect customers to when they click on the portal's link to return to your website.
	ReturnURL string `json:"return_url"`
	// The short-lived URL of the session that gives customers access to the customer portal.
	URL string `json:"url"`
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
