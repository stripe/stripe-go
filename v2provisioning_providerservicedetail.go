//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

type V2ProvisioningProviderServiceDetailAllowedUpdateDirection string

// List of values that V2ProvisioningProviderServiceDetailAllowedUpdateDirection can take
const (
	V2ProvisioningProviderServiceDetailAllowedUpdateDirectionAny  V2ProvisioningProviderServiceDetailAllowedUpdateDirection = "any"
	V2ProvisioningProviderServiceDetailAllowedUpdateDirectionDown V2ProvisioningProviderServiceDetailAllowedUpdateDirection = "down"
	V2ProvisioningProviderServiceDetailAllowedUpdateDirectionUp   V2ProvisioningProviderServiceDetailAllowedUpdateDirection = "up"
)

// Availability of the service.
type V2ProvisioningProviderServiceDetailAvailability string

// List of values that V2ProvisioningProviderServiceDetailAvailability can take
const (
	V2ProvisioningProviderServiceDetailAvailabilityAvailable    V2ProvisioningProviderServiceDetailAvailability = "available"
	V2ProvisioningProviderServiceDetailAvailabilityNotInCountry V2ProvisioningProviderServiceDetailAvailability = "not_in_country"
	V2ProvisioningProviderServiceDetailAvailabilityUnavailable  V2ProvisioningProviderServiceDetailAvailability = "unavailable"
)

type V2ProvisioningProviderServiceDetailConstraintType string

// List of values that V2ProvisioningProviderServiceDetailConstraintType can take
const (
	V2ProvisioningProviderServiceDetailConstraintTypeCount                         V2ProvisioningProviderServiceDetailConstraintType = "count"
	V2ProvisioningProviderServiceDetailConstraintTypeMutualExclusionAllowedUpdates V2ProvisioningProviderServiceDetailConstraintType = "mutual_exclusion_allowed_updates"
)

// Kind of the service.
type V2ProvisioningProviderServiceDetailKind string

// List of values that V2ProvisioningProviderServiceDetailKind can take
const (
	V2ProvisioningProviderServiceDetailKindDeployable V2ProvisioningProviderServiceDetailKind = "deployable"
	V2ProvisioningProviderServiceDetailKindPlan       V2ProvisioningProviderServiceDetailKind = "plan"
)

type V2ProvisioningProviderServiceDetailPricingComponentOptionPaidType string

// List of values that V2ProvisioningProviderServiceDetailPricingComponentOptionPaidType can take
const (
	V2ProvisioningProviderServiceDetailPricingComponentOptionPaidTypeFree     V2ProvisioningProviderServiceDetailPricingComponentOptionPaidType = "free"
	V2ProvisioningProviderServiceDetailPricingComponentOptionPaidTypeFreeform V2ProvisioningProviderServiceDetailPricingComponentOptionPaidType = "freeform"
)

type V2ProvisioningProviderServiceDetailPricingComponentOptionType string

// List of values that V2ProvisioningProviderServiceDetailPricingComponentOptionType can take
const (
	V2ProvisioningProviderServiceDetailPricingComponentOptionTypeFree V2ProvisioningProviderServiceDetailPricingComponentOptionType = "free"
	V2ProvisioningProviderServiceDetailPricingComponentOptionTypePaid V2ProvisioningProviderServiceDetailPricingComponentOptionType = "paid"
)

type V2ProvisioningProviderServiceDetailPricingPaidType string

// List of values that V2ProvisioningProviderServiceDetailPricingPaidType can take
const (
	V2ProvisioningProviderServiceDetailPricingPaidTypeFree     V2ProvisioningProviderServiceDetailPricingPaidType = "free"
	V2ProvisioningProviderServiceDetailPricingPaidTypeFreeform V2ProvisioningProviderServiceDetailPricingPaidType = "freeform"
)

type V2ProvisioningProviderServiceDetailPricingPaidPricingType string

// List of values that V2ProvisioningProviderServiceDetailPricingPaidPricingType can take
const (
	V2ProvisioningProviderServiceDetailPricingPaidPricingTypeFree     V2ProvisioningProviderServiceDetailPricingPaidPricingType = "free"
	V2ProvisioningProviderServiceDetailPricingPaidPricingTypeFreeform V2ProvisioningProviderServiceDetailPricingPaidPricingType = "freeform"
)

type V2ProvisioningProviderServiceDetailPricingType string

// List of values that V2ProvisioningProviderServiceDetailPricingType can take
const (
	V2ProvisioningProviderServiceDetailPricingTypeComponent V2ProvisioningProviderServiceDetailPricingType = "component"
	V2ProvisioningProviderServiceDetailPricingTypeFree      V2ProvisioningProviderServiceDetailPricingType = "free"
	V2ProvisioningProviderServiceDetailPricingTypePaid      V2ProvisioningProviderServiceDetailPricingType = "paid"
)

// Scope of the service.
type V2ProvisioningProviderServiceDetailScope string

// List of values that V2ProvisioningProviderServiceDetailScope can take
const (
	V2ProvisioningProviderServiceDetailScopeAccount V2ProvisioningProviderServiceDetailScope = "account"
	V2ProvisioningProviderServiceDetailScopeProject V2ProvisioningProviderServiceDetailScope = "project"
)

