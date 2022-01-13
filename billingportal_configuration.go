//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The types of customer updates that are supported. When empty, customers are not updateable.
type BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate string

// List of values that BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate can take
const (
	BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateAddress  BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate = "address"
	BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateEmail    BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate = "email"
	BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdatePhone    BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate = "phone"
	BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateShipping BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate = "shipping"
	BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateTaxID    BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate = "tax_id"
)

// Which cancellation reasons will be given as options to the customer.
type BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonOption string

// List of values that BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonOption can take
const (
	BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonOptionCustomerService BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonOption = "customer_service"
	BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonOptionLowQuality      BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonOption = "low_quality"
	BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonOptionMissingFeatures BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonOption = "missing_features"
	BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonOptionOther           BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonOption = "other"
	BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonOptionSwitchedService BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonOption = "switched_service"
	BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonOptionTooComplex      BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonOption = "too_complex"
	BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonOptionTooExpensive    BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonOption = "too_expensive"
	BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonOptionUnused          BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonOption = "unused"
)

// Whether to cancel subscriptions immediately or at the end of the billing period.
type BillingPortalConfigurationFeaturesSubscriptionCancelMode string

// List of values that BillingPortalConfigurationFeaturesSubscriptionCancelMode can take
const (
	BillingPortalConfigurationFeaturesSubscriptionCancelModeAtPeriodEnd BillingPortalConfigurationFeaturesSubscriptionCancelMode = "at_period_end"
	BillingPortalConfigurationFeaturesSubscriptionCancelModeImmediately BillingPortalConfigurationFeaturesSubscriptionCancelMode = "immediately"
)

// Whether to create prorations when canceling subscriptions. Possible values are `none` and `create_prorations`.
type BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior string

// List of values that BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior can take
const (
	BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehaviorAlwaysInvoice    BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior = "always_invoice"
	BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehaviorCreateProrations BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior = "create_prorations"
	BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehaviorNone             BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior = "none"
)

// The types of subscription updates that are supported for items listed in the `products` attribute. When empty, subscriptions are not updateable.
type BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdate string

// List of values that BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdate can take
const (
	BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdatePrice         BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdate = "price"
	BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdatePromotionCode BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdate = "promotion_code"
	BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdateQuantity      BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdate = "quantity"
)

// Determines how to handle prorations resulting from subscription updates. Valid values are `none`, `create_prorations`, and `always_invoice`.
type BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior string

// List of values that BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior can take
const (
	BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehaviorAlwaysInvoice    BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior = "always_invoice"
	BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehaviorCreateProrations BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior = "create_prorations"
	BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehaviorNone             BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior = "none"
)

// Returns a list of configurations that describe the functionality of the customer portal.
type BillingPortalConfigurationListParams struct {
	ListParams `form:"*"`
	Active     *bool `form:"active"`
	IsDefault  *bool `form:"is_default"`
}

// The business information shown to customers in the portal.
type BillingPortalConfigurationBusinessProfileParams struct {
	Headline          *string `form:"headline"`
	PrivacyPolicyURL  *string `form:"privacy_policy_url"`
	TermsOfServiceURL *string `form:"terms_of_service_url"`
}

// Information about updating the customer details in the portal.
type BillingPortalConfigurationFeaturesCustomerUpdateParams struct {
	AllowedUpdates []*string `form:"allowed_updates"`
	Enabled        *bool     `form:"enabled"`
}

// Information about showing the billing history in the portal.
type BillingPortalConfigurationFeaturesInvoiceHistoryParams struct {
	Enabled *bool `form:"enabled"`
}

// Information about updating payment methods in the portal.
type BillingPortalConfigurationFeaturesPaymentMethodUpdateParams struct {
	Enabled *bool `form:"enabled"`
}

// Whether the cancellation reasons will be collected in the portal and which options are exposed to the customer
type BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonParams struct {
	Enabled *bool     `form:"enabled"`
	Options []*string `form:"options"`
}

// Information about canceling subscriptions in the portal.
type BillingPortalConfigurationFeaturesSubscriptionCancelParams struct {
	CancellationReason *BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonParams `form:"cancellation_reason"`
	Enabled            *bool                                                                         `form:"enabled"`
	Mode               *string                                                                       `form:"mode"`
	ProrationBehavior  *string                                                                       `form:"proration_behavior"`
}

// Information about pausing subscriptions in the portal.
type BillingPortalConfigurationFeaturesSubscriptionPauseParams struct {
	Enabled *bool `form:"enabled"`
}

// The list of products that support subscription updates.
type BillingPortalConfigurationFeaturesSubscriptionUpdateProductParams struct {
	Prices  []*string `form:"prices"`
	Product *string   `form:"product"`
}

