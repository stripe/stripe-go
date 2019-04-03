// Package usagerecordsummary provides the /subscription_items/{SUBSCRIPTION_ITEM_ID}/usage_record_summaries APIs
package usagerecordsummary

import (
	"net/http"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke APIs related to usage record summaries.
type Client struct {
	B   stripe.Backend
	Key string
}

// List returns an iterator that iterates all usage record summaries.
func List(params *stripe.UsageRecordSummaryListParams) *Iter {
	return getC().List(params)
}

// List returns an iterator that iterates all usage record summaries.
func (c Client) List(listParams *stripe.UsageRecordSummaryListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		path := stripe.FormatURLPath("/v1/subscription_items/%s/usage_record_summaries", stripe.StringValue(listParams.SubscriptionItem))
		list := &stripe.UsageRecordSummaryList{}
		err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for usage record summaries.
type Iter struct {
	*stripe.Iter
}

// UsageRecordSummary returns the usage record summary which the iterator is currently pointing to.
func (i *Iter) UsageRecordSummary() *stripe.UsageRecordSummary {
	return i.Current().(*stripe.UsageRecordSummary)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
