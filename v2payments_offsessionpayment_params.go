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

// Details about the capture configuration for the OffSessionPayment.
type V2PaymentsOffSessionPaymentCaptureParams struct {
	Params `form:"*"`
	// The amount to capture.
	AmountToCapture *int64 `form:"amount_to_capture" json:"amount_to_capture,omitempty"`
	// The method to use to capture the payment.
	CaptureMethod *string `form:"capture_method" json:"capture_method,omitempty"`
	// Set of [key-value pairs](https://docs.corp.stripe.com/api/metadata) that you can
	// attach to an object. This can be useful for storing additional information about
	// the object in a structured format. Learn more about
	// [storing information in metadata](https://docs.corp.stripe.com/payments/payment-intents#storing-information-in-metadata).
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Text that appears on the customer's statement as the statement descriptor for a
	// non-card charge. This value overrides the account's default statement descriptor.
	// For information about requirements, including the 22-character limit, see the
	// [Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// Provides information about a card charge. Concatenated to the account's
	// [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static)
	// to form the complete statement descriptor that appears on the customer's statement.
	StatementDescriptorSuffix *string `form:"statement_descriptor_suffix" json:"statement_descriptor_suffix,omitempty"`
	// The data that automatically creates a Transfer after the payment finalizes. Learn more about the use case for [connected accounts](https://docs.corp.stripe.com/payments/connected-accounts).
	TransferData *V2PaymentsOffSessionPaymentCaptureTransferDataParams `form:"transfer_data" json:"transfer_data,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2PaymentsOffSessionPaymentCaptureParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Payment method options for the card payment type.
type V2PaymentsOffSessionPaymentPaymentMethodOptionsCardParams struct {
	// If you are making a Credential On File transaction with a previously saved card, you should pass the Network Transaction ID
	// from a prior initial authorization on Stripe (from a successful SetupIntent or a PaymentIntent with `setup_future_usage` set),
	// or one that you have obtained from another payment processor. This is a token from the network which uniquely identifies the transaction.
	// Visa calls this the Transaction ID, Mastercard calls this the Trace ID, and American Express calls this the Acquirer Reference Data.
	// Note that you should pass in a Network Transaction ID if you have one, regardless of whether this is a
	// Customer-Initiated Transaction (CIT) or a Merchant-Initiated Transaction (MIT).
	NetworkTransactionID *string `form:"network_transaction_id" json:"network_transaction_id"`
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

// The data that automatically creates a Transfer after the payment finalizes. Learn more about the use case for [connected accounts](https://docs.corp.stripe.com/payments/connected-accounts).
type V2PaymentsOffSessionPaymentTransferDataParams struct {
	// The amount transferred to the destination account. This transfer will occur
	// automatically after the payment succeeds. If no amount is specified, by default
	// the entire payment amount is transferred to the destination account. The amount
	// must be less than or equal to the
	// [amount_requested](https://docs.corp.stripe.com/api/v2/off-session-payments/object?api-version=2025-05-28.preview#v2_off_session_payment_object-amount_requested),
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
	// The frequency of the underlying payment.
	Cadence *string `form:"cadence" json:"cadence,omitempty"`
	// Details about the capture configuration for the OffSessionPayment.
	Capture *V2PaymentsOffSessionPaymentCaptureParams `form:"capture" json:"capture,omitempty"`
	// ID of the Customer to which this OffSessionPayment belongs.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// Set of [key-value pairs](https://docs.corp.stripe.com/api/metadata) that you can
	// attach to an object. This can be useful for storing additional information about
	// the object in a structured format. Learn more about
	// [storing information in metadata](https://docs.corp.stripe.com/payments/payment-intents#storing-information-in-metadata).
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The account (if any) for which the funds of the OffSessionPayment are intended.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
	// ID of the payment method used in this OffSessionPayment.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
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
	// The data that automatically creates a Transfer after the payment finalizes. Learn more about the use case for [connected accounts](https://docs.corp.stripe.com/payments/connected-accounts).
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

// The data that automatically creates a Transfer after the payment finalizes. Learn more about the use case for [connected accounts](https://docs.corp.stripe.com/payments/connected-accounts).
type V2PaymentsOffSessionPaymentCaptureTransferDataParams struct {
	// The amount transferred to the destination account. This transfer will occur
	// automatically after the payment succeeds. If no amount is specified, by default
	// the entire payment amount is transferred to the destination account. The amount
	// must be less than or equal to the
	// [amount_requested](https://docs.corp.stripe.com/api/v2/off-session-payments/object?api-version=2025-05-28.preview#v2_off_session_payment_object-amount_requested),
	// and must be a positive integer representing how much to transfer in the smallest
	// currency unit (e.g., 100 cents to charge $1.00).
	Amount *int64 `form:"amount" json:"amount,omitempty"`
}

// Details about the capture configuration for the OffSessionPayment.
type V2PaymentsOffSessionPaymentCreateCaptureParams struct {
	// The method to use to capture the payment.
	CaptureMethod *string `form:"capture_method" json:"capture_method"`
}

// Payment method options for the card payment type.
type V2PaymentsOffSessionPaymentCreatePaymentMethodOptionsCardParams struct {
	// If you are making a Credential On File transaction with a previously saved card, you should pass the Network Transaction ID
	// from a prior initial authorization on Stripe (from a successful SetupIntent or a PaymentIntent with `setup_future_usage` set),
	// or one that you have obtained from another payment processor. This is a token from the network which uniquely identifies the transaction.
	// Visa calls this the Transaction ID, Mastercard calls this the Trace ID, and American Express calls this the Acquirer Reference Data.
	// Note that you should pass in a Network Transaction ID if you have one, regardless of whether this is a
	// Customer-Initiated Transaction (CIT) or a Merchant-Initiated Transaction (MIT).
	NetworkTransactionID *string `form:"network_transaction_id" json:"network_transaction_id"`
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

// The data that automatically creates a Transfer after the payment finalizes. Learn more about the use case for [connected accounts](https://docs.corp.stripe.com/payments/connected-accounts).
type V2PaymentsOffSessionPaymentCreateTransferDataParams struct {
	// The amount transferred to the destination account. This transfer will occur
	// automatically after the payment succeeds. If no amount is specified, by default
	// the entire payment amount is transferred to the destination account. The amount
	// must be less than or equal to the
	// [amount_requested](https://docs.corp.stripe.com/api/v2/off-session-payments/object?api-version=2025-05-28.preview#v2_off_session_payment_object-amount_requested),
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
	// The frequency of the underlying payment.
	Cadence *string `form:"cadence" json:"cadence"`
	// Details about the capture configuration for the OffSessionPayment.
	Capture *V2PaymentsOffSessionPaymentCreateCaptureParams `form:"capture" json:"capture,omitempty"`
	// ID of the Customer to which this OffSessionPayment belongs.
	Customer *string `form:"customer" json:"customer"`
	// Set of [key-value pairs](https://docs.corp.stripe.com/api/metadata) that you can
	// attach to an object. This can be useful for storing additional information about
	// the object in a structured format. Learn more about
	// [storing information in metadata](https://docs.corp.stripe.com/payments/payment-intents#storing-information-in-metadata).
	Metadata map[string]string `form:"metadata" json:"metadata"`
	// The account (if any) for which the funds of the OffSessionPayment are intended.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
	// ID of the payment method used in this OffSessionPayment.
	PaymentMethod *string `form:"payment_method" json:"payment_method"`
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
	// The data that automatically creates a Transfer after the payment finalizes. Learn more about the use case for [connected accounts](https://docs.corp.stripe.com/payments/connected-accounts).
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
