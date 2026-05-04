//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of the attribution source.
type DelegatedCheckoutRequestedSessionAffiliateAttributionSourceType string

// List of values that DelegatedCheckoutRequestedSessionAffiliateAttributionSourceType can take
const (
	DelegatedCheckoutRequestedSessionAffiliateAttributionSourceTypePlatform DelegatedCheckoutRequestedSessionAffiliateAttributionSourceType = "platform"
	DelegatedCheckoutRequestedSessionAffiliateAttributionSourceTypeURL      DelegatedCheckoutRequestedSessionAffiliateAttributionSourceType = "url"
)

// Whether this is the first or last touchpoint.
type DelegatedCheckoutRequestedSessionAffiliateAttributionTouchpoint string

// List of values that DelegatedCheckoutRequestedSessionAffiliateAttributionTouchpoint can take
const (
	DelegatedCheckoutRequestedSessionAffiliateAttributionTouchpointFirst DelegatedCheckoutRequestedSessionAffiliateAttributionTouchpoint = "first"
	DelegatedCheckoutRequestedSessionAffiliateAttributionTouchpointLast  DelegatedCheckoutRequestedSessionAffiliateAttributionTouchpoint = "last"
)

// The marketing consent channel.
type DelegatedCheckoutRequestedSessionBuyerConsentsMarketingConsentChannel string

// List of values that DelegatedCheckoutRequestedSessionBuyerConsentsMarketingConsentChannel can take
const (
	DelegatedCheckoutRequestedSessionBuyerConsentsMarketingConsentChannelEmail DelegatedCheckoutRequestedSessionBuyerConsentsMarketingConsentChannel = "email"
	DelegatedCheckoutRequestedSessionBuyerConsentsMarketingConsentChannelSms   DelegatedCheckoutRequestedSessionBuyerConsentsMarketingConsentChannel = "sms"
)

// The consent status. 'granted' means the buyer opted in, 'none' means they did not.
type DelegatedCheckoutRequestedSessionBuyerConsentsMarketingConsentStatus string

// List of values that DelegatedCheckoutRequestedSessionBuyerConsentsMarketingConsentStatus can take
const (
	DelegatedCheckoutRequestedSessionBuyerConsentsMarketingConsentStatusGranted DelegatedCheckoutRequestedSessionBuyerConsentsMarketingConsentStatus = "granted"
	DelegatedCheckoutRequestedSessionBuyerConsentsMarketingConsentStatusNone    DelegatedCheckoutRequestedSessionBuyerConsentsMarketingConsentStatus = "none"
)

// The marketing channel type.
type DelegatedCheckoutRequestedSessionBuyerConsentsMarketingOptionChannel string

// List of values that DelegatedCheckoutRequestedSessionBuyerConsentsMarketingOptionChannel can take
const (
	DelegatedCheckoutRequestedSessionBuyerConsentsMarketingOptionChannelEmail DelegatedCheckoutRequestedSessionBuyerConsentsMarketingOptionChannel = "email"
	DelegatedCheckoutRequestedSessionBuyerConsentsMarketingOptionChannelSms   DelegatedCheckoutRequestedSessionBuyerConsentsMarketingOptionChannel = "sms"
)

// The type of the fulfillment option.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptionType string

// List of values that DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptionType can take
const (
	DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptionTypeDigital  DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptionType = "digital"
	DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptionTypeShipping DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptionType = "shipping"
)

// The type of the selected fulfillment option.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionType string

// List of values that DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionType can take
const (
	DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionTypeDigital  DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionType = "digital"
	DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionTypeShipping DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionType = "shipping"
)

// The type of the selected fulfillment option.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionOverrideType string

// List of values that DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionOverrideType can take
const (
	DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionOverrideTypeDigital  DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionOverrideType = "digital"
	DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionOverrideTypeShipping DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionOverrideType = "shipping"
)

// The content type of the disclosure.
type DelegatedCheckoutRequestedSessionLineItemDetailProductDetailsDisclosureContentType string

// List of values that DelegatedCheckoutRequestedSessionLineItemDetailProductDetailsDisclosureContentType can take
const (
	DelegatedCheckoutRequestedSessionLineItemDetailProductDetailsDisclosureContentTypeLink     DelegatedCheckoutRequestedSessionLineItemDetailProductDetailsDisclosureContentType = "link"
	DelegatedCheckoutRequestedSessionLineItemDetailProductDetailsDisclosureContentTypeMarkdown DelegatedCheckoutRequestedSessionLineItemDetailProductDetailsDisclosureContentType = "markdown"
	DelegatedCheckoutRequestedSessionLineItemDetailProductDetailsDisclosureContentTypePlain    DelegatedCheckoutRequestedSessionLineItemDetailProductDetailsDisclosureContentType = "plain"
)

// The type of disclosure.
type DelegatedCheckoutRequestedSessionLineItemDetailProductDetailsDisclosureType string

// List of values that DelegatedCheckoutRequestedSessionLineItemDetailProductDetailsDisclosureType can take
const (
	DelegatedCheckoutRequestedSessionLineItemDetailProductDetailsDisclosureTypeDisclaimer DelegatedCheckoutRequestedSessionLineItemDetailProductDetailsDisclosureType = "disclaimer"
)

// The card brands blocked by the agent.
type DelegatedCheckoutRequestedSessionPaymentMethodOptionsCardBrandsBlocked string

// List of values that DelegatedCheckoutRequestedSessionPaymentMethodOptionsCardBrandsBlocked can take
const (
	DelegatedCheckoutRequestedSessionPaymentMethodOptionsCardBrandsBlockedAmericanExpress DelegatedCheckoutRequestedSessionPaymentMethodOptionsCardBrandsBlocked = "american_express"
	DelegatedCheckoutRequestedSessionPaymentMethodOptionsCardBrandsBlockedMastercard      DelegatedCheckoutRequestedSessionPaymentMethodOptionsCardBrandsBlocked = "mastercard"
	DelegatedCheckoutRequestedSessionPaymentMethodOptionsCardBrandsBlockedVisa            DelegatedCheckoutRequestedSessionPaymentMethodOptionsCardBrandsBlocked = "visa"
)

// The computed displayable card brands.
type DelegatedCheckoutRequestedSessionPaymentMethodOptionsDisplayableCardBrand string

// List of values that DelegatedCheckoutRequestedSessionPaymentMethodOptionsDisplayableCardBrand can take
const (
	DelegatedCheckoutRequestedSessionPaymentMethodOptionsDisplayableCardBrandAmericanExpress DelegatedCheckoutRequestedSessionPaymentMethodOptionsDisplayableCardBrand = "american_express"
	DelegatedCheckoutRequestedSessionPaymentMethodOptionsDisplayableCardBrandMastercard      DelegatedCheckoutRequestedSessionPaymentMethodOptionsDisplayableCardBrand = "mastercard"
	DelegatedCheckoutRequestedSessionPaymentMethodOptionsDisplayableCardBrandVisa            DelegatedCheckoutRequestedSessionPaymentMethodOptionsDisplayableCardBrand = "visa"
)

