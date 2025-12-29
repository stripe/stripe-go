//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v84/form"
)

// v1PaymentIntentAmountDetailsLineItemService is used to invoke /v1/payment_intents/{intent}/amount_details_line_items APIs.
type v1PaymentIntentAmountDetailsLineItemService struct {
	B   Backend
	Key string
}

// Lists all LineItems of a given PaymentIntent.
func (c v1PaymentIntentAmountDetailsLineItemService) List(ctx context.Context, listParams *PaymentIntentAmountDetailsLineItemListParams) Seq2[*PaymentIntentAmountDetailsLineItem, error] {
	if listParams == nil {
		listParams = &PaymentIntentAmountDetailsLineItemListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/payment_intents/%s/amount_details_line_items", StringValue(
			listParams.Intent))
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*PaymentIntentAmountDetailsLineItem], error) {
		list := &v1Page[*PaymentIntentAmountDetailsLineItem]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
