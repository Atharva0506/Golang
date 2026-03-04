package websocket

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)

// Client represents a single WebSocket connection.
type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

// Hub maintains the set of active clients and broadcasts messages.
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

// NewHub creates a new Hub with initialized channels.
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// Run starts the Hub's event loop. Must be called in a goroutine.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			slog.Info("client connected", "total", len(h.clients))

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				slog.Info("client disconnected", "total", len(h.clients))
			}

		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

// writePump pumps messages from the send channel to the WebSocket connection.
func (c *Client) writePump() {
	defer c.conn.Close()
	for message := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			return
		}
	}
}

// readPump reads messages from the WebSocket to detect disconnects.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
	}
}

// upgrader upgrades HTTP connections to WebSocket.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// ServeWs handles a WebSocket upgrade request and registers the new client.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("websocket upgrade failed", "error", err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	hub.register <- client
	go client.readPump()
	go client.writePump()
}

// BroadcastSignal marshals a signal to JSON and sends it to all connected clients.
func (h *Hub) BroadcastSignal(signal interface{}) {
	data, err := json.Marshal(signal)
	if err != nil {
		slog.Error("failed to marshal signal", "error", err)
		return
	}
	h.broadcast <- data
}
