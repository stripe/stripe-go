//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Object containing details about where the funds will originate from.
type V2MoneyManagementInboundTransferFromParams struct {
	// An optional currency field used to specify which currency is debited from the Payment Method.
	// Since many Payment Methods support only one currency, this field is optional.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// ID of the Payment Method using which IBT will be made.
	PaymentMethod *string `form:"payment_method" json:"payment_method"`
}

// Object containing details about where the funds will land.
type V2MoneyManagementInboundTransferToParams struct {
	// The currency in which funds will land in.
	Currency *string `form:"currency" json:"currency"`
	// The FinancialAccount that funds will land in.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
}

// InboundTransfers APIs are used to create, retrieve or list InboundTransfers.
type V2MoneyManagementInboundTransferParams struct {
	Params `form:"*"`
	// The amount, in specified currency, by which the FinancialAccount balance will increase due to the InboundTransfer.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
	// An optional, freeform description field intended to store metadata.
	Description *string `form:"description" json:"description,omitempty"`
	// Object containing details about where the funds will originate from.
	From *V2MoneyManagementInboundTransferFromParams `form:"from" json:"from,omitempty"`
	// Object containing details about where the funds will land.
	To *V2MoneyManagementInboundTransferToParams `form:"to" json:"to,omitempty"`
}

// Retrieves a list of InboundTransfers.
type V2MoneyManagementInboundTransferListParams struct {
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
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Object containing details about where the funds will originate from.
type V2MoneyManagementInboundTransferCreateFromParams struct {
	// An optional currency field used to specify which currency is debited from the Payment Method.
	// Since many Payment Methods support only one currency, this field is optional.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// ID of the Payment Method using which IBT will be made.
	PaymentMethod *string `form:"payment_method" json:"payment_method"`
}

// Object containing details about where the funds will land.
type V2MoneyManagementInboundTransferCreateToParams struct {
	// The currency in which funds will land in.
	Currency *string `form:"currency" json:"currency"`
	// The FinancialAccount that funds will land in.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
}

// InboundTransfers APIs are used to create, retrieve or list InboundTransfers.
type V2MoneyManagementInboundTransferCreateParams struct {
	Params `form:"*"`
	// The amount, in specified currency, by which the FinancialAccount balance will increase due to the InboundTransfer.
	Amount *Amount `form:"amount" json:"amount"`
	// An optional, freeform description field intended to store metadata.
	Description *string `form:"description" json:"description,omitempty"`
	// Object containing details about where the funds will originate from.
	From *V2MoneyManagementInboundTransferCreateFromParams `form:"from" json:"from"`
	// Object containing details about where the funds will land.
	To *V2MoneyManagementInboundTransferCreateToParams `form:"to" json:"to"`
}

// Retrieve an InboundTransfer by ID.
type V2MoneyManagementInboundTransferRetrieveParams struct {
	Params `form:"*"`
}