// The computed displayable payment method types.
type DelegatedCheckoutRequestedSessionPaymentMethodOptionsDisplayablePaymentMethodType string

// List of values that DelegatedCheckoutRequestedSessionPaymentMethodOptionsDisplayablePaymentMethodType can take
const (
	DelegatedCheckoutRequestedSessionPaymentMethodOptionsDisplayablePaymentMethodTypeAffirm DelegatedCheckoutRequestedSessionPaymentMethodOptionsDisplayablePaymentMethodType = "affirm"
	DelegatedCheckoutRequestedSessionPaymentMethodOptionsDisplayablePaymentMethodTypeCard   DelegatedCheckoutRequestedSessionPaymentMethodOptionsDisplayablePaymentMethodType = "card"
	DelegatedCheckoutRequestedSessionPaymentMethodOptionsDisplayablePaymentMethodTypeKlarna DelegatedCheckoutRequestedSessionPaymentMethodOptionsDisplayablePaymentMethodType = "klarna"
)

// The payment method types excluded by the agent.
type DelegatedCheckoutRequestedSessionPaymentMethodOptionsExcludedPaymentMethodType string

// List of values that DelegatedCheckoutRequestedSessionPaymentMethodOptionsExcludedPaymentMethodType can take
const (
	DelegatedCheckoutRequestedSessionPaymentMethodOptionsExcludedPaymentMethodTypeAffirm DelegatedCheckoutRequestedSessionPaymentMethodOptionsExcludedPaymentMethodType = "affirm"
	DelegatedCheckoutRequestedSessionPaymentMethodOptionsExcludedPaymentMethodTypeCard   DelegatedCheckoutRequestedSessionPaymentMethodOptionsExcludedPaymentMethodType = "card"
	DelegatedCheckoutRequestedSessionPaymentMethodOptionsExcludedPaymentMethodTypeKlarna DelegatedCheckoutRequestedSessionPaymentMethodOptionsExcludedPaymentMethodType = "klarna"
)

// The card brands supported by the seller.
type DelegatedCheckoutRequestedSessionSellerDetailsCardBrand string

// List of values that DelegatedCheckoutRequestedSessionSellerDetailsCardBrand can take
const (
	DelegatedCheckoutRequestedSessionSellerDetailsCardBrandAmericanExpress DelegatedCheckoutRequestedSessionSellerDetailsCardBrand = "american_express"
	DelegatedCheckoutRequestedSessionSellerDetailsCardBrandMastercard      DelegatedCheckoutRequestedSessionSellerDetailsCardBrand = "mastercard"
	DelegatedCheckoutRequestedSessionSellerDetailsCardBrandVisa            DelegatedCheckoutRequestedSessionSellerDetailsCardBrand = "visa"
)

// The payment method types supported by the seller.
type DelegatedCheckoutRequestedSessionSellerDetailsPaymentMethodType string

// List of values that DelegatedCheckoutRequestedSessionSellerDetailsPaymentMethodType can take
const (
	DelegatedCheckoutRequestedSessionSellerDetailsPaymentMethodTypeAffirm DelegatedCheckoutRequestedSessionSellerDetailsPaymentMethodType = "affirm"
	DelegatedCheckoutRequestedSessionSellerDetailsPaymentMethodTypeCard   DelegatedCheckoutRequestedSessionSellerDetailsPaymentMethodType = "card"
	DelegatedCheckoutRequestedSessionSellerDetailsPaymentMethodTypeKlarna DelegatedCheckoutRequestedSessionSellerDetailsPaymentMethodType = "klarna"
)

// Whether or not the payment method should be saved for future use.
type DelegatedCheckoutRequestedSessionSetupFutureUsage string

// List of values that DelegatedCheckoutRequestedSessionSetupFutureUsage can take
const (
	DelegatedCheckoutRequestedSessionSetupFutureUsageOnSession DelegatedCheckoutRequestedSessionSetupFutureUsage = "on_session"
)

// The status of the requested session.
type DelegatedCheckoutRequestedSessionStatus string

// List of values that DelegatedCheckoutRequestedSessionStatus can take
const (
	DelegatedCheckoutRequestedSessionStatusCompleted      DelegatedCheckoutRequestedSessionStatus = "completed"
	DelegatedCheckoutRequestedSessionStatusExpired        DelegatedCheckoutRequestedSessionStatus = "expired"
	DelegatedCheckoutRequestedSessionStatusOpen           DelegatedCheckoutRequestedSessionStatus = "open"
	DelegatedCheckoutRequestedSessionStatusRequiresAction DelegatedCheckoutRequestedSessionStatus = "requires_action"
)

// Retrieves a requested session
type DelegatedCheckoutRequestedSessionParams struct {
	Params `form:"*"`
	// Affiliate attribution data associated with this requested session.
	AffiliateAttribution *DelegatedCheckoutRequestedSessionAffiliateAttributionParams `form:"affiliate_attribution" json:"affiliate_attribution,omitempty"`
	// The currency for this requested session.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// The customer for this requested session.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The details of the fulfillment.
	FulfillmentDetails *DelegatedCheckoutRequestedSessionFulfillmentDetailsParams `form:"fulfillment_details" json:"fulfillment_details,omitempty"`
	// The details of the line items.
	LineItemDetails []*DelegatedCheckoutRequestedSessionLineItemDetailParams `form:"line_item_details" json:"line_item_details,omitempty"`
	// The metadata for this requested session.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The payment method for this requested session.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// The payment method options for this requested session.
	PaymentMethodOptions *DelegatedCheckoutRequestedSessionPaymentMethodOptionsParams `form:"payment_method_options" json:"payment_method_options,omitempty"`
	// The details of the seller.
	SellerDetails *DelegatedCheckoutRequestedSessionSellerDetailsParams `form:"seller_details" json:"seller_details,omitempty"`
	// The setup future usage for this requested session.
	SetupFutureUsage *string `form:"setup_future_usage" json:"setup_future_usage,omitempty"`
	// The shared metadata for this requested session.
	SharedMetadata map[string]string                                   `form:"shared_metadata" json:"shared_metadata,omitempty"`
	UnsetFields    []DelegatedCheckoutRequestedSessionParamsUnsetField `form:"-" json:"-"`
}

// DelegatedCheckoutRequestedSessionParamsUnsetField is the list of fields that can be cleared/unset on DelegatedCheckoutRequestedSessionParams.
type DelegatedCheckoutRequestedSessionParamsUnsetField string

const (
	DelegatedCheckoutRequestedSessionParamsUnsetFieldMetadata       DelegatedCheckoutRequestedSessionParamsUnsetField = "metadata"
	DelegatedCheckoutRequestedSessionParamsUnsetFieldSharedMetadata DelegatedCheckoutRequestedSessionParamsUnsetField = "shared_metadata"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *DelegatedCheckoutRequestedSessionParams) AddUnsetField(field DelegatedCheckoutRequestedSessionParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *DelegatedCheckoutRequestedSessionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *DelegatedCheckoutRequestedSessionParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The digital fulfillment option.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionDigitalParams struct {
	// The digital option identifier.
	DigitalOption *string `form:"digital_option" json:"digital_option"`
}

// The shipping fulfillment option.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionShippingParams struct {
	// The shipping option identifier.
	ShippingOption *string `form:"shipping_option" json:"shipping_option"`
}

