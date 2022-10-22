import "./section-card.css"

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
    fieldset.classList.add("card")
    fieldset.innerHTML = `
        <legend>${device.host} <code>[${section.id}]</code></legend>
        <pre class="offlineIndicator">Offline</pre>
        <div class="content">
            <div class="colorPicker" style="margin: 0.25rem; margin-left: 1rem;">
                <!-- TODO: colorPicker component -->
                <!--ColorPicker bind:color on:change={colorChange} bind:ddActive /-->
            </div>
            <!-- TODO: pulseSlider component -->
            <div class="colorPicker"></div>
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
        </div>
        <div class="actions">
            <!-- TODO: powerSwitch component -->
            <div class="powerSwitch">
            </div>
        </div>
    `;

    container.appendChild(fieldset)
}
