package main

import (
	"sort"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type (
	BinarySearchRequest struct {
		Numbers []int `json:"numbers"`
		Target  int   `json:"target"`
	}

	BinarySearchResponse struct {
		TargetIndex int    `json:"target_index"`
		Error       string `json:"error,omitempty"`
	}
)

const targetNotFound = -1

func main() {
	webApp := fiber.New()
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	// BEGIN (write your solution here)
	webApp.Post("/search", func(c *fiber.Ctx) error {
		requestStuct := BinarySearchRequest{}
		responseStruct := BinarySearchResponse{}

		if err := c.BodyParser(&requestStuct); err != nil {
			responseStruct.TargetIndex = targetNotFound
			responseStruct.Error = "Invalid JSON"

			return c.Status(400).JSON(responseStruct)
		}

		responseStruct.TargetIndex = sort.SearchInts(requestStuct.Numbers, requestStuct.Target)

		if responseStruct.TargetIndex == len(requestStuct.Numbers) {
			responseStruct.TargetIndex = targetNotFound
			responseStruct.Error = "Target was not found"

			return c.Status(404).JSON(responseStruct)
		}

		return c.Status(200).JSON(responseStruct)
	})
	// END

	logrus.Fatal(webApp.Listen(":8080"))
}