// The fulfillment option to select.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionParams struct {
	// The digital fulfillment option.
	Digital *DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionDigitalParams `form:"digital" json:"digital,omitempty"`
	// The shipping fulfillment option.
	Shipping *DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionShippingParams `form:"shipping" json:"shipping,omitempty"`
	// The type of fulfillment option.
	Type *string `form:"type" json:"type"`
}

// The digital fulfillment option.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionOverrideDigitalParams struct {
	// The digital option identifier.
	DigitalOption *string `form:"digital_option" json:"digital_option"`
}

// The shipping fulfillment option.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionOverrideShippingParams struct {
	// The shipping option identifier.
	ShippingOption *string `form:"shipping_option" json:"shipping_option"`
}

// The fulfillment option overrides for specific line items.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionOverrideParams struct {
	// The digital fulfillment option.
	Digital *DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionOverrideDigitalParams `form:"digital" json:"digital,omitempty"`
	// The line item keys that this fulfillment option override applies to.
	LineItemKeys []*string `form:"line_item_keys" json:"line_item_keys"`
	// The shipping fulfillment option.
	Shipping *DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionOverrideShippingParams `form:"shipping" json:"shipping,omitempty"`
	// The type of fulfillment option.
	Type *string `form:"type" json:"type"`
}

// The details of the fulfillment.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsParams struct {
	// The customer's address.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// The customer's email address.
	Email *string `form:"email" json:"email,omitempty"`
	// The customer's name.
	Name *string `form:"name" json:"name,omitempty"`
	// The customer's phone number.
	Phone *string `form:"phone" json:"phone,omitempty"`
	// The fulfillment option to select.
	SelectedFulfillmentOption *DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionParams `form:"selected_fulfillment_option" json:"selected_fulfillment_option,omitempty"`
	// The fulfillment option overrides for specific line items.
	SelectedFulfillmentOptionOverrides []*DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionOverrideParams `form:"selected_fulfillment_option_overrides" json:"selected_fulfillment_option_overrides,omitempty"`
}

// The details of the line items.
type DelegatedCheckoutRequestedSessionLineItemDetailParams struct {
	// The key of the line item.
	Key *string `form:"key" json:"key,omitempty"`
	// The quantity of the line item.
	Quantity *int64 `form:"quantity" json:"quantity"`
	// The SKU ID of the line item.
	SKUID *string `form:"sku_id" json:"sku_id,omitempty"`
}

// Card-specific payment method options.
type DelegatedCheckoutRequestedSessionPaymentMethodOptionsCardParams struct {
	// The card brands to exclude from the session.
	BrandsBlocked []*string `form:"brands_blocked" json:"brands_blocked,omitempty"`
}

// The payment method options for this requested session.
type DelegatedCheckoutRequestedSessionPaymentMethodOptionsParams struct {
	// Card-specific payment method options.
	Card *DelegatedCheckoutRequestedSessionPaymentMethodOptionsCardParams `form:"card" json:"card,omitempty"`
	// The payment method types to exclude from the session.
	ExcludedPaymentMethodTypes []*string `form:"excluded_payment_method_types" json:"excluded_payment_method_types,omitempty"`
}

// Context about where the attribution originated.
type DelegatedCheckoutRequestedSessionAffiliateAttributionSourceParams struct {
	// The platform where the attribution originated.
	Platform *string `form:"platform" json:"platform,omitempty"`
	// The type of the attribution source.
	Type *string `form:"type" json:"type"`
	// The URL where the attribution originated.
	URL *string `form:"url" json:"url,omitempty"`
}

// Affiliate attribution data associated with this requested session.
type DelegatedCheckoutRequestedSessionAffiliateAttributionParams struct {
	// Agent-scoped campaign identifier.
	CampaignID *string `form:"campaign_id" json:"campaign_id,omitempty"`
	// Agent-scoped creative identifier.
	CreativeID *string `form:"creative_id" json:"creative_id,omitempty"`
	// Timestamp when the attribution token expires.
	ExpiresAt *int64 `form:"expires_at" json:"expires_at"`
	// Agent-issued secret to validate the legitimacy of the source of this data.
	IdentificationToken *string `form:"identification_token" json:"identification_token"`
	// Timestamp for when the attribution token was issued.
	IssuedAt *int64 `form:"issued_at" json:"issued_at"`
	// Identifier for the attribution agent / affiliate network namespace.
	Provider *string `form:"provider" json:"provider"`
	// Agent-scoped affiliate/publisher identifier.
	PublisherID *string `form:"publisher_id" json:"publisher_id,omitempty"`
	// Freeform key/value pairs for additional non-sensitive per-agent data.
	SharedMetadata map[string]string `form:"shared_metadata" json:"shared_metadata,omitempty"`
	// Context about where the attribution originated.
	Source *DelegatedCheckoutRequestedSessionAffiliateAttributionSourceParams `form:"source" json:"source,omitempty"`
	// Agent-scoped sub-tracking identifier.
	SubID *string `form:"sub_id" json:"sub_id,omitempty"`
	// Whether this is the first or last touchpoint.
	Touchpoint *string `form:"touchpoint" json:"touchpoint"`
}

// The details of the seller.
type DelegatedCheckoutRequestedSessionSellerDetailsParams struct {
	// The network profile for the seller.
	NetworkProfile *string `form:"network_profile" json:"network_profile"`
}

// Context about where the attribution originated.
type DelegatedCheckoutRequestedSessionConfirmAffiliateAttributionSourceParams struct {
	// The platform where the attribution originated.
	Platform *string `form:"platform" json:"platform,omitempty"`
	// The type of the attribution source.
	Type *string `form:"type" json:"type"`
	// The URL where the attribution originated.
	URL *string `form:"url" json:"url,omitempty"`
}

// Affiliate attribution data associated with this requested session.
type DelegatedCheckoutRequestedSessionConfirmAffiliateAttributionParams struct {
	// Agent-scoped campaign identifier.
	CampaignID *string `form:"campaign_id" json:"campaign_id,omitempty"`
	// Agent-scoped creative identifier.
	CreativeID *string `form:"creative_id" json:"creative_id,omitempty"`
	// Timestamp when the attribution token expires.
	ExpiresAt *int64 `form:"expires_at" json:"expires_at"`
	// Agent-issued secret to validate the legitimacy of the source of this data.
	IdentificationToken *string `form:"identification_token" json:"identification_token"`
	// Timestamp for when the attribution token was issued.
	IssuedAt *int64 `form:"issued_at" json:"issued_at"`
	// Identifier for the attribution agent / affiliate network namespace.
	Provider *string `form:"provider" json:"provider"`
	// Agent-scoped affiliate/publisher identifier.
	PublisherID *string `form:"publisher_id" json:"publisher_id,omitempty"`
	// Freeform key/value pairs for additional non-sensitive per-agent data.
	SharedMetadata map[string]string `form:"shared_metadata" json:"shared_metadata,omitempty"`
	// Context about where the attribution originated.
	Source *DelegatedCheckoutRequestedSessionConfirmAffiliateAttributionSourceParams `form:"source" json:"source,omitempty"`
	// Agent-scoped sub-tracking identifier.
	SubID *string `form:"sub_id" json:"sub_id,omitempty"`
	// Whether this is the first or last touchpoint.
	Touchpoint *string `form:"touchpoint" json:"touchpoint"`
}

