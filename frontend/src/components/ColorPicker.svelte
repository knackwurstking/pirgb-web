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
  let ddActive = false

  // dropdown height
  let ddHeight = 158

  let inputHeight

  /**
   * Dispatch event on click outside of node
   */
  export function clickOutside(node) {
    const handleClick = (event) => {
      if (node && !node.contains(event.target) && !event.defaultPrevented) {
        node.dispatchEvent(
          new CustomEvent('clickoutside', node)
        )
      }
    }

    document.addEventListener("click", handleClick, true)

    return {
      destroy() {
        document.removeEventListener('click', handleClick)
      }
    }
  }

  /** @param {MouseEvent} ev */
  async function toggleDropdown(ev) {
    if ((ev.clientY + inputHeight) < ddHeight
      || (windowHeight - ddHeight - inputHeight - ev.clientY) > 0) {

      top = false
    } else {
      top = true
    }

    ddActive = !ddActive
  }

  function clickOutsideDropdown() {
    ddActive = false
  }

  /** @param {string} innerValue */
  function changeValue(innerValue) {
    color = innerValue
    ddActive = false

    if (color !== innerValue) {
      dispatch("change", { color: innerValue })
    }
  }
</script>

<svelte:window bind:innerHeight={windowHeight} on:keydown={handleKeydown} />

<div class="color-picker-holder">
  <div class="color-picker-inner">
    <button
      bind:clientHeight={inputHeight}
      class="select-color"
      class:fake-focus={ddActive}
      on:click={toggleDropdown}
    >
      <div style="display: flex;">
        <div style="background: {color};" class="color-block"></div>
        <div style="margin-right: 0.2rem" class="caret" class:top={top}></div>
      </div>
    </button>
  </div>

  {#if ddActive}
    <div
      class="values-dropdown elevate-200"
      class:top 
      bind:clientHeight={ddHeight}
      use:clickOutside
      on:clickoutside={clickOutsideDropdown}
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

  .color-picker-inner {
    display: flex;
    height: 2.1875rem;
  }

  .select-color {
    border: var(--border-width, 0.0625rem) solid var(--border-color, #ccc);
    padding: 0.1875rem;
    border-radius: .2rem;
    margin-right: .4rem;
    background: var(--input-bg, #fff);
    height: 2.1875rem;
  }

  .caret {
    width: 0;
    height: 0;
    border-left: 0.25rem solid transparent;
    border-right: 0.25rem solid transparent;
    border-top: 0.25rem solid var(--text-color-light, #555);
    position: relative;
    top: 0.625rem;
    margin-left: 0.25rem;
  }

  .caret.top {
    border-left: 0.25rem solid transparent;
    border-right: 0.25rem solid transparent;
    border-bottom: 0.25rem solid var(--text-color-light, #555);
    border-top: none;
  }

  .active {
      box-shadow: inset 0 0 0 0.125rem var(--special-color, #FFF), 0 0 0.275rem 0.125rem var(--special-color, rgba(0,0,0,0.25));
  }
  
  .fake-focus, button:focus  {
      outline: 0;
      box-shadow: 0 0 0 0.0625rem var(--special-color, #18A0FB);
      border-color: var(--border-color, #18A0FB);
  }
  
  .color-block {
    border-radius: .2rem;
    width: 1.5rem;
    height: 1.5rem;
    line-height: 0;
    font-size: 0;
  }
  
  .values-dropdown {
      padding: 1rem;
      position: absolute;
      z-index: 1;
      top: 2.5rem;
      background: var(--input-bg, white);
      border: var(--border-width) var(--border-style) var(--border-color, #CCC);
      border-radius: .3rem;
  }
	
  .values-dropdown-grid {
      grid-template-columns: repeat(6, 1.5rem);
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
