package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

func echo(ws *websocket.Conn) {
	var msg string
	for {
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			break
		}
		fmt.Println("Received:", msg)
		websocket.Message.Send(ws, "Echo: "+msg)
	}
}

func main() {
	http.Handle("/ws", websocket.Handler(echo))
	fmt.Println("WebSocket server on :8080")
	http.ListenAndServe(":8080", nil)
}
