import getDownloadForSupportUrl from '@salesforce/apex/setupAssistant.getDownloadForSupportUrl';
import { LightningElement, api, track, wire } from 'lwc';
import { ShowToastEvent } from 'lightning/platformShowToastEvent'
import { getErrorMessage } from 'c/utils'
import { CurrentPageReference } from 'lightning/navigation';


export default class DownloadForSupportQuickAction extends LightningElement {
    @api recordId;
    @wire(CurrentPageReference) pageRef;
    @wire(getDownloadForSupportUrl, { recordId: "$recordId" }) exportableSupportConfigDownloadUrl

    @api async invoke() {
        try { 
            window.open(this.exportableSupportConfigDownloadUrl.data, '"_blank"');
        } catch (error) {
            let errorMessage = getErrorMessage(error);
            this.showToast('Download Error', errorMessage, 'error');
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