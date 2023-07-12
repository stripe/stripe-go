/**
 * Created by jocelyntseng on 7/5/23.
 */

trigger MergeAccountTrigger on Account (before delete, after delete) {
    MergeAccountTriggerHandler handler = new MergeAccountTriggerHandler();

    if (Trigger.isBefore) {
        // store the order ids of the orders that belong
        // to the account that is about to be deleted
        handler.saveOriginalOrdersFromAccounts(Trigger.old);
    } else {
        // get a list of "reparented" orders
        // aka the subset of the master account's orders that need to be re-synced
        handler.processDeletedAccounts(Trigger.old);
    }
}