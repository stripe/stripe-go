//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate describes a
// type of customer updates that may be supported on a portal configuration
type BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate string

// List of values that BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate can take.
const (
	BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateAddress  BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate = "address"
	BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateEmail    BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate = "email"
	BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdatePhone    BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate = "phone"
	BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateShipping BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate = "shipping"
	BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateTaxID    BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate = "tax_id"
)

// BillingPortalConfigurationFeaturesSubscriptionCancelMode describes whether
// to cancel subscriptions immediately or at the end of the billing period.
type BillingPortalConfigurationFeaturesSubscriptionCancelMode string

// List of values that BillingPortalConfigurationFeaturesSubscriptionCancelMode can take.
const (
	BillingPortalConfigurationFeaturesSubscriptionCancelModeAtPeriodEnd BillingPortalConfigurationFeaturesSubscriptionCancelMode = "at_period_end"
	BillingPortalConfigurationFeaturesSubscriptionCancelModeImmediately BillingPortalConfigurationFeaturesSubscriptionCancelMode = "immediately"
)

// BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior describes
// whether to create prorations when canceling subscriptions.
type BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior string

// List of values that BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior can take.
const (
	BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehaviorAlwaysInvoice    BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior = "always_invoice"
	BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehaviorCreateProrations BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior = "create_prorations"
	BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehaviorNone             BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior = "none"
)

// BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdate describes
// a type of subscription update that may be supported on a portal configuration.
type BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdate string

// List of values that BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdate can take.
const (
	BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdatePrice         BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdate = "price"
	BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdatePromotionCode BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdate = "promotion_code"
	BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdateQuantity      BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdate = "quantity"
)

// BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior determines
// how to handle prorations resulting from subscription updates.
type BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior string

// List of values that BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior can take.
const (
	BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehaviorAlwaysInvoice    BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior = "always_invoice"
	BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehaviorCreateProrations BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior = "create_prorations"
	BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehaviorNone             BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior = "none"
)

// BillingPortalConfigurationListParams is the set of parameters that can be
// used when listing portal configurations.
type BillingPortalConfigurationListParams struct {
	ListParams `form:"*"`
	Active     *bool `form:"active"`
	IsDefault  *bool `form:"is_default"`
}

// BillingPortalConfigurationBusinessProfileParams lets you pass the business
// profile details associated with a portal configuration.
type BillingPortalConfigurationBusinessProfileParams struct {
	Headline          *string `form:"headline"`
	PrivacyPolicyURL  *string `form:"privacy_policy_url"`
	TermsOfServiceURL *string `form:"terms_of_service_url"`
}

// BillingPortalConfigurationFeaturesCustomerUpdateParams lets you pass the
// customer update details on a portal configuration.
type BillingPortalConfigurationFeaturesCustomerUpdateParams struct {
	AllowedUpdates []*string `form:"allowed_updates"`
	Enabled        *bool     `form:"enabled"`
}

// BillingPortalConfigurationFeaturesInvoiceHistoryParams lets you pass the
// invoice history details on a portal configuration.
type BillingPortalConfigurationFeaturesInvoiceHistoryParams struct {
	Enabled *bool `form:"enabled"`
}

// BillingPortalConfigurationFeaturesPaymentMethodUpdateParams lets you pass
// the payment method update details on a portal configuration.
type BillingPortalConfigurationFeaturesPaymentMethodUpdateParams struct {
	Enabled *bool `form:"enabled"`
}

// BillingPortalConfigurationFeaturesSubscriptionCancelParams lets you pass the
// subscription cancel deetails on a portal configuration.
type BillingPortalConfigurationFeaturesSubscriptionCancelParams struct {
	Enabled           *bool   `form:"enabled"`
	Mode              *string `form:"mode"`
	ProrationBehavior *string `form:"proration_behavior"`
}

// BillingPortalConfigurationFeaturesSubscriptionPauseParams lets you pass details on the
// subscription pause on a portal configuration.
type BillingPortalConfigurationFeaturesSubscriptionPauseParams struct {
	Enabled *bool `form:"enabled"`
}

// BillingPortalConfigurationFeaturesSubscriptionUpdateProductParams lets you
// pass product details on the subscription update on a portal configuration.
type BillingPortalConfigurationFeaturesSubscriptionUpdateProductParams struct {
	Prices  []*string `form:"prices"`
	Product *string   `form:"product"`
}

// BillingPortalConfigurationFeaturesSubscriptionUpdateParams lets you pass
// subscription update details on a portal configuration.
type BillingPortalConfigurationFeaturesSubscriptionUpdateParams struct {
	DefaultAllowedUpdates []*string                                                            `form:"default_allowed_updates"`
	Enabled               *bool                                                                `form:"enabled"`
	Products              []*BillingPortalConfigurationFeaturesSubscriptionUpdateProductParams `form:"products"`
	ProrationBehavior     *string                                                              `form:"proration_behavior"`
}

// BillingPortalConfigurationFeaturesParams lets you pass details about the
// features available in the portal.
type BillingPortalConfigurationFeaturesParams struct {
	CustomerUpdate      *BillingPortalConfigurationFeaturesCustomerUpdateParams      `form:"customer_update"`
	InvoiceHistory      *BillingPortalConfigurationFeaturesInvoiceHistoryParams      `form:"invoice_history"`
	PaymentMethodUpdate *BillingPortalConfigurationFeaturesPaymentMethodUpdateParams `form:"payment_method_update"`
	SubscriptionCancel  *BillingPortalConfigurationFeaturesSubscriptionCancelParams  `form:"subscription_cancel"`
	SubscriptionPause   *BillingPortalConfigurationFeaturesSubscriptionPauseParams   `form:"subscription_pause"`
	SubscriptionUpdate  *BillingPortalConfigurationFeaturesSubscriptionUpdateParams  `form:"subscription_update"`
}

