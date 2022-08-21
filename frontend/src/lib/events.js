// TODO: global event handling

/**
 * @typedef {{
 *  name: "change",
 *  data: ChangeEventData,
 * }} EventData
 *
 * @typedef {{
 *  host: string,
 *  port: number,
 *  id: number,
 *  pulse: number,
 *  lastPulse: number,
 *  color: number[],
 * }} ChangeEventData
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
      console.log("[events] [onopen]", ev)
    }
    this.ws.onclose = (ev) => {
      console.log("[events] [onclose]", ev)
      // TODO: auto reconnect on close every 2.5 seconds
    }
    this.ws.onerror = (ev) => {
      console.log("[events] [onerror]", ev)
    }
    this.ws.onmessage = (ev) => {
      console.log("[events] [onmessage]", ev)

      /** @type {EventData} */
      const eventData = JSON.parse(ev.data)
      this.dispatchCustomEvent(eventData.name, eventData.data)
    }
  }

  /**
   * @param {"change"} type
   * @param {(ev: CustomEvent<ChangeEventData>) => (Promise<void>|void)} callback
   */
  addEventListener(type, callback) {
    super.addEventListener(type, callback)
  }

  /**
   * @param {"change"} type
   * @param {(ev: CustomEvent<ChangeEventData>) => (Promise<void>|void)} callback
   */
  removeEventListener(type, callback) {
    super.removeEventListener(type, callback)
  }

  /**
   * @param {"change"} type
   * @param {ChangeEventData} detail
   */
  dispatchCustomEvent(type, detail) {
    super.dispatchEvent(new CustomEvent(type, { detail }))
  }
}

export const global = new GlobalEvents()
