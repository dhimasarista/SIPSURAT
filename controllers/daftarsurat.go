package controllers

import (
	"sipsurat/models"

	"github.com/gofiber/fiber/v2"
)

func DaftarSuratRender(c *fiber.Ctx) error {

	var path string = c.Path()

	users := models.User{}
	data := users.FindAll()

	// Mengirimkan halaman HTML yang dihasilkan ke browser
	return c.Render("daftarsurat", fiber.Map{
		"Users": data,
		"path":  path,
	})
}
