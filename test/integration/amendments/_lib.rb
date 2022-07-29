# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::OrderAmendmentFunctionalTest < Critic::FunctionalTest
  def create_contract_from_order(sf_order)
    sf.update!(SF_ORDER, {
      SF_ID => sf_order.Id,
      SF_ORDER_CONTRACTED => true,
    })

    # the contracted order puts the ID on the contract itself
    # this operation executes async in SF, so we need to wait for the contract to be created

    related_contracts = T.let(nil, T.untyped)

    wait_until do
      related_contracts = sf.query("SELECT Id FROM #{SF_CONTRACT} WHERE SBQQ__Order__c = '#{sf_order.Id}'")
      related_contracts.count > 0
    end

    assert_equal(1, related_contracts.count)

    sf_contract = sf_get(related_contracts.first.Id)
  end

  # this API is tricky: requires empty JSON object and Content-Type set correctly
  def create_quote_data_from_contract_amendment(sf_contract)
    JSON.parse(sf.patch("services/apexrest/SBQQ/ServiceRouter?loader=SBQQ.ContractManipulationAPI.ContractAmender&uid=#{sf_contract.Id}", {}, {"Content-Type" => "application/json"}).body)
  end

  def create_order_from_quote_data(sf_quote_data)
    sf_quote_id = calculate_and_save_cpq_quote(sf_quote_data)

    # the amendment quote process doesn't seem to pick a pricebook, so we need to manually do this
    sf.update!(CPQ_QUOTE, {
      SF_ID => sf_quote_id,
      CPQ_QUOTE_PRICEBOOK => default_pricebook_id,
    })

    create_order_from_cpq_quote(sf_quote_id)
  end
end
