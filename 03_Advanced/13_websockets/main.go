package main

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)

// 1. The Upgrader defines the rules for upgrading a standard HTTP connection to a WebSocket connection.
var upgrader = websocket.Upgrader{
	// In production, we check if the origin URL is allowed (CORS). For testing, we allow everyone!
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// 2. UPGRADE! The user knocks with HTTP, and we instantly respond: "Let's switch to WebSockets!"
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("Error: ", slog.Any("Messgae", err))
		return
	}
	// Don't forget to close the persistent connection when the user leaves!
	defer conn.Close()

	// 3. THE INFINITE LOOP
	// Because the connection stays open forever, we constantly listen for messages from the user!
	for {
		// BLOCKS right here until the user sends a message
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			slog.Error("Error: ", slog.Any("Messgae", err))
			break // Break the loop and disconnect the user on error
		}
		slog.Info("Received: ", slog.String("Messgae", string(msg)))

		// 4. ECHO BACK TO THE USER
		reply := []byte("Server received: " + string(msg))
		err = conn.WriteMessage(mt, reply)
		if err != nil {
			slog.Error("Error: ", slog.Any("Messgae", err))
			break
		}
	}
}
func main() {
	http.HandleFunc("/ws", wsHandler)
	slog.Info("WebSocket Server listening on :8080")
	http.ListenAndServe(":8080", nil)

}
