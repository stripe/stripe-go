// Package charge provides the /charges APIs
package charge

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
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
	charge := &stripe.Charge{}
	err := c.B.Call2("POST", "/charges", c.Key, params, charge)
	return charge, err
}

// Get returns the details of a charge.
// For more details see https://stripe.com/docs/api#retrieve_charge.
func Get(id string, params *stripe.ChargeParams) (*stripe.Charge, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.ChargeParams) (*stripe.Charge, error) {
	path := stripe.FormatURLPath("/charges/%s", id)
	charge := &stripe.Charge{}
	err := c.B.Call2("GET", path, c.Key, params, charge)
	return charge, err
}

// Update updates a charge's properties.
// For more details see https://stripe.com/docs/api#update_charge.
func Update(id string, params *stripe.ChargeParams) (*stripe.Charge, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.ChargeParams) (*stripe.Charge, error) {
	path := stripe.FormatURLPath("/charges/%s", id)
	charge := &stripe.Charge{}
	err := c.B.Call2("POST", path, c.Key, params, charge)
	return charge, err
}

// Capture captures a charge not yet captured.
// For more details see https://stripe.com/docs/api#charge_capture.
func Capture(id string, params *stripe.CaptureParams) (*stripe.Charge, error) {
	return getC().Capture(id, params)
}

func (c Client) Capture(id string, params *stripe.CaptureParams) (*stripe.Charge, error) {
	path := stripe.FormatURLPath("/charges/%s/capture", id)
	charge := &stripe.Charge{}
	err := c.B.Call2("POST", path, c.Key, params, charge)
	return charge, err
}

// List returns a list of charges.
// For more details see https://stripe.com/docs/api#list_charges.
func List(params *stripe.ChargeListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.ChargeListParams) *Iter {
	return &Iter{stripe.GetIter2(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.ChargeList{}
		err := c.B.CallRaw("GET", "/charges", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
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
		&stripe.ChargeParams{
			FraudDetails: &stripe.FraudDetailsParams{
				UserReport: stripe.String(string(stripe.ChargeFraudUserReportFraudulent)),
			},
		},
	)
}

// MarkSafe reports the charge as not-fraudulent.
func MarkSafe(id string) (*stripe.Charge, error) {
	return getC().MarkSafe(id)
}

func (c Client) MarkSafe(id string) (*stripe.Charge, error) {
	return c.Update(
		id,
		&stripe.ChargeParams{
			FraudDetails: &stripe.FraudDetailsParams{
				UserReport: stripe.String(string(stripe.ChargeFraudUserReportSafe)),
			},
		},
	)
}

// Update updates a charge's dispute.
// For more details see https://stripe.com/docs/api#update_dispute.
func UpdateDispute(id string, params *stripe.DisputeParams) (*stripe.Dispute, error) {
	return getC().UpdateDispute(id, params)
}

func (c Client) UpdateDispute(id string, params *stripe.DisputeParams) (*stripe.Dispute, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	dispute := &stripe.Dispute{}
	err := c.B.Call("POST", stripe.FormatURLPath("/charges/%s/dispute", id), c.Key, body, commonParams, dispute)

	return dispute, err
}

// Close dismisses a charge's dispute in the customer's favor.
// For more details see https://stripe.com/docs/api#close_dispute.
func CloseDispute(id string) (*stripe.Dispute, error) {
	return getC().CloseDispute(id)
}

func (c Client) CloseDispute(id string) (*stripe.Dispute, error) {
	dispute := &stripe.Dispute{}
	err := c.B.Call2("POST", stripe.FormatURLPath("/charges/%s/dispute/close", id), c.Key, nil, dispute)
	return dispute, err
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
