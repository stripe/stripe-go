//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

type TaxCodeListParams struct {
	ListParams `form:"*"`
}
type TaxCodeParams struct {
	Params `form:"*"`
}
type TaxCode struct {
	APIResource
	Description string `json:"description"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Object      string `json:"object"`
}
type TaxCodeList struct {
	APIResource
	ListMeta
	Data []*TaxCode `json:"data"`
}

// UnmarshalJSON handles deserialization of a TaxCode.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (t *TaxCode) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		t.ID = id
		return nil
	}

	type taxCode TaxCode
	var v taxCode
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*t = TaxCode(v)
	return nil
}
