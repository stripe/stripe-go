//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The rails used to send funds.
type TreasuryOutboundPaymentDestinationPaymentMethodDetailsFinancialAccountNetwork string

// List of values that TreasuryOutboundPaymentDestinationPaymentMethodDetailsFinancialAccountNetwork can take
const (
	TreasuryOutboundPaymentDestinationPaymentMethodDetailsFinancialAccountNetworkStripe TreasuryOutboundPaymentDestinationPaymentMethodDetailsFinancialAccountNetwork = "stripe"
)

// The type of the payment method used in the OutboundPayment.
type TreasuryOutboundPaymentDestinationPaymentMethodDetailsType string

// List of values that TreasuryOutboundPaymentDestinationPaymentMethodDetailsType can take
const (
	TreasuryOutboundPaymentDestinationPaymentMethodDetailsTypeFinancialAccount TreasuryOutboundPaymentDestinationPaymentMethodDetailsType = "financial_account"
	TreasuryOutboundPaymentDestinationPaymentMethodDetailsTypeUSBankAccount    TreasuryOutboundPaymentDestinationPaymentMethodDetailsType = "us_bank_account"
)

// Account holder type: individual or company.
type TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountAccountHolderType string

// List of values that TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountAccountHolderType can take
const (
	TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountAccountHolderTypeCompany    TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountAccountHolderType = "company"
	TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountAccountHolderTypeIndividual TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountAccountHolderType = "individual"
)

// Account type: checkings or savings. Defaults to checking if omitted.
type TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountAccountType string

// List of values that TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountAccountType can take
const (
	TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountAccountTypeChecking TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountAccountType = "checking"
	TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountAccountTypeSavings  TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountAccountType = "savings"
)

// The network rails used. See the [docs](https://docs.stripe.com/treasury/money-movement/timelines) to learn more about money movement timelines for each network type.
type TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountNetwork string

// List of values that TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountNetwork can take
const (
	TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountNetworkACH            TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountNetwork = "ach"
	TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountNetworkUSDomesticWire TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountNetwork = "us_domestic_wire"
)

// The purpose of the OutboundPayment, if applicable.
type TreasuryOutboundPaymentPurpose string

// List of values that TreasuryOutboundPaymentPurpose can take
const (
	TreasuryOutboundPaymentPurposePayroll TreasuryOutboundPaymentPurpose = "payroll"
)

// Reason for the return.
type TreasuryOutboundPaymentReturnedDetailsCode string

// List of values that TreasuryOutboundPaymentReturnedDetailsCode can take
const (
	TreasuryOutboundPaymentReturnedDetailsCodeAccountClosed              TreasuryOutboundPaymentReturnedDetailsCode = "account_closed"
	TreasuryOutboundPaymentReturnedDetailsCodeAccountFrozen              TreasuryOutboundPaymentReturnedDetailsCode = "account_frozen"
	TreasuryOutboundPaymentReturnedDetailsCodeBankAccountRestricted      TreasuryOutboundPaymentReturnedDetailsCode = "bank_account_restricted"
	TreasuryOutboundPaymentReturnedDetailsCodeBankOwnershipChanged       TreasuryOutboundPaymentReturnedDetailsCode = "bank_ownership_changed"
	TreasuryOutboundPaymentReturnedDetailsCodeDeclined                   TreasuryOutboundPaymentReturnedDetailsCode = "declined"
	TreasuryOutboundPaymentReturnedDetailsCodeIncorrectAccountHolderName TreasuryOutboundPaymentReturnedDetailsCode = "incorrect_account_holder_name"
	TreasuryOutboundPaymentReturnedDetailsCodeInvalidAccountNumber       TreasuryOutboundPaymentReturnedDetailsCode = "invalid_account_number"
	TreasuryOutboundPaymentReturnedDetailsCodeInvalidCurrency            TreasuryOutboundPaymentReturnedDetailsCode = "invalid_currency"
	TreasuryOutboundPaymentReturnedDetailsCodeNoAccount                  TreasuryOutboundPaymentReturnedDetailsCode = "no_account"
	TreasuryOutboundPaymentReturnedDetailsCodeOther                      TreasuryOutboundPaymentReturnedDetailsCode = "other"
)

