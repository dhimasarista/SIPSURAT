package controllers

import (
	"sipsurat/models"

	"github.com/gofiber/fiber/v2"
)

func DashboardRender(c *fiber.Ctx) error {

	users := models.User{}
	data := users.FindAll()

	// Mengirimkan halaman HTML yang dihasilkan ke browser
	return c.Render("dashboard", fiber.Map{
		"Users": data,
	})
}
