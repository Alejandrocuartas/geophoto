package sockets

import (
	//"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Message string

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	send chan []byte
}

func NewClient(hub *Hub, conn *websocket.Conn) *Client {
	newClient := &Client{
		Hub:  hub,
		Conn: conn,
		send: make(chan []byte),
	}
	hub.Register <- newClient
	return newClient
}

func (c *Client) Read() {
	defer func() {
		c.Hub.unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		//var msg Message
		//err = json.Unmarshal(message, &msg)
		//if err != nil {
		//		log.Println("Error decoding JSON message:", err)
		//	continue
		//}

		log.Printf("Received message from %s:", message)

		// Broadcast message to all connected clients
		c.Hub.broadcast <- message
	}
}

func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.send {
		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			return
		}
	}
}

func (c *Client) Serve() {
	go c.Write()
	c.Read()
}

type Hub struct {
	clients    map[*Client]bool
	Register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					fmt.Println("here2")
					close(client.send)
					//delete(h.clients, client)
				}
			}
		}
	}
}

////////////

func HandleWebSocketConnection(conn *websocket.Conn) {
	defer conn.Close()

	// Read messages from the WebSocket connection
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Failed to read message from WebSocket connection: %v", err)
			break
		}

		log.Printf("Received message: %s", message)

		// TODO: Handle the WebSocket message
	}
}
