//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Controls whether the Payment Element allows the removal of a saved payment method.
type CustomerSessionComponentsPaymentElementFeaturesPaymentMethodRemove string

// List of values that CustomerSessionComponentsPaymentElementFeaturesPaymentMethodRemove can take
const (
	CustomerSessionComponentsPaymentElementFeaturesPaymentMethodRemoveAuto  CustomerSessionComponentsPaymentElementFeaturesPaymentMethodRemove = "auto"
	CustomerSessionComponentsPaymentElementFeaturesPaymentMethodRemoveNever CustomerSessionComponentsPaymentElementFeaturesPaymentMethodRemove = "never"
)

// Controls whether the Payment Element offers to save a new payment method.
type CustomerSessionComponentsPaymentElementFeaturesPaymentMethodSave string

// List of values that CustomerSessionComponentsPaymentElementFeaturesPaymentMethodSave can take
const (
	CustomerSessionComponentsPaymentElementFeaturesPaymentMethodSaveAuto  CustomerSessionComponentsPaymentElementFeaturesPaymentMethodSave = "auto"
	CustomerSessionComponentsPaymentElementFeaturesPaymentMethodSaveNever CustomerSessionComponentsPaymentElementFeaturesPaymentMethodSave = "never"
)

// Controls whether the Payment Element offers to set a payment method as the default.
type CustomerSessionComponentsPaymentElementFeaturesPaymentMethodSetAsDefault string

// List of values that CustomerSessionComponentsPaymentElementFeaturesPaymentMethodSetAsDefault can take
const (
	CustomerSessionComponentsPaymentElementFeaturesPaymentMethodSetAsDefaultAuto  CustomerSessionComponentsPaymentElementFeaturesPaymentMethodSetAsDefault = "auto"
	CustomerSessionComponentsPaymentElementFeaturesPaymentMethodSetAsDefaultNever CustomerSessionComponentsPaymentElementFeaturesPaymentMethodSetAsDefault = "never"
)

// Controls whether the Payment Element allows the updating of a saved payment method.
type CustomerSessionComponentsPaymentElementFeaturesPaymentMethodUpdate string

// List of values that CustomerSessionComponentsPaymentElementFeaturesPaymentMethodUpdate can take
const (
	CustomerSessionComponentsPaymentElementFeaturesPaymentMethodUpdateAuto  CustomerSessionComponentsPaymentElementFeaturesPaymentMethodUpdate = "auto"
	CustomerSessionComponentsPaymentElementFeaturesPaymentMethodUpdateNever CustomerSessionComponentsPaymentElementFeaturesPaymentMethodUpdate = "never"
)

// This hash defines whether the payment element supports certain features.
type CustomerSessionComponentsPaymentElementFeaturesParams struct {
	// Controls whether the Payment Element allows the removal of a saved payment method.
	PaymentMethodRemove *string `form:"payment_method_remove"`
	// Controls whether the Payment Element offers to save a new payment method.
	PaymentMethodSave *string `form:"payment_method_save"`
	// Controls whether the Payment Element offers to set a payment method as the default.
	PaymentMethodSetAsDefault *string `form:"payment_method_set_as_default"`
	// Controls whether the Payment Element allows the updating of a saved payment method.
	PaymentMethodUpdate *string `form:"payment_method_update"`
}

// Configuration for the payment element.
type CustomerSessionComponentsPaymentElementParams struct {
	// Whether the payment element is enabled.
	Enabled *bool `form:"enabled"`
	// This hash defines whether the payment element supports certain features.
	Features *CustomerSessionComponentsPaymentElementFeaturesParams `form:"features"`
}

// Configuration for the pricing table.
type CustomerSessionComponentsPricingTableParams struct {
	// Whether the pricing table is enabled.
	Enabled *bool `form:"enabled"`
}

// Configuration for each component. 1 component must be enabled.
type CustomerSessionComponentsParams struct {
	// Configuration for the payment element.
	PaymentElement *CustomerSessionComponentsPaymentElementParams `form:"payment_element"`
	// Configuration for the pricing table.
	PricingTable *CustomerSessionComponentsPricingTableParams `form:"pricing_table"`
}

// Creates a customer session object that includes a single-use client secret that you can use on your front-end to grant client-side API access for certain customer resources.
type CustomerSessionParams struct {
	Params `form:"*"`
	// Configuration for each component. 1 component must be enabled.
	Components *CustomerSessionComponentsParams `form:"components"`
	// The ID of an existing customer for which to create the customer session.
	Customer *string `form:"customer"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *CustomerSessionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// This hash defines whether the payment element supports certain features.
type CustomerSessionComponentsPaymentElementFeatures struct {
	// Controls whether the Payment Element allows the removal of a saved payment method.
	PaymentMethodRemove CustomerSessionComponentsPaymentElementFeaturesPaymentMethodRemove `json:"payment_method_remove"`
	// Controls whether the Payment Element offers to save a new payment method.
	PaymentMethodSave CustomerSessionComponentsPaymentElementFeaturesPaymentMethodSave `json:"payment_method_save"`
	// Controls whether the Payment Element offers to set a payment method as the default.
	PaymentMethodSetAsDefault CustomerSessionComponentsPaymentElementFeaturesPaymentMethodSetAsDefault `json:"payment_method_set_as_default"`
	// Controls whether the Payment Element allows the updating of a saved payment method.
	PaymentMethodUpdate CustomerSessionComponentsPaymentElementFeaturesPaymentMethodUpdate `json:"payment_method_update"`
}

// This hash contains whether the payment element is enabled and the features it supports.
type CustomerSessionComponentsPaymentElement struct {
	// Whether the payment element is enabled.
	Enabled bool `json:"enabled"`
	// This hash defines whether the payment element supports certain features.
	Features *CustomerSessionComponentsPaymentElementFeatures `json:"features"`
}

// This hash contains whether the pricing table is enabled.
type CustomerSessionComponentsPricingTable struct {
	// Whether the pricing table is enabled.
	Enabled bool `json:"enabled"`
}

// Configuration for the components supported by this customer session.
type CustomerSessionComponents struct {
	// This hash contains whether the payment element is enabled and the features it supports.
	PaymentElement *CustomerSessionComponentsPaymentElement `json:"payment_element"`
	// This hash contains whether the pricing table is enabled.
	PricingTable *CustomerSessionComponentsPricingTable `json:"pricing_table"`
}

// A customer session allows you to grant client access to Stripe's frontend SDKs (like StripeJs)
// control over a customer.
type CustomerSession struct {
	APIResource
	// The client secret of this customer session. Used on the client to set up secure access to the given `customer`.
	//
	// The client secret can be used to provide access to `customer` from your frontend. It should not be stored, logged, or exposed to anyone other than the relevant customer. Make sure that you have TLS enabled on any page that includes the client secret.
	ClientSecret string `json:"client_secret"`
	// Configuration for the components supported by this customer session.
	Components *CustomerSessionComponents `json:"components"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The customer the customer session was created for.
	Customer *Customer `json:"customer"`
	// The timestamp at which this customer session will expire.
	ExpiresAt int64 `json:"expires_at"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
