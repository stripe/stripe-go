//
//
// File generated from our OpenAPI spec
//
//

// Package customer provides the /customers APIs
package customer

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /customers APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new customer.
func New(params *stripe.CustomerParams) (*stripe.Customer, error) {
	return getC().New(params)
}

// New creates a new customer.
func (c Client) New(params *stripe.CustomerParams) (*stripe.Customer, error) {
	customer := &stripe.Customer{}
	err := c.B.Call(http.MethodPost, "/v1/customers", c.Key, params, customer)
	return customer, err
}

// Get returns the details of a customer.
func Get(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	return getC().Get(id, params)
}

// Get returns the details of a customer.
func (c Client) Get(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	path := stripe.FormatURLPath("/v1/customers/%s", id)
	customer := &stripe.Customer{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, customer)
	return customer, err
}

// Update updates a customer's properties.
func Update(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	return getC().Update(id, params)
}

// Update updates a customer's properties.
func (c Client) Update(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	path := stripe.FormatURLPath("/v1/customers/%s", id)
	customer := &stripe.Customer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, customer)
	return customer, err
}

// Del removes a customer.
func Del(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	return getC().Del(id, params)
}

// Del removes a customer.
func (c Client) Del(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	path := stripe.FormatURLPath("/v1/customers/%s", id)
	customer := &stripe.Customer{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, customer)
	return customer, err
}

// CreateFundingInstructions is the method for the `POST /v1/customers/{customer}/funding_instructions` API.
func CreateFundingInstructions(id string, params *stripe.CustomerCreateFundingInstructionsParams) (*stripe.FundingInstructions, error) {
	return getC().CreateFundingInstructions(id, params)
}

// CreateFundingInstructions is the method for the `POST /v1/customers/{customer}/funding_instructions` API.
func (c Client) CreateFundingInstructions(id string, params *stripe.CustomerCreateFundingInstructionsParams) (*stripe.FundingInstructions, error) {
	path := stripe.FormatURLPath("/v1/customers/%s/funding_instructions", id)
	fundinginstructions := &stripe.FundingInstructions{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, fundinginstructions)
	return fundinginstructions, err
}

// DeleteDiscount is the method for the `DELETE /v1/customers/{customer}/discount` API.
func DeleteDiscount(id string, params *stripe.CustomerDeleteDiscountParams) (*stripe.Customer, error) {
	return getC().DeleteDiscount(id, params)
}

// DeleteDiscount is the method for the `DELETE /v1/customers/{customer}/discount` API.
func (c Client) DeleteDiscount(id string, params *stripe.CustomerDeleteDiscountParams) (*stripe.Customer, error) {
	path := stripe.FormatURLPath("/v1/customers/%s/discount", id)
	customer := &stripe.Customer{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, customer)
	return customer, err
}

// RetrievePaymentMethod is the method for the `GET /v1/customers/{customer}/payment_methods/{payment_method}` API.
func RetrievePaymentMethod(id string, params *stripe.CustomerRetrievePaymentMethodParams) (*stripe.PaymentMethod, error) {
	return getC().RetrievePaymentMethod(id, params)
}

// RetrievePaymentMethod is the method for the `GET /v1/customers/{customer}/payment_methods/{payment_method}` API.
func (c Client) RetrievePaymentMethod(id string, params *stripe.CustomerRetrievePaymentMethodParams) (*stripe.PaymentMethod, error) {
	path := stripe.FormatURLPath(
		"/v1/customers/%s/payment_methods/%s",
		stripe.StringValue(params.Customer),
		id,
	)
	paymentmethod := &stripe.PaymentMethod{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentmethod)
	return paymentmethod, err
}

// List returns a list of customers.
func List(params *stripe.CustomerListParams) *Iter {
	return getC().List(params)
}

// List returns a list of customers.
func (c Client) List(listParams *stripe.CustomerListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CustomerList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/customers", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for customers.
type Iter struct {
	*stripe.Iter
}

// Customer returns the customer which the iterator is currently pointing to.
func (i *Iter) Customer() *stripe.Customer {
	return i.Current().(*stripe.Customer)
}

// CustomerList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CustomerList() *stripe.CustomerList {
	return i.List().(*stripe.CustomerList)
}

// ListPaymentMethods is the method for the `GET /v1/customers/{customer}/payment_methods` API.
func ListPaymentMethods(params *stripe.CustomerListPaymentMethodsParams) *PaymentMethodIter {
	return getC().ListPaymentMethods(params)
}

// ListPaymentMethods is the method for the `GET /v1/customers/{customer}/payment_methods` API.
func (c Client) ListPaymentMethods(listParams *stripe.CustomerListPaymentMethodsParams) *PaymentMethodIter {
	path := stripe.FormatURLPath(
		"/v1/customers/%s/payment_methods",
		stripe.StringValue(listParams.Customer),
	)
	return &PaymentMethodIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.PaymentMethodList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// PaymentMethodIter is an iterator for payment methods.
type PaymentMethodIter struct {
	*stripe.Iter
}

// PaymentMethod returns the payment method which the iterator is currently pointing to.
func (i *PaymentMethodIter) PaymentMethod() *stripe.PaymentMethod {
	return i.Current().(*stripe.PaymentMethod)
}

// PaymentMethodList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *PaymentMethodIter) PaymentMethodList() *stripe.PaymentMethodList {
	return i.List().(*stripe.PaymentMethodList)
}

// Search returns a search result containing customers.
func Search(params *stripe.CustomerSearchParams) *SearchIter {
	return getC().Search(params)
}

// Search returns a search result containing customers.
func (c Client) Search(params *stripe.CustomerSearchParams) *SearchIter {
	return &SearchIter{
		SearchIter: stripe.GetSearchIter(params, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.SearchContainer, error) {
			list := &stripe.CustomerSearchResult{}
			err := c.B.CallRaw(http.MethodGet, "/v1/customers/search", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// SearchIter is an iterator for customers.
type SearchIter struct {
	*stripe.SearchIter
}

// Customer returns the customer which the iterator is currently pointing to.
func (i *SearchIter) Customer() *stripe.Customer {
	return i.Current().(*stripe.Customer)
}

// CustomerSearchResult returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *SearchIter) CustomerSearchResult() *stripe.CustomerSearchResult {
	return i.SearchResult().(*stripe.CustomerSearchResult)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
