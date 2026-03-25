//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Issuers are available to this restaurant via Conecs, will be blank if the onboarding to Conecs is not complete or unsuccessful
type FRMealVouchersOnboardingProvidersConecsIssuersAvailable string

// List of values that FRMealVouchersOnboardingProvidersConecsIssuersAvailable can take
const (
	FRMealVouchersOnboardingProvidersConecsIssuersAvailableBimpli  FRMealVouchersOnboardingProvidersConecsIssuersAvailable = "bimpli"
	FRMealVouchersOnboardingProvidersConecsIssuersAvailableEdenred FRMealVouchersOnboardingProvidersConecsIssuersAvailable = "edenred"
	FRMealVouchersOnboardingProvidersConecsIssuersAvailablePluxee  FRMealVouchersOnboardingProvidersConecsIssuersAvailable = "pluxee"
	FRMealVouchersOnboardingProvidersConecsIssuersAvailableUp      FRMealVouchersOnboardingProvidersConecsIssuersAvailable = "up"
)

// The code for the type of error.
type FRMealVouchersOnboardingProvidersConecsRequirementsErrorCode string

// List of values that FRMealVouchersOnboardingProvidersConecsRequirementsErrorCode can take
const (
	FRMealVouchersOnboardingProvidersConecsRequirementsErrorCodePostalCodeInvalid FRMealVouchersOnboardingProvidersConecsRequirementsErrorCode = "postal_code_invalid"
	FRMealVouchersOnboardingProvidersConecsRequirementsErrorCodeSiretInvalid      FRMealVouchersOnboardingProvidersConecsRequirementsErrorCode = "siret_invalid"
)

// The specific onboarding requirement field (in the requirements hash) that needs to be resolved.
type FRMealVouchersOnboardingProvidersConecsRequirementsErrorRequirement string

// List of values that FRMealVouchersOnboardingProvidersConecsRequirementsErrorRequirement can take
const (
	FRMealVouchersOnboardingProvidersConecsRequirementsErrorRequirementPostalCode FRMealVouchersOnboardingProvidersConecsRequirementsErrorRequirement = "postal_code"
	FRMealVouchersOnboardingProvidersConecsRequirementsErrorRequirementSiret      FRMealVouchersOnboardingProvidersConecsRequirementsErrorRequirement = "siret"
)

// Fields that need to be provided to complete the onboarding to Conecs.
type FRMealVouchersOnboardingProvidersConecsRequirementsPastDue string

// List of values that FRMealVouchersOnboardingProvidersConecsRequirementsPastDue can take
const (
	FRMealVouchersOnboardingProvidersConecsRequirementsPastDuePostalCode FRMealVouchersOnboardingProvidersConecsRequirementsPastDue = "postal_code"
	FRMealVouchersOnboardingProvidersConecsRequirementsPastDueSiret      FRMealVouchersOnboardingProvidersConecsRequirementsPastDue = "siret"
)

// Status of the restaurant's onboarding to Conecs
type FRMealVouchersOnboardingProvidersConecsStatus string

// List of values that FRMealVouchersOnboardingProvidersConecsStatus can take
const (
	FRMealVouchersOnboardingProvidersConecsStatusActionRequired FRMealVouchersOnboardingProvidersConecsStatus = "action_required"
	FRMealVouchersOnboardingProvidersConecsStatusActive         FRMealVouchersOnboardingProvidersConecsStatus = "active"
	FRMealVouchersOnboardingProvidersConecsStatusDisentitled    FRMealVouchersOnboardingProvidersConecsStatus = "disentitled"
	FRMealVouchersOnboardingProvidersConecsStatusPending        FRMealVouchersOnboardingProvidersConecsStatus = "pending"
)

// Lists French Meal Vouchers Onboarding objects. The objects are returned in sorted order, with the most recently created objects appearing first.
type FRMealVouchersOnboardingListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *FRMealVouchersOnboardingListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Creates a French Meal Vouchers Onboarding object that represents a restaurant's onboarding status and starts the onboarding process.
type FRMealVouchersOnboardingParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`. This cannot be changed after creation of this object.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Name of the restaurant. This cannot be changed after creation of this object.
	Name *string `form:"name" json:"name,omitempty"`
	// Corrected Postal code of the restaurant.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// SIRET number associated with the restaurant. This cannot be changed after creation of this object.
	Siret *string `form:"siret" json:"siret,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *FRMealVouchersOnboardingParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *FRMealVouchersOnboardingParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Creates a French Meal Vouchers Onboarding object that represents a restaurant's onboarding status and starts the onboarding process.
type FRMealVouchersOnboardingCreateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`. This cannot be changed after creation of this object.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Name of the restaurant. This cannot be changed after creation of this object.
	Name *string `form:"name" json:"name"`
	// Postal code of the restaurant.
	PostalCode *string `form:"postal_code" json:"postal_code"`
	// SIRET number associated with the restaurant. This cannot be changed after creation of this object.
	Siret *string `form:"siret" json:"siret"`
}

// AddExpand appends a new field to expand.
func (p *FRMealVouchersOnboardingCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *FRMealVouchersOnboardingCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves the details of a previously created French Meal Vouchers Onboarding object.
//
// Supply the unique French Meal Vouchers Onboarding ID that was returned from your previous request,
// and Stripe returns the corresponding onboarding information.
type FRMealVouchersOnboardingRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *FRMealVouchersOnboardingRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Updates the details of a restaurant's French Meal Vouchers Onboarding object by
// setting the values of the parameters passed. Any parameters not provided are left unchanged.
// After you update the object, the onboarding process automatically restarts.
//
// You can only update French Meal Vouchers Onboarding objects with the postal_code field requirement in past_due.
type FRMealVouchersOnboardingUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Corrected Postal code of the restaurant.
	PostalCode *string `form:"postal_code" json:"postal_code"`
}

// AddExpand appends a new field to expand.
func (p *FRMealVouchersOnboardingUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// This represents information which issuers are available to this restaurant via Conecs
type FRMealVouchersOnboardingProvidersConecsIssuers struct {
	// Issuers are available to this restaurant via Conecs, will be blank if the onboarding to Conecs is not complete or unsuccessful
	Available []FRMealVouchersOnboardingProvidersConecsIssuersAvailable `json:"available"`
}

// Information any errors that are preventing the onboarding to Conecs from being completed.
type FRMealVouchersOnboardingProvidersConecsRequirementsError struct {
	// The code for the type of error.
	Code FRMealVouchersOnboardingProvidersConecsRequirementsErrorCode `json:"code"`
	// An informative message that provides additional details about the error.
	Message string `json:"message"`
	// The specific onboarding requirement field (in the requirements hash) that needs to be resolved.
	Requirement FRMealVouchersOnboardingProvidersConecsRequirementsErrorRequirement `json:"requirement"`
}

// This represents information about outstanding requirements for this restaurant to onboard to Conecs
type FRMealVouchersOnboardingProvidersConecsRequirements struct {
	// Information any errors that are preventing the onboarding to Conecs from being completed.
	Errors []*FRMealVouchersOnboardingProvidersConecsRequirementsError `json:"errors"`
	// Fields that need to be provided to complete the onboarding to Conecs.
	PastDue []FRMealVouchersOnboardingProvidersConecsRequirementsPastDue `json:"past_due"`
}

// This represents the onboarding state of the restaurant on Conecs.
type FRMealVouchersOnboardingProvidersConecs struct {
	// This represents information which issuers are available to this restaurant via Conecs
	Issuers *FRMealVouchersOnboardingProvidersConecsIssuers `json:"issuers"`
	// This represents information about outstanding requirements for this restaurant to onboard to Conecs
	Requirements *FRMealVouchersOnboardingProvidersConecsRequirements `json:"requirements"`
	// Status of the restaurant's onboarding to Conecs
	Status FRMealVouchersOnboardingProvidersConecsStatus `json:"status"`
}

// This represents the onboarding state of the restaurant on different providers.
type FRMealVouchersOnboardingProviders struct {
	// This represents the onboarding state of the restaurant on Conecs.
	Conecs *FRMealVouchersOnboardingProvidersConecs `json:"conecs"`
}

// The `French Meal Vouchers Onboarding` resource encapsulates the onboarding status and other related information
// for a single restaurant (SIRET number) in the context of the French Meal Vouchers program.
//
// To onboard a restaurant for the French Meal Vouchers program, you create a `French Meal Vouchers Onboarding` object.
// You can retrieve individual objects, list all such objects, or update objects to correct the postal code of the restaurant.
// We identify `French Meal Vouchers Onboarding` objects with a unique, random ID.
//
// Related guide: [Set up a restaurant for titres-restaurant payments](https://docs.stripe.com/payments/meal-vouchers/fr-meal-vouchers/set-up-restaurant)
type FRMealVouchersOnboarding struct {
	APIResource
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// Name of the restaurant.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Postal code of the restaurant.
	PostalCode string `json:"postal_code"`
	// This represents the onboarding state of the restaurant on different providers.
	Providers *FRMealVouchersOnboardingProviders `json:"providers"`
	// SIRET number associated with the restaurant.
	Siret string `json:"siret"`
}

// FRMealVouchersOnboardingList is a list of FrMealVouchersOnboardings as retrieved from a list endpoint.
type FRMealVouchersOnboardingList struct {
	APIResource
	ListMeta
	Data []*FRMealVouchersOnboarding `json:"data"`
}
