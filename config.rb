require 'dotenv/load'

require 'rubygems'
require 'bundler'

module StripeForce

end

Bundler.require(:default, :development)

# CREATE DATABASE stripeforce
DB = Sequel.connect(ENV.fetch('DATABASE'))

require_relative 'user'
require_relative 'translate'
require_relative 'polling'
