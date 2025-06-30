//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List OSPs matching filter.
type V2PaymentsOffSessionPaymentListParams struct {
	Params `form:"*"`
	// The page size limit, if not provided the default is 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// How you want stripe to retry the OSP.
type V2PaymentsOffSessionPaymentRetryDetailsParams struct {
	// How you want Stripe to retry the payment.
	RetryStrategy *string `form:"retry_strategy" json:"retry_strategy"`
}

// How you want to transfer the funds to your connected accounts.
type V2PaymentsOffSessionPaymentTransferDataParams struct {
	// Amount in minor units that you want to transfer.
	Amount *int64 `form:"amount" json:"amount,omitempty"`
	// ID of the connected account where you want money to go.
	Destination *string `form:"destination" json:"destination"`
}

// Create OSP.
type V2PaymentsOffSessionPaymentParams struct {
	Params `form:"*"`
	// Amount you want to collect.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
	// The frequency of the OSP.
	Cadence *string `form:"cadence" json:"cadence,omitempty"`
	// Customer that owns the provided payment method.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// Any of your internal data you want to track here.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The OBO merchant you want to use.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
	// Payment method you want to debit. Must be attached to a customer and set up for off-session usage.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// How you want stripe to retry the OSP.
	RetryDetails *V2PaymentsOffSessionPaymentRetryDetailsParams `form:"retry_details" json:"retry_details,omitempty"`
	// String you want to appear on your customer's statement.
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// Suffix appended to your account level descriptor.
	StatementDescriptorSuffix *string `form:"statement_descriptor_suffix" json:"statement_descriptor_suffix,omitempty"`
	// Test clock to be used for testing your retry handling. Only usable in a sandbox.
	TestClock *string `form:"test_clock" json:"test_clock,omitempty"`
	// How you want to transfer the funds to your connected accounts.
	TransferData *V2PaymentsOffSessionPaymentTransferDataParams `form:"transfer_data" json:"transfer_data,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2PaymentsOffSessionPaymentParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Cancel OSP.
type V2PaymentsOffSessionPaymentCancelParams struct {
	Params `form:"*"`
}

// How you want stripe to retry the OSP.
type V2PaymentsOffSessionPaymentCreateRetryDetailsParams struct {
	// How you want Stripe to retry the payment.
	RetryStrategy *string `form:"retry_strategy" json:"retry_strategy"`
}

// How you want to transfer the funds to your connected accounts.
type V2PaymentsOffSessionPaymentCreateTransferDataParams struct {
	// Amount in minor units that you want to transfer.
	Amount *int64 `form:"amount" json:"amount,omitempty"`
	// ID of the connected account where you want money to go.
	Destination *string `form:"destination" json:"destination"`
}

// Create OSP.
type V2PaymentsOffSessionPaymentCreateParams struct {
	Params `form:"*"`
	// Amount you want to collect.
	Amount *Amount `form:"amount" json:"amount"`
	// The frequency of the OSP.
	Cadence *string `form:"cadence" json:"cadence"`
	// Customer that owns the provided payment method.
	Customer *string `form:"customer" json:"customer"`
	// Any of your internal data you want to track here.
	Metadata map[string]string `form:"metadata" json:"metadata"`
	// The OBO merchant you want to use.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
	// Payment method you want to debit. Must be attached to a customer and set up for off-session usage.
	PaymentMethod *string `form:"payment_method" json:"payment_method"`
	// How you want stripe to retry the OSP.
	RetryDetails *V2PaymentsOffSessionPaymentCreateRetryDetailsParams `form:"retry_details" json:"retry_details,omitempty"`
	// String you want to appear on your customer's statement.
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// Suffix appended to your account level descriptor.
	StatementDescriptorSuffix *string `form:"statement_descriptor_suffix" json:"statement_descriptor_suffix,omitempty"`
	// Test clock to be used for testing your retry handling. Only usable in a sandbox.
	TestClock *string `form:"test_clock" json:"test_clock,omitempty"`
	// How you want to transfer the funds to your connected accounts.
	TransferData *V2PaymentsOffSessionPaymentCreateTransferDataParams `form:"transfer_data" json:"transfer_data,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2PaymentsOffSessionPaymentCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieve OSP by ID.
type V2PaymentsOffSessionPaymentRetrieveParams struct {
	Params `form:"*"`
}
