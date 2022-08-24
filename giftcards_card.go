//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of event that created this object.
type GiftCardsCardCreatedByType string

// List of values that GiftCardsCardCreatedByType can take
const (
	GiftCardsCardCreatedByTypeCheckout GiftCardsCardCreatedByType = "checkout"
	GiftCardsCardCreatedByTypeOrder    GiftCardsCardCreatedByType = "order"
	GiftCardsCardCreatedByTypePayment  GiftCardsCardCreatedByType = "payment"
)

// The details for the payment that created this object.
type GiftCardsCardCreatedByPaymentParams struct {
	// The PaymentIntent used to collect payment for this object.
	PaymentIntent *string `form:"payment_intent"`
}

// Related objects which created this gift card.
type GiftCardsCardCreatedByParams struct {
	// The details for the payment that created this object.
	Payment *GiftCardsCardCreatedByPaymentParams `form:"payment"`
	// The type of event that created this object.
	Type *string `form:"type"`
}

// Creates a new gift card object.
type GiftCardsCardParams struct {
	Params `form:"*"`
	// The new active state for the gift card.
	Active *bool `form:"active"`
	// Related objects which created this gift card.
	CreatedBy *GiftCardsCardCreatedByParams `form:"created_by"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The initial amount to load onto the new gift card, defaults to 0.
	InitialAmount *int64 `form:"initial_amount"`
}

// List gift cards for an account
type GiftCardsCardListParams struct {
	ListParams `form:"*"`
}

// Validates a gift card code, returning the matching gift card object if it exists.
type GiftCardsCardValidateParams struct {
	Params `form:"*"`
	// The gift card code to be validated.
	Code *string `form:"code"`
	// The pin associated with the gift card. Not all gift cards have pins.
	PIN *string `form:"pin"`
}
type GiftCardsCardCreatedByCheckout struct {
	// The Stripe CheckoutSession that created this object.
	CheckoutSession string `json:"checkout_session"`
	// The Stripe CheckoutSession LineItem that created this object.
	LineItem string `json:"line_item"`
}
type GiftCardsCardCreatedByOrder struct {
	// The Stripe Order LineItem that created this object.
	LineItem string `json:"line_item"`
	// The Stripe Order that created this object.
	Order string `json:"order"`
}
type GiftCardsCardCreatedByPayment struct {
	// The PaymentIntent that created this object.
	PaymentIntent string `json:"payment_intent"`
}

// The related Stripe objects that created this gift card.
type GiftCardsCardCreatedBy struct {
	Checkout *GiftCardsCardCreatedByCheckout `json:"checkout"`
	Order    *GiftCardsCardCreatedByOrder    `json:"order"`
	Payment  *GiftCardsCardCreatedByPayment  `json:"payment"`
	// The type of event that created this object.
	Type GiftCardsCardCreatedByType `json:"type"`
}

// A gift card represents a single gift card owned by a customer, including the
// remaining balance, gift card code, and whether or not it is active.
type GiftCardsCard struct {
	APIResource
	// Whether this gift card can be used or not.
	Active bool `json:"active"`
	// The amount of funds available for new transactions.
	AmountAvailable int64 `json:"amount_available"`
	// The amount of funds marked as held.
	AmountHeld int64 `json:"amount_held"`
	// Code used to redeem this gift card.
	Code string `json:"code"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The related Stripe objects that created this gift card.
	CreatedBy *GiftCardsCardCreatedBy `json:"created_by"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Transactions on this gift card.
	Transactions *GiftCardsTransactionList `json:"transactions"`
}

// GiftCardsCardList is a list of Cards as retrieved from a list endpoint.
type GiftCardsCardList struct {
	APIResource
	ListMeta
	Data []*GiftCardsCard `json:"data"`
}
