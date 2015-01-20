// Package charge provides the /charges APIs
package charge

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"

	stripe "github.com/stripe/stripe-go"
)

const (
	ReportFraudulent stripe.FraudReport = "fraudulent"
	ReportSafe       stripe.FraudReport = "safe"
)

// Client is used to invoke /charges APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs new charges.
// For more details see https://stripe.com/docs/api#create_charge.
func New(params *stripe.ChargeParams) (*stripe.Charge, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.ChargeParams) (*stripe.Charge, error) {
	body := &url.Values{
		"amount":   {strconv.FormatUint(params.Amount, 10)},
		"currency": {string(params.Currency)},
	}

	// Handle multiple ways of specifying the Chargeable object:
	// 1) Passed card details in Card params.Card
	// 2) Passed card token in params.Token
	// 3) Passed card details in params.Source.Card
	// 4) Passed card token in params.Source.Token
	// 5) Passed source ID in params.Source.ID

	var tokenParam string
	var card *stripe.CardParams
	var sourceId string

	if len(params.Token) > 0 {
		tokenParam = params.Token
	} else if params.Source != nil && len(params.Source.Token) > 0 {
		tokenParam = params.Source.Token
	}

	if params.Card != nil {
		card = params.Card
	} else if params.Source != nil && params.Source.Card != nil {
		card = params.Source.Card
	}

	if params.Source != nil && len(params.Source.ID) > 0 {
		sourceId = params.Source.ID
	}

	isCard := !(len(sourceId) > 0)

	if len(tokenParam) > 0 {
		body.Add("source", tokenParam)
	} else if len(params.Customer) > 0 {
		body.Add("customer", params.Customer)
		if card != nil && len(card.Token) > 0 {
			body.Add("source", card.Token)
		}
	} else if card != nil {
		if len(card.Token) > 0 {
			body.Add("card", card.Token)
		} else {
			card.AppendDetails(body, true)
		}
	} else if len(sourceId) > 0 {
		body.Add("source", sourceId)
	} else {
		err := errors.New("Invalid charge params: either customer, card Tok, card, source Tok, or source need to be set")
		return nil, err
	}

	if len(params.Desc) > 0 {
		body.Add("description", params.Desc)
	}

	if len(params.Statement) > 0 {
		body.Add("statement_descriptor", params.Statement)
	}

	if len(params.Email) > 0 {
		body.Add("receipt_email", params.Email)
	}

	if isCard {
		body.Add("capture", strconv.FormatBool(!params.NoCapture))
	}

	token := c.Key
	if params.Fee > 0 {
		body.Add("application_fee", strconv.FormatUint(params.Fee, 10))
	}

	params.AppendTo(body)

	charge := &stripe.Charge{}
	err := c.B.Call("POST", "/charges", token, body, &params.Params, charge)

	return charge, err
}

// Get returns the details of a charge.
// For more details see https://stripe.com/docs/api#retrieve_charge.
func Get(id string, params *stripe.ChargeParams) (*stripe.Charge, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.ChargeParams) (*stripe.Charge, error) {
	var body *url.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &url.Values{}
		params.AppendTo(body)
	}

	charge := &stripe.Charge{}
	err := c.B.Call("GET", "/charges/"+id, c.Key, body, commonParams, charge)

	return charge, err
}

// Update updates a charge's properties.
// For more details see https://stripe.com/docs/api#update_charge.
func Update(id string, params *stripe.ChargeParams) (*stripe.Charge, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.ChargeParams) (*stripe.Charge, error) {
	var body *url.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &url.Values{}

		if len(params.Desc) > 0 {
			body.Add("description", params.Desc)
		}

		if len(params.Fraud) > 0 {
			body.Add("fraud_details[user_report]", string(params.Fraud))
		}

		params.AppendTo(body)
	}

	charge := &stripe.Charge{}
	err := c.B.Call("POST", "/charges/"+id, c.Key, body, commonParams, charge)

	return charge, err
}

// Capture captures a previously created charge with NoCapture set to true.
// For more details see https://stripe.com/docs/api#charge_capture.
func Capture(id string, params *stripe.CaptureParams) (*stripe.Charge, error) {
	return getC().Capture(id, params)
}

func (c Client) Capture(id string, params *stripe.CaptureParams) (*stripe.Charge, error) {
	var body *url.Values
	token := c.Key
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &url.Values{}

		if params.Amount > 0 {
			body.Add("amount", strconv.FormatUint(params.Amount, 10))
		}

		if len(params.Email) > 0 {
			body.Add("receipt_email", params.Email)
		}

		if params.Fee > 0 {
			body.Add("application_fee", strconv.FormatUint(params.Fee, 10))
		}

		params.AppendTo(body)
	}

	charge := &stripe.Charge{}

	err := c.B.Call("POST", fmt.Sprintf("/charges/%v/capture", id), token, body, commonParams, charge)

	return charge, err
}

// List returns a list of charges.
// For more details see https://stripe.com/docs/api#list_charges.
func List(params *stripe.ChargeListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.ChargeListParams) *Iter {
	type chargeList struct {
		stripe.ListMeta
		Values []*stripe.Charge `json:"data"`
	}

	var body *url.Values
	var lp *stripe.ListParams

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}

		if len(params.Customer) > 0 {
			body.Add("customer", params.Customer)
		}

		params.AppendTo(body)
		lp = &params.ListParams
	}

	return &Iter{stripe.GetIter(lp, body, func(b url.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &chargeList{}
		err := c.B.Call("GET", "/charges", c.Key, &b, nil, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// MarkFraudulent reports the charge as fraudulent.
func MarkFraudulent(id string) (*stripe.Charge, error) {
	return getC().MarkFraudulent(id)
}

func (c Client) MarkFraudulent(id string) (*stripe.Charge, error) {
	return c.Update(
		id,
		&stripe.ChargeParams{Fraud: ReportFraudulent})
}

// MarkSafe reports the charge as not-fraudulent.
func MarkSafe(id string) (*stripe.Charge, error) {
	return getC().MarkSafe(id)
}

func (c Client) MarkSafe(id string) (*stripe.Charge, error) {
	return c.Update(
		id,
		&stripe.ChargeParams{Fraud: ReportSafe})
}

// Iter is an iterator for lists of Charges.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Charge returns the most recent Charge
// visited by a call to Next.
func (i *Iter) Charge() *stripe.Charge {
	return i.Current().(*stripe.Charge)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
