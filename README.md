# pirgb-web (v0.1.0)

Client side web server, controls the [pirgb-server](http://gitlab.com/knackwurstking/pirgb-server.git).

## Getting started

> run without `sudo`

```bash
git clone https://gitlab.com/knackwurstking/pirgb-web
cd pirgb-web
make
make install # no sudo required
```

> replace **localhost** with hostname (or ip) where the server is running

Open a browser and go to _http://**localhost**:50831/_

## Endpoints (Notes)

### HTTP

@TODO: endpoints table...

| Method | Url | Request | Response | Description |
| ------ | --- | ------- | -------- | ----------- |
| GET    | /   | -       | -        | serve ui    |

### Websocket

| Url                           | Description                              |
| ----------------------------- | ---------------------------------------- |
| [/api/events](#server-events) | server events: _change, online, offline_ |

<a id="server-events" />

#### /api/events

> The server will simply echo all data received from the client.  

Data types:

```go
type Events interface {
	DeviceEventData | ChangeEventData
}

type BaseEventData[T Events] struct {
	Name string `json:"name"`
	Data T      `json:"data"`
}

type DeviceEventData struct { // Used for "offline" and "online" events
	Host string `json:"host"`
	Port int    `json:"port"`
	ID   int    `json:"id" yaml:"id"`
}

type ChangeEventData struct {
	DeviceEventData
	Pulse     int   `json:"pulse" yaml:"pulse"`
	LastPulse int   `json:"lastPulse"`
	Color     []int `json:"color" yaml:"color"`
}
```

##### _Event Name: `"change"`_

**Sends data:**: `BaseEventData[ChangeEventData]`

<a id="online-server-event" />

##### _Event Name: `"online"`_

Fired when a device ([pirgb-server](https://gitlab.com/knackwurstking/pirgb-server.git)) is online.

**Sends data**: `BaseEventData[DeviceEventData]`

<a id="offline-server-event" />

##### _Event Name: `"offline"`_

Fired when a device ([pirgb-server](https://gitlab.com/knackwurstking/pirgb-server.git)) is offline.

**Sends data**: `BaseEventData[DeviceEventData]` (*Just like the "**online**" event*)
