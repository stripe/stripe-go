# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::OrderAmendmentFunctionalTest < Critic::VCRTest
  def create_contract_from_order(sf_order)
    log.info 'creating contract'

    sf.update!(SF_ORDER, {
      SF_ID => sf_order.Id,
      SF_ORDER_CONTRACTED => true,
    })

    # the contracted order puts the order ID on the contract itself
    # this operation executes async in SF, so we need to wait for the contract to be created
    related_contracts = T.let(nil, T.untyped)
    wait_until do
      related_contracts = sf.query("SELECT Id FROM #{SF_CONTRACT} WHERE #{SF_CONTRACT_ORDER_ID} = '#{sf_order.Id}'")
      related_contracts.count > 0
    end

    assert_equal(1, related_contracts.count)

    sf_contract = sf_get(related_contracts.first.Id)
    sf_contract
  end

  # this API is tricky: requires empty JSON object and Content-Type set correctly
  def create_quote_data_from_contract_amendment(sf_contract)
    log.info 'creating amendment quote'
    JSON.parse(sf.patch("services/apexrest/SBQQ/ServiceRouter?loader=SBQQ.ContractManipulationAPI.ContractAmender&uid=#{sf_contract.Id}", {}, {"Content-Type" => "application/json"}).body)
  end

  def create_renewal_quote_from_contract(sf_contract)
    sf_contract_id = sf_contract.Id

    log.info 'creating order renewal'

    # this operation executes async in SF and generates a renewal quote and opportunity
    sf.update!(SF_CONTRACT, {
      SF_ID => sf_contract_id,
      CPQ_CONTRACT_RENEWAL_QUOTED => true,
      # by default, the renewal term matches the contract’s term.
      CPQ_CONTRACT_RENEWAL_TERM => TEST_DEFAULT_CONTRACT_TERM,
    })
    sf_contract.refresh

    # we need to wait for the renewal quote to be created
    renewal_opportunities = T.let(nil, T.untyped)
    wait_until do
      renewal_opportunities = sf.query("SELECT Id, SBQQ__Renewal__c, SBQQ__PrimaryQuote__c FROM #{SF_OPPORTUNITY} WHERE SBQQ__RenewedContract__c = '#{sf_contract_id}' AND SBQQ__PrimaryQuote__c != NULL")
      renewal_opportunities.count > 0
    end

    # confirm this is indeed the renewal opportunity
    assert_equal(1, renewal_opportunities.count)

    renewal_opportunity = renewal_opportunities.first
    # this checkbox is marked TRUE if the Opportunity was generated from the CPQ package from a Contract
    assert(renewal_opportunity.SBQQ__Renewal__c)

    # grab and return the renewal quote linked in the renewal opportunity
    sf_renewal_quote = sf_get(renewal_opportunity.SBQQ__PrimaryQuote__c)
    sf_renewal_quote
  end

  def create_order_from_quote_data(sf_quote_data, activate_order: true)
    log.info 'creating order from quote'

    sf_quote_id = calculate_and_save_cpq_quote(sf_quote_data)

    # the amendment quote process doesn't seem to pick a pricebook, so we need to manually do this
    sf.update!(CPQ_QUOTE, {
      SF_ID => sf_quote_id,
      CPQ_QUOTE_PRICEBOOK => default_pricebook_id,
    })

    create_order_from_cpq_quote(sf_quote_id, activate_order: activate_order)
  end

  # this create a Salesforce order without 'Activating' it
  def create_draft_order_from_quote(sf_quote_id)
    log.info 'creating draft order from quote'

    # manually add the pricebook
    sf.update!(CPQ_QUOTE, {
      SF_ID => sf_quote_id,
      CPQ_QUOTE_PRICEBOOK => default_pricebook_id,
    })

    # order the quote
    sf.update!(CPQ_QUOTE, {
      SF_ID => sf_quote_id,
      CPQ_QUOTE_ORDERED => true,
    })

    # grab the order from the quote
    related_orders = sf.get("/services/data/v52.0/sobjects/#{CPQ_QUOTE}/#{sf_quote_id}/SBQQ__Orders__r")
    sf_order = related_orders.body.first
    sf_order.refresh
    sf_order
  end
end
