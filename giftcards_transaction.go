//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of event that created this object.
type GiftCardsTransactionCreatedByType string

// List of values that GiftCardsTransactionCreatedByType can take
const (
	GiftCardsTransactionCreatedByTypeCheckout GiftCardsTransactionCreatedByType = "checkout"
	GiftCardsTransactionCreatedByTypeOrder    GiftCardsTransactionCreatedByType = "order"
	GiftCardsTransactionCreatedByTypePayment  GiftCardsTransactionCreatedByType = "payment"
)

// Status of this transaction, one of `held`, `confirmed`, or `canceled`.
type GiftCardsTransactionStatus string

// List of values that GiftCardsTransactionStatus can take
const (
	GiftCardsTransactionStatusCanceled  GiftCardsTransactionStatus = "canceled"
	GiftCardsTransactionStatusConfirmed GiftCardsTransactionStatus = "confirmed"
	GiftCardsTransactionStatusHeld      GiftCardsTransactionStatus = "held"
	GiftCardsTransactionStatusInvalid   GiftCardsTransactionStatus = "invalid"
)

// List gift card transactions for a gift card
type GiftCardsTransactionListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The gift card to list transactions for.
	GiftCard *string `form:"gift_card"`
	// A string that identifies this transaction as part of a group. See the [Connect documentation](https://stripe.com/docs/connect/separate-charges-and-transfers) for details.
	TransferGroup *string `form:"transfer_group"`
}

// AddExpand appends a new field to expand.
func (p *GiftCardsTransactionListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The details for the payment that created this object.
type GiftCardsTransactionCreatedByPaymentParams struct {
	// The PaymentIntent used to collect payment for this object.
	PaymentIntent *string `form:"payment_intent"`
}

// Related objects which created this transaction.
type GiftCardsTransactionCreatedByParams struct {
	// The details for the payment that created this object.
	Payment *GiftCardsTransactionCreatedByPaymentParams `form:"payment"`
	// The type of event that created this object.
	Type *string `form:"type"`
}

// Create a gift card transaction
type GiftCardsTransactionParams struct {
	Params `form:"*"`
	// The amount of the transaction. A negative amount deducts funds, and a positive amount adds funds.
	Amount *int64 `form:"amount"`
	// Whether this is a confirmed transaction. A confirmed transaction immediately deducts from/adds to the `amount_available` on the gift card. Otherwise, it creates a held transaction that increments the `amount_held` on the gift card.
	Confirm *bool `form:"confirm"`
	// Related objects which created this transaction.
	CreatedBy *GiftCardsTransactionCreatedByParams `form:"created_by"`
	// The currency of the transaction. This must match the currency of the gift card.
	Currency *string `form:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The gift card to create a new transaction on.
	GiftCard *string `form:"gift_card"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// A string that identifies this transaction as part of a group. See the [Connect documentation](https://stripe.com/docs/connect/separate-charges-and-transfers) for details.
	TransferGroup *string `form:"transfer_group"`
}

// AddExpand appends a new field to expand.
func (p *GiftCardsTransactionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *GiftCardsTransactionParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Cancel a gift card transaction
type GiftCardsTransactionCancelParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *GiftCardsTransactionCancelParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Confirm a gift card transaction
type GiftCardsTransactionConfirmParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *GiftCardsTransactionConfirmParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type GiftCardsTransactionCreatedByCheckout struct {
	// The Stripe CheckoutSession that created this object.
	CheckoutSession string `json:"checkout_session"`
	// The Stripe CheckoutSession LineItem that created this object.
	LineItem string `json:"line_item"`
}
type GiftCardsTransactionCreatedByOrder struct {
	// The Stripe Order LineItem that created this object.
	LineItem string `json:"line_item"`
	// The Stripe Order that created this object.
	Order string `json:"order"`
}
type GiftCardsTransactionCreatedByPayment struct {
	// The PaymentIntent that created this object.
	PaymentIntent string `json:"payment_intent"`
}

// The related Stripe objects that created this gift card transaction.
type GiftCardsTransactionCreatedBy struct {
	Checkout *GiftCardsTransactionCreatedByCheckout `json:"checkout"`
	Order    *GiftCardsTransactionCreatedByOrder    `json:"order"`
	Payment  *GiftCardsTransactionCreatedByPayment  `json:"payment"`
	// The type of event that created this object.
	Type GiftCardsTransactionCreatedByType `json:"type"`
}

// A gift card transaction represents a single transaction on a referenced gift card.
// A transaction is in one of three states, `confirmed`, `held` or `canceled`. A `confirmed`
// transaction is one that has added/deducted funds. A `held` transaction has created a
// temporary hold on funds, which can then be cancelled or confirmed. A `held` transaction
// can be confirmed into a `confirmed` transaction, or canceled into a `canceled` transaction.
// A `canceled` transaction has no effect on a gift card's balance.
type GiftCardsTransaction struct {
	APIResource
	// The amount of this transaction. A positive value indicates that funds were added to the gift card. A negative value indicates that funds were removed from the gift card.
	Amount int64 `json:"amount"`
	// Time at which the transaction was confirmed. Measured in seconds since the Unix epoch.
	ConfirmedAt int64 `json:"confirmed_at"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The related Stripe objects that created this gift card transaction.
	CreatedBy *GiftCardsTransactionCreatedBy `json:"created_by"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// The gift card that this transaction occurred on
	GiftCard string `json:"gift_card"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Status of this transaction, one of `held`, `confirmed`, or `canceled`.
	Status GiftCardsTransactionStatus `json:"status"`
	// A string that identifies this transaction as part of a group. See the [Connect documentation](https://stripe.com/docs/connect/separate-charges-and-transfers) for details.
	TransferGroup string `json:"transfer_group"`
}

// GiftCardsTransactionList is a list of Transactions as retrieved from a list endpoint.
type GiftCardsTransactionList struct {
	APIResource
	ListMeta
	Data []*GiftCardsTransaction `json:"data"`
}
