# frozen_string_literal: true
# typed: true

require_relative './_lib'

class Critic::OrderAmendmentTranslation < Critic::OrderAmendmentFunctionalTest
  before do
    @user = make_user(save: true)
  end

  # this can occur if the SBQQ__RevisedOrderProduct__c is negative because of permission issues in salesforce
  it 'throws an error when a subscription item is negative' do
    # manually nil out the
  end
end
