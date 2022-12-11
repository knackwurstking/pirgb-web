<script>
  import { onMount, onDestroy } from "svelte";
  //import {  } from "svelte-heros-v2";

  import Section from "../lib/Section.svelte";

  import events from "../lib/events";

  /** @type {Sections} */
  let sections = [];

  /** @type {boolean} */
  let open = false;

  const onOpen = async () => {
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

    // parse sections
    const _sections = [];

    for (const device of respData) {
      _sections.push(...device.sections);
    }

    sections = _sections;
  };

  const onClose = async () => {
    open = false;
  };

  /** @param {CustomEvent<DeviceEvent>} event */
  const onOnline = async (event) => {
    // ...
  };

  /** @param {CustomEvent<DeviceEvent>} event */
  const onOffline = async (event) => {
    // ...
  };

  /** @param {CustomEvent<ChangeEvent>} event */
  const onChange = async (event) => {
    // ...
  };

  onMount(() => {
    events.addEventListener("open", onOpen);
    events.addEventListener("close", onClose);
    events.addEventListener("online", onOnline);
    events.addEventListener("offline", onOffline);
    events.addEventListener("change", onChange);
  });

  onDestroy(() => {
    events.removeEventListener("open", onOpen);
    events.removeEventListener("close", onClose);
    events.removeEventListener("online", onOnline);
    events.removeEventListener("offline", onOffline);
    events.removeEventListener("change", onChange);
  });
</script>

<div class="sections-list">
  {#each sections as section}
    <Section {section} {open} />
  {/each}
</div>

<style>
  div.sections-list {
    width: 100%;
    display: flex;
    flex-direction: column;
    place-items: center;
    overflow-x: hidden;
    overflow-y: auto;
    scroll-behavior: smooth;
  }
</style>
