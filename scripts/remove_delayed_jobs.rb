Resque.remove_delayed_selection(SalesforceTranslateRecordJob) do |args|
  puts args.inspect
  username = args[0]
  args.size == 5
end