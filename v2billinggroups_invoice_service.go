//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v2BillingGroupsInvoiceService is used to invoke invoice related APIs.
type v2BillingGroupsInvoiceService struct {
	B   Backend
	Key string
}

// Create an empty invoice using the billing group's settings.
func (c v2BillingGroupsInvoiceService) Create(ctx context.Context, params *V2BillingGroupsInvoiceCreateParams) (*V2BillingGroupInvoice, error) {
	if params == nil {
		params = &V2BillingGroupsInvoiceCreateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/groups/%s/invoices", StringValue(params.ID))
	groupinvoice := &V2BillingGroupInvoice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, groupinvoice)
	return groupinvoice, err
}
