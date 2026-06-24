package auth

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	svc *Service
}

func NewHandler(db *pgxpool.Pool) *Handler {
	return &Handler{
		svc: NewService(NewRepository(db)),
	}
}

func (h *Handler) Health(c fiber.Ctx) error {
	version, err := h.svc.repo.CheckVersion(c.Context())
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

func (h *Handler) Login(c fiber.Ctx) error {
	request := new(LoginRequest)
	if err := c.Bind().JSON(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request, Please try again",
		})
	}
	if request.Email == "" || request.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "details are missing, Please try again",
		})
	}

	token, user, err := h.svc.Login(c.Context(), request.Email, request.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid Email and Password, Please try again",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		HTTPOnly: true,
		Path:     "/",
		MaxAge:   60 * 60 * 24,
		Secure:   false,
	})

	return c.Status(200).JSON(fiber.Map{
		"message": "Welecome to Chat App" + " " + user.Username,
		"name":    user.Username,
		"id":      user.ID,
	})
}

func (h *Handler) Register(c fiber.Ctx) error {
	request := new(RegisterRequest)
	if err := c.Bind().JSON(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request, Please try again",
		})
	}
	if request.Username == "" || request.Email == "" || request.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "credentials are missing, Please try again",
		})
	}

	token, id, err := h.svc.Register(c.Context(), *request)
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
			log.Printf("Failed to Register User: %v", err)
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to Register User, Please Try again",
			})
		}
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		HTTPOnly: true,
		Path:     "/",
		MaxAge:   60 * 60 * 24,
		Secure:   false,
	})

	return c.Status(200).JSON(fiber.Map{
		"message":  "Welcome to Chat App" + " " + request.Username,
		"id":       id,
		"username": request.Username,
	})
}
