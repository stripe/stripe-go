# frozen_string_literal: true
# typed: true

# Design doc:
# https://docs.google.com/document/d/1jE0vB0CWeww-rMMhDpPSE7RZbU8_6wj-6Q5BFTKfMRM/edit#heading=h.k69akfih61q
module StripeForce
  module ProrationAutoBill
    extend T::Sig

    include StripeForce::Constants
    include Integrations::Log

    sig { params(user: StripeForce::User, invoice_item_event: Stripe::Event).returns(T.nilable(Stripe::Invoice)) }
    def self.create_invoice_from_invoice_item_event(user, invoice_item_event)
      invoice_item = T.unsafe(invoice_item_event.data.object)
      subscription_id = T.cast(invoice_item.subscription, T.nilable(String))
      customer_id = T.cast(invoice_item.customer, String)

      # TODO should set error and log context
      invoice_item_id = invoice_item.id

      if subscription_id.blank?
        log.info 'no subscription reference, not billing', invoice_item_id: invoice_item_id
        nil
      end

      if invoice_item.period.start == invoice_item.period.end
        log.info 'period start and end are equal, not proration', invoice_item_id: invoice_item_id
        return
      end

      if invoice_item.metadata[Translate::Metadata.metadata_key(user, MetadataKeys::PRORATION)] != "true"
        log.info 'does not contain proration metadata, skipping', invoice_item_id: invoice_item_id
        return
      end

      invoice_item = Stripe::InvoiceItem.retrieve(invoice_item_id, user.stripe_credentials)

      if invoice_item.invoice.present?
        log.info 'already invoiced, skipping', invoice_item_id: invoice_item_id
        return
      end

      # TODO we need billing to offer the ability to pass in a list of invoice items,
      #      this logic could "scoop up" invoice items which the user does NOT intend to bill
      auto_advance = user.feature_enabled? FeatureFlags::AUTO_ADVANCE_PRORATION_INVOICE
      invoice = Stripe::Invoice.create({
        customer: customer_id,
        subscription: subscription_id,
        metadata: {
          # TODO should we link to the originating order? Can we extract that from the line item metadata when we have it?
          Translate::Metadata.metadata_key(user, MetadataKeys::PRORATION_INVOICE) => "true",
        },
        auto_advance: auto_advance,
      }, user.stripe_credentials)

      log.info 'invoice created for proration', invoice_id: invoice.id

      invoice
    end

  end
end
