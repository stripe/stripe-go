# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  sig do
    params(subscription_schedule: Stripe::SubscriptionSchedule,
           sf_order: Restforce::SObject,
           invoice_items: T::Array[ContractItemStructure],
           subscription_items: T::Array[ContractItemStructure])
    .returns(T.nilable(Stripe::RevenueContract))
  end
  def create_revenue_contract_from_sub_schedule(subscription_schedule, sf_order, invoice_items, subscription_items)
    contract = nil
    begin
      request = generate_create_revenue_contract_request(subscription_schedule, sf_order, invoice_items, subscription_items)
      if request.nil?
         # TODO: Add some logging here.
         return
      else
        contract = create_revenue_contract(request, sf_order)
      end

    rescue => e
      # TODO: Currently we silently fail at creating the contrac to not fail the order sync.
      # Once this code path is tested and stable, we should allow failure and retry.
      log.info 'Exception caught during Contract creation.',
        message: e.message
    end

    contract
  end

  sig do
    params(request_json: T.untyped, sf_order: Restforce::SObject).returns(Stripe::RevenueContract)
  end
  private def create_revenue_contract(request_json, sf_order)
    # TODO: Eventually should add the Contracts API into the stripe gems once the API available to public.
    response = Stripe::APIResource.request(
      :post, "/v1/revenue_recognition/contracts",
      request_json,
        @user.stripe_credentials
    )

    responseObj = T.let(response.first, Stripe::StripeResponse)
    if responseObj.http_status != 200
      throw "Unexpected http status code: " + responseObj.http_status
    end

    contract = Stripe::RevenueContract.construct_from(responseObj.data)
    log.info 'Contract has been succesfully created.', contract_id: contract.id

    update_sf_stripe_revenue_contract_id(sf_order, contract.id)
    contract
  end

  sig do
    params(subscription_schedule: Stripe::SubscriptionSchedule,
           sf_order: T.untyped,
           invoice_items: T::Array[ContractItemStructure],
           subscription_items: T::Array[ContractItemStructure])
    .returns(T.untyped)
  end
  private def generate_create_revenue_contract_request(subscription_schedule, sf_order, invoice_items, subscription_items)
    combined_items = invoice_items + subscription_items
    if combined_items.count == 0
      raise Integrations::Errors::ImpossibleState.new("There must be at least 1 item for Contract creation.")
    end

    # We currently have all items with the same subscription start/end date during creation.
    # TODO: This might need to change when multi phase during creation is implemented.
    subscription_schedule_params = StripeForce::Utilities::SalesforceUtil.extract_salesforce_params!(mapper, sf_order, Stripe::SubscriptionSchedule)
    subscription_schedule_start_date = subscription_schedule_params["start_date"]
    subscription_term_from_salesforce = subscription_schedule_params['iterations'].to_i
    start_date = StripeForce::Utilities::SalesforceUtil.salesforce_date_to_unix_timestamp(subscription_schedule_start_date)
    end_date = StripeForce::Utilities::SalesforceUtil.datetime_to_unix_timestamp(
      DateTime.parse(subscription_schedule_start_date) + subscription_term_from_salesforce.months
    )

    contract_items = []
    combined_items.each do |item|
      contract_item = create_revenue_contract_item(item, start_date, end_date)
      if !contract_item.nil?
        contract_items << contract_item
      end
    end

    # The order only has metered items in this case as well and we will skip creating the contract.
    if contract_items.count == 0
      return nil
    end

    # TODO: This date is specific to Cloudflare right now, we default to ActivationDate, but not sure if this
    # is the correct one for future merchants. Will need to revisit.
    signed_date = subscription_schedule.metadata["contract_cf_signed_date"]
    signed_date ||= sf_order[SF_ORDER_ACTIVATED_DATE].to_s

    currency = T.must(combined_items.first).price(@user).currency # this should be the same across all item
    {
      customer: subscription_schedule.customer,
      currency: currency,
      signed_at: StripeForce::Utilities::SalesforceUtil.salesforce_date_to_unix_timestamp(signed_date),
      billing_models: [{
        type: "subscription_schedule",
        subscription_schedule: subscription_schedule.id,
      }],
      items: contract_items,
      metadata: subscription_schedule[:metadata].to_h,
    }
  end

  sig { params(phase_item: ContractItemStructure, start_date: Integer, end_date: Integer).returns(T.untyped) }
  private def create_revenue_contract_item(phase_item, start_date, end_date)
    price = phase_item.price(@user)

    # We currently do not create contract items for metered billing Price.
    if PriceHelpers.metered_price?(price)
      return nil
    end

    sf_order_item = T.let(phase_item.order_line, T.untyped)

    tfc_days = phase_item.stripe_params[:metadata][:contract_tfc_duration]
    tfc_object = if tfc_days.nil?
     nil
    else
      {
        # NOTE: We add + 1 day to expiry date as this date is exclusive.
        expires_at: StripeForce::Utilities::SalesforceUtil.datetime_to_unix_timestamp(
          Time.at(start_date) + (tfc_days.to_i + 1).days
        ),
      }
    end

    # Cloudflare has confirmed that the item_contract_value will always be populated on all order lines. The source is:log
    # "This is a formula field calculated by getting various information from Quote line. Usually this is calculated as MRR * Prorate multiplier"
    amount_subtotal = phase_item.stripe_params[:metadata][:item_contract_value]
    amount_subtotal ||= sf_order_item[SF_ORDER_ITEM_TOTAL_PRICE]

    {
      price: price.id,
      amount_subtotal: normalize_float_amount_in_currency_for_stripe(price.currency, amount_subtotal.to_s, as_decimal: false),
      quantity: phase_item.quantity,
      period: {
        start: start_date,
        end: end_date,
      },
      termination_for_convenience: tfc_object,
      metadata: phase_item.stripe_params[:metadata].to_h,
    }
  end

  sig do
    params(subscription_schedule: Stripe::SubscriptionSchedule,
           initial_sf_order: Restforce::SObject,
           amendment_sf_order: Restforce::SObject,
           aggregate_new_phase_items: T::Array[ContractItemStructure],
           new_invoice_items_in_order: T::Array[ContractItemStructure],
           other_invoice_items: T::Array[T.untyped],
           is_order_terminated: T::Boolean)
    .returns(T.nilable(Stripe::RevenueContract))
  end
  def adjust_revenue_contract_from_sub_schedule(
    subscription_schedule,
    initial_sf_order,
    amendment_sf_order,
    aggregate_new_phase_items,
    new_invoice_items_in_order,
    other_invoice_items,
    is_order_terminated
  )
  updated_contract = nil
  begin
    revenue_contract = retrieve_revenue_contract_from_stripe(initial_sf_order)
    if revenue_contract.nil?
      create_request = generate_create_revenue_contract_request_from_amendment(
        subscription_schedule,
        initial_sf_order,
        aggregate_new_phase_items,
        new_invoice_items_in_order,
        other_invoice_items)

      if create_request.nil?
        # TODO: Add some logging here.
        return
      else
        updated_contract = create_revenue_contract(create_request, initial_sf_order)
      end
    else
      request = generate_adjust_revenue_contract_request(
        subscription_schedule,
        revenue_contract,
        aggregate_new_phase_items,
        new_invoice_items_in_order,
        other_invoice_items,
        is_order_terminated)

      if request[:adjustments].nil? || request[:adjustments].count == 0
        log.info "No adjustments were generated for the revenue contract."
        # TODO: Might need some more logging here.
        return
      end

      # TODO: Eventually should add the Contracts API into the stripe gems once the API available to public.
      response = Stripe::APIResource.request(
        :post, "/v1/revenue_recognition/contracts/#{revenue_contract.id}/adjust",
          request,
          @user.stripe_credentials
      )

      responseObj = T.let(response.first, Stripe::StripeResponse)
      if responseObj.http_status != 200
        throw "Unexpected http status code: " + responseObj.http_status
      end

      updated_contract = Stripe::RevenueContract.construct_from(responseObj.data)
      log.info 'Contract has been adjusted succesfully.', contract_id: updated_contract.id
    end

    update_sf_stripe_revenue_contract_id(amendment_sf_order, updated_contract.id)
   rescue => e
     # TODO: Currently we silently fail at creating the contract to not fail the order sync.
     # Once this code path is tested and stable, we should allow failure and retry.
     log.info 'Exception caught during Contract adjustment.',
       message: e.message
   end

  updated_contract
  end

  sig do
    params(subscription_schedule: Stripe::SubscriptionSchedule,
           initial_sf_order: T.untyped,
           aggregate_new_phase_items: T::Array[ContractItemStructure],
           new_invoice_items_in_order: T::Array[ContractItemStructure],
           other_invoice_items: T::Array[T.untyped])
    .returns(T.untyped)
  end
  private def generate_create_revenue_contract_request_from_amendment(
    subscription_schedule,
    initial_sf_order,
    aggregate_new_phase_items,
    new_invoice_items_in_order,
    other_invoice_items
  )

    # All previous phases must have only metered items for us to create a revenue contract during amendment.
    subscription_schedule.phases.each_with_index do |phase, index|
      break if index == subscription_schedule.phases.length - 1
      if !all_metered_items_on_phase(phase)
        raise Integrations::Errors::ImpossibleState.new("Revenue contract cannot be created for this amendment.")
      end
    end

    new_phase = T.must(subscription_schedule.phases.last)

    contract_items = []
    (aggregate_new_phase_items + new_invoice_items_in_order).each do |item|
      contract_item = create_revenue_contract_item(item, new_phase.start_date, new_phase.end_date)
      if !contract_item.nil?
        contract_items << contract_item
      end
    end

    other_invoice_items.each do |other_item|
      contract_items << create_item_revenue_contract_for_amendments_from_price_info(other_item, new_phase)
    end

    # The amendment only has metered items in this case as well and we will skip creating the contract.
    if contract_items.count == 0
      return nil
    end

    # TODO: This date is specific to Cloudflare right now, we default to ActivationDate, but not sure if this
    # is the correct one for future merchants. Will need to revisit.
    signed_date = subscription_schedule.metadata["contract_cf_signed_date"]
    signed_date ||= initial_sf_order.ActivatedDate

    first_price = T.let(contract_items.first, T.untyped)[:price]
    stripe_price = Stripe::Price.retrieve(first_price, @user.stripe_credentials)

    currency = stripe_price.currency # this should be the same across all item
    {
      customer: subscription_schedule.customer,
      currency: currency,
      signed_at: StripeForce::Utilities::SalesforceUtil.salesforce_date_to_unix_timestamp(signed_date),
      billing_models: [{
        type: "subscription_schedule",
        subscription_schedule: subscription_schedule.id,
      }],
      items: contract_items,
      metadata: subscription_schedule[:metadata].to_h,
    }
  end

  sig do
    params(subscription_schedule: Stripe::SubscriptionSchedule,
           revenue_contract: Stripe::RevenueContract,
           aggregate_new_phase_items: T::Array[ContractItemStructure],
           new_invoice_items_in_order: T::Array[ContractItemStructure],
           other_invoice_items: T::Array[T.untyped],
           is_order_terminated: T::Boolean)
    .returns(T.untyped)
  end
  private def generate_adjust_revenue_contract_request(
    subscription_schedule,
    revenue_contract,
    aggregate_new_phase_items,
    new_invoice_items_in_order,
    other_invoice_items,
    is_order_terminated
  )
    new_phase = T.must(subscription_schedule.phases.last)

    update_contract_items = []
    add_contract_items = []

    if !is_order_terminated
      # TODO: add collect items here from rev contract pages
      (aggregate_new_phase_items + new_invoice_items_in_order).each do |item|
        price = item.price(@user)
        contract_item = revenue_contract.items.data.find {|c_item| price.id == c_item.price }
        if contract_item.nil?
          new_item = create_revenue_contract_item(item, new_phase.start_date, new_phase.end_date)
          if !new_item.nil?
            add_contract_items << {
              type: "add_contract_item",
              add_contract_item: new_item,
            }
          end
        else
          update_adjustment = create_update_adjustment_for_item_contract_item(contract_item, item, new_phase)
          if !update_adjustment.nil?
            update_contract_items << update_adjustment
          end
        end
      end

      other_invoice_items.each do |other_item|
        add_contract_items << {
          type: "add_contract_item",
          add_contract_item: create_item_revenue_contract_for_amendments_from_price_info(other_item, new_phase),
        }
      end
    end

    if subscription_schedule.phases.count >= 2 || is_order_terminated
      previous_phase = is_order_terminated ? new_phase : T.must(subscription_schedule.phases[subscription_schedule.phases.count - 2])
      current_phase = is_order_terminated ? nil : new_phase
      update_contract_items += create_update_adjustment_for_removed_items(previous_phase, current_phase, revenue_contract.items.data)
    end

    {
      # adjusted_at: 11, ned to put the order amendment date here, let's pull that
      adjusted_at: new_phase.start_date,
      adjustments: update_contract_items + add_contract_items,
    }
  end

  sig do
    params(contract_item: T.untyped,
           updated_sf_item: ContractItemStructure,
           new_phase: T.untyped)
    .returns(T.untyped)
  end
  private def create_update_adjustment_for_item_contract_item(contract_item, updated_sf_item, new_phase)
    updated_order_line = T.let(updated_sf_item.order_line, T.untyped)

    new_qty = updated_sf_item.quantity
    new_end = new_phase.end_date

    amount_subtotal = updated_sf_item.stripe_params[:metadata][:item_contract_value]
    amount_subtotal ||= normalize_float_amount_in_currency_for_stripe(updated_sf_item.price.currency, updated_order_line.TotalPrice.to_s, as_decimal: false)
    new_amount = amount_subtotal

    if new_end != contract_item.period.end || new_qty != contract_item.quantity || new_amount != contract_item.amount_subtotal
      {
        type: "update_contract_item",
        update_contract_item:
        {
          id: contract_item.id,
          amount_subtotal: new_amount,
          quantity: new_qty,
          period: {
            start: contract_item.period.start,
            end: new_end,
          },
        },
      }
    else
      nil
    end
  end

  sig do
    params(previous_phase: T.untyped,
           current_phase: T.untyped,
           revenue_contract_items: T::Array[T.untyped])
    .returns(T::Array[T.untyped])
  end
  private def create_update_adjustment_for_removed_items(previous_phase, current_phase, revenue_contract_items)
    update_items = []
    previous_phase.items.each do |prev_item|
      item_existing = current_phase && current_phase.items && current_phase.items.find {|it| it.price == prev_item.price }
      contract_item = revenue_contract_items.find {|c_item| prev_item.price == c_item.price }
      if item_existing.nil? && !contract_item.nil? && contract_item.period.end.to_i != previous_phase.end_date
        stripe_price = Stripe::Price.retrieve(contract_item.price, @user.stripe_credentials)
        start_date = Time.at(contract_item.period.start).to_date
        end_date = Time.at(previous_phase.end_date).to_date
        month_terms = (end_date.year * 12 + end_date.month) - (start_date.year * 12 + start_date.month)

        # TODO: This feels a bit hacky, need to validate if there is a better way to get the new full contract
        # item value here if the item is removed.
        new_amount = stripe_price.unit_amount * contract_item.quantity.to_i * month_terms
        update_items << {
          type: "update_contract_item",
          update_contract_item:
          {
            id: contract_item.id,
            amount_subtotal: new_amount,
            period: {
              start: contract_item.period.start,
              end: previous_phase.end_date,
            },
          },
        }
      end
    end

    update_items
  end

  sig do
    params(add_item: T.untyped,
           new_phase: T.untyped)
    .returns(T.untyped)
  end
  def create_item_revenue_contract_for_amendments_from_price_info(add_item, new_phase)
    if add_item[:price_data]
      price_data = T.must(add_item[:price_data])
      unit_amount = price_data[:unit_amount_decimal].to_i

      # Best effort to identify the phase invoice item that created the price from this price data
      # TODO: Revisit and verify the corectness of this.
      phase_item = new_phase.add_invoice_items.find do |item|
        # compare metadata
        item_metadata = item[:metadata] && item[:metadata].to_h.transform_values(&:to_s)
        add_item_metadata = add_item[:metadata] && add_item[:metadata].to_h.transform_values(&:to_s)
        metadata_equal = item_metadata == add_item_metadata

        # compare period
        item_start = item[:period] && item[:period][:start] && item[:period][:start].to_h.transform_values(&:to_s)
        add_item_start = add_item[:period] && add_item[:period][:start] && add_item[:period][:start].to_h.transform_values(&:to_s)
        item_end = item[:period] && item[:period][:end] && item[:period][:end].to_h.transform_values(&:to_s)
        add_item_end = add_item[:period] && add_item[:period][:end] && add_item[:period][:end].to_h.transform_values(&:to_s)
        period_equal = item_start == add_item_start && item_end == add_item_end

        metadata_equal && period_equal && item[:quantity] == add_item[:quantity]
      end
    else
      stripe_price_id = add_item[:price].to_s
      stripe_price = Stripe::Price.retrieve(stripe_price_id, @user.stripe_credentials)
      unit_amount = stripe_price.unit_amount

      phase_item = new_phase.add_invoice_items.find {|item| item.price == stripe_price_id }
    end

    start_date = new_phase.start_date
    end_date = new_phase.end_date
    if add_item[:period]
      start_date = convert_period_date_to_unix_timestamp(add_item[:period][:start], new_phase)
      end_date = convert_period_date_to_unix_timestamp(add_item[:period][:end], new_phase)
    end

    tfc_days = T.must(phase_item)[:metadata][:contract_tfc_duration]
    tfc_object = if tfc_days.nil?
     nil
    else
      {
        # NOTE: We add + 1 day to expiry date as this date is exclusive.
        expires_at: StripeForce::Utilities::SalesforceUtil.datetime_to_unix_timestamp(
          Time.at(start_date) + (tfc_days.to_i + 1).days
        ),
      }
    end

    amount_subtotal = phase_item[:metadata][:item_contract_value]
    amount_subtotal ||= unit_amount * add_item[:quantity].to_i

    {
      price: phase_item.price,
      amount_subtotal: amount_subtotal,
      quantity: add_item[:quantity].to_i,
      period: {
        start: start_date,
        end: end_date,
      },
      termination_for_convenience: tfc_object,
      metadata: phase_item[:metadata].to_h,
    }
  end

  sig do
    params(period_date: T.untyped, phase: T.untyped).returns(Integer)
  end
  def convert_period_date_to_unix_timestamp(period_date, phase)
     period_type = period_date[:type].to_s
     case period_type
     when "timestamp"
      T.must(period_date[:timestamp]).to_i
     when "subscription_period_end", "phase_end"
      # TODO: Verify again if sub period end should map to phase end. This is the latest phase.
      phase.end_date
     when "phase_start"
      phase.start_date
     else
      raise Integrations::Errors::ImpossibleState.new("Unexpected period type.")
     end
  end

  sig { params(phase: T.untyped).returns(T::Boolean) }
  def all_metered_items_on_phase(phase)
    return false if phase.add_invoice_items.count != 0
    all_items_metered = T.let(true, T::Boolean)
    phase.items.each do |phase_item|
      stripe_price = Stripe::Price.retrieve(phase_item.price.to_s, @user.stripe_credentials)
      if !PriceHelpers.metered_price?(stripe_price)
        all_items_metered = false
        break
      end
    end

    all_items_metered
  end

  sig do
    params(subscription_schedule: Stripe::SubscriptionSchedule,
           initial_sf_order: Restforce::SObject,
           amendment_sf_order: Restforce::SObject)
    .returns(T.nilable(Stripe::RevenueContract))
  end
  def terminate_revenue_contract(
    subscription_schedule,
    initial_sf_order,
    amendment_sf_order
  )
  updated_contract = nil
  begin
    revenue_contract = retrieve_revenue_contract_from_stripe(initial_sf_order)
    if revenue_contract.nil?
      # Some logging here
      return
    end

    if subscription_schedule.status != "canceled"
      raise Integrations::Errors::ImpossibleState.new("Subscription schedule must be canceled before terminating revenue contract.")
    end


    # TODO: Cloudflare said there is a case they would prefer calling mark_uncollectible instaed of void based on a
    # property on salesforce. Need to follow up on that and change here.
    request =
      {
        voided_at: T.let(subscription_schedule, T.untyped).canceled_at,
      }

    # TODO: Eventually should add the Contracts API into the stripe gems once the API available to public.
    response = Stripe::APIResource.request(
      :post, "/v1/revenue_recognition/contracts/#{revenue_contract.id}/void",
        request,
        @user.stripe_credentials
    )

    responseObj = T.let(response.first, Stripe::StripeResponse)
    if responseObj.http_status != 200
      throw "Unexpected http status code: " + responseObj.http_status
    end

    updated_contract = Stripe::RevenueContract.construct_from(responseObj.data)
    log.info 'Contract has been voided succesfully.', contract_id: updated_contract.id

    update_sf_stripe_revenue_contract_id(amendment_sf_order, updated_contract.id)
   rescue => e
     # TODO: Currently we silently fail at creating the contract to not fail the order sync.
     # Once this code path is tested and stable, we should allow failure and retry.
     log.info 'Exception caught during Contract termination.',
       message: e.message
   end

  updated_contract
  end


  sig { params(sf_object: T.untyped, stripe_contract_id: String).void }
  def update_sf_stripe_revenue_contract_id(sf_object, stripe_contract_id)
    stripe_contract_id_field = prefixed_stripe_field(GENERIC_STRIPE_REVENUE_CONTRACT_ID)

    if sf_object[stripe_contract_id_field]
      if sf_object[stripe_contract_id_field] == stripe_contract_id
        log.info 'stripe contract id already exists in salesforce, no update',
          stripe_contract_id: sf_object[stripe_contract_id_field],
          field_name: stripe_contract_id_field
        return
      end

      log.info 'stripe contract id is different than existing stripe contract id in salesforce, overwriting',
        old_stripe_contract_id: sf_object[stripe_contract_id_field],
        new_stripe_contract_id: stripe_contract_id,
        field_name: stripe_contract_id_field
    end

    backoff do
      sf.update!(sf_object.sobject_type, {
        SF_ID => sf_object.Id,
        stripe_contract_id_field => stripe_contract_id,
      })
    end
  end

  sig { params(sf_object: Restforce::SObject).returns(T.nilable(Stripe::RevenueContract)) }
  def retrieve_revenue_contract_from_stripe(sf_object)
    revenue_contract_id = sf_object[prefixed_stripe_field(GENERIC_STRIPE_REVENUE_CONTRACT_ID)]
    return if revenue_contract_id.nil?

    response = Stripe::APIResource.request(
      :get,
      "/v1/revenue_recognition/contracts/#{revenue_contract_id}",
      nil,
      @user.stripe_credentials
    )
    responseObj = T.let(response.first, Stripe::StripeResponse)
    Stripe::RevenueContract.construct_from(responseObj.data)
  end
end