// Information about updating subscriptions in the portal.
type BillingPortalConfigurationFeaturesSubscriptionUpdateParams struct {
	DefaultAllowedUpdates []*string                                                            `form:"default_allowed_updates"`
	Enabled               *bool                                                                `form:"enabled"`
	Products              []*BillingPortalConfigurationFeaturesSubscriptionUpdateProductParams `form:"products"`
	ProrationBehavior     *string                                                              `form:"proration_behavior"`
}

// Information about the features available in the portal.
type BillingPortalConfigurationFeaturesParams struct {
	CustomerUpdate      *BillingPortalConfigurationFeaturesCustomerUpdateParams      `form:"customer_update"`
	InvoiceHistory      *BillingPortalConfigurationFeaturesInvoiceHistoryParams      `form:"invoice_history"`
	PaymentMethodUpdate *BillingPortalConfigurationFeaturesPaymentMethodUpdateParams `form:"payment_method_update"`
	SubscriptionCancel  *BillingPortalConfigurationFeaturesSubscriptionCancelParams  `form:"subscription_cancel"`
	SubscriptionPause   *BillingPortalConfigurationFeaturesSubscriptionPauseParams   `form:"subscription_pause"`
	SubscriptionUpdate  *BillingPortalConfigurationFeaturesSubscriptionUpdateParams  `form:"subscription_update"`
}

// Creates a configuration that describes the functionality and behavior of a PortalSession
type BillingPortalConfigurationParams struct {
	Params           `form:"*"`
	Active           *bool                                            `form:"active"`
	BusinessProfile  *BillingPortalConfigurationBusinessProfileParams `form:"business_profile"`
	DefaultReturnURL *string                                          `form:"default_return_url"`
	Features         *BillingPortalConfigurationFeaturesParams        `form:"features"`
}
type BillingPortalConfigurationBusinessProfile struct {
	Headline          string `json:"headline"`
	PrivacyPolicyURL  string `json:"privacy_policy_url"`
	TermsOfServiceURL string `json:"terms_of_service_url"`
}
type BillingPortalConfigurationFeaturesCustomerUpdate struct {
	AllowedUpdates []BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate `json:"allowed_updates"`
	Enabled        bool                                                            `json:"enabled"`
}
type BillingPortalConfigurationFeaturesInvoiceHistory struct {
	Enabled bool `json:"enabled"`
}
type BillingPortalConfigurationFeaturesPaymentMethodUpdate struct {
	Enabled bool `json:"enabled"`
}
type BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReason struct {
	Enabled bool                                                                           `json:"enabled"`
	Options []BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonOption `json:"options"`
}
type BillingPortalConfigurationFeaturesSubscriptionCancel struct {
	CancellationReason *BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReason `json:"cancellation_reason"`
	Enabled            bool                                                                    `json:"enabled"`
	Mode               BillingPortalConfigurationFeaturesSubscriptionCancelMode                `json:"mode"`
	ProrationBehavior  BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior   `json:"proration_behavior"`
}
type BillingPortalConfigurationFeaturesSubscriptionPause struct {
	Enabled bool `json:"enabled"`
}

// The list of products that support subscription updates.
type BillingPortalConfigurationFeaturesSubscriptionUpdateProduct struct {
	Prices  []string `json:"prices"`
	Product string   `json:"product"`
}
type BillingPortalConfigurationFeaturesSubscriptionUpdate struct {
	DefaultAllowedUpdates []BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdate `json:"default_allowed_updates"`
	Enabled               bool                                                                       `json:"enabled"`
	Products              []*BillingPortalConfigurationFeaturesSubscriptionUpdateProduct             `json:"products"`
	ProrationBehavior     BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior      `json:"proration_behavior"`
}
type BillingPortalConfigurationFeatures struct {
	CustomerUpdate      *BillingPortalConfigurationFeaturesCustomerUpdate      `json:"customer_update"`
	InvoiceHistory      *BillingPortalConfigurationFeaturesInvoiceHistory      `json:"invoice_history"`
	PaymentMethodUpdate *BillingPortalConfigurationFeaturesPaymentMethodUpdate `json:"payment_method_update"`
	SubscriptionCancel  *BillingPortalConfigurationFeaturesSubscriptionCancel  `json:"subscription_cancel"`
	SubscriptionPause   *BillingPortalConfigurationFeaturesSubscriptionPause   `json:"subscription_pause"`
	SubscriptionUpdate  *BillingPortalConfigurationFeaturesSubscriptionUpdate  `json:"subscription_update"`
}

// A portal configuration describes the functionality and behavior of a portal session.
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
	Metadata         map[string]string                          `json:"metadata"`
	Object           string                                     `json:"object"`
	Updated          int64                                      `json:"updated"`
}

// BillingPortalConfigurationList is a list of Configurations as retrieved from a list endpoint.
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
