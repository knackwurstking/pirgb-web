/**
 * @typedef {{
 *  pulse: number,
 *  color: number[], // TODO: renamed: rgbw => color
 * }} PWM
 *
 * @typedef {{
 *  id: number,
 *  pulse: number,
 *  lastPulse: number,
 *  color: number,
 * }} Section // TODO: Section type changed
 *
 * @typedef {{
 *  host: string,
 *  port: number,
 *  groups: string[],
 *  sections: Section[],
 * }} Device
 *
 * @typedef {Device[]} Devices
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
    throw await responseError(resp)
  }

  return
}

/**
 * @param {string} host
 * @param {number} section
 * @returns {Promise<Section>}
 */
export async function getPWM(host, section) {
  const resp = await fetch(`/api/devices/${host}/pwm/${section}`)

  if (!resp.ok) {
    throw await responseError(resp)
  }

  return await resp.json()
}

/**
 * @param {Response} resp
 * @returns {Promise<string>}
 */
async function responseError(resp) {
  return `resp: ${resp.statusText}: ${await resp.text()}`
}
