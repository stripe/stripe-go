//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The precision level of the resolved address.
type V2TaxOperationsResolveAddressResultPrecision string

// List of values that V2TaxOperationsResolveAddressResultPrecision can take
const (
	V2TaxOperationsResolveAddressResultPrecisionNone       V2TaxOperationsResolveAddressResultPrecision = "none"
	V2TaxOperationsResolveAddressResultPrecisionAddress    V2TaxOperationsResolveAddressResultPrecision = "address"
	V2TaxOperationsResolveAddressResultPrecisionCity       V2TaxOperationsResolveAddressResultPrecision = "city"
	V2TaxOperationsResolveAddressResultPrecisionCountry    V2TaxOperationsResolveAddressResultPrecision = "country"
	V2TaxOperationsResolveAddressResultPrecisionPostalCode V2TaxOperationsResolveAddressResultPrecision = "postal_code"
	V2TaxOperationsResolveAddressResultPrecisionState      V2TaxOperationsResolveAddressResultPrecision = "state"
	V2TaxOperationsResolveAddressResultPrecisionStreet     V2TaxOperationsResolveAddressResultPrecision = "street"
)

// A code describing the issue.
type V2TaxOperationsResolveAddressResultPrecisionDetailsIssueCode string

// List of values that V2TaxOperationsResolveAddressResultPrecisionDetailsIssueCode can take
const (
	V2TaxOperationsResolveAddressResultPrecisionDetailsIssueCodeRequiredForImprovedPrecision V2TaxOperationsResolveAddressResultPrecisionDetailsIssueCode = "required_for_improved_precision"
)

// The address field with the issue.
type V2TaxOperationsResolveAddressResultPrecisionDetailsIssueField string

// List of values that V2TaxOperationsResolveAddressResultPrecisionDetailsIssueField can take
const (
	V2TaxOperationsResolveAddressResultPrecisionDetailsIssueFieldCity       V2TaxOperationsResolveAddressResultPrecisionDetailsIssueField = "city"
	V2TaxOperationsResolveAddressResultPrecisionDetailsIssueFieldCountry    V2TaxOperationsResolveAddressResultPrecisionDetailsIssueField = "country"
	V2TaxOperationsResolveAddressResultPrecisionDetailsIssueFieldLine1      V2TaxOperationsResolveAddressResultPrecisionDetailsIssueField = "line1"
	V2TaxOperationsResolveAddressResultPrecisionDetailsIssueFieldPostalCode V2TaxOperationsResolveAddressResultPrecisionDetailsIssueField = "postal_code"
	V2TaxOperationsResolveAddressResultPrecisionDetailsIssueFieldState      V2TaxOperationsResolveAddressResultPrecisionDetailsIssueField = "state"
)

// The normalized form of the input address.
type V2TaxOperationsResolveAddressResultAddress struct {
	// The city.
	City string `json:"city,omitempty"`
	// The two-letter country code.
	Country string `json:"country,omitempty"`
	// The first line of the street address.
	Line1 string `json:"line1,omitempty"`
	// The postal code.
	PostalCode string `json:"postal_code,omitempty"`
	// The state or province.
	State string `json:"state,omitempty"`
}

// Issues preventing higher precision.
type V2TaxOperationsResolveAddressResultPrecisionDetailsIssue struct {
	// A code describing the issue.
	Code V2TaxOperationsResolveAddressResultPrecisionDetailsIssueCode `json:"code"`
	// The address field with the issue.
	Field V2TaxOperationsResolveAddressResultPrecisionDetailsIssueField `json:"field"`
}

// Details about the precision, including any issues.
type V2TaxOperationsResolveAddressResultPrecisionDetails struct {
	// Issues preventing higher precision.
	Issues []*V2TaxOperationsResolveAddressResultPrecisionDetailsIssue `json:"issues"`
}

// The result of resolving an address to its tax precision level.
type V2TaxOperationsResolveAddressResult struct {
	APIResource
	// The normalized form of the input address.
	Address *V2TaxOperationsResolveAddressResultAddress `json:"address"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The precision level of the resolved address.
	Precision V2TaxOperationsResolveAddressResultPrecision `json:"precision"`
	// Details about the precision, including any issues.
	PrecisionDetails *V2TaxOperationsResolveAddressResultPrecisionDetails `json:"precision_details"`
}