// The list of marketing consent entries.
type DelegatedCheckoutRequestedSessionConfirmBuyerConsentsMarketingConsentParams struct {
	// The marketing consent channel.
	Channel *string `form:"channel" json:"channel"`
	// The consent status. Use 'granted' to indicate the buyer has opted in.
	Status *string `form:"status" json:"status"`
}

// The marketing consent data for the buyer.
type DelegatedCheckoutRequestedSessionConfirmBuyerConsentsMarketingParams struct {
	// The list of marketing consent entries.
	Consents []*DelegatedCheckoutRequestedSessionConfirmBuyerConsentsMarketingConsentParams `form:"consents" json:"consents,omitempty"`
}

// The buyer's consent choices for marketing communications.
type DelegatedCheckoutRequestedSessionConfirmBuyerConsentsParams struct {
	// The marketing consent data for the buyer.
	Marketing *DelegatedCheckoutRequestedSessionConfirmBuyerConsentsMarketingParams `form:"marketing" json:"marketing,omitempty"`
}

// The client device metadata details for this requested session.
type DelegatedCheckoutRequestedSessionConfirmRiskDetailsClientDeviceMetadataDetailsParams struct {
	// The radar session.
	RadarSession *string `form:"radar_session" json:"radar_session,omitempty"`
	// The referrer of the client device.
	Referrer *string `form:"referrer" json:"referrer,omitempty"`
	// The remote IP address of the client device.
	RemoteIP *string `form:"remote_ip" json:"remote_ip,omitempty"`
	// The time on page in milliseconds.
	TimeOnPageMS *int64 `form:"time_on_page_ms" json:"time_on_page_ms,omitempty"`
	// The user agent of the client device.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Risk details/signals associated with the requested session
type DelegatedCheckoutRequestedSessionConfirmRiskDetailsParams struct {
	// The client device metadata details for this requested session.
	ClientDeviceMetadataDetails *DelegatedCheckoutRequestedSessionConfirmRiskDetailsClientDeviceMetadataDetailsParams `form:"client_device_metadata_details" json:"client_device_metadata_details,omitempty"`
}

// Confirms a requested session
type DelegatedCheckoutRequestedSessionConfirmParams struct {
	Params `form:"*"`
	// Affiliate attribution data associated with this requested session.
	AffiliateAttribution *DelegatedCheckoutRequestedSessionConfirmAffiliateAttributionParams `form:"affiliate_attribution" json:"affiliate_attribution,omitempty"`
	// The buyer's consent choices for marketing communications.
	BuyerConsents *DelegatedCheckoutRequestedSessionConfirmBuyerConsentsParams `form:"buyer_consents" json:"buyer_consents,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The metadata for this requested session.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The PaymentMethod to use with the requested session.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// The URL to redirect your customer back to after they authenticate or complete a payment action. Required for redirect-based payment methods such as Affirm or Klarna.
	ReturnURL *string `form:"return_url" json:"return_url,omitempty"`
	// Risk details/signals associated with the requested session
	RiskDetails *DelegatedCheckoutRequestedSessionConfirmRiskDetailsParams `form:"risk_details" json:"risk_details,omitempty"`
	UnsetFields []DelegatedCheckoutRequestedSessionConfirmParamsUnsetField `form:"-" json:"-"`
}

// DelegatedCheckoutRequestedSessionConfirmParamsUnsetField is the list of fields that can be cleared/unset on DelegatedCheckoutRequestedSessionConfirmParams.
type DelegatedCheckoutRequestedSessionConfirmParamsUnsetField string

const (
	DelegatedCheckoutRequestedSessionConfirmParamsUnsetFieldMetadata DelegatedCheckoutRequestedSessionConfirmParamsUnsetField = "metadata"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *DelegatedCheckoutRequestedSessionConfirmParams) AddUnsetField(field DelegatedCheckoutRequestedSessionConfirmParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *DelegatedCheckoutRequestedSessionConfirmParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *DelegatedCheckoutRequestedSessionConfirmParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Expires a requested session
type DelegatedCheckoutRequestedSessionExpireParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *DelegatedCheckoutRequestedSessionExpireParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves a requested session
type DelegatedCheckoutRequestedSessionRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *DelegatedCheckoutRequestedSessionRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The digital fulfillment option.
type DelegatedCheckoutRequestedSessionUpdateFulfillmentDetailsSelectedFulfillmentOptionDigitalParams struct {
	// The digital option identifier.
	DigitalOption *string `form:"digital_option" json:"digital_option"`
}

// The shipping fulfillment option.
type DelegatedCheckoutRequestedSessionUpdateFulfillmentDetailsSelectedFulfillmentOptionShippingParams struct {
	// The shipping option identifier.
	ShippingOption *string `form:"shipping_option" json:"shipping_option"`
}

// The fulfillment option to select.
type DelegatedCheckoutRequestedSessionUpdateFulfillmentDetailsSelectedFulfillmentOptionParams struct {
	// The digital fulfillment option.
	Digital *DelegatedCheckoutRequestedSessionUpdateFulfillmentDetailsSelectedFulfillmentOptionDigitalParams `form:"digital" json:"digital,omitempty"`
	// The shipping fulfillment option.
	Shipping *DelegatedCheckoutRequestedSessionUpdateFulfillmentDetailsSelectedFulfillmentOptionShippingParams `form:"shipping" json:"shipping,omitempty"`
	// The type of fulfillment option.
	Type *string `form:"type" json:"type"`
}

// The digital fulfillment option.
type DelegatedCheckoutRequestedSessionUpdateFulfillmentDetailsSelectedFulfillmentOptionOverrideDigitalParams struct {
	// The digital option identifier.
	DigitalOption *string `form:"digital_option" json:"digital_option"`
}

// The shipping fulfillment option.
type DelegatedCheckoutRequestedSessionUpdateFulfillmentDetailsSelectedFulfillmentOptionOverrideShippingParams struct {
	// The shipping option identifier.
	ShippingOption *string `form:"shipping_option" json:"shipping_option"`
}

// The fulfillment option overrides for specific line items.
type DelegatedCheckoutRequestedSessionUpdateFulfillmentDetailsSelectedFulfillmentOptionOverrideParams struct {
	// The digital fulfillment option.
	Digital *DelegatedCheckoutRequestedSessionUpdateFulfillmentDetailsSelectedFulfillmentOptionOverrideDigitalParams `form:"digital" json:"digital,omitempty"`
	// The line item keys that this fulfillment option override applies to.
	LineItemKeys []*string `form:"line_item_keys" json:"line_item_keys"`
	// The shipping fulfillment option.
	Shipping *DelegatedCheckoutRequestedSessionUpdateFulfillmentDetailsSelectedFulfillmentOptionOverrideShippingParams `form:"shipping" json:"shipping,omitempty"`
	// The type of fulfillment option.
	Type *string `form:"type" json:"type"`
}

// The details of the fulfillment.
type DelegatedCheckoutRequestedSessionUpdateFulfillmentDetailsParams struct {
	// The customer's address.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// The customer's email address.
	Email *string `form:"email" json:"email,omitempty"`
	// The customer's name.
	Name *string `form:"name" json:"name,omitempty"`
	// The customer's phone number.
	Phone *string `form:"phone" json:"phone,omitempty"`
	// The fulfillment option to select.
	SelectedFulfillmentOption *DelegatedCheckoutRequestedSessionUpdateFulfillmentDetailsSelectedFulfillmentOptionParams `form:"selected_fulfillment_option" json:"selected_fulfillment_option,omitempty"`
	// The fulfillment option overrides for specific line items.
	SelectedFulfillmentOptionOverrides []*DelegatedCheckoutRequestedSessionUpdateFulfillmentDetailsSelectedFulfillmentOptionOverrideParams `form:"selected_fulfillment_option_overrides" json:"selected_fulfillment_option_overrides,omitempty"`
}

// The details of the line items.
type DelegatedCheckoutRequestedSessionUpdateLineItemDetailParams struct {
	// The key of the line item.
	Key *string `form:"key" json:"key"`
	// The quantity of the line item.
	Quantity *int64 `form:"quantity" json:"quantity"`
}

// Card-specific payment method options.
type DelegatedCheckoutRequestedSessionUpdatePaymentMethodOptionsCardParams struct {
	// The card brands to exclude from the session.
	BrandsBlocked []*string `form:"brands_blocked" json:"brands_blocked,omitempty"`
}

// The payment method options for this requested session.
type DelegatedCheckoutRequestedSessionUpdatePaymentMethodOptionsParams struct {
	// Card-specific payment method options.
	Card *DelegatedCheckoutRequestedSessionUpdatePaymentMethodOptionsCardParams `form:"card" json:"card,omitempty"`
	// The payment method types to exclude from the session.
	ExcludedPaymentMethodTypes []*string `form:"excluded_payment_method_types" json:"excluded_payment_method_types,omitempty"`
}

// Updates a requested session
type DelegatedCheckoutRequestedSessionUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The details of the fulfillment.
	FulfillmentDetails *DelegatedCheckoutRequestedSessionUpdateFulfillmentDetailsParams `form:"fulfillment_details" json:"fulfillment_details,omitempty"`
	// The details of the line items.
	LineItemDetails []*DelegatedCheckoutRequestedSessionUpdateLineItemDetailParams `form:"line_item_details" json:"line_item_details,omitempty"`
	// The metadata for this requested session.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The payment method for this requested session.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// The payment method options for this requested session.
	PaymentMethodOptions *DelegatedCheckoutRequestedSessionUpdatePaymentMethodOptionsParams `form:"payment_method_options" json:"payment_method_options,omitempty"`
	// The shared metadata for this requested session.
	SharedMetadata map[string]string                                         `form:"shared_metadata" json:"shared_metadata,omitempty"`
	UnsetFields    []DelegatedCheckoutRequestedSessionUpdateParamsUnsetField `form:"-" json:"-"`
}

// DelegatedCheckoutRequestedSessionUpdateParamsUnsetField is the list of fields that can be cleared/unset on DelegatedCheckoutRequestedSessionUpdateParams.
type DelegatedCheckoutRequestedSessionUpdateParamsUnsetField string

const (
	DelegatedCheckoutRequestedSessionUpdateParamsUnsetFieldMetadata       DelegatedCheckoutRequestedSessionUpdateParamsUnsetField = "metadata"
	DelegatedCheckoutRequestedSessionUpdateParamsUnsetFieldSharedMetadata DelegatedCheckoutRequestedSessionUpdateParamsUnsetField = "shared_metadata"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *DelegatedCheckoutRequestedSessionUpdateParams) AddUnsetField(field DelegatedCheckoutRequestedSessionUpdateParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *DelegatedCheckoutRequestedSessionUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *DelegatedCheckoutRequestedSessionUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Context about where the attribution originated.
type DelegatedCheckoutRequestedSessionCreateAffiliateAttributionSourceParams struct {
	// The platform where the attribution originated.
	Platform *string `form:"platform" json:"platform,omitempty"`
	// The type of the attribution source.
	Type *string `form:"type" json:"type"`
	// The URL where the attribution originated.
	URL *string `form:"url" json:"url,omitempty"`
}

// Affiliate attribution data associated with this requested session.
type DelegatedCheckoutRequestedSessionCreateAffiliateAttributionParams struct {
	// Agent-scoped campaign identifier.
	CampaignID *string `form:"campaign_id" json:"campaign_id,omitempty"`
	// Agent-scoped creative identifier.
	CreativeID *string `form:"creative_id" json:"creative_id,omitempty"`
	// Timestamp when the attribution token expires.
	ExpiresAt *int64 `form:"expires_at" json:"expires_at"`
	// Agent-issued secret to validate the legitimacy of the source of this data.
	IdentificationToken *string `form:"identification_token" json:"identification_token"`
	// Timestamp for when the attribution token was issued.
	IssuedAt *int64 `form:"issued_at" json:"issued_at"`
	// Identifier for the attribution agent / affiliate network namespace.
	Provider *string `form:"provider" json:"provider"`
	// Agent-scoped affiliate/publisher identifier.
	PublisherID *string `form:"publisher_id" json:"publisher_id,omitempty"`
	// Freeform key/value pairs for additional non-sensitive per-agent data.
	SharedMetadata map[string]string `form:"shared_metadata" json:"shared_metadata,omitempty"`
	// Context about where the attribution originated.
	Source *DelegatedCheckoutRequestedSessionCreateAffiliateAttributionSourceParams `form:"source" json:"source,omitempty"`
	// Agent-scoped sub-tracking identifier.
	SubID *string `form:"sub_id" json:"sub_id,omitempty"`
	// Whether this is the first or last touchpoint.
	Touchpoint *string `form:"touchpoint" json:"touchpoint"`
}

// The details of the fulfillment.
type DelegatedCheckoutRequestedSessionCreateFulfillmentDetailsParams struct {
	// The customer's address.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// The customer's email address.
	Email *string `form:"email" json:"email,omitempty"`
	// The customer's name.
	Name *string `form:"name" json:"name,omitempty"`
	// The customer's phone number.
	Phone *string `form:"phone" json:"phone,omitempty"`
}

// The details of the line items.
type DelegatedCheckoutRequestedSessionCreateLineItemDetailParams struct {
	// The quantity of the line item.
	Quantity *int64 `form:"quantity" json:"quantity"`
	// The SKU ID of the line item.
	SKUID *string `form:"sku_id" json:"sku_id"`
}

// Card-specific payment method options.
type DelegatedCheckoutRequestedSessionCreatePaymentMethodOptionsCardParams struct {
	// The card brands to exclude from the session.
	BrandsBlocked []*string `form:"brands_blocked" json:"brands_blocked,omitempty"`
}

// The payment method options for this requested session.
type DelegatedCheckoutRequestedSessionCreatePaymentMethodOptionsParams struct {
	// Card-specific payment method options.
	Card *DelegatedCheckoutRequestedSessionCreatePaymentMethodOptionsCardParams `form:"card" json:"card,omitempty"`
	// The payment method types to exclude from the session.
	ExcludedPaymentMethodTypes []*string `form:"excluded_payment_method_types" json:"excluded_payment_method_types,omitempty"`
}

// The details of the seller.
type DelegatedCheckoutRequestedSessionCreateSellerDetailsParams struct {
	// The network profile for the seller.
	NetworkProfile *string `form:"network_profile" json:"network_profile"`
}

// Creates a requested session
type DelegatedCheckoutRequestedSessionCreateParams struct {
	Params `form:"*"`
	// Affiliate attribution data associated with this requested session.
	AffiliateAttribution *DelegatedCheckoutRequestedSessionCreateAffiliateAttributionParams `form:"affiliate_attribution" json:"affiliate_attribution,omitempty"`
	// The currency for this requested session.
	Currency *string `form:"currency" json:"currency"`
	// The customer for this requested session.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The details of the fulfillment.
	FulfillmentDetails *DelegatedCheckoutRequestedSessionCreateFulfillmentDetailsParams `form:"fulfillment_details" json:"fulfillment_details,omitempty"`
	// The details of the line items.
	LineItemDetails []*DelegatedCheckoutRequestedSessionCreateLineItemDetailParams `form:"line_item_details" json:"line_item_details"`
	// The metadata for this requested session.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The payment method for this requested session.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// The payment method options for this requested session.
	PaymentMethodOptions *DelegatedCheckoutRequestedSessionCreatePaymentMethodOptionsParams `form:"payment_method_options" json:"payment_method_options,omitempty"`
	// The details of the seller.
	SellerDetails *DelegatedCheckoutRequestedSessionCreateSellerDetailsParams `form:"seller_details" json:"seller_details"`
	// The setup future usage for this requested session.
	SetupFutureUsage *string `form:"setup_future_usage" json:"setup_future_usage,omitempty"`
	// The shared metadata for this requested session.
	SharedMetadata map[string]string `form:"shared_metadata" json:"shared_metadata,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *DelegatedCheckoutRequestedSessionCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *DelegatedCheckoutRequestedSessionCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Context about where the attribution originated.
type DelegatedCheckoutRequestedSessionAffiliateAttributionSource struct {
	// The platform of the attribution source.
	Platform string `json:"platform"`
	// The type of the attribution source.
	Type DelegatedCheckoutRequestedSessionAffiliateAttributionSourceType `json:"type"`
	// The URL of the attribution source.
	URL string `json:"url"`
}

// Affiliate attribution data associated with this requested session.
type DelegatedCheckoutRequestedSessionAffiliateAttribution struct {
	// Agent-scoped campaign identifier.
	CampaignID string `json:"campaign_id"`
	// Agent-scoped creative identifier.
	CreativeID string `json:"creative_id"`
	// Timestamp when the attribution token expires.
	ExpiresAt int64 `json:"expires_at"`
	// Agent-issued secret to validate the legitimacy of the source of this data.
	IdentificationToken string `json:"identification_token"`
	// Timestamp for when the attribution token was issued.
	IssuedAt int64 `json:"issued_at"`
	// Identifier for the attribution agent / affiliate network namespace.
	Provider string `json:"provider"`
	// Agent-scoped affiliate/publisher identifier.
	PublisherID string `json:"publisher_id"`
	// Freeform key/value pairs for additional non-sensitive per-agent data.
	SharedMetadata map[string]string `json:"shared_metadata"`
	// Context about where the attribution originated.
	Source *DelegatedCheckoutRequestedSessionAffiliateAttributionSource `json:"source"`
	// Agent-scoped sub-tracking identifier.
	SubID string `json:"sub_id"`
	// Whether this is the first or last touchpoint.
	Touchpoint DelegatedCheckoutRequestedSessionAffiliateAttributionTouchpoint `json:"touchpoint"`
}

// The buyer's marketing consent choices.
type DelegatedCheckoutRequestedSessionBuyerConsentsMarketingConsent struct {
	// The marketing consent channel.
	Channel DelegatedCheckoutRequestedSessionBuyerConsentsMarketingConsentChannel `json:"channel"`
	// The consent status. 'granted' means the buyer opted in, 'none' means they did not.
	Status DelegatedCheckoutRequestedSessionBuyerConsentsMarketingConsentStatus `json:"status"`
}

// The available marketing consent options.
type DelegatedCheckoutRequestedSessionBuyerConsentsMarketingOption struct {
	// The marketing channel type.
	Channel DelegatedCheckoutRequestedSessionBuyerConsentsMarketingOptionChannel `json:"channel"`
	// The description of the marketing consent option.
	Description string `json:"description"`
	// The privacy policy URL for this marketing channel.
	PrivacyPolicyURL string `json:"privacy_policy_url"`
}

// The marketing consent options.
type DelegatedCheckoutRequestedSessionBuyerConsentsMarketing struct {
	// The buyer's marketing consent choices.
	Consents []*DelegatedCheckoutRequestedSessionBuyerConsentsMarketingConsent `json:"consents"`
	// The available marketing consent options.
	Options []*DelegatedCheckoutRequestedSessionBuyerConsentsMarketingOption `json:"options"`
}

// The buyer consent options for this requested session, including marketing preferences.
type DelegatedCheckoutRequestedSessionBuyerConsents struct {
	// The marketing consent options.
	Marketing *DelegatedCheckoutRequestedSessionBuyerConsentsMarketing `json:"marketing"`
}

// The digital options.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptionDigitalDigitalOption struct {
	// The description of the digital fulfillment option.
	Description string `json:"description"`
	// The digital amount of the digital fulfillment option.
	DigitalAmount int64 `json:"digital_amount"`
	// The display name of the digital fulfillment option.
	DisplayName string `json:"display_name"`
	// The key of the digital fulfillment option.
	Key string `json:"key"`
	// The line item keys associated with this digital fulfillment option.
	LineItemKeys []string `json:"line_item_keys"`
}

// The digital fulfillment option.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptionDigital struct {
	// The digital options.
	DigitalOptions []*DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptionDigitalDigitalOption `json:"digital_options"`
}

