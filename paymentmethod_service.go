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

// v1PaymentMethodService is used to invoke /v1/payment_methods APIs.
type v1PaymentMethodService struct {
	B   Backend
	Key string
}

// Creates a PaymentMethod object. Read the [Stripe.js reference](https://docs.stripe.com/docs/stripe-js/reference#stripe-create-payment-method) to learn how to create PaymentMethods via Stripe.js.
//
// Instead of creating a PaymentMethod directly, we recommend using the [PaymentIntents API to accept a payment immediately or the <a href="/docs/payments/save-and-reuse">SetupIntent](https://docs.stripe.com/docs/payments/accept-a-payment) API to collect payment method details ahead of a future payment.
func (c v1PaymentMethodService) Create(ctx context.Context, params *PaymentMethodCreateParams) (*PaymentMethod, error) {
	if params == nil {
		params = &PaymentMethodCreateParams{}
	}
	params.Context = ctx
	paymentmethod := &PaymentMethod{}
	err := c.B.Call(
		http.MethodPost, "/v1/payment_methods", c.Key, params, paymentmethod)
	return paymentmethod, err
}

// Retrieves a PaymentMethod object attached to the StripeAccount. To retrieve a payment method attached to a Customer, you should use [Retrieve a Customer's PaymentMethods](https://docs.stripe.com/docs/api/payment_methods/customer)
func (c v1PaymentMethodService) Retrieve(ctx context.Context, id string, params *PaymentMethodRetrieveParams) (*PaymentMethod, error) {
	if params == nil {
		params = &PaymentMethodRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_methods/%s", id)
	paymentmethod := &PaymentMethod{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentmethod)
	return paymentmethod, err
}

// Updates a PaymentMethod object. A PaymentMethod must be attached to a customer to be updated.
func (c v1PaymentMethodService) Update(ctx context.Context, id string, params *PaymentMethodUpdateParams) (*PaymentMethod, error) {
	if params == nil {
		params = &PaymentMethodUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_methods/%s", id)
	paymentmethod := &PaymentMethod{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentmethod)
	return paymentmethod, err
}

// Attaches a PaymentMethod object to a Customer.
//
// To attach a new PaymentMethod to a customer for future payments, we recommend you use a [SetupIntent](https://docs.stripe.com/docs/api/setup_intents)
// or a PaymentIntent with [setup_future_usage](https://docs.stripe.com/docs/api/payment_intents/create#create_payment_intent-setup_future_usage).
// These approaches will perform any necessary steps to set up the PaymentMethod for future payments. Using the /v1/payment_methods/:id/attach
// endpoint without first using a SetupIntent or PaymentIntent with setup_future_usage does not optimize the PaymentMethod for
// future use, which makes later declines and payment friction more likely.
// See [Optimizing cards for future payments](https://docs.stripe.com/docs/payments/payment-intents#future-usage) for more information about setting up
// future payments.
//
// To use this PaymentMethod as the default for invoice or subscription payments,
// set [invoice_settings.default_payment_method](https://docs.stripe.com/docs/api/customers/update#update_customer-invoice_settings-default_payment_method),
// on the Customer to the PaymentMethod's ID.
func (c v1PaymentMethodService) Attach(ctx context.Context, id string, params *PaymentMethodAttachParams) (*PaymentMethod, error) {
	if params == nil {
		params = &PaymentMethodAttachParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_methods/%s/attach", id)
	paymentmethod := &PaymentMethod{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentmethod)
	return paymentmethod, err
}

// Detaches a PaymentMethod object from a Customer. After a PaymentMethod is detached, it can no longer be used for a payment or re-attached to a Customer.
func (c v1PaymentMethodService) Detach(ctx context.Context, id string, params *PaymentMethodDetachParams) (*PaymentMethod, error) {
	if params == nil {
		params = &PaymentMethodDetachParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_methods/%s/detach", id)
	paymentmethod := &PaymentMethod{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentmethod)
	return paymentmethod, err
}

// Returns a list of PaymentMethods for Treasury flows. If you want to list the PaymentMethods attached to a Customer for payments, you should use the [List a Customer's PaymentMethods](https://docs.stripe.com/docs/api/payment_methods/customer_list) API instead.
func (c v1PaymentMethodService) List(ctx context.Context, listParams *PaymentMethodListParams) Seq2[*PaymentMethod, error] {
	if listParams == nil {
		listParams = &PaymentMethodListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*PaymentMethod, ListContainer, error) {
		list := &PaymentMethodList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/payment_methods", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
