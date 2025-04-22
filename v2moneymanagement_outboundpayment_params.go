//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Cancels an OutboundPayment. Only processing OutboundPayments can be canceled.
type V2MoneyManagementOutboundPaymentCancelParams struct {
	Params `form:"*"`
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
	// Do not send notifications when setting is NONE. Default to account setting when setting is CONFIGURED or not set.
	Setting *string `form:"setting" json:"setting"`
}

// To which payout method to send the OutboundPayment.
type V2MoneyManagementOutboundPaymentToParams struct {
	// Describes the currency to send to the recipient.
	// If included, this currency must match a currency supported by the destination.
	// Can be omitted in the following cases:
	// - destination only supports one currency
	// - destination supports multiple currencies and one of the currencies matches the FA currency
	// - destination supports multiple currencies and one of the currencies matches the presentment currency
	// Note - when both FA currency and presentment currency are supported, we pick the FA currency to minimize FX.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// The payout method which the OutboundPayment uses to send payout.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
	// To which account the OutboundPayment is sent.
	Recipient *string `form:"recipient" json:"recipient"`
}

// Creates an OutboundPayment.
type V2MoneyManagementOutboundPaymentParams struct {
	Params `form:"*"`
	// The "presentment amount" to be sent to the recipient.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
	// Delivery options to be used to send the OutboundPayment.
	DeliveryOptions *V2MoneyManagementOutboundPaymentDeliveryOptionsParams `form:"delivery_options" json:"delivery_options,omitempty"`
	// An arbitrary string attached to the OutboundPayment. Often useful for displaying to users. The description can not be longer than 100 characters and can only contain basic Latin characters and spaces. The following special characters are not allowed: <>\'"* .
	Description *string `form:"description" json:"description,omitempty"`
	// From which FinancialAccount and BalanceType to pull funds from.
	From *V2MoneyManagementOutboundPaymentFromParams `form:"from" json:"from,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The quote for this OutboundPayment. Only required for countries with regulatory mandates to display fee estimates before OutboundPayment creation.
	OutboundPaymentQuote *string `form:"outbound_payment_quote" json:"outbound_payment_quote,omitempty"`
	// Details about the notification settings for the OutboundPayment recipient.
	RecipientNotification *V2MoneyManagementOutboundPaymentRecipientNotificationParams `form:"recipient_notification" json:"recipient_notification,omitempty"`
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

// Returns a list of OutboundPayments that match the provided filters.
type V2MoneyManagementOutboundPaymentListParams struct {
	Params `form:"*"`
	// Filter for objects created at the specified timestamp.
	// Must be an RFC 3339 date & time value, for example: 2022-09-18T13:22:00Z.
	Created *time.Time `form:"created" json:"created,omitempty"`
	// Filter for objects created after the specified timestamp.
	// Must be an RFC 3339 date & time value, for example: 2022-09-18T13:22:00Z.
	CreatedGt *time.Time `form:"created_gt" json:"created_gt,omitempty"`
	// Filter for objects created on or after the specified timestamp.
	// Must be an RFC 3339 date & time value, for example: 2022-09-18T13:22:00Z.
	CreatedGTE *time.Time `form:"created_gte" json:"created_gte,omitempty"`
	// Filter for objects created before the specified timestamp.
	// Must be an RFC 3339 date & time value, for example: 2022-09-18T13:22:00Z.
	CreatedLT *time.Time `form:"created_lt" json:"created_lt,omitempty"`
	// Filter for objects created on or before the specified timestamp.
	// Must be an RFC 3339 date & time value, for example: 2022-09-18T13:22:00Z.
	CreatedLte *time.Time `form:"created_lte" json:"created_lte,omitempty"`
	// The maximum number of results to return.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Only return OutboundPayments sent to this recipient.
	Recipient *string `form:"recipient" json:"recipient,omitempty"`
	// Closed Enum. Only return OutboundPayments with this status.
	Status []*string `form:"status" json:"status,omitempty"`
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
	// Do not send notifications when setting is NONE. Default to account setting when setting is CONFIGURED or not set.
	Setting *string `form:"setting" json:"setting"`
}

// To which payout method to send the OutboundPayment.
type V2MoneyManagementOutboundPaymentCreateToParams struct {
	// Describes the currency to send to the recipient.
	// If included, this currency must match a currency supported by the destination.
	// Can be omitted in the following cases:
	// - destination only supports one currency
	// - destination supports multiple currencies and one of the currencies matches the FA currency
	// - destination supports multiple currencies and one of the currencies matches the presentment currency
	// Note - when both FA currency and presentment currency are supported, we pick the FA currency to minimize FX.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// The payout method which the OutboundPayment uses to send payout.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
	// To which account the OutboundPayment is sent.
	Recipient *string `form:"recipient" json:"recipient"`
}

// Creates an OutboundPayment.
type V2MoneyManagementOutboundPaymentCreateParams struct {
	Params `form:"*"`
	// The "presentment amount" to be sent to the recipient.
	Amount *Amount `form:"amount" json:"amount"`
	// Delivery options to be used to send the OutboundPayment.
	DeliveryOptions *V2MoneyManagementOutboundPaymentCreateDeliveryOptionsParams `form:"delivery_options" json:"delivery_options,omitempty"`
	// An arbitrary string attached to the OutboundPayment. Often useful for displaying to users. The description can not be longer than 100 characters and can only contain basic Latin characters and spaces. The following special characters are not allowed: <>\'"* .
	Description *string `form:"description" json:"description,omitempty"`
	// From which FinancialAccount and BalanceType to pull funds from.
	From *V2MoneyManagementOutboundPaymentCreateFromParams `form:"from" json:"from"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The quote for this OutboundPayment. Only required for countries with regulatory mandates to display fee estimates before OutboundPayment creation.
	OutboundPaymentQuote *string `form:"outbound_payment_quote" json:"outbound_payment_quote,omitempty"`
	// Details about the notification settings for the OutboundPayment recipient.
	RecipientNotification *V2MoneyManagementOutboundPaymentCreateRecipientNotificationParams `form:"recipient_notification" json:"recipient_notification,omitempty"`
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
