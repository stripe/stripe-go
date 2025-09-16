//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1BillingMeterEventSummaryService is used to invoke /v1/billing/meters/{id}/event_summaries APIs.
type v1BillingMeterEventSummaryService struct {
	B   Backend
	Key string
}

// Retrieve a list of billing meter event summaries.
func (c v1BillingMeterEventSummaryService) List(ctx context.Context, listParams *BillingMeterEventSummaryListParams) Seq2[*BillingMeterEventSummary, error] {
	if listParams == nil {
		listParams = &BillingMeterEventSummaryListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/billing/meters/%s/event_summaries", StringValue(listParams.ID))
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*BillingMeterEventSummary, ListContainer, error) {
		list := &BillingMeterEventSummaryList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
