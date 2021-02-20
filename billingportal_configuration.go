package stripe

import "encoding/json"

// BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate TODO docstring
type BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate string

// List of values that BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate can take.
const (
	BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateAddress  BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate = "address"
	BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateEmail    BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate = "email"
	BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdatePhone    BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate = "phone"
	BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateShipping BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate = "shipping"
	BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateTaxID    BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate = "tax_id"
)

// BillingPortalConfigurationFeaturesSubscriptionCancelMode TODO docstring
type BillingPortalConfigurationFeaturesSubscriptionCancelMode string

// List of values that BillingPortalConfigurationFeaturesSubscriptionCancelMode can take.
const (
	BillingPortalConfigurationFeaturesSubscriptionCancelModeAtPeriodEnd BillingPortalConfigurationFeaturesSubscriptionCancelMode = "at_period_end"
	BillingPortalConfigurationFeaturesSubscriptionCancelModeImmediately BillingPortalConfigurationFeaturesSubscriptionCancelMode = "immediately"
)

// BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior TODO docstring
type BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior string

// List of values that BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior can take.
const (
	BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehaviorAlwaysInvoice    BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior = "always_invoice"
	BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehaviorCreateProrations BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior = "create_prorations"
	BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehaviorNone             BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior = "none"
)

// BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdate TODO docstring
type BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdate string

// List of values that BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdate can take.
const (
	BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdatePrice         BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdate = "price"
	BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdatePromotionCode BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdate = "promotion_code"
	BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdateQuantity      BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdate = "quantity"
)

// BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior TODO docstring
type BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior string

// List of values that BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior can take.
const (
	BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehaviorAlwaysInvoice    BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior = "always_invoice"
	BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehaviorCreateProrations BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior = "create_prorations"
	BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehaviorNone             BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior = "none"
)

// BillingPortalConfigurationListParams TODO docstring
type BillingPortalConfigurationListParams struct {
	ListParams `form:"*"`
	Active     *bool `form:"active"`
	IsDefault  *bool `form:"is_default"`
}

// BillingPortalConfigurationBusinessProfileParams TODO docstring
type BillingPortalConfigurationBusinessProfileParams struct {
	Headline          *string `form:"headline"`
	PrivacyPolicyURL  *string `form:"privacy_policy_url"`
	TermsOfServiceURL *string `form:"terms_of_service_url"`
}

// BillingPortalConfigurationFeaturesCustomerUpdateParams TODO docstring
type BillingPortalConfigurationFeaturesCustomerUpdateParams struct {
	AllowedUpdates []*string `form:"allowed_updates"`
	Enabled        *bool     `form:"enabled"`
}

// BillingPortalConfigurationFeaturesInvoiceHistoryParams TODO docstring
type BillingPortalConfigurationFeaturesInvoiceHistoryParams struct {
	Enabled *bool `form:"enabled"`
}

// BillingPortalConfigurationFeaturesPaymentMethodUpdateParams TODO docstring
type BillingPortalConfigurationFeaturesPaymentMethodUpdateParams struct {
	Enabled *bool `form:"enabled"`
}

// BillingPortalConfigurationFeaturesSubscriptionCancelParams TODO docstring
type BillingPortalConfigurationFeaturesSubscriptionCancelParams struct {
	Enabled           *bool   `form:"enabled"`
	Mode              *string `form:"mode"`
	ProrationBehavior *string `form:"proration_behavior"`
}

// BillingPortalConfigurationFeaturesSubscriptionUpdateProductParams TODO docstring
type BillingPortalConfigurationFeaturesSubscriptionUpdateProductParams struct {
	Prices  []*string `form:"prices"`
	Product *string   `form:"product"`
}

// BillingPortalConfigurationFeaturesSubscriptionUpdateParams TODO docstring
type BillingPortalConfigurationFeaturesSubscriptionUpdateParams struct {
	DefaultAllowedUpdates []*string                                                            `form:"default_allowed_updates"`
	Enabled               *bool                                                                `form:"enabled"`
	Products              []*BillingPortalConfigurationFeaturesSubscriptionUpdateProductParams `form:"products"`
	ProrationBehavior     *string                                                              `form:"proration_behavior"`
}

