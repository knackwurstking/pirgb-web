# pirgb-web

Just some server for handling
[pirgb-server(s)](https://github.com/knackwurstking/pirgb-server.git)
and serving a simple ui for control.

## Getting Started

> @todo how to install all this shit (including the [pirgb-server](https://github.com/knackwurstking/pirgb-server.git) project)

## Endpoints

> @todo: "/api/v1/events"  
> @todo: what is: ":section", ":devices"

| Method | Endpoint                             | Description                                    |
| ------ | ------------------------------------ | ---------------------------------------------- |
| GET    | "/"                                  | access the ui (svelte)                         |
| GET    | "/api/v1/devices"                    | get (all) devices data                         |
| GET    | "/api/v1/devices/:host"              | get (all) device data                          |
| GET    | "/api/v1/devices/:host/:section"     | get section data from device                   |
| GET    | "/api/v1/devices/:host/:section/pwm" | get current pwm data for section               |
| POST   | "/api/v1/devices/:host/:section/pwm" | set new pwm data for section (pulse and color) |

> @todo: info about data types to send and receive (some examples?)
