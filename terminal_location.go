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
	Address *AddressParams `form:"address"`
	// The Kana variation of the full address of the location (Japan only).
	AddressKana *TerminalLocationAddressKanaParams `form:"address_kana"`
	// The Kanji variation of the full address of the location (Japan only).
	AddressKanji *TerminalLocationAddressKanjiParams `form:"address_kanji"`
	// The ID of a configuration that will be used to customize all readers in this location.
	ConfigurationOverrides *string `form:"configuration_overrides"`
	// A name for the location. Maximum length is 1000 characters.
	DisplayName *string `form:"display_name"`
	// The Kana variation of the name for the location (Japan only). Maximum length is 1000 characters.
	DisplayNameKana *string `form:"display_name_kana"`
	// The Kanji variation of the name for the location (Japan only). Maximum length is 1000 characters.
	DisplayNameKanji *string `form:"display_name_kanji"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The phone number for the location.
	Phone *string `form:"phone"`
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
	City *string `form:"city"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country"`
	// Block or building number.
	Line1 *string `form:"line1"`
	// Building details.
	Line2 *string `form:"line2"`
	// Postal code.
	PostalCode *string `form:"postal_code"`
	// Prefecture.
	State *string `form:"state"`
	// Town or cho-me.
	Town *string `form:"town"`
}

// The Kanji variation of the full address of the location (Japan only).
type TerminalLocationAddressKanjiParams struct {
	// City or ward.
	City *string `form:"city"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country"`
	// Block or building number.
	Line1 *string `form:"line1"`
	// Building details.
	Line2 *string `form:"line2"`
	// Postal code.
	PostalCode *string `form:"postal_code"`
	// Prefecture.
	State *string `form:"state"`
	// Town or cho-me.
	Town *string `form:"town"`
}

// Returns a list of Location objects.
type TerminalLocationListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
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
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *TerminalLocationRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The Kana variation of the full address of the location (Japan only).
type TerminalLocationUpdateAddressKanaParams struct {
	// City or ward.
	City *string `form:"city"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country"`
	// Block or building number.
	Line1 *string `form:"line1"`
	// Building details.
	Line2 *string `form:"line2"`
	// Postal code.
	PostalCode *string `form:"postal_code"`
	// Prefecture.
	State *string `form:"state"`
	// Town or cho-me.
	Town *string `form:"town"`
}

// The Kanji variation of the full address of the location (Japan only).
type TerminalLocationUpdateAddressKanjiParams struct {
	// City or ward.
	City *string `form:"city"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country"`
	// Block or building number.
	Line1 *string `form:"line1"`
	// Building details.
	Line2 *string `form:"line2"`
	// Postal code.
	PostalCode *string `form:"postal_code"`
	// Prefecture.
	State *string `form:"state"`
	// Town or cho-me.
	Town *string `form:"town"`
}

// Updates a Location object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
type TerminalLocationUpdateParams struct {
	Params `form:"*"`
	// The full address of the location. You can't change the location's `country`. If you need to modify the `country` field, create a new `Location` object and re-register any existing readers to that location.
	Address *AddressParams `form:"address"`
	// The Kana variation of the full address of the location (Japan only).
	AddressKana *TerminalLocationUpdateAddressKanaParams `form:"address_kana"`
	// The Kanji variation of the full address of the location (Japan only).
	AddressKanji *TerminalLocationUpdateAddressKanjiParams `form:"address_kanji"`
	// The ID of a configuration that will be used to customize all readers in this location.
	ConfigurationOverrides *string `form:"configuration_overrides"`
	// A name for the location.
	DisplayName *string `form:"display_name"`
	// The Kana variation of the name for the location (Japan only).
	DisplayNameKana *string `form:"display_name_kana"`
	// The Kanji variation of the name for the location (Japan only).
	DisplayNameKanji *string `form:"display_name_kanji"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The phone number for the location.
	Phone *string `form:"phone"`
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
	City *string `form:"city"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country"`
	// Block or building number.
	Line1 *string `form:"line1"`
	// Building details.
	Line2 *string `form:"line2"`
	// Postal code.
	PostalCode *string `form:"postal_code"`
	// Prefecture.
	State *string `form:"state"`
	// Town or cho-me.
	Town *string `form:"town"`
}

// The Kanji variation of the full address of the location (Japan only).
type TerminalLocationCreateAddressKanjiParams struct {
	// City or ward.
	City *string `form:"city"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country"`
	// Block or building number.
	Line1 *string `form:"line1"`
	// Building details.
	Line2 *string `form:"line2"`
	// Postal code.
	PostalCode *string `form:"postal_code"`
	// Prefecture.
	State *string `form:"state"`
	// Town or cho-me.
	Town *string `form:"town"`
}

// Creates a new Location object.
// For further details, including which address fields are required in each country, see the [Manage locations](https://docs.stripe.com/docs/terminal/fleet/locations) guide.
type TerminalLocationCreateParams struct {
	Params `form:"*"`
	// The full address of the location.
	Address *AddressParams `form:"address"`
	// The Kana variation of the full address of the location (Japan only).
	AddressKana *TerminalLocationCreateAddressKanaParams `form:"address_kana"`
	// The Kanji variation of the full address of the location (Japan only).
	AddressKanji *TerminalLocationCreateAddressKanjiParams `form:"address_kanji"`
	// The ID of a configuration that will be used to customize all readers in this location.
	ConfigurationOverrides *string `form:"configuration_overrides"`
	// A name for the location. Maximum length is 1000 characters.
	DisplayName *string `form:"display_name"`
	// The Kana variation of the name for the location (Japan only). Maximum length is 1000 characters.
	DisplayNameKana *string `form:"display_name_kana"`
	// The Kanji variation of the name for the location (Japan only). Maximum length is 1000 characters.
	DisplayNameKanji *string `form:"display_name_kanji"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The phone number for the location.
	Phone *string `form:"phone"`
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
// Related guide: [Fleet management](https://stripe.com/docs/terminal/fleet/locations)
type TerminalLocation struct {
	APIResource
	Address      *Address                      `json:"address"`
	AddressKana  *TerminalLocationAddressKana  `json:"address_kana"`
	AddressKanji *TerminalLocationAddressKanji `json:"address_kanji"`
	// The ID of a configuration that will be used to customize all readers in this location.
	ConfigurationOverrides string `json:"configuration_overrides"`
	Deleted                bool   `json:"deleted"`
	// The display name of the location.
	DisplayName string `json:"display_name"`
	// The Kana variation of the display name of the location.
	DisplayNameKana string `json:"display_name_kana"`
	// The Kanji variation of the display name of the location.
	DisplayNameKanji string `json:"display_name_kanji"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The phone number of the location.
	Phone string `json:"phone"`
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
