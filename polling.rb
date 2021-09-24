class DateTime
  def utc
    utc = new_offset(0)

    Time.utc(
      utc.year, utc.month, utc.day,
      utc.hour, utc.min, utc.sec + utc.sec_fraction
    )
  end
end

class StripeForce::OrderPoller
  def self.perform(user:)
    interactor = self.new(user)
    interactor.perform
  end

  def initialize(user)
    @user = user
  end

  def perform
    sf = @user.sf_client

    updated_orders = sf.get_updated('Order', DateTime.now - 1, DateTime.now)
    # anything else but "ids" in the hash?
    updated_orders = updated_orders["ids"] if updated_orders.is_a?(Hash)

    updated_orders.each do |sf_order_id|
      sf_order = sf.find('Order', sf_order_id)
      StripeForce::Translate.perform(user: @user, sf_object: sf_order)
    end
  end
end