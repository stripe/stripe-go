//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

type PlatformTaxFee struct {
	Account           string `json:"account"`
	ID                string `json:"id"`
	Object            string `json:"object"`
	SourceTransaction string `json:"source_transaction"`
	Type              string `json:"type"`
}

// UnmarshalJSON handles deserialization of a PlatformTaxFee.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *PlatformTaxFee) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type platformTaxFee PlatformTaxFee
	var v platformTaxFee
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = PlatformTaxFee(v)
	return nil
}
