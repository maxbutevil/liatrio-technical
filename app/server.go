package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {

		response := fiber.Map{
			"name":      "My name is Max van der Veen",
			"timestamp": time.Now().Format(time.Stamp),
		}

		return c.JSON(response)
	})

	log.Fatal(app.Listen(":3000"))
}
