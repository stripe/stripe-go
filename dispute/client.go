package dispute

import (
	"fmt"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

const (
	CreditNotProcessed   stripe.DisputeReason = "credit_not_processed"
	Duplicate            stripe.DisputeReason = "duplicate"
	Fraudulent           stripe.DisputeReason = "fraudulent"
	General              stripe.DisputeReason = "general"
	ProductNotReceived   stripe.DisputeReason = "product_not_received"
	ProductUnacceptable  stripe.DisputeReason = "product_unacceptable"
	SubscriptionCanceled stripe.DisputeReason = "subscription_canceled"
	Unrecognized         stripe.DisputeReason = "unrecognized"

	Lost                 stripe.DisputeStatus = "lost"
	NeedsResponse        stripe.DisputeStatus = "needs_response"
	ChargeRefunded       stripe.DisputeStatus = "charge_refunded"
	UnderReview          stripe.DisputeStatus = "under_review"
	WarningClosed        stripe.DisputeStatus = "warning_closed"
	WarningNeedsResponse stripe.DisputeStatus = "warning_needs_response"
	WarningUnderReview   stripe.DisputeStatus = "warning_under_review"
	Won                  stripe.DisputeStatus = "won"
)

// Client is used to invoke dispute-related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a dispute.
// For more details see https://stripe.com/docs/api#retrieve_dispute.
func Get(id string, params *stripe.DisputeParams) (*stripe.Dispute, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.DisputeParams) (*stripe.Dispute, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	dispute := &stripe.Dispute{}
	err := c.B.Call("GET", "/disputes/"+id, c.Key, body, commonParams, dispute)

	return dispute, err
}

// List returns a list of disputes.
// For more details see https://stripe.com/docs/api#list_disputes.
func List(params *stripe.DisputeListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.DisputeListParams) *Iter {
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
		list := &stripe.DisputeList{}
		err := c.B.Call("GET", "/disputes", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Disputes.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Dispute returns the most recent Dispute
// visited by a call to Next.
func (i *Iter) Dispute() *stripe.Dispute {
	return i.Current().(*stripe.Dispute)
}

// Update updates a dispute.
// For more details see https://stripe.com/docs/api#update_dispute.
func Update(id string, params *stripe.DisputeParams) (*stripe.Dispute, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.DisputeParams) (*stripe.Dispute, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	dispute := &stripe.Dispute{}
	err := c.B.Call("POST", fmt.Sprintf("/disputes/%v", id), c.Key, body, commonParams, dispute)

	return dispute, err
}

// Close dismisses a dispute in the customer's favor.
// For more details see https://stripe.com/docs/api#close_dispute.
func Close(id string) (*stripe.Dispute, error) {
	return getC().Close(id)
}

func (c Client) Close(id string) (*stripe.Dispute, error) {
	dispute := &stripe.Dispute{}
	err := c.B.Call("POST", fmt.Sprintf("/disputes/%v/close", id), c.Key, nil, nil, dispute)

	return dispute, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
