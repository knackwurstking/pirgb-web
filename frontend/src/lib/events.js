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

    // TODO: connect websocket to "/api/events"
  }
}

export const global = new GlobalEvents()
