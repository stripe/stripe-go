import {LightningElement} from 'lwc';

export default class StripeAccountConnectionModes extends LightningElement {
    doSelect(event) {
        this.dispatchEvent(new CustomEvent('select', { detail: { value: event.currentTarget.value } }));
    }
}