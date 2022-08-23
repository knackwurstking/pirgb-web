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

    /** @type {number|NodeJS.Timeout} */
    this._heartbeatTimeout = 0

    this.heartbeatTimeoutValue = 2500


    this.NONE = 0
    this.SEND = 1
    this.RECEIVED = 2
    this.FAILED = 3

    /**
     * 0: none
     * 1: send
     * 2: received
     * 3; fail
     */
    this.heartbeatState = 0 

    this.connect()
  }

  // Connect to "/api/events"
  connect() {
    if (this.ws) {
      this.ws.close()
      if (this._heartbeatTimeout) clearTimeout(this._heartbeatTimeout)
      this.heartbeatState = 0 
    }

    this.ws = new WebSocket(
      `${location.protocol === "https:" ? "wss" : "ws"}://${window.location.host}/api/events`
    )

    this.ws.onopen = (ev) => {
      console.log("[events] [onopen]", ev)
      this.dispatchCustomEvent("open", null)
      this.heartbeat()
    }

    this.ws.onclose = (ev) => {
      console.log("[events] [onclose]", ev)
      clearTimeout(this._heartbeatTimeout)
      this.dispatchCustomEvent("close", null)
    }

    this.ws.onmessage = (ev) => {
      if (ev.data == "heartbeat") {
        if (this.heartbeatState === this.FAILED) {
          this.dispatchCustomEvent("open", null)
        }

        this.heartbeatState = this.RECEIVED
        return
      }

      //console.log("[events] [onmessage]", ev)

      /** @type {ServerEventData} */
      const eventData = JSON.parse(ev.data)
      this.dispatchCustomEvent(eventData.name, eventData.data)
    }
  }

  heartbeat() {
    if (this._heartbeatTimeout) {
      clearTimeout(this._heartbeatTimeout)
      this._heartbeatTimeout = 0
    }

    if (!this.ws) return
    if (this.ws.readyState !== this.ws.OPEN) return

    if (this.heartbeatState === this.SEND) {
      this.heartbeatState = this.FAILED
      this.closed = true
      this.dispatchCustomEvent("close", null)
    }

    //console.log("[events] send a heartbeat ...")
    if (this.heartbeatState !== this.FAILED) this.heartbeatState = this.SEND
    this.ws.send("heartbeat")
    this._heartbeatTimeout = setTimeout(this.heartbeat.bind(this), this.heartbeatTimeoutValue);
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
