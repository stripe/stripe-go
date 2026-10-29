//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of OutboundPayments that match the provided filters.
type V2MoneyManagementOutboundPaymentListParams struct {
	Params `form:"*"`
	// Filter OutboundPayments by their creation time.
	Created *RangeQueryParams `form:"created" json:"created,omitempty"`
	// The maximum number of results to return.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Only return OutboundPayments sent to this recipient.
	Recipient *string `form:"recipient" json:"recipient,omitempty"`
	// Closed Enum. Only return OutboundPayments with this status.
	Status []*string `form:"status" json:"status,omitempty"`
}

// Delivery options to be used to send the OutboundPayment.
type V2MoneyManagementOutboundPaymentDeliveryOptionsParams struct {
	// Open Enum. Method for bank account.
	BankAccount *string `form:"bank_account" json:"bank_account,omitempty"`
}

// From which FinancialAccount and BalanceType to pull funds from.
type V2MoneyManagementOutboundPaymentFromParams struct {
	// Describes the FinancialAmount's currency drawn from.
	Currency *string `form:"currency" json:"currency"`
	// The FinancialAccount that funds were pulled from.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
}

// Details about the notification settings for the OutboundPayment recipient.
type V2MoneyManagementOutboundPaymentRecipientNotificationParams struct {
	// Closed Enum. Configuration option to enable or disable notifications to recipients.
	// Do not send notifications when `setting` is `none`. Use the account setting when `setting` is `configured` or omitted.
	Setting *string `form:"setting" json:"setting"`
}

// ACH-specific network options.
type V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHParams struct {
	// Open Enum. ACH submission timing.
	Submission *string `form:"submission" json:"submission,omitempty"`
	// The transaction purpose for this ACH payment.
	TransactionPurpose *string `form:"transaction_purpose" json:"transaction_purpose,omitempty"`
}

// Per-network configuration options.
type V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsParams struct {
	// ACH-specific network options.
	ACH *V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHParams `form:"ach" json:"ach,omitempty"`
}

// Options for bank account payout methods.
type V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountParams struct {
	// Per-network configuration options.
	PreferredNetworkOptions *V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsParams `form:"preferred_network_options" json:"preferred_network_options,omitempty"`
	// The preferred networks to use for this OutboundPayment.
	PreferredNetworks []*string `form:"preferred_networks" json:"preferred_networks,omitempty"`
}

// Payout method options for the OutboundPayment.
type V2MoneyManagementOutboundPaymentToPayoutMethodOptionsParams struct {
	// Options for bank account payout methods.
	BankAccount *V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountParams `form:"bank_account" json:"bank_account,omitempty"`
}

// To which payout method to send the OutboundPayment.
type V2MoneyManagementOutboundPaymentToParams struct {
	// Describes the currency to send to the recipient.
	// If included, this currency must match a currency supported by the destination.
	// Can be omitted in the following cases:
	// - destination only supports one currency
	// - destination supports multiple currencies and one matches `from.currency`
	// - destination supports multiple currencies and one matches `amount.currency`
	// When both `from.currency` and `amount.currency` are supported, Stripe selects `from.currency` to minimize currency conversion.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// The payout method which the OutboundPayment uses to send payout.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
	// Payout method options for the OutboundPayment.
	PayoutMethodOptions *V2MoneyManagementOutboundPaymentToPayoutMethodOptionsParams `form:"payout_method_options" json:"payout_method_options,omitempty"`
	// To which account the OutboundPayment is sent.
	Recipient *string `form:"recipient" json:"recipient"`
}

