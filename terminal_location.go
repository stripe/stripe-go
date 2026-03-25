//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Deletes a Location object.
type TerminalLocationParams struct {
	Params `form:"*"`
	// The full address of the location. You can't change the location's `country`. If you need to modify the `country` field, create a new `Location` object and re-register any existing readers to that location.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// The Kana variation of the full address of the location (Japan only).
	AddressKana *TerminalLocationAddressKanaParams `form:"address_kana" json:"address_kana,omitempty"`
	// The Kanji variation of the full address of the location (Japan only).
	AddressKanji *TerminalLocationAddressKanjiParams `form:"address_kanji" json:"address_kanji,omitempty"`
	// The ID of a configuration that will be used to customize all readers in this location.
	ConfigurationOverrides *string `form:"configuration_overrides" json:"configuration_overrides,omitempty"`
	// A name for the location. Maximum length is 1000 characters.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// The Kana variation of the name for the location (Japan only). Maximum length is 1000 characters.
	DisplayNameKana *string `form:"display_name_kana" json:"display_name_kana,omitempty"`
	// The Kanji variation of the name for the location (Japan only). Maximum length is 1000 characters.
	DisplayNameKanji *string `form:"display_name_kanji" json:"display_name_kanji,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The phone number for the location.
	Phone       *string                            `form:"phone" json:"phone,omitempty"`
	UnsetFields []TerminalLocationParamsUnsetField `form:"-" json:"-"`
}

// TerminalLocationParamsUnsetField is the list of fields that can be cleared/unset on TerminalLocationParams.
type TerminalLocationParamsUnsetField string

const (
	TerminalLocationParamsUnsetFieldConfigurationOverrides TerminalLocationParamsUnsetField = "configuration_overrides"
	TerminalLocationParamsUnsetFieldDisplayName            TerminalLocationParamsUnsetField = "display_name"
	TerminalLocationParamsUnsetFieldDisplayNameKana        TerminalLocationParamsUnsetField = "display_name_kana"
	TerminalLocationParamsUnsetFieldDisplayNameKanji       TerminalLocationParamsUnsetField = "display_name_kanji"
	TerminalLocationParamsUnsetFieldMetadata               TerminalLocationParamsUnsetField = "metadata"
	TerminalLocationParamsUnsetFieldPhone                  TerminalLocationParamsUnsetField = "phone"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *TerminalLocationParams) AddUnsetField(field TerminalLocationParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *TerminalLocationParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *TerminalLocationParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The Kana variation of the full address of the location (Japan only).
type TerminalLocationAddressKanaParams struct {
	// City or ward.
	City *string `form:"city" json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country,omitempty"`
	// Block or building number.
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Building details.
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// Postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// Prefecture.
	State *string `form:"state" json:"state,omitempty"`
	// Town or cho-me.
	Town *string `form:"town" json:"town,omitempty"`
}

// The Kanji variation of the full address of the location (Japan only).
type TerminalLocationAddressKanjiParams struct {
	// City or ward.
	City *string `form:"city" json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country,omitempty"`
	// Block or building number.
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Building details.
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// Postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// Prefecture.
	State *string `form:"state" json:"state,omitempty"`
	// Town or cho-me.
	Town *string `form:"town" json:"town,omitempty"`
}

// Returns a list of Location objects.
type TerminalLocationListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TerminalLocationListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Deletes a Location object.
type TerminalLocationDeleteParams struct {
	Params `form:"*"`
}

// Retrieves a Location object.
type TerminalLocationRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TerminalLocationRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The Kana variation of the full address of the location (Japan only).
type TerminalLocationUpdateAddressKanaParams struct {
	// City or ward.
	City *string `form:"city" json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country,omitempty"`
	// Block or building number.
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Building details.
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// Postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// Prefecture.
	State *string `form:"state" json:"state,omitempty"`
	// Town or cho-me.
	Town *string `form:"town" json:"town,omitempty"`
}

