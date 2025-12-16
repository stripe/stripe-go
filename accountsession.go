//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The list of features enabled in the embedded component.
type AccountSessionComponentsAccountManagementFeaturesParams struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication *bool `form:"disable_stripe_user_authentication"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection *bool `form:"external_account_collection"`
}

// Configuration for the [account management](https://docs.stripe.com/connect/supported-embedded-components/account-management/) embedded component.
type AccountSessionComponentsAccountManagementParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsAccountManagementFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionComponentsAccountOnboardingFeaturesParams struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication *bool `form:"disable_stripe_user_authentication"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection *bool `form:"external_account_collection"`
}

// Configuration for the [account onboarding](https://docs.stripe.com/connect/supported-embedded-components/account-onboarding/) embedded component.
type AccountSessionComponentsAccountOnboardingParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsAccountOnboardingFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionComponentsAppInstallFeaturesParams struct {
	// The list of apps allowed to be enabled in the embedded component.
	AllowedApps []*string `form:"allowed_apps"`
}

// Configuration for the [app install](https://docs.stripe.com/connect/supported-embedded-components/app-install/) embedded component.
type AccountSessionComponentsAppInstallParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsAppInstallFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionComponentsAppViewportFeaturesParams struct {
	// The list of apps allowed to be enabled in the embedded component.
	AllowedApps []*string `form:"allowed_apps"`
}

// Configuration for the [app viewport](https://docs.stripe.com/connect/supported-embedded-components/app-viewport/) embedded component.
type AccountSessionComponentsAppViewportParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsAppViewportFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionComponentsBalancesFeaturesParams struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication *bool `form:"disable_stripe_user_authentication"`
	// Whether to allow payout schedule to be changed. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	EditPayoutSchedule *bool `form:"edit_payout_schedule"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection *bool `form:"external_account_collection"`
	// Whether to allow creation of instant payouts. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	InstantPayouts *bool `form:"instant_payouts"`
	// Whether to allow creation of standard payouts. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	StandardPayouts *bool `form:"standard_payouts"`
}

// Configuration for the [balances](https://docs.stripe.com/connect/supported-embedded-components/balances/) embedded component.
type AccountSessionComponentsBalancesParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsBalancesFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionComponentsCapitalFinancingFeaturesParams struct{}

// Configuration for the [Capital financing](https://docs.stripe.com/connect/supported-embedded-components/capital-financing/) embedded component.
type AccountSessionComponentsCapitalFinancingParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionComponentsCapitalFinancingFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionComponentsCapitalFinancingApplicationFeaturesParams struct{}

// Configuration for the [Capital financing application](https://docs.stripe.com/connect/supported-embedded-components/capital-financing-application/) embedded component.
type AccountSessionComponentsCapitalFinancingApplicationParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionComponentsCapitalFinancingApplicationFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionComponentsCapitalFinancingPromotionFeaturesParams struct{}

// Configuration for the [Capital financing promotion](https://docs.stripe.com/connect/supported-embedded-components/capital-financing-promotion/) embedded component.
type AccountSessionComponentsCapitalFinancingPromotionParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionComponentsCapitalFinancingPromotionFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionComponentsCapitalOverviewFeaturesParams struct{}

// Configuration for the [Capital overview](https://docs.stripe.com/connect/supported-embedded-components/capital-overview/) embedded component.
type AccountSessionComponentsCapitalOverviewParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionComponentsCapitalOverviewFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionComponentsDisputesListFeaturesParams struct {
	// Whether to allow capturing and cancelling payment intents. This is `true` by default.
	CapturePayments *bool `form:"capture_payments"`
	// Whether connected accounts can manage destination charges that are created on behalf of them. This is `false` by default.
	DestinationOnBehalfOfChargeManagement *bool `form:"destination_on_behalf_of_charge_management"`
	// Whether responding to disputes is enabled, including submitting evidence and accepting disputes. This is `true` by default.
	DisputeManagement *bool `form:"dispute_management"`
	// Whether sending refunds is enabled. This is `true` by default.
	RefundManagement *bool `form:"refund_management"`
}

// Configuration for the [disputes list](https://docs.stripe.com/connect/supported-embedded-components/disputes-list/) embedded component.
type AccountSessionComponentsDisputesListParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsDisputesListFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionComponentsDocumentsFeaturesParams struct{}

// Configuration for the [documents](https://docs.stripe.com/connect/supported-embedded-components/documents/) embedded component.
type AccountSessionComponentsDocumentsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionComponentsDocumentsFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionComponentsExportTaxTransactionsFeaturesParams struct{}

// Configuration for the [export tax transactions](https://docs.stripe.com/connect/supported-embedded-components/export-tax-transactions/) embedded component.
type AccountSessionComponentsExportTaxTransactionsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionComponentsExportTaxTransactionsFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionComponentsFinancialAccountFeaturesParams struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication *bool `form:"disable_stripe_user_authentication"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection *bool `form:"external_account_collection"`
	// Whether to allow sending money.
	SendMoney *bool `form:"send_money"`
	// Whether to allow transferring balance.
	TransferBalance *bool `form:"transfer_balance"`
}

// Configuration for the [financial account](https://docs.stripe.com/connect/supported-embedded-components/financial-account/) embedded component.
type AccountSessionComponentsFinancialAccountParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsFinancialAccountFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionComponentsFinancialAccountTransactionsFeaturesParams struct {
	// Whether to allow card spend dispute management features.
	CardSpendDisputeManagement *bool `form:"card_spend_dispute_management"`
}

// Configuration for the [financial account transactions](https://docs.stripe.com/connect/supported-embedded-components/financial-account-transactions/) embedded component.
type AccountSessionComponentsFinancialAccountTransactionsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsFinancialAccountTransactionsFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionComponentsInstantPayoutsPromotionFeaturesParams struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication *bool `form:"disable_stripe_user_authentication"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection *bool `form:"external_account_collection"`
	// Whether to allow creation of instant payouts. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	InstantPayouts *bool `form:"instant_payouts"`
}

// Configuration for the [instant payouts promotion](https://docs.stripe.com/connect/supported-embedded-components/instant-payouts-promotion/) embedded component.
type AccountSessionComponentsInstantPayoutsPromotionParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsInstantPayoutsPromotionFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionComponentsIssuingCardFeaturesParams struct {
	// Whether to allow cardholder management features.
	CardholderManagement *bool `form:"cardholder_management"`
	// Whether to allow card management features.
	CardManagement *bool `form:"card_management"`
	// Whether to allow card spend dispute management features.
	CardSpendDisputeManagement *bool `form:"card_spend_dispute_management"`
	// Whether to allow spend control management features.
	SpendControlManagement *bool `form:"spend_control_management"`
}

// Configuration for the [issuing card](https://docs.stripe.com/connect/supported-embedded-components/issuing-card/) embedded component.
type AccountSessionComponentsIssuingCardParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsIssuingCardFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionComponentsIssuingCardsListFeaturesParams struct {
	// Whether to allow cardholder management features.
	CardholderManagement *bool `form:"cardholder_management"`
	// Whether to allow card management features.
	CardManagement *bool `form:"card_management"`
	// Whether to allow card spend dispute management features.
	CardSpendDisputeManagement *bool `form:"card_spend_dispute_management"`
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication *bool `form:"disable_stripe_user_authentication"`
	// Whether to allow spend control management features.
	SpendControlManagement *bool `form:"spend_control_management"`
}

