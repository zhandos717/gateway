// websockets.go
package main

import (
	"fmt"
	"github.com/zhandos717/gateway/app/database"
	"github.com/zhandos717/gateway/app/websocket"
	"github.com/zhandos717/gateway/app/window"
	"log"
)

func main() {

	err := database.Connect()
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных:", err)
	}

	go websocket.StartServer(messageHandler)

	window.Run()
}

func messageHandler(message []byte) {

	fmt.Println(string(message))
}
