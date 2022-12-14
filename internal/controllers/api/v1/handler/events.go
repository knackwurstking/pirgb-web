package handler

import (
	"net/http"

	"github.com/knackwurstking/pirgb-web/internal/events"
	"github.com/knackwurstking/pirgb-web/pkg/log"
	"nhooyr.io/websocket"
)

func Events(w http.ResponseWriter, r *http.Request) {
	// TODO: don't use websockets any more - use something where the `r.BasicAuth` will work
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Error.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer conn.Close(websocket.StatusNormalClosure, websocket.StatusNormalClosure.String())

	events.Global.AddClient(r.Context(), conn, r.RemoteAddr)
	defer events.Global.RemoveClientAddr(r.RemoteAddr)

	for {
		msgType, msg, err := conn.Read(r.Context())
		if err != nil {
			log.Warn.Printf("Connection read error: \"%s\"", err.Error())
			return
		}

		// echo - (#heartbeat)
		go func() {
			err = conn.Write(r.Context(), msgType, msg)
			if err != nil {
				log.Warn.Printf("Send heartbeat failed: %s", err.Error())
			}
		}()
	}
}
