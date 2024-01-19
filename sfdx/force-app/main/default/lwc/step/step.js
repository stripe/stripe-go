import { LightningElement, api, track, wire } from 'lwc';
import getExportableConfigDownloadUrl from '@salesforce/apex/setupAssistant.getExportableConfigDownloadUrl';
import LightningConfirm from 'lightning/confirm';
import LightningAlert from "lightning/alert";
import Debugger from 'c/debugger';

export default class SetupStep extends LightningElement {
    @api stepName = '';
    @api nextLabel = 'Next';
    @api backLabel = 'Back';
    @api backDisabled = false;
    @api nextDisabled = false;
    @api saveDisabled = false;
    @api setupComplete = false;
    @api footerHidden = false;
    @api introButtonLabel = 'Get Started';
    @api illustrationUrl = '';
    @api loading = false;
    @track _showIntro;
    @track useStandardFooter = true;
    @track exportLabel = 'Export Config';
    @wire(getExportableConfigDownloadUrl)
    exportableConfigDownloadUrl

    async exportAction() {
        if (this.exportDisabled) {
            const result = await LightningAlert.open({
                message: 'There are unsaved changes on this page. Please save and then export, or refresh the page to discard changes.',
                theme: 'texture',
                label: 'Unsaved Changes', // this is the header text
            });

            return;
        }

        console.log (this.exportableConfigDownloadUrl)
        window.open(this.exportableConfigDownloadUrl.data, '"_blank"');
    }

    get exportTitle() {
        return this.exportDisabled ? 'Download a copy of the previously saved config' : 'Download a copy of the current config';
    }

    // follow the opposite of next disabled, so we're clear what config is downloading.
    get exportDisabled() {
        return !this.saveDisabled || !this.exportableConfigDownloadUrl || !this.exportableConfigDownloadUrl.data;
    }

    @api 
    get showIntro() {
        return (this.setupComplete ? false : this._showIntro);
    }

    set showIntro(value) {
        this._showIntro = value;
    }

    get componentClasslist() {
        return 'stripe-settings slds-card slds-m-bottom_x-large slds-is-relative' + (this.loading ? ' stripe-is-loading' : '');
    }

    get stepContentClasslist() {
        return 'stripe-step-content' + (this.showIntro ? ' slds-hide' : '');
    }

    hideIntro() {
        this.showIntro = false;
    }

    next() {
        Debugger.log('sending next event');
        this.dispatchEvent(new CustomEvent('next', {
            bubbles: true,
            composed: true
        }));
    }

    back() {
        this.dispatchEvent(new CustomEvent('back', {
            bubbles: true,
            composed: true
        }));
    }

    save() {
        this.dispatchEvent(new CustomEvent('save', {
            bubbles: true,
            composed: true
        }));
    }

    handleFooterSlotChange() {
        this.useStandardFooter = false;
    }
}