//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The button style for the choice. Can be `primary` or `secondary`.
type TerminalReaderActionCollectInputsInputSelectionChoiceStyle string

// List of values that TerminalReaderActionCollectInputsInputSelectionChoiceStyle can take
const (
	TerminalReaderActionCollectInputsInputSelectionChoiceStylePrimary   TerminalReaderActionCollectInputsInputSelectionChoiceStyle = "primary"
	TerminalReaderActionCollectInputsInputSelectionChoiceStyleSecondary TerminalReaderActionCollectInputsInputSelectionChoiceStyle = "secondary"
)

// The toggle's default value. Can be `enabled` or `disabled`.
type TerminalReaderActionCollectInputsInputToggleDefaultValue string

// List of values that TerminalReaderActionCollectInputsInputToggleDefaultValue can take
const (
	TerminalReaderActionCollectInputsInputToggleDefaultValueDisabled TerminalReaderActionCollectInputsInputToggleDefaultValue = "disabled"
	TerminalReaderActionCollectInputsInputToggleDefaultValueEnabled  TerminalReaderActionCollectInputsInputToggleDefaultValue = "enabled"
)

// The toggle's collected value. Can be `enabled` or `disabled`.
type TerminalReaderActionCollectInputsInputToggleValue string

// List of values that TerminalReaderActionCollectInputsInputToggleValue can take
const (
	TerminalReaderActionCollectInputsInputToggleValueDisabled TerminalReaderActionCollectInputsInputToggleValue = "disabled"
	TerminalReaderActionCollectInputsInputToggleValueEnabled  TerminalReaderActionCollectInputsInputToggleValue = "enabled"
)

// Type of input being collected.
type TerminalReaderActionCollectInputsInputType string

// List of values that TerminalReaderActionCollectInputsInputType can take
const (
	TerminalReaderActionCollectInputsInputTypeEmail     TerminalReaderActionCollectInputsInputType = "email"
	TerminalReaderActionCollectInputsInputTypeNumeric   TerminalReaderActionCollectInputsInputType = "numeric"
	TerminalReaderActionCollectInputsInputTypePhone     TerminalReaderActionCollectInputsInputType = "phone"
	TerminalReaderActionCollectInputsInputTypeSelection TerminalReaderActionCollectInputsInputType = "selection"
	TerminalReaderActionCollectInputsInputTypeSignature TerminalReaderActionCollectInputsInputType = "signature"
	TerminalReaderActionCollectInputsInputTypeText      TerminalReaderActionCollectInputsInputType = "text"
)

// The reason for the refund.
type TerminalReaderActionRefundPaymentReason string

// List of values that TerminalReaderActionRefundPaymentReason can take
const (
	TerminalReaderActionRefundPaymentReasonDuplicate           TerminalReaderActionRefundPaymentReason = "duplicate"
	TerminalReaderActionRefundPaymentReasonFraudulent          TerminalReaderActionRefundPaymentReason = "fraudulent"
	TerminalReaderActionRefundPaymentReasonRequestedByCustomer TerminalReaderActionRefundPaymentReason = "requested_by_customer"
)

// Type of information to be displayed by the reader. Only `cart` is currently supported.
type TerminalReaderActionSetReaderDisplayType string

// List of values that TerminalReaderActionSetReaderDisplayType can take
const (
	TerminalReaderActionSetReaderDisplayTypeCart TerminalReaderActionSetReaderDisplayType = "cart"
)

// Status of the action performed by the reader.
type TerminalReaderActionStatus string

// List of values that TerminalReaderActionStatus can take
const (
	TerminalReaderActionStatusFailed     TerminalReaderActionStatus = "failed"
	TerminalReaderActionStatusInProgress TerminalReaderActionStatus = "in_progress"
	TerminalReaderActionStatusSucceeded  TerminalReaderActionStatus = "succeeded"
)

// Type of action performed by the reader.
type TerminalReaderActionType string

// List of values that TerminalReaderActionType can take
const (
	TerminalReaderActionTypeCollectInputs        TerminalReaderActionType = "collect_inputs"
	TerminalReaderActionTypeCollectPaymentMethod TerminalReaderActionType = "collect_payment_method"
	TerminalReaderActionTypeConfirmPaymentIntent TerminalReaderActionType = "confirm_payment_intent"
	TerminalReaderActionTypeProcessPaymentIntent TerminalReaderActionType = "process_payment_intent"
	TerminalReaderActionTypeProcessSetupIntent   TerminalReaderActionType = "process_setup_intent"
	TerminalReaderActionTypeRefundPayment        TerminalReaderActionType = "refund_payment"
	TerminalReaderActionTypeSetReaderDisplay     TerminalReaderActionType = "set_reader_display"
)

