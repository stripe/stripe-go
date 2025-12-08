//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The balance to which funds are transferred.
type BalanceTransferDestinationBalanceParams struct {
	// Destination balance type to push funds into for the Balance Transfer.
	Type *string `form:"type"`
}
type BalanceTransferSourceBalanceAllocatedFundsParams struct {
	// The charge ID that the funds are originally sourced from. Required if `type` is `charge`.
	Charge *string `form:"charge"`
	// The type of object that the funds are originally sourced from. One of `charge`.
	Type *string `form:"type"`
}

// The balance from which funds are transferred, including details specific to the balance you choose.
type BalanceTransferSourceBalanceParams struct {
	AllocatedFunds *BalanceTransferSourceBalanceAllocatedFundsParams `form:"allocated_funds"`
	// Source balance type to pull funds from for the Balance Transfer.
	Type *string `form:"type"`
}

// Creates a balance transfer. For Issuing use cases, funds will be debited immediately from the source balance and credited to the destination balance immediately (if your account is based in the US) or next-business-day (if your account is based in the EU). For Segregated Separate Charges and Transfers use cases, funds will be debited immediately from the source balance and credited immediately to the destination balance.
type BalanceTransferParams struct {
	Params `form:"*"`
	// A positive integer representing how much to transfer in the smallest currency unit.
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The balance to which funds are transferred.
	DestinationBalance *BalanceTransferDestinationBalanceParams `form:"destination_balance"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The balance from which funds are transferred, including details specific to the balance you choose.
	SourceBalance *BalanceTransferSourceBalanceParams `form:"source_balance"`
}

// AddExpand appends a new field to expand.
func (p *BalanceTransferParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *BalanceTransferParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The balance to which funds are transferred.
type BalanceTransferCreateDestinationBalanceParams struct {
	// Destination balance type to push funds into for the Balance Transfer.
	Type *string `form:"type"`
}
type BalanceTransferCreateSourceBalanceAllocatedFundsParams struct {
	// The charge ID that the funds are originally sourced from. Required if `type` is `charge`.
	Charge *string `form:"charge"`
	// The type of object that the funds are originally sourced from. One of `charge`.
	Type *string `form:"type"`
}

// The balance from which funds are transferred, including details specific to the balance you choose.
type BalanceTransferCreateSourceBalanceParams struct {
	AllocatedFunds *BalanceTransferCreateSourceBalanceAllocatedFundsParams `form:"allocated_funds"`
	// Source balance type to pull funds from for the Balance Transfer.
	Type *string `form:"type"`
}

// Creates a balance transfer. For Issuing use cases, funds will be debited immediately from the source balance and credited to the destination balance immediately (if your account is based in the US) or next-business-day (if your account is based in the EU). For Segregated Separate Charges and Transfers use cases, funds will be debited immediately from the source balance and credited immediately to the destination balance.
type BalanceTransferCreateParams struct {
	Params `form:"*"`
	// A positive integer representing how much to transfer in the smallest currency unit.
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The balance to which funds are transferred.
	DestinationBalance *BalanceTransferCreateDestinationBalanceParams `form:"destination_balance"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The balance from which funds are transferred, including details specific to the balance you choose.
	SourceBalance *BalanceTransferCreateSourceBalanceParams `form:"source_balance"`
}

// AddExpand appends a new field to expand.
func (p *BalanceTransferCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *BalanceTransferCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

type BalanceTransferDestinationBalanceIssuing struct {
	// Identifier for the balance_transaction that increased the destination balance.
	BalanceTransaction *BalanceTransaction `json:"balance_transaction"`
}
type BalanceTransferDestinationBalancePayments struct {
	// Identifier for the balance_transaction that increased the destination balance.
	BalanceTransaction *BalanceTransaction `json:"balance_transaction"`
}

// The balance that funds were transferred to.
type BalanceTransferDestinationBalance struct {
	Issuing  *BalanceTransferDestinationBalanceIssuing  `json:"issuing"`
	Payments *BalanceTransferDestinationBalancePayments `json:"payments"`
	// Destination balance type to adjust for the Balance Transfer. One of `payments`, `issuing`, or `allocated_funds`.
	Type string `json:"type"`
}
type BalanceTransferSourceBalanceIssuing struct {
	// Identifier for the balance_transaction that decreased the source balance.
	BalanceTransaction *BalanceTransaction `json:"balance_transaction"`
}
type BalanceTransferSourceBalancePayments struct {
	// Identifier for the balance_transaction that decreased the source balance.
	BalanceTransaction *BalanceTransaction `json:"balance_transaction"`
	// The payments balance type that this BalanceTransfer pulled funds from. One of `card`, `fpx`, or `bank_account`.
	SourceType string `json:"source_type"`
}

// The balance that funds were transferred from. One of `card`, `fpx`, or `bank_account`.
type BalanceTransferSourceBalance struct {
	Issuing  *BalanceTransferSourceBalanceIssuing  `json:"issuing"`
	Payments *BalanceTransferSourceBalancePayments `json:"payments"`
	// Source balance type to adjust for the Balance Transfer. One of `payments`, `issuing`, or `allocated_funds`.
	Type string `json:"type"`
}

// Balance transfers represent funds moving between balance types on your Stripe account.
// They currently support moving funds between your Stripe balance and your [Issuing](https://docs.stripe.com/issuing) balance and between your [Allocated Funds](https://docs.stripe.com/connect/funds-segregation) balance and your Stripe balance.
type BalanceTransfer struct {
	APIResource
	// A positive integer representing how much was transferred in the smallest currency unit.
	Amount int64 `json:"amount"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The balance that funds were transferred to.
	DestinationBalance *BalanceTransferDestinationBalance `json:"destination_balance"`
	// A [hosted transaction receipt](https://docs.stripe.com/treasury/moving-money/regulatory-receipts) URL that is provided when money movement is considered regulated under Stripe's money transmission licenses.
	HostedRegulatoryReceiptURL string `json:"hosted_regulatory_receipt_url"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The balance that funds were transferred from. One of `card`, `fpx`, or `bank_account`.
	SourceBalance *BalanceTransferSourceBalance `json:"source_balance"`
}
