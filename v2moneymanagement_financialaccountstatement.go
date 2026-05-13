//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The status of the statement. A statement is "active" by default.
// When a statement is replaced by a restatement, its status becomes "restated".
type V2MoneyManagementFinancialAccountStatementStatus string

// List of values that V2MoneyManagementFinancialAccountStatementStatus can take
const (
	V2MoneyManagementFinancialAccountStatementStatusActive   V2MoneyManagementFinancialAccountStatementStatus = "active"
	V2MoneyManagementFinancialAccountStatementStatusRestated V2MoneyManagementFinancialAccountStatementStatus = "restated"
)

// The download URL and expiration.
type V2MoneyManagementFinancialAccountStatementFilesByCurrencyDownloadURL struct {
	// The time at which the URL expires, in ISO 8601 format (UTC).
	ExpiresAt time.Time `json:"expires_at"`
	// The URL to download the file.
	URL Currency `json:"url"`
}

// Currency-specific files and file metadata. Null by default, populated by specifying include=files_by_currency in the Retrieve endpoint.
type V2MoneyManagementFinancialAccountStatementFilesByCurrency struct {
	// The MIME type of the file.
	ContentType Currency `json:"content_type"`
	// The download URL and expiration.
	DownloadURL *V2MoneyManagementFinancialAccountStatementFilesByCurrencyDownloadURL `json:"download_url"`
	// The size of the file in bytes.
	Size int64 `json:"size,string"`
}

// The time period covered by this statement.
type V2MoneyManagementFinancialAccountStatementPeriod struct {
	// The end of the statement period (exclusive), as a UTC-aligned ISO 8601 date
	// (e.g., "2025-02-01"). For example, a January statement has end_date "2025-02-01",
	// meaning all transactions up to but not including February 1st UTC are included.
	EndDate string `json:"end_date"`
	// The start of the statement period (inclusive), as a UTC-aligned ISO 8601 date (e.g., "2025-01-01").
	StartDate string `json:"start_date"`
}

// A Financial Account Statement represents a monthly statement for a Financial Account.
type V2MoneyManagementFinancialAccountStatement struct {
	APIResource
	// Time at which the statement was created, in ISO 8601 format (UTC).
	Created time.Time `json:"created"`
	// Available balance at the end of the statement period.
	EndingBalance map[string]Amount `json:"ending_balance"`
	// Currency-specific files and file metadata. Null by default, populated by specifying include=files_by_currency in the Retrieve endpoint.
	FilesByCurrency map[Currency]*V2MoneyManagementFinancialAccountStatementFilesByCurrency `json:"files_by_currency,omitempty"`
	// The Financial Account this statement belongs to.
	FinancialAccount string `json:"financial_account"`
	// Unique identifier for the statement.
	ID string `json:"id"`
	// True if the object exists in live mode, false if in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The time period covered by this statement.
	Period *V2MoneyManagementFinancialAccountStatementPeriod `json:"period"`
	// The ID of the statement that replaced this one. Only present on statements that have been restated.
	RestatedBy string `json:"restated_by,omitempty"`
	// The ID of the statement this one replaces. Only present on restatements.
	RestatedStatement string `json:"restated_statement,omitempty"`
	// Available balance at the start of the statement period.
	StartingBalance map[string]Amount `json:"starting_balance"`
	// The status of the statement. A statement is "active" by default.
	// When a statement is replaced by a restatement, its status becomes "restated".
	Status V2MoneyManagementFinancialAccountStatementStatus `json:"status"`
}