// Device type of the reader.
type TerminalReaderDeviceType string

// List of values that TerminalReaderDeviceType can take
const (
	TerminalReaderDeviceTypeBBPOSChipper2X      TerminalReaderDeviceType = "bbpos_chipper2x"
	TerminalReaderDeviceTypeBBPOSWisePad3       TerminalReaderDeviceType = "bbpos_wisepad3"
	TerminalReaderDeviceTypeBBPOSWisePOSE       TerminalReaderDeviceType = "bbpos_wisepos_e"
	TerminalReaderDeviceTypeMobilePhoneReader   TerminalReaderDeviceType = "mobile_phone_reader"
	TerminalReaderDeviceTypeSimulatedStripeS700 TerminalReaderDeviceType = "simulated_stripe_s700"
	TerminalReaderDeviceTypeSimulatedStripeS710 TerminalReaderDeviceType = "simulated_stripe_s710"
	TerminalReaderDeviceTypeSimulatedWisePOSE   TerminalReaderDeviceType = "simulated_wisepos_e"
	TerminalReaderDeviceTypeStripeM2            TerminalReaderDeviceType = "stripe_m2"
	TerminalReaderDeviceTypeStripeS700          TerminalReaderDeviceType = "stripe_s700"
	TerminalReaderDeviceTypeStripeS710          TerminalReaderDeviceType = "stripe_s710"
	TerminalReaderDeviceTypeVerifoneP400        TerminalReaderDeviceType = "verifone_P400"
)

// The networking status of the reader. We do not recommend using this field in flows that may block taking payments.
type TerminalReaderStatus string

// List of values that TerminalReaderStatus can take
const (
	TerminalReaderStatusOffline TerminalReaderStatus = "offline"
	TerminalReaderStatusOnline  TerminalReaderStatus = "online"
)

// Deletes a Reader object.
type TerminalReaderParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Custom label given to the reader for easier identification. If no label is specified, the registration code will be used.
	Label *string `form:"label"`
	// The location to assign the reader to.
	Location *string `form:"location"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// A code generated by the reader used for registering to an account.
	RegistrationCode *string `form:"registration_code"`
}

// AddExpand appends a new field to expand.
func (p *TerminalReaderParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *TerminalReaderParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Returns a list of Reader objects.
type TerminalReaderListParams struct {
	ListParams `form:"*"`
	// Filters readers by device type
	DeviceType *string `form:"device_type"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// A location ID to filter the response list to only readers at the specific location
	Location *string `form:"location"`
	// Filters readers by serial number
	SerialNumber *string `form:"serial_number"`
	// A status filter to filter readers to only offline or online readers
	Status *string `form:"status"`
}

