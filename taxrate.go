package stripe

import "encoding/json"

// TaxRateParams is the set of parameters that can be used when creating a tax rate.
// For more details see https://stripe.com/docs/api/tax_rates/create.
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
}

// TaxRateListParams is the set of parameters that can be used when listing tax rates.
// For more detail see https://stripe.com/docs/api/tax_rates/list.
type TaxRateListParams struct {
	ListParams   `form:"*"`
	Active       *bool             `form:"active"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Inclusive    *bool             `form:"inclusive"`
}

// TaxRate is the resource representing a Stripe tax rate.
// For more details see https://stripe.com/docs/api/tax_rates/object.
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
}

// TaxRateList is a list of tax rates as retrieved from a list endpoint.
type TaxRateList struct {
	APIResource
	ListMeta
	Data []*TaxRate `json:"data"`
}

// UnmarshalJSON handles deserialization of a TaxRate.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *TaxRate) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type taxrate TaxRate
	var v taxrate
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = TaxRate(v)
	return nil
}
