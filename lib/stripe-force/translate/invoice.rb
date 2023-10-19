# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  sig { params(stripe_customer: Stripe::Customer, invoice_items: T::Array[ContractItemStructure], sf_order: Restforce::SObject).returns(Stripe::Invoice) }
  def create_one_time_stripe_invoice_from_order(stripe_customer, invoice_items, sf_order)
    log.info 'no recurring items found, creating a one-time invoice', salesforce_object_id: sf_order.Id

    order_start_date = StripeForce::Utilities::SalesforceUtil.extract_subscription_start_date_from_order(mapper, sf_order)
    subscription_term_from_salesforce = StripeForce::Utilities::SalesforceUtil.extract_subscription_term_from_order!(mapper, sf_order)
    order_end_date = StripeForce::Utilities::SalesforceUtil.datetime_to_unix_timestamp(order_start_date + subscription_term_from_salesforce.months)
    period_start = order_start_date.to_i
    period_end = order_end_date.to_i

    if @user.feature_enabled?(StripeForce::Constants::FeatureFlags::BILLING_GATE_OPEN_INVOICING_INTERVAL)
      # https://livegrep.corp.stripe.com/view/stripe-internal/pay-server/lib/subscriptions/command/invoicing_period.rb#L26
      period_end = period_end - 1 > period_start ? period_end - 1 : period_end
    end

    # TODO there has got to be a way to include the lines on the invoice item create call, report this to stripe billing
    invoice_items.each do |invoice_item|
      log.info 'adding invoice item for order line', order_line_id: invoice_item.order_line

      Stripe::InvoiceItem.create({
        customer: stripe_customer,
        period: {
          start: period_start,
          end: period_end,
        },
      }.merge(invoice_item.stripe_params),
        @user.stripe_credentials
      )
    end

    stripe_invoice, response = create_stripe_object(
      Stripe::Invoice, sf_order,
      additional_stripe_params: {
        customer: stripe_customer.id,
      }
    )
    stripe_invoice.finalize_invoice({}, @user.stripe_credentials)

    log.info 'stripe invoice created for one-time services', stripe_resource_id: stripe_invoice.id

    update_sf_stripe_id(
      sf_order, stripe_invoice,
      stripe_response: response,
      additional_salesforce_updates: {
        prefixed_stripe_field(ORDER_INVOICE_PAYMENT_LINK) => stripe_invoice.hosted_invoice_url,
      }
    )

    stripe_invoice
  end
end
