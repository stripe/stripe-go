//
//
// File generated from our OpenAPI spec
//
//

package stripe

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

// An institution represents a banking institution which may be available for an end user to select in the Financial Connections authentication flow.
type FinancialConnectionsInstitution struct {
	APIResource
	Features *FinancialConnectionsInstitutionFeatures `json:"features"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The name of this institution.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// A list of routing numbers which are known to correspond to this institution.
	RoutingNumbers []string `json:"routing_numbers"`
	// The status of this institution in the Financial Connections authentication flow.
	Status FinancialConnectionsInstitutionStatus `json:"status"`
	// The URL for this institution's website.
	URL string `json:"url"`
}

// FinancialConnectionsInstitutionList is a list of Institutions as retrieved from a list endpoint.
type FinancialConnectionsInstitutionList struct {
	APIResource
	ListMeta
	Data []*FinancialConnectionsInstitution `json:"data"`
}
