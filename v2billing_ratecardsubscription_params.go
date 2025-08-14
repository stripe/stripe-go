//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Optionally filter by the payer associated with Cadences which the RateCardSubscriptions are subscribed to.
// Mutually exclusive with `billing_cadence`, `rate_card`, and `rate_card_version`.
type V2BillingRateCardSubscriptionListPayerParams struct {
	// The ID of the Customer object. If provided, only RateCardSubscriptions that are subscribed on the Cadences with the specified Payer will be returned.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// A string identifying the type of the payer. Currently the only supported value is `customer`.
	Type *string `form:"type" json:"type"`
}

// List all RateCardSubscription objects.
type V2BillingRateCardSubscriptionListParams struct {
	Params `form:"*"`
	// Optionally filter by a BillingCadence. Mutually exclusive with `payers`, `rate_card`, and `rate_card_version`.
	BillingCadence *string `form:"billing_cadence" json:"billing_cadence,omitempty"`
	// The page size limit, if not provided the default is 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Optionally filter by the payer associated with Cadences which the RateCardSubscriptions are subscribed to.
	// Mutually exclusive with `billing_cadence`, `rate_card`, and `rate_card_version`.
	Payer *V2BillingRateCardSubscriptionListPayerParams `form:"payer" json:"payer,omitempty"`
	// Optionally filter by a RateCard. Mutually exclusive with `billing_cadence`, `payers`, and `rate_card_version`.
	RateCard *string `form:"rate_card" json:"rate_card,omitempty"`
	// Optionally filter by a RateCard version. Mutually exclusive with `billing_cadence`, `payers`, and `rate_card`.
	RateCardVersion *string `form:"rate_card_version" json:"rate_card_version,omitempty"`
	// Optionally filter by servicing status.
	ServicingStatus *string `form:"servicing_status" json:"servicing_status,omitempty"`
}

// Create a RateCardSubscription to bill a RateCard on a specified billing Cadence.
type V2BillingRateCardSubscriptionParams struct {
	Params `form:"*"`
	// The ID of the billing Cadence.
	BillingCadence *string `form:"billing_cadence" json:"billing_cadence,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The ID of the RateCard.
	RateCard *string `form:"rate_card" json:"rate_card,omitempty"`
	// The ID of the RateCardVersion. If not specified, defaults to the "live_version" of the RateCard at the time of creation.
	RateCardVersion *string `form:"rate_card_version" json:"rate_card_version,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingRateCardSubscriptionParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Cancel an existing, active RateCardSubscription.
type V2BillingRateCardSubscriptionCancelParams struct {
	Params `form:"*"`
}

// Create a RateCardSubscription to bill a RateCard on a specified billing Cadence.
type V2BillingRateCardSubscriptionCreateParams struct {
	Params `form:"*"`
	// The ID of the billing Cadence.
	BillingCadence *string `form:"billing_cadence" json:"billing_cadence"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The ID of the RateCard.
	RateCard *string `form:"rate_card" json:"rate_card"`
	// The ID of the RateCardVersion. If not specified, defaults to the "live_version" of the RateCard at the time of creation.
	RateCardVersion *string `form:"rate_card_version" json:"rate_card_version,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingRateCardSubscriptionCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieve a RateCardSubscription by ID.
type V2BillingRateCardSubscriptionRetrieveParams struct {
	Params `form:"*"`
}

// Update fields on an existing, active RateCardSubscription.
type V2BillingRateCardSubscriptionUpdateParams struct {
	Params `form:"*"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingRateCardSubscriptionUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}
