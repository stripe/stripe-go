# HACKY SCRIPT
# Just use this as a starting place to do this manually via a console
include StripeForce::Constants
include StripeForce::Utilities::SalesforceUtil

sf_account_id = "XXX"

@user, user, u = StripeForce::User.find(salesforce_account_id: sf_account_id)

orders = @user.sf_client.query("SELECT Id FROM Order WHERE #{prefixed_stripe_field(GENERIC_STRIPE_ID)} != null")

puts "Wiping Metadata from #{orders.size} Orders"

orders.each do |order|
    puts "Wiping Metadat from Order #{order.Id}"
    wipe_object(order.Id)
end

order_items = @user.sf_client.query("SELECT Id FROM OrderItem WHERE #{prefixed_stripe_field(GENERIC_STRIPE_ID)} != null")

puts "Wiping Metadata from #{order_items.size} OrderItem's"

order_items.each do |order_item|
    puts "Wiping Metadata from OrderItem #{order_item.Id}"
    wipe_object(order_item.Id)
end

products = @user.sf_client.query("SELECT Id FROM Product2 WHERE #{prefixed_stripe_field(GENERIC_STRIPE_ID)} != null")

puts "Wiping Metadata from #{products.size} Product2's"

products.each do |product|
    puts "Wiping Metadata from Product2 #{product.Id}"
    wipe_object(product.Id)
end

prices = @user.sf_client.query("SELECT Id FROM PricebookEntry WHERE #{prefixed_stripe_field(GENERIC_STRIPE_ID)} != null")

puts "Wiping Metadata from #{prices.size} PricebookEntry's"

prices.each do |price|
    puts "Wiping Metadata from PricebookEntry #{price.Id}"
    wipe_object(price.Id)
end

accounts = @user.sf_client.query("SELECT Id FROM Account WHERE #{prefixed_stripe_field(GENERIC_STRIPE_ID)} != null")

puts "Wiping Metadata from #{accounts.size} Accounts"

accounts.each do |account|
    wipe_object(account.Id)
    puts "Wiping Metadata from Account #{account.Id}"
end