// The Kanji variation of the full address of the location (Japan only).
type TerminalLocationUpdateAddressKanjiParams struct {
	// City or ward.
	City *string `form:"city" json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country,omitempty"`
	// Block or building number.
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Building details.
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// Postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// Prefecture.
	State *string `form:"state" json:"state,omitempty"`
	// Town or cho-me.
	Town *string `form:"town" json:"town,omitempty"`
}

// Updates a Location object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
type TerminalLocationUpdateParams struct {
	Params `form:"*"`
	// The full address of the location. You can't change the location's `country`. If you need to modify the `country` field, create a new `Location` object and re-register any existing readers to that location.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// The Kana variation of the full address of the location (Japan only).
	AddressKana *TerminalLocationUpdateAddressKanaParams `form:"address_kana" json:"address_kana,omitempty"`
	// The Kanji variation of the full address of the location (Japan only).
	AddressKanji *TerminalLocationUpdateAddressKanjiParams `form:"address_kanji" json:"address_kanji,omitempty"`
	// The ID of a configuration that will be used to customize all readers in this location.
	ConfigurationOverrides *string `form:"configuration_overrides" json:"configuration_overrides,omitempty"`
	// A name for the location.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// The Kana variation of the name for the location (Japan only).
	DisplayNameKana *string `form:"display_name_kana" json:"display_name_kana,omitempty"`
	// The Kanji variation of the name for the location (Japan only).
	DisplayNameKanji *string `form:"display_name_kanji" json:"display_name_kanji,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The phone number for the location.
	Phone       *string                                  `form:"phone" json:"phone,omitempty"`
	UnsetFields []TerminalLocationUpdateParamsUnsetField `form:"-" json:"-"`
}

// TerminalLocationUpdateParamsUnsetField is the list of fields that can be cleared/unset on TerminalLocationUpdateParams.
type TerminalLocationUpdateParamsUnsetField string

const (
	TerminalLocationUpdateParamsUnsetFieldConfigurationOverrides TerminalLocationUpdateParamsUnsetField = "configuration_overrides"
	TerminalLocationUpdateParamsUnsetFieldDisplayName            TerminalLocationUpdateParamsUnsetField = "display_name"
	TerminalLocationUpdateParamsUnsetFieldDisplayNameKana        TerminalLocationUpdateParamsUnsetField = "display_name_kana"
	TerminalLocationUpdateParamsUnsetFieldDisplayNameKanji       TerminalLocationUpdateParamsUnsetField = "display_name_kanji"
	TerminalLocationUpdateParamsUnsetFieldMetadata               TerminalLocationUpdateParamsUnsetField = "metadata"
	TerminalLocationUpdateParamsUnsetFieldPhone                  TerminalLocationUpdateParamsUnsetField = "phone"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *TerminalLocationUpdateParams) AddUnsetField(field TerminalLocationUpdateParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *TerminalLocationUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *TerminalLocationUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The Kana variation of the full address of the location (Japan only).
type TerminalLocationCreateAddressKanaParams struct {
	// City or ward.
	City *string `form:"city" json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country,omitempty"`
	// Block or building number.
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Building details.
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// Postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// Prefecture.
	State *string `form:"state" json:"state,omitempty"`
	// Town or cho-me.
	Town *string `form:"town" json:"town,omitempty"`
}

// The Kanji variation of the full address of the location (Japan only).
type TerminalLocationCreateAddressKanjiParams struct {
	// City or ward.
	City *string `form:"city" json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country,omitempty"`
	// Block or building number.
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Building details.
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// Postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// Prefecture.
	State *string `form:"state" json:"state,omitempty"`
	// Town or cho-me.
	Town *string `form:"town" json:"town,omitempty"`
}

