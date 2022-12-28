<svelte:options accessors={true} />

<script>
    //import {} from "svelte-heros-v2";
    import StatusLED from "./StatusLED.svelte";

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

    /** @type {boolean} */
    export let open = false;

    /** @type {boolean} */
    export let online = false;

    /* Color values */
    /** @type {number} */
    export let cR = color[0] || 0;
    /** @type {number} */
    export let cG = color[1] || 0;
    /** @type {number} */
    export let cB = color[2] || 0;
    /** @type {number} */
    export let cW = color[3] || 0;

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
                color: [cR, cG, cB, cW],
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

<div class="section" class:open class:online>
    <StatusLED
        style={`
            top: 8px;
            left: 8px;
        `}
        active={online}
    />

    <h3>
        <span style="color: var(--color-accent);">{host}</span>:{port} [ID:
        <span style="color: var(--color-accent);">{sectionId}</span>]
    </h3>

    <section class="info">
        <pre>Color: <span
        style={`color: rgb(${color.slice(0, 3).join(",")}); font-weight: bolder;`}
        >{color}</span></pre>
        <pre>Pulse: <span style="font-weight: bolder;">{pulse || 0}</span></pre>
    </section>

    <section class="info-pins">
        {#each pins || [] as pin}
            <section class="pin">
                <code><span>pin:</span> <span>{pin.pin}</span></code>
                <code><span>colorPulse:</span> <span>{pin.colorPulse}</span></code>
                <code><span>isRunning:</span> <span>{pin.isRunning}</span></code>
            </section>
        {/each}
    </section>

    <section class="actions">
        <section class="row color">
            <input
                type="number"
                bind:value={cR}
                min={0}
                max={255}
                placeholder="R"
            />
            <input
                type="number"
                bind:value={cG}
                min={0}
                max={255}
                placeholder="G"
            />
            <input
                type="number"
                bind:value={cB}
                min={0}
                max={255}
                placeholder="B"
            />
            <input
                type="number"
                bind:value={cW}
                min={0}
                max={255}
                placeholder="W"
            />
        </section>

        <section class="row pulse">
            <button class="off" on:click={toggleOff} disabled={!online}
                >OFF</button
            >

            <input type="number" bind:value={inputPulse} />

            <button class="on" on:click={toggleOn} disabled={!online}>ON</button
            >
        </section>
    </section>
</div>

<style>
    .section {
        --color-section-border: var(--color-border);
        position: relative;
        width: 100%;
        height: fit-content;
        margin: 10px 0;
        border: 1px solid var(--color-section-border);
        width: 100%;
        text-align: center;
    }

    .section h3 {
        text-decoration: underline;
    }

    /* info section */

    .section .info-pins {
        width: calc(100% - 16px);
        position: relative;
        left: 8px;
    }

    .section .info-pins section.pin {
        border-bottom: 1px solid var(--color-section-border);
    }

    .section .info-pins section.pin:first-child {
        border-top: 1px solid var(--color-section-border);
    }

    .section .info-pins section.pin code {
        font-size: 0.85rem;
        font-style: italic;
        display: inline-block;
        width: 10ch;
        text-align: left;
        margin: 4px 0.25ch;
    }

    .section .info-pins section.pin code:nth-child(1) {
        width: 9ch;
    }

    .section .info-pins section.pin code:nth-child(2) {
        width: 17ch;
    }

    .section .info-pins section.pin code:nth-child(3) {
        width: 18ch;
    }

    .section .info-pins section.pin code span:first-child {
        text-decoration: underline;
        font-style: italic;
    }

    .section .info-pins section.pin code span:last-child {
        font-weight: bolder;
    }

    .section .actions {
        padding: 4px 0;
        display: flex;
        flex-direction: column;
        justify-content: center;
    }

    /* actions section */
    .section .actions .row.pulse button.on,
    .section .actions .row.pulse button.off {
        width: 10ch;
        height: 4ch;
        margin: 4px 8px;
    }

    .section .actions .row.pulse button.off {
        background: rgba(255, 0, 0, 0.9);
    }

    .section .actions .row.pulse button.on {
        background: rgba(0, 255, 0, 0.9);
    }

    .section .actions .row input[type="number"] {
        width: 6ch;
        height: 4ch;
        margin: 4px 8px;
    }
</style>
