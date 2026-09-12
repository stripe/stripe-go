---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2407
is_breaking: true
is_stripe_api_change: true
released_in_version: 86.4.0-alpha.2
---

* Add support for new resources `BillingFeedbackOption` and `PaymentPlan`
* ⚠️ Remove support for resource `BillingFeedbackOptions`
* Add support for `Get`, `List`, `New`, and `Update` methods on resource `PaymentPlan`
* Add support for `Update` method on resource `V2MoneyManagementTransaction`
* Add support for `WeChatPayPayments` on `AccountSettingsParams` and `AccountSettings`
* ⚠️ Change type of `BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReason.FeedbackOptions` from `$Billing.FeedbackOptions` to `$Billing.FeedbackOption`
* Add support for `SubscriptionPause` on `BillingPortalSessionFlow`
* Add support for new value `subscription_pause` on enum `BillingPortalSessionFlow.Type`
* Add support for new value `usdt` on enum `CryptoOnrampSessionTransactionDetails.DestinationCurrencies`
* Add support for new value `usdt` on enum `CryptoOnrampSessionTransactionDetails.DestinationCurrency`
* Add support for `ActiveEntitlements` on `CustomerSessionComponents`
* Add support for `SharedPaymentIssuedToken` on `DelegatedCheckoutRequestedSessionConfirmParams`
* Add support for new values `payment_plan.created`, `payment_plan.installment_due`, `payment_plan.installment_paid`, `payment_plan.installment_will_be_due`, and `payment_plan.updated` on enum `Event.Type`
* Add support for `ManagedPayments` on `InvoiceItemParams`, `InvoiceItem`, `InvoiceParams`, `Invoice`, and `QuotePreviewInvoice`
* Add support for `PaymentPlan` on `Invoice`
* Add support for `EstimatedFeeDetails` and `EstimatedFee` on `IssuingAuthorizationPendingRequestHoldAmountDetails` and `IssuingAuthorizationRequestHistoryHoldAmountDetails`
* ⚠️ Remove support for `Cryptogram` on `PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecure` and `PaymentRecordPaymentMethodDetailsCardThreeDSecure`
* ⚠️ Change type of `ProductCatalogTrialOfferEndBehaviorTransition.Price` from `$Price` to `deletable($Price)`
* ⚠️ Change type of `SubscriptionCancellationDetails.FeedbackOption` from `$Billing.FeedbackOptions` to `$Billing.FeedbackOption`
* Add support for `CancelAtPeriodEnd` on `SubscriptionPendingUpdate`
* Add support for `Igic` on `TaxRegistrationCountryOptionsAt`, `TaxRegistrationCountryOptionsBe`, `TaxRegistrationCountryOptionsBg`, `TaxRegistrationCountryOptionsCy`, `TaxRegistrationCountryOptionsCz`, `TaxRegistrationCountryOptionsDe`, `TaxRegistrationCountryOptionsDk`, `TaxRegistrationCountryOptionsEe`, `TaxRegistrationCountryOptionsEs`, `TaxRegistrationCountryOptionsFi`, `TaxRegistrationCountryOptionsFr`, `TaxRegistrationCountryOptionsGr`, `TaxRegistrationCountryOptionsHr`, `TaxRegistrationCountryOptionsHu`, `TaxRegistrationCountryOptionsIe`, `TaxRegistrationCountryOptionsIt`, `TaxRegistrationCountryOptionsLt`, `TaxRegistrationCountryOptionsLu`, `TaxRegistrationCountryOptionsLv`, `TaxRegistrationCountryOptionsMt`, `TaxRegistrationCountryOptionsNl`, `TaxRegistrationCountryOptionsPl`, `TaxRegistrationCountryOptionsPt`, `TaxRegistrationCountryOptionsRo`, `TaxRegistrationCountryOptionsSe`, `TaxRegistrationCountryOptionsSi`, and `TaxRegistrationCountryOptionsSk`
* Add support for `Metadata` on `V2BillingContractParams`, `V2BillingContractPricingLineActionUpdateParams`, `V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesData`, `V2BillingContractPricingOverrideActionAddParams`, `V2BillingContractPricingOverrideActionUpdateParams`, `V2BillingContractPricingOverrideParams`, `V2BillingContractPricingOverridesData`, and `V2MoneyManagementTransaction`
* Add support for `TaxAmount` on `V2MoneyManagementOutboundPaymentQuoteEstimatedFee`
* Add support for `PayoutMethodOptions` on `V2MoneyManagementOutboundPaymentQuoteToParams` and `V2MoneyManagementOutboundPaymentQuoteTo`
* Change type of `V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionUpdateParams.Metadata` from `string` to `emptyable(string)`
* Add support for snapshot events `EventTypePaymentPlanCreated`, `EventTypePaymentPlanInstallmentDue`, `EventTypePaymentPlanInstallmentPaid`, `EventTypePaymentPlanInstallmentWillBeDue`, and `EventTypePaymentPlanUpdated` with resource `PaymentPlan`
