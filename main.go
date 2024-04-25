package main

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type User struct {
	ID      int64
	Email   string
	Age     int
	Country string
}

var users = map[int64]User{}

type (
	CreateUserRequest struct {
		// BEGIN (write your solution here)
		ID      int64  `json:"id" validate:"required,min=0"`
		Email   string `json:"email" validate:"required,email"`
		Age     int    `json:"age" validate:"required,min=18,max=130"`
		Country string `json:"country" validate:"required,allowable_text"`
		// END
	}
)

func main() {
	webApp := fiber.New()
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	// BEGIN (write your solution here) (write your solution here)
	validate := validator.New()

	allowableWords := []string{"usa", "germany", "france"}

	vErr := validate.RegisterValidation("allowable_text", func(fl validator.FieldLevel) bool {
		text := fl.Field().String()
		for _, word := range allowableWords {
			if strings.Contains(strings.ToLower(text), word) {
				return true
			}
		}
		return false
	})

	if vErr != nil {
		logrus.Fatal("register validation ", vErr)
	}

	webApp.Post("/users", func(c *fiber.Ctx) error {
		req := &CreateUserRequest{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).SendString("body parser")
		}

		err := validate.Struct(req)

		if err != nil {
			return c.Status(422).SendString(err.Error())
		}

		users[req.ID] = User{
			ID: req.ID,
			Email: req.Email,
			Age: req.Age,
			Country: req.Country,
		}


		return c.SendStatus(200)
	})
	// END
	logrus.Fatal(webApp.Listen(":8080"))
}
