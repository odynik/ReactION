package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		foo()
		uniDistroSamples(10)
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":8080"))
}
