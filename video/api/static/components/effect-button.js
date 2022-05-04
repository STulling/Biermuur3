window.customElements.define("effect-button", class extends HTMLElement {
    request = new XMLHttpRequest();

    constructor() {
        super();
    }

    connectedCallback() {
        this.addEventListener("click", this.sendRequest)
        this.innerText = this.title;
    }

    sendRequest() {
        this.request.open("GET", `api/DJ/${this.url}`, true );
        this.request.send(null);
    }

    get title() {
        return this.getAttribute('title') || '';
    }

    get url() {
        return this.getAttribute('url') || '';
    }
});