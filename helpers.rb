def create_customer(
  description: "Sample customer for NetSuite",
  metadata: {},
  email: nil,
  account_balance: nil
)
  email ||= "netsuite-#{create_id(:netsuite)}@example.com"

  Stripe::Customer.create({
    description: description,
    email: email,
    metadata: metadata,
    account_balance: account_balance,
  }, @user.stripe_client_credentials)
end