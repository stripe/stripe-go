# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  # TODO should we protect this behind a flag?
  sig { params(stripe_customer: Stripe::Customer, invoice_items: T.untyped, sf_order: Restforce::SObject).returns(Stripe::Invoice) }
  def create_stripe_invoice_from_order(stripe_customer, invoice_items, sf_order)
    log.info 'no recurring items found, creating a one-time invoice'

    # TODO there has got to be a way to include the lines on the invoice item create call
    invoice_items.each do |invoice_item_params|
      # TODO idempotency_key
      Stripe::InvoiceItem.create(
        {customer: stripe_customer}.merge(invoice_item_params.stripe_params),
        @user.stripe_credentials
      )
    end

    stripe_invoice = create_stripe_object(Stripe::Invoice, sf_order,
      additional_stripe_params: {
        customer: stripe_customer.id,
      }
    )

    # TODO idempotency_key
    stripe_invoice.finalize_invoice

    log.info 'stripe invoice created', stripe_resource_id: stripe_invoice.id

    update_sf_stripe_id(sf_order, stripe_invoice, additional_salesforce_updates: {
      prefixed_stripe_field(ORDER_INVOICE_PAYMENT_LINK) => stripe_invoice.hosted_invoice_url,
    })

    if stripe_invoice.is_a?(Stripe::SubscriptionSchedule)
      PriceHelpers.auto_archive_prices_on_subscription_schedule(@user, stripe_invoice)
    end

    stripe_invoice
  end
end
