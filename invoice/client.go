//
//
// File generated from our OpenAPI spec
//
//

// Package invoice provides the /v1/invoices APIs
package invoice

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/invoices APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// This endpoint creates a draft invoice for a given customer. The invoice remains a draft until you [finalize the invoice, which allows you to [pay](#pay_invoice) or <a href="#send_invoice">send](https://docs.stripe.com/api#finalize_invoice) the invoice to your customers.
func New(params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	return getC().New(params)
}

// This endpoint creates a draft invoice for a given customer. The invoice remains a draft until you [finalize the invoice, which allows you to [pay](#pay_invoice) or <a href="#send_invoice">send](https://docs.stripe.com/api#finalize_invoice) the invoice to your customers.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodPost, "/v1/invoices", c.Key, params, invoice)
	return invoice, err
}

// Retrieves the invoice with the given ID.
func Get(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	return getC().Get(id, params)
}

// Retrieves the invoice with the given ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/v1/invoices/%s", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, invoice)
	return invoice, err
}

// Draft invoices are fully editable. Once an invoice is [finalized](https://docs.stripe.com/docs/billing/invoices/workflow#finalized),
// monetary values, as well as collection_method, become uneditable.
//
// If you would like to stop the Stripe Billing engine from automatically finalizing, reattempting payments on,
// sending reminders for, or [automatically reconciling](https://docs.stripe.com/docs/billing/invoices/reconciliation) invoices, pass
// auto_advance=false.
func Update(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	return getC().Update(id, params)
}

// Draft invoices are fully editable. Once an invoice is [finalized](https://docs.stripe.com/docs/billing/invoices/workflow#finalized),
// monetary values, as well as collection_method, become uneditable.
//
// If you would like to stop the Stripe Billing engine from automatically finalizing, reattempting payments on,
// sending reminders for, or [automatically reconciling](https://docs.stripe.com/docs/billing/invoices/reconciliation) invoices, pass
// auto_advance=false.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/v1/invoices/%s", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// Permanently deletes a one-off invoice draft. This cannot be undone. Attempts to delete invoices that are no longer in a draft state will fail; once an invoice has been finalized or if an invoice is for a subscription, it must be [voided](https://docs.stripe.com/api#void_invoice).
func Del(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	return getC().Del(id, params)
}

// Permanently deletes a one-off invoice draft. This cannot be undone. Attempts to delete invoices that are no longer in a draft state will fail; once an invoice has been finalized or if an invoice is for a subscription, it must be [voided](https://docs.stripe.com/api#void_invoice).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Del(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/v1/invoices/%s", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, invoice)
	return invoice, err
}

// Adds multiple line items to an invoice. This is only possible when an invoice is still a draft.
func AddLines(id string, params *stripe.InvoiceAddLinesParams) (*stripe.Invoice, error) {
	return getC().AddLines(id, params)
}

// Adds multiple line items to an invoice. This is only possible when an invoice is still a draft.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) AddLines(id string, params *stripe.InvoiceAddLinesParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/v1/invoices/%s/add_lines", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// Attaches a PaymentIntent or an Out of Band Payment to the invoice, adding it to the list of payments.
//
// For the PaymentIntent, when the PaymentIntent's status changes to succeeded, the payment is credited
// to the invoice, increasing its amount_paid. When the invoice is fully paid, the
// invoice's status becomes paid.
//
// If the PaymentIntent's status is already succeeded when it's attached, it's
// credited to the invoice immediately.
//
// See: [Partial payments](https://docs.stripe.com/docs/invoicing/partial-payments) to learn more.
func AttachPayment(id string, params *stripe.InvoiceAttachPaymentParams) (*stripe.Invoice, error) {
	return getC().AttachPayment(id, params)
}

// Attaches a PaymentIntent or an Out of Band Payment to the invoice, adding it to the list of payments.
//
// For the PaymentIntent, when the PaymentIntent's status changes to succeeded, the payment is credited
// to the invoice, increasing its amount_paid. When the invoice is fully paid, the
// invoice's status becomes paid.
//
// If the PaymentIntent's status is already succeeded when it's attached, it's
// credited to the invoice immediately.
//
// See: [Partial payments](https://docs.stripe.com/docs/invoicing/partial-payments) to learn more.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) AttachPayment(id string, params *stripe.InvoiceAttachPaymentParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/v1/invoices/%s/attach_payment", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// At any time, you can preview the upcoming invoice for a subscription or subscription schedule. This will show you all the charges that are pending, including subscription renewal charges, invoice item charges, etc. It will also show you any discounts that are applicable to the invoice.
//
// You can also preview the effects of creating or updating a subscription or subscription schedule, including a preview of any prorations that will take place. To ensure that the actual proration is calculated exactly the same as the previewed proration, you should pass the subscription_details.proration_date parameter when doing the actual subscription update.
//
// The recommended way to get only the prorations being previewed on the invoice is to consider line items where parent.subscription_item_details.proration is true.
//
// Note that when you are viewing an upcoming invoice, you are simply viewing a preview – the invoice has not yet been created. As such, the upcoming invoice will not show up in invoice listing calls, and you cannot use the API to pay or edit the invoice. If you want to change the amount that your customer will be billed, you can add, remove, or update pending invoice items, or update the customer's discount.
//
// Note: Currency conversion calculations use the latest exchange rates. Exchange rates may vary between the time of the preview and the time of the actual invoice creation. [Learn more](https://docs.stripe.com/currencies/conversions)
func CreatePreview(params *stripe.InvoiceCreatePreviewParams) (*stripe.Invoice, error) {
	return getC().CreatePreview(params)
}

