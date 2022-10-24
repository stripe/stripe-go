@user = StripeForce::User.find(salesforce_account_id: "XXX")
to_requeue = []

Resque.remove_delayed_selection(SalesforceTranslateRecordJob) do |args|
    if args[0] == @user.salesforce_account_id && args[2] != @user.livemode
        to_requeue.append(args)
        next true
    end

    false
end

to_requeue.each do |job_args|
    SalesforceTranslateRecordJob.work(@user, job_args[3])
end

