/**
 * Created by jmather-c on 12/4/23.
 */

import {track} from 'lwc';
import LightningModal from 'lightning/modal';
import isUsingV2Middleware from '@salesforce/apex/utilities.isUsingV2Middleware';

export default class AddStripeAccountModal extends LightningModal {
    @track accountType = null;
    @track isV2MiddlewareEnabled;

    async connectedCallback() {
        this.isV2MiddlewareEnabled = await isUsingV2Middleware();
        // accountType will be overriden in addStripeAccount in setupAssistant.
        // leaving this as live for now as to not change too much in SFDX
        if (this.isV2MiddlewareEnabled) {
            this.accountType = 'live';
        }
    }

    get loginDisabled() {
        return !this.accountType;
    }

    updateAccountType(event) {
        this.accountType = event.detail.value;
    }

    handleCancel(event) {
        this.close(null);
    }

    handleSelectType(event) {
        this.close(this.accountType);
    }
}