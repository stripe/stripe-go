---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2283
is_breaking: true
is_stripe_api_change: true
released_in_version: 84.5.0-alpha.2
---

* Add support for new resources `BillingAlertRecovered` and `Profile`
* Add support for `Reauthorize` method on resource `PaymentIntent`
* Add support for `Settings` on `QuoteLineActionAddDiscount`, `QuoteLineActionAddItemDiscount`, `QuoteLineActionSetDiscounts`, `QuoteLineActionSetItemsDiscount`, `QuotePreviewSubscriptionSchedulePhaseDiscount`, `QuotePreviewSubscriptionSchedulePhaseItemDiscount`, `SubscriptionSchedulePhaseDiscount`, and `SubscriptionSchedulePhaseItemDiscount`
* Add support for `SmartDisputes` on `AccountSettingsParams`, `AccountSettings`, `V2CoreAccountConfigurationMerchantParams`, and `V2CoreAccountConfigurationMerchant`
* Add support for `EmailCustomersOnSuccessfulPayment` on `AccountSettingsPaymentsParams` and `AccountSettingsPayments`
* Add support for `BalanceUpdateDetails` on `BillingCreditBalanceSummaryBalance`
* Add support for `Reauthorization` and `ReauthorizeBefore` on `ChargePaymentMethodDetailsCardPresent`, `ChargePaymentMethodDetailsCard`, `ConfirmationTokenPaymentMethodPreviewCardGeneratedFromPaymentMethodDetailsCardPresent`, `PaymentAttemptRecordPaymentMethodDetailsCardPresent`, `PaymentMethodCardGeneratedFromPaymentMethodDetailsCardPresent`, and `PaymentRecordPaymentMethodDetailsCardPresent`
* Add support for `Location` and `Reader` on `ChargePaymentMethodDetailsCardPresent`, `ChargePaymentMethodDetailsInteracPresent`, `ConfirmationTokenPaymentMethodPreviewCardGeneratedFromPaymentMethodDetailsCardPresent`, `PaymentAttemptRecordPaymentMethodDetailsCardPresent`, `PaymentAttemptRecordPaymentMethodDetailsInteracPresent`, `PaymentMethodCardGeneratedFromPaymentMethodDetailsCardPresent`, `PaymentRecordPaymentMethodDetailsCardPresent`, and `PaymentRecordPaymentMethodDetailsInteracPresent`
* Add support for `ManagedPayments` on `CheckoutSessionParams`, `CheckoutSession`, `PaymentIntent`, `SetupIntent`, and `Subscription`
* Add support for new value `lk_vat` on enums `CheckoutSessionCollectedInformationTaxIds.Type`, `CheckoutSessionCustomerDetailsTaxIds.Type`, `OrderTaxDetailsTaxId.Type`, `QuotePreviewInvoiceCustomerTaxIds.Type`, `TaxCalculationCustomerDetailsTaxId.Type`, `TaxId.Type`, and `TaxTransactionCustomerDetailsTaxId.Type`
* Add support for `Digital` on `DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptions`, `DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionParams`, and `DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOption`
* Add support for `AffiliateAttributions` on `DelegatedCheckoutRequestedSessionConfirmParams`, `DelegatedCheckoutRequestedSessionParams`, and `DelegatedCheckoutRequestedSession`
* Add support for `FulfillmentType` on `DelegatedCheckoutRequestedSessionLineItemDetail`
* Add support for `MarketplaceSellerDetails`, `NetworkProfile`, `PrivacyNoticeURL`, `ReturnPolicyURL`, `StorePolicyURL`, and `TermsOfServiceURL` on `DelegatedCheckoutRequestedSessionSellerDetails`
* Add support for `AmountToCounter` on `DisputeParams`
* Add support for new values `reserve.hold.created`, `reserve.hold.updated`, `reserve.plan.created`, `reserve.plan.disabled`, `reserve.plan.expired`, `reserve.plan.updated`, and `reserve.release.created` on enum `Event.Type`
* Add support for new values `terminal_wifi_certificate` and `terminal_wifi_private_key` on enum `File.Purpose`
* Add support for new value `pay_by_bank` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
* Add support for `DisplayName` and `ServiceUserNumber` on `MandatePaymentMethodDetailsBacsDebit`
* Add support for `RequestReauthorization` on `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentConfirmPaymentMethodOptionsCardPresentParams`, `PaymentIntentPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCardPresentParams`, `PaymentIntentPaymentMethodOptionsCardPresent`, and `PaymentIntentPaymentMethodOptionsCard`
* Add support for `TransactionPurpose` on `PaymentIntentConfirmPaymentMethodOptionsUsBankAccountParams`, `PaymentIntentPaymentMethodOptionsUsBankAccountParams`, and `PaymentIntentPaymentMethodOptionsUsBankAccount`
* Add support for new value `requires_reauthorization` on enum `PaymentIntent.Status`
* Add support for `OptionalItems` on `PaymentLinkParams`
* Add support for new value `billing_schedules_invalid` on enum `QuoteStatusDetailsStaleLastReason.Type`
* ⚠️ Remove support for `CardIssuerDecline` on `RadarPaymentEvaluationInsights`
* Add support for `PaymentBehavior` on `SubscriptionItemParams`
* Add support for `BillingCycleAnchor` on `SubscriptionTrialSettingsEndBehavior`
* Add support for `Lk` on `TaxRegistrationCountryOptionsParams` and `TaxRegistrationCountryOptions`
* Add support for `Cellular` and `StripeS710` on `TerminalConfigurationParams` and `TerminalConfiguration`
* Add support for new values `simulated_stripe_s710` and `stripe_s710` on enum `TerminalReader.DeviceType`
* Add support for new values `ar_bank_account`, `bt_bank_account`, `co_bank_account`, `cr_bank_account`, `do_bank_account`, `gt_bank_account`, `md_bank_account`, `mk_bank_account`, `mo_bank_account`, `mz_bank_account`, `pe_bank_account`, `pk_bank_account`, `tw_bank_account`, and `uz_bank_account` on enums `V2AccountConfigurationRecipientDataDefaultOutboundDestination.Type` and `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
* Add support for `RecipientOnboarding` and `RecipientUpdate` on `V2CoreAccountLinkUseCaseParams` and `V2CoreAccountLinkUseCase`
* Add support for new values `recipient_onboarding` and `recipient_update` on enum `V2CoreAccountLinkUseCase.Type`
* Add support for `Consumer` on `V2CoreAccountConfigurationStorerCapabilitiesParams` and `V2CoreAccountConfigurationStorerCapabilities`
* Add support for new value `consumer.holds_currencies.usd` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
* Add support for `FundsUsageType` on `V2MoneyManagementFinancialAccountStorageParams` and `V2MoneyManagementFinancialAccountStorage`
* Add support for `Purpose` on `V2MoneyManagementOutboundPaymentParams` and `V2MoneyManagementOutboundPayment`
* Add support for `BranchNumber` and `SwiftCode` on `V2MoneyManagementPayoutMethodBankAccount`
* Add support for new values `dispute`, `inbound_payment_failure`, `inbound_payment`, `india_mdr_processing_fee`, `payment_method_passthrough_fee`, `refund`, and `tax_withholding` on enums `V2MoneyManagementTransaction.Category` and `V2MoneyManagementTransactionEntryTransactionDetails.Category`
* ⚠️ Remove support for values `charge_failure` and `charge` from enums `V2MoneyManagementTransaction.Category` and `V2MoneyManagementTransactionEntryTransactionDetails.Category`
* Add support for new value `consumer.holds_currencies.usd` on enum `EventsV2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent.UpdatedCapability`
* Add support for snapshot event `EventTypeBillingAlertRecovered` with resource `BillingAlertRecovered`
* Add support for snapshot events `EventTypeReserveHoldCreated` and `EventTypeReserveHoldUpdated` with resource `ReserveHold`
* Add support for snapshot events `EventTypeReservePlanCreated`, `EventTypeReservePlanDisabled`, `EventTypeReservePlanExpired`, and `EventTypeReservePlanUpdated` with resource `ReservePlan`
* Add support for snapshot event `EventTypeReserveReleaseCreated` with resource `ReserveRelease`
* Add support for event notification `V2BillingRateCardCustomPricingUnitOverageRateCreatedEvent` with related object `V2BillingRateCardCustomPricingUnitOverageRate`
* Add support for event notifications `V2IamStripeAccessGrantApprovedEvent`, `V2IamStripeAccessGrantCanceledEvent`, `V2IamStripeAccessGrantDeniedEvent`, `V2IamStripeAccessGrantRemovedEvent`, `V2IamStripeAccessGrantRequestedEvent`, and `V2IamStripeAccessGrantUpdatedEvent`
* Add support for error codes `storer_capability_missing` and `storer_capability_not_active` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `QuotePreviewInvoiceLastFinalizationError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`
