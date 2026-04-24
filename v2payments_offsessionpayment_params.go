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

// Contains information about the tax on the item.
type V2PaymentsOffSessionPaymentAmountDetailsLineItemTaxParams struct {
	// Total portion of the amount that is for tax.
	TotalTaxAmount *int64 `form:"total_tax_amount" json:"total_tax_amount"`
}

// A list of line items, each containing information about a product in the OffSessionPayment. There is a maximum of 10 line items.
type V2PaymentsOffSessionPaymentAmountDetailsLineItemParams struct {
	// The amount an item was discounted for. Positive integer.
	DiscountAmount *int64 `form:"discount_amount" json:"discount_amount,omitempty"`
	// Unique identifier of the product. At most 12 characters long.
	ProductCode *string `form:"product_code" json:"product_code,omitempty"`
	// Name of the product. At most 100 characters long.
	ProductName *string `form:"product_name" json:"product_name"`
	// Number of items of the product. Positive integer.
	Quantity *int64 `form:"quantity" json:"quantity"`
	// Contains information about the tax on the item.
	Tax *V2PaymentsOffSessionPaymentAmountDetailsLineItemTaxParams `form:"tax" json:"tax,omitempty"`
	// Cost of the product. Positive integer.
	UnitCost *int64 `form:"unit_cost" json:"unit_cost"`
	// A unit of measure for the line item, such as gallons, feet, meters, etc.
	// The maximum length is 12 characters.
	UnitOfMeasure *string `form:"unit_of_measure" json:"unit_of_measure,omitempty"`
}

// Contains information about the shipping portion of the amount.
type V2PaymentsOffSessionPaymentAmountDetailsShippingParams struct {
	// Portion of the amount that is for shipping.
	Amount *int64 `form:"amount" json:"amount,omitempty"`
	// The postal code that represents the shipping source.
	FromPostalCode *string `form:"from_postal_code" json:"from_postal_code,omitempty"`
	// The postal code that represents the shipping destination.
	ToPostalCode *string `form:"to_postal_code" json:"to_postal_code,omitempty"`
}

// Contains information about the tax portion of the amount.
type V2PaymentsOffSessionPaymentAmountDetailsTaxParams struct {
	// Total portion of the amount that is for tax.
	TotalTaxAmount *int64 `form:"total_tax_amount" json:"total_tax_amount"`
}

// Provides industry-specific information about the amount.
type V2PaymentsOffSessionPaymentAmountDetailsParams struct {
	// The amount the total transaction was discounted for.
	DiscountAmount *int64 `form:"discount_amount" json:"discount_amount,omitempty"`
	// Set to `false` to return arithmetic validation errors in the response without failing the request. Use this when you want the operation to proceed regardless of arithmetic errors in the line item data.
	// Omit or set to `true` to immediately return a 400 error when arithmetic validation fails. Use this for strict validation that prevents processing with line item data that has arithmetic inconsistencies.
	// For card payments, Stripe doesn't send line item data to card networks if there's an arithmetic validation error.
	EnforceArithmeticValidation *bool `form:"enforce_arithmetic_validation" json:"enforce_arithmetic_validation,omitempty"`
	// A list of line items, each containing information about a product in the OffSessionPayment. There is a maximum of 10 line items.
	LineItems []*V2PaymentsOffSessionPaymentAmountDetailsLineItemParams `form:"line_items" json:"line_items,omitempty"`
	// Contains information about the shipping portion of the amount.
	Shipping *V2PaymentsOffSessionPaymentAmountDetailsShippingParams `form:"shipping" json:"shipping,omitempty"`
	// Contains information about the tax portion of the amount.
	Tax *V2PaymentsOffSessionPaymentAmountDetailsTaxParams `form:"tax" json:"tax,omitempty"`
}

