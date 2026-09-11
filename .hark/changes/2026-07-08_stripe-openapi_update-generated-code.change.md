---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2383
is_breaking: true
is_stripe_api_change: true
released_in_version: 86.2.0-alpha.3
---

* Add support for `ActivateGiftCard`, `CashoutGiftCard`, `CheckGiftCardBalance`, and `ReloadGiftCard` methods on resource `TerminalReader`
* Add support for `AggregationPeriod` on `BillingAlertRecovered`
* Add support for new values `mass_transit_parking_tax` and `parking_tax` on enums `TaxCalculationLineItemTaxBreakdownTaxRateDetails.TaxType`, `TaxCalculationShippingCostTaxBreakdownTaxRateDetails.TaxType`, `TaxCalculationTaxBreakdownTaxRateDetails.TaxType`, and `TaxTransactionShippingCostTaxBreakdownTaxRateDetails.TaxType`
* Add support for `AdministrativeAddress` and `PrincipalPlaceOfBusiness` on `AccountCompany`
* Add support for `AddressCollectionPrecision` on `CheckoutSessionAutomaticTax`
* Add support for `TaxID` on `CheckoutSessionCollectedInformation`
* ⚠️ Remove support for `TaxIDs` on `CheckoutSessionCollectedInformation`
* Add support for `SetupFutureUsage` on `CheckoutSessionPaymentMethodOptionsPayco`, `CheckoutSessionPaymentMethodOptionsSamsungPay`, `PaymentIntentConfirmPaymentMethodOptionsPaypayParams`, `PaymentIntentPaymentMethodOptionsPayco`, `PaymentIntentPaymentMethodOptionsPaypayParams`, `PaymentIntentPaymentMethodOptionsPaypay`, and `PaymentIntentPaymentMethodOptionsSamsungPay`
* Add support for new values `bnp_paribas`, `citibank`, and `mbsb_bank` on enums `ConfirmationTokenPaymentMethodPreviewFpx.Bank`, `PaymentAttemptRecordPaymentMethodDetailsFpx.Bank`, `PaymentRecordPaymentMethodDetailsFpx.Bank`, and `SharedPaymentGrantedTokenPaymentMethodDetailsFpx.Bank`
* Add support for `Network` on `DisputePaymentMethodDetailsCard`
* Add support for `RequirePaymentMethodSupport` on `FinancialConnectionsSessionFilters`
* Add support for `NetworkData` on `IssuingAuthorizationRequestHistory`
* Add support for `AcquiringInstitutionCountry`, `AcquiringInstitutionID`, `RetrievalReferenceNumber`, `RoutedNetwork`, and `TraceID` on `IssuingTransactionNetworkData`
* Add support for `CustomFields`, `Description`, and `Footer` on `QuoteInvoiceSettings`, `QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettings`, `QuotePreviewSubscriptionSchedulePhaseInvoiceSettings`, `SubscriptionScheduleDefaultSettingsInvoiceSettings`, and `SubscriptionSchedulePhaseInvoiceSettings`
* Add support for `Paypay` on `SetupAttemptPaymentMethodDetails`
* Add support for `MassTransitParkingTax` and `ParkingTax` on `TaxRegistrationCountryOptionsUs`
* Add support for new values `mass_transit_parking_tax` and `parking_tax` on enum `TaxRegistrationCountryOptionsUs.Type`
* Add support for `GiftCardBrand` on `TerminalReaderCollectPaymentMethodCollectConfigParams` and `TerminalReaderProcessPaymentIntentProcessConfigParams`
* Add support for `ActivateGiftCard`, `CashoutGiftCard`, `CheckGiftCardBalance`, `DeactivateGiftCard`, and `ReloadGiftCard` on `TerminalReaderAction`
* Add support for new values `activate_gift_card`, `cashout_gift_card`, `check_gift_card_balance`, `deactivate_gift_card`, and `reload_gift_card` on enum `TerminalReaderAction.Type`
* Add support for `StatusTransitions` on `V2BillingContract`
* ⚠️ Remove support for `OneTimeFees` on `V2BillingContractParams` and `V2BillingContract`
* ⚠️ Remove support for `StatusDetails` on `V2BillingContract`
* Add support for `ID` and `Priority` on `V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesData`
* ⚠️ Remove support for `PricingOverride` on `V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesData`
* ⚠️ Remove support for `TieringMode` and `Tiers` on `V2BillingContractPricingLineActionAddPricingPriceDetailsPricingOverrideOverwritePriceParams`, `V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddOverwritePriceParams`, `V2BillingContractPricingLinePricingPriceDetailsPricingOverrideOverwritePriceParams`, `V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataOverwritePrice`, and `V2BillingContractPricingOverrideActionAddOverwritePriceParams`
* Add support for `MultiplyPricing` on `V2BillingContractPricingOverrideActionAddParams`, `V2BillingContractPricingOverrideParams`, and `V2BillingContractPricingOverridesData`
* ⚠️ Remove support for `Multiplier` on `V2BillingContractPricingOverrideActionAddParams`, `V2BillingContractPricingOverrideParams`, and `V2BillingContractPricingOverridesData`
* ⚠️ Change type of `V2BillingContractPricingOverrideActionAddParams.Type`, `V2BillingContractPricingOverrideParams.Type`, and `V2BillingContractPricingOverridesData.Type` from `literal('multiplier')` to `literal('multiply_pricing')`
* Add support for `RelatedNetworkObject` on `V2CoreAccountListParams` and `V2CoreAccount`
* Add support for new value `network_business_profile_wallet` on enum `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
* Add support for `NetworkBusinessProfileWallet` on `V2MoneyManagementPayoutMethod`
* Add support for new value `network_business_profile_wallet` on enum `V2MoneyManagementPayoutMethod.Type`
* Add support for `StripeNetworkTransfer` on `V2MoneyManagementReceivedCredit`
* Add support for new value `stripe_network_transfer` on enum `V2MoneyManagementReceivedCredit.Type`
* ⚠️ Change type of `V2BillingContractPricingLineActionAddPricingPriceDetailsPricingOverrideEndsAtParams.Type`, `V2BillingContractPricingLineEndsAtParams.Type`, `V2BillingContractPricingLinePricingPriceDetailsPricingOverrideEndsAtParams.Type`, and `V2BillingContractPricingOverrideEndsAtParams.Type` from `enum('contract_end'|'timestamp')` to `literal('timestamp')`
* ⚠️ Change type of `V2BillingContractPricingLineActionAddPricingPriceDetailsPricingOverrideStartsAtParams.Type`, `V2BillingContractPricingLinePricingPriceDetailsPricingOverrideStartsAtParams.Type`, `V2BillingContractPricingLineStartsAtParams.Type`, and `V2BillingContractPricingOverrideStartsAtParams.Type` from `enum('contract_start'|'timestamp')` to `literal('timestamp')`
* ⚠️ Change type of `V2BillingContractPricingLineActionAddEndsAtParams.Type`, `V2BillingContractPricingLineActionUpdateEndsAtParams.Type`, `V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddEndsAtParams.Type`, `V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionUpdateEndsAtParams.Type`, `V2BillingContractPricingOverrideActionAddEndsAtParams.Type`, and `V2BillingContractPricingOverrideActionUpdateEndsAtParams.Type` from `enum('billing_period_end'|'timestamp')` to `literal('timestamp')`
* ⚠️ Change type of `V2BillingContractPricingLineActionAddStartsAtParams.Type`, `V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddStartsAtParams.Type`, `V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionUpdateStartsAtParams.Type`, `V2BillingContractPricingLineActionUpdateStartsAtParams.Type`, `V2BillingContractPricingOverrideActionAddStartsAtParams.Type`, and `V2BillingContractPricingOverrideActionUpdateStartsAtParams.Type` from `enum('billing_period_start'|'timestamp')` to `literal('timestamp')`
* Add support for event notifications `V2BillingContractActivatedEvent`, `V2BillingContractCanceledEvent`, `V2BillingContractCreatedEvent`, `V2BillingContractEndedEvent`, and `V2BillingContractUpdatedEvent` with related object `V2BillingContract`
