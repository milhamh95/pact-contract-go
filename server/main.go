package main

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	app := fiber.New()

	posts := []Post{
		{
			UserID: 1,
			ID:     3,
			Title:  "ea molestias quasi exercitationem repellat qui ipsa sit aut",
			Body:   "ullam et saepe reiciendis voluptatem ad",
		},
		{
			UserID: 1,
			ID:     4,
			Title:  "eum et est occaecati",
			Body:   "ullam et saepe reiciendis voluptatem adipisci\\nsit amet",
		},
	}

	app.Get("/posts", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(posts)
	})

	app.Listen(":3000")
}
