package routes

import (
	"api/ws/auth"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(app *fiber.App, db *pgxpool.Pool) {

	// get all the handler cuz thats all we need
	authHandler := auth.NewHandler(db)

	// auth
	app.Get("/", authHandler.Health)
	app.Post("/login", authHandler.Login)
	app.Post("/register", authHandler.Register)
}
