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

// v1CustomerService is used to invoke /v1/customers APIs.
type v1CustomerService struct {
	B   Backend
	Key string
}

// Creates a new customer object.
func (c v1CustomerService) Create(ctx context.Context, params *CustomerCreateParams) (*Customer, error) {
	if params == nil {
		params = &CustomerCreateParams{}
	}
	params.Context = ctx
	customer := &Customer{}
	err := c.B.Call(http.MethodPost, "/v1/customers", c.Key, params, customer)
	return customer, err
}

// Retrieves a Customer object.
func (c v1CustomerService) Retrieve(ctx context.Context, id string, params *CustomerRetrieveParams) (*Customer, error) {
	if params == nil {
		params = &CustomerRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/customers/%s", id)
	customer := &Customer{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, customer)
	return customer, err
}

// Updates the specified customer by setting the values of the parameters passed. Any parameters not provided will be left unchanged. For example, if you pass the source parameter, that becomes the customer's active source (e.g., a card) to be used for all charges in the future. When you update a customer to a new valid card source by passing the source parameter: for each of the customer's current subscriptions, if the subscription bills automatically and is in the past_due state, then the latest open invoice for the subscription with automatic collection enabled will be retried. This retry will not count as an automatic retry, and will not affect the next regularly scheduled payment for the invoice. Changing the default_source for a customer will not trigger this behavior.
//
// This request accepts mostly the same arguments as the customer creation call.
func (c v1CustomerService) Update(ctx context.Context, id string, params *CustomerUpdateParams) (*Customer, error) {
	if params == nil {
		params = &CustomerUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/customers/%s", id)
	customer := &Customer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, customer)
	return customer, err
}

// Permanently deletes a customer. It cannot be undone. Also immediately cancels any active subscriptions on the customer.
func (c v1CustomerService) Delete(ctx context.Context, id string, params *CustomerDeleteParams) (*Customer, error) {
	if params == nil {
		params = &CustomerDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/customers/%s", id)
	customer := &Customer{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, customer)
	return customer, err
}

// Retrieve funding instructions for a customer cash balance. If funding instructions do not yet exist for the customer, new
// funding instructions will be created. If funding instructions have already been created for a given customer, the same
// funding instructions will be retrieved. In other words, we will return the same funding instructions each time.
func (c v1CustomerService) CreateFundingInstructions(ctx context.Context, id string, params *CustomerCreateFundingInstructionsParams) (*FundingInstructions, error) {
	if params == nil {
		params = &CustomerCreateFundingInstructionsParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/customers/%s/funding_instructions", id)
	fundinginstructions := &FundingInstructions{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, fundinginstructions)
	return fundinginstructions, err
}

// Removes the currently applied discount on a customer.
func (c v1CustomerService) DeleteDiscount(ctx context.Context, id string, params *CustomerDeleteDiscountParams) (*Customer, error) {
	if params == nil {
		params = &CustomerDeleteDiscountParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/customers/%s/discount", id)
	customer := &Customer{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, customer)
	return customer, err
}

// Retrieves a PaymentMethod object for a given Customer.
func (c v1CustomerService) RetrievePaymentMethod(ctx context.Context, id string, params *CustomerRetrievePaymentMethodParams) (*PaymentMethod, error) {
	if params == nil {
		params = &CustomerRetrievePaymentMethodParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/payment_methods/%s", StringValue(params.Customer), id)
	paymentmethod := &PaymentMethod{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentmethod)
	return paymentmethod, err
}

// Returns a list of your customers. The customers are returned sorted by creation date, with the most recent customers appearing first.
func (c v1CustomerService) List(ctx context.Context, listParams *CustomerListParams) Seq2[*Customer, error] {
	if listParams == nil {
		listParams = &CustomerListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*Customer, ListContainer, error) {
		list := &CustomerList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/customers", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}

// Returns a list of PaymentMethods for a given Customer
func (c v1CustomerService) ListPaymentMethods(ctx context.Context, listParams *CustomerListPaymentMethodsParams) Seq2[*PaymentMethod, error] {
	if listParams == nil {
		listParams = &CustomerListPaymentMethodsParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/payment_methods", StringValue(listParams.Customer))
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*PaymentMethod, ListContainer, error) {
		list := &PaymentMethodList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}

// Search for customers you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
func (c v1CustomerService) Search(ctx context.Context, params *CustomerSearchParams) Seq2[*Customer, error] {
	if params == nil {
		params = &CustomerSearchParams{}
	}
	params.Context = ctx
	return newV1SearchList(params, func(p *Params, b *form.Values) ([]*Customer, SearchContainer, error) {
		list := &CustomerSearchResult{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/customers/search", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
