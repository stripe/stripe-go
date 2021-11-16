/*** Strike Error Logger v2.0.0 ***/
public virtual with sharing class errorLogger {
    public static void create(String actionTaken, String message, String lineNumber, String dump) {
        try {
            checkAccess();
        } catch (Exception e) {
            return;
        }
        insert new Error_Log__c(
            Dump__c = dump,
            Message__c = message,
            Line_Number__c = lineNumber,
            Action_Taken__c = actionTaken
        );

        Integer dmlLimit = Limits.getLimitDMLRows() - Limits.getDMLRows();

        if (0 < dmlLimit) {
            delete [SELECT Id
                    FROM Error_Log__c
                    WHERE CreatedDate < LAST_N_DAYS:7
                    LIMIT :dmlLimit];
        }
    }

    public static void checkAccess() {
        dmlManager.checkAccess(new Map<String, Map<String, Set<String>>> {
            'Error_Log__c' => new Map<String, Set<String>> {
                'Read' => new Set<String> {
                    'Action_Taken__c', 'Message__c', 'Line_Number__c', 'Dump__c'
                },
                'Create' => new Set<String> {
                    'Action_Taken__c', 'Message__c', 'Line_Number__c', 'Dump__c'
                },
                'Delete' => new Set<String> {
                    'Action_Taken__c', 'Message__c', 'Line_Number__c', 'Dump__c'
                }
            }
        });
    }

    public class StrikeException extends Exception {
        public override String getStackTraceString() {
            return super.getCause().getStackTraceString();
        }
    }
}
