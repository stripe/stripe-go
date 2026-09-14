---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2214
is_stripe_api_change: true
released_in_version: 84.1.0-alpha.1
---

* Add support for new resources `BalanceTransfer` and `RadarAccountEvaluation`
* Add support for `New` method on resource `BalanceTransfer`
* Add support for `Get`, `New`, and `Update` methods on resource `RadarAccountEvaluation`
* Add support for `SpecifiedCommercialTransactionsActURL` on `AccountBusinessProfileParams` and `AccountBusinessProfile`
* Add support for `PaypayPayments` on `AccountSettingsParams` and `AccountSettings`
* Change type of `BillingAnalyticsMeterUsageMeterParams.DimensionFilters` from `string` to `array(string)`
* Change type of `BillingAnalyticsMeterUsageMeterParams.TenantFilters` from `string` to `array(string)`
* Add support for `PaymentMethodConfiguration` on `BillingPortalConfigurationFeaturesPaymentMethodUpdate`
* Add support for `CarRentalData`, `FlightData`, and `LodgingData` on `ChargeCapturePaymentDetailsParams`, `ChargePaymentDetailsParams`, `PaymentIntentCapturePaymentDetailsParams`, `PaymentIntentConfirmPaymentDetailsParams`, and `PaymentIntentPaymentDetailsParams`
* Add support for `TransactionID` on `ChargePaymentMethodDetailsIdeal`, `PaymentAttemptRecordPaymentMethodDetailsIdeal`, and `PaymentRecordPaymentMethodDetailsIdeal`
* Add support for new value `finom` on enums `ConfirmationTokenPaymentMethodPreviewIdeal.Bank`, `PaymentAttemptRecordPaymentMethodDetailsIdeal.Bank`, and `PaymentRecordPaymentMethodDetailsIdeal.Bank`
* Add support for new value `FNOMNL22` on enums `ConfirmationTokenPaymentMethodPreviewIdeal.BIC`, `PaymentAttemptRecordPaymentMethodDetailsIdeal.BIC`, and `PaymentRecordPaymentMethodDetailsIdeal.BIC`
* Add support for new value `tokenized_account_number_deactivated` on enums `ConfirmationTokenPaymentMethodPreviewUsBankAccountStatusDetailsBlocked.Reason` and `PaymentMethodUsBankAccountStatusDetailsBlocked.Reason`
* Add support for `Created` on `CustomerCustomerBalanceTransactionListParams` and `InvoicePaymentListParams`
* Add support for new values `capital.financing_offer.accepted_other_offer`, `financial_connections.account.account_numbers_updated`, and `financial_connections.account.upcoming_account_number_expiry` on enum `Event.Type`
* Add support for `AccountNumbers` on `FinancialConnectionsAccount`
* Add support for `FraudRisk` on `IssuingAuthorizationRiskAssessmentParams`
* Add support for `LatestFraudWarning` on `IssuingCard`
* Add support for `SupplementaryPurchaseData` on `OrderPaymentSettingsPaymentMethodOptionsKlarnaParams`, `PaymentIntentConfirmPaymentMethodOptionsKlarnaParams`, and `PaymentIntentPaymentMethodOptionsKlarnaParams`
* Add support for `CaptureMethod` on `PaymentIntentConfirmPaymentMethodOptionsCardPresentParams`, `PaymentIntentPaymentMethodOptionsCardPresentParams`, and `PaymentIntentPaymentMethodOptionsCardPresent`
* Add support for `AllowRedisplay` and `CustomerAccount` on `PaymentMethodListParams`
* Add support for `MbWay` and `TWINT` on `RefundDestinationDetails`
* Change type of `SubscriptionScheduleParams.BillingSchedules` from `array(billing_schedules_update_params)` to `emptyable(array(billing_schedules_update_params))`
* Add support for snapshot events `EventTypeFinancialConnectionsAccountAccountNumbersUpdated` and `EventTypeFinancialConnectionsAccountUpcomingAccountNumberExpiry` with resource `FinancialConnectionsAccount`
