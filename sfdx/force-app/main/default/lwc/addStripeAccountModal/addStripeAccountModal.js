/**
 * Created by jmather-c on 12/4/23.
 */

import {track} from 'lwc';
import LightningModal from 'lightning/modal';

export default class AddStripeAccountModal extends LightningModal {
    @track accountType = null;

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