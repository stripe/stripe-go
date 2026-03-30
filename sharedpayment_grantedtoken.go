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
	SharedPaymentGrantedTokenDeactivatedReasonResolved SharedPaymentGrantedTokenDeactivatedReason = "resolved"
	SharedPaymentGrantedTokenDeactivatedReasonRevoked  SharedPaymentGrantedTokenDeactivatedReason = "revoked"
)

// The recurring interval at which the shared payment token's amount usage restrictions reset.
type SharedPaymentGrantedTokenUsageLimitsRecurringInterval string

// List of values that SharedPaymentGrantedTokenUsageLimitsRecurringInterval can take
const (
	SharedPaymentGrantedTokenUsageLimitsRecurringIntervalMonth SharedPaymentGrantedTokenUsageLimitsRecurringInterval = "month"
	SharedPaymentGrantedTokenUsageLimitsRecurringIntervalWeek  SharedPaymentGrantedTokenUsageLimitsRecurringInterval = "week"
	SharedPaymentGrantedTokenUsageLimitsRecurringIntervalYear  SharedPaymentGrantedTokenUsageLimitsRecurringInterval = "year"
)

// Retrieves an existing SharedPaymentGrantedToken object
type SharedPaymentGrantedTokenParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *SharedPaymentGrantedTokenParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves an existing SharedPaymentGrantedToken object
type SharedPaymentGrantedTokenRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
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
	// The recurring interval at which the shared payment token's amount usage restrictions reset.
	RecurringInterval SharedPaymentGrantedTokenUsageLimitsRecurringInterval `json:"recurring_interval,omitempty"`
}

// Details about the agent that issued this SharedPaymentGrantedToken.
type SharedPaymentGrantedTokenAgentDetails struct {
	// The Stripe Profile ID of the agent that issued this SharedPaymentGrantedToken.
	NetworkBusinessProfile string `json:"network_business_profile"`
}

// Bot risk insight (score: Float, recommended_action).
type SharedPaymentGrantedTokenRiskDetailsInsightsBot struct {
	// Recommended action for this insight.
	RecommendedAction string `json:"recommended_action"`
	// Risk score for this insight (float).
	Score float64 `json:"score"`
}

// Card issuer decline risk insight (score: Float, recommended_action).
type SharedPaymentGrantedTokenRiskDetailsInsightsCardIssuerDecline struct {
	// Recommended action for this insight.
	RecommendedAction string `json:"recommended_action"`
	// Risk score for this insight (float).
	Score float64 `json:"score"`
}

// Card testing risk insight (score: Float, recommended_action).
type SharedPaymentGrantedTokenRiskDetailsInsightsCardTesting struct {
	// Recommended action for this insight.
	RecommendedAction string `json:"recommended_action"`
	// Risk score for this insight (float).
	Score float64 `json:"score"`
}

// Fraudulent dispute risk insight (score: Integer, recommended_action).
type SharedPaymentGrantedTokenRiskDetailsInsightsFraudulentDispute struct {
	// Recommended action for this insight.
	RecommendedAction string `json:"recommended_action"`
	// Risk score for this insight (integer).
	Score int64 `json:"score"`
}

// Stolen card risk insight (score: Integer, recommended_action).
type SharedPaymentGrantedTokenRiskDetailsInsightsStolenCard struct {
	// Recommended action for this insight.
	RecommendedAction string `json:"recommended_action"`
	// Risk score for this insight (integer).
	Score int64 `json:"score"`
}

// Risk insights for this token, including scores and recommended actions for each risk type.
type SharedPaymentGrantedTokenRiskDetailsInsights struct {
	// Bot risk insight (score: Float, recommended_action).
	Bot *SharedPaymentGrantedTokenRiskDetailsInsightsBot `json:"bot,omitempty"`
	// Card issuer decline risk insight (score: Float, recommended_action).
	CardIssuerDecline *SharedPaymentGrantedTokenRiskDetailsInsightsCardIssuerDecline `json:"card_issuer_decline,omitempty"`
	// Card testing risk insight (score: Float, recommended_action).
	CardTesting *SharedPaymentGrantedTokenRiskDetailsInsightsCardTesting `json:"card_testing,omitempty"`
	// Fraudulent dispute risk insight (score: Integer, recommended_action).
	FraudulentDispute *SharedPaymentGrantedTokenRiskDetailsInsightsFraudulentDispute `json:"fraudulent_dispute"`
	// Stolen card risk insight (score: Integer, recommended_action).
	StolenCard *SharedPaymentGrantedTokenRiskDetailsInsightsStolenCard `json:"stolen_card,omitempty"`
}

// Risk details of the SharedPaymentGrantedToken.
type SharedPaymentGrantedTokenRiskDetails struct {
	// Risk insights for this token, including scores and recommended actions for each risk type.
	Insights *SharedPaymentGrantedTokenRiskDetailsInsights `json:"insights"`
}

// SharedPaymentGrantedToken is the view-only resource of a SharedPaymentIssuedToken, which is a limited-use reference to a PaymentMethod.
// When another Stripe merchant shares a SharedPaymentIssuedToken with you, you can view attributes of the shared token using the SharedPaymentGrantedToken API, and use it with a PaymentIntent.
type SharedPaymentGrantedToken struct {
	APIResource
	// Details about the agent that issued this SharedPaymentGrantedToken.
	AgentDetails *SharedPaymentGrantedTokenAgentDetails `json:"agent_details,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Time at which this SharedPaymentGrantedToken expires and can no longer be used to confirm a PaymentIntent.
	DeactivatedAt int64 `json:"deactivated_at"`
	// The reason why the SharedPaymentGrantedToken has been deactivated.
	DeactivatedReason SharedPaymentGrantedTokenDeactivatedReason `json:"deactivated_reason"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Risk details of the SharedPaymentGrantedToken.
	RiskDetails *SharedPaymentGrantedTokenRiskDetails `json:"risk_details,omitempty"`
	// Metadata about the SharedPaymentGrantedToken.
	SharedMetadata map[string]string `json:"shared_metadata"`
	// Some details about how the SharedPaymentGrantedToken has been used already.
	UsageDetails *SharedPaymentGrantedTokenUsageDetails `json:"usage_details"`
	// Limits on how this SharedPaymentGrantedToken can be used.
	UsageLimits *SharedPaymentGrantedTokenUsageLimits `json:"usage_limits"`
}
