<script>
  import { createEventDispatcher } from "svelte"
  const dispatch = createEventDispatcher()

  export let min
  export let max
  export let value = 0
</script>

<label {...$$props}>
  <span>{value}</span>
  <input
    type="range"
    {min}
    {max}
    bind:value
    on:change={(ev) => dispatch("change", { value: parseInt(ev.currentTarget.value) })}
  />
</label>

<style>
  label {
    display: flex;
    flex-wrap: nowrap;
    flex-direction: row-reverse;
  }

  label span {
    margin: 0.25rem;
    margin-left: 0.5rem;
    width: 3em;
  }

  label input {
    font-size: inherit;
    -webkit-appearance: none;
    margin: 0.25rem;
    width: 100%;
    color: var(--input-color);
    background: var(--input-bg);
    height: 0.5em;
    transform: translateY(0.4em);
  }

  label input::-webkit-slider-thumb {
    -webkit-appearance: none;
    border: var(--border-width, 1px) var(--border-style, solid) var(--border-color, #000000);
    height: 1.5em;
    width: 0.5em;
    border-radius: var(--border-radius, 3px);
    background: var(--special-color, #ffffff);
  }
</style>
