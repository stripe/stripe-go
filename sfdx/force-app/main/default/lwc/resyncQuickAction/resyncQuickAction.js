import manualRetry from '@salesforce/apex/setupAssistant.manualRetry';
import { LightningElement, api } from "lwc";
import { ShowToastEvent } from 'lightning/platformShowToastEvent'
import { getErrorMessage } from 'c/utils'

export default class ResyncQuickAction extends LightningElement {
    @api recordId;
    isExecuting = false;

    @api async invoke() {
        // This happens when someone clicks the button multiple times
        if (this.isExecuting === true) {
            this.showToast('Resync in progress', 'A resync is already in process.', 'error');
            return;
        }
        try { 
            this.isExecuting = true;
            const manualRetryRequest = await manualRetry({
                recordId: this.recordId
            });
            const responseData = JSON.parse(manualRetryRequest);
             if(responseData.isSuccess && responseData.results.isSyncRecordDispactched) {
                this.showToast('Resync successfully scheduled', 'This record has been successfully queued for resynchronization.', 'success');
            } else {
                this.showToast('Resync Error', responseData.error, 'error');
            }
        } catch (error) {
            let errorMessage = getErrorMessage(error);
            this.showToast('Resync Error', errorMessage, 'error');
        } finally {
            this.isExecuting = false;
        }
    }
    
    showToast(name, error, type) {
        const event = new ShowToastEvent({
            title: name,
            message: error,
            variant: type,
            mode: 'dismissable'
        });
        this.dispatchEvent(event);
    }
}