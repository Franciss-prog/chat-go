package auth

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

// each hanlder have the struct in service

type Handler struct {
	db *pgxpool.Pool
}

// function for new Handler

func NewHandler(db *pgxpool.Pool) *Handler {
	return &Handler{
		db: db,
	}
}

// for testing
func (h *Handler) Health(c fiber.Ctx) error {
	// declare string to store the version
	var version string
	// query the database
	if err := h.db.QueryRow(context.Background(), "SELECT VERSION()").Scan(&version); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "DB is running",
		"version": version,
	})
}
