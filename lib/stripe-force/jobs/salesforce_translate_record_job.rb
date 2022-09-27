# frozen_string_literal: true
# typed: true

class SalesforceTranslateRecordJob < StripeForce::BaseJob

  sig { params(user: StripeForce::User, sf_record_id: String).void }
  def self.work(user, sf_record_id)
    Resque.enqueue(
      self,

      user.salesforce_account_id, user.stripe_account_id, user.livemode, sf_record_id
    )
  end

  def self.translate(user, sf_record)
    self.work(user, sf_record.to_sparam)
  end

  def self.perform(salesforce_account_id, stripe_user_id, livemode, sf_record_id)
    set_error_context

    user = user_reference(salesforce_account_id, stripe_user_id, livemode)

    set_error_context(user: user)

    valid_system_credentials!(user)

    if !user.enabled
      log.info "skipping job, account disabled"
      return
    end

    sf_record_type = StripeForce::Utilities::SalesforceUtil.salesforce_type_from_id(user, sf_record_id)

    log.info 'translation job starting'

    locker = Integrations::Locker.new(user)
    locker.lock_on_user do
      sf_object = backoff { user.sf_client.find(sf_record_type, sf_record_id) }

      StripeForce::Translate.perform(
        user: user,
        locker: locker,
        sf_object: sf_object
      )
    end

    user.persist_refreshed_credentials

    log.info 'translation job complete'
  end
end
