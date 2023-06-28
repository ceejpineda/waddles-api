package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/test-api", func(c *fiber.Ctx) error {
		return c.SendString("Test API")
	})

	app.Post("/test-api", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
		//return c.SendString("Test API")
	})

	app.Listen(":3000")
}