// BillingPortalConfigurationParams is the set of parameters that can be passed
// when creating or updating a portal configuration.
type BillingPortalConfigurationParams struct {
	Params           `form:"*"`
	Active           *bool                                            `form:"active"`
	BusinessProfile  *BillingPortalConfigurationBusinessProfileParams `form:"business_profile"`
	DefaultReturnURL *string                                          `form:"default_return_url"`
	Features         *BillingPortalConfigurationFeaturesParams        `form:"features"`
}

// BillingPortalConfiguration is a configuration that describes the
// functionality and behavior of a PortalSession. For more details see
// https://stripe.com/docs/api/customer_portal.
type BillingPortalConfiguration struct {
	APIResource
	Active           bool                                       `json:"active"`
	Application      string                                     `json:"application"`
	BusinessProfile  *BillingPortalConfigurationBusinessProfile `json:"business_profile"`
	Created          int64                                      `json:"created"`
	DefaultReturnURL string                                     `json:"default_return_url"`
	Features         *BillingPortalConfigurationFeatures        `json:"features"`
	ID               string                                     `json:"id"`
	IsDefault        bool                                       `json:"is_default"`
	Livemode         bool                                       `json:"livemode"`
	Object           string                                     `json:"object"`
	Updated          int64                                      `json:"updated"`
}

// BillingPortalConfigurationBusinessProfile represents the business profile
// details on a portal configuration.
type BillingPortalConfigurationBusinessProfile struct {
	Headline          string `json:"headline"`
	PrivacyPolicyURL  string `json:"privacy_policy_url"`
	TermsOfServiceURL string `json:"terms_of_service_url"`
}

// BillingPortalConfigurationFeaturesCustomerUpdate represents the customer
// update details on a portal configuration.
type BillingPortalConfigurationFeaturesCustomerUpdate struct {
	AllowedUpdates []BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate `json:"allowed_updates"`
	Enabled        bool                                                            `json:"enabled"`
}

// BillingPortalConfigurationFeaturesInvoiceHistory represents the invoice
// history details on a portal configuration.
type BillingPortalConfigurationFeaturesInvoiceHistory struct {
	Enabled bool `json:"enabled"`
}

// BillingPortalConfigurationFeaturesPaymentMethodUpdate represents the payment
// method update details on a portal configuration.
type BillingPortalConfigurationFeaturesPaymentMethodUpdate struct {
	Enabled bool `json:"enabled"`
}

// BillingPortalConfigurationFeaturesSubscriptionCancel represents the
// subscription cancel details on a portal configuration.
type BillingPortalConfigurationFeaturesSubscriptionCancel struct {
	Enabled           bool                                                                  `json:"enabled"`
	Mode              BillingPortalConfigurationFeaturesSubscriptionCancelMode              `json:"mode"`
	ProrationBehavior BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior `json:"proration_behavior"`
}

// BillingPortalConfigurationFeaturesSubscriptionPause lets you pass pause details
// on the subscription update on a portal configuration.
type BillingPortalConfigurationFeaturesSubscriptionPause struct {
	Enabled bool `json:"enabled"`
}

// BillingPortalConfigurationFeaturesSubscriptionUpdateProduct represents the
// subscription update details on a portal configuration.
type BillingPortalConfigurationFeaturesSubscriptionUpdateProduct struct {
	Prices  []string `json:"prices"`
	Product string   `json:"product"`
}

// BillingPortalConfigurationFeaturesSubscriptionUpdate represents the
// subscription update details on a portal configuration.
type BillingPortalConfigurationFeaturesSubscriptionUpdate struct {
	DefaultAllowedUpdates []BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdate `json:"default_allowed_updates"`
	Enabled               bool                                                                       `json:"enabled"`
	Products              []*BillingPortalConfigurationFeaturesSubscriptionUpdateProduct             `json:"products"`
	ProrationBehavior     BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior      `json:"proration_behavior"`
}

// BillingPortalConfigurationFeatures represents details about the features
// enabled in the portal configuration.
type BillingPortalConfigurationFeatures struct {
	CustomerUpdate      *BillingPortalConfigurationFeaturesCustomerUpdate      `json:"customer_update"`
	InvoiceHistory      *BillingPortalConfigurationFeaturesInvoiceHistory      `json:"invoice_history"`
	PaymentMethodUpdate *BillingPortalConfigurationFeaturesPaymentMethodUpdate `json:"payment_method_update"`
	SubscriptionCancel  *BillingPortalConfigurationFeaturesSubscriptionCancel  `json:"subscription_cancel"`
	SubscriptionPause   *BillingPortalConfigurationFeaturesSubscriptionPause   `json:"subscription_pause"`
	SubscriptionUpdate  *BillingPortalConfigurationFeaturesSubscriptionUpdate  `json:"subscription_update"`
}

// BillingPortalConfigurationList is a list of portal configurations as
// returned from a list endpoint.
type BillingPortalConfigurationList struct {
	APIResource
	ListMeta
	Data []*BillingPortalConfiguration `json:"data"`
}

// UnmarshalJSON handles deserialization of a BillingPortalConfiguration.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (b *BillingPortalConfiguration) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		b.ID = id
		return nil
	}

	type billingPortalConfiguration BillingPortalConfiguration
	var v billingPortalConfiguration
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*b = BillingPortalConfiguration(v)
	return nil
}
