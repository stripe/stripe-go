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
}
