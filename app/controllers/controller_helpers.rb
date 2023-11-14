# frozen_string_literal: true
# typed: strict

module ControllerHelpers
  extend T::Sig
  include Kernel

  include StripeForce::Constants
  include Integrations::ErrorContext

  # https://github.com/stripe/stripe-salesforce/pull/171
  sig { params(raw_namespace: T.nilable(String)).returns(String) }
  protected def subdomain_namespace_from_param(raw_namespace)
    case raw_namespace
    when nil, ""
      Integrations::ErrorContext.report_edge_case("namespace should not be empty")
      SalesforceNamespaceOptions::PRODUCTION.serialize
    when *SalesforceNamespaceOptions.values.map(&:serialize)
      raw_namespace
    else
      raise Integrations::Errors::ImpossibleInternalError.new("Unexpected namespace: #{raw_namespace}")
    end
  end

  sig { params(state: StateEncryptionAlgo::StripeOAuthState).returns(String) }
  protected def build_postmessage_domain_from_state(state)
    namespace = subdomain_namespace_from_param(state.salesforce_namespace)
    subdomain = state.salesforce_instance_subdomain
    iframe_domain = iframe_domain_from_user(state)
    "https://#{subdomain}--#{namespace}.#{iframe_domain}"
  end

  sig { params(state: StateEncryptionAlgo::StripeOAuthState).returns(String) }
  protected def iframe_domain_from_user(state)
    scratch_org_type = StripeForce::Constants::SFInstanceTypes::SCRATCH_ORG.serialize
    is_scratch_org = state.salesforce_instance_type == scratch_org_type

    if is_scratch_org
      "scratch.vf.force.com"
    else
      "visualforce.com"
    end
  end
end
