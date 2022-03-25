import manualRetry from '@salesforce/apex/setupAssistant.manualRetry';
import { LightningElement, api } from "lwc";
import { ShowToastEvent } from 'lightning/platformShowToastEvent'

export default class ResyncQuickAction extends LightningElement {
    @api recordId;
    isExecuting = false;

    @api async invoke() {
        if (this.isExecuting) {
            this.showToast('Resync Error', 'A resync is already in process', 'error');
            return;
        }
        try { 
            this.isExecuting = true;
            const manualRetryRequest = await manualRetry({
                recordId: this.recordId
            });
            this.data =  JSON.parse(manualRetryRequest);
            if(this.data.isSuccess) {
                if(this.data.results.isSyncRecordDispactched) {
                    this.showToast('Resync Successfully Scheduled', 'This record has successfully been scheduled for resynchronization', 'success');
                }
            } else {
                this.showToast('Resync Error', this.data.error, 'error');
            }
        } catch (error) {
            this.showToast('Resync Error', error.message, 'error');
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