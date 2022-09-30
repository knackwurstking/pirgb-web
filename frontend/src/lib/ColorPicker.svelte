<script>
  import { createEventDispatcher } from "svelte"
  const dispatch = createEventDispatcher()

  /**
   * Our color set
   * @type {string[][]}
   */
  export let colors = [
    ['#ff0000', '#ffff00', '#00ff00', '#00ffff', '#0000ff', '#ff00ff'], // 255 (ff)
    ['#ff3f00', '#3fff00', '#00ff3f', '#003fff', '#3f00ff', '#ff003f'], // 63  (3f)
    ['#ff7f00', '#7fff00', '#00ff7f', '#007fff', '#7f00ff', '#ff007f'], // 127 (7f)
    ['#ffbf00', '#bfff00', '#00ffbf', '#00bfff', '#bf00ff', '#ff00bf'], // 191 (bf)
    // ...
    ['#ff7f7f', '#ffff7f', '#7fff7f', '#7fffff', '#7f7fff', '#ff7fff'], // 7f
    ['#ff3f7f', '#3fff7f', '#7fff3f', '#7f3fff', '#3f7fff', '#ff7f3f'], // 3f,7f
    ['#ff3fbf', '#3fffbf', '#bfff3f', '#bf3fff', '#3fbfff', '#ffbf3f'], // 3f,bf
    ['#ff7fbf', '#7fffbf', '#bfff7f', '#bf7fff', '#7fbfff', '#ffbf7f'], // 7f,bf
    ['#ffffff'],
  ]

  // Initial value
  export let color = "#5e7abc"

  // Keyboard shortcut
  let trigger = "Escape"
  /** @param {KeyboardEvent & { currentTarget: EventTarget & Window }} ev */
  function handleKeydown(ev) {
    if (ev.key == trigger) {
      ddActive = false
    }
  }

  let windowHeight
  let top

  // dropdown active
  export let ddActive = false

  // dropdown height
  let ddHeight = 158

  let inputHeight

  /** @param {MouseEvent} ev */
  function toggleDropdown(ev) {
    console.log("[ColorPicker.svelte] toggleDropdown triggered", { ddActive })

    if ((ev.clientY + inputHeight) < ddHeight
      || (windowHeight - ddHeight - inputHeight - ev.clientY) > 0) {

      top = false
    } else {
      top = true
    }

    ddActive = !ddActive
  }

  /** @param {string} innerValue */
  async function changeValue(innerValue) {
    let change = false
    if (color !== innerValue) {
      change = true
    }

    color = innerValue
    ddActive = false

    if (change) {
      dispatch("change", { color: innerValue })
    }
  }
</script>

<svelte:window bind:innerHeight={windowHeight} on:keydown={handleKeydown} />

<div class="color-picker-holder">
  <button
    bind:clientHeight={inputHeight}
    class="select-color"
    class:fake-focus={ddActive}
    on:click={(e) => toggleDropdown(e)}
  >
    <div style="background: {color};" class="color-block"></div>
  </button>

  {#if ddActive}
    <div
      class="values-dropdown"
      class:top 
      bind:clientHeight={ddHeight}
    >
      <div class="values-dropdown-grid">
        {#each colors as val, index}
          {#each val as innerValue, innerIndex}
            <button
              id="id-{index}-{innerIndex}"
              class="color-block"
              class:active={innerValue === color}
              style="background: {innerValue}"
              on:click={() => changeValue(innerValue)}
            >
            </button>
          {/each}
        {/each}
      </div>
    </div>
  {/if}
</div>

<style>
  .color-picker-holder {
    position: relative;
  }

  button.select-color {
    margin-right: 0.4rem;
    padding: 0.25rem;
    height: 2.1875rem;
    background-color: transparent;
    filter: var(--surface);
    box-shadow: 0.05rem 0.05rem 0.1rem var(--background);
    transform: box-shadow 0.25s ease;
  }

  button.select-color:active {
    box-shadow: 0 0 0.1rem var(--background);
  }

  .active {
    box-shadow: inset 0 0 0 0.125rem var(--special-color, var(--border-color)),
      0 0 0.275rem 0.125rem var(--special-color, rgba(0, 0, 0, 0.25));
  }

  div.color-block,
  button.color-block {
    width: 1.5rem;
    height: 1.5rem;
    line-height: 0;
    font-size: 0;
  }

  .values-dropdown {
    padding: 1em;
    position: absolute;
    top: 2.5rem;
    border: var(--border-width) var(--border-style) var(--border-color);
    border-radius: var(--border-radius);
    background-color: var(--surface);
  }

  .values-dropdown-grid {
    grid-template-columns: repeat(7, 1.5rem);
    grid-template-rows: 1.5rem 1.5rem;
    grid-gap: 0.625rem;
    display: grid;
  }

  .values-dropdown.top {
    top: auto;
    bottom: 2.5rem;
  }

  .values-dropdown button {
    border: none;
  }
</style>
