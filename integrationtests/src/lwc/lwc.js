class TestComponent extends LightningElement {
    connectedCallback() {
        this.template.querySelector('div').classList.add(
            'slds-is-fixed',
            'slds-is-absolute',
            'slds-float_right',
            'slds-float_left'
        );
    }

    someFunkyStyleChanges() {
        this.refs.myElement.style.position = 'absolute';
        this.refs.myElement.style.float = 'right';
        this.refs.otherElement.style.float = "left";
        this.refs.otherElement.style.position = "fixed";
    }

    get riskySldsStyle() {
        return 'slds-is-fixed slds-is-absolute slds-float_right slds-float_left';
    }

    get riskyStyle() {
        return 'float: right; float: left; position: absolute; position: fixed';
    }
}