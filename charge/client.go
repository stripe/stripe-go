package charge

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke /charges APIs.
type Client struct {
	B   Backend
	Tok string
}

// Create POSTs new charges.
// For more details see https://stripe.com/docs/api#create_charge.
func Create(params *ChargeParams) (*Charge, error) {
	return getC().Create(params)
}

func (c Client) Create(params *ChargeParams) (*Charge, error) {
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

	token := c.Tok
	if params.Fee > 0 {
		if len(params.AccessToken) == 0 {
			err := errors.New("Invalid charge params: an access token is required for application fees")
			return nil, err
		}

		body.Add("application_fee", strconv.FormatUint(params.Fee, 10))
		token = params.AccessToken
	}

	params.AppendTo(body)

	charge := &Charge{}
	err := c.B.Call("POST", "/charges", token, body, charge)

	return charge, err
}

// Get returns the details of a charge.
// For more details see https://stripe.com/docs/api#retrieve_charge.
func Get(id string, params *ChargeParams) (*Charge, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *ChargeParams) (*Charge, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	charge := &Charge{}
	err := c.B.Call("GET", "/charges/"+id, c.Tok, body, charge)

	return charge, err
}

// Update updates a charge's properties.
// For more details see https://stripe.com/docs/api#update_charge.
func Update(id string, params *ChargeParams) (*Charge, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *ChargeParams) (*Charge, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if len(params.Desc) > 0 {
			body.Add("description", params.Desc)
		}

		params.AppendTo(body)
	}

	charge := &Charge{}
	err := c.B.Call("POST", "/charges/"+id, c.Tok, body, charge)

	return charge, err
}

// RefundCharge refunds a charge previously created.
// For more details see https://stripe.com/docs/api#refund_charge.
func RefundCharge(params *RefundParams) (*Refund, error) {
	return getC().Refund(params)
}

func (c Client) Refund(params *RefundParams) (*Refund, error) {
	body := &url.Values{}

	if params.Amount > 0 {
		body.Add("amount", strconv.FormatUint(params.Amount, 10))
	}

	if params.Fee {
		body.Add("refund_application_fee", strconv.FormatBool(params.Fee))
	}

	params.AppendTo(body)

	refund := &Refund{}
	err := c.B.Call("POST", fmt.Sprintf("/charges/%v/refunds", params.Charge), c.Tok, body, refund)

	return refund, err
}

// Capture captures a previously created charge with NoCapture set to true.
// For more details see https://stripe.com/docs/api#charge_capture.
func Capture(id string, params *CaptureParams) (*Charge, error) {
	return getC().Capture(id, params)
}

func (c Client) Capture(id string, params *CaptureParams) (*Charge, error) {
	var body *url.Values
	token := c.Tok

	if params != nil {
		body = &url.Values{}

		if params.Amount > 0 {
			body.Add("amount", strconv.FormatUint(params.Amount, 10))
		}

		if len(params.Email) > 0 {
			body.Add("receipt_email", params.Email)
		}

		if params.Fee > 0 {
			if len(params.AccessToken) == 0 {
				err := errors.New("Invalid charge params: an access token is required for application fees")
				return nil, err
			}

			body.Add("application_fee", strconv.FormatUint(params.Fee, 10))
			token = params.AccessToken
		}

		params.AppendTo(body)
	}

	charge := &Charge{}
	err := c.B.Call("POST", fmt.Sprintf("/charges/%v/capture", id), token, body, charge)

	return charge, err
}

// List returns a list of charges.
// For more details see https://stripe.com/docs/api#list_charges.
func List(params *ChargeListParams) (*ChargeList, error) {
	return getC().List(params)
}

func (c Client) List(params *ChargeListParams) (*ChargeList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}

		if len(params.Customer) > 0 {
			body.Add("customer", params.Customer)
		}

		params.AppendTo(body)
	}

	list := &ChargeList{}
	err := c.B.Call("GET", "/charges", c.Tok, body, list)

	return list, err
}

func getC() Client {
	return Client{GetBackend(), Key}
}
