<script>
    import { onMount, onDestroy } from "svelte";

    import Section from "../lib/Section.svelte";

    import events from "../lib/events";

    /** @type {Devices} */
    let devices = [];

    /** @type {boolean} */
    export let open = false;

    const onOpen = async () => {
        open = true;
        fetchDevices();
    };

    const onClose = async () => {
        open = false;
    };

    /** @param {CustomEvent<DeviceEvent>} event */
    const onOnline = async (event) => {
        console.log(`[Sections] online`, event.detail);

        const sections = getSections(event.detail.host, event.detail.port);
        for (const section of sections) {
            if (section.element && !section.element.online) {
                section.element.online = true;
            }
        }
    };

    /** @param {CustomEvent<DeviceEvent>} event */
    const onOffline = async (event) => {
        console.log(`[Sections] offline`, event.detail);

        const sections = getSections(event.detail.host, event.detail.port);
        for (const section of sections) {
            if (section.element && section.element.online) {
                section.element.online = false;
            }
        }
    };

    /** @param {CustomEvent<ChangeEvent>} event */
    const onChange = async (event) => {
        console.log(`[Sections] change`, event.detail);

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

    async function fetchDevices() {
        const resp = await fetch("/api/v1/devices", { credentials: "include" });

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
    }

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

    /**
     * @param {string} host
     * @param {number} port
     * @returns {SectionData[]}
     */
    function getSections(host, port) {
        const result = [];
        for (const device of devices) {
            if (device.host === host && device.port === port)
                result.push(...device.sections);
        }
        return result;
    }

    onMount(async () => {
        events.addEventListener("close", onClose);
        events.addEventListener("online", onOnline);
        events.addEventListener("offline", onOffline);
        events.addEventListener("change", onChange);

        try {
            await fetchDevices();
            open = true;
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
                    online={!!device.online && open}
                    {open}
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
    }

    div.device {
        position: relative;
        width: calc(100% - 16px);
        margin: 0 8px;
        padding: 8px;
        border-bottom: 1px solid var(--color-border);
    }

    div.device:last-child {
        border-bottom: none;
    }
</style>
