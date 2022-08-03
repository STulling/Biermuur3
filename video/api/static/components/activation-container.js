window.customElements.define("activation-container", class extends HTMLElement {
    loading = true;
    effects = [];

    constructor() {
        super();

        fetch('api/DJ/activations/nl_NL').then(async response => {
            this.effects = await response.json();
            this.loading = false;
            this.connectedCallback();
        });
    }

    connectedCallback() {
        if (this.loading) return;
        this.effects.map(effect => {
            const effectButton = document.createElement("activation-button")
            effectButton.setAttribute("title", effect.title);
            effectButton.setAttribute("url", effect.url);
            this.append(effectButton);
        });
    }
});