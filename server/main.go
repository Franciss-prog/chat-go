package main

import (
	"api/ws/database"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()
	// use cors
	app.Use(cors.New())

	// conect to the database
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// test route for database connection
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
			"message": "The database is running on version",
		})
	})

	// for login
	app.Post("/login", func(c fiber.Ctx) error {
		user := new(Auth)

		// parse the body use (Bind)
		if err := c.Bind().JSON(user); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Server didnt recieve the parsed data",
			})
		}

		// form validation
		if user.Email == "" || user.Password == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Email and Password is required",
			})
		}

		// check if the users password is correct based on the email
		// store the name of user selected
		var name string
		err := db.QueryRow(context.Background(), "SELECT username FROM users WHERE email = $1 AND password_hash = crypt($2, password_hash)", user.Email, user.Password).Scan(&name)

		// check the error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Invalid Email or Password",
			})
		}

		// return the data
		return c.Status(200).JSON(fiber.Map{
			"message": "Welcome to Chat App," + " " + name,
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

		// insert to the database with hashed password
		_, err := db.Exec(context.Background(), "INSERT INTO users(username, email, password_hash) VALUES($1, $2, crypt($3, gen_salt('bf')))", user.Username, user.Email, user.Password)

		// check the error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"err": err.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Successfully Registered!" + " " + user.Username,
		})
	})
	log.Fatal(app.Listen(":8080"))
}
