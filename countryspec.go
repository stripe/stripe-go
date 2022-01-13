//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Lists all Country Spec objects available in the API.
type CountrySpecListParams struct {
	ListParams `form:"*"`
}

// Country is the list of supported countries
type Country string

// Returns a Country Spec for a given Country code.
type CountrySpecParams struct {
	Params `form:"*"`
}

// VerificationFieldsList lists the fields needed for an account verification.
// For more details see https://stripe.com/docs/api#country_spec_object-verification_fields.
type VerificationFieldsList struct {
	AdditionalFields []string `json:"additional"`
	Minimum          []string `json:"minimum"`
}

// Stripe needs to collect certain pieces of information about each account
// created. These requirements can differ depending on the account's country. The
// Country Specs API makes these rules available to your integration.
//
// You can also view the information from this API call as [an online
// guide](https://stripe.com/docs/connect/required-verification-information).
type CountrySpec struct {
	APIResource
	DefaultCurrency                Currency                                        `json:"default_currency"`
	ID                             string                                          `json:"id"`
	Object                         string                                          `json:"object"`
	SupportedBankAccountCurrencies map[Currency][]Country                          `json:"supported_bank_account_currencies"`
	SupportedPaymentCurrencies     []Currency                                      `json:"supported_payment_currencies"`
	SupportedPaymentMethods        []string                                        `json:"supported_payment_methods"`
	SupportedTransferCountries     []string                                        `json:"supported_transfer_countries"`
	VerificationFields             map[AccountBusinessType]*VerificationFieldsList `json:"verification_fields"`
}

// CountrySpecList is a list of CountrySpecs as retrieved from a list endpoint.
type CountrySpecList struct {
	APIResource
	ListMeta
	Data []*CountrySpec `json:"data"`
}
