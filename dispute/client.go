// package dispute provides the dispute-related APIs
package dispute

import (
	"fmt"
	"net/url"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke dispute-related APIs.
type Client struct {
	B   Backend
	Key string
}

// Update updates a charge's dispute.
// For more details see https://stripe.com/docs/api#update_dispute.
func Update(id string, params *DisputeParams) (*Dispute, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *DisputeParams) (*Dispute, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if len(params.Evidence) > 0 {
			body.Add("evidence", params.Evidence)
		}

		params.AppendTo(body)
	}

	dispute := &Dispute{}
	err := c.B.Call("POST", fmt.Sprintf("/charges/%v/dispute", id), c.Key, body, dispute)

	return dispute, err
}

// Close dismisses a dispute in the customer's favor.
// For more details see https://stripe.com/docs/api#close_dispute.
func Close(id string) (*Dispute, error) {
	return getC().Close(id)
}

func (c Client) Close(id string) (*Dispute, error) {
	dispute := &Dispute{}
	err := c.B.Call("POST", fmt.Sprintf("/charges/%v/dispute/close", id), c.Key, nil, dispute)

	return dispute, err
}

func getC() Client {
	return Client{GetBackend(), Key}
}
