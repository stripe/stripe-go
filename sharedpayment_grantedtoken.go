//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The reason why the SharedPaymentGrantedToken has been deactivated.
type SharedPaymentGrantedTokenDeactivatedReason string

// List of values that SharedPaymentGrantedTokenDeactivatedReason can take
const (
	SharedPaymentGrantedTokenDeactivatedReasonConsumed SharedPaymentGrantedTokenDeactivatedReason = "consumed"
	SharedPaymentGrantedTokenDeactivatedReasonExpired  SharedPaymentGrantedTokenDeactivatedReason = "expired"
	SharedPaymentGrantedTokenDeactivatedReasonRevoked  SharedPaymentGrantedTokenDeactivatedReason = "revoked"
)

// Retrieves an existing SharedPaymentGrantedToken object
type SharedPaymentGrantedTokenParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *SharedPaymentGrantedTokenParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves an existing SharedPaymentGrantedToken object
type SharedPaymentGrantedTokenRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *SharedPaymentGrantedTokenRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The total amount captured using this SharedPaymentToken.
type SharedPaymentGrantedTokenUsageDetailsAmountCaptured struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Integer value of the amount in the smallest currency unit.
	Value int64 `json:"value"`
}

// Some details about how the SharedPaymentGrantedToken has been used already.
type SharedPaymentGrantedTokenUsageDetails struct {
	// The total amount captured using this SharedPaymentToken.
	AmountCaptured *SharedPaymentGrantedTokenUsageDetailsAmountCaptured `json:"amount_captured"`
}

// Limits on how this SharedPaymentGrantedToken can be used.
type SharedPaymentGrantedTokenUsageLimits struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Time at which this SharedPaymentToken expires and can no longer be used to confirm a PaymentIntent.
	ExpiresAt int64 `json:"expires_at"`
	// Max amount that can be captured using this SharedPaymentToken.
	MaxAmount int64 `json:"max_amount"`
}

// SharedPaymentGrantedToken is the view-only resource of a SharedPaymentIssuedToken, which is a limited-use reference to a PaymentMethod.
// When another Stripe merchant shares a SharedPaymentIssuedToken with you, you can view attributes of the shared token using the SharedPaymentGrantedToken API, and use it with a PaymentIntent.
type SharedPaymentGrantedToken struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Time at which this SharedPaymentGrantedToken expires and can no longer be used to confirm a PaymentIntent.
	DeactivatedAt int64 `json:"deactivated_at"`
	// The reason why the SharedPaymentGrantedToken has been deactivated.
	DeactivatedReason SharedPaymentGrantedTokenDeactivatedReason `json:"deactivated_reason"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Metadata about the SharedPaymentGrantedToken.
	SharedMetadata map[string]string `json:"shared_metadata"`
	// Some details about how the SharedPaymentGrantedToken has been used already.
	UsageDetails *SharedPaymentGrantedTokenUsageDetails `json:"usage_details"`
	// Limits on how this SharedPaymentGrantedToken can be used.
	UsageLimits *SharedPaymentGrantedTokenUsageLimits `json:"usage_limits"`
}
