# frozen_string_literal: true
# typed: true

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

      if subscription_id.blank?
        log.info 'no subscription reference, not billing'
        nil
      end

      if invoice_item.period.start == invoice_item.period.end
        log.info 'period start and end are equal, not proration'
        return
      end

      if invoice_item.metadata[StripeForce::Utilities::Metadata.metadata_key(user, MetadataKeys::PRORATION)] != "true"
        log.info 'does not contain proration metadata, skipping'
        return
      end

      invoice_item = Stripe::InvoiceItem.retrieve(invoice_item.id, user.stripe_credentials)

      if invoice_item.invoice.present?
        log.info 'already invoiced'
        return
      end

      invoice = Stripe::Invoice.create({
        customer: customer_id,
        subscription: subscription_id,
        metadata: {
          # TODO should we link to the originating order? Can we extract that from the line item metadata when we have it?
          StripeForce::Utilities::Metadata.metadata_key(user, MetadataKeys::PRORATION_INVOICE) => "true",
        },
      }, user.stripe_credentials)

      log.info 'invoice created for proration', invoice_id: invoice.id

      invoice
    end

  end
end
