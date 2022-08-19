<script>
  import * as api from "../lib/api"
  import * as utils from "../lib/utils"

  /** @type {number} */
  export let pulse = 100 // TODO: get device section data (pwm) and set Pulse and RGBW
  $: pulse > 100 ? (pulse = 100) : pulse < 0 && (pulse = 0)

  export let color = "#ffffff" // TODO: init color from server (pwm) together with pulse

  /** @type {import('../lib/api').Section} */
  export let section;
  $: {
    if (section) {
      let newPulse = 0
      // grab the biggest pulse value
      for (let pin of section.section) {
        if (pin.pulse > newPulse) {
          newPulse = pin.pulse
        }
      }

      if (pulse > 0) {
        pulse = newPulse
      }
    }
  }
</script>

<fieldset class="section card">
  <legend class="title">{section.host}</legend>
  <pre>[Section: {section.id}, Port: {section.port}]</pre>

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
        api.setPWM(section.host, section.id,
          { pulse: pulse, rgbw: utils.hexToRGBW(color),
        })
      }}
    >
      ON
    </button>
    <button
      class="off"
      on:click={() => {
        api.setPWM(section.host, section.id,
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
    padding: 0.1em;
  }
</style>
