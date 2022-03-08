import setOrgType from '@salesforce/apex/utilities.setOrgType';
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

    async renderedCallback() {
        if (this._init !== true) {
            this._init = true;
            this.template.querySelector('c-landing').init(this.wizards);
        }
        try {
            const setOrganizationType = await setOrgType();
            this.data =  JSON.parse(setOrganizationType);
            if(!this.data.isSuccess) {
                this.showToast(this.data.error, 'error');
            }
        } catch (error) {
            this.showToast(error.message, 'error');
        }
    }

    showLanding() {
        this.template.querySelector('c-landing').show(this.wizards);
    }
    showToast(event) {
        event.stopPropagation();
        let toast = event.detail.toast;
        this.template.querySelector('c-toast').show(toast.message, toast.variant, toast.mode);
    }

    fetchPicklistValues() {
        setTimeout(() => {
            this.template.querySelector('c-data-mapping-step').getPicklistValuesForMapper(true, ''); 
        }, 4000);
       
    }


}