// AddExpand appends a new field to expand.
func (p *TerminalReaderListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Cancels the current reader action. See [Programmatic Cancellation](https://docs.stripe.com/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven#programmatic-cancellation) for more details.
type TerminalReaderCancelActionParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *TerminalReaderCancelActionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Customize the text which will be displayed while collecting this input
type TerminalReaderCollectInputsInputCustomTextParams struct {
	// The description which will be displayed when collecting this input
	Description *string `form:"description"`
	// Custom text for the skip button. Maximum 14 characters.
	SkipButton *string `form:"skip_button"`
	// Custom text for the submit button. Maximum 30 characters.
	SubmitButton *string `form:"submit_button"`
	// The title which will be displayed when collecting this input
	Title *string `form:"title"`
}

// List of choices for the `selection` input
type TerminalReaderCollectInputsInputSelectionChoiceParams struct {
	// The unique identifier for this choice
	ID *string `form:"id"`
	// The style of the button which will be shown for this choice. Can be `primary` or `secondary`.
	Style *string `form:"style"`
	// The text which will be shown on the button for this choice
	Text *string `form:"text"`
}

// Options for the `selection` input
type TerminalReaderCollectInputsInputSelectionParams struct {
	// List of choices for the `selection` input
	Choices []*TerminalReaderCollectInputsInputSelectionChoiceParams `form:"choices"`
}

// List of toggles to be displayed and customization for the toggles
type TerminalReaderCollectInputsInputToggleParams struct {
	// The default value of the toggle. Can be `enabled` or `disabled`.
	DefaultValue *string `form:"default_value"`
	// The description which will be displayed for the toggle. Maximum 50 characters. At least one of title or description must be provided.
	Description *string `form:"description"`
	// The title which will be displayed for the toggle. Maximum 50 characters. At least one of title or description must be provided.
	Title *string `form:"title"`
}

// List of inputs to be collected from the customer using the Reader. Maximum 5 inputs.
type TerminalReaderCollectInputsInputParams struct {
	// Customize the text which will be displayed while collecting this input
	CustomText *TerminalReaderCollectInputsInputCustomTextParams `form:"custom_text"`
	// Indicate that this input is required, disabling the skip button
	Required *bool `form:"required"`
	// Options for the `selection` input
	Selection *TerminalReaderCollectInputsInputSelectionParams `form:"selection"`
	// List of toggles to be displayed and customization for the toggles
	Toggles []*TerminalReaderCollectInputsInputToggleParams `form:"toggles"`
	// The type of input to collect
	Type *string `form:"type"`
}

// Initiates an [input collection flow](https://docs.stripe.com/docs/terminal/features/collect-inputs) on a Reader to display input forms and collect information from your customers.
type TerminalReaderCollectInputsParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// List of inputs to be collected from the customer using the Reader. Maximum 5 inputs.
	Inputs []*TerminalReaderCollectInputsInputParams `form:"inputs"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
}

// AddExpand appends a new field to expand.
func (p *TerminalReaderCollectInputsParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *TerminalReaderCollectInputsParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Tipping configuration for this transaction.
type TerminalReaderCollectPaymentMethodCollectConfigTippingParams struct {
	// Amount used to calculate tip suggestions on tipping selection screen for this transaction. Must be a positive integer in the smallest currency unit (e.g., 100 cents to represent $1.00 or 100 to represent ¥100, a zero-decimal currency).
	AmountEligible *int64 `form:"amount_eligible"`
}

// Configuration overrides for this collection, such as tipping, surcharging, and customer cancellation settings.
type TerminalReaderCollectPaymentMethodCollectConfigParams struct {
	// This field indicates whether this payment method can be shown again to its customer in a checkout flow. Stripe products such as Checkout and Elements use this field to determine whether a payment method can be shown as a saved payment method in a checkout flow.
	AllowRedisplay *string `form:"allow_redisplay"`
	// Enables cancel button on transaction screens.
	EnableCustomerCancellation *bool `form:"enable_customer_cancellation"`
	// Override showing a tipping selection screen on this transaction.
	SkipTipping *bool `form:"skip_tipping"`
	// Tipping configuration for this transaction.
	Tipping *TerminalReaderCollectPaymentMethodCollectConfigTippingParams `form:"tipping"`
}

// Initiates a payment flow on a Reader and updates the PaymentIntent with card details before manual confirmation. See [Collecting a Payment method](https://docs.stripe.com/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=inspect#collect-a-paymentmethod) for more details.
type TerminalReaderCollectPaymentMethodParams struct {
	Params `form:"*"`
	// Configuration overrides for this collection, such as tipping, surcharging, and customer cancellation settings.
	CollectConfig *TerminalReaderCollectPaymentMethodCollectConfigParams `form:"collect_config"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The ID of the PaymentIntent to collect a payment method for.
	PaymentIntent *string `form:"payment_intent"`
}

// AddExpand appends a new field to expand.
func (p *TerminalReaderCollectPaymentMethodParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Configuration overrides for this confirmation, such as surcharge settings and return URL.
type TerminalReaderConfirmPaymentIntentConfirmConfigParams struct {
	// The URL to redirect your customer back to after they authenticate or cancel their payment on the payment method's app or site. If you'd prefer to redirect to a mobile application, you can alternatively supply an application URI scheme.
	ReturnURL *string `form:"return_url"`
}

// Finalizes a payment on a Reader. See [Confirming a Payment](https://docs.stripe.com/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=inspect#confirm-the-paymentintent) for more details.
type TerminalReaderConfirmPaymentIntentParams struct {
	Params `form:"*"`
	// Configuration overrides for this confirmation, such as surcharge settings and return URL.
	ConfirmConfig *TerminalReaderConfirmPaymentIntentConfirmConfigParams `form:"confirm_config"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The ID of the PaymentIntent to confirm.
	PaymentIntent *string `form:"payment_intent"`
}

// AddExpand appends a new field to expand.
func (p *TerminalReaderConfirmPaymentIntentParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Tipping configuration for this transaction.
type TerminalReaderProcessPaymentIntentProcessConfigTippingParams struct {
	// Amount used to calculate tip suggestions on tipping selection screen for this transaction. Must be a positive integer in the smallest currency unit (e.g., 100 cents to represent $1.00 or 100 to represent ¥100, a zero-decimal currency).
	AmountEligible *int64 `form:"amount_eligible"`
}

// Configuration overrides for this transaction, such as tipping and customer cancellation settings.
type TerminalReaderProcessPaymentIntentProcessConfigParams struct {
	// This field indicates whether this payment method can be shown again to its customer in a checkout flow. Stripe products such as Checkout and Elements use this field to determine whether a payment method can be shown as a saved payment method in a checkout flow.
	AllowRedisplay *string `form:"allow_redisplay"`
	// Enables cancel button on transaction screens.
	EnableCustomerCancellation *bool `form:"enable_customer_cancellation"`
	// The URL to redirect your customer back to after they authenticate or cancel their payment on the payment method's app or site. If you'd prefer to redirect to a mobile application, you can alternatively supply an application URI scheme.
	ReturnURL *string `form:"return_url"`
	// Override showing a tipping selection screen on this transaction.
	SkipTipping *bool `form:"skip_tipping"`
	// Tipping configuration for this transaction.
	Tipping *TerminalReaderProcessPaymentIntentProcessConfigTippingParams `form:"tipping"`
}

// Initiates a payment flow on a Reader. See [process the payment](https://docs.stripe.com/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=immediately#process-payment) for more details.
type TerminalReaderProcessPaymentIntentParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The ID of the PaymentIntent to process on the reader.
	PaymentIntent *string `form:"payment_intent"`
	// Configuration overrides for this transaction, such as tipping and customer cancellation settings.
	ProcessConfig *TerminalReaderProcessPaymentIntentProcessConfigParams `form:"process_config"`
}

// AddExpand appends a new field to expand.
func (p *TerminalReaderProcessPaymentIntentParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Configuration overrides for this setup, such as MOTO and customer cancellation settings.
type TerminalReaderProcessSetupIntentProcessConfigParams struct {
	// Enables cancel button on transaction screens.
	EnableCustomerCancellation *bool `form:"enable_customer_cancellation"`
}

// Initiates a SetupIntent flow on a Reader. See [Save directly without charging](https://docs.stripe.com/docs/terminal/features/saving-payment-details/save-directly) for more details.
type TerminalReaderProcessSetupIntentParams struct {
	Params `form:"*"`
	// This field indicates whether this payment method can be shown again to its customer in a checkout flow. Stripe products such as Checkout and Elements use this field to determine whether a payment method can be shown as a saved payment method in a checkout flow.
	AllowRedisplay *string `form:"allow_redisplay"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Configuration overrides for this setup, such as MOTO and customer cancellation settings.
	ProcessConfig *TerminalReaderProcessSetupIntentProcessConfigParams `form:"process_config"`
	// The ID of the SetupIntent to process on the reader.
	SetupIntent *string `form:"setup_intent"`
}

// AddExpand appends a new field to expand.
func (p *TerminalReaderProcessSetupIntentParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Configuration overrides for this refund, such as customer cancellation settings.
type TerminalReaderRefundPaymentRefundPaymentConfigParams struct {
	// Enables cancel button on transaction screens.
	EnableCustomerCancellation *bool `form:"enable_customer_cancellation"`
}

// Initiates an in-person refund on a Reader. See [Refund an Interac Payment](https://docs.stripe.com/docs/terminal/payments/regional?integration-country=CA#refund-an-interac-payment) for more details.
type TerminalReaderRefundPaymentParams struct {
	Params `form:"*"`
	// A positive integer in __cents__ representing how much of this charge to refund.
	Amount *int64 `form:"amount"`
	// ID of the Charge to refund.
	Charge *string `form:"charge"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// ID of the PaymentIntent to refund.
	PaymentIntent *string `form:"payment_intent"`
	// Boolean indicating whether the application fee should be refunded when refunding this charge. If a full charge refund is given, the full application fee will be refunded. Otherwise, the application fee will be refunded in an amount proportional to the amount of the charge refunded. An application fee can be refunded only by the application that created the charge.
	RefundApplicationFee *bool `form:"refund_application_fee"`
	// Configuration overrides for this refund, such as customer cancellation settings.
	RefundPaymentConfig *TerminalReaderRefundPaymentRefundPaymentConfigParams `form:"refund_payment_config"`
	// Boolean indicating whether the transfer should be reversed when refunding this charge. The transfer will be reversed proportionally to the amount being refunded (either the entire or partial amount). A transfer can be reversed only by the application that created the charge.
	ReverseTransfer *bool `form:"reverse_transfer"`
}

// AddExpand appends a new field to expand.
func (p *TerminalReaderRefundPaymentParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *TerminalReaderRefundPaymentParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Array of line items to display.
type TerminalReaderSetReaderDisplayCartLineItemParams struct {
	// The price of the item in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	Amount *int64 `form:"amount"`
	// The description or name of the item.
	Description *string `form:"description"`
	// The quantity of the line item being purchased.
	Quantity *int64 `form:"quantity"`
}

// Cart details to display on the reader screen, including line items, amounts, and currency.
type TerminalReaderSetReaderDisplayCartParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// Array of line items to display.
	LineItems []*TerminalReaderSetReaderDisplayCartLineItemParams `form:"line_items"`
	// The amount of tax in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	Tax *int64 `form:"tax"`
	// Total balance of cart due in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	Total *int64 `form:"total"`
}

// Sets the reader display to show [cart details](https://docs.stripe.com/docs/terminal/features/display).
type TerminalReaderSetReaderDisplayParams struct {
	Params `form:"*"`
	// Cart details to display on the reader screen, including line items, amounts, and currency.
	Cart *TerminalReaderSetReaderDisplayCartParams `form:"cart"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Type of information to display. Only `cart` is currently supported.
	Type *string `form:"type"`
}

// AddExpand appends a new field to expand.
func (p *TerminalReaderSetReaderDisplayParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Deletes a Reader object.
type TerminalReaderDeleteParams struct {
	Params `form:"*"`
}

// Retrieves a Reader object.
type TerminalReaderRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *TerminalReaderRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Updates a Reader object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
type TerminalReaderUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The new label of the reader.
	Label *string `form:"label"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
}

// AddExpand appends a new field to expand.
func (p *TerminalReaderUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *TerminalReaderUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Creates a new Reader object.
type TerminalReaderCreateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Custom label given to the reader for easier identification. If no label is specified, the registration code will be used.
	Label *string `form:"label"`
	// The location to assign the reader to.
	Location *string `form:"location"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// A code generated by the reader used for registering to an account.
	RegistrationCode *string `form:"registration_code"`
}

// AddExpand appends a new field to expand.
func (p *TerminalReaderCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *TerminalReaderCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Default text of input being collected.
type TerminalReaderActionCollectInputsInputCustomText struct {
	// Customize the default description for this input
	Description string `json:"description"`
	// Customize the default label for this input's skip button
	SkipButton string `json:"skip_button"`
	// Customize the default label for this input's submit button
	SubmitButton string `json:"submit_button"`
	// Customize the default title for this input
	Title string `json:"title"`
}

// Information about a email being collected using a reader
type TerminalReaderActionCollectInputsInputEmail struct {
	// The collected email address
	Value string `json:"value"`
}

// Information about a number being collected using a reader
type TerminalReaderActionCollectInputsInputNumeric struct {
	// The collected number
	Value string `json:"value"`
}

// Information about a phone number being collected using a reader
type TerminalReaderActionCollectInputsInputPhone struct {
	// The collected phone number
	Value string `json:"value"`
}

// List of possible choices to be selected
type TerminalReaderActionCollectInputsInputSelectionChoice struct {
	// The identifier for the selected choice. Maximum 50 characters.
	ID string `json:"id"`
	// The button style for the choice. Can be `primary` or `secondary`.
	Style TerminalReaderActionCollectInputsInputSelectionChoiceStyle `json:"style"`
	// The text to be selected. Maximum 30 characters.
	Text string `json:"text"`
}

// Information about a selection being collected using a reader
type TerminalReaderActionCollectInputsInputSelection struct {
	// List of possible choices to be selected
	Choices []*TerminalReaderActionCollectInputsInputSelectionChoice `json:"choices"`
	// The id of the selected choice
	ID string `json:"id"`
	// The text of the selected choice
	Text string `json:"text"`
}

// Information about a signature being collected using a reader
type TerminalReaderActionCollectInputsInputSignature struct {
	// The File ID of a collected signature image
	Value string `json:"value"`
}

// Information about text being collected using a reader
type TerminalReaderActionCollectInputsInputText struct {
	// The collected text value
	Value string `json:"value"`
}

// List of toggles being collected. Values are present if collection is complete.
type TerminalReaderActionCollectInputsInputToggle struct {
	// The toggle's default value. Can be `enabled` or `disabled`.
	DefaultValue TerminalReaderActionCollectInputsInputToggleDefaultValue `json:"default_value"`
	// The toggle's description text. Maximum 50 characters.
	Description string `json:"description"`
	// The toggle's title text. Maximum 50 characters.
	Title string `json:"title"`
	// The toggle's collected value. Can be `enabled` or `disabled`.
	Value TerminalReaderActionCollectInputsInputToggleValue `json:"value"`
}

// List of inputs to be collected.
type TerminalReaderActionCollectInputsInput struct {
	// Default text of input being collected.
	CustomText *TerminalReaderActionCollectInputsInputCustomText `json:"custom_text"`
	// Information about a email being collected using a reader
	Email *TerminalReaderActionCollectInputsInputEmail `json:"email"`
	// Information about a number being collected using a reader
	Numeric *TerminalReaderActionCollectInputsInputNumeric `json:"numeric"`
	// Information about a phone number being collected using a reader
	Phone *TerminalReaderActionCollectInputsInputPhone `json:"phone"`
	// Indicate that this input is required, disabling the skip button.
	Required bool `json:"required"`
	// Information about a selection being collected using a reader
	Selection *TerminalReaderActionCollectInputsInputSelection `json:"selection"`
	// Information about a signature being collected using a reader
	Signature *TerminalReaderActionCollectInputsInputSignature `json:"signature"`
	// Indicate that this input was skipped by the user.
	Skipped bool `json:"skipped"`
	// Information about text being collected using a reader
	Text *TerminalReaderActionCollectInputsInputText `json:"text"`
	// List of toggles being collected. Values are present if collection is complete.
	Toggles []*TerminalReaderActionCollectInputsInputToggle `json:"toggles"`
	// Type of input being collected.
	Type TerminalReaderActionCollectInputsInputType `json:"type"`
}

// Represents a reader action to collect customer inputs
type TerminalReaderActionCollectInputs struct {
	// List of inputs to be collected.
	Inputs []*TerminalReaderActionCollectInputsInput `json:"inputs"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
}

// Represents a per-transaction tipping configuration
type TerminalReaderActionCollectPaymentMethodCollectConfigTipping struct {
	// Amount used to calculate tip suggestions on tipping selection screen for this transaction. Must be a positive integer in the smallest currency unit (e.g., 100 cents to represent $1.00 or 100 to represent ¥100, a zero-decimal currency).
	AmountEligible int64 `json:"amount_eligible"`
}

// Represents a per-transaction override of a reader configuration
type TerminalReaderActionCollectPaymentMethodCollectConfig struct {
	// Enable customer-initiated cancellation when processing this payment.
	EnableCustomerCancellation bool `json:"enable_customer_cancellation"`
	// Override showing a tipping selection screen on this transaction.
	SkipTipping bool `json:"skip_tipping"`
	// Represents a per-transaction tipping configuration
	Tipping *TerminalReaderActionCollectPaymentMethodCollectConfigTipping `json:"tipping"`
}

// Represents a reader action to collect a payment method
type TerminalReaderActionCollectPaymentMethod struct {
	// Represents a per-transaction override of a reader configuration
	CollectConfig *TerminalReaderActionCollectPaymentMethodCollectConfig `json:"collect_config"`
	// Most recent PaymentIntent processed by the reader.
	PaymentIntent *PaymentIntent `json:"payment_intent"`
	// PaymentMethod objects represent your customer's payment instruments.
	// You can use them with [PaymentIntents](https://docs.stripe.com/payments/payment-intents) to collect payments or save them to
	// Customer objects to store instrument details for future payments.
	//
	// Related guides: [Payment Methods](https://docs.stripe.com/payments/payment-methods) and [More Payment Scenarios](https://docs.stripe.com/payments/more-payment-scenarios).
	PaymentMethod *PaymentMethod `json:"payment_method"`
}

// Represents a per-transaction override of a reader configuration
type TerminalReaderActionConfirmPaymentIntentConfirmConfig struct {
	// If the customer doesn't abandon authenticating the payment, they're redirected to this URL after completion.
	ReturnURL string `json:"return_url"`
}

// Represents a reader action to confirm a payment
type TerminalReaderActionConfirmPaymentIntent struct {
	// Represents a per-transaction override of a reader configuration
	ConfirmConfig *TerminalReaderActionConfirmPaymentIntentConfirmConfig `json:"confirm_config"`
	// Most recent PaymentIntent processed by the reader.
	PaymentIntent *PaymentIntent `json:"payment_intent"`
}

// Represents a per-transaction tipping configuration
type TerminalReaderActionProcessPaymentIntentProcessConfigTipping struct {
	// Amount used to calculate tip suggestions on tipping selection screen for this transaction. Must be a positive integer in the smallest currency unit (e.g., 100 cents to represent $1.00 or 100 to represent ¥100, a zero-decimal currency).
	AmountEligible int64 `json:"amount_eligible"`
}

// Represents a per-transaction override of a reader configuration
type TerminalReaderActionProcessPaymentIntentProcessConfig struct {
	// Enable customer-initiated cancellation when processing this payment.
	EnableCustomerCancellation bool `json:"enable_customer_cancellation"`
	// If the customer doesn't abandon authenticating the payment, they're redirected to this URL after completion.
	ReturnURL string `json:"return_url"`
	// Override showing a tipping selection screen on this transaction.
	SkipTipping bool `json:"skip_tipping"`
	// Represents a per-transaction tipping configuration
	Tipping *TerminalReaderActionProcessPaymentIntentProcessConfigTipping `json:"tipping"`
}

// Represents a reader action to process a payment intent
type TerminalReaderActionProcessPaymentIntent struct {
	// Most recent PaymentIntent processed by the reader.
	PaymentIntent *PaymentIntent `json:"payment_intent"`
	// Represents a per-transaction override of a reader configuration
	ProcessConfig *TerminalReaderActionProcessPaymentIntentProcessConfig `json:"process_config"`
}

// Represents a per-setup override of a reader configuration
type TerminalReaderActionProcessSetupIntentProcessConfig struct {
	// Enable customer-initiated cancellation when processing this SetupIntent.
	EnableCustomerCancellation bool `json:"enable_customer_cancellation"`
}

// Represents a reader action to process a setup intent
type TerminalReaderActionProcessSetupIntent struct {
	// ID of a card PaymentMethod generated from the card_present PaymentMethod that may be attached to a Customer for future transactions. Only present if it was possible to generate a card PaymentMethod.
	GeneratedCard string `json:"generated_card"`
	// Represents a per-setup override of a reader configuration
	ProcessConfig *TerminalReaderActionProcessSetupIntentProcessConfig `json:"process_config"`
	// Most recent SetupIntent processed by the reader.
	SetupIntent *SetupIntent `json:"setup_intent"`
}

// Represents a per-transaction override of a reader configuration
type TerminalReaderActionRefundPaymentRefundPaymentConfig struct {
	// Enable customer-initiated cancellation when refunding this payment.
	EnableCustomerCancellation bool `json:"enable_customer_cancellation"`
}

// Represents a reader action to refund a payment
type TerminalReaderActionRefundPayment struct {
	// The amount being refunded.
	Amount int64 `json:"amount"`
	// Charge that is being refunded.
	Charge *Charge `json:"charge"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// Payment intent that is being refunded.
	PaymentIntent *PaymentIntent `json:"payment_intent"`
	// The reason for the refund.
	Reason TerminalReaderActionRefundPaymentReason `json:"reason"`
	// Unique identifier for the refund object.
	Refund *Refund `json:"refund"`
	// Boolean indicating whether the application fee should be refunded when refunding this charge. If a full charge refund is given, the full application fee will be refunded. Otherwise, the application fee will be refunded in an amount proportional to the amount of the charge refunded. An application fee can be refunded only by the application that created the charge.
	RefundApplicationFee bool `json:"refund_application_fee"`
	// Represents a per-transaction override of a reader configuration
	RefundPaymentConfig *TerminalReaderActionRefundPaymentRefundPaymentConfig `json:"refund_payment_config"`
	// Boolean indicating whether the transfer should be reversed when refunding this charge. The transfer will be reversed proportionally to the amount being refunded (either the entire or partial amount). A transfer can be reversed only by the application that created the charge.
	ReverseTransfer bool `json:"reverse_transfer"`
}

// List of line items in the cart.
type TerminalReaderActionSetReaderDisplayCartLineItem struct {
	// The amount of the line item. A positive integer in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	Amount int64 `json:"amount"`
	// Description of the line item.
	Description string `json:"description"`
	// The quantity of the line item.
	Quantity int64 `json:"quantity"`
}

// Cart object to be displayed by the reader, including line items, amounts, and currency.
type TerminalReaderActionSetReaderDisplayCart struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// List of line items in the cart.
	LineItems []*TerminalReaderActionSetReaderDisplayCartLineItem `json:"line_items"`
	// Tax amount for the entire cart. A positive integer in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	Tax int64 `json:"tax"`
	// Total amount for the entire cart, including tax. A positive integer in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	Total int64 `json:"total"`
}

// Represents a reader action to set the reader display
type TerminalReaderActionSetReaderDisplay struct {
	// Cart object to be displayed by the reader, including line items, amounts, and currency.
	Cart *TerminalReaderActionSetReaderDisplayCart `json:"cart"`
	// Type of information to be displayed by the reader. Only `cart` is currently supported.
	Type TerminalReaderActionSetReaderDisplayType `json:"type"`
}

// The most recent action performed by the reader.
type TerminalReaderAction struct {
	// Represents a reader action to collect customer inputs
	CollectInputs *TerminalReaderActionCollectInputs `json:"collect_inputs"`
	// Represents a reader action to collect a payment method
	CollectPaymentMethod *TerminalReaderActionCollectPaymentMethod `json:"collect_payment_method"`
	// Represents a reader action to confirm a payment
	ConfirmPaymentIntent *TerminalReaderActionConfirmPaymentIntent `json:"confirm_payment_intent"`
	// Failure code, only set if status is `failed`.
	FailureCode string `json:"failure_code"`
	// Detailed failure message, only set if status is `failed`.
	FailureMessage string `json:"failure_message"`
	// Represents a reader action to process a payment intent
	ProcessPaymentIntent *TerminalReaderActionProcessPaymentIntent `json:"process_payment_intent"`
	// Represents a reader action to process a setup intent
	ProcessSetupIntent *TerminalReaderActionProcessSetupIntent `json:"process_setup_intent"`
	// Represents a reader action to refund a payment
	RefundPayment *TerminalReaderActionRefundPayment `json:"refund_payment"`
	// Represents a reader action to set the reader display
	SetReaderDisplay *TerminalReaderActionSetReaderDisplay `json:"set_reader_display"`
	// Status of the action performed by the reader.
	Status TerminalReaderActionStatus `json:"status"`
	// Type of action performed by the reader.
	Type TerminalReaderActionType `json:"type"`
}

// A Reader represents a physical device for accepting payment details.
//
// Related guide: [Connecting to a reader](https://docs.stripe.com/terminal/payments/connect-reader)
type TerminalReader struct {
	APIResource
	// The most recent action performed by the reader.
	Action  *TerminalReaderAction `json:"action"`
	Deleted bool                  `json:"deleted"`
	// The current software version of the reader.
	DeviceSwVersion string `json:"device_sw_version"`
	// Device type of the reader.
	DeviceType TerminalReaderDeviceType `json:"device_type"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The local IP address of the reader.
	IPAddress string `json:"ip_address"`
	// Custom label given to the reader for easier identification.
	Label string `json:"label"`
	// The last time this reader reported to Stripe backend. Timestamp is measured in milliseconds since the Unix epoch. Unlike most other Stripe timestamp fields which use seconds, this field uses milliseconds.
	LastSeenAt int64 `json:"last_seen_at"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The location identifier of the reader.
	Location *TerminalLocation `json:"location"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Serial number of the reader.
	SerialNumber string `json:"serial_number"`
	// The networking status of the reader. We do not recommend using this field in flows that may block taking payments.
	Status TerminalReaderStatus `json:"status"`
}

// TerminalReaderList is a list of Readers as retrieved from a list endpoint.
type TerminalReaderList struct {
	APIResource
	ListMeta
	Data []*TerminalReader `json:"data"`
}
