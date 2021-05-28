//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

type TaxDeductedAtSource struct {
	ID                        string `json:"id"`
	Object                    string `json:"object"`
	PeriodEnd                 int64  `json:"period_end"`
	PeriodStart               int64  `json:"period_start"`
	TaxDeductionAccountNumber string `json:"tax_deduction_account_number"`
}

// UnmarshalJSON handles deserialization of a TaxDeductedAtSource.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (t *TaxDeductedAtSource) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		t.ID = id
		return nil
	}

	type taxDeductedAtSource TaxDeductedAtSource
	var v taxDeductedAtSource
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*t = TaxDeductedAtSource(v)
	return nil
}