// The shipping options.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptionShippingShippingOption struct {
	// The description of the shipping option.
	Description string `json:"description"`
	// The display name of the shipping option.
	DisplayName string `json:"display_name"`
	// The earliest delivery time of the shipping option.
	EarliestDeliveryTime int64 `json:"earliest_delivery_time"`
	// The key of the shipping option.
	Key string `json:"key"`
	// The latest delivery time of the shipping option.
	LatestDeliveryTime int64 `json:"latest_delivery_time"`
	// The line item keys associated with this shipping option.
	LineItemKeys []string `json:"line_item_keys"`
	// The shipping amount of the shipping option.
	ShippingAmount int64 `json:"shipping_amount"`
}

// The shipping option.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptionShipping struct {
	// The shipping options.
	ShippingOptions []*DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptionShippingShippingOption `json:"shipping_options"`
}

// The fulfillment options.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOption struct {
	// The digital fulfillment option.
	Digital *DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptionDigital `json:"digital"`
	// The shipping option.
	Shipping *DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptionShipping `json:"shipping"`
	// The type of the fulfillment option.
	Type DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptionType `json:"type"`
}

// The digital fulfillment option.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionDigital struct {
	// The digital option.
	DigitalOption string `json:"digital_option"`
}

// The shipping option.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionShipping struct {
	// The shipping option.
	ShippingOption string `json:"shipping_option"`
}

