//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

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
	Configuration *BillingPortalConfiguration `json:"configuration"`
	Created       int64                       `json:"created"`
	Customer      string                      `json:"customer"`
	ID            string                      `json:"id"`
	Livemode      bool                        `json:"livemode"`
	Object        string                      `json:"object"`
	OnBehalfOf    string                      `json:"on_behalf_of"`
	ReturnURL     string                      `json:"return_url"`
	URL           string                      `json:"url"`
}

// UnmarshalJSON handles deserialization of a BillingPortalSession.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (b *BillingPortalSession) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		b.ID = id
		return nil
	}

	type billingPortalSession BillingPortalSession
	var v billingPortalSession
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*b = BillingPortalSession(v)
	return nil
}
