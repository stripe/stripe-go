//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Returns a list of ProductCatalogImport objects.
type V2CommerceProductCatalogImportListParams struct {
	Params `form:"*"`
	// Filter for objects created at the specified timestamp.
	// Must be an RFC 3339 date & time value, for example: 2022-09-18T13:22:00Z.
	Created *time.Time `form:"created" json:"created,omitempty"`
	// Filter for objects created after the specified timestamp.
	// Must be an RFC 3339 date & time value, for example: 2022-09-18T13:22:00Z.
	CreatedGt *time.Time `form:"created_gt" json:"created_gt,omitempty"`
	// Filter for objects created on or after the specified timestamp.
	// Must be an RFC 3339 date & time value, for example: 2022-09-18T13:22:00Z.
	CreatedGTE *time.Time `form:"created_gte" json:"created_gte,omitempty"`
	// Filter for objects created before the specified timestamp.
	// Must be an RFC 3339 date & time value, for example: 2022-09-18T13:22:00Z.
	CreatedLT *time.Time `form:"created_lt" json:"created_lt,omitempty"`
	// Filter for objects created on or before the specified timestamp.
	// Must be an RFC 3339 date & time value, for example: 2022-09-18T13:22:00Z.
	CreatedLte *time.Time `form:"created_lte" json:"created_lte,omitempty"`
	// Filter by the type of feed data being imported.
	FeedType *string `form:"feed_type" json:"feed_type,omitempty"`
	// The maximum number of results per page.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter by import status.
	Status *string `form:"status" json:"status,omitempty"`
}

// Creates a ProductCatalogImport.
type V2CommerceProductCatalogImportParams struct {
	Params `form:"*"`
	// The type of catalog data to import.
	FeedType *string `form:"feed_type" json:"feed_type,omitempty"`
	// Additional information about the import in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The strategy for handling existing catalog data during import.
	Mode *string `form:"mode" json:"mode,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CommerceProductCatalogImportParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Creates a ProductCatalogImport.
type V2CommerceProductCatalogImportCreateParams struct {
	Params `form:"*"`
	// The type of catalog data to import.
	FeedType *string `form:"feed_type" json:"feed_type"`
	// Additional information about the import in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata"`
	// The strategy for handling existing catalog data during import.
	Mode *string `form:"mode" json:"mode"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CommerceProductCatalogImportCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves a ProductCatalogImport by ID.
type V2CommerceProductCatalogImportRetrieveParams struct {
	Params `form:"*"`
}
