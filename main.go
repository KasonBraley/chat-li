package main

import (
	"flag"

	"github.com/KasonBraley/chat-li/socket"
)

func main() {
	flag.Parse()
	hub := socket.NewHub()
	go hub.Run()

	r := SetupRouter(hub)
	r.Run(":5000")
}
