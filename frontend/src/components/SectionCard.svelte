<script>
  import { onMount, onDestroy } from "svelte"

  import PowerSwitch from "./PowerSwitch.svelte"
  import ColorPicker from "./ColorPicker.svelte"
  import PulseSlider from "./PulseSlider.svelte"

  import * as api from "../lib/api"
  import * as utils from "../lib/utils"
  import * as events from "../lib/events"

  /** @type {string} */
  export let host
  /** @type {number} */
  export let port
  /** @type {number} */
  export let sectionID 
  /** @type {number} */
  export let pulse = 0
  $: pulse < 0 && (pulse = 0)
  export let color = "#ffffff"
  export let online = false

  let powerChecked = false
  let currentPulse = 0

  /**
   * @param {Object} ev
   * @param {import("../lib/events").ChangeEventData} ev.detail */
  const changeEventListener = ({ detail }) => {
    if (detail.host !== host
      || detail.port !== port
      || detail.id !== sectionID) {

      return
    }

    console.log(`[SectionCard] [${host}:${port} (${sectionID})] change event occured`)

    // parse data ...
    refresh({ ...detail })
  }

  /**
   * @param {Object} ev
   * @param {import("../lib/events").OfflineEventData} ev.detail */
  const offlineEventListener = ({ detail }) => {
    if (detail.host !== host
      || detail.port !== port) {

      return
    }

    console.log(`[SectionCard] [${host}:${port} (${sectionID})] offline event occured`)

    online = false
  }

  /**
   * @param {Object} ev
   * @param {import("../lib/events").OfflineEventData} ev.detail */
  const onlineEventListener = ({ detail }) => {
    if (detail.host !== host
      || detail.port !== port) {

      return
    }

    console.log(`[SectionCard] [${host}:${port} (${sectionID})] offline event occured`)

    online = true 
  }

  const openEventListener = () => {
    console.log(`[SectionCard] [${host}:${port} (${sectionID})] websocket open event`)
    refresh(null)
  }

  const closeEventListener = () => {
    console.log(`[SectionCard.svelte] [${host}:${port} (${sectionID})] websocket close event`)
    if (online) online = false
  }

  onMount(() => {
    console.log(`[SectionCard.svelte] [${host}:${port} (${sectionID})] [onMount]`)
    refresh(null)

    // @ts-ignore
    events.global.addEventListener("change", changeEventListener)
    events.global.addEventListener("offline", offlineEventListener)
    events.global.addEventListener("online", onlineEventListener)
    events.global.addEventListener("close", closeEventListener)
    events.global.addEventListener("open", openEventListener)
  })

  onDestroy(() => {
    console.log(`[SectionCard.svelte] [${host}:${port} (${sectionID})] [onDestroy]`)

    // @ts-ignore
    events.global.removeEventListener("change", changeEventListener)
    events.global.removeEventListener("offline", offlineEventListener)
    events.global.removeEventListener("online", onlineEventListener)
    events.global.removeEventListener("close", closeEventListener)
    events.global.removeEventListener("open", openEventListener)
  })

  /**
   * @param {import('../lib/api').Section | null} section
   */
  async function refresh(section = null) {
    console.log(`[SectionCard.svelte] [${host}:${port} (${sectionID})] [refresh]`)

    try {
      if (!section) section = await api.getPWM(host, sectionID)
      if (!online) online = true
    } catch (error) {
      console.warn(`[SectionCard.svelte] [${host}:${port}, id: ${sectionID}]`, error)
      if (online) online = false
      pulse = 100
      color = "#ffffff"
      return
    }

    currentPulse = section.pulse
    powerChecked = !!currentPulse
    pulse = section.pulse || section.lastPulse || 100
    color = utils.colorToHex(...section.color)
  }
</script>

<fieldset style={`--special-color: ${currentPulse > 0 && online ? color : "transparent"};`}>
  <legend class="title"> {host} <code>[{sectionID}]</code></legend>
  <pre class={`online-indicator`} class:online>offline</pre>

  <section class="content">
    <div style="margin: 0.5rem; margin-left: 1rem;">
      <ColorPicker
        bind:color
        on:change={
          async ({ detail }) => {
            if (detail.color) {
              await api.setPWM(
                host, sectionID,
                { pulse: currentPulse, color: utils.hexToColor(detail.color) }
              )
            }
          }
        }
      />
    </div>
    <PulseSlider
      style="margin: 0.5rem;width: 100%;"
      min={5} max={100}
      bind:value={pulse}
      on:change={
        async ({ detail }) => {
          if (currentPulse == 0) return
          await api.setPWM(
            host, sectionID,
            { pulse: detail.value, color: utils.hexToColor(color) }
          )
        }
      }
    />
  </section>

  <section class="actions">
    <PowerSwitch
      scale={0.7}
      {color}
      bind:checked={powerChecked}
      on:toggled={
        async ({ detail }) => {
          if (detail.checked) {
            await api.setPWM(host, sectionID, { pulse, color: utils.hexToColor(color) })
          } else {
            await api.setPWM(host, sectionID, { pulse: 0, color: utils.hexToColor(color) })
          }
        }
      }
    />
  </section>
</fieldset>

<style lang="scss">
  // TODO: update to scss
  @use "sass:color";
  @use "../sass/theme";

  $border-width: theme.$border-width;
  $border-style: theme.$border-style;
  $border-color: theme.$border-color;
  $border-radius: theme.$border-radius;

  fieldset {
    margin: 1rem 0;
    display: flex;
    padding-top: 2rem;
    place-items: center;
    box-shadow: 0 0 0.85em 0.1em var(--special-color, transparent);
    transition: box-shadow 0.5s ease-out;
    background-color: theme.$surface;
  }

  fieldset legend {
    position: absolute;
    top: 0;
    right: 0;
    padding: 0.25rem 1.3rem;
    border-bottom-left-radius: $border-radius;
    border-bottom-right-radius: 0;
    border-bottom-width: $border-width;
    border-bottom-color: $border-color;
    border-bottom-style: $border-style;
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

  section.content {
    display: flex;
    flex-direction: column;
    place-items: center;
    justify-content: space-evenly;
    width: 13rem;
  }

  section.actions {
    display: flex;
    flex-direction: column;
    place-items: center;
    justify-content: space-evenly;
    margin-left: 0.75rem;
    width: 7rem;
    overflow: visible;
  }
</style>
