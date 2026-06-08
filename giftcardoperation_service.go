//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v1GiftCardOperationService is used to invoke /v1/gift_card_operations APIs.
type v1GiftCardOperationService struct {
	B   Backend
	Key string
}

// Retrieves a third-party gift card operation object.
func (c v1GiftCardOperationService) Retrieve(ctx context.Context, id string, params *GiftCardOperationRetrieveParams) (*GiftCardOperation, error) {
	if params == nil {
		params = &GiftCardOperationRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/gift_card_operations/%s", id)
	giftcardoperation := &GiftCardOperation{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, giftcardoperation)
	return giftcardoperation, err
}
