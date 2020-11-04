package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type ErrorBodyParser struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	ErrorCatch string `json:"error_catch"`
}

func main() {
	app := fiber.New()

	app.Post("/users", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(user); err != nil {
			return c.Status(500).JSON(ErrorBodyParser{
				Success:    false,
				Message:    "Error Body Parser Data.",
				ErrorCatch: err.Error(),
			})
		}

		return c.Status(200).JSON(user)
	})

	log.Fatal(app.Listen(":8080"))
}
