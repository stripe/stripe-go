def example_sf_order
  # sf.find('Order', '8015e000000IJ1rAAG')

  # order with recurring item
  # sf.find('Order', '8015e000000IJDxAAO')

  sf.find('Order', '8015e000000IJF5AAO')
end

def create_salesforce_order
  # TODO pull these dynamically
  product_id = '01t5e000003DsarAAC'
  pricebook_entry_id = '01u5e000000jGn6AAE'

  account_id = sf.create!('Account', Name: "REST Customer #{DateTime.now}")
  opportunity_id = sf.create!('Opportunity', {Name: "REST Oppt #{DateTime.now}", "CloseDate": DateTime.now.iso8601, AccountId: account_id, StageName: "Closed/Won"})

  # you can create a quote without *any* fields, which seems completely silly
  quote_id = sf.create!(CPQ_QUOTE, {
    "SBQQ__Opportunity2__c": opportunity_id,
    CPQ_QUOTE_PRIMARY => true
  })

  # https://developer.salesforce.com/docs/atlas.en-us.cpq_api_dev.meta/cpq_api_dev/cpq_api_read_quote.htm
  # get CPQ version of the quote
  cpq_quote_representation = JSON.parse(sf.get("services/apexrest/SBQQ/ServiceRouter?reader=SBQQ.QuoteAPI.QuoteReader&uid=#{quote_id}").body)

  cpq_product_representation = JSON.parse(sf.patch("services/apexrest/SBQQ/ServiceRouter?loader=SBQQ.ProductAPI.ProductLoader&uid=#{product_id}", {
    context: {
      # productId: product_id,
      pricebookId: pricebook_entry_id,
      # currencyCode:
    }.to_json
  }).body)

  # https://gist.github.com/paustint/bd18bd281134a180e014829b49ed043a
  updated_quote = JSON.parse(sf.patch('/services/apexrest/SBQQ/ServiceRouter?loader=SBQQ.QuoteAPI.QuoteProductAdder', {
    context: {
        "quote": cpq_quote_representation,
        "products": [
          cpq_product_representation
        ],
        # "groupKey": 0,
        "ignoreCalculate": true
      # quote: {
      #   record: {
      #     Id: quote_id,
      #     attributes: {
      #       type: CPQ_QUOTE,
      #     }
      #   }
      # },
      # products: [
      #   {
      #     record: {
      #       Id: product_id
      #     }
      #   }
      # ],
      # ignoreCalculate: true
    }.to_json,
  }).body)

  # https://developer.salesforce.com/docs/atlas.en-us.cpq_dev_api.meta/cpq_dev_api/cpq_quote_api_save_final.htm
  sf.post('/services/apexrest/SBQQ/ServiceRouter', {
    "saver": "SBQQ.QuoteAPI.QuoteSaver",
    "model": updated_quote.to_json
  })

  # sf.create!(CPQ_QUOTE_LINE, {
  #   CPQ_QUOTE => quote_id,
  #   CPQ_QUOTE_LINE_PRODUCT => product_id,
  #   CPQ_QUOTE_LINE_PRICEBOOK_ENTRY => pricebook_entry_id
  # })

  # give CPQ some time to calculate...
  sleep(5)

  # it looks like there is additional field validation triggered here when `ordered` is set to true
  sf.update!(CPQ_QUOTE, 'Id' => quote_id, CPQ_QUOTE_ORDERED => true)


  binding.pry
  # contract_id = salesforce_client.create!('Contract', accountId: account_id)
  # order_id = salesforce_client.create!('Order', {Status: "Draft", EffectiveDate: "2021-09-21", AccountId: account_id, ContractId: contract_id})
end