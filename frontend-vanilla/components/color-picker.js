import "./color-picker.css"

/** Color set to pick from
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

/**
 * @param {{
 *  container?: HTMLElement,
 *  activeColor?: string,
 *  oncolorchange?: (color: string) => Promise<void>|void,
 * }} options
 */
export function colorPickerPopup({ container = null, activeColor = null, oncolorchange = null }) {
    if (!container) {
        container = document.body
    }

    // create root (outside|root) container
    const popup = document.createElement("div")
    popup.classList.add("colorPicker-popup-container")
    popup.innerHTML = `
        <div class="colorPicker-popup-content">
            <div class="colorPicker-popup-dropdownGrid"></div>
        </div>
    `
    // handle the outside click (close popup)
    popup.addEventListener("click", async (ev) => {
        // close popup on a outside click
        if (ev.target === popup || ev.target === ddGrid) {
            popup.parentNode.removeChild(popup)
        }
    })

    // create the popup content
    const ddGrid = popup.querySelector(".colorPicker-popup-dropdownGrid")
    for (let i = 0; i < colors.length; i++) {
        const colorRow = colors[i]
        for (let i2 = 0; i2 < colorRow.length; i2++) {
            const color = colorRow[i2]

            const button = document.createElement("button")
            button.id = `colorPicker-popup-${i}-${i2}`
            button.classList.add("colorPicker-popup-colorBlock")
            button.style.background = color
            button.onclick = (ev) => {
                popup.parentNode.removeChild(popup)

                if (oncolorchange) {
                    oncolorchange(color)
                }
            }
            if (color === activeColor) {
                button.classList.add("colorPicker-popup-active")
            }

            ddGrid.appendChild(button)
        }
    }

    container.appendChild(popup)
    return popup
}

/**
 * @param {HTMLDivElement} container
 * @param {{
 *  oncolorchange?: (color: string) => Promise<void>|void,
 * }} options
 */
export default function colorPicker(container, { oncolorchange = null }) {
    const colorPicker = document.createElement("div")
    let color = "#ffffff"

    colorPicker.classList.add("colorPicker-toggle")
    colorPicker.innerHTML = `
        <div class="colorPicker-colorBlock" style="background: ${color};"></div>
    `
    const colorBlock = colorPicker.querySelector(".colorPicker-colorBlock")

    colorPicker.addEventListener("click", async () => {
        colorPickerPopup({
            activeColor: color,
            oncolorchange: (c) => {
                colorBlock.style.background = (color = c)
                if (oncolorchange) {
                    oncolorchange(c)
                }
            }
        })
    })

    container.appendChild(colorPicker)
    return colorPicker
}
