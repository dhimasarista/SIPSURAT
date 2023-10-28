package routes

import (
	"sipsurat/controllers"

	"github.com/gofiber/fiber/v2"
)

func UnggahSuratRoutes(app *fiber.App) {
	app.Get("/unggahsurat", controllers.UnggahSuratRender)
}
