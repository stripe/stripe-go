# frozen_string_literal: true
# typed: true
module StripeSuite
  class Annotator
    include StripeSuite::Log
    include StripeSuite::Translator::Util::Metadata
    include StripeSuite::Translator::Util::Currency
    include StripeSuite::ErrorContext

    attr_reader :user

    def initialize(u)
      @user = u
    end

    def annotate(ns_record, stripe_record=nil)
      static_annotations_key = netsuite_class_type(ns_record.class)

      # static annotations / field defaults
      if user.netsuite_static_annotations && user.netsuite_static_annotations[static_annotations_key]
        assign_values_from_hash(ns_record, user.netsuite_static_annotations[static_annotations_key])
      end

      # in some cases, a single NetSuite record (i.e. Credit Memo) could have multiple corresponding Stripe records (Refund, Credit Note, Customer Balance)
      # in these cases, we'll need to generate 'compound keys' so users can build unique mappings
      dynamic_annotations_key = if stripe_record && stripe_record.class == Stripe::CreditNote
        static_annotations_key + "_credit_note"
      else
        static_annotations_key
      end

      # dynamic / field mapping
      stripe_record_has_metadata = stripe_record && !stripe_record.respond_to?(:deleted) &&
        # some objects don't have metadata (credit note line items) so we should check for it
        stripe_record.respond_to?(:metadata) && stripe_record.metadata

      if stripe_record_has_metadata && user.netsuite_dynamic_annotations&.fetch(dynamic_annotations_key, nil)
        assign_values_from_hash(ns_record, build_dynamic_annotations_hash(stripe_record, dynamic_annotations_key))
      end

      annotate_user_specific_hacks(ns_record, stripe_record)

      # advanced annotations
      annotate_customer_currencies(ns_record)
      annotate_item_subsidiary(ns_record, static_annotations_key)

      # NOTE add memos last in case other annotations overwrite the memo field
      annotate_memos(ns_record, stripe_record)
    end

    def add_memo(ns_record, added_memo, skip_if_exists: false)
      memo_key = if ns_record.class == NetSuite::Records::Customer
        :comments
      else
        :memo
      end

      return if skip_if_exists && ns_record.send(memo_key) && ns_record.send(memo_key).include?(added_memo)

      if ns_record.send(memo_key)
        ns_record.send(:"#{memo_key}=", "#{ns_record.send(memo_key)}. #{added_memo}")
      else
        ns_record.send(:"#{memo_key}=", added_memo.to_s)
      end

      ns_record
    end

    # https://github.com/stripe/stripe-netsuite/issues/284
    def annotate_customer_currencies(ns_record, ignore_feature_flag: false)
      if ns_record.class == NetSuite::Records::Customer && (ignore_feature_flag || feature?(:customer_currencies))
        netsuite_currency_ids = currency_manager.netsuite_currency_mapping.values

        ns_record.currency_list = {currency:
          netsuite_currency_ids.map do |id|
            {
              currency: {internal_id: id},
            }
          end}
      end
    end

    protected

    def feature?(flag)
      @user.feature_enabled?(flag)
    end

    def is_custom_field?(field_name)
      field_name =~ /^(custentity|custbody|custitem|custcol)/
    end

    def is_record_reference_field?(field_name)
      field_name =~ /_id$/
    end

    def is_date_field?(field_name)
      field_name =~ /_date$/
    end

    def build_dynamic_annotations_hash(stripe_record, netsuite_class_key)
      dynamic_assignments = {}

      # build annotations using stripe resource metadata and users table
      # of accepted stripe metadata values and their corresponding NS keys
      user.netsuite_dynamic_annotations[netsuite_class_key].each do |k, v|
        next if v.to_s.empty?

        metadata_value = if feature?(:annotator_v2)
          extract_stripe_resource_field(stripe_record, k)
        else
          stripe_record.metadata[k]
        end

        next unless metadata_value

        # TODO should we only do this if we run into a specific field suffix?
        #      I feel like this is good default behavior. When would someone NOT want this enabled?

        # convert true/false strings into boolean so the correct
        if metadata_value.to_s.downcase.strip == 'true'
          metadata_value = true
        elsif metadata_value.to_s.downcase.strip == 'false'
          metadata_value = false
        end

        dynamic_assignments[v] = metadata_value
      end

      dynamic_assignments
    end

    def assign_values_from_hash(record, field_assignments)
      field_assignments.each do |k, v|
        # TODO need to handle nil values
        # TODO think on custom fields with NetSuite suffix `_id`
        # TODO nofify when an existing field is being overwritten

        method = "#{k}=".to_sym

        if is_record_reference_field?(k)
          ref_method = "#{k[0..-4]}=".to_sym
        elsif is_date_field?(k)
          ref_method = "#{k[0..-6]}=".to_sym
        end

        if is_custom_field?(method)
          # https://jira.corp.stripe.com/browse/DATAIO-170
          if !record.respond_to?(:custom_field_list)
            log.error 'custom field mapping requested, but custom fields are not supported on the record', field: k
            next
          end

          if is_record_reference_field?(k)
            ref_value = NetSuite::Records::CustomRecordRef.new(internal_id: v)
            record.custom_field_list.send(ref_method, ref_value)
          elsif is_date_field?(k)
            ref_value = NetSuite::Utilities.normalize_time_to_netsuite_date(v.to_i)
            record.custom_field_list.send(ref_method, ref_value)
            record.custom_field_list.send(ref_method.to_s[0..-2]&.to_sym).type = 'platformCore:DateCustomFieldRef'
          else
            record.custom_field_list.send(method, v)
          end
        else
          if is_record_reference_field?(k) && record.respond_to?(ref_method)
            ref_value = NetSuite::Records::RecordRef.new(internal_id: v)
            record.send(ref_method, ref_value)
          elsif is_date_field?(k) && record.respond_to?(ref_method)
            ref_value = NetSuite::Utilities.normalize_time_to_netsuite_date(v.to_i)
            record.send(ref_method, ref_value)
          elsif record.respond_to?(method)
            record.send(method, v)
          else
            log.error 'invalid annotator field specified',
              annotator_key: k,
              annotator_value: v,
              ns_record: record
          end
        end
      end
    end

    def annotate_memos(ns_record, stripe_record)
      return if stripe_record.nil?

      # NOTE items do not have any memo field

      if [NetSuite::Records::CreditMemo,
          NetSuite::Records::Invoice,
          NetSuite::Records::Deposit,
          NetSuite::Records::CashRefund,
          NetSuite::Records::Customer,
          NetSuite::Records::CustomerPayment,
          NetSuite::Records::CustomerRefund,
          NetSuite::Records::CustomerDeposit,
          NetSuite::Records::JournalEntry,
          NetSuite::Records::SalesOrder,].include?(ns_record.class)

        stripe_record_tag = if stripe_record.respond_to?(:id)
          stripe_record.id
        elsif stripe_record.respond_to?(:charge)
          # for disputes, the API returns Stripe::Object, we need to method sniff to find the best fit
          stripe_record.charge
        end

        # NOTE we are prefixing ID memos with "Stripe: " to give more context to NS users
        #      who are not familar with Stripe
        add_memo(ns_record, "Stripe: " + stripe_record_tag, skip_if_exists: true)
      end
    end

    # https://github.com/stripe/stripe-netsuite/issues/257
    def annotate_item_subsidiary(ns_record, class_key)
      subsidiary_id = begin
                          user.netsuite_static_annotations[class_key]['subsidiary_id']
                        rescue
                          nil
                        end

      # TODO add additional items that could be passed via metadata?
      if subsidiary_id && [NetSuite::Records::ServiceSaleItem,
                           NetSuite::Records::NonInventoryResaleItem,
                           NetSuite::Records::NonInventorySaleItem,
                           NetSuite::Records::InventoryItem,
                           NetSuite::Records::DiscountItem,].include?(ns_record.class)
       ns_record.subsidiary_list = {
         record_ref: [
           {
             internal_id: subsidiary_id,
           },
         ],
       }
      end
    end

    def annotate_user_specific_hacks(ns_record, stripe_record)
      return if !stripe_record || stripe_record.respond_to?(:deleted)

      # TODO move into shared method!
      stripe_class_key = underscore(stripe_record.class.to_s.split('::').last)

      # TODO https://github.com/stripe/stripe-netsuite/issues/825

      # urbanairship:
      #    { 'customer' => { 'id' => 'custentityua_stripe_cust_id'} }
      #    { 'customer' => { 'metadata.Company_ID' => 'custentity_go_id' }}
      #    { 'invoice' => { 'id' => 'tran_id' }}

      if @user.stripe_user_id == 'acct_16dO35FhBskN50Pf'
        if stripe_class_key == 'customer'
          ns_record.custom_field_list.custentityua_stripe_cust_id = stripe_record.id

          if stripe_record.metadata['Company_ID']
            ns_record.custom_field_list.custentity_go_id = stripe_record.metadata['Company_ID']
          end
        end

        if netsuite_class_type(ns_record.class) == 'invoice'
          ns_record.tran_id = stripe_record.id
        end
      end

      # IC real time
      # dynamic annotations: { 'customer': { 'email' => 'entity_id' } }
      if @user.stripe_user_id == 'acct_1Ap0nWFK9yHq9Ci3'
        if stripe_record.class == Stripe::Customer
          ns_record.company_name = ns_record.email
        end
      end

      # hotschedules
      # dynamic annotations: { 'customer': { 'email' => 'custentity_hs_self_service_email' } }
      if @user.stripe_user_id == 'acct_1BkyC3DlzIgnrUKv'
        if stripe_record.class == Stripe::Customer
          ns_record.custom_field_list.custentity_hs_self_service_email = extract_stripe_resource_field(stripe_record, 'email')
        end

        if stripe_record.class == Stripe::Invoice
          ns_record.email = extract_stripe_resource_field(stripe_record, 'customer.email')
        end
      end

      # eero
      # dynamic annotations: { 'invoice' => { 'subscription.metadata.org_netsuite_id' => 'custbody_parent_customer_id' } }
      if @user.stripe_user_id == 'acct_16BfDrBlqR9vzFLW'
        if stripe_record.class == Stripe::Invoice
          eero_subscription_value = extract_stripe_resource_field(stripe_record, 'subscription.metadata.org_netsuite_id')

          if eero_subscription_value
            ns_record.custom_field_list.custbody_parent_customer = {internal_id: eero_subscription_value}
          end
        end
      end

      # magnify
      # dynamic annotations: { "invoice": { "customer.metadata.country_code" => "custbody_country_code" } }
      if @user.stripe_user_id == 'acct_1A637MKeMffeTwAO'
        if ns_record.class == NetSuite::Records::Invoice
          assign_values_from_hash(ns_record, {
            "custbody_country_code" => extract_stripe_resource_field(stripe_record, "customer.metadata.country_code"),
          })
        end
      end

      # qualtrics
      # dynamic annotations: { 'customer' => { 'email' => 'company_name' } }
      if @user.stripe_user_id == 'acct_1QSGLU0x2Wnatv9naCF5'
        if ns_record.class == NetSuite::Records::Customer
          assign_values_from_hash(ns_record, {
            'company_name' => extract_stripe_resource_field(stripe_record, 'email'),
          })
        end

        if ns_record.class == NetSuite::Records::InvoiceItem
          ns_record.custom_field_list.custcol_origamt = ns_record.amount
          ns_record.custom_field_list.custcol_list_rate = ns_record.amount
        end

        if ns_record.class == NetSuite::Records::Invoice
          ns_record.custom_field_list.custbody_end_user = {internal_id: ns_record.entity.internal_id}
        end
      end

      # spatial networks
      # dynamic annotations: { 'invoice': { 'friendly_invoice_id' => 'tran_id,custbody_fulcrum_friendlyid' } }
      if @user.stripe_user_id == 'sSlifq3EccrFgoEcgrznVndghLCNNrKJ'
        if ns_record.class == NetSuite::Records::Invoice
          ns_record.custom_field_list.custbody_fulcrum_friendlyid = extract_stripe_resource_field(stripe_record, 'metadata.friendly_invoice_id')
        end
      end

      # disqus
      # copy the rev rec fields to a specific custom field that they already had setup on their account :/
      if @user.stripe_user_id == 'acct_18lNBsGT5BD8rOzk'
        if ns_record.class == NetSuite::Records::InvoiceItem
          if ns_record.custom_field_list.respond_to?(:custcol_suitesync_rev_rec_start)
            ns_record.custom_field_list.custcol1 = ns_record.custom_field_list.custcol_suitesync_rev_rec_start.value
            ns_record.custom_field_list.custcol1.type = 'platformCore:DateCustomFieldRef'

            ns_record.custom_field_list.custcol2 = ns_record.custom_field_list.custcol_suitesync_rev_rec_end.value
            ns_record.custom_field_list.custcol2.type = 'platformCore:DateCustomFieldRef'
          end
        end
      end

      # off the grid / acct_1A0D8kDbfJVAfByy
      # dynamic annotations: { 'invoice_item' => { 'invoice.metadata.netsuite_product_line_id' => 'klass_id' }}
      if @user.stripe_user_id == 'acct_1A0D8kDbfJVAfByy' && netsuite_class_type(ns_record.class) == 'invoice_item'
        # NOTE if invoice items are passed in from the invoice, they don't contain a reference to the invoice!
        if stripe_record.class == Stripe::InvoiceLineItem
          stripe_record = Stripe::InvoiceItem.retrieve(stripe_record.id, @user.stripe_client_credentials)
        end

        if stripe_record.respond_to?(:invoice)
          invoice = Stripe::Invoice.retrieve(stripe_record.invoice, @user.stripe_client_credentials)

          if invoice.metadata['netsuite_product_line_id']
            ns_record.klass = {internal_id: invoice.metadata['netsuite_product_line_id']}
          end
        end
      end
    end

  end
end
