window.customElements.define("effect-container", class extends HTMLElement {
    loading = true;
    effects = [];

    constructor() {
        super();

        fetch('api/DJ/effects/nl_NL').then(async response => {
            this.effects = await response.json();
            this.loading = false;
            this.connectedCallback();
        });
    }

    connectedCallback() {
        if (this.loading) return;
        this.effects.map(effect => {
            const effectButton = document.createElement("effect-button")
            effectButton.setAttribute("title", effect.title);
            effectButton.setAttribute("url", effect.url);
            this.append(effectButton);
        });
    }
});