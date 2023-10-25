package controllers

import "github.com/gofiber/fiber/v2"

func IndexRender(c *fiber.Ctx) error {

	// Mengirimkan halaman HTML yang dihasilkan ke browser
	return c.Render("index", fiber.Map{
		"Title":   "Selamat Datang di SIPARSIP",
		"Message": "Aplikasi Manajemen Arsip Surat!",
	})
}
