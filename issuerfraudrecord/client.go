package issuerfraudrecord

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to interact with the /issuer_fraud_records API.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of an issuer fraud record.
// For more details see https://stripe.com/docs/api#retrieve_issuer_fraud_record.
func Get(id string) (*stripe.IssuerFraudRecord, error) {
	return getC().Get(id)
}

// Get returns the details of an issuer fraud record on a client.
// For more details see https://stripe.com/docs/api#retrieve_issuer_fraud_record.
func (c Client) Get(id string) (*stripe.IssuerFraudRecord, error) {
	ifr := &stripe.IssuerFraudRecord{}
	err := c.B.Call("GET", stripe.FormatURLPath("/issuer_fraud_records/%s", id), c.Key, nil, nil, ifr)
	return ifr, err
}

// List returns a list of issuer fraud records.
// For more details see https://stripe.com/docs/api#list_issuer_fraud_records.
func List(params *stripe.IssuerFraudRecordListParams) *Iter {
	return getC().List(params)
}

// List returns a list of issuer fraud records on a client.
// For more details see https://stripe.com/docs/api#list_issuer_fraud_records.
func (c Client) List(params *stripe.IssuerFraudRecordListParams) *Iter {
	var body *form.Values
	var lp *stripe.ListParams
	var p *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		lp = &params.ListParams
		p = params.ToParams()
	}

	return &Iter{stripe.GetIter(lp, body, func(b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.IssuerFraudRecordList{}
		err := c.B.Call("GET", "/issuer_fraud_records", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Topups.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// IssuerFraudRecord returns the most recent issuer fraud record visited by a
// call to Next.
func (i *Iter) IssuerFraudRecord() *stripe.IssuerFraudRecord {
	return i.Current().(*stripe.IssuerFraudRecord)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
