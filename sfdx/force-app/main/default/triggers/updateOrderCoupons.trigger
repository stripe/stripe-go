trigger updateOrderCoupons on SBQQ__Quote__c (after update) {
  try {
    // for all newly ordered quotes, check if the quote has coupon and copy/duplicate to the corresponding order
    for(SBQQ__Quote__c quote : Trigger.new) {
      if (quote.SBQQ__Ordered__c == true && Trigger.oldMap.get(quote.Id).SBQQ__Ordered__c == false) {
          // get the corresponding order for this quote 
          List<Order> order = [
            SELECT Id
            FROM Order
            WHERE SBQQ__Quote__c = :quote.Id
          ];
         
          if (order == null || order.size() == 0)
          {
            continue;
          }
          Id orderId =  order.get(0).Id;

          // fetch the Stripe Coupon Quote Associations for this quote
          List<Stripe_Coupon_Beta_Quote_Association__c> stripeCouponQuoteAssociations = [
            SELECT Id, Stripe_Coupon__c
            FROM Stripe_Coupon_Beta_Quote_Association__c
            WHERE Quote__c = :quote.Id
          ];

          if (stripeCouponQuoteAssociations == null)
          {
            continue;
          }

          // for each Stripe Coupon Quote Association
          for (Stripe_Coupon_Beta_Quote_Association__c stripeCouponQuoteAssociation: stripeCouponQuoteAssociations)
          {
            Stripe_Coupon_Beta__c quoteCoupon = [
              SELECT Amount_Off__c, Duration__c, Duration_In_Months__c, Max_Redemptions__c, Name__c, Percent_Off__c
              FROM Stripe_Coupon_Beta__c
              WHERE Id = :stripeCouponQuoteAssociation.Stripe_Coupon__c
            ].get(0);
            
            // clone the Stripe Coupon on the quote, it will have a different Id
            Stripe_Coupon_Beta_Serialized__c clonedCoupon = new Stripe_Coupon_Beta_Serialized__c(
              Amount_Off__c = quoteCoupon.Amount_Off__c,
              Duration__c = quoteCoupon.Duration__c,
              Duration_In_Months__c = quoteCoupon.Duration_In_Months__c,
              Max_Redemptions__c = quoteCoupon.Max_Redemptions__c,
              Name__c = quoteCoupon.Name__c,
              Percent_Off__c = quoteCoupon.Percent_Off__c,
              Original_Stripe_Coupon_Beta_Id__c = quoteCoupon.Id
            );
            // insert the cloned Stripe coupon
            Database.insertImmediate((sObject)clonedCoupon);
            
            // create a Stripe Coupon Order Association junction object
            Stripe_Coupon_Beta_Order_Association__c orderStripeCouponAssociation = new Stripe_Coupon_Beta_Order_Association__c(
              Stripe_Coupon__c = clonedCoupon.Id,
              Order__c = orderId
            );

            // insert this record
            Database.insertImmediate((sObject)orderStripeCouponAssociation);
          }
      }
    }
  } catch (Exception e) {
      errorLogger.create('updateOrderCouponsTrigger', e);
  }
}