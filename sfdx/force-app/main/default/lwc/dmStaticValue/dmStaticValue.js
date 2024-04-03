import {LightningElement, api} from 'lwc';

export default class DmStaticValue extends LightningElement {
    @api field;
    @api fieldIndex;
    @api sectionIndex;
    @api valueUpdate;

    get isEnum() {
        return this.field.enum !== undefined && this.field.enum instanceof Array;
    }

    get enumOptions() {
        const options = [];
        if (this.isEnum === false) {
            return options;
        }

        for (let i = 0; i < this.field.enum.length; i++) {
            const opt = this.field.enum[i];
            options.push({value: opt, label: opt});
        }

        return options;
    }

    valueChange(event) {
        this.dispatchEvent(new CustomEvent('update', {
            detail: {
                value: event.target.value,
                sectionIndex: parseInt(this.sectionIndex, 10),
                fieldIndex: parseInt(this.fieldIndex, 10),
            }
        }));
    }
}