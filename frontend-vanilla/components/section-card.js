import "./section-card.css"
import colorPicker from "./color-picker"

/**
 * @typedev Device
 * @type {import("../lib/api").Device}
 *
 * @typedev Section
 * @type {import("../lib/api").Section}
 */

/**
 * @param {HTMLDivElement} container
 * @param {Device} device
 * @param {Section} section
 */
export function sectionCard(container, device, section) {
    const fieldset = document.createElement("fieldset")

    fieldset.id = `sectionCard-${device.host}-${device.port}-${section.id}`
    fieldset.classList.add("sectionCard")

    fieldset.innerHTML = `
        <legend>${device.host} <code>[${section.id}]</code></legend>

        <pre class="sectionCard-offlineIndicator">Offline</pre>

        <div class="sectionCard-content">
            <div class="sectionCard-colorPickerContainer" style="margin: 0.25rem; margin-left: 1rem;"></div>

            <!-- TODO: pulseSlider component -->
            <!--PulseSlider
                {color}
                min={5}
                max={100}
                bind:value={pulse}
                on:change={async ({ detail }) => {
                    if (currentPulse == 0) return

                    await Api.setPWM(host, sectionID, {
                        pulse: detail.value,
                        color: Color.hexToColor(color),
                    })
                }}
            /-->

            <!-- TODO: powerSwitch component -->
            <div class="sectionCard-powerSwitch"></div>
        </div>
    `;

    colorPicker(fieldset.querySelector(".sectionCard-colorPickerContainer"), {
        oncolorchange: (color) => {
            console.log(`color changed to ${color} for ${device.host}[${section.id}]`)
        }
    })

    container.appendChild(fieldset)
}
