//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The amount and currency of the SettlementAllocationIntent.
type V2PaymentsSettlementAllocationIntentAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// Create SettlementAllocationIntent API.
type V2PaymentsSettlementAllocationIntentParams struct {
	Params `form:"*"`
	// The amount and currency of the SettlementAllocationIntent.
	Amount *V2PaymentsSettlementAllocationIntentAmountParams `form:"amount" json:"amount,omitempty"`
	// Expected date when we expect to receive the funds.
	ExpectedSettlementDate *time.Time `form:"expected_settlement_date" json:"expected_settlement_date,omitempty"`
	// FinancialAccount where the funds are expected to land / FinancialAccount to map this SettlementAllocationIntent to.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// Metadata associated with the SettlementAllocationIntent.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Reference for the settlement intent . Max 255 characters . The reference used by PSP to send funds to Stripe .
	Reference *string `form:"reference" json:"reference,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2PaymentsSettlementAllocationIntentParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Cancel SettlementAllocationIntent API.
type V2PaymentsSettlementAllocationIntentCancelParams struct {
	Params `form:"*"`
}

// Submit SettlementAllocationIntent API.
type V2PaymentsSettlementAllocationIntentSubmitParams struct {
	Params `form:"*"`
}

// The amount and currency of the SettlementAllocationIntent.
type V2PaymentsSettlementAllocationIntentCreateAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// Create SettlementAllocationIntent API.
type V2PaymentsSettlementAllocationIntentCreateParams struct {
	Params `form:"*"`
	// The amount and currency of the SettlementAllocationIntent.
	Amount *V2PaymentsSettlementAllocationIntentCreateAmountParams `form:"amount" json:"amount"`
	// Expected date when we expect to receive the funds.
	ExpectedSettlementDate *time.Time `form:"expected_settlement_date" json:"expected_settlement_date"`
	// FinancialAccount where the funds are expected to land / FinancialAccount to map this SettlementAllocationIntent to.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
	// Metadata associated with the SettlementAllocationIntent.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Reference for the settlement intent . Max 255 characters . The reference used by PSP to send funds to Stripe .
	Reference *string `form:"reference" json:"reference"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2PaymentsSettlementAllocationIntentCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieve SettlementAllocationIntent API.
type V2PaymentsSettlementAllocationIntentRetrieveParams struct {
	Params `form:"*"`
}

// The new amount for the SettlementAllocationIntent.
type V2PaymentsSettlementAllocationIntentUpdateAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// Update SettlementAllocationIntent API.
type V2PaymentsSettlementAllocationIntentUpdateParams struct {
	Params `form:"*"`
	// The new amount for the SettlementAllocationIntent.
	Amount *V2PaymentsSettlementAllocationIntentUpdateAmountParams `form:"amount" json:"amount,omitempty"`
	// The new reference for the SettlementAllocationIntent.
	Reference *string `form:"reference" json:"reference,omitempty"`
}
