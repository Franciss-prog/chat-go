package chat

import (
	"github.com/gofiber/contrib/v3/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
)

// handler for websocket connection

type Handler struct {
	repo *Repository
}

func NewHandler(db *pgxpool.Pool) *Handler {
	return &Handler{
		repo: NewRepository(db),
	}
}

// websocket connection
func (h *Handler) Connection(c *websocket.Conn) error {

	for {
		// read the message to the websocket server
		messageType, message, err := c.ReadMessage()

		// check if there is an error ocurring in the connection
		if err != nil {
			return err
		}

		// send message
		err = c.WriteMessage(messageType, message)

		// check if there is an error occurring in the writing the message
		if err != nil {
			return err
		}
	}

}
