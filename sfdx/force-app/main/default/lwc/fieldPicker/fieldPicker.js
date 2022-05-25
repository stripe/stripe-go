/* 
    * Purpose and Expected Behavior: 
    Selection input that provides a user with a filterable dropdown list of fields based on a provided Root object. Generates an object relation hash (String) based on a user's field selections. 
    A selection is not considered complete until a user selects a non-lookup field using Field Picker. Selection of a Lookup field will change the picker context to the new object and provide the relevant list of fields for the user to continue selection. 
*/
import { LightningElement, api, track } from 'lwc';

export default class fieldPicker extends LightningElement {
    @api get searchable() {
        return this._searchable;
    }
    set searchable(value) {
        this._searchable = value;
    }

    @api get required() {
        return this._required;
    }
    set required(value) {
        this._required = value;
    }

    @api get readOnly() {
        return this._readonly;
    }
    set readOnly(value) {
        this._readonly = value;
    }

    @api get disabled() {
        return this._disabled;
    }
    set disabled(value) {
        this._disabled = value;
    }
    @track _disabled;

    @api variant = "standard";
    @api placeholder = "--- Select an option ---";
    @api searchBy = "label";
    @api name = "";

    @api threshold = 1;
    @api dropdownHeight = 5;
    @api stickyDropdown = false;

    @api label = "";
    @track _options = [];
    @api
    get options() { 
        return this._options || [];
    }

    set options(options) {
        let formattedOptions = JSON.parse(JSON.stringify(options));
        
        if(this.selectionList.length >= this.maxSelectionDepth) {
            let filteredOptions = formattedOptions.filter(fieldOptions => fieldOptions.type !== 'reference');
            this._options = filteredOptions; 
        } else {
            this._options = formattedOptions;
        }
    }

    @track validity = true;
    @track showMenu = false;

    @track searchTerm = "";
    @track errorMessage = "Complete this field";

    @track focusIndex = -1;
    @track keyboardMode = false;

    @track blurTimeout;
    @track dropdownStyles = "";

    @api
    focus() {
        if (!this.showMenu) {
            this.toggleMenu();
        } else {
            this.template.querySelector(".sp-menu-button").focus();
        }
    }

    @api
    blur(event) {
        if(event) {
            this.preventDefaultStopProp(event);
        }
        this.validity = !this.required || (this.hash && this.required);
        this.blurTimeout = setTimeout(() => {
            this.hideMenu();
            this.clearPath();
            this.display = this.generateDisplay(this.hash);
        }, 100);
    }

    @api
    isValid() {
        return this.validity;
    }

    @api
    setErrorMessage(message) {
        this.errorMessage = message;
    }

    @api
    get hash() {
        return this._hash;
    }

    set hash(value) {
        this._hash = value;
        this.display = this.generateDisplay(this._hash);
    }

    @track valueLabel;

    connectedCallback() {
        this.display = this.generateDisplay(this.hash);
        this.currentObject = this.rootObject;
    }

    repositionDropdown = (() => {
        let inputBoundingRect = this.getBoundingClientRect();
        let dropdownHeight = this.template.querySelector('.slds-dropdown').getBoundingClientRect().height;
        let marginBottom = parseInt(window.getComputedStyle(this.template.querySelector('.slds-form-element')).marginBottom, 10);
        let dropdownStyles = '';
        
        this.topAdjustment = inputBoundingRect.bottom - marginBottom;
        this.translateYAdjustment = window.innerHeight - inputBoundingRect.bottom - dropdownHeight + marginBottom - 3;
        this.anchorDropdownToBottom = (inputBoundingRect.bottom + dropdownHeight - marginBottom + 3) > window.innerHeight;

        if(!this.ticking) {
            this.ticking = true;

            window.requestAnimationFrame(() => {
                dropdownStyles = 'top: ' + this.topAdjustment + 'px; width: ' + inputBoundingRect.width + 'px;';
                
                if(this.anchorDropdownToBottom) {
                    dropdownStyles += ' transform: translate(0, ' + this.translateYAdjustment + 'px);';
                }
                
                this.dropdownStyles = dropdownStyles;
                this.ticking = false;
            });
        }
    }).bind(this);

