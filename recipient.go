package stripe

import (
	"encoding/json"

	"github.com/stripe/stripe-go/form"
)

// RecipientType is the list of allowed values for the recipient's type.
// Allowed values are "individual", "corporation".
type RecipientType string

// RecipientParams is the set of parameters that can be used when creating or updating recipients.
// For more details see https://stripe.com/docs/api#create_recipient and https://stripe.com/docs/api#update_recipient.
type RecipientParams struct {
	Params `form:"*"`

	Bank        *BankAccountParams `form:"-"` // Kind of an abberation because a bank account's token will be replace the rest of its data. Keep this in a custom AppendTo for now.
	Card        *CardParams        `form:"card"`
	DefaultCard string             `form:"default_card"`
	Desc        string             `form:"description"`
	Email       string             `form:"email"`
	Name        string             `form:"name"`
	TaxID       string             `form:"tax_id"`
	Token       string             `form:"card"`
	Type        RecipientType      `form:"-"` // Doesn't seem to be used anywhere
}

// AppendTo implements some custom behavior around a recipient's bank account.
// This was probably the wrong way to go about this, but grandfather the
// behavior for the time being.
func (p *RecipientParams) AppendTo(body *form.Values, keyParts []string) {
	if p.Bank != nil {
		if len(p.Bank.Token) > 0 {
			body.Add("bank_account", p.Bank.Token)
		} else {
			form.AppendToPrefixed(body, p.Bank, append(keyParts, "bank_account"))
		}
	}
}

// RecipientListParams is the set of parameters that can be used when listing recipients.
// For more details see https://stripe.com/docs/api#list_recipients.
type RecipientListParams struct {
	ListParams `form:"*"`
	Verified   bool `form:"verified"`
}

// Recipient is the resource representing a Stripe recipient.
// For more details see https://stripe.com/docs/api#recipients.
type Recipient struct {
	Bank        *BankAccount      `json:"active_account"`
	Cards       *CardList         `json:"cards"`
	Created     int64             `json:"created"`
	DefaultCard *Card             `json:"default_card"`
	Deleted     bool              `json:"deleted"`
	Desc        string            `json:"description"`
	Email       string            `json:"email"`
	ID          string            `json:"id"`
	Live        bool              `json:"livemode"`
	Meta        map[string]string `json:"metadata"`
	MigratedTo  *Account          `json:"migrated_to"`
	Name        string            `json:"name"`
	Type        RecipientType     `json:"type"`
}

// RecipientList is a list of recipients as retrieved from a list endpoint.
type RecipientList struct {
	ListMeta
	Values []*Recipient `json:"data"`
}

// UnmarshalJSON handles deserialization of a Recipient.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (r *Recipient) UnmarshalJSON(data []byte) error {
	type recipient Recipient
	var rr recipient
	err := json.Unmarshal(data, &rr)
	if err == nil {
		*r = Recipient(rr)
	} else {
		// the id is surrounded by "\" characters, so strip them
		r.ID = string(data[1 : len(data)-1])
	}

	return nil
}
