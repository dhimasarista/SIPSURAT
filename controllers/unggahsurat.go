package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func UnggahSuratRender(c *fiber.Ctx) error {

	var path string = c.Path()

	// Mengirimkan halaman HTML yang dihasilkan ke browser
	return c.Render("unggahsurat", fiber.Map{
		"path": path,
	})
}
