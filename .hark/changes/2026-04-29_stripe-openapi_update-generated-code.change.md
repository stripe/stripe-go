---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2349
is_stripe_api_change: true
released_in_version: 85.2.0-alpha.3
---

* Add support for `DebitCard` on `V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLeadParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLead`, `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadParams`, and `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLead`
* Add support for new value `consumer.lead.debit_card` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
* Add support for new value `consumer.lead.debit_card` on enum `EventsV2CoreAccountIncludingConfigurationCardCreatorCapabilityStatusUpdatedEvent.UpdatedCapability`
