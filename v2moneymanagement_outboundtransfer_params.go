//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Returns a list of OutboundTransfers that match the provided filters.
type V2MoneyManagementOutboundTransferListParams struct {
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
	// Closed Enum. Only return OutboundTransfers with this status.
	Status []*string `form:"status,flat_array" json:"status,omitempty"`
}

// Delivery options to be used to send the OutboundTransfer.
type V2MoneyManagementOutboundTransferDeliveryOptionsParams struct {
	// Open Enum. Method for bank account.
	BankAccount *string `form:"bank_account" json:"bank_account,omitempty"`
}

// The FinancialAccount to pull funds from.
type V2MoneyManagementOutboundTransferFromParams struct {
	// Describes the FinancialAmount's currency drawn from.
	Currency *string `form:"currency" json:"currency"`
	// The FinancialAccount that funds were pulled from.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
}

// To which payout method to send the OutboundTransfer.
type V2MoneyManagementOutboundTransferToParams struct {
	// Describes the currency to send to the recipient.
	// If included, this currency must match a currency supported by the destination.
	// Can be omitted in the following cases:
	// - destination only supports one currency
	// - destination supports multiple currencies and one of the currencies matches the FA currency
	// - destination supports multiple currencies and one of the currencies matches the presentment currency
	// Note - when both FA currency and presentment currency are supported, we pick the FA currency to minimize FX.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// The payout method which the OutboundTransfer uses to send payout.
	PayoutMethod *string `form:"payout_method" json:"payout_method"`
}

// Creates an OutboundTransfer.
type V2MoneyManagementOutboundTransferParams struct {
	Params `form:"*"`
	// The "presentment amount" for the OutboundPayment.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
	// Delivery options to be used to send the OutboundTransfer.
	DeliveryOptions *V2MoneyManagementOutboundTransferDeliveryOptionsParams `form:"delivery_options" json:"delivery_options,omitempty"`
	// An arbitrary string attached to the OutboundTransfer. Often useful for displaying to users.
	Description *string `form:"description" json:"description,omitempty"`
	// The FinancialAccount to pull funds from.
	From *V2MoneyManagementOutboundTransferFromParams `form:"from" json:"from,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// To which payout method to send the OutboundTransfer.
	To *V2MoneyManagementOutboundTransferToParams `form:"to" json:"to,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementOutboundTransferParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Cancels an OutboundTransfer. Only processing OutboundTransfers can be canceled.
type V2MoneyManagementOutboundTransferCancelParams struct {
	Params `form:"*"`
}

// Delivery options to be used to send the OutboundTransfer.
type V2MoneyManagementOutboundTransferCreateDeliveryOptionsParams struct {
	// Open Enum. Method for bank account.
	BankAccount *string `form:"bank_account" json:"bank_account,omitempty"`
}

// The FinancialAccount to pull funds from.
type V2MoneyManagementOutboundTransferCreateFromParams struct {
	// Describes the FinancialAmount's currency drawn from.
	Currency *string `form:"currency" json:"currency"`
	// The FinancialAccount that funds were pulled from.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
}

// To which payout method to send the OutboundTransfer.
type V2MoneyManagementOutboundTransferCreateToParams struct {
	// Describes the currency to send to the recipient.
	// If included, this currency must match a currency supported by the destination.
	// Can be omitted in the following cases:
	// - destination only supports one currency
	// - destination supports multiple currencies and one of the currencies matches the FA currency
	// - destination supports multiple currencies and one of the currencies matches the presentment currency
	// Note - when both FA currency and presentment currency are supported, we pick the FA currency to minimize FX.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// The payout method which the OutboundTransfer uses to send payout.
	PayoutMethod *string `form:"payout_method" json:"payout_method"`
}

// Creates an OutboundTransfer.
type V2MoneyManagementOutboundTransferCreateParams struct {
	Params `form:"*"`
	// The "presentment amount" for the OutboundPayment.
	Amount *Amount `form:"amount" json:"amount"`
	// Delivery options to be used to send the OutboundTransfer.
	DeliveryOptions *V2MoneyManagementOutboundTransferCreateDeliveryOptionsParams `form:"delivery_options" json:"delivery_options,omitempty"`
	// An arbitrary string attached to the OutboundTransfer. Often useful for displaying to users.
	Description *string `form:"description" json:"description,omitempty"`
	// The FinancialAccount to pull funds from.
	From *V2MoneyManagementOutboundTransferCreateFromParams `form:"from" json:"from"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// To which payout method to send the OutboundTransfer.
	To *V2MoneyManagementOutboundTransferCreateToParams `form:"to" json:"to"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementOutboundTransferCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves the details of an existing OutboundTransfer by passing the unique OutboundTransfer ID from either the OutboundPayment create or list response.
type V2MoneyManagementOutboundTransferRetrieveParams struct {
	Params `form:"*"`
}
