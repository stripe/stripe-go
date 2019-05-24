package stripe

// RadarEarlyFraudWarningFraudType are strings that map to the type of fraud labelled by the issuer.
type RadarEarlyFraudWarningFraudType string

// List of values that RadarEarlyFraudWarningFraudType can take.
const (
	RadarEarlyFraudWarningFraudTypeCardNeverReceived         RadarEarlyFraudWarningFraudType = "card_never_received"
	RadarEarlyFraudWarningFraudTypeFraudulentCardApplication RadarEarlyFraudWarningFraudType = "fraudulent_card_application"
	RadarEarlyFraudWarningFraudTypeMadeWithCounterfeitCard   RadarEarlyFraudWarningFraudType = "made_with_counterfeit_card"
	RadarEarlyFraudWarningFraudTypeMadeWithLostCard          RadarEarlyFraudWarningFraudType = "made_with_lost_card"
	RadarEarlyFraudWarningFraudTypeMadeWithStolenCard        RadarEarlyFraudWarningFraudType = "made_with_stolen_card"
	RadarEarlyFraudWarningFraudTypeMisc                      RadarEarlyFraudWarningFraudType = "misc"
	RadarEarlyFraudWarningFraudTypeUnauthorizedUseOfCard     RadarEarlyFraudWarningFraudType = "unauthorized_use_of_card"
)

// RadarEarlyFraudWarningParams is the set of parameters that can be used when
// retrieving early fraud warnings. For more details see
// https://stripe.com/docs/api/early_fraud_warnings/retrieve.
type RadarEarlyFraudWarningParams struct {
	Params `form:"*"`
}

// RadarEarlyFraudWarningListParams is the set of parameters that can be used when
// listing early fraud warnings. For more details see
// https://stripe.com/docs/api/early_fraud_warnings/list.
type RadarEarlyFraudWarningListParams struct {
	ListParams `form:"*"`
	Charge     *string `form:"charge"`
}

// RadarEarlyFraudWarningList is a list of early fraud warnings as retrieved from a
// list endpoint.
type RadarEarlyFraudWarningList struct {
	ListMeta
	Values []*RadarEarlyFraudWarning `json:"data"`
}

// RadarEarlyFraudWarning is the resource representing an early fraud warning. For
// more details see https://stripe.com/docs/api/early_fraud_warnings/object.
type RadarEarlyFraudWarning struct {
	Actionable bool                            `json:"actionable"`
	Charge     *Charge                         `json:"charge"`
	Created    int64                           `json:"created"`
	FraudType  RadarEarlyFraudWarningFraudType `json:"fraud_type"`
	ID         string                          `json:"id"`
	Livemode   bool                            `json:"livemode"`
}
