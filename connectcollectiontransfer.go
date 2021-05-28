//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

type ConnectCollectionTransfer struct {
	Amount      int64    `json:"amount"`
	Currency    Currency `json:"currency"`
	Destination *Account `json:"destination"`
	ID          string   `json:"id"`
	Livemode    bool     `json:"livemode"`
	Object      string   `json:"object"`
}

// UnmarshalJSON handles deserialization of a ConnectCollectionTransfer.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *ConnectCollectionTransfer) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type connectCollectionTransfer ConnectCollectionTransfer
	var v connectCollectionTransfer
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = ConnectCollectionTransfer(v)
	return nil
}
