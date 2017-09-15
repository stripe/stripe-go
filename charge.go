package stripe

import (
	"encoding/json"

	"github.com/stripe/stripe-go/form"
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
	Params        `form:"*"`
	Amount        uint64             `form:"amount"`
	Currency      Currency           `form:"currency"`
	Customer      string             `form:"customer"`
	Token         string             `form:"-"` // Does not appear to be used?
	Desc          string             `form:"description"`
	Statement     string             `form:"statement_descriptor"`
	Email         string             `form:"receipt_email"`
	Dest          string             `form:"-"` // Handled in custom AppendTo below
	Destination   *DestinationParams `form:"destination"`
	NoCapture     bool               `form:"capture,invert"`
	Fee           uint64             `form:"application_fee"`
	Fraud         FraudReport        `form:"-"` // Handled in custom AppendTo below
	Source        *SourceParams      `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
	Shipping      *ShippingDetails   `form:"shipping"`
	TransferGroup string             `form:"transfer_group"`
	OnBehalfOf    string             `form:"on_behalf_of"`
}

// AppendTo implements some custom encoding logic for ChargeParams to support
// the deprecated Dest field (please use Destination instead) and the deviant
// Fraud field.
func (p *ChargeParams) AppendTo(body *form.Values, keyParts []string) {
	// TODO: Stop supporting this field.
	if len(p.Dest) > 0 {
		body.Add(form.FormatKey(append(keyParts, "destination", "account")), p.Dest)
	}

	// This is bad in that it's unnecessarily divergent. We should change it to
	// a subparams struct called FraudDetails.
	//
	// TODO: Change to a subparameter struct for FraudDetails.
	if len(p.Fraud) > 0 {
		body.Add(form.FormatKey(append(keyParts, "fraud_details", "user_report")), string(p.Fraud))
	}
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
	Params    `form:"*"`
	Amount    uint64 `form:"amount"`
	Fee       uint64 `form:"application_fee"`
	Email     string `form:"receipt_email"`
	Statement string `form:"statement_descriptor"`
}

// Charge is the resource representing a Stripe charge.
// For more details see https://stripe.com/docs/api#charges.
type Charge struct {
	Amount         uint64            `json:"amount"`
	AmountRefunded uint64            `json:"amount_refunded"`
	Application    *Application      `json:"application"`
	Captured       bool              `json:"captured"`
	Created        int64             `json:"created"`
	Currency       Currency          `json:"currency"`
	Customer       *Customer         `json:"customer"`
	Desc           string            `json:"description"`
	Dest           *Account          `json:"destination"`
	Dispute        *Dispute          `json:"dispute"`
	Email          string            `json:"receipt_email"`
	ReceiptNumber  string            `json:"receipt_number"`
	FailCode       string            `json:"failure_code"`
	FailMsg        string            `json:"failure_message"`
	Fee            *Fee              `json:"application_fee"`
	FraudDetails   *FraudDetails     `json:"fraud_details"`
	ID             string            `json:"id"`
	Invoice        *Invoice          `json:"invoice"`
	Live           bool              `json:"livemode"`
	Meta           map[string]string `json:"metadata"`
	Outcome        *ChargeOutcome    `json:"outcome"`
	Paid           bool              `json:"paid"`
	Refunded       bool              `json:"refunded"`
	Refunds        *RefundList       `json:"refunds"`
	Review         *Review           `json:"review"`
	Shipping       *ShippingDetails  `json:"shipping"`
	Source         *PaymentSource    `json:"source"`
	SourceTransfer *Transfer         `json:"source_transfer"`
	Statement      string            `json:"statement_descriptor"`
	Status         string            `json:"status"`
	Transfer       *Transfer         `json:"transfer"`
	TransferGroup  string            `json:"transfer_group"`
	Tx             *Transaction      `json:"balance_transaction"`
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
	Values []*Charge `json:"data"`
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
	Name     string  `json:"name" form:"name"`
	Address  Address `json:"address" form:"address"`
	Phone    string  `json:"phone" form:"phone"`
	Tracking string  `json:"tracking_number" form:"tracking_number"`
	Carrier  string  `json:"carrier" form:"carrier"`
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
