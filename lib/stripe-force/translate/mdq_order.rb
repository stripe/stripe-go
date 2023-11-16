# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  include StripeForce::Constants

  sig do
    params(
      sf_order: T.untyped,
      invoice_items: T::Array[ContractItemStructure],
      contract_items: T::Array[ContractItemStructure],
      subscription_schedule: Stripe::SubscriptionSchedule,
      stripe_customer: Stripe::Customer
    ).returns(T::Array[Stripe::SubscriptionSchedulePhase])
  end
  def generate_phases_for_initial_order_with_mdq(sf_order:, invoice_items:, contract_items:, subscription_schedule:, stripe_customer:)
    sub_schedule_phases = []

    # extract the quote start date and subscription term from the order
    # to calculate the end date
    string_start_date_from_salesforce = subscription_schedule['start_date']
    start_date_from_salesforce = DateTime.parse(string_start_date_from_salesforce)
    subscription_term_from_salesforce = StripeForce::Utilities::SalesforceUtil.extract_subscription_term_from_order!(mapper, sf_order)

    order_end_date = StripeForce::Utilities::SalesforceUtil.datetime_to_unix_timestamp(start_date_from_salesforce + subscription_term_from_salesforce.months)

    non_segmented_contract_items, segmented_contract_items = split_and_sort_mdq_item(contract_items)
    non_segmented_subscription_items = non_segmented_contract_items.present? ? non_segmented_contract_items.map(&:stripe_params) : []

    # CPQ Quote discounts will apply across each phase / segment
    phase_discounts = stripe_discounts_for_sf_object(sf_object: sf_order)

    # for each subscription item, add it to the corresponding sub phase
    T.must(segmented_contract_items).each do |segment_item|
        # many of the mdq fields live on the corresponding quote line
        quote_line_data = backoff { @user.sf_client.find(CPQ_QUOTE_LINE, segment_item.order_line[CPQ_QUOTE_LINE]) }
        segment_item_end_date = StripeForce::Utilities::SalesforceUtil.salesforce_date_to_unix_timestamp(quote_line_data[CPQ_QUOTE_SUBSCRIPTION_END_DATE]) + 1.day.to_i

        # adds a new phase for the newest segment
        mdq_segment_index = segment_item.mdq_segment_index
        if sub_schedule_phases.empty? || sub_schedule_phases.count < T.must(mdq_segment_index)
          proration_behavior = StripeProrationBehavior::NONE.serialize

          # only the initial phase might use Stripe prorations
          if sub_schedule_phases.empty?
            skip_initial_phase_proration = sf_order[prefixed_stripe_field(SKIP_PAST_INITIAL_INVOICES)] == true
            if !skip_initial_phase_proration
              proration_behavior = StripeProrationBehavior::CREATE_PRORATIONS.serialize
            end
          end

          # create a new phase for this segment index
          sub_schedule_phases << {
            add_invoice_items: [],
            items: non_segmented_subscription_items.dup,
            end_date: segment_item_end_date,
            metadata: Metadata.stripe_metadata_for_sf_object(@user, sf_order).merge({
              salesforce_segment_index: segment_item.order_line[CPQ_ORDER_ITEM_SEGMENT_INDEX],
              salesforce_segment_label: quote_line_data[CPQ_ORDER_ITEM_SEGMENT_LABEL],
            }),
            discounts: phase_discounts,
            proration_behavior: proration_behavior,
          }
        end

        # add the subscription item to the existing segment index phase
        last_index = sub_schedule_phases.count - 1

        # let's be defensive and throw a error if the order items in the same segment have different end dates
        if sub_schedule_phases[last_index][:end_date] != segment_item_end_date
          log.error 'sf order items in the same segment have different end dates', sf_order_item: segment_item.order_line
          raise Integrations::Errors::ImpossibleState.new('Salesforce Order Items in the same MDQ segment has different end date than the other segments.', salesforce_object: segment_item.order_line)
        end

        # add the item to the last phase / segment
        order_item_metadata = {"metadata" => {'salesforce_segment_key': segment_item.order_line[CPQ_ORDER_ITEM_SEGMENT_KEY]}}
        sub_schedule_phases[last_index][:items] << segment_item.stripe_params.merge(order_item_metadata)
    end

    # Salesforce EndDate is the end of the last day of the subscription while the calculated EndDate we send to Stripe
    # is the day after the last day of the subscription. This is because Stripe's subscription schedule end date is exclusive.
    sub_schedule_phases[sub_schedule_phases.count - 1][:end_date] = order_end_date
    sub_schedule_phases
  end

  sig { params(contract_structure: ContractStructure).returns(T.nilable(Stripe::SubscriptionSchedule)) }
  def update_subscription_schedule_phases_from_order_amendments_with_mdq(contract_structure)
    log.info 'processing mdq amendment orders'

    subscription_schedule = get_subscription_schedule_for_initial_order(contract_structure.initial)
    subscription_schedule_phases = subscription_schedule.phases

    sf_initial_order_items = order_lines_from_order(contract_structure.initial)
    _invoice_items, initial_order_contract_items = phase_items_from_order_lines(sf_initial_order_items)

    previous_phase_contract_items = initial_order_contract_items
    contract_structure.amendments.each_with_index do |sf_order_amendment, index|
      locker.lock_salesforce_record(sf_order_amendment)

      # raise error if attempting to amend a canceled subscription schedule
      if subscription_schedule.status == 'canceled'
        throw_user_failure!(
          salesforce_object: sf_order_amendment,
          message: "Stripe subscription schedule has already been cancelled and cannot be modified."
        )
      end

      log.info 'processing amendment order for mdq products', sf_order_amendment_id: sf_order_amendment.Id, index: index

      # verify that this amendment doesn't modify the segment dates
      sf_amendment_order_items = order_lines_from_order(sf_order_amendment)
      _invoice_items, amendment_order_contract_items = phase_items_from_order_lines(sf_amendment_order_items)
      if !amendment_has_same_segment_dates?(previous_phase_contract_items, amendment_order_contract_items)
        raise StripeForce::Errors::RawUserError.new("Amendment order segment dates do not match initial order segment dates.")
      end

      # this joins the initial order items with the amendment order items
      # TODO check how revisions work when there are multiple segments => is the revised order line from the previous phase?
      _invoice_items_in_order, aggregate_phase_items = build_phase_items_from_order_amendment(previous_phase_contract_items, sf_order_amendment)

      # determine if this has synced previously
      if sf_order_amendment[prefixed_stripe_field(GENERIC_STRIPE_ID)].present?
        log.info "amendment order already translated, skipping", sf_order_amendment_id: sf_order_amendment.Id, index: index

        # it's important we update the previous phase items since the latest amendment depends on previous amendments
        aggregate_phase_items, _terminated_phase_items = OrderHelpers.remove_terminated_lines(aggregate_phase_items)
        previous_phase_contract_items = aggregate_phase_items
        next
      end

      # let's filter out any previous phase contract items from past segments
      # since we don't want to amend past segments
      # note: we can't rely on the segment index since segment index will change as segments "expire"
      amendment_order_item_end_dates = sf_amendment_order_items.map {|item| StripeForce::Utilities::SalesforceUtil.salesforce_date_to_unix_timestamp(item["EndDate"]) }.uniq
      amendment_start_date_timestamp = get_amendment_start_date(mapper, sf_order_amendment)
      aggregate_phase_items = T.must(aggregate_phase_items).filter do |contract_item|
        segment_item_end_date = StripeForce::Utilities::SalesforceUtil.salesforce_date_to_unix_timestamp(contract_item.order_line["EndDate"])
        segment_item_end_date > amendment_start_date_timestamp && amendment_order_item_end_dates.include?(segment_item_end_date)
      end

      # determine if this is a full termination order
      # determine if this is a termination order
      is_order_terminated = aggregate_phase_items.all?(&:fully_terminated?)
      if is_order_terminated
        log.info 'mdq amendment order is a termination order', sf_order_amendment_id: sf_order_amendment.Id

        if contract_structure.amendments.count - 1 != index
          raise StripeForce::Errors::RawUserError.new("Processing a termination order, but there's more amendments queued.")
        end

        # the intention here is to void/reverse out the entire contract, and this is the closest API call we have
        log.info 'cancelling subscription schedule immediately', sf_order_amendment_id: sf_order_amendment
        subscription_schedule = T.cast(subscription_schedule.cancel({invoice_now: false, prorate: false}, @user.stripe_credentials), Stripe::SubscriptionSchedule)
        return subscription_schedule
      end

      # split the segmented_subscription_items by segment keys since the same segment key/index touches the same phase
      aggregate_phase_items, _terminated_phase_items = OrderHelpers.remove_terminated_lines(aggregate_phase_items)
      non_segmented_contract_items, segmented_contract_items = split_and_sort_mdq_item(aggregate_phase_items)
      if segmented_contract_items.present?
        log.info 'amendment order is amending an mdq product'

        segmented_contract_items.chunk(&:mdq_segment_index).each do |segment_index, items_by_segment_index|
          if items_by_segment_index.empty?
            next
          end

          # let's be defensive and assert that all the items in the same segment index have the same segment key
          all_items_have_same_segment_index(items_by_segment_index)

          # find the phase this item exists in by comparing the start and end dates
          # note: we can't rely on the segment index since segment index will change as segments "expire"
          # quote_line_data = backoff { @user.sf_client.find(CPQ_QUOTE_LINE, T.must(items_by_segment_index.first).order_line[CPQ_QUOTE_LINE]) }
          item_end_date = StripeForce::Utilities::SalesforceUtil.salesforce_date_to_unix_timestamp(T.must(items_by_segment_index.first).order_line["EndDate"])
          matching_subscription_schedule_phase = subscription_schedule_phases.filter do |phase|
            phase.end_date == (item_end_date + 1.day.to_i)
          end

          if matching_subscription_schedule_phase.count == 0
            raise Integrations::Errors::ImpossibleState.new("Expected to find one subscription schedule phase with segment index but found none.", salesforce_object: sf_order_amendment)
          elsif matching_subscription_schedule_phase.count > 1
            # we expect this to happen if the user has amended the current phase previously
            # for now, we don't support it
            log.info 'found more than one phase for this mdq segment', matching_phase_count: matching_subscription_schedule_phase.count, segment_index: segment_index
            raise Integrations::Errors::ImpossibleState.new("Expected to find one subscription schedule phase with segment index but found more than one.", salesforce_object: sf_order_amendment)
          end

          # find the phase this segment is amending
          matching_phase = T.must(matching_subscription_schedule_phase.last)
          insert_index = subscription_schedule_phases.find_index(matching_phase)
          if insert_index.nil?
            throw Integrations::Errors::ImpossibleState.new("Expected to find a matching phase for segment index but found none.", salesforce_object: sf_order_amendment)
          end

          new_phase = create_new_phase_from_amendment(sf_order_amendment, matching_phase, items_by_segment_index, non_segmented_contract_items.nil? ? [] : non_segmented_contract_items.dup)
          new_phase[:metadata] = new_phase[:metadata].to_h.merge(Metadata.stripe_metadata_for_sf_object(@user, sf_order_amendment))
          if is_current_phase?(subscription_schedule, matching_phase)
            log.info 'amendment order is modifying the current phase', sf_order_amendment_id: sf_order_amendment.Id, segment_index: segment_index, segment_item_diff: items_by_segment_index.count

            # we only want to update the timestamps if we are modifying the current phase
            new_phase[:start_date] = determine_order_start_date(subscription_schedule, sf_order_amendment)
            matching_phase[:end_date] = new_phase[:start_date]
            subscription_schedule_phases.insert(insert_index + 1, new_phase)
          else
            log.info 'amendment order is modifying a future phase, replacing future phase with new phase', sf_order_amendment_id: sf_order_amendment.Id, segment_index: segment_index

            # replace the future phase with this "updated" phase
            subscription_schedule_phases.insert(insert_index + 1, new_phase)
            subscription_schedule_phases.delete_at(insert_index)
          end
        end
      else
        log.info 'amendment order is amending a non mdq product on order with mdq products', sf_order_amendment_id: sf_order_amendment.Id
        raise StripeForce::Errors::UserError.new("Amending a non mdq product on an Order with mdq product is not yet supported.", salesforce_object: sf_order_amendment)
      end

      # general subscription schedule phase cleanup
      subscription_schedule_phases = OrderHelpers.sanitize_subscription_schedule_phase_params(subscription_schedule_phases)
      subscription_schedule.phases = subscription_schedule_phases

      # We do not currently map to the subscription schedule (again) when there is an amendment order
      # we should consider remapping the subscription schedule when there is an amendment order but for now we will map specific fields
      subscription_schedule = apply_amendment_order_mappings(mapper, subscription_schedule, sf_order_amendment)

      # note: to support stacked amendments, we want to update the local sub_schedule and sub_phases
      # because  Stripe converts 'now' to a timestamp and we want to use that timestamp when there is a stacked amendment
      subscription_schedule = T.cast(subscription_schedule.save({}, @user.stripe_credentials), Stripe::SubscriptionSchedule)
    end

    subscription_schedule
  end

  def update_current_phase_with_amendment(subscription_schedule, matching_phase, sf_order_amendment, items_by_segment_index, non_segmented_subscription_items)
      # extract the start date, sub term, end date of the amendment
      phase_params = StripeForce::Utilities::SalesforceUtil.extract_salesforce_params!(mapper, sf_order_amendment, Stripe::SubscriptionSchedule)
      # convert the string Salesforce start date from to a timestamp
      sf_order_amendment_start_date = phase_params['start_date']
      phase_params['start_date'] = StripeForce::Utilities::SalesforceUtil.salesforce_date_to_unix_timestamp(sf_order_amendment_start_date)

      # this is important to remove since the user can't change the segment length and corresponding end date of the phase
      _subscription_term_from_salesforce = phase_params.delete('iterations').to_i
      _end_date = phase_params.delete('end_date').to_i

      # align date boundaries of the schedules
      new_phase = Stripe::StripeObject.construct_from({
        add_invoice_items: [], # TODO invoice_items_in_order.map(&:stripe_params) + invoice_items_for_prorations + negative_invoice_items,
        items: Integrations::Utilities::StripeUtil.deep_copy(items_by_segment_index).map(&:stripe_params) + non_segmented_subscription_items,
        proration_behavior: StripeProrationBehavior::NONE.serialize,
        metadata: (matching_phase.metadata || {}).to_h.merge(Metadata.stripe_metadata_for_sf_object(@user, sf_order_amendment)),
      }.merge(phase_params))

      # determine if this is a backdated order since this has implications on how we prorate
      stripe_customer_id = T.cast(subscription_schedule.customer, String)
      current_time = OrderAmendment.determine_current_time(@user, stripe_customer_id)
      normalized_current_time = StripeForce::Utilities::SalesforceUtil.datetime_to_unix_timestamp(Time.at(current_time))

      is_order_backdated = phase_params['start_date'] < normalized_current_time
      is_same_day = normalized_current_time == new_phase.start_date
      if is_same_day
        log.info 'amendment starts on the current day, using now'
        new_phase.start_date = 'now'
      elsif is_order_backdated
        # if this is a backdated amendment, then use the current time to update the subscription schedule
        log.info 'processing a backdated amendment order for a Salesforce Order with MDQ product, using now'
        new_phase.start_date = 'now'
      end

      matching_phase.end_date = new_phase.start_date
      [matching_phase, new_phase]
  end

  sig { params(sf_order_amendment: Restforce::SObject, matching_phase: T.untyped, items_by_segment_index: T::Array[ContractItemStructure], non_segmented_subscription_items: T::Array[ContractItemStructure]).returns(T.untyped) }
  def create_new_phase_from_amendment(sf_order_amendment, matching_phase, items_by_segment_index, non_segmented_subscription_items)
    phase_items = items_by_segment_index.concat(non_segmented_subscription_items)
    dup_matching_phase = Integrations::Utilities::StripeUtil.deep_copy(matching_phase)
    dup_matching_phase[:items] = Integrations::Utilities::StripeUtil.deep_copy(phase_items).map(&:stripe_params)
    dup_matching_phase[:proration_behavior] = StripeProrationBehavior::NONE.serialize
    dup_matching_phase
  end

  sig { params(subscription_schedule: Stripe::SubscriptionSchedule, sf_order: Restforce::SObject).returns(T.untyped) }
  def determine_order_start_date(subscription_schedule, sf_order)
    # convert the Salesforce start date from a string to a timestamp
    phase_params = StripeForce::Utilities::SalesforceUtil.extract_salesforce_params!(mapper, sf_order, Stripe::SubscriptionSchedule)
    sf_order_start_date_timestamp = StripeForce::Utilities::SalesforceUtil.salesforce_date_to_unix_timestamp(phase_params['start_date'])

    # determine if this is a backdated order since this has implications downstream
    stripe_customer_id = T.cast(subscription_schedule.customer, String)
    current_time = OrderAmendment.determine_current_time(@user, stripe_customer_id)
    normalized_current_time = StripeForce::Utilities::SalesforceUtil.datetime_to_unix_timestamp(Time.at(current_time))

    # determine if this order is amendment same day or backdated
    is_same_day = normalized_current_time == sf_order_start_date_timestamp
    is_order_backdated = sf_order_start_date_timestamp < normalized_current_time
    if is_same_day
      log.info 'amendment starts on the current day, using now'
      return 'now'
    elsif is_order_backdated
      # if this is a backdated amendment, then use the current time to update the subscription schedule
      log.info 'backdated amendment, using now'
      return 'now'
    end

    sf_order_start_date_timestamp
  end

  sig { params(contract_items: T::Array[ContractItemStructure],).returns(T::Array[T::Array[ContractItemStructure]]) }
  def split_and_sort_mdq_item(contract_items)
    # split the subscription items by mdq and non-mdq products
    # non segmented items are for the entire contact so they will be added to each phase
    segmented_subscription_items, non_segmented_subscription_items = contract_items.partition(&:is_mdq_segment?)

    # sort the mdq products by segment index
    segmented_subscription_items = segmented_subscription_items.sort_by do |item|
      if item.mdq_dimension_type == 'Custom'
        raise StripeForce::Errors::RawUserError.new("MDQ products with custom segments are not yet supported.")
      end

      T.must(item.mdq_segment_index)
    end

    [non_segmented_subscription_items, segmented_subscription_items]
  end

  sig { params(items_by_index: T::Array[StripeForce::Translate::ContractItemStructure]).returns(T::Boolean) }
  def all_items_have_same_segment_index(items_by_index)
    if items_by_index.empty? || items_by_index.count == 1
      return true
    end

    first_item_segment_index = T.must(items_by_index.first).mdq_segment_index
    items_by_index.each do |item|
      if item.mdq_segment_index != first_item_segment_index
        return false
      end
    end

    true
  end

  sig { params(subscription_schedule: Stripe::SubscriptionSchedule, phase: T.untyped).returns(T::Boolean) }
  def is_current_phase?(subscription_schedule, phase)
    if subscription_schedule.current_phase.blank?
      log.info 'subscription schedule does not have a current phase', subscription_schedule_id: subscription_schedule.id
      return false
    end

    # get the current phase start and end date
    current_phase_start_date = subscription_schedule.current_phase[:start_date]
    current_phase_end_date = subscription_schedule.current_phase[:end_date]

    if phase[:start_date] == current_phase_start_date && phase[:end_date] == current_phase_end_date
      return true
    end

    false
  end

  def get_amendment_start_date(mapper, sf_order_amendment)
    # verify that the amendment order segments match the initial orders segments
    amendment_contract_data_query = StripeForce::Translate::OrderHelpers.get_amended_contract_data(mapper.user, sf_order_amendment)
    amended_contract_id = amendment_contract_data_query.dig(SF_OPPORTUNITY, CPQ_AMENDED_CONTRACT)
    amendment_contract_start_date_query = backoff { @user.sf_client.query("Select #{CONTRACT_AMENDMENT_START_DATE} from #{SF_CONTRACT} where ID = '#{amended_contract_id}'") }
    contract_amendment_start_date = amendment_contract_start_date_query.first[CONTRACT_AMENDMENT_START_DATE]
    if contract_amendment_start_date.blank?
      raise StripeForce::Errors::RawUserError.new("Field 'SBQQ__AmendmentStartDate__c' on contract is empty.")
    end

    contract_amendment_start_date_timestamp = StripeForce::Utilities::SalesforceUtil.salesforce_date_to_unix_timestamp(contract_amendment_start_date)
    # let's be defensive and ensure the amendment order start date
    # is the same as the amendment start date on the contract
    amendment_order_start_date = StripeForce::Utilities::SalesforceUtil.extract_subscription_start_date_from_order(mapper, sf_order_amendment)
    if StripeForce::Utilities::SalesforceUtil.datetime_to_unix_timestamp(amendment_order_start_date) != contract_amendment_start_date_timestamp
      raise StripeForce::Errors::RawUserError.new("Start date on amendment order does not match the corresponding contract's amendment start date field.")
    end

    contract_amendment_start_date_timestamp
  end

  sig do
    params(initial_order_contract_items: T::Array[StripeForce::Translate::ContractItemStructure], amendment_order_contract_items: T::Array[StripeForce::Translate::ContractItemStructure],).returns(T::Boolean)
  end
  def amendment_has_same_segment_dates?(initial_order_contract_items, amendment_order_contract_items)
      # get all the segment start dates from the initial order
      initial_order_segment_start_dates = initial_order_contract_items.map do |item|
        quote_line_data = backoff { @user.sf_client.find(CPQ_QUOTE_LINE, item.order_line[CPQ_QUOTE_LINE]) }
        segment_item_start_date = StripeForce::Utilities::SalesforceUtil.salesforce_date_to_unix_timestamp(quote_line_data[QUOTE_LINE_EFFECTIVE_START_DATE])
        segment_item_start_date
      end.uniq

      # get all the segment start dates from the amendment order
      # and confirm that they match the initial order
      amendment_order_contract_items.each do |item|
        quote_line_data = backoff { @user.sf_client.find(CPQ_QUOTE_LINE, item.order_line[CPQ_QUOTE_LINE]) }
        segment_item_end_date = StripeForce::Utilities::SalesforceUtil.salesforce_date_to_unix_timestamp(quote_line_data[QUOTE_LINE_EFFECTIVE_START_DATE])
        if !initial_order_segment_start_dates.include?(segment_item_end_date)
          return false
        end
      end

      true
  end
end
