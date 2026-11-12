//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of OffSessionPayments matching a filter.
type V2PaymentsOffSessionPaymentListParams struct {
	Params `form:"*"`
	// The page size limit. If not provided, the default is 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Settings for automatic tax calculation.
type V2PaymentsOffSessionPaymentAutomaticTaxParams struct {
	// Whether to automatically calculate taxes for this OffSessionPayment.
	Enabled *bool `form:"enabled" json:"enabled"`
}

// Tax details for this product, including the [tax code](https://docs.stripe.com/tax/tax-codes)
// and an optional performance location.
type V2PaymentsOffSessionPaymentItemOneTimePriceItemPriceDataProductDataTaxDetailsParams struct {
	// A tax location ID. Depending on the [tax code](https://docs.stripe.com/tax/tax-for-tickets/reference/tax-location-performance),
	// this is required, optional, or not supported.
	PerformanceLocation *string `form:"performance_location" json:"performance_location,omitempty"`
	// A [tax code](https://docs.stripe.com/tax/tax-categories) ID.
	TaxCode *string `form:"tax_code" json:"tax_code,omitempty"`
}

// Data used to generate a new [Product](https://docs.stripe.com/api/products) object inline.
// Provide exactly one of `product` or `product_data`.
type V2PaymentsOffSessionPaymentItemOneTimePriceItemPriceDataProductDataParams struct {
	// The product's description, meant to be displayable to the customer.
	// Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.
	Description *string `form:"description" json:"description,omitempty"`
	// A list of URLs of images for this product, meant to be displayable to the customer.
	Images []*string `form:"images" json:"images,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to the product.
	// This can be useful for storing additional information about the product in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The product's name, meant to be displayable to the customer.
	Name *string `form:"name" json:"name"`
	// Tax details for this product, including the [tax code](https://docs.stripe.com/tax/tax-codes)
	// and an optional performance location.
	TaxDetails *V2PaymentsOffSessionPaymentItemOneTimePriceItemPriceDataProductDataTaxDetailsParams `form:"tax_details" json:"tax_details,omitempty"`
	// A label that represents units of this product, such as seats or licenses.
	UnitLabel *string `form:"unit_label" json:"unit_label,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2PaymentsOffSessionPaymentItemOneTimePriceItemPriceDataProductDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Data used to generate a new one-time [Price](https://docs.stripe.com/api/prices) object inline.
// Provide exactly one of `price` or `price_data`.
type V2PaymentsOffSessionPaymentItemOneTimePriceItemPriceDataParams struct {
	// The ID of the [Product](https://docs.stripe.com/api/products) that this
	// [Price](https://docs.stripe.com/api/prices) will belong to.
	// Provide exactly one of `product` or `product_data`.
	Product *string `form:"product" json:"product,omitempty"`
	// Data used to generate a new [Product](https://docs.stripe.com/api/products) object inline.
	// Provide exactly one of `product` or `product_data`.
	ProductData *V2PaymentsOffSessionPaymentItemOneTimePriceItemPriceDataProductDataParams `form:"product_data" json:"product_data,omitempty"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes.
	// One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior" json:"tax_behavior,omitempty"`
	// The amount to charge per unit and its currency. Specify `value` as a non-negative integer in
	// the currency's minor unit (for example, `100` for USD 1.00).
	UnitAmount *Amount `form:"unit_amount" json:"unit_amount"`
}

// The list of one-time prices and quantities the customer will purchase. Contains 1 to 10 items.
type V2PaymentsOffSessionPaymentItemOneTimePriceItemParams struct {
	// The ID of an existing one-time [Price](https://docs.stripe.com/api/prices).
	// Provide exactly one of `price` or `price_data`.
	Price *string `form:"price" json:"price,omitempty"`
	// Data used to generate a new one-time [Price](https://docs.stripe.com/api/prices) object inline.
	// Provide exactly one of `price` or `price_data`.
	PriceData *V2PaymentsOffSessionPaymentItemOneTimePriceItemPriceDataParams `form:"price_data" json:"price_data,omitempty"`
	// Quantity for this item.
	Quantity *int64 `form:"quantity" json:"quantity"`
}

// Configuration for the one-time price item.
type V2PaymentsOffSessionPaymentItemOneTimePriceParams struct {
	// The list of one-time prices and quantities the customer will purchase. Contains 1 to 10 items.
	Items []*V2PaymentsOffSessionPaymentItemOneTimePriceItemParams `form:"items" json:"items"`
}

// A list of items the customer will purchase.
type V2PaymentsOffSessionPaymentItemParams struct {
	// Configuration for the one-time price item.
	OneTimePrice *V2PaymentsOffSessionPaymentItemOneTimePriceParams `form:"one_time_price" json:"one_time_price"`
	// The type of item. Currently, only `one_time_price` is supported.
	Type *string `form:"type" json:"type"`
}

// Billing information associated with the payment method.
type V2PaymentsOffSessionPaymentPaymentMethodDataBillingDetailsParams struct {
	// Billing address.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// Email address.
	Email *string `form:"email" json:"email,omitempty"`
	// Full name.
	Name *string `form:"name" json:"name,omitempty"`
	// Billing phone number (including extension).
	Phone *string `form:"phone" json:"phone,omitempty"`
}

// The network token credential to use with the card.
type V2PaymentsOffSessionPaymentPaymentMethodDataCardNetworkTokenParams struct {
	// Two-digit number representing the network token's expiration month.
	ExpMonth *string `form:"exp_month" json:"exp_month"`
	// Four-digit number representing the network token's expiration year.
	ExpYear *string `form:"exp_year" json:"exp_year"`
	// The full network token number.
	Number *string `form:"number" json:"number"`
	// The tokenization method used to generate the network token.
	TokenizationMethod *string `form:"tokenization_method" json:"tokenization_method"`
}

// Contains card details that can be used to create a card PaymentMethod for PCI compliant users.
type V2PaymentsOffSessionPaymentPaymentMethodDataCardParams struct {
	// The card expiration month.
	ExpMonth *string `form:"exp_month" json:"exp_month"`
	// The card expiration year.
	ExpYear *string `form:"exp_year" json:"exp_year"`
	// The last four digits of the underlying card number.
	Last4 *string `form:"last4" json:"last4,omitempty"`
	// The network token credential to use with the card.
	NetworkToken *V2PaymentsOffSessionPaymentPaymentMethodDataCardNetworkTokenParams `form:"network_token" json:"network_token,omitempty"`
	// The card number.
	Number *string `form:"number" json:"number,omitempty"`
}

// If provided, this hash will be used to create a PaymentMethod. The new PaymentMethod will appear in the payment_method property on the OffSessionPayment.
type V2PaymentsOffSessionPaymentPaymentMethodDataParams struct {
	// Billing information associated with the payment method.
	BillingDetails *V2PaymentsOffSessionPaymentPaymentMethodDataBillingDetailsParams `form:"billing_details" json:"billing_details,omitempty"`
	// Contains card details that can be used to create a card PaymentMethod for PCI compliant users.
	Card *V2PaymentsOffSessionPaymentPaymentMethodDataCardParams `form:"card" json:"card,omitempty"`
	// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
	Type *string `form:"type" json:"type"`
}

// Payment method options for the card payment type.
type V2PaymentsOffSessionPaymentPaymentMethodOptionsCardParams struct {
	// Indicates whether the saved card is used for a recurring or unscheduled payment.
	StoredCredentialUsage *string `form:"stored_credential_usage" json:"stored_credential_usage"`
}

// Payment method options for the off-session payment.
type V2PaymentsOffSessionPaymentPaymentMethodOptionsParams struct {
	// Payment method options for the card payment type.
	Card *V2PaymentsOffSessionPaymentPaymentMethodOptionsCardParams `form:"card" json:"card,omitempty"`
}

// The amount to collect without itemization.
type V2PaymentsOffSessionPaymentRawAmountParams struct {
	// The amount to collect and its currency.
	Amount *Amount `form:"amount" json:"amount"`
}

// Retry configuration for this OffSessionPayment.
type V2PaymentsOffSessionPaymentRetriesParams struct {
	// Whether Stripe should retry failed payment attempts.
	Enabled *bool `form:"enabled" json:"enabled"`
}

// The data that automatically creates a Transfer after the payment finalizes. Learn more about the use case for [connected accounts](https://docs.stripe.com/payments/connected-accounts).
type V2PaymentsOffSessionPaymentTransferDataParams struct {
	// The amount transferred to the destination account. This transfer will occur
	// automatically after the payment succeeds. If no amount is specified, by default
	// the entire payment amount is transferred to the destination account. The amount
	// must be less than or equal to the
	// [amount_requested](https://docs.stripe.com/api/v2/off-session-payments/object?api-version=2025-05-28.preview#v2_off_session_payment_object-amount_requested),
	// and must be a positive integer representing how much to transfer in the smallest
	// currency unit (e.g., 100 cents to charge $1.00).
	Amount *int64 `form:"amount" json:"amount,omitempty"`
	// The account (if any) that the payment is attributed to for tax reporting, and
	// where funds from the payment are transferred to after payment success.
	Destination *string `form:"destination" json:"destination"`
}

// Creates an OffSessionPayment object.
type V2PaymentsOffSessionPaymentParams struct {
	Params `form:"*"`
	// The amount of the application fee (if any) that will be requested to be applied to the
	// payment and transferred to the application owner's Stripe account.
	ApplicationFeeAmount *Amount `form:"application_fee_amount" json:"application_fee_amount,omitempty"`
	// Settings for automatic tax calculation.
	AutomaticTax *V2PaymentsOffSessionPaymentAutomaticTaxParams `form:"automatic_tax" json:"automatic_tax,omitempty"`
	// Three-letter ISO currency code, in lowercase.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// ID of the Customer to which this OffSessionPayment belongs.
	// Provide exactly one of customer or customer_account.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// ID of the Account representing the customer to which this OffSessionPayment belongs.
	// Provide exactly one of customer or customer_account.
	// Payment methods attached to other Accounts cannot be used with this OffSessionPayment.
	CustomerAccount *string `form:"customer_account" json:"customer_account,omitempty"`
	// An arbitrary string attached to the payment. Set to null to clear the description.
	Description *string `form:"description" json:"description,omitempty"`
	// A list of items the customer will purchase.
	Items []*V2PaymentsOffSessionPaymentItemParams `form:"items" json:"items,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can
	// attach to an object. This can be useful for storing additional information about
	// the object in a structured format. Learn more about
	// [storing information in metadata](https://docs.stripe.com/payments/payment-intents#storing-information-in-metadata).
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
	// The account (if any) for which the funds of the OffSessionPayment are intended.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
	// ID of the payment method used in this OffSessionPayment.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// If provided, this hash will be used to create a PaymentMethod. The new PaymentMethod will appear in the payment_method property on the OffSessionPayment.
	PaymentMethodData *V2PaymentsOffSessionPaymentPaymentMethodDataParams `form:"payment_method_data" json:"payment_method_data,omitempty"`
	// Payment method options for the off-session payment.
	PaymentMethodOptions *V2PaymentsOffSessionPaymentPaymentMethodOptionsParams `form:"payment_method_options" json:"payment_method_options,omitempty"`
	// The amount to collect without itemization.
	RawAmount *V2PaymentsOffSessionPaymentRawAmountParams `form:"raw_amount" json:"raw_amount,omitempty"`
	// Retry configuration for this OffSessionPayment.
	Retries *V2PaymentsOffSessionPaymentRetriesParams `form:"retries" json:"retries,omitempty"`
	// Text that appears on the customer's statement as the statement descriptor for a
	// non-card charge. This value overrides the account's default statement descriptor.
	// For information about requirements, including the 22-character limit, see the
	// [Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// Provides information about a card charge. Concatenated to the account's
	// [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static)
	// to form the complete statement descriptor that appears on the customer's statement.
	StatementDescriptorSuffix *string `form:"statement_descriptor_suffix" json:"statement_descriptor_suffix,omitempty"`
	// Provides a way to schedule deferred payments for cards and debits, in YYYY-MM-DD format.
	TargetDate *string `form:"target_date" json:"target_date,omitempty"`
	// Test clock that can be used to advance the retry attempts in a sandbox.
	TestClock *string `form:"test_clock" json:"test_clock,omitempty"`
	// The data that automatically creates a Transfer after the payment finalizes. Learn more about the use case for [connected accounts](https://docs.stripe.com/payments/connected-accounts).
	TransferData *V2PaymentsOffSessionPaymentTransferDataParams `form:"transfer_data" json:"transfer_data,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2PaymentsOffSessionPaymentParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}

// Cancel an OffSessionPayment that has previously been created.
type V2PaymentsOffSessionPaymentCancelParams struct {
	Params `form:"*"`
}

// Pauses an OffSessionPayment that has previously been created.
type V2PaymentsOffSessionPaymentPauseParams struct {
	Params `form:"*"`
}

// Resumes an OffSessionPayment that has previously been paused.
type V2PaymentsOffSessionPaymentResumeParams struct {
	Params `form:"*"`
}

// Settings for automatic tax calculation.
type V2PaymentsOffSessionPaymentCreateAutomaticTaxParams struct {
	// Whether to automatically calculate taxes for this OffSessionPayment.
	Enabled *bool `form:"enabled" json:"enabled"`
}

// Tax details for this product, including the [tax code](https://docs.stripe.com/tax/tax-codes)
// and an optional performance location.
type V2PaymentsOffSessionPaymentCreateItemOneTimePriceItemPriceDataProductDataTaxDetailsParams struct {
	// A tax location ID. Depending on the [tax code](https://docs.stripe.com/tax/tax-for-tickets/reference/tax-location-performance),
	// this is required, optional, or not supported.
	PerformanceLocation *string `form:"performance_location" json:"performance_location,omitempty"`
	// A [tax code](https://docs.stripe.com/tax/tax-categories) ID.
	TaxCode *string `form:"tax_code" json:"tax_code,omitempty"`
}

// Data used to generate a new [Product](https://docs.stripe.com/api/products) object inline.
// Provide exactly one of `product` or `product_data`.
type V2PaymentsOffSessionPaymentCreateItemOneTimePriceItemPriceDataProductDataParams struct {
	// The product's description, meant to be displayable to the customer.
	// Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.
	Description *string `form:"description" json:"description,omitempty"`
	// A list of URLs of images for this product, meant to be displayable to the customer.
	Images []*string `form:"images" json:"images,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to the product.
	// This can be useful for storing additional information about the product in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The product's name, meant to be displayable to the customer.
	Name *string `form:"name" json:"name"`
	// Tax details for this product, including the [tax code](https://docs.stripe.com/tax/tax-codes)
	// and an optional performance location.
	TaxDetails *V2PaymentsOffSessionPaymentCreateItemOneTimePriceItemPriceDataProductDataTaxDetailsParams `form:"tax_details" json:"tax_details,omitempty"`
	// A label that represents units of this product, such as seats or licenses.
	UnitLabel *string `form:"unit_label" json:"unit_label,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2PaymentsOffSessionPaymentCreateItemOneTimePriceItemPriceDataProductDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Data used to generate a new one-time [Price](https://docs.stripe.com/api/prices) object inline.
// Provide exactly one of `price` or `price_data`.
type V2PaymentsOffSessionPaymentCreateItemOneTimePriceItemPriceDataParams struct {
	// The ID of the [Product](https://docs.stripe.com/api/products) that this
	// [Price](https://docs.stripe.com/api/prices) will belong to.
	// Provide exactly one of `product` or `product_data`.
	Product *string `form:"product" json:"product,omitempty"`
	// Data used to generate a new [Product](https://docs.stripe.com/api/products) object inline.
	// Provide exactly one of `product` or `product_data`.
	ProductData *V2PaymentsOffSessionPaymentCreateItemOneTimePriceItemPriceDataProductDataParams `form:"product_data" json:"product_data,omitempty"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes.
	// One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior" json:"tax_behavior,omitempty"`
	// The amount to charge per unit and its currency. Specify `value` as a non-negative integer in
	// the currency's minor unit (for example, `100` for USD 1.00).
	UnitAmount *Amount `form:"unit_amount" json:"unit_amount"`
}

// The list of one-time prices and quantities the customer will purchase. Contains 1 to 10 items.
type V2PaymentsOffSessionPaymentCreateItemOneTimePriceItemParams struct {
	// The ID of an existing one-time [Price](https://docs.stripe.com/api/prices).
	// Provide exactly one of `price` or `price_data`.
	Price *string `form:"price" json:"price,omitempty"`
	// Data used to generate a new one-time [Price](https://docs.stripe.com/api/prices) object inline.
	// Provide exactly one of `price` or `price_data`.
	PriceData *V2PaymentsOffSessionPaymentCreateItemOneTimePriceItemPriceDataParams `form:"price_data" json:"price_data,omitempty"`
	// Quantity for this item.
	Quantity *int64 `form:"quantity" json:"quantity"`
}

// Configuration for the one-time price item.
type V2PaymentsOffSessionPaymentCreateItemOneTimePriceParams struct {
	// The list of one-time prices and quantities the customer will purchase. Contains 1 to 10 items.
	Items []*V2PaymentsOffSessionPaymentCreateItemOneTimePriceItemParams `form:"items" json:"items"`
}

// A list of items the customer will purchase.
type V2PaymentsOffSessionPaymentCreateItemParams struct {
	// Configuration for the one-time price item.
	OneTimePrice *V2PaymentsOffSessionPaymentCreateItemOneTimePriceParams `form:"one_time_price" json:"one_time_price"`
	// The type of item. Currently, only `one_time_price` is supported.
	Type *string `form:"type" json:"type"`
}

// Billing information associated with the payment method.
type V2PaymentsOffSessionPaymentCreatePaymentMethodDataBillingDetailsParams struct {
	// Billing address.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// Email address.
	Email *string `form:"email" json:"email,omitempty"`
	// Full name.
	Name *string `form:"name" json:"name,omitempty"`
	// Billing phone number (including extension).
	Phone *string `form:"phone" json:"phone,omitempty"`
}

// The network token credential to use with the card.
type V2PaymentsOffSessionPaymentCreatePaymentMethodDataCardNetworkTokenParams struct {
	// Two-digit number representing the network token's expiration month.
	ExpMonth *string `form:"exp_month" json:"exp_month"`
	// Four-digit number representing the network token's expiration year.
	ExpYear *string `form:"exp_year" json:"exp_year"`
	// The full network token number.
	Number *string `form:"number" json:"number"`
	// The tokenization method used to generate the network token.
	TokenizationMethod *string `form:"tokenization_method" json:"tokenization_method"`
}

// Contains card details that can be used to create a card PaymentMethod for PCI compliant users.
type V2PaymentsOffSessionPaymentCreatePaymentMethodDataCardParams struct {
	// The card expiration month.
	ExpMonth *string `form:"exp_month" json:"exp_month"`
	// The card expiration year.
	ExpYear *string `form:"exp_year" json:"exp_year"`
	// The last four digits of the underlying card number.
	Last4 *string `form:"last4" json:"last4,omitempty"`
	// The network token credential to use with the card.
	NetworkToken *V2PaymentsOffSessionPaymentCreatePaymentMethodDataCardNetworkTokenParams `form:"network_token" json:"network_token,omitempty"`
	// The card number.
	Number *string `form:"number" json:"number,omitempty"`
}

// If provided, this hash will be used to create a PaymentMethod. The new PaymentMethod will appear in the payment_method property on the OffSessionPayment.
type V2PaymentsOffSessionPaymentCreatePaymentMethodDataParams struct {
	// Billing information associated with the payment method.
	BillingDetails *V2PaymentsOffSessionPaymentCreatePaymentMethodDataBillingDetailsParams `form:"billing_details" json:"billing_details,omitempty"`
	// Contains card details that can be used to create a card PaymentMethod for PCI compliant users.
	Card *V2PaymentsOffSessionPaymentCreatePaymentMethodDataCardParams `form:"card" json:"card,omitempty"`
	// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
	Type *string `form:"type" json:"type"`
}

// Payment method options for the card payment type.
type V2PaymentsOffSessionPaymentCreatePaymentMethodOptionsCardParams struct {
	// Indicates whether the saved card is used for a recurring or unscheduled payment.
	StoredCredentialUsage *string `form:"stored_credential_usage" json:"stored_credential_usage"`
}

// Payment method options for the off-session payment.
type V2PaymentsOffSessionPaymentCreatePaymentMethodOptionsParams struct {
	// Payment method options for the card payment type.
	Card *V2PaymentsOffSessionPaymentCreatePaymentMethodOptionsCardParams `form:"card" json:"card,omitempty"`
}

// The amount to collect without itemization.
type V2PaymentsOffSessionPaymentCreateRawAmountParams struct {
	// The amount to collect and its currency.
	Amount *Amount `form:"amount" json:"amount"`
}

// Retry configuration for this OffSessionPayment.
type V2PaymentsOffSessionPaymentCreateRetriesParams struct {
	// Whether Stripe should retry failed payment attempts.
	Enabled *bool `form:"enabled" json:"enabled"`
}

// The data that automatically creates a Transfer after the payment finalizes. Learn more about the use case for [connected accounts](https://docs.stripe.com/payments/connected-accounts).
type V2PaymentsOffSessionPaymentCreateTransferDataParams struct {
	// The amount transferred to the destination account. This transfer will occur
	// automatically after the payment succeeds. If no amount is specified, by default
	// the entire payment amount is transferred to the destination account. The amount
	// must be less than or equal to the
	// [amount_requested](https://docs.stripe.com/api/v2/off-session-payments/object?api-version=2025-05-28.preview#v2_off_session_payment_object-amount_requested),
	// and must be a positive integer representing how much to transfer in the smallest
	// currency unit (e.g., 100 cents to charge $1.00).
	Amount *int64 `form:"amount" json:"amount,omitempty"`
	// The account (if any) that the payment is attributed to for tax reporting, and
	// where funds from the payment are transferred to after payment success.
	Destination *string `form:"destination" json:"destination"`
}

// Creates an OffSessionPayment object.
type V2PaymentsOffSessionPaymentCreateParams struct {
	Params `form:"*"`
	// The amount of the application fee (if any) that will be requested to be applied to the
	// payment and transferred to the application owner's Stripe account.
	ApplicationFeeAmount *Amount `form:"application_fee_amount" json:"application_fee_amount,omitempty"`
	// Settings for automatic tax calculation.
	AutomaticTax *V2PaymentsOffSessionPaymentCreateAutomaticTaxParams `form:"automatic_tax" json:"automatic_tax,omitempty"`
	// Three-letter ISO currency code, in lowercase.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// ID of the Customer to which this OffSessionPayment belongs.
	// Provide exactly one of customer or customer_account.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// ID of the Account representing the customer to which this OffSessionPayment belongs.
	// Provide exactly one of customer or customer_account.
	// Payment methods attached to other Accounts cannot be used with this OffSessionPayment.
	CustomerAccount *string `form:"customer_account" json:"customer_account,omitempty"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description" json:"description,omitempty"`
	// A list of items the customer will purchase.
	Items []*V2PaymentsOffSessionPaymentCreateItemParams `form:"items" json:"items,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can
	// attach to an object. This can be useful for storing additional information about
	// the object in a structured format. Learn more about
	// [storing information in metadata](https://docs.stripe.com/payments/payment-intents#storing-information-in-metadata).
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The account (if any) for which the funds of the OffSessionPayment are intended.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
	// ID of the payment method used in this OffSessionPayment.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// If provided, this hash will be used to create a PaymentMethod. The new PaymentMethod will appear in the payment_method property on the OffSessionPayment.
	PaymentMethodData *V2PaymentsOffSessionPaymentCreatePaymentMethodDataParams `form:"payment_method_data" json:"payment_method_data,omitempty"`
	// Payment method options for the off-session payment.
	PaymentMethodOptions *V2PaymentsOffSessionPaymentCreatePaymentMethodOptionsParams `form:"payment_method_options" json:"payment_method_options,omitempty"`
	// The amount to collect without itemization.
	RawAmount *V2PaymentsOffSessionPaymentCreateRawAmountParams `form:"raw_amount" json:"raw_amount,omitempty"`
	// Retry configuration for this OffSessionPayment.
	Retries *V2PaymentsOffSessionPaymentCreateRetriesParams `form:"retries" json:"retries"`
	// Text that appears on the customer's statement as the statement descriptor for a
	// non-card charge. This value overrides the account's default statement descriptor.
	// For information about requirements, including the 22-character limit, see the
	// [Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// Provides information about a card charge. Concatenated to the account's
	// [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static)
	// to form the complete statement descriptor that appears on the customer's statement.
	StatementDescriptorSuffix *string `form:"statement_descriptor_suffix" json:"statement_descriptor_suffix,omitempty"`
	// Provides a way to schedule deferred payments for cards and debits, in YYYY-MM-DD format.
	TargetDate *string `form:"target_date" json:"target_date,omitempty"`
	// Test clock that can be used to advance the retry attempts in a sandbox.
	TestClock *string `form:"test_clock" json:"test_clock,omitempty"`
	// The data that automatically creates a Transfer after the payment finalizes. Learn more about the use case for [connected accounts](https://docs.stripe.com/payments/connected-accounts).
	TransferData *V2PaymentsOffSessionPaymentCreateTransferDataParams `form:"transfer_data" json:"transfer_data,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2PaymentsOffSessionPaymentCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves the details of an OffSessionPayment that has previously been created.
type V2PaymentsOffSessionPaymentRetrieveParams struct {
	Params `form:"*"`
}

// Updates the metadata and description of an OffSessionPayment in any status.
type V2PaymentsOffSessionPaymentUpdateParams struct {
	Params `form:"*"`
	// An arbitrary string attached to the payment. Set to null to clear the description.
	Description *string `form:"description" json:"description,omitempty"`
	// Set of key-value pairs attached to the payment. Existing keys are preserved unless
	// explicitly changed. Set an individual key to null to remove it.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2PaymentsOffSessionPaymentUpdateParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}
