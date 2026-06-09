package main

import (
	"api/ws/database"
	"api/ws/routes"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	app := fiber.New()
	// use cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true,
	}))

	// database connection
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// init the routes and pass the databse pool
	routes.SetupRoutes(app, db)

	// start the server with 8080 port
	log.Fatal(app.Listen(":8080"))
}
