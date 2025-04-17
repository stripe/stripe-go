//
//
// File generated from our OpenAPI spec
//
//

// Package paymentintentamountdetailslineitem provides the /v1/payment_intents/{intent}/amount_details_line_items APIs
package paymentintentamountdetailslineitem

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/payment_intents/{intent}/amount_details_line_items APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Lists all LineItems of a given PaymentIntent.
func List(params *stripe.PaymentIntentAmountDetailsLineItemListParams) *Iter {
	return getC().List(params)
}

// Lists all LineItems of a given PaymentIntent.
func (c Client) List(listParams *stripe.PaymentIntentAmountDetailsLineItemListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/payment_intents/%s/amount_details_line_items", stripe.StringValue(
			listParams.Intent))
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.PaymentIntentAmountDetailsLineItemList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for payment intent amount details line items.
type Iter struct {
	*stripe.Iter
}

// PaymentIntentAmountDetailsLineItem returns the payment intent amount details line item which the iterator is currently pointing to.
func (i *Iter) PaymentIntentAmountDetailsLineItem() *stripe.PaymentIntentAmountDetailsLineItem {
	return i.Current().(*stripe.PaymentIntentAmountDetailsLineItem)
}

// PaymentIntentAmountDetailsLineItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PaymentIntentAmountDetailsLineItemList() *stripe.PaymentIntentAmountDetailsLineItemList {
	return i.List().(*stripe.PaymentIntentAmountDetailsLineItemList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
