//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The status of this institution in the Financial Connections authentication flow.
type FinancialConnectionsInstitutionStatus string

// List of values that FinancialConnectionsInstitutionStatus can take
const (
	FinancialConnectionsInstitutionStatusActive   FinancialConnectionsInstitutionStatus = "active"
	FinancialConnectionsInstitutionStatusDegraded FinancialConnectionsInstitutionStatus = "degraded"
	FinancialConnectionsInstitutionStatusInactive FinancialConnectionsInstitutionStatus = "inactive"
)

// Returns a list of Financial Connections Institution objects.
type FinancialConnectionsInstitutionListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *FinancialConnectionsInstitutionListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the details of a Financial Connections Institution.
type FinancialConnectionsInstitutionParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *FinancialConnectionsInstitutionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the details of a Financial Connections Institution.
type FinancialConnectionsInstitutionRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *FinancialConnectionsInstitutionRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type FinancialConnectionsInstitutionFeaturesBalances struct {
	// Whether the given feature is supported by this institution.
	Supported bool `json:"supported"`
}
type FinancialConnectionsInstitutionFeaturesOwnership struct {
	// Whether the given feature is supported by this institution.
	Supported bool `json:"supported"`
}
type FinancialConnectionsInstitutionFeaturesPaymentMethod struct {
	// Whether the given feature is supported by this institution.
	Supported bool `json:"supported"`
}
type FinancialConnectionsInstitutionFeaturesTransactions struct {
	// Whether the given feature is supported by this institution.
	Supported bool `json:"supported"`
}
type FinancialConnectionsInstitutionFeatures struct {
	Balances      *FinancialConnectionsInstitutionFeaturesBalances      `json:"balances"`
	Ownership     *FinancialConnectionsInstitutionFeaturesOwnership     `json:"ownership"`
	PaymentMethod *FinancialConnectionsInstitutionFeaturesPaymentMethod `json:"payment_method"`
	Transactions  *FinancialConnectionsInstitutionFeaturesTransactions  `json:"transactions"`
}

// An institution represents a financial institution to which an end user can connect using the Financial Connections authentication flow.
type FinancialConnectionsInstitution struct {
	APIResource
	// The list of countries supported by this institution, formatted as ISO country codes.
	Countries []string                                 `json:"countries"`
	Features  *FinancialConnectionsInstitutionFeatures `json:"features"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The name of this institution.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// A list of routing numbers which are known to correspond to this institution. Due to the many to many relationship between institutions and routing numbers, this list may not be comprehensive and routing numbers may also be shared between institutions.
	RoutingNumbers []string `json:"routing_numbers"`
	// The status of this institution in the Financial Connections authentication flow.
	Status FinancialConnectionsInstitutionStatus `json:"status"`
	// A URL corresponding to this institution. This URL is also displayed in the authentication flow to help end users confirm that they are authenticating with the right institution.
	URL string `json:"url"`
}

// FinancialConnectionsInstitutionList is a list of Institutions as retrieved from a list endpoint.
type FinancialConnectionsInstitutionList struct {
	APIResource
	ListMeta
	Data []*FinancialConnectionsInstitution `json:"data"`
}

// UnmarshalJSON handles deserialization of a FinancialConnectionsInstitution.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (f *FinancialConnectionsInstitution) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		f.ID = id
		return nil
	}

	type financialConnectionsInstitution FinancialConnectionsInstitution
	var v financialConnectionsInstitution
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*f = FinancialConnectionsInstitution(v)
	return nil
}
