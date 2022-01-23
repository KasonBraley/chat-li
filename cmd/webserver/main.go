package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/KasonBraley/chat-li/socket"
)

var addr = flag.String("addr", "localhost:5000", "http service address")

func main() {
	flag.Parse()
	flag.Parse()
	hub := socket.NewHub()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		socket.ServeWs(hub, w, r)
	})

	log.Fatal(http.ListenAndServe(*addr, nil))
}
