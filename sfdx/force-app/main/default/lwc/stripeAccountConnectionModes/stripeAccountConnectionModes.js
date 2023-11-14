/**
 * Created by jmather-c on 9/5/23.
 */

import {LightningElement} from 'lwc';

export default class StripeAccountConnectionModes extends LightningElement {
    doSelect(event) {
        this.dispatchEvent(new CustomEvent('select', { detail: { value: event.currentTarget.value } }));
    }
}