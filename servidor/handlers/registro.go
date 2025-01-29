package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sabbaggh/marketplace-backend/models"
)

func Registro(c *fiber.Ctx) error {
	user := &models.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos xdd"})
	}
	if user.Username == "" || user.Password == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Nombre de usuario o contrasena no pueden quedar vacios"})
	}
	fmt.Println(user)
	return c.JSON(user)
}
