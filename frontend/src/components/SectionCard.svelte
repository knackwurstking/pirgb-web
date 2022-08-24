<script>
  import { onMount, onDestroy } from "svelte"

  import PowerSwitch from "./PowerSwitch.svelte"

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
  $: currentPulse ? powerChecked = true : powerChecked = false

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
    pulse = section.pulse || section.lastPulse || 100
    color = utils.colorToHex(...section.color)
  }
</script>

<fieldset class="section card">
  <legend class="title"> {host} <code>[{sectionID}]</code></legend>
  <!--
  <legend class="title"> {host} <code style="font-size: 0.75em;">[{sectionID}]</code></legend>
  -->
  <pre class={`online-indicator ${online ? "online" : ""}`}>offline</pre>

  <section class="content">
    <label class="input">
      <span>Color</span>
      <input
        type="color"
        bind:value={color}
      />
    </label>

    <label class="input">
      <span>Pulse</span>
      <input
        class="pulse"
        type="number"
        min={0}
        bind:value={pulse}
      />
    </label>
  </section>

  <section class="actions">
    <PowerSwitch
      scale={0.7}
      {color}
      bind:checked={powerChecked}
      on:toggled={
        ({ detail }) => {
          console.log("toggled", detail)
          if (detail.checked) {
            api.setPWM(host, sectionID, { pulse, rgbw: utils.hexToColor(color) })
          } else {
            api.setPWM(host, sectionID, { pulse: 0, rgbw: utils.hexToColor(color) })
          }
        }
      }
    />
  </section>

</fieldset>

<style>
  fieldset {
    margin: 1em 0;
    display: flex;
    padding-top: 2em;
    place-items: center;
  }

  fieldset legend {
    position: absolute;
    top: 0;
    right: 0;
    padding: 0.25em 1.3em;
    border-bottom: var(--border-width) var(--border-style) var(--border-color);
    border-bottom-right-radius: 0;
  }

  fieldset .online-indicator {
    position: absolute;
    top: 0;
    left: 0;
    padding: 0.25em 0.75em;
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
    width: 13em;
  }

  section.content > * {
    margin: 0.5em;
  }

  section.actions {
    display: flex;
    flex-direction: column;
    place-items: center;
    justify-content: space-evenly;
    margin-left: 0.75em;
    overflow: hidden;
    width: 7em;
    overflow: visible;
  }

  /* TODO: make this a component */
  label.input {
    display: flex;
    flex-direction: column;
    place-items: center;
  }

  label.input span {
    font-size: 0.70rem;
    width: 100%;
    text-align: center;
  }

  label.input input.pulse {
    text-align: center;
    width: 5em;
    border-top: 0.1rem solid var(--border-color);
    border-right: none;
    border-left: none;
  }

  label.input input[type=color] {
    width: 9.5em;
    height: 2.5em;
    max-width: 10em;
    padding: 0.1em;
  }
</style>
