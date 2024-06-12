//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

type PlatformTaxFee struct {
	// The Connected account that incurred this charge.
	Account string `json:"account"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The payment object that caused this tax to be inflicted.
	SourceTransaction string `json:"source_transaction"`
	// The type of tax (VAT).
	Type string `json:"type"`
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
