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
        return
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

      # at this point, we know this is a valid invoice event
      # let's ensure this is not a duplicate event caused by a sub phase split
      if self.is_duplicate_invoice_item(user, subscription_id, invoice_item)
        log.info 'duplicate invoice item created, skipping', invoice_item_id: invoice_item_id
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

    def self.is_duplicate_invoice_item(user, subscription_id, invoice_item)
      subscription = Stripe::Subscription.retrieve(T.must(subscription_id), user.stripe_credentials)

      # this can't be a duplicate invoice event if this subscription has never been billed
      if subscription.latest_invoice.present?
        invoice_items_list = Stripe::InvoiceItem.list({customer: subscription.customer, created: {gt: subscription.start_date}}, user.stripe_credentials)
        invoice_items_list.auto_paging_each do |existing_invoice_item|
          if existing_invoice_item.id != invoice_item.id && # not the same event
              existing_invoice_item.subscription == subscription_id && # same subscription
              existing_invoice_item.metadata[Translate::Metadata.metadata_key(user, MetadataKeys::PRORATION)] && # created by the connector
              self.invoice_items_are_equal?(existing_invoice_item: existing_invoice_item, new_invoice_item: invoice_item)
            log.info 'duplicate invoice item detected on previous invoice', original_invoice_item: existing_invoice_item
            return true
          end
        end
      end

      false
    end

    def self.invoice_items_are_equal?(existing_invoice_item:, new_invoice_item:)
      # https://stripe.com/docs/api/invoiceitems/object
      fields_to_check_if_both_are_equal = %i{
        amount
        currency
        description
        metadata
        period
        price
        proration
        livemode
        quantity
        unit_amount
        unit_amount_decimal
        tax_rates
      }

      simple_field_check_passed = fields_to_check_if_both_are_equal.all? do |field_sym|
        existing_invoice_item[field_sym].blank? == new_invoice_item[field_sym].blank? && existing_invoice_item[field_sym].to_s == new_invoice_item[field_sym].to_s
      end

      simple_field_check_passed
    end
  end
end
