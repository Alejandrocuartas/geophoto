package sockets

import (
	"log"
	"github.com/gorilla/websocket"
)

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	send chan []byte
}

func newClient(hub *Hub, conn *websocket.Conn) *Client {
	newClient := &Client{
		Hub:  hub,
		Conn: conn,
		send: make(chan []byte),
	}
	hub.register <- newClient
	return newClient
}

func (c *Client) read() {
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
		// Broadcast message to all connected clients
		c.Hub.broadcast <- message
	}
}

func (c *Client) write() {
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

func (c *Client) serve() {
	go c.write()
	c.read()
}

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
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
					close(client.send)
				}
			}
		}
	}
}
