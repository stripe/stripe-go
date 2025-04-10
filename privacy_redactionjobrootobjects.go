//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The objects to redact, grouped by type. All redactable objects associated with these objects will be redacted as well.
type PrivacyRedactionJobRootObjects struct {
	Charges                      []string `json:"charges"`
	CheckoutSessions             []string `json:"checkout_sessions"`
	Customers                    []string `json:"customers"`
	IdentityVerificationSessions []string `json:"identity_verification_sessions"`
	Invoices                     []string `json:"invoices"`
	IssuingCardholders           []string `json:"issuing_cardholders"`
	// String representing the object's type. Objects of the same type share the same value.
	Object              string   `json:"object"`
	PaymentIntents      []string `json:"payment_intents"`
	RadarValueListItems []string `json:"radar_value_list_items"`
	SetupIntents        []string `json:"setup_intents"`
}
