package stripe

import (
	"encoding/json"
)

// BillingPortalSessionParams is the set of parameters that can be used when creating a billing portal session.
type BillingPortalSessionParams struct {
	Params        `form:"*"`
	Configuration *string `form:"configuration"`
	Customer      *string `form:"customer"`
	OnBehalfOf    *string `form:"on_behalf_of"`
	ReturnURL     *string `form:"return_url"`
}

// BillingPortalSession is the resource representing a billing portal session.
type BillingPortalSession struct {
	APIResource
	Created       int64                       `json:"created"`
	Configuration *BillingPortalConfiguration `json:"configuration"`
	Customer      string                      `json:"customer"`
	ID            string                      `json:"id"`
	Livemode      bool                        `json:"livemode"`
	Object        string                      `json:"object"`
	OnBehalfOf    string                      `json:"on_behalf_of"`
	ReturnURL     string                      `json:"return_url"`
	URL           string                      `json:"url"`
}

// UnmarshalJSON handles deserialization of a billing portal session.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *BillingPortalSession) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type session BillingPortalSession
	var v session
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = BillingPortalSession(v)
	return nil
}
