//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of tax location to be defined. Currently the only option is `performance`.
type TaxLocationType string

// List of values that TaxLocationType can take
const (
	TaxLocationTypePerformance TaxLocationType = "performance"
)

// Retrieve a list of all tax locations. Tax locations can represent the venues for services, tickets, or other product types.
//
// The response includes detailed information for each tax location, such as its address, name, description, and current operational status.
//
// You can paginate through the list by using the limit parameter to control the number of results returned in each request.
type TaxLocationListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Type of the tax location. Currently the only option is `performance`.
	Type *string `form:"type"`
}

// AddExpand appends a new field to expand.
func (p *TaxLocationListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Create a tax location to use in calculating taxes for a service, ticket, or other type of product. The resulting object contains the id, address, name, description, and current operational status of the tax location.
type TaxLocationParams struct {
	Params `form:"*"`
	// The physical address of the tax location.
	Address *AddressParams `form:"address"`
	// Details to identify the tax location by its venue, types of events held, or available services, such as "A spacious auditorium suitable for large concerts and events.".
	Description *string `form:"description"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The type of tax location. The only supported value is "performance".
	Type *string `form:"type"`
}

// AddExpand appends a new field to expand.
func (p *TaxLocationParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Create a tax location to use in calculating taxes for a service, ticket, or other type of product. The resulting object contains the id, address, name, description, and current operational status of the tax location.
type TaxLocationCreateParams struct {
	Params `form:"*"`
	// The physical address of the tax location.
	Address *AddressParams `form:"address"`
	// Details to identify the tax location by its venue, types of events held, or available services, such as "A spacious auditorium suitable for large concerts and events.".
	Description *string `form:"description"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The type of tax location. The only supported value is "performance".
	Type *string `form:"type"`
}

// AddExpand appends a new field to expand.
func (p *TaxLocationCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Fetch the details of a specific tax location using its unique identifier. Use a tax location to calculate taxes based on the location of the end product, such as a performance, instead of the customer address. For more details, check the [integration guide](https://docs.stripe.com/tax/tax-for-tickets/integration-guide).
type TaxLocationRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *TaxLocationRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Tax locations represent venues for services, tickets, or other product types.
type TaxLocation struct {
	APIResource
	Address *Address `json:"address"`
	// A descriptive text providing additional context about the tax location. This can include information about the venue, types of events held, services available, or any relevant details for better identification (e.g., "A spacious auditorium suitable for large concerts and events.").
	Description string `json:"description"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The type of tax location to be defined. Currently the only option is `performance`.
	Type TaxLocationType `json:"type"`
}

// TaxLocationList is a list of Locations as retrieved from a list endpoint.
type TaxLocationList struct {
	APIResource
	ListMeta
	Data []*TaxLocation `json:"data"`
}
