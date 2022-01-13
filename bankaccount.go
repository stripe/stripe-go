//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v72/form"
	"strconv"
)

// The type of entity that holds the account. This can be either `individual` or `company`.
type BankAccountAccountHolderType string

// List of values that BankAccountAccountHolderType can take
const (
	BankAccountAccountHolderTypeCompany    BankAccountAccountHolderType = "company"
	BankAccountAccountHolderTypeIndividual BankAccountAccountHolderType = "individual"
)

// A set of available payout methods for this bank account. Only values from this set should be passed as the `method` when creating a payout.
type BankAccountAvailablePayoutMethod string

// List of values that BankAccountAvailablePayoutMethod can take
const (
	BankAccountAvailablePayoutMethodInstant  BankAccountAvailablePayoutMethod = "instant"
	BankAccountAvailablePayoutMethodStandard BankAccountAvailablePayoutMethod = "standard"
)

// For bank accounts, possible values are `new`, `validated`, `verified`, `verification_failed`, or `errored`. A bank account that hasn't had any activity or validation performed is `new`. If Stripe can determine that the bank account exists, its status will be `validated`. Note that there often isn't enough information to know (e.g., for smaller credit unions), and the validation is not always run. If customer bank account verification has succeeded, the bank account status will be `verified`. If the verification failed for any reason, such as microdeposit failure, the status will be `verification_failed`. If a transfer sent to this bank account fails, we'll set the status to `errored` and will not continue to send transfers until the bank details are updated.
//
// For external accounts, possible values are `new` and `errored`. Validations aren't run against external accounts because they're only used for payouts. This means the other statuses don't apply. If a transfer fails, the status is set to `errored` and transfers are stopped until account details are updated.
type BankAccountStatus string

// List of values that BankAccountStatus can take
const (
	BankAccountStatusErrored            BankAccountStatus = "errored"
	BankAccountStatusNew                BankAccountStatus = "new"
	BankAccountStatusValidated          BankAccountStatus = "validated"
	BankAccountStatusVerificationFailed BankAccountStatus = "verification_failed"
	BankAccountStatusVerified           BankAccountStatus = "verified"
)

// Updates the metadata, account holder name, account holder type of a bank account belonging to a [Custom account](https://stripe.com/docs/connect/custom-accounts), and optionally sets it as the default for its currency. Other bank account details are not editable by design.
//
// You can re-enable a disabled bank account by performing an update call without providing any arguments or changes.
type BankAccountParams struct {
	Params   `form:"*"`
	Customer *string `form:"-"` // Included in URL
	// Token is a token referencing an external account like one returned from
	// Stripe.js.
	Token *string `form:"-"` // Included in URL
	// Account is the identifier of the parent account under which bank
	// accounts are nested.
	Account            *string `form:"-"` // Included in URL
	AccountHolderName  *string `form:"account_holder_name"`
	AccountHolderType  *string `form:"account_holder_type"`
	AccountNumber      *string `form:"account_number"`
	AccountType        *string `form:"account_type"`
	AddressCity        *string `form:"address_city"`
	AddressCountry     *string `form:"address_country"`
	AddressLine1       *string `form:"address_line1"`
	AddressLine2       *string `form:"address_line2"`
	AddressState       *string `form:"address_state"`
	AddressZip         *string `form:"address_zip"`
	Country            *string `form:"country"`
	Currency           *string `form:"currency"`
	DefaultForCurrency *bool   `form:"default_for_currency"`
	ExpMonth           *string `form:"exp_month"`
	ExpYear            *string `form:"exp_year"`
	Name               *string `form:"name"`
	RoutingNumber      *string `form:"routing_number"`
	// ID is used when tokenizing a bank account for shared customers
	ID *string `form:"*"`
}

