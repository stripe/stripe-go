# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  module PriceHelpers
    extend T::Sig

    include Kernel
    include StripeForce::Constants
    extend SimpleStructuredLogger

    sig { params(raw_consumption_schedule_type: String).returns(String) }
    def self.transform_salesforce_consumption_schedule_type_to_tier_mode(raw_consumption_schedule_type)
      case raw_consumption_schedule_type
      when "Range"
        'volume'
      when "Slab"
        'graduated'
      else
        # should never happen
        raise "unexpected consumption schedule type #{raw_consumption_schedule_type}"
      end
    end

    sig { params(stripe_price: Stripe::Price).returns(Stripe::Price) }
    def self.sanitize_price_tier_params(stripe_price)
      # no side effects, please!
      stripe_price = stripe_price.dup

      # if non-tiered pricing fields are included Stripe will throw an error
      Integrations::Utilities::StripeUtil.delete_field_from_stripe_object(stripe_price, :unit_amount_decimal)

      # TODO are there other pricing fields we should delete? It's unclear what the requirements are here?

      # API also requires pricing tiers to be sorted. Wat?
      # TODO https://jira.corp.stripe.com/browse/PLATINT-1573
      stripe_price.tiers.sort! do |a, b|
        # TODO can we use null instead of `inf`
        # a & b should never both be `inf`; this should be checked upstream
        if a.up_to == 'inf'
          1
        elsif b.up_to == 'inf'
          -1
        else
          a.up_to <=> b.up_to
        end
      end

      stripe_price
    end

    sig { params(raw_usage_type: T.nilable(String)).returns(String) }
    def self.transform_salesforce_billing_type_to_usage_type(raw_usage_type)
      # # https://help.salesforce.com/s/articleView?id=sf.cpq_order_product_fields.htm&type=5
      raw_usage_type ||= begin
        log.warn 'usage type not defined, defaulting to advance (licensed)'
        CPQProductBillingTypeOptions::ADVANCE.serialize
      end

      case CPQProductBillingTypeOptions.try_deserialize(raw_usage_type)
      when CPQProductBillingTypeOptions::ADVANCE
        'licensed'
      when CPQProductBillingTypeOptions::ARREARS
        'metered'
      else
        raise "unexpected billing type"
      end
    end


    # right now, this is NOT used everywhere there is a price comparison, we should leverage this more broadly across the codebase
    sig { params(price_1: Stripe::Price, price_2: Stripe::Price).returns(T::Boolean) }
    def self.price_billing_amounts_equal?(price_1, price_2)
      billing_amounts_equal = BigDecimal(price_1.unit_amount_decimal.to_s) != BigDecimal(price_2.unit_amount_decimal.to_s) ||
        price_1.recurring.present? != price_2.recurring.present? ||
        price_1.recurring.interval != price_2.recurring[:interval] ||
        price_1.recurring.interval_count != price_2.recurring[:interval_count]

      if !billing_amounts_equal
        log.info 'price not equal', diff: HashDiff::Comparison.new(price_1.to_hash, price_2.to_hash).diff
      end

      billing_amounts_equal
    end

    # TODO this is very naive: we need a better way of determining if the price field was customized
    # TODO we'll probably need some sort of feature flag for this as well
    sig { params(user: StripeForce::User).returns(T::Boolean) }
    def self.using_custom_order_line_price_field?(user)
      # NOTE this is hacky! This is all to ensure `sobject_type` is filled in properly
      sf_line_item = Restforce::SObject.new({
        "attributes" => {
          "type" => SF_ORDER_ITEM,
        },
      })

      order_line_price_key = StripeForce::Mapper.mapping_key_for_record(Stripe::Price, sf_line_item)
      default_mapping = user.required_mappings.dig(order_line_price_key, 'unit_amount_decimal')

      !user.field_defaults.dig(order_line_price_key, 'unit_amount_decimal').nil? ||
        (
          !user.field_mappings.dig(order_line_price_key, 'unit_amount_decimal').nil? &&
          user.field_mappings.dig(order_line_price_key, 'unit_amount_decimal') != default_mapping
        )
    end

    sig { params(stripe_price: Stripe::Price).returns(T::Boolean) }
    def self.metered_price?(stripe_price)
      stripe_price.recurring&.usage_type == 'metered'
    end
  end
end
