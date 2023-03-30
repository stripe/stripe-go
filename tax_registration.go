//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The status of the registration. This field is present for convenience and can be deduced from `active_from` and `expires_at`.
type TaxRegistrationStatus string

// List of values that TaxRegistrationStatus can take
const (
	TaxRegistrationStatusActive    TaxRegistrationStatus = "active"
	TaxRegistrationStatusExpired   TaxRegistrationStatus = "expired"
	TaxRegistrationStatusScheduled TaxRegistrationStatus = "scheduled"
)

// The type of the registration. See [our guide](https://stripe.com/docs/tax/registering) for more information about registration types.
type TaxRegistrationType string

// List of values that TaxRegistrationType can take
const (
	TaxRegistrationTypeDomesticSmallSeller TaxRegistrationType = "domestic_small_seller"
	TaxRegistrationTypeIoss                TaxRegistrationType = "ioss"
	TaxRegistrationTypeSimplified          TaxRegistrationType = "simplified"
	TaxRegistrationTypeStandard            TaxRegistrationType = "standard"
	TaxRegistrationTypeVATOssNonUnion      TaxRegistrationType = "vat_oss_non_union"
	TaxRegistrationTypeVATOssUnion         TaxRegistrationType = "vat_oss_union"
)

// Returns a list of Tax Registration objects.
type TaxRegistrationListParams struct {
	ListParams `form:"*"`
	// The status of the Tax Registration.
	Status *string `form:"status"`
}

// Creates a new Tax Registration object.
type TaxRegistrationParams struct {
	Params `form:"*"`
	// Time at which the registration becomes active. Measured in seconds since the Unix epoch.
	ActiveFrom *int64 `form:"active_from"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country"`
	// If set, the registration stops being active at this time. If not set, the registration will be active indefinitely. Measured in seconds since the Unix epoch.
	ExpiresAt *int64 `form:"expires_at"`
	// State, county, province, or region.
	State *string `form:"state"`
	// The type of the Tax Registration. See [our guide](https://stripe.com/docs/tax/registering) for more information about registration types.
	Type *string `form:"type"`
}

// A Tax `Registration` lets us know that your business is registered to collect tax on payments within a region, enabling you to [automatically collect tax](https://stripe.com/docs/tax).
//
// Stripe will not register on your behalf with the relevant authorities when you create a Tax `Registration` object. For more information on how to register to collect tax, see [our guide](https://stripe.com/docs/tax/registering).
type TaxRegistration struct {
	APIResource
	// Time at which the registration becomes active. Measured in seconds since the Unix epoch.
	ActiveFrom int64 `json:"active_from"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// If set, the registration stops being active at this time. If not set, the registration will be active indefinitely. Measured in seconds since the Unix epoch.
	ExpiresAt int64 `json:"expires_at"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// State, county, province, or region.
	State string `json:"state"`
	// The status of the registration. This field is present for convenience and can be deduced from `active_from` and `expires_at`.
	Status TaxRegistrationStatus `json:"status"`
	// The type of the registration. See [our guide](https://stripe.com/docs/tax/registering) for more information about registration types.
	Type TaxRegistrationType `json:"type"`
}

// TaxRegistrationList is a list of Registrations as retrieved from a list endpoint.
type TaxRegistrationList struct {
	APIResource
	ListMeta
	Data []*TaxRegistration `json:"data"`
}
