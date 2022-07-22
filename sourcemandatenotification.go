//
//
// File generated from our OpenAPI spec
//
//

package stripe

type SourceMandateNotificationACSSDebit struct {
	// The statement descriptor associate with the debit.
	StatementDescriptor string `json:"statement_descriptor"`
}
type SourceMandateNotificationBACSDebit struct {
	// Last 4 digits of the account number associated with the debit.
	Last4 string `json:"last4"`
}
type SourceMandateNotificationSEPADebit struct {
	// SEPA creditor ID.
	CreditorIdentifier string `json:"creditor_identifier"`
	// Last 4 digits of the account number associated with the debit.
	Last4 string `json:"last4"`
	// Mandate reference associated with the debit.
	MandateReference string `json:"mandate_reference"`
}

// Source mandate notifications should be created when a notification related to
// a source mandate must be sent to the payer. They will trigger a webhook or
// deliver an email to the customer.
type SourceMandateNotification struct {
	ACSSDebit *SourceMandateNotificationACSSDebit `json:"acss_debit"`
	// A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the amount associated with the mandate notification. The amount is expressed in the currency of the underlying source. Required if the notification type is `debit_initiated`.
	Amount    int64                               `json:"amount"`
	BACSDebit *SourceMandateNotificationBACSDebit `json:"bacs_debit"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The reason of the mandate notification. Valid reasons are `mandate_confirmed` or `debit_initiated`.
	Reason    string                              `json:"reason"`
	SEPADebit *SourceMandateNotificationSEPADebit `json:"sepa_debit"`
	// `Source` objects allow you to accept a variety of payment methods. They
	// represent a customer's payment instrument, and can be used with the Stripe API
	// just like a `Card` object: once chargeable, they can be charged, or can be
	// attached to customers.
	//
	// Related guides: [Sources API](https://stripe.com/docs/sources) and [Sources & Customers](https://stripe.com/docs/sources/customers).
	Source *PaymentSource `json:"source"`
	// The status of the mandate notification. Valid statuses are `pending` or `submitted`.
	Status string `json:"status"`
	// The type of source this mandate notification is attached to. Should be the source type identifier code for the payment method, such as `three_d_secure`.
	Type string `json:"type"`
}
