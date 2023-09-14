//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The reason(s) the card logo was rejected.
type IssuingCardDesignRejectionReasonsCardLogo string

// List of values that IssuingCardDesignRejectionReasonsCardLogo can take
const (
	IssuingCardDesignRejectionReasonsCardLogoGeographicLocation  IssuingCardDesignRejectionReasonsCardLogo = "geographic_location"
	IssuingCardDesignRejectionReasonsCardLogoInappropriate       IssuingCardDesignRejectionReasonsCardLogo = "inappropriate"
	IssuingCardDesignRejectionReasonsCardLogoNetworkName         IssuingCardDesignRejectionReasonsCardLogo = "network_name"
	IssuingCardDesignRejectionReasonsCardLogoNonBinaryImage      IssuingCardDesignRejectionReasonsCardLogo = "non_binary_image"
	IssuingCardDesignRejectionReasonsCardLogoNonFiatCurrency     IssuingCardDesignRejectionReasonsCardLogo = "non_fiat_currency"
	IssuingCardDesignRejectionReasonsCardLogoOther               IssuingCardDesignRejectionReasonsCardLogo = "other"
	IssuingCardDesignRejectionReasonsCardLogoOtherEntity         IssuingCardDesignRejectionReasonsCardLogo = "other_entity"
	IssuingCardDesignRejectionReasonsCardLogoPromotionalMaterial IssuingCardDesignRejectionReasonsCardLogo = "promotional_material"
)

// The reason(s) the carrier text was rejected.
type IssuingCardDesignRejectionReasonsCarrierText string

// List of values that IssuingCardDesignRejectionReasonsCarrierText can take
const (
	IssuingCardDesignRejectionReasonsCarrierTextGeographicLocation  IssuingCardDesignRejectionReasonsCarrierText = "geographic_location"
	IssuingCardDesignRejectionReasonsCarrierTextInappropriate       IssuingCardDesignRejectionReasonsCarrierText = "inappropriate"
	IssuingCardDesignRejectionReasonsCarrierTextNetworkName         IssuingCardDesignRejectionReasonsCarrierText = "network_name"
	IssuingCardDesignRejectionReasonsCarrierTextNonFiatCurrency     IssuingCardDesignRejectionReasonsCarrierText = "non_fiat_currency"
	IssuingCardDesignRejectionReasonsCarrierTextOther               IssuingCardDesignRejectionReasonsCarrierText = "other"
	IssuingCardDesignRejectionReasonsCarrierTextOtherEntity         IssuingCardDesignRejectionReasonsCarrierText = "other_entity"
	IssuingCardDesignRejectionReasonsCarrierTextPromotionalMaterial IssuingCardDesignRejectionReasonsCarrierText = "promotional_material"
)

// Whether this card design can be used to create cards.
type IssuingCardDesignStatus string

// List of values that IssuingCardDesignStatus can take
const (
	IssuingCardDesignStatusActive   IssuingCardDesignStatus = "active"
	IssuingCardDesignStatusInactive IssuingCardDesignStatus = "inactive"
	IssuingCardDesignStatusRejected IssuingCardDesignStatus = "rejected"
	IssuingCardDesignStatusReview   IssuingCardDesignStatus = "review"
)

// Only return card designs with the given preferences.
type IssuingCardDesignListPreferencesParams struct {
	// Only return the card design that is set as the account default. A connected account will use the Connect platform's default if no card design is set as the account default.
	AccountDefault *bool `form:"account_default"`
	// Only return the card design that is set as the Connect platform's default. This parameter is only applicable to connected accounts.
	PlatformDefault *bool `form:"platform_default"`
}

// Returns a list of card design objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
type IssuingCardDesignListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Only return card designs with the given lookup keys.
	LookupKeys []*string `form:"lookup_keys"`
	// Only return card designs with the given preferences.
	Preferences *IssuingCardDesignListPreferencesParams `form:"preferences"`
	// Only return card designs with the given status.
	Status *string `form:"status"`
}

