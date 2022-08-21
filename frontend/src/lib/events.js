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
      // TODO: dispatch a custom event here ...
      // ...
    }
  }

  /**
   * @param {"change"} type
   * @param {(ev: CustomEvent<any>) => (Promise<void>|void)} callback
   */
  addEventListener(type, callback) {
    super.addEventListener(type, callback)
  }

  /**
   * @param {"change"} type
   * @param {(ev: CustomEvent<any>) => (Promise<void>|void)} callback
   */
  removeEventListener(type, callback) {
    super.removeEventListener(type, callback)
  }

  /**
   * @param {"change"} type
   * @param {any} detail
   */
  dispatchCustomEvent(type, detail) {
    super.dispatchEvent(new CustomEvent(type, { detail }))
  }
}

export const global = new GlobalEvents()
