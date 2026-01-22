//
//
// File generated from our OpenAPI spec
//
//

// Package frmealvouchersonboarding provides the /v1/fr_meal_vouchers_onboardings APIs
package frmealvouchersonboarding

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/form"
)

// Client is used to invoke /v1/fr_meal_vouchers_onboardings APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a French Meal Vouchers Onboarding object that represents a restaurant's onboarding status and starts the onboarding process.
func New(params *stripe.FRMealVouchersOnboardingParams) (*stripe.FRMealVouchersOnboarding, error) {
	return getC().New(params)
}

// Creates a French Meal Vouchers Onboarding object that represents a restaurant's onboarding status and starts the onboarding process.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.FRMealVouchersOnboardingParams) (*stripe.FRMealVouchersOnboarding, error) {
	frmealvouchersonboarding := &stripe.FRMealVouchersOnboarding{}
	err := c.B.Call(
		http.MethodPost, "/v1/fr_meal_vouchers_onboardings", c.Key, params, frmealvouchersonboarding)
	return frmealvouchersonboarding, err
}

// Retrieves the details of a French Meal Vouchers Onboarding object
func Get(id string, params *stripe.FRMealVouchersOnboardingParams) (*stripe.FRMealVouchersOnboarding, error) {
	return getC().Get(id, params)
}

// Retrieves the details of a French Meal Vouchers Onboarding object
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.FRMealVouchersOnboardingParams) (*stripe.FRMealVouchersOnboarding, error) {
	path := stripe.FormatURLPath("/v1/fr_meal_vouchers_onboardings/%s", id)
	frmealvouchersonboarding := &stripe.FRMealVouchersOnboarding{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, frmealvouchersonboarding)
	return frmealvouchersonboarding, err
}

// Updates the details of a restaurant's French Meal Vouchers Onboarding object
func Update(id string, params *stripe.FRMealVouchersOnboardingParams) (*stripe.FRMealVouchersOnboarding, error) {
	return getC().Update(id, params)
}

// Updates the details of a restaurant's French Meal Vouchers Onboarding object
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.FRMealVouchersOnboardingParams) (*stripe.FRMealVouchersOnboarding, error) {
	path := stripe.FormatURLPath("/v1/fr_meal_vouchers_onboardings/%s", id)
	frmealvouchersonboarding := &stripe.FRMealVouchersOnboarding{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, frmealvouchersonboarding)
	return frmealvouchersonboarding, err
}

// Lists French Meal Vouchers Onboarding objects
func List(params *stripe.FRMealVouchersOnboardingListParams) *Iter {
	return getC().List(params)
}

// Lists French Meal Vouchers Onboarding objects
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.FRMealVouchersOnboardingListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.FRMealVouchersOnboardingList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/fr_meal_vouchers_onboardings", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for fr meal vouchers onboardings.
type Iter struct {
	*stripe.Iter
}

// FRMealVouchersOnboarding returns the fr meal vouchers onboarding which the iterator is currently pointing to.
func (i *Iter) FRMealVouchersOnboarding() *stripe.FRMealVouchersOnboarding {
	return i.Current().(*stripe.FRMealVouchersOnboarding)
}

// FRMealVouchersOnboardingList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) FRMealVouchersOnboardingList() *stripe.FRMealVouchersOnboardingList {
	return i.List().(*stripe.FRMealVouchersOnboardingList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
