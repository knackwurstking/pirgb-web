<script>
  import * as api from "../lib/api"
  import * as utils from "../lib/utils"
  import * as events from "../lib/events"

  import { onMount } from "svelte"

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

  onMount(() => {
    console.log(`[SectionCard.svelte] [onMount] ${host}:${port} [sectionID ${sectionID}]`)
    refresh(null)

    events.global.addEventListener("change", ({ detail }) => {
      if (detail.host !== host
        || detail.port !== port
        || detail.id !== sectionID) {

        return
      }

      // parse data ...
      refresh({ ...detail })
    })
  })

  /**
   * @param {import('../lib/api').Section | null} section
   */
  async function refresh(section = null) {
    console.log(`[SectionCard.svelte] [refresh] host=${host} sectionID=${sectionID}`)
    try {
      if (!section) section = await api.getPWM(host, sectionID)
    } catch (error) {
      console.warn(`[SectionCard.svelte] [${host}:${port}, id: ${sectionID}]`, error)
      pulse = 100
      color = "#ffffff"
      return
    }

    pulse = section.pulse || section.lastPulse || 100
    color = utils.colorToHex(...section.color)
  }
</script>

<fieldset class="section card">
  <legend class="title"> {host} <code style="font-size: 0.75em;">[{sectionID}]</code></legend>

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
        on:focus={(ev) => {
          // @ts-ignore
          ev.target.select()
        }}
      />
    </label>
  </section>

  <section class="actions">
    <button
      class="on"
      on:click={() => {
        api.setPWM(host, sectionID,
          { pulse: pulse, rgbw: utils.hexToColor(color),
        })
      }}
    >
      <span>ON</span>
    </button>
    <button
      class="off"
      on:click={() => {
        api.setPWM(host, sectionID,
          { pulse: 0, rgbw: utils.hexToColor(color) })
      }}
    >
      OFF
    </button>
  </section>

</fieldset>

<style>
  fieldset {
    position: relative;
    margin: 1em 0;
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    border-style: solid;
    padding-top: 2em;
  }

  fieldset legend {
    position: absolute;
    top: 0;
    right: 0;
    padding: 0.25em 1.3em;
    border-bottom: 0.1rem solid var(--border-color);
    border-bottom-right-radius: 0;
  }

  section.content {
    display: flex;
    flex-direction: column;
    place-items: center;
    justify-content: space-evenly;
    width: 10em;
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
  }

  section.actions > * {
    margin: 0.5em;
  }

  section.actions > button {
    font-size: 1rem;
    height: fit-content;
    width: 4em;
    margin-top: 1em;
    display: flex;
    place-items: center;
    justify-content: center;
  }

  label.input {
    display: flex;
    flex-direction: column;
    place-items: center;
  }

  label.input span {
    font-size: 0.70rem;
    text-decoration: underline;
  }

  label.input input.pulse {
    text-align: center;
    width: 5em;
  }

  label.input input[type=color] {
    width: 9.5em;
    height: 2.5em;
    max-width: 10em;
    padding: 0.1em;
  }
</style>
