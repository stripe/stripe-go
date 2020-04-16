package stripe

import "encoding/json"

// IssuingDisputeParams is the set of parameters that can be used when creating or updating an issuing dispute.
type IssuingDisputeParams struct {
	Params `form:"*"`
}

// IssuingDisputeListParams is the set of parameters that can be used when listing issuing dispute.
type IssuingDisputeListParams struct {
	ListParams `form:"*"`
}

// IssuingDispute is the resource representing an issuing dispute.
type IssuingDispute struct {
	APIResource
	ID       string `json:"id"`
	Livemode bool   `json:"livemode"`
	Object   string `json:"object"`
}

// IssuingDisputeList is a list of issuing disputes as retrieved from a list endpoint.
type IssuingDisputeList struct {
	APIResource
	ListMeta
	Data []*IssuingDispute `json:"data"`
}

// UnmarshalJSON handles deserialization of an IssuingDispute.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *IssuingDispute) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type issuingDispute IssuingDispute
	var v issuingDispute
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = IssuingDispute(v)
	return nil
}
