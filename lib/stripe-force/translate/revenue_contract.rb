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

      # TODO: Eventually should add the Contracts API into the stripe gems once the API available to public.
      response = Stripe::APIResource.request(
        :post, "/v1/revenue_recognition/contracts",
          request,
          @user.stripe_credentials
      )

      responseObj = T.let(response.first, Stripe::StripeResponse)
      if responseObj.http_status != 200
        throw "Unexpected http status code: " + responseObj.http_status
      end

      contract = Stripe::RevenueContract.construct_from(responseObj.data)
      log.info 'Contract has been succesfully created.', contract_id: contract.id

      update_sf_stripe_revenue_contract_id(sf_order, contract.id)
    rescue => e
      # TODO: Currently we silently fail at creating the contrac to not fail the order sync.
      # Once this code path is tested and stable, we should allow failure and retry.
      log.info 'Exception caught during Contract creation.',
        message: e.message
    end

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

    # TODO: This date is specific to Cloudflare right now, we default to ActivationDate, but not sure if this
    # is the correct one for future merchants. Will need to revisit.
    signed_date = subscription_schedule.metadata["contract_cf_signed_date"]
    signed_date ||= sf_order.ActivatedDate

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
    amount_subtotal ||= normalize_float_amount_in_currency_for_stripe(price.currency, sf_order_item.TotalPrice.to_s, as_decimal: false)

    {
      price: price.id,
      amount_subtotal: amount_subtotal,
      quantity: phase_item.quantity,
      period: {
        start: start_date,
        end: end_date,
      },
      termination_for_convenience: tfc_object,
    }
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
end