// AddExpand appends a new field to expand.
func (p *IssuingCardDesignListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Hash containing carrier text, for use with card bundles that support carrier text.
type IssuingCardDesignCarrierTextParams struct {
	// The footer body text of the carrier letter.
	FooterBody *string `form:"footer_body"`
	// The footer title text of the carrier letter.
	FooterTitle *string `form:"footer_title"`
	// The header body text of the carrier letter.
	HeaderBody *string `form:"header_body"`
	// The header title text of the carrier letter.
	HeaderTitle *string `form:"header_title"`
}

// Information on whether this card design is used to create cards when one is not specified.
type IssuingCardDesignPreferencesParams struct {
	// Whether this card design is used to create cards when one is not specified. A connected account will use the Connect platform's default if no card design is set as the account default.
	AccountDefault *bool `form:"account_default"`
}

// Creates a card design object.
type IssuingCardDesignParams struct {
	Params `form:"*"`
	// The card bundle object belonging to this card design.
	CardBundle *string `form:"card_bundle"`
	// The file for the card logo, for use with card bundles that support card logos.
	CardLogo *string `form:"card_logo"`
	// Hash containing carrier text, for use with card bundles that support carrier text.
	CarrierText *IssuingCardDesignCarrierTextParams `form:"carrier_text"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// A lookup key used to retrieve card designs dynamically from a static string. This may be up to 200 characters.
	LookupKey *string `form:"lookup_key"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Friendly display name. Providing an empty string will set the field to null.
	Name *string `form:"name"`
	// Information on whether this card design is used to create cards when one is not specified.
	Preferences *IssuingCardDesignPreferencesParams `form:"preferences"`
	// If set to true, will atomically remove the lookup key from the existing card design, and assign it to this card design.
	TransferLookupKey *bool `form:"transfer_lookup_key"`
}

// AddExpand appends a new field to expand.
func (p *IssuingCardDesignParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *IssuingCardDesignParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Hash containing carrier text, for use with card bundles that support carrier text.
type IssuingCardDesignCarrierText struct {
	// The footer body text of the carrier letter.
	FooterBody string `json:"footer_body"`
	// The footer title text of the carrier letter.
	FooterTitle string `json:"footer_title"`
	// The header body text of the carrier letter.
	HeaderBody string `json:"header_body"`
	// The header title text of the carrier letter.
	HeaderTitle string `json:"header_title"`
}
type IssuingCardDesignPreferences struct {
	// Whether this card design is used to create cards when one is not specified. A connected account will use the Connect platform's default if no card design is set as the account default.
	AccountDefault bool `json:"account_default"`
	// Whether this card design is used to create cards when one is not specified and an account default for this connected account does not exist.
	PlatformDefault bool `json:"platform_default"`
}
type IssuingCardDesignRejectionReasons struct {
	// The reason(s) the card logo was rejected.
	CardLogo []IssuingCardDesignRejectionReasonsCardLogo `json:"card_logo"`
	// The reason(s) the carrier text was rejected.
	CarrierText []IssuingCardDesignRejectionReasonsCarrierText `json:"carrier_text"`
}

// A Card Design is a logical grouping of a Card Bundle, card logo, and carrier text that represents a product line.
type IssuingCardDesign struct {
	APIResource
	// The card bundle object belonging to this card design.
	CardBundle *IssuingCardBundle `json:"card_bundle"`
	// The file for the card logo, for use with card bundles that support card logos.
	CardLogo *File `json:"card_logo"`
	// Hash containing carrier text, for use with card bundles that support carrier text.
	CarrierText *IssuingCardDesignCarrierText `json:"carrier_text"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// A lookup key used to retrieve card designs dynamically from a static string. This may be up to 200 characters.
	LookupKey string `json:"lookup_key"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// Friendly display name.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object           string                             `json:"object"`
	Preferences      *IssuingCardDesignPreferences      `json:"preferences"`
	RejectionReasons *IssuingCardDesignRejectionReasons `json:"rejection_reasons"`
	// Whether this card design can be used to create cards.
	Status IssuingCardDesignStatus `json:"status"`
}

// IssuingCardDesignList is a list of CardDesigns as retrieved from a list endpoint.
type IssuingCardDesignList struct {
	APIResource
	ListMeta
	Data []*IssuingCardDesign `json:"data"`
}

// UnmarshalJSON handles deserialization of an IssuingCardDesign.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *IssuingCardDesign) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type issuingCardDesign IssuingCardDesign
	var v issuingCardDesign
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = IssuingCardDesign(v)
	return nil
}
