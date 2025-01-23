//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"time"
)

// Specifies the requirements that Stripe collects from connected accounts in the Connect Onboarding flow.
type AccountLinkCollectionOptionsParams struct {
	// Specifies whether the platform collects only currently_due requirements (`currently_due`) or both currently_due and eventually_due requirements (`eventually_due`). If you don't specify `collection_options`, the default value is `currently_due`.
	Fields *string `form:"fields"`
	// Specifies whether the platform collects future_requirements in addition to requirements in Connect Onboarding. The default value is `omit`.
	FutureRequirements *string `form:"future_requirements"`
}

// AccountLinkType is the type of an account link.
type AccountLinkType string

// List of values that AccountLinkType can take.
const (
	AccountLinkTypeAccountOnboarding AccountLinkType = "account_onboarding"
	AccountLinkTypeAccountUpdate     AccountLinkType = "account_update"
)

// AccountLinkCollect describes what information the platform wants to collect with the account link.
type AccountLinkCollect string

// List of values that AccountLinkCollect can take.
const (
	AccountLinkCollectCurrentlyDue  AccountLinkCollect = "currently_due"
	AccountLinkCollectEventuallyDue AccountLinkCollect = "eventually_due"
)

// Creates an AccountLink object that includes a single-use Stripe URL that the platform can redirect their user to in order to take them through the Connect Onboarding flow.
type AccountLinkParams struct {
	Params `form:"*"`
	// The identifier of the account to create an account link for.
	Account *string `form:"account"`
	// The collect parameter is deprecated. Use `collection_options` instead.
	Collect *string `form:"collect"`
	// Specifies the requirements that Stripe collects from connected accounts in the Connect Onboarding flow.
	CollectionOptions *AccountLinkCollectionOptionsParams `form:"collection_options"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The URL the user will be redirected to if the account link is expired, has been previously-visited, or is otherwise invalid. The URL you specify should attempt to generate a new account link with the same parameters used to create the original account link, then redirect the user to the new account link's URL so they can continue with Connect Onboarding. If a new account link cannot be generated or the redirect fails you should display a useful error to the user.
	RefreshURL *string `form:"refresh_url"`
	// The URL that the user will be redirected to upon leaving or completing the linked flow.
	ReturnURL *string `form:"return_url"`
	// The type of account link the user is requesting. Possible values are `account_onboarding` or `account_update`.
	Type *string `form:"type"`
}

// AddExpand appends a new field to expand.
func (p *AccountLinkParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Account Links are the means by which a Connect platform grants a connected account permission to access
// Stripe-hosted applications, such as Connect Onboarding.
//
// Related guide: [Connect Onboarding](https://stripe.com/docs/connect/custom/hosted-onboarding)
type AccountLink struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created time.Time `json:"created"`
	// The timestamp at which this account link will expire.
	ExpiresAt time.Time `json:"expires_at"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The URL for the account link.
	URL string `json:"url"`
}

// UnmarshalJSON handles deserialization of an AccountLink.
// This custom unmarshaling is needed to handle the time fields correctly.
func (a *AccountLink) UnmarshalJSON(data []byte) error {
	type accountLink AccountLink
	v := struct {
		Created   int64 `json:"created"`
		ExpiresAt int64 `json:"expires_at"`
		*accountLink
	}{
		accountLink: (*accountLink)(a),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	a.Created = time.Unix(v.Created, 0)
	a.ExpiresAt = time.Unix(v.ExpiresAt, 0)
	return nil
}

// MarshalJSON handles serialization of an AccountLink.
// This custom marshaling is needed to handle the time fields correctly.
func (a AccountLink) MarshalJSON() ([]byte, error) {
	type accountLink AccountLink
	v := struct {
		Created   int64 `json:"created"`
		ExpiresAt int64 `json:"expires_at"`
		accountLink
	}{
		accountLink: (accountLink)(a),
		Created:     a.Created.Unix(),
		ExpiresAt:   a.ExpiresAt.Unix(),
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, err
}
