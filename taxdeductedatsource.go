//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"time"
)

type TaxDeductedAtSource struct {
	// Unique identifier for the object.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The end of the invoicing period. This TDS applies to Stripe fees collected during this invoicing period.
	PeriodEnd time.Time `json:"period_end"`
	// The start of the invoicing period. This TDS applies to Stripe fees collected during this invoicing period.
	PeriodStart time.Time `json:"period_start"`
	// The TAN that was supplied to Stripe when TDS was assessed
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
	v := struct {
		PeriodEnd   int64 `json:"period_end"`
		PeriodStart int64 `json:"period_start"`
		*taxDeductedAtSource
	}{
		taxDeductedAtSource: (*taxDeductedAtSource)(t),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.PeriodEnd = time.Unix(v.PeriodEnd, 0)
	t.PeriodStart = time.Unix(v.PeriodStart, 0)
	return nil
}

// MarshalJSON handles serialization of a TaxDeductedAtSource.
// This custom marshaling is needed to handle the time fields correctly.
func (t TaxDeductedAtSource) MarshalJSON() ([]byte, error) {
	type taxDeductedAtSource TaxDeductedAtSource
	v := struct {
		PeriodEnd   int64 `json:"period_end"`
		PeriodStart int64 `json:"period_start"`
		taxDeductedAtSource
	}{
		taxDeductedAtSource: (taxDeductedAtSource)(t),
		PeriodEnd:           t.PeriodEnd.Unix(),
		PeriodStart:         t.PeriodStart.Unix(),
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, err
}
