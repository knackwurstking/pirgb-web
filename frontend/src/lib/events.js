// TODO: global event handling

/**
 * @typedef {import("./api").Section} Section
 *
 * @typedef {{
 *  name: string,
 *  data: Section,
 * }} EventData
 */

class GlobalEvents extends EventTarget {
  constructor() {
    super()

    /** @type {WebSocket} */
    this.ws = null

    this.connect()
  }

  // Connect to "/api/events"
  connect() {
    this.ws = new WebSocket(`${location.protocol === "https:" ? "wss" : "ws"}://${window.location.host}/api/events`)
    this.ws.onopen = (ev) => {
      console.log("[onopen]", ev)
    }
    this.ws.onclose = (ev) => {
      console.log("[onclose]", ev)
    }
    this.ws.onerror = (ev) => {
      console.log("[onerror]", ev)
    }
    this.ws.onmessage = (ev) => {
      console.log("[onmessage]", ev)
    }
  }
}

export const global = new GlobalEvents()