// Configuration for the [issuing cards list](https://docs.stripe.com/connect/supported-embedded-components/issuing-cards-list/) embedded component.
type AccountSessionComponentsIssuingCardsListParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsIssuingCardsListFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionComponentsNotificationBannerFeaturesParams struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication *bool `form:"disable_stripe_user_authentication"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection *bool `form:"external_account_collection"`
}

// Configuration for the [notification banner](https://docs.stripe.com/connect/supported-embedded-components/notification-banner/) embedded component.
type AccountSessionComponentsNotificationBannerParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsNotificationBannerFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionComponentsPaymentDetailsFeaturesParams struct {
	// Whether to allow capturing and cancelling payment intents. This is `true` by default.
	CapturePayments *bool `form:"capture_payments"`
	// Whether connected accounts can manage destination charges that are created on behalf of them. This is `false` by default.
	DestinationOnBehalfOfChargeManagement *bool `form:"destination_on_behalf_of_charge_management"`
	// Whether responding to disputes is enabled, including submitting evidence and accepting disputes. This is `true` by default.
	DisputeManagement *bool `form:"dispute_management"`
	// Whether sending refunds is enabled. This is `true` by default.
	RefundManagement *bool `form:"refund_management"`
}

// Configuration for the [payment details](https://docs.stripe.com/connect/supported-embedded-components/payment-details/) embedded component.
type AccountSessionComponentsPaymentDetailsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsPaymentDetailsFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionComponentsPaymentDisputesFeaturesParams struct {
	// Whether connected accounts can manage destination charges that are created on behalf of them. This is `false` by default.
	DestinationOnBehalfOfChargeManagement *bool `form:"destination_on_behalf_of_charge_management"`
	// Whether responding to disputes is enabled, including submitting evidence and accepting disputes. This is `true` by default.
	DisputeManagement *bool `form:"dispute_management"`
	// Whether sending refunds is enabled. This is `true` by default.
	RefundManagement *bool `form:"refund_management"`
}

// Configuration for the [payment disputes](https://docs.stripe.com/connect/supported-embedded-components/payment-disputes/) embedded component.
type AccountSessionComponentsPaymentDisputesParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsPaymentDisputesFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionComponentsPaymentMethodSettingsFeaturesParams struct{}

// Configuration for the [payment method settings](https://docs.stripe.com/connect/supported-embedded-components/payment-method-settings/) embedded component.
type AccountSessionComponentsPaymentMethodSettingsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionComponentsPaymentMethodSettingsFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionComponentsPaymentsFeaturesParams struct {
	// Whether to allow capturing and cancelling payment intents. This is `true` by default.
	CapturePayments *bool `form:"capture_payments"`
	// Whether connected accounts can manage destination charges that are created on behalf of them. This is `false` by default.
	DestinationOnBehalfOfChargeManagement *bool `form:"destination_on_behalf_of_charge_management"`
	// Whether responding to disputes is enabled, including submitting evidence and accepting disputes. This is `true` by default.
	DisputeManagement *bool `form:"dispute_management"`
	// Whether sending refunds is enabled. This is `true` by default.
	RefundManagement *bool `form:"refund_management"`
}

// Configuration for the [payments](https://docs.stripe.com/connect/supported-embedded-components/payments/) embedded component.
type AccountSessionComponentsPaymentsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsPaymentsFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionComponentsPayoutDetailsFeaturesParams struct{}

// Configuration for the [payout details](https://docs.stripe.com/connect/supported-embedded-components/payout-details/) embedded component.
type AccountSessionComponentsPayoutDetailsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionComponentsPayoutDetailsFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionComponentsPayoutsFeaturesParams struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication *bool `form:"disable_stripe_user_authentication"`
	// Whether to allow payout schedule to be changed. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	EditPayoutSchedule *bool `form:"edit_payout_schedule"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection *bool `form:"external_account_collection"`
	// Whether to allow creation of instant payouts. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	InstantPayouts *bool `form:"instant_payouts"`
	// Whether to allow creation of standard payouts. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	StandardPayouts *bool `form:"standard_payouts"`
}

// Configuration for the [payouts](https://docs.stripe.com/connect/supported-embedded-components/payouts/) embedded component.
type AccountSessionComponentsPayoutsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsPayoutsFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionComponentsPayoutsListFeaturesParams struct{}

// Configuration for the [payouts list](https://docs.stripe.com/connect/supported-embedded-components/payouts-list/) embedded component.
type AccountSessionComponentsPayoutsListParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionComponentsPayoutsListFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionComponentsProductTaxCodeSelectorFeaturesParams struct{}

// Configuration for the [product tax code selector](https://docs.stripe.com/connect/supported-embedded-components/product-tax-code-selector/) embedded component.
type AccountSessionComponentsProductTaxCodeSelectorParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionComponentsProductTaxCodeSelectorFeaturesParams `form:"features"`
}
type AccountSessionComponentsRecipientsFeaturesParams struct {
	// Whether to allow sending money.
	SendMoney *bool `form:"send_money"`
}

// Configuration for the [recipients](https://docs.stripe.com/connect/supported-embedded-components/recipients/) embedded component.
type AccountSessionComponentsRecipientsParams struct {
	// Whether the embedded component is enabled.
	Enabled  *bool                                             `form:"enabled"`
	Features *AccountSessionComponentsRecipientsFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionComponentsReportingChartFeaturesParams struct{}

// Configuration for the [reporting chart](https://docs.stripe.com/connect/supported-embedded-components/reporting-chart/) embedded component.
type AccountSessionComponentsReportingChartParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionComponentsReportingChartFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionComponentsTaxRegistrationsFeaturesParams struct{}

// Configuration for the [tax registrations](https://docs.stripe.com/connect/supported-embedded-components/tax-registrations/) embedded component.
type AccountSessionComponentsTaxRegistrationsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionComponentsTaxRegistrationsFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionComponentsTaxSettingsFeaturesParams struct{}

// Configuration for the [tax settings](https://docs.stripe.com/connect/supported-embedded-components/tax-settings/) embedded component.
type AccountSessionComponentsTaxSettingsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionComponentsTaxSettingsFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionComponentsTaxThresholdMonitoringFeaturesParams struct{}

// Configuration for the [tax threshold monitoring](https://docs.stripe.com/connect/supported-embedded-components/tax-threshold-monitoring/) embedded component.
type AccountSessionComponentsTaxThresholdMonitoringParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionComponentsTaxThresholdMonitoringFeaturesParams `form:"features"`
}

