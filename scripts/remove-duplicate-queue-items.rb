queue_name = Resque.queue_from_class(SalesforceTranslateRecordJob)
uniques = []

# move everything from delayed to active queue
Resque.enqueue_delayed_selection { true }

# unfortunately, selecting jobs by specific arguments gets hairy: you need to iterate
# over each job, decode the job data, and then
# https://stackoverflow.com/questions/10274974/how-can-i-delete-specific-jobs-from-resque-queue-without-clearing-the-whole-queu
Resque.peek(queue_name, 0, Resque.size(queue_name)).each do |job|
  encoded_job = Resque.encode(job)

  job_parameters = job["args"]
  # ["00D5e000001Aaj4EAC", "acct_1BlklUC7EbgboeEV", false, "Order", "8017V000000gztbQAA"]
  salesforce_account_id, stripe_account_id, livemode, sf_record_type, sf_record_id = job_parameters

  next if sf_record_id.blank?

  if uniques.include?(sf_record_id)
    puts "deleting job\t#{job_parameters}"
    Resque.data_store.remove_from_queue(queue_name, encoded_job)
  else
    puts "keeping\t#{job_parameters}"
    uniques << sf_record_id
  end
end
