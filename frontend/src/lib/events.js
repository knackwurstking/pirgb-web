// TODO: global event handling

/**
 * @typedef {{
 *  name: "change",
 *  data: ChangeEventData,
 * }} ServerEventData
 *
 * @typedef {{
 *  host: string,
 *  port: number,
 *  id: number,
 *  pulse: number,
 *  lastPulse: number,
 *  color: number[],
 * }} ChangeEventData
 *
 * @typedef {null} OfflineEventData
 *
 * @typedef {"change"|"close"|"open"} GlobalEventTypes
 * @typedef {ChangeEventData|OfflineEventData} GlobalEventData
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
    // TODO: websocket ping pong for detect a disconnect
    if (this.ws) {
      this.ws.close()
    }

    this.ws = new WebSocket(`${location.protocol === "https:" ? "wss" : "ws"}://${window.location.host}/api/events`)
    window.ws = this.ws

    this.ws.onopen = (ev) => {
      console.log("[events] [onopen]", ev)
      this.dispatchCustomEvent("open", null)
    }

    this.ws.onclose = (ev) => {
      console.log("[events] [onclose]", ev)
      this.dispatchCustomEvent("close", null)
    }

    this.ws.onerror = (ev) => {
      console.log("[events] [onerror]", ev)
    }

    this.ws.onmessage = (ev) => {
      console.log("[events] [onmessage]", ev)

      /** @type {ServerEventData} */
      const eventData = JSON.parse(ev.data)
      this.dispatchCustomEvent(eventData.name, eventData.data)
    }
  }

  /**
   * @param {GlobalEventTypes} type
   * @param {(ev: CustomEvent<GlobalEventData>) => (Promise<void>|void)} callback
   */
  addEventListener(type, callback) {
    super.addEventListener(type, callback)
  }

  /**
   * @param {GlobalEventTypes} type
   * @param {(ev: CustomEvent<GlobalEventData>) => (Promise<void>|void)} callback
   */
  removeEventListener(type, callback) {
    super.removeEventListener(type, callback)
  }

  /**
   * @param {GlobalEventTypes} type
   * @param {GlobalEventData} detail
   */
  dispatchCustomEvent(type, detail) {
    super.dispatchEvent(new CustomEvent(type, { detail }))
  }
}

export const global = new GlobalEvents()
