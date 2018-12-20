package stripe

// CheckoutSessionLineItemParams is the set of parameters allowed for a line item
// on a Checkout Session.
type CheckoutSessionLineItemParams struct {
	Amount      *int64    `form:"amount"`
	Currency    *string   `form:"currency"`
	Description *string   `form:"description"`
	Name        *string   `form:"name"`
	Images      []*string `form:"images"`
	Quantity    *int64    `form:"quantity"`
}

// CheckoutSessionPaymentIntentDataParams is the set of parameters allowed for the
// payment intent creation on a Checkout Session.
type CheckoutSessionPaymentIntentDataParams struct {
	Description         *string                `form:"description"`
	ReceiptEmail        *string                `form:"receipt_email"`
	Shipping            *ShippingDetailsParams `form:"shipping"`
	StatementDescriptor *string                `form:"statement_descriptor"`
}

// CheckoutSessionParams is the set of parameters that can be used when creating
// a checkout session.
// For more details see https://stripe.com/docs/api#create_checkout_session.
type CheckoutSessionParams struct {
	Params             `form:"*"`
	AllowedSourceTypes []*string                               `form:"allowed_source_types"`
	CancelURL          *string                                 `form:"cancel_url"`
	ClientReferenceID  *string                                 `form:"client_reference_id"`
	LineItems          []*CheckoutSessionLineItemParams        `form:"line_items"`
	PaymentIntentData  *CheckoutSessionPaymentIntentDataParams `form:"payment_intent_data"`
	SuccessURL         *string                                 `form:"success_url"`
}

// CheckoutSession is the resource representing a Stripe checkout session.
// For more details see https://stripe.com/docs/api#checkout_sessions.
type CheckoutSession struct {
	ID       string `json:"id"`
	Livemode bool   `json:"livemode"`
	Object   string `json:"object"`
}
