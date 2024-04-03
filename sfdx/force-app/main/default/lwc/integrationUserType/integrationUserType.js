import {LightningElement} from 'lwc';

export default class IntegrationUserType extends LightningElement {
    doSelect(event) {
        this.dispatchEvent(new CustomEvent('select', { detail: { value: event.currentTarget.value } }));
    }
}