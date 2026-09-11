---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2106
is_stripe_api_change: true
released_in_version: 82.6.0-alpha.1
---

* Add support for `AttachCadence` method on resource `Subscription`
* Add support for `Currency` and `ExternalCustomerID` on `BillingAlertTriggered`
* Add support for `CustomPricingUnit` on `BillingAlertTriggered`, `BillingCreditBalanceSummaryBalanceAvailableBalance`, `BillingCreditBalanceSummaryBalanceLedgerBalance`, `BillingCreditBalanceTransactionCreditAmount`, `BillingCreditBalanceTransactionDebitAmount`, `BillingCreditGrantAmountParams`, and `BillingCreditGrantAmount`
* Add support for `Customer` on `BillingAlertListParams`
* Change type of `BillingAlert.AlertType`, `BillingAlertListParams.AlertType`, and `BillingAlertParams.AlertType` from `literal('usage_threshold')` to `enum('credit_balance_threshold'|'usage_threshold')`
* Add support for `CreditBalanceThreshold` on `BillingAlertParams` and `BillingAlert`
* Add support for `BillableItems` on `BillingCreditBalanceSummaryFilterApplicabilityScopeParams`, `BillingCreditGrantApplicabilityConfigScopeParams`, and `BillingCreditGrantApplicabilityConfigScope`
* Change type of `BillingCreditBalanceSummaryBalanceAvailableBalance.Type`, `BillingCreditBalanceSummaryBalanceLedgerBalance.Type`, `BillingCreditBalanceTransactionCreditAmount.Type`, `BillingCreditBalanceTransactionDebitAmount.Type`, `BillingCreditGrantAmount.Type`, and `BillingCreditGrantAmountParams.Type` from `literal('monetary')` to `enum('custom_pricing_unit'|'monetary')`
* Add support for `LicenseFeeSubscriptionDetails` and `RateCardSubscriptionDetails` on `InvoiceItemParent` and `InvoiceLineItemParent`
* Change type of `InvoiceItemParent.Type` from `literal('subscription_details')` to `enum('license_fee_subscription_details'|'rate_card_subscription_details'|'subscription_details')`
* Add support for `LicenseFeeDetails` and `RateCardRateDetails` on `InvoiceItemPricing` and `InvoiceLineItemPricing`
* Change type of `InvoiceItemPricing.Type` and `InvoiceLineItemPricing.Type` from `literal('price_details')` to `enum('license_fee_details'|'price_details'|'rate_card_rate_details')`
* Add support for `BillingCadence` on `InvoiceCreatePreviewParams`, `SubscriptionParams`, and `Subscription`
* Add support for `BillingCadenceDetails` on `InvoiceParent` and `QuotePreviewInvoiceParent`
* Add support for new value `billing_cadence_details` on enums `InvoiceParent.Type` and `QuotePreviewInvoiceParent.Type`
* Add support for new values `license_fee_subscription_details` and `rate_card_subscription_details` on enum `InvoiceLineItemParent.Type`
* Add support for new resources `V2BillingBillSettingVersion`, `V2BillingBillSetting`, `V2BillingCadence`, `V2BillingCollectionSettingVersion`, `V2BillingCollectionSetting`, `V2BillingCustomPricingUnit`, `V2BillingIntentAction`, `V2BillingIntent`, `V2BillingLicenseFeeSubscription`, `V2BillingLicenseFeeVersion`, `V2BillingLicenseFee`, `V2BillingLicensedItem`, `V2BillingMeteredItem`, `V2BillingPricingPlanComponent`, `V2BillingPricingPlanSubscription`, `V2BillingPricingPlanVersion`, `V2BillingPricingPlan`, `V2BillingProfile`, `V2BillingRateCardRate`, `V2BillingRateCardSubscription`, `V2BillingRateCardVersion`, `V2BillingRateCard`, `V2BillingServiceAction`, `V2CoreClaimableSandbox`, `V2ReportingReportRun`, `V2ReportingReport`, and `V2TaxAutomaticRule`
* Add support for `Deactivate`, `Find`, `Get`, `New`, and `Update` methods on resource `V2TaxAutomaticRule`
* Add support for `Get` and `New` methods on resources `V2BillingServiceAction` and `V2ReportingReportRun`
* Add support for `Get` method on resources `V2BillingLicenseFeeSubscription` and `V2ReportingReport`
* Add support for `New` method on resource `V2CoreClaimableSandbox`
* Add support for `Cancel`, `Get`, `List`, `New`, and `Update` methods on resources `V2BillingCadence` and `V2BillingRateCardSubscription`
* Add support for `Get`, `List`, `New`, and `Update` methods on resources `V2BillingBillSetting`, `V2BillingCollectionSetting`, `V2BillingCustomPricingUnit`, `V2BillingLicenseFee`, `V2BillingLicensedItem`, `V2BillingMeteredItem`, `V2BillingPricingPlan`, `V2BillingProfile`, and `V2BillingRateCard`
* Add support for `Get` and `List` methods on resources `V2BillingBillSettingVersion`, `V2BillingCollectionSettingVersion`, `V2BillingIntentAction`, `V2BillingLicenseFeeVersion`, `V2BillingPricingPlanSubscription`, `V2BillingPricingPlanVersion`, and `V2BillingRateCardVersion`
* Add support for `Del`, `Get`, `List`, and `New` methods on resource `V2BillingRateCardRate`
* Add support for `Del`, `Get`, `List`, `New`, and `Update` methods on resource `V2BillingPricingPlanComponent`
* Add support for `Cancel`, `Commit`, `Get`, `List`, `New`, `ReleaseReservation`, and `Reserve` methods on resource `V2BillingIntent`
* Add support for `Changes` on `V2Event`
* Add support for thin events `V2BillingCadenceBilledEvent`, `V2BillingCadenceCanceledEvent`, `V2BillingCadenceCreatedEvent`, and `V2BillingCadenceErroredEvent` with related object `V2BillingCadence`
* Add support for thin events `V2BillingLicenseFeeCreatedEvent` and `V2BillingLicenseFeeUpdatedEvent` with related object `V2BillingLicenseFee`
* Add support for thin event `V2BillingLicenseFeeVersionCreatedEvent` with related object `V2BillingLicenseFeeVersion`
* Add support for thin events `V2BillingLicensedItemCreatedEvent` and `V2BillingLicensedItemUpdatedEvent` with related object `V2BillingLicensedItem`
* Add support for thin events `V2BillingMeteredItemCreatedEvent` and `V2BillingMeteredItemUpdatedEvent` with related object `V2BillingMeteredItem`
* Add support for thin events `V2BillingPricingPlanCreatedEvent` and `V2BillingPricingPlanUpdatedEvent` with related object `V2BillingPricingPlan`
* Add support for thin events `V2BillingPricingPlanComponentCreatedEvent` and `V2BillingPricingPlanComponentUpdatedEvent` with related object `V2BillingPricingPlanComponent`
* Add support for thin events `V2BillingPricingPlanSubscriptionCollectionAwaitingCustomerActionEvent`, `V2BillingPricingPlanSubscriptionCollectionCurrentEvent`, `V2BillingPricingPlanSubscriptionCollectionPastDueEvent`, `V2BillingPricingPlanSubscriptionCollectionPausedEvent`, `V2BillingPricingPlanSubscriptionCollectionUnpaidEvent`, `V2BillingPricingPlanSubscriptionServicingActivatedEvent`, `V2BillingPricingPlanSubscriptionServicingCanceledEvent`, and `V2BillingPricingPlanSubscriptionServicingPausedEvent` with related object `V2BillingPricingPlanSubscription`
* Add support for thin event `V2BillingPricingPlanVersionCreatedEvent` with related object `V2BillingPricingPlanVersion`
* Add support for thin events `V2BillingRateCardCreatedEvent` and `V2BillingRateCardUpdatedEvent` with related object `V2BillingRateCard`
* Add support for thin event `V2BillingRateCardRateCreatedEvent` with related object `V2BillingRateCardRate`
* Add support for thin events `V2BillingRateCardSubscriptionActivatedEvent`, `V2BillingRateCardSubscriptionCanceledEvent`, `V2BillingRateCardSubscriptionCollectionAwaitingCustomerActionEvent`, `V2BillingRateCardSubscriptionCollectionCurrentEvent`, `V2BillingRateCardSubscriptionCollectionPastDueEvent`, `V2BillingRateCardSubscriptionCollectionPausedEvent`, `V2BillingRateCardSubscriptionCollectionUnpaidEvent`, `V2BillingRateCardSubscriptionServicingActivatedEvent`, `V2BillingRateCardSubscriptionServicingCanceledEvent`, and `V2BillingRateCardSubscriptionServicingPausedEvent` with related object `V2BillingRateCardSubscription`
* Add support for thin event `V2BillingRateCardVersionCreatedEvent` with related object `V2BillingRateCardVersion`
* Add support for thin events `V2CoreHealthApiErrorFiringEvent`, `V2CoreHealthApiErrorResolvedEvent`, `V2CoreHealthApiLatencyFiringEvent`, `V2CoreHealthApiLatencyResolvedEvent`, `V2CoreHealthAuthorizationRateDropFiringEvent`, `V2CoreHealthAuthorizationRateDropResolvedEvent`, `V2CoreHealthEventGenerationFailureResolvedEvent`, `V2CoreHealthFraudRateIncreasedEvent`, `V2CoreHealthIssuingAuthorizationRequestTimeoutFiringEvent`, `V2CoreHealthIssuingAuthorizationRequestTimeoutResolvedEvent`, `V2CoreHealthPaymentMethodErrorFiringEvent`, `V2CoreHealthPaymentMethodErrorResolvedEvent`, `V2CoreHealthTrafficVolumeDropFiringEvent`, `V2CoreHealthTrafficVolumeDropResolvedEvent`, `V2CoreHealthWebhookLatencyFiringEvent`, and `V2CoreHealthWebhookLatencyResolvedEvent`
* Add support for thin events `V2ReportingReportRunCreatedEvent`, `V2ReportingReportRunFailedEvent`, `V2ReportingReportRunSucceededEvent`, and `V2ReportingReportRunUpdatedEvent` with related object `V2ReportingReportRun`
* Add support for error type `RateLimitError`
