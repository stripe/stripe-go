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

// v2MoneyManagementCurrencyConversionService is used to invoke currencyconversion related APIs.
type v2MoneyManagementCurrencyConversionService struct {
	B   Backend
	Key string
}

// Create a new Currency Conversion in a Financial Account or the payments balance.
func (c v2MoneyManagementCurrencyConversionService) Create(ctx context.Context, params *V2MoneyManagementCurrencyConversionCreateParams) (*V2MoneyManagementCurrencyConversion, error) {
	if params == nil {
		params = &V2MoneyManagementCurrencyConversionCreateParams{}
	}
	params.Context = ctx
	currencyconversion := &V2MoneyManagementCurrencyConversion{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/currency_conversions", c.Key, params, currencyconversion)
	return currencyconversion, err
}

// Retrieves the details of a Currency Conversion by its ID.
func (c v2MoneyManagementCurrencyConversionService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementCurrencyConversionRetrieveParams) (*V2MoneyManagementCurrencyConversion, error) {
	if params == nil {
		params = &V2MoneyManagementCurrencyConversionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/currency_conversions/%s", id)
	currencyconversion := &V2MoneyManagementCurrencyConversion{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, currencyconversion)
	return currencyconversion, err
}

// Lists all Currency Conversions that match the provided filters.
func (c v2MoneyManagementCurrencyConversionService) List(ctx context.Context, listParams *V2MoneyManagementCurrencyConversionListParams) *V2List[*V2MoneyManagementCurrencyConversion] {
	if listParams == nil {
		listParams = &V2MoneyManagementCurrencyConversionListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/money_management/currency_conversions", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2MoneyManagementCurrencyConversion], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2MoneyManagementCurrencyConversion]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
