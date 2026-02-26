package main

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// We must attach the function directly to the struct field!
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("Error: ", slog.Any("Messgae", err))
		return
	}
	defer conn.Close()

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			slog.Error("Error: ", slog.Any("Messgae", err))
			break
		}
		slog.Info("Received: ", slog.String("Messgae", string(msg)))

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
