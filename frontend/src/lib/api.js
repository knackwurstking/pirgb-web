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
