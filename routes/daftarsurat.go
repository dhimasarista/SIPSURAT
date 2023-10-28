package routes

import (
	"sipsurat/controllers"

	"github.com/gofiber/fiber/v2"
)

func DaftarSuratRoutes(app *fiber.App) {
	app.Get("/daftarsurat", controllers.DaftarSuratRender)
}
