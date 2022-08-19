# pirgb-web

## TODO

- [ ] create a dummy server first, deliver all data

### Frontend

- [ ] basic app template with monochrome color schemes
- [ ] nav: sections, groups (top left)
- [ ] theme picker (top right)
- [ ] flex page (auto wrap)
  - [ ] transition on page swap (nav)
    - [ ] section card component
      - [ ] title - host and section
      - [ ] content: color picker (rgb and pulse)
      - [ ] actions: pulse slider (0-100)
      - [ ] actions: power toggle (on/off)
    - [ ] group card component

### Backend

- [ ] /api
  - [ ] /devices
    - [ ] GET / - returns a list of devices (`host:string, port:number, sections:number[]`)
    - [ ] GET /{host}/{id}/pwm - get pwm data from server and parse data (`pulse:number, rgbw:number[]`)
    - [ ] POST /{host}/{id}/pwm - set pwm (`pulse:number, rgbw:number[]`)