// Each key of the dictionary represents an embedded component, and each embedded component maps to its configuration (e.g. whether it has been enabled or not).
type AccountSessionComponentsParams struct {
	// Configuration for the [account management](https://docs.stripe.com/connect/supported-embedded-components/account-management/) embedded component.
	AccountManagement *AccountSessionComponentsAccountManagementParams `form:"account_management"`
	// Configuration for the [account onboarding](https://docs.stripe.com/connect/supported-embedded-components/account-onboarding/) embedded component.
	AccountOnboarding *AccountSessionComponentsAccountOnboardingParams `form:"account_onboarding"`
	// Configuration for the [app install](https://docs.stripe.com/connect/supported-embedded-components/app-install/) embedded component.
	AppInstall *AccountSessionComponentsAppInstallParams `form:"app_install"`
	// Configuration for the [app viewport](https://docs.stripe.com/connect/supported-embedded-components/app-viewport/) embedded component.
	AppViewport *AccountSessionComponentsAppViewportParams `form:"app_viewport"`
	// Configuration for the [balances](https://docs.stripe.com/connect/supported-embedded-components/balances/) embedded component.
	Balances *AccountSessionComponentsBalancesParams `form:"balances"`
	// Configuration for the [Capital financing](https://docs.stripe.com/connect/supported-embedded-components/capital-financing/) embedded component.
	CapitalFinancing *AccountSessionComponentsCapitalFinancingParams `form:"capital_financing"`
	// Configuration for the [Capital financing application](https://docs.stripe.com/connect/supported-embedded-components/capital-financing-application/) embedded component.
	CapitalFinancingApplication *AccountSessionComponentsCapitalFinancingApplicationParams `form:"capital_financing_application"`
	// Configuration for the [Capital financing promotion](https://docs.stripe.com/connect/supported-embedded-components/capital-financing-promotion/) embedded component.
	CapitalFinancingPromotion *AccountSessionComponentsCapitalFinancingPromotionParams `form:"capital_financing_promotion"`
	// Configuration for the [Capital overview](https://docs.stripe.com/connect/supported-embedded-components/capital-overview/) embedded component.
	CapitalOverview *AccountSessionComponentsCapitalOverviewParams `form:"capital_overview"`
	// Configuration for the [disputes list](https://docs.stripe.com/connect/supported-embedded-components/disputes-list/) embedded component.
	DisputesList *AccountSessionComponentsDisputesListParams `form:"disputes_list"`
	// Configuration for the [documents](https://docs.stripe.com/connect/supported-embedded-components/documents/) embedded component.
	Documents *AccountSessionComponentsDocumentsParams `form:"documents"`
	// Configuration for the [export tax transactions](https://docs.stripe.com/connect/supported-embedded-components/export-tax-transactions/) embedded component.
	ExportTaxTransactions *AccountSessionComponentsExportTaxTransactionsParams `form:"export_tax_transactions"`
	// Configuration for the [financial account](https://docs.stripe.com/connect/supported-embedded-components/financial-account/) embedded component.
	FinancialAccount *AccountSessionComponentsFinancialAccountParams `form:"financial_account"`
	// Configuration for the [financial account transactions](https://docs.stripe.com/connect/supported-embedded-components/financial-account-transactions/) embedded component.
	FinancialAccountTransactions *AccountSessionComponentsFinancialAccountTransactionsParams `form:"financial_account_transactions"`
	// Configuration for the [instant payouts promotion](https://docs.stripe.com/connect/supported-embedded-components/instant-payouts-promotion/) embedded component.
	InstantPayoutsPromotion *AccountSessionComponentsInstantPayoutsPromotionParams `form:"instant_payouts_promotion"`
	// Configuration for the [issuing card](https://docs.stripe.com/connect/supported-embedded-components/issuing-card/) embedded component.
	IssuingCard *AccountSessionComponentsIssuingCardParams `form:"issuing_card"`
	// Configuration for the [issuing cards list](https://docs.stripe.com/connect/supported-embedded-components/issuing-cards-list/) embedded component.
	IssuingCardsList *AccountSessionComponentsIssuingCardsListParams `form:"issuing_cards_list"`
	// Configuration for the [notification banner](https://docs.stripe.com/connect/supported-embedded-components/notification-banner/) embedded component.
	NotificationBanner *AccountSessionComponentsNotificationBannerParams `form:"notification_banner"`
	// Configuration for the [payment details](https://docs.stripe.com/connect/supported-embedded-components/payment-details/) embedded component.
	PaymentDetails *AccountSessionComponentsPaymentDetailsParams `form:"payment_details"`
	// Configuration for the [payment disputes](https://docs.stripe.com/connect/supported-embedded-components/payment-disputes/) embedded component.
	PaymentDisputes *AccountSessionComponentsPaymentDisputesParams `form:"payment_disputes"`
	// Configuration for the [payment method settings](https://docs.stripe.com/connect/supported-embedded-components/payment-method-settings/) embedded component.
	PaymentMethodSettings *AccountSessionComponentsPaymentMethodSettingsParams `form:"payment_method_settings"`
	// Configuration for the [payments](https://docs.stripe.com/connect/supported-embedded-components/payments/) embedded component.
	Payments *AccountSessionComponentsPaymentsParams `form:"payments"`
	// Configuration for the [payout details](https://docs.stripe.com/connect/supported-embedded-components/payout-details/) embedded component.
	PayoutDetails *AccountSessionComponentsPayoutDetailsParams `form:"payout_details"`
	// Configuration for the [payouts](https://docs.stripe.com/connect/supported-embedded-components/payouts/) embedded component.
	Payouts *AccountSessionComponentsPayoutsParams `form:"payouts"`
	// Configuration for the [payouts list](https://docs.stripe.com/connect/supported-embedded-components/payouts-list/) embedded component.
	PayoutsList *AccountSessionComponentsPayoutsListParams `form:"payouts_list"`
	// Configuration for the [product tax code selector](https://docs.stripe.com/connect/supported-embedded-components/product-tax-code-selector/) embedded component.
	ProductTaxCodeSelector *AccountSessionComponentsProductTaxCodeSelectorParams `form:"product_tax_code_selector"`
	// Configuration for the [recipients](https://docs.stripe.com/connect/supported-embedded-components/recipients/) embedded component.
	Recipients *AccountSessionComponentsRecipientsParams `form:"recipients"`
	// Configuration for the [reporting chart](https://docs.stripe.com/connect/supported-embedded-components/reporting-chart/) embedded component.
	ReportingChart *AccountSessionComponentsReportingChartParams `form:"reporting_chart"`
	// Configuration for the [tax registrations](https://docs.stripe.com/connect/supported-embedded-components/tax-registrations/) embedded component.
	TaxRegistrations *AccountSessionComponentsTaxRegistrationsParams `form:"tax_registrations"`
	// Configuration for the [tax settings](https://docs.stripe.com/connect/supported-embedded-components/tax-settings/) embedded component.
	TaxSettings *AccountSessionComponentsTaxSettingsParams `form:"tax_settings"`
	// Configuration for the [tax threshold monitoring](https://docs.stripe.com/connect/supported-embedded-components/tax-threshold-monitoring/) embedded component.
	TaxThresholdMonitoring *AccountSessionComponentsTaxThresholdMonitoringParams `form:"tax_threshold_monitoring"`
}

// Creates a AccountSession object that includes a single-use token that the platform can use on their front-end to grant client-side API access.
type AccountSessionParams struct {
	Params `form:"*"`
	// The identifier of the account to create an Account Session for.
	Account *string `form:"account"`
	// Each key of the dictionary represents an embedded component, and each embedded component maps to its configuration (e.g. whether it has been enabled or not).
	Components *AccountSessionComponentsParams `form:"components"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *AccountSessionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The list of features enabled in the embedded component.
type AccountSessionCreateComponentsAccountManagementFeaturesParams struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication *bool `form:"disable_stripe_user_authentication"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection *bool `form:"external_account_collection"`
}

