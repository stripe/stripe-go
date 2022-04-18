# frozen_string_literal: true
# typed: false

module SalesforceDebugging
  include StripeForce::Constants
  include Kernel

  def sf_jump(id_or_object)
    sf = @user.sf_client if !sf
    id_or_object = id_or_object.Id if id_or_object.is_a?(Restforce::SObject)

    @user_info = sf.user_info
    `open "#{@user_info['urls']['custom_domain']}/#{id_or_object}"`
  end

  def object_manager
    '/lightning/setup/ObjectManager/PricebookEntry/Details/view'
  end

  def salesforce_instance_type
    result = @user.sf_client.query("SELECT InstanceName, IsSandbox, TrialExpirationDate FROM Organization").first

    if !!result['IsSandbox'] && result['TrialExpirationDate'].nil?
      SFInstanceTypes::SANDBOX
    elsif !!result['IsSandbox']
      SFInstanceTypes::SCRATCH_ORG
    elsif !result['TrialExpirationDate'].nil?
      SFInstanceTypes::TRIAL
    else
      SFInstanceTypes::PRODUCTION
    end
  end
end
