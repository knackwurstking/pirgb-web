/**
 * @typedef {{}} Section
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
      // TODO: Parse devices into sections
      // ...
    }
  }

  return sections
}