    getLabel(value) {
        let valueLabel = this.placeholder;
        if(!value) return valueLabel;
        
        try {
            // proxy objects are too slow when dealing with large amounts of options
            let options = JSON.parse(JSON.stringify(this.options));

            if (options && options.length > 0) {
                for (var i = 0; i < options.length; i++) {
                    if (options[i].isOptGroup) {
                        let selectedOption = options[i].options.find(
                            (groupOption) => {
                                return groupOption.value === value;
                            }
                        );

                        if (selectedOption) {
                            valueLabel = selectedOption.label;
                            break;
                        }
                    } else {
                        let selectedOption = options.find((option) => {
                            return option.value === value;
                        });

                        if (selectedOption) {
                            valueLabel = selectedOption.label;
                            break;
                        }
                    }
                }
            }
        } catch (e) {
            console.warn(e);
        }

        return valueLabel;
    }

    select(event) {
        clearTimeout(this.blurTimeout);
        

        let selectedOption = JSON.parse(
            JSON.stringify(event.detail.selectedOption)
        );
        let index = this.focusIndex;

        this.validity = !this.required || (value && this.required);
        if(selectedOption.type !== 'reference') {

            this.hideMenu();
            
            this.appendSelection(selectedOption);
            this.hash = this.generateHash(this.selectionList);
            this.clearPath();

            let evt = new CustomEvent("select", {
                detail: {
                    value: this.hash,
                    index: index,
                    label: this.getLabel(this.hash)
                }
            });
            this.dispatchEvent(evt);
        } else {
            this.appendSelection(selectedOption);
            this.display = this.generateDisplay(this.selectionList);
            if(selectedOption.object) {
                this.currentObject = selectedOption.object;
            }
            this.searchTerm = '';
            this.template.querySelector('.sp-scrollable-list').scrollTop = 0;
            let evt = new CustomEvent("objectchange", {
                detail: {
                    object: selectedOption.object.toString()
                }
            });
            this.dispatchEvent(evt);
            this.dropdownLoading = true;
        }
    }

    handleKeyupOnInput(event) {
        // Handles Key Press events when the main input is focused 
        // Keycode Ref: 27 -> 'Escape'; 32 -> 'Space'; 38 -> 'ArrowUp'; 40 -> 'ArrowDown'
        let keyCode = event.keyCode;
        if (keyCode !== 38 && keyCode !== 40 && keyCode !== 27) {
            if (this.searchTerm !== event.currentTarget.value) {
                this.focusIndex = -1;
                this.searchTerm = event.currentTarget.value;
            }
            //stopProp on space
            if (keyCode === 32) {
                event.stopPropagation();
            }
        }
    }

    handleKeyupOnPicklist(event) {
        // Handles Key Press events when the dropdown list is focused
        // Keycode Ref: 13 -> 'Enter'; 27 -> 'Escape'; 32 -> 'Space'; 38 -> 'ArrowUp'; 40 -> 'ArrowDown'
        if (
            (event.keyCode === 13 || event.keyCode === 32) &&
            this.focusIndex < 0
        ) {
            this.preventDefaultStopProp(event);
            clearTimeout(this.blurTimeout);
            this.toggleMenu();
        } else {
            if (event.keyCode === 27) {
                //close menu on escape
                this.preventDefaultStopProp(event);
                this.hideMenu();
            } else if (event.keyCode === 38) {
                this.preventDefaultStopProp(event);
                if (this.showMenu) {
                    this.focusIndex -= 1;
                    this.navigate(event);
                }
            } else if (event.keyCode === 40) {
                this.preventDefaultStopProp(event);
                if (this.showMenu) {
                    this.focusIndex += 1;
                    this.navigate(event);
                }
            } else if (event.keyCode === 13 && this.focusIndex > -1) {
                this.preventDefaultStopProp(event);
                let picklistOptions = this.template.querySelectorAll(
                    "c-field-picker-option"
                );
                if (this.focusIndex < picklistOptions.length) {
                    picklistOptions[this.focusIndex].select();
                }
            }
        }
    }

    handleKeydownOnInput(event) {
        // Allows Picker to handle certain key press events when main input is focused, preventing unexpected browser behavior 
        // Keycode Ref: 13 -> 'Enter'; 27 -> 'Escape'; 32 -> 'Space'; 38 -> 'ArrowUp'; 40 -> 'ArrowDown'
        let keyCode = event.keyCode;
        // prevent default on enter, esc, up and down. stopProp on space
        if (
            keyCode === 13 ||
            keyCode === 27 ||
            keyCode === 38 ||
            keyCode === 40
        ) {
            event.preventDefault();
        } else if (keyCode === 32) {
            event.stopPropagation();
        }
    }

