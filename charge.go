package stripe

import (
	"encoding/json"
)

// ChargeFraudUserReport is the list of allowed values for reporting fraud.
type ChargeFraudUserReport string

const (
	ChargeFraudUserReportFraudulent ChargeFraudUserReport = "fraudulent"
	ChargeFraudUserReportSafe       ChargeFraudUserReport = "safe"
)

// ChargeFraudStripeReport is the list of allowed values for reporting fraud.
type ChargeFraudStripeReport string

const (
	ChargeFraudStripeReportFraudulent ChargeFraudStripeReport = "fraudulent"
)

// ChargeParams is the set of parameters that can be used when creating or updating a charge.
// For more details see https://stripe.com/docs/api#create_charge and https://stripe.com/docs/api#update_charge.
type ChargeParams struct {
	Params              `form:"*"`
	Amount              *int64                 `form:"amount"`
	ApplicationFee      *int64                 `form:"application_fee"`
	Capture             *bool                  `form:"capture"`
	Currency            *string                `form:"currency"`
	Customer            *string                `form:"customer"`
	Description         *string                `form:"description"`
	Destination         *DestinationParams     `form:"destination"`
	ExchangeRate        *float64               `form:"exchange_rate"`
	FraudDetails        *FraudDetailsParams    `form:"fraud_details"`
	OnBehalfOf          *string                `form:"on_behalf_of"`
	ReceiptEmail        *string                `form:"receipt_email"`
	Shipping            *ShippingDetailsParams `form:"shipping"`
	Source              *SourceParams          `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
	StatementDescriptor *string                `form:"statement_descriptor"`
	TransferGroup       *string                `form:"transfer_group"`
}

// ShippingDetailsParams is the structure containing shipping information as parameters
type ShippingDetailsParams struct {
	Address        *AddressParams `form:"address"`
	Carrier        *string        `form:"carrier"`
	Name           *string        `form:"name"`
	Phone          *string        `form:"phone"`
	TrackingNumber *string        `form:"tracking_number"`
}

// SetSource adds valid sources to a ChargeParams object,
// returning an error for unsupported sources.
func (p *ChargeParams) SetSource(sp interface{}) error {
	source, err := SourceParamsFor(sp)
	p.Source = source
	return err
}

type DestinationParams struct {
	Account *string `form:"account"`
	Amount  *int64  `form:"amount"`
}

// FraudDetailsParams provides information on the fraud details for a charge.
type FraudDetailsParams struct {
	UserReport *string `form:"user_report"`
}

// ChargeListParams is the set of parameters that can be used when listing charges.
// For more details see https://stripe.com/docs/api#list_charges.
type ChargeListParams struct {
	ListParams    `form:"*"`
	Created       *int64            `form:"created"`
	CreatedRange  *RangeQueryParams `form:"created"`
	Customer      *string           `form:"customer"`
	TransferGroup *string           `form:"transfer_group"`
}

// CaptureParams is the set of parameters that can be used when capturing a charge.
// For more details see https://stripe.com/docs/api#charge_capture.
type CaptureParams struct {
	Params              `form:"*"`
	Amount              *int64   `form:"amount"`
	ApplicationFee      *int64   `form:"application_fee"`
	ExchangeRate        *float64 `form:"exchange_rate"`
	ReceiptEmail        *string  `form:"receipt_email"`
	StatementDescriptor *string  `form:"statement_descriptor"`
}

// Charge is the resource representing a Stripe charge.
// For more details see https://stripe.com/docs/api#charges.
type Charge struct {
	Amount              int64               `json:"amount"`
	AmountRefunded      int64               `json:"amount_refunded"`
	Application         *Application        `json:"application"`
	ApplicationFee      *ApplicationFee     `json:"application_fee"`
	BalanceTransaction  *BalanceTransaction `json:"balance_transaction"`
	Captured            bool                `json:"captured"`
	Created             int64               `json:"created"`
	Currency            Currency            `json:"currency"`
	Customer            *Customer           `json:"customer"`
	Description         string              `json:"description"`
	Destination         *Account            `json:"destination"`
	Dispute             *Dispute            `json:"dispute"`
	FailureCode         string              `json:"failure_code"`
	FailureMessage      string              `json:"failure_message"`
	FraudDetails        *FraudDetails       `json:"fraud_details"`
	ID                  string              `json:"id"`
	Invoice             *Invoice            `json:"invoice"`
	Livemode            bool                `json:"livemode"`
	Metadata            map[string]string   `json:"metadata"`
	Outcome             *ChargeOutcome      `json:"outcome"`
	Paid                bool                `json:"paid"`
	ReceiptEmail        string              `json:"receipt_email"`
	ReceiptNumber       string              `json:"receipt_number"`
	Refunded            bool                `json:"refunded"`
	Refunds             *RefundList         `json:"refunds"`
	Review              *Review             `json:"review"`
	Shipping            *ShippingDetails    `json:"shipping"`
	Source              *PaymentSource      `json:"source"`
	SourceTransfer      *Transfer           `json:"source_transfer"`
	StatementDescriptor string              `json:"statement_descriptor"`
	Status              string              `json:"status"`
	Transfer            *Transfer           `json:"transfer"`
	TransferGroup       string              `json:"transfer_group"`
}

// UnmarshalJSON handles deserialization of a charge.
// This custom unmarshaling is needed because the resulting
// property may be an ID or the full struct if it was expanded.
func (c *Charge) UnmarshalJSON(data []byte) error {
	type charge Charge
	var cc charge
	err := json.Unmarshal(data, &cc)
	if err == nil {
		*c = Charge(cc)
	} else {
		// the ID is surrounded by "\" characters, so strip them
		c.ID = string(data[1 : len(data)-1])
	}
	return nil
}

// ChargeList is a list of charges as retrieved from a list endpoint.
type ChargeList struct {
	ListMeta
	Data []*Charge `json:"data"`
}

// FraudDetails is the structure detailing fraud status.
type FraudDetails struct {
	UserReport   ChargeFraudUserReport   `json:"user_report"`
	StripeReport ChargeFraudStripeReport `json:"stripe_report"`
}

// ChargeOutcomeRule tells you the Radar rule that blocked the charge, if any.
type ChargeOutcomeRule struct {
	Action    string `json:"action"`
	ID        string `json:"id"`
	Predicate string `json:"predicate"`
}

// Outcome is the charge's outcome that details whether a payment
// was accepted and why.
type ChargeOutcome struct {
	NetworkStatus string             `json:"network_status"`
	Reason        string             `json:"reason"`
	RiskLevel     string             `json:"risk_level"`
	Rule          *ChargeOutcomeRule `json:"rule"`
	SellerMessage string             `json:"seller_message"`
	Type          string             `json:"type"`
}

// ShippingDetails is the structure containing shipping information.
type ShippingDetails struct {
	Address        *Address `json:"address"`
	Carrier        string   `json:"carrier"`
	Name           string   `json:"name"`
	Phone          string   `json:"phone"`
	TrackingNumber string   `json:"tracking_number"`
}

var depth int = -1

// UnmarshalJSON handles deserialization of a ChargeOutcomeRule.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *ChargeOutcomeRule) UnmarshalJSON(data []byte) error {
	type chargeOutcomeRule ChargeOutcomeRule
	var cc chargeOutcomeRule
	err := json.Unmarshal(data, &cc)
	if err == nil {
		*c = ChargeOutcomeRule(cc)
	} else {
		// the id is surrounded by "\" characters, so strip them
		c.ID = string(data[1 : len(data)-1])
	}

	return nil
}
