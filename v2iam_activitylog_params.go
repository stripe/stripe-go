//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List activity logs of an account.
type V2IamActivityLogListParams struct {
	Params `form:"*"`
	// Filter results to only include activity logs for the specified action group types.
	ActionGroups []*string `form:"action_groups" json:"action_groups,omitempty"`
	// Filter results to only include activity logs for the specified action types.
	Actions []*string `form:"actions" json:"actions,omitempty"`
	// Maximum number of results to return per page.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Controls which activity logs an organization-level request returns. Specify `@self` to return logs recorded in the organization context and `@organization_members` to return logs recorded in accounts that belong to the organization. Specify both values to return logs from both sources. Defaults to `["@self"]` when omitted. Cannot be empty.
	LogsFrom []*string `form:"logs_from" json:"logs_from,omitempty"`
}

// Retrieve an activity log.
type V2IamActivityLogParams struct {
	Params `form:"*"`
}

// Retrieve an activity log.
type V2IamActivityLogRetrieveParams struct {
	Params `form:"*"`
}
