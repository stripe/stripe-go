//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The amount and currency of the SettlementAllocationIntent. Allowed Currencies are `gbp` | `eur`.
type V2PaymentsSettlementAllocationIntentAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// Create a new SettlementAllocationIntent.
type V2PaymentsSettlementAllocationIntentParams struct {
	Params `form:"*"`
	// The new amount for the SettlementAllocationIntent. Only amount.value can be updated and currency must remain same.
	Amount *V2PaymentsSettlementAllocationIntentAmountParams `form:"amount" json:"amount,omitempty"`
	// Date when we expect to receive the funds. Must be in future .
	ExpectedSettlementDate *time.Time `form:"expected_settlement_date" json:"expected_settlement_date,omitempty"`
	// Financial Account Id where the funds are expected for this SettlementAllocationIntent.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// Metadata associated with the SettlementAllocationIntent.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Reference for the SettlementAllocationIntent. This should be same as the transaction reference used by payments processor to send funds to Stripe. Must have length between 5 and 255 characters and it must be unique among existing SettlementAllocationIntents that have a non-terminal status (`pending`, `submitted`, `matched`, `errored`).
	Reference *string `form:"reference" json:"reference,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2PaymentsSettlementAllocationIntentParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Cancels an existing SettlementAllocationIntent. Only SettlementAllocationIntent with status `pending`, `submitted` and `errored` can be `canceled`.
type V2PaymentsSettlementAllocationIntentCancelParams struct {
	Params `form:"*"`
}

// Submits a SettlementAllocationIntent. Only SettlementAllocationIntent with status `pending` can be `submitted`. The net sum of SettlementAllocationIntentSplit amount must be equal to SettlementAllocationIntent amount to be eligible to be submitted.
type V2PaymentsSettlementAllocationIntentSubmitParams struct {
	Params `form:"*"`
}

// The amount and currency of the SettlementAllocationIntent. Allowed Currencies are `gbp` | `eur`.
type V2PaymentsSettlementAllocationIntentCreateAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// Create a new SettlementAllocationIntent.
type V2PaymentsSettlementAllocationIntentCreateParams struct {
	Params `form:"*"`
	// The amount and currency of the SettlementAllocationIntent. Allowed Currencies are `gbp` | `eur`.
	Amount *V2PaymentsSettlementAllocationIntentCreateAmountParams `form:"amount" json:"amount"`
	// Date when we expect to receive the funds. Must be in future .
	ExpectedSettlementDate *time.Time `form:"expected_settlement_date" json:"expected_settlement_date"`
	// Financial Account Id where the funds are expected for this SettlementAllocationIntent.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
	// Metadata associated with the SettlementAllocationIntent.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Reference for the SettlementAllocationIntent. This should be same as the transaction reference used by payments processor to send funds to Stripe. Must have length between 5 and 255 characters and it must be unique among existing SettlementAllocationIntents that have a non-terminal status (`pending`, `submitted`, `matched`, `errored`).
	Reference *string `form:"reference" json:"reference"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2PaymentsSettlementAllocationIntentCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieve an existing SettlementAllocationIntent.
type V2PaymentsSettlementAllocationIntentRetrieveParams struct {
	Params `form:"*"`
}

// The new amount for the SettlementAllocationIntent. Only amount.value can be updated and currency must remain same.
type V2PaymentsSettlementAllocationIntentUpdateAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// Updates SettlementAllocationIntent. Only SettlementAllocationIntent with status `pending`, `submitted` and `errored` can be updated. Only amount and reference fields can be updated for a SettlementAllocationIntent and at least one must be present. Updating an `amount` moves the SettlementAllocationIntent `pending` status and updating the `reference` for `errored` SettlementAllocationIntent moves it to `submitted`.
type V2PaymentsSettlementAllocationIntentUpdateParams struct {
	Params `form:"*"`
	// The new amount for the SettlementAllocationIntent. Only amount.value can be updated and currency must remain same.
	Amount *V2PaymentsSettlementAllocationIntentUpdateAmountParams `form:"amount" json:"amount,omitempty"`
	// The new reference for the SettlementAllocationIntent.
	Reference *string `form:"reference" json:"reference,omitempty"`
}
