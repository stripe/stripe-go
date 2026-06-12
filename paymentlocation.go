//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Supported meal voucher issuers for card payments.
type PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCard string

// List of values that PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCard can take
const (
	PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCardBimpli  PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCard = "bimpli"
	PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCardEdenred PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCard = "edenred"
	PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCardPluxee  PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCard = "pluxee"
	PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCardUp      PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCard = "up"
)

// Supported meal voucher issuers for card present payments.
type PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCardPresent string

// List of values that PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCardPresent can take
const (
	PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCardPresentBimpli  PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCardPresent = "bimpli"
	PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCardPresentEdenred PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCardPresent = "edenred"
	PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCardPresentPluxee  PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCardPresent = "pluxee"
	PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCardPresentUp      PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCardPresent = "up"
)

// Delete a Payment Location.
type PaymentLocationParams struct {
	Params `form:"*"`
	// The full address of the location.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// Identification numbers associated with the location.
	BusinessRegistration *PaymentLocationBusinessRegistrationParams `form:"business_registration" json:"business_registration,omitempty"`
	// A name for the location.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Pass true when updating location fields that trigger onboarding review for any of the location's active location capabilities. If this parameter isn't set to true, updates that would trigger onboarding review fail. Only applicable for locations with active location capabilities.
	OnboardingDataUpdateAcknowledged *bool `form:"onboarding_data_update_acknowledged" json:"onboarding_data_update_acknowledged,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *PaymentLocationParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Identification numbers associated with the location.
type PaymentLocationBusinessRegistrationParams struct {
	// 14-digit SIRET (Système d'identification du répertoire des établissements) number for the location.
	Siret       *string                                               `form:"siret" json:"siret,omitempty"`
	UnsetFields []PaymentLocationBusinessRegistrationParamsUnsetField `form:"-" json:"-"`
}

// PaymentLocationBusinessRegistrationParamsUnsetField is the list of fields that can be cleared/unset on PaymentLocationBusinessRegistrationParams.
type PaymentLocationBusinessRegistrationParamsUnsetField string

const (
	PaymentLocationBusinessRegistrationParamsUnsetFieldSiret PaymentLocationBusinessRegistrationParamsUnsetField = "siret"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *PaymentLocationBusinessRegistrationParams) AddUnsetField(field PaymentLocationBusinessRegistrationParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// List all Payment Locations.
type PaymentLocationListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *PaymentLocationListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Delete a Payment Location.
type PaymentLocationDeleteParams struct {
	Params `form:"*"`
}

// Retrieve a Payment Location.
type PaymentLocationRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *PaymentLocationRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Identification numbers associated with the location.
type PaymentLocationUpdateBusinessRegistrationParams struct {
	// 14-digit SIRET (Système d'identification du répertoire des établissements) number for the location.
	Siret       *string                                                     `form:"siret" json:"siret,omitempty"`
	UnsetFields []PaymentLocationUpdateBusinessRegistrationParamsUnsetField `form:"-" json:"-"`
}

// PaymentLocationUpdateBusinessRegistrationParamsUnsetField is the list of fields that can be cleared/unset on PaymentLocationUpdateBusinessRegistrationParams.
type PaymentLocationUpdateBusinessRegistrationParamsUnsetField string

const (
	PaymentLocationUpdateBusinessRegistrationParamsUnsetFieldSiret PaymentLocationUpdateBusinessRegistrationParamsUnsetField = "siret"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *PaymentLocationUpdateBusinessRegistrationParams) AddUnsetField(field PaymentLocationUpdateBusinessRegistrationParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Update a Payment Location.
type PaymentLocationUpdateParams struct {
	Params `form:"*"`
	// The full address of the location.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// Identification numbers associated with the location.
	BusinessRegistration *PaymentLocationUpdateBusinessRegistrationParams `form:"business_registration" json:"business_registration,omitempty"`
	// A name for the location.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Pass true when updating location fields that trigger onboarding review for any of the location's active location capabilities. If this parameter isn't set to true, updates that would trigger onboarding review fail. Only applicable for locations with active location capabilities.
	OnboardingDataUpdateAcknowledged *bool `form:"onboarding_data_update_acknowledged" json:"onboarding_data_update_acknowledged,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *PaymentLocationUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Identification numbers associated with the location.
type PaymentLocationCreateBusinessRegistrationParams struct {
	// 14-digit SIRET (Système d'identification du répertoire des établissements) number for the location.
	Siret *string `form:"siret" json:"siret,omitempty"`
}

// Create a Payment Location.
type PaymentLocationCreateParams struct {
	Params `form:"*"`
	// The full address of the location.
	Address *AddressParams `form:"address" json:"address"`
	// Identification numbers associated with the location.
	BusinessRegistration *PaymentLocationCreateBusinessRegistrationParams `form:"business_registration" json:"business_registration,omitempty"`
	// A name for the location.
	DisplayName *string `form:"display_name" json:"display_name"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *PaymentLocationCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Identification numbers associated with the location.
type PaymentLocationBusinessRegistration struct {
	// 14-digit SIRET (Système d'identification du répertoire des établissements) number for the location.
	Siret string `json:"siret"`
}

// Supported meal voucher issuers.
type PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuers struct {
	// Supported meal voucher issuers for card payments.
	Card []PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCard `json:"card"`
	// Supported meal voucher issuers for card present payments.
	CardPresent []PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuersCardPresent `json:"card_present"`
}

// Settings for Conecs French meal voucher capability.
type PaymentLocationCapabilitySettingsFRMealVouchersConecsPayments struct {
	// Supported meal voucher issuers.
	SupportedIssuers *PaymentLocationCapabilitySettingsFRMealVouchersConecsPaymentsSupportedIssuers `json:"supported_issuers"`
}

// The capability settings for the location. Only applicable for locations with requested payment location capabilities.
type PaymentLocationCapabilitySettings struct {
	// Settings for Conecs French meal voucher capability.
	FRMealVouchersConecsPayments *PaymentLocationCapabilitySettingsFRMealVouchersConecsPayments `json:"fr_meal_vouchers_conecs_payments"`
}

// A Payment Location represents a physical location where payments take place.
type PaymentLocation struct {
	APIResource
	Address *Address `json:"address"`
	// Identification numbers associated with the location.
	BusinessRegistration *PaymentLocationBusinessRegistration `json:"business_registration"`
	// The capability settings for the location. Only applicable for locations with requested payment location capabilities.
	CapabilitySettings *PaymentLocationCapabilitySettings `json:"capability_settings"`
	Deleted            bool                               `json:"deleted,omitempty"`
	// The display name of the location.
	DisplayName string `json:"display_name"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}

// PaymentLocationList is a list of PaymentLocations as retrieved from a list endpoint.
type PaymentLocationList struct {
	APIResource
	ListMeta
	Data []*PaymentLocation `json:"data"`
}
