//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The amount and currency of the SettlementAllocationIntentSplit.
type V2PaymentsSettlementAllocationIntentsSplitAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// Create SettlementAllocationIntentSplit API.
type V2PaymentsSettlementAllocationIntentsSplitParams struct {
	Params `form:"*"`
	// The ID of the SettlementAllocationIntent this split belongs to.
	SettlementAllocationIntentID *string `form:"-" json:"-"` // Included in URL
	// The target account for settling the SettlementAllocationIntentSplit.
	Account *string `form:"account" json:"account,omitempty"`
	// The amount and currency of the SettlementAllocationIntentSplit.
	Amount *V2PaymentsSettlementAllocationIntentsSplitAmountParams `form:"amount" json:"amount,omitempty"`
	// Metadata associated with the SettlementAllocationIntentSplit.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The type of the fund transfer.
	Type *string `form:"type" json:"type,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2PaymentsSettlementAllocationIntentsSplitParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Cancel SettlementAllocationIntentSplit API.
type V2PaymentsSettlementAllocationIntentsSplitCancelParams struct {
	Params `form:"*"`
	// The SettlementAllocationIntent this split belongs to.
	SettlementAllocationIntentID *string `form:"-" json:"-"` // Included in URL
}

// The amount and currency of the SettlementAllocationIntentSplit.
type V2PaymentsSettlementAllocationIntentsSplitCreateAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// Create SettlementAllocationIntentSplit API.
type V2PaymentsSettlementAllocationIntentsSplitCreateParams struct {
	Params `form:"*"`
	// The SettlementAllocationIntent this split belongs to.
	SettlementAllocationIntentID *string `form:"-" json:"-"` // Included in URL
	// The target account for settling the SettlementAllocationIntentSplit.
	Account *string `form:"account" json:"account"`
	// The amount and currency of the SettlementAllocationIntentSplit.
	Amount *V2PaymentsSettlementAllocationIntentsSplitCreateAmountParams `form:"amount" json:"amount"`
	// Metadata associated with the SettlementAllocationIntentSplit.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The type of the fund transfer.
	Type *string `form:"type" json:"type"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2PaymentsSettlementAllocationIntentsSplitCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieve SettlementAllocationIntentSplit API.
type V2PaymentsSettlementAllocationIntentsSplitRetrieveParams struct {
	Params `form:"*"`
	// The ID of the SettlementAllocationIntent this split belongs to.
	SettlementAllocationIntentID *string `form:"-" json:"-"` // Included in URL
}
