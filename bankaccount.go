package stripe

import "net/url"

// BankAccountStatus is the list of allowed values for the bank account's status.
// Allowed values are "new", "verified", "validated", "errored".
type BankAccountStatus string

// BankAccountParams is the set of parameters that can be used when creating or updating a bank account.
type BankAccountParams struct {
	Params
	AccountID, Token, Country, Routing, Account, Currency string
	Default                                               bool
}

// BankAccountListParams is the set of parameters that can be used when listing bank accounts.
type BankAccountListParams struct {
	ListParams
	AccountID string
}

// BankAccount represents a Stripe bank account.
type BankAccount struct {
	ID          string            `json:"id"`
	Name        string            `json:"bank_name"`
	Country     string            `json:"country"`
	Currency    Currency          `json:"currency"`
	LastFour    string            `json:"last4"`
	Fingerprint string            `json:"fingerprint"`
	Status      BankAccountStatus `json:"status"`
}

// AppendDetails adds the bank account's details to the query string values.
func (b *BankAccountParams) AppendDetails(values *url.Values) {
	values.Add("bank_account[country]", b.Country)
	values.Add("bank_account[routing_number]", b.Routing)
	values.Add("bank_account[account_number]", b.Account)

	if len(b.Currency) > 0 {
		values.Add("bank_account[currency]", b.Currency)
	}
}
