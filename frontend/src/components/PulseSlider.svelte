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
    margin: 0.25em;
  }

  label input {
    margin: 0.25em;
  }
</style>
