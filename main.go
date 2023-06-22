package main

import (
	"github.com/gorilla/websocket"
	"github.com/zhandos717/gateway/app/window"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {

	window.Run()
}
