//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Archive a PayoutMethod object. Archived objects cannot be used as payout methods
// and will not appear in the payout method list.
type V2MoneyManagementPayoutMethodArchiveParams struct {
	Params `form:"*"`
}

// Usage status filter.
type V2MoneyManagementPayoutMethodListUsageStatusParams struct {
	// List of payments status to filter by.
	Payments []*string `form:"payments" json:"payments,omitempty"`
	// List of transfers status to filter by.
	Transfers []*string `form:"transfers" json:"transfers,omitempty"`
}

// List objects that adhere to the PayoutMethod interface.
type V2MoneyManagementPayoutMethodListParams struct {
	Params `form:"*"`
	// The page size.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Usage status filter.
	UsageStatus *V2MoneyManagementPayoutMethodListUsageStatusParams `form:"usage_status" json:"usage_status,omitempty"`
}

// Retrieve a PayoutMethod object.
type V2MoneyManagementPayoutMethodParams struct {
	Params `form:"*"`
}

// Unarchive an PayoutMethod object.
type V2MoneyManagementPayoutMethodUnarchiveParams struct {
	Params `form:"*"`
}
