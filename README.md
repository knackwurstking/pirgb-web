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

```go
type PWM struct {
	Pulse int   `json:"pulse"`
	RGBW  []int `json:"rgbw"`
}

type Section struct {
	ID        int   `json:"id" yaml:"id"`
	Pulse     int   `json:"pulse" yaml:"pulse"`
	LastPulse int   `json:"lastPulse"`
	Color     []int `json:"color" yaml:"color"`
}

type Device struct {
	Host     string     `json:"host" yaml:"host"`
	Port     int        `json:"port" yaml:"port"`
	Sections []*Section `json:"sections" yaml:"sections"`
	Groups   []string   `json:"groups" yaml:"groups"`
}
```

| Method | Url                                 | Request | Response     | Description                                    |
| ------ | ----------------------------------- | ------- | ------------ | ---------------------------------------------- |
| GET    | /                                   | -       | -            | serve ui                                       |
| GET    | /devices                            | -       | `[]*Devices` | list available devices                         |
| GET    | /devices/{host}/{section:[0-9]}/pwm | -       | `Section`    | get pwm section data from device               |
| POST   | /devices/{host}/{section:[0-9]}/pwm | `PWM`   | -            | update pulse and/or color for a device section |

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

**Returns**: `BaseEventData[ChangeEventData]`

##### _Event Name: `"online"`_

Fired when a device ([pirgb-server](https://gitlab.com/knackwurstking/pirgb-server.git)) is online.

**Returns**: `BaseEventData[DeviceEventData]`

##### _Event Name: `"offline"`_

Fired when a device ([pirgb-server](https://gitlab.com/knackwurstking/pirgb-server.git)) is offline.

**Returns**: `BaseEventData[DeviceEventData]` (_Just like the "**online**" event_)
