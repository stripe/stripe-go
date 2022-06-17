//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of the payment method used in the OutboundTransfer.
type TreasuryOutboundTransferDestinationPaymentMethodDetailsType string

// List of values that TreasuryOutboundTransferDestinationPaymentMethodDetailsType can take
const (
	TreasuryOutboundTransferDestinationPaymentMethodDetailsTypeUSBankAccount TreasuryOutboundTransferDestinationPaymentMethodDetailsType = "us_bank_account"
)

// Account holder type: individual or company.
type TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountAccountHolderType string

// List of values that TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountAccountHolderType can take
const (
	TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountAccountHolderTypeCompany    TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountAccountHolderType = "company"
	TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountAccountHolderTypeIndividual TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountAccountHolderType = "individual"
)

// Account type: checkings or savings. Defaults to checking if omitted.
type TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountAccountType string

// List of values that TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountAccountType can take
const (
	TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountAccountTypeChecking TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountAccountType = "checking"
	TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountAccountTypeSavings  TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountAccountType = "savings"
)

// The US bank account network used to send funds.
type TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountNetwork string

// List of values that TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountNetwork can take
const (
	TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountNetworkAch            TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountNetwork = "ach"
	TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountNetworkUSDomesticWire TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountNetwork = "us_domestic_wire"
)

// Reason for the return.
type TreasuryOutboundTransferReturnedDetailsCode string

// List of values that TreasuryOutboundTransferReturnedDetailsCode can take
const (
	TreasuryOutboundTransferReturnedDetailsCodeAccountClosed              TreasuryOutboundTransferReturnedDetailsCode = "account_closed"
	TreasuryOutboundTransferReturnedDetailsCodeAccountFrozen              TreasuryOutboundTransferReturnedDetailsCode = "account_frozen"
	TreasuryOutboundTransferReturnedDetailsCodeBankAccountRestricted      TreasuryOutboundTransferReturnedDetailsCode = "bank_account_restricted"
	TreasuryOutboundTransferReturnedDetailsCodeBankOwnershipChanged       TreasuryOutboundTransferReturnedDetailsCode = "bank_ownership_changed"
	TreasuryOutboundTransferReturnedDetailsCodeDeclined                   TreasuryOutboundTransferReturnedDetailsCode = "declined"
	TreasuryOutboundTransferReturnedDetailsCodeIncorrectAccountHolderName TreasuryOutboundTransferReturnedDetailsCode = "incorrect_account_holder_name"
	TreasuryOutboundTransferReturnedDetailsCodeInvalidAccountNumber       TreasuryOutboundTransferReturnedDetailsCode = "invalid_account_number"
	TreasuryOutboundTransferReturnedDetailsCodeInvalidCurrency            TreasuryOutboundTransferReturnedDetailsCode = "invalid_currency"
	TreasuryOutboundTransferReturnedDetailsCodeNoAccount                  TreasuryOutboundTransferReturnedDetailsCode = "no_account"
	TreasuryOutboundTransferReturnedDetailsCodeOther                      TreasuryOutboundTransferReturnedDetailsCode = "other"
)

// Current status of the OutboundTransfer: `processing`, `failed`, `canceled`, `posted`, `returned`. An OutboundTransfer is `processing` if it has been created and is pending. The status changes to `posted` once the OutboundTransfer has been "confirmed" and funds have left the account, or to `failed` or `canceled`. If an OutboundTransfer fails to arrive at its destination, its status will change to `returned`.
type TreasuryOutboundTransferStatus string

// List of values that TreasuryOutboundTransferStatus can take
const (
	TreasuryOutboundTransferStatusCanceled   TreasuryOutboundTransferStatus = "canceled"
	TreasuryOutboundTransferStatusFailed     TreasuryOutboundTransferStatus = "failed"
	TreasuryOutboundTransferStatusPosted     TreasuryOutboundTransferStatus = "posted"
	TreasuryOutboundTransferStatusProcessing TreasuryOutboundTransferStatus = "processing"
	TreasuryOutboundTransferStatusReturned   TreasuryOutboundTransferStatus = "returned"
)

// Optional fields for `us_bank_account`.
type TreasuryOutboundTransferDestinationPaymentMethodOptionsUSBankAccountParams struct {
	// Designate the OutboundTransfer as using a US bank account network configuration.
	Network *string `form:"network"`
}

// Hash describing payment method configuration details.
type TreasuryOutboundTransferDestinationPaymentMethodOptionsParams struct {
	// Optional fields for `us_bank_account`.
	USBankAccount *TreasuryOutboundTransferDestinationPaymentMethodOptionsUSBankAccountParams `form:"us_bank_account"`
}

// Creates an OutboundTransfer.
type TreasuryOutboundTransferParams struct {
	Params `form:"*"`
	// Amount (in cents) to be transferred.
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// The PaymentMethod to use as the payment instrument for the OutboundTransfer.
	DestinationPaymentMethod *string `form:"destination_payment_method"`
	// Hash describing payment method configuration details.
	DestinationPaymentMethodOptions *TreasuryOutboundTransferDestinationPaymentMethodOptionsParams `form:"destination_payment_method_options"`
	// The FinancialAccount to pull funds from.
	FinancialAccount *string `form:"financial_account"`
	// Statement descriptor to be shown on the receiving end of an OutboundTransfer. Maximum 10 characters for `ach` transfers or 140 characters for `wire` transfers. The default value is `transfer`.
	StatementDescriptor *string `form:"statement_descriptor"`
}

