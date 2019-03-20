package stripe

// AccountLinkType is the type of an account link.
type AccountLinkType string

// List of values that AccountLinkType can take.
const (
	AccountLinkTypeCustomAccountUpdate       AccountLinkType = "custom_account_update"
	AccountLinkTypeCustomAccountVerification AccountLinkType = "custom_account_verification"
)

// AccountLinkCollect describes what information the platform wants to collect with the account link.
type AccountLinkCollect string

// List of values that AccountLinkCollect can take.
const (
	AccountLinkCollectCurrentlyDue  AccountLinkCollect = "currently_due"
	AccountLinkCollectEventuallyDue AccountLinkCollect = "eventually_due"
)

// AccountLinkParams are the parameters allowed during an account link creation.
type AccountLinkParams struct {
	Params     `form:"*"`
	Account    *string `form:"account"`
	Collect    *string `form:"collect"`
	FailureURL *string `form:"failure_url"`
	SuccessURL *string `form:"success_url"`
	Type       *string `form:"type"`
}

// AccountLink is the resource representing an account link.
// For more details see https://stripe.com/docs/api/#account_links.
type AccountLink struct {
	Created   int64  `json:"created"`
	ExpiresAt int64  `json:"expires_at"`
	Object    string `json:"object"`
	URL       string `json:"url"`
}
