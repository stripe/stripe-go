---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2289
is_breaking: true
is_stripe_api_change: true
released_in_version: 84.5.0-alpha.3
---

* Add support for new resource `RadarIssuingAuthorizationEvaluation`
* Add support for `New` method on resource `RadarIssuingAuthorizationEvaluation`
* Add support for new value `fee_credits` on enum `BalanceTransaction.BalanceType`
* ⚠️ Rename `AffiliateAttributions` to `AffiliateAttribution` on `DelegatedCheckoutRequestedSessionConfirmParams` and `DelegatedCheckoutRequestedSessionParams`
* Add support for `AmountToCounter` on `Dispute`
* Add support for `FrozenFields` on `InvoiceItem`
* Add support for new value `next_billing_period_start` on enum `V2BillingIntentActionApplyEffectiveAt.Type`
* Add support for `Consumer` on `V2CoreAccountConfigurationCardCreatorCapabilitiesParams`, `V2CoreAccountConfigurationCardCreatorCapabilities`, `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorParams`, and `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreator`
* Add support for `FifthThird` on `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercial`, `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialParams`, and `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercial`
* Add support for `PrepaidCard` on `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBank`, `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankParams`, and `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBank`
* Add support for new values `commercial.cross_river_bank.prepaid_card`, `commercial.fifth_third.charge_card`, `consumer.celtic.revolving_credit_card`, `consumer.cross_river_bank.prepaid_card`, and `consumer.lead.prepaid_card` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
* Add support for `PaymentMethodData` on `V2PaymentsOffSessionPaymentParams`
* Add support for new values `commercial.cross_river_bank.prepaid_card`, `commercial.fifth_third.charge_card`, `consumer.celtic.revolving_credit_card`, `consumer.cross_river_bank.prepaid_card`, and `consumer.lead.prepaid_card` on enum `EventsV2CoreAccountIncludingConfigurationCardCreatorCapabilityStatusUpdatedEvent.UpdatedCapability`
