<script>
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();

  export let min;
  export let max;
  export let value = 100;
  export let color = "var(--accent)";
</script>

<div
  {...$$props}
  class={"slider-container " + ($$props.class || "")}
  style={`--accent: ${color}`}
>
  <!-- div class="slider-value-display">{value}</div -->
  <input
    class="slider"
    type="range"
    min={min}
    max={max}
    bind:value
    on:input={() => dispatch("change", { value: value })}
  />
</div>

<style>
  .slider-container {
    width: 100%;
    max-width: 20vmin;
    display: flex;
    flex-direction: row-reverse;
    place-items: center;
    user-select: none;
  }

  .slider {
    -webkit-appearance: none;
    appearance: none;
    width: 100%;
    height: 0.15em;
    background: var(--accent);
    border-radius: var(--border-radius);
    outline: none;
    opacity: 0.7;
    -webkit-transition: 0.2s;
    transition: opacity 0.2s;
  }

  .slider::-webkit-slider-thumb {
    -webkit-appearance: none;
    appearance: none;
    width: 1.5rem;
    height: 1.5rem;
    border-radius: 50%;
    background: var(--accent);
    cursor: pointer;
  }

  .slider::-moz-range-thumb {
    width: 25px;
    height: 25px;
    border-radius: 50%;
    background: var(--accent);
    cursor: pointer;
  }

  /*
  .slider-value-display {
    font-size: 1.3em;
    width: 3em;
    text-align: center;
  }
  */
</style>
