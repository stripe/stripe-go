//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// By default, you can see the 10 most recent refunds stored directly on the application fee object, but you can also retrieve details about a specific refund stored on the application fee.
type FeeRefundParams struct {
	Params `form:"*"`
	ID     *string `form:"-"` // Included in URL
	Fee    *string `form:"-"` // Included in URL
	// A positive integer, in _cents (or local equivalent)_, representing how much of this fee to refund. Can refund only up to the remaining unrefunded amount of the fee.
	Amount *int64 `form:"amount"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
}

// AddExpand appends a new field to expand.
func (p *FeeRefundParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *FeeRefundParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// You can see a list of the refunds belonging to a specific application fee. Note that the 10 most recent refunds are always available by default on the application fee object. If you need more than those 10, you can use this API method and the limit and starting_after parameters to page through additional refunds.
type FeeRefundListParams struct {
	ListParams `form:"*"`
	ID         *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *FeeRefundListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// By default, you can see the 10 most recent refunds stored directly on the application fee object, but you can also retrieve details about a specific refund stored on the application fee.
type FeeRefundRetrieveParams struct {
	Params `form:"*"`
	Fee    *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *FeeRefundRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Updates the specified application fee refund by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// This request only accepts metadata as an argument.
type FeeRefundUpdateParams struct {
	Params `form:"*"`
	Fee    *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
}

// AddExpand appends a new field to expand.
func (p *FeeRefundUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *FeeRefundUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Refunds an application fee that has previously been collected but not yet refunded.
// Funds will be refunded to the Stripe account from which the fee was originally collected.
//
// You can optionally refund only part of an application fee.
// You can do so multiple times, until the entire fee has been refunded.
//
// Once entirely refunded, an application fee can't be refunded again.
// This method will raise an error when called on an already-refunded application fee,
// or when trying to refund more money than is left on an application fee.
type FeeRefundCreateParams struct {
	Params `form:"*"`
	ID     *string `form:"-"` // Included in URL
	// A positive integer, in _cents (or local equivalent)_, representing how much of this fee to refund. Can refund only up to the remaining unrefunded amount of the fee.
	Amount *int64 `form:"amount"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
}

// AddExpand appends a new field to expand.
func (p *FeeRefundCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *FeeRefundCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// `Application Fee Refund` objects allow you to refund an application fee that
// has previously been created but not yet refunded. Funds will be refunded to
// the Stripe account from which the fee was originally collected.
//
// Related guide: [Refunding application fees](https://stripe.com/docs/connect/destination-charges#refunding-app-fee)
type FeeRefund struct {
	APIResource
	// Amount, in cents (or local equivalent).
	Amount int64 `json:"amount"`
	// Balance transaction that describes the impact on your account balance.
	BalanceTransaction *BalanceTransaction `json:"balance_transaction"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// ID of the application fee that was refunded.
	Fee *ApplicationFee `json:"fee"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}

// FeeRefundList is a list of FeeRefunds as retrieved from a list endpoint.
type FeeRefundList struct {
	APIResource
	ListMeta
	Data []*FeeRefund `json:"data"`
}

// UnmarshalJSON handles deserialization of a FeeRefund.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (f *FeeRefund) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		f.ID = id
		return nil
	}

	type feeRefund FeeRefund
	var v feeRefund
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*f = FeeRefund(v)
	return nil
}
