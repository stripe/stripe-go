# CREATE DATABASE stripeforce
DB = Sequel.connect(ENV.fetch('DATABASE_URL'))
