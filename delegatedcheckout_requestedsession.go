//
//
// File generated from our OpenAPI spec
//
//

package stripe

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
	DelegatedCheckoutRequestedSessionStatusCompleted DelegatedCheckoutRequestedSessionStatus = "completed"
	DelegatedCheckoutRequestedSessionStatusExpired   DelegatedCheckoutRequestedSessionStatus = "expired"
	DelegatedCheckoutRequestedSessionStatusOpen      DelegatedCheckoutRequestedSessionStatus = "open"
)

// Retrieves a requested session
type DelegatedCheckoutRequestedSessionParams struct {
	Params `form:"*"`
	// The currency for this requested session.
	Currency *string `form:"currency"`
	// The customer for this requested session.
	Customer *string `form:"customer"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The details of the fulfillment.
	FulfillmentDetails *DelegatedCheckoutRequestedSessionFulfillmentDetailsParams `form:"fulfillment_details"`
	// The details of the line items.
	LineItemDetails []*DelegatedCheckoutRequestedSessionLineItemDetailParams `form:"line_item_details"`
	// The metadata for this requested session.
	Metadata map[string]string `form:"metadata"`
	// The payment method for this requested session.
	PaymentMethod *string `form:"payment_method"`
	// The payment method data for this requested session.
	PaymentMethodData *DelegatedCheckoutRequestedSessionPaymentMethodDataParams `form:"payment_method_data"`
	// The details of the seller.
	SellerDetails *DelegatedCheckoutRequestedSessionSellerDetailsParams `form:"seller_details"`
	// The setup future usage for this requested session.
	SetupFutureUsage *string `form:"setup_future_usage"`
	// The shared metadata for this requested session.
	SharedMetadata map[string]string `form:"shared_metadata"`
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

// The shipping fulfillment option.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionShippingParams struct {
	// The shipping option identifer.
	ShippingOption *string `form:"shipping_option"`
}

// The fulfillment option to select.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionParams struct {
	// The shipping fulfillment option.
	Shipping *DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionShippingParams `form:"shipping"`
	// The type of fulfillment option.
	Type *string `form:"type"`
}

// The details of the fulfillment.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsParams struct {
	// The customer's address.
	Address *AddressParams `form:"address"`
	// The customer's email address.
	Email *string `form:"email"`
	// The customer's name.
	Name *string `form:"name"`
	// The customer's phone number.
	Phone *string `form:"phone"`
	// The fulfillment option to select.
	SelectedFulfillmentOption *DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionParams `form:"selected_fulfillment_option"`
}

// The details of the line items.
type DelegatedCheckoutRequestedSessionLineItemDetailParams struct {
	// The key of the line item.
	Key *string `form:"key"`
	// The quantity of the line item.
	Quantity *int64 `form:"quantity"`
	// The SKU ID of the line item.
	SKUID *string `form:"sku_id"`
}

// The billing details for the payment method data.
type DelegatedCheckoutRequestedSessionPaymentMethodDataBillingDetailsParams struct {
	// The address for the billing details.
	Address *AddressParams `form:"address"`
	// The email for the billing details.
	Email *string `form:"email"`
	// The name for the billing details.
	Name *string `form:"name"`
	// The phone for the billing details.
	Phone *string `form:"phone"`
}

// The card for the payment method data.
type DelegatedCheckoutRequestedSessionPaymentMethodDataCardParams struct {
	// The CVC of the card.
	CVC *string `form:"cvc"`
	// The expiration month of the card.
	ExpMonth *int64 `form:"exp_month"`
	// The expiration year of the card.
	ExpYear *int64 `form:"exp_year"`
	// The number of the card.
	Number *string `form:"number"`
}

// The payment method data for this requested session.
type DelegatedCheckoutRequestedSessionPaymentMethodDataParams struct {
	// The billing details for the payment method data.
	BillingDetails *DelegatedCheckoutRequestedSessionPaymentMethodDataBillingDetailsParams `form:"billing_details"`
	// The card for the payment method data.
	Card *DelegatedCheckoutRequestedSessionPaymentMethodDataCardParams `form:"card"`
	// The type of the payment method data.
	Type *string `form:"type"`
}

// The details of the seller.
type DelegatedCheckoutRequestedSessionSellerDetailsParams struct {
	// The network profile for the seller.
	NetworkProfile *string `form:"network_profile"`
}

// The client device metadata details for this requested session.
type DelegatedCheckoutRequestedSessionConfirmRiskDetailsClientDeviceMetadataDetailsParams struct {
	// The radar session.
	RadarSession *string `form:"radar_session"`
	// The referrer of the client device.
	Referrer *string `form:"referrer"`
	// The remote IP address of the client device.
	RemoteIP *string `form:"remote_ip"`
	// The time on page in milliseconds.
	TimeOnPageMS *int64 `form:"time_on_page_ms"`
	// The user agent of the client device.
	UserAgent *string `form:"user_agent"`
}

// Risk details/signals associated with the requested session
type DelegatedCheckoutRequestedSessionConfirmRiskDetailsParams struct {
	// The client device metadata details for this requested session.
	ClientDeviceMetadataDetails *DelegatedCheckoutRequestedSessionConfirmRiskDetailsClientDeviceMetadataDetailsParams `form:"client_device_metadata_details"`
}

// Confirms a requested session
type DelegatedCheckoutRequestedSessionConfirmParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The PaymentMethod to use with the requested session.
	PaymentMethod *string `form:"payment_method"`
	// Risk details/signals associated with the requested session
	RiskDetails *DelegatedCheckoutRequestedSessionConfirmRiskDetailsParams `form:"risk_details"`
}

// AddExpand appends a new field to expand.
func (p *DelegatedCheckoutRequestedSessionConfirmParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Expires a requested session
type DelegatedCheckoutRequestedSessionExpireParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *DelegatedCheckoutRequestedSessionExpireParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves a requested session
type DelegatedCheckoutRequestedSessionRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *DelegatedCheckoutRequestedSessionRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The shipping fulfillment option.
type DelegatedCheckoutRequestedSessionUpdateFulfillmentDetailsSelectedFulfillmentOptionShippingParams struct {
	// The shipping option identifer.
	ShippingOption *string `form:"shipping_option"`
}

// The fulfillment option to select.
type DelegatedCheckoutRequestedSessionUpdateFulfillmentDetailsSelectedFulfillmentOptionParams struct {
	// The shipping fulfillment option.
	Shipping *DelegatedCheckoutRequestedSessionUpdateFulfillmentDetailsSelectedFulfillmentOptionShippingParams `form:"shipping"`
	// The type of fulfillment option.
	Type *string `form:"type"`
}

// The details of the fulfillment.
type DelegatedCheckoutRequestedSessionUpdateFulfillmentDetailsParams struct {
	// The customer's address.
	Address *AddressParams `form:"address"`
	// The customer's email address.
	Email *string `form:"email"`
	// The customer's name.
	Name *string `form:"name"`
	// The customer's phone number.
	Phone *string `form:"phone"`
	// The fulfillment option to select.
	SelectedFulfillmentOption *DelegatedCheckoutRequestedSessionUpdateFulfillmentDetailsSelectedFulfillmentOptionParams `form:"selected_fulfillment_option"`
}

// The details of the line items.
type DelegatedCheckoutRequestedSessionUpdateLineItemDetailParams struct {
	// The key of the line item.
	Key *string `form:"key"`
	// The quantity of the line item.
	Quantity *int64 `form:"quantity"`
}

// The billing details for the payment method data.
type DelegatedCheckoutRequestedSessionUpdatePaymentMethodDataBillingDetailsParams struct {
	// The address for the billing details.
	Address *AddressParams `form:"address"`
	// The email for the billing details.
	Email *string `form:"email"`
	// The name for the billing details.
	Name *string `form:"name"`
	// The phone for the billing details.
	Phone *string `form:"phone"`
}

// The card for the payment method data.
type DelegatedCheckoutRequestedSessionUpdatePaymentMethodDataCardParams struct {
	// The CVC of the card.
	CVC *string `form:"cvc"`
	// The expiration month of the card.
	ExpMonth *int64 `form:"exp_month"`
	// The expiration year of the card.
	ExpYear *int64 `form:"exp_year"`
	// The number of the card.
	Number *string `form:"number"`
}

// The payment method data for this requested session.
type DelegatedCheckoutRequestedSessionUpdatePaymentMethodDataParams struct {
	// The billing details for the payment method data.
	BillingDetails *DelegatedCheckoutRequestedSessionUpdatePaymentMethodDataBillingDetailsParams `form:"billing_details"`
	// The card for the payment method data.
	Card *DelegatedCheckoutRequestedSessionUpdatePaymentMethodDataCardParams `form:"card"`
	// The type of the payment method data.
	Type *string `form:"type"`
}

// Updates a requested session
type DelegatedCheckoutRequestedSessionUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The details of the fulfillment.
	FulfillmentDetails *DelegatedCheckoutRequestedSessionUpdateFulfillmentDetailsParams `form:"fulfillment_details"`
	// The details of the line items.
	LineItemDetails []*DelegatedCheckoutRequestedSessionUpdateLineItemDetailParams `form:"line_item_details"`
	// The metadata for this requested session.
	Metadata map[string]string `form:"metadata"`
	// The payment method for this requested session.
	PaymentMethod *string `form:"payment_method"`
	// The payment method data for this requested session.
	PaymentMethodData *DelegatedCheckoutRequestedSessionUpdatePaymentMethodDataParams `form:"payment_method_data"`
	// The shared metadata for this requested session.
	SharedMetadata map[string]string `form:"shared_metadata"`
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

// The details of the fulfillment.
type DelegatedCheckoutRequestedSessionCreateFulfillmentDetailsParams struct {
	// The customer's address.
	Address *AddressParams `form:"address"`
	// The customer's email address.
	Email *string `form:"email"`
	// The customer's name.
	Name *string `form:"name"`
	// The customer's phone number.
	Phone *string `form:"phone"`
}

// The details of the line items.
type DelegatedCheckoutRequestedSessionCreateLineItemDetailParams struct {
	// The quantity of the line item.
	Quantity *int64 `form:"quantity"`
	// The SKU ID of the line item.
	SKUID *string `form:"sku_id"`
}

// The billing details for the payment method data.
type DelegatedCheckoutRequestedSessionCreatePaymentMethodDataBillingDetailsParams struct {
	// The address for the billing details.
	Address *AddressParams `form:"address"`
	// The email for the billing details.
	Email *string `form:"email"`
	// The name for the billing details.
	Name *string `form:"name"`
	// The phone for the billing details.
	Phone *string `form:"phone"`
}

// The card for the payment method data.
type DelegatedCheckoutRequestedSessionCreatePaymentMethodDataCardParams struct {
	// The CVC of the card.
	CVC *string `form:"cvc"`
	// The expiration month of the card.
	ExpMonth *int64 `form:"exp_month"`
	// The expiration year of the card.
	ExpYear *int64 `form:"exp_year"`
	// The number of the card.
	Number *string `form:"number"`
}

// The payment method data for this requested session.
type DelegatedCheckoutRequestedSessionCreatePaymentMethodDataParams struct {
	// The billing details for the payment method data.
	BillingDetails *DelegatedCheckoutRequestedSessionCreatePaymentMethodDataBillingDetailsParams `form:"billing_details"`
	// The card for the payment method data.
	Card *DelegatedCheckoutRequestedSessionCreatePaymentMethodDataCardParams `form:"card"`
	// The type of the payment method data.
	Type *string `form:"type"`
}

// The details of the seller.
type DelegatedCheckoutRequestedSessionCreateSellerDetailsParams struct {
	// The network profile for the seller.
	NetworkProfile *string `form:"network_profile"`
}

// Creates a requested session
type DelegatedCheckoutRequestedSessionCreateParams struct {
	Params `form:"*"`
	// The currency for this requested session.
	Currency *string `form:"currency"`
	// The customer for this requested session.
	Customer *string `form:"customer"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The details of the fulfillment.
	FulfillmentDetails *DelegatedCheckoutRequestedSessionCreateFulfillmentDetailsParams `form:"fulfillment_details"`
	// The details of the line items.
	LineItemDetails []*DelegatedCheckoutRequestedSessionCreateLineItemDetailParams `form:"line_item_details"`
	// The metadata for this requested session.
	Metadata map[string]string `form:"metadata"`
	// The payment method for this requested session.
	PaymentMethod *string `form:"payment_method"`
	// The payment method data for this requested session.
	PaymentMethodData *DelegatedCheckoutRequestedSessionCreatePaymentMethodDataParams `form:"payment_method_data"`
	// The details of the seller.
	SellerDetails *DelegatedCheckoutRequestedSessionCreateSellerDetailsParams `form:"seller_details"`
	// The setup future usage for this requested session.
	SetupFutureUsage *string `form:"setup_future_usage"`
	// The shared metadata for this requested session.
	SharedMetadata map[string]string `form:"shared_metadata"`
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
	// The shipping option.
	Shipping *DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptionShipping `json:"shipping"`
	// The type of the fulfillment option.
	Type string `json:"type"`
}

