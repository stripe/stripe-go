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

// v1PaymentMethodDomainService is used to invoke /v1/payment_method_domains APIs.
type v1PaymentMethodDomainService struct {
	B   Backend
	Key string
}

// Creates a payment method domain.
func (c v1PaymentMethodDomainService) Create(ctx context.Context, params *PaymentMethodDomainCreateParams) (*PaymentMethodDomain, error) {
	if params == nil {
		params = &PaymentMethodDomainCreateParams{}
	}
	params.Context = ctx
	paymentmethoddomain := &PaymentMethodDomain{}
	err := c.B.Call(
		http.MethodPost, "/v1/payment_method_domains", c.Key, params, paymentmethoddomain)
	return paymentmethoddomain, err
}

// Retrieves the details of an existing payment method domain.
func (c v1PaymentMethodDomainService) Retrieve(ctx context.Context, id string, params *PaymentMethodDomainRetrieveParams) (*PaymentMethodDomain, error) {
	if params == nil {
		params = &PaymentMethodDomainRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_method_domains/%s", id)
	paymentmethoddomain := &PaymentMethodDomain{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentmethoddomain)
	return paymentmethoddomain, err
}

// Updates an existing payment method domain.
func (c v1PaymentMethodDomainService) Update(ctx context.Context, id string, params *PaymentMethodDomainUpdateParams) (*PaymentMethodDomain, error) {
	if params == nil {
		params = &PaymentMethodDomainUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_method_domains/%s", id)
	paymentmethoddomain := &PaymentMethodDomain{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentmethoddomain)
	return paymentmethoddomain, err
}

// Some payment methods might require additional steps to register a domain. If the requirements weren't satisfied when the domain was created, the payment method will be inactive on the domain.
// The payment method doesn't appear in Elements or Embedded Checkout for this domain until it is active.
//
// To activate a payment method on an existing payment method domain, complete the required registration steps specific to the payment method, and then validate the payment method domain with this endpoint.
//
// Related guides: [Payment method domains](https://docs.stripe.com/docs/payments/payment-methods/pmd-registration).
func (c v1PaymentMethodDomainService) Validate(ctx context.Context, id string, params *PaymentMethodDomainValidateParams) (*PaymentMethodDomain, error) {
	if params == nil {
		params = &PaymentMethodDomainValidateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_method_domains/%s/validate", id)
	paymentmethoddomain := &PaymentMethodDomain{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentmethoddomain)
	return paymentmethoddomain, err
}

// Lists the details of existing payment method domains.
func (c v1PaymentMethodDomainService) List(ctx context.Context, listParams *PaymentMethodDomainListParams) Seq2[*PaymentMethodDomain, error] {
	if listParams == nil {
		listParams = &PaymentMethodDomainListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*PaymentMethodDomain, ListContainer, error) {
		list := &PaymentMethodDomainList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/payment_method_domains", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
