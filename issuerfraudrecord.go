package stripe

// FraudType are strings that map to the fraud label category from the issuer.
type FraudType string

const (
	CardNeverReceived         FraudType = "card_never_received"
	FraudulentCardApplication FraudType = "fraudulent_card_application"
	MadeWithCounterfeitCard   FraudType = "made_with_counterfeit_card"
	MadeWithLostCard          FraudType = "made_with_lost_card"
	MadeWithStolenCard        FraudType = "made_with_stolen_card"
	Misc                      FraudType = "misc"
	UnauthorizedUseOfCard     FraudType = "unauthorized_use_of_card"
)

// IssuerFraudRecordListParams is the set of parameters that can be used when
// listing issuer fraud records. For more details see
// https://stripe.com/docs#list_issuer_fraud_records.
type IssuerFraudRecordListParams struct {
	ListParams `form:"*"`
	Charge     string `form:"-"`
}

// IssuerFraudRecordList is a list of issuer fraud records as retrieved from a
// list endpoint.
type IssuerFraudRecordList struct {
	ListMeta
	Values []*IssuerFraudRecord `json:"data"`
}

// IssuerFraudRecord is the resource representing an issuer fraud record. For
// more details see https://stripe.com/docs/api#issuer_fraud_records.
type IssuerFraudRecord struct {
	Charge    *Charge   `json:"charge"`
	Created   int64     `json:"created"`
	FraudType FraudType `json:"fraud_type"`
	ID        string    `json:"id"`
	Live      bool      `json:"livemode"`
	PostDate  int64     `json:"post_date"`
}
