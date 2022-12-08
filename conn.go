package main

import (
	"log"
	"net/http"
)

// serveWs handles websocket requests from the peer.
func serveWs(w http.ResponseWriter, r *http.Request, roomId string) {

  	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer ws.Close()
	log.Println("Connected!")
	
	c := &connection{send: make(chan []byte, 256), ws: ws}
	s := subscription{c, roomId}
	h.register <- s
	go s.writeStream()
	go s.readStream()
}