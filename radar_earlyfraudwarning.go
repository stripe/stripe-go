//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of fraud labelled by the issuer. One of `card_never_received`, `fraudulent_card_application`, `made_with_counterfeit_card`, `made_with_lost_card`, `made_with_stolen_card`, `misc`, `unauthorized_use_of_card`.
type RadarEarlyFraudWarningFraudType string

// List of values that RadarEarlyFraudWarningFraudType can take
const (
	RadarEarlyFraudWarningFraudTypeCardNeverReceived         RadarEarlyFraudWarningFraudType = "card_never_received"
	RadarEarlyFraudWarningFraudTypeFraudulentCardApplication RadarEarlyFraudWarningFraudType = "fraudulent_card_application"
	RadarEarlyFraudWarningFraudTypeMadeWithCounterfeitCard   RadarEarlyFraudWarningFraudType = "made_with_counterfeit_card"
	RadarEarlyFraudWarningFraudTypeMadeWithLostCard          RadarEarlyFraudWarningFraudType = "made_with_lost_card"
	RadarEarlyFraudWarningFraudTypeMadeWithStolenCard        RadarEarlyFraudWarningFraudType = "made_with_stolen_card"
	RadarEarlyFraudWarningFraudTypeMisc                      RadarEarlyFraudWarningFraudType = "misc"
	RadarEarlyFraudWarningFraudTypeUnauthorizedUseOfCard     RadarEarlyFraudWarningFraudType = "unauthorized_use_of_card"
)

// Returns a list of early fraud warnings.
type RadarEarlyFraudWarningListParams struct {
	ListParams    `form:"*"`
	Charge        *string `form:"charge"`
	PaymentIntent *string `form:"payment_intent"`
}

// Retrieves the details of an early fraud warning that has previously been created.
//
// Please refer to the [early fraud warning](https://stripe.com/docs/api#early_fraud_warning_object) object reference for more details.
type RadarEarlyFraudWarningParams struct {
	Params `form:"*"`
}

// An early fraud warning indicates that the card issuer has notified us that a
// charge may be fraudulent.
//
// Related guide: [Early Fraud Warnings](https://stripe.com/docs/disputes/measuring#early-fraud-warnings).
type RadarEarlyFraudWarning struct {
	APIResource
	Actionable    bool                            `json:"actionable"`
	Charge        *Charge                         `json:"charge"`
	Created       int64                           `json:"created"`
	FraudType     RadarEarlyFraudWarningFraudType `json:"fraud_type"`
	ID            string                          `json:"id"`
	Livemode      bool                            `json:"livemode"`
	Object        string                          `json:"object"`
	PaymentIntent *PaymentIntent                  `json:"payment_intent"`
}

// RadarEarlyFraudWarningList is a list of EarlyFraudWarnings as retrieved from a list endpoint.
type RadarEarlyFraudWarningList struct {
	APIResource
	ListMeta
	// TODO: rename `Values` to `Data` in a future major version for consistency with other List structs
	Values []*RadarEarlyFraudWarning `json:"data"`
}