// Creates an OutboundPayment.
type V2MoneyManagementOutboundPaymentParams struct {
	Params `form:"*"`
	// The requested amount and currency for the OutboundPayment.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
	// Delivery options to be used to send the OutboundPayment.
	DeliveryOptions *V2MoneyManagementOutboundPaymentDeliveryOptionsParams `form:"delivery_options" json:"delivery_options,omitempty"`
	// An arbitrary string attached to the OutboundPayment. Often useful for displaying to users.
	Description *string `form:"description" json:"description,omitempty"`
	// From which FinancialAccount and BalanceType to pull funds from.
	From *V2MoneyManagementOutboundPaymentFromParams `form:"from" json:"from,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The quote for this OutboundPayment. Only required for countries with regulatory mandates to display fee estimates before OutboundPayment creation.
	OutboundPaymentQuote *string `form:"outbound_payment_quote" json:"outbound_payment_quote,omitempty"`
	// The purpose of the OutboundPayment.
	Purpose *string `form:"purpose" json:"purpose,omitempty"`
	// Details about the notification settings for the OutboundPayment recipient.
	RecipientNotification *V2MoneyManagementOutboundPaymentRecipientNotificationParams `form:"recipient_notification" json:"recipient_notification,omitempty"`
	// The ID of the RecipientVerification for this OutboundPayment. Required where recipient-name verification is mandated before creating an OutboundPayment.
	RecipientVerification *string `form:"recipient_verification" json:"recipient_verification,omitempty"`
	// The description that appears on the receiving end for an OutboundPayment (for example, on a bank statement). Must be between 3 and 22 characters long for most destinations (500 for FinancialAccount destinations), and not contain profanity.
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// To which payout method to send the OutboundPayment.
	To *V2MoneyManagementOutboundPaymentToParams `form:"to" json:"to,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementOutboundPaymentParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Cancels an OutboundPayment. Only processing OutboundPayments can be canceled.
type V2MoneyManagementOutboundPaymentCancelParams struct {
	Params `form:"*"`
}

// Delivery options to be used to send the OutboundPayment.
type V2MoneyManagementOutboundPaymentCreateDeliveryOptionsParams struct {
	// Open Enum. Method for bank account.
	BankAccount *string `form:"bank_account" json:"bank_account,omitempty"`
}

// From which FinancialAccount and BalanceType to pull funds from.
type V2MoneyManagementOutboundPaymentCreateFromParams struct {
	// Describes the FinancialAmount's currency drawn from.
	Currency *string `form:"currency" json:"currency"`
	// The FinancialAccount that funds were pulled from.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
}

// Details about the notification settings for the OutboundPayment recipient.
type V2MoneyManagementOutboundPaymentCreateRecipientNotificationParams struct {
	// Closed Enum. Configuration option to enable or disable notifications to recipients.
	// Do not send notifications when `setting` is `none`. Use the account setting when `setting` is `configured` or omitted.
	Setting *string `form:"setting" json:"setting"`
}

// ACH-specific network options.
type V2MoneyManagementOutboundPaymentCreateToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHParams struct {
	// Open Enum. ACH submission timing.
	Submission *string `form:"submission" json:"submission,omitempty"`
	// The transaction purpose for this ACH payment.
	TransactionPurpose *string `form:"transaction_purpose" json:"transaction_purpose,omitempty"`
}

// Per-network configuration options.
type V2MoneyManagementOutboundPaymentCreateToPayoutMethodOptionsBankAccountPreferredNetworkOptionsParams struct {
	// ACH-specific network options.
	ACH *V2MoneyManagementOutboundPaymentCreateToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHParams `form:"ach" json:"ach,omitempty"`
}

// Options for bank account payout methods.
type V2MoneyManagementOutboundPaymentCreateToPayoutMethodOptionsBankAccountParams struct {
	// Per-network configuration options.
	PreferredNetworkOptions *V2MoneyManagementOutboundPaymentCreateToPayoutMethodOptionsBankAccountPreferredNetworkOptionsParams `form:"preferred_network_options" json:"preferred_network_options,omitempty"`
	// The preferred networks to use for this OutboundPayment.
	PreferredNetworks []*string `form:"preferred_networks" json:"preferred_networks,omitempty"`
}

// Payout method options for the OutboundPayment.
type V2MoneyManagementOutboundPaymentCreateToPayoutMethodOptionsParams struct {
	// Options for bank account payout methods.
	BankAccount *V2MoneyManagementOutboundPaymentCreateToPayoutMethodOptionsBankAccountParams `form:"bank_account" json:"bank_account,omitempty"`
}

// To which payout method to send the OutboundPayment.
type V2MoneyManagementOutboundPaymentCreateToParams struct {
	// Describes the currency to send to the recipient.
	// If included, this currency must match a currency supported by the destination.
	// Can be omitted in the following cases:
	// - destination only supports one currency
	// - destination supports multiple currencies and one matches `from.currency`
	// - destination supports multiple currencies and one matches `amount.currency`
	// When both `from.currency` and `amount.currency` are supported, Stripe selects `from.currency` to minimize currency conversion.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// The payout method which the OutboundPayment uses to send payout.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
	// Payout method options for the OutboundPayment.
	PayoutMethodOptions *V2MoneyManagementOutboundPaymentCreateToPayoutMethodOptionsParams `form:"payout_method_options" json:"payout_method_options,omitempty"`
	// To which account the OutboundPayment is sent.
	Recipient *string `form:"recipient" json:"recipient"`
}

// Creates an OutboundPayment.
type V2MoneyManagementOutboundPaymentCreateParams struct {
	Params `form:"*"`
	// The requested amount and currency for the OutboundPayment.
	Amount *Amount `form:"amount" json:"amount"`
	// Delivery options to be used to send the OutboundPayment.
	DeliveryOptions *V2MoneyManagementOutboundPaymentCreateDeliveryOptionsParams `form:"delivery_options" json:"delivery_options,omitempty"`
	// An arbitrary string attached to the OutboundPayment. Often useful for displaying to users.
	Description *string `form:"description" json:"description,omitempty"`
	// From which FinancialAccount and BalanceType to pull funds from.
	From *V2MoneyManagementOutboundPaymentCreateFromParams `form:"from" json:"from"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The quote for this OutboundPayment. Only required for countries with regulatory mandates to display fee estimates before OutboundPayment creation.
	OutboundPaymentQuote *string `form:"outbound_payment_quote" json:"outbound_payment_quote,omitempty"`
	// The purpose of the OutboundPayment.
	Purpose *string `form:"purpose" json:"purpose,omitempty"`
	// Details about the notification settings for the OutboundPayment recipient.
	RecipientNotification *V2MoneyManagementOutboundPaymentCreateRecipientNotificationParams `form:"recipient_notification" json:"recipient_notification,omitempty"`
	// The ID of the RecipientVerification for this OutboundPayment. Required where recipient-name verification is mandated before creating an OutboundPayment.
	RecipientVerification *string `form:"recipient_verification" json:"recipient_verification,omitempty"`
	// The description that appears on the receiving end for an OutboundPayment (for example, on a bank statement). Must be between 3 and 22 characters long for most destinations (500 for FinancialAccount destinations), and not contain profanity.
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// To which payout method to send the OutboundPayment.
	To *V2MoneyManagementOutboundPaymentCreateToParams `form:"to" json:"to"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementOutboundPaymentCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves the details of an existing OutboundPayment by passing the unique OutboundPayment ID from either the OutboundPayment create or list response.
type V2MoneyManagementOutboundPaymentRetrieveParams struct {
	Params `form:"*"`
}
