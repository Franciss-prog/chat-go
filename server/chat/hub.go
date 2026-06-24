package chat

import (
	"encoding/json"
	"log"
)

type Hub struct {
	Rooms      map[string]map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 256),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.JoinRoom(client, "general")

		case client := <-h.Unregister:
			h.LeaveRoom(client)

		case msg := <-h.Broadcast:
			h.broadcastToRoom(msg)
		}
	}
}

func (h *Hub) JoinRoom(client *Client, room string) {
	h.LeaveRoom(client)
	client.Room = room
	if h.Rooms[room] == nil {
		h.Rooms[room] = make(map[*Client]bool)
	}
	h.Rooms[room][client] = true

	resp := WSResponse{
		Type:     "join",
		Room:     room,
		SenderID: client.UserID,
	}
	data, _ := json.Marshal(resp)
	for c := range h.Rooms[room] {
		select {
		case c.Send <- data:
		default:
			delete(h.Rooms[room], c)
			close(c.Send)
		}
	}

	log.Printf("client %s joined room %s", client.UserID, room)
}

func (h *Hub) LeaveRoom(client *Client) {
	if client.Room == "" {
		return
	}
	room := client.Room
	if clients, ok := h.Rooms[room]; ok {
		delete(clients, client)
		if len(clients) == 0 {
			delete(h.Rooms, room)
		}
	}

	resp := WSResponse{
		Type:     "leave",
		Room:     room,
		SenderID: client.UserID,
	}
	data, _ := json.Marshal(resp)
	for c := range h.Rooms[room] {
		select {
		case c.Send <- data:
		default:
			delete(h.Rooms[room], c)
			close(c.Send)
		}
	}

	client.Room = ""
	log.Printf("client %s left room %s", client.UserID, room)
}

func (h *Hub) broadcastToRoom(msg *Message) {
	clients, ok := h.Rooms[msg.Room]
	if !ok {
		return
	}

	resp := WSResponse{
		Type:      "message",
		Room:      msg.Room,
		SenderID:  msg.SenderID,
		Username:  msg.Username,
		Content:   msg.Content,
		Timestamp: msg.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}
	data, _ := json.Marshal(resp)

	for client := range clients {
		select {
		case client.Send <- data:
		default:
			delete(clients, client)
			close(client.Send)
		}
	}
}
