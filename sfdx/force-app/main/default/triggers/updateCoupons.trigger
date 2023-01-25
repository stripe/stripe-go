trigger updateCoupons on Order (after update) {
  try {
    // for all new Orders, check if the corresponding Quote has coupons and duplicate/copy to the corresponding order
    Map<Id, Id> quoteIdsToOrderIds = new Map<Id, Id>();

    for (Order order : Trigger.new) {
      if (Trigger.oldMap.get(order.Id).Status == 'Draft' && order.Status == 'Activated' && order.SBQQ__Quote__c != null) {
        quoteIdsToOrderIds.put(order.SBQQ__Quote__c, order.Id);
      }
    }

    if (quoteIdsToOrderIds.isEmpty()) {
      return;
    }

    OrderCouponTriggerHandler handler = new OrderCouponTriggerHandler(quoteIdsToOrderIds);
    handler.process();
    handler.persistChanges();
    handler.refreshState(new Map<Id, Id>());

  } catch (Exception e) {
    errorLogger.create('updateOrderCouponsTrigger', e);
  }
}