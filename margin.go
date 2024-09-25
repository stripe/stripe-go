//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// A (partner) margin represents a specific discount distributed in partner reseller programs to business partners who
// resell products and services and earn a discount (margin) for doing so.
type Margin struct {
	// Whether the margin can be applied to invoices, invoice items, or invoice line items. Defaults to `true`.
	Active bool `json:"active"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// Name of the margin that's displayed on, for example, invoices.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Percent that will be taken off the subtotal before tax (after all other discounts and promotions) of any invoice to which the margin is applied.
	PercentOff float64 `json:"percent_off"`
	// Time at which the object was last updated. Measured in seconds since the Unix epoch.
	Updated int64 `json:"updated"`
}

// UnmarshalJSON handles deserialization of a Margin.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (m *Margin) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		m.ID = id
		return nil
	}

	type margin Margin
	var v margin
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*m = Margin(v)
	return nil
}
