/**
 * @typedef {{
 *  Host: string,
 *  Port: number,
 *  SectionID: number,
 *  Groups: string[],
 * }} Section
 *
 * @typedef {Section[]} Sections
 *
 * @typedef {{
 *  Host: string,
 *  Port: number,
 *  Sections: number[] | null,
 *  Groups: string[] | null,
 * }} Device
 *
 * @typedef {Device[]} Devices
 *
 * @typedef {{
 *  pulse: number,
 *  rgbw?: number[],
 * }} PWMRequest
 */

/**
 * @returns {Promise<Sections>}
 */
export async function getSections() {
  /** @type {Sections} */
  let sections = []

  // Get devices from the api and parse sections
  const resp = await fetch("/api/devices")

  // Check response
  if (resp.ok) {
    /** @type {Devices | null} */
    const devices = await resp.json()
    if (devices) { // devices could be null
      // Parse devices into sections
      for (let device of devices) {
        // Build sections data from this device for each section in sections
        for (let section of device.Sections) {
          sections.push({
            Host: device.Host,
            Port: device.Port,
            SectionID: section,
            Groups: device.Groups,
          })
        }
      }
    }
  }

  return sections
}

/**
 * @param {string} host
 * @param {number} section
 * @param {PWMRequest} data
 */
export async function setPWM(host, section, data) {
  const resp = await fetch(`/api/devices/${host}/pwm/${section}`, {
    method: "POST",
    body: JSON.stringify(data),
    headers: {
      "Content-Type": "application/json",
    }
  })

  if (!resp.ok) {
    throw `resp: ${resp.statusText}: ${await resp.text()}`
  }

  return
}

// TODO: Get PWM ...
// ...