// Creates a new Location object.
// For further details, including which address fields are required in each country, see the [Manage locations](https://docs.stripe.com/docs/terminal/fleet/locations) guide.
type TerminalLocationCreateParams struct {
	Params `form:"*"`
	// The full address of the location.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// The Kana variation of the full address of the location (Japan only).
	AddressKana *TerminalLocationCreateAddressKanaParams `form:"address_kana" json:"address_kana,omitempty"`
	// The Kanji variation of the full address of the location (Japan only).
	AddressKanji *TerminalLocationCreateAddressKanjiParams `form:"address_kanji" json:"address_kanji,omitempty"`
	// The ID of a configuration that will be used to customize all readers in this location.
	ConfigurationOverrides *string `form:"configuration_overrides" json:"configuration_overrides,omitempty"`
	// A name for the location. Maximum length is 1000 characters.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// The Kana variation of the name for the location (Japan only). Maximum length is 1000 characters.
	DisplayNameKana *string `form:"display_name_kana" json:"display_name_kana,omitempty"`
	// The Kanji variation of the name for the location (Japan only). Maximum length is 1000 characters.
	DisplayNameKanji *string `form:"display_name_kanji" json:"display_name_kanji,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The phone number for the location.
	Phone       *string                                  `form:"phone" json:"phone,omitempty"`
	UnsetFields []TerminalLocationCreateParamsUnsetField `form:"-" json:"-"`
}

// TerminalLocationCreateParamsUnsetField is the list of fields that can be cleared/unset on TerminalLocationCreateParams.
type TerminalLocationCreateParamsUnsetField string

const (
	TerminalLocationCreateParamsUnsetFieldMetadata TerminalLocationCreateParamsUnsetField = "metadata"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *TerminalLocationCreateParams) AddUnsetField(field TerminalLocationCreateParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *TerminalLocationCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *TerminalLocationCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

type TerminalLocationAddressKana struct {
	// City/Ward.
	City string `json:"city"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country"`
	// Block/Building number.
	Line1 string `json:"line1"`
	// Building details.
	Line2 string `json:"line2"`
	// ZIP or postal code.
	PostalCode string `json:"postal_code"`
	// Prefecture.
	State string `json:"state"`
	// Town/cho-me.
	Town string `json:"town"`
}
type TerminalLocationAddressKanji struct {
	// City/Ward.
	City string `json:"city"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country"`
	// Block/Building number.
	Line1 string `json:"line1"`
	// Building details.
	Line2 string `json:"line2"`
	// ZIP or postal code.
	PostalCode string `json:"postal_code"`
	// Prefecture.
	State string `json:"state"`
	// Town/cho-me.
	Town string `json:"town"`
}

// A Location represents a grouping of readers.
//
// Related guide: [Fleet management](https://docs.stripe.com/terminal/fleet/locations)
type TerminalLocation struct {
	APIResource
	Address      *Address                      `json:"address"`
	AddressKana  *TerminalLocationAddressKana  `json:"address_kana,omitempty"`
	AddressKanji *TerminalLocationAddressKanji `json:"address_kanji,omitempty"`
	// The ID of a configuration that will be used to customize all readers in this location.
	ConfigurationOverrides string `json:"configuration_overrides,omitempty"`
	Deleted                bool   `json:"deleted,omitempty"`
	// The display name of the location.
	DisplayName string `json:"display_name"`
	// The Kana variation of the display name of the location.
	DisplayNameKana string `json:"display_name_kana,omitempty"`
	// The Kanji variation of the display name of the location.
	DisplayNameKanji string `json:"display_name_kanji,omitempty"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The phone number of the location.
	Phone string `json:"phone,omitempty"`
}

// TerminalLocationList is a list of Locations as retrieved from a list endpoint.
type TerminalLocationList struct {
	APIResource
	ListMeta
	Data []*TerminalLocation `json:"data"`
}

// UnmarshalJSON handles deserialization of a TerminalLocation.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (t *TerminalLocation) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		t.ID = id
		return nil
	}

	type terminalLocation TerminalLocation
	var v terminalLocation
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*t = TerminalLocation(v)
	return nil
}
