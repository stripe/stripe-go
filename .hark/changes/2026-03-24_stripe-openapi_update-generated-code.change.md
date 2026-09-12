---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2305
is_breaking: true
is_stripe_api_change: true
released_in_version: 85.1.0-alpha.1
---

* Add support for new resource `V2CoreAccountEvaluation`
* ⚠️ Remove support for resources `V2BillingLicenseFeeSubscription` and `V2BillingPricingPlanSubscriptionComponents`
* Add support for `New` method on resource `V2CoreAccountEvaluation`
* ⚠️ Remove support for `Get` method on resources `V2BillingLicenseFeeSubscription` and `V2BillingPricingPlanSubscriptionComponents`
* Add support for `ModifyRates` method on resource `V2BillingRateCard`
* Add support for `RemoveDiscounts` method on resource `V2BillingPricingPlanSubscription`
* Add support for new value `eg_bank_account` on enum `V2AccountConfigurationRecipientDataDefaultOutboundDestination.Type`
* Add support for `InvoiceResources` on `V2BillingIntent`
* Add support for `AmountDue` and `CustomerBalanceApplied` on `V2BillingIntentAmountDetails`
* Add support for `ExpiresAt` on `V2BillingIntentStatusTransitions`
* Add support for `Discount` on `V2BillingIntentActionApplyParams` and `V2BillingIntentActionApply`
* Add support for `Timestamp` on `V2BillingIntentActionApplyEffectiveAtParams` and `V2BillingIntentActionApplyEffectiveAt`
* Add support for new values `current_billing_period_start` and `timestamp` on enum `V2BillingIntentActionApplyEffectiveAt.Type`
* Add support for new value `discount` on enum `V2BillingIntentActionApply.Type`
* ⚠️ Change type of `V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehavior.Type`, `V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams.Type`, `V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehavior.Type`, `V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams.Type`, `V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehavior.Type`, and `V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams.Type` from `literal('license_fee')` to `enum('license_fee'|'recurring_credit_grant')`
* Add support for `ServiceCycle` on `V2BillingLicenseFee` and `V2BillingRateCard`
* ⚠️ Remove support for `LatestVersion` on `V2BillingLicenseFee`, `V2BillingPricingPlan`, and `V2BillingRateCard`
* ⚠️ Remove support for `ServiceIntervalCount` and `ServiceInterval` on `V2BillingLicenseFee` and `V2BillingRateCard`
* ⚠️ Change type of `V2BillingLicenseFeeTransformQuantity.DivideBy`, `V2BillingLicenseFeeTransformQuantityParams.DivideBy`, `V2BillingLicenseFeeVersionTransformQuantity.DivideBy`, `V2BillingRateCardRateTransformQuantity.DivideBy`, and `V2BillingRateCardRateTransformQuantityParams.DivideBy` from `longInteger` to `int64_string`
* Add support for `DiscountDetails` and `PricingPlanComponentDetails` on `V2BillingPricingPlanSubscription`
* Add support for new value `crypto_wallets` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
* ⚠️ Remove support for value `crypto` from enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
* Add support for `BalanceByFundsType` on `V2MoneyManagementFinancialAccountPayments`
* Add support for new value `next_day_payout_fee` on enum `V2MoneyManagementOutboundPaymentQuoteEstimatedFee.Type`
* Add support for `TreasuryTransactionEntry` on `V2MoneyManagementTransactionEntry`
* Add support for `TreasuryCreditReversal`, `TreasuryDebitReversal`, `TreasuryInboundTransfer`, `TreasuryIssuingAuthorization`, `TreasuryOutboundPayment`, `TreasuryOutboundTransfer`, `TreasuryReceivedCredit`, and `TreasuryReceivedDebit` on `V2MoneyManagementTransactionEntryTransactionDetailsFlow` and `V2MoneyManagementTransactionFlow`
* Add support for new values `treasury_credit_reversal`, `treasury_debit_reversal`, `treasury_inbound_transfer`, `treasury_issuing_authorization`, `treasury_other`, `treasury_outbound_payment`, `treasury_outbound_transfer`, `treasury_received_credit`, and `treasury_received_debit` on enums `V2MoneyManagementTransactionEntryTransactionDetailsFlow.Type` and `V2MoneyManagementTransactionFlow.Type`
* Add support for `TreasuryTransaction` on `V2MoneyManagementTransaction`
* Add support for new value `no_valid_payment_method` on enum `V2PaymentsOffSessionPayment.FailureReason`
* Add support for `Metadata` on `V2PaymentsSettlementAllocationIntentSplit`
* ⚠️ Change type of `V2ReportingReportRunResultFile.Size` from `longInteger` to `int64_string`
* Add support for `StatementDescriptor` on `V2MoneyManagementOutboundPaymentParams` and `V2MoneyManagementOutboundTransferParams`
* Add support for `Include` on `V2BillingIntentParams`, `V2BillingIntentReserveParams`, `V2BillingPricingPlanSubscriptionListParams`, `V2BillingPricingPlanSubscriptionParams`, `V2MoneyManagementFinancialAccountListParams`, and `V2MoneyManagementFinancialAccountParams`
* Add support for event notifications `V1AccountSignalsIncludingDelinquencyCreatedEvent`, `V2CoreAccountSignalsFraudulentWebsiteReadyEvent`, and `V2SignalsAccountSignalFraudulentMerchantReadyEvent`
