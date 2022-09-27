# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  module PriceHelpers
    extend T::Sig

    include Kernel
    include StripeForce::Constants
    extend SimpleStructuredLogger

    sig { params(user: StripeForce::User, subscription_schedule: Stripe::SubscriptionSchedule).void }
    def self.auto_archive_prices_on_subscription_schedule(user, subscription_schedule)
      log.info 'attempting to auto archive stripe prices'

      expanded_subscription_schedule = Stripe::SubscriptionSchedule.retrieve({
        id: subscription_schedule.id,
        expand: %w{
          phases.items.price
          phases.add_invoice_items.price
        },
      }, user.stripe_credentials)

      prices_on_subscription = OrderHelpers.extract_all_items_from_subscription_schedule(expanded_subscription_schedule).map(&:price)
      prices_on_subscription = T.cast(prices_on_subscription, T::Array[Stripe::Price])

      prices_on_subscription.select(&:active).each do |active_price|
        if active_price.metadata[Metadata.metadata_key(user, "auto_archive")]
          log.info 'archiving price', archive_price_id: active_price.id
          active_price.active = false
          active_price.save
        end
      end
    end

    sig { params(raw_billing_frequency: T.nilable(String)).returns(Integer) }
    def self.transform_salesforce_billing_frequency_to_recurring_interval(raw_billing_frequency)
      raw_billing_frequency ||= begin
        log.warn 'interval_count not defined via mapping, using monthly fallback'
        CPQBillingFrequencyOptions::MONTHLY.serialize
      end

      # convert picklist description of frequency to month integers
      case CPQBillingFrequencyOptions.try_deserialize(raw_billing_frequency)
      when CPQBillingFrequencyOptions::MONTHLY
        1
      when CPQBillingFrequencyOptions::QUARTERLY
        3
      when CPQBillingFrequencyOptions::SEMIANNUAL
        6
      when CPQBillingFrequencyOptions::ANNUAL
        12
      else
        raise StripeForce::Errors::RawUserError.new("Unexpected billing frequency #{raw_billing_frequency}. Must use default CPQ billing frequencies.")
      end
    end

    sig { params(raw_consumption_schedule_type: String).returns(String) }
    def self.transform_salesforce_consumption_schedule_type_to_tier_mode(raw_consumption_schedule_type)
      case raw_consumption_schedule_type
      when "Range"
        'volume'
      when "Slab"
        'graduated'
      else
        # should never happen
        raise StripeForce::Errors::RawUserError.new("unexpected consumption schedule type #{raw_consumption_schedule_type}")
      end
    end

    sig { params(stripe_price: Stripe::Price).returns(T::Boolean) }
    def self.tiered_price?(stripe_price)
      stripe_price.billing_scheme == "tiered"
    end

    sig { params(stripe_price: Stripe::Price).returns(T::Boolean) }
    def self.recurring_price?(stripe_price)
      stripe_price.type != "one_time"
    end

    sig { params(original_price_1: Stripe::Price, original_price_2: Stripe::Price).returns(T::Boolean) }
    def self.pricing_tiers_equal?(original_price_1, original_price_2)
      # to-be-created object:
      #   {"flat_amount":null,"flat_amount_decimal":null,"unit_amount":3000,"unit_amount_decimal":"3000","up_to":null}
      #
      # existing stripe object:
      #   {"up_to":"inf","unit_amount_decimal":"3000.0"}

      # problems:
      #   1. Tiers could be in different order
      #   2. Tiers from stripe could have null fields
      #   3. Tiers from stripe use `null` for `inf`

      # TODO we will probably need to do some sort of normalization here on tiering amounts

      # maybe the values are both nil, this should satisfy this condition
      if original_price_1[:tiers] == original_price_2[:tiers]
        return true
      end

      price_1 = Integrations::Utilities::StripeUtil.deep_copy(original_price_1)
      price_1_tiers = price_1.tiers.map {|t| Integrations::Utilities::StripeUtil.delete_nil_fields_from_stripe_object(t) }

      price_2 = Integrations::Utilities::StripeUtil.deep_copy(original_price_2)
      price_2_tiers = price_2.tiers.map {|t| Integrations::Utilities::StripeUtil.delete_nil_fields_from_stripe_object(t) }

      # normalize pricing field format between price API response and our generated price structure
      (price_1_tiers + price_2_tiers).each do |tier|
        # this is a subhash item and sorbet doesn't have these fields typed
        tier = T.unsafe(tier)

        # flat_amount_decimal should take precedence
        if tier[:flat_amount].present? && tier[:flat_amount_decimal].present?
          Integrations::Utilities::StripeUtil.delete_field_from_stripe_object(
            tier,
            :flat_amount
          )
        end

        # unit_amount_decimal should take precedence
        if tier[:unit_amount].present? && tier[:unit_amount_decimal].present?
          Integrations::Utilities::StripeUtil.delete_field_from_stripe_object(
            tier,
            :unit_amount
          )
        end

        tier[:unit_amount_decimal] = normalize_unit_amount_decimal_for_comparison(tier[:unit_amount_decimal])
        tier[:flat_amount_decimal] = normalize_unit_amount_decimal_for_comparison(tier[:flat_amount_decimal])

        # convert infinity format to API input, not output format
        if tier.up_to.nil?
          tier.up_to = 'inf'
        end
      end

      # https://jira.corp.stripe.com/browse/PLATINT-1817
      # tier sorting is not gaurenteed upstream, sort to ensure ordering does not cause inequality
      tier_sort = proc do |a, b|
        if a[:up_to] == 'inf' && b[:up_to] == 'inf'
          raise Integrations::Errors::ImpossibleState.new("two tiers should never both have an infinite boundary")
        end

        if a[:up_to] == 'inf'
          next 1
        end

        if b[:up_to] == 'inf'
          next -1
        end

        a[:up_to] <=> b[:up_to]
      end

      price_1_tiers.sort!(&tier_sort)
      price_2_tiers.sort!(&tier_sort)

      if price_1_tiers != price_2_tiers
        log.info 'tiers are not equal', diff: HashDiff::Comparison.new(price_1_tiers.map(&:to_hash), price_2_tiers.map(&:to_hash)).diff
        false
      else
        true
      end
    end

    sig { params(stripe_price: Stripe::Price).returns(Stripe::Price) }
    def self.sanitize_price_tier_params(stripe_price)
      # no side effects, please!
      stripe_price = Integrations::Utilities::StripeUtil.deep_copy(stripe_price)

      # if non-tiered pricing fields are included Stripe will throw an error
      Integrations::Utilities::StripeUtil.delete_field_from_stripe_object(stripe_price, :unit_amount_decimal)

      # TODO are there other pricing fields we should delete? It's unclear what the requirements are here?

      # if the field types are not the same, you will get the following error:
      # Caused by ArgumentError: comparison of Stripe::StripeObject with Stripe::StripeObject failed

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
        log.warn 'usage type not defined, defaulting to advance'
        CPQProductBillingTypeOptions::ADVANCE.serialize
      end

      case CPQProductBillingTypeOptions.try_deserialize(raw_usage_type)
      when CPQProductBillingTypeOptions::ADVANCE
        'licensed'
      when CPQProductBillingTypeOptions::ARREARS
        'metered'
      else
        raise StripeForce::Errors::RawUserError.new("unexpected billing type #{raw_usage_type}")
      end
    end

    sig { params(unit_amount_decimal: T.nilable(T.any(String, BigDecimal))).returns(T.nilable(BigDecimal)) }
    def self.normalize_unit_amount_decimal_for_comparison(unit_amount_decimal)
      return nil if unit_amount_decimal.nil?

      if !unit_amount_decimal.is_a?(BigDecimal)
        unit_amount_decimal = BigDecimal(unit_amount_decimal)
      end

      unit_amount_decimal.round(MAX_STRIPE_PRICE_PRECISION)
    end

    sig { params(price_1: Stripe::Price, price_2: Stripe::Price).returns(T::Boolean) }
    def self.price_billing_amounts_equal?(price_1, price_2)
      # metadata *could* have important financial data for downstream systems
      # however, most of the time users will not pull this information directly from prices
      # (they may use products instead) and a users mappings could change over time, so we don't
      # want to factor it in to our equality test. If we did, prices would never match and it would
      # create a massive amount of price objects in their Stripe account.

      # these values could be nil, which is fine as long as they are both nil
      # if they are not both nil (comparing an existing stripe price to a to-be-created price)
      # we need to relax the checks a bit: the comparisons should below should infer the correct
      # pricing types (metered vs licensed)

      fields_to_check_if_both_are_set = %i{
        product
        tax_behavior
        billing_scheme
        type
        currency
      }

      simple_field_check_passed = fields_to_check_if_both_are_set.all? do |field_sym|
        # if one of the fields is set, skip the comparison
        price_1[field_sym].blank? != price_2[field_sym].blank? ||
          price_1[field_sym] == price_2[field_sym]
      end

      if !simple_field_check_passed
        log.info 'price not equal, simple field comparison failed', diff: HashDiff::Comparison.new(price_1.to_hash, price_2.to_hash).diff
        return false
      end

      # billing_scheme is important in price comparison because it indicates to us *how* prices should be compared:
      # i.e. should the tiers be compared? Should the top-level unit_amount_decimal be compared? etc. We use the billing_scheme
      # field to branch the comparison logic below. However, this field may not always be set when comparing an existing Stripe price object
      # with the generated params from the connector. The connector does *not* set the billing_scheme if not mapped explicitly
      # so the stripe default value is used instead. Currently (09/22) the default billing_scheme value is `per_unit` which is why
      # we assume that value below if the `billing_scheme` field is nil.

      price_1_billing_scheme = if price_1[:billing_scheme].nil?
        'per_unit'
      else
        price_1.billing_scheme
      end

      price_2_billing_scheme = if price_2[:billing_scheme].nil?
        'per_unit'
      else
        price_2.billing_scheme
      end

      if price_1_billing_scheme != price_2_billing_scheme
        log.info 'price not equal, billing scheme comparison failed'
        return false
      end

      is_price_equal = if price_1_billing_scheme == 'per_unit'
        # TODO we do not expect this occur, if it does we'll need to improve the error message here
        if price_1.unit_amount_decimal.nil? || price_2.unit_amount_decimal.nil?
          raise StripeForce::Errors::RawUserError.new("unit_amount_decimal nil on price objects")
        end

        price_1_decimal = normalize_unit_amount_decimal_for_comparison(price_1.unit_amount_decimal)
        price_2_decimal = normalize_unit_amount_decimal_for_comparison(price_2.unit_amount_decimal)

        price_1_decimal == price_2_decimal
      elsif price_1_billing_scheme == 'tiered'
        # TODO probably need to think through transformed_quantity here and if it could effect this
        price_1[:tiers_mode] == price_2[:tiers_mode] && PriceHelpers.pricing_tiers_equal?(price_1, price_2)
      else
        # stripe currently (10/22) only allows for two values here (tiered, per_unit), throw an error in case the API shape changes in the future
        raise StripeForce::Errors::RawUserError.new("Unexpected billing_scheme on price encountered #{price_1.billing_scheme}")
      end

      if !is_price_equal
        log.info 'price comparison is not equal', diff: HashDiff::Comparison.new(price_1.to_hash, price_2.to_hash).diff
        return false
      end

      if (price_1.recurring.present? && !price_1.recurring.to_hash.empty?) != (price_2.recurring.present? && !price_2.recurring.to_hash.empty?)
        log.info 'price not equal, recurring is not consistent', diff: HashDiff::Comparison.new(price_1.to_hash, price_2.to_hash).diff
        return false
      end

      has_recurring_price = price_1.recurring.present? && !price_1.recurring.to_hash.empty?

      if has_recurring_price
        recurring_billing_amounts_equal = price_1.recurring[:interval] == price_2.recurring[:interval] &&
          price_1.recurring[:interval_count] == price_2.recurring[:interval_count]

        if !recurring_billing_amounts_equal
          log.info 'price not equal, recurring params not the same', diff: HashDiff::Comparison.new(price_1.to_hash, price_2.to_hash).diff
        end

        recurring_billing_amounts_equal
      else
        true
      end
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