    handleKeydownOnPicklist(event) {
        // Allows Picker to handle certain key press events when dropdown list is focused, preventing unexpected browser behavior 
        // Keycode Ref: 13 -> 'Enter'; 27 -> 'Escape'; 32 -> 'Space'; 38 -> 'ArrowUp'; 40 -> 'ArrowDown'
        let keyCode = event.keyCode;
        //prevent default on enter, esc, space, up and down
        if (
            keyCode === 13 ||
            keyCode === 27 ||
            keyCode === 32 ||
            keyCode === 38 ||
            keyCode === 40
        ) {
            event.preventDefault();
        }
    }

    focus() {
        this.dispatchEvent(new Event("focus"));
    }

    click(event) {
        clearTimeout(this.blurTimeout);

        this.preventDefaultStopProp(event);

        this.dispatchEvent(new Event("picklistclick"));
        this.display = this.generateDisplay(this.selectionList);
        let evt = new CustomEvent("objectchange", {
            detail: {
                object: this.rootObject
            }
        });
        this.dispatchEvent(evt);
        this.dropdownLoading = true;
        this.toggleMenu();
    }

    disableKeyboardMode() {
        this.keyboardMode = false;
    }

    navigate(event) {
        this.keyboardMode = true;

        let picklistOptions = this.template.querySelectorAll(
            "c-field-picker-option"
        );
        let focusIndex = this.focusIndex;

        if (focusIndex < 0) {
            this.focusIndex = picklistOptions.length - 1;
        } else if (focusIndex === picklistOptions.length) {
            this.focusIndex = 0;
        }
        focusIndex = this.focusIndex;
        let selected = picklistOptions.forEach((option, index) => {
            if (index === focusIndex) {
                option.addFocus();
                this.template.querySelector('.sp-scrollable-list').scrollTop = option.offsetTop;
            } else {
                option.removeFocus();
            }
        });
    }

    updateFocusIndex(event) {
        if (!this.keyboardMode) {
            let focusedValue = event.detail;
            let picklistOptions = this.template.querySelectorAll(
                "c-field-picker-option"
            );
            picklistOptions.forEach((option, index) => {
                if (option.getValue() === focusedValue) {
                    this.focusIndex = index;
                    option.addFocus();
                } else {
                    option.removeFocus();
                }
            });
        }
    }

    preventDefaultStopProp(event) {
        event.preventDefault();
        event.stopPropagation();
    }

    toggleMenu() {
        clearTimeout(this.blurTimeout);

        if (!this.showMenu) {
            this.showMenu = true;

            setTimeout(() => {
                //focus and add slds-has-focus to correct option
                let picklistOptions = this.template.querySelectorAll(
                    "c-field-picker-option"
                );
                picklistOptions.forEach((option) => {
                    option.removeFocus();
                });

                if(this.stickyDropdown) {
                    this.repositionDropdown();

                    window.addEventListener("scroll", this.repositionDropdown, true);
                    window.addEventListener("resize", this.repositionDropdown, true);
                }

                setTimeout(() => {
                    this.template.querySelector(".sp-input").focus();
                }, 50);
            }, 10);
        } else {
            this.hideMenu();
        }
    }

    hideMenu() {
        this.showMenu = false;
        if (this.template.querySelector(".sp-input")) {
            this.template.querySelector(".sp-input").value = "";
        }
        this.focusIndex = -1;
        this.searchTerm = "";
        this.dropdownStyles = "";

        if(this.stickyDropdown) {
            window.removeEventListener("scroll", this.repositionDropdown, true);
            window.removeEventListener("resize", this.repositionDropdown, true);
        }
    }

    clearBlurTimeout() {
        clearTimeout(this.blurTimeout);
    }

    get hasResults() {
        return this.showMenu ? this.filteredOptions.length > 0 : false;
    }

