<svelte:options accessors={true} />

<script>
    //import {} from "svelte-heros-v2";

    /** @type {string} */
    export let host;
    /** @type {number} */
    export let port;

    /** @type {number} */
    export let sectionId;
    /** @type {number[]} */
    export let color;
    /** @type {number} */
    export let pulse;
    /** @type {number} */
    export let lastPulse;
    /** @type {Pin[]} */
    export let pins;
    $: console.log("pins:", pins);

    /** @type {boolean} */
    export let open = false;

    /** @type {boolean} */
    export let online = false;

    /** @type {number} */
    let inputPulse = pulse || lastPulse || 100;

    async function toggleOn() {
        const resp = await fetch(`/api/v1/devices/${host}/${sectionId}/pwm`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                pulse: inputPulse || pulse || lastPulse || 100,
            }),
        });

        if (!resp.ok) {
            console.warn(
                `toggle "${host}/${sectionId}" on failed: ${resp.status} (${resp.statusText})`
            );
        }
    }

    async function toggleOff() {
        const resp = await fetch(`/api/v1/devices/${host}/${sectionId}/pwm`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                pulse: 0,
            }),
        });

        if (!resp.ok) {
            console.warn(
                `toggle "${host}/${sectionId}" on failed: ${resp.status} (${resp.statusText})`
            );
        }
    }
</script>

<!-- TODO: online indicator -->

<div class="section" class:open class:online>
    <h3>
        <span style="color: var(--color-accent);">{host}</span>:{port} [ID:
        <span style="color: var(--color-accent);">{sectionId}</span>]
    </h3>
    <section class="info">
        <pre>Color: {color}</pre>
        <pre>Pulse: {pulse || 0}</pre>
    </section>
    <section class="info-pins">
        {#each (pins || []) as pin}
            <section class="pin">
                <code>Pin:{pin.pin}</code>
                <code>Pulse:{pin.pulse}</code>
                <code>Range:{pin.range}</code>
                <code>isRunning:{pin.isRunning}</code>
                <code>colorPulse:{pin.colorPulse}</code>
                <code>colorValue:{pin.colorValue}</code>
            </section>
        {/each}
    </section>
    <div class="actions">
        <button class="off" on:click={toggleOff} disabled={!online}>OFF</button>
        <button class="on" on:click={toggleOn} disabled={!online}>ON</button>
        <input type="number" bind:value={inputPulse} />
    </div>
</div>

<style>
    div.section {
        --color-section-border: rgba(255, 0, 0, 0.75);
        position: relative;
        width: 100%;
        height: fit-content;
        margin: 8px 0;
        border: 1px solid var(--color-section-border);
        width: 100%;
        text-align: center;
    }

    div.section.open {
        --color-section-border: var(--color-border);
    }

    div.section:not(.online) {
        --color-section-border: rgba(255, 255, 0, 0.75);
    }

    div.section h3 {
        text-decoration: underline;
    }

    /* info section */

    div.section section.info-pins section.pin {
        border-bottom: 1px solid var(--color-section-border);
    }

    div.section section.info-pins section.pin:first-child {
        border-top: 1px solid var(--color-section-border);
    }

    div.section section.info-pins section.pin code {
        font-size: 0.85rem;
        font-style: italic;
    }

    div.section div.actions {
        padding: 4px 0;
        display: flex;
        justify-content: center;
    }

    /* actions section */

    div.section div.actions button.on,
    div.section div.actions button.off {
        width: 10ch;
        height: 4ch;
        margin: 4px 8px;
    }

    div.section div.actions button.off {
        background: rgba(255, 0, 0, 0.9);
    }

    div.section div.actions button.on {
        background: rgba(0, 255, 0, 0.9);
    }

    div.section div.actions input {
        width: 8ch;
        height: 4ch;
        margin: 4px 8px;
        background: transparent;
        color: var(--color-primary);
    }
</style>