// Configuration for the [account management](https://docs.stripe.com/connect/supported-embedded-components/account-management/) embedded component.
type AccountSessionCreateComponentsAccountManagementParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionCreateComponentsAccountManagementFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionCreateComponentsAccountOnboardingFeaturesParams struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication *bool `form:"disable_stripe_user_authentication"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection *bool `form:"external_account_collection"`
}

// Configuration for the [account onboarding](https://docs.stripe.com/connect/supported-embedded-components/account-onboarding/) embedded component.
type AccountSessionCreateComponentsAccountOnboardingParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionCreateComponentsAccountOnboardingFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionCreateComponentsAppInstallFeaturesParams struct {
	// The list of apps allowed to be enabled in the embedded component.
	AllowedApps []*string `form:"allowed_apps"`
}

// Configuration for the [app install](https://docs.stripe.com/connect/supported-embedded-components/app-install/) embedded component.
type AccountSessionCreateComponentsAppInstallParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionCreateComponentsAppInstallFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionCreateComponentsAppViewportFeaturesParams struct {
	// The list of apps allowed to be enabled in the embedded component.
	AllowedApps []*string `form:"allowed_apps"`
}

// Configuration for the [app viewport](https://docs.stripe.com/connect/supported-embedded-components/app-viewport/) embedded component.
type AccountSessionCreateComponentsAppViewportParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionCreateComponentsAppViewportFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionCreateComponentsBalancesFeaturesParams struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication *bool `form:"disable_stripe_user_authentication"`
	// Whether to allow payout schedule to be changed. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	EditPayoutSchedule *bool `form:"edit_payout_schedule"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection *bool `form:"external_account_collection"`
	// Whether to allow creation of instant payouts. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	InstantPayouts *bool `form:"instant_payouts"`
	// Whether to allow creation of standard payouts. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	StandardPayouts *bool `form:"standard_payouts"`
}

// Configuration for the [balances](https://docs.stripe.com/connect/supported-embedded-components/balances/) embedded component.
type AccountSessionCreateComponentsBalancesParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionCreateComponentsBalancesFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionCreateComponentsCapitalFinancingFeaturesParams struct{}

// Configuration for the [Capital financing](https://docs.stripe.com/connect/supported-embedded-components/capital-financing/) embedded component.
type AccountSessionCreateComponentsCapitalFinancingParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionCreateComponentsCapitalFinancingFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionCreateComponentsCapitalFinancingApplicationFeaturesParams struct{}

// Configuration for the [Capital financing application](https://docs.stripe.com/connect/supported-embedded-components/capital-financing-application/) embedded component.
type AccountSessionCreateComponentsCapitalFinancingApplicationParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionCreateComponentsCapitalFinancingApplicationFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionCreateComponentsCapitalFinancingPromotionFeaturesParams struct{}

// Configuration for the [Capital financing promotion](https://docs.stripe.com/connect/supported-embedded-components/capital-financing-promotion/) embedded component.
type AccountSessionCreateComponentsCapitalFinancingPromotionParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionCreateComponentsCapitalFinancingPromotionFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionCreateComponentsCapitalOverviewFeaturesParams struct{}

// Configuration for the [Capital overview](https://docs.stripe.com/connect/supported-embedded-components/capital-overview/) embedded component.
type AccountSessionCreateComponentsCapitalOverviewParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionCreateComponentsCapitalOverviewFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionCreateComponentsDisputesListFeaturesParams struct {
	// Whether to allow capturing and cancelling payment intents. This is `true` by default.
	CapturePayments *bool `form:"capture_payments"`
	// Whether connected accounts can manage destination charges that are created on behalf of them. This is `false` by default.
	DestinationOnBehalfOfChargeManagement *bool `form:"destination_on_behalf_of_charge_management"`
	// Whether responding to disputes is enabled, including submitting evidence and accepting disputes. This is `true` by default.
	DisputeManagement *bool `form:"dispute_management"`
	// Whether sending refunds is enabled. This is `true` by default.
	RefundManagement *bool `form:"refund_management"`
}

// Configuration for the [disputes list](https://docs.stripe.com/connect/supported-embedded-components/disputes-list/) embedded component.
type AccountSessionCreateComponentsDisputesListParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionCreateComponentsDisputesListFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionCreateComponentsDocumentsFeaturesParams struct{}

// Configuration for the [documents](https://docs.stripe.com/connect/supported-embedded-components/documents/) embedded component.
type AccountSessionCreateComponentsDocumentsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionCreateComponentsDocumentsFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionCreateComponentsExportTaxTransactionsFeaturesParams struct{}

// Configuration for the [export tax transactions](https://docs.stripe.com/connect/supported-embedded-components/export-tax-transactions/) embedded component.
type AccountSessionCreateComponentsExportTaxTransactionsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionCreateComponentsExportTaxTransactionsFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionCreateComponentsFinancialAccountFeaturesParams struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication *bool `form:"disable_stripe_user_authentication"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection *bool `form:"external_account_collection"`
	// Whether to allow sending money.
	SendMoney *bool `form:"send_money"`
	// Whether to allow transferring balance.
	TransferBalance *bool `form:"transfer_balance"`
}

// Configuration for the [financial account](https://docs.stripe.com/connect/supported-embedded-components/financial-account/) embedded component.
type AccountSessionCreateComponentsFinancialAccountParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionCreateComponentsFinancialAccountFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionCreateComponentsFinancialAccountTransactionsFeaturesParams struct {
	// Whether to allow card spend dispute management features.
	CardSpendDisputeManagement *bool `form:"card_spend_dispute_management"`
}

// Configuration for the [financial account transactions](https://docs.stripe.com/connect/supported-embedded-components/financial-account-transactions/) embedded component.
type AccountSessionCreateComponentsFinancialAccountTransactionsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionCreateComponentsFinancialAccountTransactionsFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionCreateComponentsInstantPayoutsPromotionFeaturesParams struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication *bool `form:"disable_stripe_user_authentication"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection *bool `form:"external_account_collection"`
	// Whether to allow creation of instant payouts. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	InstantPayouts *bool `form:"instant_payouts"`
}

// Configuration for the [instant payouts promotion](https://docs.stripe.com/connect/supported-embedded-components/instant-payouts-promotion/) embedded component.
type AccountSessionCreateComponentsInstantPayoutsPromotionParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionCreateComponentsInstantPayoutsPromotionFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionCreateComponentsIssuingCardFeaturesParams struct {
	// Whether to allow cardholder management features.
	CardholderManagement *bool `form:"cardholder_management"`
	// Whether to allow card management features.
	CardManagement *bool `form:"card_management"`
	// Whether to allow card spend dispute management features.
	CardSpendDisputeManagement *bool `form:"card_spend_dispute_management"`
	// Whether to allow spend control management features.
	SpendControlManagement *bool `form:"spend_control_management"`
}

