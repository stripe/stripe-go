---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2233
is_stripe_api_change: true
released_in_version: 84.2.0-alpha.1
---

* Add support for new resources `SharedPaymentGrantedToken`, `V2IamAPIKey`, `V2PaymentsSettlementAllocationIntentSplit`, `V2PaymentsSettlementAllocationIntent`, and `V2TaxManualRule`
* Add support for `Get` method on resource `SharedPaymentGrantedToken`
* Add support for `New` and `Update` test helper methods on resource `SharedPaymentGrantedToken`
* Add support for `Deactivate`, `Get`, `List`, `New`, and `Update` methods on resource `V2TaxManualRule`
* Add support for `Cancel`, `Get`, `New`, `Submit`, and `Update` methods on resource `V2PaymentsSettlementAllocationIntent`
* Add support for `Cancel`, `Get`, and `New` methods on resource `V2PaymentsSettlementAllocationIntentSplit`
* Add support for `Expire`, `Get`, `List`, `New`, `Rotate`, and `Update` methods on resource `V2IamAPIKey`
* Add support for `CheckScanning` on `AccountSessionComponentsParams`
* Add support for `TaxDetails` on `CheckoutSessionLineItemPriceDataProductDataParams`, `InvoiceAddLinesLinePriceDataProductDataParams`, `InvoiceLineItemPriceDataProductDataParams`, `InvoiceUpdateLinesLinePriceDataProductDataParams`, `PaymentLinkLineItemPriceDataProductDataParams`, and `ProductParams`
* Add support for `PaymentMethodData` on `DelegatedCheckoutRequestedSessionConfirmParams`
* Add support for `ProductDetails` on `DelegatedCheckoutRequestedSessionLineItemDetail`
* Add support for `Wallets` on `IssuingCardListParams`
* Add support for `PrimaryAccountIdentifier` on `IssuingCardWalletsApplePay` and `IssuingCardWalletsGooglePay`
* Add support for `SharedPaymentGrantedToken` on `PaymentIntentConfirmParams`, `PaymentIntentParams`, and `PaymentIntent`
* Add support for new values `al_bank_account`, `am_bank_account`, `bn_bank_account`, `bw_bank_account`, `dz_bank_account`, `gy_bank_account`, `jm_bank_account`, `jo_bank_account`, `kw_bank_account`, `lk_bank_account`, `ma_bank_account`, `om_bank_account`, and `tz_bank_account` on enum `V2AccountConfigurationRecipientDataDefaultOutboundDestination.Type`
* Add support for `Instant` on `V2AccountConfigurationRecipientDataFeaturesBankAccountsParams`, `V2AccountConfigurationRecipientDataFeaturesBankAccounts`, `V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsParams`, and `V2CoreAccountConfigurationRecipientCapabilitiesBankAccounts`
* Add support for new value `bank_accounts.instant` on enum `V2AccountRequirementImpact.RequiredForFeatures`
* Add support for `CollectAt` on `V2BillingIntentActionDeactivateParams`, `V2BillingIntentActionDeactivate`, `V2BillingIntentActionModifyParams`, `V2BillingIntentActionModify`, `V2BillingIntentActionSubscribeParams`, and `V2BillingIntentActionSubscribe`
* Remove support for `BillingDetails` on `V2BillingIntentActionDeactivateParams`, `V2BillingIntentActionDeactivate`, `V2BillingIntentActionModifyParams`, `V2BillingIntentActionModify`, `V2BillingIntentActionSubscribeParams`, and `V2BillingIntentActionSubscribe`
* Add support for `Overrides` on `V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsParams`, `V2BillingIntentActionDeactivatePricingPlanSubscriptionDetails`, `V2BillingIntentActionModifyPricingPlanSubscriptionDetailsParams`, `V2BillingIntentActionModifyPricingPlanSubscriptionDetails`, `V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsParams`, and `V2BillingIntentActionSubscribePricingPlanSubscriptionDetails`
* Remove support for `Requested` on `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticChargeCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticSpendCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripeChargeCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCard`, `V2CoreAccountConfigurationRecipientCapabilitiesCryptoWallets`, `V2CoreAccountConfigurationStorerCapabilitiesFinancialAddressesCryptoWallets`, `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesUsdc`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsCryptoWallets`, and `V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfersCryptoWallets`
* Add support for new value `bank_accounts.instant` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
* Add support for `AlternativeReference` on `V2CoreVaultGbBankAccount`, `V2CoreVaultUsBankAccount`, and `V2MoneyManagementPayoutMethod`
* Add support for `ManagedBy` and `Payments` on `V2MoneyManagementFinancialAccount`
* Add support for new value `payments` on enum `V2MoneyManagementFinancialAccount.Type`
* Add support for `Speed` on `V2MoneyManagementOutboundPaymentDeliveryOptionsParams`, `V2MoneyManagementOutboundPaymentDeliveryOptions`, `V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsParams`, and `V2MoneyManagementOutboundPaymentQuoteDeliveryOptions`
* Add support for new value `real_time_payout_fee` on enum `V2MoneyManagementOutboundPaymentQuoteEstimatedFee.Type`
* Add support for `Types` on `V2MoneyManagementFinancialAccountListParams`
* Add support for new value `bank_accounts.instant` on enum `EventsV2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent.UpdatedCapability`
* Add support for `TopImpactedAccounts` on `EventsV2CoreHealthApiErrorFiringEventImpact`, `EventsV2CoreHealthApiErrorResolvedEventImpact`, `EventsV2CoreHealthApiLatencyFiringEventImpact`, `EventsV2CoreHealthApiLatencyResolvedEventImpact`, `EventsV2CoreHealthPaymentMethodErrorFiringEventImpact`, and `EventsV2CoreHealthPaymentMethodErrorResolvedEventImpact`
* Add support for event notifications `V2CoreHealthSepaDebitDelayedFiringEvent`, `V2CoreHealthSepaDebitDelayedResolvedEvent`, and `V2PaymentsSettlementAllocationIntentNotFoundEvent`
* Add support for event notifications `V2PaymentsSettlementAllocationIntentCanceledEvent`, `V2PaymentsSettlementAllocationIntentCreatedEvent`, `V2PaymentsSettlementAllocationIntentErroredEvent`, `V2PaymentsSettlementAllocationIntentFundsNotReceivedEvent`, `V2PaymentsSettlementAllocationIntentMatchedEvent`, `V2PaymentsSettlementAllocationIntentSettledEvent`, and `V2PaymentsSettlementAllocationIntentSubmittedEvent` with related object `V2PaymentsSettlementAllocationIntent`
* Add support for event notifications `V2PaymentsSettlementAllocationIntentSplitCanceledEvent`, `V2PaymentsSettlementAllocationIntentSplitCreatedEvent`, and `V2PaymentsSettlementAllocationIntentSplitSettledEvent` with related object `V2PaymentsSettlementAllocationIntentSplit`
* Remove support for error code `account_rate_limit_exceeded` on `RateLimitError`
