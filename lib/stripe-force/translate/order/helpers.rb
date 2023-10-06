# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  module OrderHelpers
    extend T::Sig

    include Kernel
    include StripeForce::Constants
    extend SimpleStructuredLogger

    sig do
      params(
        phase_items: T::Array[ContractItemStructure],
        subscription_term: Integer,
        billing_frequency: Integer,
      ).returns(T::Array[T::Boolean])
    end
    def self.prorated_initial_order?(phase_items:, subscription_term:, billing_frequency:)
      log.info 'determining if initial order is prorated'

      if phase_items.empty?
        log.info 'no subscription items, cannot be prorated order'
        return false, false
      end

      OrderHelpers.prorated_order?(subscription_term: subscription_term, billing_frequency: billing_frequency)
    end

    sig do
      params(
        subscription_term: Integer,
        billing_frequency: Integer,
      ).returns(T::Array[T::Boolean])
    end
    def self.prorated_order?(subscription_term:, billing_frequency:)
      log.info 'determining if an order is prorated'

      # if the subscription term does not match the billing frequency of the stripe item, then there will be some proration
      if (subscription_term % billing_frequency) != 0

        is_initial_order_backend_prorated = subscription_term > billing_frequency
        is_initial_order_frontend_prorated = subscription_term < billing_frequency

        log.info 'billing frequency is not divisible by subscription term, assuming prorated initial order',
          subscription_term: subscription_term,
          billing_frequency: billing_frequency,
          is_initial_order_backend_prorated: is_initial_order_backend_prorated,
          is_initial_order_frontend_prorated: is_initial_order_frontend_prorated

        return is_initial_order_backend_prorated, is_initial_order_frontend_prorated
      end

      [false, false]
    end

    sig { params(subscription_schedule: Stripe::SubscriptionSchedule).returns(T::Array[T.any(Stripe::SubscriptionSchedulePhaseSubscriptionItem, Stripe::SubscriptionSchedulePhaseInvoiceItem)]) }
    def self.extract_all_items_from_subscription_schedule(subscription_schedule)
      subscription_schedule.phases.map(&:items).flatten +
        subscription_schedule.phases.map(&:add_invoice_items).flatten
    end

    # after lines have been adjusted with termination line, they should be removed
    # however, having terminated lines in the phase items is helpful *right* until
    # the sub schedule API call is made, which is why this is pulled into a separate method
    # terminated lines are returned after being removed from incoming phase items so they can easily be used by other logic
    sig { params(original_phase_items: T::Array[ContractItemStructure]).returns([T::Array[ContractItemStructure], T::Array[ContractItemStructure]]) }
    def self.remove_terminated_lines(original_phase_items)
      # no mutations
      phase_items = original_phase_items.dup
      terminated_items = []

      phase_items.delete_if do |phase_item|
        if phase_item.fully_terminated?
          log.info 'line item fully terminated, removing', terminated_order_line_id: phase_item.order_line_id
          terminated_items.append(phase_item)
          true
        else
          false
        end
      end

      [phase_items, terminated_items]
    end

    sig { params(raw_days_until_due: T.any(String, Integer, Float)).returns(Integer) }
    def self.transform_payment_terms_to_days_until_due(raw_days_until_due)
      log.info "transforming payment terms to days_until_due", raw_days_until_due: raw_days_until_due
      if raw_days_until_due.is_a?(Integer)
        return raw_days_until_due
      end

      # https://sentry.corp.stripe.com/organizations/stripe/issues/2252023/?project=610&query=is%3Aunresolved
      if raw_days_until_due.is_a?(Float)
        if raw_days_until_due % 1 == 0
          return raw_days_until_due.to_i
        end
        raise StripeForce::Errors::RawUserError.new("Unexpected float for days_until_due option: #{raw_days_until_due}")
      end

      if raw_days_until_due.strip =~ /^[0-9]+$/
        return raw_days_until_due.strip.to_i
      end

      # Get the rest of the string after "Net " or "Net-"
      # https://stackoverflow.com/questions/5006716/getting-the-text-that-follows-after-the-regex-match
      raw_days = raw_days_until_due[/(?<=Net[- ]).*/, 0]

      unless raw_days
        raise StripeForce::Errors::RawUserError.new("Recieved unexpected days_until_due option: #{raw_days_until_due}")
      end

      unless [15, 30, 45, 60, 90].include?(raw_days.to_i)
        raise StripeForce::Errors::RawUserError.new(
          "Recieved unexpected days_until_due option: #{raw_days_until_due}",
           metadata: {
             raw_days_until_due: raw_days_until_due,
             raw_days: raw_days,
           }
        )
      end

      raw_days.to_i
    end

    # TODO this should move to the price helpers
    sig do
      params(
        user: StripeForce::User,
        original_stripe_price: Stripe::Price,
        block: T.nilable(T.proc.params(arg0: Stripe::Price).void)
      ).returns(Stripe::Price)
    end
    def self.duplicate_stripe_price(user, original_stripe_price, &block)
      log.info 'duplicating price', stripe_price_id: original_stripe_price.id

      stripe_price_params = original_stripe_price.to_hash
      stripe_price_params.delete(:id)

      # Stripe::InvalidRequestError: Received unknown parameters: type, object, livemode, created
      stripe_price_params.delete(:type)
      stripe_price_params.delete(:object)
      stripe_price_params.delete(:livemode)
      stripe_price_params.delete(:created)

      # Stripe::InvalidRequestError: You may only specify one of these parameters: unit_amount, unit_amount_decimal.
      stripe_price_params.delete(:unit_amount)

      # same error as above
      # TODO should we conditionally remove the unit amount? Should this be done in the sanitize price params instead?
      stripe_price_params[:tiers]&.each do |tier|
        tier.delete(:unit_amount)

        # the API returns nil for `inf` boundary, but API expects `inf` :(
        if tier.key?(:up_to) && tier[:up_to].nil?
          tier[:up_to] = 'inf'
        end
      end

      stripe_price_params.delete(:active)
      stripe_price = Stripe::Price.construct_from(stripe_price_params)

      if PriceHelpers.tiered_price?(stripe_price)
        stripe_price = PriceHelpers.sanitize_price_tier_params(stripe_price)
      end

      # indicate that this price was created as a duplicate for avoid Stripe API errors
      stripe_price.metadata[Metadata.metadata_key(user, MetadataKeys::DUPLICATE_PRICE)] = true

      # indicates that this price will be auto-archived after created
      stripe_price.metadata[Metadata.metadata_key(user, MetadataKeys::AUTO_ARCHIVE_PRICE)] = true

      # links this price to the original price which this was generated from
      stripe_price.metadata[Metadata.metadata_key(user, MetadataKeys::ORIGINAL_PRICE_ID)] = original_stripe_price.id

      if block_given?
        yield(stripe_price)
      end

      new_stripe_price = Stripe::Price.create(
        stripe_price.to_hash,
        user.stripe_credentials
      )

      log.info 'duplicated price',
        original_stripe_price: original_stripe_price.id,
        new_stripe_price: new_stripe_price.id

      new_stripe_price
    end

    # since the input and output is "fake" stripe subhash, typing here doesn't work
    sig { params(user: StripeForce::User, original_phase_items: T::Array[ContractItemStructure]).returns(T::Array[ContractItemStructure]) }
    def self.ensure_unique_phase_item_prices(user, original_phase_items)
      phase_items = T.cast(Integrations::Utilities::StripeUtil.deep_copy(original_phase_items), T::Array[ContractItemStructure])

      price_ids = []
      phase_items.each do |phase_item|
        price_id = phase_item.stripe_params[:price]

        if !price_ids.include?(price_id)
          price_ids << price_id
          next
        end

        log.debug "price id is duplicated, creating additional price"
        original_price = Stripe::Price.retrieve({id: price_id, expand: %w{tiers}}, user.stripe_credentials)
        new_price = OrderHelpers.duplicate_stripe_price(user, original_price)
        phase_item.stripe_params[:price] = new_price.id
        log.info "new price created to avoid duplicate price", new_price_id: new_price.id
      end

      phase_items
    end

    # You can't pass back the phase items in their original format. Must be modified to avoid:
    # Stripe::InvalidRequestError by specifying both {coupon: discount:}
    def self.sanitize_subscription_schedule_phase_discounts(phase)
      # sanitize the phase discounts
      unless phase["discounts"].nil?
        phase["discounts"].each do |discount|
          Integrations::Utilities::StripeUtil.delete_nil_fields_from_stripe_object(discount)
        end
      end

      unless phase["coupon"].nil?
        phase["coupon"] = nil
      end

      # sanitize the phase item discounts
      phase.items.each do |item|
        unless item["discounts"].nil?
          item["discounts"].each do |discount|
            Integrations::Utilities::StripeUtil.delete_nil_fields_from_stripe_object(discount)
          end
        end
      end
    end

    sig { params(original_phases: T::Array[Stripe::SubscriptionSchedulePhase]).returns(T::Array[Stripe::SubscriptionSchedulePhase]) }
    def self.sanitize_subscription_schedule_phase_params(original_phases)
      # without deep dupping this will mutate the input
      phases = Integrations::Utilities::StripeUtil.deep_copy(original_phases)

      # TODO report https://jira.corp.stripe.com/browse/PLATINT-1479
      # You can't pass back the phase in it's original format, it must be modified to avoid:
      # 'You passed an empty string for 'phases[0][collection_method]'. We assume empty values are an attempt to unset a parameter; however 'phases[0][collection_method]' cannot be unset. You should remove 'phases[0][collection_method]' from your request or supply a non-empty value.'
      phases.each do |phase|
        self.sanitize_subscription_schedule_phase_discounts(phase)
        Integrations::Utilities::StripeUtil.delete_nil_fields_from_stripe_object(phase)
      end

      # (Status 400) (Request req_6sXw1ulKg8naEO) You may only specify one of these parameters: end_date, iterations.>
      phases.each_with_index do |phase, _i|
        # `iterations` is not returned by the API once it's set, instead end_date is
        # however, when we are amending a contract we will set the iterations value as
        # we "regenerate" each phase (even if we don't send it to stripe)
        # because of this, we remove the iterations field if the end_date is set (which indicates)
        # stripe has already consumed the phase change.
        if phase[:end_date] && phase[:iterations]
          Integrations::Utilities::StripeUtil.delete_field_from_stripe_object(
            phase,
            :iterations
          )
        end
      end

      phases
    end

    # if an order does not have a 'AmendedContract' relationship than it is a initial order
    sig { params(user: StripeForce::User, sf_order: Restforce::SObject).returns(T::Boolean) }
    def self.is_order_amendment?(user, sf_order)
      order_with_amended_contract_query = user.sf_client.query(
        # include `Type`, `OpportunityId` for debugging purposes
        <<~EOL
          SELECT Type, OpportunityId,
                Opportunity.SBQQ__AmendedContract__c
          FROM #{SF_ORDER}
          WHERE Id = '#{sf_order.Id}'
        EOL
      )

      if order_with_amended_contract_query.size.zero?
        raise Integrations::Errors::ImpossibleInternalError.new("Amendment contract query should never return an empty result.")
      end

      if order_with_amended_contract_query.size > 1
        raise Integrations::Errors::ImpossibleInternalError.new("More than one contract found for amendment.")
      end

      order_with_amended_contract = order_with_amended_contract_query.first
      amended_contract_id = order_with_amended_contract.dig(SF_OPPORTUNITY, CPQ_AMENDED_CONTRACT)
      is_order_amendment = amended_contract_id.present?

      if !OrderTypeOptions.values.map(&:serialize).include?(order_with_amended_contract.Type)
        log.warn 'salesforce order type is not standard', order_type: order_with_amended_contract.Type
      end

      if is_order_amendment && order_with_amended_contract.Type == OrderTypeOptions::NEW.serialize
        Integrations::ErrorContext.report_edge_case("order is determined to be an amendment, but type is new")
      end

      is_order_amendment
    end

    sig { params(order_date: T.any(Time, DateTime)).returns(T::Boolean) }
    def self.is_order_date_eom?(order_date)
      days_in_month = Date.new(order_date.year, order_date.month, -1).day
      days_in_month == order_date.day
    end

    # shifts the base date ahead, up to 3 days in the same month, to match the anchor order day of month
    # this is used to normalize the amendment date (base_time) to the initial order day-of-month (anchor day)
    # adapted from https://git.corp.stripe.com/stripe-internal/pay-server/blob/2c870dba2b488984042102dde344d55b83a466d9/chalk/tools/lib/chalk_tools/time_utils/time_utils.rb#L307-L320
    sig { params(base_time: T.any(Time, DateTime), anchor_day_of_month: Integer).returns(T.any(Time, DateTime)) }
    def self.anchor_time_to_day_of_month(base_time:, anchor_day_of_month:)
      normalized_base_time = base_time.clone
      normalized_base_time_day_of_month = normalized_base_time.day

      # adjust down for current month
      days_in_month = Date.new(normalized_base_time.year, normalized_base_time.month, -1).day
      days_to_add = [anchor_day_of_month, days_in_month].min - normalized_base_time_day_of_month
      if days_to_add > 0 && days_to_add <= 3
        normalized_base_time += days_to_seconds(days_to_add)
      end

      normalized_base_time
    end

    sig { params(days: Integer).returns(Integer) }
    def self.days_to_seconds(days)
      days * SECONDS_IN_DAY
    end
  end
end
