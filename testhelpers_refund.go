//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Expire a refund with a status of requires_action.
type TestHelpersRefundExpireParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// GetParams is a nil-safe implementation of the ParamsContainer interface.
func (p *TestHelpersRefundExpireParams) GetParams() *Params {
	if p == nil {
		return nil
	}
	return &p.Params
}

// AddExpand appends a new field to expand.
func (p *TestHelpersRefundExpireParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}
