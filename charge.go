package stripe

import (
	"encoding/json"
)

// Currency is the list of supported currencies.
// For more details see https://support.stripe.com/questions/which-currencies-does-stripe-support.
type Currency string

// FraudReport is the list of allowed values for reporting fraud.
// Allowed values are "fraudulent", "safe".
type FraudReport string

// ChargeParams is the set of parameters that can be used when creating or updating a charge.
// For more details see https://stripe.com/docs/api#create_charge and https://stripe.com/docs/api#update_charge.
type ChargeParams struct {
	Params              `form:"*"`
	Amount              uint64              `form:"amount"`
	ApplicationFee      uint64              `form:"application_fee"`
	Currency            Currency            `form:"currency"`
	Customer            string              `form:"customer"`
	Description         string              `form:"description"`
	Destination         *DestinationParams  `form:"destination"`
	ExchangeRate        float64             `form:"exchange_rate"`
	FraudDetails        *FraudDetailsParams `form:"fraud_details"`
	NoCapture           bool                `form:"capture,invert"`
	OnBehalfOf          string              `form:"on_behalf_of"`
	ReceiptEmail        string              `form:"receipt_email"`
	Shipping            *ShippingDetails    `form:"shipping"`
	Source              *SourceParams       `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
	StatementDescriptor string              `form:"statement_descriptor"`
	TransferGroup       string              `form:"transfer_group"`
}

// SetSource adds valid sources to a ChargeParams object,
// returning an error for unsupported sources.
func (p *ChargeParams) SetSource(sp interface{}) error {
	source, err := SourceParamsFor(sp)
	p.Source = source
	return err
}

type DestinationParams struct {
	Account string `form:"account"`
	Amount  uint64 `form:"amount"`
}

// FraudDetailsParams provides information on the fraud details for a charge.
type FraudDetailsParams struct {
	UserReport FraudReport `form:"user_report"`
}

// ChargeListParams is the set of parameters that can be used when listing charges.
// For more details see https://stripe.com/docs/api#list_charges.
type ChargeListParams struct {
	ListParams    `form:"*"`
	Created       int64             `form:"created"`
	CreatedRange  *RangeQueryParams `form:"created"`
	Customer      string            `form:"customer"`
	TransferGroup string            `form:"transfer_group"`
}

// CaptureParams is the set of parameters that can be used when capturing a charge.
// For more details see https://stripe.com/docs/api#charge_capture.
type CaptureParams struct {
	Params              `form:"*"`
	Amount              uint64  `form:"amount"`
	ApplicationFee      uint64  `form:"application_fee"`
	ExchangeRate        float64 `form:"exchange_rate"`
	ReceiptEmail        string  `form:"receipt_email"`
	StatementDescriptor string  `form:"statement_descriptor"`
}

// Charge is the resource representing a Stripe charge.
// For more details see https://stripe.com/docs/api#charges.
type Charge struct {
	Amount              uint64              `json:"amount"`
	AmountRefunded      uint64              `json:"amount_refunded"`
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
	UserReport   FraudReport `json:"user_report"`
	StripeReport FraudReport `json:"stripe_report"`
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
	Address        Address `json:"address" form:"address"`
	Carrier        string  `json:"carrier" form:"carrier"`
	Name           string  `json:"name" form:"name"`
	Phone          string  `json:"phone" form:"phone"`
	TrackingNumber string  `json:"tracking_number" form:"tracking_number"`
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
