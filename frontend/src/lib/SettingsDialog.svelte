<script>
    import Location from "../js/location";

    export let active = false;
    let savingData = false;

    /** @param {boolean} state */
    function setActive(state) {
        try {
            savingData = true

            if (!state) {
                Location.config["http:"] = {
                    hostname: httpHostname,
                    port: httpPort,
                }

                Location.config["https:"] = {
                    hostname: httpsHostname,
                    port: httpsPort,
                }

                Location.saveConfig()
            }
        } finally {
            savingData = false 
            active = state;
        }
    }

    let dialog;

    /** @type {string} */
    let httpHostname = Location.config["http:"].hostname;
    /** @type {number} */
    let httpPort = Location.config["http:"].port;

    /** @type {string} */
    let httpsHostname = Location.config["https:"].hostname;
    /** @type {number} */
    let httpsPort = Location.config["https:"].port;
</script>

<div
    bind:this={dialog}
    class="dialog"
    class:active
    on:click={(ev) => {
        if (ev.target === dialog) {
            setActive(false)
        }
    }}
>
    <div class="content" class:saving={savingData}>
        <section>
            <h1>Server Config</h1>
        </section>

        <section class="http">
            <div class="label">HTTP:</div>

            <input
                type="text"
                class="host"
                placeholder="Host"
                bind:value={httpHostname}
            />

            <input
                type="number"
                class="port"
                placeholder="Port"
                bind:value={httpPort}
            />
        </section>

        <section class="https">
            <div class="label">HTTPS:</div>

            <input
                type="text"
                class="host"
                placeholder="Host"
                bind:value={httpsHostname}
            />

            <input
                type="number"
                class="port"
                placeholder="Port"
                value={httpsPort}
            />
        </section>
    </div>
</div>

<style>
    .dialog {
        z-index: 3;
        display: none;
        justify-content: center;
        align-items: center;
        position: absolute;
        top: 0;
        right: 0;
        bottom: 0;
        left: 0;
        background: var(--dialog-background);
        filter: none;
    }

    .dialog.active {
        display: flex;
    }

    .content {
        border: var(--border);
        padding: 0.5rem;
        max-width: 25rem;
        background: var(--surface);
    }

    .content.saving {
        pointer-events: none;
        filter: opacity(0.25);
    }

    section {
        margin: 1rem 0.5rem;
    }

    section:first-child {
        margin: 0.5rem;
    }

    section:last-child {
        margin: 0.5rem;
    }

    h1 {
        width: 100%;
        text-align: center;
        user-select: none;
    }

    .label {
        text-align: right;
        width: 4rem;
        text-align: right;
        font-weight: bold;
        text-decoration: underline;
        display: inline-block;
        user-select: none;
    }

    input {
        padding: 0.25em;
        background-color: transparent;
        color: var(--color);
    }

    input[type="number"]::-webkit-outer-spin-button,
    input[type="number"]::-webkit-inner-spin-button {
        display: none;
    }

    input[type="number"] {
        -moz-appearance: textfield;
        appearance: textfield;
    }

    input.port {
        max-width: 6em;
    }
</style>
