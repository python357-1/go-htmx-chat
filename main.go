package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./wwwroot")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{}, "_layout")
	})

	app.Listen(":8000")
}
