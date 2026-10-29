//
//
// File generated from our OpenAPI spec
//
//

package stripe

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

// Archive a `PayoutMethod`. Archiving prevents the Payout Method from being used for outbound payments
// or transfers and omits it from normal list results. To restore list visibility, use the
// [unarchive endpoint](https://docs.stripe.com/api/v2/money-management/payout-methods/unarchive).
type V2MoneyManagementPayoutMethodArchiveParams struct {
	Params `form:"*"`
}

// Disable a `PayoutMethod`. Disabling temporarily prevents the Payout Method from being used for outbound
// payments or transfers while keeping it in normal list results. To re-enable it, complete setup again by
// [creating an Outbound Setup Intent](https://docs.stripe.com/api/v2/money-management/outbound-setup-intents/create).
type V2MoneyManagementPayoutMethodDisableParams struct {
	Params `form:"*"`
}

// Unarchive a `PayoutMethod`. Unarchiving restores the Payout Method to normal list results and clears
// only its archived state. It doesn't guarantee that the Payout Method can be used.
type V2MoneyManagementPayoutMethodUnarchiveParams struct {
	Params `form:"*"`
}

// Retrieve a PayoutMethod object.
type V2MoneyManagementPayoutMethodRetrieveParams struct {
	Params `form:"*"`
}
