package main

import (
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type (
	CreateItemRequest struct {
		Name  string `json:"name"`
		Price uint   `json:"price"`
	}

	Item struct {
		Name  string `json:"name"`
		Price uint   `json:"price"`
	}
)

var (
	items []Item
)

func main() {
	viewsEngine := html.New("./templates", ".tmpl")
	webApp := fiber.New(fiber.Config{
		Views: viewsEngine,
	})
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
	// BEGIN (write your solution here)

	webApp.Post("/items", func(c *fiber.Ctx) error {
		req := CreateItemRequest{}
		if err := c.BodyParser(&req); err != nil {
			return c.SendStatus(400)
		}

		items = append(items, Item(req))

		return c.SendStatus(200)
	})

	webApp.Get("/items/view", func(c *fiber.Ctx) error {
		return c.Render("item", items)
	})

	// END

	logrus.Fatal(webApp.Listen(":8080"))
}