// Details about the capture configuration for the OffSessionPayment.
type V2PaymentsOffSessionPaymentCaptureParams struct {
	Params `form:"*"`
	// Provides industry-specific information about the amount.
	AmountDetails *V2PaymentsOffSessionPaymentCaptureAmountDetailsParams `form:"amount_details" json:"amount_details,omitempty"`
	// The amount to capture.
	AmountToCapture *int64 `form:"amount_to_capture" json:"amount_to_capture,omitempty"`
	// The amount of the application fee for this capture.
	ApplicationFeeAmount *Amount `form:"application_fee_amount" json:"application_fee_amount,omitempty"`
	// The method to use to capture the payment.
	CaptureMethod *string `form:"capture_method" json:"capture_method,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can
	// attach to an object. This can be useful for storing additional information about
	// the object in a structured format. Learn more about
	// [storing information in metadata](https://docs.stripe.com/payments/payment-intents#storing-information-in-metadata).
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Provides industry-specific information about the payment.
	PaymentDetails *V2PaymentsOffSessionPaymentCapturePaymentDetailsParams `form:"payment_details" json:"payment_details,omitempty"`
	// Text that appears on the customer's statement as the statement descriptor for a
	// non-card charge. This value overrides the account's default statement descriptor.
	// For information about requirements, including the 22-character limit, see the
	// [Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// Provides information about a card charge. Concatenated to the account's
	// [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static)
	// to form the complete statement descriptor that appears on the customer's statement.
	StatementDescriptorSuffix *string `form:"statement_descriptor_suffix" json:"statement_descriptor_suffix,omitempty"`
	// The data that automatically creates a Transfer after the payment finalizes. Learn more about the use case for [connected accounts](https://docs.stripe.com/payments/connected-accounts).
	TransferData *V2PaymentsOffSessionPaymentCaptureTransferDataParams `form:"transfer_data" json:"transfer_data,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2PaymentsOffSessionPaymentCaptureParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Provides industry-specific information about the payment.
type V2PaymentsOffSessionPaymentPaymentDetailsParams struct {
	// A unique value to identify the customer. This field is applicable only for card payments. For card payments, this field is truncated to 25 alphanumeric characters, excluding spaces, before being sent to card networks.
	CustomerReference *string `form:"customer_reference" json:"customer_reference,omitempty"`
	// A unique value assigned by the business to identify the transaction. Required for L2 and L3 rates.
	// For Cards, this field is truncated to 25 alphanumeric characters, excluding spaces, before being sent to card networks.
	OrderReference *string `form:"order_reference" json:"order_reference,omitempty"`
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

// Contains card details that can be used to create a card PaymentMethod for PCI compliant users.
type V2PaymentsOffSessionPaymentPaymentMethodDataCardParams struct {
	// The card CVC.
	CVC *string `form:"cvc" json:"cvc,omitempty"`
	// The card expiration month.
	ExpMonth *string `form:"exp_month" json:"exp_month"`
	// The card expiration year.
	ExpYear *string `form:"exp_year" json:"exp_year"`
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
	// The merchant category code for this transaction. Used in interchange and authorization to improve auth rates.
	MCC *string `form:"mcc" json:"mcc,omitempty"`
	// If you are making a Credential On File transaction with a previously saved card, you should pass the Network Transaction ID
	// from a prior initial authorization on Stripe (from a successful SetupIntent or a PaymentIntent with `setup_future_usage` set),
	// or one that you have obtained from another payment processor. This is a token from the network which uniquely identifies the transaction.
	// Visa calls this the Transaction ID, Mastercard calls this the Trace ID, and American Express calls this the Acquirer Reference Data.
	// Note that you should pass in a Network Transaction ID if you have one, regardless of whether this is a
	// Customer-Initiated Transaction (CIT) or a Merchant-Initiated Transaction (MIT).
	NetworkTransactionID *string `form:"network_transaction_id" json:"network_transaction_id,omitempty"`
}

// Payment method options for the off-session payment.
type V2PaymentsOffSessionPaymentPaymentMethodOptionsParams struct {
	// Payment method options for the card payment type.
	Card *V2PaymentsOffSessionPaymentPaymentMethodOptionsCardParams `form:"card" json:"card,omitempty"`
}

// Details about the payments orchestration configuration.
type V2PaymentsOffSessionPaymentPaymentsOrchestrationParams struct {
	// True when you want to enable payments orchestration for this off-session payment. False otherwise.
	Enabled *bool `form:"enabled" json:"enabled"`
}

// Details about the OffSessionPayment retries.
type V2PaymentsOffSessionPaymentRetryDetailsParams struct {
	// The pre-configured retry policy to use for the payment.
	RetryPolicy *string `form:"retry_policy" json:"retry_policy,omitempty"`
	// Indicates the strategy for how you want Stripe to retry the payment.
	RetryStrategy *string `form:"retry_strategy" json:"retry_strategy,omitempty"`
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
	// The “presentment amount” to be collected from the customer.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
	// Provides industry-specific information about the amount.
	AmountDetails *V2PaymentsOffSessionPaymentAmountDetailsParams `form:"amount_details" json:"amount_details,omitempty"`
	// The amount of the application fee (if any) that will be requested to be applied to the
	// payment and transferred to the application owner's Stripe account.
	ApplicationFeeAmount *Amount `form:"application_fee_amount" json:"application_fee_amount,omitempty"`
	// The frequency of the underlying payment.
	Cadence *string `form:"cadence" json:"cadence,omitempty"`
	// Details about the capture configuration for the OffSessionPayment.
	Capture *V2PaymentsOffSessionPaymentCaptureParams `form:"capture" json:"capture,omitempty"`
	// ID of the Customer to which this OffSessionPayment belongs.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description" json:"description,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can
	// attach to an object. This can be useful for storing additional information about
	// the object in a structured format. Learn more about
	// [storing information in metadata](https://docs.stripe.com/payments/payment-intents#storing-information-in-metadata).
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The account (if any) for which the funds of the OffSessionPayment are intended.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
	// Provides industry-specific information about the payment.
	PaymentDetails *V2PaymentsOffSessionPaymentPaymentDetailsParams `form:"payment_details" json:"payment_details,omitempty"`
	// ID of the payment method used in this OffSessionPayment.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// If provided, this hash will be used to create a PaymentMethod. The new PaymentMethod will appear in the payment_method property on the OffSessionPayment.
	PaymentMethodData *V2PaymentsOffSessionPaymentPaymentMethodDataParams `form:"payment_method_data" json:"payment_method_data,omitempty"`
	// Payment method options for the off-session payment.
	PaymentMethodOptions *V2PaymentsOffSessionPaymentPaymentMethodOptionsParams `form:"payment_method_options" json:"payment_method_options,omitempty"`
	// Details about the payments orchestration configuration.
	PaymentsOrchestration *V2PaymentsOffSessionPaymentPaymentsOrchestrationParams `form:"payments_orchestration" json:"payments_orchestration,omitempty"`
	// Details about the OffSessionPayment retries.
	RetryDetails *V2PaymentsOffSessionPaymentRetryDetailsParams `form:"retry_details" json:"retry_details,omitempty"`
	// Text that appears on the customer's statement as the statement descriptor for a
	// non-card charge. This value overrides the account's default statement descriptor.
	// For information about requirements, including the 22-character limit, see the
	// [Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// Provides information about a card charge. Concatenated to the account's
	// [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static)
	// to form the complete statement descriptor that appears on the customer's statement.
	StatementDescriptorSuffix *string `form:"statement_descriptor_suffix" json:"statement_descriptor_suffix,omitempty"`
	// Test clock that can be used to advance the retry attempts in a sandbox.
	TestClock *string `form:"test_clock" json:"test_clock,omitempty"`
	// The data that automatically creates a Transfer after the payment finalizes. Learn more about the use case for [connected accounts](https://docs.stripe.com/payments/connected-accounts).
	TransferData *V2PaymentsOffSessionPaymentTransferDataParams `form:"transfer_data" json:"transfer_data,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2PaymentsOffSessionPaymentParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Cancel an OffSessionPayment that has previously been created.
type V2PaymentsOffSessionPaymentCancelParams struct {
	Params `form:"*"`
}

// Contains information about the tax on the item.
type V2PaymentsOffSessionPaymentCaptureAmountDetailsLineItemTaxParams struct {
	// Total portion of the amount that is for tax.
	TotalTaxAmount *int64 `form:"total_tax_amount" json:"total_tax_amount"`
}

// A list of line items, each containing information about a product in the OffSessionPayment. There is a maximum of 10 line items.
type V2PaymentsOffSessionPaymentCaptureAmountDetailsLineItemParams struct {
	// The amount an item was discounted for. Positive integer.
	DiscountAmount *int64 `form:"discount_amount" json:"discount_amount,omitempty"`
	// Unique identifier of the product. At most 12 characters long.
	ProductCode *string `form:"product_code" json:"product_code,omitempty"`
	// Name of the product. At most 100 characters long.
	ProductName *string `form:"product_name" json:"product_name"`
	// Number of items of the product. Positive integer.
	Quantity *int64 `form:"quantity" json:"quantity"`
	// Contains information about the tax on the item.
	Tax *V2PaymentsOffSessionPaymentCaptureAmountDetailsLineItemTaxParams `form:"tax" json:"tax,omitempty"`
	// Cost of the product. Positive integer.
	UnitCost *int64 `form:"unit_cost" json:"unit_cost"`
	// A unit of measure for the line item, such as gallons, feet, meters, etc.
	// The maximum length is 12 characters.
	UnitOfMeasure *string `form:"unit_of_measure" json:"unit_of_measure,omitempty"`
}

// Contains information about the shipping portion of the amount.
type V2PaymentsOffSessionPaymentCaptureAmountDetailsShippingParams struct {
	// Portion of the amount that is for shipping.
	Amount *int64 `form:"amount" json:"amount,omitempty"`
	// The postal code that represents the shipping source.
	FromPostalCode *string `form:"from_postal_code" json:"from_postal_code,omitempty"`
	// The postal code that represents the shipping destination.
	ToPostalCode *string `form:"to_postal_code" json:"to_postal_code,omitempty"`
}

// Contains information about the tax portion of the amount.
type V2PaymentsOffSessionPaymentCaptureAmountDetailsTaxParams struct {
	// Total portion of the amount that is for tax.
	TotalTaxAmount *int64 `form:"total_tax_amount" json:"total_tax_amount"`
}

// Provides industry-specific information about the amount.
type V2PaymentsOffSessionPaymentCaptureAmountDetailsParams struct {
	// The amount the total transaction was discounted for.
	DiscountAmount *int64 `form:"discount_amount" json:"discount_amount,omitempty"`
	// Set to `false` to return arithmetic validation errors in the response without failing the request. Use this when you want the operation to proceed regardless of arithmetic errors in the line item data.
	// Omit or set to `true` to immediately return a 400 error when arithmetic validation fails. Use this for strict validation that prevents processing with line item data that has arithmetic inconsistencies.
	// For card payments, Stripe doesn't send line item data to card networks if there's an arithmetic validation error.
	EnforceArithmeticValidation *bool `form:"enforce_arithmetic_validation" json:"enforce_arithmetic_validation,omitempty"`
	// A list of line items, each containing information about a product in the OffSessionPayment. There is a maximum of 10 line items.
	LineItems []*V2PaymentsOffSessionPaymentCaptureAmountDetailsLineItemParams `form:"line_items" json:"line_items,omitempty"`
	// Contains information about the shipping portion of the amount.
	Shipping *V2PaymentsOffSessionPaymentCaptureAmountDetailsShippingParams `form:"shipping" json:"shipping,omitempty"`
	// Contains information about the tax portion of the amount.
	Tax *V2PaymentsOffSessionPaymentCaptureAmountDetailsTaxParams `form:"tax" json:"tax,omitempty"`
}

// Provides industry-specific information about the payment.
type V2PaymentsOffSessionPaymentCapturePaymentDetailsParams struct {
	// A unique value to identify the customer. This field is applicable only for card payments. For card payments, this field is truncated to 25 alphanumeric characters, excluding spaces, before being sent to card networks.
	CustomerReference *string `form:"customer_reference" json:"customer_reference,omitempty"`
	// A unique value assigned by the business to identify the transaction. Required for L2 and L3 rates.
	// For Cards, this field is truncated to 25 alphanumeric characters, excluding spaces, before being sent to card networks.
	OrderReference *string `form:"order_reference" json:"order_reference,omitempty"`
}

// The data that automatically creates a Transfer after the payment finalizes. Learn more about the use case for [connected accounts](https://docs.stripe.com/payments/connected-accounts).
type V2PaymentsOffSessionPaymentCaptureTransferDataParams struct {
	// The amount transferred to the destination account. This transfer will occur
	// automatically after the payment succeeds. If no amount is specified, by default
	// the entire payment amount is transferred to the destination account. The amount
	// must be less than or equal to the
	// [amount_requested](https://docs.stripe.com/api/v2/off-session-payments/object?api-version=2025-05-28.preview#v2_off_session_payment_object-amount_requested),
	// and must be a positive integer representing how much to transfer in the smallest
	// currency unit (e.g., 100 cents to charge $1.00).
	Amount *int64 `form:"amount" json:"amount,omitempty"`
}

// Pauses an OffSessionPayment that has previously been created.
type V2PaymentsOffSessionPaymentPauseParams struct {
	Params `form:"*"`
}

// Resumes an OffSessionPayment that has previously been paused.
type V2PaymentsOffSessionPaymentResumeParams struct {
	Params `form:"*"`
}

// Contains information about the tax on the item.
type V2PaymentsOffSessionPaymentCreateAmountDetailsLineItemTaxParams struct {
	// Total portion of the amount that is for tax.
	TotalTaxAmount *int64 `form:"total_tax_amount" json:"total_tax_amount"`
}

// A list of line items, each containing information about a product in the OffSessionPayment. There is a maximum of 10 line items.
type V2PaymentsOffSessionPaymentCreateAmountDetailsLineItemParams struct {
	// The amount an item was discounted for. Positive integer.
	DiscountAmount *int64 `form:"discount_amount" json:"discount_amount,omitempty"`
	// Unique identifier of the product. At most 12 characters long.
	ProductCode *string `form:"product_code" json:"product_code,omitempty"`
	// Name of the product. At most 100 characters long.
	ProductName *string `form:"product_name" json:"product_name"`
	// Number of items of the product. Positive integer.
	Quantity *int64 `form:"quantity" json:"quantity"`
	// Contains information about the tax on the item.
	Tax *V2PaymentsOffSessionPaymentCreateAmountDetailsLineItemTaxParams `form:"tax" json:"tax,omitempty"`
	// Cost of the product. Positive integer.
	UnitCost *int64 `form:"unit_cost" json:"unit_cost"`
	// A unit of measure for the line item, such as gallons, feet, meters, etc.
	// The maximum length is 12 characters.
	UnitOfMeasure *string `form:"unit_of_measure" json:"unit_of_measure,omitempty"`
}

// Contains information about the shipping portion of the amount.
type V2PaymentsOffSessionPaymentCreateAmountDetailsShippingParams struct {
	// Portion of the amount that is for shipping.
	Amount *int64 `form:"amount" json:"amount,omitempty"`
	// The postal code that represents the shipping source.
	FromPostalCode *string `form:"from_postal_code" json:"from_postal_code,omitempty"`
	// The postal code that represents the shipping destination.
	ToPostalCode *string `form:"to_postal_code" json:"to_postal_code,omitempty"`
}

// Contains information about the tax portion of the amount.
type V2PaymentsOffSessionPaymentCreateAmountDetailsTaxParams struct {
	// Total portion of the amount that is for tax.
	TotalTaxAmount *int64 `form:"total_tax_amount" json:"total_tax_amount"`
}

// Provides industry-specific information about the amount.
type V2PaymentsOffSessionPaymentCreateAmountDetailsParams struct {
	// The amount the total transaction was discounted for.
	DiscountAmount *int64 `form:"discount_amount" json:"discount_amount,omitempty"`
	// Set to `false` to return arithmetic validation errors in the response without failing the request. Use this when you want the operation to proceed regardless of arithmetic errors in the line item data.
	// Omit or set to `true` to immediately return a 400 error when arithmetic validation fails. Use this for strict validation that prevents processing with line item data that has arithmetic inconsistencies.
	// For card payments, Stripe doesn't send line item data to card networks if there's an arithmetic validation error.
	EnforceArithmeticValidation *bool `form:"enforce_arithmetic_validation" json:"enforce_arithmetic_validation,omitempty"`
	// A list of line items, each containing information about a product in the OffSessionPayment. There is a maximum of 10 line items.
	LineItems []*V2PaymentsOffSessionPaymentCreateAmountDetailsLineItemParams `form:"line_items" json:"line_items,omitempty"`
	// Contains information about the shipping portion of the amount.
	Shipping *V2PaymentsOffSessionPaymentCreateAmountDetailsShippingParams `form:"shipping" json:"shipping,omitempty"`
	// Contains information about the tax portion of the amount.
	Tax *V2PaymentsOffSessionPaymentCreateAmountDetailsTaxParams `form:"tax" json:"tax,omitempty"`
}

// Details about the capture configuration for the OffSessionPayment.
type V2PaymentsOffSessionPaymentCreateCaptureParams struct {
	// The method to use to capture the payment.
	CaptureMethod *string `form:"capture_method" json:"capture_method"`
}

// Provides industry-specific information about the payment.
type V2PaymentsOffSessionPaymentCreatePaymentDetailsParams struct {
	// A unique value to identify the customer. This field is applicable only for card payments. For card payments, this field is truncated to 25 alphanumeric characters, excluding spaces, before being sent to card networks.
	CustomerReference *string `form:"customer_reference" json:"customer_reference,omitempty"`
	// A unique value assigned by the business to identify the transaction. Required for L2 and L3 rates.
	// For Cards, this field is truncated to 25 alphanumeric characters, excluding spaces, before being sent to card networks.
	OrderReference *string `form:"order_reference" json:"order_reference,omitempty"`
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

// Contains card details that can be used to create a card PaymentMethod for PCI compliant users.
type V2PaymentsOffSessionPaymentCreatePaymentMethodDataCardParams struct {
	// The card CVC.
	CVC *string `form:"cvc" json:"cvc,omitempty"`
	// The card expiration month.
	ExpMonth *string `form:"exp_month" json:"exp_month"`
	// The card expiration year.
	ExpYear *string `form:"exp_year" json:"exp_year"`
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
	// The merchant category code for this transaction. Used in interchange and authorization to improve auth rates.
	MCC *string `form:"mcc" json:"mcc,omitempty"`
	// If you are making a Credential On File transaction with a previously saved card, you should pass the Network Transaction ID
	// from a prior initial authorization on Stripe (from a successful SetupIntent or a PaymentIntent with `setup_future_usage` set),
	// or one that you have obtained from another payment processor. This is a token from the network which uniquely identifies the transaction.
	// Visa calls this the Transaction ID, Mastercard calls this the Trace ID, and American Express calls this the Acquirer Reference Data.
	// Note that you should pass in a Network Transaction ID if you have one, regardless of whether this is a
	// Customer-Initiated Transaction (CIT) or a Merchant-Initiated Transaction (MIT).
	NetworkTransactionID *string `form:"network_transaction_id" json:"network_transaction_id,omitempty"`
}

// Payment method options for the off-session payment.
type V2PaymentsOffSessionPaymentCreatePaymentMethodOptionsParams struct {
	// Payment method options for the card payment type.
	Card *V2PaymentsOffSessionPaymentCreatePaymentMethodOptionsCardParams `form:"card" json:"card,omitempty"`
}

// Details about the payments orchestration configuration.
type V2PaymentsOffSessionPaymentCreatePaymentsOrchestrationParams struct {
	// True when you want to enable payments orchestration for this off-session payment. False otherwise.
	Enabled *bool `form:"enabled" json:"enabled"`
}

// Details about the OffSessionPayment retries.
type V2PaymentsOffSessionPaymentCreateRetryDetailsParams struct {
	// The pre-configured retry policy to use for the payment.
	RetryPolicy *string `form:"retry_policy" json:"retry_policy,omitempty"`
	// Indicates the strategy for how you want Stripe to retry the payment.
	RetryStrategy *string `form:"retry_strategy" json:"retry_strategy,omitempty"`
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
	// The “presentment amount” to be collected from the customer.
	Amount *Amount `form:"amount" json:"amount"`
	// Provides industry-specific information about the amount.
	AmountDetails *V2PaymentsOffSessionPaymentCreateAmountDetailsParams `form:"amount_details" json:"amount_details,omitempty"`
	// The amount of the application fee (if any) that will be requested to be applied to the
	// payment and transferred to the application owner's Stripe account.
	ApplicationFeeAmount *Amount `form:"application_fee_amount" json:"application_fee_amount,omitempty"`
	// The frequency of the underlying payment.
	Cadence *string `form:"cadence" json:"cadence"`
	// Details about the capture configuration for the OffSessionPayment.
	Capture *V2PaymentsOffSessionPaymentCreateCaptureParams `form:"capture" json:"capture,omitempty"`
	// ID of the Customer to which this OffSessionPayment belongs.
	Customer *string `form:"customer" json:"customer"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description" json:"description,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can
	// attach to an object. This can be useful for storing additional information about
	// the object in a structured format. Learn more about
	// [storing information in metadata](https://docs.stripe.com/payments/payment-intents#storing-information-in-metadata).
	Metadata map[string]string `form:"metadata" json:"metadata"`
	// The account (if any) for which the funds of the OffSessionPayment are intended.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
	// Provides industry-specific information about the payment.
	PaymentDetails *V2PaymentsOffSessionPaymentCreatePaymentDetailsParams `form:"payment_details" json:"payment_details,omitempty"`
	// ID of the payment method used in this OffSessionPayment.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// If provided, this hash will be used to create a PaymentMethod. The new PaymentMethod will appear in the payment_method property on the OffSessionPayment.
	PaymentMethodData *V2PaymentsOffSessionPaymentCreatePaymentMethodDataParams `form:"payment_method_data" json:"payment_method_data,omitempty"`
	// Payment method options for the off-session payment.
	PaymentMethodOptions *V2PaymentsOffSessionPaymentCreatePaymentMethodOptionsParams `form:"payment_method_options" json:"payment_method_options,omitempty"`
	// Details about the payments orchestration configuration.
	PaymentsOrchestration *V2PaymentsOffSessionPaymentCreatePaymentsOrchestrationParams `form:"payments_orchestration" json:"payments_orchestration,omitempty"`
	// Details about the OffSessionPayment retries.
	RetryDetails *V2PaymentsOffSessionPaymentCreateRetryDetailsParams `form:"retry_details" json:"retry_details,omitempty"`
	// Text that appears on the customer's statement as the statement descriptor for a
	// non-card charge. This value overrides the account's default statement descriptor.
	// For information about requirements, including the 22-character limit, see the
	// [Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// Provides information about a card charge. Concatenated to the account's
	// [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static)
	// to form the complete statement descriptor that appears on the customer's statement.
	StatementDescriptorSuffix *string `form:"statement_descriptor_suffix" json:"statement_descriptor_suffix,omitempty"`
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
