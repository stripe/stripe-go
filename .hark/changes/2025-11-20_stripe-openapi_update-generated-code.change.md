---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2221
is_stripe_api_change: true
released_in_version: 84.1.0-alpha.2
---

* Add support for new resources `V2CoreAccountPersonToken`, `V2CoreAccountToken`, and `V2MoneyManagementCurrencyConversion`
* Add support for `Get`, `List`, and `New` methods on resource `V2MoneyManagementCurrencyConversion`
* Add support for `Get` and `New` methods on resources `V2CoreAccountPersonToken` and `V2CoreAccountToken`
* Add support for `EffectiveAt` on `InvoiceCreatePreviewScheduleDetailsAmendmentParams`, `InvoiceCreatePreviewScheduleDetailsPhaseParams`, `QuoteLineParams`, `QuoteLine`, `QuotePreviewSubscriptionSchedulePhase`, `SubscriptionScheduleAmendAmendmentParams`, `SubscriptionSchedulePhaseParams`, and `SubscriptionSchedulePhase`
* Add support for `TrialOffer` on `InvoiceCreatePreviewScheduleDetailsAmendmentItemActionAddParams`, `InvoiceCreatePreviewScheduleDetailsAmendmentItemActionSetParams`, `InvoiceCreatePreviewScheduleDetailsPhaseItemParams`, `QuoteLineActionAddItemParams`, `QuoteLineActionAddItem`, `QuoteLineActionSetItemParams`, `QuoteLineActionSetItems`, `QuotePreviewSubscriptionSchedulePhaseItem`, `SubscriptionScheduleAmendAmendmentItemActionAddParams`, `SubscriptionScheduleAmendAmendmentItemActionSetParams`, `SubscriptionSchedulePhaseItemParams`, and `SubscriptionSchedulePhaseItem`
* Add support for `AmountDiscount`, `AmountSubtotal`, `AmountTotal`, `UnitAmountAfterDiscount`, and `UnitDiscount` on `DelegatedCheckoutRequestedSessionLineItemDetail`
* Add support for `AmountSubtotalAfterDiscount` on `DelegatedCheckoutRequestedSessionLineItemDetail` and `DelegatedCheckoutRequestedSessionTotalDetails`
* Change type of `InvoiceCreatePreviewScheduleDetailsParams.BillingSchedules` from `array(billing_schedules_update_params)` to `emptyable(array(billing_schedules_update_params))`
* Add support for `CurrentTrial` on `InvoiceCreatePreviewSubscriptionDetailsItemParams`, `SubscriptionItemParams`, and `SubscriptionItem`
* Change type of `QuoteSubscriptionDataOverrideParams.BillingSchedules` and `QuoteSubscriptionDataParams.BillingSchedules` from `emptyable(array(billing_schedules_create_specs))` to `array(billing_schedules_create_specs)`
* Change type of `QuoteSubscriptionData.BillingSchedules` and `QuoteSubscriptionDataOverrides.BillingSchedules` from `nullable(array(SubscriptionsResourceBillingSchedules))` to `array(QuotesResourceSubscriptionDataBillingSchedules)`
* Change type of `QuoteSubscriptionData.PhaseEffectiveAt` and `QuoteSubscriptionDataOverrides.PhaseEffectiveAt` from `nullable(enum('billing_period_start'|'phase_start'))` to `enum('billing_period_start'|'line_start')`
* Change type of `QuotePreviewSubscriptionSchedule.BillingSchedules` and `SubscriptionSchedule.BillingSchedules` from `nullable(array(SubscriptionsResourceBillingSchedules))` to `array(SubscriptionsResourceBillingSchedules)`
* Remove support for `AmendmentStart`, `LineStartsAt`, and `Relative` on `SubscriptionBillingScheduleBillFrom`
* Change type of `SubscriptionBillingScheduleBillFrom.Type` from `enum` to `literal('timestamp')`
* Remove support for `AmendmentEnd` and `LineEndsAt` on `SubscriptionBillingScheduleBillUntil`
* Remove support for values `amendment_end`, `line_ends_at`, `schedule_end`, and `upcoming_invoice` from enum `SubscriptionBillingScheduleBillUntil.Type`
* Change type of `V2BillingServiceActionCreditGrantAmount.Monetary`, `V2BillingServiceActionCreditGrantAmountParams.Monetary`, `V2BillingServiceActionCreditGrantPerTenantAmount.Monetary`, and `V2BillingServiceActionCreditGrantPerTenantAmountParams.Monetary` from `amount` to `an object`
* Add support for `FutureRequirements` on `V2CoreAccount`
* Add support for `KonbiniPayments` and `ScriptStatementDescriptor` on `V2CoreAccountConfigurationMerchantParams` and `V2CoreAccountConfigurationMerchant`
* Add support for `EUR` on `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesParams` and `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrencies`
* Add support for `RequirementsCollector` on `V2CoreAccountDefaultsResponsibilities`
* Add support for new value `ar_cuit` on enum `V2CoreAccountIdentityBusinessDetailsIdNumber.Type`
* Add support for new value `ar_dni` on enums `V2CoreAccountIdentityIndividualIdNumber.Type` and `V2CoreAccountPersonIdNumber.Type`
* Remove support for `Collector` on `V2CoreAccountRequirements`
* Add support for new value `holds_currencies.eur` on enum `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
* Add support for new values `payment_method` and `person` on enum `V2CoreAccountRequirementsEntryReference.Type`
* Remove support for value `resource` from enum `V2CoreAccountRequirementsEntryReference.Type`
* Remove support for value `future_requirements` from enum `V2CoreAccountRequirementsEntryRequestedReason.Code`
* Remove support for `V1EventID` on `V2CoreEvent`
* Remove support for `AmountDetails` and `CaptureMethod` on `V2PaymentsOffSessionPaymentParams` and `V2PaymentsOffSessionPayment`
* Change type of `V2PaymentsOffSessionPayment.AmountCapturable` from `amount` to `an object`
* Change type of `V2PaymentsOffSessionPayment.AmountRequested` from `amount` to `an object`
* Change type of `V2PaymentsOffSessionPaymentParams.Amount` from `amount` to `an object`
* Remove support for `Destination` on `V2PaymentsOffSessionPaymentCaptureTransferDataParams`
* Add support for `Created` on `V2CoreEventListParams`
* Remove support for `GTE`, `Gt`, `LT`, and `Lte` on `V2CoreEventListParams`
* Add support for `AccountToken` on `V2CoreAccountParams`
* Add support for `PersonToken` on `V2CoreAccountPersonParams`
* Add support for `ImpactedRequestsPercentage` on `EventsV2CoreHealthApiErrorFiringEventImpact`, `EventsV2CoreHealthApiErrorResolvedEventImpact`, `EventsV2CoreHealthApiLatencyFiringEventImpact`, `EventsV2CoreHealthApiLatencyResolvedEventImpact`, `EventsV2CoreHealthPaymentMethodErrorFiringEventImpact`, and `EventsV2CoreHealthPaymentMethodErrorResolvedEventImpact`
* Add support for `Context` and `RelatedObject` on `EventsV2CoreHealthEventGenerationFailureResolvedEventImpact`
* Remove support for `Account`, `Livemode`, `MissingDeliveryAttempts`, and `RelatedObjectID` on `EventsV2CoreHealthEventGenerationFailureResolvedEventImpact`
* Change type of `EventsV2CoreHealthFraudRateIncreasedEventImpact.RealizedFraudAmount` from `amount` to `an object`
* Change type of `EventsV2CoreHealthIssuingAuthorizationRequestErrorsFiringEventImpact.ApprovedAmount`, `EventsV2CoreHealthIssuingAuthorizationRequestErrorsResolvedEventImpact.ApprovedAmount`, `EventsV2CoreHealthIssuingAuthorizationRequestTimeoutFiringEventImpact.ApprovedAmount`, and `EventsV2CoreHealthIssuingAuthorizationRequestTimeoutResolvedEventImpact.ApprovedAmount` from `amount` to `an object`
* Change type of `EventsV2CoreHealthIssuingAuthorizationRequestErrorsFiringEventImpact.DeclinedAmount`, `EventsV2CoreHealthIssuingAuthorizationRequestErrorsResolvedEventImpact.DeclinedAmount`, `EventsV2CoreHealthIssuingAuthorizationRequestTimeoutFiringEventImpact.DeclinedAmount`, and `EventsV2CoreHealthIssuingAuthorizationRequestTimeoutResolvedEventImpact.DeclinedAmount` from `amount` to `an object`
* Add support for thin events `V2PaymentsOffSessionPaymentAttemptFailedEvent` and `V2PaymentsOffSessionPaymentAttemptStartedEvent` with related object `V2PaymentsOffSessionPayment`
* Remove support for thin event `V1AccountUpdatedEvent` with related object `Account`
* Remove support for thin events `V1ApplicationFeeCreatedEvent` and `V1ApplicationFeeRefundedEvent` with related object `ApplicationFee`
* Remove support for thin events `V1BillingPortalConfigurationCreatedEvent` and `V1BillingPortalConfigurationUpdatedEvent` with related object `BillingPortalConfiguration`
* Remove support for thin event `V1CapabilityUpdatedEvent` with related object `Capability`
* Remove support for thin events `V1ChargeCapturedEvent`, `V1ChargeExpiredEvent`, `V1ChargeFailedEvent`, `V1ChargePendingEvent`, `V1ChargeRefundedEvent`, `V1ChargeSucceededEvent`, and `V1ChargeUpdatedEvent` with related object `Charge`
* Remove support for thin events `V1ChargeDisputeClosedEvent`, `V1ChargeDisputeCreatedEvent`, `V1ChargeDisputeFundsReinstatedEvent`, `V1ChargeDisputeFundsWithdrawnEvent`, and `V1ChargeDisputeUpdatedEvent` with related object `Dispute`
* Remove support for thin events `V1ChargeRefundUpdatedEvent`, `V1RefundCreatedEvent`, `V1RefundFailedEvent`, and `V1RefundUpdatedEvent` with related object `Refund`
* Remove support for thin events `V1CheckoutSessionAsyncPaymentFailedEvent`, `V1CheckoutSessionAsyncPaymentSucceededEvent`, `V1CheckoutSessionCompletedEvent`, and `V1CheckoutSessionExpiredEvent` with related object `CheckoutSession`
* Remove support for thin events `V1ClimateOrderCanceledEvent`, `V1ClimateOrderCreatedEvent`, `V1ClimateOrderDelayedEvent`, `V1ClimateOrderDeliveredEvent`, and `V1ClimateOrderProductSubstitutedEvent` with related object `ClimateOrder`
* Remove support for thin events `V1ClimateProductCreatedEvent` and `V1ClimateProductPricingUpdatedEvent` with related object `ClimateProduct`
* Remove support for thin events `V1CouponCreatedEvent`, `V1CouponDeletedEvent`, and `V1CouponUpdatedEvent` with related object `Coupon`
* Remove support for thin events `V1CreditNoteCreatedEvent`, `V1CreditNoteUpdatedEvent`, and `V1CreditNoteVoidedEvent` with related object `CreditNote`
* Remove support for thin events `V1CustomerCreatedEvent`, `V1CustomerDeletedEvent`, and `V1CustomerUpdatedEvent` with related object `Customer`
* Remove support for thin events `V1CustomerSubscriptionCreatedEvent`, `V1CustomerSubscriptionDeletedEvent`, `V1CustomerSubscriptionPausedEvent`, `V1CustomerSubscriptionPendingUpdateAppliedEvent`, `V1CustomerSubscriptionPendingUpdateExpiredEvent`, `V1CustomerSubscriptionResumedEvent`, `V1CustomerSubscriptionTrialWillEndEvent`, and `V1CustomerSubscriptionUpdatedEvent` with related object `Subscription`
* Remove support for thin events `V1CustomerTaxIdCreatedEvent`, `V1CustomerTaxIdDeletedEvent`, and `V1CustomerTaxIdUpdatedEvent` with related object `TaxID`
* Remove support for thin event `V1FileCreatedEvent` with related object `File`
* Remove support for thin events `V1FinancialConnectionsAccountCreatedEvent`, `V1FinancialConnectionsAccountDeactivatedEvent`, `V1FinancialConnectionsAccountDisconnectedEvent`, `V1FinancialConnectionsAccountReactivatedEvent`, `V1FinancialConnectionsAccountRefreshedBalanceEvent`, `V1FinancialConnectionsAccountRefreshedOwnershipEvent`, and `V1FinancialConnectionsAccountRefreshedTransactionsEvent` with related object `FinancialConnectionsAccount`
* Remove support for thin events `V1IdentityVerificationSessionCanceledEvent`, `V1IdentityVerificationSessionCreatedEvent`, `V1IdentityVerificationSessionProcessingEvent`, `V1IdentityVerificationSessionRedactedEvent`, `V1IdentityVerificationSessionRequiresInputEvent`, and `V1IdentityVerificationSessionVerifiedEvent` with related object `IdentityVerificationSession`
* Remove support for thin events `V1InvoiceCreatedEvent`, `V1InvoiceDeletedEvent`, `V1InvoiceFinalizationFailedEvent`, `V1InvoiceFinalizedEvent`, `V1InvoiceMarkedUncollectibleEvent`, `V1InvoiceOverdueEvent`, `V1InvoiceOverpaidEvent`, `V1InvoicePaidEvent`, `V1InvoicePaymentActionRequiredEvent`, `V1InvoicePaymentFailedEvent`, `V1InvoicePaymentSucceededEvent`, `V1InvoiceSentEvent`, `V1InvoiceUpcomingEvent`, `V1InvoiceUpdatedEvent`, `V1InvoiceVoidedEvent`, and `V1InvoiceWillBeDueEvent` with related object `Invoice`
* Remove support for thin event `V1InvoicePaymentPaidEvent` with related object `InvoicePayment`
* Remove support for thin events `V1InvoiceitemCreatedEvent` and `V1InvoiceitemDeletedEvent` with related object `InvoiceItem`
* Remove support for thin events `V1IssuingAuthorizationCreatedEvent`, `V1IssuingAuthorizationRequestEvent`, and `V1IssuingAuthorizationUpdatedEvent` with related object `IssuingAuthorization`
* Remove support for thin events `V1IssuingCardCreatedEvent` and `V1IssuingCardUpdatedEvent` with related object `IssuingCard`
* Remove support for thin events `V1IssuingCardholderCreatedEvent` and `V1IssuingCardholderUpdatedEvent` with related object `IssuingCardholder`
* Remove support for thin events `V1IssuingDisputeClosedEvent`, `V1IssuingDisputeCreatedEvent`, `V1IssuingDisputeFundsReinstatedEvent`, `V1IssuingDisputeFundsRescindedEvent`, `V1IssuingDisputeSubmittedEvent`, and `V1IssuingDisputeUpdatedEvent` with related object `IssuingDispute`
* Remove support for thin events `V1IssuingPersonalizationDesignActivatedEvent`, `V1IssuingPersonalizationDesignDeactivatedEvent`, `V1IssuingPersonalizationDesignRejectedEvent`, and `V1IssuingPersonalizationDesignUpdatedEvent` with related object `IssuingPersonalizationDesign`
* Remove support for thin events `V1IssuingTokenCreatedEvent` and `V1IssuingTokenUpdatedEvent` with related object `IssuingToken`
* Remove support for thin events `V1IssuingTransactionCreatedEvent`, `V1IssuingTransactionPurchaseDetailsReceiptUpdatedEvent`, and `V1IssuingTransactionUpdatedEvent` with related object `IssuingTransaction`
* Remove support for thin event `V1MandateUpdatedEvent` with related object `Mandate`
* Remove support for thin events `V1PaymentIntentAmountCapturableUpdatedEvent`, `V1PaymentIntentCanceledEvent`, `V1PaymentIntentCreatedEvent`, `V1PaymentIntentPartiallyFundedEvent`, `V1PaymentIntentPaymentFailedEvent`, `V1PaymentIntentProcessingEvent`, `V1PaymentIntentRequiresActionEvent`, and `V1PaymentIntentSucceededEvent` with related object `PaymentIntent`
* Remove support for thin events `V1PaymentLinkCreatedEvent` and `V1PaymentLinkUpdatedEvent` with related object `PaymentLink`
* Remove support for thin events `V1PaymentMethodAttachedEvent`, `V1PaymentMethodAutomaticallyUpdatedEvent`, `V1PaymentMethodDetachedEvent`, and `V1PaymentMethodUpdatedEvent` with related object `PaymentMethod`
* Remove support for thin events `V1PayoutCanceledEvent`, `V1PayoutCreatedEvent`, `V1PayoutFailedEvent`, `V1PayoutPaidEvent`, `V1PayoutReconciliationCompletedEvent`, and `V1PayoutUpdatedEvent` with related object `Payout`
* Remove support for thin events `V1PersonCreatedEvent`, `V1PersonDeletedEvent`, and `V1PersonUpdatedEvent` with related object `Person`
* Remove support for thin events `V1PlanCreatedEvent`, `V1PlanDeletedEvent`, and `V1PlanUpdatedEvent` with related object `Plan`
* Remove support for thin events `V1PriceCreatedEvent`, `V1PriceDeletedEvent`, and `V1PriceUpdatedEvent` with related object `Price`
* Remove support for thin events `V1ProductCreatedEvent`, `V1ProductDeletedEvent`, and `V1ProductUpdatedEvent` with related object `Product`
* Remove support for thin events `V1PromotionCodeCreatedEvent` and `V1PromotionCodeUpdatedEvent` with related object `PromotionCode`
* Remove support for thin events `V1QuoteAcceptedEvent`, `V1QuoteCanceledEvent`, `V1QuoteCreatedEvent`, and `V1QuoteFinalizedEvent` with related object `Quote`
* Remove support for thin events `V1RadarEarlyFraudWarningCreatedEvent` and `V1RadarEarlyFraudWarningUpdatedEvent` with related object `RadarEarlyFraudWarning`
* Remove support for thin events `V1ReviewClosedEvent` and `V1ReviewOpenedEvent` with related object `Review`
* Remove support for thin events `V1SetupIntentCanceledEvent`, `V1SetupIntentCreatedEvent`, `V1SetupIntentRequiresActionEvent`, `V1SetupIntentSetupFailedEvent`, and `V1SetupIntentSucceededEvent` with related object `SetupIntent`
* Remove support for thin event `V1SigmaScheduledQueryRunCreatedEvent` with related object `SigmaScheduledQueryRun`
* Remove support for thin events `V1SourceCanceledEvent`, `V1SourceChargeableEvent`, `V1SourceFailedEvent`, and `V1SourceRefundAttributesRequiredEvent` with related object `Source`
* Remove support for thin events `V1SubscriptionScheduleAbortedEvent`, `V1SubscriptionScheduleCanceledEvent`, `V1SubscriptionScheduleCompletedEvent`, `V1SubscriptionScheduleCreatedEvent`, `V1SubscriptionScheduleExpiringEvent`, `V1SubscriptionScheduleReleasedEvent`, and `V1SubscriptionScheduleUpdatedEvent` with related object `SubscriptionSchedule`
* Remove support for thin events `V1TaxRateCreatedEvent` and `V1TaxRateUpdatedEvent` with related object `TaxRate`
* Remove support for thin events `V1TerminalReaderActionFailedEvent`, `V1TerminalReaderActionSucceededEvent`, and `V1TerminalReaderActionUpdatedEvent` with related object `TerminalReader`
* Remove support for thin events `V1TestHelpersTestClockAdvancingEvent`, `V1TestHelpersTestClockCreatedEvent`, `V1TestHelpersTestClockDeletedEvent`, `V1TestHelpersTestClockInternalFailureEvent`, and `V1TestHelpersTestClockReadyEvent` with related object `TestHelpersTestClock`
* Remove support for thin events `V1TopupCanceledEvent`, `V1TopupCreatedEvent`, `V1TopupFailedEvent`, `V1TopupReversedEvent`, and `V1TopupSucceededEvent` with related object `Topup`
* Remove support for thin events `V1TransferCreatedEvent`, `V1TransferReversedEvent`, and `V1TransferUpdatedEvent` with related object `Transfer`
