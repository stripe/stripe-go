package stripe

import (
	"encoding/json"
	"fmt"

	"github.com/stripe/stripe-go/form"
)

// BankAccountStatus is the list of allowed values for the bank account's status.
// Allowed values are "new", "verified", "validated", "errored".
type BankAccountStatus string

// BankAccountParams is the set of parameters that can be used when creating or updating a bank account.
type BankAccountParams struct {
	Params `form:"*"`

	// AccountID is the identifier of the parent account under which a bank
	// account may be nested.
	AccountID string

	// Customer is the identifier of the parent customer under which a bank
	// account may be nested.
	Customer string

	// A token referencing an external account like one returned from
	// Stripe.js.
	Token string

	// Information on an external account to reference. Only used if `Token`
	// is not provided.
	Account           string
	AccountHolderName string
	AccountHolderType string
	Default           bool
	Country           string
	Currency          string
	Routing           string
}

func (b *BankAccountParams) AppendTo(body *form.Values, keyParts []string) {
	var sourceType string
	if len(params.Customer) > 0 {
		sourceType = "source"
	} else {
		sourceType = "external_account"
	}

	// Use token (if exists) or a dictionary containing a user’s bank account details.
	if len(params.Token) > 0 {
		body.Add(append(keyParts, sourceType), params.Token)

		if params.Default {
			body.Add(append(keyParts, "default_for_currency"),
				strconv.FormatBool(params.Default))
		}
	} else {
		body.Add(append(keyParts, sourceType, "object"), "bank_account")
		body.Add(append(keyParts, sourceType, "country"), params.Country)
		body.Add(append(keyParts, sourceType, "account_number"), params.Account)
		body.Add(append(keyParts, sourceType, "currency"), params.Currency)

		if len(params.Customer) > 0 {
			body.Add(append(keyParts, sourceType, "account_holder_name"),
				params.AccountHolderName)
			body.Add(append(keyParts, sourceType, "account_holder_type"),
				params.AccountHolderType)
		}

		if len(params.Routing) > 0 {
			body.Add(append(keyParts, sourceType, "routing_number"), params.Routing)
		}

		if params.Default {
			body.Add(append(keyParts, sourceType, "default_for_currency"),
				strconv.FormatBool(params.Default))
		}
	}
}

// BankAccountListParams is the set of parameters that can be used when listing bank accounts.
type BankAccountListParams struct {
	ListParams `form:"*"`

	// The identifier of the parent account under which the bank accounts are
	// nested. Either AccountID or Customer should be populated.
	AccountID string

	// The identifier of the parent customer under which the bank accounts are
	// nested. Either AccountID or Customer should be populated.
	Customer string
}

// BankAccount represents a Stripe bank account.
type BankAccount struct {
	ID                string            `json:"id"`
	Name              string            `json:"bank_name"`
	AccountHolderName string            `json:"account_holder_name"`
	AccountHolderType string            `json:"account_holder_type"`
	Country           string            `json:"country"`
	Currency          Currency          `json:"currency"`
	Default           bool              `json:"default_for_currency"`
	LastFour          string            `json:"last4"`
	Fingerprint       string            `json:"fingerprint"`
	Status            BankAccountStatus `json:"status"`
	Routing           string            `json:"routing_number"`
	Deleted           bool              `json:"deleted"`
	Customer          *Customer         `json:"customer"`
	Meta              map[string]string `json:"metadata"`
}

// BankAccountList is a list object for bank accounts.
type BankAccountList struct {
	ListMeta
	Values []*BankAccount `json:"data"`
}

// Display implements Displayer.Display.
func (b *BankAccount) Display() string {
	return fmt.Sprintf("Bank account ending in %s", b.LastFour)
}

// AppendDetails adds the bank account's details to the query string values.
func (b *BankAccountParams) AppendDetails(values *form.Values) {
	values.Add("bank_account[country]", b.Country)
	if len(b.Routing) > 0 {
		values.Add("bank_account[routing_number]", b.Routing)
	}
	values.Add("bank_account[account_number]", b.Account)
	if b.AccountHolderName != "" {
		values.Add("bank_account[account_holder_name]", b.AccountHolderName)
	}
	if b.AccountHolderType != "" {
		values.Add("bank_account[account_holder_type]", b.AccountHolderType)
	}

	if len(b.Currency) > 0 {
		values.Add("bank_account[currency]", b.Currency)
	}
}

// UnmarshalJSON handles deserialization of a BankAccount.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (b *BankAccount) UnmarshalJSON(data []byte) error {
	type bankAccount BankAccount
	var bb bankAccount
	err := json.Unmarshal(data, &bb)
	if err == nil {
		*b = BankAccount(bb)
	} else {
		// the id is surrounded by "\" characters, so strip them
		b.ID = string(data[1 : len(data)-1])
	}

	return nil
}
