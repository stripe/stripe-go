import { LightningElement, api, track } from 'lwc';

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
}