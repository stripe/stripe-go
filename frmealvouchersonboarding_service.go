//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v84/form"
)

// v1FRMealVouchersOnboardingService is used to invoke /v1/fr_meal_vouchers_onboardings APIs.
type v1FRMealVouchersOnboardingService struct {
	B   Backend
	Key string
}

// Creates a French Meal Vouchers Onboarding object that represents a restaurant's onboarding status and starts the onboarding process.
func (c v1FRMealVouchersOnboardingService) Create(ctx context.Context, params *FRMealVouchersOnboardingCreateParams) (*FRMealVouchersOnboarding, error) {
	if params == nil {
		params = &FRMealVouchersOnboardingCreateParams{}
	}
	params.Context = ctx
	frmealvouchersonboarding := &FRMealVouchersOnboarding{}
	err := c.B.Call(
		http.MethodPost, "/v1/fr_meal_vouchers_onboardings", c.Key, params, frmealvouchersonboarding)
	return frmealvouchersonboarding, err
}

// Retrieves the details of a French Meal Vouchers Onboarding object
func (c v1FRMealVouchersOnboardingService) Retrieve(ctx context.Context, id string, params *FRMealVouchersOnboardingRetrieveParams) (*FRMealVouchersOnboarding, error) {
	if params == nil {
		params = &FRMealVouchersOnboardingRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/fr_meal_vouchers_onboardings/%s", id)
	frmealvouchersonboarding := &FRMealVouchersOnboarding{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, frmealvouchersonboarding)
	return frmealvouchersonboarding, err
}

// Updates the details of a restaurant's French Meal Vouchers Onboarding object
func (c v1FRMealVouchersOnboardingService) Update(ctx context.Context, id string, params *FRMealVouchersOnboardingUpdateParams) (*FRMealVouchersOnboarding, error) {
	if params == nil {
		params = &FRMealVouchersOnboardingUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/fr_meal_vouchers_onboardings/%s", id)
	frmealvouchersonboarding := &FRMealVouchersOnboarding{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, frmealvouchersonboarding)
	return frmealvouchersonboarding, err
}

// Lists French Meal Vouchers Onboarding objects
func (c v1FRMealVouchersOnboardingService) List(ctx context.Context, listParams *FRMealVouchersOnboardingListParams) Seq2[*FRMealVouchersOnboarding, error] {
	if listParams == nil {
		listParams = &FRMealVouchersOnboardingListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*FRMealVouchersOnboarding], error) {
		list := &v1Page[*FRMealVouchersOnboarding]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/fr_meal_vouchers_onboardings", c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
