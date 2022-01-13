//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The high-level tax type, such as `vat` or `sales_tax`.
type TaxRateTaxType string

// List of values that TaxRateTaxType can take
const (
	TaxRateTaxTypeGST      TaxRateTaxType = "gst"
	TaxRateTaxTypeHST      TaxRateTaxType = "hst"
	TaxRateTaxTypeJct      TaxRateTaxType = "jct"
	TaxRateTaxTypePST      TaxRateTaxType = "pst"
	TaxRateTaxTypeQST      TaxRateTaxType = "qst"
	TaxRateTaxTypeRST      TaxRateTaxType = "rst"
	TaxRateTaxTypeSalesTax TaxRateTaxType = "sales_tax"
	TaxRateTaxTypeVAT      TaxRateTaxType = "vat"
)

// Returns a list of your tax rates. Tax rates are returned sorted by creation date, with the most recently created tax rates appearing first.
type TaxRateListParams struct {
	ListParams   `form:"*"`
	Active       *bool             `form:"active"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Inclusive    *bool             `form:"inclusive"`
}

// Creates a new tax rate.
type TaxRateParams struct {
	Params       `form:"*"`
	Active       *bool    `form:"active"`
	Country      *string  `form:"country"`
	Description  *string  `form:"description"`
	DisplayName  *string  `form:"display_name"`
	Inclusive    *bool    `form:"inclusive"`
	Jurisdiction *string  `form:"jurisdiction"`
	Percentage   *float64 `form:"percentage"`
	State        *string  `form:"state"`
	TaxType      *string  `form:"tax_type"`
}

// Tax rates can be applied to [invoices](https://stripe.com/docs/billing/invoices/tax-rates), [subscriptions](https://stripe.com/docs/billing/subscriptions/taxes) and [Checkout Sessions](https://stripe.com/docs/payments/checkout/set-up-a-subscription#tax-rates) to collect tax.
//
// Related guide: [Tax Rates](https://stripe.com/docs/billing/taxes/tax-rates).
type TaxRate struct {
	APIResource
	Active       bool              `json:"active"`
	Country      string            `json:"country"`
	Created      int64             `json:"created"`
	Description  string            `json:"description"`
	DisplayName  string            `json:"display_name"`
	ID           string            `json:"id"`
	Inclusive    bool              `json:"inclusive"`
	Jurisdiction string            `json:"jurisdiction"`
	Livemode     bool              `json:"livemode"`
	Metadata     map[string]string `json:"metadata"`
	Object       string            `json:"object"`
	Percentage   float64           `json:"percentage"`
	State        string            `json:"state"`
	TaxType      TaxRateTaxType    `json:"tax_type"`
}

// TaxRateList is a list of TaxRates as retrieved from a list endpoint.
type TaxRateList struct {
	APIResource
	ListMeta
	Data []*TaxRate `json:"data"`
}

// UnmarshalJSON handles deserialization of a TaxRate.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (t *TaxRate) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		t.ID = id
		return nil
	}

	type taxRate TaxRate
	var v taxRate
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*t = TaxRate(v)
	return nil
}