// Configuration for the [issuing card](https://docs.stripe.com/connect/supported-embedded-components/issuing-card/) embedded component.
type AccountSessionCreateComponentsIssuingCardParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionCreateComponentsIssuingCardFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionCreateComponentsIssuingCardsListFeaturesParams struct {
	// Whether to allow cardholder management features.
	CardholderManagement *bool `form:"cardholder_management"`
	// Whether to allow card management features.
	CardManagement *bool `form:"card_management"`
	// Whether to allow card spend dispute management features.
	CardSpendDisputeManagement *bool `form:"card_spend_dispute_management"`
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication *bool `form:"disable_stripe_user_authentication"`
	// Whether to allow spend control management features.
	SpendControlManagement *bool `form:"spend_control_management"`
}

// Configuration for the [issuing cards list](https://docs.stripe.com/connect/supported-embedded-components/issuing-cards-list/) embedded component.
type AccountSessionCreateComponentsIssuingCardsListParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionCreateComponentsIssuingCardsListFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionCreateComponentsNotificationBannerFeaturesParams struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication *bool `form:"disable_stripe_user_authentication"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection *bool `form:"external_account_collection"`
}

// Configuration for the [notification banner](https://docs.stripe.com/connect/supported-embedded-components/notification-banner/) embedded component.
type AccountSessionCreateComponentsNotificationBannerParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionCreateComponentsNotificationBannerFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionCreateComponentsPaymentDetailsFeaturesParams struct {
	// Whether to allow capturing and cancelling payment intents. This is `true` by default.
	CapturePayments *bool `form:"capture_payments"`
	// Whether connected accounts can manage destination charges that are created on behalf of them. This is `false` by default.
	DestinationOnBehalfOfChargeManagement *bool `form:"destination_on_behalf_of_charge_management"`
	// Whether responding to disputes is enabled, including submitting evidence and accepting disputes. This is `true` by default.
	DisputeManagement *bool `form:"dispute_management"`
	// Whether sending refunds is enabled. This is `true` by default.
	RefundManagement *bool `form:"refund_management"`
}

// Configuration for the [payment details](https://docs.stripe.com/connect/supported-embedded-components/payment-details/) embedded component.
type AccountSessionCreateComponentsPaymentDetailsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionCreateComponentsPaymentDetailsFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionCreateComponentsPaymentDisputesFeaturesParams struct {
	// Whether connected accounts can manage destination charges that are created on behalf of them. This is `false` by default.
	DestinationOnBehalfOfChargeManagement *bool `form:"destination_on_behalf_of_charge_management"`
	// Whether responding to disputes is enabled, including submitting evidence and accepting disputes. This is `true` by default.
	DisputeManagement *bool `form:"dispute_management"`
	// Whether sending refunds is enabled. This is `true` by default.
	RefundManagement *bool `form:"refund_management"`
}

// Configuration for the [payment disputes](https://docs.stripe.com/connect/supported-embedded-components/payment-disputes/) embedded component.
type AccountSessionCreateComponentsPaymentDisputesParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionCreateComponentsPaymentDisputesFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionCreateComponentsPaymentMethodSettingsFeaturesParams struct{}

// Configuration for the [payment method settings](https://docs.stripe.com/connect/supported-embedded-components/payment-method-settings/) embedded component.
type AccountSessionCreateComponentsPaymentMethodSettingsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionCreateComponentsPaymentMethodSettingsFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionCreateComponentsPaymentsFeaturesParams struct {
	// Whether to allow capturing and cancelling payment intents. This is `true` by default.
	CapturePayments *bool `form:"capture_payments"`
	// Whether connected accounts can manage destination charges that are created on behalf of them. This is `false` by default.
	DestinationOnBehalfOfChargeManagement *bool `form:"destination_on_behalf_of_charge_management"`
	// Whether responding to disputes is enabled, including submitting evidence and accepting disputes. This is `true` by default.
	DisputeManagement *bool `form:"dispute_management"`
	// Whether sending refunds is enabled. This is `true` by default.
	RefundManagement *bool `form:"refund_management"`
}

// Configuration for the [payments](https://docs.stripe.com/connect/supported-embedded-components/payments/) embedded component.
type AccountSessionCreateComponentsPaymentsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionCreateComponentsPaymentsFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionCreateComponentsPayoutDetailsFeaturesParams struct{}

// Configuration for the [payout details](https://docs.stripe.com/connect/supported-embedded-components/payout-details/) embedded component.
type AccountSessionCreateComponentsPayoutDetailsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionCreateComponentsPayoutDetailsFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionCreateComponentsPayoutsFeaturesParams struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication *bool `form:"disable_stripe_user_authentication"`
	// Whether to allow payout schedule to be changed. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	EditPayoutSchedule *bool `form:"edit_payout_schedule"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection *bool `form:"external_account_collection"`
	// Whether to allow creation of instant payouts. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	InstantPayouts *bool `form:"instant_payouts"`
	// Whether to allow creation of standard payouts. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	StandardPayouts *bool `form:"standard_payouts"`
}

// Configuration for the [payouts](https://docs.stripe.com/connect/supported-embedded-components/payouts/) embedded component.
type AccountSessionCreateComponentsPayoutsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionCreateComponentsPayoutsFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionCreateComponentsPayoutsListFeaturesParams struct{}

// Configuration for the [payouts list](https://docs.stripe.com/connect/supported-embedded-components/payouts-list/) embedded component.
type AccountSessionCreateComponentsPayoutsListParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionCreateComponentsPayoutsListFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionCreateComponentsProductTaxCodeSelectorFeaturesParams struct{}

// Configuration for the [product tax code selector](https://docs.stripe.com/connect/supported-embedded-components/product-tax-code-selector/) embedded component.
type AccountSessionCreateComponentsProductTaxCodeSelectorParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionCreateComponentsProductTaxCodeSelectorFeaturesParams `form:"features"`
}
type AccountSessionCreateComponentsRecipientsFeaturesParams struct {
	// Whether to allow sending money.
	SendMoney *bool `form:"send_money"`
}

// Configuration for the [recipients](https://docs.stripe.com/connect/supported-embedded-components/recipients/) embedded component.
type AccountSessionCreateComponentsRecipientsParams struct {
	// Whether the embedded component is enabled.
	Enabled  *bool                                                   `form:"enabled"`
	Features *AccountSessionCreateComponentsRecipientsFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionCreateComponentsReportingChartFeaturesParams struct{}

// Configuration for the [reporting chart](https://docs.stripe.com/connect/supported-embedded-components/reporting-chart/) embedded component.
type AccountSessionCreateComponentsReportingChartParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionCreateComponentsReportingChartFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionCreateComponentsTaxRegistrationsFeaturesParams struct{}

// Configuration for the [tax registrations](https://docs.stripe.com/connect/supported-embedded-components/tax-registrations/) embedded component.
type AccountSessionCreateComponentsTaxRegistrationsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionCreateComponentsTaxRegistrationsFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionCreateComponentsTaxSettingsFeaturesParams struct{}

// Configuration for the [tax settings](https://docs.stripe.com/connect/supported-embedded-components/tax-settings/) embedded component.
type AccountSessionCreateComponentsTaxSettingsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionCreateComponentsTaxSettingsFeaturesParams `form:"features"`
}

// An empty list, because this embedded component has no features.
type AccountSessionCreateComponentsTaxThresholdMonitoringFeaturesParams struct{}

