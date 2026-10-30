//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Lists FundingIntents.
type V2MoneyManagementFundingIntentListParams struct {
	Params `form:"*"`
	// Filters the results by their creation time.
	Created *RangeQueryParams `form:"created" json:"created,omitempty"`
	// The maximum number of results to return.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Details for a check FundingIntent.
type V2MoneyManagementFundingIntentCheckParams struct {
	// The ID of a File containing the back image of the check.
	BackImage *string `form:"back_image" json:"back_image,omitempty"`
	// The ID of a File containing the front image of the check.
	FrontImage *string `form:"front_image" json:"front_image,omitempty"`
}

// Controls when the funds become available after processing.
type V2MoneyManagementFundingIntentSettlementTimingParams struct {
	// The number of days to delay the availability of funds after processing completes.
	DelayDaysOverride *int64 `form:"delay_days_override" json:"delay_days_override,omitempty"`
}

// The destination for the funds.
type V2MoneyManagementFundingIntentToParams struct {
	// The currency of the destination FinancialAccount balance.
	Currency *string `form:"currency" json:"currency"`
	// The ID of the FinancialAccount that receives the funds.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
}

// Creates a FundingIntent that adds funds to a FinancialAccount.
type V2MoneyManagementFundingIntentParams struct {
	Params `form:"*"`
	// The updated amount to add to the destination.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
	// Updated details for a check FundingIntent.
	Check *V2MoneyManagementFundingIntentCheckParams `form:"check" json:"check,omitempty"`
	// How to confirm the FundingIntent after validation.
	ConfirmationMethod *string `form:"confirmation_method" json:"confirmation_method,omitempty"`
	// An arbitrary string to attach to the FundingIntent.
	Description *string `form:"description" json:"description,omitempty"`
	// Updates to the key-value pairs attached to the FundingIntent.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
	// Controls when the funds become available after processing.
	SettlementTiming *V2MoneyManagementFundingIntentSettlementTimingParams `form:"settlement_timing" json:"settlement_timing,omitempty"`
	// The destination for the funds.
	To *V2MoneyManagementFundingIntentToParams `form:"to" json:"to,omitempty"`
	// The type of FundingIntent to create.
	Type *string `form:"type" json:"type,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementFundingIntentParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}

// Cancels a FundingIntent that can still be canceled.
type V2MoneyManagementFundingIntentCancelParams struct {
	Params `form:"*"`
}

// Confirms a FundingIntent that is ready for processing.
type V2MoneyManagementFundingIntentConfirmParams struct {
	Params `form:"*"`
}

// Details for a check FundingIntent.
type V2MoneyManagementFundingIntentCreateCheckParams struct {
	// The ID of a File containing the back image of the check.
	BackImage *string `form:"back_image" json:"back_image,omitempty"`
	// The ID of a File containing the front image of the check.
	FrontImage *string `form:"front_image" json:"front_image,omitempty"`
}

// Controls when the funds become available after processing.
type V2MoneyManagementFundingIntentCreateSettlementTimingParams struct {
	// The number of days to delay the availability of funds after processing completes.
	DelayDaysOverride *int64 `form:"delay_days_override" json:"delay_days_override,omitempty"`
}

// The destination for the funds.
type V2MoneyManagementFundingIntentCreateToParams struct {
	// The currency of the destination FinancialAccount balance.
	Currency *string `form:"currency" json:"currency"`
	// The ID of the FinancialAccount that receives the funds.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
}

// Creates a FundingIntent that adds funds to a FinancialAccount.
type V2MoneyManagementFundingIntentCreateParams struct {
	Params `form:"*"`
	// The amount to add to the destination.
	Amount *Amount `form:"amount" json:"amount"`
	// Details for a check FundingIntent.
	Check *V2MoneyManagementFundingIntentCreateCheckParams `form:"check" json:"check,omitempty"`
	// How to confirm the FundingIntent after validation.
	ConfirmationMethod *string `form:"confirmation_method" json:"confirmation_method,omitempty"`
	// An arbitrary string to attach to the FundingIntent.
	Description *string `form:"description" json:"description,omitempty"`
	// Set of key-value pairs to attach to the FundingIntent.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Controls when the funds become available after processing.
	SettlementTiming *V2MoneyManagementFundingIntentCreateSettlementTimingParams `form:"settlement_timing" json:"settlement_timing,omitempty"`
	// The destination for the funds.
	To *V2MoneyManagementFundingIntentCreateToParams `form:"to" json:"to"`
	// The type of FundingIntent to create.
	Type *string `form:"type" json:"type"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementFundingIntentCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves a FundingIntent by ID.
type V2MoneyManagementFundingIntentRetrieveParams struct {
	Params `form:"*"`
}

// Updated details for a check FundingIntent.
type V2MoneyManagementFundingIntentUpdateCheckParams struct {
	// The ID of a File containing the back image of the check.
	BackImage *string `form:"back_image" json:"back_image,omitempty"`
	// The ID of a File containing the front image of the check.
	FrontImage *string `form:"front_image" json:"front_image,omitempty"`
}

// Updates the mutable details of a FundingIntent.
type V2MoneyManagementFundingIntentUpdateParams struct {
	Params `form:"*"`
	// The updated amount to add to the destination.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
	// Updated details for a check FundingIntent.
	Check *V2MoneyManagementFundingIntentUpdateCheckParams `form:"check" json:"check,omitempty"`
	// The updated confirmation method.
	ConfirmationMethod *string `form:"confirmation_method" json:"confirmation_method,omitempty"`
	// The updated description.
	Description *string `form:"description" json:"description,omitempty"`
	// Updates to the key-value pairs attached to the FundingIntent.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementFundingIntentUpdateParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}
