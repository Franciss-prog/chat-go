package auth

import (
	"api/ws/middleware"
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

// each hanlder have the struct in service

type Handler struct {
	repo *Repository
}

// function for new Handler

func NewHandler(db *pgxpool.Pool) *Handler {
	return &Handler{
		repo: NewRepository(db),
	}
}

// handler for testing
func (h *Handler) Health(c fiber.Ctx) error {

	// check the version to database
	version, err := h.repo.CheckVersion(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "DB is running",
		"version": version,
	})
}

// handler for login
func (h *Handler) Login(c fiber.Ctx) error {
	// request struct for login
	request := new(LoginRequest)

	// validate the incoming request if received
	if err := c.Bind().JSON(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request, Please try again",
		})
	}

	// form validation, just incase
	if (request.Email == "") || (request.Password == "") {
		return c.Status(400).JSON(fiber.Map{
			"message": "details are missing, Please try again",
		})
	}

	// validate the user in the database
	user, err := h.repo.GetUserByUsername(c.Context(), request.Email, request.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid Email and Password, Please try again",
		})
	}
	// claim the jwt
	jwt, err := middleware.GenerateToken(user.ID)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid Email and Password, Please try again",
		})
	}

	// send cookies to client
	c.Cookie(
		&fiber.Cookie{
			Name:     "token",
			Value:    jwt,
			HTTPOnly: true,
			Path:     "/",
			MaxAge:   60 * 60 * 24,
			Secure:   false,
		})

	return c.Status(200).JSON(fiber.Map{
		"message": "Welecome to Chat App" + " " + user.Username,
		"name":    user.Username,
	})
}

// handler for register
func (h *Handler) Register(c fiber.Ctx) error {
	// request struct for register
	request := new(RegisterRequest)

	// bind the incoming request
	if err := c.Bind().JSON(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request, Please try again",
		})
	}

	// form validation, just incase
	if (request.Username == "") || (request.Email == "") || (request.Password == "") {
		return c.Status(400).JSON(fiber.Map{
			"message": "credentials are missing, Please try again",
		})
	}

	// perform register insertion
	id, err := h.repo.RegisterUser(c.Context(), request.Username, request.Email, request.Password)

	// check if the id is returned we want to know the error
	if err != nil {
		switch err {
		case ErrUsernameExists:
			return c.Status(400).JSON(fiber.Map{
				"message": "Username already exists",
			})
		case ErrEmailExists:
			return c.Status(400).JSON(fiber.Map{
				"message": "Email already exists",
			})
		default:
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to Register User, Please Try again",
			})
		}
	}

	// generate the jwt
	jwt, err := middleware.GenerateToken(id)

	// check if the token is generated
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to Generate Token, Please Try again",
		})
	}

	// send the cookies to client
	c.Cookie(
		&fiber.Cookie{
			Name:     "token",
			Value:    jwt,
			HTTPOnly: true,
			Path:     "/",
			MaxAge:   60 * 60 * 24,
			Secure:   false,
		})

	// return the message and the token
	return c.Status(200).JSON(fiber.Map{
		"message": "Welecome to Chat App" + " " + request.Username,
		"name":    request.Username,
	})
}
