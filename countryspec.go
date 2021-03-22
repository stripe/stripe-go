//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Country is the list of supported countries
type Country string

// CountrySpecParams are the parameters allowed during CountrySpec retrieval.
type CountrySpecParams struct {
	Params `form:"*"`
}

// CountrySpecListParams are the parameters allowed during CountrySpec listing.
type CountrySpecListParams struct {
	ListParams `form:"*"`
}

// VerificationFieldsList lists the fields needed for an account verification.
// For more details see https://stripe.com/docs/api#country_spec_object-verification_fields.
type VerificationFieldsList struct {
	AdditionalFields []string `json:"additional"`
	Minimum          []string `json:"minimum"`
}

// CountrySpec is the resource representing the rules required for a Stripe account.
// For more details see https://stripe.com/docs/api/#country_specs.
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

// CountrySpecList is a list of country specs as retrieved from a list endpoint.
type CountrySpecList struct {
	APIResource
	ListMeta
	Data []*CountrySpec `json:"data"`
}
