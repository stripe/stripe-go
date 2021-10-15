# frozen_string_literal: true
# CREATE DATABASE stripeforce
DB = Sequel.connect(ENV.fetch('DATABASE_URL'))
