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

// v1GiftCardService is used to invoke /v1/gift_cards APIs.
type v1GiftCardService struct {
	B   Backend
	Key string
}

// Creates a gift card object.
func (c v1GiftCardService) Create(ctx context.Context, params *GiftCardCreateParams) (*GiftCard, error) {
	if params == nil {
		params = &GiftCardCreateParams{}
	}
	params.Context = ctx
	giftcard := &GiftCard{}
	err := c.B.Call(http.MethodPost, "/v1/gift_cards", c.Key, params, giftcard)
	return giftcard, err
}

// Retrieves a third-party gift card object.
func (c v1GiftCardService) Retrieve(ctx context.Context, id string, params *GiftCardRetrieveParams) (*GiftCard, error) {
	if params == nil {
		params = &GiftCardRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/gift_cards/%s", id)
	giftcard := &GiftCard{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, giftcard)
	return giftcard, err
}

// Activates a third-party gift card and optionally sets its balance.
func (c v1GiftCardService) Activate(ctx context.Context, id string, params *GiftCardActivateParams) (*GiftCardOperation, error) {
	if params == nil {
		params = &GiftCardActivateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/gift_cards/%s/activate", id)
	giftcardoperation := &GiftCardOperation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, giftcardoperation)
	return giftcardoperation, err
}

// Cashout a third-party gift card by zeroing its balance.
func (c v1GiftCardService) Cashout(ctx context.Context, id string, params *GiftCardCashoutParams) (*GiftCardOperation, error) {
	if params == nil {
		params = &GiftCardCashoutParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/gift_cards/%s/cashout", id)
	giftcardoperation := &GiftCardOperation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, giftcardoperation)
	return giftcardoperation, err
}

// Checks the balance of a third-party gift card.
func (c v1GiftCardService) CheckBalance(ctx context.Context, id string, params *GiftCardCheckBalanceParams) (*GiftCardOperation, error) {
	if params == nil {
		params = &GiftCardCheckBalanceParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/gift_cards/%s/check_balance", id)
	giftcardoperation := &GiftCardOperation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, giftcardoperation)
	return giftcardoperation, err
}

// Reloads a third-party gift card by adding the specified amount to its balance.
func (c v1GiftCardService) Reload(ctx context.Context, id string, params *GiftCardReloadParams) (*GiftCardOperation, error) {
	if params == nil {
		params = &GiftCardReloadParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/gift_cards/%s/reload", id)
	giftcardoperation := &GiftCardOperation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, giftcardoperation)
	return giftcardoperation, err
}

// Voids a previously performed gift card operation.
func (c v1GiftCardService) VoidOperation(ctx context.Context, id string, params *GiftCardVoidOperationParams) (*GiftCardOperation, error) {
	if params == nil {
		params = &GiftCardVoidOperationParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/gift_cards/%s/void_operation", id)
	giftcardoperation := &GiftCardOperation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, giftcardoperation)
	return giftcardoperation, err
}
