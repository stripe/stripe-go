//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"encoding/json"
	"net/http"
)

// v1CustomerSessionService is used to invoke /v1/customer_sessions APIs.
type v1CustomerSessionService struct {
	B   Backend
	Key string
}

// Creates a Customer Session object that includes a single-use client secret that you can use on your front-end to grant client-side API access for certain customer resources.
func (c v1CustomerSessionService) Create(ctx context.Context, params *CustomerSessionCreateParams) (*CustomerSession, error) {
	if params == nil {
		params = &CustomerSessionCreateParams{}
	}
	params.Context = ctx
	customersession := &CustomerSession{}
	err := c.B.Call(
		http.MethodPost, "/v1/customer_sessions", c.Key, params, customersession)
	return customersession, err
}

// Serializes a CustomerSession create request into a batch job JSONL line.
func (c v1CustomerSessionService) MarshalBatchCreate(params *CustomerSessionCreateParams) (string, error) {
	itemID, err := newUUID4()
	if err != nil {
		return "", err
	}

	item := struct {
		ID            string            `json:"id"`
		Context       string            `json:"context,omitempty"`
		StripeVersion string            `json:"stripe_version,omitempty"`
		PathParams    map[string]string `json:"path_params,omitempty"`
		Params        interface{}       `json:"params"`
	}{
		ID:            itemID,
		PathParams:    nil,
		StripeVersion: APIVersion,
	}
	if params != nil {
		item.Params = params
		if params.StripeContext != nil {
			item.Context = *params.StripeContext
		}
	}
	b, err := json.Marshal(item)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
