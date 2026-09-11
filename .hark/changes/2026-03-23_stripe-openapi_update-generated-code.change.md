---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/2314
is_breaking: true
is_stripe_api_change: true
released_in_version: 85.1.0-beta.1
---

* Add support for new resources `ProductCatalogTrialOffer`, `TaxLocation`, and `V2CoreBatchJob`
* Add support for `New` method on resource `ProductCatalogTrialOffer`
* Add support for `Get`, `List`, and `New` methods on resource `TaxLocation`
* Add support for `Cancel`, `Get`, and `New` methods on resource `V2CoreBatchJob`
* Add support for `PerformanceLocation` on `TaxCalculationLineItemParams` and `TaxCalculationLineItem`
* Add support for new value `performance` on enums `TaxCalculationLineItemTaxBreakdown.Sourcing`, `TaxCalculationShippingCostTaxBreakdown.Sourcing`, and `TaxTransactionShippingCostTaxBreakdown.Sourcing`
* Add support for new values `admissions_tax`, `attendance_tax`, `entertainment_tax`, `gross_receipts_tax`, `hospitality_tax`, `luxury_tax`, `resort_tax`, and `tourism_tax` on enums `TaxCalculationLineItemTaxBreakdownTaxRateDetails.TaxType`, `TaxCalculationShippingCostTaxBreakdownTaxRateDetails.TaxType`, `TaxCalculationTaxBreakdownTaxRateDetails.TaxType`, and `TaxTransactionShippingCostTaxBreakdownTaxRateDetails.TaxType`
* Add support for `TrialOffer` on `InvoiceCreatePreviewScheduleDetailsAmendmentItemActionAddParams`, `InvoiceCreatePreviewScheduleDetailsAmendmentItemActionSetParams`, `InvoiceCreatePreviewScheduleDetailsPhaseItemParams`, `QuoteLineActionAddItemParams`, `QuoteLineActionAddItem`, `QuoteLineActionSetItemParams`, `QuoteLineActionSetItems`, `QuotePreviewSubscriptionSchedulePhaseItem`, `SubscriptionScheduleAmendAmendmentItemActionAddParams`, `SubscriptionScheduleAmendAmendmentItemActionSetParams`, `SubscriptionSchedulePhaseItemParams`, and `SubscriptionSchedulePhaseItem`
* Add support for `RiskReserved` on `Balance`
* ⚠️ Remove support for `SourceType` on `ChargePaymentMethodDetailsStripeBalance`, `ConfirmationTokenPaymentMethodDataStripeBalanceParams`, `ConfirmationTokenPaymentMethodPreviewStripeBalance`, `PaymentAttemptRecordPaymentMethodDetailsStripeBalance`, `PaymentIntentConfirmPaymentMethodDataStripeBalanceParams`, `PaymentIntentPaymentMethodDataStripeBalanceParams`, `PaymentMethodStripeBalanceParams`, `PaymentMethodStripeBalance`, `PaymentRecordPaymentMethodDetailsStripeBalance`, `SetupIntentConfirmPaymentMethodDataStripeBalanceParams`, and `SetupIntentPaymentMethodDataStripeBalanceParams`
* Add support for `TaxDetails` on `CheckoutSessionLineItemPriceDataProductDataParams`, `InvoiceAddLinesLinePriceDataProductDataParams`, `InvoiceLineItemPriceDataProductDataParams`, `InvoiceUpdateLinesLinePriceDataProductDataParams`, `PaymentLinkLineItemPriceDataProductDataParams`, `PlanProductParams`, `PriceProductDataParams`, and `ProductParams`
* Add support for `PendingInvoiceItemInterval` on `CheckoutSessionSubscriptionDataParams`
* Add support for `Hosted` and `UIMode` on `FinancialConnectionsSessionParams` and `FinancialConnectionsSession`
* Add support for `URL` on `FinancialConnectionsSession`
* Add support for `ExpiresAfterSeconds` on `InvoicePaymentSettingsPaymentMethodOptionsPixParams`, `InvoicePaymentSettingsPaymentMethodOptionsPix`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsPix`, `SubscriptionPaymentSettingsPaymentMethodOptionsPixParams`, and `SubscriptionPaymentSettingsPaymentMethodOptionsPix`
* Add support for `CurrentTrial` on `InvoiceCreatePreviewSubscriptionDetailsItemParams`, `SubscriptionItemParams`, and `SubscriptionItem`
* Add support for `Surcharge` on `PaymentIntentAmountDetailsParams`, `PaymentIntentAmountDetails`, `PaymentIntentCaptureAmountDetailsParams`, `PaymentIntentConfirmAmountDetailsParams`, and `PaymentIntentIncrementAuthorizationAmountDetailsParams`
* Add support for `AmountDetails` and `PaymentDetails` on `PaymentIntentDecrementAuthorizationParams`
* Add support for `MandateOptions` on `PaymentIntentPaymentMethodOptionsStripeBalance`
* Add support for `ManagedPayments` on `PaymentLinkParams` and `PaymentLink`
* Add support for `StripeBalance` on `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
* Add support for `BillingCycleAnchor` on `SubscriptionTrialSettingsEndBehaviorParams` and `SubscriptionTrialSettingsEndBehavior`
* Add support for `AdmissionsTax`, `AttendanceTax`, `EntertainmentTax`, `GrossReceiptsTax`, `HospitalityTax`, `LuxuryTax`, `ResortTax`, and `TourismTax` on `TaxRegistrationCountryOptionsUs`
* Add support for new values `admissions_tax`, `attendance_tax`, `entertainment_tax`, `gross_receipts_tax`, `hospitality_tax`, `luxury_tax`, `resort_tax`, and `tourism_tax` on enum `TaxRegistrationCountryOptionsUs.Type`
* Add support for `Requirements` on `TaxCode`
* ⚠️ Change type of `V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCardMandateOptions.Amount`, `V2BillingCollectionSettingPaymentMethodOptionsCardMandateOptions.Amount`, `V2BillingCollectionSettingPaymentMethodOptionsCardMandateOptionsParams.Amount`, and `V2BillingCollectionSettingVersionPaymentMethodOptionsCardMandateOptions.Amount` from `longInteger` to `int64_string`
* Add support for new values `ar_bank_account`, `co_bank_account`, and `eg_bank_account` on enum `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
* Add support for `Timezone` on `V2CoreAccountDefaultsParams` and `V2CoreAccountDefaults`
* Add support for `AzureEventGrid` on `V2CoreEventDestinationParams` and `V2CoreEventDestination`
* Add support for new value `no_azure_partner_topic_exists` on enum `V2CoreEventDestinationStatusDetailsDisabled.Reason`
* Add support for new value `azure_event_grid` on enum `V2CoreEventDestination.Type`
* Add support for `SupportedCurrencies` on `V2CoreVaultGbBankAccount`, `V2CoreVaultUsBankAccount`, and `V2MoneyManagementPayoutMethodCard`
* Add support for `Restricted` on `V2MoneyManagementPayoutMethod`
* Add support for `Currencies` on `V2MoneyManagementPayoutMethodsBankAccountSpecCountriesField`
* Add support for `Counterparty` and `Description` on `V2MoneyManagementTransaction`
* ⚠️ Add support for `Currency` on `V2CoreVaultGbBankAccountParams`, `V2CoreVaultUsBankAccountParams`, `V2MoneyManagementOutboundSetupIntentPayoutMethodDataBankAccountParams`, and `V2MoneyManagementOutboundSetupIntentPayoutMethodDataCardParams`
* Add support for `IBAN` on `V2CoreVaultGbBankAccountParams`
* Add support for new value `currency` on enum `InvalidPaymentMethodError.InvalidParam`
* Add support for event notifications `V2CoreBatchJobBatchFailedEvent`, `V2CoreBatchJobCanceledEvent`, `V2CoreBatchJobCompletedEvent`, `V2CoreBatchJobCreatedEvent`, `V2CoreBatchJobReadyForUploadEvent`, `V2CoreBatchJobTimeoutEvent`, `V2CoreBatchJobUpdatedEvent`, `V2CoreBatchJobUploadTimeoutEvent`, `V2CoreBatchJobValidatingEvent`, and `V2CoreBatchJobValidationFailedEvent` with related object `V2CoreBatchJob`
* Add support for error code `service_period_coupon_with_metered_tiered_item_unsupported` on `QuotePreviewInvoiceLastFinalizationError`
