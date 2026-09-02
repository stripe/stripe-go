//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Creates a new payment retry evaluation for a failed payment.
type V2SignalsPaymentRetryEvaluationParams struct {
	Params `form:"*"`
	// ID of the PaymentIntent to evaluate. Mutually exclusive with payment_record.
	PaymentIntent *string `form:"payment_intent" json:"payment_intent,omitempty"`
	// ID of the PaymentRecord to evaluate. Mutually exclusive with payment_intent.
	PaymentRecord *string `form:"payment_record" json:"payment_record,omitempty"`
}

// Cancels an active payment retry evaluation.
type V2SignalsPaymentRetryEvaluationCancelParams struct {
	Params `form:"*"`
	// Optional reason for canceling the evaluation.
	CancellationReason *string `form:"cancellation_reason" json:"cancellation_reason,omitempty"`
}

// Creates a new payment retry evaluation for a failed payment.
type V2SignalsPaymentRetryEvaluationCreateParams struct {
	Params `form:"*"`
	// ID of the PaymentIntent to evaluate. Mutually exclusive with payment_record.
	PaymentIntent *string `form:"payment_intent" json:"payment_intent,omitempty"`
	// ID of the PaymentRecord to evaluate. Mutually exclusive with payment_intent.
	PaymentRecord *string `form:"payment_record" json:"payment_record,omitempty"`
}

// Retrieves a payment retry evaluation by ID.
type V2SignalsPaymentRetryEvaluationRetrieveParams struct {
	Params `form:"*"`
}

// Updates an active payment retry evaluation with a replacement payment identifier.
type V2SignalsPaymentRetryEvaluationUpdateParams struct {
	Params `form:"*"`
	// PaymentIntent to update to. Must match the evaluation's signal type.
	PaymentIntent *string `form:"payment_intent" json:"payment_intent,omitempty"`
	// PaymentRecord to update to. Must match the evaluation's signal type.
	PaymentRecord *string `form:"payment_record" json:"payment_record,omitempty"`
}