// AppendToAsSourceOrExternalAccount appends the given BankAccountParams as
// either a source or external account.
//
// It may look like an AppendTo from the form package, but it's not, and is
// only used in the special case where we use `bankaccount.New`. It's needed
// because we have some weird encoding logic here that can't be handled by the
// form package (and it's special enough that it wouldn't be desirable to have
// it do so).
//
// This is not a pattern that we want to push forward, and this largely exists
// because the bank accounts endpoint is a little unusual. There is one other
// resource like it, which is cards.
func (a *BankAccountParams) AppendToAsSourceOrExternalAccount(body *form.Values) {
	// Rather than being called in addition to `AppendTo`, this function
	// *replaces* `AppendTo`, so we must also make sure to handle the encoding
	// of `Params` so metadata and the like is included in the encoded payload.
	form.AppendTo(body, a.Params)

	isCustomer := a.Customer != nil

	var sourceType string
	if isCustomer {
		sourceType = "source"
	} else {
		sourceType = "external_account"
	}

	// Use token (if exists) or a dictionary containing a user’s bank account details.
	if a.Token != nil {
		body.Add(sourceType, StringValue(a.Token))

		if a.DefaultForCurrency != nil {
			body.Add(
				"default_for_currency",
				strconv.FormatBool(BoolValue(a.DefaultForCurrency)),
			)
		}
	} else {
		body.Add(sourceType+"[object]", "bank_account")
		body.Add(sourceType+"[country]", StringValue(a.Country))
		body.Add(sourceType+"[account_number]", StringValue(a.AccountNumber))
		body.Add(sourceType+"[currency]", StringValue(a.Currency))

		// These are optional and the API will fail if we try to send empty
		// values in for them, so make sure to check that they're actually set
		// before encoding them.
		if a.AccountHolderName != nil {
			body.Add(sourceType+"[account_holder_name]", StringValue(a.AccountHolderName))
		}

		if a.AccountHolderType != nil {
			body.Add(sourceType+"[account_holder_type]", StringValue(a.AccountHolderType))
		}

		if a.RoutingNumber != nil {
			body.Add(sourceType+"[routing_number]", StringValue(a.RoutingNumber))
		}

		if a.DefaultForCurrency != nil {
			body.Add(sourceType+"[default_for_currency]", strconv.FormatBool(BoolValue(a.DefaultForCurrency)))
		}
	}
}

type BankAccountListParams struct {
	ListParams `form:"*"`
	// The identifier of the parent account under which the bank accounts are
	// nested. Either Account or Customer should be populated.
	Account *string `form:"-"` // Included in URL
	// The identifier of the parent customer under which the bank accounts are
	// nested. Either Account or Customer should be populated.
	Customer *string `form:"-"` // Included in URL
}

// AppendTo implements custom encoding logic for BankAccountListParams
// so that we can send the special required `object` field up along with the
// other specified parameters.
func (p *BankAccountListParams) AppendTo(body *form.Values, keyParts []string) {
	body.Add(form.FormatKey(append(keyParts, "object")), "bank_account")
}

// These bank accounts are payment methods on `Customer` objects.
//
// On the other hand [External Accounts](https://stripe.com/docs/api#external_accounts) are transfer
// destinations on `Account` objects for [Custom accounts](https://stripe.com/docs/connect/custom-accounts).
// They can be bank accounts or debit cards as well, and are documented in the links above.
//
// Related guide: [Bank Debits and Transfers](https://stripe.com/docs/payments/bank-debits-transfers).
type BankAccount struct {
	APIResource
	Account                *Account                           `json:"account"`
	AccountHolderName      string                             `json:"account_holder_name"`
	AccountHolderType      BankAccountAccountHolderType       `json:"account_holder_type"`
	AccountType            string                             `json:"account_type"`
	AvailablePayoutMethods []BankAccountAvailablePayoutMethod `json:"available_payout_methods"`
	BankName               string                             `json:"bank_name"`
	Country                string                             `json:"country"`
	Currency               Currency                           `json:"currency"`
	Customer               *Customer                          `json:"customer"`
	DefaultForCurrency     bool                               `json:"default_for_currency"`
	Deleted                bool                               `json:"deleted"`
	Fingerprint            string                             `json:"fingerprint"`
	ID                     string                             `json:"id"`
	Last4                  string                             `json:"last4"`
	Metadata               map[string]string                  `json:"metadata"`
	Object                 string                             `json:"object"`
	RoutingNumber          string                             `json:"routing_number"`
	Status                 BankAccountStatus                  `json:"status"`
}

// BankAccountList is a list of BankAccounts as retrieved from a list endpoint.
type BankAccountList struct {
	APIResource
	ListMeta
	Data []*BankAccount `json:"data"`
}

// UnmarshalJSON handles deserialization of a BankAccount.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (b *BankAccount) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		b.ID = id
		return nil
	}

	type bankAccount BankAccount
	var v bankAccount
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*b = BankAccount(v)
	return nil
}
