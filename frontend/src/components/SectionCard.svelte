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
  $: pulse > 100 ? (pulse = 100) : pulse < 0 && (pulse = 0)

  export let color = "#ffffff"

  onMount(() => {
    console.log(`[SectionCard.svelte] [onMount] ${host}:${port} [sectionID ${sectionID}]`)
    refresh(null)

    events.global.addEventListener("change", ({ detail }) => {
      //const { id, pulse, lastPulse, color } = detail
      console.log("[SectionCard.svelte] change event occured:", detail)
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
  <legend class="title">{host}</legend>
  <pre>[Section: {sectionID}, Port: {port}]</pre>

  <input
    type="color"
    bind:value={color}
  />

  <section class="actions">
    <label class="input">
      <span>Pulse</span>
      <input
        class="pulse"
        type="number"
        min={0} max={100}
        bind:value={pulse}
        on:focus={(ev) => {
          // @ts-ignore
          ev.target.select()
        }}
      />
    </label>

    <button
      class="on"
      on:click={() => {
        api.setPWM(host, sectionID,
          { pulse: pulse, color: utils.hexToColor(color),
        })
      }}
    >
      ON
    </button>
    <button
      class="off"
      on:click={() => {
        api.setPWM(host, sectionID,
          { pulse: 0, color: utils.hexToColor(color) })
      }}
    >
      OFF
    </button>
  </section>

</fieldset>

<style>
  fieldset {
    margin: 1em 0;
  }

  fieldset > pre {
    margin-top: 0em;
  }

  section.actions {
    display: flex;
    place-items: center;
  }

  section.actions > * {
    margin: 0.5rem;
  }

  section.actions > button {
    font-size: 1rem;
    height: fit-content;
    margin-top: 1rem;
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

  label.input input {
    text-align: center;
    width: 5em;
  }

  input[type=color] {
    width: 90%;
    height: 2.5em;
    max-width: 10em;
    padding: 0.1em;
  }
</style>
