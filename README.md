# pirgb-web (v0.1.0) - Out of Date

Client side web server, controls the [pirgb-server](http://gitlab.com/knackwurstking/pirgb-server.git).

## Getting started

> run without `sudo`

```bash
git clone https://gitlab.com/knackwurstking/pirgb-web
cd pirgb-web
make
make install
```

**NOTE:** _The [pirgb-server](https://gitlab.com/knackwurstking/pirgb-server.git) needs to be running before you start this server._

> replace **localhost** with hostname (or ip) where the server is running

Open a browser and go to _http://**localhost**:50831/_

## Endpoints (Notes)

### Index

- [GET: /](#endpoint-ui)
- [GET: /api/devices](#endpoint-api-devices)
- [GET: /api/devices/{host}/{section:\[0-9\]}/pwm](#endpoint-api-get-pwm)
- [POST: /api/devices/{host}/{section:\[0-9\]}/pwm](#endpoint-api-post-pwm)
- [WS: /api/events](#endpoint-api-events)

---

<a id="endpoint-ui"></a>

> **GET** _/_

Serve the UI

---

<a id="endpoint-api-devices"></a>

> **GET** _/api/devices_

Get available devices

Example Request

```sh
curl http://localhost:50831/api/devices | jq
```

Example Response

```json
[
  {
    "host": "pi-bed",
    "port": 50826,
    "sections": [
      { "id": 0, "pulse": 0, "lastPulse": 66, "color": [255, 0, 0, 0] }
    ],
    "groups": ["all"]
  },
  {
    "host": "pi-lr",
    "port": 50826,
    "sections": [
      { "id": 0, "pulse": 0, "lastPulse": 0, "color": [255, 255, 255, 255] },
      { "id": 1, "pulse": 0, "lastPulse": 0, "color": [255, 255, 255, 255] }
    ],
    "groups": ["all", "lr"]
  }
]
```

---

<a id="endpoint-api-get-pwm"></a>

> **GET** _/api/devices/{host}/{section:[0-9]}/pwm_

Example Request

```sh
curl http://localhost:50831/api/devices/pi-lr/1/pwm | jq
```

Example Response

```json
{ "id": 1, "pulse": 0, "lastPulse": 0, "color": [255, 255, 255, 255] }
```

---

<a id="endpoint-api-post-pwm"></a>

> **POST** _/api/devices/{host}/{section:[0-9]}/pwm_

Example Request

```sh
curl http://localhost:50831/api/devices/pi-lr/1/pwm \
  -X POST \
  --header "Content-Type: application/json" \
  --data '{"pulse": 100, "color": [255, 255, 255, 255]}'
```

Example Request Body

* rgbw: optional


```json
{
  "pulse": 100,
  "color": [255, 255, 255, 255]
}
```

---

<a id="endpoint-api-events"></a>

> **WS** _/api/events_

Emits events like "change", "online", "offline"

- _"change"_: Device section (id) color or pulse changed

```json
{
  "name": "change",
  "data": {
    "host": "pi-bed",
    "Port": 50826,
    "id": 0,
    "pulse": 100,
    "lastPulse": 0,
    "color": [255, 255, 255, 255]
  }
}
```

- _"offline"_: Device (section) is offline

```json
{
  "name": "offline",
  "data": { "host": "pi-bed", "Port": 50826, "id": 0 }
}
```

- _"online"_: Device is online

```json
{
  "name": "offline",
  "data": { "host": "pi-bed", "Port": 50826, "id": 0 }
}
```

---