// The selected fulfillment option.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOption struct {
	// The digital fulfillment option.
	Digital *DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionDigital `json:"digital"`
	// The shipping option.
	Shipping *DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionShipping `json:"shipping"`
	// The type of the selected fulfillment option.
	Type DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionType `json:"type"`
}

// The digital fulfillment option.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionOverrideDigital struct {
	// The digital option.
	DigitalOption string `json:"digital_option"`
}

// The shipping option.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionOverrideShipping struct {
	// The shipping option.
	ShippingOption string `json:"shipping_option"`
}

// Per-item fulfillment option overrides.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionOverride struct {
	// The digital fulfillment option.
	Digital *DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionOverrideDigital `json:"digital"`
	// The line items this fulfillment option applies to.
	LineItemKeys []string `json:"line_item_keys"`
	// The shipping option.
	Shipping *DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionOverrideShipping `json:"shipping"`
	// The type of the selected fulfillment option.
	Type DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionOverrideType `json:"type"`
}

// The details of the fulfillment.
type DelegatedCheckoutRequestedSessionFulfillmentDetails struct {
	// The fulfillment address.
	Address *Address `json:"address"`
	// The email address for the fulfillment details.
	Email string `json:"email"`
	// The fulfillment options.
	FulfillmentOptions []*DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOption `json:"fulfillment_options"`
	// The name for the fulfillment details.
	Name string `json:"name"`
	// The phone number for the fulfillment details.
	Phone string `json:"phone"`
	// The selected fulfillment option.
	SelectedFulfillmentOption *DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOption `json:"selected_fulfillment_option"`
	// Per-item fulfillment option overrides.
	SelectedFulfillmentOptionOverrides []*DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionOverride `json:"selected_fulfillment_option_overrides"`
}

