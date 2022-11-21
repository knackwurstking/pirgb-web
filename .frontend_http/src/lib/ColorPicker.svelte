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

  function startColorPicker() {
    const cp = document.createElement("div")
    cp.classList.add("ColorPickerPopup")
    cp.style.setProperty("--special-color", color)
    cp.innerHTML = `
      <div class="ColorPickerPopup-content">
        <div class="ColorPickerPopup-ddGrid" />
      </div>
    `

    const ddGrid = cp.querySelector(".ColorPickerPopup-ddGrid")
    for (let i = 0; i < colors.length; i++) {
      const cr = colors[i]
      for (let i2 = 0; i2 < cr.length; i2++) {
        const c = cr[i2]
        const button = document.createElement("button")
        button.id = `ColorPickerPopup-${i}-${i2}`
        button.classList.add("ColorPickerPopup-colorBlock")
        button.style.background = c
        button.onclick = () => {
          cp.parentNode.removeChild(cp)
          dispatch("colorchange", { color: c })
        }

        if (c === color) {
          button.classList.add("ColorPickerPopup-active")
        }

        ddGrid.appendChild(button)
      }
    }

    cp.addEventListener("click", async (ev) => {
      if (ev.target === cp || ev.target === ddGrid) {
        cp.parentNode.removeChild(cp)
      }
    })

    document.body.appendChild(cp)
  }
</script>

<div
  class="ColorPicker"
  on:click={async () => startColorPicker()}
>
  <div
    style="background: {color};"
    class="ColorPicker-colorBlock"
  />
</div>

<style>
  .ColorPicker,
  :global(.ColorPickerPopup-colorBlock) {
    padding: 0.16rem;
    background: var(--surface);
    width: 2.1em;
    height: 2.1em;
    position: relative;
    border: var(--border);
    border-radius: var(--border-radius);
    cursor: pointer;
  }

  .ColorPicker-colorBlock {
    width: 100%;
    height: 100%;
    border-radius: var(--border-radius);
  }

  :global(.ColorPickerPopup) {
    z-index: 4;
    display: flex;
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    justify-content: center;
    align-items: center;
    background: var(--dialog-background);
  }

  :global(.ColorPickerPopup-ddGrid) {
    width: 20em;
    max-width: 90vw;
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    justify-content: space-evenly;
  }

  :global(.ColorPickerPopup-colorBlock) {
    margin: 0.25rem;
  }

  :global(.ColorPickerPopup-active) {
    box-shadow: 0rem 0rem 0.5rem 0.05rem var(--special-color, red);
  }
</style>
