//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The endpoint configuration for the batch job.
type V2CoreBatchJobEndpointParams struct {
	// The HTTP method to use when calling the endpoint.
	HTTPMethod *string `form:"http_method" json:"http_method"`
	// The path of the endpoint to run this batch job against.
	// In the form used in the documentation. For instance, for
	// subscription migration this would be `/v1/subscriptions/:id/migrate`.
	Path *V2CoreBatchJobsPath `form:"path" json:"path"`
}

// Notification suppression settings for the batch job.
type V2CoreBatchJobNotificationSuppressionParams struct {
	// The scope of notification suppression.
	Scope *string `form:"scope" json:"scope"`
}

// Creates a new batch job.
type V2CoreBatchJobParams struct {
	Params `form:"*"`
	// The endpoint configuration for the batch job.
	Endpoint *V2CoreBatchJobEndpointParams `form:"endpoint" json:"endpoint,omitempty"`
	// The metadata of the `batch_job`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Notification suppression settings for the batch job.
	NotificationSuppression *V2CoreBatchJobNotificationSuppressionParams `form:"notification_suppression" json:"notification_suppression,omitempty"`
	// Allows the user to skip validation.
	SkipValidation *bool `form:"skip_validation" json:"skip_validation,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreBatchJobParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Cancels an existing batch job.
type V2CoreBatchJobCancelParams struct {
	Params `form:"*"`
}

// The endpoint configuration for the batch job.
type V2CoreBatchJobCreateEndpointParams struct {
	// The HTTP method to use when calling the endpoint.
	HTTPMethod *string `form:"http_method" json:"http_method"`
	// The path of the endpoint to run this batch job against.
	// In the form used in the documentation. For instance, for
	// subscription migration this would be `/v1/subscriptions/:id/migrate`.
	Path *V2CoreBatchJobsPath `form:"path" json:"path"`
}

// Notification suppression settings for the batch job.
type V2CoreBatchJobCreateNotificationSuppressionParams struct {
	// The scope of notification suppression.
	Scope *string `form:"scope" json:"scope"`
}

// Creates a new batch job.
type V2CoreBatchJobCreateParams struct {
	Params `form:"*"`
	// The endpoint configuration for the batch job.
	Endpoint *V2CoreBatchJobCreateEndpointParams `form:"endpoint" json:"endpoint"`
	// The metadata of the `batch_job`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Notification suppression settings for the batch job.
	NotificationSuppression *V2CoreBatchJobCreateNotificationSuppressionParams `form:"notification_suppression" json:"notification_suppression,omitempty"`
	// Allows the user to skip validation.
	SkipValidation *bool `form:"skip_validation" json:"skip_validation"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreBatchJobCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves an existing batch job.
type V2CoreBatchJobRetrieveParams struct {
	Params `form:"*"`
}

// V2CoreBatchJobsPath represents a batchable API endpoint for batch jobs.
type V2CoreBatchJobsPath string

// List of values that V2CoreBatchJobsPath can take.
const (
	V2CoreBatchJobsPathAccountDelete                     V2CoreBatchJobsPath = "/v1/accounts/:account"
	V2CoreBatchJobsPathAccountUpdate                     V2CoreBatchJobsPath = "/v1/accounts/:account"
	V2CoreBatchJobsPathAccountCreate                     V2CoreBatchJobsPath = "/v1/accounts"
	V2CoreBatchJobsPathBankAccountVerify                 V2CoreBatchJobsPath = "/v1/customers/:customer/sources/:id/verify"
	V2CoreBatchJobsPathCouponDelete                      V2CoreBatchJobsPath = "/v1/coupons/:coupon"
	V2CoreBatchJobsPathCouponUpdate                      V2CoreBatchJobsPath = "/v1/coupons/:coupon"
	V2CoreBatchJobsPathCouponCreate                      V2CoreBatchJobsPath = "/v1/coupons"
	V2CoreBatchJobsPathCreditNoteCreate                  V2CoreBatchJobsPath = "/v1/credit_notes"
	V2CoreBatchJobsPathCustomerDelete                    V2CoreBatchJobsPath = "/v1/customers/:customer"
	V2CoreBatchJobsPathCustomerUpdate                    V2CoreBatchJobsPath = "/v1/customers/:customer"
	V2CoreBatchJobsPathCustomerDeleteDiscount            V2CoreBatchJobsPath = "/v1/customers/:customer/discount"
	V2CoreBatchJobsPathCustomerCreate                    V2CoreBatchJobsPath = "/v1/customers"
	V2CoreBatchJobsPathCustomerCreateFundingInstructions V2CoreBatchJobsPath = "/v1/customers/:customer/funding_instructions"
	V2CoreBatchJobsPathCashBalanceUpdate                 V2CoreBatchJobsPath = "/v1/customers/:customer/cash_balance"
	V2CoreBatchJobsPathCustomerBalanceTransactionCreate  V2CoreBatchJobsPath = "/v1/customers/:customer/balance_transactions"
	V2CoreBatchJobsPathCustomerBalanceTransactionUpdate  V2CoreBatchJobsPath = "/v1/customers/:customer/balance_transactions/:transaction"
	V2CoreBatchJobsPathPaymentSourceCreate               V2CoreBatchJobsPath = "/v1/customers/:customer/sources"
	V2CoreBatchJobsPathBankAccountUpdate                 V2CoreBatchJobsPath = "/v1/customers/:customer/sources/:id"
	V2CoreBatchJobsPathBankAccountDelete                 V2CoreBatchJobsPath = "/v1/customers/:customer/sources/:id"
	V2CoreBatchJobsPathTaxIDCreateForCustomer            V2CoreBatchJobsPath = "/v1/customers/:customer/tax_ids"
	V2CoreBatchJobsPathTaxIDDelete                       V2CoreBatchJobsPath = "/v1/customers/:customer/tax_ids/:id"
	V2CoreBatchJobsPathCustomerSessionCreate             V2CoreBatchJobsPath = "/v1/customer_sessions"
	V2CoreBatchJobsPathDisputeClose                      V2CoreBatchJobsPath = "/v1/disputes/:dispute/close"
	V2CoreBatchJobsPathInvoiceDelete                     V2CoreBatchJobsPath = "/v1/invoices/:invoice"
	V2CoreBatchJobsPathInvoiceUpdate                     V2CoreBatchJobsPath = "/v1/invoices/:invoice"
	V2CoreBatchJobsPathInvoiceCreate                     V2CoreBatchJobsPath = "/v1/invoices"
	V2CoreBatchJobsPathInvoiceAddLines                   V2CoreBatchJobsPath = "/v1/invoices/:invoice/add_lines"
	V2CoreBatchJobsPathInvoiceFinalizeInvoice            V2CoreBatchJobsPath = "/v1/invoices/:invoice/finalize"
	V2CoreBatchJobsPathInvoiceMarkUncollectible          V2CoreBatchJobsPath = "/v1/invoices/:invoice/mark_uncollectible"
	V2CoreBatchJobsPathInvoicePay                        V2CoreBatchJobsPath = "/v1/invoices/:invoice/pay"
	V2CoreBatchJobsPathInvoiceRemoveLines                V2CoreBatchJobsPath = "/v1/invoices/:invoice/remove_lines"
	V2CoreBatchJobsPathInvoiceSendInvoice                V2CoreBatchJobsPath = "/v1/invoices/:invoice/send"
	V2CoreBatchJobsPathInvoiceUpdateLines                V2CoreBatchJobsPath = "/v1/invoices/:invoice/update_lines"
	V2CoreBatchJobsPathInvoiceVoidInvoice                V2CoreBatchJobsPath = "/v1/invoices/:invoice/void"
	V2CoreBatchJobsPathInvoiceCreatePreview              V2CoreBatchJobsPath = "/v1/invoices/create_preview"
	V2CoreBatchJobsPathInvoiceRenderingTemplateArchive   V2CoreBatchJobsPath = "/v1/invoice_rendering_templates/:template/archive"
	V2CoreBatchJobsPathInvoiceRenderingTemplateUnarchive V2CoreBatchJobsPath = "/v1/invoice_rendering_templates/:template/unarchive"
	V2CoreBatchJobsPathInvoiceItemDelete                 V2CoreBatchJobsPath = "/v1/invoiceitems/:invoiceitem"
	V2CoreBatchJobsPathInvoiceItemUpdate                 V2CoreBatchJobsPath = "/v1/invoiceitems/:invoiceitem"
	V2CoreBatchJobsPathInvoiceItemCreate                 V2CoreBatchJobsPath = "/v1/invoiceitems"
	V2CoreBatchJobsPathPaymentMethodAttach               V2CoreBatchJobsPath = "/v1/payment_methods/:payment_method/attach"
	V2CoreBatchJobsPathPriceCreate                       V2CoreBatchJobsPath = "/v1/prices"
	V2CoreBatchJobsPathPriceUpdate                       V2CoreBatchJobsPath = "/v1/prices/:price"
	V2CoreBatchJobsPathProductDelete                     V2CoreBatchJobsPath = "/v1/products/:id"
	V2CoreBatchJobsPathProductUpdate                     V2CoreBatchJobsPath = "/v1/products/:id"
	V2CoreBatchJobsPathProductCreate                     V2CoreBatchJobsPath = "/v1/products"
	V2CoreBatchJobsPathProductFeatureDelete              V2CoreBatchJobsPath = "/v1/products/:product/features/:id"
	V2CoreBatchJobsPathProductFeatureCreate              V2CoreBatchJobsPath = "/v1/products/:product/features"
	V2CoreBatchJobsPathPromotionCodeCreate               V2CoreBatchJobsPath = "/v1/promotion_codes"
	V2CoreBatchJobsPathPromotionCodeUpdate               V2CoreBatchJobsPath = "/v1/promotion_codes/:promotion_code"
	V2CoreBatchJobsPathRadarValueListItemCreate          V2CoreBatchJobsPath = "/v1/radar/value_list_items"
	V2CoreBatchJobsPathRefundCreate                      V2CoreBatchJobsPath = "/v1/refunds"
	V2CoreBatchJobsPathRefundCancel                      V2CoreBatchJobsPath = "/v1/refunds/:refund/cancel"
	V2CoreBatchJobsPathSubscriptionCancel                V2CoreBatchJobsPath = "/v1/subscriptions/:subscription_exposed_id"
	V2CoreBatchJobsPathSubscriptionUpdate                V2CoreBatchJobsPath = "/v1/subscriptions/:subscription_exposed_id"
	V2CoreBatchJobsPathSubscriptionCreate                V2CoreBatchJobsPath = "/v1/subscriptions"
	V2CoreBatchJobsPathSubscriptionMigrate               V2CoreBatchJobsPath = "/v1/subscriptions/:subscription/migrate"
	V2CoreBatchJobsPathSubscriptionPause                 V2CoreBatchJobsPath = "/v1/subscriptions/:subscription/pause"
	V2CoreBatchJobsPathSubscriptionResume                V2CoreBatchJobsPath = "/v1/subscriptions/:subscription/resume"
	V2CoreBatchJobsPathSubscriptionItemDelete            V2CoreBatchJobsPath = "/v1/subscription_items/:item"
	V2CoreBatchJobsPathSubscriptionItemUpdate            V2CoreBatchJobsPath = "/v1/subscription_items/:item"
	V2CoreBatchJobsPathSubscriptionItemCreate            V2CoreBatchJobsPath = "/v1/subscription_items"
	V2CoreBatchJobsPathSubscriptionScheduleCreate        V2CoreBatchJobsPath = "/v1/subscription_schedules"
	V2CoreBatchJobsPathSubscriptionScheduleUpdate        V2CoreBatchJobsPath = "/v1/subscription_schedules/:schedule"
	V2CoreBatchJobsPathSubscriptionScheduleCancel        V2CoreBatchJobsPath = "/v1/subscription_schedules/:schedule/cancel"
	V2CoreBatchJobsPathSubscriptionScheduleRelease       V2CoreBatchJobsPath = "/v1/subscription_schedules/:schedule/release"
	V2CoreBatchJobsPathTaxRegistrationCreate             V2CoreBatchJobsPath = "/v1/tax/registrations"
	V2CoreBatchJobsPathTaxRegistrationUpdate             V2CoreBatchJobsPath = "/v1/tax/registrations/:id"
	V2CoreBatchJobsPathTaxSettingsUpdate                 V2CoreBatchJobsPath = "/v1/tax/settings"
	V2CoreBatchJobsPathTaxTransactionCreateReversal      V2CoreBatchJobsPath = "/v1/tax/transactions/create_reversal"
	V2CoreBatchJobsPathTaxRateCreate                     V2CoreBatchJobsPath = "/v1/tax_rates"
	V2CoreBatchJobsPathTaxRateUpdate                     V2CoreBatchJobsPath = "/v1/tax_rates/:tax_rate"
)
