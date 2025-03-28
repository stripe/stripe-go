//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Delete a specified external account for a given account.
type ExternalAccountParams struct {
	Params `form:"*"`
	// The name of the person or business that owns the bank account.
	AccountHolderName *string `form:"account_holder_name"`
	// The type of entity that holds the account. This can be either `individual` or `company`.
	AccountHolderType *string `form:"account_holder_type"`
	// The bank account type. This can only be `checking` or `savings` in most countries. In Japan, this can only be `futsu` or `toza`.
	AccountType *string `form:"account_type"`
	// City/District/Suburb/Town/Village.
	AddressCity *string `form:"address_city"`
	// Billing address country, if provided when creating card.
	AddressCountry *string `form:"address_country"`
	// Address line 1 (Street address/PO Box/Company name).
	AddressLine1 *string `form:"address_line1"`
	// Address line 2 (Apartment/Suite/Unit/Building).
	AddressLine2 *string `form:"address_line2"`
	// State/County/Province/Region.
	AddressState *string `form:"address_state"`
	// ZIP or postal code.
	AddressZip *string `form:"address_zip"`
	// When set to true, or if this is the first external account added in this currency, this account becomes the default external account for its currency.
	DefaultForCurrency *bool `form:"default_for_currency"`
	// Documents that may be submitted to satisfy various informational requests.
	Documents *ExternalAccountDocumentsParams `form:"documents"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Two digit number representing the card's expiration month.
	ExpMonth *string `form:"exp_month"`
	// Four digit number representing the card's expiration year.
	ExpYear *string `form:"exp_year"`
	// Either a token, like the ones returned by [Stripe.js](https://stripe.com/docs/js), or a dictionary containing a user's external account details (with the options shown below).
	ExternalAccount *string `form:"external_account"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Cardholder name.
	Name *string `form:"name"`
}

// AddExpand appends a new field to expand.
func (p *ExternalAccountParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *ExternalAccountParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// One or more documents that support the [Bank account ownership verification](https://support.stripe.com/questions/bank-account-ownership-verification) requirement. Must be a document associated with the bank account that displays the last 4 digits of the account number, either a statement or a check.
type ExternalAccountDocumentsBankAccountOwnershipVerificationParams struct {
	// One or more document ids returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `account_requirement`.
	Files []*string `form:"files"`
}

// Documents that may be submitted to satisfy various informational requests.
type ExternalAccountDocumentsParams struct {
	// One or more documents that support the [Bank account ownership verification](https://support.stripe.com/questions/bank-account-ownership-verification) requirement. Must be a document associated with the bank account that displays the last 4 digits of the account number, either a statement or a check.
	BankAccountOwnershipVerification *ExternalAccountDocumentsBankAccountOwnershipVerificationParams `form:"bank_account_ownership_verification"`
}

// List external accounts for an account.
type ExternalAccountListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Filter external accounts according to a particular object type.
	Object *string `form:"object"`
}

// AddExpand appends a new field to expand.
func (p *ExternalAccountListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type ExternalAccount struct{ APIResource }

// ExternalAccountList is a list of ExternalAccounts as retrieved from a list endpoint.
type ExternalAccountList struct {
	APIResource
	ListMeta
	Data []*ExternalAccount `json:"data"`
}

// UnmarshalJSON handles deserialization of an ExternalAccount.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (e *ExternalAccount) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		e.ID = id
		return nil
	}

	type externalAccount ExternalAccount
	var v externalAccount
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*e = ExternalAccount(v)
	return nil
}