    get filteredOptions() {
        let searchTerm = this.searchTerm || this.searchTerm === 0 ? this.searchTerm.toString().toLowerCase() : "";
        let searchTermChanged = (this.previousSearchTerm && this.previousSearchTerm !== searchTerm) || (!this.previousSearchTerm && searchTerm.length);
        let searchBy = this.searchBy;
        let filteredOptions = [];
        let scrollableList = this.template.querySelector('.sp-scrollable-list');

        this.previousSearchTerm = searchTerm;

        if(this._options && this._options.length > 0 && searchTerm) {
            filteredOptions = JSON.parse(JSON.stringify(this._options)).filter(option => {
                if (option.isOptGroup) {
                    let groupOptions = option.options.filter(
                        (groupOption) => {
                            let label = groupOption.label.toString();
                            let value = groupOption.value.toString();
                            
                            return this.getOptionVisibility(searchTerm, label, value, searchBy);
                        }
                    );
                    option.options = groupOptions;
                    return groupOptions.length > 0;
                } else {
                    let label = option.label.toString();
                    let value = option.value.toString();
                    
                    return this.getOptionVisibility(searchTerm, label, value, searchBy);
                }
            });
        } else {
            filteredOptions = JSON.parse(JSON.stringify(this._options));
        }

        if(scrollableList && searchTermChanged) {
            scrollableList.scrollTop = 0;
        }

        return filteredOptions;
    }

    get labelHidden() {
        return this.variant === "label-hidden";
    }

    get outerClassList() {
        return (
            "slds-form-element" +
            (!this.validity ? " slds-has-error " : "") +
            (this.class ? " " + this.class : "")
        );
    }

    get picklistClassList() {
        return (
            "slds-picklist slds-dropdown-trigger slds-dropdown-trigger_click slds-is-relative" +
            (this.showMenu ? " slds-is-open" : "") +
            (this.disabled
                ? " sp-disabled"
                : this.readOnly
                ? " sp-readonly"
                : "") +
            (this.stickyDropdown ? " sp-picklist_sticky-dropdown" : "")
        );
    }

    get dropdownClassList() {
        return (
            "slds-listbox slds-listbox_vertical slds-dropdown" +
            (this.showMenu ? "" : " slds-hidden") +
            (this.stickyDropdown ? " slds-is-fixed" : "")
        );
    }

    get inputClassList() {
        return (
            "sp-picklist__input slds-col slds-p-horizontal_xx-small" +
            (this.searchable ? "" : " slds-hide")
        );
    }

    get dropdownMenuStyles() {
        return (
            "-webkit-overflow-scrolling: touch; overflow-y: auto;" +
            (this.dropdownHeight == 0
                ? ""
                : "max-height: calc(((0.8125rem * 1.5) + 1rem) * " +
                  this.dropdownHeight +
                  ")")
        );
    }

    get valueLabelClassList() {
        return "slds-truncate" + (this.hash ? "" : " sp-placeholder");
    }

    handleDropdownMouseDown(event) {
        event.stopPropagation();
        setTimeout(() => {
            clearTimeout(this.blurTimeout);
            window.addEventListener('mouseup', this.handleDropdownMouseUp, { once: true, capture: true });
        }, 1);
    }

    handleDropdownMouseUp = (event => {
        if(this.searchable) {
            this.template.querySelector(".sp-input").focus();
        } else {
            this.template.querySelector(".sp-menu-button").focus();
        }
    }).bind(this);

    get showDropdownIcon() {
        return !this.readOnly || (this.disabled && this.readOnly);
    }

    getOptionVisibility(searchTerm, label, value, searchBy) {
        let labelIncludes =
            label &&
            label.toLowerCase().includes(searchTerm);
        let labelExactMatch = 
            label &&
            label.toLowerCase().indexOf(searchTerm) === 0;
        let valueIncludes =
            value &&
            value.toLowerCase().includes(searchTerm);
        let valueExactMatch =
            value &&
            value.toLowerCase().indexOf(searchTerm) === 0;
        
        switch(searchBy) {
            case 'label':
                return labelIncludes;
            case 'label-exact':
                return labelExactMatch;
            case 'value':
                return valueIncludes;
            case 'value-exact':
                return valueExactMatch;
            case 'both':
                return labelIncludes || valueIncludes;
            case 'both-exact':
                return labelExactMatch || valueExactMatch;
        }
    }

    // FIELD PICKER METHODS //
    display = ''; // Display string for input
    @track selectionList = []; // Selection List being constructed; cleared when input not in use
    @track currentObject = ''; // Current object context of field picker; updates as picker traverses objects
    maxSelectionDepth = 4; // Maximum number of lookup fields that can be used in a single selection; used to disable lookup field options to prevent overly complex selections
    @api dropdownLoading; // Tracks whether field picker is waiting for a list of options while the dropdown is open