// Configuration for the [tax threshold monitoring](https://docs.stripe.com/connect/supported-embedded-components/tax-threshold-monitoring/) embedded component.
type AccountSessionCreateComponentsTaxThresholdMonitoringParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// An empty list, because this embedded component has no features.
	Features *AccountSessionCreateComponentsTaxThresholdMonitoringFeaturesParams `form:"features"`
}

// Each key of the dictionary represents an embedded component, and each embedded component maps to its configuration (e.g. whether it has been enabled or not).
type AccountSessionCreateComponentsParams struct {
	// Configuration for the [account management](https://docs.stripe.com/connect/supported-embedded-components/account-management/) embedded component.
	AccountManagement *AccountSessionCreateComponentsAccountManagementParams `form:"account_management"`
	// Configuration for the [account onboarding](https://docs.stripe.com/connect/supported-embedded-components/account-onboarding/) embedded component.
	AccountOnboarding *AccountSessionCreateComponentsAccountOnboardingParams `form:"account_onboarding"`
	// Configuration for the [app install](https://docs.stripe.com/connect/supported-embedded-components/app-install/) embedded component.
	AppInstall *AccountSessionCreateComponentsAppInstallParams `form:"app_install"`
	// Configuration for the [app viewport](https://docs.stripe.com/connect/supported-embedded-components/app-viewport/) embedded component.
	AppViewport *AccountSessionCreateComponentsAppViewportParams `form:"app_viewport"`
	// Configuration for the [balances](https://docs.stripe.com/connect/supported-embedded-components/balances/) embedded component.
	Balances *AccountSessionCreateComponentsBalancesParams `form:"balances"`
	// Configuration for the [Capital financing](https://docs.stripe.com/connect/supported-embedded-components/capital-financing/) embedded component.
	CapitalFinancing *AccountSessionCreateComponentsCapitalFinancingParams `form:"capital_financing"`
	// Configuration for the [Capital financing application](https://docs.stripe.com/connect/supported-embedded-components/capital-financing-application/) embedded component.
	CapitalFinancingApplication *AccountSessionCreateComponentsCapitalFinancingApplicationParams `form:"capital_financing_application"`
	// Configuration for the [Capital financing promotion](https://docs.stripe.com/connect/supported-embedded-components/capital-financing-promotion/) embedded component.
	CapitalFinancingPromotion *AccountSessionCreateComponentsCapitalFinancingPromotionParams `form:"capital_financing_promotion"`
	// Configuration for the [Capital overview](https://docs.stripe.com/connect/supported-embedded-components/capital-overview/) embedded component.
	CapitalOverview *AccountSessionCreateComponentsCapitalOverviewParams `form:"capital_overview"`
	// Configuration for the [disputes list](https://docs.stripe.com/connect/supported-embedded-components/disputes-list/) embedded component.
	DisputesList *AccountSessionCreateComponentsDisputesListParams `form:"disputes_list"`
	// Configuration for the [documents](https://docs.stripe.com/connect/supported-embedded-components/documents/) embedded component.
	Documents *AccountSessionCreateComponentsDocumentsParams `form:"documents"`
	// Configuration for the [export tax transactions](https://docs.stripe.com/connect/supported-embedded-components/export-tax-transactions/) embedded component.
	ExportTaxTransactions *AccountSessionCreateComponentsExportTaxTransactionsParams `form:"export_tax_transactions"`
	// Configuration for the [financial account](https://docs.stripe.com/connect/supported-embedded-components/financial-account/) embedded component.
	FinancialAccount *AccountSessionCreateComponentsFinancialAccountParams `form:"financial_account"`
	// Configuration for the [financial account transactions](https://docs.stripe.com/connect/supported-embedded-components/financial-account-transactions/) embedded component.
	FinancialAccountTransactions *AccountSessionCreateComponentsFinancialAccountTransactionsParams `form:"financial_account_transactions"`
	// Configuration for the [instant payouts promotion](https://docs.stripe.com/connect/supported-embedded-components/instant-payouts-promotion/) embedded component.
	InstantPayoutsPromotion *AccountSessionCreateComponentsInstantPayoutsPromotionParams `form:"instant_payouts_promotion"`
	// Configuration for the [issuing card](https://docs.stripe.com/connect/supported-embedded-components/issuing-card/) embedded component.
	IssuingCard *AccountSessionCreateComponentsIssuingCardParams `form:"issuing_card"`
	// Configuration for the [issuing cards list](https://docs.stripe.com/connect/supported-embedded-components/issuing-cards-list/) embedded component.
	IssuingCardsList *AccountSessionCreateComponentsIssuingCardsListParams `form:"issuing_cards_list"`
	// Configuration for the [notification banner](https://docs.stripe.com/connect/supported-embedded-components/notification-banner/) embedded component.
	NotificationBanner *AccountSessionCreateComponentsNotificationBannerParams `form:"notification_banner"`
	// Configuration for the [payment details](https://docs.stripe.com/connect/supported-embedded-components/payment-details/) embedded component.
	PaymentDetails *AccountSessionCreateComponentsPaymentDetailsParams `form:"payment_details"`
	// Configuration for the [payment disputes](https://docs.stripe.com/connect/supported-embedded-components/payment-disputes/) embedded component.
	PaymentDisputes *AccountSessionCreateComponentsPaymentDisputesParams `form:"payment_disputes"`
	// Configuration for the [payment method settings](https://docs.stripe.com/connect/supported-embedded-components/payment-method-settings/) embedded component.
	PaymentMethodSettings *AccountSessionCreateComponentsPaymentMethodSettingsParams `form:"payment_method_settings"`
	// Configuration for the [payments](https://docs.stripe.com/connect/supported-embedded-components/payments/) embedded component.
	Payments *AccountSessionCreateComponentsPaymentsParams `form:"payments"`
	// Configuration for the [payout details](https://docs.stripe.com/connect/supported-embedded-components/payout-details/) embedded component.
	PayoutDetails *AccountSessionCreateComponentsPayoutDetailsParams `form:"payout_details"`
	// Configuration for the [payouts](https://docs.stripe.com/connect/supported-embedded-components/payouts/) embedded component.
	Payouts *AccountSessionCreateComponentsPayoutsParams `form:"payouts"`
	// Configuration for the [payouts list](https://docs.stripe.com/connect/supported-embedded-components/payouts-list/) embedded component.
	PayoutsList *AccountSessionCreateComponentsPayoutsListParams `form:"payouts_list"`
	// Configuration for the [product tax code selector](https://docs.stripe.com/connect/supported-embedded-components/product-tax-code-selector/) embedded component.
	ProductTaxCodeSelector *AccountSessionCreateComponentsProductTaxCodeSelectorParams `form:"product_tax_code_selector"`
	// Configuration for the [recipients](https://docs.stripe.com/connect/supported-embedded-components/recipients/) embedded component.
	Recipients *AccountSessionCreateComponentsRecipientsParams `form:"recipients"`
	// Configuration for the [reporting chart](https://docs.stripe.com/connect/supported-embedded-components/reporting-chart/) embedded component.
	ReportingChart *AccountSessionCreateComponentsReportingChartParams `form:"reporting_chart"`
	// Configuration for the [tax registrations](https://docs.stripe.com/connect/supported-embedded-components/tax-registrations/) embedded component.
	TaxRegistrations *AccountSessionCreateComponentsTaxRegistrationsParams `form:"tax_registrations"`
	// Configuration for the [tax settings](https://docs.stripe.com/connect/supported-embedded-components/tax-settings/) embedded component.
	TaxSettings *AccountSessionCreateComponentsTaxSettingsParams `form:"tax_settings"`
	// Configuration for the [tax threshold monitoring](https://docs.stripe.com/connect/supported-embedded-components/tax-threshold-monitoring/) embedded component.
	TaxThresholdMonitoring *AccountSessionCreateComponentsTaxThresholdMonitoringParams `form:"tax_threshold_monitoring"`
}

