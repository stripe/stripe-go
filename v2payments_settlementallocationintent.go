//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// SettlementAllocationIntent status.
type V2PaymentsSettlementAllocationIntentStatus string

// List of values that V2PaymentsSettlementAllocationIntentStatus can take
const (
	V2PaymentsSettlementAllocationIntentStatusCanceled  V2PaymentsSettlementAllocationIntentStatus = "canceled"
	V2PaymentsSettlementAllocationIntentStatusErrored   V2PaymentsSettlementAllocationIntentStatus = "errored"
	V2PaymentsSettlementAllocationIntentStatusMatched   V2PaymentsSettlementAllocationIntentStatus = "matched"
	V2PaymentsSettlementAllocationIntentStatusPending   V2PaymentsSettlementAllocationIntentStatus = "pending"
	V2PaymentsSettlementAllocationIntentStatusSettled   V2PaymentsSettlementAllocationIntentStatus = "settled"
	V2PaymentsSettlementAllocationIntentStatusSubmitted V2PaymentsSettlementAllocationIntentStatus = "submitted"
)

// Open Enum. The `errored` status reason.
type V2PaymentsSettlementAllocationIntentStatusDetailsErroredReasonCode string

// List of values that V2PaymentsSettlementAllocationIntentStatusDetailsErroredReasonCode can take
const (
	V2PaymentsSettlementAllocationIntentStatusDetailsErroredReasonCodeAmountMismatch V2PaymentsSettlementAllocationIntentStatusDetailsErroredReasonCode = "amount_mismatch"
)

// The amount and currency of the SettlementAllocationIntent.
type V2PaymentsSettlementAllocationIntentAmount struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value int64 `json:"value,omitempty"`
}

// Hash that provides additional information regarding the reason behind a `errored` SettlementAllocationIntent status. It is only present when the SettlementAllocationIntent status is `errored`.
type V2PaymentsSettlementAllocationIntentStatusDetailsErrored struct {
	// Stripe doc link to debug the issue.
	DocURL string `json:"doc_url,omitempty"`
	// User Message detailing the reason code and possible resolution .
	Message string `json:"message"`
	// Open Enum. The `errored` status reason.
	ReasonCode V2PaymentsSettlementAllocationIntentStatusDetailsErroredReasonCode `json:"reason_code"`
}

// Status details for a SettlementAllocationIntent in `errored` state.
type V2PaymentsSettlementAllocationIntentStatusDetails struct {
	// Hash that provides additional information regarding the reason behind a `errored` SettlementAllocationIntent status. It is only present when the SettlementAllocationIntent status is `errored`.
	Errored *V2PaymentsSettlementAllocationIntentStatusDetailsErrored `json:"errored,omitempty"`
}

// SettlementAllocationIntent resource.
type V2PaymentsSettlementAllocationIntent struct {
	APIResource
	// The amount and currency of the SettlementAllocationIntent.
	Amount *V2PaymentsSettlementAllocationIntentAmount `json:"amount"`
	// Timestamp at which SettlementAllocationIntent was created .
	Created time.Time `json:"created"`
	// Date when we expect to receive the funds.
	ExpectedSettlementDate time.Time `json:"expected_settlement_date"`
	// FinancialAccount ID where the funds are expected.
	FinancialAccount string `json:"financial_account"`
	// Unique identifier for the SettlementAllocationIntent.
	ID string `json:"id"`
	// List of ReceivedCredits that matched with the SettlementAllocationIntent.
	LinkedCredits []string `json:"linked_credits"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Metadata associated with the SettlementAllocationIntent.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Reference for the SettlementAllocationIntent. This is the transaction reference used by payments processor to send funds to Stripe .
	Reference string `json:"reference"`
	// SettlementAllocationIntent status.
	Status V2PaymentsSettlementAllocationIntentStatus `json:"status"`
	// Status details for a SettlementAllocationIntent in `errored` state.
	StatusDetails *V2PaymentsSettlementAllocationIntentStatusDetails `json:"status_details,omitempty"`
}
