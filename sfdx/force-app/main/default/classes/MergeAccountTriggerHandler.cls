/**
 * Created by jocelyntseng on 7/10/23.
 */

public with sharing class MergeAccountTriggerHandler {
    // allows us to bypass the trigger in tests
    @TestVisible
    private static Boolean TRIGGER_ENABLED = true;

    public static Map<Id, Set<Id>> childAccountIdToOrderIdMap = new Map<Id, Set<Id>>();

    public void saveOriginalOrdersFromAccounts(List<Account> accounts) {
        if (TRIGGER_ENABLED == false) {
            return;
        }

        // for each account that is about to be deleted
        for (Account acct : accounts) {
            Id accountId = acct.Id;

            // store a list of its active orders
            List<Order> orders = [
                SELECT Id
                FROM Order
                WHERE (
                    AccountId = :accountId
                    AND Stripe_ID__c <> null
                )
            ];

            // get order ids
            List<Id> orderIds = new List<Id>();
            for (Order order : orders) {
                orderIds.add(order.Id);
            }

            childAccountIdToOrderIdMap.put(accountId, new Set<Id>(orderIds));
        }
    }

    public void processDeletedAccounts(List<Account> accounts) {
        if (TRIGGER_ENABLED == false) {
            return;
        }

        // for each account that has been deleted
        for (Account acct : accounts) {
            Id childAccountId = acct.Id;
            Id primaryAccountId = acct.MasterRecordId;

            // gate to check if account was deleted because of a merge
            if (String.isNotBlank(primaryAccountId)) {

                // fetch list of orders that used to belong to the account
                Set<Id> childOrderIds = childAccountIdToOrderIdMap.get(childAccountId);

                // fetch list of active orders that belong to primary account
                List<Order> primaryOrders = [
                    SELECT Id, Stripe_ID__c
                    FROM Order
                    WHERE (
                        AccountId = :primaryAccountId
                        AND Stripe_ID__c <> null
                    )
                ];

                List<Order> reparentedOrders = findReparentedOrders(childOrderIds, primaryOrders);

                System.debug('reparented orders = ' + reparentedOrders);
            } else {
                // account was deleted
                // question: what to do with the active orders?
                return;
            }
        }
    }

    private List<Order> findReparentedOrders(Set<Id> childOrderIds, List<Order> primaryOrders) {
        List<Order> reparentedOrders = new List<Order>();

        for (Order order : primaryOrders) {
            if (childOrderIds.contains(order.Id)) {
                reparentedOrders.add(order);
            }
        }

        return reparentedOrders;
    }
}