    @api 
    get rootObject() {
        return this._rootObject;
    }

    set rootObject(value) {
        value != undefined ? this._rootObject = value : this._rootObject = '';
    }

    clearPath() {
        this.selectionList = [];
        this.currentObject = this.rootObject;
    }

    backClick() {
        // Click handler for managing in-progress selection path
        this.selectionList.pop();
        let previousObject = '';
        if (this.selectionList.length > 0) {
            previousObject = this.selectionList[this.selectionList.length - 1].object.toString();
        } else {
            previousObject = this.rootObject;
        }
        let evt = new CustomEvent("objectchange", {
            detail: {
                object: previousObject
            }
        });
        this.dispatchEvent(evt);
        this.dropdownLoading = true;
        this.currentObject = previousObject;
        this.display = this.generateDisplay(this.selectionList);
    }

    clearPathClick() {
        // Click handler for managing in-progress selection path
        this.clearPath();
        this.display = this.generateDisplay(this.selectionList);
        let evt = new CustomEvent("objectchange", {
            detail: {
                object: this.rootObject
            }
        });
        this.dispatchEvent(evt);
        this.dropdownLoading = true;
    }

    clearValue(event) {
        // Click handler for managing a saved selection value
        event.stopPropagation();
        this.hash = '';
        this.display = this.generateDisplay();
        
        let evt = new CustomEvent("objectchange", {
            detail: {
                object: this.rootObject
            }
        });
        this.dispatchEvent(evt);
        this.dropdownLoading = true;
        this.toggleMenu();
    }

    appendSelection(selection) {
        // Remove the last (non-lookup) selection before appending a new selection //
        let lastSelection = this.selectionList[this.selectionList.length - 1];
        if (lastSelection && lastSelection.type != 'reference') {
            this.selectionList.pop();
        }
        this.selectionList.push(selection);
    }

    generateDisplay(value) {
        // Generate the string to be displayed in the UI
        // Display format: '[RootObjectName] > [ObjectName] > [...] > fieldname'
        let displayString = '';
        if(value) {
            let valueArray = [];
            // Convert a hash string to an array before reading
            if (typeof value == 'string') {
                valueArray = value.split('.');
            } else {
                valueArray = value;
            }

            // Start with Root Object for reference (not needed for Hash)
            displayString += `[${this.rootObject.toString()}] > `;
            // If Object, reading selectionList
            if (typeof valueArray[0] == 'object') {
                valueArray.forEach((item) => {
                    if (item.type == 'reference') {
                        // If lookup, use Object formatting and expect another value 
                        displayString += `[${item.relationshipName}] > `;
                    } else {
                        // Else, use Field formatting for non-lookup field
                        displayString += `${item.value}`;
                    }
                })
            } 
            // If not an Object, reading a Hash
            else {
                valueArray.forEach((item, index) => {
                    // If not the last item, separate with ' > '
                    if(index < (valueArray.length - 1)) {
                        // If more than one item in a Hash, everything before the last item is an object reference
                        displayString += `[${item}] > `;
                    } else {
                        // Else reading last item, which will always be a non-lookup field in a Hash
                        displayString += item;
                    }
                })
            }
        } else {
            // If no value provided, reset to placeholder 
            displayString = this.placeholder;
        }
        return displayString;
    }

    generateHash(selectionList) {
        // Convert current list of selected fields to hash
        // Hash format: 'objectname.objectname [...] .objectname.fieldname'
        let hashString = '';
        if(selectionList.length) {
            selectionList.forEach((selectionItem, index) => {
                hashString += `${selectionItem.value}`;

                // If not the last item, separate with '.'
                if(index < (selectionList.length - 1)) {
                    hashString += '.';
                }
            });
        } 
        return hashString;
    }

    get showClearPath() {
        return (this.selectionList.length);
    }

    get showClearValue() {
        return (this.hash);
    }

    get backDisabled() {
        return this.selectionList.length == 0;
    }

    get currentObjectIcon() {
        return (this.currentObject ? `standard:${this.currentObject.toLowerCase()}` : '');
    }

    get enablePopover() {
        return (this.hash && !this.showMenu);
    }

    get isMaxSelection() {
        return (this.selectionList.length >= this.maxSelectionDepth);
    }
}