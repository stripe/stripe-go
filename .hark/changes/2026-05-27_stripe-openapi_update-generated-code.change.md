---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/2359
is_breaking: true
is_stripe_api_change: true
released_in_version: 85.2.0
---

* Add support for new resource `V2CommerceProductCatalogImport`
* Add support for `Get` and `New` methods on resource `V2CommerceProductCatalogImport`
* Add support for `BizumPayments` and `ScalapayPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `AutomaticTransferRulesByCurrency` on `BalanceSettingsPaymentsPayoutsParams` and `BalanceSettingsPaymentsPayouts`
* Add support for `StartOfDay` on `BalanceSettingsPaymentsSettlementTimingParams` and `BalanceSettingsPaymentsSettlementTiming`
* Add support for `Description` on `ChargeTransferDataParams`, `PaymentIntentTransferDataParams`, and `PaymentIntentTransferData`
* Add support for `Bizum` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentAttemptRecordPaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `PaymentRecordPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodDataParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
* Add support for `Scalapay` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptionsParams`, `CheckoutSessionPaymentMethodOptions`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentAttemptRecordPaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `PaymentRecordPaymentMethodDetails`, `RefundDestinationDetails`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `Mandate` on `ChargePaymentMethodDetailsTwint`, `PaymentAttemptRecordPaymentMethodDetailsTwint`, and `PaymentRecordPaymentMethodDetailsTwint`
* Change type of `CheckoutSessionPaymentMethodOptionsTwintParams.SetupFutureUsage`, `PaymentIntentConfirmPaymentMethodOptionsTwintParams.SetupFutureUsage`, and `PaymentIntentPaymentMethodOptionsTwintParams.SetupFutureUsage` from `literal('none')` to `enum('none'|'off_session')`
* ⚠️ Change type of `CheckoutSessionPaymentMethodOptionsTwint.SetupFutureUsage` and `PaymentIntentPaymentMethodOptionsTwint.SetupFutureUsage` from `literal('none')` to `enum('none'|'off_session')`
* Add support for new values `bizum` and `scalapay` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
* Add support for `CreditedItems` on `InvoiceItemProrationDetails`
* Add support for `Discountable` on `InvoiceCreatePreviewScheduleDetailsPhaseAddInvoiceItemParams`, `SubscriptionAddInvoiceItemParams`, `SubscriptionSchedulePhaseAddInvoiceItemParams`, and `SubscriptionSchedulePhaseAddInvoiceItem`
* Add support for `BillingSchedules` on `InvoiceCreatePreviewSubscriptionDetailsParams`, `SubscriptionParams`, and `Subscription`
* Add support for `AmountPaidOffStripe` on `Invoice`
* Add support for new value `twint` on enums `InvoicePaymentSettings.PaymentMethodTypes` and `SubscriptionPaymentSettings.PaymentMethodTypes`
* Add support for `TWINT` on `MandatePaymentMethodDetails` and `SetupAttemptPaymentMethodDetails`
* Add support for `Metadata` on `PaymentIntentTransferDataParams`, `PaymentIntentTransferData`, and `SubscriptionPendingUpdate`
* Add support for `PaymentData` on `PaymentIntentTransferDataParams` and `PaymentIntentTransferData`
* Add support for new values `bizum` and `scalapay` on enums `PaymentIntent.ExcludedPaymentMethodTypes` and `SetupIntent.ExcludedPaymentMethodTypes`
* Add support for `BLIKAuthorize` on `PaymentIntentNextAction` and `SetupIntentNextAction`
* Add support for `PaymentMethodOptions` on `PaymentLinkParams` and `PaymentLink`
* Add support for new value `bizum` on enum `PaymentLink.PaymentMethodTypes`
* Add support for `Active` on `PaymentMethodConfigurationListParams`
* Add support for `BilledUntil` on `SubscriptionItem`
* Add support for `Discount` and `Discounts` on `SubscriptionPendingUpdate`
* Add support for `VerifoneM425`, `VerifoneP630`, `VerifoneUx700`, and `VerifoneV660p` on `TerminalConfigurationParams` and `TerminalConfiguration`
* Add support for `APIError` and `PrintContent` on `TerminalReaderAction`
* Add support for new value `print_content` on enum `TerminalReaderAction.Type`
* Add support for new values `simulated_verifone_m425`, `simulated_verifone_p630`, `simulated_verifone_ux700`, `simulated_verifone_v660p`, `verifone_m425`, `verifone_p630`, `verifone_ux700`, and `verifone_v660p` on enum `TerminalReader.DeviceType`
* Add support for `Customer` on `TestHelpersTestClockParams`
* Add support for `Signer` on `V2CoreAccountIdentityBusinessDetailsDocumentsProofOfRegistrationParams`, `V2CoreAccountIdentityBusinessDetailsDocumentsProofOfRegistration`, `V2CoreAccountIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipParams`, `V2CoreAccountIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnership`, `V2CoreAccountTokenIdentityBusinessDetailsDocumentsProofOfRegistrationParams`, and `V2CoreAccountTokenIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipParams`
* Add support for `AzureEventGrid` on `V2CoreEventDestinationParams` and `V2CoreEventDestination`
* Add support for new value `no_azure_partner_topic_exists` on enum `V2CoreEventDestinationStatusDetailsDisabled.Reason`
* Add support for new value `azure_event_grid` on enum `V2CoreEventDestination.Type`
* Add support for new value `meter_event_value_too_many_digits` on enums `EventsV1BillingMeterErrorReportTriggeredEventReasonErrorType.Code` and `EventsV1BillingMeterNoMeterFoundEventReasonErrorType.Code`
* Add support for event notifications `V2CommerceProductCatalogImportsFailedEvent`, `V2CommerceProductCatalogImportsProcessingEvent`, `V2CommerceProductCatalogImportsSucceededEvent`, and `V2CommerceProductCatalogImportsSucceededWithErrorsEvent` with related object `V2CommerceProductCatalogImport`
* Add support for error codes `payment_method_microdeposit_processing_error` and `siret_invalid` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`
