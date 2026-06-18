//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of PayoutIntents.
type V2MoneyManagementPayoutIntentListParams struct {
	Params `form:"*"`
	// Maximum number of objects to return. Defaults to 10. Maximum is 100.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// The FinancialAccount that funds are pulled from.
type V2MoneyManagementPayoutIntentFromParams struct {
	// The currency of the financial account.
	Currency *string `form:"currency" json:"currency"`
	// The FinancialAccount that funds are pulled from.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
}

// Details about the OutboundPayment notification settings for recipient. Only applicable to OutboundPayment.
type V2MoneyManagementPayoutIntentRecipientNotificationParams struct {
	// Closed Enum. Configuration option to enable or disable notifications to recipients.
	// Do not send notifications when setting is NONE. Default to account setting when setting is CONFIGURED or not set.
	Setting *string `form:"setting" json:"setting"`
}

// Scheduling options for the payout. If this is nil, we assume immediate execution.
type V2MoneyManagementPayoutIntentScheduleOptionsParams struct {
	// The date when the payout should be executed, in YYYY-MM-DD format.
	ExecuteOn *string `form:"execute_on" json:"execute_on,omitempty"`
}

// ACH-specific network options.
type V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHParams struct {
	// Open Enum. ACH submission timing.
	Submission *string `form:"submission" json:"submission,omitempty"`
	// The transaction purpose for this ACH payment.
	TransactionPurpose *string `form:"transaction_purpose" json:"transaction_purpose,omitempty"`
}

// Per-network configuration options.
type V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsParams struct {
	// ACH-specific network options.
	ACH *V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHParams `form:"ach" json:"ach,omitempty"`
}

// Options for bank account payout methods.
type V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountParams struct {
	// Per-network configuration options.
	PreferredNetworkOptions *V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsParams `form:"preferred_network_options" json:"preferred_network_options,omitempty"`
	// The preferred networks to use for this PayoutIntent.
	PreferredNetworks []*string `form:"preferred_networks" json:"preferred_networks"`
}

// Payout method options for the PayoutIntent.
type V2MoneyManagementPayoutIntentToPayoutMethodOptionsParams struct {
	// Options for bank account payout methods.
	BankAccount *V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountParams `form:"bank_account" json:"bank_account,omitempty"`
}

// To which payout method the payout is sent.
type V2MoneyManagementPayoutIntentToParams struct {
	// The currency to send to the recipient.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// The payout method ID. Optional for OutboundPayment if recipient has default payment method. Required for OutboundTransfer.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
	// Payout method options for the PayoutIntent.
	PayoutMethodOptions *V2MoneyManagementPayoutIntentToPayoutMethodOptionsParams `form:"payout_method_options" json:"payout_method_options,omitempty"`
	// The recipient ID. Only relevant for OutboundPayment.
	Recipient *string `form:"recipient" json:"recipient,omitempty"`
}

// Creates a PayoutIntent.
type V2MoneyManagementPayoutIntentParams struct {
	Params `form:"*"`
	// The monetary amount to be sent.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
	// An arbitrary string attached to the PayoutIntent. Often useful for displaying to users.
	Description *string `form:"description" json:"description,omitempty"`
	// The FinancialAccount that funds are pulled from.
	From *V2MoneyManagementPayoutIntentFromParams `form:"from" json:"from,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Details about the OutboundPayment notification settings for recipient. Only applicable to OutboundPayment.
	RecipientNotification *V2MoneyManagementPayoutIntentRecipientNotificationParams `form:"recipient_notification" json:"recipient_notification,omitempty"`
	// Scheduling options for the payout. If this is nil, we assume immediate execution.
	ScheduleOptions *V2MoneyManagementPayoutIntentScheduleOptionsParams `form:"schedule_options" json:"schedule_options,omitempty"`
	// The description that appears on the receiving end for the payout (for example, on a bank statement).
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// To which payout method the payout is sent.
	To *V2MoneyManagementPayoutIntentToParams `form:"to" json:"to,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementPayoutIntentParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Cancels a PayoutIntent. Only pending PayoutIntents or processing PayoutIntents with cancelable OutboundPayment/Transfer can be canceled.
type V2MoneyManagementPayoutIntentCancelParams struct {
	Params `form:"*"`
}

// The FinancialAccount that funds are pulled from.
type V2MoneyManagementPayoutIntentCreateFromParams struct {
	// The currency of the financial account.
	Currency *string `form:"currency" json:"currency"`
	// The FinancialAccount that funds are pulled from.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
}

// Details about the OutboundPayment notification settings for recipient. Only applicable to OutboundPayment.
type V2MoneyManagementPayoutIntentCreateRecipientNotificationParams struct {
	// Closed Enum. Configuration option to enable or disable notifications to recipients.
	// Do not send notifications when setting is NONE. Default to account setting when setting is CONFIGURED or not set.
	Setting *string `form:"setting" json:"setting"`
}

// Scheduling options for the payout. If this is nil, we assume immediate execution.
type V2MoneyManagementPayoutIntentCreateScheduleOptionsParams struct {
	// The date when the payout should be executed, in YYYY-MM-DD format.
	ExecuteOn *string `form:"execute_on" json:"execute_on,omitempty"`
}

// ACH-specific network options.
type V2MoneyManagementPayoutIntentCreateToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHParams struct {
	// Open Enum. ACH submission timing.
	Submission *string `form:"submission" json:"submission,omitempty"`
	// The transaction purpose for this ACH payment.
	TransactionPurpose *string `form:"transaction_purpose" json:"transaction_purpose,omitempty"`
}

// Per-network configuration options.
type V2MoneyManagementPayoutIntentCreateToPayoutMethodOptionsBankAccountPreferredNetworkOptionsParams struct {
	// ACH-specific network options.
	ACH *V2MoneyManagementPayoutIntentCreateToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHParams `form:"ach" json:"ach,omitempty"`
}

// Options for bank account payout methods.
type V2MoneyManagementPayoutIntentCreateToPayoutMethodOptionsBankAccountParams struct {
	// Per-network configuration options.
	PreferredNetworkOptions *V2MoneyManagementPayoutIntentCreateToPayoutMethodOptionsBankAccountPreferredNetworkOptionsParams `form:"preferred_network_options" json:"preferred_network_options,omitempty"`
	// The preferred networks to use for this PayoutIntent.
	PreferredNetworks []*string `form:"preferred_networks" json:"preferred_networks"`
}

// Payout method options for the PayoutIntent.
type V2MoneyManagementPayoutIntentCreateToPayoutMethodOptionsParams struct {
	// Options for bank account payout methods.
	BankAccount *V2MoneyManagementPayoutIntentCreateToPayoutMethodOptionsBankAccountParams `form:"bank_account" json:"bank_account,omitempty"`
}

// To which payout method the payout is sent.
type V2MoneyManagementPayoutIntentCreateToParams struct {
	// The currency to send to the recipient.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// The payout method ID. Optional for OutboundPayment if recipient has default payment method. Required for OutboundTransfer.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
	// Payout method options for the PayoutIntent.
	PayoutMethodOptions *V2MoneyManagementPayoutIntentCreateToPayoutMethodOptionsParams `form:"payout_method_options" json:"payout_method_options,omitempty"`
	// The recipient ID. Only relevant for OutboundPayment.
	Recipient *string `form:"recipient" json:"recipient,omitempty"`
}

// Creates a PayoutIntent.
type V2MoneyManagementPayoutIntentCreateParams struct {
	Params `form:"*"`
	// The monetary amount to be sent.
	Amount *Amount `form:"amount" json:"amount"`
	// An arbitrary string attached to the PayoutIntent. Often useful for displaying to users.
	Description *string `form:"description" json:"description,omitempty"`
	// The FinancialAccount that funds are pulled from.
	From *V2MoneyManagementPayoutIntentCreateFromParams `form:"from" json:"from"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Details about the OutboundPayment notification settings for recipient. Only applicable to OutboundPayment.
	RecipientNotification *V2MoneyManagementPayoutIntentCreateRecipientNotificationParams `form:"recipient_notification" json:"recipient_notification,omitempty"`
	// Scheduling options for the payout. If this is nil, we assume immediate execution.
	ScheduleOptions *V2MoneyManagementPayoutIntentCreateScheduleOptionsParams `form:"schedule_options" json:"schedule_options,omitempty"`
	// The description that appears on the receiving end for the payout (for example, on a bank statement).
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// To which payout method the payout is sent.
	To *V2MoneyManagementPayoutIntentCreateToParams `form:"to" json:"to"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementPayoutIntentCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves the details of an existing PayoutIntent by passing the unique PayoutIntent ID.
type V2MoneyManagementPayoutIntentRetrieveParams struct {
	Params `form:"*"`
}

// From which FinancialAccount to pull funds.
type V2MoneyManagementPayoutIntentUpdateFromParams struct {
	// The currency of the financial account.
	Currency *string `form:"currency" json:"currency"`
	// The FinancialAccount that funds are pulled from.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
}

// Details about the OutboundPayment notification settings for recipient. Only applicable to OutboundPayment.
type V2MoneyManagementPayoutIntentUpdateRecipientNotificationParams struct {
	// Closed Enum. Configuration option to enable or disable notifications to recipients.
	// Do not send notifications when setting is NONE. Default to account setting when setting is CONFIGURED or not set.
	Setting *string `form:"setting" json:"setting"`
}

// Scheduling options for the payout. If this is nil, we assume immediate execution.
type V2MoneyManagementPayoutIntentUpdateScheduleOptionsParams struct {
	// The date when the payout should be executed, in YYYY-MM-DD format.
	ExecuteOn *string `form:"execute_on" json:"execute_on,omitempty"`
}

// ACH-specific network options.
type V2MoneyManagementPayoutIntentUpdateToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHParams struct {
	// Open Enum. ACH submission timing.
	Submission *string `form:"submission" json:"submission,omitempty"`
	// The transaction purpose for this ACH payment.
	TransactionPurpose *string `form:"transaction_purpose" json:"transaction_purpose,omitempty"`
}

// Per-network configuration options.
type V2MoneyManagementPayoutIntentUpdateToPayoutMethodOptionsBankAccountPreferredNetworkOptionsParams struct {
	// ACH-specific network options.
	ACH *V2MoneyManagementPayoutIntentUpdateToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHParams `form:"ach" json:"ach,omitempty"`
}

// Options for bank account payout methods.
type V2MoneyManagementPayoutIntentUpdateToPayoutMethodOptionsBankAccountParams struct {
	// Per-network configuration options.
	PreferredNetworkOptions *V2MoneyManagementPayoutIntentUpdateToPayoutMethodOptionsBankAccountPreferredNetworkOptionsParams `form:"preferred_network_options" json:"preferred_network_options,omitempty"`
	// The preferred networks to use for this PayoutIntent.
	PreferredNetworks []*string `form:"preferred_networks" json:"preferred_networks"`
}

// Payout method options for the PayoutIntent.
type V2MoneyManagementPayoutIntentUpdateToPayoutMethodOptionsParams struct {
	// Options for bank account payout methods.
	BankAccount *V2MoneyManagementPayoutIntentUpdateToPayoutMethodOptionsBankAccountParams `form:"bank_account" json:"bank_account,omitempty"`
}

// To which payout method the payout is sent.
type V2MoneyManagementPayoutIntentUpdateToParams struct {
	// The currency to send to the recipient.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// The payout method ID. Optional for OutboundPayment if recipient has default payment method. Required for OutboundTransfer.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
	// Payout method options for the PayoutIntent.
	PayoutMethodOptions *V2MoneyManagementPayoutIntentUpdateToPayoutMethodOptionsParams `form:"payout_method_options" json:"payout_method_options,omitempty"`
	// The recipient ID. Only relevant for OutboundPayment.
	Recipient *string `form:"recipient" json:"recipient,omitempty"`
}

// Updates a PayoutIntent. Only pending or requires_action PayoutIntents that are editable can be updated.
type V2MoneyManagementPayoutIntentUpdateParams struct {
	Params `form:"*"`
	// The monetary amount to be sent.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
	// An arbitrary string attached to the PayoutIntent. Often useful for displaying to users.
	Description *string `form:"description" json:"description,omitempty"`
	// From which FinancialAccount to pull funds.
	From *V2MoneyManagementPayoutIntentUpdateFromParams `form:"from" json:"from,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Details about the OutboundPayment notification settings for recipient. Only applicable to OutboundPayment.
	RecipientNotification *V2MoneyManagementPayoutIntentUpdateRecipientNotificationParams `form:"recipient_notification" json:"recipient_notification,omitempty"`
	// Scheduling options for the payout. If this is nil, we assume immediate execution.
	ScheduleOptions *V2MoneyManagementPayoutIntentUpdateScheduleOptionsParams `form:"schedule_options" json:"schedule_options,omitempty"`
	// The description that appears on the receiving end for the payout (for example, on a bank statement).
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// To which payout method the payout is sent.
	To *V2MoneyManagementPayoutIntentUpdateToParams `form:"to" json:"to,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementPayoutIntentUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}
