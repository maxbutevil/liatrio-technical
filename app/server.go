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
			"message":   "My name is Max van der Veen",
			"timestamp": time.Now().UnixMilli(),
		}

		return c.JSON(response)
	})

	log.Fatal(app.Listen(":3000"))
}
