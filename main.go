package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
)

func main() {
	engine := mustache.New("./views", ".mustache")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Rute untuk menampilkan halaman HTML
	app.Get("/", func(c *fiber.Ctx) error {

		// Mengirimkan halaman HTML yang dihasilkan ke browser
		return c.Render("index", fiber.Map{
			"Title":   "Selamat Datang di SIPARSIP",
			"Message": "Aplikasi Manajemen Arsip Surat!",
		})
	})

	// Menjalankan server pada port 3000
	log.Fatal(app.Listen(":3000"))
}
