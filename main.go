package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var postLikes = map[string]int64{}

func main() {
	webApp := fiber.New(fiber.Config{Immutable: true})
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Go to /likes/12345")
	})
	// BEGIN (write your solution here)
	webApp.Get("/likes/:id", func(c *fiber.Ctx) error {
		id := c.Params("id", "")
		postID, ok := postLikes[id]

		if !ok {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.SendString(strconv.FormatInt(postID, 10))
	})

	webApp.Post("/likes/:id", func(c *fiber.Ctx) error {
		id := c.Params("id", "")
		postID, ok := postLikes[id]

		postID++
		postLikes[id] += 1

		if !ok {
			return c.Status(fiber.StatusCreated).SendString(strconv.FormatInt(postID, 10))
		}

		return c.Status(fiber.StatusOK).SendString(strconv.FormatInt(postID, 10))
	})

	// END

	logrus.Fatal(webApp.Listen(":8080"))
}
