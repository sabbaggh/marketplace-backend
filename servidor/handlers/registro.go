package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sabbaggh/marketplace-backend/models"
	"github.com/sabbaggh/marketplace-backend/queries"
)

func Registro(c *fiber.Ctx) error {
	user := &models.User{}
	fmt.Println("HOLA")
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos xdd"})
	}
	if user.Username == "" || user.Password == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Nombre de usuario o contrasena no pueden quedar vacios"})
	}
	id, err := queries.Registrar(user.Username, user.Password, user.Fecha, user.Tecnologia, user.Ropa, user.Perfumes, user.Videojuegos, user.Coleccionables)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Error al registrar el usuario"})
	}
	fmt.Println(user)
	return c.Status(200).JSON(fiber.Map{"success": "Usuario registrado correctamente", "id": id})
}
