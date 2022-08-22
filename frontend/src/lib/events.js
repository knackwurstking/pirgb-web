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
    this._timeout = 0

    this._heartbeatTimeoutValue = 2500


    this._NONE = 0
    this._SEND = 1
    this._RECEIVED = 2
    this._FAILED = 3

    /**
     * 0: none
     * 1: send
     * 2: received
     * 3; fail
     */
    this._heartbeat = 0 

    this.connect()
  }

  // Connect to "/api/events"
  connect() {
    if (this.ws) {
      this.ws.close()
      if (this._timeout) clearTimeout(this._timeout)
      this._heartbeat = 0 
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
      clearTimeout(this._timeout)
      this.dispatchCustomEvent("close", null)
    }

    this.ws.onmessage = (ev) => {
      if (ev.data == "heartbeat") {
        if (this._heartbeat === this._FAILED) {
          this.dispatchCustomEvent("open", null)
        }

        this._heartbeat = this._RECEIVED
        return
      }

      //console.log("[events] [onmessage]", ev)

      /** @type {ServerEventData} */
      const eventData = JSON.parse(ev.data)
      this.dispatchCustomEvent(eventData.name, eventData.data)
    }
  }

  heartbeat() {
    if (this._timeout) {
      clearTimeout(this._timeout)
      this._timeout = 0
    }

    if (!this.ws) return
    if (this.ws.readyState !== this.ws.OPEN) return

    if (this._heartbeat === this._SEND) {
      this._heartbeat = this._FAILED
      this.closed = true
      this.dispatchCustomEvent("close", null)
    }

    //console.log("[events] send a heartbeat ...")
    if (this._heartbeat !== this._FAILED) this._heartbeat = this._SEND
    this.ws.send("heartbeat")
    this._timeout = setTimeout(this.heartbeat.bind(this), this._heartbeatTimeoutValue);
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
