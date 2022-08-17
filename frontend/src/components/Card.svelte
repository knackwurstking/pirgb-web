<script>
  import * as api from "../lib/api"

  export let section;

  /** @type {number} */
  export let pulse = 100 // TODO: get device section data (pwm) and set Pulse and RGBW
  $: pulse > 100 ? (pulse = 100) : pulse < 0 && (pulse = 0)
</script>

<fieldset class="section card">
  <legend class="title">{section.Host}</legend>
  <pre>[Section: {section.SectionID}, Port: {section.Port}]</pre>

  <!-- TODO: Actions: Pulse / RGBW -->
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
        api.setPWM(section.Host, section.SectionID, {
          pulse: pulse, rgbw: [255,255,255,255],
        })
      }}
    >
      ON
    </button>
    <button
      class="off"
      on:click={() => {
        api.setPWM(section.Host, section.SectionID, { pulse: 0 })
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
</style>
