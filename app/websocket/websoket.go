package websoket

import (
	"fmt"
	"net/http"
)

http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// Print the message to the console
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// Write message back to browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
})

http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "websockets.html")
})

err := http.ListenAndServe(":8080", nil)

if err != nil {
return
}