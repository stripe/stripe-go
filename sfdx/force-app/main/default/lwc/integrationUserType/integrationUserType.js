/**
 * Created by jmather-c on 9/5/23.
 *
 * Note: Currently unused, but may be useful in the future.
 */

import {LightningElement} from 'lwc';

export default class IntegrationUserType extends LightningElement {
    doSelect(event) {
        this.dispatchEvent(new CustomEvent('select', { detail: { value: event.currentTarget.value } }));
    }
}