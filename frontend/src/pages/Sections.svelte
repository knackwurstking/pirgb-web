<script>
    import { onMount, onDestroy } from "svelte";

    import Section from "../lib/Section.svelte";

    import events from "../lib/events";

    /** @type {Devices} */
    let devices = [];

    /** @type {boolean} */
    let open = false;

    const onOpen = async () => {
        console.log(`wss: open`);

        open = true;
        const resp = await fetch("/api/v1/devices");

        if (!resp.ok) {
            console.warn(
                `request devices from "${resp.url}" failed with ${resp.status} (${resp.statusText})`
            );
            return;
        }

        /** @type {Devices} */
        const respData = await resp.json();
        if (!Array.isArray(respData)) {
            console.warn(`unexpected data returned from ${resp.url}`);
            return;
        }

        devices = respData;
    };

    const onClose = async () => {
        console.log(`wss: close`);
        open = false;
    };

    /** @param {CustomEvent<DeviceEvent>} event */
    const onOnline = async (event) => {
        console.log(`wss: online`, event);
        // TODO: ...
    };

    /** @param {CustomEvent<DeviceEvent>} event */
    const onOffline = async (event) => {
        console.log(`wss: offline`, event);
        // TODO: ...
    };

    /** @param {CustomEvent<ChangeEvent>} event */
    const onChange = async (event) => {
        console.log(`wss: change`, event);

        const section = getSection(
            event.detail.host,
            event.detail.port,
            event.detail.id
        );

        if (section) {
            if (section.element) {
                // update section
                section.element.color = event.detail.color;
                section.element.pulse = event.detail.pulse;
                section.element.lastPulse = event.detail.lastPulse;
                section.element.online = true;
                section.element.pins = event.detail.pins;
            }
        }
    };

    /**
     * @param {string} host
     * @param {number} port
     * @param {number} id
     * @returns {SectionData|null}
     */
    function getSection(host, port, id) {
        for (const device of devices) {
            if (device.host !== host || device.port !== port) continue;

            for (const section of device.sections) {
                if (section.id === id) return section;
            }
        }

        return null;
    }

    onMount(async () => {
        events.addEventListener("close", onClose);
        events.addEventListener("online", onOnline);
        events.addEventListener("offline", onOffline);
        events.addEventListener("change", onChange);

        try {
            await onOpen();
        } finally {
            events.addEventListener("open", onOpen);
        }
    });

    onDestroy(() => {
        events.removeEventListener("open", onOpen);
        events.removeEventListener("close", onClose);
        events.removeEventListener("online", onOnline);
        events.removeEventListener("offline", onOffline);
        events.removeEventListener("change", onChange);
    });
</script>

<li class="sections-list">
    {#each devices as device}
        <div class="device">
            {#each device.sections as section}
                <Section
                    bind:this={section.element}
                    host={device.host}
                    port={device.port}
                    sectionId={section.id}
                    color={section.color}
                    pulse={section.pulse}
                    lastPulse={section.lastPulse}
                    pins={section.pins}
                    {open}
                    online={true}
                />
            {/each}
        </div>
    {/each}
</li>

<style>
    li.sections-list {
        width: 100%;
        display: flex;
        flex-direction: column;
        place-items: center;
        overflow-x: hidden;
        overflow-y: auto;
        scroll-behavior: smooth;
    }

    div.device {
        position: relative;
        width: calc(100% - 16px);
        margin: 0 8px;
        border-bottom: 1px solid var(--color-border);
    }
</style>
