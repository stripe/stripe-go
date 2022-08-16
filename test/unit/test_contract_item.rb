# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::ContractItemTest < Critic::UnitTest
  before do
    @translator = make_translator
  end

  it 'terminates a single line with no additive revisions' do
    order_line = create_mock_salesforce_order_item
    termination_line = create_mock_salesforce_order_item(additional_attributes: {
      SF_ORDER_ITEM_REVISED_ORDER_PRODUCT => order_line.Id,
    })

    contract_item = StripeForce::Translate::ContractItemStructure.from_order_line_and_params(
      order_line,
      {
        quantity: 2,
      }
    )

    termination_contract_item = StripeForce::Translate::ContractItemStructure.from_order_line_and_params(
      termination_line,
      {
        quantity: -2,
      }
    )

    # this is to test my assumptions around object mutation in ruby
    lines = [contract_item]
    terminations = [termination_contract_item]

    calculated_lines = @translator.terminate_subscription_line_items(lines, terminations)
    assert(calculated_lines.all?(&:fully_terminated?))
  end

  it 'terminates a single lines with multiple revisions' do
    order_line = create_mock_salesforce_order_item
    order_line_2 = create_mock_salesforce_order_item(additional_attributes: {
      SF_ORDER_ITEM_REVISED_ORDER_PRODUCT => order_line.Id,
    })
    termination_line = create_mock_salesforce_order_item(additional_attributes: {
      SF_ORDER_ITEM_REVISED_ORDER_PRODUCT => order_line.Id,
    })

    contract_item = StripeForce::Translate::ContractItemStructure.from_order_line_and_params(
      order_line,
      {
        quantity: 1,
      }
    )
    assert_nil(contract_item.revised_order_line_id)
    assert(contract_item.new_order_line?)

    contract_item_2 = StripeForce::Translate::ContractItemStructure.from_order_line_and_params(
      order_line_2,
      {
        quantity: 1,
      }
    )
    assert_equal(order_line.Id, contract_item_2.revised_order_line_id)

    termination_contract_item = StripeForce::Translate::ContractItemStructure.from_order_line_and_params(
      termination_line,
      {
        quantity: -2,
      }
    )

    calculated_lines = @translator.terminate_subscription_line_items(
      [contract_item, contract_item_2],
      [termination_contract_item]
    )

    assert(calculated_lines.all?(&:fully_terminated?))
  end
end
