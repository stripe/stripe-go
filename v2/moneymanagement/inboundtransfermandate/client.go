//
//
// File generated from our OpenAPI spec
//
//

// Package inboundtransfermandate provides the inboundtransfermandate related APIs
package inboundtransfermandate

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke inboundtransfermandate related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create an InboundTransferMandate for a v2 credential. If a pending or
// active mandate already exists for the same user and credential, that
// mandate is returned instead of creating a new one.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2MoneyManagementInboundTransferMandateParams) (*stripe.V2MoneyManagementInboundTransferMandate, error) {
	inboundtransfermandate := &stripe.V2MoneyManagementInboundTransferMandate{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/inbound_transfer_mandates", c.Key, params, inboundtransfermandate)
	return inboundtransfermandate, err
}

// Retrieve an InboundTransferMandate by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementInboundTransferMandateParams) (*stripe.V2MoneyManagementInboundTransferMandate, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/inbound_transfer_mandates/%s", id)
	inboundtransfermandate := &stripe.V2MoneyManagementInboundTransferMandate{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, inboundtransfermandate)
	return inboundtransfermandate, err
}

// Cancel a pending or active InboundTransferMandate.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.V2MoneyManagementInboundTransferMandateCancelParams) (*stripe.V2MoneyManagementInboundTransferMandate, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/inbound_transfer_mandates/%s/cancel", id)
	inboundtransfermandate := &stripe.V2MoneyManagementInboundTransferMandate{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, inboundtransfermandate)
	return inboundtransfermandate, err
}

// Retrieve a list of InboundTransferMandates for the authenticated compartment.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementInboundTransferMandateListParams) stripe.Seq2[*stripe.V2MoneyManagementInboundTransferMandate, error] {
	if listParams == nil {
		listParams = &stripe.V2MoneyManagementInboundTransferMandateListParams{}
	}
	return stripe.NewV2List("/v2/money_management/inbound_transfer_mandates", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementInboundTransferMandate], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementInboundTransferMandate]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