// Current status of the OutboundPayment: `processing`, `failed`, `posted`, `returned`, `canceled`. An OutboundPayment is `processing` if it has been created and is pending. The status changes to `posted` once the OutboundPayment has been "confirmed" and funds have left the account, or to `failed` or `canceled`. If an OutboundPayment fails to arrive at its destination, its status will change to `returned`.
type TreasuryOutboundPaymentStatus string

// List of values that TreasuryOutboundPaymentStatus can take
const (
	TreasuryOutboundPaymentStatusCanceled   TreasuryOutboundPaymentStatus = "canceled"
	TreasuryOutboundPaymentStatusFailed     TreasuryOutboundPaymentStatus = "failed"
	TreasuryOutboundPaymentStatusPosted     TreasuryOutboundPaymentStatus = "posted"
	TreasuryOutboundPaymentStatusProcessing TreasuryOutboundPaymentStatus = "processing"
	TreasuryOutboundPaymentStatusReturned   TreasuryOutboundPaymentStatus = "returned"
)

// The US bank account network used to send funds.
type TreasuryOutboundPaymentTrackingDetailsType string

// List of values that TreasuryOutboundPaymentTrackingDetailsType can take
const (
	TreasuryOutboundPaymentTrackingDetailsTypeACH            TreasuryOutboundPaymentTrackingDetailsType = "ach"
	TreasuryOutboundPaymentTrackingDetailsTypeUSDomesticWire TreasuryOutboundPaymentTrackingDetailsType = "us_domestic_wire"
)