// Creates a AccountSession object that includes a single-use token that the platform can use on their front-end to grant client-side API access.
type AccountSessionCreateParams struct {
	Params `form:"*"`
	// The identifier of the account to create an Account Session for.
	Account *string `form:"account"`
	// Each key of the dictionary represents an embedded component, and each embedded component maps to its configuration (e.g. whether it has been enabled or not).
	Components *AccountSessionCreateComponentsParams `form:"components"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *AccountSessionCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type AccountSessionComponentsAccountManagementFeatures struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication bool `json:"disable_stripe_user_authentication"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection bool `json:"external_account_collection"`
}
type AccountSessionComponentsAccountManagement struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                               `json:"enabled"`
	Features *AccountSessionComponentsAccountManagementFeatures `json:"features"`
}
type AccountSessionComponentsAccountOnboardingFeatures struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication bool `json:"disable_stripe_user_authentication"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection bool `json:"external_account_collection"`
}
type AccountSessionComponentsAccountOnboarding struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                               `json:"enabled"`
	Features *AccountSessionComponentsAccountOnboardingFeatures `json:"features"`
}
type AccountSessionComponentsBalancesFeatures struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication bool `json:"disable_stripe_user_authentication"`
	// Whether to allow payout schedule to be changed. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	EditPayoutSchedule bool `json:"edit_payout_schedule"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection bool `json:"external_account_collection"`
	// Whether to allow creation of instant payouts. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	InstantPayouts bool `json:"instant_payouts"`
	// Whether to allow creation of standard payouts. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	StandardPayouts bool `json:"standard_payouts"`
}
type AccountSessionComponentsBalances struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                      `json:"enabled"`
	Features *AccountSessionComponentsBalancesFeatures `json:"features"`
}
type AccountSessionComponentsCapitalFinancingFeatures struct{}
type AccountSessionComponentsCapitalFinancing struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                              `json:"enabled"`
	Features *AccountSessionComponentsCapitalFinancingFeatures `json:"features"`
}
type AccountSessionComponentsCapitalFinancingApplicationFeatures struct{}
type AccountSessionComponentsCapitalFinancingApplication struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                                         `json:"enabled"`
	Features *AccountSessionComponentsCapitalFinancingApplicationFeatures `json:"features"`
}
type AccountSessionComponentsCapitalFinancingPromotionFeatures struct{}
type AccountSessionComponentsCapitalFinancingPromotion struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                                       `json:"enabled"`
	Features *AccountSessionComponentsCapitalFinancingPromotionFeatures `json:"features"`
}
type AccountSessionComponentsDisputesListFeatures struct {
	// Whether to allow capturing and cancelling payment intents. This is `true` by default.
	CapturePayments bool `json:"capture_payments"`
	// Whether connected accounts can manage destination charges that are created on behalf of them. This is `false` by default.
	DestinationOnBehalfOfChargeManagement bool `json:"destination_on_behalf_of_charge_management"`
	// Whether responding to disputes is enabled, including submitting evidence and accepting disputes. This is `true` by default.
	DisputeManagement bool `json:"dispute_management"`
	// Whether sending refunds is enabled. This is `true` by default.
	RefundManagement bool `json:"refund_management"`
}
type AccountSessionComponentsDisputesList struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                          `json:"enabled"`
	Features *AccountSessionComponentsDisputesListFeatures `json:"features"`
}
type AccountSessionComponentsDocumentsFeatures struct{}
type AccountSessionComponentsDocuments struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                       `json:"enabled"`
	Features *AccountSessionComponentsDocumentsFeatures `json:"features"`
}
type AccountSessionComponentsFinancialAccountFeatures struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication bool `json:"disable_stripe_user_authentication"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection bool `json:"external_account_collection"`
	// Whether to allow sending money.
	SendMoney bool `json:"send_money"`
	// Whether to allow transferring balance.
	TransferBalance bool `json:"transfer_balance"`
}
type AccountSessionComponentsFinancialAccount struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                              `json:"enabled"`
	Features *AccountSessionComponentsFinancialAccountFeatures `json:"features"`
}
type AccountSessionComponentsFinancialAccountTransactionsFeatures struct {
	// Whether to allow card spend dispute management features.
	CardSpendDisputeManagement bool `json:"card_spend_dispute_management"`
}
type AccountSessionComponentsFinancialAccountTransactions struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                                          `json:"enabled"`
	Features *AccountSessionComponentsFinancialAccountTransactionsFeatures `json:"features"`
}
type AccountSessionComponentsInstantPayoutsPromotionFeatures struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication bool `json:"disable_stripe_user_authentication"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection bool `json:"external_account_collection"`
	// Whether to allow creation of instant payouts. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	InstantPayouts bool `json:"instant_payouts"`
}
type AccountSessionComponentsInstantPayoutsPromotion struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                                     `json:"enabled"`
	Features *AccountSessionComponentsInstantPayoutsPromotionFeatures `json:"features"`
}
type AccountSessionComponentsIssuingCardFeatures struct {
	// Whether to allow cardholder management features.
	CardholderManagement bool `json:"cardholder_management"`
	// Whether to allow card management features.
	CardManagement bool `json:"card_management"`
	// Whether to allow card spend dispute management features.
	CardSpendDisputeManagement bool `json:"card_spend_dispute_management"`
	// Whether to allow spend control management features.
	SpendControlManagement bool `json:"spend_control_management"`
}
type AccountSessionComponentsIssuingCard struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                         `json:"enabled"`
	Features *AccountSessionComponentsIssuingCardFeatures `json:"features"`
}
type AccountSessionComponentsIssuingCardsListFeatures struct {
	// Whether to allow cardholder management features.
	CardholderManagement bool `json:"cardholder_management"`
	// Whether to allow card management features.
	CardManagement bool `json:"card_management"`
	// Whether to allow card spend dispute management features.
	CardSpendDisputeManagement bool `json:"card_spend_dispute_management"`
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication bool `json:"disable_stripe_user_authentication"`
	// Whether to allow spend control management features.
	SpendControlManagement bool `json:"spend_control_management"`
}
type AccountSessionComponentsIssuingCardsList struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                              `json:"enabled"`
	Features *AccountSessionComponentsIssuingCardsListFeatures `json:"features"`
}
type AccountSessionComponentsNotificationBannerFeatures struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication bool `json:"disable_stripe_user_authentication"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection bool `json:"external_account_collection"`
}
type AccountSessionComponentsNotificationBanner struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                                `json:"enabled"`
	Features *AccountSessionComponentsNotificationBannerFeatures `json:"features"`
}
type AccountSessionComponentsPaymentDetailsFeatures struct {
	// Whether to allow capturing and cancelling payment intents. This is `true` by default.
	CapturePayments bool `json:"capture_payments"`
	// Whether connected accounts can manage destination charges that are created on behalf of them. This is `false` by default.
	DestinationOnBehalfOfChargeManagement bool `json:"destination_on_behalf_of_charge_management"`
	// Whether responding to disputes is enabled, including submitting evidence and accepting disputes. This is `true` by default.
	DisputeManagement bool `json:"dispute_management"`
	// Whether sending refunds is enabled. This is `true` by default.
	RefundManagement bool `json:"refund_management"`
}
type AccountSessionComponentsPaymentDetails struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                            `json:"enabled"`
	Features *AccountSessionComponentsPaymentDetailsFeatures `json:"features"`
}
type AccountSessionComponentsPaymentDisputesFeatures struct {
	// Whether connected accounts can manage destination charges that are created on behalf of them. This is `false` by default.
	DestinationOnBehalfOfChargeManagement bool `json:"destination_on_behalf_of_charge_management"`
	// Whether responding to disputes is enabled, including submitting evidence and accepting disputes. This is `true` by default.
	DisputeManagement bool `json:"dispute_management"`
	// Whether sending refunds is enabled. This is `true` by default.
	RefundManagement bool `json:"refund_management"`
}
type AccountSessionComponentsPaymentDisputes struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                             `json:"enabled"`
	Features *AccountSessionComponentsPaymentDisputesFeatures `json:"features"`
}
type AccountSessionComponentsPaymentsFeatures struct {
	// Whether to allow capturing and cancelling payment intents. This is `true` by default.
	CapturePayments bool `json:"capture_payments"`
	// Whether connected accounts can manage destination charges that are created on behalf of them. This is `false` by default.
	DestinationOnBehalfOfChargeManagement bool `json:"destination_on_behalf_of_charge_management"`
	// Whether responding to disputes is enabled, including submitting evidence and accepting disputes. This is `true` by default.
	DisputeManagement bool `json:"dispute_management"`
	// Whether sending refunds is enabled. This is `true` by default.
	RefundManagement bool `json:"refund_management"`
}
type AccountSessionComponentsPayments struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                      `json:"enabled"`
	Features *AccountSessionComponentsPaymentsFeatures `json:"features"`
}
type AccountSessionComponentsPayoutDetailsFeatures struct{}
type AccountSessionComponentsPayoutDetails struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                           `json:"enabled"`
	Features *AccountSessionComponentsPayoutDetailsFeatures `json:"features"`
}
type AccountSessionComponentsPayoutsFeatures struct {
	// Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`.
	DisableStripeUserAuthentication bool `json:"disable_stripe_user_authentication"`
	// Whether to allow payout schedule to be changed. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	EditPayoutSchedule bool `json:"edit_payout_schedule"`
	// Whether external account collection is enabled. This feature can only be `false` for accounts where you're responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`.
	ExternalAccountCollection bool `json:"external_account_collection"`
	// Whether to allow creation of instant payouts. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	InstantPayouts bool `json:"instant_payouts"`
	// Whether to allow creation of standard payouts. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`.
	StandardPayouts bool `json:"standard_payouts"`
}
type AccountSessionComponentsPayouts struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                     `json:"enabled"`
	Features *AccountSessionComponentsPayoutsFeatures `json:"features"`
}
type AccountSessionComponentsPayoutsListFeatures struct{}
type AccountSessionComponentsPayoutsList struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                         `json:"enabled"`
	Features *AccountSessionComponentsPayoutsListFeatures `json:"features"`
}
type AccountSessionComponentsTaxRegistrationsFeatures struct{}
type AccountSessionComponentsTaxRegistrations struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                              `json:"enabled"`
	Features *AccountSessionComponentsTaxRegistrationsFeatures `json:"features"`
}
type AccountSessionComponentsTaxSettingsFeatures struct{}
type AccountSessionComponentsTaxSettings struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                         `json:"enabled"`
	Features *AccountSessionComponentsTaxSettingsFeatures `json:"features"`
}
type AccountSessionComponents struct {
	AccountManagement            *AccountSessionComponentsAccountManagement            `json:"account_management"`
	AccountOnboarding            *AccountSessionComponentsAccountOnboarding            `json:"account_onboarding"`
	Balances                     *AccountSessionComponentsBalances                     `json:"balances"`
	CapitalFinancing             *AccountSessionComponentsCapitalFinancing             `json:"capital_financing"`
	CapitalFinancingApplication  *AccountSessionComponentsCapitalFinancingApplication  `json:"capital_financing_application"`
	CapitalFinancingPromotion    *AccountSessionComponentsCapitalFinancingPromotion    `json:"capital_financing_promotion"`
	DisputesList                 *AccountSessionComponentsDisputesList                 `json:"disputes_list"`
	Documents                    *AccountSessionComponentsDocuments                    `json:"documents"`
	FinancialAccount             *AccountSessionComponentsFinancialAccount             `json:"financial_account"`
	FinancialAccountTransactions *AccountSessionComponentsFinancialAccountTransactions `json:"financial_account_transactions"`
	InstantPayoutsPromotion      *AccountSessionComponentsInstantPayoutsPromotion      `json:"instant_payouts_promotion"`
	IssuingCard                  *AccountSessionComponentsIssuingCard                  `json:"issuing_card"`
	IssuingCardsList             *AccountSessionComponentsIssuingCardsList             `json:"issuing_cards_list"`
	NotificationBanner           *AccountSessionComponentsNotificationBanner           `json:"notification_banner"`
	PaymentDetails               *AccountSessionComponentsPaymentDetails               `json:"payment_details"`
	PaymentDisputes              *AccountSessionComponentsPaymentDisputes              `json:"payment_disputes"`
	Payments                     *AccountSessionComponentsPayments                     `json:"payments"`
	PayoutDetails                *AccountSessionComponentsPayoutDetails                `json:"payout_details"`
	Payouts                      *AccountSessionComponentsPayouts                      `json:"payouts"`
	PayoutsList                  *AccountSessionComponentsPayoutsList                  `json:"payouts_list"`
	TaxRegistrations             *AccountSessionComponentsTaxRegistrations             `json:"tax_registrations"`
	TaxSettings                  *AccountSessionComponentsTaxSettings                  `json:"tax_settings"`
}

// An AccountSession allows a Connect platform to grant access to a connected account in Connect embedded components.
//
// We recommend that you create an AccountSession each time you need to display an embedded component
// to your user. Do not save AccountSessions to your database as they expire relatively
// quickly, and cannot be used more than once.
//
// Related guide: [Connect embedded components](https://docs.stripe.com/connect/get-started-connect-embedded-components)
type AccountSession struct {
	APIResource
	// The ID of the account the AccountSession was created for
	Account string `json:"account"`
	// The client secret of this AccountSession. Used on the client to set up secure access to the given `account`.
	//
	// The client secret can be used to provide access to `account` from your frontend. It should not be stored, logged, or exposed to anyone other than the connected account. Make sure that you have TLS enabled on any page that includes the client secret.
	//
	// Refer to our docs to [setup Connect embedded components](https://docs.stripe.com/connect/get-started-connect-embedded-components) and learn about how `client_secret` should be handled.
	ClientSecret string                    `json:"client_secret"`
	Components   *AccountSessionComponents `json:"components"`
	// The timestamp at which this AccountSession will expire.
	ExpiresAt int64 `json:"expires_at"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
