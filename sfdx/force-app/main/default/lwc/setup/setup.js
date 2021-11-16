import isCpqEnabled from '@salesforce/apex/setupAssistant.isCpqEnabled';
import { LightningElement } from 'lwc';

export default class Setup extends LightningElement {
    get wizards() {
        if (!Array.isArray(this._wizards) || !this._wizards.length) {
            this._wizards = Array.from(this.template.querySelectorAll('c-wizard'));
        }
        return this._wizards;
    }

    constructor() {
        super();
        this.template.addEventListener('exit', this.showLanding.bind(this));
        this.template.addEventListener('finish', this.showLanding.bind(this));
    }

    connectedCallback() {
        isCpqEnabled()
        .then(response => {
           let responseData = JSON.parse(response);
            if(responseData.isSuccess) {
                /* let isCpqInstalled = responseData.results.isCpqInstalled;
                if(isCpqInstalled === false) {
                    //Arnold Sync with sean and see how we want to handle this
                    this.showToast('The CPQ package is not installed in this org', 'error');
                } */
            } else { 
                this.showToast(responseData.error, 'error');
            }
        }).catch(error => {
            this.showToast(error.body.message, 'error');
            this.loading = false;
        });
    }

    renderedCallback() {
        if (this._init !== true) {
            this._init = true;
            this.template.querySelector('c-landing').init(this.wizards);
        }
    }

    showLanding() {
        this.template.querySelector('c-landing').show(this.wizards);
    }
    showToast(message, variant, mode) {
        this.dispatchEvent(new CustomEvent('showtoast', {
            bubbles: true,
            composed: true,
            detail: { 
                toast: {
                    message: message,
                    variant: variant,
                    mode: mode
                }
            }
        })); 
    }
}