// At any time, you can preview the upcoming invoice for a subscription or subscription schedule. This will show you all the charges that are pending, including subscription renewal charges, invoice item charges, etc. It will also show you any discounts that are applicable to the invoice.
//
// You can also preview the effects of creating or updating a subscription or subscription schedule, including a preview of any prorations that will take place. To ensure that the actual proration is calculated exactly the same as the previewed proration, you should pass the subscription_details.proration_date parameter when doing the actual subscription update.
//
// The recommended way to get only the prorations being previewed on the invoice is to consider line items where parent.subscription_item_details.proration is true.
//
// Note that when you are viewing an upcoming invoice, you are simply viewing a preview – the invoice has not yet been created. As such, the upcoming invoice will not show up in invoice listing calls, and you cannot use the API to pay or edit the invoice. If you want to change the amount that your customer will be billed, you can add, remove, or update pending invoice items, or update the customer's discount.
//
// Note: Currency conversion calculations use the latest exchange rates. Exchange rates may vary between the time of the preview and the time of the actual invoice creation. [Learn more](https://docs.stripe.com/currencies/conversions)
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) CreatePreview(params *stripe.InvoiceCreatePreviewParams) (*stripe.Invoice, error) {
	invoice := &stripe.Invoice{}
	err := c.B.Call(
		http.MethodPost, "/v1/invoices/create_preview", c.Key, params, invoice)
	return invoice, err
}

// Stripe automatically finalizes drafts before sending and attempting payment on invoices. However, if you'd like to finalize a draft invoice manually, you can do so using this method.
func FinalizeInvoice(id string, params *stripe.InvoiceFinalizeInvoiceParams) (*stripe.Invoice, error) {
	return getC().FinalizeInvoice(id, params)
}

// Stripe automatically finalizes drafts before sending and attempting payment on invoices. However, if you'd like to finalize a draft invoice manually, you can do so using this method.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) FinalizeInvoice(id string, params *stripe.InvoiceFinalizeInvoiceParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/v1/invoices/%s/finalize", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// Marking an invoice as uncollectible is useful for keeping track of bad debts that can be written off for accounting purposes.
func MarkUncollectible(id string, params *stripe.InvoiceMarkUncollectibleParams) (*stripe.Invoice, error) {
	return getC().MarkUncollectible(id, params)
}

// Marking an invoice as uncollectible is useful for keeping track of bad debts that can be written off for accounting purposes.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) MarkUncollectible(id string, params *stripe.InvoiceMarkUncollectibleParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/v1/invoices/%s/mark_uncollectible", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// Stripe automatically creates and then attempts to collect payment on invoices for customers on subscriptions according to your [subscriptions settings](https://dashboard.stripe.com/account/billing/automatic). However, if you'd like to attempt payment on an invoice out of the normal collection schedule or for some other reason, you can do so.
func Pay(id string, params *stripe.InvoicePayParams) (*stripe.Invoice, error) {
	return getC().Pay(id, params)
}

// Stripe automatically creates and then attempts to collect payment on invoices for customers on subscriptions according to your [subscriptions settings](https://dashboard.stripe.com/account/billing/automatic). However, if you'd like to attempt payment on an invoice out of the normal collection schedule or for some other reason, you can do so.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Pay(id string, params *stripe.InvoicePayParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/v1/invoices/%s/pay", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// Removes multiple line items from an invoice. This is only possible when an invoice is still a draft.
func RemoveLines(id string, params *stripe.InvoiceRemoveLinesParams) (*stripe.Invoice, error) {
	return getC().RemoveLines(id, params)
}

// Removes multiple line items from an invoice. This is only possible when an invoice is still a draft.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) RemoveLines(id string, params *stripe.InvoiceRemoveLinesParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/v1/invoices/%s/remove_lines", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// Stripe will automatically send invoices to customers according to your [subscriptions settings](https://dashboard.stripe.com/account/billing/automatic). However, if you'd like to manually send an invoice to your customer out of the normal schedule, you can do so. When sending invoices that have already been paid, there will be no reference to the payment in the email.
//
// Requests made in test-mode result in no emails being sent, despite sending an invoice.sent event.
func SendInvoice(id string, params *stripe.InvoiceSendInvoiceParams) (*stripe.Invoice, error) {
	return getC().SendInvoice(id, params)
}

