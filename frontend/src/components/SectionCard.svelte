<script>
  import { onMount, onDestroy } from "svelte"

  import PowerSwitch from "./PowerSwitch.svelte"
  import ColorPicker from "./ColorPicker.svelte"
  import PulseInput from "./PulseInput.svelte"

  import * as api from "../lib/api"
  import * as utils from "../lib/utils"
  import * as events from "../lib/events"
import PulseSlider from "./PulseSlider.svelte"

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
  $: color && console.log(`color ${color} for host ${host} and section ${sectionID}`)
  export let online = false

  let powerChecked = false
  let currentPulse = 0

  /**
   * @param {Object} ev
   * @param {import("../lib/events").ChangeEventData} ev.detail */
  const changeEventListener = ({ detail }) => {
    console.log(`[SectionCard.svelte] change event occured ${host}:${port}`)

    if (detail.host !== host
      || detail.port !== port
      || detail.id !== sectionID) {

      return
    }

    // parse data ...
    refresh({ ...detail })
  }

  const openEventListener = () => {
    console.log(`[SectionCard.svelte] websocket open event ${host}:${port}`)
    refresh(null)
  }

  const closeEventListener = () => {
    console.log(`[SectionCard.svelte] websocket close event ${host}:${port}`)
    if (online) online = false
  }

  onMount(() => {
    console.log(`[SectionCard.svelte] [onMount] ${host}:${port} [sectionID ${sectionID}]`)
    refresh(null)

    events.global.addEventListener("change", changeEventListener)
    events.global.addEventListener("close", closeEventListener)
    events.global.addEventListener("open", openEventListener)
  })

  onDestroy(() => {
    console.log(`[SectionCard.svelte] [onDestroy] ${host}:${port} [sectionID ${sectionID}]`)

    events.global.removeEventListener("change", changeEventListener)
    events.global.removeEventListener("close", closeEventListener)
    events.global.removeEventListener("open", openEventListener)
  })

  /**
   * @param {import('../lib/api').Section | null} section
   */
  async function refresh(section = null) {
    console.log(`[SectionCard.svelte] [refresh] host=${host} sectionID=${sectionID}`)
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

<fieldset style={`--special-color: ${color};`} class="section elevate-300">
  <legend class="title"> {host} <code>[{sectionID}]</code></legend>
  <pre class={`online-indicator`} class:online>offline</pre>

  <section class="content">
    <div style="margin: 0.5rem; margin-left: 1rem;">
      <ColorPicker
        bind:color
        on:change={
          async ({ detail }) => {
            if (detail.color) {
              await api.setPWM(host, sectionID, { pulse: currentPulse, rgbw: utils.hexToColor(detail.color) })
            }
          }
        }
      />
    </div>
    <!-- TODO: Send button or change input to a range slider (0-100)
    <PulseInput style="margin: 0.5rem;" min={0} bind:value={pulse} />
    -->
    <PulseSlider min={0} max={100} bind:value={pulse} />
  </section>

  <section class="actions">
    <PowerSwitch
      scale={0.7}
      {color}
      bind:checked={powerChecked}
      on:toggled={
        async ({ detail }) => {
          if (detail.checked) {
            await api.setPWM(host, sectionID, { pulse, rgbw: utils.hexToColor(color) })
          } else {
            await api.setPWM(host, sectionID, { pulse: 0, rgbw: utils.hexToColor(color) })
          }
        }
      }
    />
  </section>
</fieldset>

<style>
  fieldset {
    margin: 1rem 0;
    display: flex;
    padding-top: 2rem;
    place-items: center;
  }

  fieldset legend {
    position: absolute;
    top: 0;
    right: 0;
    padding: 0.25rem 1.3rem;
    border-bottom: var(--border-width) var(--border-style) var(--border-color);
    border-bottom-right-radius: 0;
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
