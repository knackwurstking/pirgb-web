# pirgb-web (v0.1.0)

Client side web server for control [pirgb-server](http://gitlab.com/knackwurstking/pirgb-server).

## Getting started

> run without `sudo`

```bash
git clone https://gitlab.com/knackwurstking/pirgb-web
cd pirgb-web
make
make install # no sudo required
make edit_config # this will open the systemd file for adding parameters
make start_service
```

> replace **localhost** with hostname (or ip) where the server is running

Open a browser and go to _http://**localhost**:50831/_
