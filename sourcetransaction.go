//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List source transactions for a given source.
type SourceTransactionListParams struct {
	ListParams `form:"*"`
	Source     *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *SourceTransactionListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type SourceTransactionACHCreditTransfer struct {
	// Customer data associated with the transfer.
	CustomerData string `json:"customer_data,omitempty"`
	// Bank account fingerprint associated with the transfer.
	Fingerprint string `json:"fingerprint,omitempty"`
	// Last 4 digits of the account number associated with the transfer.
	Last4 string `json:"last4,omitempty"`
	// Routing number associated with the transfer.
	RoutingNumber string `json:"routing_number,omitempty"`
}
type SourceTransactionCHFCreditTransfer struct {
	// Reference associated with the transfer.
	Reference string `json:"reference,omitempty"`
	// Sender's country address.
	SenderAddressCountry string `json:"sender_address_country,omitempty"`
	// Sender's line 1 address.
	SenderAddressLine1 string `json:"sender_address_line1,omitempty"`
	// Sender's bank account IBAN.
	SenderIBAN string `json:"sender_iban,omitempty"`
	// Sender's name.
	SenderName string `json:"sender_name,omitempty"`
}
type SourceTransactionGBPCreditTransfer struct {
	// Bank account fingerprint associated with the Stripe owned bank account receiving the transfer.
	Fingerprint string `json:"fingerprint,omitempty"`
	// The credit transfer rails the sender used to push this transfer. The possible rails are: Faster Payments, BACS, CHAPS, and wire transfers. Currently only Faster Payments is supported.
	FundingMethod string `json:"funding_method,omitempty"`
	// Last 4 digits of sender account number associated with the transfer.
	Last4 string `json:"last4,omitempty"`
	// Sender entered arbitrary information about the transfer.
	Reference string `json:"reference,omitempty"`
	// Sender account number associated with the transfer.
	SenderAccountNumber string `json:"sender_account_number,omitempty"`
	// Sender name associated with the transfer.
	SenderName string `json:"sender_name,omitempty"`
	// Sender sort code associated with the transfer.
	SenderSortCode string `json:"sender_sort_code,omitempty"`
}
type SourceTransactionPaperCheck struct {
	// Time at which the deposited funds will be available for use. Measured in seconds since the Unix epoch.
	AvailableAt string `json:"available_at,omitempty"`
	// Comma-separated list of invoice IDs associated with the paper check.
	Invoices string `json:"invoices,omitempty"`
}
type SourceTransactionSEPACreditTransfer struct {
	// Reference associated with the transfer.
	Reference string `json:"reference,omitempty"`
	// Sender's bank account IBAN.
	SenderIBAN string `json:"sender_iban,omitempty"`
	// Sender's name.
	SenderName string `json:"sender_name,omitempty"`
}

// Some payment methods have no required amount that a customer must send.
// Customers can be instructed to send any amount, and it can be made up of
// multiple transactions. As such, sources can have multiple associated
// transactions.
type SourceTransaction struct {
	ACHCreditTransfer *SourceTransactionACHCreditTransfer `json:"ach_credit_transfer,omitempty"`
	// A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the amount your customer has pushed to the receiver.
	Amount            int64                               `json:"amount"`
	CHFCreditTransfer *SourceTransactionCHFCreditTransfer `json:"chf_credit_transfer,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency          Currency                            `json:"currency"`
	GBPCreditTransfer *SourceTransactionGBPCreditTransfer `json:"gbp_credit_transfer,omitempty"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object             string                               `json:"object"`
	PaperCheck         *SourceTransactionPaperCheck         `json:"paper_check,omitempty"`
	SEPACreditTransfer *SourceTransactionSEPACreditTransfer `json:"sepa_credit_transfer,omitempty"`
	// The ID of the source this transaction is attached to.
	Source string `json:"source"`
	// The status of the transaction, one of `succeeded`, `pending`, or `failed`.
	Status string `json:"status"`
	// The type of source this transaction is attached to.
	Type string `json:"type"`
}

// SourceTransactionList is a list of SourceTransactions as retrieved from a list endpoint.
type SourceTransactionList struct {
	APIResource
	ListMeta
	Data []*SourceTransaction `json:"data"`
}