// Updates allowed for resources using this service.
type V2ProvisioningProviderServiceDetailAllowedUpdate struct {
	Direction V2ProvisioningProviderServiceDetailAllowedUpdateDirection `json:"direction"`
	Service   string                                                    `json:"service"`
}
type V2ProvisioningProviderServiceDetailConstraintCount struct {
	AtMost int64 `json:"at_most"`
}

// Constraints on resources using this service.
type V2ProvisioningProviderServiceDetailConstraint struct {
	Count                         *V2ProvisioningProviderServiceDetailConstraintCount `json:"count,omitempty"`
	MutualExclusionAllowedUpdates bool                                                `json:"mutual_exclusion_allowed_updates,omitempty"`
	Type                          V2ProvisioningProviderServiceDetailConstraintType   `json:"type"`
}
type V2ProvisioningProviderServiceDetailPricingComponentOptionPaid struct {
	Description string                                                            `json:"description,omitempty"`
	Freeform    string                                                            `json:"freeform,omitempty"`
	Type        V2ProvisioningProviderServiceDetailPricingComponentOptionPaidType `json:"type"`
}
type V2ProvisioningProviderServiceDetailPricingComponentOption struct {
	IsDefault      bool                                                           `json:"is_default,omitempty"`
	Paid           *V2ProvisioningProviderServiceDetailPricingComponentOptionPaid `json:"paid"`
	ParentServices []string                                                       `json:"parent_services"`
	Type           V2ProvisioningProviderServiceDetailPricingComponentOptionType  `json:"type"`
}
type V2ProvisioningProviderServiceDetailPricingComponent struct {
	Options []*V2ProvisioningProviderServiceDetailPricingComponentOption `json:"options"`
}

// Legacy compatibility field for top-level paid pricing.
// Mirrors the single paid pricing entry when only one exists, or the entry marked
// `is_default`. If multiple paid pricing entries exist and none is default, this field
// is unset.
type V2ProvisioningProviderServiceDetailPricingPaid struct {
	Description string                                             `json:"description,omitempty"`
	Freeform    string                                             `json:"freeform,omitempty"`
	Type        V2ProvisioningProviderServiceDetailPricingPaidType `json:"type"`
}

// Canonical top-level paid pricing entries for this service.
// When multiple entries are present, callers should read this field instead of `paid`.
type V2ProvisioningProviderServiceDetailPricingPaidPricing struct {
	Configuration map[string]any                                            `json:"configuration"`
	Description   string                                                    `json:"description,omitempty"`
	Freeform      string                                                    `json:"freeform,omitempty"`
	IsDefault     bool                                                      `json:"is_default,omitempty"`
	Type          V2ProvisioningProviderServiceDetailPricingPaidPricingType `json:"type"`
}

// Pricing details for the service.
type V2ProvisioningProviderServiceDetailPricing struct {
	Component *V2ProvisioningProviderServiceDetailPricingComponent `json:"component"`
	// Legacy compatibility field for top-level paid pricing.
	// Mirrors the single paid pricing entry when only one exists, or the entry marked
	// `is_default`. If multiple paid pricing entries exist and none is default, this field
	// is unset.
	Paid *V2ProvisioningProviderServiceDetailPricingPaid `json:"paid"`
	// Canonical top-level paid pricing entries for this service.
	// When multiple entries are present, callers should read this field instead of `paid`.
	PaidPricing []*V2ProvisioningProviderServiceDetailPricingPaidPricing `json:"paid_pricing"`
	Type        V2ProvisioningProviderServiceDetailPricingType           `json:"type"`
}

// The `ProviderServiceDetail` resource represents a service offered by a
// provider in the catalog.
type V2ProvisioningProviderServiceDetail struct {
	APIResource
	// Updates allowed for resources using this service.
	AllowedUpdates []*V2ProvisioningProviderServiceDetailAllowedUpdate `json:"allowed_updates"`
	// Availability of the service.
	Availability V2ProvisioningProviderServiceDetailAvailability `json:"availability"`
	// Categories the service belongs to.
	Categories []string `json:"categories"`
	// Schema describing the configuration accepted by this service.
	ConfigurationSchema map[string]any `json:"configuration_schema"`
	// Constraints on resources using this service.
	Constraints []*V2ProvisioningProviderServiceDetailConstraint `json:"constraints"`
	// Time at which the service was created.
	Created time.Time `json:"created"`
	// Description of the service.
	Description string `json:"description"`
	// Denormalized from the parent Provider. If a Provider's partition changes, re-sync its services.
	// proto3 scalar defaults apply: if unset, this value is `false`.
	Development bool `json:"development"`
	// Group the service belongs to, used to organize related services.
	Group string `json:"group,omitempty"`
	// Unique identifier for the provider service.
	ID string `json:"id"`
	// Kind of the service.
	Kind V2ProvisioningProviderServiceDetailKind `json:"kind"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// URL of additional context about the service intended for LLM consumption.
	LlmContext string `json:"llm_context,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Pricing details for the service.
	Pricing *V2ProvisioningProviderServiceDetailPricing `json:"pricing"`
	// Identifier of the provider that offers this service.
	Provider string `json:"provider"`
	// Human-readable name of the provider that offers this service.
	ProviderName string `json:"provider_name"`
	// Scope of the service.
	Scope V2ProvisioningProviderServiceDetailScope `json:"scope"`
	// Identifier of the service, unique within its provider.
	ServiceID string `json:"service_id"`
	// Deprecated: use allowed_updates instead.
	UpdateableTo []string `json:"updateable_to"`
}