// The shipping option.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionShipping struct {
	// The shipping option.
	ShippingOption string `json:"shipping_option"`
}

// The fulfillment option.
type DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOption struct {
	// The shipping option.
	Shipping *DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionShipping `json:"shipping"`
	// The type of the selected fulfillment option.
	Type string `json:"type"`
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
	// The fulfillment option.
	SelectedFulfillmentOption *DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOption `json:"selected_fulfillment_option"`
}

// The line items to be purchased.
type DelegatedCheckoutRequestedSessionLineItemDetail struct {
	// The total discount for this line item. If no discount were applied, defaults to 0.
	AmountDiscount int64 `json:"amount_discount"`
	// The total before any discounts or taxes are applied.
	AmountSubtotal int64 `json:"amount_subtotal"`
	// The total after discounts but before taxes are applied.
	AmountSubtotalAfterDiscount int64 `json:"amount_subtotal_after_discount"`
	// The total after discounts and taxes.
	AmountTotal int64 `json:"amount_total"`
	// The description of the line item.
	Description string `json:"description"`
	// The images of the line item.
	Images []string `json:"images"`
	// The key of the line item.
	Key string `json:"key"`
	// The name of the line item.
	Name string `json:"name"`
	// The quantity of the line item.
	Quantity int64 `json:"quantity"`
	// The SKU ID of the line item.
	SKUID string `json:"sku_id"`
	// The per-unit amount of the item before any discounts or taxes are applied.
	UnitAmount int64 `json:"unit_amount"`
	// The per-unit amount of the item after discounts but before taxes are applied.
	UnitAmountAfterDiscount int64 `json:"unit_amount_after_discount"`
	// The per-unit discount amount. If no discount were applied, defaults to 0.
	UnitDiscount int64 `json:"unit_discount"`
}

