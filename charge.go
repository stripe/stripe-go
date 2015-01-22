package stripe

import "encoding/json"

// Currency is the list of supported currencies.
// For more details see https://support.stripe.com/questions/which-currencies-does-stripe-support.
type Currency string

// FraudReport is the list of allowed values for reporting fraud.
// Allowed values are "fraudulent", "safe".
type FraudReport string

// ChargeParams is the set of parameters that can be used when creating or updating a charge.
// For more details see https://stripe.com/docs/api#create_charge and https://stripe.com/docs/api#update_charge.
type ChargeParams struct {
	Params
	Amount                 uint64
	Currency               Currency
	Customer, Token        string
	Card                   *CardParams
	Desc, Statement, Email string
	NoCapture              bool
	Fee                    uint64
	Fraud                  FraudReport
	Source                 *SourceParams
}

// ChargeListParams is the set of parameters that can be used when listing charges.
// For more details see https://stripe.com/docs/api#list_charges.
type ChargeListParams struct {
	ListParams
	Created  int64
	Customer string
}

// CaptureParams is the set of parameters that can be used when capturing a charge.
// For more details see https://stripe.com/docs/api#charge_capture.
type CaptureParams struct {
	Params
	Amount, Fee uint64
	Email       string
}

// Charge is the resource representing a Stripe charge.
// For more details see https://stripe.com/docs/api#charges.
type Charge struct {
	ID             string            `json:"id"`
	Live           bool              `json:"livemode"`
	Amount         uint64            `json:"amount"`
	Captured       bool              `json:"captured"`
	Card           *Card             `json:"card"`
	Created        int64             `json:"created"`
	Currency       Currency          `json:"currency"`
	Paid           bool              `json:"paid"`
	Refunded       bool              `json:"refunded"`
	Refunds        *RefundList       `json:"refunds"`
	AmountRefunded uint64            `json:"amount_refunded"`
	Tx             *Transaction      `json:"balance_transaction"`
	Customer       *Customer         `json:"customer"`
	Desc           string            `json:"description"`
	Dispute        *Dispute          `json:"dispute"`
	FailMsg        string            `json:"failure_message"`
	FailCode       string            `json:"failure_code"`
	Invoice        *Invoice          `json:"invoice"`
	Meta           map[string]string `json:"metadata"`
	Email          string            `json:"receipt_email"`
	Statement      string            `json:"statement_descriptor"`
	FraudDetails   *FraudDetails     `json:"fraud_details"`
	Status         string            `json:"status"`
	Source         *PaymentSource    `json:"source"`
}

// FraudDetails is the structure detailing fraud status.
type FraudDetails struct {
	UserReport   FraudReport `json:"user_report"`
	StripeReport FraudReport `json:"stripe_report"`
}

// UnmarshalJSON handles deserialization of a Charge.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *Charge) UnmarshalJSON(data []byte) error {
	type charge Charge
	var cc charge
	err := json.Unmarshal(data, &cc)
	if err == nil {
		*c = Charge(cc)
	} else {
		// the id is surrounded by "\" characters, so strip them
		c.ID = string(data[1 : len(data)-1])
	}

	return nil
}
