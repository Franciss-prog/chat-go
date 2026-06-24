package chat

import (
	"log"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	svc *Service
	hub *Hub
}

func NewHandler(db *pgxpool.Pool) *Handler {
	h := &Handler{
		svc: NewService(NewRepository(db)),
		hub: NewHub(),
	}
	go h.hub.Run()
	return h
}

func (h *Handler) Connection(c *websocket.Conn) {
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		log.Println("connection rejected: missing userID")
		return
	}
	username, _ := c.Locals("username").(string)

	client := &Client{
		Hub:      h.hub,
		Conn:     c,
		Send:     make(chan []byte, 256),
		UserID:   userID,
		Username: username,
	}

	h.hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