// BillingPortalConfigurationFeaturesParams TODO docstring
type BillingPortalConfigurationFeaturesParams struct {
	CustomerUpdate      *BillingPortalConfigurationFeaturesCustomerUpdateParams      `form:"customer_update"`
	InvoiceHistory      *BillingPortalConfigurationFeaturesInvoiceHistoryParams      `form:"invoice_history"`
	PaymentMethodUpdate *BillingPortalConfigurationFeaturesPaymentMethodUpdateParams `form:"payment_method_update"`
	SubscriptionCancel  *BillingPortalConfigurationFeaturesSubscriptionCancelParams  `form:"subscription_cancel"`
	SubscriptionUpdate  *BillingPortalConfigurationFeaturesSubscriptionUpdateParams  `form:"subscription_update"`
}

// BillingPortalConfigurationParams TODO docstring
type BillingPortalConfigurationParams struct {
	Params           `form:"*"`
	Active           *bool                                            `form:"active"`
	BusinessProfile  *BillingPortalConfigurationBusinessProfileParams `form:"business_profile"`
	DefaultReturnURL *string                                          `form:"default_return_url"`
	Features         *BillingPortalConfigurationFeaturesParams        `form:"features"`
}

// BillingPortalConfiguration TODO docstring
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

// BillingPortalConfigurationBusinessProfile TODO docstring
type BillingPortalConfigurationBusinessProfile struct {
	Headline          string `json:"headline"`
	PrivacyPolicyURL  string `json:"privacy_policy_url"`
	TermsOfServiceURL string `json:"terms_of_service_url"`
}

// BillingPortalConfigurationFeaturesCustomerUpdate TODO docstring
type BillingPortalConfigurationFeaturesCustomerUpdate struct {
	AllowedUpdates []*BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdate `json:"allowed_updates"`
	Enabled        bool                                                             `json:"enabled"`
}

// BillingPortalConfigurationFeaturesInvoiceHistory TODO docstring
type BillingPortalConfigurationFeaturesInvoiceHistory struct {
	Enabled bool `json:"enabled"`
}

// BillingPortalConfigurationFeaturesPaymentMethodUpdate TODO docstring
type BillingPortalConfigurationFeaturesPaymentMethodUpdate struct {
	Enabled bool `json:"enabled"`
}

// BillingPortalConfigurationFeaturesSubscriptionCancel TODO docstring
type BillingPortalConfigurationFeaturesSubscriptionCancel struct {
	Enabled           bool                                                                   `json:"enabled"`
	Mode              *BillingPortalConfigurationFeaturesSubscriptionCancelMode              `json:"mode"`
	ProrationBehavior *BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior `json:"proration_behavior"`
}

// BillingPortalConfigurationFeaturesSubscriptionUpdateProduct TODO docstring
type BillingPortalConfigurationFeaturesSubscriptionUpdateProduct struct {
	Prices  []string `json:"prices"`
	Product string   `json:"product"`
}

// BillingPortalConfigurationFeaturesSubscriptionUpdate TODO docstring
type BillingPortalConfigurationFeaturesSubscriptionUpdate struct {
	DefaultAllowedUpdates []*BillingPortalConfigurationFeaturesSubscriptionUpdateDefaultAllowedUpdate `json:"default_allowed_updates"`
	Enabled               bool                                                                        `json:"enabled"`
	Products              []*BillingPortalConfigurationFeaturesSubscriptionUpdateProduct              `json:"products"`
	ProrationBehavior     *BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior      `json:"proration_behavior"`
}

// BillingPortalConfigurationFeatures TODO docstring
type BillingPortalConfigurationFeatures struct {
	CustomerUpdate      *BillingPortalConfigurationFeaturesCustomerUpdate      `json:"customer_update"`
	InvoiceHistory      *BillingPortalConfigurationFeaturesInvoiceHistory      `json:"invoice_history"`
	PaymentMethodUpdate *BillingPortalConfigurationFeaturesPaymentMethodUpdate `json:"payment_method_update"`
	SubscriptionCancel  *BillingPortalConfigurationFeaturesSubscriptionCancel  `json:"subscription_cancel"`
	SubscriptionUpdate  *BillingPortalConfigurationFeaturesSubscriptionUpdate  `json:"subscription_update"`
}

// BillingPortalConfigurationList TODO docstring
type BillingPortalConfigurationList struct {
	APIResource
	ListMeta
	Data []*BillingPortalConfiguration `json:"data"`
}

// UnmarshalJSON handles deserialization of a BillingPortalConfiguration.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *BillingPortalConfiguration) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type billingPortalConfiguration BillingPortalConfiguration
	var v billingPortalConfiguration
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = BillingPortalConfiguration(v)
	return nil
}
