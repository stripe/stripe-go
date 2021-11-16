import { LightningElement, api } from 'lwc';

export default class Step extends LightningElement {
    @api header = 'Base Step';
    @api containerSize = 'medium';
    @api backLabel = 'Back';
    @api backIconName = 'utility:chevronleft';
    @api backHidden = false;
    @api nextLabel = 'Next';
    @api nextIconName = 'utility:chevronright';
    @api nextDisabled = false;

    get containerClassList() {
        return 'slds-container slds-container_center slds-container_' + this.containerSize;
    }

    fireBack() {
        this.dispatchEvent(new CustomEvent('back', {
            bubbles: true,
            composed: true
        }));
    }

    fireNext() {
        this.dispatchEvent(new CustomEvent('next', {
            bubbles: true,
            composed: true
        }));
    }
}
