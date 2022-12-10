<script>
  import { onMount, onDestroy } from "svelte";
  //import {  } from "svelte-heros-v2";

  import events from "../lib/events";

  /** @param {CustomEvent<DeviceEvent>} ev */
  const onOnline = (ev) => {
    // TODO: request sections
  };

  /** @param {CustomEvent<ChangeEvent>} ev */
  const onChange = (ev) => {
    // TODO: update sections data?
  };

  onMount(() => {
    // TODO: read sections from cache
    events.addEventListener("online", onOnline);
    events.addEventListener("change", onChange);
  });

  onDestroy(() => {
    events.removeEventListener("online", onOnline);
    events.removeEventListener("change", onChange);
  });
</script>

<!-- TODO: render sections here (list style) -->
