package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	app := fiber.New()
	// use cors
	app.Use(cors.New())

	// conect to the database
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// test routes
	app.Get("/", func(c fiber.Ctx) error {
		// test the connection
		var version string

		if err := db.QueryRow(context.Background(), "SELECT VERSION()").Scan(&version); err != nil {
			fmt.Println(err.Error())
			return c.Status(500).JSON(fiber.Map{
				"err": err.Error(),
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"version": version,
		})
	})

	// for login
	app.Post("/login", func(c fiber.Ctx) error {
		user := new(Auth)

		// parse the body use (Bind)
		if err := c.Bind().JSON(user); err != nil {
			return c.Status(400).SendString("Json data didnt come to the server")
		}

		// form validation
		if user.Email == "" || user.Password == "" {
			return c.Status(400).SendString("Bad Request")
		}
		return c.Status(200).JSON(fiber.Map{
			"message": "Succcessfully Recieve the expected Data",
		})
	})

	// for register
	app.Post("/register", func(c fiber.Ctx) error {
		user := new(Register)

		// parse the body use (Bind)
		if err := c.Bind().JSON(user); err != nil {
			return c.Status(400).SendString("Json data didnt come to the server")
		}

		// form validation
		if user.Username == "" || user.Email == "" || user.Password == "" {
			return c.Status(400).SendString("Fill all the fields")
		}
		return c.Status(200).JSON(fiber.Map{
			"message": "Succcessfully Recieve the Data",
		})
	})
	log.Fatal(app.Listen(":8080"))
}
