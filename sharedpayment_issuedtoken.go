//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The reason why the SharedPaymentIssuedToken has been deactivated.
type SharedPaymentIssuedTokenDeactivatedReason string

// List of values that SharedPaymentIssuedTokenDeactivatedReason can take
const (
	SharedPaymentIssuedTokenDeactivatedReasonConsumed SharedPaymentIssuedTokenDeactivatedReason = "consumed"
	SharedPaymentIssuedTokenDeactivatedReasonExpired  SharedPaymentIssuedTokenDeactivatedReason = "expired"
	SharedPaymentIssuedTokenDeactivatedReasonResolved SharedPaymentIssuedTokenDeactivatedReason = "resolved"
	SharedPaymentIssuedTokenDeactivatedReasonRevoked  SharedPaymentIssuedTokenDeactivatedReason = "revoked"
)

// Specifies the type of next action required. Determines which child attribute contains action details.
type SharedPaymentIssuedTokenNextActionType string

// List of values that SharedPaymentIssuedTokenNextActionType can take
const (
	SharedPaymentIssuedTokenNextActionTypeUseStripeSDK SharedPaymentIssuedTokenNextActionType = "use_stripe_sdk"
)

// Indicates that you intend to save the PaymentMethod of this SharedPaymentToken to a customer later.
type SharedPaymentIssuedTokenSetupFutureUsage string

// List of values that SharedPaymentIssuedTokenSetupFutureUsage can take
const (
	SharedPaymentIssuedTokenSetupFutureUsageOnSession SharedPaymentIssuedTokenSetupFutureUsage = "on_session"
)

// Status of this SharedPaymentIssuedToken, one of `active`, `requires_action`, or `deactivated`.
type SharedPaymentIssuedTokenStatus string

// List of values that SharedPaymentIssuedTokenStatus can take
const (
	SharedPaymentIssuedTokenStatusActive         SharedPaymentIssuedTokenStatus = "active"
	SharedPaymentIssuedTokenStatusDeactivated    SharedPaymentIssuedTokenStatus = "deactivated"
	SharedPaymentIssuedTokenStatusRequiresAction SharedPaymentIssuedTokenStatus = "requires_action"
)

// The recurring interval at which the shared payment token's amount usage restrictions reset.
type SharedPaymentIssuedTokenUsageLimitsRecurringInterval string

// List of values that SharedPaymentIssuedTokenUsageLimitsRecurringInterval can take
const (
	SharedPaymentIssuedTokenUsageLimitsRecurringIntervalMonth SharedPaymentIssuedTokenUsageLimitsRecurringInterval = "month"
	SharedPaymentIssuedTokenUsageLimitsRecurringIntervalWeek  SharedPaymentIssuedTokenUsageLimitsRecurringInterval = "week"
	SharedPaymentIssuedTokenUsageLimitsRecurringIntervalYear  SharedPaymentIssuedTokenUsageLimitsRecurringInterval = "year"
)

// Contains details for handling the next action using Stripe.js, iOS, or Android SDKs. Present when `next_action.type` is `use_stripe_sdk`.
type SharedPaymentIssuedTokenNextActionUseStripeSDK struct {
	// A base64-encoded string used by Stripe.js and the iOS and Android client SDKs to handle the next action. Its content is subject to change.
	Value string `json:"value"`
}

// If present, describes the action required to make this `SharedPaymentIssuedToken` usable for payments. Present when the token is in `requires_action` state.
type SharedPaymentIssuedTokenNextAction struct {
	// Specifies the type of next action required. Determines which child attribute contains action details.
	Type SharedPaymentIssuedTokenNextActionType `json:"type"`
	// Contains details for handling the next action using Stripe.js, iOS, or Android SDKs. Present when `next_action.type` is `use_stripe_sdk`.
	UseStripeSDK *SharedPaymentIssuedTokenNextActionUseStripeSDK `json:"use_stripe_sdk"`
}

// Bot risk insight.
type SharedPaymentIssuedTokenRiskDetailsInsightsBot struct {
	// Recommended action for this insight.
	RecommendedAction string `json:"recommended_action"`
	// Risk score for this insight.
	Score float64 `json:"score"`
}

// Card issuer decline risk insight.
type SharedPaymentIssuedTokenRiskDetailsInsightsCardIssuerDecline struct {
	// Recommended action for this insight.
	RecommendedAction string `json:"recommended_action"`
	// Risk score for this insight.
	Score float64 `json:"score"`
}

// Card testing risk insight.
type SharedPaymentIssuedTokenRiskDetailsInsightsCardTesting struct {
	// Recommended action for this insight.
	RecommendedAction string `json:"recommended_action"`
	// Risk score for this insight.
	Score float64 `json:"score"`
}

// Fraudulent dispute risk insight.
type SharedPaymentIssuedTokenRiskDetailsInsightsFraudulentDispute struct {
	// Recommended action for this insight.
	RecommendedAction string `json:"recommended_action"`
	// Risk score for this insight.
	Score int64 `json:"score"`
}

// Stolen card risk insight.
type SharedPaymentIssuedTokenRiskDetailsInsightsStolenCard struct {
	// Recommended action for this insight.
	RecommendedAction string `json:"recommended_action"`
	// Risk score for this insight.
	Score int64 `json:"score"`
}

