package sockets

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocketConnection(conn *websocket.Conn, hub *Hub) {
	defer conn.Close()
	wsClient := newClient(hub, conn)
	wsClient.serve()
}

func SocketsHTTP(w http.ResponseWriter, r *http.Request, hub *Hub) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade WebSocket connection: %v", err)
		return
	}
	// Handle the WebSocket connection
	go handleWebSocketConnection(conn, hub)
}
