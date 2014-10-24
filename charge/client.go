// Package charge provides the /charges APIs
package charge

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"

	stripe "github.com/stripe/stripe-go"
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

	if len(params.Customer) > 0 {
		body.Add("customer", params.Customer)
	} else if len(params.Token) > 0 {
		body.Add("card", params.Token)
	} else if params.Card != nil {
		if len(params.Card.Token) > 0 {
			body.Add("card", params.Card.Token)
		} else {
			params.Card.AppendDetails(body, true)
		}
	} else {
		err := errors.New("Invalid charge params: either customer, card Tok or card need to be set")
		return nil, err
	}

	if len(params.Desc) > 0 {
		body.Add("description", params.Desc)
	}

	if len(params.Statement) > 0 {
		body.Add("statement_description", params.Statement)
	}

	if len(params.Email) > 0 {
		body.Add("receipt_email", params.Email)
	}

	body.Add("capture", strconv.FormatBool(!params.NoCapture))

	token := c.Key
	if params.Fee > 0 {
		body.Add("application_fee", strconv.FormatUint(params.Fee, 10))
	}

	params.AppendTo(body)

	charge := &stripe.Charge{}
	err := c.B.Call("POST", "/charges", token, body, charge)

	return charge, err
}

// Get returns the details of a charge.
// For more details see https://stripe.com/docs/api#retrieve_charge.
func Get(id string, params *stripe.ChargeParams) (*stripe.Charge, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.ChargeParams) (*stripe.Charge, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	charge := &stripe.Charge{}
	err := c.B.Call("GET", "/charges/"+id, c.Key, body, charge)

	return charge, err
}

// Update updates a charge's properties.
// For more details see https://stripe.com/docs/api#update_charge.
func Update(id string, params *stripe.ChargeParams) (*stripe.Charge, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.ChargeParams) (*stripe.Charge, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if len(params.Desc) > 0 {
			body.Add("description", params.Desc)
		}

		params.AppendTo(body)
	}

	charge := &stripe.Charge{}
	err := c.B.Call("POST", "/charges/"+id, c.Key, body, charge)

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

	if params != nil {
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
	err := c.B.Call("POST", fmt.Sprintf("/charges/%v/capture", id), token, body, charge)

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
		err := c.B.Call("GET", "/charges", c.Key, &b, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
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
	return Client{stripe.GetBackend(), stripe.Key}
}
