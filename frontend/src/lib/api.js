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

  // TODO: get devices from the api and parse sections
  const resp = await fetch("/api/devices")

  // check response
  if (resp.ok) {
    /** @type {Devices} */
    const devices = await resp.json()
    // TODO: parse devices into sections
    // ...
  }

  return sections
}
