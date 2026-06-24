package routes

import (
	"api/ws/auth"
	"api/ws/chat"
	"api/ws/middleware"
	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(app *fiber.App, db *pgxpool.Pool) {
	authHandler := auth.NewHandler(db)
	chatHandler := chat.NewHandler(db)

	app.Get("/", authHandler.Health)
	app.Post("/login", authHandler.Login)
	app.Post("/register", authHandler.Register)

	app.Get("/ws", middleware.AuthMiddleware, websocket.New(chatHandler.Connection))
}
