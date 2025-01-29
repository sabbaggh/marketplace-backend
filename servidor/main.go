package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sabbaggh/marketplace-backend/handlers"
)

func main() {
	fmt.Println("Hola")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "holaaa"})
	})
	app.Post("/register", handlers.Registro)

	log.Fatal(app.Listen(":3001"))
}