// Custom attributes for the product.
type DelegatedCheckoutRequestedSessionLineItemDetailProductDetailsCustomAttribute struct {
	// The display name of the custom attribute.
	DisplayName string `json:"display_name"`
	// The value of the custom attribute.
	Value string `json:"value"`
}

// Disclosures for the product.
type DelegatedCheckoutRequestedSessionLineItemDetailProductDetailsDisclosure struct {
	// The content of the disclosure.
	Content string `json:"content"`
	// The content type of the disclosure.
	ContentType DelegatedCheckoutRequestedSessionLineItemDetailProductDetailsDisclosureContentType `json:"content_type"`
	// The type of disclosure.
	Type DelegatedCheckoutRequestedSessionLineItemDetailProductDetailsDisclosureType `json:"type"`
}
type DelegatedCheckoutRequestedSessionLineItemDetailProductDetails struct {
	// Custom attributes for the product.
	CustomAttributes []*DelegatedCheckoutRequestedSessionLineItemDetailProductDetailsCustomAttribute `json:"custom_attributes"`
	// The description of the product.
	Description string `json:"description"`
	// Disclosures for the product.
	Disclosures []*DelegatedCheckoutRequestedSessionLineItemDetailProductDetailsDisclosure `json:"disclosures"`
	// The images of the product.
	Images []string `json:"images"`
	// The title of the product.
	Title string `json:"title"`
}

// The line items to be purchased.
type DelegatedCheckoutRequestedSessionLineItemDetail struct {
	// The total discount for this line item. If no discount were applied, defaults to 0.
	AmountDiscount int64 `json:"amount_discount"`
	// The total before any discounts or taxes are applied.
	AmountSubtotal int64 `json:"amount_subtotal"`
	// The fulfillment type of the line item.
	FulfillmentType string `json:"fulfillment_type"`
	// The key of the line item.
	Key            string                                                         `json:"key"`
	ProductDetails *DelegatedCheckoutRequestedSessionLineItemDetailProductDetails `json:"product_details,omitempty"`
	// The quantity of the line item.
	Quantity int64 `json:"quantity"`
	// The SKU ID of the line item.
	SKUID string `json:"sku_id"`
	// The per-unit amount of the item before any discounts or taxes are applied.
	UnitAmount int64 `json:"unit_amount"`
}

// The details of the order.
type DelegatedCheckoutRequestedSessionOrderDetails struct {
	// The seller's order identifier.
	OrderID string `json:"order_id"`
	// The URL to the order status.
	OrderStatusURL string `json:"order_status_url"`
}

// Card-specific payment method options.
type DelegatedCheckoutRequestedSessionPaymentMethodOptionsCard struct {
	// The card brands blocked by the agent.
	BrandsBlocked []DelegatedCheckoutRequestedSessionPaymentMethodOptionsCardBrandsBlocked `json:"brands_blocked"`
}

// The payment method options for this requested session.
type DelegatedCheckoutRequestedSessionPaymentMethodOptions struct {
	// Card-specific payment method options.
	Card *DelegatedCheckoutRequestedSessionPaymentMethodOptionsCard `json:"card"`
	// The computed displayable card brands.
	DisplayableCardBrands []DelegatedCheckoutRequestedSessionPaymentMethodOptionsDisplayableCardBrand `json:"displayable_card_brands"`
	// The computed displayable payment method types.
	DisplayablePaymentMethodTypes []DelegatedCheckoutRequestedSessionPaymentMethodOptionsDisplayablePaymentMethodType `json:"displayable_payment_method_types"`
	// The payment method types excluded by the agent.
	ExcludedPaymentMethodTypes []DelegatedCheckoutRequestedSessionPaymentMethodOptionsExcludedPaymentMethodType `json:"excluded_payment_method_types"`
}

// The billing details of the payment method.
type DelegatedCheckoutRequestedSessionPaymentMethodPreviewBillingDetails struct {
	// The billing address.
	Address *Address `json:"address"`
	// The email address for the billing details.
	Email string `json:"email"`
	// The name for the billing details.
	Name string `json:"name"`
	// The phone number for the billing details.
	Phone string `json:"phone"`
}

