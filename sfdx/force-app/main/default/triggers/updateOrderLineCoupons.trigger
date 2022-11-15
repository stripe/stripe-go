trigger updateOrderLineCoupons on SBQQ__Quote__c (after update) {
  try {
    // for all newly ordered quotes, check if the quote lines have coupon and duplicate to the corresponbding order lines
    for(SBQQ__Quote__c quote : Trigger.new) {
      if (quote.SBQQ__Ordered__c == true && Trigger.oldMap.get(quote.Id).SBQQ__Ordered__c == false) {
        // fetch the related quote lines
        List<SBQQ__QuoteLine__c> quoteLines = [
          SELECT Id
          FROM SBQQ__QuoteLine__c
          WHERE SBQQ__Quote__c = :quote.Id
        ];

        for (SBQQ__QuoteLine__c quoteLine : quoteLines) {   
          // get the corresponding order line for this quote line
          List<OrderItem> orderItem = [
            SELECT Id
            FROM OrderItem
            WHERE SBQQ__QuoteLine__c = :quoteLine.Id
          ];
         
          if (orderItem == null || orderItem.size() == 0)
          {
            continue;
          }
          Id orderItemId =  orderItem.get(0).Id;

          // fetch the Stripe Coupon Quote Line Associations for this quote line
          List<Stripe_Coupon_Beta_Quote_Line_Associatio__c> stripeCouponQuoteLineAssociations = [
            SELECT Id, Stripe_Coupon__c
            FROM Stripe_Coupon_Beta_Quote_Line_Associatio__c
            WHERE Quote_Line__c = :quoteLine.Id
          ];

          if (stripeCouponQuoteLineAssociations == null)
          {
            continue;
          }

          // for each Stripe Coupon Quote Line Association
          for (Stripe_Coupon_Beta_Quote_Line_Associatio__c stripeCouponQuoteLineAssociation: stripeCouponQuoteLineAssociations)
          {
            Stripe_Coupon_Beta__c quoteLineCoupon = [
              SELECT Amount_Off__c, Duration__c, Duration_In_Months__c, Max_Redemptions__c, Name__c, Percent_Off__c
              FROM Stripe_Coupon_Beta__c
              WHERE Id = :stripeCouponQuoteLineAssociation.Stripe_Coupon__c
            ].get(0);
            
            // clone the Stripe Coupon on the quote line, it will have a different Id
            Stripe_Coupon_Beta_Serialized__c clonedCoupon = new Stripe_Coupon_Beta_Serialized__c(
              Amount_Off__c = quoteLineCoupon.Amount_Off__c,
              Duration__c = quoteLineCoupon.Duration__c,
              Duration_In_Months__c = quoteLineCoupon.Duration_In_Months__c,
              Max_Redemptions__c = quoteLineCoupon.Max_Redemptions__c,
              Name__c = quoteLineCoupon.Name__c,
              Percent_Off__c = quoteLineCoupon.Percent_Off__c,
              Original_Stripe_Coupon_Beta_Id__c = quoteLineCoupon.Id
            );
            // insert the cloned Stripe coupon
            Database.insertImmediate((sObject)clonedCoupon);
            
            // create a Stripe Coupon Order Line Association
            Stripe_Coupon_Beta_Order_Item_Associatio__c orderLineStripeCouponAssociation = new Stripe_Coupon_Beta_Order_Item_Associatio__c(
              Stripe_Coupon__c = clonedCoupon.Id,
              Order_Item__c = orderItemId
            );

            // insert this record
            Database.insertImmediate((sObject)orderLineStripeCouponAssociation);
          }
        }
      }
    }
  } catch (Exception e) {
      errorLogger.create('updateOrderLineCouponsTrigger', e);
  }
}