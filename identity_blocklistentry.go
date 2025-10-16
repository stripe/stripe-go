//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The current status of the BlocklistEntry.
type IdentityBlocklistEntryStatus string

// List of values that IdentityBlocklistEntryStatus can take
const (
	IdentityBlocklistEntryStatusActive   IdentityBlocklistEntryStatus = "active"
	IdentityBlocklistEntryStatusDisabled IdentityBlocklistEntryStatus = "disabled"
	IdentityBlocklistEntryStatusRedacted IdentityBlocklistEntryStatus = "redacted"
)

// The type of BlocklistEntry.
type IdentityBlocklistEntryType string

// List of values that IdentityBlocklistEntryType can take
const (
	IdentityBlocklistEntryTypeDocument IdentityBlocklistEntryType = "document"
	IdentityBlocklistEntryTypeSelfie   IdentityBlocklistEntryType = "selfie"
)

// Returns a list of BlocklistEntry objects associated with your account.
//
// The blocklist entries are returned sorted by creation date, with the most recently created
// entries appearing first.
//
// Related guide: [Identity Verification Blocklist](https://docs.stripe.com/docs/identity/review-tools#block-list)
type IdentityBlocklistEntryListParams struct {
	ListParams `form:"*"`
	// Only return BlocklistEntries that were created during the given date interval.
	Created *int64 `form:"created"`
	// Only return BlocklistEntries that were created during the given date interval.
	CreatedRange *RangeQueryParams `form:"created"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Only return blocklist entries with the specified status.
	Status *string `form:"status"`
	// Only return blocklist entries of the specified type.
	Type *string `form:"type"`
	// Only return blocklist entries created from this verification report.
	VerificationReport *string `form:"verification_report"`
}

// AddExpand appends a new field to expand.
func (p *IdentityBlocklistEntryListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Creates a BlocklistEntry object from a verification report.
//
// A blocklist entry prevents future identity verifications that match the same identity information.
// You can create blocklist entries from verification reports that contain document extracted data
// or a selfie.
//
// Related guide: [Identity Verification Blocklist](https://docs.stripe.com/docs/identity/review-tools#add-a-blocklist-entry)
type IdentityBlocklistEntryParams struct {
	Params `form:"*"`
	// When true, the created BlocklistEntry will be used to retroactively unverify matching verifications.
	CheckExistingVerifications *bool `form:"check_existing_verifications"`
	// The type of blocklist entry to be created.
	EntryType *string `form:"entry_type"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The identifier of the VerificationReport to create the BlocklistEntry from.
	VerificationReport *string `form:"verification_report"`
}

// AddExpand appends a new field to expand.
func (p *IdentityBlocklistEntryParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Disables a BlocklistEntry object.
//
// After a BlocklistEntry is disabled, it will no longer block future verifications that match
// the same information. This action is irreversible. To re-enable it, a new BlocklistEntry
// must be created using the same verification report.
//
// Related guide: [Identity Verification Blocklist](https://docs.stripe.com/docs/identity/review-tools#disable-a-blocklist-entry)
type IdentityBlocklistEntryDisableParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *IdentityBlocklistEntryDisableParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Creates a BlocklistEntry object from a verification report.
//
// A blocklist entry prevents future identity verifications that match the same identity information.
// You can create blocklist entries from verification reports that contain document extracted data
// or a selfie.
//
// Related guide: [Identity Verification Blocklist](https://docs.stripe.com/docs/identity/review-tools#add-a-blocklist-entry)
type IdentityBlocklistEntryCreateParams struct {
	Params `form:"*"`
	// When true, the created BlocklistEntry will be used to retroactively unverify matching verifications.
	CheckExistingVerifications *bool `form:"check_existing_verifications"`
	// The type of blocklist entry to be created.
	EntryType *string `form:"entry_type"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The identifier of the VerificationReport to create the BlocklistEntry from.
	VerificationReport *string `form:"verification_report"`
}

// AddExpand appends a new field to expand.
func (p *IdentityBlocklistEntryCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves a BlocklistEntry object by its identifier.
//
// Related guide: [Identity Verification Blocklist](https://docs.stripe.com/docs/identity/review-tools#block-list)
type IdentityBlocklistEntryRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *IdentityBlocklistEntryRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// A BlocklistEntry represents an entry in our identity verification blocklist.
// It helps prevent fraudulent users from repeatedly attempting verification with similar information.
// When you create a BlocklistEntry, we store data from a specified VerificationReport,
// such as document details or facial biometrics.
// This allows us to compare future verification attempts against these entries.
// If a match is found, we categorize the new verification as unverified.
//
// To learn more, see [Identity Verification Blocklist](https://stripe.com/docs/identity/review-tools#block-list)
type IdentityBlocklistEntry struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Time at which you disabled the BlocklistEntry. Measured in seconds since the Unix epoch.
	DisabledAt int64 `json:"disabled_at"`
	// Time at which the BlocklistEntry expires. Measured in seconds since the Unix epoch.
	ExpiresAt int64 `json:"expires_at"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The current status of the BlocklistEntry.
	Status IdentityBlocklistEntryStatus `json:"status"`
	// The type of BlocklistEntry.
	Type IdentityBlocklistEntryType `json:"type"`
	// The verification report the BlocklistEntry was created from.
	VerificationReport *IdentityVerificationReport `json:"verification_report"`
	// The verification session the BlocklistEntry was created from.
	VerificationSession *IdentityVerificationSession `json:"verification_session"`
}

// IdentityBlocklistEntryList is a list of BlocklistEntries as retrieved from a list endpoint.
type IdentityBlocklistEntryList struct {
	APIResource
	ListMeta
	Data []*IdentityBlocklistEntry `json:"data"`
}

// UnmarshalJSON handles deserialization of an IdentityBlocklistEntry.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *IdentityBlocklistEntry) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type identityBlocklistEntry IdentityBlocklistEntry
	var v identityBlocklistEntry
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = IdentityBlocklistEntry(v)
	return nil
}
