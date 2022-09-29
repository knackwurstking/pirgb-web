<script>
  import { onMount, onDestroy } from "svelte"

  import PowerSwitch from "./PowerSwitch.svelte"
  import ColorPicker from "./ColorPicker.svelte"
  import PulseSlider from "./PulseSlider.svelte"

  import Api from "../js/api";
  import Color from "../js/color";
  import Events from "../js/events";

  /** @type {string} */
  export let host

  /** @type {number} */
  export let port

  /** @type {number} */
  export let sectionID 

  /** @type {number} */
  export let pulse = 0

  /** @type {string} */
  export let color = "#ffffff"

  /** @type {boolean} */
  export let online = false

  /** @type {boolean} */
  let powerChecked = false

  /** @type {number} */
  let currentPulse = 0

  let active = false

  /**
   * @param {Object} ev
   * @param {import("../js/events").ChangeEventData} ev.detail */
  const changeEventListener = ({ detail }) => {
    if (detail.host !== host
      || detail.port !== port
      || detail.id !== sectionID) {

      return
    }

    console.log(`[SectionCard.svelte] ${host}:${port} (${sectionID}): "change" event occured`)

    // parse data ...
    refresh({ ...detail })
  }

  /**
   * @param {Object} ev
   * @param {import("../js/events").OfflineEventData} ev.detail */
  const offlineEventListener = ({ detail }) => {
    if (detail.host !== host
      || detail.port !== port) {

      return
    }

    console.log(`[SectionCard.svelte] ${host}:${port} (${sectionID}): offline event occured`)

    online = false
  }

  /**
   * @param {Object} ev
   * @param {import("../js/events").OfflineEventData} ev.detail */
  const onlineEventListener = ({ detail }) => {
    if (detail.host !== host
      || detail.port !== port) {

      return
    }

    console.log(`[SectionCard.svelte] ${host}:${port} (${sectionID}): online event occured`)

    online = true 
  }

  const openEventListener = () => {
    console.log(`[SectionCard.svelte] ${host}:${port} (${sectionID}): websocket open event`)
    refresh(null)
  }

  const closeEventListener = () => {
    if (online) {
      console.log(`[SectionCard.svelte] ${host}:${port} (${sectionID}): websocket close event`)
      online = false
    }
  }

  onMount(() => {
    console.log(`[SectionCard.svelte] ${host}:${port} (${sectionID}): onMount`)
    refresh(null)

    Events.addEventListener("change", changeEventListener)
    Events.addEventListener("offline", offlineEventListener)
    Events.addEventListener("online", onlineEventListener)
    Events.addEventListener("close", closeEventListener)
    Events.addEventListener("open", openEventListener)
  })

  onDestroy(() => {
    console.log(`[SectionCard.svelte] ${host}:${port} (${sectionID}): onDestroy`)

    Events.removeEventListener("change", changeEventListener)
    Events.removeEventListener("offline", offlineEventListener)
    Events.removeEventListener("online", onlineEventListener)
    Events.removeEventListener("close", closeEventListener)
    Events.removeEventListener("open", openEventListener)
  })

  /**
   * @param {import('../js/api').Section | null} section
   */
  async function refresh(section = null) {
    console.log(`[SectionCard.svelte] ${host}:${port} (${sectionID}): refresh`)

    try {
      if (!section) section = await Api.getPWM(host, sectionID)
      if (!online) online = true
    } catch (error) {
      console.warn(`[SectionCard.svelte] ${host}:${port} (${sectionID}):`, error)
      if (online) online = false
      pulse = 100
      color = "#ffffff"
      active = false
      return
    }

    currentPulse = section.pulse
    powerChecked = !!currentPulse
    pulse = section.pulse || section.lastPulse || 100
    color = Color.colorToHex(...section.color)
    active = currentPulse > 0 && powerChecked
  }

  async function colorChange({ detail }) {
    if (detail.color) {
      await Api.setPWM(
        host, sectionID,
        { pulse: currentPulse, color: Color.hexToColor(detail.color) }
      )
    }
  }
</script>

<fieldset style={`--special-color: ${(currentPulse > 0 && online) ? color : "transparent"};`} class:active>
  <legend> {host} <code>[{sectionID}]</code></legend>
  <pre class={`online-indicator`} class:online>offline</pre>

  <section class="content">
    <div style="margin: 0.5rem; margin-left: 1rem;">
      <ColorPicker bind:color on:change={colorChange} />
    </div>
    <PulseSlider
      style="margin-left: 1rem;"
      {color}
      min={5} max={100}
      bind:value={pulse}
      on:change={
        async ({ detail }) => {
          if (currentPulse == 0) return

          await Api.setPWM(
            host, sectionID,
            { pulse: detail.value, color: Color.hexToColor(color) }
          )
        }
      }
    />
  </section>

  <section class="actions">
    <PowerSwitch
      scale={0.5}
      {color}
      bind:checked={powerChecked}
      on:toggled={
        async ({ detail }) => {
          if (detail.checked) {
            await Api.setPWM(host, sectionID, { pulse, color: Color.hexToColor(color) })
          } else {
            await Api.setPWM(host, sectionID, { pulse: 0, color: Color.hexToColor(color) })
          }
        }
      }
    />
  </section>
</fieldset>

<style>
  fieldset {
    --surface: transparent;

    display: flex;
    place-items: center;
    transition: box-shadow 0.5s ease-out;
    background-color: var(--surface);
    border: var(--border);
    border-radius: var(--border-radius);
    padding: 0 !important;
    width: fit-content;
    height: 100%;
    transform: scale(0.95);
transition: transform 0.25s ease;
  }

  fieldset::after {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    z-index: -1;
    width: 100%;
    height: 100%;
    transition: opacity 0.5s ease, background-color 0.5s ease;
    filter: blur(2rem);
    opacity: 0.5;
    background-color: var(--special-color);
    animation: pulsate 5s infinite;
  }

  fieldset.active {
    transform: scale(1);
  }

  fieldset legend {
    position: absolute;
    top: 0;
    right: 0;
    padding: 0.2rem 1.3rem;
    border-bottom-left-radius: var(--border-radius);
    border-bottom-right-radius: 0;
    border-bottom: var(--border);
    text-shadow: 0.1em 0.1em 0.2em var(--background)
  }

  fieldset .online-indicator {
    position: absolute;
    top: 0;
    left: 0;
    padding: 0.25rem 0.75rem;
    margin: 0;
    font-size: 0.7rem;
    text-decoration: underline;
    font-style: italic;
    color: red;
    opacity: 1;
    transition: opacity ease 0.4s;
  }

  fieldset .online-indicator.online {
    opacity: 0;
  }

  section {
    display: flex;
    place-items: center;
    justify-content: space-evenly;
    margin: 1.8rem 0.1rem 0.1rem 0.1rem;
    height: calc(100% - 2rem);
    border-radius: var(--border-radius);
  }

  section.content {
    width: 100%;
  }

  section.actions {
    margin-left: 0.25rem;
    place-items: center;
  }

  @keyframes pulsate {
    0% {
      filter: blur(2rem);
    }
    50% {
      filter: blur(1.5rem);
      opacity: 0.25;
    }
    100% {
      filter: blur(2rem);
    }
  }
</style>