// Stripe will automatically send invoices to customers according to your [subscriptions settings](https://dashboard.stripe.com/account/billing/automatic). However, if you'd like to manually send an invoice to your customer out of the normal schedule, you can do so. When sending invoices that have already been paid, there will be no reference to the payment in the email.
//
// Requests made in test-mode result in no emails being sent, despite sending an invoice.sent event.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) SendInvoice(id string, params *stripe.InvoiceSendInvoiceParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/v1/invoices/%s/send", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// Updates multiple line items on an invoice. This is only possible when an invoice is still a draft.
func UpdateLines(id string, params *stripe.InvoiceUpdateLinesParams) (*stripe.Invoice, error) {
	return getC().UpdateLines(id, params)
}

// Updates multiple line items on an invoice. This is only possible when an invoice is still a draft.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) UpdateLines(id string, params *stripe.InvoiceUpdateLinesParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/v1/invoices/%s/update_lines", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// Mark a finalized invoice as void. This cannot be undone. Voiding an invoice is similar to [deletion](https://docs.stripe.com/api#delete_invoice), however it only applies to finalized invoices and maintains a papertrail where the invoice can still be found.
//
// Consult with local regulations to determine whether and how an invoice might be amended, canceled, or voided in the jurisdiction you're doing business in. You might need to [issue another invoice or <a href="#create_credit_note">credit note](https://docs.stripe.com/api#create_invoice) instead. Stripe recommends that you consult with your legal counsel for advice specific to your business.
func VoidInvoice(id string, params *stripe.InvoiceVoidInvoiceParams) (*stripe.Invoice, error) {
	return getC().VoidInvoice(id, params)
}

// Mark a finalized invoice as void. This cannot be undone. Voiding an invoice is similar to [deletion](https://docs.stripe.com/api#delete_invoice), however it only applies to finalized invoices and maintains a papertrail where the invoice can still be found.
//
// Consult with local regulations to determine whether and how an invoice might be amended, canceled, or voided in the jurisdiction you're doing business in. You might need to [issue another invoice or <a href="#create_credit_note">credit note](https://docs.stripe.com/api#create_invoice) instead. Stripe recommends that you consult with your legal counsel for advice specific to your business.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) VoidInvoice(id string, params *stripe.InvoiceVoidInvoiceParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/v1/invoices/%s/void", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// You can list all invoices, or list the invoices for a specific customer. The invoices are returned sorted by creation date, with the most recently created invoices appearing first.
func List(params *stripe.InvoiceListParams) *Iter {
	return getC().List(params)
}

// You can list all invoices, or list the invoices for a specific customer. The invoices are returned sorted by creation date, with the most recently created invoices appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.InvoiceListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.InvoiceList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/invoices", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for invoices.
type Iter struct {
	*stripe.Iter
}

// Invoice returns the invoice which the iterator is currently pointing to.
func (i *Iter) Invoice() *stripe.Invoice {
	return i.Current().(*stripe.Invoice)
}

// InvoiceList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) InvoiceList() *stripe.InvoiceList {
	return i.List().(*stripe.InvoiceList)
}

// When retrieving an invoice, you'll get a lines property containing the total count of line items and the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
func ListLines(params *stripe.InvoiceListLinesParams) *LineItemIter {
	return getC().ListLines(params)
}

// When retrieving an invoice, you'll get a lines property containing the total count of line items and the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListLines(listParams *stripe.InvoiceListLinesParams) *LineItemIter {
	path := stripe.FormatURLPath(
		"/v1/invoices/%s/lines", stripe.StringValue(listParams.Invoice))
	return &LineItemIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.InvoiceLineItemList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// LineItemIter is an iterator for invoice line items.
type LineItemIter struct {
	*stripe.Iter
}

// InvoiceLineItem returns the invoice line item which the iterator is currently pointing to.
func (i *LineItemIter) InvoiceLineItem() *stripe.InvoiceLineItem {
	return i.Current().(*stripe.InvoiceLineItem)
}

// InvoiceLineItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *LineItemIter) InvoiceLineItemList() *stripe.InvoiceLineItemList {
	return i.List().(*stripe.InvoiceLineItemList)
}

// Search for invoices you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
func Search(params *stripe.InvoiceSearchParams) *SearchIter {
	return getC().Search(params)
}

// Search for invoices you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Search(params *stripe.InvoiceSearchParams) *SearchIter {
	return &SearchIter{
		SearchIter: stripe.GetSearchIter(params, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.SearchContainer, error) {
			list := &stripe.InvoiceSearchResult{}
			err := c.B.CallRaw(http.MethodGet, "/v1/invoices/search", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// SearchIter is an iterator for invoices.
type SearchIter struct {
	*stripe.SearchIter
}

// Invoice returns the invoice which the iterator is currently pointing to.
func (i *SearchIter) Invoice() *stripe.Invoice {
	return i.Current().(*stripe.Invoice)
}

// InvoiceSearchResult returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *SearchIter) InvoiceSearchResult() *stripe.InvoiceSearchResult {
	return i.SearchResult().(*stripe.InvoiceSearchResult)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
