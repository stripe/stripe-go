//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v86/form"
)

// v1BillingAlertNotificationService is used to invoke /v1/billing/alerts/{id}/notifications APIs.
type v1BillingAlertNotificationService struct {
	B   Backend
	Key string
}

// Lists sent billing alert triggered and recovered notifications for a billing alert.
func (c v1BillingAlertNotificationService) List(ctx context.Context, listParams *BillingAlertNotificationListParams) *V1List[*BillingAlertNotification] {
	if listParams == nil {
		listParams = &BillingAlertNotificationListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/billing/alerts/%s/notifications", StringValue(listParams.ID))
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*BillingAlertNotification], error) {
		list := &v1Page[*BillingAlertNotification]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
