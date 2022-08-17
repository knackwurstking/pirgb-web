<script>
  import { onMount } from "svelte" 

  import * as api from "../lib/api"
  import * as utils from "../lib/utils"

  /** @type {import('../lib/api').Section} */
  export let section;

  /** @type {number} */
  export let pulse = 100 // TODO: get device section data (pwm) and set Pulse and RGBW
  $: pulse > 100 ? (pulse = 100) : pulse < 0 && (pulse = 0)

  export let color = "#ffffff" // TODO: init color from server (pwm) together with pulse
  $: console.log(
    `[host: ${section.Host}, id: ${section.SectionID}] color: ${utils.hexToRGBW(color)}`
  )

  onMount(() => {
    // TODO: get stat from device
  })
</script>

<fieldset class="section card">
  <legend class="title">{section.Host}</legend>
  <pre>[Section: {section.SectionID}, Port: {section.Port}]</pre>

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
      />
    </label>

    <button
      class="on"
      on:click={() => {
        api.setPWM(section.Host, section.SectionID,
          { pulse: pulse, rgbw: utils.hexToRGBW(color),
        })
      }}
    >
      ON
    </button>
    <button
      class="off"
      on:click={() => {
        api.setPWM(section.Host, section.SectionID,
          { pulse: 0, rgbw: utils.hexToRGBW(color) })
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
    padding: 0.4em;
  }
</style>
