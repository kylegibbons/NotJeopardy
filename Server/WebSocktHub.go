package main

import (
	"github.com/satori/go.uuid"
)

// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan Message

	// Register requests from the clients.
	register chan *Client

	joinGame chan JoinInfo

	// Unregister requests from clients.
	unregister chan *Client
}

type JoinInfo struct {
	GameID string
	Client *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		joinGame:   make(chan JoinInfo),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true

			uuid, _ := uuid.NewV4()
			client.clientID = uuid.String()

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				if message.GameID == client.gameID {
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		case info := <-h.joinGame:
			for client := range h.clients {
				if client.clientID == info.Client.clientID {
					client.gameID = info.GameID
				}
			}
		}
	}
}
