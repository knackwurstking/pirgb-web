/**
 * @typedef {{
 *  pulse: number,
 *  rgbw: number[],
 * }} PWM
 *
 * @typedef {Device[]} Devices
 *
 * @typedef {{
 *  host: string,
 *  port: number,
 *  groups: string[],
 *  sections: Section[],
 * }} Device
 *
 * @typedef {{
 *  id: number,
 *  section: Pin[],
 *  host: string,
 *  port: number,
 * }} Section
 *
 * @typedef {{
 *  pin: number,
 *  range: number,
 *  pulse: number,
 *  colorValue: number,
 *  colorPulse: number,
 *  isRunning: boolean,
 * }} Pin
 */

/**
 * @returns {Promise<Devices>}
 */
export async function getDevices() {
  let resp = await fetch("/api/devices")
  /** @type Devices */
  let devices = []

  if (resp.ok) {
    /** @type {Devices} */
    devices = await resp.json()
  }

  return devices
}

/**
 * @param {string} host
 * @param {number} section
 * @param {PWM} data
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
