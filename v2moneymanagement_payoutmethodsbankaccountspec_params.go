//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Fetch the specifications for a set of countries to know which
// credential fields are required, the validations for each fields, and how to translate these
// country-specific fields to the generic fields in the PayoutMethodBankAccount type.
type V2MoneyManagementPayoutMethodsBankAccountSpecParams struct {
	Params `form:"*"`
	// The countries to fetch the bank account spec for.
	Countries []*string `form:"countries" json:"countries,omitempty"`
}
