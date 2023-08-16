//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Whether this card design is used to create cards when one is not specified.
type IssuingCardDesignPreference string

// List of values that IssuingCardDesignPreference can take
const (
	IssuingCardDesignPreferenceDefault IssuingCardDesignPreference = "default"
	IssuingCardDesignPreferenceNone    IssuingCardDesignPreference = "none"
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

// Returns a list of card design objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
type IssuingCardDesignListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Only return card designs with the given lookup keys.
	LookupKeys []*string `form:"lookup_keys"`
	// Only return card designs with the given preference.
	Preference *string `form:"preference"`
	// Only return card designs with the given status.
	Status *string `form:"status"`
}

// AddExpand appends a new field to expand.
func (p *IssuingCardDesignListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves a card design object.
type IssuingCardDesignParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// A lookup key used to retrieve card designs dynamically from a static string. This may be up to 200 characters.
	LookupKey *string `form:"lookup_key"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Friendly display name. Providing an empty string will set the field to null.
	Name *string `form:"name"`
	// Whether this card design is used to create cards when one is not specified.
	Preference *string `form:"preference"`
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

// A Card Design is a logical grouping of a Card Bundle, card logo, and carrier text that represents a product line.
type IssuingCardDesign struct {
	APIResource
	// The card bundle object belonging to this card design.
	CardBundle *IssuingCardBundle `json:"card_bundle"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// A lookup key used to retrieve card designs dynamically from a static string. This may be up to 200 characters.
	LookupKey string `json:"lookup_key"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// Friendly display name.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Whether this card design is used to create cards when one is not specified.
	Preference IssuingCardDesignPreference `json:"preference"`
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
