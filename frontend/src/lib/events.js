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

    /**
     * This will be `null` every time the websocket receives a close event
     * @type {WebSocket|null}
     */
    this.ws = null

    /**
     * The heartbeat timeout is currently running
     * @type {number|NodeJS.Timeout}
     */
    this._heartbeatTimeout = 0

    /** Time to wait before sending the next heartbeat to the server */
    this.heartbeatTimeoutValue = 2500


    /** Heartbeat state value: not started */
    this.NONE = 0
    /** Heartbeat state value: heartbeat send, wait for echo */
    this.SEND = 1
    /** Heartbeat state value: echo received, server reachable */
    this.RECEIVED = 2
    /** Heartbeat state value: server is gone? */
    this.FAILED = 3

    /**
     * The current heartbeat state
     * 0: none
     * 1: send
     * 2: received
     * 3; fail
     */
    this.heartbeatState = 0 

    /** Auto reconnect if the websocket receives a close event */
    this.autoReconnect = true

    /**
     * The running `setInterval`
     * @type {NodeJS.Timer|null}
     */
    this._autoReconnectInterval = null

    this.connect()
  }

  /** Connect to server: "/api/events" */
  connect() {
    if (this.ws) {
      this.ws.close()
    }

    if (this._heartbeatTimeout) clearTimeout(this._heartbeatTimeout)
    this.heartbeatState = 0 

    this.ws = new WebSocket(
      `${location.protocol === "https:" ? "wss" : "ws"}://${window.location.host}/api/events`
    )

    this.ws.onopen = (ev) => {
      console.log("[events] [onopen]", ev)
      if (this._autoReconnectInterval) {
        clearInterval(this._autoReconnectInterval)
        this._autoReconnectInterval = null
      }
      this.dispatchCustomEvent("open", null)
      this.heartbeat()
    }

    this.ws.onclose = (ev) => {
      console.log("[events] [onclose]", ev)
      clearTimeout(this._heartbeatTimeout)

      this.ws.close()
      this.ws = null

      this.dispatchCustomEvent("close", null)

      if (this.autoReconnect) {
        if (!this._autoReconnectInterval) {
          this._autoReconnectInterval = setInterval(() => {
            this.connect()
          }, 2500)
        }
      }
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

  /** Heartbeat for checking if the server is still reachable */
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
