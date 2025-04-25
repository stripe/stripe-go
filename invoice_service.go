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

// v1InvoiceService is used to invoke /v1/invoices APIs.
type v1InvoiceService struct {
	B   Backend
	Key string
}

// This endpoint creates a draft invoice for a given customer. The invoice remains a draft until you [finalize the invoice, which allows you to [pay](#pay_invoice) or <a href="#send_invoice">send](https://stripe.com/docs/api#finalize_invoice) the invoice to your customers.
func (c v1InvoiceService) Create(ctx context.Context, params *InvoiceCreateParams) (*Invoice, error) {
	invoice := &Invoice{}
	if params == nil {
		params = &InvoiceCreateParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, "/v1/invoices", c.Key, params, invoice)
	return invoice, err
}

// Retrieves the invoice with the given ID.
func (c v1InvoiceService) Retrieve(ctx context.Context, id string, params *InvoiceRetrieveParams) (*Invoice, error) {
	path := FormatURLPath("/v1/invoices/%s", id)
	invoice := &Invoice{}
	if params == nil {
		params = &InvoiceRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, invoice)
	return invoice, err
}

// Draft invoices are fully editable. Once an invoice is [finalized](https://stripe.com/docs/billing/invoices/workflow#finalized),
// monetary values, as well as collection_method, become uneditable.
//
// If you would like to stop the Stripe Billing engine from automatically finalizing, reattempting payments on,
// sending reminders for, or [automatically reconciling](https://stripe.com/docs/billing/invoices/reconciliation) invoices, pass
// auto_advance=false.
func (c v1InvoiceService) Update(ctx context.Context, id string, params *InvoiceUpdateParams) (*Invoice, error) {
	path := FormatURLPath("/v1/invoices/%s", id)
	invoice := &Invoice{}
	if params == nil {
		params = &InvoiceUpdateParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// Permanently deletes a one-off invoice draft. This cannot be undone. Attempts to delete invoices that are no longer in a draft state will fail; once an invoice has been finalized or if an invoice is for a subscription, it must be [voided](https://stripe.com/docs/api#void_invoice).
func (c v1InvoiceService) Delete(ctx context.Context, id string, params *InvoiceDeleteParams) (*Invoice, error) {
	path := FormatURLPath("/v1/invoices/%s", id)
	invoice := &Invoice{}
	if params == nil {
		params = &InvoiceDeleteParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodDelete, path, c.Key, params, invoice)
	return invoice, err
}

// Adds multiple line items to an invoice. This is only possible when an invoice is still a draft.
func (c v1InvoiceService) AddLines(ctx context.Context, id string, params *InvoiceAddLinesParams) (*Invoice, error) {
	path := FormatURLPath("/v1/invoices/%s/add_lines", id)
	invoice := &Invoice{}
	if params == nil {
		params = &InvoiceAddLinesParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// At any time, you can preview the upcoming invoice for a subscription or subscription schedule. This will show you all the charges that are pending, including subscription renewal charges, invoice item charges, etc. It will also show you any discounts that are applicable to the invoice.
//
// Note that when you are viewing an upcoming invoice, you are simply viewing a preview – the invoice has not yet been created. As such, the upcoming invoice will not show up in invoice listing calls, and you cannot use the API to pay or edit the invoice. If you want to change the amount that your customer will be billed, you can add, remove, or update pending invoice items, or update the customer's discount.
//
// You can preview the effects of updating a subscription, including a preview of what proration will take place. To ensure that the actual proration is calculated exactly the same as the previewed proration, you should pass the subscription_details.proration_date parameter when doing the actual subscription update. The recommended way to get only the prorations being previewed is to consider only proration line items where period[start] is equal to the subscription_details.proration_date value passed in the request.
//
// Note: Currency conversion calculations use the latest exchange rates. Exchange rates may vary between the time of the preview and the time of the actual invoice creation. [Learn more](https://docs.stripe.com/currencies/conversions)
func (c v1InvoiceService) CreatePreview(ctx context.Context, params *InvoiceCreatePreviewParams) (*Invoice, error) {
	invoice := &Invoice{}
	if params == nil {
		params = &InvoiceCreatePreviewParams{}
	}
	params.Context = ctx
	err := c.B.Call(
		http.MethodPost, "/v1/invoices/create_preview", c.Key, params, invoice)
	return invoice, err
}

// Stripe automatically finalizes drafts before sending and attempting payment on invoices. However, if you'd like to finalize a draft invoice manually, you can do so using this method.
func (c v1InvoiceService) FinalizeInvoice(ctx context.Context, id string, params *InvoiceFinalizeInvoiceParams) (*Invoice, error) {
	path := FormatURLPath("/v1/invoices/%s/finalize", id)
	invoice := &Invoice{}
	if params == nil {
		params = &InvoiceFinalizeInvoiceParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// Marking an invoice as uncollectible is useful for keeping track of bad debts that can be written off for accounting purposes.
func (c v1InvoiceService) MarkUncollectible(ctx context.Context, id string, params *InvoiceMarkUncollectibleParams) (*Invoice, error) {
	path := FormatURLPath("/v1/invoices/%s/mark_uncollectible", id)
	invoice := &Invoice{}
	if params == nil {
		params = &InvoiceMarkUncollectibleParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// Stripe automatically creates and then attempts to collect payment on invoices for customers on subscriptions according to your [subscriptions settings](https://dashboard.stripe.com/account/billing/automatic). However, if you'd like to attempt payment on an invoice out of the normal collection schedule or for some other reason, you can do so.
func (c v1InvoiceService) Pay(ctx context.Context, id string, params *InvoicePayParams) (*Invoice, error) {
	path := FormatURLPath("/v1/invoices/%s/pay", id)
	invoice := &Invoice{}
	if params == nil {
		params = &InvoicePayParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// Removes multiple line items from an invoice. This is only possible when an invoice is still a draft.
func (c v1InvoiceService) RemoveLines(ctx context.Context, id string, params *InvoiceRemoveLinesParams) (*Invoice, error) {
	path := FormatURLPath("/v1/invoices/%s/remove_lines", id)
	invoice := &Invoice{}
	if params == nil {
		params = &InvoiceRemoveLinesParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// Stripe will automatically send invoices to customers according to your [subscriptions settings](https://dashboard.stripe.com/account/billing/automatic). However, if you'd like to manually send an invoice to your customer out of the normal schedule, you can do so. When sending invoices that have already been paid, there will be no reference to the payment in the email.
//
// Requests made in test-mode result in no emails being sent, despite sending an invoice.sent event.
func (c v1InvoiceService) SendInvoice(ctx context.Context, id string, params *InvoiceSendInvoiceParams) (*Invoice, error) {
	path := FormatURLPath("/v1/invoices/%s/send", id)
	invoice := &Invoice{}
	if params == nil {
		params = &InvoiceSendInvoiceParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// Updates multiple line items on an invoice. This is only possible when an invoice is still a draft.
func (c v1InvoiceService) UpdateLines(ctx context.Context, id string, params *InvoiceUpdateLinesParams) (*Invoice, error) {
	path := FormatURLPath("/v1/invoices/%s/update_lines", id)
	invoice := &Invoice{}
	if params == nil {
		params = &InvoiceUpdateLinesParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// Mark a finalized invoice as void. This cannot be undone. Voiding an invoice is similar to [deletion](https://stripe.com/docs/api#delete_invoice), however it only applies to finalized invoices and maintains a papertrail where the invoice can still be found.
//
// Consult with local regulations to determine whether and how an invoice might be amended, canceled, or voided in the jurisdiction you're doing business in. You might need to [issue another invoice or <a href="#create_credit_note">credit note](https://stripe.com/docs/api#create_invoice) instead. Stripe recommends that you consult with your legal counsel for advice specific to your business.
func (c v1InvoiceService) VoidInvoice(ctx context.Context, id string, params *InvoiceVoidInvoiceParams) (*Invoice, error) {
	path := FormatURLPath("/v1/invoices/%s/void", id)
	invoice := &Invoice{}
	if params == nil {
		params = &InvoiceVoidInvoiceParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// You can list all invoices, or list the invoices for a specific customer. The invoices are returned sorted by creation date, with the most recently created invoices appearing first.
func (c v1InvoiceService) List(ctx context.Context, listParams *InvoiceListParams) Seq2[*Invoice, error] {
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*Invoice, ListContainer, error) {
		list := &InvoiceList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/invoices", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}

// When retrieving an invoice, you'll get a lines property containing the total count of line items and the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
func (c v1InvoiceService) ListLines(ctx context.Context, listParams *InvoiceListLinesParams) Seq2[*InvoiceLineItem, error] {
	path := FormatURLPath(
		"/v1/invoices/%s/lines", StringValue(listParams.Invoice))
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*InvoiceLineItem, ListContainer, error) {
		list := &InvoiceLineItemList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}

// Search for invoices you've previously created using Stripe's [Search Query Language](https://stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
func (c v1InvoiceService) Search(ctx context.Context, params *InvoiceSearchParams) Seq2[*Invoice, error] {
	return newV1SearchList(params, func(p *Params, b *form.Values) ([]*Invoice, SearchContainer, error) {
		list := &InvoiceSearchResult{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/invoices/search", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
