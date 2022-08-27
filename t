[33m64b3875[m[33m ([m[1;36mHEAD -> [m[1;32mmain[m[33m, [m[1;33mtag: v0.1.0[m[33m, [m[1;31mlab/main[m[33m, [m[1;31mhome/main[m[33m)[m update...
[33m240125e[m update index
[33mbab23f5[m Change box shadow on hover ".menu-button"
[33m8f45637[m rgbw field now color
[33mfd55d78[m id field removed from online/offline event data
[33m7e1378e[m ID field removed from DeviceEventData
[33m3e73013[m update...
[33m46bee21[m pwm field RGBW renamed to color
[33m35c8fc7[m Update Endpoints Notes...
[33m1350911[m update box-shadow (for nav menu buttons) on :active
[33mab1ead7[m Moving a few structs from pirgb to servertypes (internal) and from config to pirgb
[33mf024f20[m Update ...
[33m9a51c47[m Update ...
[33m8d2128f[m Update readme (some endpoints still missing)
[33mbae00e9[m add todo (endpoints)
[33mb46ace2[m reomve todos (wrong repo)
[33m6b8d90d[m install script (makefile)
[33md3782ef[m rename config.dev.yaml to config.yaml
[33m66cdf74[m Update and notes
[33mdf1211e[m Road to version 0.1.0 - some short instructions to get started
[33me3d579a[m handling config path when running as root
[33m1670e14[m Add chi middleware `Recoverer`
[33m0bb63c8[m Moving /event and /devices handler into functions
[33m73a9728[m Code clean up; Events interface ("change","offline","online")
[33mb0c3bab[m remove todo
[33m2eb331b[m handle "offline" and "online" events
[33mdf71054[m Adding todo
[33mf8a997d[m loop forever until reconnect returns error == nil
[33mf0f1ee6[m Oops - wrong defer statement after connection read error
[33mfd9793c[m Do this reconnect thing on connection read error
[33ma53fbf3[m update logs and remove wait group done call (negative counter error)
[33md1deffc[m move servertypes package to pkg/pirgb
[33m4d2f527[m PulseSlider styles; update logs
[33m2bda487[m add vite plugin "remove console logs"
[33me2b3467[m dispatch change event (PulseSlider) and setPWM
[33m5685aec[m Remove debugging border
[33mb1234d8[m forward props to PulseSlider
[33me079d77[m using `em` from padding; replace PulseInput with PulseSlider
[33me4c3218[m seletton
[33m6bff1e8[m Add my next todo
[33m1a14330[m outsource pulse input
[33m5c2f519[m Misplaced dispatch ("change")
[33mabb79dd[m dispatch "change" event (color); send pwm
[33m6f5e6dd[m remove unused css selector
[33m458c915[m undo this white color stuff
[33m5eb8972[m await ...
[33m40a7ae3[m Remove white row and add to column (last element)
[33m533433d[m disable overflow (x)
[33m8a86c4f[m Complete range of colors (ColorPicker)
[33m24aae9d[m Change colors (still a few missing)
[33m917ff79[m remove useless stuff (css)
[33m1ca85af[m Update border colors (add opacity)
[33m8831490[m border style solid (remove input border style)
[33m111c4de[m border width
[33m8abace9[m fix the border color (var) and update width
[33m7c90ead[m ColorPicker upate...
[33mb44b02e[m Add new ColorPicker component (@TODO: replace with old colorpicker)
[33m7d54c77[m update class .online
[33m122e1c8[m Update color schemes (elevate*)
[33m79ffd7f[m Adding some color schemes
[33m96a722c[m Set powerChecked in refresh (function)
[33m2b0dc16[m async callback (PowerSwitch on:toggled)
[33m01b49b2[m Remove comment
[33mf17d2c1[m Kick the Theme picker, light not working with the PowerSwitch
[33mda6cf70[m remove comment
[33m47cc002[m Remove logging
[33ma9858a0[m Using the PowerSwitch component
[33mc977a04[m Handle WebSocket close events (auto reconnect) ...
[33m04dfa13[m Adding the power switch component for SectionCard
[33me477d07[m add todo; remove tab highlight
[33m90c1be1[m Rename some vars
[33mf57744c[m Update theme (special color; shadow)
[33mf7e3236[m Remove some todos
[33m99cb26a[m Remove old app.css (replaced with the new one)
[33m43a5547[m New color schemes (themes)
[33m605ff23[m Add color scheme picker; update color schemes; add blue-dark
[33m95b7c71[m remove some logs for now
[33m3fe8576[m Move log
[33m2e55809[m Update border style (fieldset)
[33m79a1168[m Heartbeat for checking if still connected
[33m63d5e4c[m Echo all text messages received (#heartbeat)
[33maa5de4b[m Add a offline indicator
[33md2d1477[m Add a debug log
[33med4bb55[m set duration to 800
[33mea4c993[m Remove unused import
[33m65aa7d3[m Add todo
[33m7b09bf9[m popover panel styling and adding delay for fly in
[33mcd2a911[m Remove todo
[33mded47ae[m Let the main pages fly in and out
[33mf2dc659[m fieldset > legend styling
[33mdecabd9[m Update section card styling (and positioning)
[33m117a8cc[m Fix PWM typedef; check for status 200 in getDevices
[33m4358b27[m update data types and add get / post pwm to api endpoint
[33m1beee96[m Update api url(s)
[33m78def94[m Change address to store to <host>:<port>
[33m36ffecd[m Set LastPulse for section on "change" event
[33m2cfc449[m refresh section card on "change" event
[33m02cce36[m Add todo
[33m70d6e39[m change event data type changed and update event listener callbacks
[33mf154e0d[m change event data type changed
[33m7a2a973[m dispatch custom event onmessage
[33me2f2428[m Adding todo
[33m523ddf2[m JSON parse event (message) data
[33mc386245[m Set LastPulse json tag
[33mf50baad[m Update logs
[33mce9601b[m Add todo
[33m728013a[m Some basic methods for add*, remove* and dispatch*
[33m90bcbf9[m Add "change" event listener
[33mdacf257[m Update logs, WaitGroup type, forgot to Start the event handler
[33ma03745c[m Update logs
[33m268a834[m Update log
[33m08d07d6[m Update logs
[33mcf82f44[m Update logs
[33m6315ee7[m Update logs
[33me9b916e[m [events] Remove connection after a client disconnect
[33mc2a8d1e[m Update cors
[33mcf7472c[m backend event handler class (connect) and some fixes
[33mbfd41ed[m this will handle all the events (ex: pirgb-server "change event")
[33mc8010e3[m @TODO SectionCard will replace the old Card component
[33mac531a0[m changed data types for device sections
[33m6e03108[m Just "/api/events" for receiving all types of events
[33mbeb278c[m change data type to write
[33m5ecf076[m Add/Remove client to/from register
[33m9cba842[m Register client for receiving events
[33m3a9a9da[m Adding my next todos, some client register for events.go
[33m1f79430[m Remove unused stuff here
[33mc793825[m Update typedefs
[33mdcd7900[m Add route GET "/api/devices"
[33m70901d8[m Update example configuration
[33mf545134[m Start event handler for each section
[33m11a5f3d[m config device sections changed, "change" event handler
[33m8a603c6[m Initialize global
[33m4fbf175[m more stuff on event handler
[33m721003a[m working on stuff like data storage, event handling
[33ma09cb79[m Remove todo(s)
[33md914c4a[m Remove notes (todos)
[33m111b21d[m Simplify the configuraton stuff
[33mb066f77[m Rewrite the router (/api) stuff, middleware for parsing device url param
[33mc80c51e[m Fixing pirgb import
[33mab75dda[m Update some (console) logs
[33m5cbf960[m Section data (field; section) renamed to pins
[33mc4c6830[m Remove some todos
[33m044af78[m only set pulse if bigger than 0
[33m61c2354[m load and set pulse from section data
[33mcf07c04[m remove unused comment
[33m6736b2e[m remove onMount
[33mf0976e6[m fix type (@param)
[33m08667bb[m Ignore this stupid linter error
[33m1633dd8[m add todo
[33mdcf3e77[m Return data (getPWM)
[33me68b7d8[m getPWM function
[33md78df09[m Backend data changed (Device, Sections)
[33mc16dfc6[m Update device, using pirgb.Device now
[33me9c6371[m Devcie removed, using pirgb.Device version ...
[33ma86679b[m Remove lang on script and css tag(s)
[33m4510a99[m Add one more TODO
[33m59f1926[m Update todos
[33m907f28f[m Disable this stupid timestamp
[33m3fc5ec8[m Auto select input (pulse) on click
[33mc4e2209[m Adding a todo and remove log
[33m8245095[m Change the color picker (input) padding
[33mc896602[m I think i have some problem with my svelte plugin
[33m7bb4cca[m Send color on button click (ON/OFF)
[33m7698317[m Write header (status code) after copy headers
[33m090650f[m Move handler to separate functions
[33mb2ae3eb[m Update routes info
[33m33a446d[m Update routes for /devices ...
[33m5ba59f0[m Clean up a bit
[33md3d557f[m Fix html color picker white values
[33m301453e[m calc hex to rgbw and leave a comment (todo)
[33mf844686[m Add color picker
[33m3b6be12[m Update color schemes dark and light
[33m308a5f2[m Some flex page wrap thing
[33mf8c78fc[m Update href links (nav) and leave a comment (todo)
[33m95505a4[m Some styling on the input field
[33me10641c[m Outsource fieldset stuff to card component
[33mc7920a5[m (fieldset) card like component
[33mea2f7ae[m code formatting
[33m641a1b8[m Send rgbw (color: white) on button click ON
[33m9dea081[m Enable ForceColors (logrus)
[33mc265351[m Add todo
[33mafb318f[m Update all
[33m7ec0f23[m Add: make build_frontend
[33mc49d2fe[m systemd file
[33m1d45fed[m Adding some simple Makefile
[33m17d862a[m Add current version
[33m52ced3a[m Adding card class to fieldset
[33mc83164c[m Add ON/OFF buttons for section cards
[33m797241c[m Add setPWM function
[33m2f9550f[m Don't copy header anymore
[33m0829f7c[m Forward server response to frontend
[33mf56b897[m Add CopyHelper helper function
[33m48e23b1[m Try out this http redirect thing ...
[33md2b5105[m Update router and flags ...
[33ma22b108[m Add notes/todos
[33md4fa9f3[m Update router endpoint info (print) stuff
[33m65d9eb7[m Render sections - using Transition component (headlessui)
[33mfa76487[m Make the popover fixed, and the rest scrollable
[33me79b073[m Add some padding to fieldset
[33m73fb240[m Update default config
[33m443ca5e[m Enable sorting (logrus fields)
[33m6763873[m Using a fieldset as a section card, display sections
[33meda0968[m Changed (global) border radius
[33m7b4eab8[m Parse device data and return all sections
[33m78830b9[m Paper template?
[33mf45b0b4[m Add vars for scenes and groups too
[33mcf47385[m sections var for main tag
[33md25cfa3[m Update color schemes (max/low background color switch)
[33m8d7e277[m Update schemes ...
[33m6fc0b2b[m Update comments, todos and types
[33m9aeeb0a[m Add api.js to lib/ (get sections helper)
[33m153f917[m Add @smui/paper components
[33me7c08f4[m Adding a logger middleware to chi
[33ma53a2fb[m Adding server proxy for "/api"
[33m0ec6714[m Update font-face for Sora
[33mbac2a9a[m Update embed line
[33mf5017e1[m Router "/api" endpoint "/sections" renamed to "/devices"
[33m33fe61c[m Rename GetGroups and add GetDevices
[33m8b3c0d6[m Remove todo
[33m74091a8[m Update api "/groups" return
[33mf714ae8[m Add `GetGroups` method
[33m3c34ebf[m Print out some info about (available) devices
[33m50e1386[m Changed groups handling, and remove groups from config
[33mc842ec9[m Devices and Groups type update
[33md883335[m Ignore: todo.txt
[33m89c6c7a[m Moving "Sora" fonts to assets/
[33m723b0e2[m Popover positioning
[33m675e27d[m Get the embedded file system running for the UI
[33mda27f34[m - enable http server (default configuration) - remove Scan field from Config - update some logging messages
[33mf59f25d[m Add README
[33me10cb33[m No need for a Makefile
[33m81751d4[m Update loggings
[33m070aedc[m FIX the embed error on dist folder (frontend)
[33m12f931f[m Update config ...
[33m877f054[m Scan sections if len == 0
[33m4dea5c5[m Update todo
[33m25202a8[m Just do a simple http request for the "info" endpoint to get  sections (if device exists)
[33m8ac7c5f[m rmove todo
[33m3e64b05[m Add default (servre) port and leave a note (FIX) on frontend.go
[33m1bb244d[m Update todos
[33mac54494[m Change the config entries for Sections and Groups (adding some todos)
[33mddc7043[m change server port
[33maf2c7b0[m Add some padding to (router endpoints) info keys
[33mb82e6bf[m Handling cors for api endpoint(s)
[33m177e294[m handle api: GET sections and groups
[33m513837e[m Documenting code (doing some go linter things)
[33m86afa17[m Remove GetMux function (just use global Mux router variable)
[33mc9c092d[m Print endpoints function for router package
[33mf675b2a[m a one line list of sections in groups
[33m09fb872[m handle frontend dist directory
[33mb4df500[m add todos
[33mabbbf30[m work on initializing the router and config before server start
[33m8d71aac[m return the global Mux router
[33m741eb00[m adding todos
[33m2e5f74f[m server + router - basic sturr
[33m5ccd8f9[m Update tags
[33m30995f5[m Ignore executable
[33m8ecb776[m clean up
[33m5358b6b[m +todo +note
[33me1689a5[m Configuration handling (yaml)
[33m6bf3bc9[m +server.go
[33m0500497[m update popover to fit my screen (on my phone)
[33m0f219f6[m Adding some notes after TODO
[33m733c8da[m - adding font "Sora" - update scheme variable name - add special color variables
[33m91add3b[m update popover z-index(es)
[33m7667cde[m adding todo
[33m75ffec6[m rename light and dark themes to mono-{dark/light}
[33m9536d24[m formatting
[33m16ec82c[m adding color transition
[33m7e968fd[m update colorschemes (monochromatic)
[33me098572[m do some css for popover (menu) items
[33mf6dc016[m some color scheme stuff
[33m5e1a032[m Remove Counter component
[33m29ddb42[m Remove svelte.svg
[33mb0f3d9a[m fiddle with the "app" layout?
[33m05ad94e[m disclosure: sections, groups and scripts (testing, styling, basics)
[33m154c507[m create vite (svelte) project
[33m20727d4[m adding packages i thing i need for this
[33m1b8b426[m basic projet layout