// Risk insights for this token, including scores and recommended actions for each risk type.
type SharedPaymentIssuedTokenRiskDetailsInsights struct {
	// Bot risk insight.
	Bot *SharedPaymentIssuedTokenRiskDetailsInsightsBot `json:"bot,omitempty"`
	// Card issuer decline risk insight.
	CardIssuerDecline *SharedPaymentIssuedTokenRiskDetailsInsightsCardIssuerDecline `json:"card_issuer_decline,omitempty"`
	// Card testing risk insight.
	CardTesting *SharedPaymentIssuedTokenRiskDetailsInsightsCardTesting `json:"card_testing,omitempty"`
	// Fraudulent dispute risk insight.
	FraudulentDispute *SharedPaymentIssuedTokenRiskDetailsInsightsFraudulentDispute `json:"fraudulent_dispute"`
	// Stolen card risk insight.
	StolenCard *SharedPaymentIssuedTokenRiskDetailsInsightsStolenCard `json:"stolen_card,omitempty"`
}

// Risk details of the SharedPaymentIssuedToken.
type SharedPaymentIssuedTokenRiskDetails struct {
	// Risk insights for this token, including scores and recommended actions for each risk type.
	Insights *SharedPaymentIssuedTokenRiskDetailsInsights `json:"insights"`
}

// Seller details of the SharedPaymentIssuedToken, including network_id and external_id.
type SharedPaymentIssuedTokenSellerDetails struct {
	// A unique id within a network that identifies a logical seller. This should usually be the merchant id on the seller platform.
	ExternalID string `json:"external_id"`
	// The unique and logical string that identifies the seller platform that this SharedToken is being created for.
	NetworkBusinessProfile string `json:"network_business_profile"`
}

// The total amount captured using this SharedPaymentToken.
type SharedPaymentIssuedTokenUsageDetailsAmountCaptured struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Integer value of the amount in the smallest currency unit.
	Value int64 `json:"value"`
}

// Usage details of the SharedPaymentIssuedToken
type SharedPaymentIssuedTokenUsageDetails struct {
	// The total amount captured using this SharedPaymentToken.
	AmountCaptured *SharedPaymentIssuedTokenUsageDetailsAmountCaptured `json:"amount_captured"`
}

// Usage limits of the SharedPaymentIssuedToken.
type SharedPaymentIssuedTokenUsageLimits struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Time at which this SharedPaymentToken expires and can no longer be used to confirm a PaymentIntent.
	ExpiresAt int64 `json:"expires_at"`
	// Max amount that can be captured using this SharedPaymentToken.
	MaxAmount int64 `json:"max_amount"`
	// The recurring interval at which the shared payment token's amount usage restrictions reset.
	RecurringInterval SharedPaymentIssuedTokenUsageLimitsRecurringInterval `json:"recurring_interval,omitempty"`
}

// A SharedPaymentIssuedToken is a limited-use reference to a PaymentMethod that can be created with a secret key. When shared with another Stripe account (Seller), it enables that account to either process a payment on Stripe against a PaymentMethod that your Stripe account owns, or to forward a usable credential created against the originalPaymentMethod to then process the payment off-Stripe.
type SharedPaymentIssuedToken struct {
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// ID of an existing Customer.
	Customer string `json:"customer"`
	// Time at which this SharedPaymentIssuedToken was deactivated.
	DeactivatedAt int64 `json:"deactivated_at"`
	// The reason why the SharedPaymentIssuedToken has been deactivated.
	DeactivatedReason SharedPaymentIssuedTokenDeactivatedReason `json:"deactivated_reason"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// If present, describes the action required to make this `SharedPaymentIssuedToken` usable for payments. Present when the token is in `requires_action` state.
	NextAction *SharedPaymentIssuedTokenNextAction `json:"next_action"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// ID of an existing PaymentMethod.
	PaymentMethod string `json:"payment_method"`
	// If the customer does not exit their browser while authenticating, they will be redirected to this specified URL after completion.
	ReturnURL string `json:"return_url,omitempty"`
	// Risk details of the SharedPaymentIssuedToken.
	RiskDetails *SharedPaymentIssuedTokenRiskDetails `json:"risk_details,omitempty"`
	// Seller details of the SharedPaymentIssuedToken, including network_id and external_id.
	SellerDetails *SharedPaymentIssuedTokenSellerDetails `json:"seller_details"`
	// Indicates that you intend to save the PaymentMethod of this SharedPaymentToken to a customer later.
	SetupFutureUsage SharedPaymentIssuedTokenSetupFutureUsage `json:"setup_future_usage"`
	// Metadata about the SharedPaymentIssuedToken.
	SharedMetadata map[string]string `json:"shared_metadata"`
	// Status of this SharedPaymentIssuedToken, one of `active`, `requires_action`, or `deactivated`.
	Status SharedPaymentIssuedTokenStatus `json:"status"`
	// Usage details of the SharedPaymentIssuedToken
	UsageDetails *SharedPaymentIssuedTokenUsageDetails `json:"usage_details"`
	// Usage limits of the SharedPaymentIssuedToken.
	UsageLimits *SharedPaymentIssuedTokenUsageLimits `json:"usage_limits"`
}

// UnmarshalJSON handles deserialization of a SharedPaymentIssuedToken.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *SharedPaymentIssuedToken) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type sharedPaymentIssuedToken SharedPaymentIssuedToken
	var v sharedPaymentIssuedToken
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = SharedPaymentIssuedToken(v)
	return nil
}
