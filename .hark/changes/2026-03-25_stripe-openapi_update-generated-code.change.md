---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2334
is_breaking: true
is_stripe_api_change: true
released_in_version: 85.1.0-alpha.1
---

* Add support for new resource `RiskSignals`
* Add support for `FinancialAccountRewards` and `NestingDemo` on `AccountSessionComponents`
* Add support for `UpiPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `RiskSignals` on `Account`
* Add support for `FraudIntent` on `AccountSignals`
* Add support for new value `related_accounts` on enum `AccountSignalsDelinquencyIndicators.Indicator`
* Add support for `RiskReserved` on `Balance`
* ⚠️ Remove support for `BillableItems` on `BillingAlertSpendThresholdFilters`
* Add support for `Upi` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptionsParams`, `CheckoutSessionPaymentMethodOptions`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `MandatePaymentMethodDetails`, `PaymentAttemptRecordPaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `PaymentRecordPaymentMethodDetails`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodDataParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
* Add support for new value `tempo` on enums `ChargePaymentMethodDetailsCrypto.Network`, `PaymentAttemptRecordPaymentMethodDetailsCrypto.Network`, and `PaymentRecordPaymentMethodDetailsCrypto.Network`
* ⚠️ Remove support for `SourceType` on `ChargePaymentMethodDetailsStripeBalance`, `ConfirmationTokenPaymentMethodDataStripeBalanceParams`, `ConfirmationTokenPaymentMethodPreviewStripeBalance`, `PaymentAttemptRecordPaymentMethodDetailsStripeBalance`, `PaymentIntentConfirmPaymentMethodDataStripeBalanceParams`, `PaymentIntentPaymentMethodDataStripeBalanceParams`, `PaymentMethodStripeBalanceParams`, `PaymentMethodStripeBalance`, `PaymentRecordPaymentMethodDetailsStripeBalance`, `SetupIntentConfirmPaymentMethodDataStripeBalanceParams`, and `SetupIntentPaymentMethodDataStripeBalanceParams`
* Add support for `IntegrationIdentifier` on `CheckoutSessionParams` and `CheckoutSession`
* Change type of `CheckoutSessionLineItemPriceDataProductDataTaxDetailsParams.TaxCode`, `InvoiceAddLinesLinePriceDataProductDataTaxDetailsParams.TaxCode`, `InvoiceLineItemPriceDataProductDataTaxDetailsParams.TaxCode`, `InvoiceUpdateLinesLinePriceDataProductDataTaxDetailsParams.TaxCode`, `PaymentLinkLineItemPriceDataProductDataTaxDetailsParams.TaxCode`, `PlanProductTaxDetailsParams.TaxCode`, `PriceProductDataTaxDetailsParams.TaxCode`, and `ProductTaxDetailsParams.TaxCode` from `string` to `emptyable(string)`
* Add support for `Crypto` on `CheckoutSessionPaymentMethodOptionsParams`
* Add support for `PendingInvoiceItemInterval` on `CheckoutSessionSubscriptionDataParams`
* Add support for new value `application` on enums `CheckoutSessionAutomaticTaxLiability.Type`, `CheckoutSessionInvoiceCreationInvoiceDataIssuer.Type`, `InvoiceAutomaticTaxLiability.Type`, `InvoiceIssuer.Type`, `PaymentLinkAutomaticTaxLiability.Type`, `PaymentLinkInvoiceCreationInvoiceDataIssuer.Type`, `PaymentLinkSubscriptionDataInvoiceSettingsIssuer.Type`, `QuoteAutomaticTaxLiability.Type`, `QuoteInvoiceSettingsIssuer.Type`, `QuotePreviewInvoiceAutomaticTaxLiability.Type`, `QuotePreviewInvoiceIssuer.Type`, `QuotePreviewSubscriptionScheduleDefaultSettingsAutomaticTaxLiability.Type`, `QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettingsIssuer.Type`, `QuotePreviewSubscriptionSchedulePhaseAutomaticTaxLiability.Type`, `QuotePreviewSubscriptionSchedulePhaseInvoiceSettingsIssuer.Type`, `SubscriptionAutomaticTaxLiability.Type`, `SubscriptionInvoiceSettingsIssuer.Type`, `SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuer.Type`, and `SubscriptionSchedulePhaseInvoiceSettingsIssuer.Type`
* Add support for `AUBECSDebit`, `BACSDebit`, `Boleto`, `Link`, `SEPADebit`, and `USBankAccount` on `CheckoutSessionCurrentAttemptPaymentMethodDetails`
* Add support for new values `elements`, `embedded_page`, `form`, and `hosted_page` on enum `CheckoutSession.UIMode`
* ⚠️ Remove support for values `custom`, `embedded`, and `hosted` from enum `CheckoutSession.UIMode`
* Add support for new value `marine_carbon_removal` on enum `ClimateSupplier.RemovalPathway`
* Add support for new value `upi` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
* Add support for `Metadata` on `CreditNoteLineItem`, `CreditNoteLineParams`, `CreditNotePreviewLineParams`, and `CreditNotePreviewLinesLineParams`
* Add support for `SelectedFulfillmentOptionOverrides` on `DelegatedCheckoutRequestedSessionFulfillmentDetails`
* Add support for `LineItemKeys` on `DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptionsDigitalDigitalOptions` and `DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptionsShippingShippingOptions`
* Add support for `QuantityDecimal` on `InvoiceAddLinesLineParams`, `InvoiceCreatePreviewInvoiceItemParams`, `InvoiceItemParams`, `InvoiceItem`, `InvoiceLineItemParams`, `InvoiceLineItem`, and `InvoiceUpdateLinesLineParams`
* Add support for `ExpiresAfterSeconds` on `InvoicePaymentSettingsPaymentMethodOptionsPixParams`, `InvoicePaymentSettingsPaymentMethodOptionsPix`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsPix`, `SubscriptionPaymentSettingsPaymentMethodOptionsPixParams`, and `SubscriptionPaymentSettingsPaymentMethodOptionsPix`
* ⚠️ Add support for `Level` on `IssuingAuthorizationRiskAssessmentCardTestingRiskParams` and `IssuingAuthorizationRiskAssessmentMerchantDisputeRiskParams`
* ⚠️ Remove support for `RiskLevel` on `IssuingAuthorizationRiskAssessmentCardTestingRiskParams` and `IssuingAuthorizationRiskAssessmentMerchantDisputeRiskParams`
* Add support for new values `da`, `pl`, and `sv` on enum `IssuingCardholder.PreferredLocales`
* Add support for `LifecycleControls` on `IssuingCardParams` and `IssuingCard`
* Add support for `Cryptogram`, `ElectronicCommerceIndicator`, `ExemptionIndicatorApplied`, and `ExemptionIndicator` on `PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecure` and `PaymentRecordPaymentMethodDetailsCardThreeDSecure`
* Add support for `Surcharge` on `PaymentIntentAmountDetailsParams`, `PaymentIntentAmountDetails`, `PaymentIntentCaptureAmountDetailsParams`, `PaymentIntentConfirmAmountDetailsParams`, and `PaymentIntentIncrementAuthorizationAmountDetailsParams`
* Add support for `MandateOptions` on `PaymentIntentConfirmPaymentMethodOptionsStripeBalanceParams`, `PaymentIntentPaymentMethodOptionsStripeBalanceParams`, and `PaymentIntentPaymentMethodOptionsStripeBalance`
* Add support for `AmountDetails` and `PaymentDetails` on `PaymentIntentDecrementAuthorizationParams`
* Add support for new value `upi` on enums `PaymentIntent.ExcludedPaymentMethodTypes` and `SetupIntent.ExcludedPaymentMethodTypes`
* Add support for `UpiHandleRedirectOrDisplayQRCode` on `PaymentIntentNextAction` and `SetupIntentNextAction`
* Add support for `ManagedPayments` on `PaymentLinkParams` and `PaymentLink`
* Add support for new value `upi` on enum `PaymentLink.PaymentMethodTypes`
* Add support for `RecommendedAction` and `Signals` on `RadarPaymentEvaluation`
* ⚠️ Remove support for `Insights` on `RadarPaymentEvaluation`
* Add support for new value `crypto_fingerprint` on enum `RadarValueList.ItemType`
* Add support for `StripeBalance` on `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
* Add support for new value `resolved` on enum `SharedPaymentGrantedToken.DeactivatedReason`
* Add support for `RecurringInterval` on `SharedPaymentGrantedTokenUsageLimits`
* Add support for `PresentmentDetails` on `Subscription`
* Add support for new value `canceled_by_retention_policy` on enum `SubscriptionCancellationDetails.Reason`
* ⚠️ Remove support for `InvoiceResources` on `V2BillingIntent`
* ⚠️ Remove support for `AmountDue` and `CustomerBalanceApplied` on `V2BillingIntentAmountDetails`
* Add support for `RecurringCreditGrant` on `V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams`, `V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehavior`, `V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams`, and `V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehavior`
* Add support for `ConsumerPrivacyDisclosures` and `ConsumerStorer` on `V2CoreAccountIdentityAttestationsTermsOfServiceParams` and `V2CoreAccountIdentityAttestationsTermsOfService`
* ⚠️ Remove support for `Include` on `V2BillingIntentParams` and `V2BillingIntentReserveParams`
* Add support for error code `service_period_coupon_with_metered_tiered_item_unsupported` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `QuotePreviewInvoiceLastFinalizationError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`
