package main

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type (
	CreateLinkRequest struct {
		External string `json:"external"`
		Internal string `json:"internal"`
	}

	GetLinkResponse struct {
		Internal string `json:"internal"`
	}
)

var links = make(map[string]string)

func main() {
	webApp := fiber.New()
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
	// BEGIN (write your solution here)
	webApp.Post("/links", func(c *fiber.Ctx) error {
		req := CreateLinkRequest{}

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
		}

		links[req.External] = req.Internal

		return c.SendStatus(200)
	})

	webApp.Get("/links/:external", func(c *fiber.Ctx) error {
		req := c.Params("external", "")
		request, err := url.QueryUnescape(req)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Link not found")
		}

		internal, ok := links[request]

		if !ok {
			return c.Status(fiber.StatusNotFound).SendString("Link not found")
		}

		res := GetLinkResponse{Internal: internal}
		return c.Status(200).JSON(res)
	})
	// END

	logrus.Fatal(webApp.Listen(":8080"))
}