// Returns a list of OutboundTransfers sent from the specified FinancialAccount.
type TreasuryOutboundTransferListParams struct {
	ListParams `form:"*"`
	// Returns objects associated with this FinancialAccount.
	FinancialAccount *string `form:"financial_account"`
	// Only return OutboundTransfers that have the given status: `processing`, `canceled`, `failed`, `posted`, or `returned`.
	Status *string `form:"status"`
}

// An OutboundTransfer can be canceled if the funds have not yet been paid out.
type TreasuryOutboundTransferCancelParams struct {
	Params `form:"*"`
}
type TreasuryOutboundTransferDestinationPaymentMethodDetailsBillingDetails struct {
	Address *Address `json:"address"`
	// Email address.
	Email string `json:"email"`
	// Full name.
	Name string `json:"name"`
}
type TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccount struct {
	// Account holder type: individual or company.
	AccountHolderType TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountAccountHolderType `json:"account_holder_type"`
	// Account type: checkings or savings. Defaults to checking if omitted.
	AccountType TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountAccountType `json:"account_type"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// The US bank account network used to send funds.
	Network TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccountNetwork `json:"network"`
	// Routing number of the bank account.
	RoutingNumber string `json:"routing_number"`
}
type TreasuryOutboundTransferDestinationPaymentMethodDetails struct {
	BillingDetails *TreasuryOutboundTransferDestinationPaymentMethodDetailsBillingDetails `json:"billing_details"`
	// The type of the payment method used in the OutboundTransfer.
	Type          TreasuryOutboundTransferDestinationPaymentMethodDetailsType           `json:"type"`
	USBankAccount *TreasuryOutboundTransferDestinationPaymentMethodDetailsUSBankAccount `json:"us_bank_account"`
}

// Details about a returned OutboundTransfer. Only set when the status is `returned`.
type TreasuryOutboundTransferReturnedDetails struct {
	// Reason for the return.
	Code TreasuryOutboundTransferReturnedDetailsCode `json:"code"`
	// The Transaction associated with this object.
	Transaction *TreasuryTransaction `json:"transaction"`
}
type TreasuryOutboundTransferStatusTransitions struct {
	// Timestamp describing when an OutboundTransfer changed status to `canceled`
	CanceledAt int64 `json:"canceled_at"`
	// Timestamp describing when an OutboundTransfer changed status to `failed`
	FailedAt int64 `json:"failed_at"`
	// Timestamp describing when an OutboundTransfer changed status to `posted`
	PostedAt int64 `json:"posted_at"`
	// Timestamp describing when an OutboundTransfer changed status to `returned`
	ReturnedAt int64 `json:"returned_at"`
}

// Use OutboundTransfers to transfer funds from a [FinancialAccount](https://stripe.com/docs/api#financial_accounts) to a PaymentMethod belonging to the same entity. To send funds to a different party, use [OutboundPayments](https://stripe.com/docs/api#outbound_payments) instead. You can send funds over ACH rails or through a domestic wire transfer to a user's own external bank account.
//
// Simulate OutboundTransfer state changes with the `/v1/test_helpers/treasury/outbound_transfers` endpoints. These methods can only be called on test mode objects.
type TreasuryOutboundTransfer struct {
	APIResource
	// Amount (in cents) transferred.
	Amount int64 `json:"amount"`
	// Returns `true` if the object can be canceled, and `false` otherwise.
	Cancelable bool `json:"cancelable"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// The PaymentMethod used as the payment instrument for an OutboundTransfer.
	DestinationPaymentMethod        string                                                   `json:"destination_payment_method"`
	DestinationPaymentMethodDetails *TreasuryOutboundTransferDestinationPaymentMethodDetails `json:"destination_payment_method_details"`
	// The date when funds are expected to arrive in the destination account.
	ExpectedArrivalDate int64 `json:"expected_arrival_date"`
	// The FinancialAccount that funds were pulled from.
	FinancialAccount string `json:"financial_account"`
	// A [hosted transaction receipt](https://stripe.com/docs/treasury/moving-money/regulatory-receipts) URL that is provided when money movement is considered regulated under Stripe's money transmission licenses.
	HostedRegulatoryReceiptURL string `json:"hosted_regulatory_receipt_url"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Details about a returned OutboundTransfer. Only set when the status is `returned`.
	ReturnedDetails *TreasuryOutboundTransferReturnedDetails `json:"returned_details"`
	// Information about the OutboundTransfer to be sent to the recipient account.
	StatementDescriptor string `json:"statement_descriptor"`
	// Current status of the OutboundTransfer: `processing`, `failed`, `canceled`, `posted`, `returned`. An OutboundTransfer is `processing` if it has been created and is pending. The status changes to `posted` once the OutboundTransfer has been "confirmed" and funds have left the account, or to `failed` or `canceled`. If an OutboundTransfer fails to arrive at its destination, its status will change to `returned`.
	Status            TreasuryOutboundTransferStatus             `json:"status"`
	StatusTransitions *TreasuryOutboundTransferStatusTransitions `json:"status_transitions"`
	// The Transaction associated with this object.
	Transaction *TreasuryTransaction `json:"transaction"`
}

// TreasuryOutboundTransferList is a list of OutboundTransfers as retrieved from a list endpoint.
type TreasuryOutboundTransferList struct {
	APIResource
	ListMeta
	Data []*TreasuryOutboundTransfer `json:"data"`
}