// The card details of the payment method.
type DelegatedCheckoutRequestedSessionPaymentMethodPreviewCard struct {
	// The expiry month of the card.
	ExpMonth int64 `json:"exp_month"`
	// The expiry year of the card.
	ExpYear int64 `json:"exp_year"`
	// The last 4 digits of the card number.
	Last4 string `json:"last4"`
}

// The preview of the payment method to be created when the requested session is confirmed.
type DelegatedCheckoutRequestedSessionPaymentMethodPreview struct {
	// The billing details of the payment method.
	BillingDetails *DelegatedCheckoutRequestedSessionPaymentMethodPreviewBillingDetails `json:"billing_details"`
	// The card details of the payment method.
	Card *DelegatedCheckoutRequestedSessionPaymentMethodPreviewCard `json:"card"`
	// The type of the payment method.
	Type string `json:"type"`
}

// The risk metadata for the client device.
type DelegatedCheckoutRequestedSessionRiskDetailsClientDeviceMetadataDetails struct {
	// The radar session for the client device.
	RadarSession string `json:"radar_session"`
	// The referrer of the client device.
	Referrer string `json:"referrer"`
	// The remote IP address of the client device.
	RemoteIP string `json:"remote_ip"`
	// The time spent on the page by the client device.
	TimeOnPageMS int64 `json:"time_on_page_ms"`
	// The user agent of the client device.
	UserAgent string `json:"user_agent"`
}

// The risk details of the requested session.
type DelegatedCheckoutRequestedSessionRiskDetails struct {
	// The risk metadata for the client device.
	ClientDeviceMetadataDetails *DelegatedCheckoutRequestedSessionRiskDetailsClientDeviceMetadataDetails `json:"client_device_metadata_details"`
}

// The marketplace seller details.
type DelegatedCheckoutRequestedSessionSellerDetailsMarketplaceSellerDetails struct{}
type DelegatedCheckoutRequestedSessionSellerDetails struct {
	// The card brands supported by the seller.
	CardBrands []DelegatedCheckoutRequestedSessionSellerDetailsCardBrand `json:"card_brands"`
	// The marketplace seller details.
	MarketplaceSellerDetails *DelegatedCheckoutRequestedSessionSellerDetailsMarketplaceSellerDetails `json:"marketplace_seller_details"`
	// The network profile of the seller.
	NetworkProfile *Profile `json:"network_profile"`
	// The payment method types supported by the seller.
	PaymentMethodTypes []DelegatedCheckoutRequestedSessionSellerDetailsPaymentMethodType `json:"payment_method_types"`
	// The URL to the seller's privacy notice.
	PrivacyNoticeURL string `json:"privacy_notice_url"`
	// The URL to the seller's return policy.
	ReturnPolicyURL string `json:"return_policy_url"`
	// The URL to the seller's store policy.
	StorePolicyURL string `json:"store_policy_url"`
	// The URL to the seller's terms of service.
	TermsOfServiceURL string `json:"terms_of_service_url"`
}

// The applicable fees of the total details.
type DelegatedCheckoutRequestedSessionTotalDetailsApplicableFee struct {
	// The amount of the applicable fee.
	Amount int64 `json:"amount"`
	// The description of the applicable fee.
	Description string `json:"description"`
	// The display name of the applicable fee.
	DisplayName string `json:"display_name"`
}
type DelegatedCheckoutRequestedSessionTotalDetails struct {
	// The amount of order-level discounts applied to the cart. The total discount amount for this session can be computed by summing the cart discount and the item discounts.
	AmountCartDiscount int64 `json:"amount_cart_discount"`
	// The amount fulfillment of the total details.
	AmountFulfillment int64 `json:"amount_fulfillment"`
	// The amount of item-level discounts applied to the cart. The total discount amount for this session can be computed by summing the cart discount and the item discounts.
	AmountItemsDiscount int64 `json:"amount_items_discount"`
	// The amount tax of the total details.
	AmountTax int64 `json:"amount_tax"`
	// The applicable fees of the total details.
	ApplicableFees []*DelegatedCheckoutRequestedSessionTotalDetailsApplicableFee `json:"applicable_fees"`
}

// A requested session is a session that has been requested by a customer.
type DelegatedCheckoutRequestedSession struct {
	APIResource
	// Affiliate attribution data associated with this requested session.
	AffiliateAttributions []*DelegatedCheckoutRequestedSessionAffiliateAttribution `json:"affiliate_attributions,omitempty"`
	// The subtotal amount of the requested session.
	AmountSubtotal int64 `json:"amount_subtotal"`
	// The total amount of the requested session.
	AmountTotal int64 `json:"amount_total"`
	// The buyer consent options for this requested session, including marketing preferences.
	BuyerConsents *DelegatedCheckoutRequestedSessionBuyerConsents `json:"buyer_consents"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	CreatedAt int64 `json:"created_at"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The customer for this requested session.
	Customer string `json:"customer"`
	// Time at which the requested session expires. Measured in seconds since the Unix epoch.
	ExpiresAt int64 `json:"expires_at"`
	// The details of the fulfillment.
	FulfillmentDetails *DelegatedCheckoutRequestedSessionFulfillmentDetails `json:"fulfillment_details"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The line items to be purchased.
	LineItemDetails []*DelegatedCheckoutRequestedSessionLineItemDetail `json:"line_item_details"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The details of the order.
	OrderDetails *DelegatedCheckoutRequestedSessionOrderDetails `json:"order_details"`
	// The payment method used for the requested session.
	PaymentMethod string `json:"payment_method"`
	// The payment method options for this requested session.
	PaymentMethodOptions *DelegatedCheckoutRequestedSessionPaymentMethodOptions `json:"payment_method_options"`
	// The preview of the payment method to be created when the requested session is confirmed.
	PaymentMethodPreview *DelegatedCheckoutRequestedSessionPaymentMethodPreview `json:"payment_method_preview"`
	// The risk details of the requested session.
	RiskDetails   *DelegatedCheckoutRequestedSessionRiskDetails   `json:"risk_details"`
	SellerDetails *DelegatedCheckoutRequestedSessionSellerDetails `json:"seller_details"`
	// Whether or not the payment method should be saved for future use.
	SetupFutureUsage DelegatedCheckoutRequestedSessionSetupFutureUsage `json:"setup_future_usage"`
	// The metadata shared with the seller.
	SharedMetadata map[string]string `json:"shared_metadata"`
	// The SPT used for payment.
	SharedPaymentIssuedToken *SharedPaymentIssuedToken `json:"shared_payment_issued_token"`
	// The status of the requested session.
	Status       DelegatedCheckoutRequestedSessionStatus        `json:"status"`
	TotalDetails *DelegatedCheckoutRequestedSessionTotalDetails `json:"total_details"`
	// Time at which the object was last updated. Measured in seconds since the Unix epoch.
	UpdatedAt int64 `json:"updated_at"`
}