// Returns a list of OutboundPayments sent from the specified FinancialAccount.
type TreasuryOutboundPaymentListParams struct {
	ListParams `form:"*"`
	// Only return OutboundPayments that were created during the given date interval.
	Created *int64 `form:"created" json:"created,omitempty"`
	// Only return OutboundPayments that were created during the given date interval.
	CreatedRange *RangeQueryParams `form:"created" json:"-"`
	// Only return OutboundPayments sent to this customer.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Returns objects associated with this FinancialAccount.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
	// Only return OutboundPayments that have the given status: `processing`, `failed`, `posted`, `returned`, or `canceled`.
	Status *string `form:"status" json:"status,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TreasuryOutboundPaymentListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
type TreasuryOutboundPaymentDestinationPaymentMethodDataBillingDetailsParams struct {
	// Billing address.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// Email address.
	Email *string `form:"email" json:"email,omitempty"`
	// Full name.
	Name *string `form:"name" json:"name,omitempty"`
	// Billing phone number (including extension).
	Phone       *string                                                                             `form:"phone" json:"phone,omitempty"`
	UnsetFields []TreasuryOutboundPaymentDestinationPaymentMethodDataBillingDetailsParamsUnsetField `form:"-" json:"-"`
}

// TreasuryOutboundPaymentDestinationPaymentMethodDataBillingDetailsParamsUnsetField is the list of fields that can be cleared/unset on TreasuryOutboundPaymentDestinationPaymentMethodDataBillingDetailsParams.
type TreasuryOutboundPaymentDestinationPaymentMethodDataBillingDetailsParamsUnsetField string

const (
	TreasuryOutboundPaymentDestinationPaymentMethodDataBillingDetailsParamsUnsetFieldAddress TreasuryOutboundPaymentDestinationPaymentMethodDataBillingDetailsParamsUnsetField = "address"
	TreasuryOutboundPaymentDestinationPaymentMethodDataBillingDetailsParamsUnsetFieldEmail   TreasuryOutboundPaymentDestinationPaymentMethodDataBillingDetailsParamsUnsetField = "email"
	TreasuryOutboundPaymentDestinationPaymentMethodDataBillingDetailsParamsUnsetFieldName    TreasuryOutboundPaymentDestinationPaymentMethodDataBillingDetailsParamsUnsetField = "name"
	TreasuryOutboundPaymentDestinationPaymentMethodDataBillingDetailsParamsUnsetFieldPhone   TreasuryOutboundPaymentDestinationPaymentMethodDataBillingDetailsParamsUnsetField = "phone"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *TreasuryOutboundPaymentDestinationPaymentMethodDataBillingDetailsParams) AddUnsetField(field TreasuryOutboundPaymentDestinationPaymentMethodDataBillingDetailsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Required hash if type is set to `us_bank_account`.
type TreasuryOutboundPaymentDestinationPaymentMethodDataUSBankAccountParams struct {
	// Account holder type: individual or company.
	AccountHolderType *string `form:"account_holder_type" json:"account_holder_type,omitempty"`
	// Account number of the bank account.
	AccountNumber *string `form:"account_number" json:"account_number,omitempty"`
	// Account type: checkings or savings. Defaults to checking if omitted.
	AccountType *string `form:"account_type" json:"account_type,omitempty"`
	// The ID of a Financial Connections Account to use as a payment method.
	FinancialConnectionsAccount *string `form:"financial_connections_account" json:"financial_connections_account,omitempty"`
	// Routing number of the bank account.
	RoutingNumber *string `form:"routing_number" json:"routing_number,omitempty"`
}

// Hash used to generate the PaymentMethod to be used for this OutboundPayment. Exclusive with `destination_payment_method`.
type TreasuryOutboundPaymentDestinationPaymentMethodDataParams struct {
	// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
	BillingDetails *TreasuryOutboundPaymentDestinationPaymentMethodDataBillingDetailsParams `form:"billing_details" json:"billing_details,omitempty"`
	// Required if type is set to `financial_account`. The FinancialAccount ID to send funds to.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
	Type *string `form:"type" json:"type"`
	// Required hash if type is set to `us_bank_account`.
	USBankAccount *TreasuryOutboundPaymentDestinationPaymentMethodDataUSBankAccountParams `form:"us_bank_account" json:"us_bank_account,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *TreasuryOutboundPaymentDestinationPaymentMethodDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Optional fields for `us_bank_account`.
type TreasuryOutboundPaymentDestinationPaymentMethodOptionsUSBankAccountParams struct {
	// Specifies the network rails to be used. If not set, will default to the PaymentMethod's preferred network. See the [docs](https://docs.stripe.com/treasury/money-movement/timelines) to learn more about money movement timelines for each network type.
	Network *string `form:"network" json:"network,omitempty"`
}

// Payment method-specific configuration for this OutboundPayment.
type TreasuryOutboundPaymentDestinationPaymentMethodOptionsParams struct {
	// Optional fields for `us_bank_account`.
	USBankAccount *TreasuryOutboundPaymentDestinationPaymentMethodOptionsUSBankAccountParams `form:"us_bank_account" json:"us_bank_account,omitempty"`
	UnsetFields   []TreasuryOutboundPaymentDestinationPaymentMethodOptionsParamsUnsetField   `form:"-" json:"-"`
}

// TreasuryOutboundPaymentDestinationPaymentMethodOptionsParamsUnsetField is the list of fields that can be cleared/unset on TreasuryOutboundPaymentDestinationPaymentMethodOptionsParams.
type TreasuryOutboundPaymentDestinationPaymentMethodOptionsParamsUnsetField string

const (
	TreasuryOutboundPaymentDestinationPaymentMethodOptionsParamsUnsetFieldUSBankAccount TreasuryOutboundPaymentDestinationPaymentMethodOptionsParamsUnsetField = "us_bank_account"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *TreasuryOutboundPaymentDestinationPaymentMethodOptionsParams) AddUnsetField(field TreasuryOutboundPaymentDestinationPaymentMethodOptionsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// End user details.
type TreasuryOutboundPaymentEndUserDetailsParams struct {
	// IP address of the user initiating the OutboundPayment. Must be supplied if `present` is set to `true`.
	IPAddress *string `form:"ip_address" json:"ip_address,omitempty"`
	// `True` if the OutboundPayment creation request is being made on behalf of an end user by a platform. Otherwise, `false`.
	Present *bool `form:"present" json:"present"`
}

// Creates an OutboundPayment.
type TreasuryOutboundPaymentParams struct {
	Params `form:"*"`
	// Amount (in cents) to be transferred.
	Amount *int64 `form:"amount" json:"amount,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// ID of the customer to whom the OutboundPayment is sent. Must match the Customer attached to the `destination_payment_method` passed in.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description" json:"description,omitempty"`
	// The PaymentMethod to use as the payment instrument for the OutboundPayment. Exclusive with `destination_payment_method_data`.
	DestinationPaymentMethod *string `form:"destination_payment_method" json:"destination_payment_method,omitempty"`
	// Hash used to generate the PaymentMethod to be used for this OutboundPayment. Exclusive with `destination_payment_method`.
	DestinationPaymentMethodData *TreasuryOutboundPaymentDestinationPaymentMethodDataParams `form:"destination_payment_method_data" json:"destination_payment_method_data,omitempty"`
	// Payment method-specific configuration for this OutboundPayment.
	DestinationPaymentMethodOptions *TreasuryOutboundPaymentDestinationPaymentMethodOptionsParams `form:"destination_payment_method_options" json:"destination_payment_method_options,omitempty"`
	// End user details.
	EndUserDetails *TreasuryOutboundPaymentEndUserDetailsParams `form:"end_user_details" json:"end_user_details,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The FinancialAccount to pull funds from.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The purpose of the OutboundPayment, if applicable. This list is not exhaustive, do not specify this parameter if your purpose does not match any that are provided.
	Purpose *string `form:"purpose" json:"purpose,omitempty"`
	// The description that appears on the receiving end for this OutboundPayment (for example, bank statement for external bank transfer). Maximum 10 characters for `ach` payments, 140 characters for `us_domestic_wire` payments, or 500 characters for `stripe` network transfers. Can only include -#.$&*, spaces, and alphanumeric characters. The default value is "payment".
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TreasuryOutboundPaymentParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *TreasuryOutboundPaymentParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Cancel an OutboundPayment.
type TreasuryOutboundPaymentCancelParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TreasuryOutboundPaymentCancelParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
type TreasuryOutboundPaymentCreateDestinationPaymentMethodDataBillingDetailsParams struct {
	// Billing address.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// Email address.
	Email *string `form:"email" json:"email,omitempty"`
	// Full name.
	Name *string `form:"name" json:"name,omitempty"`
	// Billing phone number (including extension).
	Phone       *string                                                                                   `form:"phone" json:"phone,omitempty"`
	UnsetFields []TreasuryOutboundPaymentCreateDestinationPaymentMethodDataBillingDetailsParamsUnsetField `form:"-" json:"-"`
}

// TreasuryOutboundPaymentCreateDestinationPaymentMethodDataBillingDetailsParamsUnsetField is the list of fields that can be cleared/unset on TreasuryOutboundPaymentCreateDestinationPaymentMethodDataBillingDetailsParams.
type TreasuryOutboundPaymentCreateDestinationPaymentMethodDataBillingDetailsParamsUnsetField string

const (
	TreasuryOutboundPaymentCreateDestinationPaymentMethodDataBillingDetailsParamsUnsetFieldAddress TreasuryOutboundPaymentCreateDestinationPaymentMethodDataBillingDetailsParamsUnsetField = "address"
	TreasuryOutboundPaymentCreateDestinationPaymentMethodDataBillingDetailsParamsUnsetFieldEmail   TreasuryOutboundPaymentCreateDestinationPaymentMethodDataBillingDetailsParamsUnsetField = "email"
	TreasuryOutboundPaymentCreateDestinationPaymentMethodDataBillingDetailsParamsUnsetFieldName    TreasuryOutboundPaymentCreateDestinationPaymentMethodDataBillingDetailsParamsUnsetField = "name"
	TreasuryOutboundPaymentCreateDestinationPaymentMethodDataBillingDetailsParamsUnsetFieldPhone   TreasuryOutboundPaymentCreateDestinationPaymentMethodDataBillingDetailsParamsUnsetField = "phone"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *TreasuryOutboundPaymentCreateDestinationPaymentMethodDataBillingDetailsParams) AddUnsetField(field TreasuryOutboundPaymentCreateDestinationPaymentMethodDataBillingDetailsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Required hash if type is set to `us_bank_account`.
type TreasuryOutboundPaymentCreateDestinationPaymentMethodDataUSBankAccountParams struct {
	// Account holder type: individual or company.
	AccountHolderType *string `form:"account_holder_type" json:"account_holder_type,omitempty"`
	// Account number of the bank account.
	AccountNumber *string `form:"account_number" json:"account_number,omitempty"`
	// Account type: checkings or savings. Defaults to checking if omitted.
	AccountType *string `form:"account_type" json:"account_type,omitempty"`
	// The ID of a Financial Connections Account to use as a payment method.
	FinancialConnectionsAccount *string `form:"financial_connections_account" json:"financial_connections_account,omitempty"`
	// Routing number of the bank account.
	RoutingNumber *string `form:"routing_number" json:"routing_number,omitempty"`
}

// Hash used to generate the PaymentMethod to be used for this OutboundPayment. Exclusive with `destination_payment_method`.
type TreasuryOutboundPaymentCreateDestinationPaymentMethodDataParams struct {
	// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
	BillingDetails *TreasuryOutboundPaymentCreateDestinationPaymentMethodDataBillingDetailsParams `form:"billing_details" json:"billing_details,omitempty"`
	// Required if type is set to `financial_account`. The FinancialAccount ID to send funds to.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
	Type *string `form:"type" json:"type"`
	// Required hash if type is set to `us_bank_account`.
	USBankAccount *TreasuryOutboundPaymentCreateDestinationPaymentMethodDataUSBankAccountParams `form:"us_bank_account" json:"us_bank_account,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *TreasuryOutboundPaymentCreateDestinationPaymentMethodDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Optional fields for `us_bank_account`.
type TreasuryOutboundPaymentCreateDestinationPaymentMethodOptionsUSBankAccountParams struct {
	// Specifies the network rails to be used. If not set, will default to the PaymentMethod's preferred network. See the [docs](https://docs.stripe.com/treasury/money-movement/timelines) to learn more about money movement timelines for each network type.
	Network *string `form:"network" json:"network,omitempty"`
}

// Payment method-specific configuration for this OutboundPayment.
type TreasuryOutboundPaymentCreateDestinationPaymentMethodOptionsParams struct {
	// Optional fields for `us_bank_account`.
	USBankAccount *TreasuryOutboundPaymentCreateDestinationPaymentMethodOptionsUSBankAccountParams `form:"us_bank_account" json:"us_bank_account,omitempty"`
	UnsetFields   []TreasuryOutboundPaymentCreateDestinationPaymentMethodOptionsParamsUnsetField   `form:"-" json:"-"`
}

// TreasuryOutboundPaymentCreateDestinationPaymentMethodOptionsParamsUnsetField is the list of fields that can be cleared/unset on TreasuryOutboundPaymentCreateDestinationPaymentMethodOptionsParams.
type TreasuryOutboundPaymentCreateDestinationPaymentMethodOptionsParamsUnsetField string

const (
	TreasuryOutboundPaymentCreateDestinationPaymentMethodOptionsParamsUnsetFieldUSBankAccount TreasuryOutboundPaymentCreateDestinationPaymentMethodOptionsParamsUnsetField = "us_bank_account"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *TreasuryOutboundPaymentCreateDestinationPaymentMethodOptionsParams) AddUnsetField(field TreasuryOutboundPaymentCreateDestinationPaymentMethodOptionsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// End user details.
type TreasuryOutboundPaymentCreateEndUserDetailsParams struct {
	// IP address of the user initiating the OutboundPayment. Must be supplied if `present` is set to `true`.
	IPAddress *string `form:"ip_address" json:"ip_address,omitempty"`
	// `True` if the OutboundPayment creation request is being made on behalf of an end user by a platform. Otherwise, `false`.
	Present *bool `form:"present" json:"present"`
}

// Creates an OutboundPayment.
type TreasuryOutboundPaymentCreateParams struct {
	Params `form:"*"`
	// Amount (in cents) to be transferred.
	Amount *int64 `form:"amount" json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency"`
	// ID of the customer to whom the OutboundPayment is sent. Must match the Customer attached to the `destination_payment_method` passed in.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description" json:"description,omitempty"`
	// The PaymentMethod to use as the payment instrument for the OutboundPayment. Exclusive with `destination_payment_method_data`.
	DestinationPaymentMethod *string `form:"destination_payment_method" json:"destination_payment_method,omitempty"`
	// Hash used to generate the PaymentMethod to be used for this OutboundPayment. Exclusive with `destination_payment_method`.
	DestinationPaymentMethodData *TreasuryOutboundPaymentCreateDestinationPaymentMethodDataParams `form:"destination_payment_method_data" json:"destination_payment_method_data,omitempty"`
	// Payment method-specific configuration for this OutboundPayment.
	DestinationPaymentMethodOptions *TreasuryOutboundPaymentCreateDestinationPaymentMethodOptionsParams `form:"destination_payment_method_options" json:"destination_payment_method_options,omitempty"`
	// End user details.
	EndUserDetails *TreasuryOutboundPaymentCreateEndUserDetailsParams `form:"end_user_details" json:"end_user_details,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The FinancialAccount to pull funds from.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The purpose of the OutboundPayment, if applicable. This list is not exhaustive, do not specify this parameter if your purpose does not match any that are provided.
	Purpose *string `form:"purpose" json:"purpose,omitempty"`
	// The description that appears on the receiving end for this OutboundPayment (for example, bank statement for external bank transfer). Maximum 10 characters for `ach` payments, 140 characters for `us_domestic_wire` payments, or 500 characters for `stripe` network transfers. Can only include -#.$&*, spaces, and alphanumeric characters. The default value is "payment".
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TreasuryOutboundPaymentCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *TreasuryOutboundPaymentCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves the details of an existing OutboundPayment by passing the unique OutboundPayment ID from either the OutboundPayment creation request or OutboundPayment list.
type TreasuryOutboundPaymentRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TreasuryOutboundPaymentRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type TreasuryOutboundPaymentDestinationPaymentMethodDetailsBillingDetails struct {
	Address *Address `json:"address"`
	// Email address.
	Email string `json:"email"`
	// Full name.
	Name string `json:"name"`
}
type TreasuryOutboundPaymentDestinationPaymentMethodDetailsFinancialAccount struct {
	// Token of the FinancialAccount.
	ID string `json:"id"`
	// The rails used to send funds.
	Network TreasuryOutboundPaymentDestinationPaymentMethodDetailsFinancialAccountNetwork `json:"network"`
}
type TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccount struct {
	// Account holder type: individual or company.
	AccountHolderType TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountAccountHolderType `json:"account_holder_type"`
	// Account type: checkings or savings. Defaults to checking if omitted.
	AccountType TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountAccountType `json:"account_type"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// ID of the mandate used to make this payment.
	Mandate *Mandate `json:"mandate,omitempty"`
	// The network rails used. See the [docs](https://docs.stripe.com/treasury/money-movement/timelines) to learn more about money movement timelines for each network type.
	Network TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccountNetwork `json:"network"`
	// Routing number of the bank account.
	RoutingNumber string `json:"routing_number"`
}

// Details about the PaymentMethod for an OutboundPayment.
type TreasuryOutboundPaymentDestinationPaymentMethodDetails struct {
	BillingDetails   *TreasuryOutboundPaymentDestinationPaymentMethodDetailsBillingDetails   `json:"billing_details"`
	FinancialAccount *TreasuryOutboundPaymentDestinationPaymentMethodDetailsFinancialAccount `json:"financial_account,omitempty"`
	// The type of the payment method used in the OutboundPayment.
	Type          TreasuryOutboundPaymentDestinationPaymentMethodDetailsType           `json:"type"`
	USBankAccount *TreasuryOutboundPaymentDestinationPaymentMethodDetailsUSBankAccount `json:"us_bank_account,omitempty"`
}

// Details about the end user.
type TreasuryOutboundPaymentEndUserDetails struct {
	// IP address of the user initiating the OutboundPayment. Set if `present` is set to `true`. IP address collection is required for risk and compliance reasons. This will be used to help determine if the OutboundPayment is authorized or should be blocked.
	IPAddress string `json:"ip_address"`
	// `true` if the OutboundPayment creation request is being made on behalf of an end user by a platform. Otherwise, `false`.
	Present bool `json:"present"`
}

// Details about a returned OutboundPayment. Only set when the status is `returned`.
type TreasuryOutboundPaymentReturnedDetails struct {
	// Reason for the return.
	Code TreasuryOutboundPaymentReturnedDetailsCode `json:"code"`
	// The Transaction associated with this object.
	Transaction *TreasuryTransaction `json:"transaction"`
}
type TreasuryOutboundPaymentStatusTransitions struct {
	// Timestamp describing when an OutboundPayment changed status to `canceled`.
	CanceledAt int64 `json:"canceled_at"`
	// Timestamp describing when an OutboundPayment changed status to `failed`.
	FailedAt int64 `json:"failed_at"`
	// Timestamp describing when an OutboundPayment changed status to `posted`.
	PostedAt int64 `json:"posted_at"`
	// Timestamp describing when an OutboundPayment changed status to `returned`.
	ReturnedAt int64 `json:"returned_at"`
}
type TreasuryOutboundPaymentTrackingDetailsACH struct {
	// ACH trace ID of the OutboundPayment for payments sent over the `ach` network.
	TraceID string `json:"trace_id"`
}
type TreasuryOutboundPaymentTrackingDetailsUSDomesticWire struct {
	// CHIPS System Sequence Number (SSN) of the OutboundPayment for payments sent over the `us_domestic_wire` network.
	Chips string `json:"chips"`
	// IMAD of the OutboundPayment for payments sent over the `us_domestic_wire` network.
	Imad string `json:"imad"`
	// OMAD of the OutboundPayment for payments sent over the `us_domestic_wire` network.
	Omad string `json:"omad"`
}

// Details about network-specific tracking information if available.
type TreasuryOutboundPaymentTrackingDetails struct {
	ACH *TreasuryOutboundPaymentTrackingDetailsACH `json:"ach,omitempty"`
	// The US bank account network used to send funds.
	Type           TreasuryOutboundPaymentTrackingDetailsType            `json:"type"`
	USDomesticWire *TreasuryOutboundPaymentTrackingDetailsUSDomesticWire `json:"us_domestic_wire,omitempty"`
}

// Use [OutboundPayments](https://docs.stripe.com/docs/treasury/moving-money/financial-accounts/out-of/outbound-payments) to send funds to another party's external bank account or [FinancialAccount](https://api.stripe.com#financial_accounts). To send money to an account belonging to the same user, use an [OutboundTransfer](https://api.stripe.com#outbound_transfers).
//
// Simulate OutboundPayment state changes with the `/v1/test_helpers/treasury/outbound_payments` endpoints. These methods can only be called on test mode objects.
//
// Related guide: [Moving money with Treasury using OutboundPayment objects](https://docs.stripe.com/docs/treasury/moving-money/financial-accounts/out-of/outbound-payments)
type TreasuryOutboundPayment struct {
	APIResource
	// Amount (in cents) transferred.
	Amount int64 `json:"amount"`
	// Returns `true` if the object can be canceled, and `false` otherwise.
	Cancelable bool `json:"cancelable"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// ID of the [customer](https://docs.stripe.com/api/customers) to whom an OutboundPayment is sent.
	Customer string `json:"customer"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// The PaymentMethod via which an OutboundPayment is sent. This field can be empty if the OutboundPayment was created using `destination_payment_method_data`.
	DestinationPaymentMethod string `json:"destination_payment_method"`
	// Details about the PaymentMethod for an OutboundPayment.
	DestinationPaymentMethodDetails *TreasuryOutboundPaymentDestinationPaymentMethodDetails `json:"destination_payment_method_details"`
	// Details about the end user.
	EndUserDetails *TreasuryOutboundPaymentEndUserDetails `json:"end_user_details"`
	// The date when funds are expected to arrive in the destination account.
	ExpectedArrivalDate int64 `json:"expected_arrival_date"`
	// The FinancialAccount that funds were pulled from.
	FinancialAccount string `json:"financial_account"`
	// A [hosted transaction receipt](https://docs.stripe.com/treasury/moving-money/regulatory-receipts) URL that is provided when money movement is considered regulated under Stripe's money transmission licenses.
	HostedRegulatoryReceiptURL string `json:"hosted_regulatory_receipt_url"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The purpose of the OutboundPayment, if applicable.
	Purpose TreasuryOutboundPaymentPurpose `json:"purpose,omitempty"`
	// Details about a returned OutboundPayment. Only set when the status is `returned`.
	ReturnedDetails *TreasuryOutboundPaymentReturnedDetails `json:"returned_details"`
	// The description that appears on the receiving end for an OutboundPayment (for example, bank statement for external bank transfer).
	StatementDescriptor string `json:"statement_descriptor"`
	// Current status of the OutboundPayment: `processing`, `failed`, `posted`, `returned`, `canceled`. An OutboundPayment is `processing` if it has been created and is pending. The status changes to `posted` once the OutboundPayment has been "confirmed" and funds have left the account, or to `failed` or `canceled`. If an OutboundPayment fails to arrive at its destination, its status will change to `returned`.
	Status            TreasuryOutboundPaymentStatus             `json:"status"`
	StatusTransitions *TreasuryOutboundPaymentStatusTransitions `json:"status_transitions"`
	// Details about network-specific tracking information if available.
	TrackingDetails *TreasuryOutboundPaymentTrackingDetails `json:"tracking_details"`
	// The Transaction associated with this object.
	Transaction *TreasuryTransaction `json:"transaction"`
}

// TreasuryOutboundPaymentList is a list of OutboundPayments as retrieved from a list endpoint.
type TreasuryOutboundPaymentList struct {
	APIResource
	ListMeta
	Data []*TreasuryOutboundPayment `json:"data"`
}
