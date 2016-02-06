package stripe

// CountrySpec is the resource representing the rules required for a Stripe account.
// For more details see https://stripe.com/docs/api/#country_specs.
type CountrySpec struct {
	ID                             string                         `json:"id"`
	SupportedBankAccountCurrencies map[string][]string            `json:"supported_bank_account_currencies"`
	SupportedPaymentCurrencies     []string                       `json:"supported_payment_currencies"`
	SupportedPaymentMethods        []string                       `json:"supported_payment_methods"`
	VerificationFields             map[string]map[string][]string `json:"verification_fields"`
}

// CountrySpecListParams are the parameters allowed during CountrySpec listing.
type CountrySpecListParams struct {
	ListParams
}
