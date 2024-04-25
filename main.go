package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type (
	GetTaskResponse struct {
		ID       int64  `json:"id"`
		Desc     string `json:"description"`
		Deadline int64  `json:"deadline"`
	}

	CreateTaskRequest struct {
		Desc     string `json:"description"`
		Deadline int64  `json:"deadline"`
	}

	CreateTaskResponse struct {
		ID int64 `json:"id"`
	}

	UpdateTaskRequest struct {
		Desc     string `json:"description"`
		Deadline int64  `json:"deadline"`
	}

	Task struct {
		ID       int64
		Desc     string
		Deadline int64
	}
)

var (
	taskIDCounter int64 = 1
	tasks               = make(map[int64]Task)
)

func main() {
	webApp := fiber.New()
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	// BEGIN (write your solution here) (write your solution here)
	webApp.Post("/tasks", func(c *fiber.Ctx) error {
		req := CreateTaskRequest{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).SendString("Invalid JSON")
		}

		CreatedTask := Task{
			ID:       taskIDCounter,
			Desc:     req.Desc,
			Deadline: req.Deadline,
		}
		taskIDCounter += 1
		tasks[CreatedTask.ID] = CreatedTask

		res := CreateTaskResponse{ID: CreatedTask.ID}

		return c.Status(200).JSON(res)
	})

	webApp.Patch("/tasks/:id", func(c *fiber.Ctx) error {
		req := c.Params("id", "")
		reqBody := GetTaskResponse{}
		if err := c.BodyParser(&reqBody); err != nil {
			return c.SendStatus(404)
		}

		reqID, err := strconv.Atoi(req)

		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}

		PatchTask, ok := tasks[int64(reqID)]

		if !ok {
			return c.SendStatus(404)
		}

		if reqBody.Desc != "" {
			PatchTask.Desc = reqBody.Desc
		}
		if reqBody.Deadline != 0 {
			PatchTask.Deadline = reqBody.Deadline
		}

		tasks[int64(reqID)] = PatchTask


		return c.SendStatus(200)
	})

	webApp.Get("/tasks/:id", func(c *fiber.Ctx) error {
		req := c.Params("id", "")
		reqID, err := strconv.Atoi(req)

		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}

		GetTask, ok := tasks[int64(reqID)]

		if !ok {
			return c.SendStatus(404)
		}

		return c.Status(200).JSON(GetTaskResponse(GetTask))
	})

	webApp.Delete("/tasks/:id", func(c *fiber.Ctx) error {
		req := c.Params("id", "")
		reqID, err := strconv.Atoi(req)

		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}

		_, ok := tasks[int64(reqID)]

		if !ok {
			return c.SendStatus(404)
		}

		delete(tasks, int64(reqID))


		return c.SendStatus(200)
	})

	// END

	logrus.Fatal(webApp.Listen(":8080"))
}
