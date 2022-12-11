export class GlobalEvents extends EventTarget {
  constructor() {
    console.log("[events.js] initializing...");
    super();

    /**
     * This will be `null` every time the websocket receives a close event
     * @type {WebSocket|null}
     */
    this.ws = null;

    /**
     * The heartbeat timeout is currently running
     * @type {number|NodeJS.Timeout}
     */
    this._heartbeatTimeout = 0;

    /** Time to wait before sending the next heartbeat to the server */
    this.heartbeatTimeoutValue = 2500;

    /** Heartbeat state value: not started */
    this.NONE = 0;
    /** Heartbeat state value: heartbeat send, wait for echo */
    this.SEND = 1;
    /** Heartbeat state value: echo received, server reachable */
    this.RECEIVED = 2;
    /** Heartbeat state value: server is gone? */
    this.FAILED = 3;

    /**
     * The current heartbeat state
     * 0: none
     * 1: send
     * 2: received
     * 3; fail
     */
    this.heartbeatState = 0;

    /** Auto reconnect if the websocket receives a close event */
    this.autoReconnect = true;

    /**
     * The running `setInterval`
     * @type {NodeJS.Timer|null}
     */
    this._autoReconnectInterval = null;

    this.connect();
  }

  /** Connect to server: "/api/events" */
  connect() {
    if (this.ws) this.ws.close();
    if (this._heartbeatTimeout) clearTimeout(this._heartbeatTimeout);
    this.heartbeatState = 0;

    this.ws = new WebSocket(`wss://${location.host}/api/v1/events`);

    this.ws.onopen = () => {
      console.log("[events] onopen");
      if (this._autoReconnectInterval) {
        clearInterval(this._autoReconnectInterval);
        this._autoReconnectInterval = null;
      }
      this.dispatchCustomEvent("open", null);
      this.heartbeat();
    };

    this.ws.onclose = () => {
      console.log("[events] onclose");
      clearTimeout(this._heartbeatTimeout);

      if (this.ws) this.ws.close();
      this.ws = null;

      this.dispatchCustomEvent("close", null);

      if (this.autoReconnect) {
        if (!this._autoReconnectInterval) {
          this._autoReconnectInterval = setInterval(() => {
            this.connect();
          }, 2500);
        }
      }
    };

    /** @param {MessageEvent<string>} ev */
    this.ws.onmessage = (ev) => {
      if (ev.data == "heartbeat") {
        if (this.heartbeatState === this.FAILED) {
          this.dispatchCustomEvent("open", null);
        }

        this.heartbeatState = this.RECEIVED;
        return;
      }

      /** @type {BaseEvent<BaseEventTypes, DeviceEvent | ChangeEvent>} */
      const eventData = JSON.parse(ev.data);
      this.dispatchCustomEvent(eventData.name, eventData.data);
    };
  }

  /** Heartbeat for checking if the server is still reachable */
  heartbeat() {
    if (this._heartbeatTimeout) {
      clearTimeout(this._heartbeatTimeout);
      this._heartbeatTimeout = 0;
    }

    if (!this.ws) return;
    if (this.ws.readyState !== this.ws.OPEN) return;

    if (this.heartbeatState === this.SEND) {
      this.heartbeatState = this.FAILED;
      this.closed = true;
      this.dispatchCustomEvent("close", null);
    }

    if (this.heartbeatState !== this.FAILED) this.heartbeatState = this.SEND;
    this.ws.send("heartbeat");
    this._heartbeatTimeout = setTimeout(
      this.heartbeat.bind(this),
      this.heartbeatTimeoutValue
    );
  }

  /**
   * @param {BaseEventTypes | "open" | "close"} type
   * @param {(ev: CustomEvent<DeviceEvent | ChangeEvent | null>) => (Promise<void>|void)} callback
   */
  addEventListener(type, callback) {
    super.addEventListener(type, callback);
  }

  /**
   * @param {BaseEventTypes | "open" | "close"} type
   * @param {(ev: CustomEvent<DeviceEvent | ChangeEvent | null>) => (Promise<void>|void)} callback
   */
  removeEventListener(type, callback) {
    super.removeEventListener(type, callback);
  }

  /**
   * @param {BaseEventTypes | "open" | "close"} type
   * @param {DeviceEvent | ChangeEvent | null} detail
   */
  dispatchCustomEvent(type, detail) {
    super.dispatchEvent(new CustomEvent(type, { detail }));
  }
}

export default new GlobalEvents();
