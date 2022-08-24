<script>
  //import { tick } from "svelte"

  /**
   * Our color set
   * @type {string[][]}
   */
  export let colors = [
    ['#DAAFE9', '#C7DBF5', '#AAD5FB', '#ADE5DA', '#B0EDC3', '#FDF0A4', '#F8D6A2'],
    ['#C47ADA', '#90BAEE', '#75BAFA', '#72D5BF', '#73DE8C', '#FBE66E', '#F5B969'],
    ['#AE44B7', '#5E7ABC', '#5E7ABC', '#4DACA9', '#63B75A', '#EDBD4A', '#EC9740'],
    ['#501B87', '#021B6B', '#0C2794', '#337277', '#2F6A52', '#AE802F', '#AD6127']
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

    //await tick()
    //if (ddActive) {
    //  //document.querySelector('.color-block.active').focus()
    //}
  }

  function clickOutsideDropdown() {
    ddActive = false
  }

  /** @param {string} innerValue */
  function changeValue(innerValue) {
    color = innerValue
    ddActive = false
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
    <input type="text" bind:value={color} />
  </div>

  {#if ddActive}
    <div
      class="values-dropdown"
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
    border: 0.0625rem solid #ccc;
    padding: 0.1875rem;
    border-radius: .2rem;
    margin-right: .4rem;
    background: #fff;
    height: 2.1875rem;
  }

  .caret {
    width: 0;
    height: 0;
    border-left: 0.25rem solid transparent;
    border-right: 0.25rem solid transparent;
    border-top: 0.25rem solid #555;
    position: relative;
    top: 0.625rem;
    margin-left: 0.25rem;
  }

  .caret.top {
    border-left: 0.25rem solid transparent;
    border-right: 0.25rem solid transparent;
    border-bottom: 0.25rem solid #555;
    border-top: none;
  }

  .active {
      box-shadow: inset 0 0 0 0.0625rem #FFF, 0 0 0.1875rem 0.0625rem rgba(0,0,0,0.25);
  }
  
  .fake-focus, input:focus, button:focus  {
      outline: 0;
      box-shadow: 0 0 0 0.125rem #18A0FB;
      border-color: #18A0FB;
  }
  
  input {
    border: 0.0625rem solid #CCC;
    height: 2.1875rem;
    border-radius: .2rem;
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
      background: white;
      border: 0.0625rem solid #CCC;
      border-radius: .3rem;
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