// The details of the order.
type DelegatedCheckoutRequestedSessionOrderDetails struct {
	// The seller's order identifier.
	OrderID string `json:"order_id"`
	// The URL to the order status.
	OrderStatusURL string `json:"order_status_url"`
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
type DelegatedCheckoutRequestedSessionSellerDetails struct{}

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
	// The amount discount of the total details.
	AmountDiscount int64 `json:"amount_discount"`
	// The amount fulfillment of the total details.
	AmountFulfillment int64 `json:"amount_fulfillment"`
	// Total of all items after discounts but before taxes are applied.
	AmountSubtotalAfterDiscount int64 `json:"amount_subtotal_after_discount"`
	// The amount tax of the total details.
	AmountTax int64 `json:"amount_tax"`
	// The applicable fees of the total details.
	ApplicableFees []*DelegatedCheckoutRequestedSessionTotalDetailsApplicableFee `json:"applicable_fees"`
}

// A requested session is a session that has been requested by a customer.
type DelegatedCheckoutRequestedSession struct {
	APIResource
	// The subtotal amount of the requested session.
	AmountSubtotal int64 `json:"amount_subtotal"`
	// The total amount of the requested session.
	AmountTotal int64 `json:"amount_total"`
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
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The details of the order.
	OrderDetails *DelegatedCheckoutRequestedSessionOrderDetails `json:"order_details"`
	// The payment method used for the requested session.
	PaymentMethod string `json:"payment_method"`
	// The preview of the payment method to be created when the requested session is confirmed.
	PaymentMethodPreview *DelegatedCheckoutRequestedSessionPaymentMethodPreview `json:"payment_method_preview"`
	SellerDetails        *DelegatedCheckoutRequestedSessionSellerDetails        `json:"seller_details"`
	// Whether or not the payment method should be saved for future use.
	SetupFutureUsage DelegatedCheckoutRequestedSessionSetupFutureUsage `json:"setup_future_usage"`
	// The metadata shared with the seller.
	SharedMetadata map[string]string `json:"shared_metadata"`
	// The SPT used for payment.
	SharedPaymentIssuedToken string `json:"shared_payment_issued_token"`
	// The status of the requested session.
	Status       DelegatedCheckoutRequestedSessionStatus        `json:"status"`
	TotalDetails *DelegatedCheckoutRequestedSessionTotalDetails `json:"total_details"`
	// Time at which the object was last updated. Measured in seconds since the Unix epoch.
	UpdatedAt int64 `json:"updated_at"`
}
